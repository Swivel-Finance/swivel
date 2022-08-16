package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/swivel-finance/swivel/deploy/internal/marketplace"
)

func createMarket(a *bind.TransactOpts, c *ethclient.Client, m common.Address) {
	opts := &bind.TransactOpts{
		From:   a.From,
		Signer: a.Signer,
	}

	fmt.Println("Creating Market...")

	// get the deployed marketplace...
	marketCont, err := marketplace.NewMarketPlace(m, c)
	if err != nil {
		log.Fatal(err)
	}

	// ******************* RINKEBY DAI **********************************************
	// cTokenAddr := common.HexToAddress("0x6d7f0754ffeb405d23c51ce938289d4835be3b14")
	// maturity - 9/1/22
	// maturity := big.NewInt(1662089767)
	// name := "DAI-1662089767"
	// symbol := "zcDAI"

	// ******************* RINKEBY USDC *********************************************
	// cTokenAddr := common.HexToAddress("0x5b281a6dda0b271e91ae35de655ad301c976edb1")
	// maturity - <>
	// maturity := big.NewInt(1633059407)
	// name := "USDC-1633059407"
	// symbol := "zcUSDC"

	// ******************* KOVAN DAI **********************************************
	// cTokenAddr := common.HexToAddress("0xf0d0eb522cfa50b716b3b1604c4f0fa6f04376ad")
	// maturity := big.NewInt(1671157528) // thurs dec 15 22
	// name := "zcDAI-1671157528"
	// symbol := "zcDAI-Dec"

	// ******************* KOVAN USDC **********************************************
	proto := uint8(0)
	cTokenAddr := common.HexToAddress("0x4a92e71227d294f041bd82dd8f78591b75140d63")
	maturity := big.NewInt(1642299928) // sat jan 15 22
	name := "zcUSDC-1642299928"
	symbol := "zcUSDC-Jan"

	tx, err := marketCont.CreateMarket(opts, proto, maturity, cTokenAddr, name, symbol)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transaction hash: %v\n", tx.Hash().Hex())
}
