import { Contract, ContractTransaction } from 'ethers';

export interface ContractDeployer<C extends Contract> {
  deploy(): Promise<C>;
}

export async function deployContract<C extends Contract>(
  deployerPromise: Promise<ContractDeployer<C>>,
  displayLogs?: Boolean,
  name?: String
): Promise<C> {
  if (displayLogs == undefined) {
    displayLogs = false;
  }

  if (name === undefined) {
    name = 'Contract';
  }

  if (displayLogs) {
    console.log(`Deploying ${name}`);
  }

  const deployer = await deployerPromise;
  const contract = await deployer.deploy();

  if (displayLogs) {
    console.log(`Deploying ${name} with address ${contract.address}`);
  }

  await contract.deployed();

  if (displayLogs) {
    console.log(`Deployed contract\n`);
  }

  return contract;
}

export async function waitTransaction(
  txPromise: Promise<ContractTransaction>,
  displayLogs?: Boolean,
  message?: String
): Promise<void> {
  if (displayLogs == undefined || message === undefined) {
    displayLogs = false;
  }

  if (displayLogs) {
    console.log(message);
  }

  const tx = await txPromise;

  if (displayLogs) {
    console.log(`Sending transaction ${tx.hash}`);
  }

  await tx.wait();

  if (displayLogs) {
    console.log(`Confirmed transaction\n`);
  }

  return;
}
