import { ethers, network } from 'hardhat';
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
  ValidatorRewardDistributionPool,
} from '../../typechain';

const defaultSystemDeploymentParameters: SystemDeploymentParameters = {
  minimalStake: ethers.utils.parseEther('100'),
  minimalValidators: BigNumber.from(2),
  stakeWithdrawalPeriod: BigNumber.from(60),
  router: '0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0',
  erc20BridgeMediator: '0x0000000000000000000000000000000000000001',
  slashingEpochs: BigNumber.from(3),
  slashingEpochPeriod: BigNumber.from(1000),
  slashingBansThresold: BigNumber.from(10),

  dkgDeadlinePeriod: BigNumber.from(100),

  displayLogs: false,
  parallelDeployment: false,
  verify: false,
  stakingKeys: [],
};

export async function deploySystemContracts(options?: SystemDeploymentOptions): Promise<SystemDeploymentResult> {
  const params = resolveParameters(options);
  const deployer = new Deployer(params.displayLogs, params.parallelDeployment);

  deployer.log(`🧾 Deploying ChainFusion contracts (${network.name} chain)\n`);

  const [deployerSigner] = await ethers.getSigners();
  const deployerAddress = await deployerSigner.getAddress();
  deployer.setNonce(await ethers.provider.getTransactionCount(deployerAddress));

  const contractRegistryPromise = deployer.deploy(ethers.getContractFactory('ContractRegistry'), 'ContractRegistry');
  const eventRegistryPromise = deployer.deploy(ethers.getContractFactory('EventRegistry'), 'EventRegistry');
  const addressStoragePromise = deployer.deploy(ethers.getContractFactory('AddressStorage'), 'AddressStorage');
  const stakingPromise = deployer.deploy(ethers.getContractFactory('Staking'), 'Staking');
  const dkgPromise = deployer.deploy(ethers.getContractFactory('DKG'), 'DKG');
  const stakingVotingPromise = deployer.deploy(ethers.getContractFactory('SlashingVoting'), 'SlashingVoting');
  const bridgeAppFactoryPromise = deployer.deploy(ethers.getContractFactory('BridgeAppFactory'), 'BridgeAppFactory');
  const validatorRewardDistributionPoolPromise = deployer.deploy(
    ethers.getContractFactory('ValidatorRewardDistributionPool'),
    'ValidatorRewardDistributionPool'
  );

  const res: SystemDeployment = {
    contractRegistry: await contractRegistryPromise,
    eventRegistry: await eventRegistryPromise,
    addressStorage: await addressStoragePromise,
    staking: await stakingPromise,
    dkg: await dkgPromise,
    slashingVoting: await stakingVotingPromise,
    bridgeAppFactory: await bridgeAppFactoryPromise,
    validatorRewardDistributionPool: await validatorRewardDistributionPoolPromise,
  };

  deployer.log('Successfully deployed ChainFusion contracts\n');

  deployer.log(`🔄 Initializing ChainFusion contracts (${network.name} chain)\n`);

  await deployer.waitPromises([
    deployer.sendTransaction(res.addressStorage.initialize([], deployer.getOverrides()), 'Initializing AddressStorage'),
    deployer.sendTransaction(
      res.dkg.initialize(
        res.contractRegistry.address,
        params.dkgDeadlinePeriod,
        params.minimalValidators,
        deployer.getOverrides()
      ),
      'Initializing DKG'
    ),
    deployer.sendTransaction(
      res.contractRegistry.initialize(res.dkg.address, deployer.getOverrides()),
      'Initializing ContractRegistry'
    ),
    deployer.sendTransaction(
      res.staking.initialize(
        res.dkg.address,
        params.minimalStake,
        params.stakeWithdrawalPeriod,
        res.contractRegistry.address,
        res.addressStorage.address,
        deployer.getOverrides()
      ),
      'Initializing Staking'
    ),
    deployer.sendTransaction(
      res.slashingVoting.initialize(
        res.dkg.address,
        res.staking.address,
        params.slashingEpochPeriod,
        params.slashingBansThresold,
        params.slashingEpochs,
        res.contractRegistry.address,
        deployer.getOverrides()
      ),
      'Initializing SlashingVoting'
    ),
    deployer.sendTransaction(
      res.eventRegistry.initialize(res.staking.address, deployer.getOverrides()),
      'Initializing EventRegistry'
    ),
    deployer.sendTransaction(
      res.validatorRewardDistributionPool.initialize(
        res.contractRegistry.address,
        params.router,
        res.dkg.address,
        deployer.getOverrides()
      ),
      'Initializing ValidatorRewardDistributionPool'
    ),
  ]);

  await deployer.sendTransaction(
    res.addressStorage.transferOwnership(res.staking.address),
    'Transferring ownership of AddressStorage'
  );

  deployer.setNonce(await ethers.provider.getTransactionCount(deployerAddress));

  const registryTxs = [
    res.contractRegistry.setContract(
      await res.slashingVoting.SLASHING_VOTING_KEY(),
      res.slashingVoting.address,
      deployer.getOverrides()
    ),
    res.contractRegistry.setContract(await res.staking.STAKING_KEY(), res.staking.address, deployer.getOverrides()),
    res.contractRegistry.setContract(await res.dkg.DKG_KEY(), res.dkg.address, deployer.getOverrides()),
    res.contractRegistry.setContract(
      await res.eventRegistry.EVENT_REGISTRY_KEY(),
      res.eventRegistry.address,
      deployer.getOverrides()
    ),
    res.contractRegistry.setContract(
      await res.validatorRewardDistributionPool.VALIDATOR_REWARD_DISTRIBUTION_POOL_KEY(),
      res.validatorRewardDistributionPool.address,
      deployer.getOverrides()
    ),
  ];

  await deployer.waitTransactions(registryTxs);

  deployer.log('Successfully initialized ChainFusion contracts\n');

  if (params.stakingKeys.length > 0 && network.name !== 'hardhat') {
    deployer.log(`👤 Staking for initial validators (${network.name} chain)\n`);

    let validatorId: number = 1;
    let stakeTxs: Promise<void>[] = [];
    for (const privateKey of params.stakingKeys) {
      const signer = new ethers.Wallet(privateKey, ethers.provider);
      const signerStaking = await ethers.getContractAt('Staking', res.staking.address, signer);
      const msg = `Staking ${ethers.utils.formatEther(params.minimalStake)} CFN for validator ${validatorId++}`;
      stakeTxs.push(
        deployer.sendTransaction(signerStaking.stake({ value: params.minimalStake, gasLimit: 1000000 }), msg)
      );
    }

    await deployer.waitPromises(stakeTxs);

    const targetGeneration = BigNumber.from(params.stakingKeys.length - 2);
    deployer.log('Successfully staked\n');

    deployer.log(`⏳ Waiting for distributed key generation to complete, this could take a while\n`);
    await waitSignerAddressUpdated(res.dkg, targetGeneration);
    deployer.log(`✅ Distributed key generation complete\n`);
  }

  if (params.verify) {
    await deployer.verifyObjectValues(res);
  }

  return {
    ...res,
    ...params,
  };
}

