# MetricsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]MetricsLogs**](MetricsLogs.md) |  | [optional] 
**TotalExports** | Pointer to **int32** |  | [optional] 
**TotalReplays** | Pointer to **int32** |  | [optional] 

## Methods

### NewMetricsData

`func NewMetricsData() *MetricsData`

NewMetricsData instantiates a new MetricsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsDataWithDefaults

`func NewMetricsDataWithDefaults() *MetricsData`

NewMetricsDataWithDefaults instantiates a new MetricsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MetricsData) GetData() []MetricsLogs`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MetricsData) GetDataOk() (*[]MetricsLogs, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MetricsData) SetData(v []MetricsLogs)`

SetData sets Data field to given value.

### HasData

`func (o *MetricsData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotalExports

`func (o *MetricsData) GetTotalExports() int32`

GetTotalExports returns the TotalExports field if non-nil, zero value otherwise.

### GetTotalExportsOk

`func (o *MetricsData) GetTotalExportsOk() (*int32, bool)`

GetTotalExportsOk returns a tuple with the TotalExports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExports

`func (o *MetricsData) SetTotalExports(v int32)`

SetTotalExports sets TotalExports field to given value.

### HasTotalExports

`func (o *MetricsData) HasTotalExports() bool`

HasTotalExports returns a boolean if a field has been set.

### GetTotalReplays

`func (o *MetricsData) GetTotalReplays() int32`

GetTotalReplays returns the TotalReplays field if non-nil, zero value otherwise.

### GetTotalReplaysOk

`func (o *MetricsData) GetTotalReplaysOk() (*int32, bool)`

GetTotalReplaysOk returns a tuple with the TotalReplays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReplays

`func (o *MetricsData) SetTotalReplays(v int32)`

SetTotalReplays sets TotalReplays field to given value.

### HasTotalReplays

`func (o *MetricsData) HasTotalReplays() bool`

HasTotalReplays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


