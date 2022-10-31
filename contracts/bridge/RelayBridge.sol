// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./SignerOwnable.sol";
import "./TokenManager.sol";
import "./ERC20Bridge.sol";
import "./FeeManager.sol";
import "./Globals.sol";
import "./BridgeValidatorFeePool.sol";
import "../interfaces/IBridgeApp.sol";
import "../interfaces/IBridgeMediator.sol";

contract RelayBridge is Initializable, SignerOwnable {
    mapping(bytes32 => bytes) public sentData;

    mapping(bytes32 => bool) public sent;
    mapping(bytes32 => bool) public executed;
    mapping(bytes32 => bool) public reverted;
    mapping(bytes32 => bool) public failed;

    address[] public leaderHistory;

    BridgeValidatorFeePool public bridgeValidatorFeePool;

    uint256 public nonce;

    event Sent(
        bytes32 hash,
        address appContract,
        uint256 sourceChain,
        uint256 destinationChain,
        bytes data,
        uint256 gasLimit,
        uint256 nonce,
        uint256 validatorFee
    );
    event FailedSend(
        bytes32 hash,
        address appContract,
        uint256 sourceChain,
        uint256 destinationChain,
        bytes data,
        uint256 gasLimit,
        uint256 nonce
    );
    event Reverted(bytes32 hash, uint256 sourceChain, uint256 destinationChain);
    event Executed(bytes32 hash, uint256 sourceChain, uint256 destinationChain);

    function initialize(address _signerStorage, address payable _bridgeValidatorFeePool) external initializer {
        _setSignerStorage(_signerStorage);
        bridgeValidatorFeePool = BridgeValidatorFeePool(_bridgeValidatorFeePool);
    }

    function send(
        uint256 destinationChain,
        uint256 gasLimit,
        bytes memory data
    ) external payable {
        bytes32 hash = dataHash(msg.sender, block.chainid, destinationChain, gasLimit, data, nonce);
        require(sentData[hash].length == 0, "RelayBridge: data already send");

        sent[hash] = true;
        sentData[hash] = data;

        // solhint-disable-next-line avoid-low-level-calls
        (bool success, ) = address(bridgeValidatorFeePool).call{value: msg.value, gas: 21000}("");
        require(success, "RelayBridge: transfer value failed");

        emit Sent(hash, msg.sender, block.chainid, destinationChain, data, gasLimit, nonce, msg.value);

        nonce++;
    }

    function failedSend(
        address _appContract,
        uint256 _sourceChain,
        uint256 _destinationChain,
        bytes memory _data,
        uint256 _gasLimit,
        uint256 _nonce
    ) external onlySigner {
        bytes32 hash = dataHash(_appContract, _destinationChain, _sourceChain, _gasLimit, _data, _nonce);
        require(!failed[hash], "RelayBridge: data already failed");

        failed[hash] = true;
        emit FailedSend(hash, _appContract, _destinationChain, _sourceChain, _data, _gasLimit, _nonce);
    }

    function revertSend(
        address appContract,
        uint256 destinationChain,
        uint256 gasLimit,
        bytes memory data,
        uint256 _nonce,
        address leader
    ) external onlySigner {
        bytes32 hash = dataHash(appContract, block.chainid, destinationChain, gasLimit, data, _nonce);
        require(sent[hash], "RelayBridge: data never sent");
        require(!reverted[hash], "RelayBridge: data already reverted");

        reverted[hash] = true;
        leaderHistory.push(leader);

        IBridgeApp(appContract).revertSend(destinationChain, data);

        emit Reverted(hash, block.chainid, destinationChain);
    }

    function execute(
        address appContract,
        uint256 sourceChain,
        uint256 gasLimit,
        bytes memory data,
        uint256 _nonce,
        address leader
    ) external onlySigner {
        bytes32 hash = dataHash(appContract, sourceChain, block.chainid, gasLimit, data, _nonce);
        require(!executed[hash], "RelayBridge: data already executed");

        executed[hash] = true;
        leaderHistory.push(leader);

        IBridgeApp(appContract).execute(sourceChain, data);

        emit Executed(hash, sourceChain, block.chainid);
    }

    function leaderHistoryLength() external view returns (uint256) {
        return leaderHistory.length;
    }

    function dataHash(
        address appContract,
        uint256 sourceChain,
        uint256 destinationChain,
        uint256 gasLimit,
        bytes memory data,
        uint256 _nonce
    ) public pure returns (bytes32) {
        return keccak256(abi.encode(appContract, sourceChain, destinationChain, gasLimit, data, _nonce));
    }
}
