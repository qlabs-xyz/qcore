# Interact with Node

## Getting The Client

To be able to interact with the network, i.e., execute transactions and perform queries, you need to
have a client installed. Please reference [this repo](https://github.com/outbe/outbe-node) to install the client locally.

At this point we assume that you have `outbe-noded` available in your terminal.

## Working with outbe-noded client

`outbe-noded` is a special binary for accessing OutBe network and it has the same functionality as 
a standard client. It means that it provides the same
commands and args, for example, to generate a new keypair:

```shell
outbe-noded keys add my-wallet 
```

Please see manual for all available commands: 

```shell
outbe-noded --help 
```

## Working with devnet

Currently, only a devnet environment is supported which is publicly available. 
Here are details of the denvet environment. We recommend creating `.env` file in your workdir 
and put there the following params so you can always have them via `source .env`:

```dotenv
CHAIN_ID="outbe-devnet-1"
FEE_DENOM="qnc"
STAKE_DENOM="qnc"
BECH32_HRP="qnc"
WASMD_VERSION="v0.52.0"
CONFIG_DIR=".outbe"
BINARY="outbe-noded"

RPC="https://rpc.dev.outbe.io:26657"
API="https://rpc.dev.outbe.io:1317"

NODE=(--node $RPC)
TXFLAG=($NODE --chain-id $CHAIN_ID --gas-prices 0.25$FEE_DENOM --gas auto --gas-adjustment 1.3 --output json)
```

### Querying a smart contract

To query a smart contract, you can use the following command:

```shell
outbe-noded $NODE query wasm contract-state smart [CONTRACT_ADDRESS] [QUERY]
```

Query is a JSON object that you create based on the query message. Let's say you have this query message:

```rust
pub enum QueryMsg {
    Tokens {
        owner: Addr,
        start_after: Option<String>,
        limit: Option<u32>,
    },
}
```

You can create a query message like this:

```shell
outbe-noded $NODE query wasm contract-state smart [CONTRACT_ADDRESS] '{"tokens": {"owner": "$owner", "start_after": null, "limit": 10}}'
```

### Execute a smart contract

To execute a smart contract, you can use the following command:

```shell
outbe-noded tx wasm execute [CONTRACT_ADDRESS] [EXECUTE_MESSAGE] --from [YOUR_ADDRESS] $TXFLAG
```

Execute message is a JSON object that you create based on the execute message. 
Let's say you have this an execute message:

```rust
pub enum ExecuteMsg {
    Mint {
        recipient: Addr,
    },
}
```

You can create an execute message like this:

```shell
outbe-noded tx wasm execute [CONTRACT_ADDRESS] '{"mint": {"recipient": "$address"}}' --from my-wallet $TXFLAG
```

