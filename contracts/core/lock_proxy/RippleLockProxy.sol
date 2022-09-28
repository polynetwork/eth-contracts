pragma solidity ^0.5.0;

import "./../../libs/ownership/Ownable.sol";
import "./../../libs/common/ZeroCopySource.sol";
import "./../../libs/common/ZeroCopySink.sol";
import "./../../libs/utils/Utils.sol";
import "./../../libs/token/ERC20/SafeERC20.sol";
import "./../../libs/token/ERC20/ERC20.sol";
import "./../../libs/token/ERC20/ERC20Detailed.sol";
import "./../cross_chain_manager/interface/IEthCrossChainManager.sol";
import "./../cross_chain_manager/interface/IEthCrossChainManagerProxy.sol";

contract bridgeAsset is Context, ERC20, ERC20Detailed {

    address public bridge;

    constructor (string memory name, string memory symbol, uint8 decimals, address bridge_) 
    public ERC20Detailed(name, symbol, decimals) {
        bridge = bridge_;
    }

    modifier onlyBridge() {
        require(_msgSender() == bridge, "msgSender is not Bridge!");
        _;
    }

    function mint(address to, uint256 amount) public onlyBridge {
        _mint(to, amount);
    }
    
    function burnFrom(address account, uint256 amount) public onlyBridge {
        _burnFrom(account, amount);
    }
}

contract RippleLockProxy is Ownable {
    using SafeMath for uint;
    using SafeERC20 for IERC20;

    struct TxArgs {
        bytes toAddress;
        uint256 amount;
    }
    bridgeAsset public token;
    address public managerProxyContract;
    mapping(uint64 => bytes) public proxyHashMap;

    uint public rippleMinAmount = 30000000;
    uint64 public rippleChainId = 39;
    uint public rippleAddressLength = 20;

    event SetManagerProxyEvent(address manager);
    event BindProxyEvent(uint64 toChainId, bytes targetProxyHash);
    event UnlockEvent(address toAssetHash, address toAddress, uint256 amount);
    event LockEvent(address fromAssetHash, address fromAddress, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount);
    
    constructor(string memory name, string memory symbol, uint8 decimals) public {
        token = new bridgeAsset(name, symbol, decimals, address(this));
    }

    modifier onlyManagerContract() {
        IEthCrossChainManagerProxy ieccmp = IEthCrossChainManagerProxy(managerProxyContract);
        require(_msgSender() == ieccmp.getEthCrossChainManager(), "msgSender is not EthCrossChainManagerContract");
        _;
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

    function rippleSetup(uint64 _rippleChainId, uint _rippleMinAmount, uint _rippleAddressLength) external onlyOwner {
        rippleChainId = _rippleChainId;
        rippleAddressLength = _rippleAddressLength;
        rippleMinAmount = _rippleMinAmount;
    }

    function rippleSetup(uint64 _rippleChainId) external onlyOwner {
        rippleChainId = _rippleChainId;
    }
    
    /* @notice                  This function is meant to be invoked by the user,
    *                           a certin amount teokens will be locked in the proxy contract the invoker/msg.sender immediately.
    *                           Then the same amount of tokens will be unloked from target chain proxy contract at the target chain with chainId later.
    *  @param fromAssetHash     The asset address in current chain, uniformly named as `fromAssetHash`
    *  @param toChainId         The target chain id
    *                           
    *  @param toAddress         The address in bytes format to receive same amount of tokens in target chain 
    *  @param amount            The amount of tokens to be crossed from ethereum to the chain with chainId
    */
    function lock(uint64 toChainId, bytes memory toAddress, uint256 amount) public payable returns (bool) {
        _rippleCheck(toChainId, toAddress, amount);
        require(amount != 0, "amount cannot be zero!");
        
        bridgeAsset(token).burnFrom(_msgSender(), amount);

        TxArgs memory txArgs = TxArgs({
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

        emit LockEvent(address(token), _msgSender(), toChainId, toProxyHash, toAddress, amount);
        
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
        TxArgs memory args = _deserializeTxArgs(argsBs);

        require(fromContractAddr.length != 0, "from proxy contract address cannot be empty");
        require(Utils.equalStorage(proxyHashMap[fromChainId], fromContractAddr), "From Proxy contract address error!");

        require(args.toAddress.length != 0, "toAddress cannot be empty");
        address toAddress = Utils.bytesToAddress(args.toAddress);
        
        bridgeAsset(token).mint(toAddress, args.amount);
        
        emit UnlockEvent(address(token), toAddress, args.amount);
        return true;
    }
    
    function _rippleCheck(uint64 toChainId, bytes memory toAddress, uint amount) internal view {
        if (toChainId == rippleChainId) {
            require(toAddress.length == rippleAddressLength, "invalid ripple address");
            require(amount >= rippleMinAmount, "amount less than the minimum");
        }
    }
    
    function _serializeTxArgs(TxArgs memory args) internal pure returns (bytes memory) {
        bytes memory buff;
        buff = abi.encodePacked(
            ZeroCopySink.WriteVarBytes(args.toAddress),
            ZeroCopySink.WriteUint255(args.amount)
            );
        return buff;
    }

    function _deserializeTxArgs(bytes memory valueBs) internal pure returns (TxArgs memory) {
        TxArgs memory args;
        uint256 off = 0;
        (args.toAddress, off) = ZeroCopySource.NextVarBytes(valueBs, off);
        (args.amount, off) = ZeroCopySource.NextUint255(valueBs, off);
        return args;
    }
}