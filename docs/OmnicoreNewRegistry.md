# OmnicoreNewRegistry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**[]OmnicoreRegistryCredential**](OmnicoreRegistryCredential.md) |  | [optional] 
**EventNotificationConfigs** | Pointer to [**[]OmnicoreEventNotificationConfig**](OmnicoreEventNotificationConfig.md) |  | [optional] 
**HttpConfig** | Pointer to [**OmnicoreHttpConfig**](OmnicoreHttpConfig.md) |  | [optional] 
**Id** | **string** |  | 
**LogLevel** | Pointer to **string** |  | [optional] 
**LogNotificationConfig** | Pointer to [**OmnicoreStateNotificationConfig**](OmnicoreStateNotificationConfig.md) |  | [optional] 
**MqttConfig** | Pointer to [**OmnicoreMqttConfig**](OmnicoreMqttConfig.md) |  | [optional] 
**StateNotificationConfig** | Pointer to [**OmnicoreStateNotificationConfig**](OmnicoreStateNotificationConfig.md) |  | [optional] 

## Methods

### NewOmnicoreNewRegistry

`func NewOmnicoreNewRegistry(id string, ) *OmnicoreNewRegistry`

NewOmnicoreNewRegistry instantiates a new OmnicoreNewRegistry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOmnicoreNewRegistryWithDefaults

`func NewOmnicoreNewRegistryWithDefaults() *OmnicoreNewRegistry`

NewOmnicoreNewRegistryWithDefaults instantiates a new OmnicoreNewRegistry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *OmnicoreNewRegistry) GetCredentials() []OmnicoreRegistryCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *OmnicoreNewRegistry) GetCredentialsOk() (*[]OmnicoreRegistryCredential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *OmnicoreNewRegistry) SetCredentials(v []OmnicoreRegistryCredential)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *OmnicoreNewRegistry) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetEventNotificationConfigs

`func (o *OmnicoreNewRegistry) GetEventNotificationConfigs() []OmnicoreEventNotificationConfig`

GetEventNotificationConfigs returns the EventNotificationConfigs field if non-nil, zero value otherwise.

### GetEventNotificationConfigsOk

`func (o *OmnicoreNewRegistry) GetEventNotificationConfigsOk() (*[]OmnicoreEventNotificationConfig, bool)`

GetEventNotificationConfigsOk returns a tuple with the EventNotificationConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotificationConfigs

`func (o *OmnicoreNewRegistry) SetEventNotificationConfigs(v []OmnicoreEventNotificationConfig)`

SetEventNotificationConfigs sets EventNotificationConfigs field to given value.

### HasEventNotificationConfigs

`func (o *OmnicoreNewRegistry) HasEventNotificationConfigs() bool`

HasEventNotificationConfigs returns a boolean if a field has been set.

### GetHttpConfig

`func (o *OmnicoreNewRegistry) GetHttpConfig() OmnicoreHttpConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *OmnicoreNewRegistry) GetHttpConfigOk() (*OmnicoreHttpConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *OmnicoreNewRegistry) SetHttpConfig(v OmnicoreHttpConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *OmnicoreNewRegistry) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetId

`func (o *OmnicoreNewRegistry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OmnicoreNewRegistry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OmnicoreNewRegistry) SetId(v string)`

SetId sets Id field to given value.


### GetLogLevel

`func (o *OmnicoreNewRegistry) GetLogLevel() string`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *OmnicoreNewRegistry) GetLogLevelOk() (*string, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *OmnicoreNewRegistry) SetLogLevel(v string)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *OmnicoreNewRegistry) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetLogNotificationConfig

`func (o *OmnicoreNewRegistry) GetLogNotificationConfig() OmnicoreStateNotificationConfig`

GetLogNotificationConfig returns the LogNotificationConfig field if non-nil, zero value otherwise.

### GetLogNotificationConfigOk

`func (o *OmnicoreNewRegistry) GetLogNotificationConfigOk() (*OmnicoreStateNotificationConfig, bool)`

GetLogNotificationConfigOk returns a tuple with the LogNotificationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogNotificationConfig

`func (o *OmnicoreNewRegistry) SetLogNotificationConfig(v OmnicoreStateNotificationConfig)`

SetLogNotificationConfig sets LogNotificationConfig field to given value.

### HasLogNotificationConfig

`func (o *OmnicoreNewRegistry) HasLogNotificationConfig() bool`

HasLogNotificationConfig returns a boolean if a field has been set.

### GetMqttConfig

`func (o *OmnicoreNewRegistry) GetMqttConfig() OmnicoreMqttConfig`

GetMqttConfig returns the MqttConfig field if non-nil, zero value otherwise.

### GetMqttConfigOk

`func (o *OmnicoreNewRegistry) GetMqttConfigOk() (*OmnicoreMqttConfig, bool)`

GetMqttConfigOk returns a tuple with the MqttConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttConfig

`func (o *OmnicoreNewRegistry) SetMqttConfig(v OmnicoreMqttConfig)`

SetMqttConfig sets MqttConfig field to given value.

### HasMqttConfig

`func (o *OmnicoreNewRegistry) HasMqttConfig() bool`

HasMqttConfig returns a boolean if a field has been set.

### GetStateNotificationConfig

`func (o *OmnicoreNewRegistry) GetStateNotificationConfig() OmnicoreStateNotificationConfig`

GetStateNotificationConfig returns the StateNotificationConfig field if non-nil, zero value otherwise.

### GetStateNotificationConfigOk

`func (o *OmnicoreNewRegistry) GetStateNotificationConfigOk() (*OmnicoreStateNotificationConfig, bool)`

GetStateNotificationConfigOk returns a tuple with the StateNotificationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateNotificationConfig

`func (o *OmnicoreNewRegistry) SetStateNotificationConfig(v OmnicoreStateNotificationConfig)`

SetStateNotificationConfig sets StateNotificationConfig field to given value.

### HasStateNotificationConfig

`func (o *OmnicoreNewRegistry) HasStateNotificationConfig() bool`

HasStateNotificationConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


