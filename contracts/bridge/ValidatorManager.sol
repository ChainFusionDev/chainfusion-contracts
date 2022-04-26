// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./Bridge.sol";

contract ValidatorManager is Ownable, Initializable {
    address[] public validators;
    uint256 public requiredApprovals;
    mapping(address => bool) public isValidator;

    event ValidatorsUpdated(address[] validators);
    event RequiredApprovalsUpdated(uint256 requiredApprovals);

    function setValidators(address[] calldata _validators) external onlyOwner {
        for (uint256 i = 0; i < validators.length; i++) {
            isValidator[validators[i]] = false;
        }
        for (uint256 i = 0; i < _validators.length; i++) {
            isValidator[_validators[i]] = true;
        }
        validators = _validators;
        emit ValidatorsUpdated(_validators);
    }

    function setRequiredApprovals(uint256 _requiredApprovals) external onlyOwner {
        require(_requiredApprovals > 0, "ValidatorManager: required approvals too small");
        requiredApprovals = _requiredApprovals;
        emit RequiredApprovalsUpdated(_requiredApprovals);
    }
}
