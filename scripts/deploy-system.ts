import { ethers } from 'hardhat';

async function main() {
  const withdrawalPeriod = 1;
  const minimalStake: number = 1;
  const validators: string[] = [
    '0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266',
    '0x70997970c51812dc3a010c7d01b50e0d17dc79c8',
    '0x3c44cdddb6a900fa2b585dd299e03d12fa4293bc',
    '0x90f79bf6eb2c4f870365e785982e1f101e93b906',
    '0x15d34aaf54267db7d7c367839aaf71a00a2c6a65',
  ];

  const AddressStorage = await ethers.getContractFactory('AddressStorage');
  const addressStorage = await AddressStorage.deploy();
  await addressStorage.deployed();
  await (await addressStorage.initialize(validators)).wait();

  console.log('AddressStorage deployed to:', addressStorage.address);

  const ValidatorStaking = await ethers.getContractFactory('ValidatorStaking');
  const validatorStaking = await ValidatorStaking.deploy();
  await validatorStaking.deployed();
  await (await validatorStaking.initialize(minimalStake, withdrawalPeriod, addressStorage.address)).wait();

  await addressStorage.transferOwnership(validatorStaking.address);

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
