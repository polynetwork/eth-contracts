pragma solidity ^0.5.0;

import "./LockProxyPausable.sol";

contract LockProxyReviewable is LockProxyPausable {

    enum RequestState { Invalid, Pending, Executed, Banned, Approved }

    struct UnlockRequest {
        bytes argsBs;
        bytes fromContractAddr;
        uint64 fromChainId;
        RequestState state;
    }

    uint public latestRequestId;

    mapping(uint => UnlockRequest) public requests;

    mapping(address => uint) public limits;

    mapping(address => bool) public censors;

    event NewRequestEvent(uint requestId, bytes argsBs, bytes fromContractAddr, uint64 fromChainId);
    event ExecuteRequestEvent(uint requestId, address toAsset, address toAddress, uint amount);
    event BanRequestEvent(uint requestId, string reason);
    event UnbanRequestEvent(uint requestId, string reason);

    function unlock(bytes memory argsBs, bytes memory fromContractAddr, uint64 fromChainId) onlyManagerContract whenNotPaused public returns (bool) {
        TxArgs memory args = _deserializeTxArgs(argsBs);

        require(fromContractAddr.length != 0, "from proxy contract address cannot be empty");
        require(Utils.equalStorage(proxyHashMap[fromChainId], fromContractAddr), "From Proxy contract address error!");
        
        require(args.toAssetHash.length != 0, "toAssetHash cannot be empty");
        address toAsset = Utils.bytesToAddress(args.toAssetHash);
        require(assetHashMap[toAsset][fromChainId].length != 0, "toAsset not bind");

        require(args.toAddress.length != 0, "toAddress cannot be empty");
        address toAddress = Utils.bytesToAddress(args.toAddress);
        
        if (limits[toAsset] != 0 && args.amount >= limits[toAsset]) {
            latestRequestId += 1;
            requests[latestRequestId] = UnlockRequest({
                argsBs: argsBs,
                fromContractAddr: fromContractAddr,
                fromChainId: fromChainId,
                state: RequestState.Pending
            });
            emit NewRequestEvent(latestRequestId, argsBs, fromContractAddr, fromChainId);
        } else {
            _unlock(toAsset, toAddress, args.amount);
        }
        return true;
    } 

    function _unlock(address toAsset, address toAddress, uint256 amount) internal {
        require(_transferFromContract(toAsset, toAddress, amount), "transfer asset from lock_proxy contract to toAddress failed!");
        
        emit UnlockEvent(toAsset, toAddress, amount);
    }

    /*
     admin functions
    */
    function setLimitForToken(address toAsset, uint limit) public onlyOwner() {
        limits[toAsset] = limit;
    }

    function removeLimitForToken(address toAsset) public onlyOwner() {
        limits[toAsset] = 0;
    }

    // !! remove can not be undo
    function removeBannedRequest(uint requestId) public onlyOwner() {
        require(requests[requestId].state == RequestState.Banned, "this is not a banned request");
        requests[requestId].state = RequestState.Invalid;
    }

    function addCensor(address newCensor) public onlyOwner() {
        require(!censors[newCensor], "already added");
        censors[newCensor] = true;
    }

    function removeCensor(address censor) public onlyOwner() {
        require(censors[censor], "already removed");
        censors[censor] = false;
    }

    /*
     censor functions
    */
    modifier onlyCensor() {
        require(censors[msg.sender], "Access denied: not censor");
        _;
    }

    function approve(uint requestId, address toAsset, address toAddress, uint amount) public onlyCensor() {
        UnlockRequest memory request = requests[requestId];
        require(request.state == RequestState.Pending, "this is not a pending request");
        TxArgs memory args = _deserializeTxArgs(request.argsBs);
        address _toAsset = Utils.bytesToAddress(args.toAssetHash);
        address _toAddress = Utils.bytesToAddress(args.toAddress);
        require(_toAsset == toAsset && _toAddress == toAddress && args.amount == amount, "invalid TxArgs");

        requests[requestId].state = RequestState.Executed;
        _unlock(toAsset, toAddress, amount);

        emit ExecuteRequestEvent(requestId, toAsset, toAddress, amount);
    }

    function ban(uint requestId, string memory reason) public onlyCensor() {
        require(requests[requestId].state == RequestState.Pending, "this is not a pending request");
        requests[requestId].state = RequestState.Banned;

        emit BanRequestEvent(requestId, reason);
    }

    function unban(uint requestId, string memory reason , address toAsset, address toAddress, uint amount) public onlyCensor() {
        UnlockRequest memory request = requests[requestId];
        require(request.state == RequestState.Banned, "this is not a banned request");
        TxArgs memory args = _deserializeTxArgs(request.argsBs);
        address _toAsset = Utils.bytesToAddress(args.toAssetHash);
        address _toAddress = Utils.bytesToAddress(args.toAddress);
        require(_toAsset == toAsset && _toAddress == toAddress && args.amount == amount, "invalid TxArgs");

        requests[requestId].state = RequestState.Executed;
        _unlock(toAsset, toAddress, amount);

        emit UnbanRequestEvent(requestId, reason);
        emit ExecuteRequestEvent(requestId, toAsset, toAddress, amount);
    }
}