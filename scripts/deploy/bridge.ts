import { BigNumber } from 'ethers';
import { ethers } from 'hardhat';
import { RelayBridge, SignerStorage } from '../../typechain';
import { Deployer } from './deployer';

const defaultBridgeDeploymentParameters: BridgeDeploymentParameters = {
  feePercentage: BigNumber.from('10000000000000000'),
  validatorRefundFee: BigNumber.from('10000000000000000'),
  foundationAddress: '0x0A97ddbac3A97693a75C727D4D8D6Ab4F5a22d43',
  bridgeAppAddress: '0x3B02fF1e626Ed7a8fd6eC5299e2C54e1421B626B',
  bridgeValidatorFeePool: '0x0000000000000000000000000000000000000001',

  displayLogs: false,
  verify: false,
};

export async function deployBridgeContracts(options?: BridgeDeploymentOptions): Promise<BridgeDeploymentResult> {
  const params = resolveParameters(options);
  const deployer = new Deployer(params.displayLogs);

  const [owner] = await ethers.getSigners();

  deployer.log('Deploying contracts\n');

  const res: BridgeDeployment = {
    signerStorage: await deployer.deploy(ethers.getContractFactory('SignerStorage'), 'SignerStorage'),
    relayBridge: await deployer.deploy(ethers.getContractFactory('RelayBridge'), 'RelayBridge'),
  };

  deployer.log('Successfully deployed contracts\n');

  deployer.log('Initializing contracts\n');

  await deployer.sendTransaction(res.signerStorage.initialize(owner.address), 'Initializing SignerStorage');

  await deployer.sendTransaction(
    res.relayBridge.initialize(res.signerStorage.address, params.bridgeValidatorFeePool),
    'Initializing RelayBridge'
  );

  deployer.log('Successfully initialized contracts\n');

  if (params.verify) {
    await deployer.verifyObjectValues(res);
  }

  return {
    ...res,
    ...params,
  };
}

function resolveParameters(options?: BridgeDeploymentOptions): BridgeDeploymentParameters {
  let parameters = defaultBridgeDeploymentParameters;

  if (options === undefined) {
    return parameters;
  }

  if (options.feePercentage !== undefined) {
    parameters.feePercentage = options.feePercentage;
  }

  if (options.validatorRefundFee !== undefined) {
    parameters.validatorRefundFee = options.validatorRefundFee;
  }

  if (options.foundationAddress !== undefined) {
    parameters.foundationAddress = options.foundationAddress;
  }

  if (options.displayLogs !== undefined) {
    parameters.displayLogs = options.displayLogs;
  }

  if (options.verify !== undefined) {
    parameters.verify = options.verify;
  }

  if (options.bridgeAppAddress !== undefined) {
    parameters.bridgeAppAddress = options.bridgeAppAddress;
  }

  if (options.bridgeValidatorFeePool !== undefined) {
    parameters.bridgeValidatorFeePool = options.bridgeValidatorFeePool;
  }

  return parameters;
}

export interface BridgeDeploymentResult extends BridgeDeployment, BridgeDeploymentParameters {}

export interface BridgeDeployment {
  signerStorage: SignerStorage;
  relayBridge: RelayBridge;
}

export interface BridgeDeploymentParameters {
  feePercentage: BigNumber;
  validatorRefundFee: BigNumber;
  foundationAddress: string;
  bridgeAppAddress: string;
  bridgeValidatorFeePool: string;
  displayLogs: boolean;
  verify: boolean;
}

export interface BridgeDeploymentOptions {
  feePercentage?: BigNumber;
  validatorRefundFee?: BigNumber;
  foundationAddress?: string;
  bridgeAppAddress?: string;
  bridgeValidatorFeePool?: string;
  displayLogs?: boolean;
  verify?: boolean;
  deployMocks?: boolean;
}
