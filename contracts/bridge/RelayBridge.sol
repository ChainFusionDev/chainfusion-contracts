// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./TokenManager.sol";
import "./Bridge.sol";
import "./FeeManager.sol";
import "./Globals.sol";
import "hardhat/console.sol";
import "../interfaces/IRelayBridgeApp.sol";

contract RelayBridge is Initializable, Ownable {
    mapping(bytes32 => bytes) public sendData;
    mapping(bytes32 => bool) public transmitted;

    address public validator;

    event SentData(bytes32 hash);
    event TransmittedData(bytes32 hash);

    modifier onlyValidator() {
        require(validator == msg.sender, "RelayBridge: only validator");
        _;
    }

    function initialize(address _validator) external initializer {
        setValidator(_validator);
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
    ) external onlyValidator {
        bytes32 hash = dataHash(fromChainId, data);
        require(!transmitted[hash], "RelayBridge: data already transmitted");

        IRelayBridgeApp(appContract).process(fromChainId, data);

        transmitted[hash] = true;

        emit TransmittedData(hash);
    }

    function setValidator(address _validator) public onlyOwner {
        validator = _validator;
    }

    function dataHash(uint256 chainId, bytes memory data) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(chainId, data));
    }
}
