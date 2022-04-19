// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

struct BroacastInfo {
    uint256 count;
    mapping(address => BroadcastData) data;
}

struct BroadcastData {
    bool provided;
    bytes[] publicData;
    bytes[] privateData;
}

contract DKG is Ownable, Initializable {
    address[][] public validators;
    mapping(address => bool) public isValidator;

    mapping(uint256 => BroacastInfo) private round1BroadcastData;
    mapping(uint256 => BroacastInfo) private round2BroadcastData;
    mapping(uint256 => BroacastInfo) private round3BroadcastData;

    event Round1Provided(uint256 id, address validator);
    event Round2Provided(uint256 id, address validator);
    event Round3Provided(uint256 id, address validator);

    event Round1Filled(uint256 id);
    event Round2Filled(uint256 id);
    event Round3Filled(uint256 id);

    event ValidatorsUpdated(address[] validators);

    modifier onlyValidator() {
        require(isValidator[msg.sender], "not a validator");
        _;
    }

    function initialize(address[] memory _validators) external initializer {
        _setValidators(_validators);
    }

    function setValidators(address[] memory _validators) external onlyOwner {
        _setValidators(_validators);
    }

    function round1Broadcast(
        uint256 _id,
        bytes[] memory _public,
        bytes[] memory _private
    ) external onlyValidator {
        require(!round1BroadcastData[_id].data[msg.sender].provided, "data already provided");
        round1BroadcastData[_id].count++;
        round1BroadcastData[_id].data[msg.sender] = BroadcastData(true, _public, _private);
        emit Round1Provided(_id, msg.sender);
        if (round1BroadcastData[_id].count == validators.length) {
            emit Round1Filled(_id);
        }
    }

    function round2Broadcast(uint256 _id, bytes[] memory _public) external onlyValidator {
        require(round1BroadcastData[_id].count == validators.length, "round 1 not finished");
        require(!round2BroadcastData[_id].data[msg.sender].provided, "data already provided");
        round2BroadcastData[_id].count++;
        round2BroadcastData[_id].data[msg.sender] = BroadcastData(true, _public, new bytes[](0));
        emit Round2Provided(_id, msg.sender);
        if (round2BroadcastData[_id].count == validators.length) {
            emit Round2Filled(_id);
        }
    }

    function round3Broadcast(uint256 _id, bytes[] memory _public) external onlyValidator {
        require(round1BroadcastData[_id].count == validators.length, "round 1 not finished");
        require(round2BroadcastData[_id].count == validators.length, "round 2 not finished");
        require(!round3BroadcastData[_id].data[msg.sender].provided, "data already provided");
        round3BroadcastData[_id].count++;
        round3BroadcastData[_id].data[msg.sender] = BroadcastData(true, _public, new bytes[](0));
        emit Round3Provided(_id, msg.sender);
        if (round3BroadcastData[_id].count == validators.length) {
            emit Round3Filled(_id);
        }
    }

    function getRound1BroadcastCount(uint256 _id) external view returns (uint256) {
        return round1BroadcastData[_id].count;
    }

    function getRound2BroadcastCount(uint256 _id) external view returns (uint256) {
        return round2BroadcastData[_id].count;
    }

    function getRound3BroadcastCount(uint256 _id) external view returns (uint256) {
        return round3BroadcastData[_id].count;
    }

    function getRound1BroadcastData(uint256 _id, address _validator) external view returns (BroadcastData memory) {
        return round1BroadcastData[_id].data[_validator];
    }

    function getRound2BroadcastData(uint256 _id, address _validator) external view returns (BroadcastData memory) {
        return round2BroadcastData[_id].data[_validator];
    }

    function getRound3BroadcastData(uint256 _id, address _validator) external view returns (BroadcastData memory) {
        return round3BroadcastData[_id].data[_validator];
    }

    function getValidators() external view returns (address[] memory) {
        if (validators.length == 0) {
            return new address[](0);
        }

        return validators[validators.length - 1];
    }

    function _setValidators(address[] memory _validators) private {
        address[] memory currentValidators = this.getValidators();
        for (uint256 i = 0; i < validators.length; i++) {
            isValidator[currentValidators[i]] = false;
        }

        for (uint256 i = 0; i < _validators.length; i++) {
            isValidator[_validators[i]] = true;
        }

        validators.push(_validators);
        emit ValidatorsUpdated(_validators);
    }
}
