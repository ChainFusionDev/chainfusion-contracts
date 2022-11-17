import { ethers } from 'hardhat';

async function main() {
  const bridgeAppAddress = '0x3B02fF1e626Ed7a8fd6eC5299e2C54e1421B626B';
  const mediatorAddress = '0xa513E6E4b8f2a923D98304ec87F64353C4D5C853';

  const tokenAddressBG1 = '0x56FA06B9Fd91845Ca48f49c6cB05a966a509CbBE'; // set address
  const erc20BridgeAddressBG1 = '0x672aBf9EC1c9DacB74721DE96bedFf051dFd85F6'; // set address

  const tokenAddressBG2 = '0x9564cec2C8b2Da954D2F0B1D4A777978759B945c'; // set address
  const erc20BridgeAddressBG2 = '0x98dC62904Ff76ab5eBa77B664A0d7f488Ec4AA94'; // set address

  const chainIdBG1 = '5001';
  const chainIdBG2 = '5002';

  const bridgeApp = await ethers.getContractAt('BridgeApp', bridgeAppAddress);
  const mediator = await ethers.getContractAt('ERC20BridgeMediator', mediatorAddress);

  console.log('\nSetting contract address');

  await (await bridgeApp.setContractAddress(chainIdBG1, erc20BridgeAddressBG1)).wait();
  await (await bridgeApp.setContractAddress(chainIdBG2, erc20BridgeAddressBG2)).wait();
  console.log('✅DONE');

  console.log('\nSetting token');

  await (await mediator.addToken('CFN', chainIdBG1, tokenAddressBG1)).wait();
  await (await mediator.addToken('CFN', chainIdBG2, tokenAddressBG2)).wait();
  console.log('✅DONE');
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
