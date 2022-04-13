// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract TokenManager is Initializable, Ownable {
    mapping(uint256 => mapping(address => address)) public supportedTokens;

    function initialize(address _owner) external initializer {
        _transferOwnership(_owner);
    }

    function addSupportedToken(
        uint256 _chainId,
        address _token,
        address _destinationToken
    ) external onlyOwner {
        supportedTokens[_chainId][_token] = _destinationToken;
    }
}
