import { expect } from 'chai';
import { ethers } from 'hardhat';
import { deployBridge } from '../utils/deploy';

describe('RelayBridge', function () {
  it('should send data', async function () {
    const [owner] = await ethers.getSigners();
    const chainId = 2;
    const data = ethers.utils.keccak256('0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF');

    const { relayBridge } = await deployBridge(owner.address);

    const hash = await relayBridge.dataHash(chainId, data);

    await expect(relayBridge.send(chainId, data)).to.emit(relayBridge, 'SentData').withArgs(hash);
    await expect(relayBridge.send(chainId, data)).to.be.revertedWith('RelayBridge: data already send');

    expect(await relayBridge.sendData(hash)).to.equals(data);
  });

  it('should transmit data', async function () {
    const [owner, user] = await ethers.getSigners();
    const fromChainId = 2;
    var abiCoder = ethers.utils.defaultAbiCoder;
    const data = abiCoder.encode(['string'], ['datafortransmit']);

    const { relayBridge, mockRelayBridgeApp } = await deployBridge(owner.address);

    const hash = await relayBridge.dataHash(fromChainId, data);

    expect(await relayBridge.validator()).to.equal(owner.address);

    await expect(relayBridge.transmit(mockRelayBridgeApp.address, fromChainId, data))
      .to.emit(relayBridge, 'TransmittedData')
      .withArgs(hash);
    await expect(relayBridge.transmit(mockRelayBridgeApp.address, fromChainId, data)).to.be.revertedWith(
      'RelayBridge: data already transmitted'
    );

    const bridgeUser = await ethers.getContractAt('RelayBridge', relayBridge.address, user);
    await expect(bridgeUser.transmit(mockRelayBridgeApp.address, fromChainId, data)).to.be.revertedWith(
      'RelayBridge: only validator'
    );

    expect(await relayBridge.transmitted(hash)).to.equals(true);
    expect(await mockRelayBridgeApp.canProcess()).to.equals(true);
    expect(await mockRelayBridgeApp.value()).to.equals('datafortransmit');
  });
});
