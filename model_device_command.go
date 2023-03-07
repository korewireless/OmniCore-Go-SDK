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

// checks if the DeviceCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceCommand{}

// DeviceCommand struct for DeviceCommand
type DeviceCommand struct {
	// Base64 Encoded Command String
	BinaryData string `json:"binaryData"`
	Subfolder *string `json:"subfolder,omitempty"`
}

// NewDeviceCommand instantiates a new DeviceCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceCommand(binaryData string) *DeviceCommand {
	this := DeviceCommand{}
	this.BinaryData = binaryData
	return &this
}

// NewDeviceCommandWithDefaults instantiates a new DeviceCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceCommandWithDefaults() *DeviceCommand {
	this := DeviceCommand{}
	return &this
}

// GetBinaryData returns the BinaryData field value
func (o *DeviceCommand) GetBinaryData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BinaryData
}

// GetBinaryDataOk returns a tuple with the BinaryData field value
// and a boolean to check if the value has been set.
func (o *DeviceCommand) GetBinaryDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BinaryData, true
}

// SetBinaryData sets field value
func (o *DeviceCommand) SetBinaryData(v string) {
	o.BinaryData = v
}

// GetSubfolder returns the Subfolder field value if set, zero value otherwise.
func (o *DeviceCommand) GetSubfolder() string {
	if o == nil || IsNil(o.Subfolder) {
		var ret string
		return ret
	}
	return *o.Subfolder
}

// GetSubfolderOk returns a tuple with the Subfolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCommand) GetSubfolderOk() (*string, bool) {
	if o == nil || IsNil(o.Subfolder) {
		return nil, false
	}
	return o.Subfolder, true
}

// HasSubfolder returns a boolean if a field has been set.
func (o *DeviceCommand) HasSubfolder() bool {
	if o != nil && !IsNil(o.Subfolder) {
		return true
	}

	return false
}

// SetSubfolder gets a reference to the given string and assigns it to the Subfolder field.
func (o *DeviceCommand) SetSubfolder(v string) {
	o.Subfolder = &v
}

func (o DeviceCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["binaryData"] = o.BinaryData
	if !IsNil(o.Subfolder) {
		toSerialize["subfolder"] = o.Subfolder
	}
	return toSerialize, nil
}

type NullableDeviceCommand struct {
	value *DeviceCommand
	isSet bool
}

func (v NullableDeviceCommand) Get() *DeviceCommand {
	return v.value
}

func (v *NullableDeviceCommand) Set(val *DeviceCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceCommand(val *DeviceCommand) *NullableDeviceCommand {
	return &NullableDeviceCommand{value: val, isSet: true}
}

func (v NullableDeviceCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