async function waitSignerAddressUpdated(dkg: DKG, generation: BigNumber): Promise<string> {
  return new Promise<string>((resolve) => {
    const eventName = 'SignerAddressUpdated';
    const listener = (gen: BigNumber, signerAddress: string) => {
      if (gen.eq(generation)) {
        dkg.off(eventName, listener);
        resolve(signerAddress);
        return false;
      }
    };

    dkg.on(eventName, listener);
  });
}

function resolveParameters(options?: SystemDeploymentOptions): SystemDeploymentParameters {
  let parameters = defaultSystemDeploymentParameters;

  if (options === undefined) {
    return parameters;
  }

  if (options.minimalStake !== undefined) {
    parameters.minimalStake = options.minimalStake;
  }

  if (options.minimalValidators !== undefined) {
    parameters.minimalValidators = options.minimalValidators;
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

  if (options.parallelDeployment !== undefined) {
    parameters.parallelDeployment = options.parallelDeployment;
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

  if (options.erc20BridgeMediator !== undefined) {
    parameters.erc20BridgeMediator = options.erc20BridgeMediator;
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
  validatorRewardDistributionPool: ValidatorRewardDistributionPool;
}

export interface SystemDeploymentParameters {
  minimalStake: BigNumber;
  minimalValidators: BigNumber;
  stakeWithdrawalPeriod: BigNumber;

  slashingEpochs: BigNumber;
  slashingEpochPeriod: BigNumber;
  slashingBansThresold: BigNumber;

  dkgDeadlinePeriod: BigNumber;

  router: string;
  erc20BridgeMediator: string;

  displayLogs: boolean;
  parallelDeployment: boolean;
  verify: boolean;
  stakingKeys: string[];
}

export interface SystemDeploymentOptions {
  minimalStake?: BigNumber;
  minimalValidators?: BigNumber;
  stakeWithdrawalPeriod?: BigNumber;

  slashingEpochs?: BigNumber;
  slashingEpochPeriod?: BigNumber;
  slashingBansThresold?: BigNumber;

  dkgDeadlinePeriod?: BigNumber;

  router?: string;
  erc20BridgeMediator?: string;

  displayLogs?: boolean;
  parallelDeployment?: boolean;
  verify?: boolean;
  stakingKeys?: string[];
}
