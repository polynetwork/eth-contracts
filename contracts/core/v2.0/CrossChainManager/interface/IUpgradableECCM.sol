pragma solidity ^0.5.0;

/**
 * @dev Interface of the ERC20 standard as defined in the EIP. Does not include
 * the optional functions; to access them see {ERC20Detailed}.
 */
interface IUpgradableECCM {
    function pause() external returns (bool);
    function unpause() external returns (bool);
    function paused() external view returns (bool);
    function upgradeToNew(address) external returns (bool);
    function isOwner() external view returns (bool);
}
