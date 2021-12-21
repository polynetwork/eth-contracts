//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract Context {
    function _msgSender() internal view returns (address) {
        return msg.sender;
    }

    function _msgData() internal view returns (bytes memory) {
        this; // silence state mutability warning without generating bytecode - see https://github.com/ethereum/solidity/issues/2691
        return msg.data;
    }
}

contract Ownable is Context {
    address private _owner;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    /**
     * @dev Initializes the contract setting the deployer as the initial owner.
     */
    constructor () internal {
        address msgSender = _msgSender();
        _owner = msgSender;
        emit OwnershipTransferred(address(0), msgSender);
    }

    /**
     * @dev Returns the address of the current owner.
     */
    function owner() public view returns (address) {
        return _owner;
    }

    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyOwner() {
        require(isOwner(), "Ownable: caller is not the owner");
        _;
    }

    /**
     * @dev Returns true if the caller is the current owner.
     */
    function isOwner() public view returns (bool) {
        return _msgSender() == _owner;
    }

    /**
     * @dev Leaves the contract without owner. It will not be possible to call
     * `onlyOwner` functions anymore. Can only be called by the current owner.
     *
     * NOTE: Renouncing ownership will leave the contract without an owner,
     * thereby removing any functionality that is only available to the owner.
     */
    function renounceOwnership() public onlyOwner {
        emit OwnershipTransferred(_owner, address(0));
        _owner = address(0);
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Can only be called by the current owner.
     */
    function transferOwnership(address newOwner) public  onlyOwner {
        _transferOwnership(newOwner);
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     */
    function _transferOwnership(address newOwner) internal {
        require(newOwner != address(0), "Ownable: new owner is the zero address");
        emit OwnershipTransferred(_owner, newOwner);
        _owner = newOwner;
    }
}

contract ISwapProxy is Ownable {

    address public managerProxyContract;
    mapping(uint64 => bytes) public proxyHashMap;
    mapping(uint64 => bytes) public swapperHashMap;
    mapping(uint64 => address) public poolAddressMap;
    mapping(uint64 => mapping(uint64 => mapping(bytes => address))) public assetPoolMap;  // assetInPoolMap[poolId][chainId][sourceChainAssetHash] = swapAssetAddress
    mapping(address => mapping(uint64 => bytes)) public assetHashMap;  // assetHashMap[fromAssetAddress][toChainId] = toAssetHash

    function equalStorage(bytes storage _preBytes, bytes memory _postBytes) internal view returns (bool) {
        bool success = true;

        assembly {
            // we know _preBytes_offset is 0
            let fslot := sload(_preBytes_slot)
            // Arrays of 31 bytes or less have an even value in their slot,
            // while longer arrays have an odd value. The actual length is
            // the slot divided by two for odd values, and the lowest order
            // byte divided by two for even values.
            // If the slot is even, bitwise and the slot with 255 and divide by
            // two to get the length. If the slot is odd, bitwise and the slot
            // with -1 and divide by two.
            let slength := div(and(fslot, sub(mul(0x100, iszero(and(fslot, 1))), 1)), 2)
            let mlength := mload(_postBytes)

            // if lengths don't match the arrays are not equal
            switch eq(slength, mlength)
            case 1 {
                // fslot can contain both the length and contents of the array
                // if slength < 32 bytes so let's prepare for that
                // v. http://solidity.readthedocs.io/en/latest/miscellaneous.html#layout-of-state-variables-in-storage
                // slength != 0
                if iszero(iszero(slength)) {
                    switch lt(slength, 32)
                    case 1 {
                        // blank the last byte which is the length
                        fslot := mul(div(fslot, 0x100), 0x100)

                        if iszero(eq(fslot, mload(add(_postBytes, 0x20)))) {
                            // unsuccess:
                            success := 0
                        }
                    }
                    default {
                        // cb is a circuit breaker in the for loop since there's
                        //  no said feature for inline assembly loops
                        // cb = 1 - don't breaker
                        // cb = 0 - break
                        let cb := 1

                        // get the keccak hash to get the contents of the array
                        mstore(0x0, _preBytes_slot)
                        let sc := keccak256(0x0, 0x20)

                        let mc := add(_postBytes, 0x20)
                        let end := add(mc, mlength)

                        // the next line is the loop condition:
                        // while(uint(mc < end) + cb == 2)
                        for {} eq(add(lt(mc, end), cb), 2) {
                            sc := add(sc, 1)
                            mc := add(mc, 0x20)
                        } {
                            if iszero(eq(sload(sc), mload(mc))) {
                                // unsuccess:
                                success := 0
                                cb := 0
                            }
                        }
                    }
                }
            }
            default {
                // unsuccess:
                success := 0
            }
        }

        return success;
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
        require(equalStorage(assetHashMap[assetAddress][chainId], rawAssetHash), "invalid chain-asset pair");
        assetPoolMap[poolId][chainId][rawAssetHash] = assetAddress;
        return true;
    }
    
    function bindAssetHash(address fromAssetHash, uint64 toChainId, bytes memory toAssetHash) onlyOwner public returns (bool) {
        assetHashMap[fromAssetHash][toChainId] = toAssetHash;
        return true;
    }

}

contract IPool {
    address public lp_token;
    constructor(address lp) public {
        lp_token = lp;
    }
}

contract bindProxy is Ownable{
    
    ISwapProxy public sp;
    
    constructor(address swapProxyAddr) public {
        sp = ISwapProxy(swapProxyAddr);
    }
    
    function setSwapProxyAddr(address swapProxyAddr) public onlyOwner {
        sp = ISwapProxy(swapProxyAddr);
    }
    
    function transferSwapProxyOwner(address new_Owner) public onlyOwner {
        require(address(this)==sp.owner(), "not SwapProxy owner");
        sp.transferOwnership(new_Owner);
    }
    
    function set4Pool(
        uint64 poolId, 
        address poolAddress, 
        uint64[4] memory chainIds, 
        address[4] memory assets, 
        bytes[4] memory rawAssets, 
        bytes[4] memory rawLps
        ) public onlyOwner {
        require(address(this)==sp.owner(), "not SwapProxy owner");
        address lp = IPool(poolAddress).lp_token();
        require(lp != address(0), "given poolAddress is not a pool");
        
        sp.bindPoolAddress(poolId, poolAddress);
        
        // bind lpToken
        sp.bindAssetHash(lp,chainIds[0],rawLps[0]);
        sp.bindAssetHash(lp,chainIds[1],rawLps[1]);
        sp.bindAssetHash(lp,chainIds[2],rawLps[2]);
        sp.bindAssetHash(lp,chainIds[3],rawLps[3]);
        
        // bind token in pool
        sp.bindAssetHash(assets[0],chainIds[0],rawAssets[0]);
        sp.bindPoolAssetAddress(poolId,chainIds[0],assets[0],rawAssets[0]);
        sp.bindAssetHash(assets[1],chainIds[1],rawAssets[1]);
        sp.bindPoolAssetAddress(poolId,chainIds[1],assets[1],rawAssets[1]);
        sp.bindAssetHash(assets[2],chainIds[2],rawAssets[2]);
        sp.bindPoolAssetAddress(poolId,chainIds[2],assets[2],rawAssets[2]);
        sp.bindAssetHash(assets[3],chainIds[3],rawAssets[3]);
        sp.bindPoolAssetAddress(poolId,chainIds[3],assets[3],rawAssets[3]);
        
    }
    
    function set3Pool(
        uint64 poolId, 
        address poolAddress, 
        uint64[3] memory chainIds, 
        address[3] memory assets, 
        bytes[3] memory rawAssets, 
        bytes[3] memory rawLps
        ) public onlyOwner {
        require(address(this)==sp.owner(), "not SwapProxy owner");
        address lp = IPool(poolAddress).lp_token();
        require(lp != address(0), "given poolAddress is not a pool");
        
        sp.bindPoolAddress(poolId, poolAddress);
        
        // bind lpToken
        sp.bindAssetHash(lp,chainIds[0],rawLps[0]);
        sp.bindAssetHash(lp,chainIds[1],rawLps[1]);
        sp.bindAssetHash(lp,chainIds[2],rawLps[2]);
      
        // bind token in pool
        sp.bindAssetHash(assets[0],chainIds[0],rawAssets[0]);
        sp.bindPoolAssetAddress(poolId,chainIds[0],assets[0],rawAssets[0]);
        sp.bindAssetHash(assets[1],chainIds[1],rawAssets[1]);
        sp.bindPoolAssetAddress(poolId,chainIds[1],assets[1],rawAssets[1]);
        sp.bindAssetHash(assets[2],chainIds[2],rawAssets[2]);
        sp.bindPoolAssetAddress(poolId,chainIds[2],assets[2],rawAssets[2]);
    }
    
    function set2Pool(
        uint64 poolId, 
        address poolAddress, 
        uint64[2] memory chainIds, 
        address[2] memory assets, 
        bytes[2] memory rawAssets, 
        bytes[2] memory rawLps) 
        public onlyOwner {
        require(address(this)==sp.owner(), "not SwapProxy owner");
        address lp = IPool(poolAddress).lp_token();
        require(lp != address(0), "given poolAddress is not a pool");
        
        sp.bindPoolAddress(poolId, poolAddress);
        
        // bind lpToken
        sp.bindAssetHash(lp,chainIds[0],rawLps[0]);
        sp.bindAssetHash(lp,chainIds[1],rawLps[1]);
        
        // bind token in pool
        sp.bindAssetHash(assets[0],chainIds[0],rawAssets[0]);
        sp.bindPoolAssetAddress(poolId,chainIds[0],assets[0],rawAssets[0]);
        sp.bindAssetHash(assets[1],chainIds[1],rawAssets[1]);
        sp.bindPoolAssetAddress(poolId,chainIds[1],assets[1],rawAssets[1]);

    }
}

