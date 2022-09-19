import { expect } from 'chai';
import { utils } from 'ethers';
import { ethers } from 'hardhat';
import { deployBridgeWithMocks } from '../utils/deploy';

describe('FeeManager', function () {
  it('should check if fee is transferred to FeeManager contract', async function () {
    const [, receiver] = await ethers.getSigners();
    const amount = utils.parseEther('1');
    const fee = utils.parseEther('0.01');

    const { mockToken, mockChainId, liquidityPools, tokenManager, bridge, feeManager } = await deployBridgeWithMocks();

    expect(await tokenManager.isTokenEnabled(mockToken.address)).to.equal(true);

    await mockToken.approve(bridge.address, amount);
    await mockToken.approve(liquidityPools.address, amount);
    await liquidityPools.addLiquidity(mockToken.address, amount);

    await bridge.deposit(mockToken.address, mockChainId, receiver.address, amount);
    expect(await mockToken.balanceOf(feeManager.address)).to.equal(fee);
  });

  it('should collect fee in native currency', async function () {
    const [deployer, foundation, receiver, user] = await ethers.getSigners();
    const amount = '1000000000000000000000';
    const tokenFee = '10000000000000000';
    const fee = '10000000000000000';
    const validatorReward = '300000000000000000';
    const liquidityReward = '300000000000000000';
    const tokenLimit = '10000000000000';
    const NATIVE_TOKEN = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const { mockChainId, liquidityPools, tokenManager, bridge, feeManager, bridgeValidatorFeePool } =
      await deployBridgeWithMocks({
        foundationAddress: foundation.address,
      });

    await tokenManager.setEnabled(NATIVE_TOKEN, true);
    expect(await tokenManager.isTokenEnabled(NATIVE_TOKEN)).to.equal(true);

    await bridgeValidatorFeePool.setLimitPerToken(NATIVE_TOKEN, tokenLimit);

    await liquidityPools.addNativeLiquidity({ value: amount });
    await bridge.depositNative(mockChainId, receiver.address, { value: amount });

    await feeManager.setTokenFee(NATIVE_TOKEN, tokenFee, validatorReward, liquidityReward);

    const balanceValidatorFeePoolBefore = await ethers.provider.getBalance(bridgeValidatorFeePool.address);
    const balanceLiquidityPoolsBefore = await ethers.provider.getBalance(feeManager.liquidityPools());
    const balanceFoundationAddressBefore = await ethers.provider.getBalance(foundation.address);

    const balance = await ethers.provider.getBalance(feeManager.address);
    expect(balance).to.equal(fee);

    const userFeeManager = await ethers.getContractAt('FeeManager', feeManager.address, user);
    await userFeeManager.distributeRewards(NATIVE_TOKEN);

    const balanceValidatorFeePoolAfter = await ethers.provider.getBalance(bridgeValidatorFeePool.address);
    const balanceLiquidityPoolsAfter = await ethers.provider.getBalance(feeManager.liquidityPools());
    const balanceFoundationAddressAfter = await ethers.provider.getBalance(foundation.address);

    let validatorRewards = utils.parseEther('0.003');
    let liquidityRewards = utils.parseEther('0.003');
    let foundationRewards = utils.parseEther('0.004');

    expect(balanceValidatorFeePoolAfter).to.equal(balanceValidatorFeePoolBefore.add(validatorRewards));
    expect(balanceLiquidityPoolsAfter).to.equal(balanceLiquidityPoolsBefore.add(liquidityRewards));
    expect(balanceFoundationAddressAfter).to.equal(balanceFoundationAddressBefore.add(foundationRewards));

    expect(await (await liquidityPools.liquidityPositions(NATIVE_TOKEN, deployer.address)).balance).to.equal(
      liquidityRewards.add(amount)
    );
  });
});
