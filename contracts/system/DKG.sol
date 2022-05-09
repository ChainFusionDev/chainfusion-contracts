// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./ThresholdSigner.sol";
import "./ValidatorStaking.sol";

struct BroacastData {
    uint256 count;
    mapping(address => bytes) data;
}

contract DKG is Ownable, Initializable {
    ThresholdSigner public thresholdSigner;
    ValidatorStaking public validatorStaking;

    address[][] public validators;
    mapping(address => bool) public isValidator;

    // DKG rounds data
    mapping(uint256 => mapping(uint256 => BroacastData)) private roundBroadcastData;

    // Signer address voting
    mapping(uint256 => mapping(address => address)) private signerVotes;
    mapping(uint256 => mapping(address => uint256)) private signerVoteCounts;

    event RoundDataProvided(uint256 generation, uint256 round, address validator);
    event RoundDataFilled(uint256 generation, uint256 round);
    event ValidatorsUpdated(uint256 generation, address[] validators);
    event SignerVoted(uint256 generation, address validator, address collectiveSigner);
    event ThresholdSignerUpdated(address signer);
    event ValidatorStakingUpdated(address validatorStaking);

    modifier onlyValidator() {
        require(isValidator[msg.sender], "DKG: not a validator");
        _;
    }

    modifier onlyValidatorStaking() {
        require(msg.sender == address(validatorStaking), "DKG: not a validatorStaking");
        _;
    }

    modifier roundIsCorrect(uint256 _generation, uint256 _round) {
        require(
            _round == 1 || roundBroadcastData[_generation][_round - 1].count == this.getValidatorsCount(_generation),
            "DKG: previous round was not filled"
        );
        _;
    }

    function initialize(address[] memory _validators, address _validatorStaking) external initializer {
        _setValidators(_validators);
        validatorStaking = ValidatorStaking(_validatorStaking);
    }

    function setThresholdSigner(address _thresholdSigner) external onlyOwner {
        thresholdSigner = ThresholdSigner(_thresholdSigner);
        emit ThresholdSignerUpdated(_thresholdSigner);
    }

    function setValidators(address[] memory _validators) external onlyValidatorStaking {
        _setValidators(_validators);
    }

    function setValidatorStaking(address _validatorStaking) external onlyOwner {
        validatorStaking = ValidatorStaking(_validatorStaking);
        emit ValidatorStakingUpdated(_validatorStaking);
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

    function voteSigner(uint256 _generation, address _signerAddress)
        external
        onlyValidator
        roundIsCorrect(_generation, 4)
    {
        require(signerVotes[_generation][msg.sender] == address(0), "DKG: already voted");
        signerVotes[_generation][msg.sender] = _signerAddress;
        signerVoteCounts[_generation][_signerAddress]++;

        emit SignerVoted(_generation, msg.sender, _signerAddress);

        bool enoughVotes = _enoughVotes(_generation, signerVoteCounts[_generation][_signerAddress]);
        bool signerChanged = thresholdSigner.activeAddress() != _signerAddress;
        if (enoughVotes && signerChanged) {
            thresholdSigner.updateSignerAddress(_generation, _signerAddress);
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
        emit RoundDataFilled(validators.length - 1, 0);
    }

    function _enoughVotes(uint256 _generation, uint256 votes) private view returns (bool) {
        return votes > (validators[_generation].length / 2);
    }
}
