// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract TokenManager is Initializable, Ownable {
    struct TokenInfo {
        mapping(uint256 => address) chainToToken;
        bool isSupported;
    }

    mapping(address => TokenInfo) public supportedTokens;

    function initialize(address _owner) external initializer {
        _transferOwnership(_owner);
    }

    function addSupportedToken(
        uint256 _chainId,
        address _token,
        address _destinationToken
    ) external onlyOwner {
        supportedTokens[_token].chainToToken[_chainId] = _destinationToken;
        supportedTokens[_token].isSupported = true;
    }

    function isTokenSupported(address _token) public view returns (bool) {
        return supportedTokens[_token].isSupported;
    }

    function getDestinationToken(address _token, uint256 _chainId) public view returns (address) {
        return supportedTokens[_token].chainToToken[_chainId];
    }
}
