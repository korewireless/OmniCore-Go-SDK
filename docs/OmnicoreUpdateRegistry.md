# OmnicoreUpdateRegistry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**[]OmnicoreRegistryCredential**](OmnicoreRegistryCredential.md) |  | [optional] 
**EventNotificationConfigs** | Pointer to [**[]OmnicoreEventNotificationConfig**](OmnicoreEventNotificationConfig.md) |  | [optional] 
**HttpConfig** | Pointer to [**OmnicoreHttpConfig**](OmnicoreHttpConfig.md) |  | [optional] 
**LogLevel** | Pointer to **string** |  | [optional] 
**LogNotificationConfig** | Pointer to [**OmnicoreStateNotificationConfig**](OmnicoreStateNotificationConfig.md) |  | [optional] 
**MqttConfig** | Pointer to [**OmnicoreMqttConfig**](OmnicoreMqttConfig.md) |  | [optional] 
**StateNotificationConfig** | Pointer to [**OmnicoreStateNotificationConfig**](OmnicoreStateNotificationConfig.md) |  | [optional] 

## Methods

### NewOmnicoreUpdateRegistry

`func NewOmnicoreUpdateRegistry() *OmnicoreUpdateRegistry`

NewOmnicoreUpdateRegistry instantiates a new OmnicoreUpdateRegistry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOmnicoreUpdateRegistryWithDefaults

`func NewOmnicoreUpdateRegistryWithDefaults() *OmnicoreUpdateRegistry`

NewOmnicoreUpdateRegistryWithDefaults instantiates a new OmnicoreUpdateRegistry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *OmnicoreUpdateRegistry) GetCredentials() []OmnicoreRegistryCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *OmnicoreUpdateRegistry) GetCredentialsOk() (*[]OmnicoreRegistryCredential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *OmnicoreUpdateRegistry) SetCredentials(v []OmnicoreRegistryCredential)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *OmnicoreUpdateRegistry) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetEventNotificationConfigs

`func (o *OmnicoreUpdateRegistry) GetEventNotificationConfigs() []OmnicoreEventNotificationConfig`

GetEventNotificationConfigs returns the EventNotificationConfigs field if non-nil, zero value otherwise.

### GetEventNotificationConfigsOk

`func (o *OmnicoreUpdateRegistry) GetEventNotificationConfigsOk() (*[]OmnicoreEventNotificationConfig, bool)`

GetEventNotificationConfigsOk returns a tuple with the EventNotificationConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotificationConfigs

`func (o *OmnicoreUpdateRegistry) SetEventNotificationConfigs(v []OmnicoreEventNotificationConfig)`

SetEventNotificationConfigs sets EventNotificationConfigs field to given value.

### HasEventNotificationConfigs

`func (o *OmnicoreUpdateRegistry) HasEventNotificationConfigs() bool`

HasEventNotificationConfigs returns a boolean if a field has been set.

### GetHttpConfig

`func (o *OmnicoreUpdateRegistry) GetHttpConfig() OmnicoreHttpConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *OmnicoreUpdateRegistry) GetHttpConfigOk() (*OmnicoreHttpConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *OmnicoreUpdateRegistry) SetHttpConfig(v OmnicoreHttpConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *OmnicoreUpdateRegistry) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetLogLevel

`func (o *OmnicoreUpdateRegistry) GetLogLevel() string`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *OmnicoreUpdateRegistry) GetLogLevelOk() (*string, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *OmnicoreUpdateRegistry) SetLogLevel(v string)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *OmnicoreUpdateRegistry) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetLogNotificationConfig

`func (o *OmnicoreUpdateRegistry) GetLogNotificationConfig() OmnicoreStateNotificationConfig`

GetLogNotificationConfig returns the LogNotificationConfig field if non-nil, zero value otherwise.

### GetLogNotificationConfigOk

`func (o *OmnicoreUpdateRegistry) GetLogNotificationConfigOk() (*OmnicoreStateNotificationConfig, bool)`

GetLogNotificationConfigOk returns a tuple with the LogNotificationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogNotificationConfig

`func (o *OmnicoreUpdateRegistry) SetLogNotificationConfig(v OmnicoreStateNotificationConfig)`

SetLogNotificationConfig sets LogNotificationConfig field to given value.

### HasLogNotificationConfig

`func (o *OmnicoreUpdateRegistry) HasLogNotificationConfig() bool`

HasLogNotificationConfig returns a boolean if a field has been set.

### GetMqttConfig

`func (o *OmnicoreUpdateRegistry) GetMqttConfig() OmnicoreMqttConfig`

GetMqttConfig returns the MqttConfig field if non-nil, zero value otherwise.

### GetMqttConfigOk

`func (o *OmnicoreUpdateRegistry) GetMqttConfigOk() (*OmnicoreMqttConfig, bool)`

GetMqttConfigOk returns a tuple with the MqttConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttConfig

`func (o *OmnicoreUpdateRegistry) SetMqttConfig(v OmnicoreMqttConfig)`

SetMqttConfig sets MqttConfig field to given value.

### HasMqttConfig

`func (o *OmnicoreUpdateRegistry) HasMqttConfig() bool`

HasMqttConfig returns a boolean if a field has been set.

### GetStateNotificationConfig

`func (o *OmnicoreUpdateRegistry) GetStateNotificationConfig() OmnicoreStateNotificationConfig`

GetStateNotificationConfig returns the StateNotificationConfig field if non-nil, zero value otherwise.

### GetStateNotificationConfigOk

`func (o *OmnicoreUpdateRegistry) GetStateNotificationConfigOk() (*OmnicoreStateNotificationConfig, bool)`

GetStateNotificationConfigOk returns a tuple with the StateNotificationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateNotificationConfig

`func (o *OmnicoreUpdateRegistry) SetStateNotificationConfig(v OmnicoreStateNotificationConfig)`

SetStateNotificationConfig sets StateNotificationConfig field to given value.

### HasStateNotificationConfig

`func (o *OmnicoreUpdateRegistry) HasStateNotificationConfig() bool`

HasStateNotificationConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


