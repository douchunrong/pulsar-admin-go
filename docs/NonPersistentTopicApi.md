# \NonPersistentTopicApi

All URIs are relative to *http://localhost/admin/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Compact**](NonPersistentTopicApi.md#Compact) | **Put** /non-persistent/{tenant}/{namespace}/{topic}/compaction | Trigger a compaction operation on a topic.
[**CompactionStatus**](NonPersistentTopicApi.md#CompactionStatus) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/compaction | Get the status of a compaction operation for a topic.
[**CreateNonPartitionedTopic**](NonPersistentTopicApi.md#CreateNonPartitionedTopic) | **Put** /non-persistent/{tenant}/{namespace}/{topic} | Create a non-partitioned topic.
[**CreatePartitionedTopic**](NonPersistentTopicApi.md#CreatePartitionedTopic) | **Put** /non-persistent/{tenant}/{namespace}/{topic}/partitions | Create a partitioned topic.
[**CreateSubscription**](NonPersistentTopicApi.md#CreateSubscription) | **Put** /non-persistent/{tenant}/{namespace}/{topic}/subscription/{subscriptionName} | Reset subscription to message position closest to given position.
[**DeletePartitionedTopic**](NonPersistentTopicApi.md#DeletePartitionedTopic) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/partitions | Delete a partitioned topic.
[**DeleteSubscription**](NonPersistentTopicApi.md#DeleteSubscription) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/subscription/{subName} | Delete a subscription.
[**DeleteTopic**](NonPersistentTopicApi.md#DeleteTopic) | **Delete** /non-persistent/{tenant}/{namespace}/{topic} | Delete a topic.
[**ExpireMessagesForAllSubscriptions**](NonPersistentTopicApi.md#ExpireMessagesForAllSubscriptions) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/all_subscription/expireMessages/{expireTimeInSeconds} | Expiry messages on all subscriptions of topic.
[**ExpireTopicMessages**](NonPersistentTopicApi.md#ExpireTopicMessages) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/subscription/{subName}/expireMessages/{expireTimeInSeconds} | Expiry messages on a topic subscription.
[**GetBacklog**](NonPersistentTopicApi.md#GetBacklog) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/backlog | Get estimated backlog for offline topic.
[**GetInternalStats**](NonPersistentTopicApi.md#GetInternalStats) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/internalStats | Get the internal stats for the topic.
[**GetLastMessageId**](NonPersistentTopicApi.md#GetLastMessageId) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/lastMessageId | Return the last commit message id of topic
[**GetList**](NonPersistentTopicApi.md#GetList) | **Get** /non-persistent/{tenant}/{namespace} | Get the list of non-persistent topics under a namespace.
[**GetListFromBundle**](NonPersistentTopicApi.md#GetListFromBundle) | **Get** /non-persistent/{tenant}/{namespace}/{bundle} | Get the list of non-persistent topics under a namespace bundle.
[**GetManagedLedgerInfo**](NonPersistentTopicApi.md#GetManagedLedgerInfo) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/internal-info | Get the internal stats for the topic.
[**GetPartitionedMetadata**](NonPersistentTopicApi.md#GetPartitionedMetadata) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/partitions | Get partitioned topic metadata.
[**GetPartitionedStats**](NonPersistentTopicApi.md#GetPartitionedStats) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/partitioned-stats | Get the stats for the partitioned topic.
[**GetPartitionedTopicList**](NonPersistentTopicApi.md#GetPartitionedTopicList) | **Get** /non-persistent/{tenant}/{namespace}/partitioned | Get the list of partitioned topics under a namespace.
[**GetPermissionsOnTopic**](NonPersistentTopicApi.md#GetPermissionsOnTopic) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/permissions | Get permissions on a topic.
[**GetStats**](NonPersistentTopicApi.md#GetStats) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/stats | Get the stats for the topic.
[**GetSubscriptions**](NonPersistentTopicApi.md#GetSubscriptions) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/subscriptions | Get the list of persistent subscriptions for a given topic.
[**GrantPermissionsOnTopic**](NonPersistentTopicApi.md#GrantPermissionsOnTopic) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/permissions/{role} | Grant a new permission to a role on a single topic.
[**OffloadStatus**](NonPersistentTopicApi.md#OffloadStatus) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/offload | Offload a prefix of a topic to long term storage
[**PeekNthMessage**](NonPersistentTopicApi.md#PeekNthMessage) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/subscription/{subName}/position/{messagePosition} | Peek nth message on a topic subscription.
[**ResetCursor**](NonPersistentTopicApi.md#ResetCursor) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/subscription/{subName}/resetcursor/{timestamp} | Reset subscription to message position closest to absolute timestamp (in ms).
[**ResetCursorOnPosition**](NonPersistentTopicApi.md#ResetCursorOnPosition) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/subscription/{subName}/resetcursor | Reset subscription to message position closest to given position.
[**RevokePermissionsOnTopic**](NonPersistentTopicApi.md#RevokePermissionsOnTopic) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/permissions/{role} | Revoke permissions on a topic.
[**SkipAllMessages**](NonPersistentTopicApi.md#SkipAllMessages) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/subscription/{subName}/skip_all | Skip all messages on a topic subscription.
[**SkipMessages**](NonPersistentTopicApi.md#SkipMessages) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/subscription/{subName}/skip/{numMessages} | Skipping messages on a topic subscription.
[**Terminate**](NonPersistentTopicApi.md#Terminate) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/terminate | Terminate a topic. A topic that is terminated will not accept any more messages to be published and will let consumer to drain existing messages in backlog
[**TriggerOffload**](NonPersistentTopicApi.md#TriggerOffload) | **Put** /non-persistent/{tenant}/{namespace}/{topic}/offload | Offload a prefix of a topic to long term storage
[**UnloadTopic**](NonPersistentTopicApi.md#UnloadTopic) | **Put** /non-persistent/{tenant}/{namespace}/{topic}/unload | Unload a topic
[**UpdatePartitionedTopic**](NonPersistentTopicApi.md#UpdatePartitionedTopic) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/partitions | Increment partitons of an existing partitioned topic.



## Compact

> Compact(ctx, tenant, namespace, topic, optional)
Trigger a compaction operation on a topic.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 
 **optional** | ***CompactOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CompactOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **optional.Bool**| Is authentication required to perform this operation | [default to false]

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


## CompactionStatus

> LongRunningProcessStatus CompactionStatus(ctx, tenant, namespace, topic, optional)
Get the status of a compaction operation for a topic.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 
 **optional** | ***CompactionStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CompactionStatusOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **optional.Bool**| Is authentication required to perform this operation | [default to false]

### Return type

[**LongRunningProcessStatus**](LongRunningProcessStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNonPartitionedTopic

> CreateNonPartitionedTopic(ctx, tenant, namespace, topic, optional)
Create a non-partitioned topic.

This is the only REST endpoint from which non-partitioned topics could be created.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 
 **optional** | ***CreateNonPartitionedTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateNonPartitionedTopicOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **optional.Bool**| Is authentication required to perform this operation | [default to false]

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


## CreatePartitionedTopic

> CreatePartitionedTopic(ctx, tenant, namespace, topic, body)
Create a partitioned topic.

It needs to be called before creating a producer on a partitioned topic.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 
**body** | **int32**| The number of partitions for the topic | 

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


## CreateSubscription

> CreateSubscription(ctx, tenant, namespace, topic, subscriptionName, optional)
Reset subscription to message position closest to given position.

Creates a subscription on the topic at the specified message id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 
**subscriptionName** | **string**| Subscription to create position on | 
 **optional** | ***CreateSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSubscriptionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authoritative** | **optional.Bool**| messageId where to create the subscription. It can be &#39;latest&#39;, &#39;earliest&#39; or (ledgerId:entryId) | [default to false]
 **replicated** | **optional.Bool**| Is authentication required to perform this operation | 

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


## DeletePartitionedTopic

> DeletePartitionedTopic(ctx, tenant, namespace, topic, optional)
Delete a partitioned topic.

It will also delete all the partitions of the topic if it exists.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 
 **optional** | ***DeletePartitionedTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeletePartitionedTopicOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **force** | **optional.Bool**| Stop all producer/consumer/replicator and delete topic forcefully | [default to false]
 **authoritative** | **optional.Bool**| Is authentication required to perform this operation | [default to false]

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


## DeleteSubscription

> DeleteSubscription(ctx, tenant, namespace, topic, subName, optional)
Delete a subscription.

There should not be any active consumers on the subscription.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 
**subName** | **string**| Subscription to be deleted | 
 **optional** | ***DeleteSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteSubscriptionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authoritative** | **optional.Bool**| Is authentication required to perform this operation | [default to false]

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


## DeleteTopic

> DeleteTopic(ctx, tenant, namespace, topic, optional)
Delete a topic.

The topic cannot be deleted if delete is not forcefully and there's any active subscription or producer connected to the it. Force delete ignores connected clients and deletes topic by explicitly closing them.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 
 **optional** | ***DeleteTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteTopicOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **force** | **optional.Bool**| Stop all producer/consumer/replicator and delete topic forcefully | [default to false]
 **authoritative** | **optional.Bool**| Is authentication required to perform this operation | [default to false]

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


## ExpireMessagesForAllSubscriptions

> ExpireMessagesForAllSubscriptions(ctx, tenant, namespace, topic, expireTimeInSeconds, optional)
Expiry messages on all subscriptions of topic.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 
**expireTimeInSeconds** | **int32**| Expires beyond the specified number of seconds | [default to 0]
 **optional** | ***ExpireMessagesForAllSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExpireMessagesForAllSubscriptionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authoritative** | **optional.Bool**| Is authentication required to perform this operation | [default to false]

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


## ExpireTopicMessages

> ExpireTopicMessages(ctx, tenant, namespace, topic, subName, expireTimeInSeconds, optional)
Expiry messages on a topic subscription.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 
**subName** | **string**| Subscription to be Expiry messages on | 
**expireTimeInSeconds** | **int32**| Expires beyond the specified number of seconds | [default to 0]
 **optional** | ***ExpireTopicMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExpireTopicMessagesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **authoritative** | **optional.Bool**| Is authentication required to perform this operation | [default to false]

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


## GetBacklog

> PersistentOfflineTopicStats GetBacklog(ctx, tenant, namespace, topic, optional)
Get estimated backlog for offline topic.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 
 **optional** | ***GetBacklogOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBacklogOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **optional.Bool**| Is authentication required to perform this operation | [default to false]

### Return type

[**PersistentOfflineTopicStats**](PersistentOfflineTopicStats.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInternalStats

> PersistentTopicInternalStats GetInternalStats(ctx, tenant, namespace, topic, optional)
Get the internal stats for the topic.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 
 **optional** | ***GetInternalStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInternalStatsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **optional.Bool**| Is authentication required to perform this operation | [default to false]

### Return type

[**PersistentTopicInternalStats**](PersistentTopicInternalStats.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLastMessageId

> map[string]interface{} GetLastMessageId(ctx, tenant, namespace, topic, optional)
Return the last commit message id of topic

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 
 **optional** | ***GetLastMessageIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLastMessageIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **optional.Bool**| Is authentication required to perform this operation | [default to false]

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


## GetList

> []string GetList(ctx, tenant, namespace)
Get the list of non-persistent topics under a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 

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


## GetListFromBundle

> []string GetListFromBundle(ctx, tenant, namespace, bundle)
Get the list of non-persistent topics under a namespace bundle.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 
**bundle** | **string**|  | 

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


## GetManagedLedgerInfo

> GetManagedLedgerInfo(ctx, tenant, namespace, topic)
Get the internal stats for the topic.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 

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


## GetPartitionedMetadata

> PartitionedTopicMetadata GetPartitionedMetadata(ctx, tenant, namespace, topic, optional)
Get partitioned topic metadata.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 
 **optional** | ***GetPartitionedMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPartitionedMetadataOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **optional.Bool**| Is authentication required to perform this operation | [default to false]

### Return type

[**PartitionedTopicMetadata**](PartitionedTopicMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartitionedStats

> GetPartitionedStats(ctx, tenant, namespace, topic, optional)
Get the stats for the partitioned topic.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 
 **optional** | ***GetPartitionedStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPartitionedStatsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **optional.Bool**| Is authentication required to perform this operation | [default to false]

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


## GetPartitionedTopicList

> []string GetPartitionedTopicList(ctx, tenant, namespace)
Get the list of partitioned topics under a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 

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


## GetPermissionsOnTopic

> map[string]map[string]interface{} GetPermissionsOnTopic(ctx, tenant, namespace, topic)
Get permissions on a topic.

Retrieve the effective permissions for a topic. These permissions are defined by the permissions set at thenamespace level combined (union) with any eventual specific permission set on the topic.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 

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


## GetStats

> TopicStats GetStats(ctx, tenant, namespace, topic, optional)
Get the stats for the topic.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 
 **optional** | ***GetStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetStatsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **optional.Bool**| Is authentication required to perform this operation | [default to false]

### Return type

[**TopicStats**](TopicStats.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriptions

> []map[string]interface{} GetSubscriptions(ctx, tenant, namespace, topic, optional)
Get the list of persistent subscriptions for a given topic.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 
 **optional** | ***GetSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSubscriptionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **optional.Bool**| Is authentication required to perform this operation | [default to false]

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


## GrantPermissionsOnTopic

> GrantPermissionsOnTopic(ctx, tenant, namespace, topic, role, optional)
Grant a new permission to a role on a single topic.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 
**role** | **string**| Client role to which grant permissions | 
 **optional** | ***GrantPermissionsOnTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GrantPermissionsOnTopicOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**optional.Interface of []string**](string.md)| Actions to be granted (produce,functions,consume) | 

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


## OffloadStatus

> OffloadProcessStatus OffloadStatus(ctx, tenant, namespace, topic, optional)
Offload a prefix of a topic to long term storage

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 
 **optional** | ***OffloadStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OffloadStatusOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **optional.Bool**| Is authentication required to perform this operation | [default to false]

### Return type

[**OffloadProcessStatus**](OffloadProcessStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PeekNthMessage

> PeekNthMessage(ctx, tenant, namespace, topic, subName, messagePosition, optional)
Peek nth message on a topic subscription.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 
**subName** | **string**| Subscribed message expired | 
**messagePosition** | **int32**| The number of messages (default 1) | [default to 1]
 **optional** | ***PeekNthMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PeekNthMessageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **authoritative** | **optional.Bool**| Is authentication required to perform this operation | [default to false]

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


## ResetCursor

> ResetCursor(ctx, tenant, namespace, topic, subName, timestamp, optional)
Reset subscription to message position closest to absolute timestamp (in ms).

It fence cursor and disconnects all active consumers before reseting cursor.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 
**subName** | **string**| Subscription to reset position on | 
**timestamp** | **int64**| time in minutes to reset back to (or minutes, hours,days,weeks eg:100m, 3h, 2d, 5w) | 
 **optional** | ***ResetCursorOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ResetCursorOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **authoritative** | **optional.Bool**| Is authentication required to perform this operation | [default to false]

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


## ResetCursorOnPosition

> ResetCursorOnPosition(ctx, tenant, namespace, topic, subName, optional)
Reset subscription to message position closest to given position.

It fence cursor and disconnects all active consumers before reseting cursor.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 
**subName** | **string**| Subscription to reset position on | 
 **optional** | ***ResetCursorOnPositionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ResetCursorOnPositionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authoritative** | **optional.Bool**| Is authentication required to perform this operation | [default to false]
 **messageId** | [**optional.Interface of MessageIdImpl**](MessageIdImpl.md)| messageId to reset back to (ledgerId:entryId) | 

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


## RevokePermissionsOnTopic

> RevokePermissionsOnTopic(ctx, tenant, namespace, topic, role)
Revoke permissions on a topic.

Revoke permissions to a role on a single topic. If the permission was not set at the topiclevel, but rather at the namespace level, this operation will return an error (HTTP status code 412).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 
**role** | **string**| Client role to which grant permissions | 

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


## SkipAllMessages

> SkipAllMessages(ctx, tenant, namespace, topic, subName, optional)
Skip all messages on a topic subscription.

Completely clears the backlog on the subscription.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 
**subName** | **string**| Name of subscription | 
 **optional** | ***SkipAllMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SkipAllMessagesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authoritative** | **optional.Bool**| Is authentication required to perform this operation | [default to false]

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


## SkipMessages

> SkipMessages(ctx, tenant, namespace, topic, subName, numMessages, optional)
Skipping messages on a topic subscription.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 
**subName** | **string**| Name of subscription | 
**numMessages** | **int32**| The number of messages to skip | [default to 0]
 **optional** | ***SkipMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SkipMessagesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **authoritative** | **optional.Bool**| Is authentication required to perform this operation | [default to false]

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


## Terminate

> map[string]interface{} Terminate(ctx, tenant, namespace, topic, optional)
Terminate a topic. A topic that is terminated will not accept any more messages to be published and will let consumer to drain existing messages in backlog

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 
 **optional** | ***TerminateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TerminateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **optional.Bool**| Is authentication required to perform this operation | [default to false]

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


## TriggerOffload

> TriggerOffload(ctx, tenant, namespace, topic, optional)
Offload a prefix of a topic to long term storage

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 
 **optional** | ***TriggerOffloadOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TriggerOffloadOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **optional.Bool**| Is authentication required to perform this operation | [default to false]

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


## UnloadTopic

> UnloadTopic(ctx, tenant, namespace, topic, optional)
Unload a topic

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 
 **optional** | ***UnloadTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UnloadTopicOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **optional.Bool**| Is authentication required to perform this operation | [default to false]

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


## UpdatePartitionedTopic

> UpdatePartitionedTopic(ctx, tenant, namespace, topic, body)
Increment partitons of an existing partitioned topic.

It only increments partitions of existing non-global partitioned-topic

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**| Specify the tenant | 
**namespace** | **string**| Specify the namespace | 
**topic** | **string**| Specify topic name | 
**body** | **int32**| The number of partitions for the topic | 

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

