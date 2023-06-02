import { ethers, config, network } from 'hardhat';
import { HttpNetworkConfig } from 'hardhat/types';
import { RelayBridge, SignerStorage } from '../../typechain';
import { Deployer } from './deployer';

const defaultBridgeDeploymentParameters: BridgeDeploymentParameters = {
  bridgeValidatorFeePool: '0x0000000000000000000000000000000000000001',

  displayLogs: false,
  parallelDeployment: false,
  verify: false,
};

export async function deployBridgeContracts(options?: BridgeDeploymentOptions): Promise<BridgeDeploymentResult> {
  const params = resolveParameters(options);
  const deployer = new Deployer(params.displayLogs, params.parallelDeployment);

  const [owner] = await ethers.getSigners();

  let initialSignerAddress = owner.address;
  const initialSignerBalance = ethers.utils.parseEther('0.05');

  if (options?.homeNetwork !== undefined && options?.homeDKGAddress !== undefined && network.name !== 'hardhat') {
    deployer.log('Receiving DKG signer address\n');

    const networkConfig = config.networks[options.homeNetwork] as HttpNetworkConfig;
    const homeProvider = new ethers.providers.JsonRpcProvider(networkConfig.url, networkConfig.chainId);
    const homeSigner = homeProvider.getSigner('0x0000000000000000000000000000000000000001');

    const dkgFactory = await ethers.getContractFactory('DKG', homeSigner);
    const dkg = dkgFactory.attach(options.homeDKGAddress);

    deployer.log(`Using DKG address: ${options.homeDKGAddress}`);

    const dkgSignerAddress = await dkg.getSignerAddress();

    if (dkgSignerAddress === '0x0000000000000000000000000000000000000000') {
      throw new Error('DKG signer address it not awailable');
    } else {
      initialSignerAddress = dkgSignerAddress;
    }

    deployer.log(`Received signer address: ${initialSignerAddress}\n`);

    if ((await ethers.provider.getBalance(initialSignerAddress)).lt(initialSignerBalance)) {
      deployer.log('Depositing initial signer balance');

      await (
        await owner.sendTransaction({
          to: initialSignerAddress,
          value: initialSignerBalance,
        })
      ).wait();

      deployer.log('Deposited initial signer balance\n');
    }
  }

  deployer.log(`🧾 Deploying ChainFusion contracts on foreign chain (${network.name} chain)\n`);

  const [deployerSigner] = await ethers.getSigners();
  const deployerAddress = await deployerSigner.getAddress();
  deployer.setNonce(await ethers.provider.getTransactionCount(deployerAddress));

  const signerStoragePromise = deployer.deploy(ethers.getContractFactory('SignerStorage'), 'SignerStorage');
  const relayBridgePromise = deployer.deploy(ethers.getContractFactory('RelayBridge'), 'RelayBridge');

  const res: BridgeDeployment = {
    signerStorage: await signerStoragePromise,
    relayBridge: await relayBridgePromise,
  };

  deployer.log('Successfully deployed ChainFusion contracts on foreign chain\n');

  deployer.log(`🔄 Initializing ChainFusion contracts on foreign chain (${network.name} chain)\n`);

  await deployer.waitPromises([
    deployer.sendTransaction(
      res.signerStorage.initialize(initialSignerAddress, deployer.getOverrides()),
      'Initializing SignerStorage'
    ),
    deployer.sendTransaction(
      res.relayBridge.initialize(res.signerStorage.address, params.bridgeValidatorFeePool, deployer.getOverrides()),
      'Initializing RelayBridge'
    ),
  ]);

  deployer.log('Successfully initialized ChainFusion contracts on foreign chain\n');

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

  if (options.bridgeValidatorFeePool !== undefined) {
    parameters.bridgeValidatorFeePool = options.bridgeValidatorFeePool;
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

  return parameters;
}

export interface BridgeDeploymentResult extends BridgeDeployment, BridgeDeploymentParameters {}

export interface BridgeDeployment {
  signerStorage: SignerStorage;
  relayBridge: RelayBridge;
}

export interface BridgeDeploymentParameters {
  bridgeValidatorFeePool: string;
  displayLogs: boolean;
  parallelDeployment: boolean;
  verify: boolean;
}

export interface BridgeDeploymentOptions {
  bridgeValidatorFeePool?: string;
  homeDKGAddress?: string;
  homeNetwork?: string;
  displayLogs?: boolean;
  parallelDeployment?: boolean;
  verify?: boolean;
}
