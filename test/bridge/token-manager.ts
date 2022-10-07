import { expect } from 'chai';
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
});
