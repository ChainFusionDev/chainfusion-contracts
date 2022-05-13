import { expect } from 'chai';
import { ethers } from 'hardhat';
import { utils } from 'ethers';
import { deployBridge } from '../utils/deploy';

describe('Bridge', function () {
  it('should change required approvals', async function () {
    const [owner, v1] = await ethers.getSigners();
    const initialRequiredApprovals = 1;

    const newRequiredApprovals = 2;

    const { validatorManager } = await deployBridge(owner.address, [v1.address], initialRequiredApprovals);

    await expect(validatorManager.setRequiredApprovals(newRequiredApprovals))
      .to.emit(validatorManager, 'RequiredApprovalsUpdated')
      .withArgs(newRequiredApprovals);

    expect(await validatorManager.requiredApprovals()).to.equal(newRequiredApprovals);
  });

  it('should change validators', async function () {
    const [owner, v1, v2] = await ethers.getSigners();
    const initialRequiredApprovals = 1;

    const { validatorManager } = await deployBridge(owner.address, [v1.address], initialRequiredApprovals);

    await expect(validatorManager.setValidators([v2.address]))
      .to.emit(validatorManager, 'ValidatorsUpdated')
      .withArgs([v2.address]);

    expect(await validatorManager.validators(0)).to.equal(v2.address);
  });

  it('should change token manager', async function () {
    const [owner, v1] = await ethers.getSigners();
    const initialRequiredApprovals = 1;

    const { bridge, tokenManager } = await deployBridge(owner.address, [v1.address], initialRequiredApprovals);
    const newTokenManager = await ethers.getContractAt('TokenManager', tokenManager.address, v1);

    await expect(bridge.setTokenManager(newTokenManager.address))
      .to.emit(bridge, 'TokenManagerUpdated')
      .withArgs(newTokenManager.address);

    expect(await bridge.tokenManager()).to.equal(newTokenManager.address);
  });

  it('should change validator manager', async function () {
    const [owner, v1] = await ethers.getSigners();
    const initialRequiredApprovals = 1;

    const { bridge, validatorManager } = await deployBridge(owner.address, [v1.address], initialRequiredApprovals);
    const newValidatorManager = await ethers.getContractAt('ValidatorManager', validatorManager.address, v1);

    await expect(bridge.setValidatorManager(newValidatorManager.address))
      .to.emit(bridge, 'ValidatorManagerUpdated')
      .withArgs(newValidatorManager.address);

    expect(await bridge.validatorManager()).to.equal(newValidatorManager.address);
  });

  it('should change liquidity pools', async function () {
    const [owner, v1] = await ethers.getSigners();
    const initialRequiredApprovals = 1;

    const { bridge, liquidityPools } = await deployBridge(owner.address, [v1.address], initialRequiredApprovals);
    const newLiquidityPools = await ethers.getContractAt('LiquidityPools', liquidityPools.address, v1);

    await expect(bridge.setLiquidityPools(newLiquidityPools.address))
      .to.emit(bridge, 'LiquidityPoolsUpdated')
      .withArgs(newLiquidityPools.address);

    expect(await bridge.liquidityPools()).to.equal(newLiquidityPools.address);
  });

  it('should deposit tokens to bridge', async function () {
    const [owner, v1, receiver] = await ethers.getSigners();
    const initialRequiredApprovals = 1;

    const depositAmount = '1000000000000000000';
    const transferAmount = '990000000000000000';
    const { mockToken, bridge, chainId, liquidityPools } = await deployBridge(
      owner.address,
      [v1.address],
      initialRequiredApprovals
    );

    await mockToken.approve(bridge.address, depositAmount);
    await mockToken.approve(liquidityPools.address, depositAmount);
    await bridge.deposit(mockToken.address, chainId, receiver.address, depositAmount);

    expect(await mockToken.balanceOf(liquidityPools.address)).to.equal(transferAmount);
  });

  it('should execute transfer', async function () {
    const [owner, v1, v2, v3, receiver] = await ethers.getSigners();
    const initialRequiredApprovals = 2;
    const depositAmount = '10000000000000000000';
    const sourceChainId = 123;

    const { mockToken, bridge, chainId, liquidityPools, tokenManager } = await deployBridge(
      owner.address,
      [v1.address, v2.address, v3.address],
      initialRequiredApprovals
    );

    const txHash = '0x54c96e7f79d5fd653951c49783fc2fa7299f14c01a5a3a03f8bfb55eecb2751f';

    expect(await tokenManager.isTokenEnabled(mockToken.address)).to.equal(true);

    await mockToken.approve(bridge.address, depositAmount);
    await mockToken.approve(liquidityPools.address, depositAmount);
    await liquidityPools.addLiquidity(mockToken.address, depositAmount);
    await bridge.deposit(mockToken.address, chainId, receiver.address, depositAmount);

    const id = utils.solidityKeccak256(
      ['bytes', 'address', 'address', 'uint256'],
      [txHash, mockToken.address, receiver.address, depositAmount]
    );

    const bridge1 = await ethers.getContractAt('Bridge', bridge.address, v1);

    await expect(bridge1.approveTransfer(txHash, mockToken.address, sourceChainId, receiver.address, depositAmount))
      .to.emit(bridge1, 'Approved')
      .withArgs(id, v1.address);

    const bridge2 = await ethers.getContractAt('Bridge', bridge.address, v2);
    await expect(bridge2.approveTransfer(txHash, mockToken.address, sourceChainId, receiver.address, depositAmount))
      .emit(bridge2, 'Transferred')
      .withArgs(mockToken.address, chainId, receiver.address, depositAmount, v2.address);

    const bridge3 = await ethers.getContractAt('Bridge', bridge.address, v3);
    await bridge3.approveTransfer(txHash, mockToken.address, sourceChainId, receiver.address, depositAmount);

    await expect(
      bridge.approveTransfer(txHash, mockToken.address, sourceChainId, receiver.address, depositAmount)
    ).to.be.revertedWith('only validator');

    expect(await bridge.executed(id)).to.equal(true);
  });

  it('should deposit supported tokens to bridge', async function () {
    const [owner, v1, receiver] = await ethers.getSigners();
    const initialRequiredApprovals = 2;
    const depositAmount = '10000000000000000000';
    const mintAmount = '100000000000000000000';
    const { mockToken, bridge, chainId } = await deployBridge(owner.address, [v1.address], initialRequiredApprovals);

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

  it('should check if transfer was already approved', async function () {
    const [owner, v1, receiver] = await ethers.getSigners();
    const initialRequiredApprovals = 2;
    const depositAmount = '10000000000000000000';
    const sourceChainId = 123;

    const { mockToken, bridge, chainId, liquidityPools } = await deployBridge(
      owner.address,
      [owner.address, v1.address],
      initialRequiredApprovals
    );

    const txHash = '0x54c96e7f79d5fd653951c49783fc2fa7299f14c01a5a3a03f8bfb55eecb2751f';

    await mockToken.approve(bridge.address, depositAmount);
    await mockToken.approve(liquidityPools.address, depositAmount);

    await liquidityPools.addLiquidity(mockToken.address, depositAmount);

    await bridge.deposit(mockToken.address, chainId, receiver.address, depositAmount);

    const bridge1 = await ethers.getContractAt('Bridge', bridge.address, v1);

    await bridge1.approveTransfer(txHash, mockToken.address, sourceChainId, receiver.address, depositAmount);
    await bridge.approveTransfer(txHash, mockToken.address, sourceChainId, receiver.address, depositAmount);

    expect(await bridge.isApproved(txHash, mockToken.address, receiver.address, depositAmount)).to.equal(true);
  });

  it('should mint and burn tokens', async function () {
    const [owner, receiver] = await ethers.getSigners();
    const initialRequiredApprovals = 1;
    const depositAmount = '10000000000000000000';
    const initialSupply = '100000000000000000000';
    const transferAmount = '9990000000000000000';
    const sourceChainId = 123;

    const { bridge, chainId, tokenManager } = await deployBridge(
      owner.address,
      [owner.address],
      initialRequiredApprovals
    );

    const txHash = '0x54c96e7f79d5fd653951c49783fc2fa7299f14c01a5a3a03f8bfb55eecb2751f';

    const MintableBurnableMockToken = await ethers.getContractFactory('MintableBurnableMockToken');
    const mintableBurnableMockToken = await MintableBurnableMockToken.deploy('Token', 'TKN');
    await mintableBurnableMockToken.deployed();

    await tokenManager.setEnabled(mintableBurnableMockToken.address, true);
    await tokenManager.setMintable(mintableBurnableMockToken.address, true);
    await mintableBurnableMockToken.mint(owner.address, initialSupply);

    await mintableBurnableMockToken.transferOwnership(bridge.address);

    await expect(
      bridge.approveTransfer(txHash, mintableBurnableMockToken.address, sourceChainId, receiver.address, depositAmount)
    )
      .emit(mintableBurnableMockToken, 'Transfer')
      .withArgs('0x0000000000000000000000000000000000000000', receiver.address, depositAmount);

    await mintableBurnableMockToken.approve(bridge.address, depositAmount);
    await expect(bridge.deposit(mintableBurnableMockToken.address, chainId, receiver.address, depositAmount))
      .emit(mintableBurnableMockToken, 'Transfer')
      .withArgs(owner.address, '0x0000000000000000000000000000000000000000', transferAmount);
  });

  it('should deposit and approve transfer using native currency', async function () {
    const [owner, receiver] = await ethers.getSigners();
    const initialRequiredApprovals = 1;
    const NATIVE_TOKEN = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
    const amount = '1000000000000000000';
    const fee = '990000000000000000';

    const { liquidityPools, tokenManager, bridge, chainId } = await deployBridge(
      owner.address,
      [owner.address],
      initialRequiredApprovals
    );

    await tokenManager.setEnabled(NATIVE_TOKEN, true);

    expect(await tokenManager.isTokenEnabled(NATIVE_TOKEN)).to.equal(true);

    await liquidityPools.addNativeLiquidity({ value: amount });

    const txHash = '0x54c96e7f79d5fd653951c49783fc2fa7299f14c01a5a3a03f8bfb55eecb2751f';

    await bridge.depositNative(chainId, receiver.address, { value: amount });
    const balanceReceiverTo = await receiver.provider!.getBalance(receiver.address);

    const balance = await owner.provider!.getBalance(liquidityPools.address);
    expect(Number(balance) - Number(amount)).to.equal(Number(fee));

    await bridge.approveTransfer(txHash, NATIVE_TOKEN, chainId, receiver.address, amount);
    const balanceReceiverAfter = await owner.provider!.getBalance(receiver.address);
    expect(Number(balanceReceiverAfter)).to.equal(Number(balanceReceiverTo) + Number(amount));
    expect(await bridge.isApproved(txHash, NATIVE_TOKEN, receiver.address, amount)).to.equal(true);
  });
});
