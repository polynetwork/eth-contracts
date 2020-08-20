pragma solidity ^0.5.0;

import "./../../libs/GSN/Context.sol";
import "./../../libs/common/ZeroCopySource.sol";
import "./../../libs/common/ZeroCopySink.sol";
import "./../../libs/utils/Utils.sol";
import "./../../libs/math/SafeMath.sol";
import "./../cross_chain_manager/interface/IEthCrossChainManager.sol";
import "./../cross_chain_manager/interface/IEthCrossChainManagerProxy.sol";

interface ERC20Interface {
    function transfer(address _to, uint256 _value) external returns (bool);
    function transferFrom(address _from, address _to, uint _value) external returns (bool success);
    function balanceOf(address account) external view returns (uint256);
}

contract LockProxyPip1 is Context {
    using SafeMath for uint;

    struct RegisterAssetTxArgs {
        bytes assetHash;
        bytes nativeAssetHash;
    }

    struct TxArgs {
        bytes fromAssetHash;
        bytes toAssetHash;
        bytes toAddress;
        uint256 amount;
        uint256 feeAmount;
        bytes feeAddress;
    }

    address public managerProxyContract;
    mapping(bytes32 => bool) public registry;
    mapping(bytes32 => uint256) public balances;

    event SetManagerProxyEvent(address manager);
    event DelegateAssetEvent(address assetHash, uint64 nativeChainId, bytes nativeLockProxy, bytes nativeAssetHash);
    event UnlockEvent(address toAssetHash, address toAddress, uint256 amount);
    event LockEvent(address fromAssetHash, address fromAddress, uint64 toChainId, bytes toAssetHash, bytes toAddress, bytes txArgs);

    constructor(address ethCCMProxyAddr) public {
        managerProxyContract = ethCCMProxyAddr;
        emit SetManagerProxyEvent(managerProxyContract);
    }

    modifier onlyManagerContract() {
        IEthCrossChainManagerProxy ieccmp = IEthCrossChainManagerProxy(managerProxyContract);
        require(_msgSender() == ieccmp.getEthCrossChainManager(), "msgSender is not EthCrossChainManagerContract");
        _;
    }

    function delegateAsset(uint64 nativeChainId, bytes memory nativeLockProxy, bytes memory nativeAssetHash, uint256 delegatedSupply) public {
        require(nativeChainId != 0, "nativeChainId cannot be zero");
        require(nativeLockProxy.length != 0, "empty nativeLockProxy");
        require(nativeAssetHash.length != 0, "empty nativeAssetHash");

        address assetHash = _msgSender();
        bytes32 key = _getRegistryKey(assetHash, nativeChainId, nativeLockProxy, nativeAssetHash);

        require(registry[key] != true, "asset already registered");
        require(balances[key] == 0, "balance is not zero");
        require(_balanceFor(assetHash) == delegatedSupply, "controlled balance does not match delegatedSupply");

        registry[key] = true;

        RegisterAssetTxArgs memory txArgs = RegisterAssetTxArgs({
            assetHash: Utils.addressToBytes(assetHash),
            nativeAssetHash: nativeAssetHash
        });

        bytes memory txData = _serializeRegisterAssetTxArgs(txArgs);

        IEthCrossChainManager eccm = _getEccm();
        require(eccm.crossChain(nativeChainId, nativeLockProxy, "registerAsset", txData), "EthCrossChainManager crossChain executed error!");
        balances[key] = delegatedSupply;

        emit DelegateAssetEvent(assetHash, nativeChainId, nativeLockProxy, nativeAssetHash);
    }

    function registerAsset(bytes memory argsBs, bytes memory fromContractAddr, uint64 fromChainId) onlyManagerContract public returns (bool) {
        RegisterAssetTxArgs memory args = _deserializeRegisterAssetTxArgs(argsBs);

        bytes32 key = _getRegistryKey(Utils.bytesToAddress(args.nativeAssetHash), fromChainId, fromContractAddr, args.assetHash);

        require(registry[key] != true, "asset already registerd");
        registry[key] = true;

        return true;
    }

    /* @notice                  This function is meant to be invoked by the user,
    *                           a certain amount teokens will be locked in the proxy contract the invoker/msg.sender immediately.
    *                           Then the same amount of tokens will be unloked from target chain proxy contract at the target chain with chainId later.
    *  @param fromAssetHash     The asset hash in current chain
    *  @param toChainId         The target chain id
    *
    *  @param toAddress         The address in bytes format to receive same amount of tokens in target chain
    *  @param amount            The amount of tokens to be crossed from ethereum to the chain with chainId
    */
    function lock(
        address fromAssetHash,
        uint64 toChainId,
        bytes memory targetProxyHash,
        bytes memory toAssetHash,
        bytes memory toAddress,
        uint256 amount,
        bool deductFeeInLock,
        uint256 feeAmount,
        bytes memory feeAddress
    )
        public
        payable
        returns (bool)
    {
        require(toChainId != 0, "toChainId cannot be zero");
        require(targetProxyHash.length != 0, "empty targetProxyHash");
        require(toAssetHash.length != 0, "empty toAssetHash");
        require(toAddress.length != 0, "empty toAddress");
        require(amount != 0, "amount must be more than zero!");

        require(_transferToContract(fromAssetHash, amount), "transfer asset from fromAddress to lock_proxy contract  failed!");

        bytes32 key = _getRegistryKey(fromAssetHash, toChainId, targetProxyHash, toAssetHash);
        require(registry[key] == true, "asset not registered");

        TxArgs memory txArgs = TxArgs({
            fromAssetHash: Utils.addressToBytes(fromAssetHash),
            toAssetHash: toAssetHash,
            toAddress: toAddress,
            amount: amount,
            feeAmount: feeAmount,
            feeAddress: feeAddress
        });

        if (feeAmount != 0 && deductFeeInLock) {
            require(feeAddress.length != 0, "empty fee address");
            uint256 afterFeeAmount = amount.sub(feeAmount);
            require(_transferFromContract(fromAssetHash, Utils.bytesToAddress(feeAddress), feeAmount), "transfer asset from lock_proxy contract to feeAddress failed!");

            // set feeAmount to zero as fee has already been transferred
            txArgs.feeAmount = 0;
            txArgs.amount = afterFeeAmount;
        }

        bytes memory txData = _serializeTxArgs(txArgs);
        IEthCrossChainManager eccm = _getEccm();

        require(eccm.crossChain(toChainId, targetProxyHash, "unlock", txData), "EthCrossChainManager crossChain executed error!");
        balances[key] = balances[key].add(txArgs.amount);

        emit LockEvent(fromAssetHash, _msgSender(), toChainId, toAssetHash, toAddress, txData);

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
        address toAssetHash = Utils.bytesToAddress(args.toAssetHash);
        address toAddress = Utils.bytesToAddress(args.toAddress);

        bytes32 key = _getRegistryKey(toAssetHash, fromChainId, fromContractAddr, args.fromAssetHash);

        require(registry[key] == true, "asset not registered");
        require(balances[key] >= args.amount, "insufficient balance in registry");

        balances[key] = balances[key].sub(args.amount);

        uint256 afterFeeAmount = args.amount;
        if (args.feeAmount != 0) {
            afterFeeAmount = args.amount.sub(args.feeAmount);
            address feeAddress = Utils.bytesToAddress(args.feeAddress);

            // transfer feeAmount to feeAddress
            require(_transferFromContract(toAssetHash, feeAddress, args.feeAmount), "transfer asset from lock_proxy contract to feeAddress failed!");
            emit UnlockEvent(toAssetHash, feeAddress, args.feeAmount);
        }

        require(_transferFromContract(toAssetHash, toAddress, afterFeeAmount), "transfer asset from lock_proxy contract to toAddress failed!");

        emit UnlockEvent(toAssetHash, toAddress, args.amount);
        return true;
    }

    function _balanceFor(address fromAssetHash) public view returns (uint256) {
        if (fromAssetHash == address(0)) {
            // return address(this).balance; // this expression would result in error: Failed to decode output: Error: insufficient data for uint256 type
            address selfAddr = address(this);
            return selfAddr.balance;
        } else {
            ERC20Interface erc20Token = ERC20Interface(fromAssetHash);
            return erc20Token.balanceOf(address(this));
        }
    }
    function _getEccm() internal view returns (IEthCrossChainManager) {
      IEthCrossChainManagerProxy eccmp = IEthCrossChainManagerProxy(managerProxyContract);
      address eccmAddr = eccmp.getEthCrossChainManager();
      IEthCrossChainManager eccm = IEthCrossChainManager(eccmAddr);
      return eccm;
    }
    function _getRegistryKey(address assetHash, uint64 nativeChainId, bytes memory nativeLockProxy, bytes memory nativeAssetHash) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(
            assetHash,
            nativeChainId,
            nativeLockProxy,
            nativeAssetHash
        ));
    }
    function _transferToContract(address fromAssetHash, uint256 amount) internal returns (bool) {
        if (fromAssetHash == address(0)) {
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
            ZeroCopySink.WriteVarBytes(args.fromAssetHash),
            ZeroCopySink.WriteVarBytes(args.toAssetHash),
            ZeroCopySink.WriteVarBytes(args.toAddress),
            ZeroCopySink.WriteUint255(args.amount),
            ZeroCopySink.WriteUint255(args.feeAmount),
            ZeroCopySink.WriteVarBytes(args.feeAddress)
        );
        return buff;
    }

    function _serializeRegisterAssetTxArgs(RegisterAssetTxArgs memory args) internal pure returns (bytes memory) {
        bytes memory buff;
        buff = abi.encodePacked(
            ZeroCopySink.WriteVarBytes(args.assetHash),
            ZeroCopySink.WriteVarBytes(args.nativeAssetHash)
        );
        return buff;
    }

    function _deserializeRegisterAssetTxArgs(bytes memory valueBs) internal pure returns (RegisterAssetTxArgs memory) {
        RegisterAssetTxArgs memory args;
        uint256 off = 0;
        (args.assetHash, off) = ZeroCopySource.NextVarBytes(valueBs, off);
        (args.nativeAssetHash, off) = ZeroCopySource.NextVarBytes(valueBs, off);
        return args;
    }

    function _deserializeTxArgs(bytes memory valueBs) internal pure returns (TxArgs memory) {
        TxArgs memory args;
        uint256 off = 0;
        (args.fromAssetHash, off) = ZeroCopySource.NextVarBytes(valueBs, off);
        (args.toAssetHash, off) = ZeroCopySource.NextVarBytes(valueBs, off);
        (args.toAddress, off) = ZeroCopySource.NextVarBytes(valueBs, off);
        (args.amount, off) = ZeroCopySource.NextUint255(valueBs, off);
        (args.feeAmount, off) = ZeroCopySource.NextUint255(valueBs, off);
        (args.feeAddress, off) = ZeroCopySource.NextVarBytes(valueBs, off);
        return args;
    }
}