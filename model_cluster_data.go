/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// The configuration data for a cluster
type ClusterData struct {
	// The HTTP rest service URL (for admin operations)
	ServiceUrl string `json:"serviceUrl,omitempty"`
	// The HTTPS rest service URL (for admin operations)
	ServiceUrlTls string `json:"serviceUrlTls,omitempty"`
	// The broker service url (for produce and consume operations)
	BrokerServiceUrl string `json:"brokerServiceUrl,omitempty"`
	// The secured broker service url (for produce and consume operations)
	BrokerServiceUrlTls string `json:"brokerServiceUrlTls,omitempty"`
	// A set of peer cluster names
	PeerClusterNames []string `json:"peerClusterNames,omitempty"`
}
