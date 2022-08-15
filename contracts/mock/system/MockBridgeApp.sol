// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "../../interfaces/IBridgeApp.sol";
import "../../bridge/RelayBridge.sol";

contract MockBridgeApp is IBridgeApp, Initializable {
    string public value;
    RelayBridge public relayBridge;

    event Reverted(bytes32 hash);

    function initialize(address _relayBridgeAddress) external initializer {
        relayBridge = RelayBridge(_relayBridgeAddress);
    }

    function send(
        uint256 destinationChain,
        uint256 gasLimit,
        bytes memory data
    ) public {
        relayBridge.send(destinationChain, gasLimit, data);
    }

    function execute(uint256, bytes memory data) public {
        value = abi.decode(data, (string));
    }

    function revertSend(uint256, bytes memory data) public {
        emit Reverted(keccak256(data));
    }

    function bridgeAppAddress() public pure returns (address) {
        return address(0);
    }
}
