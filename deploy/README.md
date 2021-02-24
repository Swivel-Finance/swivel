# Deploying
* export a private key, `main.go` expects it as `PRIVATE_KEY` available to `os.Getenv`
* export an infura project id, `main.go` expects it as `INFURA_PROJECT_ID` available to `os.Getenv`
* edit cmd/main.go for the `chainId` you want to deploy to
* `cd cmd && go run main.go`

### Notes
You may want to adjust the `auth.GasLimit` and set a `verifier` (you can see it is, ATM, set to `zero`)
