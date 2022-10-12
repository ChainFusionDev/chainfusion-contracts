// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "../interfaces/IERC20MintableBurnable.sol";

contract ERC20Bridgeable is ERC20, IERC20MintableBurnable, ERC20Burnable, Ownable {
    constructor(string memory name, string memory symbol)
        ERC20(name, symbol)
    // solhint-disable-next-line no-empty-blocks
    {

    }

    function mint(address to, uint256 amount) public onlyOwner {
        _mint(to, amount);
    }

    function burnFrom(address account, uint256 amount) public override(IERC20MintableBurnable, ERC20Burnable) {
        ERC20Burnable.burnFrom(account, amount);
    }
}
