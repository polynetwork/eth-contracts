pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./../../libs/ownership/Ownable.sol";
import "./../../libs/common/ZeroCopySource.sol";
import "./../../libs/common/ZeroCopySink.sol";
import "./../../libs/utils/Utils.sol";
import "./../../libs/token/ERC20/SafeERC20.sol";
import "./../../libs/lifecycle/Pausable.sol";
import "./../cross_chain_manager/interface/IEthCrossChainManager.sol";
import "./../cross_chain_manager/interface/IEthCrossChainManagerProxy.sol";

contract LockProxyWithLP is Ownable, Pausable {
    using SafeMath for uint;
    using SafeERC20 for IERC20;

    struct TxArgs {
        bytes toAssetHash;
        bytes toAddress;
        uint256 amount;
    }
    address public managerProxyContract;

    mapping(uint64 => bytes) public proxyHashMap;
    mapping(address => mapping(uint64 => bytes)) public assetHashMap;
    mapping(address => address) public assetLPMap;

    event SetManagerProxyEvent(address manager);
    event BindProxyEvent(uint64 toChainId, bytes targetProxyHash);
    event BindAssetEvent(address fromAssetHash, uint64 toChainId, bytes targetProxyHash, uint initialAmount);
    event UnlockEvent(address toAssetHash, address toAddress, uint256 amount);
    event LockEvent(address fromAssetHash, address fromAddress, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount);

    event depositEvent(address toAddress, address fromAssetHash, address fromLPHash, uint256 amount);
    event withdrawEvent(address toAddress, address fromAssetHash, address fromLPHash, uint256 amount);
    event BindLPToAssetEvent(address originAssetAddress, address LPTokenAddress);

    modifier onlyManagerContract() {
        IEthCrossChainManagerProxy ieccmp = IEthCrossChainManagerProxy(managerProxyContract);
        require(_msgSender() == ieccmp.getEthCrossChainManager(), "msgSender is not EthCrossChainManagerContract");
        _;
    }

    function pause() onlyOwner whenNotPaused public returns (bool) {
        _pause();
        return true;
    }
    function unpause() onlyOwner whenPaused public returns (bool) {
        _unpause();
        return true;
    }
    
    function setManagerProxy(address ethCCMProxyAddr) onlyOwner public {
        managerProxyContract = ethCCMProxyAddr;
        emit SetManagerProxyEvent(managerProxyContract);
    }
    
    function bindProxyHash(uint64 toChainId, bytes memory targetProxyHash) onlyOwner public returns (bool) {
        proxyHashMap[toChainId] = targetProxyHash;
        emit BindProxyEvent(toChainId, targetProxyHash);
        return true;
    }

    function bindAssetHash(address fromAssetHash, uint64 toChainId, bytes memory toAssetHash) onlyOwner public returns (bool) {
        assetHashMap[fromAssetHash][toChainId] = toAssetHash;
        emit BindAssetEvent(fromAssetHash, toChainId, toAssetHash, getBalanceFor(fromAssetHash));
        return true;
    }

    function bindLPToAsset(address originAssetAddress, address LPTokenAddress) onlyOwner public returns (bool) {
        assetLPMap[originAssetAddress] = LPTokenAddress;
        emit BindLPToAssetEvent(originAssetAddress, LPTokenAddress);
        return true;
    }

    function bindLPAndAsset(address fromAssetHash, address fromLPHash, uint64 toChainId, bytes memory toAssetHash, bytes memory toLPHash) onlyOwner public returns (bool) {
        assetHashMap[fromAssetHash][toChainId] = toAssetHash;
        assetHashMap[fromLPHash][toChainId] = toLPHash;
        assetLPMap[fromAssetHash] = fromLPHash;
        emit BindAssetEvent(fromAssetHash, toChainId, toAssetHash, getBalanceFor(fromAssetHash));
        emit BindAssetEvent(fromLPHash, toChainId, toLPHash, getBalanceFor(fromLPHash));
        emit BindLPToAssetEvent(fromAssetHash, fromLPHash);
        return true;
    }
    
    function bindProxyHashBatch(uint64[] memory toChainId, bytes[] memory targetProxyHash) onlyOwner public returns (bool) {
        require(toChainId.length == targetProxyHash.length, "bindProxyHashBatch: args length diff");
        for (uint i = 0; i < toChainId.length; i++) {
            proxyHashMap[toChainId[i]] = targetProxyHash[i];
            emit BindProxyEvent(toChainId[i], targetProxyHash[i]);
        }
        return true;
    }

    function bindAssetHashBatch(address[] memory fromAssetHash, uint64[] memory toChainId, bytes[] memory toAssetHash) onlyOwner public returns (bool) {
        require(toChainId.length == fromAssetHash.length, "bindAssetHashBatch: args length diff");
        require(toChainId.length == toAssetHash.length, "bindAssetHashBatch: args length diff");
        for (uint i = 0; i < toChainId.length; i++) {
            assetHashMap[fromAssetHash[i]][toChainId[i]] = toAssetHash[i];
            emit BindAssetEvent(fromAssetHash[i], toChainId[i], toAssetHash[i], getBalanceFor(fromAssetHash[i]));
        }
        return true;
    }

    function bindLPToAssetBatch(address[] memory originAssetAddress, address[] memory LPTokenAddress) onlyOwner public returns (bool) {
        require(originAssetAddress.length == LPTokenAddress.length, "bindLPToAssetBatch: args length diff");
        for (uint i = 0; i < originAssetAddress.length; i++) {
            assetLPMap[originAssetAddress[i]] = LPTokenAddress[i];
            emit BindLPToAssetEvent(originAssetAddress[i], LPTokenAddress[i]);
        }
        return true;
    }
 
    function bindLPAndAssetBatch(address[] memory fromAssetHash, address[] memory fromLPHash, uint64[] memory toChainId, bytes[] memory toAssetHash, bytes[] memory toLPHash) onlyOwner public returns (bool) {
        require(fromAssetHash.length == fromLPHash.length, "bindLPAndAssetBatch: args length diff");
        require(toAssetHash.length == toLPHash.length, "bindLPAndAssetBatch: args length diff");
        for(uint256 i = 0; i < fromLPHash.length; i++) {
            assetHashMap[fromAssetHash[i]][toChainId[i]] = toAssetHash[i];
            assetHashMap[fromLPHash[i]][toChainId[i]] = toLPHash[i];
            assetLPMap[fromAssetHash[i]] = fromLPHash[i];
            emit BindAssetEvent(fromAssetHash[i], toChainId[i], toAssetHash[i], getBalanceFor(fromAssetHash[i]));
            emit BindAssetEvent(fromLPHash[i], toChainId[i], toLPHash[i], getBalanceFor(fromLPHash[i]));
            emit BindLPToAssetEvent(fromAssetHash[i], fromLPHash[i]);
        }
        return true;
    } 

    /* @notice                  This function is meant to be invoked by the user,
    *                           a certain amount tokens will be locked in the proxy contract the invoker/msg.sender immediately.
    *                           Then the same amount of tokens will be unloked from target chain proxy contract at the target chain with chainId later.
    *  @param fromAssetHash     The asset address in current chain, uniformly named as `fromAssetHash`
    *  @param toChainId         The target chain id
    *                           
    *  @param toAddress         The address in bytes format to receive same amount of tokens in target chain 
    *  @param amount            The amount of tokens to be crossed from ethereum to the chain with chainId
    */
    function lock(address fromAssetHash, uint64 toChainId, bytes memory toAddress, uint256 amount) public payable returns (bool) {
        require(amount != 0, "amount cannot be zero!");
        require(_transferToContract(fromAssetHash, amount), "transfer asset from fromAddress to lock_proxy contract failed!");
        bytes memory toAssetHash = assetHashMap[fromAssetHash][toChainId];
        require(toAssetHash.length != 0, "empty illegal toAssetHash");

        TxArgs memory txArgs = TxArgs({
            toAssetHash: toAssetHash,
            toAddress: toAddress,
            amount: amount
        });
        bytes memory txData = _serializeTxArgs(txArgs);
        
        IEthCrossChainManagerProxy eccmp = IEthCrossChainManagerProxy(managerProxyContract);
        address eccmAddr = eccmp.getEthCrossChainManager();
        IEthCrossChainManager eccm = IEthCrossChainManager(eccmAddr);
        
        bytes memory toProxyHash = proxyHashMap[toChainId];
        require(toProxyHash.length != 0, "empty illegal toProxyHash");
        require(eccm.crossChain(toChainId, toProxyHash, "unlock", txData), "EthCrossChainManager crossChain executed error!");

        emit LockEvent(fromAssetHash, _msgSender(), toChainId, toAssetHash, toAddress, amount);
        
        return true;
    }
    
    /* @notice                  This function is meant to be invoked by the ETH crosschain management contract,
    *                           then mint a certin amount of tokens to the designated address since a certain amount 
    *                           was burnt from the source chain invoker.
    *  @param argsBs            The argument bytes recevied by the ethereum lock proxy contract, need to be deserialized.
    *                           based on the way of serialization in the source chain proxy contract.
    *  @param fromContractAddr  The source chain contract address
    *  @param fromChainId       The source chain id
    */
    function unlock(bytes memory argsBs, bytes memory fromContractAddr, uint64 fromChainId) onlyManagerContract public returns (bool) {
        TxArgs memory args = _deserializeTxArgs(argsBs);

        require(fromContractAddr.length != 0, "from proxy contract address cannot be empty");
        require(Utils.equalStorage(proxyHashMap[fromChainId], fromContractAddr), "From Proxy contract address error!");
        
        require(args.toAssetHash.length != 0, "toAssetHash cannot be empty");
        address toAssetHash = Utils.bytesToAddress(args.toAssetHash);

        require(args.toAddress.length != 0, "toAddress cannot be empty");
        address toAddress = Utils.bytesToAddress(args.toAddress);
        
        require(_transferFromContract(toAssetHash, toAddress, args.amount), "transfer asset from lock_proxy contract to toAddress failed!");
        
        emit UnlockEvent(toAssetHash, toAddress, args.amount);
        return true;
    }


    function deposit(address originAssetAddress, uint amount) whenNotPaused payable public returns (bool) {
        require(amount != 0, "amount cannot be zero!");

        require(_transferToContract(originAssetAddress, amount), "transfer asset from fromAddress to lock_proxy contract failed!");

        address LPTokenAddress = assetLPMap[originAssetAddress];
        require(LPTokenAddress != address(0), "do not support deposite this token");
        require(_transferFromContract(LPTokenAddress, msg.sender, amount), "transfer proof of liquidity from lock_proxy contract to fromAddress failed!");
        
        emit depositEvent(msg.sender, originAssetAddress, LPTokenAddress, amount);
        return true;
    }

    function withdraw(address targetTokenAddress, uint amount) whenNotPaused public returns (bool) {
        require(amount != 0, "amount cannot be zero!");

        address LPTokenAddress = assetLPMap[targetTokenAddress];
        require(LPTokenAddress != address(0), "do not support withdraw this token");
        require(_transferToContract(LPTokenAddress, amount), "transfer proof of liquidity from fromAddress to lock_proxy contract failed!");

        require(_transferFromContract(targetTokenAddress, msg.sender, amount), "transfer asset from lock_proxy contract to fromAddress failed!");
        
        emit withdrawEvent(msg.sender, targetTokenAddress, LPTokenAddress, amount);
        return true;
    }

    function getBalanceFor(address fromAssetHash) public view returns (uint256) {
        if (fromAssetHash == address(0)) {
            // return address(this).balance; // this expression would result in error: Failed to decode output: Error: insufficient data for uint256 type
            address selfAddr = address(this);
            return selfAddr.balance;
        } else {
            IERC20 erc20Token = IERC20(fromAssetHash);
            return erc20Token.balanceOf(address(this));
        }
    }

    function _transferToContract(address fromAssetHash, uint256 amount) internal returns (bool) {
        if (fromAssetHash == address(0)) {
            // fromAssetHash === address(0) denotes user choose to lock ether
            // passively check if the received msg.value equals amount
            require(msg.value != 0, "transferred ether cannot be zero!");
            require(msg.value == amount, "transferred ether is not equal to amount!");
        } else {
            // make sure lockproxy contract will decline any received ether
            require(msg.value == 0, "there should be no ether transfer!");
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
         IERC20 erc20Token = IERC20(fromAssetHash);
        //  require(erc20Token.transferFrom(fromAddress, toAddress, amount), "trasnfer ERC20 Token failed!");
         erc20Token.safeTransferFrom(fromAddress, toAddress, amount);
         return true;
    }
    function _transferERC20FromContract(address toAssetHash, address toAddress, uint256 amount) internal returns (bool) {
         IERC20 erc20Token = IERC20(toAssetHash);
        //  require(erc20Token.transfer(toAddress, amount), "transfer ERC20 Token failed!");
         erc20Token.safeTransfer(toAddress, amount);
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

    function _deserializeTxArgs(bytes memory valueBs) internal pure returns (TxArgs memory) {
        TxArgs memory args;
        uint256 off = 0;
        (args.toAssetHash, off) = ZeroCopySource.NextVarBytes(valueBs, off);
        (args.toAddress, off) = ZeroCopySource.NextVarBytes(valueBs, off);
        (args.amount, off) = ZeroCopySource.NextUint255(valueBs, off);
        return args;
    }
}