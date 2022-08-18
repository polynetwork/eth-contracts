pragma solidity ^0.5.0;

import "./../../../../libs/math/SafeMath.sol";

library ECCUtils {
    using SafeMath for uint256;

    uint constant ZION_SEAL_LEN = 67; // rlpPrefix: 2 , r: 32 , s:32 , v:1
    uint constant ZION_PEER_LEN = 93; // rlpPrefix: 2 , pk_rlp: 70 , address_rlp: 21
    uint constant ZION_ADDRESS_LEN = 21; // rlpPrefix: 1 , address: 20

    struct Header {
        bytes    root;
        uint256  number;
    }

    struct CrossTx {
        bytes  txHash;    // Zion txhash
        uint64 fromChainID;
        TxParam crossTxParam;
    }

    struct TxParam {
        bytes txHash;     // source chain txhash
        bytes crossChainId;
        bytes fromContract;
        uint64 toChainId;
        bytes toContract;
        bytes method;
        bytes args;
    }
    
    // will change memory (rawSeals)
    function verifyHeader(bytes32 headerHash, bytes memory rawSeals, address[] memory validators) internal pure returns(bool) {
        uint offset = 0x20;
        bytes memory seal;

        (rawSeals,) = rlpSplit(rawSeals, 0x20); 
        uint sealCount = rawSeals.length.div(ZION_SEAL_LEN);
        address[] memory signers = new address[](sealCount);

        for (uint i=0; i<sealCount; i++) {
            (seal,offset) = rlpSplit(rawSeals, offset);
            signers[i] = verifySeal(keccak256(abi.encodePacked(headerHash)), seal);
            if (signers[i] == address(0)) {
                return false;
            }
        }
        return hasEnoughSigners(validators, signers);
    }
    
    function verifySeal(bytes32 sigHash, bytes memory seal) internal pure returns(address signer) {
        if (seal.length != 65) {
            return (address(0));
        }

        bytes32 r;
        bytes32 s;
        uint8 v;

        assembly {
            r := mload(add(seal, 0x20))
            s := mload(add(seal, 0x40))
            v := add(27, byte(0, mload(add(seal, 0x60))))
        }

        // uncomment codes below if needed
        // if (uint256(s) > 0x7FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF5D576E7357A4501DDFE92F46681B20A0) {
        //     return address(0);
        // }

        // if (v != 27 && v != 28) {
        //     return address(0);
        // }

        return ecrecover(sigHash, v, r, s);
    }

    function hasEnoughSigners(address[] memory _validators, address[] memory _signers) internal pure returns(bool) {
        uint _m = _validators.length.mul(2).div(3).add(1);
        bool[] memory _checked = new bool[](_validators.length);
        uint m = 0;
        for(uint i = 0; i < _signers.length; i++){
            for (uint j = 0; j < _validators.length; j++) {
                if (!_checked[j] && _signers[i]==_validators[j]) {
                    m++;
                    _checked[j] = true;
                    break;
                }
            }
        }
        return m >= _m;
    }
    
    // []byte("request") = 72657175657374
    function getCrossTxStorageSlot(CrossTx memory ctx) internal pure returns(bytes memory slotIndex) {
        return bytes32ToBytes(keccak256(abi.encodePacked(bytes7(0x72657175657374), getUint64Bytes(ctx.crossTxParam.toChainId), ctx.txHash)));
    }
    
    // little endian
    function getUint64Bytes(uint64 num) internal pure returns(bytes8 res) {
        assembly {
            res := shl(56, and(num, 0xff))
            res := add(res, shl(40, and(num, 0xff00)))
            res := add(res, shl(24, and(num, 0xff0000)))
            res := add(res, shl(8,  and(num, 0xff000000)))
            res := add(res, shr(8,  and(num, 0xff00000000)))
            res := add(res, shr(24, and(num, 0xff0000000000)))
            res := add(res, shr(40, and(num, 0xff000000000000)))
            res := add(res, shr(56, num))
            res := shl(192, res)
        }
    }
    
    function bytes32ToBytes(bytes32 raw) internal pure returns(bytes memory res) {
        assembly {
            res := mload(0x40)
            mstore(0x40,add(res,0x40))
            mstore(res,0x20)
            mstore(add(res,0x20),raw)
        }
    }

    function uint256ToBytes(uint256 _value) internal pure returns (bytes memory bs) {
        require(_value <= 0x7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff, "Value exceeds the range");
        assembly {
            bs := mload(0x40)
            mstore(bs, 0x20)
            mstore(add(bs, 0x20), _value)
            mstore(0x40, add(bs, 0x40))
        }
    }
    
    function addressToBytes(address _addr) internal pure returns (bytes memory bs){
        assembly {
            bs := mload(0x40)
            mstore(bs, 0x14)
            mstore(add(bs, 0x20), shl(96, _addr))
            mstore(0x40, add(bs, 0x40))
       }
    }

    function bytesToBytes32(bytes memory _bs) internal pure returns (bytes32 value) {
        require(_bs.length == 32, "bytes length is not 32.");
        assembly {
            value := mload(add(_bs, 0x20))
        }
    }

    function bytesToUint256(bytes memory _bs) internal pure returns (uint256 value) {
        require(_bs.length == 32, "bytes length is not 32.");
        assembly {
            value := mload(add(_bs, 0x20))
        }
        require(value <= 0x7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff, "Value exceeds the range");
    }

    function bytesToAddress(bytes memory _bs) internal pure returns (address addr) {
        require(_bs.length == 20, "bytes length does not match address");
        assembly {
            addr := mload(add(_bs, 0x14))
        }
    }

    function getHeaderValidatorsAndEpochEndHeight(bytes memory rawHeader) internal view returns(uint64 endHeight,address[] memory validators) {
        uint size;
        
        (,uint offset) = rlpReadKind(rawHeader,0x20);
        (size,offset) = rlpReadKind(rawHeader, offset + 445 ); // position of Difficulty
        (size,offset) = rlpReadKind(rawHeader, offset + size); // position of Number
        (size,offset) = rlpReadKind(rawHeader, offset + size); // position of GasLimit
        (size,offset) = rlpReadKind(rawHeader, offset + size); // position of GasUsed
        (size,offset) = rlpReadKind(rawHeader, offset + size); // position of Time
        (size,offset) = rlpReadKind(rawHeader, offset + size); // position of Extra(with digest)
        (size,offset) = rlpReadKind(rawHeader, offset + 0x20); // position of Extra(without digest) , a bytes32 digest is appended before extra
        (size,offset) = rlpReadKind(rawHeader, offset);        // position of Extra.EpochStartHeight
        (endHeight,offset) = rlpGetNextUint64(rawHeader, offset); // position of Extra.EpochEndHeight
        (bytes memory validatorsBytes,) = rlpGetNextBytes(rawHeader, offset + size);
        size = validatorsBytes.length;
        require(size%ZION_ADDRESS_LEN==0,"invalid header extra validatorSet");
        validators = new address[](size/ZION_ADDRESS_LEN);
        for (uint i = 0; i*ZION_ADDRESS_LEN<size; i++) {
            (address valAddr,) = rlpGetNextAddress(validatorsBytes, 0x20 + i*ZION_ADDRESS_LEN);
            validators[i] = valAddr;
        }
    }
    
    function decodeHeader(bytes memory rawHeader) internal view returns(Header memory header) {
        uint size;
        
        (,uint offset) = rlpReadKind(rawHeader,0x20);
        (header.root,) = rlpGetNextBytes(rawHeader, offset + 87); // position of Root
        (size,offset) = rlpReadKind(rawHeader, offset + 445); // position of Difficulty
        (header.number,) = rlpGetNextUint256(rawHeader, offset + size);  // position of Number
    }
    
    function decodeValidators(bytes memory validatorBytes) internal pure returns(address[] memory validators) {
        validators = abi.decode(validatorBytes,(address[]));
    }
    
    function encodeValidators(address[] memory validators) internal pure returns(bytes memory validatorBytes) {
        validatorBytes = abi.encode(validators);
    }
    
    function encodeTxParam(
        bytes memory txHash,
        bytes memory crossChainId,
        bytes memory fromContract,
        uint64 toChainId,
        bytes memory toContract,
        bytes memory method,
        bytes memory args
    ) internal pure returns(bytes memory rawParam) {
        rawParam = abi.encode(txHash,
                crossChainId,
                fromContract,
                toChainId,
                toContract,
                method,
                args
            );
    }
    
    function decodeTxParam(bytes memory rawParam) internal view returns(TxParam memory param) {
        uint offset = 0x20;
        (param.txHash,offset) = rlpGetNextBytes(rawParam, offset);
        (param.crossChainId,offset) = rlpGetNextBytes(rawParam, offset);
        (param.fromContract,offset) = rlpGetNextBytes(rawParam, offset);
        (param.toChainId,offset) = rlpGetNextUint64(rawParam, offset);
        (param.toContract,offset) = rlpGetNextBytes(rawParam, offset);
        (param.method,offset) = rlpGetNextBytes(rawParam, offset);
        (param.args,offset) = rlpGetNextBytes(rawParam, offset);
        // (param.txHash, 
        // param.crossChainId,
        // param.fromContract,
        // param.toChainId,
        // param.toContract,
        // param.method,
        // param.args) = abi.decode(rawParam,(bytes,bytes,bytes,uint64,bytes,bytes,bytes));
    }
    
    function decodeCrossTx(bytes memory rawTx) internal view returns(CrossTx memory crossTx) {
        uint offset = 0x20;
        (,offset) = rlpReadKind(rawTx, offset);
        bytes memory crossTxParam;
        (crossTx.txHash,offset) = rlpGetNextBytes(rawTx, offset);
        (crossTx.fromChainID,offset) = rlpGetNextUint64(rawTx, offset);
        (crossTxParam,offset) = rlpGetNextBytes(rawTx, offset);
        crossTx.crossTxParam = decodeTxParam(crossTxParam);
    }
    
    // won't change memory
    function rlpGetNextBytes(bytes memory raw, uint offset) internal view returns (bytes memory res, uint _offset){
        uint size;
        (size,offset) = rlpReadKind(raw, offset);
        _offset = size + offset;
        assembly {
            res := mload(0x40)
            let alloc := add(0x20, mul(0x20, div(add(size,0x1f), 0x20)))
            mstore(0x40, add(res,alloc))
            mstore(res,size)
            pop(staticcall(gas(),0x4,add(raw,offset), size, add(res,0x20), size))
        }
    }  
    
    // won't change memory
    function rlpGetNextBytes32(bytes memory raw, uint offset) internal pure returns (bytes32 res, uint _offset){
        uint size;
        (size,offset) = rlpReadKind(raw, offset);
        require(size<=0x20,"rlpGetNextBytes32: data longer than 32 bytes");
        _offset = size + offset;
        assembly {
            let pad := sub(0x20,size)
            res := shl(mul(pad,8), mload(add(raw, sub(offset,pad))))
        }
    } 
    
    // won't change memory
    function rlpGetNextAddress(bytes memory raw, uint offset) internal pure returns (address res, uint _offset){
        uint size;
        (size,offset) = rlpReadKind(raw, offset);
        require(size<=0x14,"rlpGetNextAddress: data longer than 20 bytes");
        _offset = size + offset;
        assembly {
            let pad := sub(0x20,size)
            res := shr(mul(pad,8), shl(mul(pad,8), mload(add(raw,sub(offset,pad)))))
        }
    } 
    
    // won't change memory
    function rlpGetNextUint64(bytes memory raw, uint offset) internal pure returns (uint64 res, uint _offset){
        uint size;
        (size,offset) = rlpReadKind(raw, offset);
        require(size<=8,"rlpGetNextUint64: data longer than 8 bytes");
        _offset = size + offset;
        assembly {
            let pad := sub(0x20,size)
            res := shr(mul(pad,8), shl(mul(pad,8), mload(add(raw,sub(offset,pad)))))
        }
    } 
    
    // won't change memory
    function rlpGetNextUint256(bytes memory raw, uint offset) internal pure returns (uint256 res, uint _offset){
        uint size;
        (size,offset) = rlpReadKind(raw, offset);
        require(size<=32,"rlpGetNextUint256: data longer than 32 bytes");
        _offset = size + offset;
        assembly {
            let pad := sub(0x20,size)
            res := shr(mul(pad,8), shl(mul(pad,8), mload(add(raw,sub(offset,pad)))))
        }
    } 
    
    /*
    Account Proof
     -key : keccak256(address[:])
     -val : rlp(state-obj)
     -root : header.Root
     
    Storage Proof
     -key : keccak256(slot-index) // 32bytes index
     -val : 32 bytes data
     -root : Account.storageHash
     
    rlp(state-obj) == rlp(Account)
     type Account struct {
	  Nonce    uint64
	  Balance  *big.Int
	  Root     bytes32   // also : storageHash
	  CodeHash []byte
    }
    */
    function verifyAccountProof(
        bytes memory _accountProof, 
        bytes memory _headerRoot,
        bytes memory _address,
        bytes memory _storageProof,
        bytes memory _storageIndex
    ) internal pure returns(bytes memory value) {
         // verify account proof
        bytes memory _accountKey = _address;
        assembly { 
            mstore(add(_accountKey,0x20), keccak256(add(_accountKey, 0x20), mload(_accountKey)))
            mstore(_accountKey, 0x20)
        }
        bytes memory account = verifyProof(_accountProof, _accountKey, _headerRoot);
        
        // decode state object
        (account,) = rlpSplit(account,0x20);
        (,uint offset) = rlpSplit(account,0x20); // nonce
        (,offset) = rlpSplit(account,offset);    // balance
        (bytes memory storageRoot,) = rlpSplit(account,offset);
        
        // verify storage proof
        bytes memory _storageKey = _storageIndex;
        assembly { 
            mstore(add(_storageKey,0x20), keccak256(add(_storageKey, 0x20), mload(_storageKey)))
            mstore(_storageKey, 0x20)
        }
        value = verifyProof(_storageProof, _storageKey, storageRoot);
        (value,) = rlpSplit(value, 0x20);
    }
    
    // no copy , parse proof in-place , so if _proof is needed after , back up it.
    // for proof like: rlp([rlp(node1),rlp(node2),...,rlp(nodeN)])
    function verifyProof(bytes memory _proof, bytes memory _key, bytes memory _root) internal pure returns(bytes memory value) {
        _key = bytesToHex(_key);
        (_proof,) = rlpSplit(_proof,0x20); // get *[rlp(node1),rlp(node2),...,rlp(nodeN)]*
        uint offset = 0x20;
        uint size = _proof.length;
        value = _root; // *value* = *_root* ( which equal *hash(rlp(node1))* )
        
        for (;offset<size+0x20;) {
            bytes memory node; // memory for *nodeI*
            require(checkNodeHash(_proof,offset,value),"proof:unequal node hash"); // compare *nodeI-1.value* and *hash(rlp(nodeI))*
            (node,offset) = rlpSplit(_proof,offset); // get *nodeI*; _proof now: *[rlp(nodeI+1),rlp(nodeI+2),...,rlp(nodeN)]*
            
            (uint size_tmp,uint offset_tmp) = rlpReadKind(node,0x20);
            (size_tmp,offset_tmp) = rlpReadKind(node,offset_tmp+size_tmp);
            offset_tmp += size_tmp; // length of first two node elements
            if (offset_tmp==node.length+0x20) { // shortNode = [rlp(key),rlp(value)]
                (bytes memory keyElement,uint _offset) = rlpSplit(node,0x20);
                bool isIn;
                bytes memory subKey = compactToHex(keyElement);
                (_key,isIn) = compareKey(_key, subKey);
                require(isIn,"proof:invalid key");
                (value,) = rlpSplit(node,_offset);
            } else { // fullNode = [hash(child[0]),hash(child[1]),...,hash(child[15]),emptyKeyValue]
                uint index;
                if (_key.length==0) {
                    index = 17;
                } else {
                    (_key,index) = takeOneByte(_key);
                }
                uint _offset = 0x20;
                for (uint i=0;i<17;i++) {
                    if (index==i) {
                        (value,) = rlpSplit(node,_offset);
                        break;
                    } else {
                        (size_tmp,offset_tmp) = rlpReadKind(node,_offset);
                        _offset = offset_tmp + size_tmp;
                    }
                }
            } 
        }
        
        require(_key.length==0,"proof:invalid key");
    }
    
    // won't change memory
    function bytesToHex(bytes memory keyBytes) internal pure returns(bytes memory keyHex) {
        uint len = keyBytes.length*2 + 1;
        assembly {
            keyHex := mload(0x40)
            mstore(0x40, add(add(keyHex,len),0x20))
            mstore(keyHex, len)
        }
        for (uint i=0;i<keyBytes.length;i++) {
            keyHex[i*2] = keyBytes[i] >> 4;
            keyHex[i*2+1] = keyBytes[i] & 0x0f;
        }
        keyHex[len-1] = 0x10;
    }
    
    // won't change memory
    function compactToHex(bytes memory keyCompact) internal pure returns(bytes memory keyHex) {
        bytes1 t = keyCompact[0] >> 4;
        uint compactLen = keyCompact.length;
        uint hexLen;
        uint hasTerm;
        uint isOdd;
        assembly {
            isOdd := mod(shr(0xf8,t),2)
            hasTerm := div(shr(0xf8,t),2)
            hexLen := add(add(sub(mul(compactLen,2), 2), hasTerm), isOdd)
            keyHex := mload(0x40)
            mstore(0x40, add(add(keyHex,hexLen),0x20))
            mstore(keyHex, hexLen)
        }
        if (isOdd==1) {
            keyHex[0] = keyCompact[0] & 0x0f;
        }
        if (hasTerm==1) {
            keyHex[hexLen-1] = 0x10;
        }
        for (uint i=0;i<compactLen-1;i++) {
            keyHex[2*i+isOdd] = keyCompact[i+1] >> 4;
            keyHex[2*i+1+isOdd] = keyCompact[i+1] & 0x0f;
        }
    }
    
    // won't change memory
    function hexToCompact(bytes memory keyHex) internal pure returns(bytes memory keyCompact) {
        if (keyHex.length == 0) { return hex"00"; }
        uint hexLen = keyHex.length;
        uint compactLen = hexLen/2 + 1;
        uint t = 0;
        uint isOdd = hexLen%2;
        if (keyHex[hexLen-1]==0x10) { // has terminator
            t = 2;
            compactLen = (hexLen + 1)/2;
            isOdd = 1 - isOdd;
        }
        t += isOdd;
        assembly {
            keyCompact := mload(0x40)
            mstore(0x40, add(add(keyCompact,compactLen),0x20))
            mstore(keyCompact, compactLen)
        }
        {
            bytes1 right;
            if (isOdd == 1) { right = keyHex[0]; }
            bytes1 first;
            assembly { first := add(shl(0xfc,t), right) }
            keyCompact[0] = first;
        }
        for (uint i=0;i<compactLen-1;i++) {
            bytes1 left = keyHex[2*i+isOdd];
            bytes1 right = keyHex[2*i+1+isOdd];
            bytes1 bt;
            assembly { bt := add(shl(4,left), right) }
            keyCompact[i+1] = bt;
        }
    }
    
    // won't change memory
    function checkNodeHash(bytes memory raw,uint offset,bytes memory hash) internal pure returns(bool flag) {
        (uint size,uint offset_) = rlpReadKind(raw,offset);
        assembly {
            let fullSize := add(size,sub(offset_,offset))
            switch lt(fullSize,0x20) // for value less than 32 bytes , hash(value) = value
            case 1 {
                let equalVal := iszero(shr(mul(8,sub(0x20,fullSize)), xor(mload(add(raw,offset)),mload(add(hash,0x20)))))
                flag := and(equalVal, eq(fullSize,mload(hash)))
            }
            case 0 {
                flag := eq(keccak256(add(raw,offset), fullSize), mload(add(hash,0x20)))
            }
        }
    }
    
    // ! will change memory !
    // _buf = i + buf
    function takeOneByte(bytes memory _buf) internal pure returns(bytes memory buf,uint i) {
        require(_buf.length!=0,"takeOneByte: empty input");
        assembly {
            i := and(mload(add(_buf,1)),0xff)
            mstore(add(_buf,1),sub(mload(_buf),1))
            buf := add(_buf,1)
        }
    }
    
    // ! will change memory !
    // element must less than 32 bytes , key = element + keySlice (if isIn)
    function compareKey(bytes memory key,bytes memory element) internal pure returns(bytes memory keySlice,bool isIn) {
        if (key.length<element.length) {
            return (key,false);
        }
        uint slotCnt = element.length/32;
        isIn = true;
        for (uint i=0;i<slotCnt;i++) {
            assembly {
                let offset := add(0x20, mul(i, 0x20))
                isIn := and(isIn, eq(mload(add(key, offset)), mload(add(element, offset))))
            }
        }
        assembly{
            let offset := add(0x20, mul(slotCnt, 0x20))
            isIn := and(isIn, iszero(and(xor(mload(add(key,offset)),mload(add(element,offset))),shl(mul(8,sub(offset,mload(element))),0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff))))
        }
        if (!isIn) {
            return (key,false);
        }
        assembly{
            mstore(add(key,mload(element)),sub(mload(key),mload(element)))
            keySlice := add(key,mload(element))
        }
    }
    
    // ! will change memory !
    function rlpSplit(bytes memory raw,uint offset) internal pure returns(bytes memory res, uint offset_) {
        uint size;
        (size,offset_) = rlpReadKind(raw,offset);
        assembly {
            mstore(add(raw,sub(offset_,0x20)),size)
            res := add(raw,sub(offset_,0x20))
            offset_ := add(offset_,size)
        }
    }
    
    // won't change memory
    function rlpReadKind(bytes memory raw, uint offset) internal pure returns(uint size ,uint offset_) {
        uint k;
        assembly {
            k := shr(0xf8,mload(add(raw,offset)))
        }
        if (k<0x80) {
            assembly {
                size := 1
                offset_ := offset
            }
            return (size,offset_);
        }
        if (k<0xb7) {
            assembly {
                size := sub(k,0x80)
                offset_ := add(offset,1)
            }
            return (size,offset_);
        }
        if (k<0xc0) {
            assembly {
                size := shr(mul(8,sub(0xd7,k)),mload(add(add(raw,offset),0x01)))
                offset_ := add(offset,sub(k,0xb6))
            }
            return (size,offset_);
        }
        if (k<0xf7) {
            assembly {
                size := sub(k,0xc0)                    
                offset_ := add(offset,1)
            }
            return (size,offset_);
        }
        {
            assembly {
                size := shr(mul(8,sub(0x117,k)),mload(add(add(raw,offset),0x01)))
                offset_ := add(offset,sub(k,0xf6))
            }
        }
            
    }

    // won't change memory
    // function rlpReadKind(bytes memory raw, uint offset) internal pure returns(uint size ,uint offset_) {
    //     assembly {
    //         let k := shr(0xf8,mload(add(raw,offset)))
    //         if lt(k,0x80) { // Byte
    //             size := 1
    //             offset_ := offset
    //         }
    //         if and(lt(k,0xb7), gt(k,0x7f)) { // StringShort
    //             size := sub(k,0x80)
    //             offset_ := add(offset,1)
    //         }
    //         if and(lt(k,0xc0), gt(k,0xb6)) { // StringLong
    //             size := shr(mul(8,sub(0xd7,k)),mload(add(add(raw,offset),0x01)))
    //             offset_ := add(offset,sub(k,0xb6))
    //         }
    //         if and(lt(k,0xf7), gt(k,0xbf)) { // ListShort
    //             size := sub(k,0xc0)                    
    //             offset_ := add(offset,1)
    //         }
    //         if gt(k,0xf6) { // ListLong
    //             size := shr(mul(8,sub(0x117,k)),mload(add(add(raw,offset),0x01)))
    //             offset_ := add(offset,sub(k,0xf6))
    //         }
    //     }
    // }
}