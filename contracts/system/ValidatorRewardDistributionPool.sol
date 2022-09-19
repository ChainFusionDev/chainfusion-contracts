// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./ContractRegistry.sol";
import "./Staking.sol";
import "./ContractKeys.sol";
import "../bridge/Globals.sol";
import "hardhat/console.sol";

contract ValidatorRewardDistributionPool is Initializable, ContractKeys {
    struct RewardPosition {
        uint256 stake;
        uint256 reward;
        uint256 lastRewardPoints;
    }

    ContractRegistry public contractRegistry;

    uint256 public totalRewardPoints;
    uint256 public providedStake;

    mapping(address => RewardPosition) public rewardPositions;

    event CollectRewards(address _validator, uint256 amount);

    // solhint-disable-next-line no-empty-blocks
    receive() external payable {}

    function initialize(address _contractRegistry) external initializer {
        contractRegistry = ContractRegistry(_contractRegistry);
    }

    function distributeRewards(uint256 _amount) public {
        require(_amount > 0, "ValidatorRewardDistributionPool: amount must be greater than 0");

        providedStake = _stakingContract().getTotalStake();
        totalRewardPoints = (_amount * BASE_DIVISOR) / providedStake;

        _updateValidatorStake();
        _distributeRewards();
    }

    function collectRewards() public {
        _sendRewards(msg.sender);
    }

    function reinvestRewards() public {
        uint256 reward = rewardPositions[msg.sender].reward;

        _sendRewards(address(_stakingContract()));
        _stakingContract().addRewardsToStake(msg.sender, reward);
    }

    function rewardsOwing(address validator) public view returns (uint256) {
        return (rewardPositions[validator].stake * totalRewardPoints) / BASE_DIVISOR;
    }

    function _updateValidatorStake() private {
        address[] memory validators = _stakingContract().getValidators();
        for (uint256 i = 0; i < validators.length; i++) {
            uint256 stake = _stakingContract().getStake(validators[i]);
            rewardPositions[validators[i]].stake = stake;
        }
    }

    function _distributeRewards() private {
        address[] memory validators = _stakingContract().getValidators();
        for (uint256 i = 0; i < validators.length; i++) {
            uint256 rewardsOwingAmount = rewardsOwing(validators[i]);
            rewardPositions[validators[i]].reward += rewardsOwingAmount;
        }
    }

    function _sendRewards(address _receiver) private {
        uint256 reward = rewardPositions[msg.sender].reward;

        require(reward > 0, "ValidatorRewardDistributionPool: reward must be greater than 0");

        rewardPositions[msg.sender].reward -= reward;

        // solhint-disable-next-line avoid-low-level-calls
        (bool success, ) = _receiver.call{value: reward, gas: 21000}("");
        require(success, "ValidatorRewardDistributionPool: transfer failed");

        emit CollectRewards(msg.sender, reward);
    }

    function _stakingContract() private view returns (Staking) {
        return Staking(payable(contractRegistry.getContract(STAKING_KEY)));
    }
}
