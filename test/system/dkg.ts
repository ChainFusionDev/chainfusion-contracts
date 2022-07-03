import { ethers } from 'hardhat';
import { expect } from 'chai';
import { deploySystem } from '../utils/deploy';

describe('DKG', function () {
  it('should broadcast all rounds', async function () {
    const generation = 0;
    const signerAddress = '0x0000000000000000000000000000000000000001';
    const data1 = ethers.utils.keccak256([1]);
    const data2 = ethers.utils.keccak256([2]);
    const data3 = ethers.utils.keccak256([3]);

    const [, v1, v2] = await ethers.getSigners();
    const initialMinimalStake = ethers.utils.parseEther('3');

    const { dkg, validatorStaking } = await deploySystem(initialMinimalStake);

    expect(await dkg.getGenerationsCount()).to.equal(0);

    const validatorStaking1 = await ethers.getContractAt('ValidatorStaking', validatorStaking.address, v1);
    const validatorStaking2 = await ethers.getContractAt('ValidatorStaking', validatorStaking.address, v2);

    await dkg.setValidatorStaking(validatorStaking.address);

    await expect(dkg.roundBroadcast(generation, 1, data1)).to.be.revertedWith('DKG: not a validator');

    expect(await dkg.isValidator(generation, v1.address)).to.equal(false);
    expect(await dkg.isValidator(generation, v2.address)).to.equal(false);
    expect(await dkg.getValidators(generation)).to.deep.equal([]);
    expect(await dkg.getValidatorsCount(generation)).to.equal(0);

    await validatorStaking1.stake({ value: initialMinimalStake });
    await validatorStaking2.stake({ value: initialMinimalStake });

    expect(await dkg.getGenerationsCount()).to.equal(1);
    expect(await dkg.isValidator(generation, v1.address)).to.equal(true);
    expect(await dkg.isValidator(generation, v2.address)).to.equal(true);
    expect(await dkg.getValidators(generation)).to.deep.equal([v1.address, v2.address]);
    expect(await dkg.getValidatorsCount(generation)).to.equal(2);

    const dkgV1 = await ethers.getContractAt('DKG', dkg.address, v1);
    const dkgV2 = await ethers.getContractAt('DKG', dkg.address, v2);

    await expect(dkgV1.roundBroadcast(generation, 2, data2)).to.be.revertedWith('DKG: round was not filled');
    await expect(dkgV2.roundBroadcast(generation, 2, data2)).to.be.revertedWith('DKG: round was not filled');

    // round1 - v1

    await expect(dkgV1.roundBroadcast(generation, 1, data1))
      .to.emit(dkgV1, 'RoundDataProvided')
      .withArgs(generation, 1, v1.address);

    await expect(dkgV1.roundBroadcast(generation, 1, data1)).to.be.revertedWith('DKG: round data already provided');

    expect(await dkg.getRoundBroadcastData(generation, 1, v1.address)).to.equal(data1);
    expect(await dkg.getRoundBroadcastCount(generation, 1)).to.equal(1);
    expect(await dkg.getRoundBroadcastCount(generation, 2)).to.equal(0);
    expect(await dkg.getRoundBroadcastCount(generation, 3)).to.equal(0);
    expect(await dkg.isRoundFilled(generation, 1)).to.equal(false);

    // round1 - v2

    await expect(dkgV2.roundBroadcast(generation, 1, data1)).to.emit(dkgV2, 'RoundDataFilled').withArgs(generation, 1);

    expect(await dkg.getRoundBroadcastData(generation, 1, v2.address)).to.equal(data1);
    expect(await dkg.getRoundBroadcastCount(generation, 1)).to.equal(2);
    expect(await dkg.getRoundBroadcastCount(generation, 2)).to.equal(0);
    expect(await dkg.getRoundBroadcastCount(generation, 3)).to.equal(0);
    expect(await dkg.isRoundFilled(generation, 1)).to.equal(true);

    // round2 - v1

    await expect(dkgV1.roundBroadcast(generation, 3, data2)).to.be.revertedWith('DKG: round was not filled');
    await expect(dkgV1.roundBroadcast(generation, 2, data2))
      .to.emit(dkgV1, 'RoundDataProvided')
      .withArgs(generation, 2, v1.address);

    expect(await dkg.getRoundBroadcastData(generation, 2, v1.address)).to.equal(data2);
    expect(await dkg.getRoundBroadcastCount(generation, 1)).to.equal(2);
    expect(await dkg.getRoundBroadcastCount(generation, 2)).to.equal(1);
    expect(await dkg.getRoundBroadcastCount(generation, 3)).to.equal(0);
    expect(await dkg.isRoundFilled(generation, 2)).to.equal(false);

    // round2 - v2

    await expect(dkgV2.roundBroadcast(generation, 3, data2)).to.be.revertedWith('DKG: round was not filled');
    await expect(dkgV2.roundBroadcast(generation, 2, data2)).to.emit(dkgV2, 'RoundDataFilled').withArgs(generation, 2);

    expect(await dkg.getRoundBroadcastData(generation, 2, v2.address)).to.equal(data2);
    expect(await dkg.getRoundBroadcastCount(generation, 1)).to.equal(2);
    expect(await dkg.getRoundBroadcastCount(generation, 2)).to.equal(2);
    expect(await dkg.getRoundBroadcastCount(generation, 3)).to.equal(0);
    expect(await dkg.isRoundFilled(generation, 2)).to.equal(true);

    // round3 - v1

    await expect(dkgV1.voteSigner(generation, signerAddress)).to.be.revertedWith('DKG: round was not filled');
    await expect(dkgV1.roundBroadcast(generation, 3, data3))
      .to.emit(dkgV1, 'RoundDataProvided')
      .withArgs(generation, 3, v1.address);

    expect(await dkgV1.getRoundBroadcastData(generation, 3, v1.address)).to.equal(data3);
    expect(await dkgV1.getRoundBroadcastCount(generation, 1)).to.equal(2);
    expect(await dkgV1.getRoundBroadcastCount(generation, 2)).to.equal(2);
    expect(await dkgV1.getRoundBroadcastCount(generation, 3)).to.equal(1);
    expect(await dkg.isRoundFilled(generation, 3)).to.equal(false);

    // round3 - v2

    await expect(dkgV2.voteSigner(generation, signerAddress)).to.be.revertedWith('DKG: round was not filled');
    await expect(dkgV2.roundBroadcast(generation, 3, data3)).to.emit(dkgV2, 'RoundDataFilled').withArgs(generation, 3);

    expect(await dkgV2.getRoundBroadcastData(generation, 3, v1.address)).to.equal(data3);
    expect(await dkgV2.getRoundBroadcastCount(generation, 1)).to.equal(2);
    expect(await dkgV2.getRoundBroadcastCount(generation, 2)).to.equal(2);
    expect(await dkgV2.getRoundBroadcastCount(generation, 3)).to.equal(2);
    expect(await dkg.isRoundFilled(generation, 3)).to.equal(true);

    await expect(dkgV1.voteSigner(generation, signerAddress))
      .to.emit(dkgV1, 'SignerVoted')
      .withArgs(generation, v1.address, signerAddress);

    await expect(dkgV1.voteSigner(generation, signerAddress)).to.be.revertedWith('DKG: already voted');

    await expect(dkgV2.voteSigner(generation, signerAddress))
      .to.emit(dkgV2, 'SignerAddressUpdated')
      .withArgs(generation, signerAddress);
  });

  it('should set validators by validatorStaking', async function () {
    const [, other] = await ethers.getSigners();
    const initialMinimalStake = ethers.utils.parseEther('3');

    const { dkg } = await deploySystem(initialMinimalStake);

    const dkgOther = await ethers.getContractAt('DKG', dkg.address, other);
    await expect(dkgOther.setValidators([other.address])).to.be.revertedWith('DKG: not a validatorStaking');
  });
});
