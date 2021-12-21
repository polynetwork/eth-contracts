pragma solidity ^0.5.0;

contract CCMMock {
    
    function iAmMock() public pure returns(bool) {
        return true;
    }
    
    uint64 public _toChainId;
    bytes public _toContract; 
    bytes public _method; 
    bytes public _txData;
    bytes public _fromContract;
    uint64 public _fromChainId;

    function crossChain(
        uint64 toChainId, 
        bytes calldata toContract, 
        bytes calldata method, 
        bytes calldata txData
    ) external returns (bool) {
        _toChainId = toChainId;
        _toContract = toContract;
        _method = method;
        _txData = txData;
        _fromContract = addressToBytes(msg.sender);
        _fromChainId = toChainId;
        return true;
    }

    function verifyHeaderAndExecuteTx(
        bytes memory toContract,
        bytes memory method,
        bytes memory args,
        bytes memory fromContract,
        uint64 fromChainId
    ) public returns (bool)
    {
        (bool success, bytes memory returnData) = bytesToAddress(toContract).call(abi.encodePacked(bytes4(keccak256(abi.encodePacked(method, "(bytes,bytes,uint64)"))), abi.encode(args, fromContract, fromChainId)));
        
        require(success == true, "EthCrossChain call business contract failed");
        
        require(returnData.length != 0, "No return value from business contract!");
        bool res = abi.decode(returnData, (bool));
        require(res == true, "EthCrossChain call business contract return is not true");
        
        return true;
    }

    function clearStorage() public {
        delete _toChainId;
        delete _toContract; 
        delete _method; 
        delete _txData;
        delete _fromContract;
        delete _fromChainId;
    }

    function bytesToAddress(bytes memory _bs) internal pure returns (address addr)
    {
        require(_bs.length == 20, "bytes length does not match address");
        assembly {
            addr := mload(add(_bs, 0x14))
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
}