pragma solidity ^0.5.0;

import "./LockProxyPausable.sol";

contract LockProxyLimited is LockProxyPausable {

    mapping(address => uint) public quota;
    mapping(address => uint) public usedQuota;
    mapping(address => uint) public refreshPeriod;
    mapping(address => uint) public refreshTimestamp;

    function unlock(bytes memory argsBs, bytes memory fromContractAddr, uint64 fromChainId) onlyManagerContract whenNotPaused public returns (bool) {
        TxArgs memory args = _deserializeTxArgs(argsBs);

        require(fromContractAddr.length != 0, "from proxy contract address cannot be empty");
        require(Utils.equalStorage(proxyHashMap[fromChainId], fromContractAddr), "From Proxy contract address error!");
        
        require(args.toAssetHash.length != 0, "toAssetHash cannot be empty");
        address toAsset = Utils.bytesToAddress(args.toAssetHash);
        require(assetHashMap[toAsset][fromChainId].length != 0, "toAsset not bind");

        require(args.toAddress.length != 0, "toAddress cannot be empty");
        address toAddress = Utils.bytesToAddress(args.toAddress);
 
        if (refreshPeriod[toAsset] != 0) {
            _refreshQuota(toAsset);
            require(usedQuota[toAsset] + args.amount <= quota[toAsset], "limit reached");
            usedQuota[toAsset] += args.amount;
        } else { // refreshPeriod = 0 means quota becomes single-time quota
            refreshTimestamp[toAsset] = block.timestamp;
            require(args.amount <= quota[toAsset], "limit reached");
        }

        require(_transferFromContract(toAsset, toAddress, args.amount), "transfer asset from lock_proxy contract to toAddress failed!");
        
        emit UnlockEvent(toAsset, toAddress, args.amount);
        
        return true;
    }

    function _refreshQuota(address asset) internal {
        if (refreshTimestamp[asset] + refreshPeriod[asset] <= block.timestamp) {
            refreshTimestamp[asset] = block.timestamp;
            usedQuota[asset] = 0;
        }
    }

    /*
     admin functions
    */
    function setQuota(address asset, uint _quota) public onlyOwner() {
        quota[asset] = _quota;
    }

    function setRefreshPeriod(address asset, uint _refreshPeriod) public onlyOwner() {
        refreshPeriod[asset] = _refreshPeriod;
        _refreshQuota(asset);
    }
}
