# OmnicoreMqttConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MqttEnabledState** | Pointer to **string** | MqttEnabledState: If enabled, allows connections using the MQTT protocol. Otherwise, MQTT connections to this registry will fail.  Possible values:   \&quot;MQTT_STATE_UNSPECIFIED\&quot; - No MQTT state specified. If not specified, MQTT will be enabled by default.   \&quot;MQTT_ENABLED\&quot; - Enables a MQTT connection.   \&quot;MQTT_DISABLED\&quot; - Disables a MQTT connection. | [optional] 

## Methods

### NewOmnicoreMqttConfig

`func NewOmnicoreMqttConfig() *OmnicoreMqttConfig`

NewOmnicoreMqttConfig instantiates a new OmnicoreMqttConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOmnicoreMqttConfigWithDefaults

`func NewOmnicoreMqttConfigWithDefaults() *OmnicoreMqttConfig`

NewOmnicoreMqttConfigWithDefaults instantiates a new OmnicoreMqttConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMqttEnabledState

`func (o *OmnicoreMqttConfig) GetMqttEnabledState() string`

GetMqttEnabledState returns the MqttEnabledState field if non-nil, zero value otherwise.

### GetMqttEnabledStateOk

`func (o *OmnicoreMqttConfig) GetMqttEnabledStateOk() (*string, bool)`

GetMqttEnabledStateOk returns a tuple with the MqttEnabledState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttEnabledState

`func (o *OmnicoreMqttConfig) SetMqttEnabledState(v string)`

SetMqttEnabledState sets MqttEnabledState field to given value.

### HasMqttEnabledState

`func (o *OmnicoreMqttConfig) HasMqttEnabledState() bool`

HasMqttEnabledState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


