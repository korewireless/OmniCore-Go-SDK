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

// checks if the OmnicoreGatewayConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OmnicoreGatewayConfig{}

// OmnicoreGatewayConfig struct for OmnicoreGatewayConfig
type OmnicoreGatewayConfig struct {
	GatewayAuthMethod *string `json:"gatewayAuthMethod,omitempty"`
	GatewayType *string `json:"gatewayType,omitempty"`
}

// NewOmnicoreGatewayConfig instantiates a new OmnicoreGatewayConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOmnicoreGatewayConfig() *OmnicoreGatewayConfig {
	this := OmnicoreGatewayConfig{}
	return &this
}

// NewOmnicoreGatewayConfigWithDefaults instantiates a new OmnicoreGatewayConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOmnicoreGatewayConfigWithDefaults() *OmnicoreGatewayConfig {
	this := OmnicoreGatewayConfig{}
	return &this
}

// GetGatewayAuthMethod returns the GatewayAuthMethod field value if set, zero value otherwise.
func (o *OmnicoreGatewayConfig) GetGatewayAuthMethod() string {
	if o == nil || IsNil(o.GatewayAuthMethod) {
		var ret string
		return ret
	}
	return *o.GatewayAuthMethod
}

// GetGatewayAuthMethodOk returns a tuple with the GatewayAuthMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreGatewayConfig) GetGatewayAuthMethodOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayAuthMethod) {
		return nil, false
	}
	return o.GatewayAuthMethod, true
}

// HasGatewayAuthMethod returns a boolean if a field has been set.
func (o *OmnicoreGatewayConfig) HasGatewayAuthMethod() bool {
	if o != nil && !IsNil(o.GatewayAuthMethod) {
		return true
	}

	return false
}

// SetGatewayAuthMethod gets a reference to the given string and assigns it to the GatewayAuthMethod field.
func (o *OmnicoreGatewayConfig) SetGatewayAuthMethod(v string) {
	o.GatewayAuthMethod = &v
}

// GetGatewayType returns the GatewayType field value if set, zero value otherwise.
func (o *OmnicoreGatewayConfig) GetGatewayType() string {
	if o == nil || IsNil(o.GatewayType) {
		var ret string
		return ret
	}
	return *o.GatewayType
}

// GetGatewayTypeOk returns a tuple with the GatewayType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreGatewayConfig) GetGatewayTypeOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayType) {
		return nil, false
	}
	return o.GatewayType, true
}

// HasGatewayType returns a boolean if a field has been set.
func (o *OmnicoreGatewayConfig) HasGatewayType() bool {
	if o != nil && !IsNil(o.GatewayType) {
		return true
	}

	return false
}

// SetGatewayType gets a reference to the given string and assigns it to the GatewayType field.
func (o *OmnicoreGatewayConfig) SetGatewayType(v string) {
	o.GatewayType = &v
}

func (o OmnicoreGatewayConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OmnicoreGatewayConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GatewayAuthMethod) {
		toSerialize["gatewayAuthMethod"] = o.GatewayAuthMethod
	}
	if !IsNil(o.GatewayType) {
		toSerialize["gatewayType"] = o.GatewayType
	}
	return toSerialize, nil
}

type NullableOmnicoreGatewayConfig struct {
	value *OmnicoreGatewayConfig
	isSet bool
}

func (v NullableOmnicoreGatewayConfig) Get() *OmnicoreGatewayConfig {
	return v.value
}

func (v *NullableOmnicoreGatewayConfig) Set(val *OmnicoreGatewayConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableOmnicoreGatewayConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableOmnicoreGatewayConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOmnicoreGatewayConfig(val *OmnicoreGatewayConfig) *NullableOmnicoreGatewayConfig {
	return &NullableOmnicoreGatewayConfig{value: val, isSet: true}
}

func (v NullableOmnicoreGatewayConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOmnicoreGatewayConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


