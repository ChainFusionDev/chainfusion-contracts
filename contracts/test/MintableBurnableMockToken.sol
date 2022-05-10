// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "../interfaces/IERC20MintableBurnable.sol";
import "hardhat/console.sol";

contract MintableBurnableMockToken is ERC20, IERC20MintableBurnable, ERC20Burnable, Ownable {
    constructor(string memory name, string memory symbol)
        // solhint-disable-next-line no-empty-blocks
        ERC20(name, symbol)
    {}

    function mint(address to, uint256 amount) public onlyOwner {
        _mint(to, amount);
    }

    function burn(uint256 amount) public override(IERC20MintableBurnable, ERC20Burnable) {
        ERC20Burnable.burn(amount);
    }

    function burnFrom(address account, uint256 amount) public override(IERC20MintableBurnable, ERC20Burnable) {
        ERC20Burnable.burnFrom(account, amount);
    }
}
