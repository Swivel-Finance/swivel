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

func createMarket(a *bind.TransactOpts, c *ethclient.Client, m common.Address, d uint8) {
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

	// ******************* DAI **********************************************
	underlying := common.HexToAddress("0x5592ec0cfb4dbc12d3ab100b257153436a1f0fea")
	cTokenAddr := common.HexToAddress("0x6d7f0754ffeb405d23c51ce938289d4835be3b14")
	// maturity - 9/1/22
	maturity := big.NewInt(1662089767)
	name := "DAI-1662089767"
	symbol := "zcDAI"

	// ******************* USDC *********************************************
	// underlying := common.HexToAddress("0x4dbcdf9b62e891a7cec5a2568c3f4faf9e8abe2b")
	// cTokenAddr := common.HexToAddress("0x5b281a6dda0b271e91ae35de655ad301c976edb1")
	// maturity - 9/31/21
	// maturity := big.NewInt(1633059407)
	// name := "USDC-1633059407"
	// symbol := "zcUSDC"

	tx, err := marketCont.CreateMarket(opts, underlying, maturity, cTokenAddr, name, symbol, d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transaction hash: %v\n", tx.Hash().Hex())
}
