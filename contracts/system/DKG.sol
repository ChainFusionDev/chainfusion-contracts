// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "./Staking.sol";
import "./ContractKeys.sol";
import "./ContractRegistry.sol";

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

contract DKG is ContractKeys, Initializable {
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

    modifier onlyDKGValidator(uint256 _generation) {
        require(
            generations.length > _generation && generations[_generation].isGenerationValidator[msg.sender],
            "DKG: not a validator"
        );
        _;
    }

    modifier onlyValidatorStaking() {
        require(msg.sender == address(_stakingContract()), "DKG: not a staking");
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

    modifier onlyPending(uint256 _generation) {
        require(getStatus(_generation) == GenerationStatus.PENDING, "DKG: not a pending generation");
        _;
    }

    function initialize(address _contractRegistry, uint256 _deadlinePeriod) external initializer {
        contractRegistry = ContractRegistry(_contractRegistry);
        deadlinePeriod = _deadlinePeriod;

        generations.push();
        generations[0].signer = msg.sender;
    }

    function setValidators(address[] memory _validators) external onlyValidatorStaking {
        _setValidators(_validators);
    }

    function roundBroadcast(
        uint256 _generation,
        uint256 _round,
        bytes memory _rawData
    )
        external
        onlyDKGValidator(_generation)
        roundIsFilled(_generation, _round - 1)
        roundNotProvided(_generation, _round)
        onlyPending(_generation)
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
    ) external onlyDKGValidator(_generation) roundIsFilled(_generation, 3) onlyPending(_generation) {
        address recoveredSigner = bytes("verify").toEthSignedMessageHash().recover(_signature);
        require(recoveredSigner == _signerAddress, "DKG: signature is invalid");
        require(generations[_generation].signerVotes[msg.sender] == address(0), "DKG: already voted");

        generations[_generation].signerVotes[msg.sender] = _signerAddress;
        generations[_generation].signerVoteCounts[_signerAddress]++;

        emit SignerVoted(_generation, msg.sender, _signerAddress);

        bool enoughVotes = _enoughVotes(_generation, generations[_generation].signerVoteCounts[_signerAddress]);
        if (enoughVotes) {
            generations[_generation].signer = _signerAddress;
            emit SignerAddressUpdated(_generation, _signerAddress);
        }
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
        return generations[generations.length - 1].validators;
    }

    function getSignerAddresses() external view returns (address) {
        return generations[lastActiveGeneration].signer;
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

    function getValidators(uint256 _generation) external view returns (address[] memory) {
        if (generations.length > _generation) {
            return generations[_generation].validators;
        }

        return new address[](0);
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

    function getStatus(uint256 _generation) public view returns (GenerationStatus) {
        if (generations[_generation].signer != address(0)) {
            return GenerationStatus.ACTIVE;
        }

        if (generations[_generation].deadline >= block.number) {
            return GenerationStatus.PENDING;
        }

        return GenerationStatus.EXPIRED;
    }

    function _setValidators(address[] memory _validators) private {
        if (_validators.length < 2) {
            // dkg requires at least 2 validators
            return;
        }

        uint256 newGeneration = generations.length;

        if (newGeneration > lastActiveGeneration) {
            lastActiveGeneration = newGeneration;
        }

        generations.push();
        for (uint256 i = 0; i < _validators.length; i++) {
            generations[newGeneration].isGenerationValidator[_validators[i]] = true;
        }

        generations[newGeneration].validators = _validators;
        generations[newGeneration].deadline = block.number + deadlinePeriod;

        emit ValidatorsUpdated(newGeneration, _validators);
        emit RoundDataFilled(newGeneration, 0);
    }

    function _enoughVotes(uint256 _generation, uint256 votes) private view returns (bool) {
        return votes > (generations[_generation].validators.length / 2);
    }

    function _stakingContract() private view returns (Staking) {
        return Staking(contractRegistry.getContract(STAKING_KEY));
    }
}
