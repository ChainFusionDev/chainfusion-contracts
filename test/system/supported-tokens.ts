import { ethers } from 'hardhat';
import { expect } from 'chai';
import { deploySystem } from '../utils/deploy';

describe('SupportedTokens', function () {
  it('should only owner should add/remove tokens', async function () {
    const [, user] = await ethers.getSigners();
    const symbol = 'ABC';
    const chainId = 1;
    const token = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
    const tokenType = 0;

    const { supportedTokens } = await deploySystem();

    const systemUser = await ethers.getContractAt('SupportedTokens', supportedTokens.address, user);
    await expect(systemUser.addToken(symbol, chainId, token, tokenType)).to.be.revertedWith(
      'SignerOwnable: only signer'
    );
    await expect(systemUser.removeToken(symbol, chainId)).to.be.revertedWith('SignerOwnable: only signer');
  });

  it('should emit the AddedToken event with proper arguments', async function () {
    const symbol = 'ABC';
    const chainId = 1;
    const token = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
    const tokenType = 0;

    const { supportedTokens } = await deploySystem();

    await expect(supportedTokens.addToken(symbol, chainId, token, tokenType))
      .to.emit(supportedTokens, 'AddedToken')
      .withArgs('ABC', 1, '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF', 0);
  });

  it('should emit the RemovedToken event with proper arguments', async function () {
    const symbol = 'ABC';
    const chainId = 1;
    const token = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
    const tokenType = 0;

    const { supportedTokens } = await deploySystem();

    await supportedTokens.addToken(symbol, chainId, token, tokenType);

    await expect(supportedTokens.removeToken(symbol, chainId))
      .to.emit(supportedTokens, 'RemovedToken')
      .withArgs('ABC', 1, '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF');
  });

  it('should add token to token mapping', async function () {
    const symbol = 'ABC';
    const chainId = 1;
    const token = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
    const tokenType = 0;

    const { supportedTokens } = await deploySystem();

    await supportedTokens.addToken(symbol, chainId, token, tokenType);
    expect((await supportedTokens.tokens('ABC', 1)).token).to.equals(token);

    await expect(supportedTokens.addToken(symbol, chainId, token, tokenType)).to.be.revertedWith(
      'SupportedTokens: token already added'
    );
  });

  it('should remove token from token mapping', async function () {
    const symbol = 'ABC';
    const chainId = 1;
    const token = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
    const tokenZero = '0x0000000000000000000000000000000000000000';
    const tokenType = 0;

    const { supportedTokens } = await deploySystem();

    await supportedTokens.addToken(symbol, chainId, token, tokenType);

    await supportedTokens.removeToken(symbol, chainId);
    expect((await supportedTokens.tokens('ABC', 1)).token).to.equals(tokenZero);

    await expect(supportedTokens.removeToken(symbol, chainId)).to.be.revertedWith(
      'SupportedTokens: token does not exist'
    );
  });

  it('should add address and type of token in mapping', async function () {
    const symbol = 'ABC';
    const chainId = 1;
    const token = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
    const tokenType = 0;

    const { supportedTokens } = await deploySystem();

    await supportedTokens.addToken(symbol, chainId, token, tokenType);

    expect((await supportedTokens.tokens('ABC', 1)).token).to.equals(token);
    expect((await supportedTokens.tokens('ABC', 1)).tokenType).to.equals(tokenType);
  });
});
