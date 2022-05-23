// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./DKG.sol";

struct SignatureData {
    mapping(address => bytes) rounds;
    bytes digest;
    bytes signature;
}

contract ThresholdSigner is Ownable, Initializable {
    DKG public dkg;

    mapping(address => mapping(uint256 => SignatureData)) public signatures;
    mapping(address => uint256) public preparedSignatures;
    mapping(address => mapping(bytes => uint256)) public digestToSignature;

    event DKGUpdated(address dkg);

    function initialize(address _dkg) external initializer {
        dkg = DKG(_dkg);
    }

    function setDKG(address _dkg) external onlyOwner {
        dkg = DKG(_dkg);
        emit DKGUpdated(_dkg);
    }
}
