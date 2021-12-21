package main

import (
	"fmt"
	"log"
	// "math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/swivel-finance/swivel/deploy/internal/destributor"
)

func deployDestributor(a *bind.TransactOpts, c *ethclient.Client) {
	fmt.Println("Deploying Destributor...")

	// current swiv token...
	token := common.HexToAddress("0xbf30461210b37012783957D90dC26B95Ce3b6f2d")
	// initial merkle root as empty...
	byteRoot := common.FromHex("0x0000000000000000000000000000000000000000000000000000000000000000")
	root := [32]byte{}
	copy(root[:], byteRoot)

	destributorAddr, tx, _, err := destributor.DeployDestributor(a, c, token, root)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deployed to: %v\n", destributorAddr.Hex())
	fmt.Printf("Transaction hash: %v\n", tx.Hash().Hex())
}
