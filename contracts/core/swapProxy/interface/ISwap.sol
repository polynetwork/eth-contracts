// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.5.0;

contract ISwap {
    int128 public N_COINS;
    address[] public coins;
    address public lp_token;
    uint public fee;
    uint256[] public balances;
    
    function add_liquidity_one_coin(uint256 _amount, address tokenIn, uint256 _min_mint_amount) public returns(uint256);
    function remove_liquidity_one_coin(uint256 _amount, address tokenOut, uint256 _min_token_out_amount) public returns(uint256);
    function exchange(address tokenIn, address tokenOut, uint256 _dx, uint256 _min_dy) public returns(uint256);
    
    function get_virtual_price() public view returns(uint256);
    function calc_token_amount(uint256[3] memory _amounts, bool _is_deposit) view public returns(uint256);
    function get_dy(int128 i,int128 j,uint256 _dx) view public returns(uint256);
    function calc_withdraw_one_coin(uint256 amount, int128 i) view public returns(uint256);
    
}

interface Ip {
    function add_liquidity_one_coin(uint256 _amount, address tokenIn, uint256 _min_mint_amount) external returns(uint256);
    function remove_liquidity_one_coin(uint256 _amount, address tokenOut, uint256 _min_token_out_amount) external returns(uint256);
    function exchange(address tokenIn, address tokenOut, uint256 _dx, uint256 _min_dy) external returns(uint256);
}