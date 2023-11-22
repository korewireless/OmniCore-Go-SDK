# MetricsResponseDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | Pointer to **string** |  | [optional] 
**NoOfFiles** | Pointer to **int32** |  | [optional] 
**FileSize** | Pointer to **float32** |  | [optional] 
**Noofoperations** | Pointer to **int32** |  | [optional] 
**NoOfReplays** | Pointer to **int32** |  | [optional] 
**NoOfExports** | Pointer to **int32** |  | [optional] 
**Operations** | Pointer to [**[]OperationMetrics**](OperationMetrics.md) |  | [optional] 
**DetailsForTimePeriod** | Pointer to [**MetricsData**](MetricsData.md) |  | [optional] 

## Methods

### NewMetricsResponseDetails

`func NewMetricsResponseDetails() *MetricsResponseDetails`

NewMetricsResponseDetails instantiates a new MetricsResponseDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsResponseDetailsWithDefaults

`func NewMetricsResponseDetailsWithDefaults() *MetricsResponseDetails`

NewMetricsResponseDetailsWithDefaults instantiates a new MetricsResponseDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *MetricsResponseDetails) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *MetricsResponseDetails) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *MetricsResponseDetails) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *MetricsResponseDetails) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetNoOfFiles

`func (o *MetricsResponseDetails) GetNoOfFiles() int32`

GetNoOfFiles returns the NoOfFiles field if non-nil, zero value otherwise.

### GetNoOfFilesOk

`func (o *MetricsResponseDetails) GetNoOfFilesOk() (*int32, bool)`

GetNoOfFilesOk returns a tuple with the NoOfFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfFiles

`func (o *MetricsResponseDetails) SetNoOfFiles(v int32)`

SetNoOfFiles sets NoOfFiles field to given value.

### HasNoOfFiles

`func (o *MetricsResponseDetails) HasNoOfFiles() bool`

HasNoOfFiles returns a boolean if a field has been set.

### GetFileSize

`func (o *MetricsResponseDetails) GetFileSize() float32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *MetricsResponseDetails) GetFileSizeOk() (*float32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *MetricsResponseDetails) SetFileSize(v float32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *MetricsResponseDetails) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetNoofoperations

`func (o *MetricsResponseDetails) GetNoofoperations() int32`

GetNoofoperations returns the Noofoperations field if non-nil, zero value otherwise.

### GetNoofoperationsOk

`func (o *MetricsResponseDetails) GetNoofoperationsOk() (*int32, bool)`

GetNoofoperationsOk returns a tuple with the Noofoperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoofoperations

`func (o *MetricsResponseDetails) SetNoofoperations(v int32)`

SetNoofoperations sets Noofoperations field to given value.

### HasNoofoperations

`func (o *MetricsResponseDetails) HasNoofoperations() bool`

HasNoofoperations returns a boolean if a field has been set.

### GetNoOfReplays

`func (o *MetricsResponseDetails) GetNoOfReplays() int32`

GetNoOfReplays returns the NoOfReplays field if non-nil, zero value otherwise.

### GetNoOfReplaysOk

`func (o *MetricsResponseDetails) GetNoOfReplaysOk() (*int32, bool)`

GetNoOfReplaysOk returns a tuple with the NoOfReplays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfReplays

`func (o *MetricsResponseDetails) SetNoOfReplays(v int32)`

SetNoOfReplays sets NoOfReplays field to given value.

### HasNoOfReplays

`func (o *MetricsResponseDetails) HasNoOfReplays() bool`

HasNoOfReplays returns a boolean if a field has been set.

### GetNoOfExports

`func (o *MetricsResponseDetails) GetNoOfExports() int32`

GetNoOfExports returns the NoOfExports field if non-nil, zero value otherwise.

### GetNoOfExportsOk

`func (o *MetricsResponseDetails) GetNoOfExportsOk() (*int32, bool)`

GetNoOfExportsOk returns a tuple with the NoOfExports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfExports

`func (o *MetricsResponseDetails) SetNoOfExports(v int32)`

SetNoOfExports sets NoOfExports field to given value.

### HasNoOfExports

`func (o *MetricsResponseDetails) HasNoOfExports() bool`

HasNoOfExports returns a boolean if a field has been set.

### GetOperations

`func (o *MetricsResponseDetails) GetOperations() []OperationMetrics`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *MetricsResponseDetails) GetOperationsOk() (*[]OperationMetrics, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *MetricsResponseDetails) SetOperations(v []OperationMetrics)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *MetricsResponseDetails) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetDetailsForTimePeriod

`func (o *MetricsResponseDetails) GetDetailsForTimePeriod() MetricsData`

GetDetailsForTimePeriod returns the DetailsForTimePeriod field if non-nil, zero value otherwise.

### GetDetailsForTimePeriodOk

`func (o *MetricsResponseDetails) GetDetailsForTimePeriodOk() (*MetricsData, bool)`

GetDetailsForTimePeriodOk returns a tuple with the DetailsForTimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailsForTimePeriod

`func (o *MetricsResponseDetails) SetDetailsForTimePeriod(v MetricsData)`

SetDetailsForTimePeriod sets DetailsForTimePeriod field to given value.

### HasDetailsForTimePeriod

`func (o *MetricsResponseDetails) HasDetailsForTimePeriod() bool`

HasDetailsForTimePeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


