// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./TokenManager.sol";
import "./ValidatorManager.sol";
import "./LiquidityPools.sol";
import "./Globals.sol";
import "../interfaces/IERC20MintableBurnable.sol";

contract Bridge is Initializable, Ownable {
    mapping(bytes32 => mapping(address => bool)) public approvals;
    mapping(bytes32 => uint256) public approvalsCount;
    mapping(bytes32 => bool) public executed;
    TokenManager public tokenManager;
    ValidatorManager public validatorManager;
    LiquidityPools public liquidityPools;

    event Approved(bytes32 id, address validator);
    event Deposited(
        address token,
        address destinationToken,
        uint256 destinationChainId,
        address receiver,
        uint256 amount
    );
    event Transferred(
        address token,
        uint256 sourceChainId,
        address receiver,
        uint256 fee,
        uint256 transferAmount,
        address validator
    );
    event TokenManagerUpdated(address _tokenManager);
    event ValidatorManagerUpdated(address _validatorManager);
    event LiquidityPoolsUpdated(address _liquidityPools);

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
        emit TokenManagerUpdated(_tokenManager);
    }

    function setValidatorManager(address _validatorManager) external onlyOwner {
        validatorManager = ValidatorManager(_validatorManager);
        emit ValidatorManagerUpdated(_validatorManager);
    }

    function setLiquidityPools(address _liquidityPools) external onlyOwner {
        liquidityPools = LiquidityPools(_liquidityPools);
        emit LiquidityPoolsUpdated(_liquidityPools);
    }

    function deposit(
        address _token,
        uint256 _chainId,
        address _receiver,
        uint256 _amount
    ) external {
        require(_amount != 0, "Bridge: amount cannot be equal to 0.");
        require(tokenManager.isTokenEnabled(_token), "TokenManager: token is not supported");

        if (tokenManager.isTokenMintable(_token)) {
            IERC20MintableBurnable(_token).burnFrom(msg.sender, _amount);
        } else {
            require(
                IERC20(_token).transferFrom(msg.sender, address(liquidityPools), _amount),
                "IERC20: transfer failed"
            );
        }
        emit Deposited(_token, tokenManager.getDestinationToken(_token, _chainId), _chainId, _receiver, _amount);
    }

    function approveTransfer(
        bytes calldata _txHash,
        address _token,
        uint256 _sourceChainId,
        address _receiver,
        uint256 _amount
    ) external onlyValidator {
        require(tokenManager.isTokenEnabled(_token), "TokenManager: token is not supported");
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
            uint256 fee = (_amount * liquidityPools.feePercentage()) / BASE_DIVISOR;
            uint256 transferAmount = _amount - fee;

            if (tokenManager.isTokenMintable(_token)) {
                IERC20MintableBurnable(_token).mint(_receiver, _amount);
            } else {
                liquidityPools.transfer(_token, _receiver, fee, transferAmount);
            }

            emit Transferred(_token, _sourceChainId, _receiver, fee, transferAmount, msg.sender);
        }
    }

    function isApproved(
        bytes calldata _txHash,
        address _token,
        address _receiver,
        uint256 _amount
    ) public view returns (bool) {
        bytes32 id = keccak256(abi.encodePacked(_txHash, _token, _receiver, _amount));
        return executed[id];
    }
}
