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

	token := common.HexToAddress("0x46378144d3f9a2a61470e10cbddcd4d933b96f94")
	byteRoot := common.FromHex("0x22f89e6494aca95b9e00cd0897af1784c2cddd80e092d0fa5d54cb4297894002")
	root := [32]byte{}
	copy(root[:], byteRoot)

	destributorAddr, tx, _, err := destributor.DeployDestributor(a, c, token, root)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deployed to: %v\n", destributorAddr.Hex())
	fmt.Printf("Transaction hash: %v\n", tx.Hash().Hex())
}
