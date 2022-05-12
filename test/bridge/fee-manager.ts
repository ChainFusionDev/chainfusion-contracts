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

  it('should distribute the reward equally', async function () {
    const [owner, v1, v2, receiver] = await ethers.getSigners();
    const initialRequiredApprovals = 2;
    const amountTransfer = '1000000000000000000';
    const amount = '100000000000000000';
    const tokenFee = '10000';
    const validatorReward = '10000';
    const liquidityReward = '10000';
    const foundationReward = '10000';
    const fee = '100';
    const sourceChainId = 5;
    const txHash = '0x54c96e7f79d5fd653951c49783fc2fa7299f14c01a5a3a03f8bfb55eecb2751f';

    const { liquidityPools, mockToken, bridge, feeManager } = await deployBridge(
      owner.address,
      [v1.address, v2.address],
      initialRequiredApprovals
    );

    await mockToken.transfer(v1.address, amountTransfer);
    await mockToken.transfer(v2.address, amountTransfer);

    const mockToken1 = await ethers.getContractAt('MockToken', mockToken.address, v1);
    const mockToken2 = await ethers.getContractAt('MockToken', mockToken.address, v2);

    const bridge1 = await ethers.getContractAt('Bridge', bridge.address, v1);
    const bridge2 = await ethers.getContractAt('Bridge', bridge.address, v2);

    const liquidityPools1 = await ethers.getContractAt('LiquidityPools', liquidityPools.address, v1);
    const liquidityPools2 = await ethers.getContractAt('LiquidityPools', liquidityPools.address, v2);

    await mockToken1.approve(liquidityPools1.address, amount);
    await mockToken2.approve(liquidityPools2.address, amount);

    await mockToken1.approve(bridge1.address, amount);
    await mockToken2.approve(bridge2.address, amount);

    await liquidityPools1.addLiquidity(mockToken1.address, amount);
    await liquidityPools2.addLiquidity(mockToken2.address, amount);

    await bridge1.deposit(mockToken1.address, sourceChainId, receiver.address, amount);

    await bridge1.approveTransfer(txHash, mockToken.address, sourceChainId, receiver.address, amount);
    await bridge2.approveTransfer(txHash, mockToken.address, sourceChainId, receiver.address, amount);

    await feeManager.setTokenFee(mockToken.address, tokenFee, validatorReward, liquidityReward, foundationReward);
    await feeManager.distributeRewards(mockToken.address);

    expect(await liquidityPools.collectedFees(mockToken.address)).to.equal(fee);

    await liquidityPools1.claimRewards(mockToken1.address);
    expect(await liquidityPools.collectedFees(mockToken.address)).to.equal(Number(fee) / 2);

    await liquidityPools2.claimRewards(mockToken2.address);
    expect(await liquidityPools.collectedFees(mockToken.address)).to.equal(0);
  });

  it('should fees are collected in native currency', async function () {
    const [owner, receiver] = await ethers.getSigners();
    const initialRequiredApprovals = 1;
    const amount = '10000000000000000000';
    const fee = '10000000000000000';
    const NATIVE_TOKEN = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const { liquidityPools, tokenManager, bridge, chainId, feeManager } = await deployBridge(
      owner.address,
      [owner.address],
      initialRequiredApprovals
    );
    await tokenManager.setEnabled(NATIVE_TOKEN, true);

    expect(await tokenManager.isTokenEnabled(NATIVE_TOKEN)).to.equal(true);

    await liquidityPools.addNativeLiquidity({ value: amount });
    await bridge.depositNative(chainId, receiver.address, { value: amount });

    const balance = await owner.provider!.getBalance(feeManager.address);

    expect(balance).to.equal(fee);
  });
});
