# Swivel

<img src="https://user-images.githubusercontent.com/62613746/111923535-1cddbd00-8a76-11eb-80fa-853acfc789e3.png" width="400px">

## contracts
Swivel smart contracts are located here, stored by version. Their associated `abi` and `bin` compilation artifacts are also here for convenience.

### v2
The currently deployed *v2* protocol is located here:
* swivel/
  * The `.sol`, `.abi` and `.bin` files for Swivel.sol and its dependencies
* marketplace/
  * The `.sol`, `.abi` and `.bin` files for MarketPlace.sol and its dependencies

#### deployment
We use the Geth ABIGEN tool to generate golang bindings for our smart contracts. We use those bindings to deploy the contracts with the
`deploy` project in this repo. see `/deploy/cmd/main.go` for the script itself or `/deploy/internal/*` to view the bindings themselves.

## current deployments
* Rinkeby:
  * MarketPlace: 0x9fa54f942D8b8e992501952C3e6E67F1A42595b8
  * Swivel: 0x4ccD4C002216f08218EdE1B13621faa80CecfC98
  * Destributor: 0x4644f0a61d823D635397317aFA06b51f57d4Eb33
* Kovan:
  * MarketPlace: 0xE7601e3FCB9bD0e948554F59DffC9D428E5091d4
  * Swivel: 0x301292f76885b5a20c7dbd0e06F093E9D4e5fA3F
  * Destributor: 0xe59Da40627588136f2b5525797c2789162C86b29
* Mainnet:
  * MarketPlace: 0x998689650D4d55822b4bDd4B7DB5F596bf6b3570
  * Swivel: 0x3b983B701406010866bD68331aAed374fb9f50C9
  * Destributor: 0x40CbFf2619b72f1F8e788fb7792142BA58bdc27C
