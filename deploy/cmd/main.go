package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	chainId := big.NewInt(421611)

	// whichever fully qualified url your probject uses to establish connection to your node...
	client, err := ethclient.Dial(os.Getenv("CLIENT_URL"))

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

	fmt.Printf("Current Nonce: %v\n", nonce)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Gas estimate: %v\n", gasPrice)
	// fmt.Println("Attempting to use 70 gwei as gas price")

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)          // in wei
	auth.GasLimit = uint64(80000000000) // let's see if 8 mill will go...
	//auth.GasPrice = big.NewInt(10000000) // if u wanna just hardcode it - use gwei
	auth.GasPrice = gasPrice // let geth estimate
	fmt.Printf("Using gas price: %v\n", gasPrice)

	fmt.Printf("Transaction options: %v\n", auth)

	// if transferring admin, address here...
	// admin := common.HexToAddress(os.Getenv("ADMIN"))

	/*
		We simply turn these steps on and off by commenting them.
		TODO we _could_ automate it by waiting for receipts etc...

		1. deploy marketplace
		   a. update marketplace address var
		2. deploy swivel
		   a. update swivel address var
		3. set swivel address in deployed marketplace
		4. create any desired markets
		5. deploy a destributor if desired
		6. Transfer admin in contracts if desired
		7. Deploy a PErc20 token if desired
	*/

	// TODO we dont return the address here as we don't try to chain them atm
	// deployMarketplace(auth, client)
	// marketAddr := common.HexToAddress("0xFFA3fcAB9ed73a526C7621dB73615AaD1e49265B")

	// deploySwivel(auth, client, marketAddr)
	// swivelAddr := common.HexToAddress("0x3C01fb861501428cDc6F067461E8866b0542FAbE")

	//setSwivelAddress(auth, client, marketAddr, swivelAddr)

	// createMarket(auth, client, marketAddr)

	// deployDestributor(auth, client)
	// destributorAddr := common.HexToAddress("0x57E18D9F50F3Fd0894c8436BC84D2f523A8d0968")

	// transferAdminMarketplace(auth, client, marketAddr, admin)

	// transferAdminSwivel(auth, client, swivelAddr, admin)

	// transferAdminDestributor(auth, client, destributorAddr, admin)

	name := "Arbitrum Test 2"
	symbol := "AT2"
	decimals := uint8(18)
	deployPErc20(auth, client, name, symbol, decimals)
}
