# OmnicoreUpdateDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocked** | Pointer to **bool** |  | [optional] 
**Credentials** | Pointer to [**[]OmnicoreDeviceCredential**](OmnicoreDeviceCredential.md) |  | [optional] 
**LogLevel** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewOmnicoreUpdateDevice

`func NewOmnicoreUpdateDevice() *OmnicoreUpdateDevice`

NewOmnicoreUpdateDevice instantiates a new OmnicoreUpdateDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOmnicoreUpdateDeviceWithDefaults

`func NewOmnicoreUpdateDeviceWithDefaults() *OmnicoreUpdateDevice`

NewOmnicoreUpdateDeviceWithDefaults instantiates a new OmnicoreUpdateDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocked

`func (o *OmnicoreUpdateDevice) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *OmnicoreUpdateDevice) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *OmnicoreUpdateDevice) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *OmnicoreUpdateDevice) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetCredentials

`func (o *OmnicoreUpdateDevice) GetCredentials() []OmnicoreDeviceCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *OmnicoreUpdateDevice) GetCredentialsOk() (*[]OmnicoreDeviceCredential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *OmnicoreUpdateDevice) SetCredentials(v []OmnicoreDeviceCredential)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *OmnicoreUpdateDevice) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetLogLevel

`func (o *OmnicoreUpdateDevice) GetLogLevel() string`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *OmnicoreUpdateDevice) GetLogLevelOk() (*string, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *OmnicoreUpdateDevice) SetLogLevel(v string)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *OmnicoreUpdateDevice) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetMetadata

`func (o *OmnicoreUpdateDevice) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OmnicoreUpdateDevice) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OmnicoreUpdateDevice) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OmnicoreUpdateDevice) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


