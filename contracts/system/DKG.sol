// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

struct BroacastData {
    uint256 count;
    mapping(address => bytes) data;
}

contract DKG is Ownable, Initializable {
    address[][] public validators;
    mapping(address => bool) public isValidator;

    mapping(uint256 => mapping(uint256 => BroacastData)) private roundBroadcastData;

    event RoundDataProvided(uint256 generation, uint256 round, address validator);
    event RoundDataFilled(uint256 generation, uint256 round);

    event ValidatorsUpdated(uint256 generation, address[] validators);

    modifier onlyValidator() {
        require(isValidator[msg.sender], "DKG: not a validator");
        _;
    }

    function initialize(address[] memory _validators) external initializer {
        _setValidators(_validators);
    }

    function setValidators(address[] memory _validators) external onlyOwner {
        _setValidators(_validators);
    }

    function roundBroadcast(
        uint256 _generation,
        uint256 _round,
        bytes memory _rawData
    ) external onlyValidator {
        require(
            roundBroadcastData[_generation][_round].data[msg.sender].length == 0,
            "DKG: round data already provided"
        );
        roundBroadcastData[_generation][_round].count++;
        roundBroadcastData[_generation][_round].data[msg.sender] = _rawData;
        emit RoundDataProvided(_generation, _round, msg.sender);
        if (roundBroadcastData[_generation][_round].count == validators[_generation].length) {
            emit RoundDataFilled(_generation, _round);
        }
    }

    function isRoundFilled(uint256 _generation, uint256 _round) external view returns (bool) {
        return roundBroadcastData[_generation][_round].count == validators[_generation].length;
    }

    function getRoundBroadcastCount(uint256 _generation, uint256 _round) external view returns (uint256) {
        return roundBroadcastData[_generation][_round].count;
    }

    function getRoundBroadcastData(
        uint256 _generation,
        uint256 _round,
        address _validator
    ) external view returns (bytes memory) {
        return roundBroadcastData[_generation][_round].data[_validator];
    }

    function getCurrentValidators() external view returns (address[] memory) {
        if (validators.length == 0) {
            return new address[](0);
        }

        return validators[validators.length - 1];
    }

    function getValidators(uint256 _generation) external view returns (address[] memory) {
        return validators[_generation];
    }

    function _setValidators(address[] memory _validators) private {
        address[] memory currentValidators = this.getCurrentValidators();
        for (uint256 i = 0; i < currentValidators.length; i++) {
            isValidator[currentValidators[i]] = false;
        }

        for (uint256 i = 0; i < _validators.length; i++) {
            isValidator[_validators[i]] = true;
        }

        validators.push(_validators);
        emit ValidatorsUpdated(validators.length - 1, _validators);
    }
}
