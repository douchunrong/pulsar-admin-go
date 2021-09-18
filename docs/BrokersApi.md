# \BrokersApi

All URIs are relative to *http://localhost/admin/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActiveBrokers**](BrokersApi.md#GetActiveBrokers) | **Get** /brokers/{cluster} | Get the list of active brokers (web service addresses) in the cluster.If authorization is not enabled, any cluster name is valid.
[**GetAllDynamicConfigurations**](BrokersApi.md#GetAllDynamicConfigurations) | **Get** /brokers/configuration/values | Get value of all dynamic configurations&#39; value overridden on local config
[**GetDynamicConfigurationName**](BrokersApi.md#GetDynamicConfigurationName) | **Get** /brokers/configuration | Get all updatable dynamic configurations&#39;s name
[**GetInternalConfigurationData**](BrokersApi.md#GetInternalConfigurationData) | **Get** /brokers/internal-configuration | Get the internal configuration data
[**GetOwnedNamespaes**](BrokersApi.md#GetOwnedNamespaes) | **Get** /brokers/{clusterName}/{broker-webserviceurl}/ownedNamespaces | Get the list of namespaces served by the specific broker
[**GetRuntimeConfiguration**](BrokersApi.md#GetRuntimeConfiguration) | **Get** /brokers/configuration/runtime | Get all runtime configurations. This operation requires Pulsar super-user privileges.
[**Healthcheck**](BrokersApi.md#Healthcheck) | **Get** /brokers/health | Run a healthcheck against the broker
[**UpdateDynamicConfiguration**](BrokersApi.md#UpdateDynamicConfiguration) | **Post** /brokers/configuration/{configName}/{configValue} | Update dynamic serviceconfiguration into zk only. This operation requires Pulsar super-user privileges.



## GetActiveBrokers

> []string GetActiveBrokers(ctx, cluster)
Get the list of active brokers (web service addresses) in the cluster.If authorization is not enabled, any cluster name is valid.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string**|  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllDynamicConfigurations

> map[string]map[string]interface{} GetAllDynamicConfigurations(ctx, )
Get value of all dynamic configurations' value overridden on local config

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**map[string]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDynamicConfigurationName

> []map[string]interface{} GetDynamicConfigurationName(ctx, )
Get all updatable dynamic configurations's name

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInternalConfigurationData

> InternalConfigurationData GetInternalConfigurationData(ctx, )
Get the internal configuration data

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InternalConfigurationData**](InternalConfigurationData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOwnedNamespaes

> map[string]NamespaceOwnershipStatus GetOwnedNamespaes(ctx, clusterName, brokerWebserviceurl)
Get the list of namespaces served by the specific broker

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string**|  | 
**brokerWebserviceurl** | **string**|  | 

### Return type

[**map[string]NamespaceOwnershipStatus**](NamespaceOwnershipStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuntimeConfiguration

> map[string]map[string]interface{} GetRuntimeConfiguration(ctx, )
Get all runtime configurations. This operation requires Pulsar super-user privileges.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**map[string]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Healthcheck

> Healthcheck(ctx, )
Run a healthcheck against the broker

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDynamicConfiguration

> UpdateDynamicConfiguration(ctx, configName, configValue)
Update dynamic serviceconfiguration into zk only. This operation requires Pulsar super-user privileges.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configName** | **string**|  | 
**configValue** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

