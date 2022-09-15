import { expect } from 'chai';
import { utils } from 'ethers';
import { ethers } from 'hardhat';
import { deployBridge, deployBridgeWithMocks } from '../utils/deploy';

describe('Bridge', function () {
  it('should change token manager', async function () {
    const { bridge, tokenManager } = await deployBridge();

    await expect(bridge.setTokenManager(tokenManager.address))
      .to.emit(bridge, 'TokenManagerUpdated')
      .withArgs(tokenManager.address);
    expect(await bridge.tokenManager()).to.equal(tokenManager.address);
  });

  it('should change liquidity pools', async function () {
    const [, v1] = await ethers.getSigners();

    const { bridge, liquidityPools } = await deployBridge();
    const newLiquidityPools = await ethers.getContractAt('LiquidityPools', liquidityPools.address, v1);

    await expect(bridge.setLiquidityPools(newLiquidityPools.address))
      .to.emit(bridge, 'LiquidityPoolsUpdated')
      .withArgs(newLiquidityPools.address);

    expect(await bridge.liquidityPools()).to.equal(newLiquidityPools.address);
  });

  it('should deposit tokens to bridge', async function () {
    const [sender, receiver] = await ethers.getSigners();

    const depositAmount = utils.parseEther('1');
    const transferAmount = utils.parseEther('0.99');
    const { mockChainId, mockToken, bridge, liquidityPools, relayBridge } = await deployBridgeWithMocks();
    const sourceChain = ethers.provider.network.chainId;
    const gasLimit = (await ethers.provider.getBlock(0)).gasLimit;
    const nonce = 0;

    await mockToken.approve(bridge.address, depositAmount);
    await mockToken.approve(liquidityPools.address, depositAmount);
    await bridge.deposit(mockToken.address, mockChainId, receiver.address, depositAmount);

    const abiCoder = ethers.utils.defaultAbiCoder;
    const data = abiCoder.encode(
      ['address', 'address', 'uint256', 'address', 'uint256'],
      [sender.address, mockToken.address, mockChainId, receiver.address, transferAmount]
    );

    const dataForHash = abiCoder.encode(
      ['address', 'uint256', 'uint256', 'uint256', 'bytes', 'uint256'],
      [bridge.address, sourceChain, mockChainId, gasLimit, data, nonce]
    );
    const hash = ethers.utils.keccak256(dataForHash);

    expect(await mockToken.balanceOf(liquidityPools.address)).to.equal(transferAmount);
    expect(await relayBridge.sentData(hash)).to.equals(data);
  });

  it('should revert deposit', async function () {
    const [sender, receiver] = await ethers.getSigners();

    const initialSupply = utils.parseEther('1');
    const depositAmount = utils.parseEther('1');
    const transferAmount = utils.parseEther('0.99');
    const NATIVE_TOKEN = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
    const leader = '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266';
    const sourceChainId = ethers.provider.network.chainId;
    const gasLimit = (await ethers.provider.getBlock(0)).gasLimit;

    const { mockChainId, mockToken, bridge, liquidityPools, mockMintableBurnableToken, relayBridge, tokenManager } =
      await deployBridgeWithMocks();

    const abiCoder = ethers.utils.defaultAbiCoder;
    const data = abiCoder.encode(
      ['address', 'address', 'uint256', 'address', 'uint256'],
      [sender.address, mockToken.address, mockChainId, receiver.address, transferAmount]
    );

    const dataMintableToken = abiCoder.encode(
      ['address', 'address', 'uint256', 'address', 'uint256'],
      [sender.address, mockMintableBurnableToken.address, mockChainId, receiver.address, transferAmount]
    );

    const dataNativeToken = abiCoder.encode(
      ['address', 'address', 'uint256', 'address', 'uint256'],
      [sender.address, NATIVE_TOKEN, mockChainId, receiver.address, transferAmount]
    );

    const hashToken = await relayBridge.dataHash(bridge.address, sourceChainId, mockChainId, gasLimit, data, 0);

    const hashMintToken = await relayBridge.dataHash(
      bridge.address,
      sourceChainId,
      mockChainId,
      gasLimit,
      dataMintableToken,
      1
    );

    const hashNativeToken = await relayBridge.dataHash(
      bridge.address,
      sourceChainId,
      mockChainId,
      gasLimit,
      dataNativeToken,
      2
    );

    const mockSenderBalanceBeforeDeposit = await mockToken.balanceOf(sender.address);
    const mockLiquidityBalanceBeforeDeposit = await mockToken.balanceOf(liquidityPools.address);

    await mockToken.approve(bridge.address, depositAmount);
    await mockToken.approve(liquidityPools.address, depositAmount);
    await bridge.deposit(mockToken.address, mockChainId, receiver.address, depositAmount);

    const mockLiquidityBalanceAfterDeposit = await mockToken.balanceOf(liquidityPools.address);
    expect(mockLiquidityBalanceAfterDeposit.sub(mockLiquidityBalanceBeforeDeposit)).to.equal(transferAmount);

    await relayBridge.revertSend(bridge.address, mockChainId, gasLimit, data, 0, leader);

    expect(await relayBridge.sent(hashToken)).to.equals(true);

    const mockExpectedBalance = mockSenderBalanceBeforeDeposit.sub(depositAmount).add(transferAmount);
    expect(await mockToken.balanceOf(sender.address)).to.equal(mockExpectedBalance);

    const mintableSenderBalanceBeforeDeposit = await mockMintableBurnableToken.balanceOf(sender.address);

    await mockMintableBurnableToken.mint(sender.address, initialSupply);
    await mockMintableBurnableToken.transferOwnership(bridge.address);
    await mockMintableBurnableToken.approve(bridge.address, depositAmount);
    await bridge.deposit(mockMintableBurnableToken.address, mockChainId, receiver.address, depositAmount);

    await expect(relayBridge.revertSend(bridge.address, mockChainId, gasLimit, dataMintableToken, 1, leader))
      .emit(relayBridge, 'Reverted')
      .withArgs(hashMintToken, sourceChainId, mockChainId);

    expect(await relayBridge.sent(hashMintToken)).to.equals(true);

    const mintableExpectedBalance = mintableSenderBalanceBeforeDeposit
      .add(initialSupply)
      .sub(depositAmount)
      .add(transferAmount);
    expect(await mockMintableBurnableToken.balanceOf(sender.address)).to.equal(mintableExpectedBalance);

    await tokenManager.setEnabled(NATIVE_TOKEN, true);
    expect(await tokenManager.isTokenEnabled(NATIVE_TOKEN)).to.equal(true);
    await liquidityPools.addNativeLiquidity({ value: depositAmount });

    const nativeSenderBalanceBeforeDeposit = await ethers.provider.getBalance(sender.address);
    const nativeLiquidityBalanceBeforeDeposit = await ethers.provider.getBalance(liquidityPools.address);

    const nativeDepositTx = await (
      await bridge.depositNative(mockChainId, receiver.address, { value: depositAmount })
    ).wait();
    const depositTxFee = nativeDepositTx.gasUsed.mul(nativeDepositTx.effectiveGasPrice);

    const nativeSenderBalanceAfterDeposit = await ethers.provider.getBalance(sender.address);
    const nativeExpectedBalanceDeposit = nativeSenderBalanceBeforeDeposit.sub(depositAmount.add(depositTxFee));
    expect(nativeSenderBalanceAfterDeposit).to.equal(nativeExpectedBalanceDeposit);

    const nativeLiquidityBalanceAfterDeposit = await ethers.provider.getBalance(liquidityPools.address);
    expect(nativeLiquidityBalanceAfterDeposit.sub(nativeLiquidityBalanceBeforeDeposit)).to.equal(transferAmount);

    const nativeSenderBalanceBeforeRevert = await ethers.provider.getBalance(sender.address);

    const nativeRevertTx = await (
      await relayBridge.revertSend(bridge.address, mockChainId, gasLimit, dataNativeToken, 2, leader)
    ).wait();
    const revertTxFee = nativeRevertTx.gasUsed.mul(nativeRevertTx.effectiveGasPrice);

    expect(await relayBridge.sent(hashNativeToken)).to.equals(true);

    const nativeSenderBalanceAfterRevert = await ethers.provider.getBalance(sender.address);
    const nativeExpectedBalanceRevert = nativeSenderBalanceBeforeRevert.add(transferAmount.sub(revertTxFee));
    expect(nativeSenderBalanceAfterRevert).to.equal(nativeExpectedBalanceRevert);
  });

  it('should emit event Reverted native token', async function () {
    const [sender, receiver] = await ethers.getSigners();

    const depositAmount = utils.parseEther('1');
    const transferAmount = utils.parseEther('0.99');
    const NATIVE_TOKEN = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
    const leader = '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266';
    const sourceChainId = ethers.provider.network.chainId;
    const gasLimit = (await ethers.provider.getBlock(0)).gasLimit;
    const nonce = 0;

    const { mockChainId, bridge, liquidityPools, relayBridge, tokenManager } = await deployBridgeWithMocks();

    const abiCoder = ethers.utils.defaultAbiCoder;
    const dataNativeToken = abiCoder.encode(
      ['address', 'address', 'uint256', 'address', 'uint256'],
      [sender.address, NATIVE_TOKEN, mockChainId, receiver.address, transferAmount]
    );

    const hashNativeToken = await relayBridge.dataHash(
      bridge.address,
      sourceChainId,
      mockChainId,
      gasLimit,
      dataNativeToken,
      nonce
    );

    await tokenManager.setEnabled(NATIVE_TOKEN, true);
    expect(await tokenManager.isTokenEnabled(NATIVE_TOKEN)).to.equal(true);
    await liquidityPools.addNativeLiquidity({ value: depositAmount });
    await bridge.depositNative(mockChainId, receiver.address, { value: depositAmount });

    await expect(relayBridge.revertSend(bridge.address, mockChainId, gasLimit, dataNativeToken, nonce, leader))
      .emit(relayBridge, 'Reverted')
      .withArgs(hashNativeToken, sourceChainId, mockChainId);
  });

  it('should execute', async function () {
    const [sender, receiver] = await ethers.getSigners();
    const depositAmount = utils.parseEther('1');
    const transferAmount = utils.parseEther('0.99');
    const leader = '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266';
    const sourceChain = ethers.provider.network.chainId;
    const gasLimit = (await ethers.provider.getBlock(0)).gasLimit;
    const nonce = 0;

    const { mockChainId, mockToken, bridge, liquidityPools, tokenManager, relayBridge } = await deployBridgeWithMocks();

    const abiCoder = ethers.utils.defaultAbiCoder;
    const data = abiCoder.encode(
      ['address', 'address', 'uint256', 'address', 'uint256'],
      [sender.address, mockToken.address, mockChainId, receiver.address, transferAmount]
    );
    const txHash = data;

    const dataForHash = abiCoder.encode(
      ['address', 'uint256', 'uint256', 'uint256', 'bytes', 'uint256'],
      [bridge.address, mockChainId, sourceChain, gasLimit, data, nonce]
    );
    const hash = ethers.utils.keccak256(dataForHash);

    expect(await tokenManager.isTokenEnabled(mockToken.address)).to.equal(true);

    await mockToken.approve(bridge.address, depositAmount);
    await mockToken.approve(liquidityPools.address, depositAmount);
    await liquidityPools.addLiquidity(mockToken.address, depositAmount);
    await bridge.deposit(mockToken.address, mockChainId, receiver.address, depositAmount);

    const receiverBalanceBeforeExecute = await mockToken.balanceOf(receiver.address);

    await expect(relayBridge.execute(bridge.address, mockChainId, gasLimit, data, nonce, leader))
      .to.emit(relayBridge, 'Executed')
      .withArgs(hash, mockChainId, sourceChain);

    const receiverBalanceAfterExecute = await mockToken.balanceOf(receiver.address);
    expect(receiverBalanceAfterExecute.sub(receiverBalanceBeforeExecute)).to.equal(transferAmount);

    expect(
      await bridge.isExecuted(sender.address, txHash, mockToken.address, receiver.address, transferAmount)
    ).to.equal(true);
  });

  it('should execute transfer', async function () {
    const [sender, receiver] = await ethers.getSigners();
    const depositAmount = utils.parseEther('1');
    const transferAmount = utils.parseEther('0.99');
    const leader = '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266';
    const gasLimit = (await ethers.provider.getBlock(0)).gasLimit;
    const nonce = 1;

    const { mockChainId, mockToken, bridge, liquidityPools, tokenManager, relayBridge } = await deployBridgeWithMocks();

    const abiCoder = ethers.utils.defaultAbiCoder;
    const data = abiCoder.encode(
      ['address', 'address', 'uint256', 'address', 'uint256'],
      [sender.address, mockToken.address, mockChainId, receiver.address, transferAmount]
    );
    const txHash = data;

    expect(await tokenManager.isTokenEnabled(mockToken.address)).to.equal(true);

    await mockToken.approve(bridge.address, depositAmount);
    await mockToken.approve(liquidityPools.address, depositAmount);
    await liquidityPools.addLiquidity(mockToken.address, depositAmount);
    await bridge.deposit(mockToken.address, mockChainId, receiver.address, depositAmount);

    await expect(relayBridge.execute(bridge.address, mockChainId, gasLimit, data, nonce, leader))
      .to.emit(bridge, 'Transferred')
      .withArgs(sender.address, mockToken.address, mockChainId, receiver.address, transferAmount);

    expect(
      await bridge.isExecuted(sender.address, txHash, mockToken.address, receiver.address, transferAmount)
    ).to.equal(true);
  });

  it('should deposit supported tokens to bridge', async function () {
    const [, receiver] = await ethers.getSigners();
    const depositAmount = utils.parseEther('1');
    const mintAmount = utils.parseEther('1');
    const { mockChainId, mockToken, bridge } = await deployBridgeWithMocks();

    const MockToken = await ethers.getContractFactory('MockToken');
    const mockToken2 = await MockToken.deploy('Token2', 'TKN2', mintAmount);
    await mockToken2.deployed();

    await mockToken.approve(bridge.address, depositAmount);
    await mockToken2.approve(bridge.address, depositAmount);

    await bridge.deposit(mockToken.address, mockChainId, receiver.address, depositAmount);

    await expect(bridge.deposit(mockToken2.address, mockChainId, receiver.address, depositAmount)).to.be.revertedWith(
      'TokenManager: token is not enabled'
    );
  });

  it('should mint and burn tokens', async function () {
    const [sender, receiver] = await ethers.getSigners();
    const depositAmount = utils.parseEther('1');
    const initialSupply = utils.parseEther('1');
    const transferAmount = utils.parseEther('0.99');
    const leader = '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266';
    const gasLimit = (await ethers.provider.getBlock(0)).gasLimit;
    const nonce = 1;

    const { mockChainId, mockMintableBurnableToken, bridge, relayBridge } = await deployBridgeWithMocks();

    const abiCoder = ethers.utils.defaultAbiCoder;
    const dataMintableToken = abiCoder.encode(
      ['address', 'address', 'uint256', 'address', 'uint256'],
      [sender.address, mockMintableBurnableToken.address, mockChainId, receiver.address, transferAmount]
    );

    await mockMintableBurnableToken.mint(sender.address, initialSupply);
    await mockMintableBurnableToken.transferOwnership(bridge.address);

    await expect(relayBridge.execute(bridge.address, mockChainId, gasLimit, dataMintableToken, nonce, leader))
      .to.emit(mockMintableBurnableToken, 'Transfer')
      .withArgs('0x0000000000000000000000000000000000000000', receiver.address, transferAmount);

    await mockMintableBurnableToken.approve(bridge.address, depositAmount);
    await expect(bridge.deposit(mockMintableBurnableToken.address, mockChainId, receiver.address, depositAmount))
      .emit(mockMintableBurnableToken, 'Transfer')
      .withArgs(sender.address, '0x0000000000000000000000000000000000000000', transferAmount);
  });

  it('should deposit and transfer using native currency', async function () {
    const [sender, receiver] = await ethers.getSigners();

    const depositAmount = utils.parseEther('1');
    const transferAmount = utils.parseEther('0.99');
    const NATIVE_TOKEN = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
    const leader = '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266';
    const gasLimit = (await ethers.provider.getBlock(0)).gasLimit;
    const nonce = 1;

    const { mockChainId, bridge, liquidityPools, relayBridge, tokenManager } = await deployBridgeWithMocks();

    const abiCoder = ethers.utils.defaultAbiCoder;
    const data = abiCoder.encode(
      ['address', 'address', 'uint256', 'address', 'uint256'],
      [sender.address, NATIVE_TOKEN, mockChainId, receiver.address, transferAmount]
    );
    const txHash = data;

    await tokenManager.setEnabled(NATIVE_TOKEN, true);
    expect(await tokenManager.isTokenEnabled(NATIVE_TOKEN)).to.equal(true);

    await tokenManager.setDestinationToken(mockChainId, NATIVE_TOKEN, NATIVE_TOKEN);

    await liquidityPools.addNativeLiquidity({ value: depositAmount });
    await bridge.depositNative(mockChainId, receiver.address, { value: depositAmount });

    const balanceReceiverBefore = await ethers.provider.getBalance(receiver.address);

    await relayBridge.execute(bridge.address, mockChainId, gasLimit, data, nonce, leader);
    await expect(relayBridge.execute(bridge.address, mockChainId, gasLimit, data, nonce, leader)).to.be.revertedWith(
      'RelayBridge: data already executed'
    );

    const balanceReceiverAfter = await ethers.provider.getBalance(receiver.address);
    expect(balanceReceiverAfter).to.equal(balanceReceiverBefore.add(transferAmount));

    expect(await bridge.isExecuted(sender.address, txHash, NATIVE_TOKEN, receiver.address, transferAmount)).to.equal(
      true
    );
  });
});
