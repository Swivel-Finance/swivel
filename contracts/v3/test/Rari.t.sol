// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.16;

import 'forge-std/Test.sol';

import 'src/Swivel.sol';
import 'src/MarketPlace.sol' as mp;
import 'src/Creator.sol' as c;
import 'src/VaultTracker.sol' as vt;

import 'test/Constants.sol';
import 'test/TestHash.sol' as th;

contract RariTest is Test {
    c.Creator creator;
    mp.MarketPlace marketplace;
    Swivel swivel;
    IERC20 zcToken;
    vt.VaultTracker vaultTracker;

    address cUSDC = 0x39AA39c021dfbaE8faC545936693aC917d5E7563;
    uint256 maturity = 1672549200; // Jan 1, 2023 | 00:00:00 GMT +0
    uint256 startingBalance = 10000000;

    // common order info
    bytes32 key = bytes32('hello');
    uint256 principal = 1000;
    uint256 premium = 100;

    // test order
    Hash.Order order =
        Hash.Order(
            key,
            uint8(Protocols.Rari),
            Constants.userPublicKey,
            Constants.USDC,
            false,
            false,
            principal,
            premium,
            maturity,
            maturity
        );

    function setUp() public {
        // Deploy swivel contracts
        creator = new c.Creator();
        marketplace = new mp.MarketPlace(address(creator));
        swivel = new Swivel(address(marketplace), address(0));

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
            uint8(Protocols.Rari),
            maturity,
            cUSDC,
            'TEST-Constants.USDC',
            'zcUSD'
        );

        address[] memory underlyings = new address[](1);
        address[] memory compoundings = new address[](1);
        underlyings[0] = Constants.USDC;
        compoundings[0] = cUSDC;

        (, address zcTokenAddr, address vaultTrackerAddr, ) = marketplace
            .markets(uint8(Protocols.Rari), Constants.USDC, maturity);
        zcToken = IERC20(zcTokenAddr);
        vaultTracker = vt.VaultTracker(vaultTrackerAddr);
    }

    function testInitiateVaultFillingZcTokenInitiate() public {
        deployMarket();
        order.vault = false;
        order.exit = false;
        Hash.Order[] memory o = new Hash.Order[](1);
        o[0] = order;

        uint256[] memory a = new uint256[](1);
        uint256 amount = 50;
        a[0] = amount;

        bytes32 messageDigest = Hash.message(
            swivel.domain(),
            th.TestHash.order(order)
        );

        Sig.Components[] memory signatures = new Sig.Components[](1);
        (uint8 v, bytes32 r1, bytes32 s) = vm.sign(
            Constants.userPrivateKey,
            messageDigest
        );
        signatures[0] = Sig.Components(v, r1, s);

        swivel.initiate(o, a, signatures);

        uint256 transactionSize = (amount * principal) / premium;
        assertEq(zcToken.balanceOf(Constants.userPublicKey), transactionSize);
        (uint256 notionalBalance, ) = vaultTracker.balancesOf(address(this));
        uint256 fee = transactionSize / swivel.feenominators(2);
        assertEq(notionalBalance, transactionSize - fee);
    }

    function testInitiateZcTokenFillingVaultInitiate() public {
        deployMarket();
        order.vault = true;
        order.exit = false;
        Hash.Order[] memory o = new Hash.Order[](1);
        o[0] = order;

        uint256[] memory a = new uint256[](1);
        uint256 amount = 50;
        a[0] = amount;

        bytes32 messageDigest = Hash.message(
            swivel.domain(),
            th.TestHash.order(order)
        );

        Sig.Components[] memory signatures = new Sig.Components[](1);
        (uint8 v, bytes32 r1, bytes32 s) = vm.sign(
            Constants.userPrivateKey,
            messageDigest
        );
        signatures[0] = Sig.Components(v, r1, s);

        swivel.initiate(o, a, signatures);

        assertEq(zcToken.balanceOf(address(this)), amount);
        (uint256 notionalBalance, ) = vaultTracker.balancesOf(
            Constants.userPublicKey
        );
        assertEq(notionalBalance, amount);
    }

    function testInitiateZcTokenFillingZcTokenExit() public {
        deployMarket();
        uint256 amount = 50;
        deal(address(zcToken), Constants.userPublicKey, startingBalance, true);

        order.vault = false;
        order.exit = true;
        Hash.Order[] memory o = new Hash.Order[](1);
        o[0] = order;

        uint256[] memory a = new uint256[](1);
        a[0] = amount;

        bytes32 messageDigest = Hash.message(
            swivel.domain(),
            th.TestHash.order(order)
        );

        Sig.Components[] memory signatures = new Sig.Components[](1);
        (uint8 v, bytes32 r1, bytes32 s) = vm.sign(
            Constants.userPrivateKey,
            messageDigest
        );
        signatures[0] = Sig.Components(v, r1, s);

        swivel.initiate(o, a, signatures);

        assertEq(
            zcToken.balanceOf(Constants.userPublicKey),
            startingBalance - amount
        );
        assertEq(zcToken.balanceOf(address(this)), amount);
    }

    function testInitiateVaultFillingVaultExit() public {
        deployMarket();
        uint256 amount = 50;
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
            th.TestHash.order(order)
        );

        Sig.Components[] memory signatures = new Sig.Components[](1);
        (uint8 v, bytes32 r1, bytes32 s) = vm.sign(
            Constants.userPrivateKey,
            messageDigest
        );
        signatures[0] = Sig.Components(v, r1, s);

        swivel.initiate(o, a, signatures);

        (uint256 notionalBalance, ) = vaultTracker.balancesOf(
            Constants.userPublicKey
        );
        assertEq(notionalBalance, 0);
        (notionalBalance, ) = vaultTracker.balancesOf(address(this));
        uint256 fee = transactionSize / swivel.feenominators(2);
        assertEq(notionalBalance, transactionSize - fee);
    }

    function testExitZcTokenFillingZcTokenInitiate() public {
        deployMarket();
        deal(cUSDC, address(swivel), startingBalance);
        deal(address(zcToken), address(this), startingBalance);
        uint256 amount = 50;
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
            th.TestHash.order(order)
        );

        Sig.Components[] memory signatures = new Sig.Components[](1);
        (uint8 v, bytes32 r1, bytes32 s) = vm.sign(
            Constants.userPrivateKey,
            messageDigest
        );
        signatures[0] = Sig.Components(v, r1, s);

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

    function testExitVaultFillingVaultInitiate() public {
        deployMarket();
        deal(cUSDC, address(swivel), startingBalance);
        uint256 amount = 50;
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
            th.TestHash.order(order)
        );

        Sig.Components[] memory signatures = new Sig.Components[](1);
        (uint8 v, bytes32 r1, bytes32 s) = vm.sign(
            Constants.userPrivateKey,
            messageDigest
        );
        signatures[0] = Sig.Components(v, r1, s);

        swivel.exit(o, a, signatures);

        (uint256 notionalBalance, ) = vaultTracker.balancesOf(
            Constants.userPublicKey
        );
        assertEq(notionalBalance, amount);
        assertGt(
            IERC20(Constants.USDC).balanceOf(address(this)),
            startingBalance
        );
    }

    function testExitVaultFillingZcTokenExit() public {
        deployMarket();
        deal(cUSDC, address(swivel), startingBalance);
        uint256 amount = 50;
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
            th.TestHash.order(order)
        );

        Sig.Components[] memory signatures = new Sig.Components[](1);
        (uint8 v, bytes32 r1, bytes32 s) = vm.sign(
            Constants.userPrivateKey,
            messageDigest
        );
        signatures[0] = Sig.Components(v, r1, s);

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

    function testExitZcTokenFillingVaultExit() public {
        deployMarket();
        deal(cUSDC, address(swivel), startingBalance);
        deal(address(zcToken), address(this), startingBalance);
        uint256 amount = 50;
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
            th.TestHash.order(order)
        );

        Sig.Components[] memory signatures = new Sig.Components[](1);
        (uint8 v, bytes32 r1, bytes32 s) = vm.sign(
            Constants.userPrivateKey,
            messageDigest
        );
        signatures[0] = Sig.Components(v, r1, s);

        swivel.exit(o, a, signatures);

        (uint256 notionalBalance, ) = vaultTracker.balancesOf(
            Constants.userPublicKey
        );
        uint256 transactionSize = (amount * principal) / premium;
        assertEq(notionalBalance, startingBalance - transactionSize);
        assertGt(
            IERC20(Constants.USDC).balanceOf(address(this)),
            startingBalance
        );
    }
}
