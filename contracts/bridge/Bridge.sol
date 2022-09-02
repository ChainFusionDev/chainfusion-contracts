// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./SignerOwnable.sol";
import "./TokenManager.sol";
import "./LiquidityPools.sol";
import "./FeeManager.sol";
import "./Globals.sol";
import "../interfaces/IERC20MintableBurnable.sol";
import "./RelayBridge.sol";
import "../interfaces/IBridgeApp.sol";
import "hardhat/console.sol";

contract Bridge is Initializable, SignerOwnable, IBridgeApp {
    mapping(bytes32 => bool) public executed;

    address public bridgeAppAddress;

    TokenManager public tokenManager;
    LiquidityPools public liquidityPools;
    FeeManager public feeManager;
    RelayBridge public relayBridge;

    event Deposited(
        address token,
        address destinationToken,
        uint256 destinationChainId,
        address receiver,
        uint256 fee,
        uint256 transferAmount
    );
    event DepositedNative(
        address token,
        uint256 destinationChainId,
        address receiver,
        uint256 fee,
        uint256 transferAmount
    );
    event Transferred(address token, uint256 sourceChainId, address receiver, uint256 amount);
    event TokenManagerUpdated(address _tokenManager);
    event ValidatorAddressUpdated(address _validatorAddress);
    event LiquidityPoolsUpdated(address _liquidityPools);
    event FeeManagerUpdated(address _feeManager);
    event Reverted(address token, uint256 sourceChainId, address receiver, uint256 amount);

    modifier onlyRelayBridge() {
        require(msg.sender == address(relayBridge), "Bridge: only RelayBridge");
        _;
    }

    function initialize(
        address _relayBridgeAddress,
        address _signerStorage,
        address _tokenManager,
        address payable _liquidityPools,
        address payable _feeManager,
        address _bridgeAppAddress
    ) external initializer {
        _setSignerStorage(_signerStorage);
        setTokenManager(_tokenManager);
        setLiquidityPools(_liquidityPools);
        setFeeManager(_feeManager);
        relayBridge = RelayBridge(_relayBridgeAddress);
        bridgeAppAddress = _bridgeAppAddress;
    }

    function deposit(
        address _token,
        uint256 _chainId,
        address _receiver,
        uint256 _amount
    ) external {
        require(_amount != 0, "Bridge: amount cannot be equal to 0.");
        require(tokenManager.isTokenEnabled(_token), "TokenManager: token is not enabled");

        uint256 fee = feeManager.calculateFee(_token, _amount);
        require(IERC20(_token).transferFrom(msg.sender, address(feeManager), fee), "IERC20: transfer failed");

        uint256 transferAmount = _amount - fee;

        if (tokenManager.isTokenMintable(_token)) {
            IERC20MintableBurnable(_token).burnFrom(msg.sender, transferAmount);
        } else {
            require(
                IERC20(_token).transferFrom(msg.sender, address(liquidityPools), transferAmount),
                "IERC20: transfer failed"
            );
        }

        address destinationToken = tokenManager.getDestinationToken(_token, _chainId);

        emit Deposited(_token, destinationToken, _chainId, _receiver, fee, transferAmount);

        bytes memory data = abi.encode(_token, _chainId, _receiver, transferAmount);

        // solhint-disable-next-line check-send-result
        relayBridge.send(_chainId, block.gaslimit, data);
    }

    function execute(uint256, bytes memory data) external onlyRelayBridge {
        (address _token, uint256 _chainId, address _receiver, uint256 transferAmount) = abi.decode(
            data,
            (address, uint256, address, uint256)
        );

        address destinationToken = tokenManager.getDestinationToken(_token, _chainId);

        _executeTransfer(data, destinationToken, _chainId, _receiver, transferAmount);
    }

    function executeTransfer(
        bytes calldata _txHash,
        address _token,
        uint256 _sourceChainId,
        address _receiver,
        uint256 _amount
    ) external onlySigner {
        _executeTransfer(_txHash, _token, _sourceChainId, _receiver, _amount);
    }

    function revertSend(uint256, bytes memory data) external onlyRelayBridge {
        (address _token, uint256 _chainId, address _receiver, uint256 _amount) = abi.decode(
            data,
            (address, uint256, address, uint256)
        );

        require(tokenManager.isTokenEnabled(_token), "TokenManager: token is not enabled");

        if (tokenManager.isTokenMintable(_token)) {
            IERC20MintableBurnable(_token).mint(_receiver, _amount);
        } else if (_token == NATIVE_TOKEN) {
            liquidityPools.transferNative(_receiver, _amount);
        } else {
            liquidityPools.transfer(_token, _receiver, _amount);
        }

        emit Reverted(_token, _chainId, _receiver, _amount);
    }

    function setTokenManager(address _tokenManager) public onlySigner {
        tokenManager = TokenManager(_tokenManager);
        emit TokenManagerUpdated(_tokenManager);
    }

    function setLiquidityPools(address payable _liquidityPools) public onlySigner {
        liquidityPools = LiquidityPools(_liquidityPools);
        emit LiquidityPoolsUpdated(_liquidityPools);
    }

    function setFeeManager(address payable _feeManager) public onlySigner {
        feeManager = FeeManager(_feeManager);
        emit FeeManagerUpdated(_feeManager);
    }

    function depositNative(uint256 _chainId, address _receiver) public payable {
        uint256 _amount = msg.value;
        require(_amount != 0, "Bridge: amount cannot be equal to 0.");
        require(tokenManager.isTokenEnabled(NATIVE_TOKEN), "TokenManager: token is not enabled");

        uint256 fee = feeManager.calculateFee(NATIVE_TOKEN, _amount);

        // solhint-disable-next-line avoid-low-level-calls
        (bool success, ) = address(feeManager).call{value: fee, gas: 21000}("");
        require(success, "Bridge: transfer native token failed");

        uint256 transferAmount = _amount - fee;

        // solhint-disable-next-line avoid-low-level-calls
        (success, ) = address(liquidityPools).call{value: transferAmount, gas: 21000}("");
        require(success, "Bridge: transfer native token failed");

        emit DepositedNative(NATIVE_TOKEN, _chainId, _receiver, fee, transferAmount);

        bytes memory data = abi.encode(NATIVE_TOKEN, _chainId, _receiver, transferAmount);

        // solhint-disable-next-line check-send-result
        relayBridge.send(_chainId, block.gaslimit, data);
    }

    function isExecuted(
        bytes calldata _txHash,
        address _token,
        address _receiver,
        uint256 _amount
    ) public view returns (bool) {
        bytes32 id = keccak256(abi.encodePacked(_txHash, _token, _receiver, _amount));
        return executed[id];
    }

    function _executeTransfer(
        bytes memory _txHash,
        address _token,
        uint256 _sourceChainId,
        address _receiver,
        uint256 _amount
    ) private {
        require(tokenManager.isTokenEnabled(_token), "TokenManager: token is not enabled");
        bytes32 id = keccak256(abi.encodePacked(_txHash, _token, _receiver, _amount));

        if (executed[id]) {
            return;
        }

        executed[id] = true;

        if (tokenManager.isTokenMintable(_token)) {
            IERC20MintableBurnable(_token).mint(_receiver, _amount);
        } else if (_token == NATIVE_TOKEN) {
            liquidityPools.transferNative(_receiver, _amount);
        } else {
            liquidityPools.transfer(_token, _receiver, _amount);
        }

        emit Transferred(_token, _sourceChainId, _receiver, _amount);
    }
}
