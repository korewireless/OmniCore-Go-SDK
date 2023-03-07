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

// checks if the OmnicoreListDeviceConfigVersionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OmnicoreListDeviceConfigVersionsResponse{}

// OmnicoreListDeviceConfigVersionsResponse struct for OmnicoreListDeviceConfigVersionsResponse
type OmnicoreListDeviceConfigVersionsResponse struct {
	DeviceConfigs []OmnicoreDeviceConfig `json:"deviceConfigs,omitempty"`
}

// NewOmnicoreListDeviceConfigVersionsResponse instantiates a new OmnicoreListDeviceConfigVersionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOmnicoreListDeviceConfigVersionsResponse() *OmnicoreListDeviceConfigVersionsResponse {
	this := OmnicoreListDeviceConfigVersionsResponse{}
	return &this
}

// NewOmnicoreListDeviceConfigVersionsResponseWithDefaults instantiates a new OmnicoreListDeviceConfigVersionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOmnicoreListDeviceConfigVersionsResponseWithDefaults() *OmnicoreListDeviceConfigVersionsResponse {
	this := OmnicoreListDeviceConfigVersionsResponse{}
	return &this
}

// GetDeviceConfigs returns the DeviceConfigs field value if set, zero value otherwise.
func (o *OmnicoreListDeviceConfigVersionsResponse) GetDeviceConfigs() []OmnicoreDeviceConfig {
	if o == nil || IsNil(o.DeviceConfigs) {
		var ret []OmnicoreDeviceConfig
		return ret
	}
	return o.DeviceConfigs
}

// GetDeviceConfigsOk returns a tuple with the DeviceConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreListDeviceConfigVersionsResponse) GetDeviceConfigsOk() ([]OmnicoreDeviceConfig, bool) {
	if o == nil || IsNil(o.DeviceConfigs) {
		return nil, false
	}
	return o.DeviceConfigs, true
}

// HasDeviceConfigs returns a boolean if a field has been set.
func (o *OmnicoreListDeviceConfigVersionsResponse) HasDeviceConfigs() bool {
	if o != nil && !IsNil(o.DeviceConfigs) {
		return true
	}

	return false
}

// SetDeviceConfigs gets a reference to the given []OmnicoreDeviceConfig and assigns it to the DeviceConfigs field.
func (o *OmnicoreListDeviceConfigVersionsResponse) SetDeviceConfigs(v []OmnicoreDeviceConfig) {
	o.DeviceConfigs = v
}

func (o OmnicoreListDeviceConfigVersionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OmnicoreListDeviceConfigVersionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceConfigs) {
		toSerialize["deviceConfigs"] = o.DeviceConfigs
	}
	return toSerialize, nil
}

type NullableOmnicoreListDeviceConfigVersionsResponse struct {
	value *OmnicoreListDeviceConfigVersionsResponse
	isSet bool
}

func (v NullableOmnicoreListDeviceConfigVersionsResponse) Get() *OmnicoreListDeviceConfigVersionsResponse {
	return v.value
}

func (v *NullableOmnicoreListDeviceConfigVersionsResponse) Set(val *OmnicoreListDeviceConfigVersionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOmnicoreListDeviceConfigVersionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOmnicoreListDeviceConfigVersionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOmnicoreListDeviceConfigVersionsResponse(val *OmnicoreListDeviceConfigVersionsResponse) *NullableOmnicoreListDeviceConfigVersionsResponse {
	return &NullableOmnicoreListDeviceConfigVersionsResponse{value: val, isSet: true}
}

func (v NullableOmnicoreListDeviceConfigVersionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOmnicoreListDeviceConfigVersionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


