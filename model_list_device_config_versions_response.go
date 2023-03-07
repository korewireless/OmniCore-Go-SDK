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

// checks if the ListDeviceConfigVersionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDeviceConfigVersionsResponse{}

// ListDeviceConfigVersionsResponse struct for ListDeviceConfigVersionsResponse
type ListDeviceConfigVersionsResponse struct {
	DeviceConfigs []DeviceConfig `json:"deviceConfigs,omitempty"`
}

// NewListDeviceConfigVersionsResponse instantiates a new ListDeviceConfigVersionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDeviceConfigVersionsResponse() *ListDeviceConfigVersionsResponse {
	this := ListDeviceConfigVersionsResponse{}
	return &this
}

// NewListDeviceConfigVersionsResponseWithDefaults instantiates a new ListDeviceConfigVersionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDeviceConfigVersionsResponseWithDefaults() *ListDeviceConfigVersionsResponse {
	this := ListDeviceConfigVersionsResponse{}
	return &this
}

// GetDeviceConfigs returns the DeviceConfigs field value if set, zero value otherwise.
func (o *ListDeviceConfigVersionsResponse) GetDeviceConfigs() []DeviceConfig {
	if o == nil || IsNil(o.DeviceConfigs) {
		var ret []DeviceConfig
		return ret
	}
	return o.DeviceConfigs
}

// GetDeviceConfigsOk returns a tuple with the DeviceConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDeviceConfigVersionsResponse) GetDeviceConfigsOk() ([]DeviceConfig, bool) {
	if o == nil || IsNil(o.DeviceConfigs) {
		return nil, false
	}
	return o.DeviceConfigs, true
}

// HasDeviceConfigs returns a boolean if a field has been set.
func (o *ListDeviceConfigVersionsResponse) HasDeviceConfigs() bool {
	if o != nil && !IsNil(o.DeviceConfigs) {
		return true
	}

	return false
}

// SetDeviceConfigs gets a reference to the given []DeviceConfig and assigns it to the DeviceConfigs field.
func (o *ListDeviceConfigVersionsResponse) SetDeviceConfigs(v []DeviceConfig) {
	o.DeviceConfigs = v
}

func (o ListDeviceConfigVersionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDeviceConfigVersionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceConfigs) {
		toSerialize["deviceConfigs"] = o.DeviceConfigs
	}
	return toSerialize, nil
}

type NullableListDeviceConfigVersionsResponse struct {
	value *ListDeviceConfigVersionsResponse
	isSet bool
}

func (v NullableListDeviceConfigVersionsResponse) Get() *ListDeviceConfigVersionsResponse {
	return v.value
}

func (v *NullableListDeviceConfigVersionsResponse) Set(val *ListDeviceConfigVersionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListDeviceConfigVersionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListDeviceConfigVersionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDeviceConfigVersionsResponse(val *ListDeviceConfigVersionsResponse) *NullableListDeviceConfigVersionsResponse {
	return &NullableListDeviceConfigVersionsResponse{value: val, isSet: true}
}

func (v NullableListDeviceConfigVersionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDeviceConfigVersionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


