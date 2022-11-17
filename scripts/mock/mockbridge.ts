import { ethers } from 'hardhat';
import { BigNumber } from 'ethers';
import { deployBridgeContracts } from '../deploy/bridge';
import { DKG__factory } from '../../typechain';

async function main() {
  const initialSupply = BigNumber.from('10000000000000000000000000');
  const approve = '1000000000000000000000';
  const tokenFee = '1000000000000000';
  const tokenLimit = '100000000000000';
  const liquidityAmount = '100000000000000000000';

  const deployment = await deployBridgeContracts({ displayLogs: true, verify: false });

  const MockToken = await ethers.getContractFactory('MockToken');
  const mockToken = await MockToken.deploy('Mock Token', 'MOCK', initialSupply);
  await mockToken.deployed();

  console.log('Deploying token with address', mockToken.address);

  await (await mockToken.approve(deployment.liquidityPools.address, approve)).wait();
  await (await mockToken.approve(deployment.erc20Bridge.address, approve)).wait();

  await (await deployment.tokenManager.setToken(mockToken.address, '1')).wait();
  await (await deployment.feeManager.setTokenFee(mockToken.address, tokenFee, tokenFee, tokenFee)).wait();
  await (await deployment.bridgeValidatorFeePool.setLimitPerToken(mockToken.address, tokenLimit)).wait();

  await (await deployment.liquidityPools.addLiquidity(mockToken.address, liquidityAmount)).wait();

  const urlProvider = 'ws://localhost:8546';
  const provider = new ethers.providers.WebSocketProvider(urlProvider);
  const dkgAddress = '0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9';

  const dkg = new ethers.Contract(dkgAddress, DKG__factory.abi, provider);

  console.log('\nListening SignerAddressUpdated event');

  dkg.once('SignerAddressUpdated', (generation, signerAddress, event) => {
    const signer = deployment.signerStorage.getAddress();
    console.log(event);
    if (signerAddress != signer) {
      const signerBalance = '2000000000000000000';

      deployment.signerStorage.setAddress(signerAddress, { value: signerBalance });
      console.log('✅New signer: ', signerAddress, ', generation: ', generation);
    }
  });
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
