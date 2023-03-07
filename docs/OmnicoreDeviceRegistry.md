# OmnicoreDeviceRegistry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedOn** | Pointer to **string** |  | [optional] 
**Credentials** | Pointer to [**[]OmnicoreRegistryCredential**](OmnicoreRegistryCredential.md) |  | [optional] 
**EventNotificationConfigs** | Pointer to [**[]OmnicoreEventNotificationConfig**](OmnicoreEventNotificationConfig.md) |  | [optional] 
**HttpConfig** | Pointer to [**OmnicoreHttpConfig**](OmnicoreHttpConfig.md) |  | [optional] 
**Id** | **string** |  | 
**LogLevel** | Pointer to **string** |  | [optional] 
**LogNotificationConfig** | Pointer to [**OmnicoreStateNotificationConfig**](OmnicoreStateNotificationConfig.md) |  | [optional] 
**MqttConfig** | Pointer to [**OmnicoreMqttConfig**](OmnicoreMqttConfig.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NumberOfDevices** | Pointer to **int32** |  | [optional] 
**NumberOfGateways** | Pointer to **int32** |  | [optional] 
**Parent** | **string** |  | 
**StateNotificationConfig** | Pointer to [**OmnicoreStateNotificationConfig**](OmnicoreStateNotificationConfig.md) |  | [optional] 
**UpdatedOn** | Pointer to **string** |  | [optional] 

## Methods

### NewOmnicoreDeviceRegistry

`func NewOmnicoreDeviceRegistry(id string, parent string, ) *OmnicoreDeviceRegistry`

NewOmnicoreDeviceRegistry instantiates a new OmnicoreDeviceRegistry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOmnicoreDeviceRegistryWithDefaults

`func NewOmnicoreDeviceRegistryWithDefaults() *OmnicoreDeviceRegistry`

NewOmnicoreDeviceRegistryWithDefaults instantiates a new OmnicoreDeviceRegistry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedOn

`func (o *OmnicoreDeviceRegistry) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *OmnicoreDeviceRegistry) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *OmnicoreDeviceRegistry) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *OmnicoreDeviceRegistry) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetCredentials

`func (o *OmnicoreDeviceRegistry) GetCredentials() []OmnicoreRegistryCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *OmnicoreDeviceRegistry) GetCredentialsOk() (*[]OmnicoreRegistryCredential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *OmnicoreDeviceRegistry) SetCredentials(v []OmnicoreRegistryCredential)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *OmnicoreDeviceRegistry) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetEventNotificationConfigs

`func (o *OmnicoreDeviceRegistry) GetEventNotificationConfigs() []OmnicoreEventNotificationConfig`

GetEventNotificationConfigs returns the EventNotificationConfigs field if non-nil, zero value otherwise.

### GetEventNotificationConfigsOk

`func (o *OmnicoreDeviceRegistry) GetEventNotificationConfigsOk() (*[]OmnicoreEventNotificationConfig, bool)`

GetEventNotificationConfigsOk returns a tuple with the EventNotificationConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotificationConfigs

`func (o *OmnicoreDeviceRegistry) SetEventNotificationConfigs(v []OmnicoreEventNotificationConfig)`

SetEventNotificationConfigs sets EventNotificationConfigs field to given value.

### HasEventNotificationConfigs

`func (o *OmnicoreDeviceRegistry) HasEventNotificationConfigs() bool`

HasEventNotificationConfigs returns a boolean if a field has been set.

### GetHttpConfig

