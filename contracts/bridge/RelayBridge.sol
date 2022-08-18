// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./SignerOwnable.sol";
import "./TokenManager.sol";
import "./Bridge.sol";
import "./FeeManager.sol";
import "./Globals.sol";
import "hardhat/console.sol";
import "../interfaces/IBridgeApp.sol";
import "hardhat/console.sol";

contract RelayBridge is Initializable, SignerOwnable {
    mapping(bytes32 => bytes) public sentData;

    mapping(bytes32 => bool) public sent;
    mapping(bytes32 => bool) public executed;
    mapping(bytes32 => bool) public reverted;

    event Sent(bytes32 hash, uint256 sourceChain, uint256 destinationChain);
    event Reverted(bytes32 hash, uint256 sourceChain, uint256 destinationChain);
    event Executed(bytes32 hash, uint256 sourceChain, uint256 destinationChain);

    function initialize(address _signerStorage) external initializer {
        _setSignerStorage(_signerStorage);
    }

    function send(
        uint256 destinationChain,
        uint256 gasLimit,
        bytes memory data
    ) external {
        bytes32 hash = dataHash(msg.sender, block.chainid, destinationChain, gasLimit, data);
        require(sentData[hash].length == 0, "RelayBridge: data already send");

        sent[hash] = true;
        sentData[hash] = data;

        emit Sent(hash, block.chainid, destinationChain);
    }

    function revertSend(
        address appContract,
        uint256 destinationChain,
        uint256 gasLimit,
        bytes memory data
    ) external onlySigner {
        bytes32 hash = dataHash(appContract, block.chainid, destinationChain, gasLimit, data);
        require(sent[hash], "RelayBridge: data never sent");
        require(!reverted[hash], "RelayBridge: data already reverted");

        reverted[hash] = true;

        IBridgeApp(appContract).revertSend(destinationChain, data);

        emit Reverted(hash, block.chainid, destinationChain);
    }

    function execute(
        address appContract,
        uint256 sourceChain,
        uint256 destinationChain,
        uint256 gasLimit,
        bytes memory data
    ) external onlySigner {
        bytes32 hash = dataHash(appContract, sourceChain, destinationChain, gasLimit, data);
        require(!executed[hash], "RelayBridge: data already executed");

        IBridgeApp(appContract).execute(sourceChain, data);

        executed[hash] = true;

        emit Executed(hash, sourceChain, destinationChain);
    }

    function dataHash(
        address appContract,
        uint256 sourceChain,
        uint256 destinationChain,
        uint256 gasLimit,
        bytes memory data
    ) public pure returns (bytes32) {
        return keccak256(abi.encode(appContract, sourceChain, destinationChain, gasLimit, data));
    }
}
