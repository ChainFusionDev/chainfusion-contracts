import { ContractTransaction } from "ethers";
import { ethers } from "hardhat";

async function main() {
  const [owner] = await ethers.getSigners();
  const ownerAddress = owner.address;

  const Bridge = await ethers.getContractFactory("Bridge");
  const bridge = await Bridge.deploy();
  await bridge.deployed();
  await (await bridge.initialize(ownerAddress, [ownerAddress], 1)).wait();

  console.log("Bridge deployed to:", bridge.address);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
