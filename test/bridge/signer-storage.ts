import { expect } from 'chai';
import { utils } from 'ethers';
import { ethers } from 'hardhat';
import { deployBridge } from '../utils/deploy';

describe('SignerStorage', function () {
  it('should set, get and emit event', async function () {
    const [, newSigner] = await ethers.getSigners();

    const { signerStorage } = await deployBridge();

    const balanceNewSignerBefore = await ethers.provider.getBalance(newSigner.address);
    const value = utils.parseEther('1');

    await expect(signerStorage.setAddress(newSigner.address, { value }))
      .to.emit(signerStorage, 'SignerUpdated')
      .withArgs(newSigner.address);
    expect(await signerStorage.getAddress()).to.equal(newSigner.address);

    const balanceNewSignerAfter = await ethers.provider.getBalance(newSigner.address);
    expect(balanceNewSignerAfter).to.be.equal(balanceNewSignerBefore.add(value));
  });
});
