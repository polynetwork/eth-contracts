pragma solidity >=0.4.25;

/**
 * @dev Interface of the EthCrossChainManagerProxy for business contract like LockProxy to obtain the reliable EthCrossChainManager contract hash.
 */
interface IEthCrossChainManagerProxy {
    function getEthCrossChainManager() external view returns (address);
}
