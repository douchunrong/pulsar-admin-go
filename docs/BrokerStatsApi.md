# \BrokerStatsApi

All URIs are relative to *http://localhost/admin/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllocatorStats**](BrokerStatsApi.md#GetAllocatorStats) | **Get** /broker-stats/allocator-stats/{allocator} | Get the stats for the Netty allocator. Available allocators are &#39;default&#39; and &#39;ml-cache&#39;
[**GetBrokerResourceAvailability**](BrokerStatsApi.md#GetBrokerResourceAvailability) | **Get** /broker-stats/broker-resource-availability/{tenant}/{namespace} | Broker availability report
[**GetLoadReport**](BrokerStatsApi.md#GetLoadReport) | **Get** /broker-stats/load-report | Get Load for this broker
[**GetMBeans**](BrokerStatsApi.md#GetMBeans) | **Get** /broker-stats/mbeans | Get all the mbean details of this broker JVM
[**GetMetrics**](BrokerStatsApi.md#GetMetrics) | **Get** /broker-stats/metrics | Gets the metrics for Monitoring
[**GetPendingBookieOpsStats**](BrokerStatsApi.md#GetPendingBookieOpsStats) | **Get** /broker-stats/bookieops | Get pending bookie client op stats by namesapce
[**GetTopics2**](BrokerStatsApi.md#GetTopics2) | **Get** /broker-stats/topics | Get all the topic stats by namespace



## GetAllocatorStats

> AllocatorStats GetAllocatorStats(ctx, allocator)
Get the stats for the Netty allocator. Available allocators are 'default' and 'ml-cache'

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allocator** | **string**|  | 

### Return type

[**AllocatorStats**](AllocatorStats.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrokerResourceAvailability

> map[string]ResourceUnit GetBrokerResourceAvailability(ctx, tenant, namespace)
Broker availability report

This API gives the current broker availability in percent, each resource percentage usage is calculated and thensum of all of the resource usage percent is called broker-resource-availability<br/><br/>THIS API IS ONLY FOR USE BY TESTING FOR CONFIRMING NAMESPACE ALLOCATION ALGORITHM

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

### Return type

[**map[string]ResourceUnit**](ResourceUnit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoadReport

> LoadReport GetLoadReport(ctx, )
Get Load for this broker

consists of topics stats & systemResourceUsage

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**LoadReport**](LoadReport.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMBeans

> []Metrics GetMBeans(ctx, )
Get all the mbean details of this broker JVM

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Metrics**](Metrics.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetrics

> []Metrics GetMetrics(ctx, )
Gets the metrics for Monitoring

Requested should be executed by Monitoring agent on each broker to fetch the metrics

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Metrics**](Metrics.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPendingBookieOpsStats

> map[string]PendingBookieOpsStats GetPendingBookieOpsStats(ctx, )
Get pending bookie client op stats by namesapce

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**map[string]PendingBookieOpsStats**](PendingBookieOpsStats.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopics2

> map[string]interface{} GetTopics2(ctx, )
Get all the topic stats by namespace

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

