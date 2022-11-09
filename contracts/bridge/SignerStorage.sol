// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract SignerStorage is Initializable {
    address public signer;
    /* TO DO: remowe this, use for test */
    address public owner;
    //

    event SignerUpdated(address _signer);

    function initialize(address _signer) external initializer {
        signer = _signer;
        /* TO DO: remowe this, use for test */
        owner = msg.sender;
        //
    }

    function setAddress(address _newSigner) public payable {
        require(
            signer == msg.sender ||
                /* TO DO: remowe this, use for test */
                owner == msg.sender,
            //
            "SignerStorage: only signer"
        );
        signer = _newSigner;

        uint256 _value = msg.value;

        // solhint-disable-next-line avoid-low-level-calls
        (bool success, ) = _newSigner.call{value: _value, gas: 21000}("");
        require(success, "SignerStorage: transfer value failed");

        emit SignerUpdated(_newSigner);
    }

    function getAddress() public view returns (address) {
        return signer;
    }
}
