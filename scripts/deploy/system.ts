import { ethers } from 'hardhat';
import { BigNumber } from 'ethers';
import { Deployer } from './deployer';

import {
  Staking,
  AddressStorage,
  DKG,
  SlashingVoting,
  ContractRegistry,
  EventRegistry,
  BridgeAppFactory,
  ERC20BridgeMediator,
  ValidatorRewardDistributionPool,
} from '../../typechain';

const defaultSystemDeploymentParameters: SystemDeploymentParameters = {
  minimalStake: ethers.utils.parseEther('100'),
  stakeWithdrawalPeriod: BigNumber.from(60),
  router: '0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0',
  slashingEpochs: BigNumber.from(3),
  slashingEpochPeriod: BigNumber.from(1000),
  slashingBansThresold: BigNumber.from(10),

  dkgDeadlinePeriod: BigNumber.from(100),

  displayLogs: false,
  verify: false,
  stakingKeys: [],
};

export async function deploySystemContracts(options?: SystemDeploymentOptions): Promise<SystemDeploymentResult> {
  const params = resolveParameters(options);
  const deployer = new Deployer(params.displayLogs);

  const tockenSymbol = 'NFT';
  const tockenChainId = 5;
  const tockenAddress = '0x7f4C3165388a79a35DE397f88b89b59de0934bcd';

  deployer.log('Deploying contracts\n');

  const res: SystemDeployment = {
    contractRegistry: await deployer.deploy(ethers.getContractFactory('ContractRegistry'), 'ContractRegistry'),
    eventRegistry: await deployer.deploy(ethers.getContractFactory('EventRegistry'), 'EventRegistry'),
    addressStorage: await deployer.deploy(ethers.getContractFactory('AddressStorage'), 'AddressStorage'),
    staking: await deployer.deploy(ethers.getContractFactory('Staking'), 'Staking'),
    dkg: await deployer.deploy(ethers.getContractFactory('DKG'), 'DKG'),
    slashingVoting: await deployer.deploy(ethers.getContractFactory('SlashingVoting'), 'SlashingVoting'),
    bridgeAppFactory: await deployer.deploy(ethers.getContractFactory('BridgeAppFactory'), 'BridgeAppFactory'),
    erc20BridgeMediator: await deployer.deploy(ethers.getContractFactory('ERC20BridgeMediator'), 'ERC20BridgeMediator'),
    validatorRewardDistributionPool: await deployer.deploy(
      ethers.getContractFactory('ValidatorRewardDistributionPool'),
      'ValidatorRewardDistributionPool'
    ),
  };

  deployer.log('Successfully deployed contracts\n');

  deployer.log('Initializing contracts\n');

  await deployer.sendTransaction(res.addressStorage.initialize([]), 'Initializing AddressStorage');

  await deployer.sendTransaction(
    res.addressStorage.transferOwnership(res.staking.address),
    'Transferring ownership of AddressStorage'
  );

  await deployer.sendTransaction(
    res.dkg.initialize(res.contractRegistry.address, params.dkgDeadlinePeriod),
    'Initializing DKG'
  );

  await deployer.sendTransaction(res.contractRegistry.initialize(res.dkg.address), 'Initializing ContractRegistry');

  await deployer.sendTransaction(
    res.staking.initialize(
      res.dkg.address,
      params.minimalStake,
      params.stakeWithdrawalPeriod,
      res.contractRegistry.address,
      res.addressStorage.address
    ),
    'Initializing Staking'
  );

  await deployer.sendTransaction(
    res.slashingVoting.initialize(
      res.dkg.address,
      res.staking.address,
      params.slashingEpochPeriod,
      params.slashingBansThresold,
      params.slashingEpochs,
      res.contractRegistry.address
    ),
    'Initializing SlashingVoting'
  );

  await deployer.sendTransaction(res.eventRegistry.initialize(res.staking.address), 'Initializing EventRegistry');
  await deployer.sendTransaction(
    res.validatorRewardDistributionPool.initialize(res.contractRegistry.address, params.router, res.dkg.address),
    'Initializing ValidatorRewardDistributionPool'
  );

  await deployer.sendTransaction(res.bridgeAppFactory.createApp(), 'Creating BridgeApp');
  const bridgeApp = await ethers.getContractAt('BridgeApp', await res.bridgeAppFactory.apps(0));
  await deployer.sendTransaction(
    bridgeApp.setMediator(res.erc20BridgeMediator.address),
    'Seting BridgeMediator to BridgeApp'
  );
  await deployer.sendTransaction(
    res.erc20BridgeMediator.addToken(tockenSymbol, tockenChainId, tockenAddress),
    'adde token to BridgeMediator'
  );

  await res.contractRegistry.setContract(await res.slashingVoting.SLASHING_VOTING_KEY(), res.slashingVoting.address);
  await res.contractRegistry.setContract(await res.staking.STAKING_KEY(), res.staking.address);
  await res.contractRegistry.setContract(await res.dkg.DKG_KEY(), res.dkg.address);
  await res.contractRegistry.setContract(await res.eventRegistry.EVENT_REGISTRY_KEY(), res.eventRegistry.address);
  await res.contractRegistry.setContract(
    await res.validatorRewardDistributionPool.VALIDATOR_REWARD_DISTRIBUTION_POOL_KEY(),
    res.validatorRewardDistributionPool.address
  );

  deployer.log('Successfully initialized contracts\n');

  if (params.stakingKeys.length > 0) {
    console.log('Staking for initial validators\n');

    for (const privateKey of params.stakingKeys) {
      const signer = new ethers.Wallet(privateKey, ethers.provider);
      const signerStaking = await ethers.getContractAt('Staking', res.staking.address, signer);
      const msg = `Staking ${ethers.utils.formatEther(params.minimalStake)} CFN for: ${signer.address}`;
      await deployer.sendTransaction(signerStaking.stake({ value: params.minimalStake }), msg);
    }

    deployer.log('Successfully staked\n');
  }

  if (params.verify) {
    await deployer.verifyObjectValues(res);
  }

  return {
    ...res,
    ...params,
  };
}

