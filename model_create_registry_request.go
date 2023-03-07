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

// checks if the CreateRegistryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRegistryRequest{}

// CreateRegistryRequest struct for CreateRegistryRequest
type CreateRegistryRequest struct {
	Credentials []OmnicoreRegistryCredential `json:"credentials,omitempty"`
	EventNotificationConfigs []OmnicoreEventNotificationConfig `json:"eventNotificationConfigs,omitempty"`
	HttpConfig *OmnicoreHttpConfig `json:"httpConfig,omitempty"`
	Id string `json:"id"`
	LogLevel *string `json:"logLevel,omitempty"`
	LogNotificationConfig *OmnicoreStateNotificationConfig `json:"logNotificationConfig,omitempty"`
	MqttConfig *OmnicoreMqttConfig `json:"mqttConfig,omitempty"`
	StateNotificationConfig *OmnicoreStateNotificationConfig `json:"stateNotificationConfig,omitempty"`
}

// NewCreateRegistryRequest instantiates a new CreateRegistryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRegistryRequest(id string) *CreateRegistryRequest {
	this := CreateRegistryRequest{}
	this.Id = id
	return &this
}

// NewCreateRegistryRequestWithDefaults instantiates a new CreateRegistryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRegistryRequestWithDefaults() *CreateRegistryRequest {
	this := CreateRegistryRequest{}
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *CreateRegistryRequest) GetCredentials() []OmnicoreRegistryCredential {
	if o == nil || IsNil(o.Credentials) {
		var ret []OmnicoreRegistryCredential
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRegistryRequest) GetCredentialsOk() ([]OmnicoreRegistryCredential, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *CreateRegistryRequest) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given []OmnicoreRegistryCredential and assigns it to the Credentials field.
func (o *CreateRegistryRequest) SetCredentials(v []OmnicoreRegistryCredential) {
	o.Credentials = v
}

// GetEventNotificationConfigs returns the EventNotificationConfigs field value if set, zero value otherwise.
func (o *CreateRegistryRequest) GetEventNotificationConfigs() []OmnicoreEventNotificationConfig {
	if o == nil || IsNil(o.EventNotificationConfigs) {
		var ret []OmnicoreEventNotificationConfig
		return ret
	}
	return o.EventNotificationConfigs
}

// GetEventNotificationConfigsOk returns a tuple with the EventNotificationConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRegistryRequest) GetEventNotificationConfigsOk() ([]OmnicoreEventNotificationConfig, bool) {
	if o == nil || IsNil(o.EventNotificationConfigs) {
		return nil, false
	}
	return o.EventNotificationConfigs, true
}

// HasEventNotificationConfigs returns a boolean if a field has been set.
func (o *CreateRegistryRequest) HasEventNotificationConfigs() bool {
	if o != nil && !IsNil(o.EventNotificationConfigs) {
		return true
	}

	return false
}

// SetEventNotificationConfigs gets a reference to the given []OmnicoreEventNotificationConfig and assigns it to the EventNotificationConfigs field.
func (o *CreateRegistryRequest) SetEventNotificationConfigs(v []OmnicoreEventNotificationConfig) {
	o.EventNotificationConfigs = v
}

// GetHttpConfig returns the HttpConfig field value if set, zero value otherwise.
func (o *CreateRegistryRequest) GetHttpConfig() OmnicoreHttpConfig {
	if o == nil || IsNil(o.HttpConfig) {
		var ret OmnicoreHttpConfig
		return ret
	}
	return *o.HttpConfig
}

// GetHttpConfigOk returns a tuple with the HttpConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRegistryRequest) GetHttpConfigOk() (*OmnicoreHttpConfig, bool) {
	if o == nil || IsNil(o.HttpConfig) {
		return nil, false
	}
	return o.HttpConfig, true
}

// HasHttpConfig returns a boolean if a field has been set.
func (o *CreateRegistryRequest) HasHttpConfig() bool {
	if o != nil && !IsNil(o.HttpConfig) {
		return true
	}

	return false
}

// SetHttpConfig gets a reference to the given OmnicoreHttpConfig and assigns it to the HttpConfig field.
func (o *CreateRegistryRequest) SetHttpConfig(v OmnicoreHttpConfig) {
	o.HttpConfig = &v
}

// GetId returns the Id field value
func (o *CreateRegistryRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateRegistryRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateRegistryRequest) SetId(v string) {
	o.Id = v
}

// GetLogLevel returns the LogLevel field value if set, zero value otherwise.
func (o *CreateRegistryRequest) GetLogLevel() string {
	if o == nil || IsNil(o.LogLevel) {
		var ret string
		return ret
	}
	return *o.LogLevel
}

// GetLogLevelOk returns a tuple with the LogLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRegistryRequest) GetLogLevelOk() (*string, bool) {
	if o == nil || IsNil(o.LogLevel) {
		return nil, false
	}
	return o.LogLevel, true
}

// HasLogLevel returns a boolean if a field has been set.
func (o *CreateRegistryRequest) HasLogLevel() bool {
	if o != nil && !IsNil(o.LogLevel) {
		return true
	}

	return false
}

