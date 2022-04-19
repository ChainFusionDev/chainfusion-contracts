// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./TokenManager.sol";
import "./Bridge.sol";

contract LiquidityPools is Initializable, Ownable {
    mapping(address => uint256) public providedLiquidity;
    mapping(address => uint256) public availableLiquidity;
    mapping(address => mapping(address => uint256)) public liquidityPositions;
    TokenManager public tokenManager;
    Bridge public bridge;

    event LiquidityAdded(address token, address account, uint256 amount);
    event LiquidityRemoved(address token, address account, uint256 amount);

    modifier onlyBridge() {
        require(msg.sender == address(bridge), "LiquidityPools: only bridge");
        _;
    }

    function initialize(address _tokenManager, address _bridge) external initializer {
        tokenManager = TokenManager(_tokenManager);
        bridge = Bridge(_bridge);
    }

    function setTokenManager(address _tokenManager) external onlyOwner {
        tokenManager = TokenManager(_tokenManager);
    }

    function transfer(
        address _token,
        address _receiver,
        uint256 _amount
    ) external onlyBridge {
        // solhint-disable-next-line reason-string
        require(tokenManager.isTokenSupported(_token), "TokenManager: token is not supported");
        // solhint-disable-next-line reason-string
        require(IERC20(_token).balanceOf(address(this)) >= _amount, "IERC20: amount more than contract balance");
        require(IERC20(_token).transfer(_receiver, _amount), "IERC20: transfer failed");
    }

    function addLiquidity(address _token, uint256 _amount) public {
        // solhint-disable-next-line reason-string
        require(tokenManager.isTokenSupported(_token), "TokenManager: token is not supported");
        require(IERC20(_token).transferFrom(msg.sender, address(this), _amount), "IERC20: transfer failed");

        providedLiquidity[_token] += _amount;
        availableLiquidity[_token] += _amount;
        liquidityPositions[_token][msg.sender] += _amount;

        emit LiquidityAdded(_token, msg.sender, _amount);
    }

    function removeLiquidity(address _token, uint256 _amount) public {
        // solhint-disable-next-line reason-string
        require(tokenManager.isTokenSupported(_token), "TokenManager: token is not supported");
        // solhint-disable-next-line reason-string
        require(IERC20(_token).balanceOf(address(this)) >= _amount, "IERC20: amount more than contract balance");
        require(liquidityPositions[_token][msg.sender] >= _amount, "LiquidityPools: too much amount");

        providedLiquidity[_token] -= _amount;
        availableLiquidity[_token] -= _amount;
        liquidityPositions[_token][msg.sender] -= _amount;

        require(IERC20(_token).transfer(msg.sender, _amount), "IERC20: transfer failed");

        emit LiquidityRemoved(_token, msg.sender, _amount);
    }
}
