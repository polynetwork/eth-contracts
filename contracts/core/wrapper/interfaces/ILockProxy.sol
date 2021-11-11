// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.5.0;

interface ILockProxy {
    function managerProxyContract() external view returns (address);
    function proxyHashMap(uint64) external view returns (bytes memory);
    function assetHashMap(address, uint64) external view returns (bytes memory);
    function getBalanceFor(address) external view returns (uint256);
    function setManagerProxy(
        address eccmpAddr
    ) external;
    
    function bindProxyHash(
        uint64 toChainId, 
        bytes calldata targetProxyHash
    ) external returns (bool);

    function bindAssetHash(
        address fromAssetHash, 
        uint64 toChainId, 
        bytes calldata toAssetHash
    ) external returns (bool);

    function lock(
        address fromAssetHash, 
        uint64 toChainId, 
        bytes calldata toAddress, 
        uint256 amount
    ) external payable returns (bool);
}