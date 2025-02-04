// SPDX-License-Identifier: UNLICENSED
pragma solidity =0.8.12;

import {IServiceManager} from "eigenlayer-middleware/interfaces/IServiceManager.sol";
import {BLSSignatureChecker} from "eigenlayer-middleware/BLSSignatureChecker.sol";
import {IMachOptimism} from "../interfaces/IMachOptimism.sol";

interface ITotalAlertsLegacy {
    // Legacy
    /// @notice Returns the length of total alerts
    function totalAlerts() external view returns (uint256);
}

interface ITotalAlerts {
    /// @notice Returns the length of total alerts
    function totalAlerts(uint256 rollupChainId) external view returns (uint256);
}

/**
 * @title Interface for the MachServiceManager contract.
 * @author Altlayer, Inc.
 */
interface IMachServiceManager is IServiceManager {
    struct AlertHeader {
        bytes32 messageHash;
        // for BLS verification
        bytes quorumNumbers; // each byte is a different quorum number
        bytes quorumThresholdPercentages; // every bytes is an amount less than 100 specifying the percentage of stake
            // the must have signed in the corresponding quorum in `quorumNumbers`
        uint32 referenceBlockNumber;
        uint256 rollupChainID;
    }

    struct ReducedAlertHeader {
        bytes32 messageHash;
        uint32 referenceBlockNumber;
        uint256 rollupChainID;
    }

    event AllowlistUpdated(address[] operators, bool[] status);

    /**
     * @notice Emitted when the alert confirmer is changed.
     * @param previousAddress The address of the previous alert confirmer
     * @param newAddress The address of the new alert confirmer
     */
    event AlertConfirmerChanged(address previousAddress, address newAddress);

    /**
     * @notice Emitted when the whitelister is changed.
     * @param previousAddress The address of the previous whitelister
     * @param newAddress The address of the new whitelister
     */
    event WhitelisterChanged(address previousAddress, address newAddress);

    /**
     * @notice Emitted when the quorum threshold percentage is changed.
     * @param thresholdPercentages The new quorum threshold percentage
     */
    event QuorumThresholdPercentageChanged(uint8 thresholdPercentages);

    /**
     * @notice Emitted when an operator is added to the allowlist.
     * @param operator The operator
     */
    event OperatorAllowed(address operator);

    /**
     * @notice Emitted when an operator is removed from the allowlist.
     * @param operator The operator
     */
    event OperatorDisallowed(address operator);

    /**
     * @notice Emitted when the allowlist is enabled.
     */
    event AllowlistEnabled();

    /**
     * @notice Emitted when the allowlist is disabled.
     */
    event AllowlistDisabled();

    /**
     * @notice Emitted when rollup chain id is updated
     */
    event RollupChainIDUpdated(uint256 rollupChainId, bool status);

    /**
     * @notice Emitted when a Alert is confirmed.
     * @param alertHeaderHash The hash of the alert header
     * @param messageHash The message hash
     */
    event AlertConfirmed(bytes32 indexed alertHeaderHash, bytes32 messageHash);

    /**
     * @notice Emitted when a Alert is removed.
     * @param messageHash The message hash
     * @param messageHash The sender address
     */
    event AlertRemoved(bytes32 messageHash, address sender);

    /**
     * @notice Enable the allowlist.
     */
    function enableAllowlist() external;

    /**
     * @notice Disable the allowlist.
     */
    function disableAllowlist() external;

    /**
     * @notice Set confirmer address.
     */
    function setConfirmer(address confirmer) external;

    /**
     * @notice Set whitelister address.
     */
    function setWhitelister(address whitelister) external;

    /**
     * @notice Set the status of a rollup chain ID
     * @param rollupChainId The ID of the rollup chain to be updated
     * @param status The new status for the rollup chain ID (true for active, false for inactive)
     */
    function setRollupChainID(uint256 rollupChainId, bool status) external;

    /**
     * @notice Remove an Alert.
     * @param messageHash The message hash of the alert
     */
    function removeAlert(uint256 rollupChainId, bytes32 messageHash) external;

    /**
     * @notice Update quorum threshold percentage
     * @param thresholdPercentage The new quorum threshold percentage
     */
    function updateQuorumThresholdPercentage(uint8 thresholdPercentage) external;

    /**
     * @notice This function is used for
     * - submitting alert,
     * - check that the aggregate signature is valid,
     * - and check whether quorum has been achieved or not.
     */
    function confirmAlert(
        AlertHeader calldata alertHeader,
        BLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) external;

    /// @notice Returns the length of total alerts
    function totalAlerts(uint256 rollupChainId) external view returns (uint256);

    /// @notice Checks if messageHash exists
    function contains(uint256 rollupChainId, bytes32 messageHash) external view returns (bool);

    /// @notice Returns an array of messageHash
    function queryMessageHashes(uint256 rollupChainId, uint256 start, uint256 querySize)
        external
        view
        returns (bytes32[] memory);
}
