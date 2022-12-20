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
      accounts,
      gasPrice: 10000000000,
    },
    ternopil: {
      chainId: 953842,
      url: 'https://rpc.chainfusion.org',
      accounts,
      gasPrice: 10000000000,
    },
    bg1: {
      chainId: 5001,
      url: 'https://bg1-rpc.chainfusion.org',
      accounts,
      gasPrice: 10000000000,
    },
    bg2: {
      chainId: 5002,
      url: 'https://bg2-rpc.chainfusion.org',
      accounts,
      gasPrice: 10000000000,
    },
    bg3: {
      chainId: 5003,
      url: 'https://bg3-rpc.chainfusion.org',
      accounts,
      gasPrice: 10000000000,
    },
    sepolia: {
      chainId: 11155111,
      url: 'https://rpc.sepolia.org',
      accounts,
    },
    goerli: {
      chainId: 5,
      url: `https://goerli.infura.io/v3/${process.env.INFURA_PROJECT_ID}`,
      accounts,
    },
    'bsc-testnet': {
      chainId: 97,
      url: 'https://data-seed-prebsc-1-s1.binance.org:8545',
      accounts,
    },
    mumbai: {
      chainId: 80001,
      url: 'https://rpc-mumbai.maticvigil.com',
      accounts,
    },
    fuji: {
      chainId: 43113,
      url: 'https://api.avax-test.network/ext/bc/C/rpc',
      accounts,
    },
    'q-testnet': {
      chainId: 35443,
      url: 'https://rpc.qtestnet.org',
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
      {
        network: "sepolia",
        chainId: 11155111,
        urls: {
          apiURL: "https://sepolia.etherscan.io/api",
          browserURL: "https://sepolia.etherscan.io"
        }
      },
      {
        network: "goerli",
        chainId: 5,
        urls: {
          apiURL: "https://goerli.etherscan.io/api",
          browserURL: "https://goerli.etherscan.io"
        }
      },
      {
        network: "bsc-testnet",
        chainId: 97,
        urls: {
          apiURL: "https://bscscan.com/api",
          browserURL: "https://bscscan.com"
        }
      },
      {
        network: "mumbai",
        chainId: 80001,
        urls: {
          apiURL: "https://mumbai.polygonscan.com/api",
          browserURL: "https://mumbai.polygonscan.com"
        }
      },
      {
        network: 'fuji',
        chainId: 43113,
        urls: {
          apiURL: "https://testnet.snowtrace.io/api",
          browserURL: "https://testnet.snowtrace.io"
        }
      },
      {
        network: "q-testnet",
        chainId: 35443,
        urls: {
          apiURL: "https://explorer.qtestnet.org/api",
          browserURL: "https://explorer.qtestnet.org"
        }
      },
    ]
  },
};

export default config;
