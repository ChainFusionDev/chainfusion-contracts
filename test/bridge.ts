import { expect } from 'chai';
import { ethers } from 'hardhat';
import { utils } from 'ethers';
import { deployBridge } from './utils/deploy';

describe('Bridge', function () {
  it('Should change required approvals', async function () {
    const [owner, v1] = await ethers.getSigners();
    const initialRequiredApprovals = 1;

    const newRequiredApprovals = 2;

    const {validatorManager} = await deployBridge(owner.address, [v1.address], initialRequiredApprovals);

    await validatorManager.setRequiredApprovals(newRequiredApprovals);
    expect(await validatorManager.requiredApprovals()).to.equal(newRequiredApprovals);
  });

  it('Should deposit tokens to bridge', async function () {
    const [owner, v1] = await ethers.getSigners();
    const initialRequiredApprovals = 1;

    const depositAmount = '10000000000000000000';

    const { mockToken, bridge, chainId } = await deployBridge(owner.address, [v1.address], initialRequiredApprovals);

    await mockToken.approve(bridge.address, depositAmount);
    await bridge.deposit(
      mockToken.address,
      chainId,
      depositAmount
    );

    expect(await mockToken.balanceOf(bridge.address)).to.equal(depositAmount);
  });

  it('Should execute transfer', async function () {
    const [owner, v1, v2, v3, receiver] = await ethers.getSigners();
    const initialRequiredApprovals = 2;
    const depositAmount = '10000000000000000000';

    const {mockToken, bridge, chainId} = await deployBridge(owner.address, [v1.address, v2.address, v3.address], initialRequiredApprovals);

    const txHash = '0x54c96e7f79d5fd653951c49783fc2fa7299f14c01a5a3a03f8bfb55eecb2751f';

    await mockToken.approve(bridge.address, depositAmount);
    await bridge.deposit
      (
        mockToken.address,
        chainId,
        depositAmount
      );
    expect(
      await mockToken.balanceOf(bridge.address)
    ).to.equal(depositAmount);

    const id = utils.solidityKeccak256(
      ['bytes', 'address', 'address', 'uint256'],
      [txHash, mockToken.address, receiver.address, depositAmount]
    );

    const bridge1 = await ethers.getContractAt('Bridge', bridge.address, v1);

    await expect(
      bridge1.approveTransfer(txHash, mockToken.address, receiver.address, depositAmount)
    )
      .to.emit(bridge1, 'Approved')
      .withArgs(id, v1.address);

    const bridge2 = await ethers.getContractAt('Bridge', bridge.address, v2);
    await expect(
      bridge2.approveTransfer(txHash, mockToken.address, receiver.address, depositAmount)
    )
      .to.emit(bridge2, 'Approved')
      .withArgs(id, v2.address)
      .emit(bridge2, 'Transferred')
      .withArgs(mockToken.address, receiver.address, depositAmount, v2.address);

    const bridge3 = await ethers.getContractAt('Bridge', bridge.address, v3);
    await bridge3.approveTransfer(
      txHash,
      mockToken.address,
      receiver.address,
      depositAmount
    );

    await expect(
      bridge.approveTransfer(
        txHash,
        mockToken.address,
        receiver.address,
        depositAmount
      )
    ).to.be.revertedWith('only validator');

    expect(await bridge.executed(id)).to.equal(true);
  });

  it('Should deposit supported tokens to bridge', async function () {
    const [owner, v1] = await ethers.getSigners();
    const initialRequiredApprovals = 2;
    const depositAmount = '10000000000000000000';
    const mintAmount = '100000000000000000000';
    const {mockToken, bridge, chainId}= await deployBridge(owner.address, [v1.address], initialRequiredApprovals);

    const MockToken = await ethers.getContractFactory('MockToken');
    const mockToken2 = await MockToken.deploy('Token2', 'TKN2', mintAmount);
    await mockToken2.deployed();

    await mockToken.approve(bridge.address, depositAmount);
    await mockToken2.approve(bridge.address, depositAmount);

    await bridge.deposit(
      mockToken.address,
      chainId,
      depositAmount
    );

    await expect(
      bridge.deposit(mockToken2.address, chainId, depositAmount)
    ).to.be.revertedWith('Token is not supported');
  });
});
