package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/swivel-finance/swivel/deploy/internal/marketplace"
	"github.com/swivel-finance/swivel/deploy/internal/swivel"
)

func querySwivelPublicFields(c *ethclient.Client, s common.Address) {

	fmt.Println("querying swivel public fields...")

	// get the deployed swivel...
	swivelCont, err := swivel.NewSwivel(s, c)
	if err != nil {
		log.Fatal(err)
	}

	name, err := swivelCont.NAME(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name: %v\n", name)

	verz, err := swivelCont.VERSION(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("version: %v\n", verz)

	hold, err := swivelCont.HOLD(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("holding period: %v\n", hold)

	domain, err := swivelCont.Domain(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("EIP 7xx domain separator: %v\n", hex.EncodeToString(domain[:]))

	mPlace, err := swivelCont.MarketPlace(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("marketplace address: %v\n", mPlace.Hex())
}

func queryMarketPlacePublicFields(c *ethclient.Client, m common.Address) {

	fmt.Println("querying market place public fields...")

	// get the deployed market place...
	marketCont, err := marketplace.NewMarketPlace(m, c)
	if err != nil {
		log.Fatal(err)
	}

	swivel, err := marketCont.Swivel(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("swivel address: %v\n", swivel.Hex())
}
