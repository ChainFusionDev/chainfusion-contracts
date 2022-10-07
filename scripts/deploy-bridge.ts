import { deployBridgeContracts } from './deploy/bridge';

async function main() {
  const verify = (process.env.VERIFY || '').trim().toLowerCase() === 'true';
  await deployBridgeContracts({ displayLogs: true, verify: verify });
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
