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
    const initialminimalStake = ethers.utils.parseEther('3');

    const { addressStorage } = await deploySystem(initialminimalStake);
    const DKG = await ethers.getContractFactory('DKG');
    const ThresholdSigner = await ethers.getContractFactory('ThresholdSigner');
    const dkg = await DKG.deploy();
    const thresholdSigner = await ThresholdSigner.deploy();

    await expect(dkg.initialize([v1.address, v2.address], addressStorage.address))
      .to.emit(dkg, 'ValidatorsUpdated')
      .withArgs(generation, [v1.address, v2.address]);

    await thresholdSigner.setDKG(dkg.address);
    await dkg.setThresholdSigner(thresholdSigner.address);

    await expect(dkg.roundBroadcast(generation, 1, data1)).to.be.revertedWith('DKG: not a validator');

    const dkgV1 = await ethers.getContractAt('DKG', dkg.address, v1);
    const dkgV2 = await ethers.getContractAt('DKG', dkg.address, v2);

    await expect(dkgV1.roundBroadcast(generation, 2, data2)).to.be.revertedWith('DKG: previous round was not filled');
    await expect(dkgV2.roundBroadcast(generation, 2, data2)).to.be.revertedWith('DKG: previous round was not filled');

    // round1 - v1

    await expect(dkgV1.roundBroadcast(generation, 1, data1))
      .to.emit(dkgV1, 'RoundDataProvided')
      .withArgs(generation, 1, v1.address);

    expect(await dkg.getRoundBroadcastData(generation, 1, v1.address)).to.equal(data1);
    expect(await dkg.getRoundBroadcastCount(generation, 1)).to.equal(1);
    expect(await dkg.getRoundBroadcastCount(generation, 2)).to.equal(0);
    expect(await dkg.getRoundBroadcastCount(generation, 3)).to.equal(0);

    // round1 - v2

    await expect(dkgV2.roundBroadcast(generation, 1, data1)).to.emit(dkgV2, 'RoundDataFilled').withArgs(generation, 1);

    expect(await dkg.getRoundBroadcastData(generation, 1, v2.address)).to.equal(data1);
    expect(await dkg.getRoundBroadcastCount(generation, 1)).to.equal(2);
    expect(await dkg.getRoundBroadcastCount(generation, 2)).to.equal(0);
    expect(await dkg.getRoundBroadcastCount(generation, 3)).to.equal(0);

    // round2 - v1

    await expect(dkgV1.roundBroadcast(generation, 3, data2)).to.be.revertedWith('DKG: previous round was not filled');
    await expect(dkgV1.roundBroadcast(generation, 2, data2))
      .to.emit(dkgV1, 'RoundDataProvided')
      .withArgs(generation, 2, v1.address);

    expect(await dkg.getRoundBroadcastData(generation, 2, v1.address)).to.equal(data2);
    expect(await dkg.getRoundBroadcastCount(generation, 1)).to.equal(2);
    expect(await dkg.getRoundBroadcastCount(generation, 2)).to.equal(1);
    expect(await dkg.getRoundBroadcastCount(generation, 3)).to.equal(0);

    // round2 - v2

    await expect(dkgV2.roundBroadcast(generation, 3, data2)).to.be.revertedWith('DKG: previous round was not filled');
    await expect(dkgV2.roundBroadcast(generation, 2, data2)).to.emit(dkgV2, 'RoundDataFilled').withArgs(generation, 2);

    expect(await dkg.getRoundBroadcastData(generation, 2, v2.address)).to.equal(data2);
    expect(await dkg.getRoundBroadcastCount(generation, 1)).to.equal(2);
    expect(await dkg.getRoundBroadcastCount(generation, 2)).to.equal(2);
    expect(await dkg.getRoundBroadcastCount(generation, 3)).to.equal(0);

    // round3 - v1

    await expect(dkgV1.voteSigner(generation, signerAddress)).to.be.revertedWith('DKG: previous round was not filled');
    await expect(dkgV1.roundBroadcast(generation, 3, data3))
      .to.emit(dkgV1, 'RoundDataProvided')
      .withArgs(generation, 3, v1.address);

    expect(await dkgV1.getRoundBroadcastData(generation, 3, v1.address)).to.equal(data3);
    expect(await dkgV1.getRoundBroadcastCount(generation, 1)).to.equal(2);
    expect(await dkgV1.getRoundBroadcastCount(generation, 2)).to.equal(2);
    expect(await dkgV1.getRoundBroadcastCount(generation, 3)).to.equal(1);

    // round3 - v2

    await expect(dkgV2.voteSigner(generation, signerAddress)).to.be.revertedWith('DKG: previous round was not filled');
    await expect(dkgV2.roundBroadcast(generation, 3, data3)).to.emit(dkgV2, 'RoundDataFilled').withArgs(generation, 3);

    expect(await dkgV2.getRoundBroadcastData(generation, 3, v1.address)).to.equal(data3);
    expect(await dkgV2.getRoundBroadcastCount(generation, 1)).to.equal(2);
    expect(await dkgV2.getRoundBroadcastCount(generation, 2)).to.equal(2);
    expect(await dkgV2.getRoundBroadcastCount(generation, 3)).to.equal(2);

    await expect(dkgV1.voteSigner(generation, signerAddress))
      .to.emit(dkgV1, 'SignerVoted')
      .withArgs(generation, v1.address, signerAddress);

    await expect(dkgV2.voteSigner(generation, signerAddress))
      .to.emit(thresholdSigner, 'SignerAddressUpdated')
      .withArgs(generation, signerAddress);
  });

  it('should set validators by validatorStaking', async function () {
    const [, other] = await ethers.getSigners();
    const initialminimalStake = ethers.utils.parseEther('3');

    const { dkg } = await deploySystem(initialminimalStake);

    const dkgOther = await ethers.getContractAt('DKG', dkg.address, other);
    await expect(dkgOther.setValidators([other.address])).to.be.revertedWith('DKG: not a validatorStaking');
  });
});
