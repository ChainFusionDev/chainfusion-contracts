// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./ContractKeys.sol";
import "./ContractRegistry.sol";
import "./Staking.sol";

contract SlashingVoting is ContractKeys, Ownable, Initializable {
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

    ContractRegistry public contractRegistry;

    uint256 public epochPeriod;
    uint256 public slashingThresold;
    uint256 public slashingEpochs;

    // Votes
    mapping(bytes32 => mapping(address => bool)) public votes;
    mapping(bytes32 => uint256) public voteCounts;

    // Bans
    mapping(bytes32 => bool) public ban;
    mapping(uint256 => mapping(address => uint256)) public banCounts;
    mapping(uint256 => mapping(address => mapping(SlashingReasonGroup => bool))) public bansByGroup;
    mapping(uint256 => mapping(SlashingReason => SlashingReasonGroup)) public banGroupByReasons;

    event VotedWithReason(address voter, address validator, SlashingReason reason);
    event BannedWithReason(address validator, SlashingReason reason);
    event SlashedWithReason(address validator);

    modifier onlyValidator() {
        require(_stakingContract().isValidatorActive(msg.sender) == true, "SlashingVoting: only active validator");
        _;
    }

    function initialize(
        uint256 _epochPeriod,
        uint256 _slashingThresold,
        uint256 _lashingEpochs,
        address _contractRegistry
    ) external initializer {
        setEpochPeriod(_epochPeriod);
        setSlashingThresold(_slashingThresold);
        setSlashingEpochs(_lashingEpochs);
        contractRegistry = ContractRegistry(_contractRegistry);
        _setBanGroupByReasons();
    }

    function voteWithReason(
        address _validator,
        SlashingReason _reason,
        bytes calldata _nonse
    ) external onlyValidator {
        Staking staking = _stakingContract();
        bytes32 voteHash = votingHashWithReason(_validator, _reason, _nonse);

        require(staking.isValidatorActive(_validator) == true, "SlashingVoting: target is not active validator");
        require(ban[voteHash] == false, "SlashingVoting: validator is already banned");
        require(
            votes[voteHash][msg.sender] == false,
            "SlashingVoting: voter is already voted against given validator "
        );

        votes[voteHash][msg.sender] = true;
        voteCounts[voteHash]++;
        emit VotedWithReason(msg.sender, _validator, _reason);

        address[] memory validatorsAddresses = staking.getValidators();
        if (voteCounts[voteHash] >= (validatorsAddresses.length / 2 + 1)) {
            ban[voteHash] = true;
            banCounts[currentEpoch()][_validator]++;
            _setBanStatus(_reason, _validator);
            emit BannedWithReason(_validator, _reason);
        }

        if (isBanNeed(currentEpoch(), _validator)) {
            _stakingContract().slash(_validator);
            emit SlashedWithReason(_validator);
        }
    }

    function setEpochPeriod(uint256 _epochPeriod) public onlyOwner {
        epochPeriod = _epochPeriod;
    }

    function setSlashingThresold(uint256 _slashingThresold) public onlyOwner {
        slashingThresold = _slashingThresold;
    }

    function setSlashingEpochs(uint256 _slashingEpochs) public onlyOwner {
        slashingEpochs = _slashingEpochs;
    }

    function isBanByDKGGroup(address _validator) public view returns (bool) {
        return bansByGroup[currentEpoch()][_validator][SlashingReasonGroup.REASON_GROUP_DKG];
    }

    function isBanNeed(uint256 _epoch, address _validator) public view returns (bool) {
        uint256 bansNubmer;
        uint256 bansSeries;
        for (int256 i = int256(_epoch); i > int256(_epoch) - int256((slashingEpochs)); i--) {
            if (i < 0) {
                return false;
            }

            uint256 bansInEpoch = bansByEpoch(uint256(i), _validator);
            if (bansInEpoch != 0) {
                bansNubmer += bansInEpoch;
                bansSeries++;
            }
        }

        if (bansSeries == slashingEpochs && bansNubmer >= slashingEpochs) {
            return true;
        }

        return false;
    }

    function bansByEpoch(uint256 _epoch, address _validator) public view returns (uint256) {
        return banCounts[_epoch][_validator];
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
        bytes calldata _nonse
    ) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(_validator, _reason, _nonse));
    }

    function _setBanStatus(SlashingReason _reason, address _validator) private {
        SlashingReasonGroup bansGroup = banGroupByReasons[currentEpoch()][_reason];
        bansByGroup[currentEpoch()][_validator][bansGroup] = true;
    }

    function _setBanGroupByReasons() private {
        banGroupByReasons[currentEpoch()][SlashingReason.REASON_NO_RECENT_BLOCKS] = SlashingReasonGroup
            .REASON_GROUP_BLOCKS;
        banGroupByReasons[currentEpoch()][SlashingReason.REASON_DKG_INACTIVITY] = SlashingReasonGroup.REASON_GROUP_DKG;
        banGroupByReasons[currentEpoch()][SlashingReason.REASON_DKG_VIOLATION] = SlashingReasonGroup.REASON_GROUP_DKG;
        banGroupByReasons[currentEpoch()][SlashingReason.REASON_SIGNING_INACTIVITY] = SlashingReasonGroup
            .REASON_GROUP_SIGNING;
        banGroupByReasons[currentEpoch()][SlashingReason.REASON_SIGNING_VIOLATION] = SlashingReasonGroup
            .REASON_GROUP_SIGNING;
    }

    function _stakingContract() private view returns (Staking) {
        return Staking(contractRegistry.getContract(STAKING_KEY));
    }
}
