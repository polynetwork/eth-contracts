pragma solidity ^0.5.0;

/**
 * @dev Interface of the EthCrossChainData contract, the implementation is in EthCrossChainData.sol
 */
interface IEthCrossChainData {
    function putCurEpochStartHeight(uint32 curEpochStartHeight) external returns (bool);
    function getCurEpochStartHeight() external view returns (uint32);
    function putCurEpochConPubKeyBytes(bytes calldata curEpochPkBytes) external returns (bool);
    function getCurEpochConPubKeyBytes() external view returns (bytes memory);
    function markFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) external returns (bool);
    function checkIfFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) external view returns (bool);
    function getEthTxHashIndex() external view returns (uint256);
    function putEthTxHash(bytes32 ethTxHash) external returns (bool);
    function putExtraData(bytes32 key1, bytes32 key2, bytes calldata value) external returns (bool);
    function getExtraData(bytes32 key1, bytes32 key2) external view returns (bytes memory);
    function transferOwnership(address newOwner) external;
    function pause() external returns (bool);
    function unpause() external returns (bool);
    function paused() external view returns (bool);
    // Not used currently by ECCM
    function getEthTxHash(uint256 ethTxHashIndex) external view returns (bytes32);
}