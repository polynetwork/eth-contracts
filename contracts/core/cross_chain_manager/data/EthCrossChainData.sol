pragma solidity ^0.5.0;
import "./../../../libs/ownership/Ownable.sol";
import "./../../../libs/lifecycle/Pausable.sol";
import "./../interface/IEthCrossChainData.sol";

contract EthCrossChainData is IEthCrossChainData, Ownable, Pausable{
    /*
     Ethereum cross chain tx hash indexed by the automatically increased index.
     This map exists for the reason that Poly chain can verify the existence of 
     cross chain request tx coming from Ethereum
    */
    mapping(uint256 => bytes32) public EthToPolyTxHashMap;
    // This index records the current Map length
    uint256 public EthToPolyTxHashIndex;

    /* 
     When Poly chain switches the consensus epoch book keepers, the consensus peers public keys of Poly chain should be 
     changed into no-compressed version so that solidity smart contract can convert it to address type and 
     verify the signature derived from Poly chain account signature.
     ConKeepersPkBytes means Consensus book Keepers Public Key Bytes
    */
    bytes public ConKeepersPkBytes;
    
    // CurEpochStartHeight means Current Epoch Start Height of Poly chain block
    uint32 public CurEpochStartHeight;
    
    // Record the from chain txs that have been processed
    mapping(uint64 => mapping(bytes32 => bool)) FromChainTxExist;
    
    // Extra map for the usage of future potentially
    mapping(bytes32 => mapping(bytes32 => bytes)) public ExtraData;
    
    // Store Current Epoch Start Height of Poly chain block
    function putCurEpochStartHeight(uint32 curEpochStartHeight) public whenNotPaused onlyOwner returns (bool) {
        CurEpochStartHeight = curEpochStartHeight;
        return true;
    }

    // Get Current Epoch Start Height of Poly chain block
    function getCurEpochStartHeight() public view returns (uint32) {
        return CurEpochStartHeight;
    }

    // Store Consensus book Keepers Public Key Bytes
    function putCurEpochConPubKeyBytes(bytes memory curEpochPkBytes) public whenNotPaused onlyOwner returns (bool) {
        ConKeepersPkBytes = curEpochPkBytes;
        return true;
    }

    // Get Consensus book Keepers Public Key Bytes
    function getCurEpochConPubKeyBytes() public view returns (bytes memory) {
        return ConKeepersPkBytes;
    }

    // Mark from chain tx fromChainTx as exist or processed
    function markFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) public whenNotPaused onlyOwner returns (bool) {
        FromChainTxExist[fromChainId][fromChainTx] = true;
        return true;
    }

    // Check if from chain tx fromChainTx has been processed before
    function checkIfFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) public view returns (bool) {
        return FromChainTxExist[fromChainId][fromChainTx];
    }

    // Get current recorded index of cross chain txs requesting from Ethereum to other public chains
    // in order to help cross chain manager contract differenciate two cross chain tx requests
    function getEthTxHashIndex() public view returns (uint256) {
        return EthToPolyTxHashIndex;
    }

    // Store Ethereum cross chain tx hash, increase the index record by 1
    function putEthTxHash(bytes32 ethTxHash) public whenNotPaused onlyOwner returns (bool) {
        EthToPolyTxHashMap[EthToPolyTxHashIndex] = ethTxHash;
        EthToPolyTxHashIndex = EthToPolyTxHashIndex + 1;
        return true;
    }

    // Get Ethereum cross chain tx hash indexed by ethTxHashIndex
    function getEthTxHash(uint256 ethTxHashIndex) public view returns (bytes32) {
        return EthToPolyTxHashMap[ethTxHashIndex];
    }

    // Store extra data, which may be used in the future
    function putExtraData(bytes32 key1, bytes32 key2, bytes memory value) public whenNotPaused onlyOwner returns (bool) {
        ExtraData[key1][key2] = value;
        return true;
    }
    // Get extra data, which may be used in the future
    function getExtraData(bytes32 key1, bytes32 key2) public view returns (bytes memory) {
        return ExtraData[key1][key2];
    }
    
    function pause() onlyOwner whenNotPaused public returns (bool) {
        _pause();
        return true;
    }
    
    function unpause() onlyOwner whenPaused public returns (bool) {
        _unpause();
        return true;
    }
}