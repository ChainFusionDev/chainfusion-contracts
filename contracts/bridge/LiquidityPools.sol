// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "./TokenManager.sol";
import "./Bridge.sol";
import "./FeeManager.sol";
import "./Globals.sol";

contract LiquidityPools is Initializable, Ownable {
    struct LiquidityPosition {
        uint256 balance;
        uint256 lastRewardPoints;
    }

    TokenManager public tokenManager;
    Bridge public bridge;
    FeeManager public feeManager;

    uint256 public feePercentage;

    mapping(address => uint256) public providedLiquidity;
    mapping(address => uint256) public availableLiquidity;
    mapping(address => mapping(address => LiquidityPosition)) public liquidityPositions;
    mapping(address => uint256) public collectedFees;
    mapping(address => uint256) public totalRewardPoints;

    event TokenManagerUpdated(address tokenManager);
    event BridgeUpdated(address bridge);
    event FeeManagerUpdated(address feeManager);
    event FeePercentageUpdated(uint256 feePercentage);

    event LiquidityAdded(address token, address account, uint256 amount);
    event LiquidityRemoved(address token, address account, uint256 amount);

    modifier onlyBridge() {
        require(msg.sender == address(bridge), "LiquidityPools: only bridge");
        _;
    }

    modifier onlyFeeManager() {
        require(msg.sender == address(feeManager), "LiquidityPools: only feeManager");
        _;
    }

    // solhint-disable-next-line no-empty-blocks
    receive() external payable {}

    function initialize(
        address _tokenManager,
        address _bridge,
        address payable _feeManager,
        uint256 _feePercentage
    ) external initializer {
        setTokenManager(_tokenManager);
        setBridge(_bridge);
        setFeeManager(_feeManager);
        setFeePercentage(_feePercentage);
    }

    function transfer(
        address _token,
        address _receiver,
        uint256 _transferAmount
    ) external onlyBridge {
        require(tokenManager.isTokenEnabled(_token), "TokenManager: token is not supported");
        require(
            IERC20(_token).balanceOf(address(this)) >= _transferAmount,
            "IERC20: amount more than contract balance"
        );
        require(ERC20(_token).transfer(_receiver, _transferAmount), "ERC20: transfer failed");
    }

    function distributeFee(address _token, uint256 _amount) external onlyFeeManager {
        require(_amount > 0, "LiquidityPools: amount must be greater than zero");
        totalRewardPoints[_token] += (_amount * BASE_DIVISOR) / providedLiquidity[_token];
        providedLiquidity[_token] += _amount;
        collectedFees[_token] += _amount;
    }

    function setTokenManager(address _tokenManager) public onlyOwner {
        tokenManager = TokenManager(_tokenManager);
        emit TokenManagerUpdated(_tokenManager);
    }

    function setBridge(address _bridge) public onlyOwner {
        bridge = Bridge(_bridge);
        emit BridgeUpdated(_bridge);
    }

    function setFeeManager(address payable _feeManager) public onlyOwner {
        feeManager = FeeManager(_feeManager);
        emit FeeManagerUpdated(_feeManager);
    }

    function setFeePercentage(uint256 _feePercentage) public onlyOwner {
        feePercentage = _feePercentage;
        emit FeePercentageUpdated(_feePercentage);
    }

    function addLiquidity(address _token, uint256 _amount) public {
        claimRewards(_token);

        require(tokenManager.isTokenEnabled(_token), "TokenManager: token is not supported");
        require(IERC20(_token).transferFrom(msg.sender, address(this), _amount), "IERC20: transfer failed");

        _addLiquidity(_token, _amount);
    }

    function removeLiquidity(address _token, uint256 _amount) public payable {
        claimRewards(_token);

        require(tokenManager.isTokenEnabled(_token), "TokenManager: token is not supported");
        require(liquidityPositions[_token][msg.sender].balance >= _amount, "LiquidityPools: too much amount");

        _removeLiquidity(_token, _amount);

        if (_token == NATIVE_TOKEN) {
            // solhint-disable-next-line avoid-low-level-calls
            (bool success, ) = msg.sender.call{value: _amount, gas: 21000}("");
            require(success, "LiquidityPools: transfer native token failed");
        } else {
            require(IERC20(_token).balanceOf(address(this)) >= _amount, "IERC20: amount more than contract balance");
            require(IERC20(_token).transfer(msg.sender, _amount), "IERC20: transfer failed");
        }
    }

    function claimRewards(address _token) public {
        uint256 rewardsOwingAmount = rewardsOwing(_token);
        if (rewardsOwingAmount > 0) {
            collectedFees[_token] -= rewardsOwingAmount;
            liquidityPositions[_token][msg.sender].balance += rewardsOwingAmount;
            liquidityPositions[_token][msg.sender].lastRewardPoints = totalRewardPoints[_token];
        }
    }

    function transferNative(address _receiver, uint256 _amount) public payable onlyBridge {
        require(tokenManager.isTokenEnabled(NATIVE_TOKEN), "TokenManager: token is not supported");

        // solhint-disable-next-line avoid-low-level-calls
        (bool success, ) = _receiver.call{value: _amount, gas: 21000}("");
        require(success, "LiquidityPools: transfer native token failed");
    }

    function addNativeLiquidity() public payable {
        claimRewards(NATIVE_TOKEN);

        _addLiquidity(NATIVE_TOKEN, msg.value);
    }

    function rewardsOwing(address _token) public view returns (uint256) {
        uint256 newRewardPoints = totalRewardPoints[_token] - liquidityPositions[_token][msg.sender].lastRewardPoints;
        return (liquidityPositions[_token][msg.sender].balance * newRewardPoints) / BASE_DIVISOR;
    }

    function _addLiquidity(address _token, uint256 _amount) private {
        providedLiquidity[_token] += _amount;
        availableLiquidity[_token] += _amount;
        liquidityPositions[_token][msg.sender].balance += _amount;

        emit LiquidityAdded(_token, msg.sender, _amount);
    }

    function _removeLiquidity(address _token, uint256 _amount) private {
        providedLiquidity[_token] -= _amount;
        availableLiquidity[_token] -= _amount;
        liquidityPositions[_token][msg.sender].balance -= _amount;

        emit LiquidityRemoved(_token, msg.sender, _amount);
    }
}
