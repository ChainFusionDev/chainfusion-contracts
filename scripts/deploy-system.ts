import { ethers } from 'hardhat';

async function main() {
  const withdrawalPeriod = 1;
  const minimalStake: number = 1;
  const validators: string[] = [
    '0x2915862D6994C595eB5015b38217542B443A8D73',
    '0xA06C1d549cB28dFb54490219981a3175f3DE9108',
  ];

  const AddressStorage = await ethers.getContractFactory('AddressStorage');
  const addressStorage = await AddressStorage.deploy();
  await addressStorage.deployed();
  await (await addressStorage.initialize(validators)).wait();

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
  await (await dkg.initialize(validators, validatorStaking.address)).wait();

  console.log('DKG deployed to:', dkg.address);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
