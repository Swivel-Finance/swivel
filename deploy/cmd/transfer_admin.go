package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/swivel-finance/swivel/deploy/internal/destributor"
	"github.com/swivel-finance/swivel/deploy/internal/marketplace"
	"github.com/swivel-finance/swivel/deploy/internal/swivel"
)

// @param m Address of the deployed Marketplace contract
// @param p Address (public key) of the new admin
func transferAdminMarketplace(a *bind.TransactOpts, c *ethclient.Client, m common.Address, p common.Address) {
	opts := &bind.TransactOpts{
		From:   a.From,
		Signer: a.Signer,
	}

	fmt.Println("Transferring admin in Marketplace...")

	// get the deployed marketplace...
	marketCont, err := marketplace.NewMarketPlace(m, c)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := marketCont.TransferAdmin(opts, p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transaction hash: %v\n", tx.Hash().Hex())
}

// @param s Address of the deployed Swivel contract
// @param p Address (public key) of the new admin
func transferAdminSwivel(a *bind.TransactOpts, c *ethclient.Client, s common.Address, p common.Address) {
	opts := &bind.TransactOpts{
		From:   a.From,
		Signer: a.Signer,
	}

	fmt.Println("Transferring admin in Swivel...")

	// get the deployed swivel...
	swivelCont, err := swivel.NewSwivel(s, c)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := swivelCont.TransferAdmin(opts, p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transaction hash: %v\n", tx.Hash().Hex())
}

// @param d Address of the deployed Destributor contract
// @param p Address (public key) of the new admin
func transferAdminDestributor(a *bind.TransactOpts, c *ethclient.Client, d common.Address, p common.Address) {
	opts := &bind.TransactOpts{
		From:   a.From,
		Signer: a.Signer,
	}

	fmt.Println("Transferring admin in Destributor...")

	// get the deployed contract...
	destributorCont, err := destributor.NewDestributor(d, c)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := destributorCont.TransferAdmin(opts, p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transaction hash: %v\n", tx.Hash().Hex())
}
