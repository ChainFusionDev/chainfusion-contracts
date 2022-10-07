// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./SignerOwnable.sol";

enum TokenType {
    DISABLED,
    PROVIDED,
    MINTED
}

struct TokenInfo {
    address token;
    TokenType tokenType;
}

contract TokenManager is Initializable, SignerOwnable {
    mapping(address => TokenInfo) public supportedTokens;

    function initialize(address _signerStorage) external initializer {
        _setSignerStorage(_signerStorage);
    }

    function setToken(address _token, TokenType tokenType) external onlySigner {
        supportedTokens[_token].tokenType = tokenType;
    }

    function getType(address _token) public view returns (TokenType) {
        return supportedTokens[_token].tokenType;
    }
}
