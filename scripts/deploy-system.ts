import { ethers } from 'hardhat';
import hre from 'hardhat';

const withdrawalPeriod = 60;
const minimalStake = ethers.utils.parseEther('100');

const VERIFY = (process.env.VERIFY || '').trim().toLowerCase();
const VALIDATOR_KEYS =
  process.env.VALIDATOR_KEYS === undefined ? [] : (process.env.VALIDATOR_KEYS || '').trim().split(',');

async function main() {
  console.log('\nDeploying contracts\n');

  const ContractRegistry = await ethers.getContractFactory('ContractRegistry');
  const contractRegistry = await ContractRegistry.deploy();
  await contractRegistry.deployed();

  console.log('ContractRegistry deployed to:', contractRegistry.address);

  const AddressStorage = await ethers.getContractFactory('AddressStorage');
  const addressStorage = await AddressStorage.deploy();
  await addressStorage.deployed();

  await (await addressStorage.initialize([])).wait();

  console.log('AddressStorage deployed to:', addressStorage.address);

  const Staking = await ethers.getContractFactory('Staking');
  const staking = await Staking.deploy();
  await staking.deployed();

  await contractRegistry.setContract(await staking.STAKING_KEY(), staking.address);
  await addressStorage.transferOwnership(staking.address);

  console.log('Staking deployed to:', staking.address);

  const DKG = await ethers.getContractFactory('DKG');
  const dkg = await DKG.deploy();
  await dkg.deployed();

  await contractRegistry.setContract(await dkg.DKG_KEY(), dkg.address);

  await (
    await staking.initialize(minimalStake, withdrawalPeriod, contractRegistry.address, addressStorage.address)
  ).wait();
  await (await dkg.initialize(contractRegistry.address)).wait();

  console.log('DKG deployed to:', dkg.address);

  const SupportedTokens = await ethers.getContractFactory('SupportedTokens');
  const supportedTokens = await SupportedTokens.deploy();
  await supportedTokens.deployed();

  await contractRegistry.setContract(await supportedTokens.SUPPORTED_TOKENS_KEY(), supportedTokens.address);

  console.log('SupportedTokens deployed to:', supportedTokens.address);

  if (VALIDATOR_KEYS.length > 0) {
    console.log('\nStaking\n');

    for (const privateKey of VALIDATOR_KEYS) {
      const signer = new ethers.Wallet(privateKey, ethers.provider);
      const signerStaking = await ethers.getContractAt('Staking', staking.address, signer);
      console.log('Staking', ethers.utils.formatEther(minimalStake), 'CFN for:', signer.address);
      await (await signerStaking.stake({ value: minimalStake })).wait();
    }
  }

  if (VERIFY === 'true') {
    console.log('\nVerifying contracts\n');

    // Sometimes fails, false positively, likely it's a bug in verify plugin or blockscout
    await ignoreError(hre.run('verify:verify', { address: addressStorage.address }));
    await ignoreError(hre.run('verify:verify', { address: staking.address }));
    await ignoreError(hre.run('verify:verify', { address: dkg.address }));
  }
}

async function ignoreError(callback: Promise<any>): Promise<any> {
  try {
    await callback;
  } catch {
    return;
  }
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
