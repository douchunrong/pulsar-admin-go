# \SchemasApi

All URIs are relative to *http://localhost/admin/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSchema**](SchemasApi.md#DeleteSchema) | **Delete** /schemas/{tenant}/{namespace}/{topic}/schema | Delete the schema of a topic
[**GetSchema**](SchemasApi.md#GetSchema) | **Get** /schemas/{tenant}/{namespace}/{topic}/schema | Get the schema of a topic
[**GetSchema_0**](SchemasApi.md#GetSchema_0) | **Get** /schemas/{tenant}/{namespace}/{topic}/schema/{version} | Get the schema of a topic at a given version
[**PostSchema**](SchemasApi.md#PostSchema) | **Post** /schemas/{tenant}/{namespace}/{topic}/schema | Update the schema of a topic



## DeleteSchema

> DeleteSchemaResponse DeleteSchema(ctx, tenant, namespace, topic, optional)
Delete the schema of a topic

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 
**topic** | **string**|  | 
 **optional** | ***DeleteSchemaOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteSchemaOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **optional.Bool**|  | [default to false]

### Return type

[**DeleteSchemaResponse**](DeleteSchemaResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchema

> GetSchemaResponse GetSchema(ctx, tenant, namespace, topic, optional)
Get the schema of a topic

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 
**topic** | **string**|  | 
 **optional** | ***GetSchemaOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSchemaOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **optional.Bool**|  | [default to false]

### Return type

[**GetSchemaResponse**](GetSchemaResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchema_0

> GetSchemaResponse GetSchema_0(ctx, tenant, namespace, topic, version, optional)
Get the schema of a topic at a given version

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 
**topic** | **string**|  | 
**version** | **string**|  | 
 **optional** | ***GetSchema_1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSchema_1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authoritative** | **optional.Bool**|  | [default to false]

### Return type

[**GetSchemaResponse**](GetSchemaResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSchema

> PostSchemaResponse PostSchema(ctx, tenant, namespace, topic, optional)
Update the schema of a topic

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 
**topic** | **string**|  | 
 **optional** | ***PostSchemaOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostSchemaOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **optional.Bool**|  | [default to false]
 **body** | [**optional.Interface of PostSchemaPayload**](PostSchemaPayload.md)| A JSON value presenting a schema playload. An example of the expected schema can be found down here. | 

### Return type

[**PostSchemaResponse**](PostSchemaResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

