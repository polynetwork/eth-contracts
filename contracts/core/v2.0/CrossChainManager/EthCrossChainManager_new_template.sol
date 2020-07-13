pragma solidity ^0.5.0;

import "./../../../libs/math/SafeMath.sol";
import "./../../../libs/common/ZeroCopySource.sol";
import "./../../../libs/common/ZeroCopySink.sol";
import "./../../../libs/utils/Utils.sol";
import "./UpgradableECCM.sol";
import "./libs/EthCrossChainUtils.sol";
import "./interface/IEthCrossChainManager.sol";
import "./interface/IEthCrossChainData.sol";
contract NewEthCrossChainManager is IEthCrossChainManager, UpgradableECCM {
    using SafeMath for uint256;

    event InitGenesisBlockEvent(uint256  height, bytes rawHeader);
    event ChangeBookKeeperEvent(uint256 height, bytes rawHeader);
    event SyncBlockHeaderEvent(uint256 height, bytes rawHeader);
    event CrossChainEvent(address indexed sender, bytes txId, address proxyOrAssetContract, uint64 toChainId, bytes toContract, bytes rawdata);
    event VerifyAndExecuteTxEvent(uint64 fromChainID, bytes toContract, bytes crossChainTxHash, bytes fromChainTxHash);
    event AddFunctionTest1(bytes rawHeader);
    event AddFunctionTest2(uint256 height, bytes toMerkleValueBs);
    constructor(address _eccd) UpgradableECCM(_eccd) public {}
    
    /* @notice      sync multi-chain genesis block header to smart contrat
    *  @dev         this function can only be called once, the rawheader
    *               nextbookkeeper can't be empty
    *  @param _b    multichain genesis block raw header
    *  @return      none
    */
    function InitGenesisBlock(bytes memory _rawHeader, bytes memory _pubKeyList) whenNotPaused public returns(bool) {
        
        IEthCrossChainData eccd = IEthCrossChainData(EthCrossChainDataAddress);
        require(!eccd.isGenesisBlockInited(), "EthCrossChainData contract has already been initialized!");
        ECCUtils.Header memory header = ECCUtils.deserializeHeader(_rawHeader);
        
        (bytes20 nextBookKeeper, address[] memory keepers) = ECCUtils.verifyPubkey(_pubKeyList);
        require(header.nextBookkeeper == nextBookKeeper, "NextBookers illegal");
        
        require(eccd.putMCHeaderBytes(header.height, _rawHeader), "Save MCHeader bytes to Data contract failed!");
        require(eccd.putLatestHeight(header.height), "Save MC LatestHeight to Data contract failed!");
        require(eccd.putLatestBookKeeperHeight(header.height), "Save MC LatestBookKeeperHeight to Data contract failed!");
        require(eccd.initGenesisBlock(), "Save initGenesisBlock to Data contract failed!");

        //governance
        require(eccd.putMCKeeperHeight(header.height), "Save MCKeeperHeight to Data contract failed!");
        require(eccd.putMCKeeperPubBytes(header.height, ECCUtils.serializeKeepers(keepers)), "Save MCKeeperBytes to Data contract failed!");
        
        emit InitGenesisBlockEvent(header.height, _rawHeader);
        return true;
    }
    
    /* @notice               change multi-chain consensus book keeper
    *  @param _rawHeader     multichain change book keeper block raw header
    *  @param _pubKeyList    multichain consensus nodes public key list
    *  @param _sigList       multichain consensus nodes signature list
    *  @return true or false
    */
    function ChangeBookKeeper(bytes memory _rawHeader, bytes memory _pubKeyList, bytes memory _sigList) whenNotPaused public returns(bool) {
        ECCUtils.Header memory header = ECCUtils.deserializeHeader(_rawHeader);
        //prevent sync duplicate blocks
        IEthCrossChainData eccd = IEthCrossChainData(EthCrossChainDataAddress);
        uint64 _latestHeight = eccd.getLatestHeight();
        require(header.height > _latestHeight, "The height of header illegal");
        require(header.nextBookkeeper != bytes20(0), "The nextBookKeeper of header is empty");
        bytes memory mcKeeperPubBytes = eccd.getMCKeeperPubBytes(eccd.getLatestBookKeeperHeight());
        
        address[] memory McPubKeys = ECCUtils.deserializeKeepers(mcKeeperPubBytes);
        uint n = McPubKeys.length;

        require(ECCUtils.verifySig(_rawHeader, _sigList, McPubKeys, n - (2 * n - 1) / 3), "Verify signature failed!");

        (bytes20 nextBookKeeper, address[] memory keepers) = ECCUtils.verifyPubkey(_pubKeyList);

        require(header.nextBookkeeper == nextBookKeeper, "NextBookers illegal");
        
        require(eccd.putLatestHeight(header.height), "Save MC LatestHeight to Data contract failed!");
        require(eccd.putMCHeaderBytes(header.height, _rawHeader), "Save MCHeader bytes to Data contract failed!");

        require(eccd.putLatestBookKeeperHeight(header.height), "Save MC LatestBookKeeperHeight to Data contract failed!");
        
        //governance
        require(eccd.putMCKeeperHeight(header.height), "Save MCKeeperHeight to Data contract failed!");
        require(eccd.putMCKeeperPubBytes(header.height, ECCUtils.serializeKeepers(keepers)), "Save MCKeeperBytes to Data contract failed!");
        
        emit ChangeBookKeeperEvent(header.height, _rawHeader);
        return true;
    }

    /* @notice               sync multi-chain none-genesis and none change book keeper block header to smart contract
    *  @dev                  the same block height can only be synchronized once
    *  @param _rawHeader     multichain block raw header
    *  @param _sigList       multichain consensus nodes signature list
    *  @return true or false
    */
    function SyncBlockHeader(bytes memory _rawHeader, bytes memory _sigList) whenNotPaused public returns(bool) {
        ECCUtils.Header memory header = ECCUtils.deserializeHeader(_rawHeader);
        //multi-chain none change book keeper block header, the next book keeper is null
        require(header.nextBookkeeper == bytes20(0), "Header.nextnextBookkeeper should be empty");
        require(header.height > 0, "block height should > 0");
        IEthCrossChainData eccd = IEthCrossChainData(EthCrossChainDataAddress);
        bytes memory headerBytes = eccd.getMCHeaderBytes(header.height);
        if (headerBytes.length > 0) {
            // header has been synced
            return true;  
        }
        //governance
        uint64 targetBlockHeight;
        uint64 _latestBookKeeperHeight= eccd.getLatestBookKeeperHeight();
        if (header.height >= _latestBookKeeperHeight){
            targetBlockHeight = _latestBookKeeperHeight;
        } else {
            uint64[] memory _mcKeeperHeight = eccd.getMCKeeperHeight();
            (uint64 index, bool b) = Utils.findBookKeeper(_mcKeeperHeight, uint64(_mcKeeperHeight.length), header.height);
            require(b, "can not find block height in book keeper height list");
            targetBlockHeight = _mcKeeperHeight[index]; 
        }
        address[] memory McPubKeys = ECCUtils.deserializeKeepers(eccd.getMCKeeperPubBytes(targetBlockHeight));

        uint n = McPubKeys.length;
        require(ECCUtils.verifySig(_rawHeader, _sigList, McPubKeys, n - (2 * n - 1) / 3), "Verify header signature failed!");
        require(eccd.putMCHeaderBytes(header.height, _rawHeader), "Save MCHeader bytes to Data contract failed!");
        uint64 _latestHeight = eccd.getLatestHeight();
        if(header.height > _latestHeight){
            require(eccd.putLatestHeight(header.height), "Save MC LatestHeight to Data contract failed!");
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
    function crossChain(uint64 _toChainId, bytes calldata _toContract, bytes calldata _method, bytes calldata _txData) whenNotPaused external returns (bool) {
        IEthCrossChainData eccd = IEthCrossChainData(EthCrossChainDataAddress);
        uint256 ethTxHashIndex = eccd.getEthTxHashIndex();
        bytes memory paramTxHash = Utils.uint256ToBytes(ethTxHashIndex);

        bytes memory rawParam = abi.encodePacked(ZeroCopySink.WriteVarBytes(paramTxHash),
            ZeroCopySink.WriteVarBytes(abi.encodePacked(sha256(abi.encodePacked(address(this), Utils.uint256ToBytes(ethTxHashIndex))))),
            ZeroCopySink.WriteVarBytes(Utils.addressToBytes(msg.sender)),
            ZeroCopySink.WriteUint64(_toChainId),
            ZeroCopySink.WriteVarBytes(_toContract),
            ZeroCopySink.WriteVarBytes(_method),
            ZeroCopySink.WriteVarBytes(_txData)
        );
        // must save it in the storage to be included in the proof to be verified.
        require(eccd.putEthTxHash(keccak256(rawParam)), "Save ethTxHash by index to Data contract failed!");

        emit CrossChainEvent(tx.origin, paramTxHash, msg.sender, _toChainId, _toContract,  rawParam);
        
        return true;
    }

    /* @notice              verify multi-chan transaction, if success, mint ERC20 token
    *  @param _proof        multi-chain merkle proof
    *  @param _position     left or right
    *  @param _toMerkleValueBs  multi-chain transaction value
    *  @param _blockHeight    multi-chain block height
    *  @return true or false
    */
    function verifyAndExecuteTx(bytes32[] memory _proof, ECCUtils.POS[] memory _position, bytes memory _toMerkleValueBs, uint64 _blockHeight) whenNotPaused public returns (bool){
        IEthCrossChainData eccd = IEthCrossChainData(EthCrossChainDataAddress);
        uint64 _latestHeight = eccd.getLatestHeight();
        
        require(_blockHeight <= _latestHeight, "blockHeight > LatestHeight!");
        
        bytes memory headerBytes = eccd.getMCHeaderBytes(_blockHeight);
        
        ECCUtils.Header memory header = ECCUtils.deserializeHeader(headerBytes);
        
        require(ECCUtils.verifyCrossChainMerkleProof(_proof, _position, header.crossStatesRoot, Utils.hashLeaf(_toMerkleValueBs)), "Verify crosschain merkle proof failed!");

        ECCUtils.ToMerkleValue memory toMerkleValue = ECCUtils.deserializMerkleValue(_toMerkleValueBs);

        require(!eccd.getCrossChainTxExist(Utils.bytesToBytes32(toMerkleValue.txHash)), "the transaction has been executed!");

        require(eccd.putCrossChainTxExist(Utils.bytesToBytes32(toMerkleValue.txHash)), "Save crosschain tx exist failed!");


        address toContract = Utils.bytesToAddress(toMerkleValue.makeTxParam.toContract);
        // Ethereum ChainId is 1, we need to check the transaction is for Ethereum network
        require(toMerkleValue.makeTxParam.toChainId == uint64(2), "This Tx is not aiming at Ethereeum network!");
        //TODO: check this part to make sure we commit the next line when doing local net UT test
        require(_executeCrossChainTx(toContract, toMerkleValue.makeTxParam.method, toMerkleValue.makeTxParam.args, toMerkleValue.makeTxParam.fromContract, toMerkleValue.fromChainID), "Execute CrossChain Tx failed!");

        emit VerifyAndExecuteTxEvent(toMerkleValue.fromChainID, toMerkleValue.makeTxParam.toContract, toMerkleValue.txHash, toMerkleValue.makeTxParam.txHash);

        return true;
    }

    /* @notice  combine SyncBlockHeader and verifyAndExecuteTx function,see above function comment
    */
    function SyncAndVerify(bytes memory rawHeader, bytes memory sigList, bytes32[] memory proof, ECCUtils.POS[] memory position, bytes memory toMerkleValueBs, uint64 blockHeight) whenNotPaused public returns(bool) {
        require(SyncBlockHeader(rawHeader, sigList), "SynSyncBlockHeader failed!");
        require(verifyAndExecuteTx(proof, position, toMerkleValueBs, blockHeight), "verifyAndExecuteTx failed!");
        return true;
    }
    function addFunctionTest1(bytes memory _rawHeader, bytes memory _sigList) whenNotPaused public returns(bool) {
        require(SyncBlockHeader(_rawHeader, _sigList), "addFunctionTest invoke 'SyncBlockHeader' failed!");
        emit AddFunctionTest1(_rawHeader);
    }
    function addFunctionTest2(bytes32[] memory _proof, ECCUtils.POS[] memory _position, bytes memory _toMerkleValueBs, uint64 _blockHeight) whenNotPaused public returns (bool) {
        require(verifyAndExecuteTx(_proof, _position, _toMerkleValueBs, _blockHeight), "addFunctionTest invoke 'verifyAndExecuteTx' failed!");
        emit AddFunctionTest2(_blockHeight, _toMerkleValueBs);
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
}