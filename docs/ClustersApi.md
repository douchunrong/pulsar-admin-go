# \ClustersApi

All URIs are relative to *http://localhost/admin/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCluster**](ClustersApi.md#CreateCluster) | **Put** /clusters/{cluster} | Create a new cluster.
[**DeleteCluster**](ClustersApi.md#DeleteCluster) | **Delete** /clusters/{cluster} | Delete an existing cluster.
[**DeleteFailureDomain**](ClustersApi.md#DeleteFailureDomain) | **Delete** /clusters/{cluster}/failureDomains/{domainName} | Delete the failure domain of the cluster
[**DeleteNamespaceIsolationPolicy**](ClustersApi.md#DeleteNamespaceIsolationPolicy) | **Delete** /clusters/{cluster}/namespaceIsolationPolicies/{policyName} | Delete namespace isolation policy.
[**GetBrokerWithNamespaceIsolationPolicy**](ClustersApi.md#GetBrokerWithNamespaceIsolationPolicy) | **Get** /clusters/{cluster}/namespaceIsolationPolicies/brokers/{broker} | Get a broker with namespace-isolation policies attached to it.
[**GetBrokersWithNamespaceIsolationPolicy**](ClustersApi.md#GetBrokersWithNamespaceIsolationPolicy) | **Get** /clusters/{cluster}/namespaceIsolationPolicies/brokers | Get list of brokers with namespace-isolation policies attached to them.
[**GetCluster**](ClustersApi.md#GetCluster) | **Get** /clusters/{cluster} | Get the configuration for the specified cluster.
[**GetClusters**](ClustersApi.md#GetClusters) | **Get** /clusters | Get the list of all the Pulsar clusters.
[**GetDomain**](ClustersApi.md#GetDomain) | **Get** /clusters/{cluster}/failureDomains/{domainName} | Get a domain in a cluster
[**GetFailureDomains**](ClustersApi.md#GetFailureDomains) | **Get** /clusters/{cluster}/failureDomains | Get the cluster failure domains.
[**GetNamespaceIsolationPolicies**](ClustersApi.md#GetNamespaceIsolationPolicies) | **Get** /clusters/{cluster}/namespaceIsolationPolicies | Get the namespace isolation policies assigned to the cluster.
[**GetNamespaceIsolationPolicy**](ClustersApi.md#GetNamespaceIsolationPolicy) | **Get** /clusters/{cluster}/namespaceIsolationPolicies/{policyName} | Get the single namespace isolation policy assigned to the cluster.
[**GetPeerCluster**](ClustersApi.md#GetPeerCluster) | **Get** /clusters/{cluster}/peers | Get the peer-cluster data for the specified cluster.
[**SetFailureDomain**](ClustersApi.md#SetFailureDomain) | **Post** /clusters/{cluster}/failureDomains/{domainName} | Set the failure domain of the cluster.
[**SetNamespaceIsolationPolicy**](ClustersApi.md#SetNamespaceIsolationPolicy) | **Post** /clusters/{cluster}/namespaceIsolationPolicies/{policyName} | Set namespace isolation policy.
[**SetPeerClusterNames**](ClustersApi.md#SetPeerClusterNames) | **Post** /clusters/{cluster}/peers | Update peer-cluster-list for a cluster.
[**UpdateCluster**](ClustersApi.md#UpdateCluster) | **Post** /clusters/{cluster} | Update the configuration for a cluster.



## CreateCluster

> CreateCluster(ctx, cluster, body)
Create a new cluster.

This operation requires Pulsar superuser privileges, and the name cannot contain the '/' characters.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string**| The cluster name | 
**body** | [**ClusterData**](ClusterData.md)| The cluster data | 

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


## DeleteCluster

> DeleteCluster(ctx, cluster)
Delete an existing cluster.

This operation requires Pulsar superuser privileges.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string**| The cluster name | 

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


## DeleteFailureDomain

> DeleteFailureDomain(ctx, cluster, domainName)
Delete the failure domain of the cluster

This operation requires Pulsar superuser privileges.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string**| The cluster name | 
**domainName** | **string**| The failure domain name | 

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


## DeleteNamespaceIsolationPolicy

> DeleteNamespaceIsolationPolicy(ctx, cluster, policyName)
Delete namespace isolation policy.

This operation requires Pulsar superuser privileges.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string**| The cluster name | 
**policyName** | **string**| The namespace isolation policy name | 

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


## GetBrokerWithNamespaceIsolationPolicy

> BrokerNamespaceIsolationData GetBrokerWithNamespaceIsolationPolicy(ctx, cluster, broker)
Get a broker with namespace-isolation policies attached to it.

This operation requires Pulsar superuser privileges.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string**| The cluster name | 
**broker** | **string**| The broker name (&lt;broker-hostname&gt;:&lt;web-service-port&gt;) | 

### Return type

[**BrokerNamespaceIsolationData**](BrokerNamespaceIsolationData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrokersWithNamespaceIsolationPolicy

> []BrokerNamespaceIsolationData GetBrokersWithNamespaceIsolationPolicy(ctx, cluster)
Get list of brokers with namespace-isolation policies attached to them.

This operation requires Pulsar superuser privileges.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string**| The cluster name | 

### Return type

[**[]BrokerNamespaceIsolationData**](BrokerNamespaceIsolationData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCluster

> ClusterData GetCluster(ctx, cluster)
Get the configuration for the specified cluster.

This operation requires Pulsar superuser privileges.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string**| The cluster name | 

### Return type

[**ClusterData**](ClusterData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusters

> []string GetClusters(ctx, )
Get the list of all the Pulsar clusters.

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


## GetDomain

> FailureDomain GetDomain(ctx, cluster, domainName)
Get a domain in a cluster

This operation requires Pulsar superuser privileges.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string**| The cluster name | 
**domainName** | **string**| The failure domain name | 

### Return type

[**FailureDomain**](FailureDomain.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFailureDomains

> map[string]FailureDomain GetFailureDomains(ctx, cluster)
Get the cluster failure domains.

This operation requires Pulsar superuser privileges.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string**| The cluster name | 

### Return type

[**map[string]FailureDomain**](FailureDomain.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespaceIsolationPolicies

> map[string]NamespaceIsolationData GetNamespaceIsolationPolicies(ctx, cluster)
Get the namespace isolation policies assigned to the cluster.

This operation requires Pulsar superuser privileges.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string**| The cluster name | 

### Return type

[**map[string]NamespaceIsolationData**](NamespaceIsolationData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespaceIsolationPolicy

> NamespaceIsolationData GetNamespaceIsolationPolicy(ctx, cluster, policyName)
Get the single namespace isolation policy assigned to the cluster.

This operation requires Pulsar superuser privileges.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string**| The cluster name | 
**policyName** | **string**| The name of the namespace isolation policy | 

### Return type

[**NamespaceIsolationData**](NamespaceIsolationData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPeerCluster

> []string GetPeerCluster(ctx, cluster)
Get the peer-cluster data for the specified cluster.

This operation requires Pulsar superuser privileges.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string**| The cluster name | 

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


## SetFailureDomain

> SetFailureDomain(ctx, cluster, domainName, body)
Set the failure domain of the cluster.

This operation requires Pulsar superuser privileges.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string**| The cluster name | 
**domainName** | **string**| The failure domain name | 
**body** | [**FailureDomain**](FailureDomain.md)| The configuration data of a failure domain | 

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


## SetNamespaceIsolationPolicy

> SetNamespaceIsolationPolicy(ctx, cluster, policyName, body)
Set namespace isolation policy.

This operation requires Pulsar superuser privileges.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string**| The cluster name | 
**policyName** | **string**| The namespace isolation policy name | 
**body** | [**NamespaceIsolationData**](NamespaceIsolationData.md)| The namespace isolation policy data | 

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


## SetPeerClusterNames

> SetPeerClusterNames(ctx, cluster, body)
Update peer-cluster-list for a cluster.

This operation requires Pulsar superuser privileges.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string**| The cluster name | 
**body** | [**[]string**](string.md)| The list of peer cluster names | 

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


## UpdateCluster

> UpdateCluster(ctx, cluster, body)
Update the configuration for a cluster.

This operation requires Pulsar superuser privileges.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string**| The cluster name | 
**body** | [**ClusterData**](ClusterData.md)| The cluster data | 

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

