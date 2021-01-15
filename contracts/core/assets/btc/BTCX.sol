pragma solidity ^0.5.0;
import "./../../../libs/GSN/Context.sol";
import "./../../../libs/common/ZeroCopySource.sol";
import "./../../../libs/common/ZeroCopySink.sol";
import "./../../../libs/utils/Utils.sol";
import "./../../cross_chain_manager/interface/IEthCrossChainManager.sol";
import "./../../cross_chain_manager/interface/IEthCrossChainManagerProxy.sol";
import "./../../../libs/token/ERC20/ERC20Extended.sol";
contract BTCX is ERC20Extended {
    struct TxArgs {
        bytes toAddress;
        uint64 amount;
    }

    bytes public redeemScript;
    uint64 public minimumLimit;
    event UnlockEvent(address toAssetHash, address toAddress, uint64 amount);
    event LockEvent(address fromAssetHash, address fromAddress, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint64 amount);
    
    constructor (bytes memory _redeemScript) public ERC20Detailed("BTC Token", "BTCX", 8) {
        operator = _msgSender();
        redeemScript = _redeemScript;
    }
    function setMinimumLimit(uint64 minimumTransferLimit) onlyOperator public returns (bool) {
        minimumLimit = minimumTransferLimit;
        return true;
    }
    /* @notice                  This function is meant to be invoked by the ETH crosschain management contract,
    *                           then mint a certin amount of tokens to the designated address since a certain amount 
    *                           was burnt from the source chain invoker.
    *  @param argsBs            The argument bytes recevied by the ethereum business contract, need to be deserialized.
    *                           based on the way of serialization in the source chain contract.
    *  @param fromContractAddr  The source chain contract address
    *  @param fromChainId       The source chain id
    */
    function unlock(bytes memory argsBs, bytes memory fromContractAddr, uint64 fromChainId) onlyManagerContract public returns (bool) {
        TxArgs memory args = _deserializeTxArgs(argsBs);
        require(fromContractAddr.length != 0, "from asset contract address cannot be empty");
        require(Utils.equalStorage(bondAssetHashes[fromChainId], fromContractAddr), "From contract address error!");
        
        address toAddress = Utils.bytesToAddress(args.toAddress);
        require(mint(toAddress, uint256(args.amount)), "mint BTCX in unlock method failed!");
        
        emit UnlockEvent(address(this), toAddress, args.amount);
        return true;
    }

    /* @notice                  This function is meant to be invoked by the user,
    *                           a certin amount teokens will be burnt from the invoker/msg.sender immediately.
    *                           Then the same amount of tokens will be mint at the target chain with chainId later.
    *  @param toChainId         The target chain id
    *                           
    *  @param toUserAddr        The address in bytes format to receive same amount of tokens in target chain 
    *  @param amount            The amount of tokens to be crossed from ethereum to the chain with chainId
    */
    function lock(uint64 toChainId, bytes memory toUserAddr, uint64 amount) public returns (bool) {
        TxArgs memory txArgs = TxArgs({
            toAddress: toUserAddr,
            amount: amount
        });

        bytes memory txData;
        // if toChainId is BTC chain, put redeemScript into Args
        if (toChainId == 1) {
            require(amount >= minimumLimit, "btcx amount should be greater than 2000");
            txData = _serializeToBtcTxArgs(txArgs, redeemScript);
        } else {
            txData = _serializeTxArgs(txArgs);
        }
        

        require(burn(uint256(amount)), "Burn msg.sender BTCX tokens failed");
    
        IEthCrossChainManagerProxy eccmp = IEthCrossChainManagerProxy(managerProxyContract);
        address eccmAddr = eccmp.getEthCrossChainManager();
        IEthCrossChainManager eccm = IEthCrossChainManager(eccmAddr);
        
        require(eccm.crossChain(toChainId, bondAssetHashes[toChainId], "unlock", txData), "EthCrossChainManager crossChain executed error!");

        emit LockEvent(address(this), _msgSender(), toChainId, bondAssetHashes[toChainId], toUserAddr, amount);
        
        return true;

    }

    function _serializeToBtcTxArgs(TxArgs memory args, bytes memory redeemScript) internal pure returns (bytes memory) {
        bytes memory buff;
        buff = abi.encodePacked(
            ZeroCopySink.WriteVarBytes(args.toAddress),
            ZeroCopySink.WriteUint64(args.amount),
            ZeroCopySink.WriteVarBytes(redeemScript)
            );
        return buff;
    }
    function _serializeTxArgs(TxArgs memory args) internal pure returns (bytes memory) {
        bytes memory buff;
        buff = abi.encodePacked(
            ZeroCopySink.WriteVarBytes(args.toAddress),
            ZeroCopySink.WriteUint64(args.amount)
            );
        return buff;
    }

    function _deserializeTxArgs(bytes memory valueBs) internal pure returns (TxArgs memory) {
        TxArgs memory args;
        uint256 off = 0;
        (args.toAddress, off) = ZeroCopySource.NextVarBytes(valueBs, off);
        (args.amount, off) = ZeroCopySource.NextUint64(valueBs, off);
        return args;
    }
}