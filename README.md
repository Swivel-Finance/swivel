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
