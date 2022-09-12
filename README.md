# CW-CONTRACTS-RESOLVER

CW Contracts Resolver is a simple golang server and client library for the CosmWasm contract registry. It provides a convenient
API, automatically updating itself to any changes in the github repo once a day.

## Usage

Install the binary

```cli
go install github.com/teamscanworks/cw-contracts-resolver@latest
```

and run the server (this is for port :8080)

```cli
cw-contract-resolver teamscanworks/cw-contract-registry <rpc-endpoint> :8080
```

## API Reference


| Query                                    | Description                                   | Response Type |
|------------------------------------------|-----------------------------------------------|---------------|
| `/v1/chains`                             | Returns an array of registered chains by name | `[]string`    |
| `/v1/chain/{chain}/contracts/{codeId}`   | Returns the contract info by code id          | `Contract`    |
| `/v1/chain/{chain}/contracts/{addreess}` | Returns the contract info by address          | `Contract`    |

## Acknowledgements

The cw-contracts-resolver is a modified version of Skychart, by @cmwaters