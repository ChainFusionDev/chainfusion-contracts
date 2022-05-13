import { ethers } from 'hardhat';
import { BigNumber } from 'ethers';
import {
  Bridge,
  MockToken,
  FeeManager,
  TokenManager,
  ValidatorManager,
  ValidatorStaking,
  LiquidityPools,
  AddressStorage,
  DKG,
} from '../../typechain';

interface BridgeDeployment {
  bridge: Bridge;
  tokenManager: TokenManager;
  feeManager: FeeManager;
  validatorManager: ValidatorManager;
  mockToken: MockToken;
  liquidityPools: LiquidityPools;
  chainId: number;
}

interface SystemDeployment {
  validatorStaking: ValidatorStaking;
  addressStorage: AddressStorage;
  dkg: DKG;
}

export async function deployBridge(
  owner: string,
  validators: string[],
  requiredSignatures: number,
  chainId: number = 123,
  validatorAddress?: string,
  foundationAddress?: string
): Promise<BridgeDeployment> {
  const mintAmount = '10000000000000000000000';
  const validatorRefundFee = '10000000000000000';
  const destinationToken = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

  if (validatorAddress === undefined) {
    validatorAddress = owner;
  }

  if (foundationAddress === undefined) {
    foundationAddress = owner;
  }

  const MockToken = await ethers.getContractFactory('MockToken');
  const mockToken = await MockToken.deploy('Token', 'TKN', mintAmount);
  await mockToken.deployed();

  const TokenManager = await ethers.getContractFactory('TokenManager');
  const tokenManager = await TokenManager.deploy();
  await tokenManager.deployed();
  await tokenManager.initialize(owner);

  await tokenManager.setDestinationToken(chainId, mockToken.address, destinationToken);
  await tokenManager.setEnabled(mockToken.address, true);

  const LiquidityPools = await ethers.getContractFactory('LiquidityPools');
  const liquidityPools = await LiquidityPools.deploy();
  await liquidityPools.deployed();

  const ValidatorManager = await ethers.getContractFactory('ValidatorManager');
  const validatorManager = await ValidatorManager.deploy();
  await validatorManager.deployed();

  const FeeManager = await ethers.getContractFactory('FeeManager');
  const feeManager = await FeeManager.deploy();
  await feeManager.deployed();
  await feeManager.initialize(liquidityPools.address, validatorAddress, foundationAddress, validatorRefundFee);

  const Bridge = await ethers.getContractFactory('Bridge');
  const bridge = await Bridge.deploy();
  await bridge.deployed();
  await bridge.initialize(
    owner,
    validatorManager.address,
    tokenManager.address,
    liquidityPools.address,
    feeManager.address
  );

  const feePercentage = '10000000000000000';
  await liquidityPools.initialize(tokenManager.address, bridge.address, feeManager.address, feePercentage);

  await validatorManager.setRequiredApprovals(requiredSignatures);
  await validatorManager.setValidators(validators);

  return {
    bridge: bridge,
    tokenManager: tokenManager,
    feeManager: feeManager,
    validatorManager: validatorManager,
    mockToken: mockToken,
    liquidityPools: liquidityPools,
    chainId: chainId,
  };
}

export async function deploySystem(initialminimalStake: BigNumber): Promise<SystemDeployment> {
  const withdrawalPeriod = 1;
  const validators: string[] = [];

  const AddressStorage = await ethers.getContractFactory('AddressStorage');
  const addressStorage = await AddressStorage.deploy();
  await addressStorage.deployed();
  await (await addressStorage.initialize(validators)).wait();

  const ValidatorStaking = await ethers.getContractFactory('ValidatorStaking');
  const validatorStaking = await ValidatorStaking.deploy();
  await validatorStaking.deployed();

  const DKG = await ethers.getContractFactory('DKG');
  const dkg = await DKG.deploy();
  await dkg.deployed();

  await (
    await validatorStaking.initialize(initialminimalStake, withdrawalPeriod, addressStorage.address, dkg.address)
  ).wait();
  await (await dkg.initialize(validators, validatorStaking.address)).wait();

  await addressStorage.transferOwnership(validatorStaking.address);

  return {
    validatorStaking: validatorStaking,
    addressStorage: addressStorage,
    dkg: dkg,
  };
}
