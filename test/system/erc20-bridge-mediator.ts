import { expect } from 'chai';
import { utils } from 'ethers';
import { deploySystem, deployBridgeWithMocks } from '../utils/deploy';
import { ethers } from 'hardhat';

describe('ERC20BridgeMediator', function () {
  it('should add token', async function () {
    const symbol1 = 'CFN';
    const symbol2 = 'NFN';
    const chainId = 111;
    const token1 = '0x0000000000000000000000000000000000000001';
    const token2 = '0x0000000000000000000000000000000000000002';

    const { erc20BridgeMediator } = await deploySystem();

    await expect(await erc20BridgeMediator.addToken(symbol1, chainId, token1))
      .to.emit(erc20BridgeMediator, 'AddedToken')
      .withArgs(symbol1, chainId, token1);

    expect(await erc20BridgeMediator.tokenToSymbol(chainId, token1)).to.equal(symbol1);
    expect(await erc20BridgeMediator.symbolToToken(chainId, symbol1)).to.equal(token1);

    await expect(erc20BridgeMediator.addToken(symbol2, chainId, token1)).to.be.revertedWith(
      'ERC20BridgeMediator: symbol already added'
    );

    await expect(erc20BridgeMediator.addToken(symbol1, chainId, token2)).to.be.revertedWith(
      'ERC20BridgeMediator: token already added'
    );
  });

  it('should remove token', async function () {
    const symbol = 'CFN';
    const chainId = 111;
    const token = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
    const tokenZero = '0x0000000000000000000000000000000000000000';

    const { erc20BridgeMediator } = await deploySystem();

    await erc20BridgeMediator.addToken(symbol, chainId, token);

    await expect(await erc20BridgeMediator.removeToken(symbol, chainId))
      .to.emit(erc20BridgeMediator, 'RemovedToken')
      .withArgs(symbol, chainId, token);

    expect(await erc20BridgeMediator.tokenToSymbol(chainId, token)).to.equal('');
    expect(await erc20BridgeMediator.symbolToToken(chainId, symbol)).to.equal(tokenZero);

    await expect(erc20BridgeMediator.removeToken(symbol, chainId)).to.be.revertedWith(
      'ERC20BridgeMediator: token does not exist'
    );
  });

  it('should only owner should add/remove tokens', async function () {
    const [, user] = await ethers.getSigners();
    const symbol = 'CFN';
    const chainId = 111;
    const token = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const { erc20BridgeMediator } = await deploySystem();

    const systemUser = await ethers.getContractAt('ERC20BridgeMediator', erc20BridgeMediator.address, user);
    await expect(systemUser.addToken(symbol, chainId, token)).to.be.revertedWith('Ownable: caller is not the owner');
    await expect(systemUser.removeToken(symbol, chainId)).to.be.revertedWith('Ownable: caller is not the owner');
  });

  it('should mediate encoded data', async function () {
    const [sender, receiver] = await ethers.getSigners();
    const transferAmount = utils.parseEther('0.99');
    const symbol = 'CFN';
    const sourceChain = 111;
    const destinationChain = 222;
    const destinationToken = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const { erc20BridgeMediator } = await deploySystem();
    const { mockChainId, mockToken } = await deployBridgeWithMocks();

    const abiCoder = ethers.utils.defaultAbiCoder;
    const data = abiCoder.encode(
      ['address', 'address', 'uint256', 'address', 'uint256'],
      [sender.address, mockToken.address, mockChainId, receiver.address, transferAmount]
    );

    await expect(erc20BridgeMediator.mediate(sourceChain, destinationChain, data)).to.be.revertedWith(
      "ERC20BridgeMediator: can't find token symbol"
    );

    await erc20BridgeMediator.addToken(symbol, sourceChain, mockToken.address);

    await expect(erc20BridgeMediator.mediate(sourceChain, destinationChain, data)).to.be.revertedWith(
      "ERC20BridgeMediator: can't find token by chain and symbol"
    );

    await erc20BridgeMediator.addToken(symbol, destinationChain, destinationToken);

    await expect(await erc20BridgeMediator.mediate(sourceChain, destinationChain, data))
      .to.emit(erc20BridgeMediator, 'MediatorAddress')
      .withArgs(erc20BridgeMediator.address);
  });
});
