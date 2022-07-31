import { ethers } from 'hardhat';
import { BigNumber } from 'ethers';
import { deployContract, waitTransaction } from './utils';
import { SystemDeployment, SystemDeploymentOptions, SystemDeploymentParameters, SystemDeploymentResult } from './types';

const defaultSystemDeploymentParameters: SystemDeploymentParameters = {
  minimalStake: ethers.utils.parseEther('100'),
  stakeWithdrawalPeriod: BigNumber.from(60),

  slashingEpochs: BigNumber.from(3),
  slashingEpochPeriod: BigNumber.from(1000),
  slashingBansThresold: BigNumber.from(10),

  dkgDeadlinePeriod: BigNumber.from(20),

  displayLogs: false,
};

export async function deploySystemContracts(options?: SystemDeploymentOptions): Promise<SystemDeploymentResult> {
  const params = resolveParameters(options);

  if (params.displayLogs) {
    console.log('Deploying contracts\n');
  }

  const res: SystemDeployment = {
    contractRegistry: await deployContract(
      ethers.getContractFactory('ContractRegistry'),
      params.displayLogs,
      'ContractRegistry'
    ),
    eventRegistry: await deployContract(
      ethers.getContractFactory('EventRegistry'),
      params.displayLogs,
      'EventRegistry'
    ),
    addressStorage: await deployContract(
      ethers.getContractFactory('AddressStorage'),
      params.displayLogs,
      'AddressStorage'
    ),
    staking: await deployContract(ethers.getContractFactory('Staking'), params.displayLogs, 'Staking'),
    dkg: await deployContract(ethers.getContractFactory('DKG'), params.displayLogs, 'DKG'),
    supportedTokens: await deployContract(
      ethers.getContractFactory('SupportedTokens'),
      params.displayLogs,
      'SupportedTokens'
    ),
    slashingVoting: await deployContract(
      ethers.getContractFactory('SlashingVoting'),
      params.displayLogs,
      'SlashingVoting'
    ),
  };

  if (params.displayLogs) {
    console.log('Successfully deployed contracts\n');
  }

  if (params.displayLogs) {
    console.log('Initializing contracts\n');
  }

  await waitTransaction(res.addressStorage.initialize([]), params.displayLogs, 'Initializing AddressStorage');

  await waitTransaction(
    res.addressStorage.transferOwnership(res.staking.address),
    params.displayLogs,
    'Transferring ownership of AddressStorage'
  );

  await waitTransaction(
    res.dkg.initialize(res.contractRegistry.address, params.dkgDeadlinePeriod),
    params.displayLogs,
    'Initializing DKG'
  );

  await waitTransaction(
    res.contractRegistry.initialize(res.dkg.address),
    params.displayLogs,
    'Initializing ContractRegistry'
  );

  await waitTransaction(
    res.eventRegistry.initialize(res.dkg.address),
    params.displayLogs,
    'Initializing EventRegistry'
  );

  await waitTransaction(
    res.supportedTokens.initialize(res.dkg.address),
    params.displayLogs,
    'Initializing SupportedTokens'
  );

  await waitTransaction(
    res.staking.initialize(
      res.dkg.address,
      params.minimalStake,
      params.stakeWithdrawalPeriod,
      res.contractRegistry.address,
      res.addressStorage.address
    ),
    params.displayLogs,
    'Initializing Staking'
  );

  await waitTransaction(
    res.slashingVoting.initialize(
      res.dkg.address,
      params.slashingEpochPeriod,
      params.slashingBansThresold,
      params.slashingEpochs,
      res.contractRegistry.address
    ),
    params.displayLogs,
    'Initializing SlashingVoting'
  );

  await res.contractRegistry.setContract(await res.slashingVoting.SLASHING_VOTING_KEY(), res.slashingVoting.address);
  await res.contractRegistry.setContract(await res.staking.STAKING_KEY(), res.staking.address);
  await res.contractRegistry.setContract(await res.dkg.DKG_KEY(), res.dkg.address);
  await res.contractRegistry.setContract(await res.supportedTokens.SUPPORTED_TOKENS_KEY(), res.supportedTokens.address);

  if (params.displayLogs) {
    console.log('Successfully initialized contracts\n');
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

  return parameters;
}
