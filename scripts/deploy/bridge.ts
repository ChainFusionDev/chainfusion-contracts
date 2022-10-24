import { BigNumber } from 'ethers';
import { ethers } from 'hardhat';
import {
  ERC20Bridge,
  FeeManager,
  BridgeValidatorFeePool,
  LiquidityPools,
  RelayBridge,
  SignerStorage,
  TokenManager,
} from '../../typechain';
import { Deployer } from './deployer';

const defaultBridgeDeploymentParameters: BridgeDeploymentParameters = {
  feePercentage: BigNumber.from('10000000000000000'),
  validatorRefundFee: BigNumber.from('10000000000000000'),
  foundationAddress: '0xd13F66863ED91704e386C57501F00b5307CAbA18',
  bridgeAppAddress: '0x3B02fF1e626Ed7a8fd6eC5299e2C54e1421B626B',
  displayLogs: false,
  verify: false,
};

export async function deployBridgeContracts(options?: BridgeDeploymentOptions): Promise<BridgeDeploymentResult> {
  const params = resolveParameters(options);
  const deployer = new Deployer(params.displayLogs);

  const [validator] = await ethers.getSigners();

  deployer.log('Deploying contracts\n');

  const res: BridgeDeployment = {
    signerStorage: await deployer.deploy(ethers.getContractFactory('SignerStorage'), 'SignerStorage'),
    tokenManager: await deployer.deploy(ethers.getContractFactory('TokenManager'), 'TokenManager'),
    feeManager: await deployer.deploy(ethers.getContractFactory('FeeManager'), 'FeeManager'),
    bridgeValidatorFeePool: await deployer.deploy(
      ethers.getContractFactory('BridgeValidatorFeePool'),
      'BridgeValidatorFeePool'
    ),
    liquidityPools: await deployer.deploy(ethers.getContractFactory('LiquidityPools'), 'LiquidityPools'),
    erc20Bridge: await deployer.deploy(ethers.getContractFactory('ERC20Bridge'), 'ERC20Bridge'),
    relayBridge: await deployer.deploy(ethers.getContractFactory('RelayBridge'), 'RelayBridge'),
  };

  deployer.log('Successfully deployed contracts\n');

  deployer.log('Initializing contracts\n');

  await deployer.sendTransaction(res.signerStorage.initialize(validator.address), 'Initializing ValidatorStorage');
  await deployer.sendTransaction(res.tokenManager.initialize(res.signerStorage.address), 'Initializing TokenManager');

  await deployer.sendTransaction(
    res.feeManager.initialize(
      res.signerStorage.address,
      res.liquidityPools.address,
      params.foundationAddress,
      res.bridgeValidatorFeePool.address,
      params.validatorRefundFee
    ),
    'Initializing FeeManager'
  );

  await deployer.sendTransaction(
    res.liquidityPools.initialize(
      res.signerStorage.address,
      res.tokenManager.address,
      res.erc20Bridge.address,
      res.feeManager.address,
      params.feePercentage
    ),
    'Initializing LiquidityPools'
  );

  await deployer.sendTransaction(
    res.erc20Bridge.initialize(
      res.relayBridge.address,
      res.signerStorage.address,
      res.tokenManager.address,
      res.liquidityPools.address,
      res.feeManager.address,
      params.bridgeAppAddress
    ),
    'Initializing ERC20Bridge'
  );

  await deployer.sendTransaction(
    res.relayBridge.initialize(res.signerStorage.address, res.bridgeValidatorFeePool.address),
    'Initializing RelayBridge'
  );
  await deployer.sendTransaction(
    res.bridgeValidatorFeePool.initialize(res.signerStorage.address, res.erc20Bridge.address, validator.address),
    'Initializing BridgeValidatorFeePool'
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

  return parameters;
}

export interface BridgeDeploymentResult extends BridgeDeployment, BridgeDeploymentParameters {}

export interface BridgeDeployment {
  signerStorage: SignerStorage;
  tokenManager: TokenManager;
  feeManager: FeeManager;
  bridgeValidatorFeePool: BridgeValidatorFeePool;
  liquidityPools: LiquidityPools;
  erc20Bridge: ERC20Bridge;
  relayBridge: RelayBridge;
}

export interface BridgeDeploymentParameters {
  feePercentage: BigNumber;
  validatorRefundFee: BigNumber;
  foundationAddress: string;
  bridgeAppAddress: string;
  displayLogs: boolean;
  verify: boolean;
}

export interface BridgeDeploymentOptions {
  feePercentage?: BigNumber;
  validatorRefundFee?: BigNumber;
  foundationAddress?: string;
  bridgeAppAddress?: string;
  displayLogs?: boolean;
  verify?: boolean;
  deployMocks?: boolean;
}
