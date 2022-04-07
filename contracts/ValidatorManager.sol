// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./Bridge.sol";

contract ValidatorManager is Ownable, Initializable {
    address[] public validators;
    mapping(address => bool) public isValidator;
    uint256 public requiredApprovals;

    function setValidators(address[] calldata _validators) external onlyOwner {
        for (uint256 i = 0; i < validators.length; i++) {
            isValidator[validators[i]] = false;
        }
        for (uint256 i = 0; i < _validators.length; i++) {
            isValidator[_validators[i]] = true;
        }
        validators = _validators;
    }

    function setRequiredApprovals(uint256 _requiredApprovals) external onlyOwner {
        require(_requiredApprovals > 0, "required approvals too small");
        requiredApprovals = _requiredApprovals;
    }
}
