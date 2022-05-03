// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./Signer.sol";

struct BroacastData {
    uint256 count;
    mapping(address => bytes) data;
}

contract DKG is Ownable, Initializable {
    Signer public signer;

    address[][] public validators;
    mapping(address => bool) public isValidator;

    // DKG rounds data
    mapping(uint256 => mapping(uint256 => BroacastData)) private roundBroadcastData;

    // Collective address voting
    mapping(uint256 => mapping(address => address)) private roundVotes;
    mapping(uint256 => mapping(address => uint256)) private roundVoteCounts;

    event RoundDataProvided(uint256 generation, uint256 round, address validator);
    event RoundDataFilled(uint256 generation, uint256 round);
    event ValidatorsUpdated(uint256 generation, address[] validators);
    event CollectiveSignerVoted(uint256 generation, address validator, address collectiveSigner);
    event SigningUpdated(address signing);

    modifier onlyValidator() {
        require(isValidator[msg.sender], "DKG: not a validator");
        _;
    }

    modifier roundIsCorrect(uint256 _generation, uint256 _round) {
        require(
            _round == 1 || roundBroadcastData[_generation][_round - 1].count == this.getValidatorsCount(_generation),
            "DKG: previous round was not filled"
        );
        _;
    }

    function initialize(address[] memory _validators) external initializer {
        _setValidators(_validators);
    }

    function setSigner(address _signing) external onlyOwner {
        signer = Signer(_signing);
        emit SigningUpdated(_signing);
    }

    function setValidators(address[] memory _validators) external onlyOwner {
        _setValidators(_validators);
    }

    function roundBroadcast(
        uint256 _generation,
        uint256 _round,
        bytes memory _rawData
    ) external onlyValidator roundIsCorrect(_generation, _round) {
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

    function voteCollectiveSigner(uint256 _generation, address _collectiveSigner)
        external
        onlyValidator
        roundIsCorrect(_generation, 4)
    {
        require(roundVotes[_generation][msg.sender] == address(0), "DKG: already voted");
        roundVotes[_generation][msg.sender] = _collectiveSigner;
        roundVoteCounts[_generation][_collectiveSigner]++;

        emit CollectiveSignerVoted(_generation, msg.sender, _collectiveSigner);

        bool enoughVotes = _enoughVotes(_generation, roundVoteCounts[_generation][_collectiveSigner]);
        bool signerChanged = signer.activeSigner() != _collectiveSigner;
        if (enoughVotes && signerChanged) {
            signer.updateCollectiveSigner(_generation, _collectiveSigner);
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

    function getValidatorsCount(uint256 _generation) external view returns (uint256) {
        return validators[_generation].length;
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

    function _enoughVotes(uint256 _generation, uint256 votes) private view returns (bool) {
        return votes > (validators[_generation].length / 2);
    }
}
