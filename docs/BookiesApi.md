# \BookiesApi

All URIs are relative to *http://localhost/admin/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBookieRackInfo**](BookiesApi.md#DeleteBookieRackInfo) | **Delete** /bookies/racks-info/{bookie} | Removed the rack placement information for a specific bookie in the cluster
[**GetBookieRackInfo**](BookiesApi.md#GetBookieRackInfo) | **Get** /bookies/racks-info/{bookie} | Gets the rack placement information for a specific bookie in the cluster
[**GetBookiesRackInfo**](BookiesApi.md#GetBookiesRackInfo) | **Get** /bookies/racks-info | Gets the rack placement information for all the bookies in the cluster
[**UpdateBookieRackInfo**](BookiesApi.md#UpdateBookieRackInfo) | **Post** /bookies/racks-info/{bookie} | Updates the rack placement information for a specific bookie in the cluster



## DeleteBookieRackInfo

> DeleteBookieRackInfo(ctx, bookie)
Removed the rack placement information for a specific bookie in the cluster

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookie** | **string**|  | 

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


## GetBookieRackInfo

> BookieInfo GetBookieRackInfo(ctx, bookie)
Gets the rack placement information for a specific bookie in the cluster

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookie** | **string**|  | 

### Return type

[**BookieInfo**](BookieInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBookiesRackInfo

> map[string]map[string]BookieInfo GetBookiesRackInfo(ctx, )
Gets the rack placement information for all the bookies in the cluster

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**map[string]map[string]BookieInfo**](map.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBookieRackInfo

> UpdateBookieRackInfo(ctx, bookie, optional)
Updates the rack placement information for a specific bookie in the cluster

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookie** | **string**|  | 
 **optional** | ***UpdateBookieRackInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateBookieRackInfoOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **group** | **optional.String**|  | 

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

