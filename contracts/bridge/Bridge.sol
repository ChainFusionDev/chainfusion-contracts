// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./TokenManager.sol";
import "./ValidatorManager.sol";
import "./LiquidityPools.sol";

contract Bridge is Initializable, Ownable {
    mapping(bytes32 => mapping(address => bool)) public approvals;
    mapping(bytes32 => uint256) public approvalsCount;
    mapping(bytes32 => bool) public executed;
    TokenManager public tokenManager;
    ValidatorManager public validatorManager;
    LiquidityPools public liquidityPools;

    event Approved(bytes32 id, address validator);
    event Deposited(address token, address destinationToken, uint256 chainId, uint256 amount);
    event Transferred(address token, address receiver, uint256 amount, address validator);

    modifier onlyValidator() {
        require(validatorManager.isValidator(msg.sender), "Bridge: only validator");
        _;
    }

    function initialize(
        address _owner,
        ValidatorManager _validatorManager,
        address _tokenManager,
        address _liquidityPools
    ) external initializer {
        _transferOwnership(_owner);
        validatorManager = _validatorManager;
        tokenManager = TokenManager(_tokenManager);
        liquidityPools = LiquidityPools(_liquidityPools);
    }

    function setTokenManager(address _tokenManager) external onlyOwner {
        tokenManager = TokenManager(_tokenManager);
    }

    function setValidatorManager(address _validatorManager) external onlyOwner {
        validatorManager = ValidatorManager(_validatorManager);
    }

    function setLiquidityPools(address _liquidityPools) external onlyOwner {
        liquidityPools = LiquidityPools(_liquidityPools);
    }

    function deposit(
        address _token,
        uint256 _chainId,
        uint256 _amount
    ) external {
        require(_amount != 0, "Bridge: amount cannot be equal to 0.");
        require(IERC20(_token).transferFrom(msg.sender, address(liquidityPools), _amount), "IERC20: transfer failed");
        require(tokenManager.isTokenSupported(_token), "TokenManager: token is not supported");
        emit Deposited(_token, tokenManager.getDestinationToken(_token, _chainId), _chainId, _amount);
    }

    function approveTransfer(
        bytes calldata _txHash,
        address _token,
        address _receiver,
        uint256 _amount
    ) external onlyValidator {
        bytes32 id = keccak256(abi.encodePacked(_txHash, _token, _receiver, _amount));

        if (!approvals[id][msg.sender]) {
            approvals[id][msg.sender] = true;
            approvalsCount[id]++;
            emit Approved(id, msg.sender);
        }

        if (executed[id]) {
            return;
        }

        if (approvalsCount[id] >= validatorManager.requiredApprovals()) {
            executed[id] = true;
            liquidityPools.transfer(_token, _receiver, _amount);
            emit Transferred(_token, _receiver, _amount, msg.sender);
        }
    }
}
