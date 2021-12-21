pragma solidity ^0.5.0;

import "../../core/cross_chain_manager/libs/upgradeability/InitializableAdminUpgradeabilityProxy.sol";

contract CallerImplementationMock {
    
    address public ccm;
    bytes4 constant crossChainSelector = 0xbd5cf625; //"crossChain(uint64,bytes,bytes,bytes)": "bd5cf625",
    event Unlock(address ccm, bytes args, bytes fromContract, uint64 fromChainId);

    function initialize(address _ccm) public {
        require(ccm==address(0),"already initialized");
        ccm = _ccm;
    }
    
    function lock(bytes memory args) public {
        (bool success,) = ccm.call(abi.encodeWithSelector(crossChainSelector,2,addressToBytes(address(this)),"unlock",args));
        require(success,"call ccm failed");
    }
    
    function unlock(bytes memory args, bytes memory fromContract, uint64 fromChainId) public returns(bool){
        require(msg.sender==ccm, "unlock caller not cmm");
        emit Unlock(msg.sender, args, fromContract, fromChainId);
    }
    
    function addressToBytes(address _addr) internal pure returns (bytes memory bs){
        assembly {
            bs := mload(0x40)
            mstore(bs, 0x14)
            mstore(add(bs, 0x20), shl(96, _addr))
            mstore(0x40, add(bs, 0x40))
       }
    }

    function whoAmI() public pure returns(uint) {
        return 1;
    }

    function fallback() public payable {
        revert("unknown function");
    }
}

contract CallerImplementationMock_2 {
    
    function whoAmI() public pure returns(uint) {
        return 2;
    }

}

contract CallerSigMsgGen {
    function getSigMsg(uint256 _salt, address _logic, address _admin, bytes memory _data, address _factory) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(_salt, _logic, _admin, _data, _factory));
    }
}

contract EthSigMsgGen {
    function msgToEthSignedMessageHash(bytes memory _msg) public pure returns (bytes32) {
        return toEthSignedMessageHash(hashMsg(_msg));
    }

    function toEthSignedMessageHash(bytes32 hash) public pure returns (bytes32) {
        return keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hash));
    }

    function hashMsg(bytes memory _msg) public pure returns (bytes32) {
        return keccak256(_msg);
    }
}

