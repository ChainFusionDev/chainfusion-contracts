import { ethers } from 'hardhat';
import { expect } from 'chai';
import { deploySystem } from '../utils/deploy';

describe('SupportedTokens', function () {
  it('should only owner should add/remove tokens', async function () {
    const [, user] = await ethers.getSigners();
    const symbol = 'ABC';
    const chainId = 1;
    const token = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const { supportedTokens } = await deploySystem();

    const systemUser = await ethers.getContractAt('SupportedTokens', supportedTokens.address, user);
    await expect(systemUser.addToken(symbol, chainId, token)).to.be.revertedWith('Ownable: caller is not the owner');
    await expect(systemUser.removeToken(symbol, chainId)).to.be.revertedWith('Ownable: caller is not the owner');
  });

  it('should emit the AddedToken event with proper arguments', async function () {
    const symbol = 'ABC';
    const chainId = 1;
    const token = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const { supportedTokens } = await deploySystem();

    await expect(await supportedTokens.addToken(symbol, chainId, token))
      .to.emit(supportedTokens, 'AddedToken')
      .withArgs('ABC', 1, '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF');
  });

  it('should emit the RemovedToken event with proper arguments', async function () {
    const symbol = 'ABC';
    const chainId = 1;
    const token = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const { supportedTokens } = await deploySystem();

    expect(await supportedTokens.addToken(symbol, chainId, token));

    await expect(await supportedTokens.removeToken(symbol, chainId))
      .to.emit(supportedTokens, 'RemovedToken')
      .withArgs('ABC', 1, '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF');
  });

  it('should add token to token mapping', async function () {
    const symbol = 'ABC';
    const chainId = 1;
    const token = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const { supportedTokens } = await deploySystem();

    expect(await supportedTokens.addToken(symbol, chainId, token));
    expect(await supportedTokens.tokens('ABC', 1)).to.equals(token);

    await expect(supportedTokens.addToken(symbol, chainId, token)).to.be.revertedWith(
      'SupportedTokens: token already added'
    );
  });

  it('should remove token from token mapping', async function () {
    const symbol = 'ABC';
    const chainId = 1;
    const token = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
    const tokenZero = '0x0000000000000000000000000000000000000000';

    const { supportedTokens } = await deploySystem();

    expect(await supportedTokens.addToken(symbol, chainId, token));

    expect(await supportedTokens.removeToken(symbol, chainId));
    expect(await supportedTokens.tokens('ABC', 1)).to.equals(tokenZero);

    await expect(supportedTokens.removeToken(symbol, chainId)).to.be.revertedWith(
      'SupportedTokens: token does not exist'
    );
  });
});
