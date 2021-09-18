# NonPersistentTopicStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MsgRateIn** | **float64** |  | [optional] 
**MsgThroughputIn** | **float64** |  | [optional] 
**MsgRateOut** | **float64** |  | [optional] 
**MsgThroughputOut** | **float64** |  | [optional] 
**AverageMsgSize** | **float64** |  | [optional] 
**StorageSize** | **int64** |  | [optional] 
**Publishers** | [**[]NonPersistentPublisherStats**](NonPersistentPublisherStats.md) |  | [optional] 
**Subscriptions** | [**map[string]NonPersistentSubscriptionStats**](NonPersistentSubscriptionStats.md) |  | [optional] 
**Replication** | [**map[string]NonPersistentReplicatorStats**](NonPersistentReplicatorStats.md) |  | [optional] 
**DeduplicationStatus** | **string** |  | [optional] 
**MsgDropRate** | **float64** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


