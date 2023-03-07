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

// checks if the OmnicoreMqttConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OmnicoreMqttConfig{}

// OmnicoreMqttConfig struct for OmnicoreMqttConfig
type OmnicoreMqttConfig struct {
	// MqttEnabledState: If enabled, allows connections using the MQTT protocol. Otherwise, MQTT connections to this registry will fail.  Possible values:   \"MQTT_STATE_UNSPECIFIED\" - No MQTT state specified. If not specified, MQTT will be enabled by default.   \"MQTT_ENABLED\" - Enables a MQTT connection.   \"MQTT_DISABLED\" - Disables a MQTT connection.
	MqttEnabledState *string `json:"mqttEnabledState,omitempty"`
}

// NewOmnicoreMqttConfig instantiates a new OmnicoreMqttConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOmnicoreMqttConfig() *OmnicoreMqttConfig {
	this := OmnicoreMqttConfig{}
	return &this
}

// NewOmnicoreMqttConfigWithDefaults instantiates a new OmnicoreMqttConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOmnicoreMqttConfigWithDefaults() *OmnicoreMqttConfig {
	this := OmnicoreMqttConfig{}
	return &this
}

// GetMqttEnabledState returns the MqttEnabledState field value if set, zero value otherwise.
func (o *OmnicoreMqttConfig) GetMqttEnabledState() string {
	if o == nil || IsNil(o.MqttEnabledState) {
		var ret string
		return ret
	}
	return *o.MqttEnabledState
}

// GetMqttEnabledStateOk returns a tuple with the MqttEnabledState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreMqttConfig) GetMqttEnabledStateOk() (*string, bool) {
	if o == nil || IsNil(o.MqttEnabledState) {
		return nil, false
	}
	return o.MqttEnabledState, true
}

// HasMqttEnabledState returns a boolean if a field has been set.
func (o *OmnicoreMqttConfig) HasMqttEnabledState() bool {
	if o != nil && !IsNil(o.MqttEnabledState) {
		return true
	}

	return false
}

// SetMqttEnabledState gets a reference to the given string and assigns it to the MqttEnabledState field.
func (o *OmnicoreMqttConfig) SetMqttEnabledState(v string) {
	o.MqttEnabledState = &v
}

func (o OmnicoreMqttConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OmnicoreMqttConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MqttEnabledState) {
		toSerialize["mqttEnabledState"] = o.MqttEnabledState
	}
	return toSerialize, nil
}

type NullableOmnicoreMqttConfig struct {
	value *OmnicoreMqttConfig
	isSet bool
}

func (v NullableOmnicoreMqttConfig) Get() *OmnicoreMqttConfig {
	return v.value
}

func (v *NullableOmnicoreMqttConfig) Set(val *OmnicoreMqttConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableOmnicoreMqttConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableOmnicoreMqttConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOmnicoreMqttConfig(val *OmnicoreMqttConfig) *NullableOmnicoreMqttConfig {
	return &NullableOmnicoreMqttConfig{value: val, isSet: true}
}

func (v NullableOmnicoreMqttConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOmnicoreMqttConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


