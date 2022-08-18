pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "../../core/cross_chain_manager/libs/ECCUtils/EthCrossChainUtils.sol";

contract ECCUtilsMock {
    
    function k256(bytes memory raw) public pure returns(bytes32) {
        return keccak256(raw);
    }
    
    function verifyHeader(bytes32 headerHash, bytes memory rawSeals, address[] memory validators) public pure returns(bool) {
        return ECCUtils.verifyHeader(headerHash, rawSeals, validators);
    }
    
    function verifySeal(bytes32 headerHash, bytes memory seal) public pure returns(address signer) {
        return ECCUtils.verifySeal(headerHash, seal);
    }

    function hasEnoughSigners(address[] memory _validators, address[] memory _signers) public pure returns(bool) {
        return ECCUtils.hasEnoughSigners(_validators, _signers);         
    }
    
    function decodeHeader(bytes memory rawHeader) public view returns(bytes memory root, uint256 number) {
        ECCUtils.Header memory header = ECCUtils.decodeHeader(rawHeader);
        return (header.root, header.number);
    }
    
    function getCrossTxStorageSlot(
        bytes memory ziontxHash,
        uint64 toChainId
    ) public pure returns(bytes memory slotIndex) {
        ECCUtils.TxParam memory param;
        ECCUtils.CrossTx memory ctx;
        // param.txHash = sourceTxHash;
        // param.crossChainId = crossChainId;
        // param.fromContract = fromContract;
        param.toChainId = toChainId;
        // param.toContract = toContract;
        // param.method = method;
        // param.args = args;
        ctx.txHash = ziontxHash;
        // ctx.fromChainID = fromChainID;
        ctx.crossTxParam = param;
        return ECCUtils.getCrossTxStorageSlot(ctx);
    }

    // little endian
    function getUint64Bytes(uint64 num) public pure returns(bytes8 res) {
        return ECCUtils.getUint64Bytes(num);
    }
    
    function bytes32ToBytes(bytes32 raw) public pure returns(bytes memory res) {
        return ECCUtils.bytes32ToBytes(raw);
    }

    function uint256ToBytes(uint256 _value) public pure returns (bytes memory bs) {
        return ECCUtils.uint256ToBytes(_value);
    }
    
    function addressToBytes(address _addr) public pure returns (bytes memory bs){
        return ECCUtils.addressToBytes(_addr);
    }

    function bytesToBytes32(bytes memory _bs) public pure returns (bytes32 value) {
        return ECCUtils.bytesToBytes32(_bs);
    }

    function bytesToUint256(bytes memory _bs) public pure returns (uint256 value) {
        return ECCUtils.bytesToUint256(_bs);
    }

    function bytesToAddress(bytes memory _bs) public pure returns (address addr) {
        return ECCUtils.bytesToAddress(_bs);
    }

    function getHeaderValidatorsAndEpochEndHeight(bytes memory rawHeader) public view returns(uint64 endHeight, address[] memory validators) {
        return ECCUtils.getHeaderValidatorsAndEpochEndHeight(rawHeader);
    }
    
    function decodeValidators(bytes memory validatorBytes) public pure returns(address[] memory validators) {
        return ECCUtils.decodeValidators(validatorBytes);
    }
    
    function encodeValidators(address[] memory validators) public pure returns(bytes memory validatorBytes) {
        return ECCUtils.encodeValidators(validators);
    }
    
    function encodeTxParam(
        bytes memory txHash,
        bytes memory crossChainId,
        bytes memory fromContract,
        uint64 toChainId,
        bytes memory toContract,
        bytes memory method,
        bytes memory args
    ) public pure returns(bytes memory rawParam) {
        return ECCUtils.encodeTxParam(txHash, crossChainId, fromContract, toChainId, toContract, method, args);
    }
    
    function decodeTxParam(bytes memory rawParam) public view returns(
        bytes memory sourceTxHash,
        bytes memory crossChainId,
        bytes memory fromContract,
        uint64 toChainId,
        bytes memory toContract,
        bytes memory method,
        bytes memory args
    ) {
        ECCUtils.TxParam memory tmp = ECCUtils.decodeTxParam(rawParam);
        return (tmp.txHash, tmp.crossChainId, tmp.fromContract, tmp.toChainId, tmp.toContract, tmp.method, tmp.args);
    }
    
    function decodeCrossTx(bytes memory rawTx) public view returns(
        bytes memory zionTxHash,
        uint64 fromChainID,
        bytes memory sourceTxHash,
        bytes memory crossChainId,
        bytes memory fromContract,
        uint64 toChainId,
        bytes memory toContract,
        bytes memory method,
        bytes memory args
    ) { 
        ECCUtils.CrossTx memory tmp = ECCUtils.decodeCrossTx(rawTx);
        zionTxHash = tmp.txHash;
        fromChainID = tmp.fromChainID;
        ECCUtils.TxParam memory txData = tmp.crossTxParam;
        sourceTxHash = txData.txHash; 
        crossChainId = txData.crossChainId;
        fromContract = txData.fromContract;
        toChainId = txData.toChainId;
        toContract = txData.toContract;
        method = txData.method;
        args = txData.args;
    }
    
    function rlpGetNextBytes(bytes memory raw, uint offset) public view returns (bytes memory res, uint _offset, bytes memory _raw){
        (res,_offset) = ECCUtils.rlpGetNextBytes(raw, offset);
        return (res,_offset,raw);
    } 
    
    function rlpGetNextBytes32(bytes memory raw, uint offset) public pure returns (bytes32 res, uint _offset, bytes memory _raw){
        (res,_offset) = ECCUtils.rlpGetNextBytes32(raw, offset);
        return (res,_offset,raw);
    } 
    
    function rlpGetNextAddress(bytes memory raw, uint offset) public pure returns (address res, uint _offset, bytes memory _raw){
        (res,_offset) = ECCUtils.rlpGetNextAddress(raw, offset);
        return (res,_offset,raw);
    } 
    
    function rlpGetNextUint64(bytes memory raw, uint offset) public pure returns (uint64 res, uint _offset, bytes memory _raw){
        (res,_offset) = ECCUtils.rlpGetNextUint64(raw, offset);
        return (res,_offset,raw);
    } 
    
    function rlpGetNextUint256(bytes memory raw, uint offset) public pure returns (uint256 res, uint _offset, bytes memory _raw){
        (res,_offset) = ECCUtils.rlpGetNextUint256(raw, offset);
        return (res,_offset,raw);
    } 
    
    function verifyAccountProof(
        bytes memory _accountProof, 
        bytes memory _headerRoot,
        bytes memory _address,
        bytes memory _storageProof,
        bytes memory _storageIndex
    ) public pure returns(bytes memory value) {
        return ECCUtils.verifyAccountProof(_accountProof,_headerRoot,_address,_storageProof,_storageIndex);
    }
    
    function verifyProof(bytes memory _proof, bytes memory _key, bytes memory _root) public pure returns(bytes memory value) {
        return ECCUtils.verifyProof(_proof,_key,_root);
    }
    
    function bytesToHex(bytes memory keyBytes) public pure returns(bytes memory keyHex) {
        return ECCUtils.bytesToHex(keyBytes);
    }
    
    function compactToHex(bytes memory keyCompact) public pure returns(bytes memory keyHex) {
        return ECCUtils.compactToHex(keyCompact);
    }
    
    function hexToCompact(bytes memory keyHex) public pure returns(bytes memory keyCompact) {
        return ECCUtils.hexToCompact(keyHex);
    }
    
    function checkNodeHash(bytes memory raw,uint offset,bytes memory hash) public pure returns(bool flag) {
        return ECCUtils.checkNodeHash(raw, offset, hash);
    }
    
    function takeOneByte(bytes memory _buf) public pure returns(bytes memory buf,uint i) {
        return ECCUtils.takeOneByte(_buf);
    }
    
    function compareKey(bytes memory key,bytes memory element) public pure returns(bytes memory keySlice,bool isIn) {
        return ECCUtils.compareKey(key,element);
    }
    
    function rlpSplit(bytes memory raw,uint offset) public pure returns(bytes memory res, uint offset_) {
        return ECCUtils.rlpSplit(raw,offset);
    }

    function rlpReadKind(bytes memory raw, uint offset) public pure returns(uint size ,uint offset_) {
        return ECCUtils.rlpReadKind(raw, offset);
    }
}
