import { ethers } from 'hardhat';
import { expect } from 'chai';
import { ValidatorStatusActive } from '../utils/helpers';
import { deployValidatorStaking } from '../utils/deploy';

describe('ValidatorStaking', function () {
  it('should change minimal stake', async function () {
    const initialminimalStake = ethers.utils.parseEther('1');
    const newMinimalStake = ethers.utils.parseEther('2');

    const { validatorStaking } = await deployValidatorStaking(initialminimalStake);

    expect(await validatorStaking.setMinimalStake(newMinimalStake))
      .to.emit(validatorStaking, 'MinimalStakeUpdated')
      .withArgs(newMinimalStake);
    expect(await validatorStaking.minimalStake()).to.equal(newMinimalStake);
  });

  it('should change setwithdrawal period', async function () {
    const initialminimalStake = ethers.utils.parseEther('1');
    const newWithdrawalPeriod = 2;

    const { validatorStaking } = await deployValidatorStaking(initialminimalStake);

    expect(await validatorStaking.setWithdrawalPeriod(newWithdrawalPeriod))
      .to.emit(validatorStaking, 'WithdrawalPeriodUpdated')
      .withArgs(newWithdrawalPeriod);
    expect(await validatorStaking.withdrawalPeriod()).to.equal(newWithdrawalPeriod);
  });

  it('should verify error on staking less than minimal stake', async function () {
    const initialminimalStake = ethers.utils.parseEther('3');

    const { validatorStaking } = await deployValidatorStaking(initialminimalStake);

    await expect(validatorStaking.stake({ value: ethers.utils.parseEther('1') })).to.be.revertedWith(
      'insufficient stake provided'
    );
  });

  it('should verify staking more than minimal stake', async function () {
    const [owner] = await ethers.getSigners();
    const initialminimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');

    const { validatorStaking } = await deployValidatorStaking(initialminimalStake);

    await validatorStaking.stake({ value: value });
    const { validator, stake, status } = await validatorStaking.stakes(owner.address);
    expect(validator).to.equal(owner.address);
    expect(stake).to.equal(value);
    expect(status).to.equal(ValidatorStatusActive);
  });

  it('should check if it is possible to stake several times', async function () {
    const [owner] = await ethers.getSigners();
    const initialminimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');

    const { validatorStaking } = await deployValidatorStaking(initialminimalStake);

    await validatorStaking.stake({ value: value });
    await validatorStaking.stake({ value: value });

    const { validator, stake, status } = await validatorStaking.stakes(owner.address);
    expect(validator).to.equal(owner.address);
    expect(stake).to.equal(ethers.utils.parseEther('10'));
    expect(status).to.equal(ValidatorStatusActive);
  });

  it('should check if only validator can slash', async function () {
    const [v1, v2] = await ethers.getSigners();
    const initialminimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');

    const { validatorStaking } = await deployValidatorStaking(initialminimalStake);

    const validatorStaking2 = await ethers.getContractAt('ValidatorStaking', validatorStaking.address, v2);

    await validatorStaking.stake({ value: value });

    await expect(validatorStaking2.slash(v1.address)).to.be.revertedWith('only active validator');
  });

  it('should check if slashingCount is not incremented if slash() is called several times', async function () {
    const [, v2] = await ethers.getSigners();
    const initialminimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');

    const { validatorStaking } = await deployValidatorStaking(initialminimalStake);

    const validatorStaking2 = await ethers.getContractAt('ValidatorStaking', validatorStaking.address, v2);

    await validatorStaking.stake({ value: value });
    await validatorStaking2.stake({ value: value });

    await validatorStaking.slash(v2.address);
    await validatorStaking.slash(v2.address);
    expect(await validatorStaking.slashingCount(v2.address)).to.be.equal(1);
  });

  it('should check if slashingCount is incremented if called by different validators', async function () {
    const [, v2, v3] = await ethers.getSigners();
    const initialminimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');

    const { validatorStaking } = await deployValidatorStaking(initialminimalStake);

    const validatorStaking2 = await ethers.getContractAt('ValidatorStaking', validatorStaking.address, v2);
    const validatorStaking3 = await ethers.getContractAt('ValidatorStaking', validatorStaking.address, v3);

    await validatorStaking.stake({ value: value });
    await validatorStaking2.stake({ value: value });
    await validatorStaking3.stake({ value: value });

    await validatorStaking.slash(v3.address);
    await validatorStaking2.slash(v3.address);
    expect(await validatorStaking.slashingCount(v3.address)).to.be.equal(2);
  });

  it('should check if validatorCount is decremented after slashing', async function () {
    const [, v2, v3, v4, v5] = await ethers.getSigners();
    const initialminimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');

    const { validatorStaking } = await deployValidatorStaking(initialminimalStake);

    const validatorStaking2 = await ethers.getContractAt('ValidatorStaking', validatorStaking.address, v2);
    const validatorStaking3 = await ethers.getContractAt('ValidatorStaking', validatorStaking.address, v3);
    const validatorStaking4 = await ethers.getContractAt('ValidatorStaking', validatorStaking.address, v4);
    const validatorStaking5 = await ethers.getContractAt('ValidatorStaking', validatorStaking.address, v5);

    await validatorStaking.stake({ value: value });
    await validatorStaking2.stake({ value: value });
    await validatorStaking3.stake({ value: value });
    await validatorStaking4.stake({ value: value });
    await validatorStaking5.stake({ value: value });

    await validatorStaking.slash(v5.address);
    await validatorStaking2.slash(v5.address);
    await validatorStaking3.slash(v5.address);
    await validatorStaking4.slash(v5.address);

    expect(await validatorStaking.validatorCount()).to.be.equal(3);
  });

  it('should check if only validator can announceWithdrawal', async function () {
    const [, v2] = await ethers.getSigners();
    const initialminimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');

    const { validatorStaking } = await deployValidatorStaking(initialminimalStake);

    const validatorStaking2 = await ethers.getContractAt('ValidatorStaking', validatorStaking.address, v2);

    await validatorStaking.stake({ value: value });

    await expect(validatorStaking2.announceWithdrawal(value)).to.be.revertedWith('only active validator');
  });

  it('should check if only validator can announceWithdrawal', async function () {
    const [, v2] = await ethers.getSigners();
    const initialminimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');

    const { validatorStaking } = await deployValidatorStaking(initialminimalStake);

    const validatorStaking2 = await ethers.getContractAt('ValidatorStaking', validatorStaking.address, v2);

    await validatorStaking.stake({ value: value });

    await expect(validatorStaking2.announceWithdrawal(value)).to.be.revertedWith('only active validator');
  });

  it('should check if amount more than stake', async function () {
    const initialminimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');

    const { validatorStaking } = await deployValidatorStaking(initialminimalStake);

    await validatorStaking.stake({ value: value });
    await expect(validatorStaking.announceWithdrawal(ethers.utils.parseEther('10'))).to.be.revertedWith(
      'amount must be <= to stake'
    );
  });

  it('should check if amount must be less or equal to stake', async function () {
    const [owner] = await ethers.getSigners();
    const initialminimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');

    const { validatorStaking } = await deployValidatorStaking(initialminimalStake);

    await validatorStaking.stake({ value: value });
    await validatorStaking.announceWithdrawal(value);

    const { amount } = await validatorStaking.withdrawalAnnouncements(owner.address);

    expect(amount).to.equal(value);
  });

  it('should check if only validator can withdraw', async function () {
    const [, v2] = await ethers.getSigners();
    const initialminimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');

    const { validatorStaking } = await deployValidatorStaking(initialminimalStake);

    const validatorStaking2 = await ethers.getContractAt('ValidatorStaking', validatorStaking.address, v2);

    await validatorStaking.stake({ value: value });
    await validatorStaking.announceWithdrawal(value);

    await expect(validatorStaking2.withdraw()).to.be.revertedWith('only active validator');
  });

  it('should check if withdrawalPeriod not passed', async function () {
    const initialminimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');

    const { validatorStaking } = await deployValidatorStaking(initialminimalStake);
    await validatorStaking.setWithdrawalPeriod(1000000);
    await validatorStaking.stake({ value: value });
    await validatorStaking.announceWithdrawal(value);

    await expect(validatorStaking.withdraw()).to.be.revertedWith('withdrawalPeriod not passed');
  });

  it('should validator can withdraw', async function () {
    const [owner] = await ethers.getSigners();
    const initialminimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');

    const { validatorStaking } = await deployValidatorStaking(initialminimalStake);

    await validatorStaking.stake({ value: value });
    await validatorStaking.announceWithdrawal(value);
    await validatorStaking.withdraw();

    const { amount, time } = await validatorStaking.withdrawalAnnouncements(owner.address);

    expect(amount).to.equal(0);
    expect(time).to.equal(0);
  });
});
