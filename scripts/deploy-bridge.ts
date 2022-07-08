import { ethers } from 'hardhat';

async function main() {
  const [owner] = await ethers.getSigners();
  const ownerAddress = owner.address;
  const feePercentage = '10000000000000000';
  const validatorAddress = '0x70997970c51812dc3a010c7d01b50e0d17dc79c8';
  const foundationAddress = '0x3c44cdddb6a900fa2b585dd299e03d12fa4293bc';
  const validatorRefundFee = '10000000000000000';

  const TokenManager = await ethers.getContractFactory('TokenManager');
  const tokenManager = await TokenManager.deploy();
  await tokenManager.deployed();
  await (await tokenManager.initialize(ownerAddress)).wait();

  console.log('TokenManager deployed to:', tokenManager.address);

  const LiquidityPools = await ethers.getContractFactory('LiquidityPools');
  const liquidityPools = await LiquidityPools.deploy();
  await liquidityPools.deployed();

  console.log('LiquidityPools deployed to:', liquidityPools.address);

  const FeeManager = await ethers.getContractFactory('FeeManager');
  const feeManager = await FeeManager.deploy();
  await feeManager.deployed();
  await (
    await feeManager.initialize(liquidityPools.address, validatorAddress, foundationAddress, validatorRefundFee)
  ).wait();

  console.log('FeeManager deployed to:', feeManager.address);

  const Bridge = await ethers.getContractFactory('Bridge');
  const bridge = await Bridge.deploy();
  await bridge.deployed();
  await (
    await bridge.initialize(
      ownerAddress,
      validatorAddress,
      tokenManager.address,
      liquidityPools.address,
      feeManager.address
    )
  ).wait();

  await (
    await liquidityPools.initialize(tokenManager.address, bridge.address, feeManager.address, feePercentage)
  ).wait();

  console.log('Bridge deployed to:', bridge.address);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
