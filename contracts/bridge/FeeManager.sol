// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./LiquidityPools.sol";
import "./Globals.sol";
import "hardhat/console.sol";

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
        uint256 totalRewards = IERC20(token).balanceOf(address(this));
        uint256 validatorRewards = (validatorRewardPercentage[token] * totalRewards) / BASE_DIVISOR;
        uint256 liquidityRewards = (liquidityRewardPercentage[token] * totalRewards) / BASE_DIVISOR;
        uint256 foundationReward = totalRewards - validatorRewards - liquidityRewards;

        require(IERC20(token).transfer(validatorAddress, validatorRewards), "IERC20: transfer failed");
        require(IERC20(token).transfer(address(liquidityPools), liquidityRewards), "IERC20: transfer failed");
        require(IERC20(token).transfer(foundationAddress, foundationReward), "IERC20: transfer failed");

        liquidityPools.distributeFee(token, liquidityRewards);
    }

    function calculateFee(address token, uint256 amount) public view returns (uint256 fee) {
        fee = validatorRefundFee + (tokenFeePercentage[token] * amount) / BASE_DIVISOR;
        console.log(fee);
        require(fee <= amount, "FeeManager: fee to be less than or equal to amount");

        return fee;
    }
}
