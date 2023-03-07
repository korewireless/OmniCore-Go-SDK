# OmnicoreListDeviceRegistriesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceRegistries** | [**[]OmnicoreDeviceRegistry**](OmnicoreDeviceRegistry.md) |  | 
**PageNumber** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewOmnicoreListDeviceRegistriesResponse

`func NewOmnicoreListDeviceRegistriesResponse(deviceRegistries []OmnicoreDeviceRegistry, ) *OmnicoreListDeviceRegistriesResponse`

NewOmnicoreListDeviceRegistriesResponse instantiates a new OmnicoreListDeviceRegistriesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOmnicoreListDeviceRegistriesResponseWithDefaults

`func NewOmnicoreListDeviceRegistriesResponseWithDefaults() *OmnicoreListDeviceRegistriesResponse`

NewOmnicoreListDeviceRegistriesResponseWithDefaults instantiates a new OmnicoreListDeviceRegistriesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceRegistries

`func (o *OmnicoreListDeviceRegistriesResponse) GetDeviceRegistries() []OmnicoreDeviceRegistry`

GetDeviceRegistries returns the DeviceRegistries field if non-nil, zero value otherwise.

### GetDeviceRegistriesOk

`func (o *OmnicoreListDeviceRegistriesResponse) GetDeviceRegistriesOk() (*[]OmnicoreDeviceRegistry, bool)`

GetDeviceRegistriesOk returns a tuple with the DeviceRegistries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistries

`func (o *OmnicoreListDeviceRegistriesResponse) SetDeviceRegistries(v []OmnicoreDeviceRegistry)`

SetDeviceRegistries sets DeviceRegistries field to given value.


### GetPageNumber

`func (o *OmnicoreListDeviceRegistriesResponse) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *OmnicoreListDeviceRegistriesResponse) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *OmnicoreListDeviceRegistriesResponse) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *OmnicoreListDeviceRegistriesResponse) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *OmnicoreListDeviceRegistriesResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *OmnicoreListDeviceRegistriesResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *OmnicoreListDeviceRegistriesResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *OmnicoreListDeviceRegistriesResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalCount

`func (o *OmnicoreListDeviceRegistriesResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *OmnicoreListDeviceRegistriesResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *OmnicoreListDeviceRegistriesResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *OmnicoreListDeviceRegistriesResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


