/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type CursorStats struct {
	MarkDeletePosition string `json:"markDeletePosition,omitempty"`
	ReadPosition string `json:"readPosition,omitempty"`
	WaitingReadOp bool `json:"waitingReadOp,omitempty"`
	PendingReadOps int32 `json:"pendingReadOps,omitempty"`
	MessagesConsumedCounter int64 `json:"messagesConsumedCounter,omitempty"`
	CursorLedger int64 `json:"cursorLedger,omitempty"`
	CursorLedgerLastEntry int64 `json:"cursorLedgerLastEntry,omitempty"`
	IndividuallyDeletedMessages string `json:"individuallyDeletedMessages,omitempty"`
	LastLedgerSwitchTimestamp string `json:"lastLedgerSwitchTimestamp,omitempty"`
	State string `json:"state,omitempty"`
	NumberOfEntriesSinceFirstNotAckedMessage int64 `json:"numberOfEntriesSinceFirstNotAckedMessage,omitempty"`
	TotalNonContiguousDeletedMessagesRange int32 `json:"totalNonContiguousDeletedMessagesRange,omitempty"`
	Properties map[string]int64 `json:"properties,omitempty"`
}
