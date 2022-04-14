// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract ValidatorStaking is Ownable, Initializable {
    enum ValidatorStatus {
        INACTIVE,
        ACTIVE,
        SLASHED,
        WITHDRAWING
    }

    struct ValidatorInfo {
        address validator;
        uint256 stake;
        ValidatorStatus status;
    }

    uint256 public minimalStake;
    mapping(address => ValidatorInfo) public stakes;

    function initialize(uint256 _minimalStake) external initializer {
        minimalStake = _minimalStake;
    }

    function setMinimalStake(uint256 _minimalStake) external onlyOwner {
        minimalStake = _minimalStake;
    }

    function stake() public payable {
        require(msg.value >= minimalStake, "insufficient stake provided");
        require(stakes[msg.sender].status != ValidatorStatus.SLASHED, "validator is slashed");
        stakes[msg.sender].validator = msg.sender;
        stakes[msg.sender].stake += msg.value;
        stakes[msg.sender].status = ValidatorStatus.ACTIVE;
    }
}
