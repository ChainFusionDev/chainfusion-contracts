// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract TokenManager is Initializable, Ownable {
    struct TokenInfo {
        mapping(uint256 => address) chainToToken;
        bool isEnabled;
        bool isMintable;
    }

    mapping(address => TokenInfo) public supportedTokens;

    function initialize(address _owner) external initializer {
        _transferOwnership(_owner);
    }

    function setDestinationToken(
        uint256 _chainId,
        address _token,
        address _destinationToken
    ) external onlyOwner {
        supportedTokens[_token].chainToToken[_chainId] = _destinationToken;
    }

    function setMintable(address _token, bool _isMintable) public {
        supportedTokens[_token].isMintable = _isMintable;
    }

    function setEnabled(address _token, bool _isEnabled) public {
        supportedTokens[_token].isEnabled = _isEnabled;
    }

    function isTokenEnabled(address _token) public view returns (bool) {
        return supportedTokens[_token].isEnabled;
    }

    function isTokenMintable(address _token) public view returns (bool) {
        return supportedTokens[_token].isMintable;
    }

    function getDestinationToken(address _token, uint256 _chainId) public view returns (address) {
        return supportedTokens[_token].chainToToken[_chainId];
    }
}
