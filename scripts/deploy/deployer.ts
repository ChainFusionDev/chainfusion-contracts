import hre from 'hardhat';
import { BaseContract, Contract, ContractTransaction, Overrides } from 'ethers';

export interface ContractsObject {
  [key: string]: Contract;
}

export interface ContractFactory<C extends Contract> {
  // eslint-disable-next-line no-unused-vars
  deploy(overrides?: Overrides & { from?: string | Promise<string> }): Promise<C>;
}

export class Deployer {
  readonly displayLogs: boolean;
  readonly parallelDeployment: boolean;
  private nonce: number;

  constructor(displayLogs: boolean, parallelDeployment: boolean) {
    this.displayLogs = displayLogs;
    this.parallelDeployment = parallelDeployment;
    this.nonce = 0;
  }

  public setNonce(nonce: number) {
    this.nonce = nonce;
  }

  public nextNonce(): number {
    return this.nonce++;
  }

  public getOverrides(): Overrides {
    if (this.parallelDeployment) {
      return {
        nonce: this.nextNonce(),
        gasLimit: 1000000,
      };
    }

    return {};
  }

  public async deploy<C extends Contract>(factoryPromise: Promise<ContractFactory<C>>, name?: String): Promise<C> {
    if (name === undefined) {
      name = 'Contract';
    }

    const factory = await factoryPromise;
    const contract = this.parallelDeployment
      ? await factory.deploy({ nonce: this.nextNonce() })
      : await factory.deploy();

    if (this.displayLogs) {
      console.log(`Deploying ${name} ${contract.address}`);
    }

    await contract.deployed();

    if (this.displayLogs) {
      console.log(`Deployed ${name} ${contract.address}`);
    }

    return contract;
  }

  public async sendTransaction(txPromise: Promise<ContractTransaction>, message?: String): Promise<void> {
    const tx = await txPromise;

    if (this.displayLogs) {
      console.log(`${message} ${tx.hash}`);
    }

    await tx.wait();

    if (this.displayLogs) {
      console.log(`Confirmed transaction ${tx.hash}`);
    }

    return;
  }

  public async verify(contract: BaseContract): Promise<void> {
    try {
      this.log(`Verifying ${contract.address}\n`);
      await hre.run('verify:verify', { address: contract.address });
    } catch (e) {
      if (this.displayLogs) {
        this.log(`Failed to verify contract ${contract.address}: ${e}`);
      }
    }

    this.log('');
  }

  public async verifyObjectValues(obj: Object): Promise<void> {
    this.log('Verifying contracts\n');

    for (const contract of Object.values(obj)) {
      if (contract instanceof BaseContract) {
        await this.verify(contract);
      }
    }

    this.log('Successfully verified contracts\n');
  }

  public async waitTransactions(txs: Promise<ContractTransaction>[]): Promise<void> {
    for (const tx of txs) {
      const res = await tx;
      await res.wait();
    }
  }

  public async waitPromises(txs: Promise<void>[]): Promise<void> {
    for (const tx of txs) {
      await tx;
    }
  }

  public log(message?: any, ...optionalParams: any[]): void {
    if (this.displayLogs) {
      console.log(message, ...optionalParams);
    }
  }
}
