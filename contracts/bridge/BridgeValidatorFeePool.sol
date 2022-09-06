// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./SignerOwnable.sol";
import "./Bridge.sol";
import "./Globals.sol";

contract BridgeValidatorFeePool is Initializable, SignerOwnable {
    Bridge public bridge;

    mapping(address => uint256) public limitPerToken;

    address public validatorFeeReceiver;

    event ValidatorFeeReceiverUpdated(address validatorFeeReceiver);
    event LimitPerTokenUpdated(address token, uint256 limit);
    event BridgeUpdated(address bridge);
    event Collected(address token, uint256 amount);

    // solhint-disable-next-line no-empty-blocks
    receive() external payable {}

    function initialize(
        address _signerStorage,
        address _bridge,
        address _validatorFeeReceiver
    ) external initializer {
        _setSignerStorage(_signerStorage);
        setBridge(_bridge);
        setValidatorFeeReceiver(_validatorFeeReceiver);
    }

    function setValidatorFeeReceiver(address _validatorFeeReceiver) public onlySigner {
        validatorFeeReceiver = _validatorFeeReceiver;
        emit ValidatorFeeReceiverUpdated(_validatorFeeReceiver);
    }

    function setLimitPerToken(address _token, uint256 _limit) public onlySigner {
        limitPerToken[_token] = _limit;
        emit LimitPerTokenUpdated(_token, _limit);
    }

    function setBridge(address _bridge) public onlySigner {
        bridge = Bridge(_bridge);
        emit BridgeUpdated(_bridge);
    }

    function collect(address _token) public {
        require(limitPerToken[_token] > 0, "BridgeValidatorFeePool: no limit for this token");

        uint256 balanceAmount;

        if (_token == NATIVE_TOKEN) {
            balanceAmount = address(this).balance;

            require(limitPerToken[_token] < balanceAmount, "BridgeValidatorFeePool: insufficient funds");
            bridge.depositNative{value: balanceAmount}(block.chainid, validatorFeeReceiver);
        } else {
            balanceAmount = IERC20(_token).balanceOf(address(this));

            require(limitPerToken[_token] < balanceAmount, "BridgeValidatorFeePool: insufficient funds");

            IERC20(_token).approve(address(bridge), balanceAmount);

            bridge.deposit(_token, block.chainid, validatorFeeReceiver, balanceAmount);
        }

        emit Collected(_token, balanceAmount);
    }
}
