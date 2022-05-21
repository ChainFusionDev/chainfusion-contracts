import { ethers } from 'hardhat';

async function main() {
  const withdrawalPeriod = 1;
  const minimalStake: number = 1;

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
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
