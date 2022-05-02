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

    event RoundDataProvided(uint256 id, uint256 round, address validator);
    event RoundDataFilled(uint256 id, uint256 round);

    event ValidatorsUpdated(uint256 id, address[] validators);

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
        uint256 _id,
        uint256 _round,
        bytes memory _rawData
    ) external onlyValidator {
        require(roundBroadcastData[_id][_round].data[msg.sender].length == 0, "DKG: round data already provided");
        roundBroadcastData[_id][_round].count++;
        roundBroadcastData[_id][_round].data[msg.sender] = _rawData;
        emit RoundDataProvided(_id, _round, msg.sender);
        if (roundBroadcastData[_id][_round].count == validators[_id].length) {
            emit RoundDataFilled(_id, _round);
        }
    }

    function isRoundFilled(uint256 _id, uint256 _round) external view returns (bool) {
        return roundBroadcastData[_id][_round].count == validators[_id].length;
    }

    function getRoundBroadcastCount(uint256 _id, uint256 _round) external view returns (uint256) {
        return roundBroadcastData[_id][_round].count;
    }

    function getRoundBroadcastData(
        uint256 _id,
        uint256 _round,
        address _validator
    ) external view returns (bytes memory) {
        return roundBroadcastData[_id][_round].data[_validator];
    }

    function getCurrentValidators() external view returns (address[] memory) {
        if (validators.length == 0) {
            return new address[](0);
        }

        return validators[validators.length - 1];
    }

    function getValidators(uint256 _id) external view returns (address[] memory) {
        return validators[_id];
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
