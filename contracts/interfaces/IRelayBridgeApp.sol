// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

interface IRelayBridgeApp {
    function canProcess() external returns (bool);

    function process(uint256 sourceChainId, bytes memory data) external;

    function revertSend(uint256 destinationChainId, bytes memory data) external;
}
