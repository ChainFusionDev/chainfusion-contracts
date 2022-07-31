import { expect } from 'chai';
import { deployBridge } from '../utils/deploy';

describe('ValidatorStorage', function () {
  it('should set, get and emit event ValidatorUpdated', async function () {
    const newValidator = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const { validatorStorage } = await deployBridge();

    await expect(validatorStorage.setAddress(newValidator))
      .to.emit(validatorStorage, 'ValidatorUpdated')
      .withArgs(newValidator);
    expect(await validatorStorage.getAddress()).to.equal(newValidator);
  });
});
