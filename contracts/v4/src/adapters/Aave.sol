// SPDX-License-Identifier: AGPL-3.0

pragma solidity 0.8.16;

import {IAave} from '../interfaces/IAave.sol';
import {IAaveToken} from '../interfaces/IAaveToken.sol';
import {IAavePool} from '../interfaces/IAavePool.sol';
import {IAdapter} from '../interfaces/IAdapter.sol';
import {IERC20} from '../interfaces/IERC20.sol';

contract Aave is IAdapter {
    function underlying(address c) external view returns (address) {
        return IAaveToken(c).UNDERLYING_ASSET_ADDRESS();
    }

    function exchangeRate(address c) external view returns (uint256) {
        IAaveToken aToken = IAaveToken(c);
        return
            IAavePool(aToken.POOL()).getReserveNormalizedIncome(
                aToken.UNDERLYING_ASSET_ADDRESS()
            );
    }

    function deposit(address c, uint256 a) external returns (uint256) {
        // Get the starting balance of the compounding asset
        uint256 starting = IERC20(c).balanceOf(address(this));

        address underlyingToken = IAaveToken(c).UNDERLYING_ASSET_ADDRESS();

        IAave(0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9).deposit(
            underlyingToken,
            a,
            address(this),
            0
        );

        return IERC20(c).balanceOf(address(this)) - starting;
    }

    function withdraw(address c, uint256 a) external returns (uint256) {
        // Get the starting balance of the compounding token
        uint256 starting = IERC20(c).balanceOf(address(this));

        // Withdraw the underlying amount
        address underlyingToken = IAaveToken(c).UNDERLYING_ASSET_ADDRESS();
        IAave(0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9).withdraw(
            underlyingToken,
            a,
            address(this)
        );

        return starting - IERC20(c).balanceOf(address(this));
    }
}
