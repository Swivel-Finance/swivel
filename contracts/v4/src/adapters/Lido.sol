// SPDX-License-Identifier: AGPL-3.0

pragma solidity 0.8.16;

import {IAdapter} from '../interfaces/IAdapter.sol';
import {ILido} from '../interfaces/ILido.sol';
import {ILidoToken} from '../interfaces/ILidoToken.sol';

contract Lido is IAdapter {
    function approve(address a, address c) external {}

    function underlying(address c) external view returns (address) {
        return ILidoToken(c).stETH();
    }

    function exchangeRate(address c) external view returns (uint256) {
        return ILidoToken(c).stEthPerToken();
    }

    function deposit(address c, uint256 a) external returns (uint256) {
        return ILido(c).wrap(a);
    }

    function withdraw(address c, uint256 a) external returns (uint256) {
        ILido wstEth = ILido(c);

        return wstEth.unwrap(wstEth.getWstETHByStETH(a));
    }
}
