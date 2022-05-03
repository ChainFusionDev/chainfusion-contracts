// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./DKG.sol";

contract ThresholdSigner is Ownable, Initializable {
    DKG public dkg;

    mapping(uint256 => address) public signerAddresses;
    uint256 public activeGeneration;

    event SignerAddressUpdated(uint256 generation, address signerAddress);
    event DKGUpdated(address dkg);

    modifier onlySigner() {
        require(msg.sender == activeAddress(), "Signer: only active collective signer allowed");
        _;
    }

    modifier onlyDKG() {
        require(msg.sender == address(dkg), "Signer: only DKG allowed");
        _;
    }

    function initialize(address _dkg, address _signerAddress) external initializer {
        dkg = DKG(_dkg);
        _updateSignerAddress(0, _signerAddress);
    }

    function setDKG(address _dkg) external onlyOwner {
        dkg = DKG(_dkg);
        emit DKGUpdated(_dkg);
    }

    function updateSignerAddress(uint256 _generation, address _signerAddress) external onlyDKG {
        if (signerAddresses[_generation] == address(0)) {
            _updateSignerAddress(_generation, _signerAddress);
        }
    }

    function activeAddress() public view returns (address) {
        return signerAddresses[activeGeneration];
    }

    function _updateSignerAddress(uint256 _generation, address _signerAddress) private {
        if (_generation > activeGeneration) {
            activeGeneration = _generation;
        }

        signerAddresses[_generation] = _signerAddress;
        emit SignerAddressUpdated(_generation, _signerAddress);
    }
}
