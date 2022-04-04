import { expect } from "chai";
import { assert } from "console";
import { ethers } from "hardhat";

describe("Bridge", function () {
  it("Should change required approvals", async function () {
    const [owner] = await ethers.getSigners();
    const ownerAddress = owner.address;
    const validators = [ownerAddress];

    const initialRequiredApprovals = 1;
    const newRequiredApprovals = 2;

    const Bridge = await ethers.getContractFactory("Bridge");
    const bridge = await Bridge.deploy();
    await bridge.deployed();
    await (await bridge.initialize(ownerAddress, validators, initialRequiredApprovals)).wait();
    expect(await bridge.requiredApprovals()).to.equal(initialRequiredApprovals);

    await (await bridge.setRequiredApprovals(newRequiredApprovals)).wait();
    expect(await bridge.requiredApprovals()).to.equal(newRequiredApprovals);
  });
    
  it("Should deposit tokens to bridge", async function () {
    const chainId = 123;
    const mintAmount = "100000000000000000000";
    const depositAmount = "10000000000000000000";

    const Bridge = await ethers.getContractFactory("Bridge");
    const bridge = await Bridge.deploy();
    await bridge.deployed();

    const MockToken = await ethers.getContractFactory("MockToken");
    const mockToken = await MockToken.deploy("Token", "TKN", mintAmount);
    await mockToken.deployed();

    await mockToken.approve(bridge.address, depositAmount);
    await bridge.deposit(mockToken.address, chainId, depositAmount);
    expect(await mockToken.balanceOf(bridge.address)).to.equal(depositAmount);
  });
    
  it("Should execute transfer", async function () {
    const [owner, v1, v2, v3, receiver] = await ethers.getSigners();
    const chainId = 123;
    const mintAmount = "100000000000000000000";
    const depositAmount = "10000000000000000000";
    
    const txHash1 = "0x54c96e7f79d5fd653951c49783fc2fa7299f14c01a5a3a03f8bfb55eecb2751f";
    const txHash2 = "0x54c96e7f79d5fd653951c49343fc2fa7299f14c01a5a3a03f8bfb55eecb2711f";

    const Bridge = await ethers.getContractFactory("Bridge");
    const bridge = await Bridge.deploy();
    await bridge.deployed();
    await (await bridge.initialize(owner.address, [v1.address, v2.address, v3.address], 2)).wait();

    const MockToken = await ethers.getContractFactory("MockToken");
    const mockToken = await MockToken.deploy("Token", "TKN", mintAmount);
    await mockToken.deployed();

    await mockToken.approve(bridge.address, depositAmount);
    await bridge.deposit(mockToken.address, chainId, depositAmount);
    expect(await mockToken.balanceOf(bridge.address)).to.equal(depositAmount);

    const bridge1 = await ethers.getContractAt("Bridge", bridge.address, v1);
    await bridge1.approveTransfer(txHash1, mockToken.address, receiver.address, depositAmount);

    const bridge2 = await ethers.getContractAt("Bridge", bridge.address, v2);
    await bridge2.approveTransfer(txHash1, mockToken.address, receiver.address, depositAmount);

    const bridge3 = await ethers.getContractAt("Bridge", bridge.address, v3);    
    await bridge3.approveTransfer(txHash2, mockToken.address, receiver.address, depositAmount);

    try {
      const bridgeOther = await ethers.getContractAt("Bridge", bridge.address, owner);
      await bridgeOther.approveTransfer(txHash2, mockToken.address, receiver.address, depositAmount);
    } catch (ex) {
      console.log("Only Validator");
    }   
  });
});
