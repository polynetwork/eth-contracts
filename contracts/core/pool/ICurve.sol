// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.5.0;

contract LPToken {
    function name() external view returns (string memory) {}
    function symbol() external view returns (string memory) {}
    function set_minter(address _minter) external {}
    function approve(address _spender, uint256 _vaule) external returns(bool) {}
}

contract Pool3 {
    address[3] public coins;
    uint256[3] public balances;
    address public owner;
    address public lp_token;

    function A() external view returns(uint256) {}
    function A_precise() external view returns(uint256) {}
    function get_virtual_price() external view returns(uint256) {}
    function calc_token_amount(uint256[3] calldata _amounts, bool _is_deposit) external view returns(uint256) {}
    function add_liquidity_one_coin(uint256 _token_amount, address  _coin, uint256 _min_mint_amount) external returns(uint256) {}
    function add_liquidity(uint256[3] calldata _amounts, uint256 _min_mint_amount) external returns(uint256) {}
    function get_dy(int128 i, int128 j, uint256 _dx) external view returns(uint256) {}
    function exchange(address xxx, address yyy, uint256 _dx, uint256 _min_dy) external returns(uint256) {}
    function remove_liquidity(uint256 _amount, uint256[3] calldata _min_amounts) external returns(uint256[3] memory) {}
    function remove_liquidity_imbalance(uint256[3] calldata _amounts, uint256 _max_burn_amount) external returns(uint256) {}
    function calc_withdraw_one_coin(uint256 _token_amount, int128 i) external view returns(uint256) {}
    function remove_liquidity_one_coin(uint256 _token_amount, address  _coin, uint256 _min_amount) external returns(uint256) {}
    
}

contract Pool4 {
    address[4] public coins;
    uint256[4] public balances;
    address public owner;
    address public lp_token;

    function A() external view returns(uint256) {}
    function A_precise() external view returns(uint256) {}
    function get_virtual_price() external view returns(uint256) {}
    function calc_token_amount(uint256[4] calldata _amounts, bool _is_deposit) external view returns(uint256) {}
    function add_liquidity_one_coin(uint256 _token_amount, address  _coin, uint256 _min_mint_amount) external returns(uint256) {}
    function add_liquidity(uint256[4] calldata _amounts, uint256 _min_mint_amount) external returns(uint256) {}
    function get_dy(int128 i, int128 j, uint256 _dx) external view returns(uint256) {}
    function exchange(address xxx, address yyy, uint256 _dx, uint256 _min_dy) external returns(uint256) {}
    function remove_liquidity(uint256 _amount, uint256[4] calldata _min_amounts) external returns(uint256[4] memory) {}
    function remove_liquidity_imbalance(uint256[4] calldata _amounts, uint256 _max_burn_amount) external returns(uint256) {}
    function calc_withdraw_one_coin(uint256 _token_amount, int128 i) external view returns(uint256) {}
    function remove_liquidity_one_coin(uint256 _token_amount, address  _coin, uint256 _min_amount) external returns(uint256) {}
}