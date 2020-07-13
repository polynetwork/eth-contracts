pragma solidity ^0.5.0;

import "../../../../contracts/libs/common/ZeroCopySink.sol";

contract ZeroCopySinkMock {

    function WriteBool(bool _b) public pure returns (bytes memory) {
        return ZeroCopySink.WriteBool(_b);
    }

    function WriteByte(byte _b) public pure returns (bytes memory) {
        return ZeroCopySink.WriteByte(_b);
    }

    function WriteUint8(uint8 _v) public pure returns (bytes memory) {
        return ZeroCopySink.WriteUint8(_v);
    }

    function WriteUint16(uint16 _v) public pure returns (bytes memory) {
        return ZeroCopySink.WriteUint16(_v);
    }

    function WriteUint32(uint32 _v) public pure returns (bytes memory) {
        return ZeroCopySink.WriteUint32(_v);
    }

    function WriteUint64(uint64 _v) public pure returns (bytes memory) {
        return ZeroCopySink.WriteUint64(_v);
    }
    function WriteUint255(uint256 _v) public pure returns (bytes memory) {
        return ZeroCopySink.WriteUint255(_v);
    }

    function WriteVarBytes(bytes memory _b) public pure returns (bytes memory) {
        return ZeroCopySink.WriteVarBytes(_b);
    }

    function WriteVarUint(uint64 _v) public pure returns (bytes memory) {
        return ZeroCopySink.WriteVarUint(_v);
    }
}
