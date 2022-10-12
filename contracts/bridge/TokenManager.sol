// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./SignerOwnable.sol";
import "./ERC20Bridgeable.sol";

enum TokenType {
    DISABLED,
    PROVIDED,
    MINTED
}

contract TokenManager is Initializable, SignerOwnable {
    mapping(address => TokenType) public supportedTokens;

    function initialize(address _signerStorage) external initializer {
        _setSignerStorage(_signerStorage);
    }

    function setToken(address _token, TokenType _tokenType) external onlySigner {
        _setToken(_token, _tokenType);
    }

    function deployToken(string memory name, string memory symbol) external onlySigner returns (address) {
        ERC20Bridgeable erc20Bridgeable = new ERC20Bridgeable(name, symbol);
        _setToken(address(erc20Bridgeable), TokenType.MINTED);

        return address(erc20Bridgeable);
    }

    function getType(address _token) public view returns (TokenType) {
        return supportedTokens[_token];
    }

    function _setToken(address _token, TokenType _tokenType) private {
        supportedTokens[_token] = _tokenType;
    }
}
