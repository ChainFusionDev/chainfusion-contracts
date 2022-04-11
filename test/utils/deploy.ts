import { ethers } from 'hardhat';
import { Bridge, MockToken, TokenManager, ValidatorManager } from '../../typechain';

interface BridgeDeployment {
  bridge: Bridge;
  tokenManager: TokenManager;
  validatorManager: ValidatorManager;
  mockToken: MockToken;
  chainId: number;
}

export async function deployBridge(
  owner: string,
  validators: string[],
  requiredSignatures: number,
  chainId: number = 123,
): Promise<BridgeDeployment> {
  const mintAmount = '100000000000000000000';
  const destinationToken = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

  const MockToken = await ethers.getContractFactory('MockToken');
  const mockToken = await MockToken.deploy('Token', 'TKN', mintAmount);
  await mockToken.deployed();

  const TokenManager = await ethers.getContractFactory('TokenManager');
  const tokenManager = await TokenManager.deploy();
  await tokenManager.deployed();
  await tokenManager.initialize(owner);
  await tokenManager.addSupportedToken(chainId, mockToken.address, destinationToken);

  const ValidatorManager = await ethers.getContractFactory('ValidatorManager');
  const validatorManager = await ValidatorManager.deploy();
  await validatorManager.deployed();

  const Bridge = await ethers.getContractFactory('Bridge');
  const bridge = await Bridge.deploy();
  await bridge.deployed();
  await bridge.initialize(owner, validatorManager.address, tokenManager.address);

  await validatorManager.setRequiredApprovals(requiredSignatures);
  await validatorManager.setValidators(validators);

  return {
    bridge: bridge,
    tokenManager: tokenManager,
    validatorManager: validatorManager,
    mockToken: mockToken,
    chainId: chainId,
  };
}
