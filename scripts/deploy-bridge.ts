import { deployBridgeContracts } from './deploy/bridge';

async function main() {
  await deployBridgeContracts({displayLogs: true});
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
