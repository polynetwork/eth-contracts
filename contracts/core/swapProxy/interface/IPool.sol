// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.5.0;

interface IPool {
    // ---- Pool
    function isPublicSwap() external view returns (bool);
    function isFinalized() external view returns (bool);
    function isBound(address t) external view returns (bool);
    function getNumTokens() external view returns (uint256);
    function getCurrentTokens() external view returns (address[] memory tokens);
    function getFinalTokens() external view returns (address[] memory tokens);
    function getDenormalizedWeight(address token) external view  returns (uint256);
    function getTotalDenormalizedWeight() external view returns (uint256);
    function getNormalizedWeight(address token) external view returns (uint256);
    function getBalance(address token) external view returns (uint256);
    function getSwapFee() external view returns (uint256);
    function getController() external view returns (address);
    
    function getSpotPrice(address tokenIn, address tokenOut) external view returns (uint256 spotPrice);
    function getSpotPriceSansFee(address tokenIn, address tokenOut) external view returns (uint256 spotPrice);
    
    function joinPool(uint256 poolAmountOut, uint256[] calldata maxAmountsIn) external ;
    function exitPool(uint256 poolAmountIn, uint256[] calldata minAmountsOut) external ;
    
    function swapExactAmountIn(
        address tokenIn, 
        uint256 tokenAmountIn, 
        address tokenOut, 
        uint256 minAmountOut, 
        uint256 maxPrice
    ) 
        external 
        returns (uint256 tokenAmountOut, uint256 spotPriceAfter);
    function swapExactAmountOut(
        address tokenIn, 
        uint256 maxAmountIn, 
        address tokenOut, 
        uint256 tokenAmountOut, 
        uint256 maxPrice
    ) 
        external 
        returns (uint256 tokenAmountIn, uint256 spotPriceAfter);
    function joinswapExternAmountIn(
        address tokenIn, 
        uint256 tokenAmountIn, 
        uint256 minPoolAmountOut
    ) 
        external 
        returns (uint256 poolAmountOut);
    function joinswapPoolAmountOut(
        address tokenIn, 
        uint256 poolAmountOut, 
        uint256 maxAmountIn
    ) 
        external 
        returns (uint256 tokenAmountIn);
    function exitswapPoolAmountIn(
        address tokenOut,
        uint256 poolAmountIn,
        uint256 minAmountOut
    ) 
        external 
        returns (uint256 tokenAmountOut);
    function exitswapExternAmountOut(
        address tokenOut, 
        uint256 tokenAmountOut, 
        uint256 maxPoolAmountIn
    ) 
        external 
        returns (uint256 poolAmountIn);
    
    // ---- Math
    function calcSpotPrice(
        uint tokenBalanceIn,
        uint tokenWeightIn,
        uint tokenBalanceOut,
        uint tokenWeightOut,
        uint swapFee
    )
        external pure
        returns (uint spotPrice);
    function calcOutGivenIn(
        uint tokenBalanceIn,
        uint tokenWeightIn,
        uint tokenBalanceOut,
        uint tokenWeightOut,
        uint tokenAmountIn,
        uint swapFee
    )
        external pure
        returns (uint tokenAmountOut);
    function calcInGivenOut(
        uint tokenBalanceIn,
        uint tokenWeightIn,
        uint tokenBalanceOut,
        uint tokenWeightOut,
        uint tokenAmountOut,
        uint swapFee
    )
        external pure
        returns (uint tokenAmountIn);
    function calcPoolOutGivenSingleIn(
        uint tokenBalanceIn,
        uint tokenWeightIn,
        uint poolSupply,
        uint totalWeight,
        uint tokenAmountIn,
        uint swapFee
    )
        external pure
        returns (uint poolAmountOut);
    function calcSingleInGivenPoolOut(
        uint tokenBalanceIn,
        uint tokenWeightIn,
        uint poolSupply,
        uint totalWeight,
        uint poolAmountOut,
        uint swapFee
    )
        external pure
        returns (uint tokenAmountIn);
    function calcSingleOutGivenPoolIn(
        uint tokenBalanceOut,
        uint tokenWeightOut,
        uint poolSupply,
        uint totalWeight,
        uint poolAmountIn,
        uint swapFee
    )
        external pure
        returns (uint tokenAmountOut);
    function calcPoolInGivenSingleOut(
        uint tokenBalanceOut,
        uint tokenWeightOut,
        uint poolSupply,
        uint totalWeight,
        uint tokenAmountOut,
        uint swapFee
    )
        external pure
        returns (uint poolAmountIn);
    
    // ---- Token 
    function name() external view returns (string memory);
    function symbol() external view returns (string memory);
    function decimals() external view returns(uint8);
    function allowance(address src, address dst) external view returns (uint);
    function balanceOf(address whom) external view returns (uint);
    function totalSupply() external view returns (uint);
    function approve(address dst, uint amt) external returns (bool);
    function increaseApproval(address dst, uint amt) external returns (bool);
    function decreaseApproval(address dst, uint amt) external returns (bool);
    function transfer(address dst, uint amt) external returns (bool);
    function transferFrom(address src, address dst, uint amt) external returns (bool);
    
}