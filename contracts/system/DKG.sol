// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "./Staking.sol";
import "./ContractKeys.sol";
import "./ContractRegistry.sol";
import "./SlashingVoting.sol";

struct GenerationInfo {
    address[] validators;
    mapping(address => bool) isGenerationValidator;
    address signer;
    mapping(address => address) signerVotes;
    mapping(address => uint256) signerVoteCounts;
    mapping(uint256 => BroadcastData) roundBroadcastData;
    uint256 deadline;
}

struct BroadcastData {
    uint256 count;
    mapping(address => bytes) data;
}

contract DKG is ContractKeys, Ownable, Initializable {
    using ECDSA for bytes;
    using ECDSA for bytes32;

    enum GenerationStatus {
        PENDING,
        EXPIRED,
        ACTIVE
    }

    ContractRegistry public contractRegistry;

    GenerationInfo[] public generations;
    uint256 public lastActiveGeneration;
    uint256 public deadlinePeriod;

    event RoundDataProvided(uint256 generation, uint256 round, address validator);
    event RoundDataFilled(uint256 generation, uint256 round);

    event ValidatorsUpdated(uint256 generation, address[] validators);
    event SignerVoted(uint256 generation, address validator, address collectiveSigner);
    event SignerAddressUpdated(uint256 generation, address signerAddress);

    event ThresholdSignerUpdated(address signer);

    modifier onlyValidator(uint256 _generation) {
        require(
            generations.length > _generation && generations[_generation].isGenerationValidator[msg.sender],
            "DKG: not a validator"
        );
        _;
    }

    modifier onlyContract() {
        require(
            msg.sender == address(_stakingContract()) || msg.sender == address(_slashingVotingContract()),
            "DKG: not contract"
        );
        _;
    }

    modifier roundIsFilled(uint256 _generation, uint256 _round) {
        require(
            _round == 0 ||
                generations[_generation].roundBroadcastData[_round].count == generations[_generation].validators.length,
            "DKG: round was not filled"
        );
        _;
    }

    modifier roundNotProvided(uint256 _generation, uint256 _round) {
        require(
            generations[_generation].roundBroadcastData[_round].data[msg.sender].length == 0,
            "DKG: round data already provided"
        );
        _;
    }

    modifier onlyActiveSigner() {
        require(msg.sender == generations[lastActiveGeneration].signer, "DKG: not a active signer");
        _;
    }

    modifier deadlineViolated(uint256 _generation) {
        require(generations[_generation].deadline >= block.number, "DKG: deadline is violated");
        _;
    }

    function initialize(address _contractRegistry, uint256 _deadlinePeriod) external initializer {
        contractRegistry = ContractRegistry(_contractRegistry);
        deadlinePeriod = _deadlinePeriod;
    }

    function updateGeneration(address _validator) external onlyContract {
        uint256 newGeneration = generations.length;
        uint256 lastgeneration;

        if (newGeneration == 0) {
            lastgeneration = 0;
        }

        if (newGeneration > 0) {
            lastgeneration = newGeneration - 1;
        }

        if (newGeneration > lastActiveGeneration) {
            lastActiveGeneration = newGeneration;
        }

        generations.push();
        generations[newGeneration].validators = getValidators(lastgeneration);
        address[] storage validators = generations[newGeneration].validators;

        if (
            _stakingContract().isValidatorActive(_validator) == true &&
            _slashingVotingContract().isBanStatusNone(_validator)
        ) {
            generations[newGeneration].validators.push(_validator);
        }

        if (
            _stakingContract().isValidatorActive(_validator) != true ||
            !_slashingVotingContract().isBanStatusNone(_validator)
        ) {
            for (uint256 i = 0; i < validators.length; i++) {
                if (validators[i] == _validator) {
                    validators[i] = validators[validators.length - 1];
                    validators.pop();
                    break;
                }
            }
        }

        for (uint256 i = 0; i < validators.length; i++) {
            generations[newGeneration].isGenerationValidator[validators[i]] = true;
        }

        generations[newGeneration].validators = validators;
        generations[newGeneration].deadline = block.number + deadlinePeriod;
        emit ValidatorsUpdated(newGeneration, validators);
        emit RoundDataFilled(newGeneration, 0);
    }

    function roundBroadcast(
        uint256 _generation,
        uint256 _round,
        bytes memory _rawData
    )
        external
        onlyValidator(_generation)
        roundIsFilled(_generation, _round - 1)
        roundNotProvided(_generation, _round)
        deadlineViolated(_generation)
    {
        generations[_generation].roundBroadcastData[_round].count++;
        generations[_generation].roundBroadcastData[_round].data[msg.sender] = _rawData;
        emit RoundDataProvided(_generation, _round, msg.sender);
        if (generations[_generation].roundBroadcastData[_round].count == generations[_generation].validators.length) {
            emit RoundDataFilled(_generation, _round);
        }
    }

    function voteSigner(
        uint256 _generation,
        address _signerAddress,
        bytes memory _signature
    ) external onlyValidator(_generation) roundIsFilled(_generation, 3) deadlineViolated(_generation) {
        address recoveredSigner = bytes("verify").toEthSignedMessageHash().recover(_signature);
        require(recoveredSigner == _signerAddress, "DKG: signature is invalid");
        require(generations[_generation].signerVotes[msg.sender] == address(0), "DKG: already voted");

        generations[_generation].signerVotes[msg.sender] = _signerAddress;
        generations[_generation].signerVoteCounts[_signerAddress]++;

        emit SignerVoted(_generation, msg.sender, _signerAddress);

        bool enoughVotes = _enoughVotes(_generation, generations[_generation].signerVoteCounts[_signerAddress]);
        bool signerChanged = generations[_generation].signer != _signerAddress;
        if (enoughVotes && signerChanged) {
            generations[_generation].signer = _signerAddress;
            emit SignerAddressUpdated(_generation, _signerAddress);
        }
    }

    function getStatus(uint256 _generation) external view returns (GenerationStatus) {
        if (generations[_generation].signer != address(0)) {
            return GenerationStatus.ACTIVE;
        }

        if (generations[_generation].deadline >= block.number) {
            return GenerationStatus.PENDING;
        }

        return GenerationStatus.EXPIRED;
    }

    function isRoundFilled(uint256 _generation, uint256 _round) external view returns (bool) {
        return generations[_generation].roundBroadcastData[_round].count == generations[_generation].validators.length;
    }

    function getRoundBroadcastCount(uint256 _generation, uint256 _round) external view returns (uint256) {
        return generations[_generation].roundBroadcastData[_round].count;
    }

    function getRoundBroadcastData(
        uint256 _generation,
        uint256 _round,
        address _validator
    ) external view returns (bytes memory) {
        return generations[_generation].roundBroadcastData[_round].data[_validator];
    }

    function getCurrentValidators() external view returns (address[] memory) {
        if (generations.length == 0) {
            return new address[](0);
        }

        return generations[generations.length - 1].validators;
    }

    function getGenerationsCount() external view returns (uint256) {
        return generations.length;
    }

    function isValidator(uint256 _generation, address _validator) external view returns (bool) {
        if (generations.length > _generation) {
            return generations[_generation].isGenerationValidator[_validator];
        }

        return false;
    }

    function getValidatorsCount(uint256 _generation) external view returns (uint256) {
        if (generations.length > _generation) {
            return generations[_generation].validators.length;
        }

        return 0;
    }

    function setDeadlinePeriod(uint256 _deadlinePeriod) public onlyActiveSigner {
        deadlinePeriod = _deadlinePeriod;
    }

    function getValidators(uint256 _generation) public view returns (address[] memory) {
        if (generations.length > _generation) {
            return generations[_generation].validators;
        }

        return new address[](0);
    }

    function _removeValidators(address _validator) private {
        uint256 newGeneration = generations.length;

        if (newGeneration > lastActiveGeneration) {
            lastActiveGeneration = newGeneration;
        }

        generations.push();
        address[] storage newValidators = generations[newGeneration - 1].validators;
        newValidators.push(_validator);
        generations[newGeneration].validators = newValidators;
        for (uint256 i = 0; i < generations[newGeneration].validators.length; i++) {
            generations[newGeneration].isGenerationValidator[generations[newGeneration].validators[i]] = true;
        }

        generations[newGeneration].deadline = block.number + deadlinePeriod;
        emit ValidatorsUpdated(newGeneration, generations[newGeneration].validators);
        emit RoundDataFilled(newGeneration, 0);
    }

    function _enoughVotes(uint256 _generation, uint256 votes) private view returns (bool) {
        return votes > (generations[_generation].validators.length / 2);
    }

    function _stakingContract() private view returns (Staking) {
        return Staking(contractRegistry.getContract(STAKING_KEY));
    }

    function _slashingVotingContract() private view returns (SlashingVoting) {
        return SlashingVoting(contractRegistry.getContract(SLASHING_VOTING_KEY));
    }
}
