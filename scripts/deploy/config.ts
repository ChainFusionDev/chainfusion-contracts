import { BaseContract } from 'ethers';
import { promises as fs } from 'fs';
import { parse, stringify } from 'yaml';

export interface ContractsConfig {
  [key: string]: string | undefined;
}

export async function readContractsConfig(): Promise<ContractsConfig> {
  try {
    const fileData = await fs.readFile('contracts.json');
    const config = JSON.parse(fileData.toString());

    return config;
  } catch (error) {
    return {};
  }
}

export async function writeContractsConfig(config: ContractsConfig): Promise<void> {
  try {
    await fs.writeFile('contracts.json', JSON.stringify(config, null, 2));
  } catch (error) {
    console.error(error);
  }
}

export async function updateContractsConfig(config: ContractsConfig, data: Object) {
  for (const [key, value] of Object.entries(data)) {
    if (value !== undefined && value instanceof BaseContract) {
      config[key] = value.address;
    }
  }
}

export async function readChainContractsConfig(chainId: number): Promise<ContractsConfig> {
  try {
    const fileData = await fs.readFile(`contracts-${chainId}.json`);
    const config = JSON.parse(fileData.toString());

    return config;
  } catch (error) {
    return {};
  }
}

export async function writeChainContractsConfig(chainId: number, config: ContractsConfig): Promise<void> {
  try {
    await fs.writeFile(`contracts-${chainId}.json`, JSON.stringify(config, null, 2));
  } catch (error) {
    console.error(error);
  }
}

export interface BridgeConfig {
  [key: string]: Config | undefined;
}

export interface Config {
  chainId: Number;
  rpcURL: string;
  bridgeContract: string;
  signerStorageContract: string;
  startBlock: Number;
}

export async function processBridgeConfig(
  bridgeContract: string,
  signerStorageContract: string,
  startBlock: Number
): Promise<void> {
  const bridgeConfig = await readBridgeConfig();

  updateBridgeConfig(bridgeConfig, bridgeContract, signerStorageContract, startBlock);
  await writeBridgeConfig(bridgeConfig);
}

export async function updateBridgeConfig(
  bridgeConfig: BridgeConfig,
  bridgeContract: string,
  signerStorageContract: string,
  startBlock: Number
): Promise<BridgeConfig> {
  const hardhat = require('hardhat');

  const rpcURL = hardhat.network.config.url;
  const chainId = hardhat.network.config.chainId;
  const name = hardhat.network.name;

  var con: Config = {
    chainId: chainId,
    rpcURL: rpcURL,
    bridgeContract: bridgeContract,
    signerStorageContract: signerStorageContract,
    startBlock: startBlock,
  };

  bridgeConfig[name] = con;

  return bridgeConfig;
}

export async function readBridgeConfig(): Promise<BridgeConfig> {
  try {
    const fileData = await fs.readFile('bridgeconfig.yaml');
    const config = parse(fileData.toString());

    return config;
  } catch (error) {
    return {};
  }
}

export async function writeBridgeConfig(bg: BridgeConfig): Promise<void> {
  try {
    await fs.writeFile('bridgeconfig.yaml', stringify(bg, null, 2));
  } catch (error) {
    console.error(error);
  }
}
