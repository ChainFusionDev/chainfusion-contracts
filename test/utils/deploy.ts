import { ethers } from 'hardhat';
import { BigNumber } from 'ethers';
import { deploySystemContracts, SystemDeploymentOptions, SystemDeploymentResult } from '../../scripts/deploy/system';
import { BridgeDeploymentOptions, BridgeDeploymentResult, deployBridgeContracts } from '../../scripts/deploy/bridge';
import { MockMintableBurnableToken, MockRelayBridgeApp, MockToken } from '../../typechain';

export async function deploySystem(options?: SystemDeploymentOptions): Promise<SystemDeploymentResult> {
  return await deploySystemContracts(options);
}

export async function deployBridge(options?: BridgeDeploymentOptions): Promise<BridgeDeploymentResult> {
  return await deployBridgeContracts(options);
}

export async function deployBridgeWithMocks(
  options?: BridgeDeploymentOptions
): Promise<BridgeWithMocksDeploymentResult> {
  const chainId = BigNumber.from(853);

  const deployment = await deployBridgeContracts(options);

  const MockToken = await ethers.getContractFactory('MockToken');
  const mockToken = await MockToken.deploy('Mock Token', 'MOCK', BigNumber.from('10000000000000000000000'));
  await mockToken.deployed();

  await deployment.tokenManager.setDestinationToken(chainId, mockToken.address, mockToken.address);
  await deployment.tokenManager.setEnabled(mockToken.address, true);

  const MockMintableBurnableToken = await ethers.getContractFactory('MockMintableBurnableToken');
  const mockMintableBurnableToken = await MockMintableBurnableToken.deploy('Mintable Mock Token', 'MINT');
  await mockMintableBurnableToken.deployed();

  await deployment.tokenManager.setDestinationToken(
    chainId,
    mockMintableBurnableToken.address,
    mockMintableBurnableToken.address
  );
  await deployment.tokenManager.setEnabled(mockMintableBurnableToken.address, true);
  await deployment.tokenManager.setMintable(mockMintableBurnableToken.address, true);

  const MockRelayBridgeApp = await ethers.getContractFactory('MockRelayBridgeApp');
  const mockRelayBridgeApp = await MockRelayBridgeApp.deploy();
  await mockRelayBridgeApp.deployed();

  return {
    mockChainId: chainId,
    mockToken: mockToken,
    mockMintableBurnableToken: mockMintableBurnableToken,
    mockRelayBridgeApp: mockRelayBridgeApp,
    ...deployment,
  };
}

export interface BridgeWithMocksDeploymentResult extends BridgeDeploymentResult {
  mockChainId: BigNumber;
  mockToken: MockToken;
  mockMintableBurnableToken: MockMintableBurnableToken;
  mockRelayBridgeApp: MockRelayBridgeApp;
}
