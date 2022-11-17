import { deploySystemContracts } from './deploy/system';
import { readContractsConfig, updateContractsConfig, writeContractsConfig } from './deploy/config';

async function main() {
  const verify = (process.env.VERIFY || '').trim().toLowerCase() === 'true';

  const contractsConfig = await readContractsConfig();
  const stakingKeys = !process.env.STAKING_KEYS ? [] : (process.env.STAKING_KEYS).trim().split(',');
  const router = contractsConfig.router ?? process.env.ROUTER_ADDRESS;

  const res = await deploySystemContracts({ displayLogs: true, verify, stakingKeys, router });

  updateContractsConfig(contractsConfig, res);
  await writeContractsConfig(contractsConfig);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
