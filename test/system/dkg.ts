import { ethers } from 'hardhat';
import { expect } from 'chai';
import { deploySystem } from '../utils/deploy';

describe('DKG', function () {
  it('should broadcast all rounds', async function () {
    const generation = 1;

    const data1 = ethers.utils.keccak256([1]);
    const data2 = ethers.utils.keccak256([2]);
    const data3 = ethers.utils.keccak256([3]);

    const [, v1, v2, signer, other] = await ethers.getSigners();

    const signerAddress = signer.address;
    const message = 'verify';

    const signature = await signer.signMessage(message);
    const signatureOther = await other.signMessage(message);

    const initialMinimalStake = ethers.utils.parseEther('3');
    const { dkg, staking } = await deploySystem();

    expect(await dkg.getGenerationsCount()).to.equal(1);

    const staking1 = await ethers.getContractAt('Staking', staking.address, v1);
    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);

    await expect(dkg.roundBroadcast(generation, 1, data1)).to.be.revertedWith('DKG: not a validator');

    expect(await dkg.isValidator(generation, v1.address)).to.equal(false);
    expect(await dkg.isValidator(generation, v2.address)).to.equal(false);
    expect(await dkg.getValidators(generation)).to.deep.equal([]);
    expect(await dkg.getValidatorsCount(generation)).to.equal(0);

    await staking1.stake({ value: initialMinimalStake });
    await staking2.stake({ value: initialMinimalStake });

    expect(await dkg.getGenerationsCount()).to.equal(2);

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

    await expect(dkgV1.voteSigner(generation, signerAddress, signature)).to.be.revertedWith(
      'DKG: round was not filled'
    );
    await expect(dkgV1.roundBroadcast(generation, 3, data3))
      .to.emit(dkgV1, 'RoundDataProvided')
      .withArgs(generation, 3, v1.address);

    expect(await dkgV1.getRoundBroadcastData(generation, 3, v1.address)).to.equal(data3);
    expect(await dkgV1.getRoundBroadcastCount(generation, 1)).to.equal(2);
    expect(await dkgV1.getRoundBroadcastCount(generation, 2)).to.equal(2);
    expect(await dkgV1.getRoundBroadcastCount(generation, 3)).to.equal(1);
    expect(await dkg.isRoundFilled(generation, 3)).to.equal(false);

    // round3 - v2

    await expect(dkgV2.voteSigner(generation, signerAddress, signature)).to.be.revertedWith(
      'DKG: round was not filled'
    );
    await expect(dkgV2.roundBroadcast(generation, 3, data3)).to.emit(dkgV2, 'RoundDataFilled').withArgs(generation, 3);

    expect(await dkgV2.getRoundBroadcastData(generation, 3, v1.address)).to.equal(data3);
    expect(await dkgV2.getRoundBroadcastCount(generation, 1)).to.equal(2);
    expect(await dkgV2.getRoundBroadcastCount(generation, 2)).to.equal(2);
    expect(await dkgV2.getRoundBroadcastCount(generation, 3)).to.equal(2);
    expect(await dkg.isRoundFilled(generation, 3)).to.equal(true);

    await expect(dkgV1.voteSigner(generation, signerAddress, signature))
      .to.emit(dkgV1, 'SignerVoted')
      .withArgs(generation, v1.address, signerAddress);

    await expect(dkgV1.voteSigner(generation, signerAddress, signatureOther)).to.be.revertedWith(
      'DKG: signature is invalid'
    );
    await expect(dkgV2.voteSigner(generation, signerAddress, signatureOther)).to.be.revertedWith(
      'DKG: signature is invalid'
    );

    await expect(dkgV1.voteSigner(generation, signerAddress, signature)).to.be.revertedWith('DKG: already voted');
    await expect(dkgV2.voteSigner(generation, signerAddress, signature))
      .to.emit(dkgV2, 'SignerAddressUpdated')
      .withArgs(generation, signerAddress);
  });

  it('should set validators by staking', async function () {
    const [, other] = await ethers.getSigners();

    const { dkg } = await deploySystem();

    const dkgOther = await ethers.getContractAt('DKG', dkg.address, other);
    await expect(dkgOther.setValidators([other.address])).to.be.revertedWith('DKG: not a staking');
  });

  it('should get active and pending status ', async function () {
    const initialMinimalStake = ethers.utils.parseEther('3');

    const [, signer] = await ethers.getSigners();
    const { dkg, staking } = await deploySystem(initialMinimalStake);

    const PENDING: number = 0;
    const ACTIVE: number = 2;

    const generation = 1;
    const message = 'verify';
    const signature = await signer.signMessage(message);
    const data1 = ethers.utils.keccak256([1]);
    const data2 = ethers.utils.keccak256([2]);
    const data3 = ethers.utils.keccak256([3]);

    const dkgSigner = await ethers.getContractAt('DKG', dkg.address, signer);
    const stakingSigner = await ethers.getContractAt('Staking', staking.address, signer);

    await staking.stake({ value: initialMinimalStake });
    await stakingSigner.stake({ value: initialMinimalStake });

    expect(await dkgSigner.getStatus(generation)).to.equal(PENDING);

    await dkg.roundBroadcast(generation, 1, data1);
    await dkgSigner.roundBroadcast(generation, 1, data1);

    await dkg.roundBroadcast(generation, 2, data2);
    await dkgSigner.roundBroadcast(generation, 2, data2);

    await dkg.roundBroadcast(generation, 3, data3);
    await dkgSigner.roundBroadcast(generation, 3, data3);

    expect(await dkgSigner.getStatus(generation)).to.equal(0);

    await dkg.voteSigner(generation, signer.address, signature);
    await dkgSigner.voteSigner(generation, signer.address, signature);

    expect(await dkgSigner.getStatus(generation)).to.equal(ACTIVE);
  });

  it('should get expired status', async function () {
    const initialMinimalStake = ethers.utils.parseEther('3');

    const [, signer] = await ethers.getSigners();
    const { dkg, staking } = await deploySystem(initialMinimalStake);

    const EXPIRED: number = 1;
    const generation = 1;

    const stakingSigner = await ethers.getContractAt('Staking', staking.address, signer);

    await staking.stake({ value: initialMinimalStake });
    await stakingSigner.stake({ value: initialMinimalStake });

    const hre = require('hardhat');
    await hre.network.provider.send('hardhat_mine', ['0x64']);

    expect(await dkg.getStatus(generation)).to.equal(EXPIRED);
  });

  it('should check if we can set deadline period ', async function () {
    const initialMinimalStake = ethers.utils.parseEther('3');

    const [, signer] = await ethers.getSigners();
    const { dkg, staking } = await deploySystem(initialMinimalStake);

    const generation = 1;
    const message = 'verify';
    const signature = await signer.signMessage(message);
    const data1 = ethers.utils.keccak256([1]);
    const data2 = ethers.utils.keccak256([2]);
    const data3 = ethers.utils.keccak256([3]);

    const dkgSigner = await ethers.getContractAt('DKG', dkg.address, signer);
    const stakingSigner = await ethers.getContractAt('Staking', staking.address, signer);

    await staking.stake({ value: initialMinimalStake });
    await stakingSigner.stake({ value: initialMinimalStake });

    await dkg.roundBroadcast(generation, 1, data1);
    await dkgSigner.roundBroadcast(generation, 1, data1);

    await dkg.roundBroadcast(generation, 2, data2);
    await dkgSigner.roundBroadcast(generation, 2, data2);

    await dkg.roundBroadcast(generation, 3, data3);
    await dkgSigner.roundBroadcast(generation, 3, data3);

    await expect(dkgSigner.setDeadlinePeriod(100)).to.be.revertedWith('DKG: not a active signer');

    await dkg.voteSigner(generation, signer.address, signature);
    await dkgSigner.voteSigner(generation, signer.address, signature);

    await dkgSigner.setDeadlinePeriod(100);

    expect(await dkgSigner.deadlinePeriod()).to.equals(100);
  });
});
