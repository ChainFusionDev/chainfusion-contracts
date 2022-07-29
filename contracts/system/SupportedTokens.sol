// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./ValidatorOwnable.sol";
import "./ContractKeys.sol";

contract SupportedTokens is ContractKeys, ValidatorOwnable, Initializable {
    enum TokenType {
        PROVIDED,
        MINTED
    }

    struct TokenInfo {
        address token;
        TokenType tokenType;
    }

    mapping(string => mapping(uint256 => TokenInfo)) public tokens;

    event AddedToken(string symbol, uint256 chainId, address token, TokenType tokenType);
    event RemovedToken(string symbol, uint256 chainId, address token);

    function initialize(address _signerGetterAddress) external initializer {
        _setSignerGetter(_signerGetterAddress);
    }

    function addToken(
        string memory symbol,
        uint256 chainId,
        address token,
        TokenType tokenType
    ) public onlyValidator {
        require(tokens[symbol][chainId].token == address(0), "SupportedTokens: token already added");
        tokens[symbol][chainId].token = token;
        tokens[symbol][chainId].tokenType = tokenType;

        emit AddedToken(symbol, chainId, token, tokenType);
    }

    function removeToken(string memory symbol, uint256 chainId) public onlyValidator {
        require(tokens[symbol][chainId].token != address(0), "SupportedTokens: token does not exist");
        address token = tokens[symbol][chainId].token;
        delete tokens[symbol][chainId];

        emit RemovedToken(symbol, chainId, token);
    }
}
