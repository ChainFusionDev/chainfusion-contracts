import { Contract, ContractTransaction } from 'ethers';

export interface ContractDeployer<C extends Contract> {
  deploy(): Promise<C>;
}

export async function deployContract<C extends Contract>(
  deployerPromise: Promise<ContractDeployer<C>>,
  showLogs?: Boolean,
  name?: String
): Promise<C> {
  if (showLogs == undefined) {
    showLogs = false;
  }

  if (name === undefined) {
    name = 'Contract';
  }

  if (showLogs) {
    console.log(`Deploying ${name}`);
  }

  const deployer = await deployerPromise;
  const contract = await deployer.deploy();

  if (showLogs) {
    console.log(`Deploying ${name} with address ${contract.address}`);
  }

  await contract.deployed();

  if (showLogs) {
    console.log(`Deployed contract\n`);
  }

  return contract;
}

export async function waitTransaction(
  txPromise: Promise<ContractTransaction>,
  showLogs?: Boolean,
  message?: String
): Promise<void> {
  if (showLogs == undefined || message === undefined) {
    showLogs = false;
  }

  if (showLogs) {
    console.log(message);
  }

  const tx = await txPromise;

  if (showLogs) {
    console.log(`Sending transaction ${tx.hash}`);
  }

  await tx.wait();

  if (showLogs) {
    console.log(`Confirmed transaction\n`);
  }

  return;
}
