# \ResourceQuotasApi

All URIs are relative to *http://localhost/admin/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDefaultResourceQuota**](ResourceQuotasApi.md#GetDefaultResourceQuota) | **Get** /resource-quotas | Get the default quota
[**GetNamespaceBundleResourceQuota**](ResourceQuotasApi.md#GetNamespaceBundleResourceQuota) | **Get** /resource-quotas/{tenant}/{namespace}/{bundle} | Get resource quota of a namespace bundle.
[**RemoveNamespaceBundleResourceQuota**](ResourceQuotasApi.md#RemoveNamespaceBundleResourceQuota) | **Delete** /resource-quotas/{tenant}/{namespace}/{bundle} | Remove resource quota for a namespace.
[**SetDefaultResourceQuota**](ResourceQuotasApi.md#SetDefaultResourceQuota) | **Post** /resource-quotas | Set the default quota
[**SetNamespaceBundleResourceQuota**](ResourceQuotasApi.md#SetNamespaceBundleResourceQuota) | **Post** /resource-quotas/{tenant}/{namespace}/{bundle} | Set resource quota on a namespace.



## GetDefaultResourceQuota

> []string GetDefaultResourceQuota(ctx, )
Get the default quota

### Required Parameters

This endpoint does not need any parameter.

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


## GetNamespaceBundleResourceQuota

> ResourceQuota GetNamespaceBundleResourceQuota(ctx, tenant, namespace, bundle)
Get resource quota of a namespace bundle.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 
**bundle** | **string**|  | 

### Return type

[**ResourceQuota**](ResourceQuota.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveNamespaceBundleResourceQuota

> RemoveNamespaceBundleResourceQuota(ctx, tenant, namespace, bundle)
Remove resource quota for a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 
**bundle** | **string**|  | 

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


## SetDefaultResourceQuota

> []string SetDefaultResourceQuota(ctx, )
Set the default quota

### Required Parameters

This endpoint does not need any parameter.

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


## SetNamespaceBundleResourceQuota

> SetNamespaceBundleResourceQuota(ctx, tenant, namespace, bundle)
Set resource quota on a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 
**bundle** | **string**|  | 

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

