# OmnicoreNewDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocked** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to [**OmnicoreDeviceConfig**](OmnicoreDeviceConfig.md) |  | [optional] 
**Credentials** | Pointer to [**[]OmnicoreDeviceCredential**](OmnicoreDeviceCredential.md) |  | [optional] 
**GatewayConfig** | Pointer to [**OmnicoreGatewayConfig**](OmnicoreGatewayConfig.md) |  | [optional] 
**Id** | **string** |  | 
**LogLevel** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewOmnicoreNewDevice

`func NewOmnicoreNewDevice(id string, ) *OmnicoreNewDevice`

NewOmnicoreNewDevice instantiates a new OmnicoreNewDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOmnicoreNewDeviceWithDefaults

`func NewOmnicoreNewDeviceWithDefaults() *OmnicoreNewDevice`

NewOmnicoreNewDeviceWithDefaults instantiates a new OmnicoreNewDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocked

`func (o *OmnicoreNewDevice) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *OmnicoreNewDevice) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *OmnicoreNewDevice) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *OmnicoreNewDevice) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetConfig

`func (o *OmnicoreNewDevice) GetConfig() OmnicoreDeviceConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *OmnicoreNewDevice) GetConfigOk() (*OmnicoreDeviceConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *OmnicoreNewDevice) SetConfig(v OmnicoreDeviceConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *OmnicoreNewDevice) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCredentials

`func (o *OmnicoreNewDevice) GetCredentials() []OmnicoreDeviceCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *OmnicoreNewDevice) GetCredentialsOk() (*[]OmnicoreDeviceCredential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *OmnicoreNewDevice) SetCredentials(v []OmnicoreDeviceCredential)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *OmnicoreNewDevice) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetGatewayConfig

`func (o *OmnicoreNewDevice) GetGatewayConfig() OmnicoreGatewayConfig`

GetGatewayConfig returns the GatewayConfig field if non-nil, zero value otherwise.

### GetGatewayConfigOk

`func (o *OmnicoreNewDevice) GetGatewayConfigOk() (*OmnicoreGatewayConfig, bool)`

GetGatewayConfigOk returns a tuple with the GatewayConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayConfig

`func (o *OmnicoreNewDevice) SetGatewayConfig(v OmnicoreGatewayConfig)`

SetGatewayConfig sets GatewayConfig field to given value.

### HasGatewayConfig

`func (o *OmnicoreNewDevice) HasGatewayConfig() bool`

HasGatewayConfig returns a boolean if a field has been set.

### GetId

`func (o *OmnicoreNewDevice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OmnicoreNewDevice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OmnicoreNewDevice) SetId(v string)`

SetId sets Id field to given value.


### GetLogLevel

`func (o *OmnicoreNewDevice) GetLogLevel() string`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *OmnicoreNewDevice) GetLogLevelOk() (*string, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *OmnicoreNewDevice) SetLogLevel(v string)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *OmnicoreNewDevice) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetMetadata

`func (o *OmnicoreNewDevice) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OmnicoreNewDevice) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OmnicoreNewDevice) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OmnicoreNewDevice) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


