// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./ValidatorStorage.sol";

abstract contract ValidatorOwnable {
    ValidatorStorage public validatorStorage;

    modifier onlyValidator() {
        require(validatorStorage.getAddress() == msg.sender, "ValidatorOwnable: only validator");
        _;
    }

    function _setValidatorStorage(address _validatorStorage) internal virtual {
        validatorStorage = ValidatorStorage(_validatorStorage);
    }
}
