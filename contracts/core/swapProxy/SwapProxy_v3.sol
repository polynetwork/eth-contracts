// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.5.0;

import "../../libs/ownership/Ownable.sol";
import "../../libs/common/ZeroCopySource.sol";
import "../../libs/common/ZeroCopySink.sol";
import "../../libs/utils/Utils.sol";
import "../../libs/token/ERC20/SafeERC20.sol";
import "../../libs/token/ERC20/IERC20.sol";
import "./interface/IEthCrossChainManager.sol";
import "./interface/IEthCrossChainManagerProxy.sol";
import "./interface/ISwap.sol";


contract SwapProxy is Ownable {
    using SafeMath for uint;
    using SafeERC20 for IERC20;

    struct TxArgs {
        bytes toAssetHash;
        bytes toAddress;
        uint256 amount;
    }
    
    struct SwapArgs {
        uint amount;
        uint minOut;
        uint64 toPoolId;
        uint64 toChainId;
        bytes fromAssetHash;
        bytes fromAddress;
        bytes toAssetHash;
        bytes toAddress;
    }
    
    address public managerProxyContract;
    mapping(uint64 => bytes) public proxyHashMap;
    mapping(uint64 => bytes) public swapperHashMap;
    mapping(uint64 => address) public poolAddressMap;
    mapping(uint64 => mapping(uint64 => mapping(bytes => address))) public assetPoolMap;  // assetInPoolMap[poolId][chainId][sourceChainAssetHash] = swapAssetAddress
    mapping(address => mapping(uint64 => bytes)) public assetHashMap;  // assetHashMap[fromAssetAddress][toChainId] = toAssetHash

    event RollBackEvent(uint64 backChainId, bytes backAssetHash, bytes backAddress, uint amount);
    event UnlockEvent(address toAssetHash, address toAddress, uint256 amount);
    event LockEvent(address fromAssetHash, address fromAddress, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount);
    event SwapEvent(uint64 toPoolId, address inAssetAddress, uint inAmount, address outAssetAddress, uint outAmount, uint64 toChainId, bytes toAssetHash, bytes toAddress);
    event AddLiquidityEvent(uint64 toPoolId, address inAssetAddress, uint inAmount, address poolTokenAddress, uint outLPAmount, uint64 toChainId, bytes toAssetHash, bytes toAddress);
    event RemoveLiquidityEvent(uint64 toPoolId, address poolTokenAddress, uint inLPAmount, address outAssetAddress, uint outAmount, uint64 toChainId, bytes toAssetHash, bytes toAddress);
    
    modifier onlyManagerContract() {
        IEthCrossChainManagerProxy ieccmp = IEthCrossChainManagerProxy(managerProxyContract);
        require(_msgSender() == ieccmp.getEthCrossChainManager(), "msgSender is not EthCrossChainManagerContract");
        _;
    }

    modifier onlyThis() {
        require(_msgSender() == address(this), "this is an internal_function in the form of public_function for catch error");
        _;
    }
    
    function update(address[] memory tokens, address newContract) onlyOwner public {
        for (uint i=0;i<tokens.length;i++) {
            require(_transferERC20FromContract(tokens[i], newContract, IERC20(tokens[i]).balanceOf(address(this))),
            "transfer asset to newContract failed!");
        }
    }
    
    function setManagerProxy(address ethCCMProxyAddr) onlyOwner public {
        managerProxyContract = ethCCMProxyAddr;
    }
    
    function bindProxyHash(uint64 toChainId, bytes memory targetProxyHash) onlyOwner public returns (bool) {
        proxyHashMap[toChainId] = targetProxyHash;
        return true;
    }
    
    function bindSwapperHash(uint64 toChainId, bytes memory targetSwapperHash) onlyOwner public returns (bool) {
        swapperHashMap[toChainId] = targetSwapperHash;
        return true;
    }
    
    function bindPoolAddress(uint64 poolId, address poolAddress) onlyOwner public returns (bool) {
        poolAddressMap[poolId] = poolAddress;
        return true;
    }
    
    function bindPoolAssetAddress(uint64 poolId, uint64 chainId, address assetAddress, bytes memory rawAssetHash) onlyOwner public returns (bool) {
        require(Utils.equalStorage(assetHashMap[assetAddress][chainId], rawAssetHash), "invalid chain-asset pair");
        assetPoolMap[poolId][chainId][rawAssetHash] = assetAddress;
        return true;
    }
    
    function bindAssetHash(address fromAssetHash, uint64 toChainId, bytes memory toAssetHash) onlyOwner public returns (bool) {
        assetHashMap[fromAssetHash][toChainId] = toAssetHash;
        return true;
    }
    
    
    function swap(bytes memory argsBs, bytes memory fromContractAddr, uint64 fromChainId) onlyManagerContract public returns (bool) {
        (bool success,) = address(this).call(abi.encodeWithSelector(this.swapUnderlying.selector, argsBs, fromContractAddr, fromChainId));
        
        if (!success) {
            _rollBack(argsBs, fromChainId);
        }
        return true;
    }

    function swapUnderlying(bytes memory argsBs, bytes memory fromContractAddr, uint64 fromChainId) onlyThis public returns (bool) {
        SwapArgs memory args = _deserializeSwapArgs(argsBs);
        
        require(fromContractAddr.length != 0, "from contract address cannot be empty");
        require(Utils.equalStorage(swapperHashMap[fromChainId], fromContractAddr), "from swapper contract address error!");
       
        address poolAddress = poolAddressMap[args.toPoolId];
        require(poolAddress != address(0), "pool do not exsit");
        
        address inAssetAddress = assetPoolMap[args.toPoolId][fromChainId][args.fromAssetHash];
        require(inAssetAddress != address(0), "inAssetHash cannot be empty");
        
        address outAssetAddress = assetPoolMap[args.toPoolId][args.toChainId][args.toAssetHash];
        require(outAssetAddress != address(0), "target asset do not exsit");

        require(args.toAddress.length != 0, "toAddress cannot be empty");
        
        require(args.toAssetHash.length != 0, "empty illegal toAssetHash");


        uint outAmount = _swapInPool(poolAddress, inAssetAddress, args.amount, outAssetAddress, args.minOut);
        
        require(_crossChain(args.toChainId, args.toAddress, args.toAssetHash, outAmount));


        emit UnlockEvent(inAssetAddress, address(this), args.amount);
        emit SwapEvent(args.toPoolId, inAssetAddress, args.amount, outAssetAddress, outAmount, args.toChainId, args.toAssetHash, args.toAddress);        
        emit LockEvent(outAssetAddress, address(this), args.toChainId, args.toAssetHash, args.toAddress, outAmount);
        
        return true;
        
    }
    
    function _swapInPool(address poolAddress, address tokenIn, uint inAmount, address tokenOut, uint minOutAmount) internal returns (uint) {
        ISwap pool = ISwap(poolAddress);
        
        IERC20(tokenIn).safeApprove(poolAddress, 0);
        IERC20(tokenIn).safeApprove(poolAddress, inAmount);
        uint outAmount = pool.exchange(tokenIn, tokenOut, inAmount, minOutAmount);
        
        return outAmount;
    }
    
    function add(bytes memory argsBs, bytes memory fromContractAddr, uint64 fromChainId) onlyManagerContract public returns (bool) {
        (bool success,) = address(this).call(abi.encodeWithSelector(this.addUnderlying.selector, argsBs, fromContractAddr, fromChainId));
        
        if (!success) {
            _rollBack(argsBs, fromChainId);
        }
        return true;
    }

    function addUnderlying(bytes memory argsBs, bytes memory fromContractAddr, uint64 fromChainId) onlyThis public returns (bool) {
        SwapArgs memory args = _deserializeSwapArgs(argsBs);
        
        require(fromContractAddr.length != 0, "from contract address cannot be empty");
        require(Utils.equalStorage(swapperHashMap[fromChainId], fromContractAddr), "from swapper contract address error!");
        
        address poolAddress = poolAddressMap[args.toPoolId];
        require(poolAddress != address(0), "pool do not exsit");
        
        address inAssetAddress = assetPoolMap[args.toPoolId][fromChainId][args.fromAssetHash];
        require(inAssetAddress != address(0), "inAssetHash cannot be empty");
        
        address outAssetAddress = ISwap(poolAddress).lp_token();
        
        require(args.toAddress.length != 0, "toAddress cannot be empty");
       
        bytes memory toAssetHash = assetHashMap[outAssetAddress][args.toChainId];
        require(toAssetHash.length != 0, "empty illegal toAssetHash");


        uint outAmount = _addInPool(poolAddress, inAssetAddress, args.amount, args.minOut);

        require(_crossChain(args.toChainId, args.toAddress, toAssetHash, outAmount));


        emit UnlockEvent(inAssetAddress, address(this), args.amount);
        emit AddLiquidityEvent(args.toPoolId, inAssetAddress, args.amount, poolAddress, outAmount, args.toChainId, toAssetHash, args.toAddress);
        emit LockEvent(ISwap(poolAddress).lp_token(), address(this), args.toChainId, toAssetHash, args.toAddress, outAmount);
        
        return true;
    }
    
    function _addInPool(address poolAddress, address tokenIn, uint inAmount, uint minOutAmount) internal returns (uint) {
        ISwap pool = ISwap(poolAddress);
        
        IERC20(tokenIn).safeApprove(poolAddress, 0);
        IERC20(tokenIn).safeApprove(poolAddress, inAmount);
        uint outAmount = pool.add_liquidity_one_coin(inAmount, tokenIn, minOutAmount);
        
        return outAmount;
    }
 
 
    function remove(bytes memory argsBs, bytes memory fromContractAddr, uint64 fromChainId) onlyManagerContract public returns (bool) {
        (bool success,) = address(this).call(abi.encodeWithSelector(this.removeUnderlying.selector, argsBs, fromContractAddr, fromChainId));
        
        if (!success) {
            _rollBack(argsBs, fromChainId);
        }
        return true;
    }

    function removeUnderlying(bytes memory argsBs, bytes memory fromContractAddr, uint64 fromChainId) onlyThis public returns (bool) {
        SwapArgs memory args = _deserializeSwapArgs(argsBs);
        
        require(fromContractAddr.length != 0, "from contract address cannot be empty");
        require(Utils.equalStorage(swapperHashMap[fromChainId], fromContractAddr), "from swapper contract address error!");
        
        address poolAddress = poolAddressMap[args.toPoolId];
        require(poolAddress != address(0), "pool do not exsit");
        require(Utils.equalStorage(assetHashMap[ISwap(poolAddress).lp_token()][fromChainId], args.fromAssetHash), "from Asset do not match pool token address");
        
        address outAssetAddress = assetPoolMap[args.toPoolId][args.toChainId][args.toAssetHash];
        require(outAssetAddress != address(0), "target asset do not exsit");
        
        require(args.toAddress.length != 0, "toAddress cannot be empty");

        uint outAmount = _removeInPool(poolAddress, args.amount, outAssetAddress, args.minOut);
        
        require(_crossChain(args.toChainId, args.toAddress, args.toAssetHash, outAmount));


        emit UnlockEvent(ISwap(poolAddress).lp_token(), address(this), args.amount);
        emit RemoveLiquidityEvent(args.toPoolId, poolAddress, args.amount, outAssetAddress, outAmount, args.toChainId, args.toAssetHash, args.toAddress);
        emit LockEvent(outAssetAddress, address(this), args.toChainId, args.toAssetHash, args.toAddress, outAmount);
        
        return true;
    }
    
    function _removeInPool(address poolAddress, uint inAmount, address tokenOut, uint minOutAmount) internal returns (uint) {
        ISwap pool = ISwap(poolAddress);
        
        IERC20(pool.lp_token()).safeApprove(poolAddress, 0);
        IERC20(pool.lp_token()).safeApprove(poolAddress, inAmount);
        uint outAmount = pool.remove_liquidity_one_coin(inAmount, tokenOut, minOutAmount);
        
        return outAmount;
    }
    
    function _rollBack(bytes memory argsBs, uint64 fromChainId) internal {
        SwapArgs memory args = _deserializeSwapArgs(argsBs);
        require(_crossChain(fromChainId, args.fromAddress, args.fromAssetHash, args.amount));
        
        emit UnlockEvent(address(0), address(0), 0);
        emit RollBackEvent(fromChainId, args.fromAssetHash, args.fromAddress, args.amount);
        emit LockEvent(address(0), address(0), fromChainId,args.fromAssetHash, args.fromAddress, args.amount);
    }
    
    // when get token out from pool,call _crossChain() to send outToken to the final chain
    // to avoid StackTooDeep error
    function _crossChain(uint64 toChainId, bytes memory toAddress, bytes memory toAssetHash, uint amount) internal returns (bool) {
        TxArgs memory txArgs = TxArgs({
            toAssetHash: toAssetHash,
            toAddress: toAddress,
            amount: amount
        });
        bytes memory txData = _serializeTxArgs(txArgs);
        
        address eccmAddr = IEthCrossChainManagerProxy(managerProxyContract).getEthCrossChainManager();
        IEthCrossChainManager eccm = IEthCrossChainManager(eccmAddr);
        
        bytes memory toProxyHash = proxyHashMap[toChainId];
        require(toProxyHash.length != 0, "empty illegal toProxyHash");
        require(eccm.crossChain(toChainId, toProxyHash, "unlock", txData), "EthCrossChainManager crossChain executed error!");
        
        return true;
    }
 
 
 
 
 
 
    
    /* @notice                  This function is meant to be invoked by the user,
    *                           a certin amount teokens will be locked in the proxy contract the invoker/msg.sender immediately.
    *                           Then the same amount of tokens will be unloked from target chain proxy contract at the target chain with chainId later.
    *  @param fromAssetHash     The asset address in current chain, uniformly named as `fromAssetHash`
    *  @param toChainId         The target chain id
    *                           
    *  @param toAddress         The address in bytes format to receive same amount of tokens in target chain 
    *  @param amount            The amount of tokens to be crossed from ethereum to the chain with chainId
    */
    function lock(address fromAssetHash, uint64 toChainId, bytes memory toAddress, uint256 amount) public payable returns (bool) {
        require(amount != 0, "amount cannot be zero!");
        
        
        require(_transferToContract(fromAssetHash, amount), "transfer asset from fromAddress to lock_proxy contract  failed!");
        
        bytes memory toAssetHash = assetHashMap[fromAssetHash][toChainId];
        require(toAssetHash.length != 0, "empty illegal toAssetHash");
        
        require(_crossChain(toChainId, toAddress, toAssetHash, amount));

        // TxArgs memory txArgs = TxArgs({
        //     toAssetHash: toAssetHash,
        //     toAddress: toAddress,
        //     amount: amount
        // });
        // bytes memory txData = _serializeTxArgs(txArgs);
        
        // IEthCrossChainManagerProxy eccmp = IEthCrossChainManagerProxy(managerProxyContract);
        // address eccmAddr = eccmp.getEthCrossChainManager();
        // IEthCrossChainManager eccm = IEthCrossChainManager(eccmAddr);
        
        // bytes memory toProxyHash = proxyHashMap[toChainId];
        // require(toProxyHash.length != 0, "empty illegal toProxyHash");
        // require(eccm.crossChain(toChainId, toProxyHash, "unlock", txData), "EthCrossChainManager crossChain executed error!");

        emit LockEvent(fromAssetHash, _msgSender(), toChainId, toAssetHash, toAddress, amount);
        
        return true;

    }
    
    // /* @notice                  This function is meant to be invoked by the ETH crosschain management contract,
    // *                           then mint a certin amount of tokens to the designated address since a certain amount 
    // *                           was burnt from the source chain invoker.
    // *  @param argsBs            The argument bytes recevied by the ethereum lock proxy contract, need to be deserialized.
    // *                           based on the way of serialization in the source chain proxy contract.
    // *  @param fromContractAddr  The source chain contract address
    // *  @param fromChainId       The source chain id
    // */
    function unlock(bytes memory argsBs, bytes memory fromContractAddr, uint64 fromChainId) onlyManagerContract public returns (bool) {
        TxArgs memory args = _deserializeTxArgs(argsBs);

        require(fromContractAddr.length != 0, "from proxy contract address cannot be empty");
        require(Utils.equalStorage(proxyHashMap[fromChainId], fromContractAddr), "From Proxy contract address error!");
        
        require(args.toAssetHash.length != 0, "toAssetHash cannot be empty");
        address toAssetHash = Utils.bytesToAddress(args.toAssetHash);

        require(args.toAddress.length != 0, "toAddress cannot be empty");
        address toAddress = Utils.bytesToAddress(args.toAddress);
        
        
        require(_transferFromContract(toAssetHash, toAddress, args.amount), "transfer asset from lock_proxy contract to toAddress failed!");
        
        emit UnlockEvent(toAssetHash, toAddress, args.amount);
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
    
    function _transferToContract(address fromAssetHash, uint256 amount) internal returns (bool) {
        if (fromAssetHash == address(0)) {
            // fromAssetHash === address(0) denotes user choose to lock ether
            // passively check if the received msg.value equals amount
            require(msg.value != 0, "transferred ether cannot be zero!");
            require(msg.value == amount, "transferred ether is not equal to amount!");
        } else {
            // make sure lockproxy contract will decline any received ether
            require(msg.value == 0, "there should be no ether transfer!");
            // actively transfer amount of asset from msg.sender to lock_proxy contract
            require(_transferERC20ToContract(fromAssetHash, _msgSender(), address(this), amount), "transfer erc20 asset to lock_proxy contract failed!");
        }
        return true;
    }
    
    function _transferFromContract(address toAssetHash, address toAddress, uint256 amount) internal returns (bool) {
        if (toAssetHash == address(0x0000000000000000000000000000000000000000)) {
            // toAssetHash === address(0) denotes contract needs to unlock ether to toAddress
            // convert toAddress from 'address' type to 'address payable' type, then actively transfer ether
            address(uint160(toAddress)).transfer(amount);
        } else {
            // actively transfer amount of asset from msg.sender to lock_proxy contract 
            require(_transferERC20FromContract(toAssetHash, toAddress, amount), "transfer erc20 asset to lock_proxy contract failed!");
        }
        return true;
    }
    
    
    function _transferERC20ToContract(address fromAssetHash, address fromAddress, address toAddress, uint256 amount) internal returns (bool) {
         IERC20 erc20Token = IERC20(fromAssetHash);
        //  require(erc20Token.transferFrom(fromAddress, toAddress, amount), "trasnfer ERC20 Token failed!");
         erc20Token.safeTransferFrom(fromAddress, toAddress, amount);
         return true;
    }
    function _transferERC20FromContract(address toAssetHash, address toAddress, uint256 amount) internal returns (bool) {
         IERC20 erc20Token = IERC20(toAssetHash);
        //  require(erc20Token.transfer(toAddress, amount), "trasnfer ERC20 Token failed!");
         erc20Token.safeTransfer(toAddress, amount);
         return true;
    }
    
    function _serializeTxArgs(TxArgs memory args) internal pure returns (bytes memory) {
        bytes memory buff;
        buff = abi.encodePacked(
            ZeroCopySink.WriteVarBytes(args.toAssetHash),
            ZeroCopySink.WriteVarBytes(args.toAddress),
            ZeroCopySink.WriteUint255(args.amount)
            );
        return buff;
    }

    function _deserializeTxArgs(bytes memory valueBs) internal pure returns (TxArgs memory) {
        TxArgs memory args;
        uint256 off = 0;
        (args.toAssetHash, off) = ZeroCopySource.NextVarBytes(valueBs, off);
        (args.toAddress, off) = ZeroCopySource.NextVarBytes(valueBs, off);
        (args.amount, off) = ZeroCopySource.NextUint255(valueBs, off);
        return args;
    }

    function _deserializeSwapArgs(bytes memory valueBs) internal pure returns (SwapArgs memory) {
        SwapArgs memory args;
        uint256 off = 0;
        (args.amount, off) = ZeroCopySource.NextUint255(valueBs, off);
        (args.minOut, off) = ZeroCopySource.NextUint255(valueBs, off);
        (args.toPoolId, off) = ZeroCopySource.NextUint64(valueBs, off);
        (args.toChainId, off) = ZeroCopySource.NextUint64(valueBs, off);
        (args.fromAssetHash, off) = ZeroCopySource.NextVarBytes(valueBs, off);
        (args.fromAddress, off) = ZeroCopySource.NextVarBytes(valueBs, off);
        (args.toAssetHash, off) = ZeroCopySource.NextVarBytes(valueBs, off);
        (args.toAddress, off) = ZeroCopySource.NextVarBytes(valueBs, off);
        return args;
    }
}