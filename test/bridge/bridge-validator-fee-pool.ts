import { expect } from 'chai';
import { BigNumber } from 'ethers';
import { ethers } from 'hardhat';
import { deployBridgeWithMocks } from '../utils/deploy';

describe('BridgeValidatorFeePool', function () {
  it('should change all setters', async function () {
    const [, validator, user] = await ethers.getSigners();
    const NATIVE_TOKEN = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
    const tokenLimit = '1000000000000';

    const { mockToken, bridgeValidatorFeePool, bridge } = await deployBridgeWithMocks();
    const bridgeValidatorFeePoolByUser = await ethers.getContractAt(
      'BridgeValidatorFeePool',
      bridgeValidatorFeePool.address,
      user
    );

    await expect(bridgeValidatorFeePoolByUser.setBridge(bridge.address)).to.be.revertedWith(
      'SignerOwnable: only signer'
    );
    await expect(bridgeValidatorFeePool.setBridge(bridge.address))
      .to.emit(bridgeValidatorFeePool, 'BridgeUpdated')
      .withArgs(bridge.address);

    expect(await bridgeValidatorFeePool.bridge()).to.be.equal(bridge.address);

    await expect(bridgeValidatorFeePoolByUser.setValidatorFeeReceiver(validator.address)).to.be.revertedWith(
      'SignerOwnable: only signer'
    );
    await expect(bridgeValidatorFeePool.setValidatorFeeReceiver(validator.address))
      .to.emit(bridgeValidatorFeePool, 'ValidatorFeeReceiverUpdated')
      .withArgs(validator.address);

    expect(await bridgeValidatorFeePool.validatorFeeReceiver()).to.be.equal(validator.address);

    await expect(bridgeValidatorFeePoolByUser.setLimitPerToken(mockToken.address, tokenLimit)).to.be.revertedWith(
      'SignerOwnable: only signer'
    );
    await expect(bridgeValidatorFeePool.setLimitPerToken(mockToken.address, tokenLimit))
      .to.emit(bridgeValidatorFeePool, 'LimitPerTokenUpdated')
      .withArgs(mockToken.address, tokenLimit);
    await expect(bridgeValidatorFeePool.setLimitPerToken(NATIVE_TOKEN, tokenLimit))
      .to.emit(bridgeValidatorFeePool, 'LimitPerTokenUpdated')
      .withArgs(NATIVE_TOKEN, tokenLimit);

    expect(await bridgeValidatorFeePool.limitPerToken(mockToken.address)).to.be.equal(tokenLimit);
    expect(await bridgeValidatorFeePool.limitPerToken(NATIVE_TOKEN)).to.be.equal(tokenLimit);
  });

  it('should collect fees', async function () {
    const [, receiver] = await ethers.getSigners();
    const NATIVE_TOKEN = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
    const depositAmount = '10000000000000000000';

    const tokenLimit = '10000000';
    const tokenFee = '10000000000000000';
    const validatorReward = '300000000000000000';
    const liquidityReward = '300000000000000000';

    const { tokenManager, liquidityPools, mockChainId, mockToken, bridgeValidatorFeePool, bridge, feeManager } =
      await deployBridgeWithMocks();

    const MockToken = await ethers.getContractFactory('MockToken');
    const mockToken2 = await MockToken.deploy('Token2', 'TKN2', depositAmount);
    await mockToken2.deployed();

    await bridgeValidatorFeePool.setLimitPerToken(mockToken.address, tokenLimit);
    await bridgeValidatorFeePool.setLimitPerToken(NATIVE_TOKEN, tokenLimit);
    expect(await bridgeValidatorFeePool.limitPerToken(mockToken.address)).to.be.equal(tokenLimit);
    expect(await bridgeValidatorFeePool.limitPerToken(NATIVE_TOKEN)).to.be.equal(tokenLimit);

    await tokenManager.setEnabled(NATIVE_TOKEN, true);
    await tokenManager.setEnabled(mockToken.address, true);
    expect(await tokenManager.isTokenEnabled(NATIVE_TOKEN)).to.equal(true);
    expect(await tokenManager.isTokenEnabled(mockToken.address)).to.equal(true);

    await mockToken.approve(bridge.address, depositAmount);
    await mockToken.approve(liquidityPools.address, depositAmount);

    await liquidityPools.addNativeLiquidity({ value: depositAmount });
    await liquidityPools.addLiquidity(mockToken.address, depositAmount);
    expect(await liquidityPools.providedLiquidity(NATIVE_TOKEN)).to.equal(depositAmount);
    expect(await liquidityPools.providedLiquidity(mockToken.address)).to.equal(depositAmount);

    await feeManager.setTokenFee(NATIVE_TOKEN, tokenFee, validatorReward, liquidityReward);
    await feeManager.setTokenFee(mockToken.address, tokenFee, validatorReward, liquidityReward);
    expect(await bridgeValidatorFeePool.limitPerToken(NATIVE_TOKEN)).to.be.equal(tokenLimit);
    expect(await bridgeValidatorFeePool.limitPerToken(mockToken.address)).to.be.equal(tokenLimit);

    await bridge.depositNative(mockChainId, receiver.address, { value: depositAmount });
    await bridge.deposit(mockToken.address, mockChainId, receiver.address, depositAmount);

    await feeManager.distributeRewards(NATIVE_TOKEN);
    await feeManager.distributeRewards(mockToken.address);

    let reward = BigNumber.from('33000000000000000');

    await expect(bridgeValidatorFeePool.collect(mockToken2.address)).to.be.revertedWith(
      'BridgeValidatorFeePool: no limit for this token'
    );

    let nativeBalanceValidatorFeePool = await ethers.provider.getBalance(bridgeValidatorFeePool.address);
    let mockBalanceValidatorFeePool = await mockToken.balanceOf(bridgeValidatorFeePool.address);

    expect(nativeBalanceValidatorFeePool).to.equal(reward);
    expect(mockBalanceValidatorFeePool).to.equal(reward);

    await expect(bridgeValidatorFeePool.collect(NATIVE_TOKEN))
      .to.emit(bridgeValidatorFeePool, 'Collected')
      .withArgs(NATIVE_TOKEN, reward);
    await expect(bridgeValidatorFeePool.collect(mockToken.address))
      .to.emit(bridgeValidatorFeePool, 'Collected')
      .withArgs(mockToken.address, reward);

    nativeBalanceValidatorFeePool = await ethers.provider.getBalance(bridgeValidatorFeePool.address);
    mockBalanceValidatorFeePool = await mockToken.balanceOf(bridgeValidatorFeePool.address);

    expect(nativeBalanceValidatorFeePool).to.equal(0);
    expect(mockBalanceValidatorFeePool).to.equal(0);
  });
});
