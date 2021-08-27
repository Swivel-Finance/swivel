package main

import (
	"fmt"
	"log"
	"math/big"
	"time"

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

	underlying := common.HexToAddress("0x5592ec0cfb4dbc12d3ab100b257153436a1f0fea")
	cTokenAddr := common.HexToAddress("0x6d7f0754ffeb405d23c51ce938289d4835be3b14")
	now := big.NewInt(time.Now().Unix())
	// year := big.NewInt(31536000)
	sixMo := big.NewInt(15780000)
	maturity := now.Add(now, sixMo)
	name := "Test Market 3"
	symbol := "TM3"

	tx, err := marketCont.CreateMarket(opts, underlying, maturity, cTokenAddr, name, symbol, d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transaction hash: %v\n", tx.Hash().Hex())
}
