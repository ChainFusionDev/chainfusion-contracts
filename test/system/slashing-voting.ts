import { ethers } from 'hardhat';
import { expect } from 'chai';
import { deploySystem } from '../utils/deploy';
import { BigNumber } from 'ethers';

describe('SlashingVoting', function () {
  it('should vote only by validator', async function () {
    const [, v2] = await ethers.getSigners();
    const nonce = ethers.utils.arrayify(0);
    const reason = 0;

    const { slashingVoting } = await deploySystem();

    await expect(slashingVoting.voteWithReason(v2.address, reason, nonce)).to.be.revertedWith(
      'ValidatorOwnable: only validator'
    );
  });

  it('should check if we can vote against non active validator', async function () {
    const [, v1, v2, v3] = await ethers.getSigners();
    const nonce = ethers.utils.arrayify(0);
    const reason = 0;

    const { slashingVoting, staking, minimalStake } = await deploySystem();

    const staking1 = await ethers.getContractAt('Staking', staking.address, v1);
    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);

    await staking1.stake({ value: minimalStake });
    await staking2.stake({ value: minimalStake });

    const slashingVoting1 = await ethers.getContractAt('SlashingVoting', slashingVoting.address, v1);

    await expect(slashingVoting1.voteWithReason(v3.address, reason, nonce)).to.be.revertedWith(
      'SlashingVoting: target is not active validator'
    );
  });

  it('should check if we can ban already banned validator', async function () {
    const [, v2, v3] = await ethers.getSigners();
    const nonce = ethers.utils.arrayify(0);
    const reason: number = 0;

    const { slashingVoting, staking, minimalStake } = await deploySystem();

    const slashingVoting2 = await ethers.getContractAt('SlashingVoting', slashingVoting.address, v2);
    const slashingVoting3 = await ethers.getContractAt('SlashingVoting', slashingVoting.address, v3);
    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);
    const staking3 = await ethers.getContractAt('Staking', staking.address, v3);

    await staking.stake({ value: minimalStake });
    await staking2.stake({ value: minimalStake });
    await staking3.stake({ value: minimalStake });

    await slashingVoting.voteWithReason(v3.address, reason, nonce);
    await slashingVoting2.voteWithReason(v3.address, reason, nonce);
    await expect(slashingVoting3.voteWithReason(v3.address, reason, nonce)).to.be.revertedWith(
      'SlashingVoting: validator is already banned'
    );
  });

  it('should check if we can ban slashed validator', async function () {
    const [, v1, v2, v3] = await ethers.getSigners();
    const reason: number = 0;
    const hre = require('hardhat');

    const slashingEpochs = 3;
    const slashingThreshold = 4;

    const { slashingVoting, staking, minimalStake, slashingEpochPeriod } = await deploySystem({
      slashingEpochs: BigNumber.from(slashingEpochs),
      slashingBansThresold: BigNumber.from(slashingThreshold),
    });

    const slashingVoting1 = await ethers.getContractAt('SlashingVoting', slashingVoting.address, v1);
    const slashingVoting2 = await ethers.getContractAt('SlashingVoting', slashingVoting.address, v2);

    const staking1 = await ethers.getContractAt('Staking', staking.address, v1);
    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);
    const staking3 = await ethers.getContractAt('Staking', staking.address, v3);

    await staking1.stake({ value: minimalStake });
    await staking2.stake({ value: minimalStake });
    await staking3.stake({ value: minimalStake });

    expect(await staking.isValidatorSlashed(v3.address)).to.equal(false);

    // three epochs with one ban each
    for (let i = 0; i < slashingEpochs; i++) {
      let blockInThisEpoch = BigNumber.from(await ethers.provider.getBlockNumber()).mod(slashingEpochPeriod);
      let blocksToNextEpoch = slashingEpochPeriod.sub(blockInThisEpoch);
      await hre.network.provider.send('hardhat_mine', [ethers.utils.hexValue(blocksToNextEpoch)]);

      const banNonce = ethers.utils.randomBytes(32);
      await slashingVoting1.voteWithReason(v3.address, reason, banNonce);
      await slashingVoting2.voteWithReason(v3.address, reason, banNonce);

      const epoch = await slashingVoting.currentEpoch();
      expect(await slashingVoting.getBansByEpoch(epoch, v3.address)).to.equal(1);
      expect(await staking.isValidatorSlashed(v3.address)).to.equal(false);
    }

    // second ban in last epoch
    const banNonce = ethers.utils.randomBytes(32);
    await slashingVoting1.voteWithReason(v3.address, reason, banNonce);
    await slashingVoting2.voteWithReason(v3.address, reason, banNonce);

    const epoch = await slashingVoting.currentEpoch();
    expect(await slashingVoting.getBansByEpoch(epoch, v3.address)).to.equal(2);
    expect(await staking.isValidatorSlashed(v3.address)).to.equal(true);

    await expect(slashingVoting.voteWithReason(v3.address, reason, '0x00')).to.be.revertedWith(
      'ValidatorOwnable: only validator'
    );
  });

  it('should check if we can ban twice of the same validator', async function () {
    const [, v2] = await ethers.getSigners();
    const nonce = ethers.utils.arrayify(0);
    const reason: number = 0;

    const { slashingVoting, staking, minimalStake } = await deploySystem();
    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);

    await staking.stake({ value: minimalStake });
    await staking2.stake({ value: minimalStake });

    await slashingVoting.voteWithReason(v2.address, reason, nonce);
    await expect(slashingVoting.voteWithReason(v2.address, reason, nonce)).to.be.revertedWith(
      'SlashingVoting: voter is already voted against given validator'
    );
  });
});
