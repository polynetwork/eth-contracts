pragma solidity ^0.5.0;

import "./../../../libs/token/ERC20/ERC20Detailed.sol";
import "./../../../libs/GSN/Context.sol";
import "./../../../libs/token/ERC20/ERC20.sol";

contract eNEO is Context, ERC20, ERC20Detailed {
    
    constructor (address lockProxyContractAddress) public ERC20Detailed("NEO Token", "eNEO", 8) {
        _mint(lockProxyContractAddress, 10000000000000000);
    }
}