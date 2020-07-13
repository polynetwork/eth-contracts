pragma solidity ^0.5.0;


import "../../../../contracts/libs/math/SafeMath.sol";


contract SafeMathMock {
    function sub(uint256 a, uint256 b) public pure returns (uint256) {
        return SafeMath.sub(a, b);
    }

    function add(uint256 a, uint256 b) public pure returns (uint256) {
        return SafeMath.add(a, b);
    }
    function mul(uint256 a, uint256 b) public pure returns (uint256) {
        return SafeMath.mul(a, b);
    }

    function div(uint256 a, uint256 b) public pure returns (uint256) {
        return SafeMath.div(a, b);
    }
    function mod(uint256 a, uint256 b) public pure returns (uint256) {
        return SafeMath.mod(a, b);
    }
}
