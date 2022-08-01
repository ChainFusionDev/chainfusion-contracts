// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./ValidatorOwnable.sol";
import "./ContractKeys.sol";
import "./ContractRegistry.sol";
import "./Staking.sol";

enum SlashingReason {
    REASON_NO_RECENT_BLOCKS,
    REASON_DKG_INACTIVITY,
    REASON_DKG_VIOLATION,
    REASON_SIGNING_INACTIVITY,
    REASON_SIGNING_VIOLATION
}

enum SlashingReasonGroup {
    NONE,
    REASON_GROUP_BLOCKS,
    REASON_GROUP_DKG,
    REASON_GROUP_SIGNING
}

contract SlashingVoting is ContractKeys, ValidatorOwnable, Initializable {
    ContractRegistry public contractRegistry;

    uint256 public epochPeriod;
    uint256 public slashingThresold;
    uint256 public slashingEpochs;

    // Votes
    mapping(bytes32 => mapping(address => bool)) public votes;
    mapping(bytes32 => uint256) public voteCounts;

    // Bans
    mapping(bytes32 => bool) public bans;
    mapping(uint256 => mapping(address => mapping(SlashingReason => bool))) public bansByReason;
    mapping(uint256 => mapping(address => uint256)) public bansByEpoch;

    event VotedWithReason(address voter, address validator, SlashingReason reason);
    event BannedWithReason(address validator, SlashingReason reason);
    event SlashedWithReason(address validator);

    modifier onlyActiveValidator() {
        require(_stakingContract().isValidatorActive(msg.sender) == true, "SlashingVoting: only active validator");
        _;
    }

    function initialize(
        address _signerGetterAddress,
        uint256 _epochPeriod,
        uint256 _slashingThresold,
        uint256 _lashingEpochs,
        address _contractRegistry
    ) external initializer {
        _setSignerGetter(_signerGetterAddress);
        setEpochPeriod(_epochPeriod);
        setSlashingThresold(_slashingThresold);
        setSlashingEpochs(_lashingEpochs);
        contractRegistry = ContractRegistry(_contractRegistry);
    }

    function voteWithReason(
        address _validator,
        SlashingReason _reason,
        bytes calldata _nonce
    ) external onlyActiveValidator {
        Staking staking = _stakingContract();
        bytes32 voteHash = votingHashWithReason(_validator, _reason, _nonce);

        require(staking.isValidatorActive(_validator) == true, "SlashingVoting: target is not active validator");
        require(bans[voteHash] == false, "SlashingVoting: validator is already banned");
        require(votes[voteHash][msg.sender] == false, "SlashingVoting: voter is already voted against given validator");

        votes[voteHash][msg.sender] = true;
        voteCounts[voteHash]++;
        emit VotedWithReason(msg.sender, _validator, _reason);

        uint256 epoch = currentEpoch();
        address[] memory validators = staking.getValidators();
        if (voteCounts[voteHash] >= (validators.length / 2 + 1)) {
            bans[voteHash] = true;
            bansByReason[epoch][_validator][_reason] = true;
            bansByEpoch[epoch][_validator]++;
            emit BannedWithReason(_validator, _reason);
        }

        if (shouldShash(epoch, _validator)) {
            _stakingContract().slash(_validator);
            emit SlashedWithReason(_validator);
        }
    }

    function setEpochPeriod(uint256 _epochPeriod) public onlyValidator {
        epochPeriod = _epochPeriod;
    }

    function setSlashingThresold(uint256 _slashingThresold) public onlyValidator {
        slashingThresold = _slashingThresold;
    }

    function setSlashingEpochs(uint256 _slashingEpochs) public onlyValidator {
        slashingEpochs = _slashingEpochs;
    }

    function isBannedByReason(address _validator, SlashingReason _reason) public view returns (bool) {
        return bansByReason[currentEpoch()][_validator][_reason];
    }

    function shouldShash(uint256 _epoch, address _validator) public view returns (bool) {
        if (_epoch < slashingEpochs) {
            return false;
        }

        uint256 totalBans;
        for (uint256 i = _epoch; i > _epoch - slashingEpochs; i--) {
            uint256 bansInEpoch = bansByEpoch[i][_validator];
            if (bansInEpoch == 0) {
                return false;
            }

            totalBans += bansInEpoch;
        }

        return totalBans >= slashingThresold;
    }

    function getBansByEpoch(uint256 _epoch, address _validator) public view returns (uint256) {
        return bansByEpoch[_epoch][_validator];
    }

    function currentEpoch() public view returns (uint256) {
        return epochByBlock(block.number);
    }

    function epochByBlock(uint256 _blockNumber) public view returns (uint256) {
        return _blockNumber / epochPeriod;
    }

    function votingHashWithReason(
        address _validator,
        SlashingReason _reason,
        bytes calldata _nonce
    ) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(_validator, _reason, _nonce));
    }

    function _stakingContract() private view returns (Staking) {
        return Staking(contractRegistry.getContract(STAKING_KEY));
    }
}
