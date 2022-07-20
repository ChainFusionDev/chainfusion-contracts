// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../../interfaces/IRelayBridgeApp.sol";

contract MockRelayBridgeApp is IRelayBridgeApp {
    string public value;

    event Reverted(bytes32 hash);

    function canProcess() external pure returns (bool) {
        return true;
    }

    function process(uint256, bytes memory data) public {
        value = abi.decode(data, (string));
    }

    function revertSend(uint256, bytes memory data) public {
        emit Reverted(keccak256(data));
    }
}
