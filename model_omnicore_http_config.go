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

// checks if the OmnicoreHttpConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OmnicoreHttpConfig{}

// OmnicoreHttpConfig struct for OmnicoreHttpConfig
type OmnicoreHttpConfig struct {
	// HttpEnabledState: If enabled, allows devices to use DeviceService via the HTTP protocol. Otherwise, any requests to DeviceService will fail for this registry.  Possible values:   \"HTTP_STATE_UNSPECIFIED\" - No HTTP state specified. If not specified, DeviceService will be enabled by default.   \"HTTP_ENABLED\" - Enables DeviceService (HTTP) service for the registry.   \"HTTP_DISABLED\" - Disables DeviceService (HTTP) service for the registry.
	HttpEnabledState *string `json:"httpEnabledState,omitempty"`
}

// NewOmnicoreHttpConfig instantiates a new OmnicoreHttpConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOmnicoreHttpConfig() *OmnicoreHttpConfig {
	this := OmnicoreHttpConfig{}
	return &this
}

// NewOmnicoreHttpConfigWithDefaults instantiates a new OmnicoreHttpConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOmnicoreHttpConfigWithDefaults() *OmnicoreHttpConfig {
	this := OmnicoreHttpConfig{}
	return &this
}

// GetHttpEnabledState returns the HttpEnabledState field value if set, zero value otherwise.
func (o *OmnicoreHttpConfig) GetHttpEnabledState() string {
	if o == nil || IsNil(o.HttpEnabledState) {
		var ret string
		return ret
	}
	return *o.HttpEnabledState
}

// GetHttpEnabledStateOk returns a tuple with the HttpEnabledState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreHttpConfig) GetHttpEnabledStateOk() (*string, bool) {
	if o == nil || IsNil(o.HttpEnabledState) {
		return nil, false
	}
	return o.HttpEnabledState, true
}

// HasHttpEnabledState returns a boolean if a field has been set.
func (o *OmnicoreHttpConfig) HasHttpEnabledState() bool {
	if o != nil && !IsNil(o.HttpEnabledState) {
		return true
	}

	return false
}

// SetHttpEnabledState gets a reference to the given string and assigns it to the HttpEnabledState field.
func (o *OmnicoreHttpConfig) SetHttpEnabledState(v string) {
	o.HttpEnabledState = &v
}

func (o OmnicoreHttpConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OmnicoreHttpConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HttpEnabledState) {
		toSerialize["httpEnabledState"] = o.HttpEnabledState
	}
	return toSerialize, nil
}

type NullableOmnicoreHttpConfig struct {
	value *OmnicoreHttpConfig
	isSet bool
}

func (v NullableOmnicoreHttpConfig) Get() *OmnicoreHttpConfig {
	return v.value
}

func (v *NullableOmnicoreHttpConfig) Set(val *OmnicoreHttpConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableOmnicoreHttpConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableOmnicoreHttpConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOmnicoreHttpConfig(val *OmnicoreHttpConfig) *NullableOmnicoreHttpConfig {
	return &NullableOmnicoreHttpConfig{value: val, isSet: true}
}

func (v NullableOmnicoreHttpConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOmnicoreHttpConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


