pragma solidity ^0.5.0;

import "../../../../contracts/libs/common/ZeroCopySource.sol";

contract ZeroCopySourceMock {
    uint N = 10;
    function NextBool(bytes memory _b, uint256 _off) public returns (bool, uint256) {
        bool res;
        uint256 offset = 0;
        for (uint i = 0; i <= N; i++) {
            (res, offset) = ZeroCopySource.NextBool(_b, _off);
        }
        return ZeroCopySource.NextBool(_b, _off);
    }

    function NextByte(bytes memory _b, uint256 _off) public returns (byte, uint256) {
        return ZeroCopySource.NextByte(_b, _off);
    }

    function NextUint8(bytes memory _b, uint256 _off) public returns (uint8, uint256) {
        return ZeroCopySource.NextUint8(_b, _off);
    }

    function NextUint16(bytes memory _b, uint256 _off) public returns (uint16, uint256) {
        return ZeroCopySource.NextUint16(_b, _off);
    }

    function NextUint32(bytes memory _b, uint256 _off) public returns (uint32, uint256) {
        return ZeroCopySource.NextUint32(_b, _off);
    }

    function NextUint64(bytes memory _b, uint256 _off) public returns (uint64, uint256) {
        return ZeroCopySource.NextUint64(_b, _off);
    }

    function NextVarBytes(bytes memory _b, uint256 _off) public returns (bytes memory, uint256) {
        return ZeroCopySource.NextVarBytes(_b, _off);
    }

    function NextHash(bytes memory _b, uint256 _off) public returns (bytes32, uint256) {
        return ZeroCopySource.NextHash(_b, _off);
    }

//    function NextAddress(bytes _b, uint256 _off) public pure returns (bytes20, uint256) {
//        return ZeroCopySource.NextAddress(_b, _off);
//    }

    function NextBytes20(bytes memory _b, uint256 _off) public returns (bytes20, uint256) {
        return ZeroCopySource.NextBytes20(_b, _off);
    }

    function NextUint255(bytes memory _b, uint256 _off) public returns (uint256, uint256) {
        return ZeroCopySource.NextUint255(_b, _off);
    }
    function NextVarUint(bytes memory _b, uint256 _off) public returns (uint, uint256) {
        return ZeroCopySource.NextVarUint(_b, _off);
    }
}