`func (o *OmnicoreDeviceRegistry) GetHttpConfig() OmnicoreHttpConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *OmnicoreDeviceRegistry) GetHttpConfigOk() (*OmnicoreHttpConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *OmnicoreDeviceRegistry) SetHttpConfig(v OmnicoreHttpConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *OmnicoreDeviceRegistry) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetId

`func (o *OmnicoreDeviceRegistry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OmnicoreDeviceRegistry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OmnicoreDeviceRegistry) SetId(v string)`

SetId sets Id field to given value.


### GetLogLevel

`func (o *OmnicoreDeviceRegistry) GetLogLevel() string`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *OmnicoreDeviceRegistry) GetLogLevelOk() (*string, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *OmnicoreDeviceRegistry) SetLogLevel(v string)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *OmnicoreDeviceRegistry) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetLogNotificationConfig

`func (o *OmnicoreDeviceRegistry) GetLogNotificationConfig() OmnicoreStateNotificationConfig`

GetLogNotificationConfig returns the LogNotificationConfig field if non-nil, zero value otherwise.

### GetLogNotificationConfigOk

`func (o *OmnicoreDeviceRegistry) GetLogNotificationConfigOk() (*OmnicoreStateNotificationConfig, bool)`

GetLogNotificationConfigOk returns a tuple with the LogNotificationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogNotificationConfig

`func (o *OmnicoreDeviceRegistry) SetLogNotificationConfig(v OmnicoreStateNotificationConfig)`

SetLogNotificationConfig sets LogNotificationConfig field to given value.

### HasLogNotificationConfig

`func (o *OmnicoreDeviceRegistry) HasLogNotificationConfig() bool`

HasLogNotificationConfig returns a boolean if a field has been set.

### GetMqttConfig

`func (o *OmnicoreDeviceRegistry) GetMqttConfig() OmnicoreMqttConfig`

GetMqttConfig returns the MqttConfig field if non-nil, zero value otherwise.

### GetMqttConfigOk

`func (o *OmnicoreDeviceRegistry) GetMqttConfigOk() (*OmnicoreMqttConfig, bool)`

GetMqttConfigOk returns a tuple with the MqttConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttConfig

`func (o *OmnicoreDeviceRegistry) SetMqttConfig(v OmnicoreMqttConfig)`

SetMqttConfig sets MqttConfig field to given value.

### HasMqttConfig

`func (o *OmnicoreDeviceRegistry) HasMqttConfig() bool`

HasMqttConfig returns a boolean if a field has been set.

### GetName

`func (o *OmnicoreDeviceRegistry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OmnicoreDeviceRegistry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OmnicoreDeviceRegistry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OmnicoreDeviceRegistry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumberOfDevices

`func (o *OmnicoreDeviceRegistry) GetNumberOfDevices() int32`

GetNumberOfDevices returns the NumberOfDevices field if non-nil, zero value otherwise.

### GetNumberOfDevicesOk

`func (o *OmnicoreDeviceRegistry) GetNumberOfDevicesOk() (*int32, bool)`

GetNumberOfDevicesOk returns a tuple with the NumberOfDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDevices

`func (o *OmnicoreDeviceRegistry) SetNumberOfDevices(v int32)`

SetNumberOfDevices sets NumberOfDevices field to given value.

### HasNumberOfDevices

`func (o *OmnicoreDeviceRegistry) HasNumberOfDevices() bool`

HasNumberOfDevices returns a boolean if a field has been set.

### GetNumberOfGateways

`func (o *OmnicoreDeviceRegistry) GetNumberOfGateways() int32`

GetNumberOfGateways returns the NumberOfGateways field if non-nil, zero value otherwise.

### GetNumberOfGatewaysOk

`func (o *OmnicoreDeviceRegistry) GetNumberOfGatewaysOk() (*int32, bool)`

GetNumberOfGatewaysOk returns a tuple with the NumberOfGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfGateways

`func (o *OmnicoreDeviceRegistry) SetNumberOfGateways(v int32)`

SetNumberOfGateways sets NumberOfGateways field to given value.

### HasNumberOfGateways

`func (o *OmnicoreDeviceRegistry) HasNumberOfGateways() bool`

HasNumberOfGateways returns a boolean if a field has been set.

### GetParent

`func (o *OmnicoreDeviceRegistry) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *OmnicoreDeviceRegistry) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *OmnicoreDeviceRegistry) SetParent(v string)`

SetParent sets Parent field to given value.


### GetStateNotificationConfig

`func (o *OmnicoreDeviceRegistry) GetStateNotificationConfig() OmnicoreStateNotificationConfig`

GetStateNotificationConfig returns the StateNotificationConfig field if non-nil, zero value otherwise.

### GetStateNotificationConfigOk

`func (o *OmnicoreDeviceRegistry) GetStateNotificationConfigOk() (*OmnicoreStateNotificationConfig, bool)`

GetStateNotificationConfigOk returns a tuple with the StateNotificationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateNotificationConfig

`func (o *OmnicoreDeviceRegistry) SetStateNotificationConfig(v OmnicoreStateNotificationConfig)`

SetStateNotificationConfig sets StateNotificationConfig field to given value.

### HasStateNotificationConfig

`func (o *OmnicoreDeviceRegistry) HasStateNotificationConfig() bool`

HasStateNotificationConfig returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *OmnicoreDeviceRegistry) GetUpdatedOn() string`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *OmnicoreDeviceRegistry) GetUpdatedOnOk() (*string, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *OmnicoreDeviceRegistry) SetUpdatedOn(v string)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *OmnicoreDeviceRegistry) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


