pragma solidity >=0.4.25;

/**
 * @dev Interface of the EthCrossChainManager contract for business contract like LockProxy to request cross chain transaction
 */
interface IEthCrossChainManager {
    function crossChain(uint64 _toChainId, bytes _toContract, bytes _method, bytes _txData) external returns (bool);
}
