import { BigNumber } from 'ethers';

import {
  Staking,
  AddressStorage,
  DKG,
  SlashingVoting,
  SupportedTokens,
  ContractRegistry,
  EventRegistry,
  TokenManager,
  LiquidityPools,
  FeeManager,
  Bridge,
  RelayBridge,
  ValidatorStorage,
} from '../../typechain';

// System contracts types

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

export interface SystemDeploymentParameters {
  minimalStake: BigNumber;
  stakeWithdrawalPeriod: BigNumber;

  slashingEpochs: BigNumber;
  slashingEpochPeriod: BigNumber;
  slashingBansThresold: BigNumber;

  dkgDeadlinePeriod: BigNumber;

  displayLogs: boolean;
}

export interface SystemDeploymentOptions {
  minimalStake?: BigNumber;
  stakeWithdrawalPeriod?: BigNumber;

  slashingEpochs?: BigNumber;
  slashingEpochPeriod?: BigNumber;
  slashingBansThresold?: BigNumber;

  dkgDeadlinePeriod?: BigNumber;

  displayLogs?: boolean;
}

// Bridge contracts types

export interface BridgeDeploymentResult extends BridgeDeployment, BridgeDeploymentParameters {}

export interface BridgeDeployment {
  validatorStorage: ValidatorStorage;
  tokenManager: TokenManager;
  feeManager: FeeManager;
  liquidityPools: LiquidityPools;
  bridge: Bridge;
  relayBridge: RelayBridge;
}

export interface BridgeDeploymentParameters {
  feePercentage: BigNumber;
  validatorRefundFee: BigNumber;
  foundationAddress: string;

  displayLogs: boolean;
}

export interface BridgeDeploymentOptions {
  feePercentage?: BigNumber;
  validatorRefundFee?: BigNumber;
  foundationAddress?: string;

  displayLogs?: boolean;
  deployMocks?: boolean;
}
