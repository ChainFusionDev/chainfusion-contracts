// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./TokenManager.sol";
import "./Bridge.sol";

contract LiquidityPools is Initializable, Ownable {
    mapping(address => uint256) public providedLiquidity;
    mapping(address => uint256) public availableLiquidity;
    mapping(address => mapping(address => uint256)) public liquidityPositions;
    TokenManager public tokenManager;

    event LiquidityAdded(address token, address account, uint256 amount);

    function initialize(address _tokenManager) external initializer {
        tokenManager = TokenManager(_tokenManager);
    }

    function setTokenManager(address _tokenManager) external onlyOwner {
        tokenManager = TokenManager(_tokenManager);
    }

    function addLiquidity(address _token, uint256 _amount) public {
        require(tokenManager.isTokenSupported(_token), "Token is not supported");
        require(IERC20(_token).transferFrom(msg.sender, address(this), _amount), "Transfer failed");

        providedLiquidity[_token] += _amount;
        availableLiquidity[_token] += _amount;
        liquidityPositions[_token][msg.sender] += _amount;

        emit LiquidityAdded(_token, msg.sender, _amount);
    }

    function removeLiquidity(address _token, uint256 _amount) public {
        require(tokenManager.isTokenSupported(_token), "Token is not supported");
        require(liquidityPositions[_token][msg.sender] >= _amount, "Too much amount");

        providedLiquidity[_token] -= _amount;
        availableLiquidity[_token] -= _amount;
        liquidityPositions[_token][msg.sender] -= _amount;

        require(IERC20(_token).transfer(msg.sender, _amount), "Transfer failed");

        emit LiquidityAdded(_token, msg.sender, _amount);
    }
}