// SetLogLevel gets a reference to the given string and assigns it to the LogLevel field.
func (o *CreateRegistryRequest) SetLogLevel(v string) {
	o.LogLevel = &v
}

// GetLogNotificationConfig returns the LogNotificationConfig field value if set, zero value otherwise.
func (o *CreateRegistryRequest) GetLogNotificationConfig() OmnicoreStateNotificationConfig {
	if o == nil || IsNil(o.LogNotificationConfig) {
		var ret OmnicoreStateNotificationConfig
		return ret
	}
	return *o.LogNotificationConfig
}

// GetLogNotificationConfigOk returns a tuple with the LogNotificationConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRegistryRequest) GetLogNotificationConfigOk() (*OmnicoreStateNotificationConfig, bool) {
	if o == nil || IsNil(o.LogNotificationConfig) {
		return nil, false
	}
	return o.LogNotificationConfig, true
}

// HasLogNotificationConfig returns a boolean if a field has been set.
func (o *CreateRegistryRequest) HasLogNotificationConfig() bool {
	if o != nil && !IsNil(o.LogNotificationConfig) {
		return true
	}

	return false
}

// SetLogNotificationConfig gets a reference to the given OmnicoreStateNotificationConfig and assigns it to the LogNotificationConfig field.
func (o *CreateRegistryRequest) SetLogNotificationConfig(v OmnicoreStateNotificationConfig) {
	o.LogNotificationConfig = &v
}

// GetMqttConfig returns the MqttConfig field value if set, zero value otherwise.
func (o *CreateRegistryRequest) GetMqttConfig() OmnicoreMqttConfig {
	if o == nil || IsNil(o.MqttConfig) {
		var ret OmnicoreMqttConfig
		return ret
	}
	return *o.MqttConfig
}

// GetMqttConfigOk returns a tuple with the MqttConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRegistryRequest) GetMqttConfigOk() (*OmnicoreMqttConfig, bool) {
	if o == nil || IsNil(o.MqttConfig) {
		return nil, false
	}
	return o.MqttConfig, true
}

// HasMqttConfig returns a boolean if a field has been set.
func (o *CreateRegistryRequest) HasMqttConfig() bool {
	if o != nil && !IsNil(o.MqttConfig) {
		return true
	}

	return false
}

// SetMqttConfig gets a reference to the given OmnicoreMqttConfig and assigns it to the MqttConfig field.
func (o *CreateRegistryRequest) SetMqttConfig(v OmnicoreMqttConfig) {
	o.MqttConfig = &v
}

// GetStateNotificationConfig returns the StateNotificationConfig field value if set, zero value otherwise.
func (o *CreateRegistryRequest) GetStateNotificationConfig() OmnicoreStateNotificationConfig {
	if o == nil || IsNil(o.StateNotificationConfig) {
		var ret OmnicoreStateNotificationConfig
		return ret
	}
	return *o.StateNotificationConfig
}

// GetStateNotificationConfigOk returns a tuple with the StateNotificationConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRegistryRequest) GetStateNotificationConfigOk() (*OmnicoreStateNotificationConfig, bool) {
	if o == nil || IsNil(o.StateNotificationConfig) {
		return nil, false
	}
	return o.StateNotificationConfig, true
}

// HasStateNotificationConfig returns a boolean if a field has been set.
func (o *CreateRegistryRequest) HasStateNotificationConfig() bool {
	if o != nil && !IsNil(o.StateNotificationConfig) {
		return true
	}

	return false
}

// SetStateNotificationConfig gets a reference to the given OmnicoreStateNotificationConfig and assigns it to the StateNotificationConfig field.
func (o *CreateRegistryRequest) SetStateNotificationConfig(v OmnicoreStateNotificationConfig) {
	o.StateNotificationConfig = &v
}

func (o CreateRegistryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRegistryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.EventNotificationConfigs) {
		toSerialize["eventNotificationConfigs"] = o.EventNotificationConfigs
	}
	if !IsNil(o.HttpConfig) {
		toSerialize["httpConfig"] = o.HttpConfig
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.LogLevel) {
		toSerialize["logLevel"] = o.LogLevel
	}
	if !IsNil(o.LogNotificationConfig) {
		toSerialize["logNotificationConfig"] = o.LogNotificationConfig
	}
	if !IsNil(o.MqttConfig) {
		toSerialize["mqttConfig"] = o.MqttConfig
	}
	if !IsNil(o.StateNotificationConfig) {
		toSerialize["stateNotificationConfig"] = o.StateNotificationConfig
	}
	return toSerialize, nil
}

type NullableCreateRegistryRequest struct {
	value *CreateRegistryRequest
	isSet bool
}

func (v NullableCreateRegistryRequest) Get() *CreateRegistryRequest {
	return v.value
}

func (v *NullableCreateRegistryRequest) Set(val *CreateRegistryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRegistryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRegistryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRegistryRequest(val *CreateRegistryRequest) *NullableCreateRegistryRequest {
	return &NullableCreateRegistryRequest{value: val, isSet: true}
}

func (v NullableCreateRegistryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRegistryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


