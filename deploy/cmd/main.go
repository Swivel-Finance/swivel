package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/swivel-finance/swivel/deploy/internal/marketplace"
	// "github.com/swivel-finance/swivel/deploy/internal/swivel"
)

func main() {
	// rinkeby
	client, err := ethclient.Dial(fmt.Sprintf("https://rinkeby.infura.io/v3/%s", os.Getenv("INFURA_PROJECT_ID")))

	// goerli
	// client, err := ethclient.Dial(fmt.Sprintf("https://goerli.infura.io/v3/%s", os.Getenv("INFURA_PROJECT_ID")))

	// kovan
	// client, err := ethclient.Dial(fmt.Sprintf("https://kovan.infura.io/v3/%s", os.Getenv("INFURA_PROJECT_ID")))

	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(8000000) // let's see if 8 mill will go...
	// auth.GasPrice = gasPrice.Mul(gasPrice, big.NewInt(2))
	auth.GasPrice = gasPrice

	// deploy marketplace first...
	// fmt.Println("Deploying MarketPlace...")

	// marketAddr, tx, _, err := marketplace.DeployMarketPlace(auth, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Deployed to: %v\n", marketAddr.Hex())
	// fmt.Printf("Transaction hash: %v\n", tx.Hash().Hex())

	// now swivel with market address... (update the nonce first)
	// nonce, err = client.PendingNonceAt(context.Background(), fromAddress)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// auth.Nonce = big.NewInt(int64(nonce))

	// fmt.Println("Deploying Swivel...")

	marketAddr := common.HexToAddress("0xc078Ae76aEdf9150359ef5470A5C0b51a508bC79")

	// swivelAddr, tx, _, err := swivel.DeploySwivel(auth, client, marketAddr)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Deployed to: %v\n", swivelAddr.Hex())
	// fmt.Printf("Transaction hash: %v\n", tx.Hash().Hex())

	// set the swivel address in the marketplace (wait and confirm the above blocks have mined)
	opts := &bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
	}

	fmt.Println("Setting Swivel address in Marketplace...")

	swivelAddr := "0x5FC6f550f2965F1AB5ACD8eA145518ddC40deE57"

	// get the deployed marketplace...
	marketCont, err := marketplace.NewMarketPlace(marketAddr, client)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := marketCont.SetSwivelAddress(opts, common.HexToAddress(swivelAddr))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transaction hash: %v\n", tx.Hash().Hex())
}
