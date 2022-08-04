// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";

contract BridgeApp is Ownable {
    mapping(uint256 => address) public contractAddresses;

    event ContractAddressUpdated(uint256 chainId, address contractAddress);

    constructor(address _owner) {
        _transferOwnership(_owner);
    }

    function setContractAddress(uint256 chainId, address contractAddress) public onlyOwner {
        contractAddresses[chainId] = contractAddress;
        emit ContractAddressUpdated(chainId, contractAddress);
    }
}
