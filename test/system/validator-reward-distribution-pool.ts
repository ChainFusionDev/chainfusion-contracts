import { ethers } from 'hardhat';
import { expect } from 'chai';
import { deploySystem } from '../utils/deploy';
import { BigNumber } from 'ethers';

describe('ValidatorRewardDistributionPool', function () {
  it('should collect rewards', async function () {
    const [signer, v2, v3] = await ethers.getSigners();
    const totalReward = ethers.utils.parseEther('1');
    const rewards = ethers.utils.parseEther('0.1');
    const { validatorRewardDistributionPool, staking, minimalStake } = await deploySystem();

    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);
    const staking3 = await ethers.getContractAt('Staking', staking.address, v3);

    const validatorRewardDistributionPool2 = await ethers.getContractAt(
      'ValidatorRewardDistributionPool',
      validatorRewardDistributionPool.address,
      v2
    );
    const validatorRewardDistributionPool3 = await ethers.getContractAt(
      'ValidatorRewardDistributionPool',
      validatorRewardDistributionPool.address,
      v3
    );

    await staking2.stake({ value: minimalStake });
    await staking3.stake({ value: minimalStake });

    await (
      await signer.sendTransaction({
        to: validatorRewardDistributionPool.address,
        value: totalReward,
      })
    ).wait();

    await expect(validatorRewardDistributionPool2.collectRewards()).to.be.revertedWith(
      'ValidatorRewardDistributionPool: reward must be greater than 0'
    );
    await expect(validatorRewardDistributionPool3.collectRewards()).to.be.revertedWith(
      'ValidatorRewardDistributionPool: reward must be greater than 0'
    );

    await expect(validatorRewardDistributionPool.distributeRewards(0)).to.be.revertedWith(
      'ValidatorRewardDistributionPool: amount must be greater than 0'
    );
    await validatorRewardDistributionPool.distributeRewards(rewards);

    const validator3BalanceBefore = await ethers.provider.getBalance(v3.address);
    const validator2BalanceBefore = await ethers.provider.getBalance(v2.address);

    const validatorReward = rewards.div(BigNumber.from('2'));

    await expect(validatorRewardDistributionPool3.collectRewards())
      .to.emit(validatorRewardDistributionPool3, 'CollectRewards')
      .withArgs(v3.address, validatorReward);
    await expect(validatorRewardDistributionPool2.collectRewards())
      .to.emit(validatorRewardDistributionPool2, 'CollectRewards')
      .withArgs(v2.address, validatorReward);

    const validator3BalanceAfte = await ethers.provider.getBalance(v3.address);
    const validator2BalanceAfte = await ethers.provider.getBalance(v2.address);

    const txFee1 = validator3BalanceBefore.add(validatorReward).sub(validator3BalanceAfte);
    const txFee2 = validator2BalanceBefore.add(validatorReward).sub(validator2BalanceAfte);

    expect(validator3BalanceAfte).to.equal(validator3BalanceBefore.add(validatorReward).sub(txFee1));
    expect(validator2BalanceAfte).to.equal(validator2BalanceBefore.add(validatorReward).sub(txFee2));

    await expect(validatorRewardDistributionPool2.collectRewards()).to.be.revertedWith(
      'ValidatorRewardDistributionPool: reward must be greater than 0'
    );
    await expect(validatorRewardDistributionPool3.collectRewards()).to.be.revertedWith(
      'ValidatorRewardDistributionPool: reward must be greater than 0'
    );

    await validatorRewardDistributionPool.distributeRewards(rewards);

    const validator3StakeBefore = await staking.getStake(v3.address);
    const validator2StakeBefore = await staking.getStake(v2.address);

    await expect(validatorRewardDistributionPool3.reinvestRewards())
      .to.emit(validatorRewardDistributionPool3, 'CollectRewards')
      .withArgs(v3.address, validatorReward);
    await expect(validatorRewardDistributionPool2.reinvestRewards())
      .to.emit(validatorRewardDistributionPool2, 'CollectRewards')
      .withArgs(v2.address, validatorReward);

    const validator3StakeAfter = await staking.getStake(v3.address);
    const validator2StakeAfter = await staking.getStake(v2.address);

    expect(validator3StakeAfter).to.equal(validator3StakeBefore.add(validatorReward));
    expect(validator2StakeAfter).to.equal(validator2StakeBefore.add(validatorReward));
  });
});
