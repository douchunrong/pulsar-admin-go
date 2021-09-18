# LoadReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [optional] 
**BrokerVersionString** | **string** |  | [optional] 
**WebServiceUrl** | **string** |  | [optional] 
**WebServiceUrlTls** | **string** |  | [optional] 
**PulsarServiceUrl** | **string** |  | [optional] 
**PulsarServiceUrlTls** | **string** |  | [optional] 
**PersistentTopicsEnabled** | **bool** |  | [optional] 
**NonPersistentTopicsEnabled** | **bool** |  | [optional] 
**Timestamp** | **int64** |  | [optional] 
**MsgRateIn** | **float64** |  | [optional] 
**MsgRateOut** | **float64** |  | [optional] 
**NumTopics** | **int32** |  | [optional] 
**NumConsumers** | **int32** |  | [optional] 
**NumProducers** | **int32** |  | [optional] 
**NumBundles** | **int32** |  | [optional] 
**SystemResourceUsage** | [**SystemResourceUsage**](SystemResourceUsage.md) |  | [optional] 
**BundleStats** | [**map[string]NamespaceBundleStats**](NamespaceBundleStats.md) |  | [optional] 
**BundleGains** | **[]string** |  | [optional] 
**BundleLosses** | **[]string** |  | [optional] 
**AllocatedCPU** | **float64** |  | [optional] 
**AllocatedMemory** | **float64** |  | [optional] 
**AllocatedBandwidthIn** | **float64** |  | [optional] 
**AllocatedBandwidthOut** | **float64** |  | [optional] 
**AllocatedMsgRateIn** | **float64** |  | [optional] 
**AllocatedMsgRateOut** | **float64** |  | [optional] 
**PreAllocatedCPU** | **float64** |  | [optional] 
**PreAllocatedMemory** | **float64** |  | [optional] 
**PreAllocatedBandwidthIn** | **float64** |  | [optional] 
**PreAllocatedBandwidthOut** | **float64** |  | [optional] 
**PreAllocatedMsgRateIn** | **float64** |  | [optional] 
**PreAllocatedMsgRateOut** | **float64** |  | [optional] 
**OverLoaded** | **bool** |  | [optional] 
**LoadReportType** | **string** |  | [optional] 
**UnderLoaded** | **bool** |  | [optional] 
**BandwidthIn** | [**ResourceUsage**](ResourceUsage.md) |  | [optional] 
**BandwidthOut** | [**ResourceUsage**](ResourceUsage.md) |  | [optional] 
**Memory** | [**ResourceUsage**](ResourceUsage.md) |  | [optional] 
**LastUpdate** | **int64** |  | [optional] 
**MsgThroughputIn** | **float64** |  | [optional] 
**MsgThroughputOut** | **float64** |  | [optional] 
**Cpu** | [**ResourceUsage**](ResourceUsage.md) |  | [optional] 
**DirectMemory** | [**ResourceUsage**](ResourceUsage.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


