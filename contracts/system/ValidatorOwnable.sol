// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./DKG.sol";

abstract contract ValidatorOwnable {
    DKG public dkgAddress;

    modifier onlyValidator() {
        require(dkgAddress.getSignerAddresses() == msg.sender, "ValidatorOwnable: only validator");
        _;
    }

    function _setDKG(address _signerAddress) internal virtual {
        dkgAddress = DKG(_signerAddress);
    }
}
