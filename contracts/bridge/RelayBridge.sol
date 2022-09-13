// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./SignerOwnable.sol";
import "./TokenManager.sol";
import "./Bridge.sol";
import "./FeeManager.sol";
import "./Globals.sol";
import "../interfaces/IBridgeApp.sol";
import "../interfaces/IBridgeMediator.sol";

contract RelayBridge is Initializable, SignerOwnable {
    mapping(bytes32 => bytes) public sentData;

    mapping(bytes32 => bool) public sent;
    mapping(bytes32 => bool) public executed;
    mapping(bytes32 => bool) public reverted;

    address[] public leaderHistory;

    uint256 public nonce;

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
        bytes32 hash = dataHash(msg.sender, block.chainid, destinationChain, gasLimit, data, nonce);
        require(sentData[hash].length == 0, "RelayBridge: data already send");

        sent[hash] = true;
        sentData[hash] = data;
        nonce++;

        emit Sent(hash, block.chainid, destinationChain);
    }

    function revertSend(
        address appContract,
        uint256 destinationChain,
        uint256 gasLimit,
        bytes memory data,
        uint256 _nonce,
        address leader
    ) external onlySigner {
        bytes32 hash = dataHash(appContract, block.chainid, destinationChain, gasLimit, data, _nonce);
        require(sent[hash], "RelayBridge: data never sent");
        require(!reverted[hash], "RelayBridge: data already reverted");

        reverted[hash] = true;
        leaderHistory.push(leader);

        IBridgeApp(appContract).revertSend(destinationChain, data);

        emit Reverted(hash, block.chainid, destinationChain);
    }

    function execute(
        address appContract,
        uint256 sourceChain,
        uint256 gasLimit,
        bytes memory data,
        uint256 _nonce,
        address leader
    ) external onlySigner {
        bytes32 hash = dataHash(appContract, sourceChain, block.chainid, gasLimit, data, _nonce);
        require(!executed[hash], "RelayBridge: data already executed");

        executed[hash] = true;
        leaderHistory.push(leader);

        IBridgeApp(appContract).execute(sourceChain, data);

        emit Executed(hash, sourceChain, block.chainid);
    }

    function leaderHistoryLength() external view returns (uint256) {
        return leaderHistory.length;
    }

    function dataHash(
        address appContract,
        uint256 sourceChain,
        uint256 destinationChain,
        uint256 gasLimit,
        bytes memory data,
        uint256 _nonce
    ) public pure returns (bytes32) {
        return keccak256(abi.encode(appContract, sourceChain, destinationChain, gasLimit, data, _nonce));
    }
}
