import { expect } from 'chai';
import { ethers } from 'hardhat';
import { deployBridge } from '../utils/deploy';

describe('TokenManager', function () {
  it('should called addSupportedToken() only by owner', async function () {
    const [, v1] = await ethers.getSigners();

    const chainId = 123;
    const tokenManagerAddress = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
    const destinationToken = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const TokenManager = await ethers.getContractFactory('TokenManager');
    const tokenManager = await TokenManager.deploy();
    await tokenManager.deployed();
    await tokenManager.initialize(v1.address);

    await expect(tokenManager.setDestinationToken(chainId, tokenManagerAddress, destinationToken)).to.be.revertedWith(
      'Ownable: caller is not the owner'
    );
  });

  it('should adds token addresses to supportedTokens mapping', async function () {
    const [owner] = await ethers.getSigners();
    const initialRequiredApprovals = 1;

    const tokenManagerAddress = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
    const destinationToken = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const { tokenManager, chainId } = await deployBridge(owner.address, [owner.address], initialRequiredApprovals);

    await tokenManager.setDestinationToken(chainId, tokenManagerAddress, destinationToken);
    expect(await tokenManager.getDestinationToken(tokenManagerAddress, chainId)).to.equal(destinationToken);
  });

  it('should check if token supported', async function () {
    const [owner] = await ethers.getSigners();
    const initialRequiredApprovals = 1;

    const tokenManagerAddress = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
    const destinationToken = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const { tokenManager, chainId } = await deployBridge(owner.address, [owner.address], initialRequiredApprovals);

    await tokenManager.setDestinationToken(chainId, tokenManagerAddress, destinationToken);

    await tokenManager.setEnabled(tokenManagerAddress, true);
    await tokenManager.setMintable(tokenManagerAddress, true);

    expect(await tokenManager.isTokenEnabled(tokenManagerAddress)).to.equal(true);
  });
});
