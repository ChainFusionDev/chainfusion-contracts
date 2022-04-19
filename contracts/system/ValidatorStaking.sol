// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract ValidatorStaking is Ownable, Initializable {
    enum ValidatorStatus {
        INACTIVE,
        ACTIVE,
        SLASHED
    }

    struct ValidatorInfo {
        address validator;
        uint256 stake;
        ValidatorStatus status;
    }

    struct WithdrawalAnnouncement {
        uint256 amount;
        uint256 time;
    }

    uint256 public minimalStake;
    uint256 public validatorCount;
    uint256 public withdrawalPeriod;
    mapping(address => ValidatorInfo) public stakes;
    mapping(address => mapping(address => bool)) public slashingVotes;
    mapping(address => uint256) public slashingCount;
    mapping(address => WithdrawalAnnouncement) public withdrawalAnnouncements;

    modifier onlyValidator() {
        require(stakes[msg.sender].status == ValidatorStatus.ACTIVE, "ValidatorStaking: only active validator");
        _;
    }

    function initialize(uint256 _minimalStake, uint256 _withdrawalPeriod) external initializer {
        minimalStake = _minimalStake;
        withdrawalPeriod = _withdrawalPeriod;
    }

    function setMinimalStake(uint256 _minimalStake) external onlyOwner {
        minimalStake = _minimalStake;
    }

    function setwithdrawalPeriod(uint256 _withdrawalPeriod) external onlyOwner {
        withdrawalPeriod = _withdrawalPeriod;
    }

    function slash(address _validator) external onlyValidator {
        if (slashingVotes[_validator][msg.sender] == false) {
            slashingVotes[_validator][msg.sender] = true;
            slashingCount[_validator] += 1;
        }

        if (slashingCount[_validator] >= ((validatorCount / 2) + 1)) {
            stakes[_validator].status = ValidatorStatus.SLASHED;
            validatorCount--;
        }
    }

    function announceWithdrawal(uint256 _amount) external onlyValidator {
        require(_amount <= stakes[msg.sender].stake, "ValidatorStaking: amount must be <= to stake");
        withdrawalAnnouncements[msg.sender].amount = _amount;
        // solhint-disable-next-line not-rely-on-time
        withdrawalAnnouncements[msg.sender].time = block.timestamp;
    }

    function withdraw() external onlyValidator {
        require(withdrawalAnnouncements[msg.sender].amount > 0, "ValidatorStaking: amount must be greater than zero");
        require(
            // solhint-disable-next-line not-rely-on-time
            withdrawalAnnouncements[msg.sender].time + withdrawalPeriod <= block.timestamp,
            "withdrawalPeriod not passed"
        );

        uint256 withdrawalAmount = withdrawalAnnouncements[msg.sender].amount;
        withdrawalAnnouncements[msg.sender].amount = 0;
        withdrawalAnnouncements[msg.sender].time = 0;

        // solhint-disable-next-line avoid-low-level-calls
        (bool success, ) = msg.sender.call{value: withdrawalAmount, gas: 21000}("");
        require(success, "ValidatorStaking: transfer failed.");
    }

    function stake() public payable {
        require(msg.value >= minimalStake, "ValidatorStaking: insufficient stake provided");
        require(stakes[msg.sender].status != ValidatorStatus.SLASHED, "ValidatorStaking: validator is slashed");

        if (stakes[msg.sender].status == ValidatorStatus.INACTIVE) {
            stakes[msg.sender].validator = msg.sender;
            stakes[msg.sender].status = ValidatorStatus.ACTIVE;
            validatorCount++;
        }

        stakes[msg.sender].stake += msg.value;
    }
}
