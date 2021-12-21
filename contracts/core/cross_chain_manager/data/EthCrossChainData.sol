pragma solidity ^0.5.0;

import "./../../../libs/ownership/Ownable.sol";
import "./../../../libs/lifecycle/Pausable.sol";
import "../interface/IEthCrossChainData.sol";

contract EthCrossChainData is IEthCrossChainData, Ownable, Pausable{

    mapping(uint256 => bytes32) public EthToPolyTxHashMap;
  
    uint256 public EthToPolyTxHashIndex;

    bytes public CurValidatorPkBytes;
    
    uint64 public CurEpochStartHeight;

    uint64 public CurEpochId;
    
    mapping(uint64 => mapping(bytes32 => bool)) FromChainTxExist;
    
    mapping(bytes32 => mapping(bytes32 => bytes)) public ExtraData;
    
    function putCurEpochStartHeight(uint64 startHeight) public whenNotPaused onlyOwner returns (bool) {
        CurEpochStartHeight = startHeight;
        return true;
    }

    function getCurEpochStartHeight() public view returns (uint64) {
        return CurEpochStartHeight;
    }
    
    function putCurEpochId(uint64 epochId) public whenNotPaused onlyOwner returns (bool) {
        CurEpochId = epochId;
        return true;
    }

    function getCurEpochId() public view returns (uint64) {
        return CurEpochId;
    }

    function putCurEpochValidatorPkBytes(bytes memory curEpochPkBytes) public whenNotPaused onlyOwner returns (bool) {
        CurValidatorPkBytes = curEpochPkBytes;
        return true;
    }

    function getCurEpochValidatorPkBytes() public view returns (bytes memory) {
        return CurValidatorPkBytes;
    }

    function markFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) public whenNotPaused onlyOwner returns (bool) {
        FromChainTxExist[fromChainId][fromChainTx] = true;
        return true;
    }

    function checkIfFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) public view returns (bool) {
        return FromChainTxExist[fromChainId][fromChainTx];
    }

    function getEthTxHashIndex() public view returns (uint256) {
        return EthToPolyTxHashIndex;
    }

    function putEthTxHash(bytes32 ethTxHash) public whenNotPaused onlyOwner returns (bool) {
        EthToPolyTxHashMap[EthToPolyTxHashIndex] = ethTxHash;
        EthToPolyTxHashIndex = EthToPolyTxHashIndex + 1;
        return true;
    }

    function getEthTxHash(uint256 ethTxHashIndex) public view returns (bytes32) {
        return EthToPolyTxHashMap[ethTxHashIndex];
    }

    function putExtraData(bytes32 key1, bytes32 key2, bytes memory value) public whenNotPaused onlyOwner returns (bool) {
        ExtraData[key1][key2] = value;
        return true;
    }
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