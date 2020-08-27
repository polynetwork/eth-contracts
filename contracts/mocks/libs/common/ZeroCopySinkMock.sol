pragma solidity ^0.5.0;

import "../../../../contracts/libs/common/ZeroCopySink.sol";

contract ZeroCopySinkMock {
    uint N = 10;
    function WriteBool(bool _b) public returns (bytes memory) {
        bytes memory executeResult;
        for (uint i = 0; i <= N; i++) {
            executeResult = ZeroCopySink.WriteBool(_b);
        }
        return ZeroCopySink.WriteBool(_b);
    }

    function WriteByte(byte _b) public returns (bytes memory) {
        bytes memory executeResult;
        for (uint i = 0; i <= N; i++) {
            executeResult = ZeroCopySink.WriteByte(_b);
        }
        return ZeroCopySink.WriteByte(_b);
    }

    function WriteUint8(uint8 _v) public returns (bytes memory) {
        bytes memory executeResult;
        for (uint i = 0; i <= N; i++) {
            executeResult = ZeroCopySink.WriteUint8(_v);
        }
        return ZeroCopySink.WriteUint8(_v);
    }

    function WriteUint16(uint16 _v) public returns (bytes memory) {
        bytes memory executeResult;
        for (uint i = 0; i <= N; i++) {
            executeResult = ZeroCopySink.WriteUint16(_v);
        }
        return ZeroCopySink.WriteUint16(_v);
    }

    function WriteUint32(uint32 _v) public returns (bytes memory) {
        bytes memory executeResult;
        for (uint i = 0; i <= N; i++) {
            executeResult = ZeroCopySink.WriteUint32(_v);
        }
        return ZeroCopySink.WriteUint32(_v);
    }

    function WriteUint64(uint64 _v) public returns (bytes memory) {
        bytes memory executeResult;
        for (uint i = 0; i <= N; i++) {
            executeResult = ZeroCopySink.WriteUint64(_v);
        }
        return ZeroCopySink.WriteUint64(_v);
    }
    function WriteUint255(uint256 _v) public returns (bytes memory) {
        bytes memory executeResult;
        for (uint i = 0; i <= N; i++) {
            executeResult = ZeroCopySink.WriteUint255(_v);
        }
        return ZeroCopySink.WriteUint255(_v);
    }

    function WriteVarBytes(bytes memory _b) public returns (bytes memory) {
        bytes memory executeResult;
        for (uint i = 0; i <= N; i++) {
            executeResult = ZeroCopySink.WriteVarBytes(_b);
        }
        return ZeroCopySink.WriteVarBytes(_b);
    }

    function WriteVarUint(uint64 _v) public returns (bytes memory) {
        bytes memory executeResult;
        for (uint i = 0; i <= N; i++) {
            executeResult = ZeroCopySink.WriteVarUint(_v);
        }
        return ZeroCopySink.WriteVarUint(_v);
    }
}
