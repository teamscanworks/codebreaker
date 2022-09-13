# Codebreaker

Codebreaker is a simple Go server and client library for the [CosmWasm Contracts Registry](https://github.com/teamscanworks/cw-contracts-registry). It provides a convenient
API, automatically updating itself to any changes in the Github repo once a day.

## Usage

Install the binary

```cli
go install github.com/teamscanworks/codebreaker@latest
```

and run the server (this is for port :8080)

```cli
codebreaker teamscanworks/cw-contract-registry <rpc-endpoint> :8080
```

## API Reference


| Query                                    | Description                                   | Response Type |
|------------------------------------------|-----------------------------------------------|---------------|
| `/v1/chains`                             | Returns an array of registered chains by name | `[]string`    |
| `/v1/chain/{chain}/code-id/{codeId}`     | Returns the contract info by code id          | `Contract`    |
| `/v1/chain/{chain}/address/{addreess}`   | Returns the contract info by address (TODO)   | `Contract`    |

## Acknowledgements

Codebreaker is a modified version of [Skychart](https://github.com/cmwaters/skychart), a Go server for the Cosmos Chain Registry by @cmwaters.
