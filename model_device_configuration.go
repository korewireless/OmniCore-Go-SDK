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

// checks if the DeviceConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceConfiguration{}

// DeviceConfiguration struct for DeviceConfiguration
type DeviceConfiguration struct {
	// Base64 Encoded Config String
	BinaryData string `json:"binaryData"`
	Subfolder *string `json:"subfolder,omitempty"`
	VersionToUpdate *string `json:"versionToUpdate,omitempty"`
}

// NewDeviceConfiguration instantiates a new DeviceConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceConfiguration(binaryData string) *DeviceConfiguration {
	this := DeviceConfiguration{}
	this.BinaryData = binaryData
	return &this
}

// NewDeviceConfigurationWithDefaults instantiates a new DeviceConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceConfigurationWithDefaults() *DeviceConfiguration {
	this := DeviceConfiguration{}
	return &this
}

// GetBinaryData returns the BinaryData field value
func (o *DeviceConfiguration) GetBinaryData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BinaryData
}

// GetBinaryDataOk returns a tuple with the BinaryData field value
// and a boolean to check if the value has been set.
func (o *DeviceConfiguration) GetBinaryDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BinaryData, true
}

// SetBinaryData sets field value
func (o *DeviceConfiguration) SetBinaryData(v string) {
	o.BinaryData = v
}

// GetSubfolder returns the Subfolder field value if set, zero value otherwise.
func (o *DeviceConfiguration) GetSubfolder() string {
	if o == nil || IsNil(o.Subfolder) {
		var ret string
		return ret
	}
	return *o.Subfolder
}

// GetSubfolderOk returns a tuple with the Subfolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceConfiguration) GetSubfolderOk() (*string, bool) {
	if o == nil || IsNil(o.Subfolder) {
		return nil, false
	}
	return o.Subfolder, true
}

// HasSubfolder returns a boolean if a field has been set.
func (o *DeviceConfiguration) HasSubfolder() bool {
	if o != nil && !IsNil(o.Subfolder) {
		return true
	}

	return false
}

// SetSubfolder gets a reference to the given string and assigns it to the Subfolder field.
func (o *DeviceConfiguration) SetSubfolder(v string) {
	o.Subfolder = &v
}

// GetVersionToUpdate returns the VersionToUpdate field value if set, zero value otherwise.
func (o *DeviceConfiguration) GetVersionToUpdate() string {
	if o == nil || IsNil(o.VersionToUpdate) {
		var ret string
		return ret
	}
	return *o.VersionToUpdate
}

// GetVersionToUpdateOk returns a tuple with the VersionToUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceConfiguration) GetVersionToUpdateOk() (*string, bool) {
	if o == nil || IsNil(o.VersionToUpdate) {
		return nil, false
	}
	return o.VersionToUpdate, true
}

// HasVersionToUpdate returns a boolean if a field has been set.
func (o *DeviceConfiguration) HasVersionToUpdate() bool {
	if o != nil && !IsNil(o.VersionToUpdate) {
		return true
	}

	return false
}

// SetVersionToUpdate gets a reference to the given string and assigns it to the VersionToUpdate field.
func (o *DeviceConfiguration) SetVersionToUpdate(v string) {
	o.VersionToUpdate = &v
}

func (o DeviceConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["binaryData"] = o.BinaryData
	if !IsNil(o.Subfolder) {
		toSerialize["subfolder"] = o.Subfolder
	}
	if !IsNil(o.VersionToUpdate) {
		toSerialize["versionToUpdate"] = o.VersionToUpdate
	}
	return toSerialize, nil
}

type NullableDeviceConfiguration struct {
	value *DeviceConfiguration
	isSet bool
}

func (v NullableDeviceConfiguration) Get() *DeviceConfiguration {
	return v.value
}

func (v *NullableDeviceConfiguration) Set(val *DeviceConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceConfiguration(val *DeviceConfiguration) *NullableDeviceConfiguration {
	return &NullableDeviceConfiguration{value: val, isSet: true}
}

func (v NullableDeviceConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


