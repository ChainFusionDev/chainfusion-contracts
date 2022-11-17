# ChainFusion Contracts

This repository contains Solidity contracts with bridging logic.

RelayBridge is a cross-chain system that acts as a bridge to transfer tokens and assets between two different blockchain platforms.

RelayBridge provides blockchain interoperability by connecting different blockchain networks. Thus, it allows users to experience the distinctive features of each network and the additional benefits of a chain of nodes.

When a user initiates a token transfer from chain A to chain B, the assets are first sent to a relay node on chain A. The relay node then sends the same number of tokens to chain B and maintains liquidity at destination B.

RelayBridge provides interaction between blockchain networks using various consensus mechanisms. They enable interoperable token transfer, data transfer, digital asset transfer, and smart contract instruction transfer between two different platforms. With RelayBridges, users can deploy digital assets stored on one chain to applications running on another, facilitating fast and low-cost transactions for tokens hosted on a less scalable blockchain.

## Development

### Install Packages

```
$ npm install
```

### Compile Contracts

```
$ npm run compile
```

### Run Tests

```
$ npm run test
```

## Deployment and Verification

First of all, copy `.env.example` into `.env` and set up all required variables inside

### Deploy Contracts

In This example we are deploying to `goerli` testnet. To deploy to different chain, `--network` parameter should be changed to `ropsten`, etc.

Deployment involves running two scripts in sequence. The first script is deployed once.

```
$ npx hardhat --network goerli deploy scripts/deploy-system.ts
```
The second script is deployed each time for a new chain.

```
$ npx hardhat --network goerli deploy scripts/deploy-bridge.ts
```

### Verify Contracts

To verify contracts, you need to specify network, contract address and constructor parameters (if present).

```
$ npx hardhat --network goerli verify <CONTRACT_ADDRESS> <CONSTRUCTOR_PARAMETERS>
```
