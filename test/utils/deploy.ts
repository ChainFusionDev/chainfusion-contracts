import { ethers } from 'hardhat';
import { BigNumber } from 'ethers';
import {
  Bridge,
  MockToken,
  FeeManager,
  TokenManager,
  Staking,
  LiquidityPools,
  AddressStorage,
  DKG,
  SlashingVoting,
  RelayBridge,
  SupportedTokens,
  MockRelayBridgeApp,
  ValidatorStorage,
  ContractRegistry,
} from '../../typechain';

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

interface SystemDeployment {
  staking: Staking;
  addressStorage: AddressStorage;
  dkg: DKG;
  slashingVoting: SlashingVoting;
  supportedTokens: SupportedTokens;
  contractRegistry: ContractRegistry;
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

export async function deploySystem(initialMinimalStake?: BigNumber): Promise<SystemDeployment> {
  const withdrawalPeriod = 1;
  const epochPeriod = 100;
  const slashingThresold = 2;
  const slashingEpochs = 2;
  const deadlinePeriod = 20;

  if (initialMinimalStake === undefined) {
    initialMinimalStake = ethers.utils.parseEther('3');
  }

  const ContractRegistry = await ethers.getContractFactory('ContractRegistry');
  const contractRegistry = await ContractRegistry.deploy();
  await contractRegistry.deployed();

  const AddressStorage = await ethers.getContractFactory('AddressStorage');
  const addressStorage = await AddressStorage.deploy();
  await addressStorage.deployed();

  await (await addressStorage.initialize([])).wait();

  const Staking = await ethers.getContractFactory('Staking');
  const staking = await Staking.deploy();
  await staking.deployed();

  await contractRegistry.setContract(await staking.STAKING_KEY(), staking.address);

  const DKG = await ethers.getContractFactory('DKG');
  const dkg = await DKG.deploy();
  await dkg.deployed();

  await contractRegistry.setContract(await dkg.DKG_KEY(), dkg.address);

  const SlashingVoting = await ethers.getContractFactory('SlashingVoting');
  const slashingVoting = await SlashingVoting.deploy();
  await slashingVoting.deployed();
  await (
    await slashingVoting.initialize(epochPeriod, slashingThresold, slashingEpochs, contractRegistry.address)
  ).wait();

  await contractRegistry.setContract(await slashingVoting.SLASHING_VOTING_KEY(), slashingVoting.address);

  await contractRegistry.setContract(await dkg.DKG_KEY(), dkg.address);

  const SupportedTokens = await ethers.getContractFactory('SupportedTokens');
  const supportedTokens = await SupportedTokens.deploy();
  await supportedTokens.deployed();

  await contractRegistry.setContract(await supportedTokens.SUPPORTED_TOKENS_KEY(), supportedTokens.address);

  await (
    await staking.initialize(initialMinimalStake, withdrawalPeriod, contractRegistry.address, addressStorage.address)
  ).wait();
  await (await dkg.initialize(contractRegistry.address, deadlinePeriod)).wait();

  await addressStorage.transferOwnership(staking.address);

  return {
    staking: staking,
    addressStorage: addressStorage,
    slashingVoting: slashingVoting,
    dkg: dkg,
    supportedTokens: supportedTokens,
    contractRegistry: contractRegistry,
  };
}
