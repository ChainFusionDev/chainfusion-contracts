import { expect } from 'chai';
import { ethers } from 'hardhat';
import { deployBridge } from '../utils/deploy';

describe('LiquidityPools', function () {
  it('should change token manager', async function () {
    const [owner, v1] = await ethers.getSigners();
    const initialRequiredApprovals = 1;

    const { tokenManager, liquidityPools } = await deployBridge(owner.address, [v1.address], initialRequiredApprovals);
    const newTokenManager = await ethers.getContractAt('TokenManager', tokenManager.address, v1);

    await expect(liquidityPools.setTokenManager(newTokenManager.address))
      .to.emit(liquidityPools, 'TokenManagerUpdated')
      .withArgs(newTokenManager.address);

    expect(await liquidityPools.tokenManager()).to.equal(newTokenManager.address);
  });

  it('should change fee percentage', async function () {
    const [owner, v1] = await ethers.getSigners();
    const initialRequiredApprovals = 1;
    const newFeePercentage = '10000';

    const { liquidityPools } = await deployBridge(owner.address, [v1.address], initialRequiredApprovals);

    await expect(liquidityPools.setFeePercentage(newFeePercentage))
      .to.emit(liquidityPools, 'FeePercentageUpdated')
      .withArgs(newFeePercentage);

    expect(await liquidityPools.feePercentage()).to.equal(newFeePercentage);
  });

  it('should add liquidity for supported token', async function () {
    const [owner] = await ethers.getSigners();
    const initialRequiredApprovals = 1;

    const amount = '100000';

    const { liquidityPools, tokenManager, mockToken } = await deployBridge(
      owner.address,
      [owner.address],
      initialRequiredApprovals
    );

    expect(await tokenManager.isTokenEnabled(mockToken.address)).to.equal(true);

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

    expect(await tokenManager.isTokenEnabled(mockToken.address)).to.equal(true);

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

    expect(await tokenManager.isTokenEnabled(mockToken.address)).to.equal(true);

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
    const amount = '10000000000000000000';

    const { liquidityPools, tokenManager, mockToken, bridge, chainId } = await deployBridge(
      owner.address,
      [owner.address],
      initialRequiredApprovals
    );

    expect(await tokenManager.isTokenEnabled(mockToken.address)).to.equal(true);

    await mockToken.approve(bridge.address, amount);
    await mockToken.approve(liquidityPools.address, amount);
    await liquidityPools.addLiquidity(mockToken.address, amount);

    await bridge.deposit(mockToken.address, chainId, receiver.address, amount);
  });

  it('should collect fees', async function () {
    const [owner, receiver] = await ethers.getSigners();
    const initialRequiredApprovals = 1;
    const amount = '10000000000000000000';
    const txHash = '0x54c96e7f79d5fd653951c49783fc2fa7299f14c01a5a3a03f8bfb55eecb2751f';
    const tokenFee = '10000';
    const validatorReward = '10000';
    const liquidityReward = '10000';
    const foundationReward = '10000';
    const fee = '100';
    const transferAmount = '10000000000000000000';
    const sourceChainId = 5;

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

    await bridge.approveTransfer(txHash, mockToken.address, sourceChainId, receiver.address, amount);
    await feeManager.setTokenFee(mockToken.address, tokenFee, validatorReward, liquidityReward, foundationReward);
    await feeManager.distributeRewards(mockToken.address);
    expect(await liquidityPools.collectedFees(mockToken.address)).to.equal(fee);
    expect(await mockToken.balanceOf(receiver.address)).to.equal(transferAmount);
  });

  it('should claim rewards', async function () {
    const [owner, receiver] = await ethers.getSigners();
    const initialRequiredApprovals = 1;
    const amount = '10000000000000000000';
    const tokenFee = '10000';
    const validatorReward = '10000';
    const liquidityReward = '10000';
    const foundationReward = '10000';
    const fee = '100';
    const sourceChainId = 5;
    const txHash = '0x54c96e7f79d5fd653951c49783fc2fa7299f14c01a5a3a03f8bfb55eecb2751f';

    const { liquidityPools, mockToken, bridge, feeManager } = await deployBridge(
      owner.address,
      [owner.address],
      initialRequiredApprovals
    );

    await mockToken.approve(bridge.address, amount);
    await mockToken.approve(liquidityPools.address, amount);
    await liquidityPools.addLiquidity(mockToken.address, amount);

    await bridge.deposit(mockToken.address, sourceChainId, receiver.address, amount);

    await bridge.approveTransfer(txHash, mockToken.address, sourceChainId, receiver.address, amount);
    await feeManager.setTokenFee(mockToken.address, tokenFee, validatorReward, liquidityReward, foundationReward);
    await feeManager.distributeRewards(mockToken.address);
    expect(await liquidityPools.collectedFees(mockToken.address)).to.equal(fee);
    expect(await liquidityPools.rewardsOwing(mockToken.address)).to.equal(fee);
    await liquidityPools.claimRewards(mockToken.address);
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
