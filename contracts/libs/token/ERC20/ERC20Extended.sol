pragma solidity ^0.5.0;

import "./ERC20Detailed.sol";
import "./../../GSN/Context.sol";
import "./ERC20.sol";

contract ERC20Extended is Context, ERC20, ERC20Detailed {

    address public managerContract; // here managerContract should only be set as the ETH cross chain managing contract address
    address public operator;    // operator should be the address who deploys this contract, and responsible for 'setManager' and 'bindContractAddrWithChainId'
    mapping(uint64 => bytes) public contractAddrBindChainId;

    event BindContractAddrWithChainIdEvent(uint64 chainId, bytes contractAddr);
    event SetManagerEvent(address managerContract);

    
    modifier onlyManagerContract() {
        require(_msgSender() == managerContract) ;
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
    function setManager(address ethCrossChainContractAddr) onlyOperator public {
        managerContract = ethCrossChainContractAddr;
        emit SetManagerEvent(managerContract);
    }
    
    /* @notice              Bind the target chain with the target chain id
    *  @param chainId       The target chain id
    *  @param contractAddr  The specific contract address in bytes format in the target chain
    */
    function bindContractAddrWithChainId(uint64 chainId, bytes memory contractAddr) onlyOperator public {
        require(chainId >= 0, "chainId illegal!");
        require(contractAddr.lenght <= 100, "contractAddr too long");
        contractAddrBindChainId[chainId] = contractAddr;
        emit BindContractAddrWithChainIdEvent(chainId, contractAddr);
    }
}