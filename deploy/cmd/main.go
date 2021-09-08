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

	/*
		We simply turn these steps on and off by commenting them.
		TODO we _could_ automate it by waiting for receipts etc...

		1. deploy marketplace
		   a. update marketplace address var
		2. deploy swivel
		   a. update swivel address var
		3. set swivel address in deployed marketplace
		4. create any desired markets
	*/

	// TODO we dont return the address here as we don't try to chain them atm
	// deployMarketplace(auth, client)
	marketAddr := common.HexToAddress("0xeb389e2796E081FBb5C032F3D307CD5E6b438D78")

	// deploySwivel(auth, client, marketAddr)
	// swivelAddr := common.HexToAddress("0xDe9a819630094dF6dA6FF7CCc77E04Fd3ad0ACFE")

	// setSwivelAddress(auth, client, marketAddr, swivelAddr)

	// NOTE be sure to set the correct number of decimals for the market (zctoken) you are creating (18 dai, 6 usdc)
	createMarket(auth, client, marketAddr, uint8(6))
}
