import { ethers } from 'hardhat';
import { expect } from 'chai';

describe('DKG', function () {
  it('should broadcast all rounds', async function () {
    const generation = 0;
    const dummyAddress = '0x0000000000000000000000000000000000000001';
    const data1 = ethers.utils.keccak256([1]);
    const data2 = ethers.utils.keccak256([2]);
    const data3 = ethers.utils.keccak256([3]);

    const [, v1, v2] = await ethers.getSigners();
    const DKG = await ethers.getContractFactory('DKG');
    const Signing = await ethers.getContractFactory('Signing');
    const dkg = await DKG.deploy();
    const signing = await Signing.deploy();

    await expect(dkg.initialize([v1.address, v2.address]))
      .to.emit(dkg, 'ValidatorsUpdated')
      .withArgs(generation, [v1.address, v2.address]);

    await signing.setDKG(dkg.address);
    await dkg.setSigning(signing.address);

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

    await expect(dkgV1.voteCollectiveAddress(generation, dummyAddress)).to.be.revertedWith(
      'DKG: previous round was not filled'
    );
    await expect(dkgV1.roundBroadcast(generation, 3, data3))
      .to.emit(dkgV1, 'RoundDataProvided')
      .withArgs(generation, 3, v1.address);

    expect(await dkgV1.getRoundBroadcastData(generation, 3, v1.address)).to.equal(data3);
    expect(await dkgV1.getRoundBroadcastCount(generation, 1)).to.equal(2);
    expect(await dkgV1.getRoundBroadcastCount(generation, 2)).to.equal(2);
    expect(await dkgV1.getRoundBroadcastCount(generation, 3)).to.equal(1);

    // round3 - v2

    await expect(dkgV2.voteCollectiveAddress(generation, dummyAddress)).to.be.revertedWith(
      'DKG: previous round was not filled'
    );
    await expect(dkgV2.roundBroadcast(generation, 3, data3)).to.emit(dkgV2, 'RoundDataFilled').withArgs(generation, 3);

    expect(await dkgV2.getRoundBroadcastData(generation, 3, v1.address)).to.equal(data3);
    expect(await dkgV2.getRoundBroadcastCount(generation, 1)).to.equal(2);
    expect(await dkgV2.getRoundBroadcastCount(generation, 2)).to.equal(2);
    expect(await dkgV2.getRoundBroadcastCount(generation, 3)).to.equal(2);

    await dkgV1.voteCollectiveAddress(generation, dummyAddress);
    await expect(dkgV2.voteCollectiveAddress(generation, dummyAddress))
      .to.emit(signing, 'CollectiveSignerUpdated')
      .withArgs(generation, dummyAddress);
  });

  it('should set validators by owner', async function () {
    const [, v1, v2, other] = await ethers.getSigners();
    const DKG = await ethers.getContractFactory('DKG');
    const dkg = await DKG.deploy();
    await dkg.initialize([v1.address]);

    await dkg.setValidators([v1.address, v2.address]);

    const dkgOther = await ethers.getContractAt('DKG', dkg.address, other);
    await expect(dkgOther.setValidators([other.address])).to.be.revertedWith('Ownable: caller is not the owner');
  });
});
