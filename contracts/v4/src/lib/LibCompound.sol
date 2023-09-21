// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.4;

import {FixedPointMathLib} from '../lib/FixedPointMathLib.sol';

import '../interfaces/ICERC20.sol';

/// @notice Get up to date cToken data without mutating state.
/// @author Transmissions11 (https://github.com/transmissions11/libcompound)
library LibCompound {
    using FixedPointMathLib for uint256;

    function viewExchangeRate(ICERC20 cToken) internal view returns (uint256) {
        uint256 accrualBlockNumberPrior = cToken.accrualBlockNumber();

        if (accrualBlockNumberPrior == block.number)
            return cToken.exchangeRateStored();

        uint256 totalCash = cToken.underlying().balanceOf(address(cToken));
        uint256 borrowsPrior = cToken.totalBorrows();
        uint256 reservesPrior = cToken.totalReserves();

        uint256 borrowRateMantissa = cToken.interestRateModel().getBorrowRate(
            totalCash,
            borrowsPrior,
            reservesPrior
        );

        require(borrowRateMantissa <= 0.0005e16, 'RATE_TOO_HIGH'); // Same as borrowRateMaxMantissa in CTokenInterfaces.sol

        uint256 interestAccumulated = (borrowRateMantissa *
            (block.number - accrualBlockNumberPrior)).mulWadDown(borrowsPrior);

        uint256 totalReserves = cToken.reserveFactorMantissa().mulWadDown(
            interestAccumulated
        ) + reservesPrior;
        uint256 totalBorrows = interestAccumulated + borrowsPrior;
        uint256 totalSupply = cToken.totalSupply();

        return
            totalSupply == 0
                ? cToken.initialExchangeRateMantissa()
                : (totalCash + totalBorrows - totalReserves).divWadDown(
                    totalSupply
                );
    }
}