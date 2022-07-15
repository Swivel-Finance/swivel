package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/swivel-finance/swivel/deploy/internal/swivel"
)

func deploySwivel(a *bind.TransactOpts, c *ethclient.Client, m common.Address, v common.Address) {
	fmt.Println("Deploying Swivel...")

	swivelAddr, tx, _, err := swivel.DeploySwivel(a, c, m, v)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deployed to: %v\n", swivelAddr.Hex())
	fmt.Printf("Transaction hash: %v\n", tx.Hash().Hex())
}
