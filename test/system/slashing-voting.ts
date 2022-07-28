import { ethers } from 'hardhat';
import { expect } from 'chai';
import { deploySystem } from '../utils/deploy';

describe('SlashingVoting', function () {
  it('should vote only by validator', async function () {
    const [, v2] = await ethers.getSigners();
    const initialMinimalStake = ethers.utils.parseEther('3');
    const nonse = ethers.utils.arrayify(0);
    const reason = 0;

    const { slashingVoting } = await deploySystem(initialMinimalStake);

    await expect(slashingVoting.voteWithReason(v2.address, reason, nonse)).to.be.revertedWith(
      'SlashingVoting: only active validator'
    );
  });

  it('should check if we can vote agains non active validator', async function () {
    const [, v2] = await ethers.getSigners();
    const initialMinimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');
    const nonse = ethers.utils.arrayify(0);
    const reason = 0;

    const { slashingVoting, staking } = await deploySystem(initialMinimalStake);
    await staking.stake({ value: value });

    await expect(slashingVoting.voteWithReason(v2.address, reason, nonse)).to.be.revertedWith(
      'SlashingVoting: target is not active validator'
    );
  });

  it('should check if we can ban already baned validator', async function () {
    const [, v2, v3] = await ethers.getSigners();
    const initialMinimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');
    const nonse = ethers.utils.arrayify(0);
    const reason: number = 0;

    const { slashingVoting, staking } = await deploySystem(initialMinimalStake);

    const slashingVoting2 = await ethers.getContractAt('SlashingVoting', slashingVoting.address, v2);
    const slashingVoting3 = await ethers.getContractAt('SlashingVoting', slashingVoting.address, v3);
    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);
    const staking3 = await ethers.getContractAt('Staking', staking.address, v3);

    await staking.stake({ value: value });
    await staking2.stake({ value: value });
    await staking3.stake({ value: value });

    await slashingVoting.voteWithReason(v3.address, reason, nonse);
    await slashingVoting2.voteWithReason(v3.address, reason, nonse);
    await expect(slashingVoting3.voteWithReason(v3.address, reason, nonse)).to.be.revertedWith(
      'SlashingVoting: validator is already banned'
    );
  });

  it('should check if we can ban slashed validator', async function () {
    const [, v2, v3] = await ethers.getSigners();
    const initialMinimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');
    const nonse = ethers.utils.arrayify(0);
    const reason: number = 0;
    const secondReason: number = 1;
    const thirsdReason: number = 2;
    const hre = require('hardhat');

    const { slashingVoting, staking } = await deploySystem(initialMinimalStake);

    const slashingVoting2 = await ethers.getContractAt('SlashingVoting', slashingVoting.address, v2);
    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);
    const staking3 = await ethers.getContractAt('Staking', staking.address, v3);

    await staking.stake({ value: value });
    await staking2.stake({ value: value });
    await staking3.stake({ value: value });

    await slashingVoting.voteWithReason(v3.address, reason, nonse);
    await slashingVoting2.voteWithReason(v3.address, reason, nonse);
    await hre.network.provider.send('hardhat_mine', ['0x64']);

    await slashingVoting.voteWithReason(v3.address, secondReason, nonse);
    await slashingVoting2.voteWithReason(v3.address, secondReason, nonse);
    await hre.network.provider.send('hardhat_mine', ['0x64']);

    expect(await staking.isValidatorSlashing(v3.address)).to.equal(true);
    await expect(slashingVoting.voteWithReason(v3.address, thirsdReason, nonse)).to.be.revertedWith(
      'SlashingVoting: target is not active validator'
    );
  });

  it('should check if we can ban twice of the same validator', async function () {
    const [, v2] = await ethers.getSigners();
    const initialMinimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');
    const nonse = ethers.utils.arrayify(0);
    const reason: number = 0;

    const { slashingVoting, staking } = await deploySystem(initialMinimalStake);
    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);

    await staking.stake({ value: value });
    await staking2.stake({ value: value });

    await slashingVoting.voteWithReason(v2.address, reason, nonse);
    await expect(slashingVoting.voteWithReason(v2.address, reason, nonse)).to.be.revertedWith(
      'SlashingVoting: voter is already voted against given validator'
    );
  });
});
