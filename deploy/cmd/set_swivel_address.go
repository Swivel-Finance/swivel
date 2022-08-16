package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/swivel-finance/swivel/deploy/internal/marketplace"
)

func setSwivelAddress(a *bind.TransactOpts, c *ethclient.Client, m common.Address, s common.Address) {
	opts := &bind.TransactOpts{
		From:   a.From,
		Signer: a.Signer,
	}

	fmt.Println("Setting Swivel address in Marketplace...")

	// get the deployed marketplace...
	marketCont, err := marketplace.NewMarketPlace(m, c)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := marketCont.SetSwivel(opts, s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transaction hash: %v\n", tx.Hash().Hex())
}
