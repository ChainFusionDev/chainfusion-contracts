// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface SignerGetter {
    function getSignerAddress() external view returns (address);

    function isCurrentValidator(address _validator) external view returns (bool);
}

// TODO: Rename to SignerOwnable
// As the result, there should be 2 contracts:
// SignerOwnable:
// - onlySigner() -> should use DKG.getSignerAddress()
//
// ValidatorOwnable:
// - onlyValidator() -> should use Staking.isValidatorActive()
abstract contract ValidatorOwnable {
    SignerGetter public signerGetter;

    // TODO: Rename to onlySigner()
    modifier onlyValidator() {
        require(signerGetter.getSignerAddress() == msg.sender, "ValidatorOwnable: only validator");
        _;
    }

    // TODO: Move to ValidatorOwnable and rename to onlyValidator()
    modifier onlyCurrentValidator() {
        require(signerGetter.isCurrentValidator(msg.sender), "ValidatorOwnable: only current validator");
        _;
    }

    function _setSignerGetter(address _signerGetterAddress) internal virtual {
        signerGetter = SignerGetter(_signerGetterAddress);
    }
}
