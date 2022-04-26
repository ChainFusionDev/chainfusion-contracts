import { ethers } from 'hardhat';

async function main() {
  const [owner] = await ethers.getSigners();
  const ownerAddress = owner.address;
  const feePercentage = '10000000000000000';

  const Globals = await ethers.getContractFactory('Globals');
  const globals = await Globals.deploy();
  await globals.deployed();
  console.log('Globals deployed to:', globals.address);

  const TokenManager = await ethers.getContractFactory('TokenManager');
  const tokenManager = await TokenManager.deploy();
  await tokenManager.deployed();
  await (await tokenManager.initialize(ownerAddress)).wait();

  console.log('TokenManager deployed to:', tokenManager.address);

  const LiquidityPools = await ethers.getContractFactory('LiquidityPools');
  const liquidityPools = await LiquidityPools.deploy();
  await liquidityPools.deployed();

  console.log('LiquidityPools deployed to:', liquidityPools.address);

  const ValidatorManager = await ethers.getContractFactory('ValidatorManager');
  const validatorManager = await ValidatorManager.deploy();
  await validatorManager.deployed();

  console.log('ValidatorManager deployed to:', validatorManager.address);

  const Bridge = await ethers.getContractFactory('Bridge');
  const bridge = await Bridge.deploy();
  await bridge.deployed();
  await (
    await bridge.initialize(
      ownerAddress,
      validatorManager.address,
      tokenManager.address,
      liquidityPools.address,
      globals.address
    )
  ).wait();

  await (await liquidityPools.initialize(tokenManager.address, bridge.address, globals.address, feePercentage)).wait();

  console.log('Bridge deployed to:', bridge.address);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
