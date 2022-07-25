// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./ValidatorOwnable.sol";
import "./LiquidityPools.sol";
import "./Globals.sol";

contract FeeManager is Initializable, ValidatorOwnable {
    LiquidityPools public liquidityPools;
    address public foundationAddress;

    uint256 public validatorRefundFee;
    mapping(address => uint256) public tokenFeePercentage;
    mapping(address => uint256) public validatorRewardPercentage;
    mapping(address => uint256) public liquidityRewardPercentage;
    mapping(address => uint256) public foundationRewardPercentage;

    event LiquidityPoolsUpdated(address _liquidityPools);
    event FoundationAddressUpdated(address _foundationAddress);
    event ValidatorRefundFeeUpdated(uint256 _validatorRefundFee);

    // solhint-disable-next-line no-empty-blocks
    receive() external payable {}

    function initialize(
        address _validatorStorage,
        address payable _liquidityPools,
        address _foundationAddress,
        uint256 _validatorRefundFee
    ) external initializer {
        _setValidatorStorage(_validatorStorage);
        setLiquidityPools(_liquidityPools);
        setFoundationAddress(_foundationAddress);
        setValidatorRefundFee(_validatorRefundFee);
    }

    function setLiquidityPools(address payable _liquidityPools) public onlyValidator {
        liquidityPools = LiquidityPools(_liquidityPools);
        emit LiquidityPoolsUpdated(_liquidityPools);
    }

    function setFoundationAddress(address _foundationAddress) public onlyValidator {
        foundationAddress = _foundationAddress;
        emit FoundationAddressUpdated(_foundationAddress);
    }

    function setValidatorRefundFee(uint256 _validatorRefundFee) public onlyValidator {
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

        address validatorAddress = validatorStorage.getAddress();

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
