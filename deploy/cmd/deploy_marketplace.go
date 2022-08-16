package main

import (
	"fmt"
	"log"

	// "math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/swivel-finance/swivel/deploy/internal/marketplace"
)

// addr is the deployed creator address
func deployMarketplace(a *bind.TransactOpts, c *ethclient.Client, addr common.Address) {
	// deploy marketplace first...
	fmt.Println("Deploying MarketPlace...")

	marketAddr, tx, _, err := marketplace.DeployMarketPlace(a, c, addr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deployed to: %v\n", marketAddr.Hex())
	fmt.Printf("Transaction hash: %v\n", tx.Hash().Hex())
}
