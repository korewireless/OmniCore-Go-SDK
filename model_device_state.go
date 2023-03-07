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

// checks if the DeviceState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceState{}

// DeviceState struct for DeviceState
type DeviceState struct {
	// Base64 Encoded State String
	BinaryData *string `json:"binaryData,omitempty"`
	UpdateTime *string `json:"updateTime,omitempty"`
}

// NewDeviceState instantiates a new DeviceState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceState() *DeviceState {
	this := DeviceState{}
	return &this
}

// NewDeviceStateWithDefaults instantiates a new DeviceState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceStateWithDefaults() *DeviceState {
	this := DeviceState{}
	return &this
}

// GetBinaryData returns the BinaryData field value if set, zero value otherwise.
func (o *DeviceState) GetBinaryData() string {
	if o == nil || IsNil(o.BinaryData) {
		var ret string
		return ret
	}
	return *o.BinaryData
}

// GetBinaryDataOk returns a tuple with the BinaryData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceState) GetBinaryDataOk() (*string, bool) {
	if o == nil || IsNil(o.BinaryData) {
		return nil, false
	}
	return o.BinaryData, true
}

// HasBinaryData returns a boolean if a field has been set.
func (o *DeviceState) HasBinaryData() bool {
	if o != nil && !IsNil(o.BinaryData) {
		return true
	}

	return false
}

// SetBinaryData gets a reference to the given string and assigns it to the BinaryData field.
func (o *DeviceState) SetBinaryData(v string) {
	o.BinaryData = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *DeviceState) GetUpdateTime() string {
	if o == nil || IsNil(o.UpdateTime) {
		var ret string
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceState) GetUpdateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *DeviceState) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given string and assigns it to the UpdateTime field.
func (o *DeviceState) SetUpdateTime(v string) {
	o.UpdateTime = &v
}

func (o DeviceState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BinaryData) {
		toSerialize["binaryData"] = o.BinaryData
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	return toSerialize, nil
}

type NullableDeviceState struct {
	value *DeviceState
	isSet bool
}

func (v NullableDeviceState) Get() *DeviceState {
	return v.value
}

func (v *NullableDeviceState) Set(val *DeviceState) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceState) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceState(val *DeviceState) *NullableDeviceState {
	return &NullableDeviceState{value: val, isSet: true}
}

func (v NullableDeviceState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


