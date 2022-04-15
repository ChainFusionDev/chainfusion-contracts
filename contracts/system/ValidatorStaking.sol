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

    uint256 public minimalStake;
    uint256 public validatorCount;
    mapping(address => ValidatorInfo) public stakes;
    mapping(address => mapping(address => bool)) public slashingVotes;
    mapping(address => uint256) public slashingCount;

    modifier onlyValidator() {
        require(stakes[msg.sender].status == ValidatorStatus.ACTIVE, "only active validator");
        _;
    }

    function initialize(uint256 _minimalStake) external initializer {
        minimalStake = _minimalStake;
    }

    function setMinimalStake(uint256 _minimalStake) external onlyOwner {
        minimalStake = _minimalStake;
    }

    function slash(address _validator) external onlyValidator {
        if (slashingVotes[_validator][msg.sender] == false) {
            slashingVotes[_validator][msg.sender] = true;
            slashingCount[_validator] += 1;
        }

        if (slashingCount[_validator] > ((validatorCount / 2) + 1)) {
            validatorCount--;
        }
    }

    function stake() public payable {
        require(msg.value >= minimalStake, "insufficient stake provided");
        require(stakes[msg.sender].status != ValidatorStatus.SLASHED, "validator is slashed");

        if (stakes[msg.sender].status == ValidatorStatus.INACTIVE) {
            stakes[msg.sender].validator = msg.sender;
            stakes[msg.sender].status = ValidatorStatus.ACTIVE;
            validatorCount++;
        }

        stakes[msg.sender].stake += msg.value;
    }
}
