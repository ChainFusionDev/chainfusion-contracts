// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "../common/AddressStorage.sol";
import "./DKG.sol";
import "./ContractKeys.sol";
import "./ContractRegistry.sol";
import "./SlashingVoting.sol";

contract Staking is ContractKeys, Ownable, Initializable {
    enum ValidatorStatus {
        INACTIVE,
        ACTIVE,
        SLASHED
    }

    struct ValidatorInfo {
        address validator;
        uint256 stake;
        ValidatorStatus status;
    }

    struct WithdrawalAnnouncement {
        uint256 amount;
        uint256 time;
    }

    uint256 public minimalStake;
    uint256 public withdrawalPeriod;
    mapping(address => ValidatorInfo) public stakes;
    mapping(address => WithdrawalAnnouncement) public withdrawalAnnouncements;
    ContractRegistry public contractRegistry;
    AddressStorage public validatorStorage;

    event MinimalStakeUpdated(uint256 minimalStake);
    event WithdrawalPeriodUpdated(uint256 withdrawalPeriod);
    event ContractRegistryUpdated(address contractRegistry);

    modifier onlyActiveValidator() {
        require(stakes[msg.sender].status == ValidatorStatus.ACTIVE, "Staking: only active validator");
        _;
    }

    modifier onlyNotSlashed() {
        require(stakes[msg.sender].status != ValidatorStatus.SLASHED, "Staking: validator is slashed");
        _;
    }

    modifier onlySlashingVoting() {
        require(msg.sender == address(_slashingVotingContract()), "Staking: not a slashing voting");
        _;
    }

    function initialize(
        uint256 _minimalStake,
        uint256 _withdrawalPeriod,
        address _contractRegistry,
        address _validatorStorage
    ) external initializer {
        setMinimalStake(_minimalStake);
        setWithdrawalPeriod(_withdrawalPeriod);
        contractRegistry = ContractRegistry(_contractRegistry);
        validatorStorage = AddressStorage(_validatorStorage);
    }

    function isValidatorActive(address _validator) external view returns (bool) {
        return (stakes[_validator].status == ValidatorStatus.ACTIVE);
    }

    function isValidatorSlashing(address _validator) external view returns (bool) {
        if (stakes[_validator].status == ValidatorStatus.SLASHED) {
            return true;
        }
        return false;
    }

    function setMinimalStake(uint256 _minimalStake) public onlyOwner {
        minimalStake = _minimalStake;
        emit MinimalStakeUpdated(_minimalStake);
    }

    function setWithdrawalPeriod(uint256 _withdrawalPeriod) public onlyOwner {
        withdrawalPeriod = _withdrawalPeriod;
        emit WithdrawalPeriodUpdated(_withdrawalPeriod);
    }

    function slash(address _validator) public onlySlashingVoting {
        stakes[_validator].status = ValidatorStatus.SLASHED;
        _removeValidator(_validator);
    }

    function announceWithdrawal(uint256 _amount) public onlyNotSlashed {
        require(_amount <= stakes[msg.sender].stake, "Staking: amount must be <= to stake");
        withdrawalAnnouncements[msg.sender].amount = _amount;
        // solhint-disable-next-line not-rely-on-time
        withdrawalAnnouncements[msg.sender].time = block.timestamp;
    }

    function withdraw() public onlyNotSlashed {
        require(withdrawalAnnouncements[msg.sender].amount > 0, "Staking: amount must be greater than zero");
        require(
            // solhint-disable-next-line not-rely-on-time
            withdrawalAnnouncements[msg.sender].time + withdrawalPeriod <= block.timestamp,
            "Staking: withdrawal period not passed"
        );

        uint256 withdrawalAmount = withdrawalAnnouncements[msg.sender].amount;

        require(withdrawalAmount <= stakes[msg.sender].stake, "Staking: amount must be <= to validator stake");

        stakes[msg.sender].stake -= withdrawalAmount;

        withdrawalAnnouncements[msg.sender].amount = 0;
        withdrawalAnnouncements[msg.sender].time = 0;

        if (stakes[msg.sender].stake < minimalStake && validatorStorage.contains(msg.sender)) {
            stakes[msg.sender].status = ValidatorStatus.INACTIVE;
            _removeValidator(msg.sender);
        }

        // solhint-disable-next-line avoid-low-level-calls
        (bool success, ) = msg.sender.call{value: withdrawalAmount, gas: 21000}("");
        require(success, "Staking: transfer failed");
    }

    function stake() public payable {
        require(msg.value + stakes[msg.sender].stake >= minimalStake, "Staking: insufficient stake provided");
        require(stakes[msg.sender].status != ValidatorStatus.SLASHED, "Staking: validator is slashed");

        if (stakes[msg.sender].status == ValidatorStatus.INACTIVE) {
            stakes[msg.sender].validator = msg.sender;
            stakes[msg.sender].status = ValidatorStatus.ACTIVE;
            _addValidator(msg.sender);
        }

        stakes[msg.sender].stake += msg.value;
    }

    function getValidators() public view returns (address[] memory) {
        return validatorStorage.getAddresses();
    }

    function _addValidator(address validator) private {
        DKG dkg = _dkgContract();

        validatorStorage.mustAdd(validator);
        dkg.updateGeneration(validator);
    }

    function _removeValidator(address validator) private {
        DKG dkg = _dkgContract();

        validatorStorage.mustRemove(validator);
        dkg.updateGeneration(validator);
    }

    function _dkgContract() private view returns (DKG) {
        return DKG(contractRegistry.getContract(DKG_KEY));
    }

    function _slashingVotingContract() private view returns (SlashingVoting) {
        return SlashingVoting(contractRegistry.getContract(SLASHING_VOTING_KEY));
    }
}
