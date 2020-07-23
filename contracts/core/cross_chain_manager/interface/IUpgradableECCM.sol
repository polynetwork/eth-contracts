pragma solidity ^0.5.0;

/**
 * @dev Interface of upgradableECCM to make ECCM be upgradable, the implementation is in UpgradableECCM.sol
 */
interface IUpgradableECCM {
    function pause() external returns (bool);
    function unpause() external returns (bool);
    function paused() external view returns (bool);
    function upgradeToNew(address) external returns (bool);
    function isOwner() external view returns (bool);
}
