// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./ValidatorOwnable.sol";
import "./ContractKeys.sol";

contract EventRegistry is ValidatorOwnable, Initializable, ContractKeys {
    mapping(bytes32 => bool) public registeredEvents;

    event EventRegistered(
        bytes32 _hash,
        uint256 _generation,
        address _appContract,
        uint256 _sourceChain,
        uint256 _destinationChain,
        bytes _data,
        uint256 _gasLimit,
        uint256 _nonce
    );

    function initialize(address _validatorGetterAddress) external initializer {
        _setValidatorGetter(_validatorGetterAddress);
    }

    function registerEvent(
        bytes32 _hash,
        uint256 _generation,
        address _appContract,
        uint256 _sourceChain,
        uint256 _destinationChain,
        bytes memory _data,
        uint256 _gasLimit,
        uint256 _nonce
    ) external onlyValidator {
        bytes32 key = this.eventKey(
            _hash,
            _generation,
            _appContract,
            _sourceChain,
            _destinationChain,
            _data,
            _gasLimit,
            _nonce
        );
        require(registeredEvents[key] == false, "EventRegistry: event is already registered");

        registeredEvents[key] = true;
        emit EventRegistered(
            _hash,
            _generation,
            _appContract,
            _sourceChain,
            _destinationChain,
            _data,
            _gasLimit,
            _nonce
        );
    }

    function eventKey(
        bytes32 _hash,
        uint256 _generation,
        address _appContract,
        uint256 _sourceChain,
        uint256 _destinationChain,
        bytes memory _data,
        uint256 _gasLimit,
        uint256 _nonce
    ) external pure returns (bytes32) {
        return
            keccak256(
                abi.encodePacked(
                    _hash,
                    _generation,
                    _appContract,
                    _sourceChain,
                    _destinationChain,
                    _data,
                    _gasLimit,
                    _nonce
                )
            );
    }
}
