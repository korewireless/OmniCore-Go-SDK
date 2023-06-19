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

// checks if the Config type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Config{}

// Config struct for Config
type Config struct {
	ConnectionParameter *string `json:"connectionParameter,omitempty"`
}

// NewConfig instantiates a new Config object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfig() *Config {
	this := Config{}
	return &this
}

// NewConfigWithDefaults instantiates a new Config object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigWithDefaults() *Config {
	this := Config{}
	return &this
}

// GetConnectionParameter returns the ConnectionParameter field value if set, zero value otherwise.
func (o *Config) GetConnectionParameter() string {
	if o == nil || IsNil(o.ConnectionParameter) {
		var ret string
		return ret
	}
	return *o.ConnectionParameter
}

// GetConnectionParameterOk returns a tuple with the ConnectionParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Config) GetConnectionParameterOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionParameter) {
		return nil, false
	}
	return o.ConnectionParameter, true
}

// HasConnectionParameter returns a boolean if a field has been set.
func (o *Config) HasConnectionParameter() bool {
	if o != nil && !IsNil(o.ConnectionParameter) {
		return true
	}

	return false
}

// SetConnectionParameter gets a reference to the given string and assigns it to the ConnectionParameter field.
func (o *Config) SetConnectionParameter(v string) {
	o.ConnectionParameter = &v
}

func (o Config) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Config) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConnectionParameter) {
		toSerialize["connectionParameter"] = o.ConnectionParameter
	}
	return toSerialize, nil
}

type NullableConfig struct {
	value *Config
	isSet bool
}

func (v NullableConfig) Get() *Config {
	return v.value
}

func (v *NullableConfig) Set(val *Config) {
	v.value = val
	v.isSet = true
}

func (v NullableConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfig(val *Config) *NullableConfig {
	return &NullableConfig{value: val, isSet: true}
}

func (v NullableConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


