import { ethers } from 'hardhat';
import { expect } from 'chai';

describe('DKG', function () {
  it('should broadcast all rounds', async function () {
    const id = ethers.utils.keccak256([0]);
    const data1 = ethers.utils.keccak256([1]);
    const data2 = ethers.utils.keccak256([2]);
    const data3 = ethers.utils.keccak256([3]);

    const [, v1] = await ethers.getSigners();
    const DKG = await ethers.getContractFactory('DKG');
    const dkg = await DKG.deploy();
    expect(await dkg.initialize([v1.address]))
      .to.emit(dkg, 'ValidatorsUpdated')
      .withArgs([v1.address]);

    await expect(dkg.round1Broadcast(id, [data1], [data1])).to.be.revertedWith('not a validator');

    const dkgV1 = await ethers.getContractAt('DKG', dkg.address, v1);
    expect(await dkgV1.round1Broadcast(id, [data1], [data1]))
      .to.emit(DKG, 'Round1Provided')
      .withArgs(id, v1.address)
      .to.emit(DKG, 'Round1Filled')
      .withArgs(id);

    const bcastData1 = await dkgV1.getRound1BroadcastData(id, v1.address);
    expect(bcastData1.provided).to.equal(true);
    expect(bcastData1.publicData.length).to.equal(1);
    expect(bcastData1.privateData.length).to.equal(1);
    expect(bcastData1.publicData[0]).to.equal(data1);
    expect(bcastData1.privateData[0]).to.equal(data1);

    expect(await dkgV1.getRound1BroadcastCount(id)).to.equal(1);
    expect(await dkgV1.getRound2BroadcastCount(id)).to.equal(0);
    expect(await dkgV1.getRound3BroadcastCount(id)).to.equal(0);

    expect(await dkgV1.round2Broadcast(id, [data2]))
      .to.emit(DKG, 'Round2Provided')
      .withArgs(id, v1.address)
      .to.emit(DKG, 'Round2Filled')
      .withArgs(id);

    const bcastData2 = await dkgV1.getRound2BroadcastData(id, v1.address);
    expect(bcastData2.provided).to.equal(true);
    expect(bcastData2.publicData.length).to.equal(1);
    expect(bcastData2.privateData.length).to.equal(0);
    expect(bcastData2.publicData[0]).to.equal(data2);

    expect(await dkgV1.getRound1BroadcastCount(id)).to.equal(1);
    expect(await dkgV1.getRound2BroadcastCount(id)).to.equal(1);
    expect(await dkgV1.getRound3BroadcastCount(id)).to.equal(0);

    expect(await dkgV1.round3Broadcast(id, [data3]))
      .to.emit(DKG, 'Round3Provided')
      .withArgs(id, v1.address)
      .to.emit(DKG, 'Round3Filled')
      .withArgs(id);

    const bcastData3 = await dkgV1.getRound3BroadcastData(id, v1.address);
    expect(bcastData3.provided).to.equal(true);
    expect(bcastData3.publicData.length).to.equal(1);
    expect(bcastData3.privateData.length).to.equal(0);
    expect(bcastData3.publicData[0]).to.equal(data3);

    expect(await dkgV1.getRound1BroadcastCount(id)).to.equal(1);
    expect(await dkgV1.getRound2BroadcastCount(id)).to.equal(1);
    expect(await dkgV1.getRound3BroadcastCount(id)).to.equal(1);
  });

  it('should broadcast all rounds one after another', async function () {
    const data = ethers.utils.keccak256([]);

    const [, v1] = await ethers.getSigners();
    const DKG = await ethers.getContractFactory('DKG');
    const dkg = await DKG.deploy();
    await dkg.initialize([v1.address]);

    const dkgV1 = await ethers.getContractAt('DKG', dkg.address, v1);
    await expect(dkgV1.round2Broadcast(data, [data])).to.be.revertedWith('round 1 not finished');
    await expect(dkgV1.round3Broadcast(data, [data])).to.be.revertedWith('round 1 not finished');

    await dkgV1.round1Broadcast(data, [data], [data]);
    await expect(dkgV1.round3Broadcast(data, [data])).to.be.revertedWith('round 2 not finished');

    await dkgV1.round2Broadcast(data, [data]);
    await dkgV1.round3Broadcast(data, [data]);
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
