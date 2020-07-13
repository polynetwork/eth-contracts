pragma solidity ^0.5.0;
import "./../../../../libs/common/ZeroCopySource.sol";
import "./../../../../libs/common/ZeroCopySink.sol";
import "./../../../../libs/utils/Utils.sol";
library ECCUtils {
    struct Header {
        uint32 version;
        uint64 chainId;
        bytes32 prevBlockHash;
        bytes32 transactionsRoot;
        bytes32 crossStatesRoot;
        bytes32 blockRoot;
        uint32 timestamp;
        uint32 height;
        uint64 consensusData;
        bytes consensusPayload;
        bytes20 nextBookkeeper;
    }

    struct ToMerkleValue {
        bytes  txHash;  // cross chain txhash
        uint64 fromChainID;
        TxParam makeTxParam;
    }

    struct TxParam {
        bytes txHash; //  source chain txhash
        bytes crossChainId;
        bytes fromContract;
        uint64 toChainId;
        bytes toContract;
        bytes method;
        bytes args;
    }

    enum POS {
        LEFT,
        RIGHT
    }
    uint constant MCCHAIN_PUBKEY_LEN = 67;
    uint constant MCCHAIN_SIGNATURE_LEN = 65;
    /* @notice                  verify multi-chain transaction whether exist or not
    *  @param _proof            multi-chain merkle proof
    *  @param _position         left or right
    *  @param _root             multi-chain block root
    *  @param _toMerkleValue    multi-chain transaction value
    *  @return true or false
    */
    function verifyCrossChainMerkleProof(bytes32[] memory _proof, POS[] memory _position, bytes32 _root, bytes32 toMerkleValueHash) internal pure returns (bool) {
        bytes32 hash = toMerkleValueHash;

        bytes32 v;
        for (uint i = 0; i < _proof.length; i++){
            v = _proof[i];
            if (_position[i] == POS.LEFT) {
                hash = Utils.hashChildren(v, hash);
            }else if (_position[i] == POS.RIGHT) {
                hash = Utils.hashChildren(hash, v);
            }
        }

        return hash == _root;
    }

    /* @notice              calculate next book keeper according to public key list
    *  @param _keyLen       consensus node number
    *  @param _m            minimum signature number
    *  @param _pubKeyList   consensus node public key list
    *  @return two element: next book keeper, consensus node signer address
    */
    function _getBookKeeper(uint _keyLen, uint _m, bytes memory _pubKeyList) internal pure returns (bytes20, address[] memory){
         bytes memory buff ;
         buff = abi.encodePacked(buff, ZeroCopySink.WriteUint16(uint16(_keyLen)));

         address[] memory keeprs = new address[](_keyLen);

         for(uint i = 0; i < _keyLen; i++){
             buff =  abi.encodePacked(buff, ZeroCopySink.WriteVarBytes(Utils.compressMCPubKey(Utils.slice(_pubKeyList, i*MCCHAIN_PUBKEY_LEN, MCCHAIN_PUBKEY_LEN))));
             bytes32 hash = keccak256(Utils.slice(Utils.slice(_pubKeyList, i*MCCHAIN_PUBKEY_LEN, MCCHAIN_PUBKEY_LEN), 3, 64));
             keeprs[i] = address(uint160(uint256(hash)));
         }

         buff = abi.encodePacked(buff, ZeroCopySink.WriteUint16(uint16(_m)));
         bytes20  nextBookKeeper = ripemd160(abi.encodePacked(sha256(buff)));
         return (nextBookKeeper, keeprs);
    }

    /* @notice              verify multi-chain public key
    *  @param _pubKeyList   consensus node public key list
    *  @param _sigList      consensus node signature list
    *  @return  return three element: next book keeper, consensus node signer address, minimum signature number
    */
    function verifyPubkey(bytes memory _pubKeyList) internal pure returns (bytes20, address[] memory) {
        require(_pubKeyList.length % MCCHAIN_PUBKEY_LEN == 0, "_pubKeyList length illegal!");
        uint n = _pubKeyList.length / MCCHAIN_PUBKEY_LEN;
        return _getBookKeeper(n, n - (n - 1) / 3, _pubKeyList);
    }
    /* @notice              verify multi-chain consensus node signature
    *  @param _rawHeader    multi-chain block header raw bytes
    *  @param _sigList      consensus node signature list
    *  @param _m            minimum signature number
    *  @return true or false
    */
    function verifySig(bytes memory _rawHeader, bytes memory _sigList, address[] memory _keepers, uint _m) internal pure returns (bool){
        bytes32 hash = sha256(abi.encodePacked(sha256(_rawHeader)));

        uint signed = 0;
        uint sigCount = _sigList.length / MCCHAIN_SIGNATURE_LEN;
        for(uint j = 0; j  < sigCount; j++){
            bytes32 r = Utils.bytesToBytes32(Utils.slice(_sigList, j*MCCHAIN_SIGNATURE_LEN, 32));
            bytes32 s =  Utils.bytesToBytes32(Utils.slice(_sigList, j*MCCHAIN_SIGNATURE_LEN + 32, 32));
            uint8 v =  uint8(_sigList[j*MCCHAIN_SIGNATURE_LEN + 64]) + 27;

            // address signer =  ecrecover(keccak256(abi.encodePacked(hash)), v, r, s);
            address signer =  ecrecover(sha256(abi.encodePacked(hash)), v, r, s);
            if (Utils.containsAddress(_keepers, signer)){
                signed += 1;
            }
        }
        return signed >= _m;
    }
    function serializeKeepers(address[] memory keepers) internal pure returns (bytes memory) {
        uint256 keeperLen = keepers.length;
        bytes memory keepersBytes = ZeroCopySink.WriteUint64(uint64(keeperLen));
        for(uint i = 0; i < keeperLen; i++) {
            keepersBytes = abi.encodePacked(keepersBytes, ZeroCopySink.WriteVarBytes(Utils.addressToBytes(keepers[i])));
        }
        return keepersBytes;
    }
    function deserializeKeepers(bytes memory keepersBytes) internal pure returns (address[] memory) {
        uint256 off = 0;
        uint64 keeperLen;
        (keeperLen, off) = ZeroCopySource.NextUint64(keepersBytes, off);
        address[] memory keepers = new address[](keeperLen);
        bytes memory keeperBytes;
        for(uint i = 0; i < keeperLen; i++) {
            (keeperBytes, off) = ZeroCopySource.NextVarBytes(keepersBytes, off);
            keepers[i] = Utils.bytesToAddress(keeperBytes);
        }
        return keepers;
    }
    /* @notice               deserialize multi-chain transaction raw value
    *  @param _valueBs       multi-chain transaction raw bytes
    *  @return ToMerkleValue
    */
    function deserializMerkleValue(bytes memory _valueBs) internal pure returns (ToMerkleValue memory) {
        ToMerkleValue memory toMerkleValue;
        uint256 off = 0;

        (toMerkleValue.txHash, off) = ZeroCopySource.NextVarBytes(_valueBs, off);

        (toMerkleValue.fromChainID, off) = ZeroCopySource.NextUint64(_valueBs, off);

        TxParam memory txParam;

        (txParam.txHash, off) = ZeroCopySource.NextVarBytes(_valueBs, off);
        
        (txParam.crossChainId, off) = ZeroCopySource.NextVarBytes(_valueBs, off);

        (txParam.fromContract, off) = ZeroCopySource.NextVarBytes(_valueBs, off);

        (txParam.toChainId, off) = ZeroCopySource.NextUint64(_valueBs, off);

        (txParam.toContract, off) = ZeroCopySource.NextVarBytes(_valueBs, off);

        (txParam.method, off) = ZeroCopySource.NextVarBytes(_valueBs, off);

        (txParam.args, off) = ZeroCopySource.NextVarBytes(_valueBs, off);
        toMerkleValue.makeTxParam = txParam;

        return toMerkleValue;
    }

    /* @notice            deserialize multi-chain block header raw bytes
    *  @param _valueBs    multi-chain block header raw bytes
    *  @return Header
    */
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

    function parseProofToAuditPath(bytes memory proof) internal pure returns ()
}