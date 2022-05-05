// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "../common/AddressStorage.sol";

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
    uint256 public withdrawalPeriod;
    mapping(address => ValidatorInfo) public stakes;
    mapping(address => mapping(address => bool)) public slashingVotes;
    mapping(address => uint256) public slashingCount;
    mapping(address => WithdrawalAnnouncement) public withdrawalAnnouncements;
    AddressStorage public validatorStorage;

    event MinimalStakeUpdated(uint256 minimalStake);
    event WithdrawalPeriodUpdated(uint256 withdrawalPeriod);

    modifier onlyValidator() {
        require(stakes[msg.sender].status == ValidatorStatus.ACTIVE, "ValidatorStaking: only active validator");
        _;
    }

    function initialize(
        uint256 _minimalStake,
        uint256 _withdrawalPeriod,
        address _validatorStorage
    ) external initializer {
        minimalStake = _minimalStake;
        withdrawalPeriod = _withdrawalPeriod;
        validatorStorage = AddressStorage(_validatorStorage);
    }

    function setMinimalStake(uint256 _minimalStake) external onlyOwner {
        minimalStake = _minimalStake;
        emit MinimalStakeUpdated(_minimalStake);
    }

    function setWithdrawalPeriod(uint256 _withdrawalPeriod) external onlyOwner {
        withdrawalPeriod = _withdrawalPeriod;
        emit WithdrawalPeriodUpdated(_withdrawalPeriod);
    }

    function slash(address _validator) external onlyValidator {
        require(stakes[_validator].status != ValidatorStatus.SLASHED, "ValidatorStaking: validator is already slashed");
        if (slashingVotes[_validator][msg.sender] == false) {
            slashingVotes[_validator][msg.sender] = true;
            slashingCount[_validator] += 1;
        }

        if (slashingCount[_validator] >= ((validatorStorage.size() / 2) + 1)) {
            stakes[_validator].status = ValidatorStatus.SLASHED;
            validatorStorage.mustRemove(_validator);
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
            "ValidatorStaking: withdrawal period not passed"
        );

        uint256 withdrawalAmount = withdrawalAnnouncements[msg.sender].amount;

        require(withdrawalAmount <= stakes[msg.sender].stake, "ValidatorStaking: amount must be <= to validator stake");

        stakes[msg.sender].stake -= withdrawalAmount;

        withdrawalAnnouncements[msg.sender].amount = 0;
        withdrawalAnnouncements[msg.sender].time = 0;

        if (stakes[msg.sender].stake < minimalStake) {
            stakes[msg.sender].status = ValidatorStatus.INACTIVE;
            validatorStorage.mustRemove(msg.sender);
        }

        // solhint-disable-next-line avoid-low-level-calls
        (bool success, ) = msg.sender.call{value: withdrawalAmount, gas: 21000}("");
        require(success, "ValidatorStaking: transfer failed");
    }

    function stake() public payable {
        require(msg.value >= minimalStake, "ValidatorStaking: insufficient stake provided");
        require(stakes[msg.sender].status != ValidatorStatus.SLASHED, "ValidatorStaking: validator is slashed");

        if (stakes[msg.sender].status == ValidatorStatus.INACTIVE) {
            stakes[msg.sender].validator = msg.sender;
            stakes[msg.sender].status = ValidatorStatus.ACTIVE;
            validatorStorage.mustAdd(msg.sender);
        }

        stakes[msg.sender].stake += msg.value;
    }

    function getValidators() public view returns (address[] memory) {
        return validatorStorage.getAddresses();
    }
}
