import { ethers } from 'hardhat';
import hre from 'hardhat';

const withdrawalPeriod = 60;
const minimalStake = ethers.utils.parseEther('100');
const epochPeriod = 1000;
const slashingThresold = 10;
const slashingEpochs = 3;
const deadlinePeriod = 20;

const VERIFY = (process.env.VERIFY || '').trim().toLowerCase();
const VALIDATOR_KEYS =
  process.env.VALIDATOR_KEYS === undefined ? [] : (process.env.VALIDATOR_KEYS || '').trim().split(',');

async function main() {
  console.log('\nDeploying contracts\n');

  const ContractRegistry = await ethers.getContractFactory('ContractRegistry');
  const contractRegistry = await ContractRegistry.deploy();
  await contractRegistry.deployed();

  console.log('ContractRegistry deployed to:', contractRegistry.address);

  const EventRegistry = await ethers.getContractFactory('EventRegistry');
  const eventRegistry = await EventRegistry.deploy();
  await eventRegistry.deployed();

  console.log('EventRegistry deployed to:', eventRegistry.address);

  const AddressStorage = await ethers.getContractFactory('AddressStorage');
  const addressStorage = await AddressStorage.deploy();
  await addressStorage.deployed();

  console.log('AddressStorage deployed to:', addressStorage.address);

  const Staking = await ethers.getContractFactory('Staking');
  const staking = await Staking.deploy();
  await staking.deployed();

  console.log('Staking deployed to:', staking.address);

  const DKG = await ethers.getContractFactory('DKG');
  const dkg = await DKG.deploy();
  await dkg.deployed();

  console.log('DKG deployed to:', dkg.address);

  const SupportedTokens = await ethers.getContractFactory('SupportedTokens');
  const supportedTokens = await SupportedTokens.deploy();
  await supportedTokens.deployed();

  console.log('SupportedTokens deployed to:', supportedTokens.address);

  const SlashingVoting = await ethers.getContractFactory('SlashingVoting');
  const slashingVoting = await SlashingVoting.deploy();
  await slashingVoting.deployed();

  console.log('SlashingVoting deployed to:', slashingVoting.address);

  console.log('Initializing contracts...');

  await (await addressStorage.initialize([])).wait();
  await addressStorage.transferOwnership(staking.address);
  await (await dkg.initialize(contractRegistry.address, deadlinePeriod)).wait();
  await (await contractRegistry.initialize(dkg.address)).wait();
  await (await eventRegistry.initialize(dkg.address)).wait();
  await (await supportedTokens.initialize(dkg.address)).wait();

  await (
    await staking.initialize(
      dkg.address,
      minimalStake,
      withdrawalPeriod,
      contractRegistry.address,
      addressStorage.address
    )
  ).wait();

  await (
    await slashingVoting.initialize(
      dkg.address,
      epochPeriod,
      slashingThresold,
      slashingEpochs,
      contractRegistry.address
    )
  ).wait();

  await contractRegistry.setContract(await slashingVoting.SLASHING_VOTING_KEY(), slashingVoting.address);
  await contractRegistry.setContract(await staking.STAKING_KEY(), staking.address);
  await contractRegistry.setContract(await dkg.DKG_KEY(), dkg.address);
  await contractRegistry.setContract(await supportedTokens.SUPPORTED_TOKENS_KEY(), supportedTokens.address);

  console.log('Initialized contracts');

  if (VALIDATOR_KEYS.length > 0) {
    console.log('\nStaking\n');

    for (const privateKey of VALIDATOR_KEYS) {
      const signer = new ethers.Wallet(privateKey, ethers.provider);
      const signerStaking = await ethers.getContractAt('Staking', staking.address, signer);
      console.log('Staking', ethers.utils.formatEther(minimalStake), 'CFN for:', signer.address);
      await (await signerStaking.stake({ value: minimalStake })).wait();
    }
  }

  if (VERIFY === 'true') {
    console.log('\nVerifying contracts\n');

    // Sometimes fails, false positively, likely it's a bug in verify plugin or blockscout
    await ignoreError(hre.run('verify:verify', { address: addressStorage.address }));
    await ignoreError(hre.run('verify:verify', { address: staking.address }));
    await ignoreError(hre.run('verify:verify', { address: dkg.address }));
    await ignoreError(hre.run('verify:verify', { address: slashingVoting.address }));
  }
}

async function ignoreError(callback: Promise<any>): Promise<any> {
  try {
    await callback;
  } catch {
    return;
  }
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
