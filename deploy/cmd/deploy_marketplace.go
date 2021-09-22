package main

import (
	"fmt"
	"log"
	// "math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/swivel-finance/swivel/deploy/internal/marketplace"
)

func deployMarketplace(a *bind.TransactOpts, c *ethclient.Client) {
	// deploy marketplace first...
	fmt.Println("Deploying MarketPlace...")

	marketAddr, tx, _, err := marketplace.DeployMarketPlace(a, c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deployed to: %v\n", marketAddr.Hex())
	fmt.Printf("Transaction hash: %v\n", tx.Hash().Hex())
}
