// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface SignerGetter {
    function getSignerAddresses() external view returns (address);
}

abstract contract ValidatorOwnable {
    SignerGetter public signerGetter;

    modifier onlyValidator() {
        require(signerGetter.getSignerAddresses() == msg.sender, "ValidatorOwnable: only validator");
        _;
    }

    function _setSignerGetter(address _signerGetterAddress) internal virtual {
        signerGetter = SignerGetter(_signerGetterAddress);
    }
}
