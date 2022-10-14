// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

interface IBridgeApp {
    function execute(uint256 sourceChain, bytes memory data) external;

    function revertSend(uint256 destinationChain, bytes memory data) external;

    // address of BridgeApp in CFN blockchain
    function bridgeAppAddress() external view returns (address);
}
