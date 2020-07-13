pragma solidity ^0.5.0;

/**
 * @dev Interface of the ERC20 standard as defined in the EIP. Does not include
 * the optional functions; to access them see {ERC20Detailed}.
 */
interface IEthCrossChainData {
    function putMCHeaderBytes(uint64 _height, bytes calldata _rawHeader) external returns (bool);
    function getMCHeaderBytes(uint64 _height) external view returns (bytes memory);
    function putMCKeeperPubKeybytes(uint64 _height, bytes calldata _keepersBytes) external returns (bool);
    function getMCKeeperPubKeybytes(uint64 _height) external view returns (bytes memory);
    function getEthTxHashIndex() external view returns (uint256);
    function putEthTxHash(bytes32 _ethTxHash) external returns (bool);
    function getEthTxHash(uint256 _ethTxHashIndex) external view returns (bytes32);
    function putCrossChainTxExist(bytes32 _crossChainTx) external returns (bool);
    function getCrossChainTxExist(bytes32 _crossChainTx) external view returns (bool);
    function putLatestHeight(uint64 _latestHeight) external returns (bool);
    function getLatestHeight() external view returns (uint64);
    function putLatestBookKeeperHeight(uint64 _latestBookKeeperHeight) external returns (bool);
    function getLatestBookKeeperHeight() external view returns (uint64);
    function putMCKeeperHeight(uint64 height) external returns (bool);
    function getMCKeeperHeight() external view returns (uint64[] memory);
    function putMCKeeperPubBytes(uint64 height, bytes calldata _MCKeeperPubBytes) external returns (bool);
    function getMCKeeperPubBytes(uint64 height) external view returns (bytes memory);
    function initGenesisBlock() external returns (bool);
    function isGenesisBlockInited() external view returns (bool);
    function putExtraData(bytes32 key1, bytes32 key2, bytes calldata value) external returns (bool);
    function getExtraData(bytes32 key1, bytes32 key2) external view returns (bytes memory);
    function transferOwnership(address newOwner) external;
    function pause() external returns (bool);
    function unpause() external returns (bool);
    function paused() external view returns (bool);
}