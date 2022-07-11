import { expect } from 'chai';
import { BigNumber } from 'ethers';
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
      initialRequiredApprovals
    );

    expect(await tokenManager.isTokenEnabled(mockToken.address)).to.equal(true);

    await mockToken.approve(bridge.address, amount);
    await mockToken.approve(liquidityPools.address, amount);
    await liquidityPools.addLiquidity(mockToken.address, amount);

    await bridge.deposit(mockToken.address, chainId, receiver.address, amount);
    expect(await mockToken.balanceOf(feeManager.address)).to.equal(fee);
  });

  it('should collect fee in native currency', async function () {
    const [owner, validator, foundation, receiver] = await ethers.getSigners();
    const amount = '1000000000000000000000';
    const tokenFee = '10000000000000000';
    const validatorReward = '10000000000000000';
    const liquidityReward = '10000000000000000';
    const foundationReward = '10000000000000000';
    const fee = '10000000000000000';
    const chainId = 2;
    const NATIVE_TOKEN = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const { liquidityPools, tokenManager, bridge, feeManager } = await deployBridge(
      owner.address,
      chainId,
      validator.address,
      foundation.address
    );
    await tokenManager.setEnabled(NATIVE_TOKEN, true);

    expect(await tokenManager.isTokenEnabled(NATIVE_TOKEN)).to.equal(true);

    await liquidityPools.addNativeLiquidity({ value: amount });
    await bridge.depositNative(chainId, receiver.address, { value: amount });

    const balanceValidatorAddressTo = await owner.provider!.getBalance(validator.address);
    const balanceFoundationAddressTo = await owner.provider!.getBalance(foundation.address);
    const balanceLiquidityPoolsTo = await owner.provider!.getBalance(feeManager.liquidityPools());

    const balance = await owner.provider!.getBalance(feeManager.address);
    expect(balance).to.equal(fee);

    await feeManager.setTokenFee(NATIVE_TOKEN, tokenFee, validatorReward, liquidityReward, foundationReward);
    await feeManager.distributeRewards(NATIVE_TOKEN);

    const balanceValidatorAddressAfter = await owner.provider!.getBalance(validator.address);
    const balanceFoundationAddressAfter = await owner.provider!.getBalance(foundation.address);
    const balanceLiquidityPoolsAfter = await owner.provider!.getBalance(feeManager.liquidityPools());

    let validatorRewards = BigNumber.from('100000000000000');
    let liquidityRewards = BigNumber.from('100000000000000');
    let foundationRewards = BigNumber.from('9800000000000000');

    expect(Number(balanceValidatorAddressAfter)).to.equal(Number(balanceValidatorAddressTo) + Number(validatorRewards));
    expect(Number(balanceFoundationAddressAfter)).to.equal(
      Number(balanceFoundationAddressTo) + Number(foundationRewards)
    );

    expect(balanceLiquidityPoolsAfter).to.equal(balanceLiquidityPoolsTo.add(liquidityRewards));

    const collectedFee = '100000000000000';
    expect(await liquidityPools.collectedFees(NATIVE_TOKEN)).to.equal(collectedFee);
    await liquidityPools.claimRewards(NATIVE_TOKEN);
    expect(await liquidityPools.collectedFees(NATIVE_TOKEN)).to.equal(0);
  });
});
