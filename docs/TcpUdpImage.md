# TcpUdpImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Valid** | **bool** | Indicates if the image is valid | 
**ContentType** | Pointer to **string** | The content type of the image | [optional] 
**StorageKey** | Pointer to **string** | The storage key of the image | [optional] 
**Link** | Pointer to **string** | The link to the image | [optional] [readonly] 
**Data** | Pointer to ***os.File** | The binary data of the image | [optional] 

## Methods

### NewTcpUdpImage

`func NewTcpUdpImage(valid bool, ) *TcpUdpImage`

NewTcpUdpImage instantiates a new TcpUdpImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTcpUdpImageWithDefaults

`func NewTcpUdpImageWithDefaults() *TcpUdpImage`

NewTcpUdpImageWithDefaults instantiates a new TcpUdpImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValid

`func (o *TcpUdpImage) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *TcpUdpImage) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *TcpUdpImage) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetContentType

`func (o *TcpUdpImage) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *TcpUdpImage) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *TcpUdpImage) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *TcpUdpImage) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetStorageKey

`func (o *TcpUdpImage) GetStorageKey() string`

GetStorageKey returns the StorageKey field if non-nil, zero value otherwise.

### GetStorageKeyOk

`func (o *TcpUdpImage) GetStorageKeyOk() (*string, bool)`

GetStorageKeyOk returns a tuple with the StorageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageKey

`func (o *TcpUdpImage) SetStorageKey(v string)`

SetStorageKey sets StorageKey field to given value.

### HasStorageKey

`func (o *TcpUdpImage) HasStorageKey() bool`

HasStorageKey returns a boolean if a field has been set.

### GetLink

`func (o *TcpUdpImage) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *TcpUdpImage) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *TcpUdpImage) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *TcpUdpImage) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetData

`func (o *TcpUdpImage) GetData() *os.File`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TcpUdpImage) GetDataOk() (**os.File, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TcpUdpImage) SetData(v *os.File)`

SetData sets Data field to given value.

### HasData

`func (o *TcpUdpImage) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


