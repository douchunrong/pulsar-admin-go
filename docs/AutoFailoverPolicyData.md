# AutoFailoverPolicyData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyType** | **string** | The auto failover policy type | [optional] 
**Parameters** | **map[string]string** | The parameters applied to the auto failover policy specified by &#x60;policy_type&#x60;. The parameters for &#39;min_available&#39; are :   - &#39;min_limit&#39;: the limit of minimal number of available brokers in primary group before auto failover   - &#39;usage_threshold&#39;: the resource usage threshold. If the usage of a broker is beyond this value, it would be marked as unavailable  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


