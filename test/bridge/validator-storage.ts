import { expect } from 'chai';
import { ethers } from 'hardhat';
import { deployBridge } from '../utils/deploy';

describe('ValidatorStorage', function () {
  it('should set, get and emit event ValidatorUpdated', async function () {
    const [validator] = await ethers.getSigners();
    const chainId = 1;
    const newValidator = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const { validatorStorage } = await deployBridge(validator.address, chainId);
    const newValidatorStorage = await ethers.getContractAt('ValidatorStorage', validatorStorage.address, validator);

    await expect(newValidatorStorage.setAddress(newValidator))
      .to.emit(newValidatorStorage, 'ValidatorUpdated')
      .withArgs(newValidator);
    expect(await newValidatorStorage.getAddress()).to.equal(newValidator);
  });
});
