import { ethers } from 'hardhat';
import hre from 'hardhat';

const withdrawalPeriod = 60;
const minimalStake = ethers.utils.parseEther('100');

const VERIFY = (process.env.VERIFY || '').trim().toLowerCase();
const VALIDATOR_KEYS =
  process.env.VALIDATOR_KEYS === undefined ? [] : (process.env.VALIDATOR_KEYS || '').trim().split(',');

async function main() {
  console.log('\nDeploying contracts\n');

  const AddressStorage = await ethers.getContractFactory('AddressStorage');
  const addressStorage = await AddressStorage.deploy();
  await addressStorage.deployed();
  await (await addressStorage.initialize([])).wait();

  console.log('AddressStorage deployed to:', addressStorage.address);

  const ValidatorStaking = await ethers.getContractFactory('ValidatorStaking');
  const validatorStaking = await ValidatorStaking.deploy();
  await validatorStaking.deployed();

  await addressStorage.transferOwnership(validatorStaking.address);

  console.log('ValidatorStaking deployed to:', validatorStaking.address);

  const DKG = await ethers.getContractFactory('DKG');
  const dkg = await DKG.deploy();
  await dkg.deployed();

  await (await validatorStaking.initialize(minimalStake, withdrawalPeriod, addressStorage.address, dkg.address)).wait();
  await (await dkg.initialize(validatorStaking.address)).wait();

  console.log('DKG deployed to:', dkg.address);

  if (VALIDATOR_KEYS.length > 0) {
    console.log('\nStaking\n');

    for (const privateKey of VALIDATOR_KEYS) {
      const signer = new ethers.Wallet(privateKey, ethers.provider);
      const staking = await ethers.getContractAt('ValidatorStaking', validatorStaking.address, signer);
      console.log('Staking', ethers.utils.formatEther(minimalStake), 'CFN for:', signer.address);
      await (await staking.stake({ value: minimalStake })).wait();
    }
  }

  if (VERIFY === 'true') {
    console.log('\nVerifying contracts\n');

    // Sometimes fails, false positively, likely it's a bug in verify plugin or blockscout
    await ignoreError(hre.run('verify:verify', { address: addressStorage.address }));
    await ignoreError(hre.run('verify:verify', { address: validatorStaking.address }));
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
