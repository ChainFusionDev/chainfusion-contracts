import { expect } from 'chai';
import { ethers } from 'hardhat';
import { deployBridge } from '../utils/deploy';

describe('Bridge', function () {
  it('should change token manager', async function () {
    const [validator] = await ethers.getSigners();

    const { bridge, tokenManager } = await deployBridge(validator.address);

    await expect(bridge.setTokenManager(tokenManager.address))
      .to.emit(bridge, 'TokenManagerUpdated')
      .withArgs(tokenManager.address);
    expect(await bridge.tokenManager()).to.equal(tokenManager.address);
  });

  it('should change liquidity pools', async function () {
    const [owner, v1] = await ethers.getSigners();

    const { bridge, liquidityPools } = await deployBridge(owner.address);
    const newLiquidityPools = await ethers.getContractAt('LiquidityPools', liquidityPools.address, v1);

    await expect(bridge.setLiquidityPools(newLiquidityPools.address))
      .to.emit(bridge, 'LiquidityPoolsUpdated')
      .withArgs(newLiquidityPools.address);

    expect(await bridge.liquidityPools()).to.equal(newLiquidityPools.address);
  });

  it('should deposit tokens to bridge', async function () {
    const [owner, receiver] = await ethers.getSigners();

    const depositAmount = '1000000000000000000';
    const transferAmount = '990000000000000000';
    const { mockToken, bridge, chainId, liquidityPools } = await deployBridge(owner.address);

    await mockToken.approve(bridge.address, depositAmount);
    await mockToken.approve(liquidityPools.address, depositAmount);
    await bridge.deposit(mockToken.address, chainId, receiver.address, depositAmount);

    expect(await mockToken.balanceOf(liquidityPools.address)).to.equal(transferAmount);
  });

  it('should execute transfer', async function () {
    const [validator, receiver] = await ethers.getSigners();
    const depositAmount = '10000000000000000000';

    const { mockToken, bridge, chainId, liquidityPools, tokenManager } = await deployBridge(validator.address);

    const txHash = '0x54c96e7f79d5fd653951c49783fc2fa7299f14c01a5a3a03f8bfb55eecb2751f';

    expect(await tokenManager.isTokenEnabled(mockToken.address)).to.equal(true);

    await mockToken.approve(bridge.address, depositAmount);
    await mockToken.approve(liquidityPools.address, depositAmount);
    await liquidityPools.addLiquidity(mockToken.address, depositAmount);
    await bridge.deposit(mockToken.address, chainId, receiver.address, depositAmount);

    await expect(bridge.executeTransfer(txHash, mockToken.address, chainId, receiver.address, depositAmount))
      .emit(bridge, 'Transferred')
      .withArgs(mockToken.address, chainId, receiver.address, depositAmount, validator.address);

    await bridge.executeTransfer(txHash, mockToken.address, chainId, receiver.address, depositAmount);

    expect(await bridge.isExecuted(txHash, mockToken.address, receiver.address, depositAmount)).to.equal(true);
  });

  it('should deposit supported tokens to bridge', async function () {
    const [validator, receiver] = await ethers.getSigners();
    const depositAmount = '10000000000000000000';
    const mintAmount = '100000000000000000000';
    const { mockToken, bridge, chainId } = await deployBridge(validator.address);

    const MockToken = await ethers.getContractFactory('MockToken');
    const mockToken2 = await MockToken.deploy('Token2', 'TKN2', mintAmount);
    await mockToken2.deployed();

    await mockToken.approve(bridge.address, depositAmount);
    await mockToken2.approve(bridge.address, depositAmount);

    await bridge.deposit(mockToken.address, chainId, receiver.address, depositAmount);

    await expect(bridge.deposit(mockToken2.address, chainId, receiver.address, depositAmount)).to.be.revertedWith(
      'TokenManager: token is not enabled'
    );
  });

  it('should mint and burn tokens', async function () {
    const [validator, receiver] = await ethers.getSigners();
    const depositAmount = '10000000000000000000';
    const initialSupply = '100000000000000000000';
    const transferAmount = '9990000000000000000';
    const sourceChainId = 123;

    const { bridge, chainId, tokenManager } = await deployBridge(validator.address);

    const txHash = '0x54c96e7f79d5fd653951c49783fc2fa7299f14c01a5a3a03f8bfb55eecb2751f';

    const MintableBurnableMockToken = await ethers.getContractFactory('MintableBurnableMockToken');
    const mintableBurnableMockToken = await MintableBurnableMockToken.deploy('Token', 'TKN');
    await mintableBurnableMockToken.deployed();

    await tokenManager.setEnabled(mintableBurnableMockToken.address, true);
    await tokenManager.setMintable(mintableBurnableMockToken.address, true);
    await mintableBurnableMockToken.mint(validator.address, initialSupply);

    await mintableBurnableMockToken.transferOwnership(bridge.address);

    await expect(
      bridge.executeTransfer(txHash, mintableBurnableMockToken.address, sourceChainId, receiver.address, depositAmount)
    )
      .emit(mintableBurnableMockToken, 'Transfer')
      .withArgs('0x0000000000000000000000000000000000000000', receiver.address, depositAmount);

    await mintableBurnableMockToken.approve(bridge.address, depositAmount);
    await expect(bridge.deposit(mintableBurnableMockToken.address, chainId, receiver.address, depositAmount))
      .emit(mintableBurnableMockToken, 'Transfer')
      .withArgs(validator.address, '0x0000000000000000000000000000000000000000', transferAmount);
  });

  it('should deposit and transfer using native currency', async function () {
    const [validator, receiver] = await ethers.getSigners();
    const NATIVE_TOKEN = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
    const amount = '1000000000000000000';
    const fee = '990000000000000000';

    const { liquidityPools, tokenManager, bridge, chainId } = await deployBridge(validator.address);

    await tokenManager.setEnabled(NATIVE_TOKEN, true);

    expect(await tokenManager.isTokenEnabled(NATIVE_TOKEN)).to.equal(true);

    await liquidityPools.addNativeLiquidity({ value: amount });

    const txHash = '0x54c96e7f79d5fd653951c49783fc2fa7299f14c01a5a3a03f8bfb55eecb2751f';

    await bridge.depositNative(chainId, receiver.address, { value: amount });
    const balanceReceiverTo = await ethers.provider.getBalance(receiver.address);

    const balance = await ethers.provider.getBalance(liquidityPools.address);
    expect(Number(balance) - Number(amount)).to.equal(Number(fee));

    await bridge.executeTransfer(txHash, NATIVE_TOKEN, chainId, receiver.address, amount);
    const balanceReceiverAfter = await ethers.provider.getBalance(receiver.address);
    expect(Number(balanceReceiverAfter)).to.equal(Number(balanceReceiverTo) + Number(amount));
    expect(await bridge.isExecuted(txHash, NATIVE_TOKEN, receiver.address, amount)).to.equal(true);
  });
});
