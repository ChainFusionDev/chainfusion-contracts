// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "./SignerOwnable.sol";
import "./TokenManager.sol";
import "./Bridge.sol";
import "./FeeManager.sol";
import "./Globals.sol";

contract LiquidityPools is Initializable, SignerOwnable {
    struct LiquidityPosition {
        uint256 balance;
        uint256 lastRewardPoints;
    }

    TokenManager public tokenManager;
    Bridge public bridge;
    FeeManager public feeManager;

    uint256 public feePercentage;

    mapping(address => uint256) public providedLiquidity;
    mapping(address => address[]) public liquidityProviders;

    mapping(address => uint256) public availableLiquidity;
    mapping(address => mapping(address => LiquidityPosition)) public liquidityPositions;
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
        address _signerStorage,
        address _tokenManager,
        address _bridge,
        address payable _feeManager,
        uint256 _feePercentage
    ) external initializer {
        _setSignerStorage(_signerStorage);
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
        require(
            IERC20(_token).balanceOf(address(this)) >= _transferAmount,
            "IERC20: amount more than contract balance"
        );
        require(ERC20(_token).transfer(_receiver, _transferAmount), "ERC20: transfer failed");
    }

    function transferNative(address _receiver, uint256 _amount) external onlyBridge {
        // solhint-disable-next-line avoid-low-level-calls
        (bool success, ) = _receiver.call{value: _amount, gas: 21000}("");
        require(success, "LiquidityPools: transfer native token failed");
    }

    function distributeFee(address _token, uint256 _amount) external onlyFeeManager {
        require(_amount > 0, "LiquidityPools: amount must be greater than zero");
        totalRewardPoints[_token] = (_amount * BASE_DIVISOR) / providedLiquidity[_token];
        providedLiquidity[_token] += _amount;
        availableLiquidity[_token] += _amount;

        _distributeFee(_token);
    }

    function setTokenManager(address _tokenManager) public onlySigner {
        tokenManager = TokenManager(_tokenManager);
        emit TokenManagerUpdated(_tokenManager);
    }

    function setBridge(address _bridge) public onlySigner {
        bridge = Bridge(_bridge);
        emit BridgeUpdated(_bridge);
    }

    function setFeeManager(address payable _feeManager) public onlySigner {
        feeManager = FeeManager(_feeManager);
        emit FeeManagerUpdated(_feeManager);
    }

    function setFeePercentage(uint256 _feePercentage) public onlySigner {
        feePercentage = _feePercentage;
        emit FeePercentageUpdated(_feePercentage);
    }

    function addLiquidity(address _token, uint256 _amount) public {
        require(tokenManager.isTokenEnabled(_token), "TokenManager: token is not supported");
        require(IERC20(_token).transferFrom(msg.sender, address(this), _amount), "IERC20: transfer failed");

        _addLiquidity(_token, _amount);
    }

    function removeLiquidity(address _token, uint256 _amount) public payable {
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

    function addNativeLiquidity() public payable {
        _addLiquidity(NATIVE_TOKEN, msg.value);
    }

    function rewardsOwing(address _token, address _provider) public view returns (uint256) {
        return (liquidityPositions[_token][_provider].balance * totalRewardPoints[_token]) / BASE_DIVISOR;
    }

    function _distributeFee(address _token) private {
        address[] memory providers = liquidityProviders[_token];
        for (uint256 i = 0; i < providers.length; i++) {
            uint256 rewardsOwingAmount = rewardsOwing(_token, providers[i]);
            liquidityPositions[_token][providers[i]].balance += rewardsOwingAmount;
        }
    }

    function _addLiquidity(address _token, uint256 _amount) private {
        providedLiquidity[_token] += _amount;
        availableLiquidity[_token] += _amount;
        liquidityPositions[_token][msg.sender].balance += _amount;
        liquidityProviders[_token].push(msg.sender);

        emit LiquidityAdded(_token, msg.sender, _amount);
    }

    function _removeLiquidity(address _token, uint256 _amount) private {
        providedLiquidity[_token] -= _amount;
        availableLiquidity[_token] -= _amount;

        if (liquidityPositions[_token][msg.sender].balance - _amount == 0) {
            for (uint256 i = 0; i < liquidityProviders[_token].length; i++) {
                if (liquidityProviders[_token][i] == msg.sender) {
                    delete (liquidityProviders[_token][i]);
                }
            }
        }

        liquidityPositions[_token][msg.sender].balance -= _amount;

        emit LiquidityRemoved(_token, msg.sender, _amount);
    }
}
