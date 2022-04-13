import { ethers } from 'hardhat';

async function main() {
  const minimalStake = 1;

  const ValidatorStaking = await ethers.getContractFactory('ValidatorStaking');
  const validatorStaking = await ValidatorStaking.deploy();
  await validatorStaking.deployed();
  await (await validatorStaking.initialize(minimalStake)).wait();

  console.log('ValidatorStaking deployed to:', validatorStaking.address);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
