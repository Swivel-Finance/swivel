package main

import (
	"fmt"
	"log"

	// "math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/swivel-finance/swivel/deploy/internal/creator"
)

func deployCreator(a *bind.TransactOpts, c *ethclient.Client) {
	// deploy creator first...
	fmt.Println("Deploying Creator...")

	creatorAddr, tx, _, err := creator.DeployCreator(a, c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deployed to: %v\n", creatorAddr.Hex())
	fmt.Printf("Transaction hash: %v\n", tx.Hash().Hex())
}
