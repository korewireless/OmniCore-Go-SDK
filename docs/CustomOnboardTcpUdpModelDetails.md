# CustomOnboardTcpUdpModelDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the TCP/UDP model | [optional] [readonly] 
**ModelName** | Pointer to **string** | The name of the model | [optional] 
**Manufacturer** | Pointer to **string** | The manufacturer of the model | [optional] 
**Image** | [**TcpUdpImage**](TcpUdpImage.md) |  | 
**TcpDetails** | [**TcpUdpPortDetail**](TcpUdpPortDetail.md) |  | 
**UdpDetails** | [**TcpUdpPortDetail**](TcpUdpPortDetail.md) |  | 
**Metadata** | Pointer to **map[string]interface{}** | Additional metadata in raw JSON format | [optional] 
**CreatedAt** | Pointer to **time.Time** | The creation timestamp of the model | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The last update timestamp of the model | [optional] [readonly] 

## Methods

### NewCustomOnboardTcpUdpModelDetails

`func NewCustomOnboardTcpUdpModelDetails(image TcpUdpImage, tcpDetails TcpUdpPortDetail, udpDetails TcpUdpPortDetail, ) *CustomOnboardTcpUdpModelDetails`

NewCustomOnboardTcpUdpModelDetails instantiates a new CustomOnboardTcpUdpModelDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomOnboardTcpUdpModelDetailsWithDefaults

`func NewCustomOnboardTcpUdpModelDetailsWithDefaults() *CustomOnboardTcpUdpModelDetails`

NewCustomOnboardTcpUdpModelDetailsWithDefaults instantiates a new CustomOnboardTcpUdpModelDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomOnboardTcpUdpModelDetails) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomOnboardTcpUdpModelDetails) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomOnboardTcpUdpModelDetails) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CustomOnboardTcpUdpModelDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModelName

`func (o *CustomOnboardTcpUdpModelDetails) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *CustomOnboardTcpUdpModelDetails) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *CustomOnboardTcpUdpModelDetails) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *CustomOnboardTcpUdpModelDetails) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### GetManufacturer

`func (o *CustomOnboardTcpUdpModelDetails) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *CustomOnboardTcpUdpModelDetails) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *CustomOnboardTcpUdpModelDetails) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *CustomOnboardTcpUdpModelDetails) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetImage

`func (o *CustomOnboardTcpUdpModelDetails) GetImage() TcpUdpImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CustomOnboardTcpUdpModelDetails) GetImageOk() (*TcpUdpImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CustomOnboardTcpUdpModelDetails) SetImage(v TcpUdpImage)`

SetImage sets Image field to given value.


### GetTcpDetails

`func (o *CustomOnboardTcpUdpModelDetails) GetTcpDetails() TcpUdpPortDetail`

GetTcpDetails returns the TcpDetails field if non-nil, zero value otherwise.

### GetTcpDetailsOk

`func (o *CustomOnboardTcpUdpModelDetails) GetTcpDetailsOk() (*TcpUdpPortDetail, bool)`

GetTcpDetailsOk returns a tuple with the TcpDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpDetails

`func (o *CustomOnboardTcpUdpModelDetails) SetTcpDetails(v TcpUdpPortDetail)`

SetTcpDetails sets TcpDetails field to given value.


### GetUdpDetails

`func (o *CustomOnboardTcpUdpModelDetails) GetUdpDetails() TcpUdpPortDetail`

GetUdpDetails returns the UdpDetails field if non-nil, zero value otherwise.

### GetUdpDetailsOk

`func (o *CustomOnboardTcpUdpModelDetails) GetUdpDetailsOk() (*TcpUdpPortDetail, bool)`

GetUdpDetailsOk returns a tuple with the UdpDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpDetails

`func (o *CustomOnboardTcpUdpModelDetails) SetUdpDetails(v TcpUdpPortDetail)`

SetUdpDetails sets UdpDetails field to given value.


### GetMetadata

`func (o *CustomOnboardTcpUdpModelDetails) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomOnboardTcpUdpModelDetails) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomOnboardTcpUdpModelDetails) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomOnboardTcpUdpModelDetails) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CustomOnboardTcpUdpModelDetails) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomOnboardTcpUdpModelDetails) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomOnboardTcpUdpModelDetails) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CustomOnboardTcpUdpModelDetails) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CustomOnboardTcpUdpModelDetails) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CustomOnboardTcpUdpModelDetails) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CustomOnboardTcpUdpModelDetails) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CustomOnboardTcpUdpModelDetails) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


