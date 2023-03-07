/* Copyright 2018-2020 KoreWireless
 *
 * This is part of the KoreWireless Omnicore SDK.
 * It is licensed under the BSD 3-Clause license; you may not use this file
 * except in compliance with the License.
 *
 * You may obtain a copy of the License at:
 *  https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */


package OmniCore

import (
	"encoding/json"
)

// checks if the NewRegistry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewRegistry{}

// NewRegistry struct for NewRegistry
type NewRegistry struct {
	Id string `json:"id"`
	Credentials []RegistryCredential `json:"credentials,omitempty"`
	HttpConfig *HttpConfig `json:"httpConfig,omitempty"`
	MqttConfig *MqttConfig `json:"mqttConfig,omitempty"`
	LogLevel *LogLevel `json:"logLevel,omitempty"`
	EventNotificationConfigs []EventNotificationConfig `json:"eventNotificationConfigs,omitempty"`
	LogNotificationConfig *NotificationConfig `json:"logNotificationConfig,omitempty"`
	StateNotificationConfig *NotificationConfig `json:"stateNotificationConfig,omitempty"`
}

// NewNewRegistry instantiates a new NewRegistry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewRegistry(id string) *NewRegistry {
	this := NewRegistry{}
	this.Id = id
	return &this
}

// NewNewRegistryWithDefaults instantiates a new NewRegistry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewRegistryWithDefaults() *NewRegistry {
	this := NewRegistry{}
	return &this
}

// GetId returns the Id field value
func (o *NewRegistry) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NewRegistry) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NewRegistry) SetId(v string) {
	o.Id = v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *NewRegistry) GetCredentials() []RegistryCredential {
	if o == nil || IsNil(o.Credentials) {
		var ret []RegistryCredential
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewRegistry) GetCredentialsOk() ([]RegistryCredential, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *NewRegistry) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given []RegistryCredential and assigns it to the Credentials field.
func (o *NewRegistry) SetCredentials(v []RegistryCredential) {
	o.Credentials = v
}

// GetHttpConfig returns the HttpConfig field value if set, zero value otherwise.
func (o *NewRegistry) GetHttpConfig() HttpConfig {
	if o == nil || IsNil(o.HttpConfig) {
		var ret HttpConfig
		return ret
	}
	return *o.HttpConfig
}

// GetHttpConfigOk returns a tuple with the HttpConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewRegistry) GetHttpConfigOk() (*HttpConfig, bool) {
	if o == nil || IsNil(o.HttpConfig) {
		return nil, false
	}
	return o.HttpConfig, true
}

// HasHttpConfig returns a boolean if a field has been set.
func (o *NewRegistry) HasHttpConfig() bool {
	if o != nil && !IsNil(o.HttpConfig) {
		return true
	}

	return false
}

// SetHttpConfig gets a reference to the given HttpConfig and assigns it to the HttpConfig field.
func (o *NewRegistry) SetHttpConfig(v HttpConfig) {
	o.HttpConfig = &v
}

// GetMqttConfig returns the MqttConfig field value if set, zero value otherwise.
func (o *NewRegistry) GetMqttConfig() MqttConfig {
	if o == nil || IsNil(o.MqttConfig) {
		var ret MqttConfig
		return ret
	}
	return *o.MqttConfig
}

// GetMqttConfigOk returns a tuple with the MqttConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewRegistry) GetMqttConfigOk() (*MqttConfig, bool) {
	if o == nil || IsNil(o.MqttConfig) {
		return nil, false
	}
	return o.MqttConfig, true
}

// HasMqttConfig returns a boolean if a field has been set.
func (o *NewRegistry) HasMqttConfig() bool {
	if o != nil && !IsNil(o.MqttConfig) {
		return true
	}

	return false
}

// SetMqttConfig gets a reference to the given MqttConfig and assigns it to the MqttConfig field.
func (o *NewRegistry) SetMqttConfig(v MqttConfig) {
	o.MqttConfig = &v
}

// GetLogLevel returns the LogLevel field value if set, zero value otherwise.
func (o *NewRegistry) GetLogLevel() LogLevel {
	if o == nil || IsNil(o.LogLevel) {
		var ret LogLevel
		return ret
	}
	return *o.LogLevel
}

// GetLogLevelOk returns a tuple with the LogLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewRegistry) GetLogLevelOk() (*LogLevel, bool) {
	if o == nil || IsNil(o.LogLevel) {
		return nil, false
	}
	return o.LogLevel, true
}

// HasLogLevel returns a boolean if a field has been set.
func (o *NewRegistry) HasLogLevel() bool {
	if o != nil && !IsNil(o.LogLevel) {
		return true
	}

	return false
}

// SetLogLevel gets a reference to the given LogLevel and assigns it to the LogLevel field.
func (o *NewRegistry) SetLogLevel(v LogLevel) {
	o.LogLevel = &v
}

