// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.16;

import 'forge-std/Test.sol';

import {Swivel} from 'src/Swivel.sol';
import {MarketPlace} from 'src/MarketPlace.sol';
import {Creator} from 'src/Creator.sol';
import {VaultTracker} from 'src/VaultTracker.sol';
import {Compound} from 'src/adapters/Compound.sol';
import {ZcToken} from 'src/tokens/ZcToken.sol';

import {IERC20} from 'src/interfaces/IERC20.sol';
import {IVaultTracker} from 'src/interfaces/IVaultTracker.sol';

import {Sig} from 'src/lib/Sig.sol';
import {Hash} from 'src/lib/Hash.sol';

import {Constants} from 'test/Constants.sol';
import {TestHash} from 'test/TestHash.sol';
import {Protocols} from 'test/Protocols.sol';

contract AdminTest is Test {
    Creator creator;
    MarketPlace marketplace;
    Swivel swivel;
    ZcToken zcToken;
    VaultTracker vaultTracker;
    Compound adapter;

    address cUSDC = 0x39AA39c021dfbaE8faC545936693aC917d5E7563;
    uint256 maturity = 1672549200; // Jan 1, 2023 | 00:00:00 GMT +0
    uint256 startingBalance = 10000000e6;

    // common order info
    bytes32 key = bytes32('hello');
    uint256 principal = 1000e6;
    uint256 premium = 100e6;

    bytes32 key2 = bytes32('bytes32');
    uint256 principal2 = 2000e6;
    uint256 premium2 = 250e6;

    // test order
    Hash.Order order =
        Hash.Order(
            key,
            uint8(Protocols.Compound),
            Constants.userPublicKey,
            Constants.USDC,
            false,
            false,
            principal,
            premium,
            maturity,
            maturity
        );

    // test order
    Hash.Order order2 =
        Hash.Order(
            key2,
            uint8(Protocols.Compound),
            Constants.userPublicKey,
            Constants.USDC,
            false,
            false,
            principal2,
            premium2,
            maturity,
            maturity
        );

    function setUp() public {
        // Deploy swivel contracts
        creator = new Creator();
        marketplace = new MarketPlace(address(creator));
        swivel = new Swivel(address(marketplace), address(0));

        // Deploy adapter contract
        adapter = new Compound();
        adapter.approve(Constants.USDC, cUSDC);

        // Approve compound to spend from swivel
        vm.startPrank(address(swivel));
        IERC20(Constants.USDC).approve(cUSDC, type(uint256).max);
        vm.stopPrank();

        // Approve swivel to spend underlying
        vm.startPrank(address(this));
        IERC20(Constants.USDC).approve(address(swivel), type(uint256).max);
        vm.stopPrank();

        vm.startPrank(Constants.userPublicKey);
        IERC20(Constants.USDC).approve(address(swivel), type(uint256).max);
        vm.stopPrank();

        // Deal balance to key addresses
        deal(Constants.USDC, address(this), startingBalance);
        deal(Constants.USDC, Constants.userPublicKey, startingBalance);
    }

    function deployMarket() internal {
        creator.setMarketPlace(address(marketplace));
        marketplace.setSwivel(address(swivel));

        marketplace.createMarket(
            uint8(Protocols.Compound),
            maturity,
            cUSDC,
            address(adapter),
            'TEST-Constants.USDC',
            'zcUSD'
        );

        address[] memory underlyings = new address[](1);
        address[] memory compoundings = new address[](1);
        underlyings[0] = Constants.USDC;
        compoundings[0] = cUSDC;

        (, , address zcTokenAddr, address vaultTrackerAddr, ) = marketplace
            .markets(uint8(Protocols.Compound), Constants.USDC, maturity);
        zcToken = ZcToken(zcTokenAddr);
        vaultTracker = VaultTracker(vaultTrackerAddr);
    }

    function testSetAdapter() public {
        deployMarket();

        (address startingAdapter, , , address vaultTrackerAddr, ) = marketplace
            .markets(uint8(Protocols.Compound), Constants.USDC, maturity);

        assertEq(startingAdapter, address(adapter));
        assertEq(IVaultTracker(vaultTrackerAddr).adapter(), address(adapter));

        Compound newAdapter = new Compound();

        marketplace.setAdapter(
            Constants.USDC,
            maturity,
            uint8(Protocols.Compound),
            address(newAdapter)
        );

        (address updatedAdapter, , , , ) = marketplace.markets(
            uint8(Protocols.Compound),
            Constants.USDC,
            maturity
        );

        assertEq(updatedAdapter, address(newAdapter));
        assertEq(
            IVaultTracker(vaultTrackerAddr).adapter(),
            address(newAdapter)
        );
    }
}
