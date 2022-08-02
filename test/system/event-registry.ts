import { ethers } from 'hardhat';
import { expect } from 'chai';
import { deploySystem } from '../utils/deploy';

describe('EventRegistry', function () {
  it('should register events', async function () {
    const eventHash1 = ethers.utils.keccak256([1, 1, 1]);
    const eventHash2 = ethers.utils.keccak256([2, 2, 2]);
    const eventHash3 = ethers.utils.keccak256([3, 3, 3]);

    const { staking, eventRegistry, minimalStake } = await deploySystem();

    const [, v1, v2, other] = await ethers.getSigners();

    const staking1 = await ethers.getContractAt('Staking', staking.address, v1);
    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);

    const eventRegistry1 = await ethers.getContractAt('EventRegistry', eventRegistry.address, v1);
    const eventRegistry2 = await ethers.getContractAt('EventRegistry', eventRegistry.address, v2);
    const eventRegistryOther = await ethers.getContractAt('EventRegistry', eventRegistry.address, other);

    await staking1.stake({ value: minimalStake });
    await staking2.stake({ value: minimalStake });

    await eventRegistry1.registerEvent(eventHash1, 1, 1, 2);
    await expect(eventRegistry1.registerEvent(eventHash1, 1, 1, 2)).to.be.revertedWith(
      'EventRegistry: event is already registered'
    );
    expect(await eventRegistry1.registeredEvents(await eventRegistry1.eventKey(eventHash1, 1, 1, 2))).to.equal(true);

    await eventRegistry1.registerEvent(eventHash1, 1, 3, 4);

    await eventRegistry2.registerEvent(eventHash2, 1, 3, 4);
    await expect(eventRegistry2.registerEvent(eventHash2, 1, 3, 4)).to.be.revertedWith(
      'EventRegistry: event is already registered'
    );

    await expect(eventRegistryOther.registerEvent(eventHash3, 1, 5, 6)).to.be.revertedWith(
      'ValidatorOwnable: only validator'
    );
    expect(await eventRegistry1.registeredEvents(await eventRegistry1.eventKey(eventHash3, 1, 5, 6))).to.equal(false);
  });
});
