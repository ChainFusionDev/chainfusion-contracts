// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract ContractKeys {
    string public constant STAKING_KEY = "staking";
    string public constant DKG_KEY = "dkg";
    string public constant SUPPORTED_TOKENS_KEY = "supported-tokens";
    string public constant SLASHING_VOTING_KEY = "slashing-voting";
    string public constant EVENT_REGISTRY_KEY = "event-registry";
    string public constant VALIDATOR_REWARD_DISTRIBUTION_POOL_KEY = "validator-reward-distribution-pool";
}
