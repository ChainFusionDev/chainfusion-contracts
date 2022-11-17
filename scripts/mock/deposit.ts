import { ethers } from 'hardhat';

async function main() {
  const tokenAddressBG1 = '0x56FA06B9Fd91845Ca48f49c6cB05a966a509CbBE'; // set address
  const erc20BridgeAddressBG1 = '0x672aBf9EC1c9DacB74721DE96bedFf051dFd85F6'; // set address

  const [receiver] = await ethers.getSigners();

  const chainIdBG2 = '5002';

  const erc20bridge = await ethers.getContractAt('ERC20Bridge', erc20BridgeAddressBG1);

  console.log('Depositing');

  await (await erc20bridge.deposit(tokenAddressBG1, chainIdBG2, receiver.address, '1000000000000000000')).wait();
  console.log('✅DONE');
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
