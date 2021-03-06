/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type OffloadProcessStatus struct {
	Status string `json:"status,omitempty"`
	LastError string `json:"lastError,omitempty"`
	FirstUnoffloadedMessage MessageIdImpl `json:"firstUnoffloadedMessage,omitempty"`
}
