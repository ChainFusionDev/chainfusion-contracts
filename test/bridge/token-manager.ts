import { expect } from 'chai';
import { utils } from 'ethers';
import { ethers } from 'hardhat';
import { deployBridgeWithMocks } from '../utils/deploy';

describe('TokenManager', function () {
  it('should adds token addresses to supportedTokens mapping', async function () {
    const tokenManagerAddress = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const { tokenManager } = await deployBridgeWithMocks();

    await tokenManager.setToken(tokenManagerAddress, 1);
    expect(await tokenManager.getType(tokenManagerAddress)).to.equal(1);
  });

  it('should check if token supported', async function () {
    const tokenManagerAddress = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const { tokenManager } = await deployBridgeWithMocks();

    await tokenManager.setToken(tokenManagerAddress, 1);
    expect(await tokenManager.getType(tokenManagerAddress)).to.equal(1);
  });

  it('should mint, deploy and burn token', async function () {
    const [sender, user] = await ethers.getSigners();
    const depositAmount = utils.parseEther('1');
    const initialSupply = utils.parseEther('1');
    const tokenName = 'CFN Token';
    const tokenSymbol = 'CFN';
    const tokenAddress = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
    const tokenType = 0;

    const { tokenManager, erc20Bridgeable } = await deployBridgeWithMocks();
    await erc20Bridgeable.mint(sender.address, initialSupply);
    await erc20Bridgeable.approve(tokenManager.address, depositAmount);

    const tokenManagerUser = await ethers.getContractAt('TokenManager', tokenManager.address, user);
    await expect(tokenManagerUser.deployToken(tokenName, tokenSymbol)).to.be.revertedWith('SignerOwnable: only signer');
    await expect(tokenManagerUser.setToken(tokenAddress, tokenType)).to.be.revertedWith('SignerOwnable: only signer');

    const erc20BridgeableAddress = await tokenManager.callStatic.deployToken(tokenName, tokenSymbol);
    await tokenManager.deployToken(tokenName, tokenSymbol);

    expect(await tokenManager.getType(erc20BridgeableAddress)).to.equal(2);

    await erc20Bridgeable.approve(sender.address, initialSupply);
    await erc20Bridgeable.burnFrom(sender.address, initialSupply);
  });
});
