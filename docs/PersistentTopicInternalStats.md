# PersistentTopicInternalStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntriesAddedCounter** | **int64** |  | [optional] 
**NumberOfEntries** | **int64** |  | [optional] 
**TotalSize** | **int64** |  | [optional] 
**CurrentLedgerEntries** | **int64** |  | [optional] 
**CurrentLedgerSize** | **int64** |  | [optional] 
**LastLedgerCreatedTimestamp** | **string** |  | [optional] 
**LastLedgerCreationFailureTimestamp** | **string** |  | [optional] 
**WaitingCursorsCount** | **int32** |  | [optional] 
**PendingAddEntriesCount** | **int32** |  | [optional] 
**LastConfirmedEntry** | **string** |  | [optional] 
**State** | **string** |  | [optional] 
**Ledgers** | [**[]LedgerInfo**](LedgerInfo.md) |  | [optional] 
**Cursors** | [**map[string]CursorStats**](CursorStats.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


