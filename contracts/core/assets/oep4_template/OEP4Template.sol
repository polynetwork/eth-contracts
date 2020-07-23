pragma solidity ^0.5.0;

import "./../../../libs/token/ERC20/ERC20Detailed.sol";
import "./../../../libs/GSN/Context.sol";
import "./../../../libs/token//ERC20/ERC20.sol";
import "./../../../libs/math/SafeMath.sol";

contract OEP4Template is Context, ERC20, ERC20Detailed {
    address public Operator;
    address public proxyHash;
    constructor (address proxyContractAddress) public ERC20Detailed("OEP4 Template", "OEP4T", 9) {
        _mint(address(this), 10000 * 10 ** 9);
        Operator = _msgSender();
        proxyHash = proxyContractAddress;
    }
    
    modifier onlyOperator() {
        require(_msgSender() == Operator, "Only Operator has access!");
        _;
    }
    
    function deletageToProxy(address _proxyHash, uint256 _amount) onlyOperator public returns (bool) {
        if (proxyHash != address(0)) {
            require(_proxyHash == proxyHash, "proxy contract address cannot be changed!");
        } else {
            proxyHash = _proxyHash;
        }
        require(this.transfer(_proxyHash, _amount), "transfer token to proxy contract failed!");
        return true;
        
    }
}