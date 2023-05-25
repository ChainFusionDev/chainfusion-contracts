import { BigNumber } from 'ethers';
import { deploySystemContracts } from './deploy/chainfusion';
import { readContractsConfig, updateContractsConfig, writeContractsConfig } from './deploy/config';
import { network } from 'hardhat';

async function main() {
  const verify = (process.env.VERIFY || '').trim().toLowerCase() === 'true';

  const contractsConfig = await readContractsConfig();
  const stakingKeys = !process.env.STAKING_KEYS ? [] : (process.env.STAKING_KEYS).trim().split(',');
  const router = contractsConfig.router ?? process.env.ROUTER_ADDRESS;

  const res = await deploySystemContracts({ displayLogs: true, parallelDeployment: true, minimalValidators: BigNumber.from(3), verify, stakingKeys, router });

  contractsConfig.networkName = network.name;
  updateContractsConfig(contractsConfig, res);
  await writeContractsConfig(contractsConfig);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
