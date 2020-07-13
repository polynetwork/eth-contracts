pragma solidity ^0.5.0;

import "./../../../libs/GSN/Context.sol";
import "./../../../libs/common/ZeroCopySource.sol";
import "./../../../libs/common/ZeroCopySink.sol";
import "./../../../libs/utils/Utils.sol";
import "./../../../libs/math/SafeMath.sol";
import "./../CrossChainManager/IEthCrossChainManager.sol";

interface ERC20Interface {
    function transfer(address _to, uint256 _value) external returns (bool);
    function transferFrom(address _from, address _to, uint _value) external returns (bool success);
}


contract LockProxy is Context {
    using SafeMath for uint;
    
    struct TxArgs {
        bytes toAssetHash;
        bytes toAddress;
        uint256 amount;
    }
    address public operator;
    address public managerContract;
    mapping(uint64 => bytes) public proxyHashMap;
    mapping(address => mapping(uint64 => bytes)) public assetHashMap;
    mapping(address => mapping(uint64 => uint256)) public crossedLimit;    // asset in Target Chain Circulating limit
    mapping(address => mapping(uint64 => uint256)) public crossedAmount;    // asset in Target Chain circulating supply
    
    event SetManagerEvent(address manager);
    event BindProxyEvent(uint64 toChainId, bytes targetProxyHash);
    event BindAssetEvent(address fromAssetHash, uint64 toChainId, bytes targetProxyHash, uint256 assetLimit);
    event UnlockEvent(bytes fromContractAddr, uint64 fromChainId, bytes toAddress, uint256 amount);
    event LockEvent(address thisContract, uint64 chainId, bytes toContract, bytes txArgs);
    
    constructor () public {
        operator = _msgSender();
    }
    modifier onlyOperator() {
        require(_msgSender() == operator) ;
        _;
    }
    modifier onlyManagerContract() {
        require(_msgSender() == managerContract) ;
        _;
    }
    
    function setManager(address ethCrossChainContractAddr) onlyOperator public {
        managerContract = ethCrossChainContractAddr;
        emit SetManagerEvent(ethCrossChainContractAddr);
    }
    
    function bindProxyHash(uint64 toChainId, bytes memory targetProxyHash) onlyOperator public returns (bool) {
        proxyHashMap[toChainId] = targetProxyHash;
        emit BindProxyEvent(toChainId, targetProxyHash);
        return true;
    }
    
    function bindAssetHash(address fromAssetHash, uint64 toChainId, bytes memory toAssetHash, uint256 assetLimit, bool isTargetChainAsset) onlyOperator public returns (bool) {
        assetHashMap[fromAssetHash][toChainId] = toAssetHash;
        if (isTargetChainAsset) {
            uint256 currentLimit = crossedLimit[fromAssetHash][toChainId];
            require(assetLimit >= currentLimit, "assetLimit can only be updated increasingly!");
            uint256 increment = assetLimit.sub(currentLimit);
            crossedAmount[fromAssetHash][toChainId] = crossedAmount[fromAssetHash][toChainId].add(increment);
        }
        crossedLimit[fromAssetHash][toChainId] = assetLimit;
        emit BindAssetEvent(fromAssetHash, toChainId, toAssetHash, assetLimit);
        return true;
    }
    
    /* @notice                  This function is meant to be invoked by the user,
    *                           a certin amount teokens will be locked in the proxy contract the invoker/msg.sender immediately.
    *                           Then the same amount of tokens will be unloked from target chain proxy contract at the target chain with chainId later.
    *  @param fromAssetHash   The asset hash in current chain
    *  @param toChainId         The target chain id
    *                           
    *  @param toAddress         The address in bytes format to receive same amount of tokens in target chain 
    *  @param amount            The amount of tokens to be crossed from ethereum to the chain with chainId
    */
    function lock(address fromAssetHash, uint64 toChainId, bytes memory toAddress, uint256 amount) public payable returns (bool) {
        require(amount >= 0, "amount is less than zero!");
        
        crossedAmount[fromAssetHash][toChainId] = crossedAmount[fromAssetHash][toChainId].add(amount);
        require(crossedAmount[fromAssetHash][toChainId] <= crossedLimit[fromAssetHash][toChainId], "asset in target chain will exceed limit control!");
        
        require(_transferToContract(fromAssetHash, amount), "transfer asset from fromAddress to lock_proxy contract  failed!");
        
        bytes memory toAssetHash = assetHashMap[fromAssetHash][toChainId];
        require(toAssetHash.length > 0, "empty illegal toAssetHash");
        TxArgs memory txArgs = TxArgs({
            toAssetHash: toAssetHash,
            toAddress: toAddress,
            amount: amount
        });
        bytes memory txData = _serializeTxArgs(txArgs);
        
        IEthCrossChainManager eccm = IEthCrossChainManager(managerContract);
        
        require(eccm.crossChain(toChainId, proxyHashMap[toChainId], "unlock", txData), "EthCrossChainManager crossChain executed error!");

        emit LockEvent(address(this), toChainId, proxyHashMap[toChainId], txData);
        
        return true;

    }
    
    // /* @notice                  This function is meant to be invoked by the ETH crosschain management contract,
    // *                           then mint a certin amount of tokens to the designated address since a certain amount 
    // *                           was burnt from the source chain invoker.
    // *  @param argsBs            The argument bytes recevied by the ethereum lock proxy contract, need to be deserialized.
    // *                           based on the way of serialization in the source chain proxy contract.
    // *  @param fromContractAddr  The source chain contract address
    // *  @param fromChainId       The source chain id
    // */
    function unlock(bytes memory argsBs, bytes memory fromContractAddr, uint64 fromChainId) onlyManagerContract public returns (bool) {
        TxArgs memory args = _deserializTxArgs(argsBs);

        require(Utils.equalStorage(proxyHashMap[fromChainId], fromContractAddr), "From Proxy contract address error!");
        
        address targetAssetAddress = Utils.bytesToAddress(args.toAssetHash);
        address toAddress = Utils.bytesToAddress(args.toAddress);
        
        crossedAmount[targetAssetAddress][fromChainId] = crossedAmount[targetAssetAddress][fromChainId].sub(args.amount);
        
        require(_transferFromContract(targetAssetAddress, toAddress, args.amount), "transfer asset from lock_proxy contract to toAddress failed!");
        
        emit UnlockEvent(fromContractAddr, fromChainId, args.toAddress, args.amount);
        return true;
    }
    
    function _transferToContract(address fromAssetHash, uint256 amount) internal returns (bool) {
        if (fromAssetHash == address(0) && msg.value > 0) {
            // fromAssetHash === address(0) denotes user choose to lock ether
            // passively check if the received msg.value equals amount
            require(msg.value == amount, "transferred ether is not equal to amount!");
        } else {
            // actively transfer amount of asset from msg.sender to lock_proxy contract 
            require(_transferERC20ToContract(fromAssetHash, _msgSender(), address(this), amount), "transfer erc20 asset to lock_proxy contract failed!");
        }
        return true;
    }
    function _transferFromContract(address toAssetHash, address toAddress, uint256 amount) internal returns (bool) {
        if (toAssetHash == address(0x0000000000000000000000000000000000000000)) {
            // toAssetHash === address(0) denotes contract needs to unlock ether to toAddress
            // convert toAddress from 'address' type to 'address payable' type, then actively transfer ether
            address(uint160(toAddress)).transfer(amount);
        } else {
            // actively transfer amount of asset from msg.sender to lock_proxy contract 
            require(_transferERC20FromContract(toAssetHash, toAddress, amount), "transfer erc20 asset to lock_proxy contract failed!");
        }
        return true;
    }
    
    
    function _transferERC20ToContract(address fromAssetHash, address fromAddress, address toAddress, uint256 amount) internal returns (bool) {
         ERC20Interface erc20Token = ERC20Interface(fromAssetHash);
         require(erc20Token.transferFrom(fromAddress, toAddress, amount), "trasnfer ERC20 Token failed!");
         return true;
    }
    function _transferERC20FromContract(address toAssetHash, address toAddress, uint256 amount) internal returns (bool) {
         ERC20Interface erc20Token = ERC20Interface(toAssetHash);
         require(erc20Token.transfer(toAddress, amount), "trasnfer ERC20 Token failed!");
         return true;
    }
    
    function _serializeTxArgs(TxArgs memory args) internal pure returns (bytes memory) {
        bytes memory buff;
        buff = abi.encodePacked(
            ZeroCopySink.WriteVarBytes(args.toAssetHash),
            ZeroCopySink.WriteVarBytes(args.toAddress),
            ZeroCopySink.WriteUint255(args.amount)
            );
        return buff;
    }

    function _deserializTxArgs(bytes memory valueBs) internal pure returns (TxArgs memory) {
        TxArgs memory args;
        uint256 off = 0;
        (args.toAssetHash, off) = ZeroCopySource.NextVarBytes(valueBs, off);
        (args.toAddress, off) = ZeroCopySource.NextVarBytes(valueBs, off);
        (args.amount, off) = ZeroCopySource.NextUint255(valueBs, off);
        return args;
    }
    
    event DebugBytes(bytes data);
    function test_serializeTxArgs(bytes memory toAssetHash, bytes memory toAddress, uint256 amount) public returns (bool) {
        TxArgs memory txArgs = TxArgs({
            toAssetHash: toAssetHash,
            toAddress: toAddress,
            amount: amount
        });
        bytes memory txData = _serializeTxArgs(txArgs);
        emit DebugBytes(txData);
        return true;
    }
    event DebugUint256(uint256 amount);
    function test_deserializeTxArgs(bytes memory txData) public returns (bool) {
        TxArgs memory args = _deserializTxArgs(txData);
        emit DebugBytes(args.toAssetHash);
        emit DebugBytes(args.toAddress);
        emit DebugUint256(args.amount);
    }
}