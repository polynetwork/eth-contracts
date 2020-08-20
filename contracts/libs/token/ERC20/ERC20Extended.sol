pragma solidity ^0.5.0;

import "./ERC20Detailed.sol";
import "./../../GSN/Context.sol";
import "./ERC20.sol";
import "./../../../core/cross_chain_manager/interface/IEthCrossChainManagerProxy.sol";

contract ERC20Extended is Context, ERC20, ERC20Detailed {

    address public managerProxyContract; // here managerContract should only be set as the ETH cross chain managing contract address
    address public operator;    // operator should be the address who deploys this contract, and responsible for 'setManager' and 'bindContractAddrWithChainId'
    mapping(uint64 => bytes) public bondAssetHashes;

    event BindAssetHash(uint64 chainId, bytes contractAddr);
    event SetManagerProxyEvent(address managerContract);

    
    modifier onlyManagerContract() {
        IEthCrossChainManagerProxy ieccmp = IEthCrossChainManagerProxy(managerProxyContract);
        require(_msgSender() == ieccmp.getEthCrossChainManager(), "msgSender is not EthCrossChainManagerContract");
        _;
    }

    modifier onlyOperator() {
        require(_msgSender() == operator) ;
        _;
    }
    
    /* @notice              Mint amount of tokens to the account
    *  @param account       The account which will receive the minted tokens
    *  @param amount        The amount of tokens to be minted
    */
    function mint(address account, uint256 amount) public onlyManagerContract returns (bool) {
        _mint(account, amount);
        return true;
    }
    
    /* @notice              Burn amount of tokens from the msg.sender's balance
    *  @param amount        The amount of tokens to be burnt
    */
    function burn(uint256 amount) public returns (bool) {
        _burn(_msgSender(), amount);
        return true;
    }
    
    /* @notice                              Set the ETH cross chain contract as the manager such that the ETH cross chain contract 
    *                                       will be able to mint tokens to the designated account after a certain amount of tokens
    *                                       are locked in the source chain
    *  @param ethCrossChainContractAddr     The ETH cross chain management contract address
    */
    function setManagerProxy(address ethCrossChainManagerProxyAddr) onlyOperator public {
        managerProxyContract = ethCrossChainManagerProxyAddr;
        emit SetManagerProxyEvent(managerProxyContract);
    }
    
    /* @notice              Bind the target chain with the target chain id
    *  @param chainId       The target chain id
    *  @param contractAddr  The specific contract address in bytes format in the target chain
    */
    function bindAssetHash(uint64 chainId, bytes memory contractAddr) onlyOperator public {
        require(chainId != 0, "chainId illegal!");
        bondAssetHashes[chainId] = contractAddr;
        emit BindAssetHash(chainId, contractAddr);
    }
}