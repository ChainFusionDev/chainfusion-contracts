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
};

export async function deploySystemContracts(
  options?: SystemDeploymentOptions,
  showLogs?: Boolean
): Promise<SystemDeploymentResult> {
  const parameters = resolveParameters(options);

  if (showLogs === undefined) {
    showLogs = false;
  }

  if (showLogs) {
    console.log('Deploying contracts\n');
  }

  const res: SystemDeployment = {
    contractRegistry: await deployContract(ethers.getContractFactory('ContractRegistry'), showLogs, 'ContractRegistry'),
    eventRegistry: await deployContract(ethers.getContractFactory('EventRegistry'), showLogs, 'EventRegistry'),
    addressStorage: await deployContract(ethers.getContractFactory('AddressStorage'), showLogs, 'AddressStorage'),
    staking: await deployContract(ethers.getContractFactory('Staking'), showLogs, 'Staking'),
    dkg: await deployContract(ethers.getContractFactory('DKG'), showLogs, 'DKG'),
    supportedTokens: await deployContract(ethers.getContractFactory('SupportedTokens'), showLogs, 'SupportedTokens'),
    slashingVoting: await deployContract(ethers.getContractFactory('SlashingVoting'), showLogs, 'SlashingVoting'),
  };

  if (showLogs) {
    console.log('Successfully deployed contracts\n');
  }

  if (showLogs) {
    console.log('Initializing contracts\n');
  }

  await waitTransaction(res.addressStorage.initialize([]), showLogs, 'Initializing AddressStorage');

  await waitTransaction(
    res.addressStorage.transferOwnership(res.staking.address),
    showLogs,
    'Transferring ownership of AddressStorage'
  );

  await waitTransaction(
    res.dkg.initialize(res.contractRegistry.address, parameters.dkgDeadlinePeriod),
    showLogs,
    'Initializing DKG'
  );

  await waitTransaction(res.contractRegistry.initialize(res.dkg.address), showLogs, 'Initializing ContractRegistry');

  await waitTransaction(res.eventRegistry.initialize(res.dkg.address), showLogs, 'Initializing EventRegistry');

  await waitTransaction(res.supportedTokens.initialize(res.dkg.address), showLogs, 'Initializing SupportedTokens');

  await waitTransaction(
    res.staking.initialize(
      res.dkg.address,
      parameters.minimalStake,
      parameters.stakeWithdrawalPeriod,
      res.contractRegistry.address,
      res.addressStorage.address
    ),
    showLogs,
    'Initializing Staking'
  );

  await waitTransaction(
    res.slashingVoting.initialize(
      res.dkg.address,
      parameters.slashingEpochPeriod,
      parameters.slashingBansThresold,
      parameters.slashingEpochs,
      res.contractRegistry.address
    ),
    showLogs,
    'Initializing SlashingVoting'
  );

  await res.contractRegistry.setContract(await res.slashingVoting.SLASHING_VOTING_KEY(), res.slashingVoting.address);
  await res.contractRegistry.setContract(await res.staking.STAKING_KEY(), res.staking.address);
  await res.contractRegistry.setContract(await res.dkg.DKG_KEY(), res.dkg.address);
  await res.contractRegistry.setContract(await res.supportedTokens.SUPPORTED_TOKENS_KEY(), res.supportedTokens.address);

  if (showLogs) {
    console.log('Successfully initialized contracts\n');
  }

  return {
    ...res,
    ...parameters,
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

  return parameters;
}
