pragma solidity ^0.5.0;

import "./../../../../libs/token/ERC20/ERC20Detailed.sol";
import "./../../../../libs/GSN/Context.sol";
import "./../../../../libs/token/ERC20/ERC20Extended_v2.0.sol";


contract ERC20Template is Context, ERC20, ERC20Detailed {
    
    constructor () public ERC20Detailed("ERC20 Template", "ERC20T", 9) {
        _mint(_msgSender(), 10000000000000);
    }
}