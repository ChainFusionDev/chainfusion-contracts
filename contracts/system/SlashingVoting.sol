// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "./Staking.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract SlashingVoting is Ownable, Initializable {
    enum SlashingReason {
        REASON_NO_RECENT_BLOCKS,
        REASON_DKG_INACTIVITY,
        REASON_DKG_VIOLATION,
        REASON_SIGNING_INACTIVITY,
        REASON_SIGNING_VIOLATION
    }

    Staking public staking;
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

    modifier onlyValidator() {
        require(staking.isValidatorActive(msg.sender) == true, "SlashingVoting: only active validator");
        _;
    }

    function initialize(
        address _staking,
        uint256 _epochPeriod,
        uint256 _slashingThresold,
        uint256 _lashingEpochs
    ) external initializer {
        setStaking(_staking);
        setEpochPeriod(_epochPeriod);
        setSlashingThresold(_slashingThresold);
        setSlashingEpochs(_lashingEpochs);
    }

    function voteWithReason(
        address _validator,
        SlashingReason _reason,
        bytes calldata _nonse
    ) external onlyValidator {
        bytes32 voteHash = votingHashWithReason(_validator, _reason, _nonse);

        require(staking.isValidatorActive(_validator) == true, "SlashingVoting: target is not active validator");
        require(ban[voteHash] == false, "SlashingVoting: validator already baned");
        require(slashes[_validator] == false, "SlashingVoting: validator already slashed");
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
            emit BannedWithReason(_validator, _reason);
        }

        if (bansTotal(currentEpoch(), _validator) >= slashingThresold) {
            slashes[_validator] = true;
            emit SlashedWithReason(_validator);
        }
    }

    function setStaking(address _staking) public onlyOwner {
        staking = Staking(_staking);
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

    function bansTotal(uint256 _epoch, address _validator) public view returns (uint256) {
        uint256 bansNubmer;
        for (int256 i = int256(_epoch); i > int256(_epoch) - int256((slashingEpochs)); i--) {
            if (i < 0) {
                break;
            }

            bansNubmer += bansByEpoch(uint256(i), _validator);
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
}
