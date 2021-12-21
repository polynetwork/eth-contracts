pragma solidity ^0.5.0;

import "./interface/IEthCrossChainManager.sol";
import "./interface/IEthCrossChainManagerProxy.sol";
import "./../../libs/utils/Utils.sol";

contract TunnelCCMCaller is IEthCrossChainManager, IEthCrossChainManagerProxy {
    
    address private _owner;
    address public realCCM;
    mapping(uint64 => bytes) public callerHashMap;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    function initialize(address newOwner) public {
        require(_owner == address(0),"already initialized");
        _owner = newOwner;
    }

    function owner() public view returns (address) {
        return _owner;
    }

    modifier onlyOwner() {
        require(isOwner(), "Ownable: caller is not the owner");
        _;
    }

    function isOwner() public view returns (bool) {
        return msg.sender == _owner;
    }

    function renounceOwnership() public onlyOwner {
        emit OwnershipTransferred(_owner, address(0));
        _owner = address(0);
    }

    function transferOwnership(address newOwner) public  onlyOwner {
        _transferOwnership(newOwner);
    }

    function _transferOwnership(address newOwner) internal {
        require(newOwner != address(0), "Ownable: new owner is the zero address");
        emit OwnershipTransferred(_owner, newOwner);
        _owner = newOwner;
    }

    function setRealCCM(address _ccm) onlyOwner public {
        realCCM = _ccm;
    }
    
    function bindCallerHash(uint64 toChainId, bytes memory targetCallerHash) onlyOwner public {
        callerHashMap[toChainId] = targetCallerHash;
    }

    
    // interface from poly1.0 EthCrossChainManagerProxy
    function getEthCrossChainManager() public view returns (address) {
        return address(this);
    }
    

    // interface from poly1.0 EthCrossChainmanager
    function crossChain(
        uint64 _toChainId, 
        bytes calldata _toContract, 
        bytes calldata _method, 
        bytes calldata _txData
    ) external returns (bool) {

        bytes memory toCaller = callerHashMap[_toChainId];
        require(toCaller.length != 0, "empty illegal toCallerHash");

        bytes memory tunnelData = abi.encode(Utils.addressToBytes(msg.sender), _toContract, _method, _txData);

        return IEthCrossChainManager(realCCM).crossChain(_toChainId, toCaller, "unwrap", tunnelData);
    }

    function unwrap(bytes memory args, bytes memory fromContractAddr, uint64 fromChainId) public returns (bool) {

        require(msg.sender == realCCM, "msgSender is not EthCrossChainManagerContract");
        require(fromContractAddr.length != 0, "from Caller contract address cannot be empty");
        require(Utils.equalStorage(callerHashMap[fromChainId], fromContractAddr), "From Caller contract address error!");

        (bytes memory _fromContractAddr, bytes memory _toContract, bytes memory _method, bytes memory _txData) = abi.decode(args,(bytes,bytes,bytes,bytes));

        (bool success, bytes memory returnData) = Utils.bytesToAddress(_toContract).call(
            abi.encodePacked(
                bytes4(keccak256(abi.encodePacked(_method, "(bytes,bytes,uint64)"))), 
                abi.encode(_txData, _fromContractAddr, fromChainId)
            )
        );

        require(success == true, "EthCrossChain call business contract failed");
        
        require(returnData.length != 0, "No return value from business contract!");
        bool res = abi.decode(returnData, (bool));
        require(res == true, "EthCrossChain call business contract return is not true");
        
        return true;
    }



}
