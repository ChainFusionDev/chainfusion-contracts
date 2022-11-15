// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

interface IERC20BridgeMediator {
    function addToken(
        string memory symbol,
        uint256 chainId,
        address token
    ) external;

    function removeToken(string memory symbol, uint256 chainId) external;

    function mediate(
        uint256 sourceChain,
        uint256 destinationChain,
        bytes memory sourceData
    ) external view returns (bytes memory);
}
