package main

import (
	"fmt"
	"log"

	// "math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/swivel-finance/swivel/deploy/internal/mocks"
)

func deployFErc20(a *bind.TransactOpts, c *ethclient.Client) {
	fmt.Println("Deploying FToken...")

	ftokenAddr, tx, _, err := mocks.DeployFErc20(a, c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deployed to: %v\n", ftokenAddr.Hex())
	fmt.Printf("Transaction hash: %v\n", tx.Hash().Hex())
}
