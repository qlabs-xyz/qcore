# Ports

This section provides essential information about the ports used by the system, their primary purposes, and recommendations for exposure settings.


Here's your converted markdown table:

|                                                                                          | Description                                                                          | Default Port |
| ---------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------ | ------------ |
| **Cosmos [gRPC](/develop/api/cosmos-grpc#cosmos-grpc)**                                  | Query or send Cosmos EVM transactions using gRPC                                     | `9090`       |
| **Cosmos REST ([gRPC-Gateway](/develop/api/cosmos-grpc#cosmos-http-rest-grpc-gateway))** | Query or send Cosmos EVM transactions using an HTTP RESTful API                      | `9091`       |
| **Ethereum [JSON-RPC](/develop/api/ethereum-json-rpc)**                                  | Query Ethereum-formatted transactions and blocks or send Ethereum txs using JSON-RPC | `8545`       |
| **Ethereum [Websocket](/develop/api/ethereum-json-rpc#ethereum-websocket)**              | Subscribe to Ethereum logs and events emitted in smart contracts.                    | `8586`       |
| **Tendermint [RPC](#tendermint-rpc)**                                                    | Query transactions, blocks, consensus state, broadcast transactions, etc.            | `26657`      |
| **Tendermint [Websocket](#tendermint-websocket)**                                        | Subscribe to Tendermint ABCI events                                                  | `26657`      |