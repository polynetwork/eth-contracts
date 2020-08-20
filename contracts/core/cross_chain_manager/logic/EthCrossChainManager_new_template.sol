pragma solidity ^0.5.0;

import "./../../../libs/math/SafeMath.sol";
import "./../../../libs/common/ZeroCopySource.sol";
import "./../../../libs/common/ZeroCopySink.sol";
import "./../../../libs/utils/Utils.sol";
import "./../upgrade/UpgradableECCM.sol";
import "./../libs/EthCrossChainUtils.sol";
import "./../interface/IEthCrossChainManager.sol";
import "./../interface/IEthCrossChainData.sol";
contract NewEthCrossChainManager is IEthCrossChainManager, UpgradableECCM {
    using SafeMath for uint256;

    event InitGenesisBlockEvent(uint256 height, bytes rawHeader);
    event ChangeBookKeeperEvent(uint256 height, bytes rawHeader);
    event CrossChainEvent(address indexed sender, bytes txId, address proxyOrAssetContract, uint64 toChainId, bytes toContract, bytes rawdata);
    event VerifyHeaderAndExecuteTxEvent(uint64 fromChainID, bytes toContract, bytes crossChainTxHash, bytes fromChainTxHash);
    constructor(address _eccd) UpgradableECCM(_eccd) public {}
    
    /* @notice              sync Poly chain genesis block header to smart contrat
    *  @dev                 this function can only be called once, nextbookkeeper of rawHeader can't be empty
    *  @param rawHeader     Poly chain genesis block raw header or raw Header including switching consensus peers info
    *  @return              true or false
    */
    function initGenesisBlock(bytes memory rawHeader, bytes memory pubKeyList) whenNotPaused public returns(bool) {
        // Load Ethereum cross chain data contract
        IEthCrossChainData eccd = IEthCrossChainData(EthCrossChainDataAddress);
        
        // Make sure the contract has not been initialized before
        require(eccd.getCurEpochConPubKeyBytes().length == 0, "EthCrossChainData contract has already been initialized!");
        
        // Parse header and convit the public keys into nextBookKeeper and compare it with header.nextBookKeeper to verify the validity of signature
        ECCUtils.Header memory header = ECCUtils.deserializeHeader(rawHeader);
        (bytes20 nextBookKeeper, address[] memory keepers) = ECCUtils.verifyPubkey(pubKeyList);
        require(header.nextBookkeeper == nextBookKeeper, "NextBookers illegal");
        
        // Record current epoch start height and public keys (by storing them in address format)
        require(eccd.putCurEpochStartHeight(header.height), "Save Poly chain current epoch start height to Data contract failed!");
        require(eccd.putCurEpochConPubKeyBytes(ECCUtils.serializeKeepers(keepers)), "Save Poly chain current epoch book keepers to Data contract failed!");
        
        // Fire the event
        emit InitGenesisBlockEvent(header.height, rawHeader);
        return true;
    }
    
    /* @notice              change Poly chain consensus book keeper
    *  @param rawHeader     Poly chain change book keeper block raw header
    *  @param pubKeyList    Poly chain consensus nodes public key list
    *  @param sigList       Poly chain consensus nodes signature list
    *  @return              true or false
    */
    function changeBookKeeper(bytes memory rawHeader, bytes memory pubKeyList, bytes memory sigList) whenNotPaused public returns(bool) {
        // Load Ethereum cross chain data contract
        ECCUtils.Header memory header = ECCUtils.deserializeHeader(rawHeader);
        IEthCrossChainData eccd = IEthCrossChainData(EthCrossChainDataAddress);
        
        // Make sure rawHeader.height is higher than recorded current epoch start height
        uint64 curEpochStartHeight = eccd.getCurEpochStartHeight();
        require(header.height > curEpochStartHeight, "The height of header is lower than current epoch start height!");
        
        // Ensure the rawHeader is the key header including info of switching consensus peers by containing non-empty nextBookKeeper field
        require(header.nextBookkeeper != bytes20(0), "The nextBookKeeper of header is empty");
        
        // Verify signature of rawHeader comes from pubKeyList
        address[] memory polyChainBKs = ECCUtils.deserializeKeepers(eccd.getCurEpochConPubKeyBytes());
        uint n = polyChainBKs.length;
        require(ECCUtils.verifySig(rawHeader, sigList, polyChainBKs, n - (n - 1) / 3), "Verify signature failed!");
        
        // Convert pubKeyList into ethereum address format and make sure the compound address from the converted ethereum addresses
        // equals passed in header.nextBooker
        (bytes20 nextBookKeeper, address[] memory keepers) = ECCUtils.verifyPubkey(pubKeyList);
        require(header.nextBookkeeper == nextBookKeeper, "NextBookers illegal");
        
        // update current epoch start height of Poly chain and current epoch consensus peers book keepers addresses
        require(eccd.putCurEpochStartHeight(header.height), "Save MC LatestHeight to Data contract failed!");
        require(eccd.putCurEpochConPubKeyBytes(ECCUtils.serializeKeepers(keepers)), "Save Poly chain book keepers bytes to Data contract failed!");
        
        // Fire the change book keeper event
        emit ChangeBookKeeperEvent(header.height, rawHeader);
        return true;
    }


    /* @notice              ERC20 token cross chain to other blockchain.
    *                       this function push tx event to blockchain
    *  @param toChainId     Target chain id
    *  @param toContract    Target smart contract address in target block chain
    *  @param txData        Transaction data for target chain, include to_address, amount
    *  @return              true or false
    */
    function crossChain(uint64 toChainId, bytes calldata toContract, bytes calldata method, bytes calldata txData) whenNotPaused external returns (bool) {
        // Load Ethereum cross chain data contract
        IEthCrossChainData eccd = IEthCrossChainData(EthCrossChainDataAddress);
        
        // To help differentiate two txs, the ethTxHashIndex is increasing automatically
        uint256 txHashIndex = eccd.getEthTxHashIndex();
        
        // Convert the uint256 into bytes
        bytes memory paramTxHash = Utils.uint256ToBytes(txHashIndex);
        
        // Construct the makeTxParam, and put the hash info storage, to help provide proof of tx existence
        bytes memory rawParam = abi.encodePacked(ZeroCopySink.WriteVarBytes(paramTxHash),
            ZeroCopySink.WriteVarBytes(abi.encodePacked(sha256(abi.encodePacked(address(this), paramTxHash)))),
            ZeroCopySink.WriteVarBytes(Utils.addressToBytes(msg.sender)),
            ZeroCopySink.WriteUint64(toChainId),
            ZeroCopySink.WriteVarBytes(toContract),
            ZeroCopySink.WriteVarBytes(method),
            ZeroCopySink.WriteVarBytes(txData)
        );
        
        // Must save it in the storage to be included in the proof to be verified.
        require(eccd.putEthTxHash(keccak256(rawParam)), "Save ethTxHash by index to Data contract failed!");
        
        // Fire the cross chain event denoting there is a cross chain request from Ethereum network to other public chains through Poly chain network
        emit CrossChainEvent(tx.origin, paramTxHash, msg.sender, toChainId, toContract, rawParam);
        return true;
    }

    /* @notice              Verify Poly chain header and proof, execute the cross chain tx from Poly chain to Ethereum
    *  @param proof         Poly chain tx merkle proof
    *  @param rawHeader     The header containing crossStateRoot to verify the above tx merkle proof
    *  @param headerProof   The header merkle proof used to verify rawHeader
    *  @param curRawHeader  Any header in current epoch consensus of Poly chain
    *  @param headerSig     The coverted signature veriable for solidity derived from Poly chain consensus nodes' signature
    *                       used to verify the validity of curRawHeader
    *  @return              true or false
    */
    function verifyHeaderAndExecuteTx(bytes memory proof, bytes memory rawHeader, bytes memory headerProof, bytes memory curRawHeader,bytes memory headerSig) whenNotPaused public returns (bool){
        ECCUtils.Header memory header = ECCUtils.deserializeHeader(rawHeader);
        // Load ehereum cross chain data contract
        IEthCrossChainData eccd = IEthCrossChainData(EthCrossChainDataAddress);
        
        // Get stored consensus public key bytes of current poly chain epoch and deserialize Poly chain consensus public key bytes to address[]
        address[] memory polyChainBKs = ECCUtils.deserializeKeepers(eccd.getCurEpochConPubKeyBytes());

        uint256 curEpochStartHeight = eccd.getCurEpochStartHeight();

        uint n = polyChainBKs.length;
        if (header.height >= curEpochStartHeight) {
            // It's enough to verify rawHeader signature
            require(ECCUtils.verifySig(rawHeader, headerSig, polyChainBKs, n - ( n - 1) / 3), "Verify poly chain header signature failed!");
        } else {
            // We need to verify the signature of curHeader 
            require(ECCUtils.verifySig(curRawHeader, headerSig, polyChainBKs, n - ( n - 1) / 3), "Verify poly chain header signature failed!");

            // Then use curHeader.StateRoot and headerProof to verify rawHeader.CrossStateRoot
            ECCUtils.Header memory curHeader = ECCUtils.deserializeHeader(curRawHeader);
            bytes memory proveValue = ECCUtils.merkleProve(headerProof, curHeader.blockRoot);
            require(ECCUtils.getHeaderHash(rawHeader) == Utils.bytesToBytes32(proveValue), "verify header proof failed!");
        }
        
        // Through rawHeader.CrossStateRoot, the toMerkleValue or cross chain msg can be verified and parsed from proof
        bytes memory toMerkleValueBs = ECCUtils.merkleProve(proof, header.crossStatesRoot);
        
        // Parse the toMerkleValue struct and make sure the tx has not been processed, then mark this tx as processed
        ECCUtils.ToMerkleValue memory toMerkleValue = ECCUtils.deserializeMerkleValue(toMerkleValueBs);
        require(!eccd.checkIfFromChainTxExist(toMerkleValue.fromChainID, Utils.bytesToBytes32(toMerkleValue.txHash)), "the transaction has been executed!");
        require(eccd.markFromChainTxExist(toMerkleValue.fromChainID, Utils.bytesToBytes32(toMerkleValue.txHash)), "Save crosschain tx exist failed!");
        
        // Ethereum ChainId is 2, we need to check the transaction is for Ethereum network
        require(toMerkleValue.makeTxParam.toChainId == uint64(2), "This Tx is not aiming at Ethereum network!");
        
        // Obtain the targeting contract, so that Ethereum cross chain manager contract can trigger the executation of cross chain tx on Ethereum side
        address toContract = Utils.bytesToAddress(toMerkleValue.makeTxParam.toContract);
        
        //TODO: check this part to make sure we commit the next line when doing local net UT test
        require(_executeCrossChainTx(toContract, toMerkleValue.makeTxParam.method, toMerkleValue.makeTxParam.args, toMerkleValue.makeTxParam.fromContract, toMerkleValue.fromChainID), "Execute CrossChain Tx failed!");

        // Fire the cross chain event denoting the executation of cross chain tx is successful,
        // and this tx is coming from other public chains to current Ethereum network
        emit VerifyHeaderAndExecuteTxEvent(toMerkleValue.fromChainID, toMerkleValue.makeTxParam.toContract, toMerkleValue.txHash, toMerkleValue.makeTxParam.txHash);

        return true;
    }


    /* @notice                  Dynamically invoke the targeting contract, and trigger executation of cross chain tx on Ethereum side
    *  @param _toContract       The targeting contract that will be invoked by the Ethereum Cross Chain Manager contract
    *  @param _method           At which method will be invoked within the targeting contract
    *  @param _args             The parameter that will be passed into the targeting contract
    *  @param _fromContractAddr From chain smart contract address
    *  @param _fromChainId      Indicate from which chain current cross chain tx comes 
    *  @return                  true or false
    */
    function _executeCrossChainTx(address _toContract, bytes memory _method, bytes memory _args, bytes memory _fromContractAddr, uint64 _fromChainId) internal returns (bool){
        // Ensure the targeting contract gonna be invoked is indeed a contract rather than a normal account address
        require(Utils.isContract(_toContract), "The passed in address is not a contract!");
        bytes memory returnData;
        bool success;
        
        // The returnData will be bytes32, the last byte must be 01;
        (success, returnData) = _toContract.call(abi.encodePacked(bytes4(keccak256(abi.encodePacked(_method, "(bytes,bytes,uint64)"))), abi.encode(_args, _fromContractAddr, _fromChainId)));
        
        // Ensure the executation is successful
        require(success == true, "EthCrossChain call business contract failed");
        
        // Ensure the returned value is true
        require(returnData.length != 0, "No return value from business contract!");
        (bool res,) = ZeroCopySource.NextBool(returnData, 31);
        require(res == true, "EthCrossChain call business contract return is not true");
        
        return true;
    }
    // new feature added in the upgraded contract
    function addFunctionTest1(uint256 a, uint256 b) whenNotPaused public view returns (uint256) {
        return a.add(b);
    }
}