pragma solidity ^0.5.0;

import "./../../../libs/GSN/Context.sol";
import "./../../../libs/token/ERC20/ERC20.sol";
import "./../../../libs/token/ERC20/ERC20Detailed.sol";

contract LPToken is Context, ERC20, ERC20Detailed {
    
    constructor(string memory name, string memory symbol, uint8 decimal, address initReceiver, uint initAmount) 
    public ERC20Detailed(name, symbol, decimal) {
        _mint(initReceiver, initAmount);
    }
}