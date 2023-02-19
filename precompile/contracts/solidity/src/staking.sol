// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

/**
 * @dev Library of all supported Cosmos events emitted by the staking module
 */
library StakingEvents {
    /**
     * @dev Emitted by the staking module when `amount` tokens are delegated to
     * `validator`
     */
    event Delegate(address indexed validator, uint256 amount);

    /**
     * @dev Emitted by the staking module when `amount` tokens are redelegated from
     * `sourceValidator` to `destinationValidator`
     */
    event Redelegate(
        address indexed sourceValidator,
        address indexed destinationValidator,
        uint256 amount
    );

    /**
     * @dev Emitted by the staking module when `amount` tokens are used to create `validator`
     */
    event CreateValidator(address indexed validator, uint256 amount);

    /**
     * @dev Emitted by the staking module when `amount` tokens are unbonded from `validator`
     */
    event Unbond(address indexed validator, uint256 amount);

    /**
     * @dev Emitted by the staking module when `amount` tokens are canceled from `delegator`'s
     * unbonding delegation with `validator`
     */
    event CancelUnbondingDelegation(
        address indexed validator,
        address indexed delegator,
        uint256 amount,
        int64 creationHeight
    );
}

/**
 * @dev Interface of the staking module's precompiled contract
 */
interface IStakingModule {
    /////////////////////////////////////// READ METHODS //////////////////////////////////////////

    /**
     * @dev Returns the `amount` of tokens currently delegated by msg.sender to `validatorAddress`
     */
    function getDelegation(
        address validatorAddress
    ) external view returns (uint256);

    /**
     * @dev Returns the `amount` of tokens currently delegated by msg.sender to `validatorAddress`
     * (at hex bech32 address)
     */
    function getDelegation(
        string calldata validatorAddress
    ) external view returns (uint256);

    /**
     * @dev Returns a time-ordered list of all UnbondingDelegationEntries between msg.sender and
     * `validatorAddress`
     */
    function getUnbondingDelegation(
        address validatorAddress
    ) external view returns (UnbondingDelegationEntry[] memory);

    /**
     * @dev Returns a time-ordered list of all UnbondingDelegationEntries between msg.sender and
     * `validatorAddress` (at hex bech32 address)
     */
    function getUnbondingDelegation(
        string calldata validatorAddress
    ) external view returns (UnbondingDelegationEntry[] memory);

    /**
     * @dev Returns a list of the msg.sender's redelegating bonds from `srcValidator` to
     * `dstValidator`
     */
    function getRedelegations(
        address srcValidator,
        address dstValidator
    ) external view returns (RedelegationEntry[] memory);

    /**
     * @dev Returns a list of the msg.sender's redelegating bonds from `srcValidator` to
     * `dstValidator` (at hex bech32 addresses)
     */
    function getRedelegations(
        string calldata srcValidator,
        string calldata dstValidator
    ) external view returns (RedelegationEntry[] memory);

    ////////////////////////////////////// WRITE METHODS //////////////////////////////////////////

    /**
     * @dev msg.sender delegates the `amount` of tokens to `validatorAddress`
     */
    function delegate(
        address validatorAddress,
        uint256 amount
    ) external payable;

    /**
     * @dev msg.sender delegates the `amount` of tokens to `validatorAddress` (at hex bech32
     * address)
     */
    function delegate(
        string calldata validatorAddress,
        uint256 amount
    ) external payable;

    /**
     * @dev msg.sender undelegates the `amount` of tokens from `validatorAddress`
     */
    function undelegate(
        address validatorAddress,
        uint256 amount
    ) external payable;

    /**
     * @dev msg.sender undelegates the `amount` of tokens from `validatorAddress` (at hex bech32
     * address)
     */
    function undelegate(
        string calldata validatorAddress,
        uint256 amount
    ) external payable;

    /**
     * @dev msg.sender redelegates the `amount` of tokens from `srcValidator` to
     * `validtorDstAddr`
     */
    function beginRedelegate(
        address srcValidator,
        address dstValidator,
        uint256 amount
    ) external payable;

    /**
     * @dev msg.sender redelegates the `amount` of tokens from `srcValidator` to
     * `validtorDstAddr` (at hex bech32 addresses)
     */
    function beginRedelegate(
        string calldata srcValidator,
        string calldata dstValidator,
        uint256 amount
    ) external payable;

    /**
     * @dev Cancels msg.sender's unbonding delegation with `validatorAddress` and delegates the
     * `amount` of tokens back to `validatorAddress`
     *
     * Provide the `creationHeight` of the original unbonding delegation
     */
    function cancelUnbondingDelegation(
        address validatorAddress,
        uint256 amount,
        int64 creationHeight
    ) external payable;

    /**
     * @dev Cancels msg.sender's unbonding delegation with `validatorAddress` and delegates the
     * `amount` of tokens back to `validatorAddress` (at hex bech32 addresses)
     *
     * Provide the `creationHeight` of the original unbonding delegation
     */
    function cancelUnbondingDelegation(
        string calldata validatorAddress,
        uint256 amount,
        int64 creationHeight
    ) external payable;

    //////////////////////////////////////////// UTILS ////////////////////////////////////////////

    /**
     * @dev Represents one entry of an unbonding delegation
     *
     * Note: the field names of the native struct should match these field names (by camelCase)
     */
    struct UnbondingDelegationEntry {
        // creationHeight is the height which the unbonding took place
        int64 creationHeight;
        // completionTime is the unix time for unbonding completion, formatted as a string
        string completionTime;
        // initialBalance defines the tokens initially scheduled to receive at completion
        uint256 initialBalance;
        // balance defines the tokens to receive at completion
        uint256 balance;
        // unbondingingId incrementing id that uniquely identifies this entry
        uint64 unbondingId;
    }

    /**
     * @dev Represents a redelegation entry with relevant metadata
     *
     * Note: the field names of the native struct should match these field names (by camelCase)
     */
    struct RedelegationEntry {
        // creationHeight is the height which the redelegation took place
        int64 creationHeight;
        // completionTime is the unix time for redelegation completion, formatted as a string
        string completionTime;
        // initialBalance defines the initial balance when redelegation started
        uint256 initialBalance;
        // sharesDst is the amount of destination-validatorAddress shares created by redelegation
        uint256 sharesDst;
        // unbondingId is the incrementing id that uniquely identifies this entry
        uint64 unbondingId;
    }
}
