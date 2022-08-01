import { ethers } from 'hardhat';
import { expect } from 'chai';
import { ValidatorStatusActive } from '../utils/helpers';
import { deploySystem } from '../utils/deploy';
import { hexValue } from 'ethers/lib/utils';

describe('Staking', function () {
  it('should change minimal stake', async function () {
    const initialMinimalStake = ethers.utils.parseEther('1');
    const newMinimalStake = ethers.utils.parseEther('2');

    const { staking } = await deploySystem(initialMinimalStake);

    await expect(staking.setMinimalStake(newMinimalStake))
      .to.emit(staking, 'MinimalStakeUpdated')
      .withArgs(newMinimalStake);

    expect(await staking.minimalStake()).to.equal(newMinimalStake);
  });

  it('should change setwithdrawal period', async function () {
    const initialMinimalStake = ethers.utils.parseEther('1');
    const newWithdrawalPeriod = 2;

    const { staking } = await deploySystem(initialMinimalStake);

    await expect(staking.setWithdrawalPeriod(newWithdrawalPeriod))
      .to.emit(staking, 'WithdrawalPeriodUpdated')
      .withArgs(newWithdrawalPeriod);

    expect(await staking.withdrawalPeriod()).to.equal(newWithdrawalPeriod);
  });

  it('should verify error on staking less than minimal stake', async function () {
    const initialMinimalStake = ethers.utils.parseEther('3');

    const { staking } = await deploySystem(initialMinimalStake);

    await expect(staking.stake({ value: ethers.utils.parseEther('1') })).to.be.revertedWith(
      'insufficient stake provided'
    );
  });

  it('should verify staking more than minimal stake', async function () {
    const [owner] = await ethers.getSigners();
    const initialMinimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');
    const { staking } = await deploySystem(initialMinimalStake);
    await staking.stake({ value: value });

    const { validator, stake, status } = await staking.stakes(owner.address);
    expect(validator).to.equal(owner.address);
    expect(stake).to.equal(value);
    expect(status).to.equal(ValidatorStatusActive);
  });

  it('should check if it is possible to stake several times', async function () {
    const [owner] = await ethers.getSigners();
    const initialMinimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');

    const { staking } = await deploySystem(initialMinimalStake);

    await staking.stake({ value: value });
    await staking.stake({ value: value });

    const { validator, stake, status } = await staking.stakes(owner.address);
    expect(validator).to.equal(owner.address);
    expect(stake).to.equal(ethers.utils.parseEther('10'));
    expect(status).to.equal(ValidatorStatusActive);
  });

  it('should check if only slashing voting can slash', async function () {
    const [v1, v2] = await ethers.getSigners();
    const initialMinimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');

    const { staking } = await deploySystem(initialMinimalStake);

    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);

    await staking.stake({ value: value });

    expect(await staking.getValidators()).to.deep.equal([v1.address]);

    await expect(staking2.slash(v1.address)).to.be.revertedWith('Staking: not a slashing voting');
  });

  it('should check if validatorCount is decremented after slashing', async function () {
    const [, v2, v3] = await ethers.getSigners();
    const initialMinimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');
    const nonse = ethers.utils.arrayify(0);
    const reason: number = 0;
    const secondReason: number = 1;
    const hre = require('hardhat');

    const { staking, addressStorage, slashingVoting } = await deploySystem(initialMinimalStake);

    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);
    const slashingVoting2 = await ethers.getContractAt('SlashingVoting', slashingVoting.address, v2);
    const staking3 = await ethers.getContractAt('Staking', staking.address, v3);

    await staking.stake({ value: value });
    await staking2.stake({ value: value });
    await staking3.stake({ value: value });

    await slashingVoting.voteWithReason(v3.address, reason, nonse);
    await slashingVoting2.voteWithReason(v3.address, reason, nonse);

    var blocksStep = (await slashingVoting.currentEpoch())
      .add(1)
      .mul(100)
      .sub(await ethers.provider.getBlockNumber());
    await hre.network.provider.send('hardhat_mine', [hexValue(blocksStep)]);

    await slashingVoting.voteWithReason(v3.address, secondReason, nonse);
    await slashingVoting2.voteWithReason(v3.address, secondReason, nonse);

    blocksStep = (await slashingVoting.currentEpoch())
      .add(1)
      .mul(100)
      .sub(await ethers.provider.getBlockNumber());
    await hre.network.provider.send('hardhat_mine', [hexValue(blocksStep)]);

    expect(await staking.isValidatorSlashing(v3.address)).to.equal(true);
    expect(await addressStorage.size()).to.be.equal(2);
  });

  it('should check if only validator with stake can announceWithdrawal', async function () {
    const [, v2] = await ethers.getSigners();
    const initialMinimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');

    const { staking } = await deploySystem(initialMinimalStake);

    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);

    await staking.stake({ value: value });

    await expect(staking2.announceWithdrawal(value)).to.be.revertedWith('Staking: amount must be <= to stake');
  });

  it('should check if amount more than stake', async function () {
    const initialMinimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');

    const { staking } = await deploySystem(initialMinimalStake);

    await staking.stake({ value: value });
    await expect(staking.announceWithdrawal(ethers.utils.parseEther('10'))).to.be.revertedWith(
      'amount must be <= to stake'
    );
  });

  it('should check if amount must be less or equal to stake', async function () {
    const [owner] = await ethers.getSigners();
    const initialMinimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');

    const { staking } = await deploySystem(initialMinimalStake);

    await staking.stake({ value: value });
    await staking.announceWithdrawal(value);

    const { amount } = await staking.withdrawalAnnouncements(owner.address);

    expect(amount).to.equal(value);
  });

  it('should check if only validator with stake can withdraw', async function () {
    const [, v2] = await ethers.getSigners();
    const initialMinimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');

    const { staking } = await deploySystem(initialMinimalStake);

    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);

    await staking.stake({ value: value });
    await staking.announceWithdrawal(value);

    await expect(staking2.withdraw()).to.be.revertedWith('Staking: amount must be greater than zero');
  });

  it('should check if slashed validator can not announce withdraw', async function () {
    const [, v2, v3] = await ethers.getSigners();
    const initialMinimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');
    const nonse = ethers.utils.arrayify(0);
    const reason: number = 0;
    const secondReason: number = 1;
    const hre = require('hardhat');
    const { staking, slashingVoting } = await deploySystem(initialMinimalStake);

    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);
    const slashingVoting2 = await ethers.getContractAt('SlashingVoting', slashingVoting.address, v2);
    const staking3 = await ethers.getContractAt('Staking', staking.address, v3);

    await staking.stake({ value: value });
    await staking2.stake({ value: value });
    await staking3.stake({ value: value });

    await slashingVoting.voteWithReason(v3.address, reason, nonse);
    await slashingVoting2.voteWithReason(v3.address, reason, nonse);

    var blocksStep = (await slashingVoting.currentEpoch())
      .add(1)
      .mul(100)
      .sub(await ethers.provider.getBlockNumber());
    await hre.network.provider.send('hardhat_mine', [hexValue(blocksStep)]);

    await slashingVoting.voteWithReason(v3.address, secondReason, nonse);
    await slashingVoting2.voteWithReason(v3.address, secondReason, nonse);

    expect(await staking.isValidatorSlashing(v3.address)).to.equal(true);

    await expect(staking3.announceWithdrawal(value)).to.be.revertedWith('Staking: validator is slashed');
  });

  it('should check if slashed validator can not withdraw', async function () {
    const [, v2, v3] = await ethers.getSigners();
    const initialMinimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');
    const nonse = ethers.utils.arrayify(0);
    const firstreason: number = 0;
    const secondReason: number = 1;
    const hre = require('hardhat');

    const { staking, slashingVoting } = await deploySystem(initialMinimalStake);

    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);
    const slashingVoting2 = await ethers.getContractAt('SlashingVoting', slashingVoting.address, v2);
    const staking3 = await ethers.getContractAt('Staking', staking.address, v3);

    await staking.stake({ value: value });
    await staking2.stake({ value: value });
    await staking3.stake({ value: value });

    await slashingVoting.voteWithReason(v3.address, firstreason, nonse);
    await slashingVoting2.voteWithReason(v3.address, firstreason, nonse);

    var blocksStep = (await slashingVoting.currentEpoch())
      .add(1)
      .mul(100)
      .sub(await ethers.provider.getBlockNumber());
    await hre.network.provider.send('hardhat_mine', [hexValue(blocksStep)]);

    await slashingVoting.voteWithReason(v3.address, secondReason, nonse);
    await slashingVoting2.voteWithReason(v3.address, secondReason, nonse);

    expect(await staking.isValidatorSlashing(v3.address)).to.equal(true);
    await expect(staking3.withdraw()).to.be.revertedWith('Staking: validator is slashed');
  });

  it('should user be able to withdraw his stake even if he is no longer a validator', async function () {
    const [owner] = await ethers.getSigners();
    const initialMinimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');
    const firstWithdrawal = ethers.utils.parseEther('3');
    const secondWithdrawal = ethers.utils.parseEther('2');

    const { staking } = await deploySystem(initialMinimalStake);

    await staking.stake({ value: value });
    await staking.announceWithdrawal(firstWithdrawal);
    await staking.withdraw();
    await staking.announceWithdrawal(secondWithdrawal);
    await staking.withdraw();

    const { stake } = await staking.stakes(owner.address);

    expect(stake).to.equal(0);
  });

  it('should check if withdrawalPeriod not passed', async function () {
    const initialMinimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');

    const { staking } = await deploySystem(initialMinimalStake);
    await staking.setWithdrawalPeriod(1000000);
    await staking.stake({ value: value });
    await staking.announceWithdrawal(value);

    await expect(staking.withdraw()).to.be.revertedWith('Staking: withdrawal period not passed');
  });

  it('should validator can withdraw', async function () {
    const [owner] = await ethers.getSigners();
    const initialMinimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');

    const { staking } = await deploySystem(initialMinimalStake);

    await staking.stake({ value: value });
    await staking.announceWithdrawal(value);
    await staking.withdraw();

    const { amount, time } = await staking.withdrawalAnnouncements(owner.address);

    expect(amount).to.equal(0);
    expect(time).to.equal(0);
  });

  it('should check if DKG validators are being updated after staking and slashing', async function () {
    const [owner, v1, v2, v3] = await ethers.getSigners();
    const initialMinimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');
    const valueStake = ethers.utils.parseEther('10');

    const { staking, dkg } = await deploySystem(initialMinimalStake);

    const staking1 = await ethers.getContractAt('Staking', staking.address, v1);
    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);
    const staking3 = await ethers.getContractAt('Staking', staking.address, v3);

    expect((await dkg.getCurrentValidators()).length).to.equal(0);

    await staking1.stake({ value: value });
    await staking1.stake({ value: value });
    await staking2.stake({ value: value });
    await staking3.stake({ value: value });

    expect((await dkg.getCurrentValidators()).length).to.equal(3);
    await staking1.announceWithdrawal(valueStake);
    await staking1.withdraw();
    expect((await dkg.getCurrentValidators()).length).to.equal(2);

    const { amount, time } = await staking.withdrawalAnnouncements(owner.address);

    expect(amount).to.equal(0);
    expect(time).to.equal(0);
  });

  it('should user can able to stake with the difference of the current stake and minimal stake', async function () {
    const [owner] = await ethers.getSigners();
    const initialMinimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');
    const withdrawal = ethers.utils.parseEther('3');
    const secondStake = ethers.utils.parseEther('2');

    const { staking } = await deploySystem(initialMinimalStake);

    await staking.stake({ value: value });
    await staking.announceWithdrawal(withdrawal);
    await staking.withdraw();
    await staking.stake({ value: secondStake });

    const { status } = await staking.stakes(owner.address);
    expect(status).to.equal(1);
  });

  it('should check if slashed validator can stake', async function () {
    const [, v2, v3] = await ethers.getSigners();
    const initialMinimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');
    const nonse = ethers.utils.arrayify(0);
    const reason: number = 0;
    const secondReason: number = 1;
    const hre = require('hardhat');
    const { staking, slashingVoting } = await deploySystem(initialMinimalStake);

    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);
    const slashingVoting2 = await ethers.getContractAt('SlashingVoting', slashingVoting.address, v2);
    const staking3 = await ethers.getContractAt('Staking', staking.address, v3);

    await staking.stake({ value: value });
    await staking2.stake({ value: value });
    await staking3.stake({ value: value });

    await slashingVoting.voteWithReason(v3.address, reason, nonse);
    await slashingVoting2.voteWithReason(v3.address, reason, nonse);

    var blocksStep = (await slashingVoting.currentEpoch())
      .add(1)
      .mul(100)
      .sub(await ethers.provider.getBlockNumber());
    await hre.network.provider.send('hardhat_mine', [hexValue(blocksStep)]);

    await slashingVoting.voteWithReason(v3.address, secondReason, nonse);
    await slashingVoting2.voteWithReason(v3.address, secondReason, nonse);

    expect(await staking.isValidatorSlashing(v3.address)).to.equal(true);
    await expect(staking3.stake({ value })).to.be.revertedWith('Staking: validator is slashed');
  });
});
