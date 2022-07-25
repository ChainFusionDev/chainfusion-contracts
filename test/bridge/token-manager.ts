import { expect } from 'chai';
import { ethers } from 'hardhat';
import { deployBridge } from '../utils/deploy';

describe('TokenManager', function () {
  it('should adds token addresses to supportedTokens mapping', async function () {
    const [validator] = await ethers.getSigners();
    const initialRequiredApprovals = 1;

    const tokenManagerAddress = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
    const destinationToken = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const { tokenManager, chainId } = await deployBridge(validator.address, initialRequiredApprovals);

    await tokenManager.setDestinationToken(chainId, tokenManagerAddress, destinationToken);
    expect(await tokenManager.getDestinationToken(tokenManagerAddress, chainId)).to.equal(destinationToken);
  });

  it('should check if token supported', async function () {
    const [validator] = await ethers.getSigners();
    const initialRequiredApprovals = 1;

    const tokenManagerAddress = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
    const destinationToken = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const { tokenManager, chainId } = await deployBridge(validator.address, initialRequiredApprovals);

    await tokenManager.setDestinationToken(chainId, tokenManagerAddress, destinationToken);

    await tokenManager.setEnabled(tokenManagerAddress, true);

    expect(await tokenManager.isTokenEnabled(tokenManagerAddress)).to.equal(true);
  });
});
