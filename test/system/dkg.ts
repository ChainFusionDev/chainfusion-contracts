import { ethers } from 'hardhat';
import { expect } from 'chai';

describe('DKG', function () {
  it('should broadcast all rounds', async function () {
    const id = 0;
    const data1 = ethers.utils.keccak256([1]);
    const data2 = ethers.utils.keccak256([2]);
    const data3 = ethers.utils.keccak256([3]);

    const [, v1] = await ethers.getSigners();
    const DKG = await ethers.getContractFactory('DKG');
    const dkg = await DKG.deploy();
    expect(await dkg.initialize([v1.address]))
      .to.emit(dkg, 'ValidatorsUpdated')
      .withArgs([v1.address]);

    await expect(dkg.roundBroadcast(id, 1, data1)).to.be.revertedWith('not a validator');

    const dkgV1 = await ethers.getContractAt('DKG', dkg.address, v1);
    expect(await dkgV1.roundBroadcast(id, 1, data1))
      .to.emit(DKG, 'RoundDataProvided')
      .withArgs(id, 1, v1.address)
      .to.emit(DKG, 'RoundDataFilled')
      .withArgs(id, 1);

    expect(await dkgV1.getRoundBroadcastData(id, 1, v1.address)).to.equal(data1);
    expect(await dkgV1.getRoundBroadcastCount(id, 1)).to.equal(1);
    expect(await dkgV1.getRoundBroadcastCount(id, 2)).to.equal(0);
    expect(await dkgV1.getRoundBroadcastCount(id, 3)).to.equal(0);

    expect(await dkgV1.roundBroadcast(id, 2, data2))
      .to.emit(DKG, 'RoundDataProvided')
      .withArgs(id, 2, v1.address)
      .to.emit(DKG, 'RoundDataFilled')
      .withArgs(id, 2);

    expect(await dkgV1.getRoundBroadcastData(id, 2, v1.address)).to.equal(data2);
    expect(await dkgV1.getRoundBroadcastCount(id, 1)).to.equal(1);
    expect(await dkgV1.getRoundBroadcastCount(id, 2)).to.equal(1);
    expect(await dkgV1.getRoundBroadcastCount(id, 3)).to.equal(0);

    expect(await dkgV1.roundBroadcast(id, 3, data3))
      .to.emit(DKG, 'RoundDataProvided')
      .withArgs(id, 3, v1.address)
      .to.emit(DKG, 'RoundDataFilled')
      .withArgs(id, 3);

    expect(await dkgV1.getRoundBroadcastData(id, 3, v1.address)).to.equal(data3);
    expect(await dkgV1.getRoundBroadcastCount(id, 1)).to.equal(1);
    expect(await dkgV1.getRoundBroadcastCount(id, 2)).to.equal(1);
    expect(await dkgV1.getRoundBroadcastCount(id, 3)).to.equal(1);
  });

  it('should set validators by owner', async function () {
    const [, v1, v2, other] = await ethers.getSigners();
    const DKG = await ethers.getContractFactory('DKG');
    const dkg = await DKG.deploy();
    await dkg.initialize([v1.address]);

    await dkg.setValidators([v1.address, v2.address]);

    const dkgOther = await ethers.getContractAt('DKG', dkg.address, other);
    expect(dkgOther.setValidators([other.address])).to.be.revertedWith('Ownable: caller is not the owner');
  });
});
