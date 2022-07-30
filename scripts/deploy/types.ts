import { BigNumber } from 'ethers';

import {
  Staking,
  AddressStorage,
  DKG,
  SlashingVoting,
  SupportedTokens,
  ContractRegistry,
  EventRegistry,
} from '../../typechain';

export interface SystemDeploymentResult extends SystemDeployment, SystemDeploymentParameters {}

export interface SystemDeployment {
  staking: Staking;
  addressStorage: AddressStorage;
  dkg: DKG;
  slashingVoting: SlashingVoting;
  supportedTokens: SupportedTokens;
  contractRegistry: ContractRegistry;
  eventRegistry: EventRegistry;
}

export interface SystemDeploymentOptions {
  minimalStake?: BigNumber;
  stakeWithdrawalPeriod?: BigNumber;

  slashingEpochs?: BigNumber;
  slashingEpochPeriod?: BigNumber;
  slashingBansThresold?: BigNumber;

  dkgDeadlinePeriod?: BigNumber;
}

export interface SystemDeploymentParameters {
  minimalStake: BigNumber;
  stakeWithdrawalPeriod: BigNumber;

  slashingEpochs: BigNumber;
  slashingEpochPeriod: BigNumber;
  slashingBansThresold: BigNumber;

  dkgDeadlinePeriod: BigNumber;
}
