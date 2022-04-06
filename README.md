# ChainFusion Contracts

This repository contains Solidity contracts with bridge logic

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

In This example we are deploying to `rinkeby` testnet. To deploy to different chain, `--network` parameter should be changed to `ropsten` or `goerli`, etc.

```
$ npx hardhat --network rinkeby deploy scripts/deploy.ts
```

### Verify Contracts

To verify contracts, you need to specify network, contract address and constructor parameters (if present).

```
$ npx hardhat --network rikneby verify <CONTRACT_ADDRESS> <CONSTRUCTOR_PARAMETERS>
```
