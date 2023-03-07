# ListDeviceRegistriesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceRegistries** | [**[]DeviceRegistry**](DeviceRegistry.md) |  | 
**PageNumber** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewListDeviceRegistriesResponse

`func NewListDeviceRegistriesResponse(deviceRegistries []DeviceRegistry, ) *ListDeviceRegistriesResponse`

NewListDeviceRegistriesResponse instantiates a new ListDeviceRegistriesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDeviceRegistriesResponseWithDefaults

`func NewListDeviceRegistriesResponseWithDefaults() *ListDeviceRegistriesResponse`

NewListDeviceRegistriesResponseWithDefaults instantiates a new ListDeviceRegistriesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceRegistries

`func (o *ListDeviceRegistriesResponse) GetDeviceRegistries() []DeviceRegistry`

GetDeviceRegistries returns the DeviceRegistries field if non-nil, zero value otherwise.

### GetDeviceRegistriesOk

`func (o *ListDeviceRegistriesResponse) GetDeviceRegistriesOk() (*[]DeviceRegistry, bool)`

GetDeviceRegistriesOk returns a tuple with the DeviceRegistries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistries

`func (o *ListDeviceRegistriesResponse) SetDeviceRegistries(v []DeviceRegistry)`

SetDeviceRegistries sets DeviceRegistries field to given value.


### GetPageNumber

`func (o *ListDeviceRegistriesResponse) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *ListDeviceRegistriesResponse) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *ListDeviceRegistriesResponse) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *ListDeviceRegistriesResponse) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *ListDeviceRegistriesResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListDeviceRegistriesResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListDeviceRegistriesResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListDeviceRegistriesResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalCount

`func (o *ListDeviceRegistriesResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListDeviceRegistriesResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListDeviceRegistriesResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ListDeviceRegistriesResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


