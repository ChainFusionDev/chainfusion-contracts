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
    address[] public validators;
    mapping(address => bool) public isValidator;

    mapping(bytes32 => BroacastInfo) private round1BroadcastData;
    mapping(bytes32 => BroacastInfo) private round2BroadcastData;
    mapping(bytes32 => BroacastInfo) private round3BroadcastData;

    event Round1Provided(bytes32 id, address validator);
    event Round2Provided(bytes32 id, address validator);
    event Round3Provided(bytes32 id, address validator);

    event Round1Filled(bytes32 id);
    event Round2Filled(bytes32 id);
    event Round3Filled(bytes32 id);

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
        bytes32 _id,
        bytes[] memory _public,
        bytes[] memory _private
    ) external onlyValidator {
        require(validators.length == _public.length, "invalid public data length");
        require(validators.length == _private.length, "invalid private data length");
        require(!round1BroadcastData[_id].data[msg.sender].provided, "data already provided");
        round1BroadcastData[_id].count++;
        round1BroadcastData[_id].data[msg.sender] = BroadcastData(true, _public, _private);
        emit Round1Provided(_id, msg.sender);
        if (round1BroadcastData[_id].count == validators.length) {
            emit Round1Filled(_id);
        }
    }

    function round2Broadcast(bytes32 _id, bytes[] memory _public) external onlyValidator {
        require(validators.length == _public.length, "invalid public data length");
        require(!round2BroadcastData[_id].data[msg.sender].provided, "data already provided");
        require(round1BroadcastData[_id].count == validators.length, "round 1 not finished");
        round2BroadcastData[_id].count++;
        round2BroadcastData[_id].data[msg.sender] = BroadcastData(true, _public, new bytes[](0));
        emit Round2Provided(_id, msg.sender);
        if (round2BroadcastData[_id].count == validators.length) {
            emit Round2Filled(_id);
        }
    }

    function round3Broadcast(bytes32 _id, bytes[] memory _public) external onlyValidator {
        require(validators.length == _public.length, "invalid public data length");
        require(!round3BroadcastData[_id].data[msg.sender].provided, "data already provided");
        require(round1BroadcastData[_id].count == validators.length, "round 1 not finished");
        require(round2BroadcastData[_id].count == validators.length, "round 2 not finished");
        round3BroadcastData[_id].count++;
        round3BroadcastData[_id].data[msg.sender] = BroadcastData(true, _public, new bytes[](0));
        emit Round3Provided(_id, msg.sender);
        if (round3BroadcastData[_id].count == validators.length) {
            emit Round3Filled(_id);
        }
    }

    function getRound1BroadcastCount(bytes32 _id) external view returns (uint256) {
        return round1BroadcastData[_id].count;
    }

    function getRound2BroadcastCount(bytes32 _id) external view returns (uint256) {
        return round2BroadcastData[_id].count;
    }

    function getRound3BroadcastCount(bytes32 _id) external view returns (uint256) {
        return round3BroadcastData[_id].count;
    }

    function getRound1BroadcastData(bytes32 _id, address _validator) external view returns (BroadcastData memory) {
        return round1BroadcastData[_id].data[_validator];
    }

    function getRound2BroadcastData(bytes32 _id, address _validator) external view returns (BroadcastData memory) {
        return round2BroadcastData[_id].data[_validator];
    }

    function getRound3BroadcastData(bytes32 _id, address _validator) external view returns (BroadcastData memory) {
        return round3BroadcastData[_id].data[_validator];
    }

    function getValidators() external view returns (address[] memory) {
        return validators;
    }

    function _setValidators(address[] memory _validators) private {
        for (uint256 i = 0; i < validators.length; i++) {
            isValidator[validators[i]] = false;
        }

        for (uint256 i = 0; i < _validators.length; i++) {
            isValidator[_validators[i]] = true;
        }

        validators = _validators;
    }
}
