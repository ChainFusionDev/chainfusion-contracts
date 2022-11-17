import { ethers } from 'hardhat';

async function main() {
  const tokenAddressBG2 = '0x9564cec2C8b2Da954D2F0B1D4A777978759B945c'; // set address
  const erc20BridgeAddressBG2 = '0x98dC62904Ff76ab5eBa77B664A0d7f488Ec4AA94'; // set address

  const [receiver] = await ethers.getSigners();

  const chainIdBG1 = '5001';

  const erc20bridge = await ethers.getContractAt('ERC20Bridge', erc20BridgeAddressBG2);

  console.log('Depositing');

  await (await erc20bridge.deposit(tokenAddressBG2, chainIdBG1, receiver.address, '1000000000000000000')).wait();
  console.log('✅DONE');
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
