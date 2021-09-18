# Policies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthPolicies** | [**AuthPolicies**](AuthPolicies.md) |  | [optional] 
**ReplicationClusters** | **[]string** |  | [optional] 
**Bundles** | [**BundlesData**](BundlesData.md) |  | [optional] 
**BacklogQuotaMap** | [**map[string]BacklogQuota**](BacklogQuota.md) |  | [optional] 
**TopicDispatchRate** | [**map[string]DispatchRate**](DispatchRate.md) |  | [optional] 
**SubscriptionDispatchRate** | [**map[string]DispatchRate**](DispatchRate.md) |  | [optional] 
**ReplicatorDispatchRate** | [**map[string]DispatchRate**](DispatchRate.md) |  | [optional] 
**ClusterSubscribeRate** | [**map[string]SubscribeRate**](SubscribeRate.md) |  | [optional] 
**Persistence** | [**PersistencePolicies**](PersistencePolicies.md) |  | [optional] 
**DeduplicationEnabled** | **bool** |  | [optional] 
**LatencyStatsSampleRate** | **map[string]int32** |  | [optional] 
**MessageTtlInSeconds** | **int32** |  | [optional] 
**RetentionPolicies** | [**RetentionPolicies**](RetentionPolicies.md) |  | [optional] 
**Deleted** | **bool** |  | [optional] 
**AntiAffinityGroup** | **string** |  | [optional] 
**EncryptionRequired** | **bool** |  | [optional] 
**SubscriptionAuthMode** | **string** |  | [optional] 
**MaxProducersPerTopic** | **int32** |  | [optional] 
**MaxConsumersPerTopic** | **int32** |  | [optional] 
**MaxConsumersPerSubscription** | **int32** |  | [optional] 
**CompactionThreshold** | **int64** |  | [optional] 
**OffloadThreshold** | **int64** |  | [optional] 
**OffloadDeletionLagMs** | **int64** |  | [optional] 
**SchemaAutoUpdateCompatibilityStrategy** | **string** |  | [optional] 
**SchemaValidationEnforced** | **bool** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


