// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./SignerOwnable.sol";

contract ContractRegistry is SignerOwnable, Initializable {
    mapping(string => address) public contracts;

    function initialize(address _signerGetterAddress) external initializer {
        _setSignerGetter(_signerGetterAddress);
    }

    function setContract(string memory _key, address _value) public onlySigner {
        contracts[_key] = _value;
    }

    function getContract(string memory _key) public view returns (address) {
        require(contracts[_key] != address(0), "ContractRegistry: no address was found for the specified key");

        return (contracts[_key]);
    }
}
