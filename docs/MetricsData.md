# MetricsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]MetricsLogs**](MetricsLogs.md) |  | [optional] 
**TotalExportSize** | Pointer to **float32** |  | [optional] 
**TotalReplaySize** | Pointer to **float32** |  | [optional] 

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

### GetTotalExportSize

`func (o *MetricsData) GetTotalExportSize() float32`

GetTotalExportSize returns the TotalExportSize field if non-nil, zero value otherwise.

### GetTotalExportSizeOk

`func (o *MetricsData) GetTotalExportSizeOk() (*float32, bool)`

GetTotalExportSizeOk returns a tuple with the TotalExportSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExportSize

`func (o *MetricsData) SetTotalExportSize(v float32)`

SetTotalExportSize sets TotalExportSize field to given value.

### HasTotalExportSize

`func (o *MetricsData) HasTotalExportSize() bool`

HasTotalExportSize returns a boolean if a field has been set.

### GetTotalReplaySize

`func (o *MetricsData) GetTotalReplaySize() float32`

GetTotalReplaySize returns the TotalReplaySize field if non-nil, zero value otherwise.

### GetTotalReplaySizeOk

`func (o *MetricsData) GetTotalReplaySizeOk() (*float32, bool)`

GetTotalReplaySizeOk returns a tuple with the TotalReplaySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReplaySize

`func (o *MetricsData) SetTotalReplaySize(v float32)`

SetTotalReplaySize sets TotalReplaySize field to given value.

### HasTotalReplaySize

`func (o *MetricsData) HasTotalReplaySize() bool`

HasTotalReplaySize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


