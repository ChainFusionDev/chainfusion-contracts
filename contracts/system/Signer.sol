// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./DKG.sol";

contract Signer is Ownable, Initializable {
    DKG public dkg;

    mapping(uint256 => address) public collectiveSigners;
    uint256 public activeGeneration;

    event CollectiveSignerUpdated(uint256 generation, address collectiveSigner);
    event DKGUpdated(address dkg);

    modifier onlySigner() {
        require(msg.sender == activeSigner(), "Signer: only active collective signer allowed");
        _;
    }

    modifier onlyDKG() {
        require(msg.sender == address(dkg), "Signer: only DKG allowed");
        _;
    }

    function initialize(address _dkg, address _collectiveSigner) external initializer {
        dkg = DKG(_dkg);
        _updateCollectiveSigner(0, _collectiveSigner);
    }

    function setDKG(address _dkg) external onlyOwner {
        dkg = DKG(_dkg);
        emit DKGUpdated(_dkg);
    }

    function updateCollectiveSigner(uint256 _generation, address _collectiveSigner) external onlyDKG {
        if (collectiveSigners[_generation] == address(0)) {
            _updateCollectiveSigner(_generation, _collectiveSigner);
        }
    }

    function activeSigner() public view returns (address) {
        return collectiveSigners[activeGeneration];
    }

    function _updateCollectiveSigner(uint256 _generation, address _collectiveSigner) private {
        if (_generation > activeGeneration) {
            activeGeneration = _generation;
        }

        collectiveSigners[_generation] = _collectiveSigner;
        emit CollectiveSignerUpdated(_generation, _collectiveSigner);
    }
}
