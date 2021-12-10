pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract AbiEncoder {
    function encodeWhiteList(address _contract, bytes[] memory _methods) public pure returns(bytes memory) {
        return abi.encode(_contract,_methods);
    }
}