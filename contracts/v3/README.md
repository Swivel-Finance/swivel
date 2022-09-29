# Swivel v3

Swivel v3 was built using [Foundry](https://book.getfoundry.sh/), a Solidity development toolkit. As such, our tests are provided and can be run using `forge`. This README provides information about running and building the project.

## Testing

We use Foundry's fork-mode in our tests. In order to run these tests, a couple environment variables need to be set by the user:

- RPC_URL: a URL to an RPC endpoint of a **mainnet** Ethereum client
- BLOCK_NUMBER: the block from which the state will be forked. Note that it is important to select a valid block (i.e. must have USDC deployed, etc.). When running tests, we currently use block 15189976.

To run the tests, execute the following command:

`forge test --fork-url ${RPC_URL} --fork-block-number ${BLOCK_NUMBER} --use solc:0.8.16 `

Please note that some tests have not been fully complete, and may be skipped for now by including the following option:

`--no-match-test "Skip\(\B"`
