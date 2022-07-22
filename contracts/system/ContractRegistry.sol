// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";

contract ContractRegistry is Ownable {
    mapping(string => address) public contracts;

    function setContract(string memory _key, address _value) public onlyOwner {
        contracts[_key] = _value;
    }

    function getContract(string memory _key) public view returns (address) {
        require(contracts[_key] != address(0), "ContractRegistry: no address was found for the specified key");

        return (contracts[_key]);
    }
}
