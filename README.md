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
* Rinkeby (development):
  * MarketPlace: 0xb029e31C252549B20a8dc502c2832F11A2EDE417
  * Swivel: 0x3a09584FF42CDFe27Fe72Da0533bba24E9C28AaD
  * Destributor: 0x0C6f4487881166dA2e26111aBC2dCC852B25ee2f (previous version)
* Rinkeby (staging):
  * MarketPlace: 0x9fa54f942D8b8e992501952C3e6E67F1A42595b8
  * Swivel: 0x4ccD4C002216f08218EdE1B13621faa80CecfC98
  * Destributor: 0x4644f0a61d823D635397317aFA06b51f57d4Eb33 (previous version)
* Mainnet (production):
  * MarketPlace: 0x998689650D4d55822b4bDd4B7DB5F596bf6b3570
  * Swivel: 0x3b983B701406010866bD68331aAed374fb9f50C9
  * Destributor: 0x57E18D9F50F3Fd0894c8436BC84D2f523A8d0968 (current version)
