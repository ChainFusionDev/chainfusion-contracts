import hre from 'hardhat';
import { ethers } from 'hardhat';
import { expect } from 'chai';
import { ValidatorStatusActive } from '../utils/helpers';
import { deploySystem } from '../utils/deploy';

describe('Staking', function () {
  it('should change minimal stake', async function () {
    const newMinimalStake = ethers.utils.parseEther('2');
    const { staking } = await deploySystem();

    await expect(staking.setMinimalStake(newMinimalStake))
      .to.emit(staking, 'MinimalStakeUpdated')
      .withArgs(newMinimalStake);

    expect(await staking.minimalStake()).to.equal(newMinimalStake);
  });

  it('should set withdrawal period', async function () {
    const newWithdrawalPeriod = 2;
    const { staking } = await deploySystem();

    await expect(staking.setWithdrawalPeriod(newWithdrawalPeriod))
      .to.emit(staking, 'WithdrawalPeriodUpdated')
      .withArgs(newWithdrawalPeriod);

    expect(await staking.withdrawalPeriod()).to.equal(newWithdrawalPeriod);
  });

  it('should verify error on staking less than minimal stake', async function () {
    const { staking, minimalStake } = await deploySystem();

    await expect(staking.stake({ value: minimalStake.sub(1) })).to.be.revertedWith('insufficient stake provided');
  });

  it('should verify staking more than minimal stake', async function () {
    const [owner] = await ethers.getSigners();
    const additionalStake = ethers.utils.parseEther('5');
    const { staking, minimalStake } = await deploySystem();
    await staking.stake({ value: minimalStake.add(additionalStake) });

    const { validator, stake, status } = await staking.stakes(owner.address);
    expect(validator).to.equal(owner.address);
    expect(stake).to.equal(minimalStake.add(additionalStake));
    expect(status).to.equal(ValidatorStatusActive);
  });

  it('should check if it is possible to stake several times', async function () {
    const [owner] = await ethers.getSigners();
    const { staking, minimalStake } = await deploySystem();

    await staking.stake({ value: minimalStake });
    await staking.stake({ value: minimalStake });

    const { validator, stake, status } = await staking.stakes(owner.address);
    expect(validator).to.equal(owner.address);
    expect(stake).to.equal(minimalStake.mul(2));
    expect(status).to.equal(ValidatorStatusActive);
  });

  it('should check if only validator can slash', async function () {
    const [v1, v2] = await ethers.getSigners();
    const { staking, minimalStake } = await deploySystem();

    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);

    await staking.stake({ value: minimalStake });

    expect(await staking.getValidators()).to.deep.equal([v1.address]);

    await expect(staking2.slash(v1.address)).to.be.revertedWith('only active validator');
  });

  it('should check if slashingCount is not incremented if slash() is called several times', async function () {
    const [, v2] = await ethers.getSigners();
    const { staking, minimalStake } = await deploySystem();

    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);

    await staking.stake({ value: minimalStake });
    await staking2.stake({ value: minimalStake });

    await staking.slash(v2.address);
    await staking.slash(v2.address);
    expect(await staking.slashingCount(v2.address)).to.be.equal(1);
  });

  it('should check if slash() already slashed validator', async function () {
    const [, v2, v3] = await ethers.getSigners();
    const { staking, minimalStake } = await deploySystem();

    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);
    const staking3 = await ethers.getContractAt('Staking', staking.address, v3);

    await staking.stake({ value: minimalStake });
    await staking2.stake({ value: minimalStake });
    await staking3.stake({ value: minimalStake });

    await staking.slash(v2.address);
    await staking3.slash(v2.address);

    await expect(staking3.slash(v2.address)).to.be.revertedWith('Staking: validator is already slashed');
  });

  it('should check if slashingCount is incremented if called by different validators', async function () {
    const [, v2, v3] = await ethers.getSigners();
    const { staking, minimalStake } = await deploySystem();

    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);
    const staking3 = await ethers.getContractAt('Staking', staking.address, v3);

    await staking.stake({ value: minimalStake });
    await staking2.stake({ value: minimalStake });
    await staking3.stake({ value: minimalStake });

    await staking.slash(v3.address);
    await staking2.slash(v3.address);
    expect(await staking.slashingCount(v3.address)).to.be.equal(2);
  });

  it('should check if validatorCount is decremented after slashing', async function () {
    const [, v2, v3, v4, v5] = await ethers.getSigners();
    const { staking, addressStorage, minimalStake } = await deploySystem();

    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);
    const staking3 = await ethers.getContractAt('Staking', staking.address, v3);
    const staking4 = await ethers.getContractAt('Staking', staking.address, v4);
    const staking5 = await ethers.getContractAt('Staking', staking.address, v5);

    await staking.stake({ value: minimalStake });
    await staking2.stake({ value: minimalStake });
    await staking3.stake({ value: minimalStake });
    await staking4.stake({ value: minimalStake });
    await staking5.stake({ value: minimalStake });

    await staking.slash(v5.address);
    await staking2.slash(v5.address);
    await staking3.slash(v5.address);

    expect(await addressStorage.size()).to.be.equal(4);
  });

  it('should check if only validator with stake can announceWithdrawal', async function () {
    const [, v2] = await ethers.getSigners();
    const { staking, minimalStake } = await deploySystem();

    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);

    await staking.stake({ value: minimalStake });

    await expect(staking2.announceWithdrawal(minimalStake)).to.be.revertedWith('Staking: amount must be <= to stake');
  });

  it('should check if amount more than stake', async function () {
    const { staking, minimalStake } = await deploySystem();

    await staking.stake({ value: minimalStake });
    await expect(staking.announceWithdrawal(minimalStake.add(ethers.utils.parseEther('1')))).to.be.revertedWith(
      'Staking: amount must be <= to stake'
    );
  });

  it('should check if amount must be less or equal to stake', async function () {
    const [owner] = await ethers.getSigners();
    const { staking, minimalStake } = await deploySystem();

    await staking.stake({ value: minimalStake });
    await staking.announceWithdrawal(minimalStake);

    const { amount } = await staking.withdrawalAnnouncements(owner.address);

    expect(amount).to.equal(minimalStake);
  });

  it('should check if only validator with stake can withdraw', async function () {
    const [, v2] = await ethers.getSigners();
    const { staking, minimalStake, stakeWithdrawalPeriod } = await deploySystem();

    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);

    await staking.stake({ value: minimalStake });
    await staking.announceWithdrawal(minimalStake);
    await hre.network.provider.send('hardhat_mine', [stakeWithdrawalPeriod.toHexString(), '0x1']);

    await expect(staking2.withdraw()).to.be.revertedWith('Staking: amount must be greater than zero');
  });

  it('should check if slashed validator can not withdraw', async function () {
    const [, v2, v3] = await ethers.getSigners();

    const { staking, minimalStake } = await deploySystem();

    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);
    const staking3 = await ethers.getContractAt('Staking', staking.address, v3);

    await staking.stake({ value: minimalStake });
    await staking2.stake({ value: minimalStake });
    await staking3.stake({ value: minimalStake });

    await staking.slash(v2.address);
    await staking3.slash(v2.address);

    await expect(staking2.announceWithdrawal(minimalStake)).to.be.revertedWith('Staking: validator is slashed');
    await expect(staking2.withdraw()).to.be.revertedWith('Staking: validator is slashed');
  });

  it('should be able to withdraw even if no longer a validator', async function () {
    const [owner] = await ethers.getSigners();
    const minimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');
    const firstWithdrawal = ethers.utils.parseEther('3');
    const secondWithdrawal = ethers.utils.parseEther('2');

    const { staking, stakeWithdrawalPeriod } = await deploySystem({ minimalStake });

    await staking.stake({ value: value });

    await staking.announceWithdrawal(firstWithdrawal);
    await hre.network.provider.send('hardhat_mine', [stakeWithdrawalPeriod.toHexString(), '0x1']);
    await staking.withdraw();

    await staking.announceWithdrawal(secondWithdrawal);
    await hre.network.provider.send('hardhat_mine', [stakeWithdrawalPeriod.toHexString(), '0x1']);
    await staking.withdraw();

    const { stake } = await staking.stakes(owner.address);

    expect(stake).to.equal(0);
  });

  it('should check if withdrawalPeriod not passed', async function () {
    const { staking, minimalStake } = await deploySystem();
    await staking.setWithdrawalPeriod(1000000);
    await staking.stake({ value: minimalStake });
    await staking.announceWithdrawal(minimalStake);

    await expect(staking.withdraw()).to.be.revertedWith('Staking: withdrawal period not passed');
  });

  it('should withdraw', async function () {
    const [owner] = await ethers.getSigners();
    const { staking, minimalStake, stakeWithdrawalPeriod } = await deploySystem();

    await staking.stake({ value: minimalStake });
    await staking.announceWithdrawal(minimalStake);
    await hre.network.provider.send('hardhat_mine', [stakeWithdrawalPeriod.toHexString(), '0x1']);
    await staking.withdraw();

    const { amount, time } = await staking.withdrawalAnnouncements(owner.address);

    expect(amount).to.equal(0);
    expect(time).to.equal(0);
  });

  it('should check if DKG validators are being updated after staking and slashing', async function () {
    const [owner, v1, v2, v3] = await ethers.getSigners();
    const minimalStake = ethers.utils.parseEther('3');
    const stake = ethers.utils.parseEther('5');
    const withdrawStake = ethers.utils.parseEther('10');

    const { staking, dkg, stakeWithdrawalPeriod } = await deploySystem({ minimalStake });
    const staking1 = await ethers.getContractAt('Staking', staking.address, v1);
    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);
    const staking3 = await ethers.getContractAt('Staking', staking.address, v3);

    expect((await dkg.getCurrentValidators()).length).to.equal(0);

    await staking1.stake({ value: stake });
    await staking1.stake({ value: stake });
    await staking2.stake({ value: stake });
    await staking3.stake({ value: stake });

    expect((await dkg.getCurrentValidators()).length).to.equal(3);
    await staking1.announceWithdrawal(withdrawStake);
    await hre.network.provider.send('hardhat_mine', [stakeWithdrawalPeriod.toHexString(), '0x1']);
    await staking1.withdraw();
    expect((await dkg.getCurrentValidators()).length).to.equal(2);

    const { amount, time } = await staking.withdrawalAnnouncements(owner.address);

    expect(amount).to.equal(0);
    expect(time).to.equal(0);
  });

  it('should user can able to stake with the difference of the current stake and minimal stake', async function () {
    const [owner] = await ethers.getSigners();
    const minimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');
    const withdrawal = ethers.utils.parseEther('3');
    const secondStake = ethers.utils.parseEther('2');

    const { staking, stakeWithdrawalPeriod } = await deploySystem({ minimalStake });

    await staking.stake({ value: value });
    await staking.announceWithdrawal(withdrawal);
    await hre.network.provider.send('hardhat_mine', [stakeWithdrawalPeriod.toHexString(), '0x1']);
    await staking.withdraw();
    await staking.stake({ value: secondStake });

    const { status } = await staking.stakes(owner.address);
    expect(status).to.equal(1);
  });
});
