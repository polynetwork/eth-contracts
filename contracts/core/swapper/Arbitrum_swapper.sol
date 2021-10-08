// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.5.0;

import "../../libs/ownership/Ownable.sol";
import "../../libs/common/ZeroCopySource.sol";
import "../../libs/common/ZeroCopySink.sol";
import "../../libs/utils/Utils.sol";
import "../../libs/token/ERC20/SafeERC20.sol";
import "../../libs/token/ERC20/IERC20.sol";
import "../../libs/lifecycle/Pausable.sol";
import "../../libs/utils/ReentrancyGuard.sol";
import "../cross_chain_manager/interface/IEthCrossChainManager.sol";
import "../cross_chain_manager/interface/IEthCrossChainManagerProxy.sol";

interface IWETH {
    function deposit() external payable;
    function transfer(address to, uint value) external returns (bool);
    function withdraw(uint) external;
}

contract Swapper is Ownable, Pausable, ReentrancyGuard {
    using SafeMath for uint;
    using SafeERC20 for IERC20;

    struct TxArgs {
        uint amount;
        uint minOut;
        uint64 toPoolId;
        uint64 toChainId;
        bytes fromAssetHash;
        bytes fromAddress;
        bytes toAssetHash;
        bytes toAddress;
    }
    
    address public WETH;
    address public feeCollector;
    address public lockProxyAddress;
    address public managerProxyContract;
    bytes public swapProxyHash;
    uint64 public swapChainId;
    uint64 public chainId;
   
    mapping(bytes => mapping(uint64 => bool)) public assetInPool;
    mapping(uint64 => address) public poolTokenMap;

    event LockEvent(address fromAssetHash, address fromAddress, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount);
    event SwapEvent(address fromAssetHash, address fromAddress, uint64 toPoolId, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint amount,uint fee, uint id);
    event AddLiquidityEvent(address fromAssetHash, address fromAddress,  uint64 toPoolId, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint amount, uint fee, uint id);
    event RemoveLiquidityEvent(address fromAssetHash, address fromAddress,  uint64 toPoolId, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint amount, uint fee, uint id);
    
    constructor(address _owner, uint64 _chainId, uint64 _swapChianId, address _lockProxy, address _CCMP, address _weth, bytes memory _swapProxyHash) public {
        require(_chainId != 0, "!legal");
        transferOwnership(_owner);
        chainId = _chainId;
        lockProxyAddress = _lockProxy;
        managerProxyContract = _CCMP;
        WETH = _weth;
        swapProxyHash = _swapProxyHash;
        swapChainId = _swapChianId;
    }

    modifier onlyManagerContract() {
        IEthCrossChainManagerProxy ieccmp = IEthCrossChainManagerProxy(managerProxyContract);
        require(_msgSender() == ieccmp.getEthCrossChainManager(), "msgSender is not EthCrossChainManagerContract");
        _;
    }

    function pause() external onlyOwner {
        _pause();
    }

    function unpause() external onlyOwner {
        _unpause();
    }
    
    function setFeeCollector(address collector) external onlyOwner {
        require(collector != address(0), "emtpy address");
        feeCollector = collector;
    }

    function setLockProxy(address _lockProxy) external onlyOwner {
        require(_lockProxy != address(0), "emtpy address");
        lockProxyAddress = _lockProxy;
    }
    
    function setManagerProxy(address ethCCMProxyAddr) onlyOwner public {
        managerProxyContract = ethCCMProxyAddr;
    }
    
    function setSwapProxyHash(bytes memory swapProxyAddr) onlyOwner public {
        swapProxyHash = swapProxyAddr;
    }
    
    function setSwapChainId(uint64 _swapChianId) onlyOwner public {
        swapChainId = _swapChianId;
    }
    
    function setWETH(address _weth) external onlyOwner {
        WETH = _weth;
    }
    
    function bindAssetAndPool(bytes memory fromAssetHash, uint64 poolId) onlyOwner public returns (bool) {
        assetInPool[fromAssetHash][poolId] = true;
        return true;
    }
    
    function bind3Asset(bytes memory asset1, bytes memory asset2, bytes memory asset3, uint64 poolId) onlyOwner public {
        assetInPool[asset1][poolId] = true;
        assetInPool[asset2][poolId] = true;
        assetInPool[asset3][poolId] = true;
    }
    
    function registerPoolWith3Assets(uint64 poolId, address poolTokenAddress, bytes memory asset1, bytes memory asset2, bytes memory asset3) onlyOwner public {
        poolTokenMap[poolId] = poolTokenAddress;
        assetInPool[asset1][poolId] = true;
        assetInPool[asset2][poolId] = true;
        assetInPool[asset3][poolId] = true;
    }
    
    function registerPool(uint64 poolId, address poolTokenAddress) onlyOwner public returns (bool) {
        poolTokenMap[poolId] = poolTokenAddress;
        return true;
    }
    
    function extractFee(address token) external {
        require(msg.sender == feeCollector, "!feeCollector");
        if (token == address(0)) {
            msg.sender.transfer(address(this).balance);
        } else {
            IERC20(token).safeTransfer(feeCollector, IERC20(token).balanceOf(address(this)));
        }
    }
    
    function swap(address fromAssetHash, uint64 toPoolId, uint64 toChainId, bytes memory toAssetHash, bytes memory toAddress, uint amount, uint minOutAmount, uint fee, uint id) public payable nonReentrant whenNotPaused returns (bool) {
        _pull(fromAssetHash, amount);
    
        amount = _checkoutFee(fromAssetHash, amount, fee);
        
        _push(fromAssetHash, amount);

        fromAssetHash = fromAssetHash==address(0) ? WETH : fromAssetHash ; 
        require(poolTokenMap[toPoolId] != address(0), "given pool do not exsit");
        require(assetInPool[Utils.addressToBytes(fromAssetHash)][toPoolId],"input token not in given pool");
        require(assetInPool[toAssetHash][toPoolId],"output token not in given pool");
        require(toAddress.length !=0, "empty toAddress");
        address addr;
        assembly { addr := mload(add(toAddress,0x14)) }
        require(addr != address(0),"zero toAddress");
         
        {
            TxArgs memory txArgs = TxArgs({
                amount: amount,
                minOut: minOutAmount,
                toPoolId: toPoolId,
                toChainId: toChainId,
                fromAssetHash: Utils.addressToBytes(fromAssetHash),
                fromAddress: Utils.addressToBytes(_msgSender()),
                toAssetHash: toAssetHash,
                toAddress: toAddress
            });
            bytes memory txData = _serializeTxArgs(txArgs);
            
            address eccmAddr = IEthCrossChainManagerProxy(managerProxyContract).getEthCrossChainManager();
            IEthCrossChainManager eccm = IEthCrossChainManager(eccmAddr);
            
            require(eccm.crossChain(swapChainId, swapProxyHash, "swap", txData), "EthCrossChainManager crossChain executed error!");
        }
        
        emit LockEvent(fromAssetHash, _msgSender(), swapChainId, Utils.addressToBytes(address(0)), swapProxyHash, amount);
        emit SwapEvent(fromAssetHash, _msgSender(), toPoolId, toChainId, toAssetHash, toAddress, amount, fee, id);
        
        return true;
    }
    
    function add_liquidity(address fromAssetHash, uint64 toPoolId, uint64 toChainId, bytes memory toAddress, uint amount, uint minOutAmount, uint fee, uint id) public payable nonReentrant whenNotPaused returns (bool) {
        _pull(fromAssetHash, amount);
            
        amount = _checkoutFee(fromAssetHash, amount, fee);
        
        _push(fromAssetHash, amount);

        fromAssetHash = fromAssetHash==address(0) ? WETH : fromAssetHash ;   
        require(poolTokenMap[toPoolId] != address(0), "given pool do not exsit");
        require(assetInPool[Utils.addressToBytes(fromAssetHash)][toPoolId],"input token not in given pool");
        require(toAddress.length !=0, "empty toAddress");
        address addr;
        assembly { addr := mload(add(toAddress,0x14)) }
        require(addr != address(0),"zero toAddress");
        
        {
            TxArgs memory txArgs = TxArgs({
                amount: amount,
                minOut: minOutAmount,
                toPoolId: toPoolId,
                toChainId: toChainId,
                fromAssetHash: Utils.addressToBytes(fromAssetHash),
                fromAddress: Utils.addressToBytes(_msgSender()),
                toAssetHash: Utils.addressToBytes(address(0)),
                toAddress: toAddress
            });
            bytes memory txData = _serializeTxArgs(txArgs);
            
            address eccmAddr = IEthCrossChainManagerProxy(managerProxyContract).getEthCrossChainManager();
            IEthCrossChainManager eccm = IEthCrossChainManager(eccmAddr);
            
            require(eccm.crossChain(swapChainId, swapProxyHash, "add", txData), "EthCrossChainManager crossChain executed error!");
        }

        emit LockEvent(fromAssetHash, _msgSender(), swapChainId, Utils.addressToBytes(address(0)), swapProxyHash, amount);
        emit AddLiquidityEvent(fromAssetHash, _msgSender(), toPoolId, toChainId, Utils.addressToBytes(address(0)), toAddress, amount, fee, id);
        
        return true;
    }
    
    function remove_liquidity(address fromAssetHash, uint64 toPoolId, uint64 toChainId, bytes memory toAssetHash, bytes memory toAddress, uint amount, uint minOutAmount, uint fee, uint id) public payable nonReentrant whenNotPaused returns (bool) {
        _pull(fromAssetHash, amount);
    
        amount = _checkoutFee(fromAssetHash, amount, fee);
        
        _push(fromAssetHash, amount);
            
        fromAssetHash = fromAssetHash==address(0) ? WETH : fromAssetHash ; 
        require(poolTokenMap[toPoolId] != address(0), "given pool do not exsit");
        require(poolTokenMap[toPoolId] == fromAssetHash,"input token is not pool LP token");
        require(assetInPool[toAssetHash][toPoolId],"output token not in given pool");
        require(toAddress.length !=0, "empty toAddress");
        address addr;
        assembly { addr := mload(add(toAddress,0x14)) }
        require(addr != address(0),"zero toAddress");
        
        {
            TxArgs memory txArgs = TxArgs({
                amount: amount,
                minOut: minOutAmount,
                toPoolId: toPoolId,
                toChainId: toChainId,
                fromAssetHash: Utils.addressToBytes(fromAssetHash),
                fromAddress: Utils.addressToBytes(_msgSender()),
                toAssetHash: toAssetHash,
                toAddress: toAddress
            });
            bytes memory txData = _serializeTxArgs(txArgs);
            
            address eccmAddr = IEthCrossChainManagerProxy(managerProxyContract).getEthCrossChainManager();
            IEthCrossChainManager eccm = IEthCrossChainManager(eccmAddr);
            
            require(eccm.crossChain(swapChainId, swapProxyHash, "remove", txData), "EthCrossChainManager crossChain executed error!");
        }
        
        emit LockEvent(fromAssetHash, _msgSender(), swapChainId, Utils.addressToBytes(address(0)), swapProxyHash, amount);
        emit RemoveLiquidityEvent(fromAssetHash, _msgSender(), toPoolId, toChainId, toAssetHash, toAddress, amount, fee, id);
        
        return true;
    }
    
    function getBalanceFor(address fromAssetHash) public view returns (uint256) {
        if (fromAssetHash == address(0)) {
            // return address(this).balance; // this expression would result in error: Failed to decode output: Error: insufficient data for uint256 type
            address selfAddr = address(this);
            return selfAddr.balance;
        } else {
            IERC20 erc20Token = IERC20(fromAssetHash);
            return erc20Token.balanceOf(address(this));
        }
    }
    
    // take input
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
            require(msg.value == amount, "insufficient ether");
            require(amount > fee, "amount less than fee");
            return amount.sub(fee);
        } else {
            require(msg.value == fee, "insufficient ether");
            return amount;
        }
    }
    
    // lock money in LockProxy, ether store in swapper 
    function _push(address fromAsset, uint amount) internal {
        if (fromAsset == address(0)) {
            // TODO: send ether to LockProxy, ** LockProxy do not have receive(),cannot send ether now
            IWETH(WETH).deposit.value(amount)();
            IWETH(WETH).transfer(lockProxyAddress, amount);
        } else {
            IERC20(fromAsset).safeTransfer(lockProxyAddress, amount);
        }
    }
    
    function _serializeTxArgs(TxArgs memory args) internal pure returns (bytes memory) {
        bytes memory buff;
        buff = abi.encodePacked(
            ZeroCopySink.WriteUint255(args.amount),
            ZeroCopySink.WriteUint255(args.minOut),
            ZeroCopySink.WriteUint64(args.toPoolId),
            ZeroCopySink.WriteUint64(args.toChainId),
            ZeroCopySink.WriteVarBytes(args.fromAssetHash),
            ZeroCopySink.WriteVarBytes(args.fromAddress),
            ZeroCopySink.WriteVarBytes(args.toAssetHash),
            ZeroCopySink.WriteVarBytes(args.toAddress)
            );
        return buff;
    }

}