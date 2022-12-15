pragma solidity ^0.5.0;

import "../interface/IEthCrossChainData.sol";
import "../libs/ECCUtils/EthCrossChainUtils.sol";
import "../../../libs/common/ZeroCopySink.sol";

contract VerifyHelper {

    address public EthCrossChainDataAddress;
    bytes public ZionCrossChainManagerAddress;

    constructor(address ethCrossChainDataAddress, bytes memory zionCrossChainManagerAddress) public {
        EthCrossChainDataAddress = ethCrossChainDataAddress;
        ZionCrossChainManagerAddress = zionCrossChainManagerAddress;
    }

    function verifyHeader(
        bytes memory rawHeader,
        bytes memory rawSeals,
        bytes memory accountProof, 
        bytes memory storageProof,
        bytes memory rawCrossTx
    ) public view returns (bytes memory) {
        ECCUtils.Header memory header = ECCUtils.decodeHeader(rawHeader);
        ECCUtils.CrossTx memory crossTx = ECCUtils.decodeCrossTx(rawCrossTx);
        IEthCrossChainData eccd = IEthCrossChainData(EthCrossChainDataAddress);
        
        address[] memory validators = ECCUtils.decodeValidators(eccd.getCurEpochValidatorPkBytes());
        
        // verify block.height
        require(header.number>=eccd.getCurEpochStartHeight(), "Invalid block height, before CurEpochStartHeight");
        require(header.number<=eccd.getCurEpochEndHeight(), "Invalid block height, after CurEpochEndHeight");
        
        // verify header
        require(ECCUtils.verifyHeader(keccak256(rawHeader), rawSeals, validators), "Verify header failed");
        
        // verify proof
        bytes memory storageIndex = ECCUtils.getCrossTxStorageSlot(crossTx);
        bytes memory storageValue = ECCUtils.verifyAccountProof(accountProof, header.root, ZionCrossChainManagerAddress, storageProof, storageIndex);
        require(ECCUtils.bytesToBytes32(storageValue) == keccak256(rawCrossTx), "Verify proof failed");

        return _serialize(crossTx);
    }

    function _serialize(
        ECCUtils.CrossTx memory crossTx
    ) internal pure returns (bytes memory) {
        bytes memory buff;
        buff = abi.encodePacked(
            ZeroCopySink.WriteVarBytes(crossTx.txHash),
            ZeroCopySink.WriteUint64(crossTx.fromChainID),
            ZeroCopySink.WriteVarBytes(crossTx.crossTxParam.txHash),
            ZeroCopySink.WriteVarBytes(crossTx.crossTxParam.crossChainId),
            ZeroCopySink.WriteVarBytes(crossTx.crossTxParam.fromContract),
            ZeroCopySink.WriteUint64(crossTx.crossTxParam.toChainId),
            ZeroCopySink.WriteVarBytes(crossTx.crossTxParam.toContract),
            ZeroCopySink.WriteVarBytes(crossTx.crossTxParam.method),
            ZeroCopySink.WriteVarBytes(crossTx.crossTxParam.args)
            );
        return buff;
    }
}