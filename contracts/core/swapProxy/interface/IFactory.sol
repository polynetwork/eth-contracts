  
// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.5.0;

interface IFactory {
    function isPool(address addr) external view returns (bool);
}