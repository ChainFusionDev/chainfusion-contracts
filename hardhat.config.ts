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
    localhost: {
      chainId: 1337,
      accounts: {
        mnemonic: 'test test test test test test test test test test test junk'
      },
      gasPrice: 100000000000,
    },
    ternopil: {
      chainId: 953842,
      url: 'https://rpc.chainfusion.org',
      gasPrice: 10000000000,
      accounts,
    },
    bg1: {
      chainId: 5001,
      url: 'https://bg1-rpc.chainfusion.org',
      gasPrice: 10000000000,
      accounts,
    },
    bg2: {
      chainId: 5002,
      url: 'https://bg2-rpc.chainfusion.org',
      gasPrice: 10000000000,
      accounts,
    },
    bg3: {
      chainId: 5003,
      url: 'https://bg3-rpc.chainfusion.org',
      gasPrice: 10000000000,
      accounts,
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
    apiKey: process.env.ETHERSCAN_API_KEY,
    customChains: [
      {
        network: "localhost",
        chainId: 1337,
        urls: {
          apiURL: "http://localhost:4000/api",
          browserURL: "http://localhost:4000"
        }
      },
      {
        network: "ternopil",
        chainId: 953842,
        urls: {
          apiURL: "https://explorer.chainfusion.org/api",
          browserURL: "https://explorer.chainfusion.org"
        }
      },
      {
        network: "bg1",
        chainId: 5001,
        urls: {
          apiURL: "https://bg1-explorer.chainfusion.org/api",
          browserURL: "https://bg1-explorer.chainfusion.org"
        }
      },
      {
        network: "bg2",
        chainId: 5002,
        urls: {
          apiURL: "https://bg2-explorer.chainfusion.org/api",
          browserURL: "https://bg2-explorer.chainfusion.org"
        }
      },
      {
        network: "bg3",
        chainId: 5003,
        urls: {
          apiURL: "https://bg3-explorer.chainfusion.org/api",
          browserURL: "https://bg3-explorer.chainfusion.org"
        }
      },
    ]
  },
};

export default config;
