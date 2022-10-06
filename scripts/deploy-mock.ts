import { deployMockBridgeAppContracts } from './deploy/mock';

async function main() {
  const verify = (process.env.VERIFY || '').trim().toLowerCase() === 'true';
  await deployMockBridgeAppContracts({displayLogs: true, verify: verify});
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
