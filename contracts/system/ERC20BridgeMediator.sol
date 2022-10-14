// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";

contract ERC20BridgeMediator is Ownable {
    mapping(uint256 => mapping(address => string)) public tokenToSymbol;
    mapping(uint256 => mapping(string => address)) public symbolToToken;

    event AddedToken(string symbol, uint256 chainId, address token);
    event RemovedToken(string symbol, uint256 chainId, address token);

    function mediate(
        uint256 sourceChain,
        uint256 destinationChain,
        bytes memory sourceData
    ) external view returns (bytes memory) {
        (address _sender, address _sourceToken, uint256 _chainId, address _receiver, uint256 _transferAmount) = abi
            .decode(sourceData, (address, address, uint256, address, uint256));

        string memory symbol = tokenToSymbol[sourceChain][_sourceToken];
        require(bytes(symbol).length > 0, "ERC20BridgeMediator: can't find token symbol");

        address destinationToken = symbolToToken[destinationChain][symbol];
        require(destinationToken != address(0), "ERC20BridgeMediator: can't find token by chain and symbol");

        bytes memory destinationData = abi.encode(_sender, destinationToken, _chainId, _receiver, _transferAmount);

        return destinationData;
    }

    function addToken(
        string memory symbol,
        uint256 chainId,
        address token
    ) public onlyOwner {
        require(bytes(tokenToSymbol[chainId][token]).length == 0, "ERC20BridgeMediator: symbol already added");
        require(symbolToToken[chainId][symbol] == address(0), "ERC20BridgeMediator: token already added");

        tokenToSymbol[chainId][token] = symbol;
        symbolToToken[chainId][symbol] = token;

        emit AddedToken(symbol, chainId, token);
    }

    function removeToken(string memory symbol, uint256 chainId) public onlyOwner {
        require(symbolToToken[chainId][symbol] != address(0), "ERC20BridgeMediator: token does not exist");

        address token = symbolToToken[chainId][symbol];
        delete tokenToSymbol[chainId][token];
        delete symbolToToken[chainId][symbol];

        emit RemovedToken(symbol, chainId, token);
    }
}
