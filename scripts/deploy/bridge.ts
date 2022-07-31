import { BigNumber } from 'ethers';
import { ethers } from 'hardhat';
import { BridgeDeployment, BridgeDeploymentOptions, BridgeDeploymentParameters, BridgeDeploymentResult } from './types';
import { deployContract, waitTransaction } from './utils';

const defaultBridgeDeploymentParameters: BridgeDeploymentParameters = {
  feePercentage: BigNumber.from('10000000000000000'),
  validatorRefundFee: BigNumber.from('10000000000000000'),
  foundationAddress: '0xd13F66863ED91704e386C57501F00b5307CAbA18',

  displayLogs: false,
};

export async function deployBridgeContracts(options?: BridgeDeploymentOptions): Promise<BridgeDeploymentResult> {
  const params = resolveParameters(options);

  const [validator] = await ethers.getSigners();

  if (params.displayLogs) {
    console.log('Deploying contracts\n');
  }

  const res: BridgeDeployment = {
    validatorStorage: await deployContract(
      ethers.getContractFactory('ValidatorStorage'),
      params.displayLogs,
      'ValidatorStorage'
    ),
    tokenManager: await deployContract(ethers.getContractFactory('TokenManager'), params.displayLogs, 'TokenManager'),
    feeManager: await deployContract(ethers.getContractFactory('FeeManager'), params.displayLogs, 'FeeManager'),
    liquidityPools: await deployContract(
      ethers.getContractFactory('LiquidityPools'),
      params.displayLogs,
      'LiquidityPools'
    ),
    bridge: await deployContract(ethers.getContractFactory('Bridge'), params.displayLogs, 'Bridge'),
    relayBridge: await deployContract(ethers.getContractFactory('RelayBridge'), params.displayLogs, 'RelayBridge'),
  };

  if (params.displayLogs) {
    console.log('Successfully deployed contracts\n');
  }

  if (params.displayLogs) {
    console.log('Initializing contracts\n');
  }

  await waitTransaction(
    res.validatorStorage.initialize(validator.address),
    params.displayLogs,
    'Initializing ValidatorStorage'
  );
  await waitTransaction(
    res.tokenManager.initialize(res.validatorStorage.address),
    params.displayLogs,
    'Initializing TokenManager'
  );

  await waitTransaction(
    res.feeManager.initialize(
      res.validatorStorage.address,
      res.liquidityPools.address,
      params.foundationAddress,
      params.validatorRefundFee
    ),
    params.displayLogs,
    'Initializing FeeManager'
  );

  await waitTransaction(
    res.liquidityPools.initialize(
      res.validatorStorage.address,
      res.tokenManager.address,
      res.bridge.address,
      res.feeManager.address,
      params.feePercentage
    ),
    params.displayLogs,
    'Initializing LiquidityPools'
  );

  await waitTransaction(
    res.bridge.initialize(
      res.validatorStorage.address,
      res.tokenManager.address,
      res.liquidityPools.address,
      res.feeManager.address
    ),
    params.displayLogs,
    'Initializing Bridge'
  );

  await waitTransaction(
    res.relayBridge.initialize(res.validatorStorage.address),
    params.displayLogs,
    'Initializing RelayBridge'
  );

  if (params.displayLogs) {
    console.log('Successfully initialized contracts\n');
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

  return parameters;
}
