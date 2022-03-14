pragma solidity 0.5.17;

import "../../core/cross_chain_manager/libs/EthCrossChainUtils.sol";

contract ECCUtilsMock {

    function merkleProve(bytes memory _auditPath, bytes32 _root) public pure returns (bytes memory) {
        return ECCUtils.merkleProve(_auditPath, _root);
    }

    function verifySig(bytes memory _rawHeader, bytes memory _sigList, address[] memory _keepers, uint _m) public pure returns (bool){
        return ECCUtils.verifySig(_rawHeader, _sigList, _keepers, _m);
    }

    function deserializeMerkleValue(bytes memory _valueBs) public pure 
    returns (
        bytes memory polyTxHash,
        uint64 fromChainID,
        bytes memory txHash,
        // bytes memory crossChainId, // STACK TO DEEP
        bytes memory fromContract,
        uint64 toChainId,
        bytes memory toContract,
        bytes memory method,
        bytes memory args
    ) {
        uint256 off = 0;

        (polyTxHash, off) = ZeroCopySource.NextVarBytes(_valueBs, off);

        (fromChainID, off) = ZeroCopySource.NextUint64(_valueBs, off);

        (txHash, off) = ZeroCopySource.NextVarBytes(_valueBs, off);
        
        (, off) = ZeroCopySource.NextVarBytes(_valueBs, off);

        (fromContract, off) = ZeroCopySource.NextVarBytes(_valueBs, off);

        (toChainId, off) = ZeroCopySource.NextUint64(_valueBs, off);

        (toContract, off) = ZeroCopySource.NextVarBytes(_valueBs, off);

        (method, off) = ZeroCopySource.NextVarBytes(_valueBs, off);

        (args, off) = ZeroCopySource.NextVarBytes(_valueBs, off);

    }

    function deserializeHeader(bytes memory _headerBs) public pure 
    returns (
        uint32 version,
        uint64 chainId,
        bytes32 prevBlockHash,
        bytes32 transactionsRoot,
        bytes32 crossStatesRoot,
        bytes32 blockRoot,
        uint32 timestamp,
        uint32 height,
        uint64 consensusData,
        bytes memory consensusPayload,
        bytes20 nextBookkeeper
    ) {
        uint256 off = 0;
        (version, off)  = ZeroCopySource.NextUint32(_headerBs, off);

        (chainId, off) = ZeroCopySource.NextUint64(_headerBs, off);

        (prevBlockHash, off) = ZeroCopySource.NextHash(_headerBs, off);

        (transactionsRoot, off) = ZeroCopySource.NextHash(_headerBs, off);

        (crossStatesRoot, off) = ZeroCopySource.NextHash(_headerBs, off);

        (blockRoot, off) = ZeroCopySource.NextHash(_headerBs, off);

        (timestamp, off) = ZeroCopySource.NextUint32(_headerBs, off);

        (height, off) = ZeroCopySource.NextUint32(_headerBs, off);

        (consensusData, off) = ZeroCopySource.NextUint64(_headerBs, off);

        (consensusPayload, off) = ZeroCopySource.NextVarBytes(_headerBs, off);

        (nextBookkeeper, off) = ZeroCopySource.NextBytes20(_headerBs, off);

    }

    function getHeaderHash(bytes memory rawHeader) public pure returns (bytes32) {
        return ECCUtils.getHeaderHash(rawHeader);
    }
}