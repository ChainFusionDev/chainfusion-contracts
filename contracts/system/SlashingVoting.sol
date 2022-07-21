// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./ValidatorStaking.sol";

contract SlashingVoting is Initializable {
    enum SlashingReason {
        REASON_NO_RECENT_BLOCKS,
        REASON_DKG_INACTIVITY,
        REASON_DKG_VIOLATION,
        REASON_SIGNING_INACTIVITY,
        REASON_SIGNING_VIOLATION
    }

    address public owner;
    ValidatorStaking public staking;
    uint256 public epochPeriod;
    uint256 public slashingThresold;
    uint256 public slashingEpochs;
    mapping(bytes32 => mapping(address => bool)) public votes;
    mapping(bytes32 => uint256) public voteCounts;
    mapping(bytes32 => bool) public ban;
    mapping(uint256 => mapping(address => uint256)) public banCounts;
    mapping(address => bool) public slashes;

    event VotedWithReason(address voter, address validator, SlashingReason reason);
    event BannedWithReason(address validator, SlashingReason reason);
    event SlashedWithReason(address validator);

    modifier onlyOwner() {
        require(msg.sender == owner, "SlashingVoting: caller is not the owner");
        _;
    }

    modifier onlyValidator() {
        require(staking.validatorOnly(msg.sender) == true, "SlashingVoting: only active validator");
        _;
    }

    function initialize() external initializer {
        owner = msg.sender;
    }

    function votedWithReason(
        address _validator,
        SlashingReason _reason,
        bytes calldata _nonse
    ) external onlyValidator {
        bytes32 voteHash = votingHashWithReason(_validator, _reason, _nonse);
        require(ban[voteHash] == false, "SlashingVoting: validator already baned");
        require(slashes[_validator] == false, "SlashingVoting: validator already slashed");
        require(
            votes[voteHash][msg.sender] == false,
            "SlashingVoting: voter is already voted against given validator "
        );
        address[] memory validatorsAddresses = staking.getValidators();
        emit VotedWithReason(msg.sender, _validator, _reason);

        if (voteCounts[voteHash] >= (validatorsAddresses.length / 2 + 1)) {
            banCounts[currentEpoch()][_validator] = banCounts[currentEpoch()][_validator]++;
            emit BannedWithReason(_validator, _reason);
        }

        if (bansTotal(currentEpoch(), _validator) >= slashingThresold) {
            slashes[_validator] = true;
            emit SlashedWithReason(_validator);
        }
    }

    function bansTotal(uint256 _epoch, address _validator) public view returns (uint256) {
        uint256 bansNubmer;
        for (uint256 i = _epoch; i > _epoch - slashingEpochs; i--) {
            bansNubmer = bansByEpoch(i, _validator);
        }

        return bansNubmer;
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

    function setStaking(ValidatorStaking _staking) internal onlyOwner {
        staking = _staking;
    }

    function setEpochPeriod(uint256 _epochPeriod) internal onlyOwner {
        epochPeriod = _epochPeriod;
    }

    function setSlashingThresold(uint256 _slashingThresold) internal onlyOwner {
        slashingThresold = _slashingThresold;
    }

    function setSlashingEpochs(uint256 _slashingEpochs) internal onlyOwner {
        slashingEpochs = _slashingEpochs;
    }
}
