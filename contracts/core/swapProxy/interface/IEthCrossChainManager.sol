pragma solidity ^0.5.0;

/**
 * @dev Interface of the EthCrossChainManager contract for business contract like LockProxy to request cross chain transaction
 */
interface IEthCrossChainManager {
    function crossChain(uint64 _toChainId, bytes calldata _toContract, bytes calldata _method, bytes calldata _txData) external returns (bool);
}
