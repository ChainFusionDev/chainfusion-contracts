import { ContractTransaction } from "ethers";
import { ethers } from "hardhat";

async function main() {
  const [owner] = await ethers.getSigners();
  const ownerAddress = owner.address;

  const TokenManager = await ethers.getContractFactory("TokenManager");
  const tokenManager = await TokenManager.deploy();
  await tokenManager.deployed();
  await (await tokenManager.initialize(ownerAddress)).wait();
  
  console.log("TokenManager deployed to:", tokenManager.address);

  const ValidatorManager = await ethers.getContractFactory("ValidatorManager");
  const validatorManager = await ValidatorManager.deploy();
  await validatorManager.deployed();
  
  console.log("ValidatorManager deployed to:", validatorManager.address);

  const Bridge = await ethers.getContractFactory("Bridge");
  const bridge = await Bridge.deploy();
  await bridge.deployed();
  await (await bridge.initialize(ownerAddress, validatorManager.address, tokenManager.address)).wait();

  console.log("Bridge deployed to:", bridge.address);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
