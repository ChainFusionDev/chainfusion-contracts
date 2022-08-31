import { expect } from 'chai';
import { ethers } from 'hardhat';
import { deployBridge, deployBridgeWithMocks } from '../utils/deploy';

describe('RelayBridge', function () {
  it('should send data', async function () {
    const [appContract] = await ethers.getSigners();

    const sourceChain = ethers.provider.network.chainId;
    const destinationChain = 1;
    const gasLimit = 1;

    const abiCoder = ethers.utils.defaultAbiCoder;
    const data = abiCoder.encode(['string'], ['dataforsend']);

    const { relayBridge } = await deployBridge();

    const hash = await relayBridge.dataHash(appContract.address, sourceChain, destinationChain, gasLimit, data);

    await expect(relayBridge.send(destinationChain, gasLimit, data)).to.emit(relayBridge, 'Sent').withArgs(hash);
    await expect(relayBridge.send(destinationChain, gasLimit, data)).to.be.revertedWith(
      'RelayBridge: data already send'
    );

    expect(await relayBridge.sentData(hash)).to.equals(data);
  });

  it('should emit event Reverted in appContract', async function () {
    const destinationChain = 1;
    const gasLimit = 1;

    const abiCoder = ethers.utils.defaultAbiCoder;
    const data = abiCoder.encode(['string'], ['dataforsend']);

    const { relayBridge, mockBridgeApp } = await deployBridgeWithMocks();

    await mockBridgeApp.send(destinationChain, gasLimit, data);

    const hash = ethers.utils.keccak256(data);

    await expect(relayBridge.revertSend(mockBridgeApp.address, destinationChain, gasLimit, data))
      .to.emit(mockBridgeApp, 'Reverted')
      .withArgs(hash);
  });

  it('should execute data', async function () {
    const sourceChain = 1;
    const destinationChain = ethers.provider.network.chainId;
    const gasLimit = 1;

    const abiCoder = ethers.utils.defaultAbiCoder;
    const data = abiCoder.encode(['string'], ['dataforsend']);

    const { relayBridge, mockBridgeApp } = await deployBridgeWithMocks();

    const hash = await relayBridge.dataHash(mockBridgeApp.address, sourceChain, destinationChain, gasLimit, data);

    await expect(relayBridge.execute(mockBridgeApp.address, sourceChain, gasLimit, data))
      .to.emit(relayBridge, 'Executed')

      .withArgs(hash);
    await expect(relayBridge.execute(mockBridgeApp.address, sourceChain, gasLimit, data)).to.be.revertedWith(
      'RelayBridge: data already executed'
    );

    expect(await relayBridge.executed(hash)).to.equals(true);
    expect(await mockBridgeApp.value()).to.equals('dataforsend');

    expect(await mockBridgeApp.bridgeAppAddress()).to.equals('0x0000000000000000000000000000000000000000');
  });
});
