// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.16;

import 'forge-std/Test.sol';

import {Swivel} from 'src/Swivel.sol';
import {MarketPlace} from 'src/MarketPlace.sol';
import {Creator} from 'src/Creator.sol';
import {VaultTracker} from 'src/VaultTracker.sol';
import {Lido} from 'src/adapters/Lido.sol';
import {ZcToken} from 'src/tokens/ZcToken.sol';

import {IERC20} from 'src/interfaces/IERC20.sol';

import {Sig} from 'src/lib/Sig.sol';
import {Hash} from 'src/lib/Hash.sol';

import {Constants} from 'test/Constants.sol';
import {TestHash} from 'test/TestHash.sol';
import {Protocols} from 'test/Protocols.sol';

contract LidoTest is Test {
    Creator creator;
    MarketPlace marketplace;
    Swivel swivel;
    ZcToken zcToken;
    VaultTracker vaultTracker;
    Lido adapter;

    address wstETH = 0x7f39C581F595B53c5cb19bD0b3f8dA6c935E2Ca0;
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
            Constants.STETH,
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
            Constants.STETH,
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
        adapter = new Lido();
        adapter.approve(Constants.STETH, wstETH);

        // Approve compound to spend from swivel
        vm.startPrank(address(swivel));
        IERC20(Constants.STETH).approve(wstETH, type(uint256).max);
        vm.stopPrank();

        // Approve swivel to spend underlying
        vm.startPrank(address(this));
        IERC20(Constants.STETH).approve(address(swivel), type(uint256).max);
        vm.stopPrank();

        vm.startPrank(Constants.userPublicKey);
        IERC20(Constants.STETH).approve(address(swivel), type(uint256).max);
        vm.stopPrank();

        // Deal balance to key addresses
        dealStakedEth(address(this), startingBalance);
        dealStakedEth(Constants.userPublicKey, startingBalance);
    }

    function deployMarket() internal {
        creator.setMarketPlace(address(marketplace));
        marketplace.setSwivel(address(swivel));

        marketplace.createMarket(
            uint8(Protocols.Compound),
            maturity,
            wstETH,
            address(adapter),
            'TEST-Constants.STETH',
            'zcSTETH'
        );

        address[] memory underlyings = new address[](1);
        address[] memory compoundings = new address[](1);
        underlyings[0] = Constants.STETH;
        compoundings[0] = wstETH;

        (, , address zcTokenAddr, address vaultTrackerAddr, ) = marketplace
            .markets(uint8(Protocols.Compound), Constants.STETH, maturity);
        zcToken = ZcToken(zcTokenAddr);
        vaultTracker = VaultTracker(vaultTrackerAddr);
    }

    // underlying
    // 1. maker receives [a] and loses [principalFilled]
    // 2. msg.sender loses [a]

    // zc tokens
    // 1. maker receives [principalFilled]

    // n tokens
    // 1. msg.sender receives [principalFilled] - [fee]
    function testInitiateVaultFillingZcTokenInitiate() public {
        deployMarket();
        order.vault = false;
        order.exit = false;
        Hash.Order[] memory o = new Hash.Order[](1);
        o[0] = order;

        uint256[] memory a = new uint256[](1);
        uint256 amount = 50e6;
        a[0] = amount;

        bytes32 messageDigest = Hash.message(
            swivel.domain(),
            TestHash.order(order)
        );

        Sig.Components[] memory signatures = new Sig.Components[](1);
        (uint8 v, bytes32 r1, bytes32 s) = vm.sign(
            Constants.userPrivateKey,
            messageDigest
        );
        signatures[0] = Sig.Components(v, r1, s);

        uint256 makerStartingUnderlyingBalance = IERC20(order.underlying)
            .balanceOf(order.maker);
        uint256 takerStartingUnderlyingBalance = IERC20(order.underlying)
            .balanceOf(address(this));
        uint256 makerStartingZcBalance = zcToken.balanceOf(order.maker);
        (uint256 takerStartingNBalance, ) = vaultTracker.balancesOf(
            address(this)
        );

        swivel.initiate(o, a, signatures);

        uint256 transactionSize = (amount * principal) / premium;
        uint256 fee = transactionSize / swivel.feenominators(2);
        // maker balances
        assertEq(
            zcToken.balanceOf(Constants.userPublicKey),
            makerStartingZcBalance + transactionSize
        );
        assertEq(
            IERC20(order.underlying).balanceOf(order.maker),
            makerStartingUnderlyingBalance + amount - transactionSize + 1
        );

        // taker balances
        assertEq(
            IERC20(order.underlying).balanceOf(address(this)),
            takerStartingUnderlyingBalance - amount + 1
        );
        (uint256 takerEndingNBalance, ) = vaultTracker.balancesOf(
            address(this)
        );
        assertEq(
            takerEndingNBalance,
            takerStartingNBalance + transactionSize - fee
        );
    }

    // underlying
    // 1. msg.sender receives [premiumFilled] and loses [a + fee]
    // 2. maker loses [premiumFilled]

    // zc tokens
    // 1. msg.sender receives [a]

    // n tokens
    // 1. maker receives [a]
    function testInitiateZcTokenFillingVaultInitiate() public {
        deployMarket();
        order.vault = true;
        order.exit = false;
        Hash.Order[] memory o = new Hash.Order[](1);
        o[0] = order;

        uint256[] memory a = new uint256[](1);
        uint256 amount = 50e6;
        a[0] = amount;

        bytes32 messageDigest = Hash.message(
            swivel.domain(),
            TestHash.order(order)
        );

        Sig.Components[] memory signatures = new Sig.Components[](1);
        (uint8 v, bytes32 r1, bytes32 s) = vm.sign(
            Constants.userPrivateKey,
            messageDigest
        );
        signatures[0] = Sig.Components(v, r1, s);

        uint256 makerStartingUnderlyingBalance = IERC20(order.underlying)
            .balanceOf(order.maker);

        swivel.initiate(o, a, signatures);

        uint256 transactionSize = (amount * premium) / principal;
        uint256 fee = transactionSize / swivel.feenominators(0);

        // maker balanaces
        assertEq(
            IERC20(order.underlying).balanceOf(order.maker),
            makerStartingUnderlyingBalance - transactionSize + 1
        );
        (uint256 notionalBalance, ) = vaultTracker.balancesOf(
            Constants.userPublicKey
        );
        assertEq(notionalBalance, amount);

        // taker balances
        assertEq(
            IERC20(order.underlying).balanceOf(address(this)),
            makerStartingUnderlyingBalance + transactionSize - amount - fee
        );
        assertEq(zcToken.balanceOf(address(this)), amount);
    }

    // underlying
    // msg.sender loses [a - premiumFilled] and fee
    // maker receives [a - premiumFilled]

    // zc tokens
    // maker loses [a]
    // msg.sender receives [a]

    // n tokens
    // no activity
    function testInitiateZcTokenFillingZcTokenExit() public {
        deployMarket();
        uint256 amount = 50e6;
        deal(address(zcToken), Constants.userPublicKey, startingBalance, true);

        order.vault = false;
        order.exit = true;
        Hash.Order[] memory o = new Hash.Order[](1);
        o[0] = order;

        uint256[] memory a = new uint256[](1);
        a[0] = amount;

        bytes32 messageDigest = Hash.message(
            swivel.domain(),
            TestHash.order(order)
        );

        Sig.Components[] memory signatures = new Sig.Components[](1);
        (uint8 v, bytes32 r1, bytes32 s) = vm.sign(
            Constants.userPrivateKey,
            messageDigest
        );
        signatures[0] = Sig.Components(v, r1, s);

        uint256 makerStartingUnderlyingBalance = IERC20(order.underlying)
            .balanceOf(order.maker);
        uint256 takerStartingUnderlyingBalance = IERC20(order.underlying)
            .balanceOf(address(this));
        uint256 makerStartingZcBalance = zcToken.balanceOf(order.maker);
        uint256 takerStartingZcBalance = zcToken.balanceOf(address(this));

        swivel.initiate(o, a, signatures);

        uint256 transactionSize = (amount * premium) / principal;

        // maker balances
        assertEq(
            zcToken.balanceOf(order.maker),
            makerStartingZcBalance - amount
        );
        assertEq(
            IERC20(Constants.STETH).balanceOf(order.maker),
            makerStartingUnderlyingBalance + (amount - transactionSize)
        );

        // taker balances
        assertEq(
            zcToken.balanceOf(address(this)),
            takerStartingZcBalance + amount
        );

        {
            uint256 fee = transactionSize / swivel.feenominators(0);
            uint256 a1 = amount;
            assertEq(
                IERC20(Constants.STETH).balanceOf(address(this)),
                takerStartingUnderlyingBalance -
                    (a1 - transactionSize) -
                    fee +
                    1
            );
        }
    }

    // underlying
    // 1. maker receives [a]
    // 2. msg.sender loses [a]

    // zc tokens
    // no activity

    // n tokens
    // 1. maker loses [principalFilled]
    // 2. msg.sender gains [principalFilled] and loses [fee]
    function testInitiateVaultFillingVaultExit() public {
        deployMarket();
        uint256 amount = 50e6;
        deal(address(zcToken), Constants.userPublicKey, startingBalance, true);
        vm.startPrank(address(marketplace));
        uint256 transactionSize = (amount * principal) / premium;
        vaultTracker.addNotional(Constants.userPublicKey, transactionSize);
        vm.stopPrank();

        order.vault = true;
        order.exit = true;
        Hash.Order[] memory o = new Hash.Order[](1);
        o[0] = order;

        uint256[] memory a = new uint256[](1);
        a[0] = amount;

        bytes32 messageDigest = Hash.message(
            swivel.domain(),
            TestHash.order(order)
        );

        Sig.Components[] memory signatures = new Sig.Components[](1);
        (uint8 v, bytes32 r1, bytes32 s) = vm.sign(
            Constants.userPrivateKey,
            messageDigest
        );
        signatures[0] = Sig.Components(v, r1, s);

        uint256 makerStartingUnderlyingBalance = IERC20(order.underlying)
            .balanceOf(order.maker);
        uint256 takerStartingUnderlyingBalance = IERC20(order.underlying)
            .balanceOf(address(this));
        (uint256 makerStartingNBalance, ) = vaultTracker.balancesOf(
            order.maker
        );
        (uint256 takerStartingNBalance, ) = vaultTracker.balancesOf(
            address(this)
        );

        swivel.initiate(o, a, signatures);

        // maker balances
        (uint256 notionalBalance, ) = vaultTracker.balancesOf(order.maker);
        assertEq(notionalBalance, makerStartingNBalance - transactionSize);
        assertEq(
            IERC20(Constants.STETH).balanceOf(order.maker),
            makerStartingUnderlyingBalance + amount
        );

        // taker balances
        (notionalBalance, ) = vaultTracker.balancesOf(address(this));
        {
            uint256 ts = transactionSize;
            uint256 fee = transactionSize / swivel.feenominators(2);
            assertEq(notionalBalance, takerStartingNBalance + ts - fee);
        }
        assertEq(
            IERC20(Constants.STETH).balanceOf(address(this)),
            takerStartingUnderlyingBalance - amount + 1
        );
    }

    // underlying
    // msg.sender recevies [principalFilled - a] and loses [fee]
    // maker loses [principalFilled - a]

    // zc tokens
    // msg.sender loses [principalFilled]
    // maker receives [principalFilled]

    // n tokens
    // no activity
    function testExitZcTokenFillingZcTokenInitiate() public {
        deployMarket();
        deal(wstETH, address(swivel), startingBalance);
        deal(address(zcToken), address(this), startingBalance);
        uint256 amount = 50e6;
        vm.startPrank(address(marketplace));
        vaultTracker.addNotional(address(this), amount);
        vm.stopPrank();
        deal(address(zcToken), Constants.userPublicKey, startingBalance, true);

        order.vault = false;
        order.exit = false;
        Hash.Order[] memory o = new Hash.Order[](1);
        o[0] = order;

        uint256[] memory a = new uint256[](1);
        a[0] = amount;

        bytes32 messageDigest = Hash.message(
            swivel.domain(),
            TestHash.order(order)
        );

        Sig.Components[] memory signatures = new Sig.Components[](1);
        (uint8 v, bytes32 r1, bytes32 s) = vm.sign(
            Constants.userPrivateKey,
            messageDigest
        );
        signatures[0] = Sig.Components(v, r1, s);

        {
            uint256 makerStartingUnderlyingBalance = IERC20(order.underlying)
                .balanceOf(order.maker);
            uint256 takerStartingUnderlyingBalance = IERC20(order.underlying)
                .balanceOf(address(this));
            uint256 makerStartingZcBalance = zcToken.balanceOf(order.maker);
            uint256 takerStartingZcBalance = zcToken.balanceOf(address(this));

            swivel.exit(o, a, signatures);

            uint256 transactionSize = (amount * principal) / premium;
            // maker balances
            assertEq(
                IERC20(order.underlying).balanceOf(order.maker),
                makerStartingUnderlyingBalance - (transactionSize - amount) + 1
            );
            assertEq(
                zcToken.balanceOf(order.maker),
                makerStartingZcBalance + transactionSize
            );

            // taker balances
            {
                uint256 fee = transactionSize / swivel.feenominators(1);
                uint256 a1 = amount;
                assertEq(
                    IERC20(order.underlying).balanceOf(address(this)),
                    takerStartingUnderlyingBalance +
                        (transactionSize - a1) -
                        fee
                );
            }
            assertEq(
                zcToken.balanceOf(address(this)),
                takerStartingZcBalance - transactionSize
            );
        }
    }

    // underlying
    // 1. maker loses [premiumFilled]
    // 2. msg.sender receives [premiumFilled] and loses [fee]

    // zc tokens
    // no activity

    // n tokens
    // 1. maker receives [a]
    // 2. msg.sender loses [a]
    function testExitVaultFillingVaultInitiate() public {
        deployMarket();
        deal(wstETH, address(swivel), startingBalance);
        uint256 amount = 50e6;
        vm.startPrank(address(marketplace));
        vaultTracker.addNotional(address(this), amount);
        vm.stopPrank();
        deal(address(zcToken), Constants.userPublicKey, startingBalance, true);

        order.vault = true;
        order.exit = false;
        Hash.Order[] memory o = new Hash.Order[](1);
        o[0] = order;

        uint256[] memory a = new uint256[](1);
        a[0] = amount;

        bytes32 messageDigest = Hash.message(
            swivel.domain(),
            TestHash.order(order)
        );

        Sig.Components[] memory signatures = new Sig.Components[](1);
        (uint8 v, bytes32 r1, bytes32 s) = vm.sign(
            Constants.userPrivateKey,
            messageDigest
        );
        signatures[0] = Sig.Components(v, r1, s);

        uint256 makerStartingUnderlyingBalance = IERC20(order.underlying)
            .balanceOf(order.maker);
        uint256 takerStartingUnderlyingBalance = IERC20(order.underlying)
            .balanceOf(address(this));
        (uint256 makerStartingNBalance, ) = vaultTracker.balancesOf(
            order.maker
        );
        (uint256 takerStartingNBalance, ) = vaultTracker.balancesOf(
            address(this)
        );

        swivel.exit(o, a, signatures);

        uint256 transactionSize = (amount * premium) / principal;
        // maker balances
        assertEq(
            IERC20(order.underlying).balanceOf(order.maker),
            makerStartingUnderlyingBalance - transactionSize + 1
        );
        (uint256 notionalBalance, ) = vaultTracker.balancesOf(order.maker);
        assertEq(notionalBalance, makerStartingNBalance + amount);

        // taker balances
        (notionalBalance, ) = vaultTracker.balancesOf(address(this));
        assertEq(notionalBalance, takerStartingNBalance - amount);
        {
            uint256 fee = transactionSize / swivel.feenominators(3);
            assertEq(
                IERC20(order.underlying).balanceOf(address(this)),
                takerStartingUnderlyingBalance + transactionSize - fee - 1
            );
        }
    }

    // underlying
    // 1. maker loses [a - premiumFilled]
    // 2. msg.sender loses [premiumFilled - fee]

    // zc tokens
    // 1. maker loses [a]

    // n tokens
    // 1. msg.sender loses [a]
    function testExitVaultFillingZcTokenExit() public {
        deployMarket();
        deal(wstETH, address(swivel), startingBalance);
        uint256 amount = 50e6;
        vm.startPrank(address(marketplace));
        vaultTracker.addNotional(address(this), amount);
        vm.stopPrank();
        deal(address(zcToken), Constants.userPublicKey, startingBalance, true);

        order.vault = false;
        order.exit = true;
        Hash.Order[] memory o = new Hash.Order[](1);
        o[0] = order;

        uint256[] memory a = new uint256[](1);
        a[0] = amount;

        bytes32 messageDigest = Hash.message(
            swivel.domain(),
            TestHash.order(order)
        );

        Sig.Components[] memory signatures = new Sig.Components[](1);
        (uint8 v, bytes32 r1, bytes32 s) = vm.sign(
            Constants.userPrivateKey,
            messageDigest
        );
        signatures[0] = Sig.Components(v, r1, s);

        uint256 makerStartingUnderlyingBalance = IERC20(order.underlying)
            .balanceOf(order.maker);
        uint256 takerStartingUnderlyingBalance = IERC20(order.underlying)
            .balanceOf(address(this));
        uint256 makerStartingZcBalance = zcToken.balanceOf(order.maker);
        (uint256 takerStartingNBalance, ) = vaultTracker.balancesOf(
            address(this)
        );

        swivel.exit(o, a, signatures);

        uint256 transactionSize = (amount * premium) / principal;
        // maker balances
        assertEq(
            IERC20(order.underlying).balanceOf(order.maker),
            makerStartingUnderlyingBalance + (amount - transactionSize)
        );
        assertEq(
            zcToken.balanceOf(Constants.userPublicKey),
            makerStartingZcBalance - amount
        );

        // taker balances
        {
            uint256 fee = transactionSize / swivel.feenominators(3);
            assertEq(
                IERC20(Constants.STETH).balanceOf(address(this)),
                takerStartingUnderlyingBalance + (transactionSize - fee) - 1
            );
        }
        (uint256 notionalBalance, ) = vaultTracker.balancesOf(address(this));
        assertEq(notionalBalance, takerStartingNBalance - amount);
    }

    // underlying
    // 1. msg.sender loses [principalFilled - a - fee]
    // 2. maker recevies [a]

    // zc tokens
    // 1. msg.sender loses [principalFilled]

    // n tokens
    // 1. maker loses [principalFilled]
    function testExitZcTokenFillingVaultExit() public {
        deployMarket();
        deal(wstETH, address(swivel), startingBalance);
        deal(address(zcToken), address(this), startingBalance);
        uint256 amount = 50e6;
        vm.startPrank(address(marketplace));
        vaultTracker.addNotional(Constants.userPublicKey, startingBalance);
        vm.stopPrank();
        deal(address(zcToken), Constants.userPublicKey, startingBalance, true);

        order.vault = true;
        order.exit = true;
        Hash.Order[] memory o = new Hash.Order[](1);
        o[0] = order;

        uint256[] memory a = new uint256[](1);
        a[0] = amount;

        bytes32 messageDigest = Hash.message(
            swivel.domain(),
            TestHash.order(order)
        );

        Sig.Components[] memory signatures = new Sig.Components[](1);
        (uint8 v, bytes32 r1, bytes32 s) = vm.sign(
            Constants.userPrivateKey,
            messageDigest
        );
        signatures[0] = Sig.Components(v, r1, s);

        uint256 makerStartingUnderlyingBalance = IERC20(order.underlying)
            .balanceOf(order.maker);
        uint256 takerStartingUnderlyingBalance = IERC20(order.underlying)
            .balanceOf(address(this));
        (uint256 makerStartingNBalance, ) = vaultTracker.balancesOf(
            order.maker
        );
        uint256 takerStartingZcBalance = zcToken.balanceOf(address(this));

        swivel.exit(o, a, signatures);

        uint256 transactionSize = (amount * principal) / premium;
        // maker balances
        (uint256 notionalBalance, ) = vaultTracker.balancesOf(order.maker);
        assertEq(notionalBalance, makerStartingNBalance - transactionSize);
        assertEq(
            IERC20(order.underlying).balanceOf(order.maker),
            makerStartingUnderlyingBalance + amount
        );
        // taker balances
        {
            uint256 fee = transactionSize / swivel.feenominators(1);
            uint256 a1 = amount;
            assertEq(
                IERC20(order.underlying).balanceOf(address(this)),
                takerStartingUnderlyingBalance +
                    (transactionSize - a1 - fee) -
                    1
            );
        }
        assertEq(
            zcToken.balanceOf(address(this)),
            takerStartingZcBalance - transactionSize
        );
    }

    function testInitiateVaultFillingZcTokenInitiateMultiple() public {
        deployMarket();
        order.vault = false;
        order.exit = false;
        Hash.Order[] memory o = new Hash.Order[](2);
        o[0] = order;
        o[1] = order2;

        uint256[] memory a = new uint256[](2);
        uint256 amount = 50e6;
        a[0] = amount;

        uint256 amount2 = 100e6;
        a[1] = amount2;

        bytes32 messageDigest = Hash.message(
            swivel.domain(),
            TestHash.order(order)
        );

        bytes32 messageDigest2 = Hash.message(
            swivel.domain(),
            TestHash.order(order2)
        );

        Sig.Components[] memory signatures = new Sig.Components[](2);
        (uint8 v, bytes32 r1, bytes32 s) = vm.sign(
            Constants.userPrivateKey,
            messageDigest
        );
        signatures[0] = Sig.Components(v, r1, s);

        (uint8 v2, bytes32 r12, bytes32 s2) = vm.sign(
            Constants.userPrivateKey,
            messageDigest2
        );
        signatures[1] = Sig.Components(v2, r12, s2);

        swivel.initiate(o, a, signatures);

        uint256 transactionSize = ((amount2 * principal2) / premium2) +
            ((amount * principal) / premium);
        assertEq(zcToken.balanceOf(Constants.userPublicKey), transactionSize);
        (uint256 notionalBalance, ) = vaultTracker.balancesOf(address(this));
        uint256 fee = transactionSize / swivel.feenominators(2);
        assertEq(notionalBalance, transactionSize - fee);
    }

    function testInitiateZcTokenFillingVaultInitiateMultple() public {
        deployMarket();
        order.vault = true;
        order.exit = false;
        order2.vault = true;
        order2.exit = false;
        Hash.Order[] memory o = new Hash.Order[](2);
        o[0] = order;
        o[1] = order2;

        uint256[] memory a = new uint256[](2);
        uint256 amount = 50e6;
        a[0] = amount;

        uint256 amount2 = 100e6;
        a[1] = amount2;

        bytes32 messageDigest = Hash.message(
            swivel.domain(),
            TestHash.order(order)
        );

        bytes32 messageDigest2 = Hash.message(
            swivel.domain(),
            TestHash.order(order2)
        );

        Sig.Components[] memory signatures = new Sig.Components[](2);
        (uint8 v, bytes32 r1, bytes32 s) = vm.sign(
            Constants.userPrivateKey,
            messageDigest
        );
        signatures[0] = Sig.Components(v, r1, s);

        (uint8 v2, bytes32 r12, bytes32 s2) = vm.sign(
            Constants.userPrivateKey,
            messageDigest2
        );
        signatures[1] = Sig.Components(v2, r12, s2);

        swivel.initiate(o, a, signatures);

        assertEq(zcToken.balanceOf(address(this)), amount + amount2);
        (uint256 notionalBalance, ) = vaultTracker.balancesOf(
            Constants.userPublicKey
        );
        assertEq(notionalBalance, amount + amount2);
    }

    function testInitiateZcTokenFillingZcTokenExitMultiple() public {
        deployMarket();
        deal(address(zcToken), Constants.userPublicKey, startingBalance, true);

        order.vault = false;
        order.exit = true;
        order2.vault = false;
        order2.exit = true;
        Hash.Order[] memory o = new Hash.Order[](2);
        o[0] = order;
        o[1] = order2;

        uint256[] memory a = new uint256[](2);
        uint256 amount = 50e6;
        a[0] = amount;

        uint256 amount2 = 100e6;
        a[1] = amount2;

        bytes32 messageDigest = Hash.message(
            swivel.domain(),
            TestHash.order(order)
        );

        bytes32 messageDigest2 = Hash.message(
            swivel.domain(),
            TestHash.order(order2)
        );

        Sig.Components[] memory signatures = new Sig.Components[](2);
        (uint8 v, bytes32 r1, bytes32 s) = vm.sign(
            Constants.userPrivateKey,
            messageDigest
        );
        signatures[0] = Sig.Components(v, r1, s);

        (uint8 v2, bytes32 r12, bytes32 s2) = vm.sign(
            Constants.userPrivateKey,
            messageDigest2
        );
        signatures[1] = Sig.Components(v2, r12, s2);

        swivel.initiate(o, a, signatures);

        assertEq(
            zcToken.balanceOf(Constants.userPublicKey),
            startingBalance - amount - amount2
        );
        assertEq(zcToken.balanceOf(address(this)), amount + amount2);
    }

    function testInitiateVaultFillingVaultExitMultiple() public {
        deployMarket();
        uint256 amount = 50e6;
        uint256 amount2 = 100e6;
        uint256 transactionSize = (amount * principal) / premium;
        transactionSize += (amount2 * principal2) / premium2;
        deal(address(zcToken), Constants.userPublicKey, startingBalance, true);
        vm.startPrank(address(marketplace));
        vaultTracker.addNotional(Constants.userPublicKey, transactionSize);
        vm.stopPrank();

        order.vault = true;
        order.exit = true;
        order2.vault = true;
        order2.exit = true;
        Hash.Order[] memory o = new Hash.Order[](2);
        o[0] = order;
        o[1] = order2;

        uint256[] memory a = new uint256[](2);
        a[0] = amount;
        a[1] = amount2;

        bytes32 messageDigest = Hash.message(
            swivel.domain(),
            TestHash.order(order)
        );

        bytes32 messageDigest2 = Hash.message(
            swivel.domain(),
            TestHash.order(order2)
        );

        Sig.Components[] memory signatures = new Sig.Components[](2);
        {
            (uint8 v, bytes32 r1, bytes32 s) = vm.sign(
                Constants.userPrivateKey,
                messageDigest
            );
            signatures[0] = Sig.Components(v, r1, s);

            (uint8 v2, bytes32 r12, bytes32 s2) = vm.sign(
                Constants.userPrivateKey,
                messageDigest2
            );
            signatures[1] = Sig.Components(v2, r12, s2);
        }

        swivel.initiate(o, a, signatures);

        (uint256 notionalBalance, ) = vaultTracker.balancesOf(
            Constants.userPublicKey
        );
        assertEq(notionalBalance, 0);
        (notionalBalance, ) = vaultTracker.balancesOf(address(this));
        uint256 fee = transactionSize / swivel.feenominators(2);
        assertEq(notionalBalance, transactionSize - fee);
    }

    function testExitZcTokenFillingZcTokenInitiateMultiple() public {
        deployMarket();
        deal(wstETH, address(swivel), startingBalance);
        deal(address(zcToken), address(this), startingBalance);
        uint256 amount = 50e6;
        uint256 amount2 = 100e6;
        vm.startPrank(address(marketplace));
        vaultTracker.addNotional(address(this), amount + amount2);
        vm.stopPrank();
        deal(address(zcToken), Constants.userPublicKey, startingBalance, true);

        order.vault = false;
        order.exit = false;
        order2.vault = false;
        order2.exit = false;
        Hash.Order[] memory o = new Hash.Order[](2);
        o[0] = order;
        o[1] = order2;

        uint256[] memory a = new uint256[](2);
        a[0] = amount;
        a[1] = amount2;

        bytes32 messageDigest = Hash.message(
            swivel.domain(),
            TestHash.order(order)
        );

        bytes32 messageDigest2 = Hash.message(
            swivel.domain(),
            TestHash.order(order2)
        );

        Sig.Components[] memory signatures = new Sig.Components[](2);
        (uint8 v, bytes32 r1, bytes32 s) = vm.sign(
            Constants.userPrivateKey,
            messageDigest
        );
        signatures[0] = Sig.Components(v, r1, s);

        (uint8 v2, bytes32 r12, bytes32 s2) = vm.sign(
            Constants.userPrivateKey,
            messageDigest2
        );
        signatures[1] = Sig.Components(v2, r12, s2);

        swivel.exit(o, a, signatures);

        (uint256 notionalBalance, ) = vaultTracker.balancesOf(
            Constants.userPublicKey
        );
        assertEq(notionalBalance, 0);
        assertGt(
            IERC20(Constants.STETH).balanceOf(address(this)),
            startingBalance
        );
    }

    function testExitVaultFillingVaultInitiateMultiple() public {
        deployMarket();
        deal(wstETH, address(swivel), startingBalance);
        uint256 amount = 50e6;
        uint256 amount2 = 50e6;
        vm.startPrank(address(marketplace));
        vaultTracker.addNotional(address(this), amount + amount2);
        vm.stopPrank();
        deal(address(zcToken), Constants.userPublicKey, startingBalance, true);

        order.vault = true;
        order.exit = false;
        order2.vault = true;
        order2.exit = false;
        Hash.Order[] memory o = new Hash.Order[](2);
        o[0] = order;
        o[1] = order2;

        uint256[] memory a = new uint256[](2);
        a[0] = amount;
        a[1] = amount2;

        bytes32 messageDigest = Hash.message(
            swivel.domain(),
            TestHash.order(order)
        );

        bytes32 messageDigest2 = Hash.message(
            swivel.domain(),
            TestHash.order(order2)
        );

        Sig.Components[] memory signatures = new Sig.Components[](2);
        {
            (uint8 v, bytes32 r1, bytes32 s) = vm.sign(
                Constants.userPrivateKey,
                messageDigest
            );
            signatures[0] = Sig.Components(v, r1, s);

            (uint8 v2, bytes32 r12, bytes32 s2) = vm.sign(
                Constants.userPrivateKey,
                messageDigest2
            );
            signatures[1] = Sig.Components(v2, r12, s2);
        }

        swivel.exit(o, a, signatures);

        (uint256 notionalBalance, ) = vaultTracker.balancesOf(
            Constants.userPublicKey
        );
        assertEq(notionalBalance, amount + amount2);
        assertGt(
            IERC20(Constants.STETH).balanceOf(address(this)),
            startingBalance
        );
    }

    function testExitVaultFillingZcTokenExitMultiple() public {
        deployMarket();
        dealStakedEth(address(swivel), startingBalance);
        uint256 amount = 50e6;
        uint256 amount2 = 100e6;
        vm.startPrank(address(marketplace));
        vaultTracker.addNotional(address(this), amount + amount2);
        vm.stopPrank();
        deal(address(zcToken), Constants.userPublicKey, startingBalance, true);

        order.vault = false;
        order.exit = true;
        order2.vault = false;
        order2.exit = true;
        Hash.Order[] memory o = new Hash.Order[](2);
        o[0] = order;
        o[1] = order2;

        uint256[] memory a = new uint256[](2);
        a[0] = amount;
        a[1] = amount2;

        bytes32 messageDigest = Hash.message(
            swivel.domain(),
            TestHash.order(order)
        );

        bytes32 messageDigest2 = Hash.message(
            swivel.domain(),
            TestHash.order(order2)
        );

        Sig.Components[] memory signatures = new Sig.Components[](2);
        (uint8 v, bytes32 r1, bytes32 s) = vm.sign(
            Constants.userPrivateKey,
            messageDigest
        );
        signatures[0] = Sig.Components(v, r1, s);

        (uint8 v2, bytes32 r12, bytes32 s2) = vm.sign(
            Constants.userPrivateKey,
            messageDigest2
        );
        signatures[1] = Sig.Components(v2, r12, s2);

        swivel.exit(o, a, signatures);

        (uint256 notionalBalance, ) = vaultTracker.balancesOf(
            Constants.userPublicKey
        );
        assertEq(notionalBalance, 0);
        assertGt(
            IERC20(Constants.STETH).balanceOf(address(this)),
            startingBalance
        );
    }

    function testExitZcTokenFillingVaultExitMultiple() public {
        deployMarket();
        deal(wstETH, address(swivel), startingBalance);
        deal(address(zcToken), address(this), startingBalance);
        uint256 amount = 50e6;
        uint256 amount2 = 100e6;
        vm.startPrank(address(marketplace));
        vaultTracker.addNotional(Constants.userPublicKey, startingBalance);
        vm.stopPrank();
        deal(address(zcToken), Constants.userPublicKey, startingBalance, true);

        order.vault = true;
        order.exit = true;
        order2.vault = true;
        order2.exit = true;
        Hash.Order[] memory o = new Hash.Order[](2);
        o[0] = order;
        o[1] = order2;

        uint256[] memory a = new uint256[](2);
        a[0] = amount;
        a[1] = amount2;

        bytes32 messageDigest = Hash.message(
            swivel.domain(),
            TestHash.order(order)
        );

        bytes32 messageDigest2 = Hash.message(
            swivel.domain(),
            TestHash.order(order2)
        );

        Sig.Components[] memory signatures = new Sig.Components[](2);
        {
            (uint8 v, bytes32 r1, bytes32 s) = vm.sign(
                Constants.userPrivateKey,
                messageDigest
            );
            signatures[0] = Sig.Components(v, r1, s);

            (uint8 v2, bytes32 r12, bytes32 s2) = vm.sign(
                Constants.userPrivateKey,
                messageDigest2
            );
            signatures[1] = Sig.Components(v2, r12, s2);
        }

        swivel.exit(o, a, signatures);

        (uint256 notionalBalance, ) = vaultTracker.balancesOf(
            Constants.userPublicKey
        );
        uint256 transactionSize = (amount * principal) / premium;
        transactionSize += (amount2 * principal2) / premium2;
        assertEq(notionalBalance, startingBalance - transactionSize);
        assertGt(
            IERC20(Constants.STETH).balanceOf(address(this)),
            startingBalance
        );
    }

    function dealStakedEth(address p, uint256 a) internal {
        deal(wstETH, p, a * 2);
        vm.startPrank(p);
        IWrappedSteth(wstETH).unwrap(a);
        vm.stopPrank();
    }
}

interface IWrappedSteth {
    function unwrap(uint256 a) external returns (uint256);
}
