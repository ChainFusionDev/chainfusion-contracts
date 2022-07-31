import { expect } from 'chai';
import { deployBridgeWithMocks } from '../utils/deploy';

describe('TokenManager', function () {
  it('should adds token addresses to supportedTokens mapping', async function () {
    const tokenManagerAddress = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
    const destinationToken = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const { mockChainId, tokenManager } = await deployBridgeWithMocks();

    await tokenManager.setDestinationToken(mockChainId, tokenManagerAddress, destinationToken);
    expect(await tokenManager.getDestinationToken(tokenManagerAddress, mockChainId)).to.equal(destinationToken);
  });

  it('should check if token supported', async function () {
    const tokenManagerAddress = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
    const destinationToken = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const { mockChainId, tokenManager } = await deployBridgeWithMocks();

    await tokenManager.setDestinationToken(mockChainId, tokenManagerAddress, destinationToken);
    await tokenManager.setEnabled(tokenManagerAddress, true);

    expect(await tokenManager.isTokenEnabled(tokenManagerAddress)).to.equal(true);
  });
});
