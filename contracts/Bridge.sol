// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract Bridge is Initializable {
    address public owner;
    address[] public validators;
    mapping(address => bool) public isValidator;
    mapping(bytes32 => mapping(address => bool)) public approvals;
    mapping(bytes32 => uint256) public approvalsCount;
    mapping(bytes32 => bool) public executed;
    uint256 public requiredApprovals;

    event Approved(bytes32 id, address validator);
    event Executed(bytes32 id, bytes data, address validator);

    modifier onlyOwner() {
        require(msg.sender == owner, "only owner");
        _;
    }

    modifier onlyValidator() {
        require(isValidator[msg.sender], "only validator");
        _;
    }

    function initialize(address _owner, address[] calldata _validators, uint256 _requiredApprovals) external initializer {
        owner = _owner;
        validators = _validators;
        for (uint256 i = 0; i < _validators.length; i++) {
            isValidator[_validators[i]] = true;
        }
        requiredApprovals = _requiredApprovals;
    }

    function setOwner(address _owner) external onlyOwner {
        owner = _owner;
    }

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

    function approve(bytes calldata _data) external onlyValidator {
        bytes32 id = keccak256(_data);
        if (!approvals[id][msg.sender]) {
            approvals[id][msg.sender] = true;
            approvalsCount[id]++;
            emit Approved(id, msg.sender);
        }

        if (executed[id]) {
            return;
        }

        if (approvalsCount[id] >= requiredApprovals) {
            _execute(_data);
        }
    }

    function _execute(bytes calldata _data) private {
        bytes32 id = keccak256(_data);
        executed[id] = true;
        emit Executed(id, _data, msg.sender);
    }
}
