pragma solidity ^0.5.0;

import "./../interface/IEthCrossChainData.sol";
import "./../interface/IUpgradableECCM.sol";
import "./../../../libs/lifecycle/Pausable.sol";
import "./../../../libs/ownership/Ownable.sol";

contract UpgradableECCM is IUpgradableECCM, Ownable, Pausable {
    address public EthCrossChainDataAddress;
    uint64 public chainId;  
    
    constructor (address ethCrossChainDataAddr, uint64 _chainId) Pausable() Ownable()  public {
        EthCrossChainDataAddress = ethCrossChainDataAddr;
        chainId = _chainId;
    }
    function pause() onlyOwner public returns (bool) {
        if (!paused()) {
            _pause();
        }
        IEthCrossChainData eccd = IEthCrossChainData(EthCrossChainDataAddress);
        if (!eccd.paused()) {
            require(eccd.pause(), "pause EthCrossChainData contract failed");
        }
        return true;
    }
    
    function unpause() onlyOwner public returns (bool) {
        if (paused()) {
            _unpause();
        }
        IEthCrossChainData eccd = IEthCrossChainData(EthCrossChainDataAddress);
        if (eccd.paused()) {
            require(eccd.unpause(), "unpause EthCrossChainData contract failed");
        }
        return true;
    }

    // if we want to upgrade this contract, we need to invoke this method 
    function upgradeToNew(address newEthCrossChainManagerAddress) whenPaused onlyOwner public returns (bool) {
        IEthCrossChainData eccd = IEthCrossChainData(EthCrossChainDataAddress);
        eccd.transferOwnership(newEthCrossChainManagerAddress);
        return true;
    }
    
    function setChainId(uint64 _newChainId) whenPaused onlyOwner public returns (bool) {
        chainId = _newChainId;
        return true;
    }
}