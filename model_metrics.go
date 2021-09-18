/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Metrics struct {
	Metrics map[string]map[string]interface{} `json:"metrics,omitempty"`
	Dimensions map[string]string `json:"dimensions,omitempty"`
}