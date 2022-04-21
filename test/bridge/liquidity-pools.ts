import { expect } from 'chai';
import { ethers } from 'hardhat';
import { deployBridge } from '../utils/deploy';

describe('LiquidityPools', function () {
  it('should add liquidity for supported token', async function () {
    const [owner] = await ethers.getSigners();
    const initialRequiredApprovals = 1;

    const amount = '100000';

    const { liquidityPools, tokenManager, mockToken } = await deployBridge(
      owner.address,
      [owner.address],
      initialRequiredApprovals
    );

    expect(await tokenManager.isTokenSupported(mockToken.address)).to.equal(true);

    await mockToken.approve(liquidityPools.address, amount);
    await expect(liquidityPools.addLiquidity(mockToken.address, amount))
      .to.emit(liquidityPools, 'LiquidityAdded')
      .withArgs(mockToken.address, owner.address, amount);

    expect(await liquidityPools.providedLiquidity(mockToken.address)).to.equal(amount);
    expect(await liquidityPools.availableLiquidity(mockToken.address)).to.equal(amount);
    expect((await liquidityPools.liquidityPositions(mockToken.address, owner.address)).balance).to.equal(amount);
  });

  it('should not add liquidity for unsupported token', async function () {
    const [owner] = await ethers.getSigners();
    const initialRequiredApprovals = 1;

    const amount = '100000';
    const mintAmount = '100000000000000000000';

    const { liquidityPools, mockToken } = await deployBridge(owner.address, [owner.address], initialRequiredApprovals);

    const MockToken = await ethers.getContractFactory('MockToken');
    const mockToken2 = await MockToken.deploy('Token2', 'TKN2', mintAmount);
    await mockToken2.deployed();

    await mockToken.approve(liquidityPools.address, amount);
    await mockToken2.approve(liquidityPools.address, amount);

    await expect(liquidityPools.addLiquidity(mockToken2.address, amount)).to.be.revertedWith(
      'TokenManager: token is not supported'
    );
  });

  it('should remove liquidity', async function () {
    const [owner] = await ethers.getSigners();
    const initialRequiredApprovals = 1;

    const amount = '100000';

    const { liquidityPools, tokenManager, mockToken } = await deployBridge(
      owner.address,
      [owner.address],
      initialRequiredApprovals
    );

    expect(await tokenManager.isTokenSupported(mockToken.address)).to.equal(true);

    await mockToken.approve(liquidityPools.address, amount);
    await expect(liquidityPools.addLiquidity(mockToken.address, amount))
      .to.emit(liquidityPools, 'LiquidityAdded')
      .withArgs(mockToken.address, owner.address, amount);

    await expect(liquidityPools.removeLiquidity(mockToken.address, amount))
      .to.emit(liquidityPools, 'LiquidityRemoved')
      .withArgs(mockToken.address, owner.address, amount);

    expect(await liquidityPools.providedLiquidity(mockToken.address)).to.equal(0);
    expect(await liquidityPools.availableLiquidity(mockToken.address)).to.equal(0);
    expect((await liquidityPools.liquidityPositions(mockToken.address, owner.address)).balance).to.equal(0);
  });

  it('should not remove liquidity more than provided', async function () {
    const [owner] = await ethers.getSigners();
    const initialRequiredApprovals = 1;

    const amount = '100000';
    const amountRemove = '1000000';

    const { liquidityPools, tokenManager, mockToken } = await deployBridge(
      owner.address,
      [owner.address],
      initialRequiredApprovals
    );

    expect(await tokenManager.isTokenSupported(mockToken.address)).to.equal(true);

    await mockToken.approve(liquidityPools.address, amount);
    await expect(liquidityPools.addLiquidity(mockToken.address, amount))
      .to.emit(liquidityPools, 'LiquidityAdded')
      .withArgs(mockToken.address, owner.address, amount);

    await expect(liquidityPools.removeLiquidity(mockToken.address, amountRemove)).to.be.revertedWith(
      'IERC20: amount more than contract balance'
    );
  });

  it('should transfer token', async function () {
    const [owner, receiver] = await ethers.getSigners();
    const initialRequiredApprovals = 1;

    const amount = '100000';

    const { liquidityPools, tokenManager, mockToken, bridge } = await deployBridge(
      owner.address,
      [owner.address],
      initialRequiredApprovals
    );

    expect(await tokenManager.isTokenSupported(mockToken.address)).to.equal(true);

    await mockToken.approve(bridge.address, amount);
    await mockToken.approve(liquidityPools.address, amount);
    await liquidityPools.addLiquidity(mockToken.address, amount);

    await bridge.deposit(mockToken.address, receiver.address, amount);
  });

  it('should collect fees', async function () {
    const [owner, v1, receiver] = await ethers.getSigners();
    const initialRequiredApprovals = 1;
    const amount = '1000000';
    const txHash = '0x54c96e7f79d5fd653951c49783fc2fa7299f14c01a5a3a03f8bfb55eecb2751f';
    const fee = '10000';
    const transferAmount = '990000';

    const { liquidityPools, tokenManager, mockToken, bridge } = await deployBridge(
      owner.address,
      [v1.address],
      initialRequiredApprovals
    );

    expect(await tokenManager.isTokenSupported(mockToken.address)).to.equal(true);

    await mockToken.approve(bridge.address, amount);
    await mockToken.approve(liquidityPools.address, amount);
    await liquidityPools.addLiquidity(mockToken.address, amount);

    await bridge.deposit(mockToken.address, receiver.address, amount);

    const bridge1 = await ethers.getContractAt('Bridge', bridge.address, v1);

    await bridge1.approveTransfer(txHash, mockToken.address, receiver.address, amount);

    expect(await liquidityPools.collectedFees(mockToken.address)).to.equal(fee);
    expect(await mockToken.balanceOf(receiver.address)).to.equal(transferAmount);
  });

  it('should claimed rewards', async function () {
    const [owner, receiver] = await ethers.getSigners();
    const initialRequiredApprovals = 1;
    const amount = '1000000';
    const fee = '10000';
    const txHash = '0x54c96e7f79d5fd653951c49783fc2fa7299f14c01a5a3a03f8bfb55eecb2751f';

    const { liquidityPools, mockToken, bridge } = await deployBridge(
      owner.address,
      [owner.address],
      initialRequiredApprovals
    );

    await mockToken.approve(bridge.address, amount);
    await mockToken.approve(liquidityPools.address, amount);
    await liquidityPools.addLiquidity(mockToken.address, amount);

    await bridge.approveTransfer(txHash, mockToken.address, receiver.address, amount);

    expect(await liquidityPools.rewardsOwing(mockToken.address)).to.equal(fee);
    await liquidityPools.claimedRewards(mockToken.address);
    expect(await liquidityPools.rewardsOwing(mockToken.address)).to.equal(0);
  });

  it('should add and remove liquidity several times', async function () {
    const [owner] = await ethers.getSigners();
    const initialRequiredApprovals = 1;
    const amount = '1000000';
    const amountLiquidity = '100';

    const { liquidityPools, mockToken, bridge } = await deployBridge(
      owner.address,
      [owner.address],
      initialRequiredApprovals
    );

    await mockToken.approve(bridge.address, amount);
    await mockToken.approve(liquidityPools.address, amount);
    await liquidityPools.addLiquidity(mockToken.address, amountLiquidity);
    await liquidityPools.addLiquidity(mockToken.address, amountLiquidity);
    await liquidityPools.addLiquidity(mockToken.address, amountLiquidity);

    expect(await liquidityPools.providedLiquidity(mockToken.address)).to.equal(3 * Number(amountLiquidity));

    await liquidityPools.removeLiquidity(mockToken.address, amountLiquidity);
    await liquidityPools.removeLiquidity(mockToken.address, amountLiquidity);
    await liquidityPools.removeLiquidity(mockToken.address, amountLiquidity);
    expect(await liquidityPools.providedLiquidity(mockToken.address)).to.equal(0);
  });
});
