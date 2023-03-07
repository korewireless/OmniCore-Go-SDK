# OmnicoreListDevicesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | [**[]OmnicoreDevice**](OmnicoreDevice.md) |  | 
**PageNumber** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewOmnicoreListDevicesResponse

`func NewOmnicoreListDevicesResponse(devices []OmnicoreDevice, ) *OmnicoreListDevicesResponse`

NewOmnicoreListDevicesResponse instantiates a new OmnicoreListDevicesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOmnicoreListDevicesResponseWithDefaults

`func NewOmnicoreListDevicesResponseWithDefaults() *OmnicoreListDevicesResponse`

NewOmnicoreListDevicesResponseWithDefaults instantiates a new OmnicoreListDevicesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *OmnicoreListDevicesResponse) GetDevices() []OmnicoreDevice`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *OmnicoreListDevicesResponse) GetDevicesOk() (*[]OmnicoreDevice, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *OmnicoreListDevicesResponse) SetDevices(v []OmnicoreDevice)`

SetDevices sets Devices field to given value.


### GetPageNumber

`func (o *OmnicoreListDevicesResponse) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *OmnicoreListDevicesResponse) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *OmnicoreListDevicesResponse) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *OmnicoreListDevicesResponse) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *OmnicoreListDevicesResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *OmnicoreListDevicesResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *OmnicoreListDevicesResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *OmnicoreListDevicesResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalCount

`func (o *OmnicoreListDevicesResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *OmnicoreListDevicesResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *OmnicoreListDevicesResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *OmnicoreListDevicesResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


