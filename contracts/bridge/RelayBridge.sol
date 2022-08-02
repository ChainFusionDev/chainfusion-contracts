// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./SignerOwnable.sol";
import "./TokenManager.sol";
import "./Bridge.sol";
import "./FeeManager.sol";
import "./Globals.sol";
import "hardhat/console.sol";
import "../interfaces/IRelayBridgeApp.sol";

contract RelayBridge is Initializable, SignerOwnable {
    mapping(bytes32 => bytes) public sendData;
    mapping(bytes32 => bool) public transmitted;

    event SentData(bytes32 hash);
    event TransmittedData(bytes32 hash);

    function initialize(address _signerStorage) external initializer {
        _setSignerStorage(_signerStorage);
    }

    function send(uint256 chainId, bytes memory data) external {
        bytes32 hash = dataHash(chainId, data);
        require(sendData[hash].length == 0, "RelayBridge: data already send");

        sendData[hash] = data;

        emit SentData(hash);
    }

    function transmit(
        address appContract,
        uint256 fromChainId,
        bytes memory data
    ) external onlySigner {
        bytes32 hash = dataHash(fromChainId, data);
        require(!transmitted[hash], "RelayBridge: data already transmitted");

        IRelayBridgeApp(appContract).process(fromChainId, data);

        transmitted[hash] = true;

        emit TransmittedData(hash);
    }

    function revertSend(
        address appContract,
        uint256 destinationChainId,
        bytes memory data
    ) public {
        IRelayBridgeApp(appContract).revertSend(destinationChainId, data);
    }

    function dataHash(uint256 chainId, bytes memory data) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(chainId, data));
    }
}
