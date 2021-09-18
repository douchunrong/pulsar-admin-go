# \NamespacesApi

All URIs are relative to *http://localhost/admin/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearNamespaceBacklog**](NamespacesApi.md#ClearNamespaceBacklog) | **Post** /namespaces/{tenant}/{namespace}/clearBacklog | Clear backlog for all topics on a namespace.
[**ClearNamespaceBacklogForSubscription**](NamespacesApi.md#ClearNamespaceBacklogForSubscription) | **Post** /namespaces/{tenant}/{namespace}/clearBacklog/{subscription} | Clear backlog for a given subscription on all topics on a namespace.
[**ClearNamespaceBundleBacklog**](NamespacesApi.md#ClearNamespaceBundleBacklog) | **Post** /namespaces/{tenant}/{namespace}/{bundle}/clearBacklog | Clear backlog for all topics on a namespace bundle.
[**ClearNamespaceBundleBacklogForSubscription**](NamespacesApi.md#ClearNamespaceBundleBacklogForSubscription) | **Post** /namespaces/{tenant}/{namespace}/{bundle}/clearBacklog/{subscription} | Clear backlog for a given subscription on all topics on a namespace bundle.
[**ClearOffloadDeletionLag**](NamespacesApi.md#ClearOffloadDeletionLag) | **Delete** /namespaces/{tenant}/{namespace}/offloadDeletionLagMs | Clear the namespace configured offload deletion lag. The topics in the namespace will fallback to using the default configured deletion lag for the broker
[**CreateNamespace**](NamespacesApi.md#CreateNamespace) | **Put** /namespaces/{tenant}/{namespace} | Creates a new namespace with the specified policies
[**DeleteBookieAffinityGroup**](NamespacesApi.md#DeleteBookieAffinityGroup) | **Delete** /namespaces/{property}/{namespace}/persistence/bookieAffinity | Delete the bookie-affinity-group from namespace-local policy.
[**DeleteNamespace**](NamespacesApi.md#DeleteNamespace) | **Delete** /namespaces/{tenant}/{namespace} | Delete a namespace and all the topics under it.
[**DeleteNamespaceBundle**](NamespacesApi.md#DeleteNamespaceBundle) | **Delete** /namespaces/{tenant}/{namespace}/{bundle} | Delete a namespace bundle and all the topics under it.
[**GetAntiAffinityNamespaces**](NamespacesApi.md#GetAntiAffinityNamespaces) | **Get** /namespaces/{cluster}/antiAffinity/{group} | Get all namespaces that are grouped by given anti-affinity group in a given cluster. api can be only accessed by admin of any of the existing tenant
[**GetBacklogQuotaMap**](NamespacesApi.md#GetBacklogQuotaMap) | **Get** /namespaces/{tenant}/{namespace}/backlogQuotaMap | Get backlog quota map on a namespace.
[**GetBookieAffinityGroup**](NamespacesApi.md#GetBookieAffinityGroup) | **Get** /namespaces/{property}/{namespace}/persistence/bookieAffinity | Get the bookie-affinity-group from namespace-local policy.
[**GetBundlesData**](NamespacesApi.md#GetBundlesData) | **Get** /namespaces/{tenant}/{namespace}/bundles | Get the bundles split data.
[**GetCompactionThreshold**](NamespacesApi.md#GetCompactionThreshold) | **Get** /namespaces/{tenant}/{namespace}/compactionThreshold | Maximum number of uncompacted bytes in topics before compaction is triggered.
[**GetDispatchRate**](NamespacesApi.md#GetDispatchRate) | **Get** /namespaces/{tenant}/{namespace}/dispatchRate | Get dispatch-rate configured for the namespace, -1 represents not configured yet
[**GetMaxConsumersPerSubscription**](NamespacesApi.md#GetMaxConsumersPerSubscription) | **Get** /namespaces/{tenant}/{namespace}/maxConsumersPerSubscription | Get maxConsumersPerSubscription config on a namespace.
[**GetMaxConsumersPerTopic**](NamespacesApi.md#GetMaxConsumersPerTopic) | **Get** /namespaces/{tenant}/{namespace}/maxConsumersPerTopic | Get maxConsumersPerTopic config on a namespace.
[**GetMaxProducersPerTopic**](NamespacesApi.md#GetMaxProducersPerTopic) | **Get** /namespaces/{tenant}/{namespace}/maxProducersPerTopic | Get maxProducersPerTopic config on a namespace.
[**GetNamespaceAntiAffinityGroup**](NamespacesApi.md#GetNamespaceAntiAffinityGroup) | **Get** /namespaces/{tenant}/{namespace}/antiAffinity | Get anti-affinity group of a namespace.
[**GetNamespaceMessageTTL**](NamespacesApi.md#GetNamespaceMessageTTL) | **Get** /namespaces/{tenant}/{namespace}/messageTTL | Get the message TTL for the namespace
[**GetNamespaceReplicationClusters**](NamespacesApi.md#GetNamespaceReplicationClusters) | **Get** /namespaces/{tenant}/{namespace}/replication | Get the replication clusters for a namespace.
[**GetOffloadDeletionLag**](NamespacesApi.md#GetOffloadDeletionLag) | **Get** /namespaces/{tenant}/{namespace}/offloadDeletionLagMs | Number of milliseconds to wait before deleting a ledger segment which has been offloaded from the Pulsar cluster&#39;s local storage (i.e. BookKeeper)
[**GetOffloadThreshold**](NamespacesApi.md#GetOffloadThreshold) | **Get** /namespaces/{tenant}/{namespace}/offloadThreshold | Maximum number of bytes stored on the pulsar cluster for a topic, before the broker will start offloading to longterm storage
[**GetPermissions**](NamespacesApi.md#GetPermissions) | **Get** /namespaces/{tenant}/{namespace}/permissions | Retrieve the permissions for a namespace.
[**GetPersistence**](NamespacesApi.md#GetPersistence) | **Get** /namespaces/{tenant}/{namespace}/persistence | Get the persistence configuration for a namespace.
[**GetPolicies**](NamespacesApi.md#GetPolicies) | **Get** /namespaces/{tenant}/{namespace} | Get the dump all the policies specified for a namespace.
[**GetReplicatorDispatchRate**](NamespacesApi.md#GetReplicatorDispatchRate) | **Get** /namespaces/{tenant}/{namespace}/replicatorDispatchRate | Get replicator dispatch-rate configured for the namespace, -1 represents not configured yet
[**GetRetention**](NamespacesApi.md#GetRetention) | **Get** /namespaces/{tenant}/{namespace}/retention | Get retention config on a namespace.
[**GetSchemaAutoUpdateCompatibilityStrategy**](NamespacesApi.md#GetSchemaAutoUpdateCompatibilityStrategy) | **Get** /namespaces/{tenant}/{namespace}/schemaAutoUpdateCompatibilityStrategy | The strategy used to check the compatibility of new schemas, provided by producers, before automatically updating the schema
[**GetSchemaValidtionEnforced**](NamespacesApi.md#GetSchemaValidtionEnforced) | **Get** /namespaces/{tenant}/{namespace}/schemaValidationEnforced | Get schema validation enforced flag for namespace.
[**GetSubscribeRate**](NamespacesApi.md#GetSubscribeRate) | **Get** /namespaces/{tenant}/{namespace}/subscribeRate | Get subscribe-rate configured for the namespace
[**GetSubscriptionDispatchRate**](NamespacesApi.md#GetSubscriptionDispatchRate) | **Get** /namespaces/{tenant}/{namespace}/subscriptionDispatchRate | Get Subscription dispatch-rate configured for the namespace, -1 represents not configured yet
[**GetTenantNamespaces**](NamespacesApi.md#GetTenantNamespaces) | **Get** /namespaces/{tenant} | Get the list of all the namespaces for a certain tenant.
[**GetTopics**](NamespacesApi.md#GetTopics) | **Get** /namespaces/{tenant}/{namespace}/topics | Get the list of all the topics under a certain namespace.
[**GrantPermissionOnNamespace**](NamespacesApi.md#GrantPermissionOnNamespace) | **Post** /namespaces/{tenant}/{namespace}/permissions/{role} | Grant a new permission to a role on a namespace.
[**ModifyDeduplication**](NamespacesApi.md#ModifyDeduplication) | **Post** /namespaces/{tenant}/{namespace}/deduplication | Enable or disable broker side deduplication for all topics in a namespace
[**ModifyEncryptionRequired**](NamespacesApi.md#ModifyEncryptionRequired) | **Post** /namespaces/{tenant}/{namespace}/encryptionRequired | Message encryption is required or not for all topics in a namespace
[**RemoveBacklogQuota**](NamespacesApi.md#RemoveBacklogQuota) | **Delete** /namespaces/{tenant}/{namespace}/backlogQuota | Remove a backlog quota policy from a namespace.
[**RemoveNamespaceAntiAffinityGroup**](NamespacesApi.md#RemoveNamespaceAntiAffinityGroup) | **Delete** /namespaces/{tenant}/{namespace}/antiAffinity | Remove anti-affinity group of a namespace.
[**RevokePermissionsOnNamespace**](NamespacesApi.md#RevokePermissionsOnNamespace) | **Delete** /namespaces/{tenant}/{namespace}/permissions/{role} | Revoke all permissions to a role on a namespace.
[**SetBacklogQuota**](NamespacesApi.md#SetBacklogQuota) | **Post** /namespaces/{tenant}/{namespace}/backlogQuota |  Set a backlog quota for all the topics on a namespace.
[**SetBookieAffinityGroup**](NamespacesApi.md#SetBookieAffinityGroup) | **Post** /namespaces/{tenant}/{namespace}/persistence/bookieAffinity | Set the bookie-affinity-group to namespace-persistent policy.
[**SetCompactionThreshold**](NamespacesApi.md#SetCompactionThreshold) | **Put** /namespaces/{tenant}/{namespace}/compactionThreshold | Set maximum number of uncompacted bytes in a topic before compaction is triggered.
[**SetDispatchRate**](NamespacesApi.md#SetDispatchRate) | **Post** /namespaces/{tenant}/{namespace}/dispatchRate | Set dispatch-rate throttling for all topics of the namespace
[**SetMaxConsumersPerSubscription**](NamespacesApi.md#SetMaxConsumersPerSubscription) | **Post** /namespaces/{tenant}/{namespace}/maxConsumersPerSubscription |  Set maxConsumersPerSubscription configuration on a namespace.
[**SetMaxConsumersPerTopic**](NamespacesApi.md#SetMaxConsumersPerTopic) | **Post** /namespaces/{tenant}/{namespace}/maxConsumersPerTopic |  Set maxConsumersPerTopic configuration on a namespace.
[**SetMaxProducersPerTopic**](NamespacesApi.md#SetMaxProducersPerTopic) | **Post** /namespaces/{tenant}/{namespace}/maxProducersPerTopic |  Set maxProducersPerTopic configuration on a namespace.
[**SetNamespaceAntiAffinityGroup**](NamespacesApi.md#SetNamespaceAntiAffinityGroup) | **Post** /namespaces/{tenant}/{namespace}/antiAffinity | Set anti-affinity group for a namespace
[**SetNamespaceMessageTTL**](NamespacesApi.md#SetNamespaceMessageTTL) | **Post** /namespaces/{tenant}/{namespace}/messageTTL | Set message TTL in seconds for namespace
[**SetNamespaceReplicationClusters**](NamespacesApi.md#SetNamespaceReplicationClusters) | **Post** /namespaces/{tenant}/{namespace}/replication | Set the replication clusters for a namespace.
[**SetOffloadDeletionLag**](NamespacesApi.md#SetOffloadDeletionLag) | **Put** /namespaces/{tenant}/{namespace}/offloadDeletionLagMs | Set number of milliseconds to wait before deleting a ledger segment which has been offloaded from the Pulsar cluster&#39;s local storage (i.e. BookKeeper)
[**SetOffloadThreshold**](NamespacesApi.md#SetOffloadThreshold) | **Put** /namespaces/{tenant}/{namespace}/offloadThreshold | Set maximum number of bytes stored on the pulsar cluster for a topic, before the broker will start offloading to longterm storage
[**SetPersistence**](NamespacesApi.md#SetPersistence) | **Post** /namespaces/{tenant}/{namespace}/persistence | Set the persistence configuration for all the topics on a namespace.
[**SetReplicatorDispatchRate**](NamespacesApi.md#SetReplicatorDispatchRate) | **Post** /namespaces/{tenant}/{namespace}/replicatorDispatchRate | Set replicator dispatch-rate throttling for all topics of the namespace
[**SetRetention**](NamespacesApi.md#SetRetention) | **Post** /namespaces/{tenant}/{namespace}/retention |  Set retention configuration on a namespace.
[**SetSchemaAutoUpdateCompatibilityStrategy**](NamespacesApi.md#SetSchemaAutoUpdateCompatibilityStrategy) | **Put** /namespaces/{tenant}/{namespace}/schemaAutoUpdateCompatibilityStrategy | Update the strategy used to check the compatibility of new schemas, provided by producers, before automatically updating the schema
[**SetSchemaValidtionEnforced**](NamespacesApi.md#SetSchemaValidtionEnforced) | **Post** /namespaces/{tenant}/{namespace}/schemaValidationEnforced | Set schema validation enforced flag on namespace.
[**SetSubscribeRate**](NamespacesApi.md#SetSubscribeRate) | **Post** /namespaces/{tenant}/{namespace}/subscribeRate | Set subscribe-rate throttling for all topics of the namespace
[**SetSubscriptionAuthMode**](NamespacesApi.md#SetSubscriptionAuthMode) | **Post** /namespaces/{tenant}/{namespace}/subscriptionAuthMode |  Set a subscription auth mode for all the topics on a namespace.
[**SetSubscriptionDispatchRate**](NamespacesApi.md#SetSubscriptionDispatchRate) | **Post** /namespaces/{tenant}/{namespace}/subscriptionDispatchRate | Set Subscription dispatch-rate throttling for all topics of the namespace
[**SplitNamespaceBundle**](NamespacesApi.md#SplitNamespaceBundle) | **Put** /namespaces/{tenant}/{namespace}/{bundle}/split | Split a namespace bundle
[**UnloadNamespace**](NamespacesApi.md#UnloadNamespace) | **Put** /namespaces/{tenant}/{namespace}/unload | Unload namespace
[**UnloadNamespaceBundle**](NamespacesApi.md#UnloadNamespaceBundle) | **Put** /namespaces/{tenant}/{namespace}/{bundle}/unload | Unload a namespace bundle
[**UnsubscribeNamespace**](NamespacesApi.md#UnsubscribeNamespace) | **Post** /namespaces/{tenant}/{namespace}/unsubscribe/{subscription} | Unsubscribes the given subscription on all topics on a namespace.
[**UnsubscribeNamespaceBundle**](NamespacesApi.md#UnsubscribeNamespaceBundle) | **Post** /namespaces/{tenant}/{namespace}/{bundle}/unsubscribe/{subscription} | Unsubscribes the given subscription on all topics on a namespace bundle.



## ClearNamespaceBacklog

> ClearNamespaceBacklog(ctx, tenant, namespace, optional)
Clear backlog for all topics on a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 
 **optional** | ***ClearNamespaceBacklogOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClearNamespaceBacklogOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authoritative** | **optional.Bool**|  | [default to false]

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


## ClearNamespaceBacklogForSubscription

> ClearNamespaceBacklogForSubscription(ctx, tenant, namespace, subscription, optional)
Clear backlog for a given subscription on all topics on a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 
**subscription** | **string**|  | 
 **optional** | ***ClearNamespaceBacklogForSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClearNamespaceBacklogForSubscriptionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **optional.Bool**|  | [default to false]

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


## ClearNamespaceBundleBacklog

> ClearNamespaceBundleBacklog(ctx, tenant, namespace, bundle, optional)
Clear backlog for all topics on a namespace bundle.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 
**bundle** | **string**|  | 
 **optional** | ***ClearNamespaceBundleBacklogOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClearNamespaceBundleBacklogOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **optional.Bool**|  | [default to false]

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


## ClearNamespaceBundleBacklogForSubscription

> ClearNamespaceBundleBacklogForSubscription(ctx, tenant, namespace, subscription, bundle, optional)
Clear backlog for a given subscription on all topics on a namespace bundle.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 
**subscription** | **string**|  | 
**bundle** | **string**|  | 
 **optional** | ***ClearNamespaceBundleBacklogForSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClearNamespaceBundleBacklogForSubscriptionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authoritative** | **optional.Bool**|  | [default to false]

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


## ClearOffloadDeletionLag

> ClearOffloadDeletionLag(ctx, tenant, namespace)
Clear the namespace configured offload deletion lag. The topics in the namespace will fallback to using the default configured deletion lag for the broker

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

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


## CreateNamespace

> CreateNamespace(ctx, tenant, namespace)
Creates a new namespace with the specified policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

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


## DeleteBookieAffinityGroup

> DeleteBookieAffinityGroup(ctx, property, namespace)
Delete the bookie-affinity-group from namespace-local policy.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**property** | **string**|  | 
**namespace** | **string**|  | 

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


## DeleteNamespace

> DeleteNamespace(ctx, tenant, namespace, optional)
Delete a namespace and all the topics under it.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 
 **optional** | ***DeleteNamespaceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteNamespaceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authoritative** | **optional.Bool**|  | [default to false]

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


## DeleteNamespaceBundle

> DeleteNamespaceBundle(ctx, tenant, namespace, bundle, optional)
Delete a namespace bundle and all the topics under it.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 
**bundle** | **string**|  | 
 **optional** | ***DeleteNamespaceBundleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteNamespaceBundleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **optional.Bool**|  | [default to false]

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


## GetAntiAffinityNamespaces

> []map[string]interface{} GetAntiAffinityNamespaces(ctx, cluster, group, optional)
Get all namespaces that are grouped by given anti-affinity group in a given cluster. api can be only accessed by admin of any of the existing tenant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string**|  | 
**group** | **string**|  | 
 **optional** | ***GetAntiAffinityNamespacesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAntiAffinityNamespacesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tenant** | **optional.String**|  | 

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


## GetBacklogQuotaMap

> map[string]map[string]interface{} GetBacklogQuotaMap(ctx, tenant, namespace)
Get backlog quota map on a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

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


## GetBookieAffinityGroup

> BookieAffinityGroupData GetBookieAffinityGroup(ctx, property, namespace)
Get the bookie-affinity-group from namespace-local policy.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**property** | **string**|  | 
**namespace** | **string**|  | 

### Return type

[**BookieAffinityGroupData**](BookieAffinityGroupData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBundlesData

> BundlesData GetBundlesData(ctx, tenant, namespace)
Get the bundles split data.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

### Return type

[**BundlesData**](BundlesData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompactionThreshold

> int64 GetCompactionThreshold(ctx, tenant, namespace)
Maximum number of uncompacted bytes in topics before compaction is triggered.

The backlog size is compared to the threshold periodically. A threshold of 0 disabled automatic compaction

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

### Return type

**int64**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDispatchRate

> DispatchRate GetDispatchRate(ctx, tenant, namespace)
Get dispatch-rate configured for the namespace, -1 represents not configured yet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

### Return type

[**DispatchRate**](DispatchRate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMaxConsumersPerSubscription

> int32 GetMaxConsumersPerSubscription(ctx, tenant, namespace)
Get maxConsumersPerSubscription config on a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMaxConsumersPerTopic

> int32 GetMaxConsumersPerTopic(ctx, tenant, namespace)
Get maxConsumersPerTopic config on a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMaxProducersPerTopic

> int32 GetMaxProducersPerTopic(ctx, tenant, namespace)
Get maxProducersPerTopic config on a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespaceAntiAffinityGroup

> string GetNamespaceAntiAffinityGroup(ctx, tenant, namespace)
Get anti-affinity group of a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespaceMessageTTL

> int32 GetNamespaceMessageTTL(ctx, tenant, namespace)
Get the message TTL for the namespace

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespaceReplicationClusters

> []string GetNamespaceReplicationClusters(ctx, tenant, namespace)
Get the replication clusters for a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

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


## GetOffloadDeletionLag

> int64 GetOffloadDeletionLag(ctx, tenant, namespace)
Number of milliseconds to wait before deleting a ledger segment which has been offloaded from the Pulsar cluster's local storage (i.e. BookKeeper)

A negative value denotes that deletion has been completely disabled. 'null' denotes that the topics in the namespace will fall back to the broker default for deletion lag.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

### Return type

**int64**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOffloadThreshold

> int64 GetOffloadThreshold(ctx, tenant, namespace)
Maximum number of bytes stored on the pulsar cluster for a topic, before the broker will start offloading to longterm storage

A negative value disables automatic offloading

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

### Return type

**int64**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPermissions

> map[string]map[string]interface{} GetPermissions(ctx, tenant, cluster, namespace)
Retrieve the permissions for a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**cluster** | **string**|  | 
**namespace** | **string**|  | 

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


## GetPersistence

> PersistencePolicies GetPersistence(ctx, tenant, namespace)
Get the persistence configuration for a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

### Return type

[**PersistencePolicies**](PersistencePolicies.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicies

> Policies GetPolicies(ctx, tenant, namespace)
Get the dump all the policies specified for a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

### Return type

[**Policies**](Policies.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReplicatorDispatchRate

> DispatchRate GetReplicatorDispatchRate(ctx, tenant, namespace)
Get replicator dispatch-rate configured for the namespace, -1 represents not configured yet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

### Return type

[**DispatchRate**](DispatchRate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRetention

> RetentionPolicies GetRetention(ctx, tenant, namespace)
Get retention config on a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

### Return type

[**RetentionPolicies**](RetentionPolicies.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchemaAutoUpdateCompatibilityStrategy

> string GetSchemaAutoUpdateCompatibilityStrategy(ctx, tenant, namespace)
The strategy used to check the compatibility of new schemas, provided by producers, before automatically updating the schema

The value AutoUpdateDisabled prevents producers from updating the schema.  If set to AutoUpdateDisabled, schemas must be updated through the REST api

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchemaValidtionEnforced

> bool GetSchemaValidtionEnforced(ctx, tenant, namespace)
Get schema validation enforced flag for namespace.

If the flag is set to true, when a producer without a schema attempts to produce to a topic with schema in this namespace, the producer will be failed to connect. PLEASE be carefully on using this, since non-java clients don't support schema.if you enable this setting, it will cause non-java clients failed to produce.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscribeRate

> SubscribeRate GetSubscribeRate(ctx, tenant, namespace)
Get subscribe-rate configured for the namespace

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

### Return type

[**SubscribeRate**](SubscribeRate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriptionDispatchRate

> DispatchRate GetSubscriptionDispatchRate(ctx, tenant, namespace)
Get Subscription dispatch-rate configured for the namespace, -1 represents not configured yet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

### Return type

[**DispatchRate**](DispatchRate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantNamespaces

> []string GetTenantNamespaces(ctx, tenant)
Get the list of all the namespaces for a certain tenant.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 

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


## GetTopics

> []string GetTopics(ctx, tenant, namespace, optional)
Get the list of all the topics under a certain namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 
 **optional** | ***GetTopicsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTopicsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mode** | **optional.String**|  | [default to PERSISTENT]

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


## GrantPermissionOnNamespace

> GrantPermissionOnNamespace(ctx, tenant, namespace, role)
Grant a new permission to a role on a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 
**role** | **string**|  | 

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


## ModifyDeduplication

> ModifyDeduplication(ctx, tenant, namespace)
Enable or disable broker side deduplication for all topics in a namespace

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

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


## ModifyEncryptionRequired

> ModifyEncryptionRequired(ctx, tenant, namespace)
Message encryption is required or not for all topics in a namespace

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

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


## RemoveBacklogQuota

> RemoveBacklogQuota(ctx, tenant, namespace, optional)
Remove a backlog quota policy from a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 
 **optional** | ***RemoveBacklogQuotaOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveBacklogQuotaOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **backlogQuotaType** | **optional.String**|  | 

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


## RemoveNamespaceAntiAffinityGroup

> RemoveNamespaceAntiAffinityGroup(ctx, tenant, namespace)
Remove anti-affinity group of a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

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


## RevokePermissionsOnNamespace

> RevokePermissionsOnNamespace(ctx, tenant, namespace, role)
Revoke all permissions to a role on a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 
**role** | **string**|  | 

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


## SetBacklogQuota

> SetBacklogQuota(ctx, tenant, namespace, optional)
 Set a backlog quota for all the topics on a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 
 **optional** | ***SetBacklogQuotaOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetBacklogQuotaOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **backlogQuotaType** | **optional.String**|  | 

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


## SetBookieAffinityGroup

> SetBookieAffinityGroup(ctx, tenant, namespace)
Set the bookie-affinity-group to namespace-persistent policy.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

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


## SetCompactionThreshold

> SetCompactionThreshold(ctx, tenant, namespace)
Set maximum number of uncompacted bytes in a topic before compaction is triggered.

The backlog size is compared to the threshold periodically. A threshold of 0 disabled automatic compaction

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

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


## SetDispatchRate

> SetDispatchRate(ctx, tenant, namespace)
Set dispatch-rate throttling for all topics of the namespace

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

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


## SetMaxConsumersPerSubscription

> SetMaxConsumersPerSubscription(ctx, tenant, namespace)
 Set maxConsumersPerSubscription configuration on a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

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


## SetMaxConsumersPerTopic

> SetMaxConsumersPerTopic(ctx, tenant, namespace)
 Set maxConsumersPerTopic configuration on a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

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


## SetMaxProducersPerTopic

> SetMaxProducersPerTopic(ctx, tenant, namespace)
 Set maxProducersPerTopic configuration on a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

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


## SetNamespaceAntiAffinityGroup

> SetNamespaceAntiAffinityGroup(ctx, tenant, namespace)
Set anti-affinity group for a namespace

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

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


## SetNamespaceMessageTTL

> SetNamespaceMessageTTL(ctx, tenant, namespace)
Set message TTL in seconds for namespace

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

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


## SetNamespaceReplicationClusters

> SetNamespaceReplicationClusters(ctx, tenant, namespace)
Set the replication clusters for a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

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


## SetOffloadDeletionLag

> SetOffloadDeletionLag(ctx, tenant, namespace)
Set number of milliseconds to wait before deleting a ledger segment which has been offloaded from the Pulsar cluster's local storage (i.e. BookKeeper)

A negative value disables the deletion completely.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

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


## SetOffloadThreshold

> SetOffloadThreshold(ctx, tenant, namespace)
Set maximum number of bytes stored on the pulsar cluster for a topic, before the broker will start offloading to longterm storage

A negative value disables automatic offloading

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

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


## SetPersistence

> SetPersistence(ctx, tenant, namespace)
Set the persistence configuration for all the topics on a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

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


## SetReplicatorDispatchRate

> SetReplicatorDispatchRate(ctx, tenant, namespace)
Set replicator dispatch-rate throttling for all topics of the namespace

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

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


## SetRetention

> SetRetention(ctx, tenant, namespace)
 Set retention configuration on a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

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


## SetSchemaAutoUpdateCompatibilityStrategy

> SetSchemaAutoUpdateCompatibilityStrategy(ctx, tenant, namespace)
Update the strategy used to check the compatibility of new schemas, provided by producers, before automatically updating the schema

The value AutoUpdateDisabled prevents producers from updating the schema.  If set to AutoUpdateDisabled, schemas must be updated through the REST api

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

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


## SetSchemaValidtionEnforced

> SetSchemaValidtionEnforced(ctx, tenant, namespace)
Set schema validation enforced flag on namespace.

If the flag is set to true, when a producer without a schema attempts to produce to a topic with schema in this namespace, the producer will be failed to connect. PLEASE be carefully on using this, since non-java clients don't support schema.if you enable this setting, it will cause non-java clients failed to produce.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

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


## SetSubscribeRate

> SetSubscribeRate(ctx, tenant, namespace)
Set subscribe-rate throttling for all topics of the namespace

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

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


## SetSubscriptionAuthMode

> SetSubscriptionAuthMode(ctx, tenant, namespace)
 Set a subscription auth mode for all the topics on a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

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


## SetSubscriptionDispatchRate

> SetSubscriptionDispatchRate(ctx, tenant, namespace)
Set Subscription dispatch-rate throttling for all topics of the namespace

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

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


## SplitNamespaceBundle

> SplitNamespaceBundle(ctx, tenant, namespace, bundle, optional)
Split a namespace bundle

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 
**bundle** | **string**|  | 
 **optional** | ***SplitNamespaceBundleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SplitNamespaceBundleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **optional.Bool**|  | [default to false]
 **unload** | **optional.Bool**|  | [default to false]

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


## UnloadNamespace

> UnloadNamespace(ctx, tenant, namespace)
Unload namespace

Unload an active namespace from the current broker serving it. Performing this operation will let the brokerremoves all producers, consumers, and connections using this namespace, and close all topics (includingtheir persistent store). During that operation, the namespace is marked as tentatively unavailable until thebroker completes the unloading action. This operation requires strictly super user privileges, since it wouldresult in non-persistent message loss and unexpected connection closure to the clients.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 

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


## UnloadNamespaceBundle

> UnloadNamespaceBundle(ctx, tenant, namespace, bundle, optional)
Unload a namespace bundle

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 
**bundle** | **string**|  | 
 **optional** | ***UnloadNamespaceBundleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UnloadNamespaceBundleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **optional.Bool**|  | [default to false]

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


## UnsubscribeNamespace

> UnsubscribeNamespace(ctx, tenant, cluster, namespace, subscription, optional)
Unsubscribes the given subscription on all topics on a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**cluster** | **string**|  | 
**namespace** | **string**|  | 
**subscription** | **string**|  | 
 **optional** | ***UnsubscribeNamespaceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UnsubscribeNamespaceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authoritative** | **optional.Bool**|  | [default to false]

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


## UnsubscribeNamespaceBundle

> UnsubscribeNamespaceBundle(ctx, tenant, namespace, subscription, bundle, optional)
Unsubscribes the given subscription on all topics on a namespace bundle.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string**|  | 
**namespace** | **string**|  | 
**subscription** | **string**|  | 
**bundle** | **string**|  | 
 **optional** | ***UnsubscribeNamespaceBundleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UnsubscribeNamespaceBundleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authoritative** | **optional.Bool**|  | [default to false]

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

