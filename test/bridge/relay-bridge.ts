import { expect } from 'chai';
import { ethers } from 'hardhat';
import { deployBridge, deployBridgeWithMocks } from '../utils/deploy';

describe('RelayBridge', function () {
  it('should send data', async function () {
    const [validator] = await ethers.getSigners();

    const chainId = 2;
    const data = ethers.utils.keccak256('0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF');

    const { relayBridge } = await deployBridge();

    const hash = await relayBridge.dataHash(chainId, data);

    await expect(relayBridge.send(chainId, data))
      .to.emit(relayBridge, 'SentData')
      .withArgs(hash, validator.getChainId, chainId);
    await expect(relayBridge.send(chainId, data)).to.be.revertedWith('RelayBridge: data already send');

    expect(await relayBridge.sendData(hash)).to.equals(data);
  });

  it('should transmit data', async function () {
    const fromChainId = 2;
    const abiCoder = ethers.utils.defaultAbiCoder;
    const data = abiCoder.encode(['string'], ['datafortransmit']);

    const { relayBridge, mockRelayBridgeApp } = await deployBridgeWithMocks();

    const hash = await relayBridge.dataHash(fromChainId, data);

    await expect(relayBridge.transmit(mockRelayBridgeApp.address, fromChainId, data))
      .to.emit(relayBridge, 'TransmittedData')
      .withArgs(hash);
    await expect(relayBridge.transmit(mockRelayBridgeApp.address, fromChainId, data)).to.be.revertedWith(
      'RelayBridge: data already transmitted'
    );

    expect(await relayBridge.transmitted(hash)).to.equals(true);
    expect(await mockRelayBridgeApp.canProcess()).to.equals(true);
    expect(await mockRelayBridgeApp.value()).to.equals('datafortransmit');
  });

  it('should emit event Reverted in appContract', async function () {
    const destinationChainId = 1;
    const abiCoder = ethers.utils.defaultAbiCoder;
    const data = abiCoder.encode(['string'], ['dataforrevertsend']);
    const hash = ethers.utils.keccak256(data);
    const { relayBridge, mockRelayBridgeApp } = await deployBridgeWithMocks();

    await expect(relayBridge.revertSend(mockRelayBridgeApp.address, destinationChainId, data))
      .to.emit(mockRelayBridgeApp, 'Reverted')
      .withArgs(hash);
  });
});
