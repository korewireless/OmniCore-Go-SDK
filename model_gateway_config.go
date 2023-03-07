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

// checks if the GatewayConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayConfig{}

// GatewayConfig struct for GatewayConfig
type GatewayConfig struct {
	GatewayAuthMethod *string `json:"gatewayAuthMethod,omitempty"`
	GatewayType *string `json:"gatewayType,omitempty"`
}

// NewGatewayConfig instantiates a new GatewayConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayConfig() *GatewayConfig {
	this := GatewayConfig{}
	return &this
}

// NewGatewayConfigWithDefaults instantiates a new GatewayConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayConfigWithDefaults() *GatewayConfig {
	this := GatewayConfig{}
	return &this
}

// GetGatewayAuthMethod returns the GatewayAuthMethod field value if set, zero value otherwise.
func (o *GatewayConfig) GetGatewayAuthMethod() string {
	if o == nil || IsNil(o.GatewayAuthMethod) {
		var ret string
		return ret
	}
	return *o.GatewayAuthMethod
}

// GetGatewayAuthMethodOk returns a tuple with the GatewayAuthMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayConfig) GetGatewayAuthMethodOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayAuthMethod) {
		return nil, false
	}
	return o.GatewayAuthMethod, true
}

// HasGatewayAuthMethod returns a boolean if a field has been set.
func (o *GatewayConfig) HasGatewayAuthMethod() bool {
	if o != nil && !IsNil(o.GatewayAuthMethod) {
		return true
	}

	return false
}

// SetGatewayAuthMethod gets a reference to the given string and assigns it to the GatewayAuthMethod field.
func (o *GatewayConfig) SetGatewayAuthMethod(v string) {
	o.GatewayAuthMethod = &v
}

// GetGatewayType returns the GatewayType field value if set, zero value otherwise.
func (o *GatewayConfig) GetGatewayType() string {
	if o == nil || IsNil(o.GatewayType) {
		var ret string
		return ret
	}
	return *o.GatewayType
}

// GetGatewayTypeOk returns a tuple with the GatewayType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayConfig) GetGatewayTypeOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayType) {
		return nil, false
	}
	return o.GatewayType, true
}

// HasGatewayType returns a boolean if a field has been set.
func (o *GatewayConfig) HasGatewayType() bool {
	if o != nil && !IsNil(o.GatewayType) {
		return true
	}

	return false
}

// SetGatewayType gets a reference to the given string and assigns it to the GatewayType field.
func (o *GatewayConfig) SetGatewayType(v string) {
	o.GatewayType = &v
}

func (o GatewayConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GatewayAuthMethod) {
		toSerialize["gatewayAuthMethod"] = o.GatewayAuthMethod
	}
	if !IsNil(o.GatewayType) {
		toSerialize["gatewayType"] = o.GatewayType
	}
	return toSerialize, nil
}

type NullableGatewayConfig struct {
	value *GatewayConfig
	isSet bool
}

func (v NullableGatewayConfig) Get() *GatewayConfig {
	return v.value
}

func (v *NullableGatewayConfig) Set(val *GatewayConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayConfig(val *GatewayConfig) *NullableGatewayConfig {
	return &NullableGatewayConfig{value: val, isSet: true}
}

func (v NullableGatewayConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


