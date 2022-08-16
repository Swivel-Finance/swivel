package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/swivel-finance/swivel/deploy/internal/creator"
)

func setMarketPlaceAddress(a *bind.TransactOpts, c *ethclient.Client, cr common.Address, m common.Address) {
	opts := &bind.TransactOpts{
		From:   a.From,
		Signer: a.Signer,
	}

	fmt.Println("Setting MarketPlace address in Creator...")

	// get the deployed creator...
	creatorCont, err := creator.NewCreator(cr, c)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := creatorCont.SetMarketPlace(opts, m)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transaction hash: %v\n", tx.Hash().Hex())
}
