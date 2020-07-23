pragma solidity ^0.5.0;
import "./../../../libs/ownership/Ownable.sol";
import "./../../../libs/lifecycle/Pausable.sol";
import "./../interface/IUpgradableECCM.sol";
import "./../interface/IEthCrossChainManagerProxy.sol";

contract EthCrossChainManagerProxy is IEthCrossChainManagerProxy, Ownable, Pausable {
    address private EthCrossChainManagerAddr_;
    
    constructor(address _ethCrossChainManagerAddr) public {
        EthCrossChainManagerAddr_ = _ethCrossChainManagerAddr;
    }
    
    function pause() onlyOwner public returns (bool) {
        if (paused()) {
            return true;
        }
        _pause();
        return true;
    }
    function unpause() onlyOwner public returns (bool) {
        if (!paused()) {
            return true;
        }
        _unpause();
        return true;
    }
    function pauseEthCrossChainManager() onlyOwner whenNotPaused public returns (bool) {
        IUpgradableECCM eccm = IUpgradableECCM(EthCrossChainManagerAddr_);
        require(pause(), "pause EthCrossChainManagerProxy contract failed!");
        require(eccm.pause(), "pause EthCrossChainManager contract failed!");
    }
    function upgradeEthCrossChainManager(address _newEthCrossChainManagerAddr) onlyOwner whenPaused public returns (bool) {
        IUpgradableECCM eccm = IUpgradableECCM(EthCrossChainManagerAddr_);
        if (!eccm.paused()) {
            require(eccm.pause(), "Pause old EthCrossChainManager contract failed!");
        }
        require(eccm.upgradeToNew(_newEthCrossChainManagerAddr), "EthCrossChainManager upgradeToNew failed!");
        IUpgradableECCM neweccm = IUpgradableECCM(_newEthCrossChainManagerAddr);
        require(neweccm.isOwner(), "EthCrossChainManagerProxy is not owner of new EthCrossChainManager contract");
        EthCrossChainManagerAddr_ = _newEthCrossChainManagerAddr;
    }
    function unpauseEthCrossChainManager() onlyOwner whenPaused public returns (bool) {
        IUpgradableECCM eccm = IUpgradableECCM(EthCrossChainManagerAddr_);
        require(eccm.unpause(), "unpause EthCrossChainManager contract failed!");
        require(unpause(), "unpause EthCrossChainManagerProxy contract failed!");
    }
    function getEthCrossChainManager() whenNotPaused public view returns (address) {
        return EthCrossChainManagerAddr_;
    }
}