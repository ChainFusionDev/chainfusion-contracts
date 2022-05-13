import { expect } from 'chai';
import { ethers } from 'hardhat';
import { deployBridge } from '../utils/deploy';

describe('FeeManager', function () {
  it('should check if fee is transferred to FeeManager contract', async function () {
    const [owner, receiver] = await ethers.getSigners();
    const initialRequiredApprovals = 1;
    const amount = '10000000000000000000';
    const fee = '10000000000000000';

    const { liquidityPools, tokenManager, mockToken, bridge, chainId, feeManager } = await deployBridge(
      owner.address,
      [owner.address],
      initialRequiredApprovals
    );

    expect(await tokenManager.isTokenEnabled(mockToken.address)).to.equal(true);

    await mockToken.approve(bridge.address, amount);
    await mockToken.approve(liquidityPools.address, amount);
    await liquidityPools.addLiquidity(mockToken.address, amount);

    await bridge.deposit(mockToken.address, chainId, receiver.address, amount);
    expect(await mockToken.balanceOf(feeManager.address)).to.equal(fee);
  });
});
