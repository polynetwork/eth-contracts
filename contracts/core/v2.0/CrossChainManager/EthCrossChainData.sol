pragma solidity ^0.5.0;
import "./../../../libs/ownership/Ownable.sol";
import "./../../../libs/lifecycle/Pausable.sol";
import "./interface/IEthCrossChainData.sol";

contract EthCrossChainData is IEthCrossChainData, Ownable, Pausable{
    mapping(uint256 => bytes32) public EthTxHash;
    
    uint256 public EthTxHashIndex;

    // Multi-chain block header
    mapping(uint64 => bytes) public MCHeaderBytes;
    
    // multi-chain book keeper node switches, put the
    // book keeper public key to the map
    // key is block height, value is public keys
    mapping(uint64 => bytes) public MCKeeperPubBytes;

    mapping(bytes32 => bool) CrossChainTxExist;

    // multi-chain book keeper node switches, put the
    // block height to array   
    uint64[] public MCKeeperHeight;

    // the latest multi-chain book keeper switch block height
    uint64 public LatestBookKeeperHeight;

    // Latest synchronized multi-chain block header
    uint64 public LatestHeight;

    bool IsInitGenesisBlock;

    mapping(bytes32 => mapping(bytes32 => bytes)) public ExtraData;
    
    constructor() Ownable() Pausable() public {}

    function putMCHeaderBytes(uint64 _height, bytes memory _rawHeader) public whenNotPaused onlyOwner returns (bool) {
        MCHeaderBytes[_height] = _rawHeader;
        return true;
    }
    function getMCHeaderBytes(uint64 _height) public view returns (bytes memory) {
        return MCHeaderBytes[_height];
    }
    function putMCKeeperPubKeybytes(uint64 _height, bytes memory _keepersBytes) public whenNotPaused onlyOwner returns (bool) {
        MCKeeperPubBytes[_height] = _keepersBytes;
        return true;
    }
    function getMCKeeperPubKeybytes(uint64 _height) public view returns (bytes memory) {
        return MCKeeperPubBytes[_height];
    }
    function getEthTxHashIndex() public view returns (uint256) {
        return EthTxHashIndex;
    }
    function putEthTxHash(bytes32 _ethTxHash) public whenNotPaused onlyOwner returns (bool) {
        EthTxHash[EthTxHashIndex] = _ethTxHash;
        EthTxHashIndex = EthTxHashIndex + 1;
        return true;
    }
    function getEthTxHash(uint256 _ethTxHashIndex) public view returns (bytes32) {
        return EthTxHash[_ethTxHashIndex];
    }

    function putCrossChainTxExist(bytes32 _crossChainTx) public whenNotPaused onlyOwner returns (bool) {
        CrossChainTxExist[_crossChainTx] = true;
        return true;
    }
    function getCrossChainTxExist(bytes32 _crossChainTx) public view returns (bool) {
        return CrossChainTxExist[_crossChainTx];
    }
    function putLatestHeight(uint64 _latestHeight) public whenNotPaused onlyOwner returns (bool) {
        LatestHeight = _latestHeight;
        return true;
    }
    function getLatestHeight() public view returns (uint64) {
        return LatestHeight;
    }
    function putLatestBookKeeperHeight(uint64 _latestBookKeeperHeight) public whenNotPaused onlyOwner returns (bool)  {
        LatestBookKeeperHeight = _latestBookKeeperHeight;
        return true;
    }
    function getLatestBookKeeperHeight() public view returns (uint64) {
        return LatestBookKeeperHeight;
    }

    function putMCKeeperHeight(uint64 height) public whenNotPaused onlyOwner returns (bool) {
        MCKeeperHeight.push(height);
        return true;
    }
    
    function getMCKeeperHeight() public view returns (uint64[] memory) {
        return MCKeeperHeight;
    }
    
    function putMCKeeperPubBytes(uint64 height, bytes memory _MCKeeperPubBytes) public whenNotPaused onlyOwner returns (bool) {
        MCKeeperPubBytes[height] = _MCKeeperPubBytes;
        return true;
    }
    function getMCKeeperPubBytes(uint64 height) public view returns (bytes memory) {
        return MCKeeperPubBytes[height];
    }

    function initGenesisBlock() public onlyOwner returns (bool)  {
        IsInitGenesisBlock = true;
        return true;
    }
    function isGenesisBlockInited() public view returns (bool) {
        return IsInitGenesisBlock;
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