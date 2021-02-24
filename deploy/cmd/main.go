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
	"github.com/swivel-finance/swivel/deploy/internal/swivel"
)

func main() {
	client, err := ethclient.Dial(fmt.Sprintf("https://goerli.infura.io/v3/%s", os.Getenv("INFURA_PROJECT_ID")))
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
	auth.GasLimit = uint64(7500000) // in units
	auth.GasPrice = gasPrice

	// goerli
	chainId := big.NewInt(5)
	// goerli
	cDai := common.HexToAddress("0x822397d9a55d0fefd20f5c4bcab33c5f65bd28eb")
	// verifier, allowing to default to `this` by passing "zero address"
	zero := common.HexToAddress("0x0")
	address, tx, _, err := swivel.DeploySwivel(auth, client, chainId, cDai, zero)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deployed to: %v\n", address.Hex())
	fmt.Printf("Transaction hash: %v\n", tx.Hash().Hex())
}
