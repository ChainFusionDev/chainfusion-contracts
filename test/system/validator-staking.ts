import { ethers } from 'hardhat';
import { expect } from 'chai';

describe('ValidatorStaking', function () {
  it('Should change minimal stake', async function () {
    const initialminimalStake = 1;
    const newminimalStake = 2;

    const ValidatorStaking = await ethers.getContractFactory('ValidatorStaking');
    const validatorStaking = await ValidatorStaking.deploy();
    await validatorStaking.deployed();
    await (await validatorStaking.initialize(initialminimalStake)).wait();

    await validatorStaking.setMinimalStake(newminimalStake);
    expect(await validatorStaking.minimalStake()).to.equal(newminimalStake);
  });

  it('Should verify error on staking less than minimal stake', async function () {
    const initialminimalStake = ethers.utils.parseEther('3');

    const ValidatorStaking = await ethers.getContractFactory('ValidatorStaking');
    const validatorStaking = await ValidatorStaking.deploy();
    await validatorStaking.deployed();
    await (await validatorStaking.initialize(initialminimalStake)).wait();

    await expect(validatorStaking.stake({ value: ethers.utils.parseEther('1') })).to.be.revertedWith(
      'insufficient stake provided'
    );
  });

  it('Should verify staking more than minimal stake', async function () {
    const [owner] = await ethers.getSigners();
    const initialminimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');

    const ValidatorStaking = await ethers.getContractFactory('ValidatorStaking');
    const validatorStaking = await ValidatorStaking.deploy();
    await validatorStaking.deployed();
    await (await validatorStaking.initialize(initialminimalStake)).wait();

    await validatorStaking.stake({ value: ethers.utils.parseEther('5') });
    await expect(await validatorStaking.stakes(owner.address)).to.equal(value);
  });
});
