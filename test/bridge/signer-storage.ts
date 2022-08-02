import { expect } from 'chai';
import { deployBridge } from '../utils/deploy';

describe('SignerStorage', function () {
  it('should set, get and emit event', async function () {
    const newSigner = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const { signerStorage } = await deployBridge();

    await expect(signerStorage.setAddress(newSigner)).to.emit(signerStorage, 'SignerUpdated').withArgs(newSigner);
    expect(await signerStorage.getAddress()).to.equal(newSigner);
  });
});