function resolveParameters(options?: SystemDeploymentOptions): SystemDeploymentParameters {
  let parameters = defaultSystemDeploymentParameters;

  if (options === undefined) {
    return parameters;
  }

  if (options.minimalStake !== undefined) {
    parameters.minimalStake = options.minimalStake;
  }

  if (options.stakeWithdrawalPeriod !== undefined) {
    parameters.stakeWithdrawalPeriod = options.stakeWithdrawalPeriod;
  }

  if (options.slashingEpochs !== undefined) {
    parameters.slashingEpochs = options.slashingEpochs;
  }

  if (options.slashingEpochPeriod !== undefined) {
    parameters.slashingEpochPeriod = options.slashingEpochPeriod;
  }

  if (options.slashingBansThresold !== undefined) {
    parameters.slashingBansThresold = options.slashingBansThresold;
  }

  if (options.dkgDeadlinePeriod !== undefined) {
    parameters.dkgDeadlinePeriod = options.dkgDeadlinePeriod;
  }

  if (options.displayLogs !== undefined) {
    parameters.displayLogs = options.displayLogs;
  }

  if (options.verify !== undefined) {
    parameters.verify = options.verify;
  }

  if (options.stakingKeys !== undefined) {
    parameters.stakingKeys = options.stakingKeys;
  }

  if (options.router !== undefined) {
    parameters.router = options.router;
  }

  return parameters;
}

export interface SystemDeploymentResult extends SystemDeployment, SystemDeploymentParameters {}

export interface SystemDeployment {
  staking: Staking;
  addressStorage: AddressStorage;
  dkg: DKG;
  slashingVoting: SlashingVoting;
  contractRegistry: ContractRegistry;
  eventRegistry: EventRegistry;
  bridgeAppFactory: BridgeAppFactory;
  erc20BridgeMediator: ERC20BridgeMediator;
  validatorRewardDistributionPool: ValidatorRewardDistributionPool;
}

export interface SystemDeploymentParameters {
  minimalStake: BigNumber;
  stakeWithdrawalPeriod: BigNumber;

  slashingEpochs: BigNumber;
  slashingEpochPeriod: BigNumber;
  slashingBansThresold: BigNumber;

  dkgDeadlinePeriod: BigNumber;

  router: string;

  displayLogs: boolean;
  verify: boolean;
  stakingKeys: string[];
}

export interface SystemDeploymentOptions {
  minimalStake?: BigNumber;
  stakeWithdrawalPeriod?: BigNumber;

  slashingEpochs?: BigNumber;
  slashingEpochPeriod?: BigNumber;
  slashingBansThresold?: BigNumber;

  dkgDeadlinePeriod?: BigNumber;

  router?: string;

  displayLogs?: boolean;
  verify?: boolean;
  stakingKeys?: string[];
}
