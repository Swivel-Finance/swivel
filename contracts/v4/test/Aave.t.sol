// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.16;
import 'forge-std/Test.sol';
import {Swivel} from 'src/Swivel.sol';
import {MarketPlace} from 'src/MarketPlace.sol';
import {Creator} from 'src/Creator.sol';
import {VaultTracker} from 'src/VaultTracker.sol';
import {Aave} from 'src/adapters/Aave.sol';
import {ZcToken} from 'src/tokens/ZcToken.sol';
import {IERC20} from 'src/interfaces/IERC20.sol';
import {IAave} from 'src/interfaces/IAave.sol';
import {Sig} from 'src/lib/Sig.sol';
import {Hash} from 'src/lib/Hash.sol';
import {Constants} from 'test/Constants.sol';
import {TestHash} from 'test/TestHash.sol';
import {Protocols} from 'test/Protocols.sol';
contract AaveTest is Test {
    Creator creator;
    MarketPlace marketplace;
    Swivel swivel;
    ZcToken zcToken;
    VaultTracker vaultTracker;
    Aave adapter;
    address aUSDC = 0xBcca60bB61934080951369a648Fb03DF4F96263C;
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
            uint8(Protocols.Aave),
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
            uint8(Protocols.Aave),
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
        adapter = new Aave();
        // Approve compound to spend from swivel
        vm.startPrank(address(swivel));
        IERC20(Constants.USDC).approve(aUSDC, type(uint256).max);
        IERC20(Constants.USDC).approve(
            0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9,
            type(uint256).max
        );
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
            uint8(Protocols.Aave),
            maturity,
            aUSDC,
            address(adapter),
            'TEST-Constants.USDC',
            'zcUSD'
        );
        address[] memory underlyings = new address[](1);
        address[] memory compoundings = new address[](1);
        underlyings[0] = Constants.USDC;
        compoundings[0] = aUSDC;
        (, , address zcTokenAddr, address vaultTrackerAddr, ) = marketplace
            .markets(uint8(Protocols.Aave), Constants.USDC, maturity);
        zcToken = ZcToken(zcTokenAddr);
        vaultTracker = VaultTracker(vaultTrackerAddr);
    }
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
        swivel.initiate(o, a, signatures);
        uint256 transactionSize = (amount * principal) / premium;
        uint256 fee = transactionSize / swivel.feenominators(2);
        // maker balances
        assertEq(zcToken.balanceOf(Constants.userPublicKey), transactionSize);
        assertEq(
            IERC20(Constants.USDC).balanceOf(Constants.userPublicKey),
            startingBalance + amount - transactionSize
        );
        // taker balances
        (uint256 notionalBalance, ) = vaultTracker.balancesOf(address(this));
        assertEq(notionalBalance, transactionSize - fee);
        assertEq(
            IERC20(Constants.USDC).balanceOf(address(this)),
            startingBalance - amount
        );
    }
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
        swivel.initiate(o, a, signatures);
        uint256 transactionSize = (amount * premium) / principal;
        uint256 fee = transactionSize / swivel.feenominators(0);
        // maker balances
        assertEq(
            IERC20(Constants.USDC).balanceOf(Constants.userPublicKey),
            startingBalance - transactionSize
        );
        (uint256 notionalBalance, ) = vaultTracker.balancesOf(
            Constants.userPublicKey
        );
        assertEq(notionalBalance, amount);
        // taker balances
        assertEq(zcToken.balanceOf(address(this)), amount);
        assertEq(
            IERC20(Constants.USDC).balanceOf(address(this)),
            startingBalance - (amount + fee) + transactionSize
        );
    }
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
        swivel.initiate(o, a, signatures);
        uint256 transactionSize = (amount * premium) / principal;
        uint256 fee = transactionSize / swivel.feenominators(0);
        // maker balances
        assertEq(
            zcToken.balanceOf(Constants.userPublicKey),
            startingBalance - amount
        );
        assertEq(
            IERC20(Constants.USDC).balanceOf(Constants.userPublicKey),
            startingBalance + (amount - transactionSize)
        );
        // taker balances
        assertEq(zcToken.balanceOf(address(this)), amount);
        assertEq(
            IERC20(Constants.USDC).balanceOf(address(this)),
            startingBalance - (amount - transactionSize) - fee
        );
    }
    function testInitiateVaultFillingVaultExit() public {
        deployMarket();
        uint256 amount = 50e6;
        uint256 transactionSize = (amount * principal) / premium;
        deal(address(zcToken), Constants.userPublicKey, startingBalance, true);
        vm.startPrank(address(marketplace));
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
        swivel.initiate(o, a, signatures);
        // make balances
        (uint256 notionalBalance, ) = vaultTracker.balancesOf(
            Constants.userPublicKey
        );
        assertEq(notionalBalance, 0);
        assertEq(
            IERC20(Constants.USDC).balanceOf(Constants.userPublicKey),
            startingBalance + amount
        );
        // taker balances
        (notionalBalance, ) = vaultTracker.balancesOf(address(this));
        uint256 fee = transactionSize / swivel.feenominators(2);
        assertEq(notionalBalance, transactionSize - fee);
        assertEq(
            IERC20(Constants.USDC).balanceOf(address(this)),
            startingBalance - amount
        );
    }
    function testExitZcTokenFillingZcTokenInitiate() public {
        deployMarket();
        aaveDeal(address(swivel), startingBalance);
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
        swivel.exit(o, a, signatures);
        uint256 transactionSize = (amount * principal) / premium;
        uint256 fee = transactionSize / swivel.feenominators(1);
        // maker balances
        assertEq(
            zcToken.balanceOf(Constants.userPublicKey),
            startingBalance + transactionSize
        );
        // taker balances
        assertEq(
            IERC20(Constants.USDC).balanceOf(address(this)),
            startingBalance + (transactionSize - amount) - fee
        );
        assertEq(
            zcToken.balanceOf(address(this)),
            startingBalance - transactionSize
        );
    }
    function testExitVaultFillingVaultInitiate() public {
        deployMarket();
        aaveDeal(address(swivel), startingBalance);
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
        swivel.exit(o, a, signatures);
        uint256 transactionSize = (amount * premium) / principal;
        uint256 fee = transactionSize / swivel.feenominators(3);
        // maker balances
        assertEq(
            IERC20(Constants.USDC).balanceOf(Constants.userPublicKey),
            startingBalance - transactionSize
        );
        (uint256 notionalBalance, ) = vaultTracker.balancesOf(
            Constants.userPublicKey
        );
        assertEq(notionalBalance, amount);
        // taker balances
        (notionalBalance, ) = vaultTracker.balancesOf(address(this));
        assertEq(notionalBalance, 0);
        assertEq(
            IERC20(Constants.USDC).balanceOf(address(this)),
            startingBalance + transactionSize - fee
        );
    }
    function testExitVaultFillingZcTokenExit() public {
        deployMarket();
        aaveDeal(address(swivel), startingBalance);
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
        swivel.exit(o, a, signatures);
        uint256 transactionSize = (amount * premium) / principal;
        uint256 fee = transactionSize / swivel.feenominators(3);
        // taker balances
        assertEq(
            IERC20(Constants.USDC).balanceOf(Constants.userPublicKey),
            startingBalance + (amount - transactionSize)
        );
        assertEq(
            zcToken.balanceOf(Constants.userPublicKey),
            startingBalance - amount
        );
        // maker balances
        assertEq(
            IERC20(Constants.USDC).balanceOf(address(this)),
            startingBalance + (transactionSize - fee)
        );
        (uint256 notionalBalance, ) = vaultTracker.balancesOf(address(this));
        assertEq(notionalBalance, 0);
    }
    function testExitZcTokenFillingVaultExit() public {
        deployMarket();
        aaveDeal(address(swivel), startingBalance);
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
        swivel.exit(o, a, signatures);
        uint256 transactionSize = (amount * principal) / premium;
        uint256 fee = transactionSize / swivel.feenominators(1);
        // maker balances
        (uint256 notionalBalance, ) = vaultTracker.balancesOf(
            Constants.userPublicKey
        );
        assertEq(notionalBalance, startingBalance - transactionSize);
        assertEq(
            IERC20(Constants.USDC).balanceOf(Constants.userPublicKey),
            startingBalance + amount
        );
        // taker balances
        assertEq(
            IERC20(Constants.USDC).balanceOf(address(this)),
            startingBalance + (transactionSize - amount - fee)
        );
        assertEq(
            zcToken.balanceOf(address(this)),
            startingBalance - transactionSize
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
        aaveDeal(address(swivel), startingBalance);
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
            IERC20(Constants.USDC).balanceOf(address(this)),
            startingBalance
        );
    }
    function testExitVaultFillingVaultInitiateMultiple() public {
        deployMarket();
        aaveDeal(address(swivel), startingBalance);
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
            IERC20(Constants.USDC).balanceOf(address(this)),
            startingBalance
        );
    }
    function testExitVaultFillingZcTokenExitMultiple() public {
        deployMarket();
        aaveDeal(address(swivel), startingBalance);
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
            IERC20(Constants.USDC).balanceOf(address(this)),
            startingBalance
        );
    }
    function testExitZcTokenFillingVaultExitMultiple() public {
        deployMarket();
        aaveDeal(address(swivel), startingBalance);
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
            IERC20(Constants.USDC).balanceOf(address(this)),
            startingBalance
        );
    }
    /// @dev implicitly deals aUSDC to recipient by calling deposit on Aave
    /// @param r intended recipient address
    /// @param a amount to deposit to the recipient
    function aaveDeal(address r, uint256 a) internal {
        uint256 starting = IERC20(Constants.USDC).balanceOf(r);
        deal(Constants.USDC, r, starting + a);
        vm.startPrank(r);
        IAave(0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9).deposit(
            Constants.USDC,
            a,
            r,
            0
        );
        vm.stopPrank();
    }
}