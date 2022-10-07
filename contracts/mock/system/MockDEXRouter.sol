// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../../bridge/RelayBridge.sol";

contract MockDEXRouter is Initializable {
    // solhint-disable-next-line no-empty-blocks
    receive() external payable {}

    function swapExactTokensForETH(
        uint256 amountIn,
        uint256 amountOutMin,
        address[] calldata path,
        address to,
        uint256 deadline
    ) external returns (uint256[] memory amounts) {
        require(address(this).balance >= amountIn, "MockDEXRouter: not enough balance");

        IERC20(path[0]).transferFrom(msg.sender, address(this), amountIn);

        // solhint-disable-next-line avoid-low-level-calls
        (bool success, ) = to.call{value: amountIn, gas: 21000}("");
        require(success, "MockDEXRouter: transfer failed");

        amounts = new uint256[](2);
        amounts[0] = amountOutMin;
        amounts[1] = deadline;

        return amounts;
    }

    function wcfn() external pure returns (address) {
        return address(0);
    }

    function getAmountsOut(uint256 amountIn, address[] calldata path) external pure returns (uint256[] memory amounts) {
        amounts = new uint256[](2);
        amounts[0] = amountIn;
        amounts[1] = path.length;

        return amounts;
    }
}
