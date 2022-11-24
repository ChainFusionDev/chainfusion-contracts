import { ethers, network } from 'hardhat';
import { deployBridgeContracts } from './deploy/chain';
import { readChainContractsConfig, readContractsConfig, updateContractsConfig, writeChainContractsConfig, processBridgeConfig } from './deploy/config';

async function main() {
  const verify = (process.env.VERIFY || '').trim().toLowerCase() === 'true';
  const chainId = network.config.chainId ?? 1;
  const blockNumber = await ethers.provider.getBlockNumber();

  const homeContractsConfig = await readContractsConfig();
  const homeNetwork = homeContractsConfig.networkName;
  const dkgAddress = homeContractsConfig.dkg;

  const contractsConfig = await readChainContractsConfig(chainId);

  const res = await deployBridgeContracts({
    homeNetwork: homeNetwork,
    homeDKGAddress: dkgAddress,
    displayLogs: true,
    verify: verify,
  });

  updateContractsConfig(contractsConfig, res);
  await writeChainContractsConfig(chainId, contractsConfig);

  await processBridgeConfig(res.relayBridge.address, res.signerStorage.address, blockNumber)
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
