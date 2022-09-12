// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

interface IBridgeMediator {
    function mediate(
        uint256 sourceChain,
        uint256 destinationChain,
        bytes memory data
    ) external view returns (bytes memory);
}
