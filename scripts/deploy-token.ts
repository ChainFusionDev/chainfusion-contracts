import { ethers } from 'hardhat';

async function main() {
  const mintAmount = '100000000000000000000';

  const MockToken = await ethers.getContractFactory('MockToken');
  const mockToken = await MockToken.deploy('Tether', 'USDT', mintAmount);
  await mockToken.deployed();

  console.log('MockToken deployed to:', mockToken.address);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
