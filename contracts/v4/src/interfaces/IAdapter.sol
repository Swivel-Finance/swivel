// SPDX-License-Identifier: AGPL-3.0

pragma solidity 0.8.16;

interface IAdapter {
    function underlying(address) external view returns (address);

    function exchangeRate(address) external view returns (uint256);

    function deposit(address, uint256) external returns (uint256);

    function withdraw(address, uint256) external returns (uint256);
}
