package main

import (
	"fmt"
	"log"

	// "math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/swivel-finance/swivel/deploy/internal/tokens"
)

func deployPErc20(a *bind.TransactOpts, c *ethclient.Client, n string, s string, d uint8) {
	fmt.Println("Deploying PErc20...")

	PErc20Addr, tx, _, err := tokens.DeployPErc20(a, c, n, s, d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deployed to: %v\n", PErc20Addr.Hex())
	fmt.Printf("Transaction hash: %v\n", tx.Hash().Hex())
}
