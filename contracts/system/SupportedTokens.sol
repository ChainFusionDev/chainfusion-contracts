// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./ContractKeys.sol";

contract SupportedTokens is ContractKeys, Ownable {
    mapping(string => mapping(uint256 => address)) public tokens;

    event AddedToken(string symbol, uint256 chainId, address token);
    event RemovedToken(string symbol, uint256 chainId, address token);

    function addToken(
        string memory symbol,
        uint256 chainId,
        address token
    ) public onlyOwner {
        require(tokens[symbol][chainId] == address(0), "SupportedTokens: token already added");
        tokens[symbol][chainId] = token;

        emit AddedToken(symbol, chainId, token);
    }

    function removeToken(string memory symbol, uint256 chainId) public onlyOwner {
        require(tokens[symbol][chainId] != address(0), "SupportedTokens: token does not exist");
        address token = tokens[symbol][chainId];
        delete tokens[symbol][chainId];

        emit RemovedToken(symbol, chainId, token);
    }
}
