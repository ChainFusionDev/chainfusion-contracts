import { ethers } from 'hardhat';
import {
  Bridge,
  MockToken,
  FeeManager,
  TokenManager,
  LiquidityPools,
  RelayBridge,
  MockRelayBridgeApp,
  ValidatorStorage,
} from '../../typechain';

import { deploySystemContracts } from '../../scripts/deploy/system';
import { SystemDeploymentOptions, SystemDeploymentResult } from '../../scripts/deploy/types';

interface BridgeDeployment {
  bridge: Bridge;
  tokenManager: TokenManager;
  feeManager: FeeManager;
  validatorAddress: string;
  mockToken: MockToken;
  liquidityPools: LiquidityPools;
  chainId: number;
  relayBridge: RelayBridge;
  mockRelayBridgeApp: MockRelayBridgeApp;
  validatorStorage: ValidatorStorage;
}

export async function deployBridge(
  validatorAddress: string,
  chainId: number = 123,
  foundationAddress?: string
): Promise<BridgeDeployment> {
  const mintAmount = '10000000000000000000000';
  const validatorRefundFee = '10000000000000000';
  const destinationToken = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
  const feePercentage = '10000000000000000';

  if (foundationAddress === undefined) {
    foundationAddress = validatorAddress;
  }

  const ValidatorStorage = await ethers.getContractFactory('ValidatorStorage');
  const validatorStorage = await ValidatorStorage.deploy();
  await validatorStorage.deployed();
  await validatorStorage.initialize(validatorAddress);

  const MockToken = await ethers.getContractFactory('MockToken');
  const mockToken = await MockToken.deploy('Token', 'TKN', mintAmount);
  await mockToken.deployed();

  const TokenManager = await ethers.getContractFactory('TokenManager');
  const tokenManager = await TokenManager.deploy();
  await tokenManager.deployed();
  await tokenManager.initialize(validatorStorage.address);

  await tokenManager.setDestinationToken(chainId, mockToken.address, destinationToken);
  await tokenManager.setEnabled(mockToken.address, true);

  const LiquidityPools = await ethers.getContractFactory('LiquidityPools');
  const liquidityPools = await LiquidityPools.deploy();
  await liquidityPools.deployed();

  const FeeManager = await ethers.getContractFactory('FeeManager');
  const feeManager = await FeeManager.deploy();
  await feeManager.deployed();
  await feeManager.initialize(validatorStorage.address, liquidityPools.address, foundationAddress, validatorRefundFee);

  const Bridge = await ethers.getContractFactory('Bridge');
  const bridge = await Bridge.deploy();
  await bridge.deployed();
  await bridge.initialize(validatorStorage.address, tokenManager.address, liquidityPools.address, feeManager.address);

  const RelayBridge = await ethers.getContractFactory('RelayBridge');
  const relayBridge = await RelayBridge.deploy();
  await relayBridge.deployed();
  await relayBridge.initialize(validatorStorage.address);

  await liquidityPools.initialize(
    validatorStorage.address,
    tokenManager.address,
    bridge.address,
    feeManager.address,
    feePercentage
  );

  const MockRelayBridgeApp = await ethers.getContractFactory('MockRelayBridgeApp');
  const mockRelayBridgeApp = await MockRelayBridgeApp.deploy();
  await mockRelayBridgeApp.deployed();

  return {
    bridge: bridge,
    tokenManager: tokenManager,
    feeManager: feeManager,
    validatorAddress: validatorAddress,
    mockToken: mockToken,
    liquidityPools: liquidityPools,
    chainId: chainId,
    relayBridge: relayBridge,
    mockRelayBridgeApp: mockRelayBridgeApp,
    validatorStorage: validatorStorage,
  };
}

export async function deploySystem(options?: SystemDeploymentOptions): Promise<SystemDeploymentResult> {
  return await deploySystemContracts(options, false);
}
