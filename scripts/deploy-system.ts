import { ethers } from 'hardhat';
import hre from 'hardhat';
import { deploySystemContracts } from './deploy/system';

const VERIFY = (process.env.VERIFY || '').trim().toLowerCase();
const VALIDATOR_KEYS =
  process.env.VALIDATOR_KEYS === undefined ? [] : (process.env.VALIDATOR_KEYS || '').trim().split(',');

async function main() {
  const deployment = await deploySystemContracts({}, true);

  if (VALIDATOR_KEYS.length > 0) {
    console.log('\nStaking\n');

    for (const privateKey of VALIDATOR_KEYS) {
      const signer = new ethers.Wallet(privateKey, ethers.provider);
      const signerStaking = await ethers.getContractAt('Staking', deployment.staking.address, signer);
      console.log('Staking', ethers.utils.formatEther(deployment.minimalStake), 'CFN for:', signer.address);
      await (await signerStaking.stake({ value: deployment.minimalStake })).wait();
    }
  }

  if (VERIFY === 'true') {
    console.log('\nVerifying contracts\n');

    // Sometimes fails, false positively, likely it's a bug in verify plugin or blockscout
    await ignoreError(hre.run('verify:verify', { address: deployment.addressStorage.address }));
    await ignoreError(hre.run('verify:verify', { address: deployment.staking.address }));
    await ignoreError(hre.run('verify:verify', { address: deployment.dkg.address }));
    await ignoreError(hre.run('verify:verify', { address: deployment.slashingVoting.address }));
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
