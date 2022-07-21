import { ethers } from 'hardhat';
import { expect } from 'chai';
import { deploySystem } from '../utils/deploy';
import { REASON_NO_RECENT_BLOCKS } from '../utils/helpers';

describe('SlashingVoting', function () {
  it('should check if only validator can use votedWithReason()', async function () {
    const [, v2] = await ethers.getSigners();
    const initialMinimalStake = ethers.utils.parseEther('3');
    const nonse = ethers.utils.arrayify(0);

    const { slashingVoting } = await deploySystem(initialMinimalStake);

    await expect(slashingVoting.votedWithReason(v2.address, REASON_NO_RECENT_BLOCKS, nonse)).to.be.revertedWith(
      'only active validator'
    );
  });

  // it('should check if we can slash slashed validator', async function () {
  //     const [, v2,v3] = await ethers.getSigners();
  //     const initialMinimalStake = ethers.utils.parseEther('3');
  //     const value = ethers.utils.parseEther('5');
  //     const nonse = ethers.utils.arrayify(0);

  //     const { slashingVoting } = await deploySystem(initialMinimalStake);

  //     const slashingVoting2 = await ethers.getContractAt('SlashingVoting', slashingVoting.address, v2);
  //     await slashingVoting2.stake({ value: value });
  //     await slashingVoting2.slashes(v3.address);

  //     await expect(slashingVoting2.votedWithReason(v3.address, REASON_NO_RECENT_BLOCKS, nonse))
  //     .to.be.revertedWith('SlashingVoting: validator already slashed');
  // });
});
