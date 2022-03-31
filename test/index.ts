import { expect } from "chai";
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
});
