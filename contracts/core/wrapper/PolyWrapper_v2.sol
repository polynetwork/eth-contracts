// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.5.0;

import "../../libs/token/ERC20/SafeERC20.sol";
import "../../libs/token/ERC20/IERC20.sol";
import "../../libs/ownership/Ownable.sol";
import "../../libs/utils/ReentrancyGuard.sol";
import "../../libs/math/SafeMath.sol";
import "../../libs/lifecycle/Pausable.sol";

import "./interfaces/ILockProxy.sol";

contract PolyWrapperV2 is Ownable, Pausable, ReentrancyGuard {
    using SafeMath for uint;
    using SafeERC20 for IERC20;

    uint public chainId;
    address public feeCollector;

    ILockProxy public lockProxy;

    constructor(address _owner, uint _chainId) public {
        require(_chainId != 0, "!legal");
        transferOwnership(_owner);
        chainId = _chainId;
    }

    function setFeeCollector(address collector) external onlyOwner {
        require(collector != address(0), "emtpy address");
        feeCollector = collector;
    }


    function setLockProxy(address _lockProxy) external onlyOwner {
        require(_lockProxy != address(0));
        lockProxy = ILockProxy(_lockProxy);
        require(lockProxy.managerProxyContract() != address(0), "not lockproxy");
    }

    function pause() external onlyOwner {
        _pause();
    }

    function unpause() external onlyOwner {
        _unpause();
    }


    function extractFee(address token) external {
        require(msg.sender == feeCollector, "!feeCollector");
        if (token == address(0)) {
            msg.sender.transfer(address(this).balance);
        } else {
            IERC20(token).safeTransfer(feeCollector, IERC20(token).balanceOf(address(this)));
        }
    }
    
    function lock(address fromAsset, uint64 toChainId, bytes memory toAddress, uint amount, uint fee, uint id) public payable nonReentrant whenNotPaused {
        
        require(toChainId != chainId && toChainId != 0, "!toChainId");
        require(toAddress.length !=0, "empty toAddress");
        address addr;
        assembly { addr := mload(add(toAddress,0x14)) }
        require(addr != address(0),"zero toAddress");
        
        _pull(fromAsset, amount);

        amount = _checkoutFee(fromAsset, amount, fee);

        _push(fromAsset, toChainId, toAddress, amount);

        emit PolyWrapperLock(fromAsset, msg.sender, toChainId, toAddress, amount, fee, id);
    }

    function speedUp(address fromAsset, bytes memory txHash, uint fee) public payable nonReentrant whenNotPaused {
        _pull(fromAsset, fee);
        emit PolyWrapperSpeedUp(fromAsset, txHash, msg.sender, fee);
    }

    function _pull(address fromAsset, uint amount) internal {
        if (fromAsset == address(0)) {
            require(msg.value == amount, "insufficient ether");
        } else {
            IERC20(fromAsset).safeTransferFrom(msg.sender, address(this), amount);
        }
    }

    // take fee in the form of ether
    function _checkoutFee(address fromAsset, uint amount, uint fee) internal view returns (uint) {
        if (fromAsset == address(0)) {
            require(msg.value >= amount, "insufficient ether");
            require(amount > fee, "amount less than fee");
            return amount.sub(fee);
        } else {
            require(msg.value >= fee, "insufficient ether");
            return amount;
        }
    }

    function _push(address fromAsset, uint64 toChainId, bytes memory toAddress, uint amount) internal {
        if (fromAsset == address(0)) {
            require(lockProxy.lock.value(amount)(fromAsset, toChainId, toAddress, amount), "lock ether fail");
        } else {
            IERC20(fromAsset).safeApprove(address(lockProxy), 0);
            IERC20(fromAsset).safeApprove(address(lockProxy), amount);
            require(lockProxy.lock(fromAsset, toChainId, toAddress, amount), "lock erc20 fail");
        }
    }

    event PolyWrapperLock(address indexed fromAsset, address indexed sender, uint64 toChainId, bytes toAddress, uint net, uint fee, uint id);
    event PolyWrapperSpeedUp(address indexed fromAsset, bytes indexed txHash, address indexed sender, uint efee);

}