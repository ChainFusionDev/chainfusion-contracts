import * as dotenv from "dotenv";

import { HardhatUserConfig, task } from "hardhat/config";
import "@nomiclabs/hardhat-etherscan";
import "@nomiclabs/hardhat-waffle";
import "@nomiclabs/hardhat-solhint";
import "@typechain/hardhat";
import "hardhat-gas-reporter";
import "solidity-coverage";

dotenv.config();

task("accounts", "Prints the list of accounts", async (taskArgs, hre) => {
  const accounts = await hre.ethers.getSigners();

  for (const account of accounts) {
    console.log(account.address);
  }
});

const accounts = process.env.PRIVATE_KEY !== undefined ? [process.env.PRIVATE_KEY] : [];

const config: HardhatUserConfig = {
  solidity: "0.8.9",
  networks: {
    hardhat: {
      blockGasLimit: 100_000_000,
    },
    ganache: {
      url: `http://ganache:8545`,
    },
    localhost: {
      accounts: {
        mnemonic: 'test test test test test test test test test test test junk'
      },
      gasPrice: 10000000000,
    },
    ternopil: {
      chainId: 953842,
      url: 'http://192.168.10.40:8545',
      accounts: {
        mnemonic: 'test test test test test test test test test test test junk'
      },
      gasPrice: 10000000000,
    },
    ropsten: {
      url: `https://ropsten.infura.io/v3/${process.env.INFURA_PROJECT_ID}`,
      accounts,
    },
    kovan: {
      url: `https://kovan.infura.io/v3/${process.env.INFURA_PROJECT_ID}`,
      accounts,
    },
    rinkeby: {
      url: `https://rinkeby.infura.io/v3/${process.env.INFURA_PROJECT_ID}`,
      accounts,
    },
    goerli: {
      url: `https://goerli.infura.io/v3/${process.env.INFURA_PROJECT_ID}`,
      accounts,
    },
  },
  gasReporter: {
    enabled: process.env.REPORT_GAS !== undefined,
    currency: "USD",
  },
  etherscan: {
    apiKey: '03316f51-9541-446f-a333-e1f712a33a75',
    customChains: [
      {
        network: "ternopil",
        chainId: 953842,
        urls: {
          apiURL: "http://192.168.10.44:4000/api",
          browserURL: "http://192.168.10.44:4000"
        }
      }
    ]
  },
};

export default config;
