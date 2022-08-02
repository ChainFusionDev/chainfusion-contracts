// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract SignerStorage is Initializable {
    address public signer;

    event SignerUpdated(address _signer);

    function initialize(address _signer) external initializer {
        signer = _signer;
    }

    function setAddress(address _newSigner) public {
        require(signer == msg.sender, "SignerStorage: only signer");
        signer = _newSigner;
        emit SignerUpdated(_newSigner);
    }

    function getAddress() public view returns (address) {
        return signer;
    }
}
