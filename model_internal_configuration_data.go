/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InternalConfigurationData struct {
	ZookeeperServers string `json:"zookeeperServers,omitempty"`
	ConfigurationStoreServers string `json:"configurationStoreServers,omitempty"`
	LedgersRootPath string `json:"ledgersRootPath,omitempty"`
	StateStorageServiceUrl string `json:"stateStorageServiceUrl,omitempty"`
}
