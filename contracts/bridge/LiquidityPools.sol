// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "./TokenManager.sol";
import "./Bridge.sol";
import "./Globals.sol";

contract LiquidityPools is Initializable, Ownable {
    struct LiquidityPosition {
        uint256 balance;
        uint256 lastRewardPoints;
    }

    mapping(address => uint256) public providedLiquidity;
    mapping(address => uint256) public availableLiquidity;
    mapping(address => mapping(address => LiquidityPosition)) public liquidityPositions;
    mapping(address => uint256) public collectedFees;
    mapping(address => uint256) public totalRewardPoints;
    TokenManager public tokenManager;
    Bridge public bridge;
    Globals public globals;
    uint256 public feePercentage;

    event LiquidityAdded(address token, address account, uint256 amount);
    event LiquidityRemoved(address token, address account, uint256 amount);
    event TokenManagerUpdated(address tokenManager);
    event FeePercentageUpdated(uint256 feePercentage);

    modifier onlyBridge() {
        require(msg.sender == address(bridge), "LiquidityPools: only bridge");
        _;
    }

    function initialize(
        address _tokenManager,
        address _bridge,
        address _globals,
        uint256 _feePercentage
    ) external initializer {
        tokenManager = TokenManager(_tokenManager);
        bridge = Bridge(_bridge);
        globals = Globals(_globals);
        feePercentage = _feePercentage;
    }

    function setTokenManager(address _tokenManager) external onlyOwner {
        tokenManager = TokenManager(_tokenManager);
        emit TokenManagerUpdated(_tokenManager);
    }

    function setFeePercentage(uint256 _feePercentage) external onlyOwner {
        feePercentage = _feePercentage;
        emit FeePercentageUpdated(_feePercentage);
    }

    function transfer(
        address _token,
        address _receiver,
        uint256 _fee,
        uint256 _transferAmount
    ) external onlyBridge {
        require(tokenManager.isTokenSupported(_token), "TokenManager: token is not supported");
        require(
            IERC20(_token).balanceOf(address(this)) >= _transferAmount + _fee,
            "IERC20: amount more than contract balance"
        );

        distributeFee(_token, _fee);

        require(ERC20(_token).transfer(_receiver, _transferAmount), "ERC20: transfer failed");
    }

    function addLiquidity(address _token, uint256 _amount) public {
        claimRewards(_token);
        require(tokenManager.isTokenSupported(_token), "TokenManager: token is not supported");
        require(IERC20(_token).transferFrom(msg.sender, address(this), _amount), "IERC20: transfer failed");

        providedLiquidity[_token] += _amount;
        availableLiquidity[_token] += _amount;
        liquidityPositions[_token][msg.sender].balance += _amount;

        emit LiquidityAdded(_token, msg.sender, _amount);
    }

    function removeLiquidity(address _token, uint256 _amount) public {
        claimRewards(_token);
        require(tokenManager.isTokenSupported(_token), "TokenManager: token is not supported");
        require(IERC20(_token).balanceOf(address(this)) >= _amount, "IERC20: amount more than contract balance");
        require(liquidityPositions[_token][msg.sender].balance >= _amount, "LiquidityPools: too much amount");

        providedLiquidity[_token] -= _amount;
        availableLiquidity[_token] -= _amount;
        liquidityPositions[_token][msg.sender].balance -= _amount;

        require(IERC20(_token).transfer(msg.sender, _amount), "IERC20: transfer failed");

        emit LiquidityRemoved(_token, msg.sender, _amount);
    }

    function claimRewards(address _token) public {
        uint256 rewardsOwingAmount = rewardsOwing(_token);
        if (rewardsOwingAmount > 0) {
            collectedFees[_token] -= rewardsOwingAmount;
            liquidityPositions[_token][msg.sender].balance += rewardsOwingAmount;
            liquidityPositions[_token][msg.sender].lastRewardPoints = totalRewardPoints[_token];
        }
    }

    function rewardsOwing(address _token) public view returns (uint256) {
        uint256 newRewardPoints = totalRewardPoints[_token] - liquidityPositions[_token][msg.sender].lastRewardPoints;
        return (liquidityPositions[_token][msg.sender].balance * newRewardPoints) / globals.BASE_DIVISOR();
    }

    function distributeFee(address _token, uint256 _amount) private {
        require(_amount > 0, "LiquidityPools: amount must be greater than zero");
        totalRewardPoints[_token] += (_amount * globals.BASE_DIVISOR()) / providedLiquidity[_token];
        providedLiquidity[_token] += _amount;
        collectedFees[_token] += _amount;
    }
}
