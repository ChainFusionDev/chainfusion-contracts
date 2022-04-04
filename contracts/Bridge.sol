// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol";

contract Bridge is Initializable {
    address public owner;
    address[] public validators;
    mapping(address => bool) public isValidator;
    mapping(bytes32 => mapping(address => bool)) public approvals;
    mapping(bytes32 => uint256) public approvalsCount;
    mapping(bytes32 => bool) public executed;
    uint256 public requiredApprovals;
    IERC20Upgradeable public token;

    event Approved(bytes32 id, address validator);
    event Executed(bytes32 id, bytes data, address validator);
    event Deposit(address token, uint256 chainId, uint256 amount);
    event Transferred(address token, address receiver, uint256 amount, address validator);

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

    function deposit (address _token, uint256 _chainId, uint256 _amount) external {
        require(_amount != 0, "Amount cannot be equal to 0.");
        require(IERC20Upgradeable(_token).transferFrom(msg.sender, address(this), _amount), "Transfer failed.");
        emit Deposit (_token, _chainId, _amount);
    }

    function approveTransfer(bytes calldata _txHash, address _token, address _receiver, uint256 _amount) external onlyValidator {
        bytes32 id = keccak256(abi.encodePacked(_txHash, _token, _receiver, _amount));
        
        if (!approvals[id][msg.sender]) {
            approvals[id][msg.sender] = true;
            approvalsCount[id]++;
            emit Approved(id, msg.sender);
        }

        if (executed[id]) {
            return;
        }

        if (approvalsCount[id] >= requiredApprovals) {
            executed[id] = true;
            require(IERC20Upgradeable(_token).transfer(_receiver, _amount), "Failed transfer");
            emit Transferred(_token, _receiver, _amount, msg.sender);
        }
    }
}
