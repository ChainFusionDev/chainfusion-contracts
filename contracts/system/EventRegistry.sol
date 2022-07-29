// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./ValidatorOwnable.sol";

contract EventRegistry is ValidatorOwnable, Initializable {
    mapping(bytes32 => bool) public registeredEvents;

    event EventRegistered(bytes32 _hash, uint256 _sourceChain, uint256 _destinationChain);

    function initialize(address _signerGetterAddress) external initializer {
        _setSignerGetter(_signerGetterAddress);
    }

    function registerEvent(
        bytes32 _hash,
        uint256 _sourceChain,
        uint256 _destinationChain
    ) external onlyCurrentValidator {
        bytes32 key = this.eventKey(_hash, _sourceChain, _destinationChain);
        require(registeredEvents[key] == false, "EventRegistry: event is already registered");

        registeredEvents[key] = true;
        emit EventRegistered(_hash, _sourceChain, _destinationChain);
    }

    function eventKey(
        bytes32 _hash,
        uint256 _sourceChain,
        uint256 _destinationChain
    ) external pure returns (bytes32) {
        return keccak256(abi.encodePacked(_hash, _sourceChain, _destinationChain));
    }
}
