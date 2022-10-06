import { ethers } from 'hardhat';
import { Deployer } from './deployer';
import { MockBridgeApp } from '../../typechain';

const defaultMockBridgeAppDeploymentParameters: MockBridgeAppDeploymentParameters = {
  displayLogs: false,
  verify: false,
};

export async function deployMockBridgeAppContracts(
  options?: MockBridgeAppDeploymentOptions
): Promise<MockBridgeAppDeploymentResult> {
  const params = resolveParameters(options);
  const deployer = new Deployer(params.displayLogs);
  const relayBridge = '0x9E03A1E783d53eE799F1c9E238De851FAF27FcfF';

  deployer.log('Deploying contracts\n');

  const res: MockBridgeAppDeployment = {
    mockBridgeApp: await deployer.deploy(ethers.getContractFactory('MockBridgeApp'), 'MockBridgeApp'),
  };

  deployer.log('Successfully deployed contracts\n');

  deployer.log('Initializing contracts\n');

  await deployer.sendTransaction(res.mockBridgeApp.initialize(relayBridge), 'Initializing DEXV2Pair');

  deployer.log('Successfully initialized contracts\n');

  if (params.verify) {
    await deployer.verifyObjectValues(res);
  }

  return {
    ...res,
    ...params,
  };
}

function resolveParameters(options?: MockBridgeAppDeploymentOptions): MockBridgeAppDeploymentParameters {
  let parameters = defaultMockBridgeAppDeploymentParameters;

  if (options === undefined) {
    return parameters;
  }

  if (options.displayLogs !== undefined) {
    parameters.displayLogs = options.displayLogs;
  }

  if (options.verify !== undefined) {
    parameters.verify = options.verify;
  }

  return parameters;
}

export interface MockBridgeAppDeploymentResult extends MockBridgeAppDeployment, MockBridgeAppDeploymentParameters {}

export interface MockBridgeAppDeployment {
  mockBridgeApp: MockBridgeApp;
}

export interface MockBridgeAppDeploymentParameters {
  displayLogs: boolean;
  verify: boolean;
}

export interface MockBridgeAppDeploymentOptions {
  displayLogs?: boolean;
  verify?: boolean;
}
