import { ethers } from 'hardhat';

async function main() {
  const minimalStake: number = 1;
  const validators: string[] = [];

  const ValidatorStaking = await ethers.getContractFactory('ValidatorStaking');
  const validatorStaking = await ValidatorStaking.deploy();
  await validatorStaking.deployed();
  await (await validatorStaking.initialize(minimalStake)).wait();

  console.log('ValidatorStaking deployed to:', validatorStaking.address);

  const DKG = await ethers.getContractFactory('DKG');
  const dkg = await DKG.deploy();
  await (await dkg.initialize(validators)).wait();

  console.log('DKG deployed to:', dkg.address);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
