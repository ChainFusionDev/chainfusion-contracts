// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract ValidatorStorage is Initializable {
    address public validator;

    event ValidatorUpdated(address validator);

    function initialize(address _validator) external initializer {
        validator = _validator;
    }

    function setAddress(address _newValidator) public {
        require(validator == msg.sender, "ValidatorStorage: only validator");
        validator = _newValidator;
        emit ValidatorUpdated(validator);
    }

    function getAddress() public view returns (address) {
        return validator;
    }
}