// GetEventNotificationConfigs returns the EventNotificationConfigs field value if set, zero value otherwise.
func (o *NewRegistry) GetEventNotificationConfigs() []EventNotificationConfig {
	if o == nil || IsNil(o.EventNotificationConfigs) {
		var ret []EventNotificationConfig
		return ret
	}
	return o.EventNotificationConfigs
}

// GetEventNotificationConfigsOk returns a tuple with the EventNotificationConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewRegistry) GetEventNotificationConfigsOk() ([]EventNotificationConfig, bool) {
	if o == nil || IsNil(o.EventNotificationConfigs) {
		return nil, false
	}
	return o.EventNotificationConfigs, true
}

// HasEventNotificationConfigs returns a boolean if a field has been set.
func (o *NewRegistry) HasEventNotificationConfigs() bool {
	if o != nil && !IsNil(o.EventNotificationConfigs) {
		return true
	}

	return false
}

// SetEventNotificationConfigs gets a reference to the given []EventNotificationConfig and assigns it to the EventNotificationConfigs field.
func (o *NewRegistry) SetEventNotificationConfigs(v []EventNotificationConfig) {
	o.EventNotificationConfigs = v
}

// GetLogNotificationConfig returns the LogNotificationConfig field value if set, zero value otherwise.
func (o *NewRegistry) GetLogNotificationConfig() NotificationConfig {
	if o == nil || IsNil(o.LogNotificationConfig) {
		var ret NotificationConfig
		return ret
	}
	return *o.LogNotificationConfig
}

// GetLogNotificationConfigOk returns a tuple with the LogNotificationConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewRegistry) GetLogNotificationConfigOk() (*NotificationConfig, bool) {
	if o == nil || IsNil(o.LogNotificationConfig) {
		return nil, false
	}
	return o.LogNotificationConfig, true
}

// HasLogNotificationConfig returns a boolean if a field has been set.
func (o *NewRegistry) HasLogNotificationConfig() bool {
	if o != nil && !IsNil(o.LogNotificationConfig) {
		return true
	}

	return false
}

// SetLogNotificationConfig gets a reference to the given NotificationConfig and assigns it to the LogNotificationConfig field.
func (o *NewRegistry) SetLogNotificationConfig(v NotificationConfig) {
	o.LogNotificationConfig = &v
}

// GetStateNotificationConfig returns the StateNotificationConfig field value if set, zero value otherwise.
func (o *NewRegistry) GetStateNotificationConfig() NotificationConfig {
	if o == nil || IsNil(o.StateNotificationConfig) {
		var ret NotificationConfig
		return ret
	}
	return *o.StateNotificationConfig
}

// GetStateNotificationConfigOk returns a tuple with the StateNotificationConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewRegistry) GetStateNotificationConfigOk() (*NotificationConfig, bool) {
	if o == nil || IsNil(o.StateNotificationConfig) {
		return nil, false
	}
	return o.StateNotificationConfig, true
}

// HasStateNotificationConfig returns a boolean if a field has been set.
func (o *NewRegistry) HasStateNotificationConfig() bool {
	if o != nil && !IsNil(o.StateNotificationConfig) {
		return true
	}

	return false
}

// SetStateNotificationConfig gets a reference to the given NotificationConfig and assigns it to the StateNotificationConfig field.
func (o *NewRegistry) SetStateNotificationConfig(v NotificationConfig) {
	o.StateNotificationConfig = &v
}

func (o NewRegistry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewRegistry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.HttpConfig) {
		toSerialize["httpConfig"] = o.HttpConfig
	}
	if !IsNil(o.MqttConfig) {
		toSerialize["mqttConfig"] = o.MqttConfig
	}
	if !IsNil(o.LogLevel) {
		toSerialize["logLevel"] = o.LogLevel
	}
	if !IsNil(o.EventNotificationConfigs) {
		toSerialize["eventNotificationConfigs"] = o.EventNotificationConfigs
	}
	if !IsNil(o.LogNotificationConfig) {
		toSerialize["logNotificationConfig"] = o.LogNotificationConfig
	}
	if !IsNil(o.StateNotificationConfig) {
		toSerialize["stateNotificationConfig"] = o.StateNotificationConfig
	}
	return toSerialize, nil
}

type NullableNewRegistry struct {
	value *NewRegistry
	isSet bool
}

func (v NullableNewRegistry) Get() *NewRegistry {
	return v.value
}

func (v *NullableNewRegistry) Set(val *NewRegistry) {
	v.value = val
	v.isSet = true
}

func (v NullableNewRegistry) IsSet() bool {
	return v.isSet
}

func (v *NullableNewRegistry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewRegistry(val *NewRegistry) *NullableNewRegistry {
	return &NullableNewRegistry{value: val, isSet: true}
}

func (v NullableNewRegistry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewRegistry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

