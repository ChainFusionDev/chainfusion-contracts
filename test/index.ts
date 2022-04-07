import { expect } from 'chai';
import { ethers } from 'hardhat';
import { utils } from 'ethers';

describe('Bridge', function () {
  it('Should change required approvals', async function () {
    const [owner, v1] = await ethers.getSigners();
    const ownerAddress = owner.address;
    const tokenManagerAddress = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const initialRequiredApprovals = 1;
    const newRequiredApprovals = 2;

    const ValidatorManager = await ethers.getContractFactory('ValidatorManager');
    const validatorManager = await ValidatorManager.deploy();
    await validatorManager.deployed();

    const Bridge = await ethers.getContractFactory('Bridge');
    const bridge = await Bridge.deploy();
    await bridge.deployed();
    await bridge.initialize(ownerAddress, validatorManager.address, tokenManagerAddress);

    await validatorManager.setRequiredApprovals(initialRequiredApprovals);
    await validatorManager.setValidators([v1.address]);

    await validatorManager.setRequiredApprovals(newRequiredApprovals);
    expect(await validatorManager.requiredApprovals()).to.equal(newRequiredApprovals);
  });

  it('Should deposit tokens to bridge', async function () {
    const [owner, v1] = await ethers.getSigners();
    const ownerAddress = owner.address;
    const destinationToken = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const chainId = 123;
    const mintAmount = '100000000000000000000';
    const depositAmount = '10000000000000000000';
    const initialRequiredApprovals = 1;

    const MockToken = await ethers.getContractFactory('MockToken');
    const mockToken = await MockToken.deploy('Token', 'TKN', mintAmount);
    await mockToken.deployed();

    const TokenManager = await ethers.getContractFactory('TokenManager');
    const tokenManager = await TokenManager.deploy();
    await tokenManager.deployed();
    await tokenManager.initialize(ownerAddress);
    await tokenManager.addSupportedToken(chainId, mockToken.address, destinationToken);

    const ValidatorManager = await ethers.getContractFactory('ValidatorManager');
    const validatorManager = await ValidatorManager.deploy();
    await validatorManager.deployed();

    const Bridge = await ethers.getContractFactory('Bridge');
    const bridge = await Bridge.deploy();
    await bridge.deployed();
    await bridge.initialize(ownerAddress, validatorManager.address, tokenManager.address);

    await validatorManager.setRequiredApprovals(initialRequiredApprovals);
    await validatorManager.setValidators([v1.address]);

    await mockToken.approve(bridge.address, depositAmount);
    await bridge.deposit(mockToken.address, chainId, depositAmount);
    expect(await mockToken.balanceOf(bridge.address)).to.equal(depositAmount);
  });

  it('Should execute transfer', async function () {
    const [owner, v1, v2, v3, receiver] = await ethers.getSigners();
    const chainId = 123;
    const mintAmount = '100000000000000000000';
    const depositAmount = '10000000000000000000';
    const initialRequiredApprovals = 2;

    const txHash = '0x54c96e7f79d5fd653951c49783fc2fa7299f14c01a5a3a03f8bfb55eecb2751f';
    const destinationToken = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const MockToken = await ethers.getContractFactory('MockToken');
    const mockToken = await MockToken.deploy('Token', 'TKN', mintAmount);
    await mockToken.deployed();

    const TokenManager = await ethers.getContractFactory('TokenManager');
    const tokenManager = await TokenManager.deploy();
    await tokenManager.deployed();

    await tokenManager.initialize(owner.address);
    await tokenManager.addSupportedToken(chainId, mockToken.address, destinationToken);

    const ValidatorManager = await ethers.getContractFactory('ValidatorManager');
    const validatorManager = await ValidatorManager.deploy();
    await validatorManager.deployed();

    const Bridge = await ethers.getContractFactory('Bridge');
    const bridge = await Bridge.deploy();
    await bridge.deployed();
    await bridge.initialize(owner.address, validatorManager.address, tokenManager.address);

    await validatorManager.setRequiredApprovals(initialRequiredApprovals);
    await validatorManager.setValidators([v1.address, v2.address, v3.address]);

    await mockToken.approve(bridge.address, depositAmount);
    await bridge.deposit(mockToken.address, chainId, depositAmount);
    expect(await mockToken.balanceOf(bridge.address)).to.equal(depositAmount);

    const id = utils.solidityKeccak256(
      ['bytes', 'address', 'address', 'uint256'],
      [txHash, mockToken.address, receiver.address, depositAmount]
    );

    const bridge1 = await ethers.getContractAt('Bridge', bridge.address, v1);

    await expect(bridge1.approveTransfer(txHash, mockToken.address, receiver.address, depositAmount))
      .to.emit(bridge1, 'Approved')
      .withArgs(id, v1.address);

    const bridge2 = await ethers.getContractAt('Bridge', bridge.address, v2);
    await expect(bridge2.approveTransfer(txHash, mockToken.address, receiver.address, depositAmount))
      .to.emit(bridge2, 'Approved')
      .withArgs(id, v2.address)
      .emit(bridge2, 'Transferred')
      .withArgs(mockToken.address, receiver.address, depositAmount, v2.address);

    const bridge3 = await ethers.getContractAt('Bridge', bridge.address, v3);
    await bridge3.approveTransfer(txHash, mockToken.address, receiver.address, depositAmount);

    await expect(bridge.approveTransfer(txHash, mockToken.address, receiver.address, depositAmount)).to.be.revertedWith(
      'only validator'
    );

    expect(await bridge.executed(id)).to.equal(true);
  });

  it('Should deposit supported tokens to bridge', async function () {
    const [owner] = await ethers.getSigners();
    const ownerAddress = owner.address;
    const destinationToken = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const chainId = 123;
    const mintAmount = '100000000000000000000';
    const depositAmount = '10000000000000000000';
    const initialRequiredApprovals = 1;

    const MockToken = await ethers.getContractFactory('MockToken');
    const mockToken = await MockToken.deploy('Token1', 'TKN1', mintAmount);
    await mockToken.deployed();

    const mockToken2 = await MockToken.deploy('Token2', 'TKN2', mintAmount);
    await mockToken2.deployed();

    const TokenManager = await ethers.getContractFactory('TokenManager');
    const tokenManager = await TokenManager.deploy();
    await tokenManager.deployed();
    await tokenManager.initialize(ownerAddress);
    await tokenManager.addSupportedToken(chainId, mockToken.address, destinationToken);

    const ValidatorManager = await ethers.getContractFactory('ValidatorManager');
    const validatorManager = await ValidatorManager.deploy();
    await validatorManager.deployed();

    const Bridge = await ethers.getContractFactory('Bridge');
    const bridge = await Bridge.deploy();
    await bridge.deployed();
    await bridge.initialize(owner.address, validatorManager.address, tokenManager.address);

    await validatorManager.setRequiredApprovals(initialRequiredApprovals);
    await validatorManager.setValidators([ownerAddress]);

    await mockToken.approve(bridge.address, depositAmount);
    await mockToken2.approve(bridge.address, depositAmount);

    await bridge.deposit(mockToken.address, chainId, depositAmount);

    await expect(bridge.deposit(mockToken2.address, chainId, depositAmount)).to.be.revertedWith(
      'Token is not supported'
    );
  });
});

describe('TokenManager', function () {
  it('Should called addSupportedToken() only by owner', async function () {
    const [, v1] = await ethers.getSigners();

    const chainId = 123;
    const tokenManagerAddress = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
    const destinationToken = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const TokenManager = await ethers.getContractFactory('TokenManager');
    const tokenManager = await TokenManager.deploy();
    await tokenManager.deployed();
    await tokenManager.initialize(v1.address);

    await expect(tokenManager.addSupportedToken(chainId, tokenManagerAddress, destinationToken)).to.be.revertedWith(
      'Ownable: caller is not the owner'
    );
  });

  it('Should adds token addresses to supportedTokens mapping', async function () {
    const [owner] = await ethers.getSigners();
    const ownerAddress = owner.address;

    const chainId = 123;
    const tokenManagerAddress = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
    const destinationToken = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const TokenManager = await ethers.getContractFactory('TokenManager');
    const tokenManager = await TokenManager.deploy();
    await tokenManager.deployed();
    await tokenManager.initialize(ownerAddress);

    await tokenManager.addSupportedToken(chainId, tokenManagerAddress, destinationToken);
    expect(await tokenManager.supportedTokens(chainId, tokenManagerAddress)).to.equal(destinationToken);
  });
});
