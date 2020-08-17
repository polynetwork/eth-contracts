pragma solidity ^0.5.0;

import "./../../../../contracts/libs/utils//Utils.sol";

contract UtilsMock {
    bytes public storageBytes;
    constructor() public {
        storageBytes = bytes("I am global in global storage");
    }
    function bytesToBytes32(bytes memory _b) public pure returns (bytes32) {
        return Utils.bytesToBytes32(_b);
    }

    function bytesToUint256(bytes memory  _v) public pure returns (uint256) {
        return Utils.bytesToUint256(_v);
    }

    function uint256ToBytes(uint256 _v) public pure returns (bytes memory) {
        return Utils.uint256ToBytes(_v);
    }

    function bytesToAddress(bytes memory _v) public pure returns (address) {
        return Utils.bytesToAddress(_v);
    }

    function addressToBytes(address _v) public pure returns (bytes memory) {
        return Utils.addressToBytes(_v);
    }
    function hashLeaf(bytes memory _data) public pure returns (bytes32) {
        return Utils.hashLeaf(_data);
    }
    function hashChildren(bytes32 left, bytes32 right) public pure returns (bytes32) {
        return Utils.hashChildren(left, right);
    }
    function equalStorage( bytes memory memoryBytes) public view returns (bool) {
        return Utils.equalStorage(storageBytes, memoryBytes);
    }
    function slice(bytes memory _bytes, uint _start, uint _length )public pure returns (bytes memory){
        return Utils.slice(_bytes, _start, _length);
    }

    function containMAddresses(address[] memory keepers, address[] memory signers, uint m) public pure returns(bool){
        return Utils.containMAddresses(keepers, signers, m);
    }
}
