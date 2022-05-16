// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./LiquidityPools.sol";
import "./Globals.sol";

contract FeeManager is Initializable, Ownable {
    mapping(address => uint256) public tokenFeePercentage;
    mapping(address => uint256) public validatorRewardPercentage;
    mapping(address => uint256) public liquidityRewardPercentage;
    mapping(address => uint256) public foundationRewardPercentage;

    LiquidityPools public liquidityPools;
    address public validatorAddress;
    address public foundationAddress;
    uint256 public validatorRefundFee;

    event LiquidityPoolsUpdated(address _liquidityPools);
    event ValidatorAddressUpdated(address _validatorAddress);
    event FoundationAddressUpdated(address _foundationAddress);
    event ValidatorRefundFeeUpdated(uint256 _validatorRefundFee);

    // solhint-disable-next-line no-empty-blocks
    receive() external payable {}

    function initialize(
        address payable _liquidityPools,
        address _validatorAddress,
        address _foundationAddress,
        uint256 _validatorRefundFee
    ) external initializer {
        liquidityPools = LiquidityPools(_liquidityPools);
        validatorAddress = _validatorAddress;
        foundationAddress = _foundationAddress;
        validatorRefundFee = _validatorRefundFee;
    }

    function setLiquidityPools(address payable _liquidityPools) external onlyOwner {
        liquidityPools = LiquidityPools(_liquidityPools);
        emit LiquidityPoolsUpdated(_liquidityPools);
    }

    function setValidatorAddress(address _validatorAddress) external onlyOwner {
        validatorAddress = _validatorAddress;
        emit ValidatorAddressUpdated(_validatorAddress);
    }

    function setFoundationAddress(address _foundationAddress) external onlyOwner {
        foundationAddress = _foundationAddress;
        emit FoundationAddressUpdated(_foundationAddress);
    }

    function setValidatorRefundFee(uint256 _validatorRefundFee) external onlyOwner {
        validatorRefundFee = _validatorRefundFee;
        emit ValidatorRefundFeeUpdated(_validatorRefundFee);
    }

    function setTokenFee(
        address token,
        uint256 tokenFee,
        uint256 validatorReward,
        uint256 liquidityReward,
        uint256 foundationReward
    ) public {
        tokenFeePercentage[token] = tokenFee;
        validatorRewardPercentage[token] = validatorReward;
        liquidityRewardPercentage[token] = liquidityReward;
        foundationRewardPercentage[token] = foundationReward;
    }

    function distributeRewards(address token) public {
        uint256 totalRewards;
        uint256 validatorRewards;
        uint256 liquidityRewards;
        uint256 foundationRewards;

        if (token == NATIVE_TOKEN) {
            totalRewards = address(this).balance;

            (validatorRewards, liquidityRewards, foundationRewards) = _calculateRewards(token, totalRewards);

            // solhint-disable-next-line avoid-low-level-calls
            (bool success, ) = validatorAddress.call{value: validatorRewards, gas: 21000}("");
            require(success, "FeeManager: transfer native token failed");

            // solhint-disable-next-line avoid-low-level-calls
            (success, ) = address(liquidityPools).call{value: liquidityRewards, gas: 21000}("");
            require(success, "FeeManager: transfer native token failed");

            // solhint-disable-next-line avoid-low-level-calls
            (success, ) = foundationAddress.call{value: foundationRewards, gas: 21000}("");
            require(success, "FeeManager: transfer native token failed");
        } else {
            totalRewards = IERC20(token).balanceOf(address(this));
            (validatorRewards, liquidityRewards, foundationRewards) = _calculateRewards(token, totalRewards);

            require(IERC20(token).transfer(validatorAddress, validatorRewards), "IERC20: transfer failed");
            require(IERC20(token).transfer(address(liquidityPools), liquidityRewards), "IERC20: transfer failed");
            require(IERC20(token).transfer(foundationAddress, foundationRewards), "IERC20: transfer failed");
        }

        liquidityPools.distributeFee(token, liquidityRewards);
    }

    function calculateFee(address token, uint256 amount) public view returns (uint256 fee) {
        fee = validatorRefundFee + (tokenFeePercentage[token] * amount) / BASE_DIVISOR;
        require(fee <= amount, "FeeManager: fee to be less than or equal to amount");

        return fee;
    }

    function _calculateRewards(address token, uint256 totalRewards)
        private
        view
        returns (
            uint256 validatorRewards,
            uint256 liquidityRewards,
            uint256 foundationRewards
        )
    {
        validatorRewards = (validatorRewardPercentage[token] * totalRewards) / BASE_DIVISOR;
        liquidityRewards = (liquidityRewardPercentage[token] * totalRewards) / BASE_DIVISOR;
        foundationRewards = totalRewards - validatorRewards - liquidityRewards;

        return (validatorRewards, liquidityRewards, foundationRewards);
    }
}
