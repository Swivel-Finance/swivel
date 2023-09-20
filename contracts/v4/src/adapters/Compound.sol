// SPDX-License-Identifier: AGPL-3.0

pragma solidity 0.8.16;

import {LibCompound} from '../lib/LibCompound.sol';
import {FixedPointMathLib} from '../lib/FixedPointMathLib.sol';

import {IAdapter} from '../interfaces/IAdapter.sol';
import {ICompound} from '../interfaces/ICompound.sol';
import {IERC20, ICERC20} from '../interfaces/ICERC20.sol';

contract Compound is IAdapter {
    using FixedPointMathLib for uint256;

    /// @notice Approves cTokens to call `transferFrom`
    /// @param a Address of the asset being approved
    /// @param c Address of cToken being approved
    function approve(address a, address c) external {
        IERC20(a).approve(c, type(uint256).max);
    }

    /// @param c Compounding token address
    function underlying(address c) external view returns (address) {
        return ICompound(c).underlying();
    }

    /// @param c Compounding token address
    function exchangeRate(address c) external view returns (uint256) {
        return LibCompound.viewExchangeRate(ICERC20(c));
    }

    /// @param c Compounding token address
    /// @param a Amount of underlying to deposit into the protocol
    function deposit(address c, uint256 a) external returns (uint256) {
        // Get the starting balance of the compounding asset
        uint256 starting = ICERC20(c).balanceOf(address(this));

        // Mint cTokens
        ICompound(c).mint(a);

        // Return how many cTokens were obtained in the mint
        return ICERC20(c).balanceOf(address(this)) - starting;
    }

    /// @param c address of the compounding asset
    /// @param a amount, in underlying, to be withdrawn from the protocol
    function withdraw(address c, uint256 a) external returns (uint256) {
        // Get the starting balance of the compounding asset
        uint256 starting = ICERC20(c).balanceOf(address(this));

        // Redeem the underlying amount
        ICompound(c).redeemUnderlying(a);

        return starting - ICERC20(c).balanceOf(address(this));
    }
}
