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
	chainId := big.NewInt(4)

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
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(8000000) // let's see if 8 mill will go...
	// auth.GasPrice = big.NewInt(60000000000) // if u wanna just hardcode it - use gwei
	auth.GasPrice = gasPrice // let geth estimate

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
		7. delploy a mock rari token
	*/

	// TODO we dont return the address here as we don't try to chain them atm
	// deployMarketplace(auth, client)
	// marketAddr := common.HexToAddress("0xb029e31C252549B20a8dc502c2832F11A2EDE417")

	// deploySwivel(auth, client, marketAddr)
	// swivelAddr := common.HexToAddress("0x3a09584FF42CDFe27Fe72Da0533bba24E9C28AaD")

	// setSwivelAddress(auth, client, marketAddr, swivelAddr)

	// createMarket(auth, client, marketAddr)

	// deployDestributor(auth, client)
	// destributorAddr := common.HexToAddress("")

	// transferAdminMarketplace(auth, client, marketAddr, admin)

	// transferAdminSwivel(auth, client, swivelAddr, admin)

	// transferAdminDestributor(auth, client, destributorAddr, admin)

	// deployFErc20(auth, client)
}
