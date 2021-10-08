pragma solidity ^0.5.0;
import "./../../../libs/common/ZeroCopySource.sol";
import "./../../../libs/common/ZeroCopySink.sol";
import "./../../../libs/utils/Utils.sol";
contract t  {
    
    uint constant POLYCHAIN_PUBKEY_LEN = 67;
    uint constant POLYCHAIN_SIGNATURE_LEN = 65;
    
    struct Header {
        uint32 version;
        uint64 chainId;
        uint32 timestamp;
        uint32 height;
        uint64 consensusData;
        bytes32 prevBlockHash;
        bytes32 transactionsRoot;
        bytes32 crossStatesRoot;
        bytes32 blockRoot;
        bytes consensusPayload;
        bytes20 nextBookkeeper;
    }
    
    constructor(bytes memory raw) public {
        ripemd160(raw);
    }
    
    function initG(bytes memory rawHeader, bytes memory pubKeyList) public pure returns(bytes20 a, bytes20 b) {
        Header memory header = deserializeHeader(rawHeader);
        return (header.nextBookkeeper,b);
        (bytes20 nextBookKeeper, ) = verifyPubkey(pubKeyList);
        return (header.nextBookkeeper, nextBookKeeper);
    }
    
    function s(bytes memory buff) public pure returns(bytes32) {
        return sha256(buff);
    }
    
    function es(bytes memory buff) public pure returns(bytes memory) {
        return abi.encodePacked(sha256(buff));
    }
    
    function res(bytes memory buff) public pure returns(bytes20) {
        return ripemd160(abi.encodePacked(sha256(buff)));
    }
    
    function _getBookKeeper(uint _keyLen, uint _m, bytes memory _pubKeyList) internal pure returns (bytes20 a, address[] memory b){
         bytes memory buff;
         buff = ZeroCopySink.WriteUint16(uint16(_keyLen));
         address[] memory keepers = new address[](_keyLen);
         bytes32 hash;
         bytes memory publicKey;
         for(uint i = 0; i < _keyLen; i++){
             publicKey = Utils.slice(_pubKeyList, i*POLYCHAIN_PUBKEY_LEN, POLYCHAIN_PUBKEY_LEN);
             buff =  abi.encodePacked(buff, ZeroCopySink.WriteVarBytes(Utils.compressMCPubKey(publicKey)));
             hash = keccak256(Utils.slice(publicKey, 3, 64));
             keepers[i] = address(uint160(uint256(hash)));
         }
         buff = abi.encodePacked(buff, ZeroCopySink.WriteUint16(uint16(_m)));
         // ERROR HERE 
         bytes20  nextBookKeeper = ripemd160(abi.encodePacked(sha256(buff)));
         // HERE ERROR
         return (nextBookKeeper, keepers);
    }

    /* @notice              Verify public key derived from Poly chain
    *  @param _pubKeyList   serialized consensus node public key list
    *  @param _sigList      consensus node signature list
    *  @return              return two element: next book keeper, consensus node signer addresses
    */
    function verifyPubkey(bytes memory _pubKeyList) internal pure returns (bytes20 a, address[] memory b) {
        require(_pubKeyList.length % POLYCHAIN_PUBKEY_LEN == 0, "_pubKeyList length illegal!");
        uint n = _pubKeyList.length / POLYCHAIN_PUBKEY_LEN;
        require(n >= 1, "too short _pubKeyList!");
        return _getBookKeeper(n, n - (n - 1) / 3, _pubKeyList);
    }
    
    function deserializeHeader(bytes memory _headerBs) internal pure returns (Header memory) {
        Header memory header;
        uint256 off = 0;
        (header.version, off)  = ZeroCopySource.NextUint32(_headerBs, off);

        (header.chainId, off) = ZeroCopySource.NextUint64(_headerBs, off);

        (header.prevBlockHash, off) = ZeroCopySource.NextHash(_headerBs, off);

        (header.transactionsRoot, off) = ZeroCopySource.NextHash(_headerBs, off);

        (header.crossStatesRoot, off) = ZeroCopySource.NextHash(_headerBs, off);

        (header.blockRoot, off) = ZeroCopySource.NextHash(_headerBs, off);

        (header.timestamp, off) = ZeroCopySource.NextUint32(_headerBs, off);

        (header.height, off) = ZeroCopySource.NextUint32(_headerBs, off);

        (header.consensusData, off) = ZeroCopySource.NextUint64(_headerBs, off);

        (header.consensusPayload, off) = ZeroCopySource.NextVarBytes(_headerBs, off);

        (header.nextBookkeeper, off) = ZeroCopySource.NextBytes20(_headerBs, off);

        return header;
    }
}

contract x {
    
    function test(bytes memory raw) public view returns(bool success, bytes32 res) {
        assembly {
            res := mload(0x40)
            mstore(0x40,add(res,0x20))
            success := staticcall(gas(),0x3,add(raw,0x20),mload(raw),res,0x20)
            res := mload(res)
        }
    }
}

