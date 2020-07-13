pragma solidity ^0.5.0;

import "./../../../libs/math/SafeMath.sol";
import "./../../../libs/common/ZeroCopySource.sol";
import "./../../../libs/common/ZeroCopySink.sol";
import "./../../../libs/utils/Utils.sol";
import "./IEthCrossChainManager.sol";


contract EthCrossChainManager is IEthCrossChainManager{
    using SafeMath for uint256;
    enum POS {
        LEFT,
        RIGHT
    }

    mapping(uint256 => bytes32) public Transactions;

    mapping(bytes32 => bool) CrossChainTxExist;

    // Multi-chain block header
    mapping(uint64 => Header) public MCBlockHeaders;

    // Latest synchronized multi-chain block header
    uint64 public LatestHeight;

    // multi-chain book keeper node switches, put the
    // block height to array 
    uint64[] public MCKeeperHeight;

    // multi-chain book keeper node switches, put the
    // book keeper public key to the map
    // key is block height, value is public keys
    mapping(uint64 => address[]) public MCKeeperPubKeys;

    // the latest multi-chain book keeper switch block height
    uint64 public LatestBookKeeperHeight;


    bool IsInitGenesisBlock;

    uint constant MCCHAIN_PUBKEY_LEN = 67;
    uint constant MCCHAIN_SIGNATURE_LEN = 65;

    uint256 public TransactionId;


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
    event InitGenesisBlockEvent (uint256  height, bytes torawHeadersken);
    event ChangeBookKeeperEvent(uint256 height, bytes rawHeaders);
    event SyncBlockHeaderEvent(uint256 height, bytes rawHeaders);
    event CrossChainEvent (address indexed sender, bytes txId, address proxyOrAssetContract, uint64 toChainId, bytes toContract, bytes rawdata);
    event VerifyAndExecuteTxEvent(uint64 fromChainID, bytes toContract, bytes crossChainTxHash, bytes fromChainTxHash);
    
    /* @notice      sync multi-chain genesis block header to smart contrat
    *  @dev         this function can only be called once, the rawheader
    *               nextbookkeeper can't be empty
    *  @param _b    multichain genesis block raw header
    *  @return      none
    */
    function InitGenesisBlock(bytes memory _rawHeader, bytes memory _pubKeyList) public returns(bool) {
        require(!IsInitGenesisBlock, "Has already been initalized!");
        Header memory header = _deserializeHeader(_rawHeader);

        require(_pubKeyList.length % MCCHAIN_PUBKEY_LEN == 0, "Length of pubKeyList is illegal! ");
        uint n = _pubKeyList.length / MCCHAIN_PUBKEY_LEN;
        uint m = n - (n - 1) / 3;
        (bytes20 nextBookKeeper, address[] memory keepers) = _getBookKeeper(n, m, _pubKeyList);
        require(header.nextBookkeeper == nextBookKeeper, "NextBookers illegal");

        MCBlockHeaders[header.height] = header;
        LatestHeight = header.height;
        LatestBookKeeperHeight = header.height;
        IsInitGenesisBlock = true;

        //governance
        MCKeeperHeight.push(header.height);
        MCKeeperPubKeys[header.height] = keepers;

        emit SyncBlockHeaderEvent(LatestHeight, _rawHeader);
        return true;
    }

    /* @notice               change multi-chain consensus book keeper
    *  @param _rawHeader     multichain change book keeper block raw header
    *  @param _pubKeyList    multichain consensus nodes public key list
    *  @param _sigList       multichain consensus nodes signature list
    *  @return true or false
    */
    function ChangeBookKeeper(bytes memory _rawHeader, bytes memory _pubKeyList, bytes memory _sigList) public returns(bool) {
        Header memory header = _deserializeHeader(_rawHeader);
        //prevent sync duplicate blocks
        require(header.height > LatestHeight, "The height of header illegal");
        require(header.nextBookkeeper != bytes20(0), "The nextBookKeeper of header is empty");

        address [] memory McPubKeys = MCKeeperPubKeys[LatestBookKeeperHeight];
        uint n = McPubKeys.length;
        // uint m = n - (n - 1) / 3;
        uint m = n - (2 * n - 1) / 3;

        require(_verifySig(_rawHeader, _sigList, McPubKeys, m), "Verify signature failed!");

        (bytes20 nextBookKeeper, address[] memory keepers,) = _verifyPubkey(_pubKeyList);

        require(header.nextBookkeeper == nextBookKeeper, "NextBookers illegal");

        LatestHeight = header.height;
        MCBlockHeaders[header.height] = header;
        LatestBookKeeperHeight = header.height;

        //governance
        MCKeeperHeight.push(header.height);
        MCKeeperPubKeys[header.height] = keepers;

        emit ChangeBookKeeperEvent(LatestHeight, _rawHeader);
        return true;
    }

    /* @notice               sync multi-chain none-genesis and none change book keeper block header to smart contract
    *  @dev                  the same block height can only be synchronized once
    *  @param _rawHeader     multichain block raw header
    *  @param _sigList       multichain consensus nodes signature list
    *  @return true or false
    */
    function SyncBlockHeader(bytes memory _rawHeader, bytes memory _sigList) public returns(bool) {
        Header memory header = _deserializeHeader(_rawHeader);
        //multi-chain none change book keeper block header, the next book keeper is null
        require(header.nextBookkeeper == bytes20(0), "Header.nextnextBookkeeper should be empty");
        require(header.height > 0, "block height should > 0");
        Header memory tmp = MCBlockHeaders[header.height];
        // header has been synced
        if (tmp.timestamp != 0) return true;

        //governance
        address[] memory McPubKeys;
        if (header.height >= LatestBookKeeperHeight){
            McPubKeys = MCKeeperPubKeys[LatestBookKeeperHeight];
        } else {
            (uint64 index, bool b) = Utils.findBookKeeper(MCKeeperHeight, uint64(MCKeeperHeight.length), header.height);
            require(b, "can not find block height in book keeper height list");
            uint64 blockHeight = MCKeeperHeight[index];
            McPubKeys = MCKeeperPubKeys[blockHeight];
        }

        uint n = McPubKeys.length;
        // uint m = n - (n - 1) / 3;
        uint m = n - (2 * n - 1) / 3;

        require(_verifySig(_rawHeader, _sigList, McPubKeys, m), "Verify header signature failed!");

        MCBlockHeaders[header.height] = header;
        if(header.height > LatestHeight){
            LatestHeight = header.height;
        }

        emit SyncBlockHeaderEvent(header.height, _rawHeader);
        return true;
    }

    /* @notice              ERC20 token cross chain to other blockchain.
    *                       this function push tx event to blockchain
    *  @param _token        ERC20 token address
    *  @param _toChainId    target chain id
    *  @param _toContract   target smart contract address in target block chain
    *  @param _txData       transaction data for target chain, include to_address, amount
    *  @return unique transaction id
    */
    function crossChain(uint64 _toChainId, bytes calldata _toContract, bytes calldata _method, bytes calldata _txData) external returns (bool) {
        TxParam memory param;
        TransactionId = TransactionId + 1;
        param.txHash = Utils.uint256ToBytes(TransactionId);
        address _proxyOrAssetContract = msg.sender;
        
        param.crossChainId = abi.encodePacked(sha256(abi.encodePacked(address(this), Utils.uint256ToBytes(TransactionId))));
        
        param.fromContract = Utils.addressToBytes(_proxyOrAssetContract);

        param.toChainId = _toChainId;

        param.toContract = _toContract;

        param.method = _method;

        param.args = _txData;

        bytes memory rawParam = abi.encodePacked(ZeroCopySink.WriteVarBytes(param.txHash),
            ZeroCopySink.WriteVarBytes(param.crossChainId),
            ZeroCopySink.WriteVarBytes(param.fromContract),
            ZeroCopySink.WriteUint64(param.toChainId),
            ZeroCopySink.WriteVarBytes(param.toContract),
            ZeroCopySink.WriteVarBytes(param.method),
            ZeroCopySink.WriteVarBytes(param.args)
        );
        // must save it in the storage to be included in the proof to be verified.
        Transactions[TransactionId] = keccak256(rawParam);
        
        emit CrossChainEvent(tx.origin, param.txHash, _proxyOrAssetContract, _toChainId, _toContract,  rawParam);
        

        return true;
    }

    /* @notice              verify multi-chan transaction, if success, mint ERC20 token
    *  @param _proof        multi-chain merkle proof
    *  @param _position     left or right
    *  @param _toMerkleValueBs  multi-chain transaction value
    *  @param _blockHeight    multi-chain block height
    *  @return true or false
    */
    function verifyAndExecuteTx(bytes32[] memory _proof, POS[] memory _position, bytes memory _toMerkleValueBs, uint64 _blockHeight) public returns (bool){

        require(_blockHeight <= LatestHeight, "blockHeight > LatestHeight!");
        bytes32 root = MCBlockHeaders[_blockHeight].crossStatesRoot;

        require(_verifyCrossChainMerkleProof(_proof, _position, root, _toMerkleValueBs), "Verify crosschain merkle proof failed!");

        ToMerkleValue memory toMerkleValue = _deserializMerkleValue(_toMerkleValueBs);

        require(!CrossChainTxExist[Utils.bytesToBytes32(toMerkleValue.txHash)], "the transaction has been executed!");

        CrossChainTxExist[Utils.bytesToBytes32(toMerkleValue.txHash)] = true;


        address toContract = Utils.bytesToAddress(toMerkleValue.makeTxParam.toContract);
        // Ethereum ChainId is 1, we need to check the transaction is for Ethereum network
        require(toMerkleValue.makeTxParam.toChainId == uint64(2), "This Tx is not aiming at Ethereeum network!");

        require(_executeCrossChainTx(toContract, toMerkleValue.makeTxParam.method, toMerkleValue.makeTxParam.args, toMerkleValue.makeTxParam.fromContract, toMerkleValue.fromChainID), "Execute CrossChain Tx failed!");

        emit VerifyAndExecuteTxEvent(toMerkleValue.fromChainID, toMerkleValue.makeTxParam.toContract, toMerkleValue.txHash, toMerkleValue.makeTxParam.txHash);

        return true;
    }

    /* @notice  combine SyncBlockHeader and verifyAndExecuteTx function,see above function comment
    */
    function SyncAndVerify(bytes memory rawHeader, bytes memory sigList, bytes32[] memory proof, POS[] memory position, bytes memory toMerkleValueBs, uint64 blockHeight) public returns(bool) {
        require(SyncBlockHeader(rawHeader, sigList), "SynSyncBlockHeader failed!");
        require(verifyAndExecuteTx(proof, position, toMerkleValueBs, blockHeight), "verifyAndExecuteTx failed!");
        return true;
    }

    /* @notice              verify multi-chain public key
    *  @param _pubKeyList   consensus node public key list
    *  @param _sigList      consensus node signature list
    *  @return  return three element: next book keeper, consensus node signer address, minimum signature number
    */
    function _verifyPubkey(bytes memory _pubKeyList) internal pure returns (bytes20, address[] memory, uint) {
        require(_pubKeyList.length % MCCHAIN_PUBKEY_LEN == 0, "_pubKeyList length illegal!");
        uint n = _pubKeyList.length / MCCHAIN_PUBKEY_LEN;
        uint m = n - (n - 1) / 3;

        (bytes20 nextBookKeeper, address[] memory keepers) = _getBookKeeper(n, m, _pubKeyList);
        return (nextBookKeeper, keepers, m);
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

    /* @notice              verify multi-chain consensus node signature
    *  @param _rawHeader    multi-chain block header raw bytes
    *  @param _sigList      consensus node signature list
    *  @param _m            minimum signature number
    *  @return true or false
    */
    function _verifySig(bytes memory _rawHeader, bytes memory _sigList, address[] memory _keepers, uint _m) internal pure returns (bool){
        bytes32 hash = sha256(abi.encodePacked(sha256(_rawHeader)));

        uint signed = 0;
        for(uint j = 0; j  < _sigList.length / MCCHAIN_SIGNATURE_LEN; j++){

            bytes32 r = Utils.bytesToBytes32(Utils.slice(_sigList, j*MCCHAIN_SIGNATURE_LEN, 32));
            bytes32 s =  Utils.bytesToBytes32(Utils.slice(_sigList, j*MCCHAIN_SIGNATURE_LEN + 32, 32));
            uint8 v =  uint8(_sigList[j*MCCHAIN_SIGNATURE_LEN + 64]) + 27;

            address signer =  ecrecover(keccak256(abi.encodePacked(hash)), v, r, s);

            if (Utils.containsAddress(_keepers, signer)){
                signed += 1;
            }
        }
        return signed >= _m;
    }

    /* @notice                  verify multi-chain transaction whether exist or not
    *  @param _proof            multi-chain merkle proof
    *  @param _position         left or right
    *  @param _root             multi-chain block root
    *  @param _toMerkleValue    multi-chain transaction value
    *  @return true or false
    */
    function _verifyCrossChainMerkleProof(bytes32[] memory _proof, POS[] memory _position, bytes32 _root, bytes memory _toMerkleValue) internal pure returns (bool) {
        bytes32 hash = Utils.hashLeaf(_toMerkleValue);

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

    /* @notice                  mint ERC20 token transaction
    *  @param _token            ERC2O token address
    *  @param _args             cross chain parameters
    *  @param _fromContractAddr from chain smart contract address
    *  @return true or false
    */
    function _executeCrossChainTx(address _token, bytes memory method, bytes memory _args, bytes memory _fromContractAddr, uint64 fromChainId) internal returns (bool){
        require(Utils.isContract(_token), "The passed in address is not a contract!");
        bytes memory returnData;
        bool success;
        // The returnData will be bytes32, the last byte must be 01;
        (success, returnData) = _token.call(abi.encodePacked(bytes4(keccak256(abi.encodePacked(method, "(bytes,bytes,uint64)"))), abi.encode(_args, _fromContractAddr, fromChainId)));
        require(success == true, "EthCrossChain call business contract failed");
        require(returnData.length > 0, "No return value from business contract!");
        (bool res,) = ZeroCopySource.NextBool(returnData, 31);
        require(res == true, "EthCrossChain call business contract return is not true");
        return true;
    }

    /* @notice               deserialize multi-chain transaction raw value
    *  @param _valueBs       multi-chain transaction raw bytes
    *  @return ToMerkleValue
    */
    function _deserializMerkleValue(bytes memory _valueBs) internal pure returns (ToMerkleValue memory) {
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
    function _deserializeHeader(bytes memory _headerBs) internal pure returns (Header memory) {
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