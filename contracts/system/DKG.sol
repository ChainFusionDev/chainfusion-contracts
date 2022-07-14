// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "./Staking.sol";

struct BroacastData {
    uint256 count;
    mapping(address => bytes) data;
}

contract DKG is Ownable, Initializable {
    Staking public staking;

    // Validators storage
    address[][] public validators;
    mapping(uint256 => mapping(address => bool)) public isGenerationValidator;

    // DKG rounds data
    mapping(uint256 => mapping(uint256 => BroacastData)) public roundBroadcastData;

    // Signer address voting
    mapping(uint256 => address) public signerAddresses;
    mapping(uint256 => mapping(address => address)) public signerVotes;
    mapping(uint256 => mapping(address => uint256)) public signerVoteCounts;

    event RoundDataProvided(uint256 generation, uint256 round, address validator);
    event RoundDataFilled(uint256 generation, uint256 round);

    event ValidatorsUpdated(uint256 generation, address[] validators);
    event SignerVoted(uint256 generation, address validator, address collectiveSigner);
    event SignerAddressUpdated(uint256 generation, address signerAddress);

    event ThresholdSignerUpdated(address signer);
    event StakingUpdated(address validatorStaking);

    modifier onlyValidator(uint256 _generation) {
        require(isGenerationValidator[_generation][msg.sender], "DKG: not a validator");
        _;
    }

    modifier onlyValidatorStaking() {
        require(msg.sender == address(staking), "DKG: not a staking");
        _;
    }

    modifier roundIsFilled(uint256 _generation, uint256 _round) {
        require(
            _round == 0 || roundBroadcastData[_generation][_round].count == validators[_generation].length,
            "DKG: round was not filled"
        );
        _;
    }

    modifier roundNotProvided(uint256 _generation, uint256 _round) {
        require(
            roundBroadcastData[_generation][_round].data[msg.sender].length == 0,
            "DKG: round data already provided"
        );
        _;
    }

    function initialize(address _staking) external initializer {
        setStaking(_staking);
    }

    function setValidators(address[] memory _validators) external onlyValidatorStaking {
        _setValidators(_validators);
    }

    function roundBroadcast(
        uint256 _generation,
        uint256 _round,
        bytes memory _rawData
    ) external onlyValidator(_generation) roundIsFilled(_generation, _round - 1) roundNotProvided(_generation, _round) {
        roundBroadcastData[_generation][_round].count++;
        roundBroadcastData[_generation][_round].data[msg.sender] = _rawData;
        emit RoundDataProvided(_generation, _round, msg.sender);
        if (roundBroadcastData[_generation][_round].count == validators[_generation].length) {
            emit RoundDataFilled(_generation, _round);
        }
    }

    function voteSigner(
        uint256 _generation,
        address _signerAddress,
        bytes memory _signature
    ) external onlyValidator(_generation) roundIsFilled(_generation, 3) {
        require(recoverSigner(_signature) == _signerAddress, "DKG: signature is invalid");
        require(signerVotes[_generation][msg.sender] == address(0), "DKG: already voted");

        signerVotes[_generation][msg.sender] = _signerAddress;
        signerVoteCounts[_generation][_signerAddress]++;

        emit SignerVoted(_generation, msg.sender, _signerAddress);

        bool enoughVotes = _enoughVotes(_generation, signerVoteCounts[_generation][_signerAddress]);
        bool signerChanged = signerAddresses[_generation] != _signerAddress;
        if (enoughVotes && signerChanged) {
            signerAddresses[_generation] = _signerAddress;
            emit SignerAddressUpdated(_generation, _signerAddress);
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

    function getGenerationsCount() external view returns (uint256) {
        return validators.length;
    }

    function isValidator(uint256 _generation, address _validator) external view returns (bool) {
        return isGenerationValidator[_generation][_validator];
    }

    function getValidators(uint256 _generation) external view returns (address[] memory) {
        if (validators.length > _generation) {
            return validators[_generation];
        }

        return new address[](0);
    }

    function getValidatorsCount(uint256 _generation) external view returns (uint256) {
        if (validators.length > _generation) {
            return validators[_generation].length;
        }

        return 0;
    }

    function setStaking(address _staking) public onlyOwner {
        staking = Staking(_staking);
        emit StakingUpdated(_staking);
    }

    function recoverSigner(bytes memory _signature) public pure returns (address) {
        bytes32 r;
        bytes32 s;
        uint8 v;

        if (_signature.length != 65) {
            return (address(0));
        }

        // solhint-disable-next-line no-inline-assembly
        assembly {
            r := mload(add(_signature, 32))
            s := mload(add(_signature, 64))
            v := and(mload(add(_signature, 65)), 255)
        }

        if (v < 27) {
            v += 27;
        }

        if (v != 27 && v != 28) {
            return (address(0));
        } else {
            string memory message = "verify";
            uint256 messageLen = bytes(message).length;

            string memory str = string(
                abi.encodePacked("\x19Ethereum Signed Message:\n", Strings.toString(messageLen), message)
            );

            bytes32 prefixedHashMessage = keccak256(abi.encodePacked(str));

            return ecrecover(prefixedHashMessage, v, r, s);
        }
    }

    function _setValidators(address[] memory _validators) private {
        if (_validators.length < 2) {
            // dkg requires at least 2 validators
            return;
        }

        uint256 newGeneration = validators.length;
        for (uint256 i = 0; i < _validators.length; i++) {
            isGenerationValidator[newGeneration][_validators[i]] = true;
        }

        validators.push(_validators);
        emit ValidatorsUpdated(newGeneration, _validators);
        emit RoundDataFilled(newGeneration, 0);
    }

    function _enoughVotes(uint256 _generation, uint256 votes) private view returns (bool) {
        return votes > (validators[_generation].length / 2);
    }
}
