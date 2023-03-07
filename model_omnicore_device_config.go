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

// checks if the OmnicoreDeviceConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OmnicoreDeviceConfig{}

// OmnicoreDeviceConfig struct for OmnicoreDeviceConfig
type OmnicoreDeviceConfig struct {
	Acknowledged *bool `json:"acknowledged,omitempty"`
	BinaryData *string `json:"binaryData,omitempty"`
	CloudUpdateTime *string `json:"cloudUpdateTime,omitempty"`
	DeviceAckTime *string `json:"deviceAckTime,omitempty"`
	Version *int32 `json:"version,omitempty"`
}

// NewOmnicoreDeviceConfig instantiates a new OmnicoreDeviceConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOmnicoreDeviceConfig() *OmnicoreDeviceConfig {
	this := OmnicoreDeviceConfig{}
	return &this
}

// NewOmnicoreDeviceConfigWithDefaults instantiates a new OmnicoreDeviceConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOmnicoreDeviceConfigWithDefaults() *OmnicoreDeviceConfig {
	this := OmnicoreDeviceConfig{}
	return &this
}

// GetAcknowledged returns the Acknowledged field value if set, zero value otherwise.
func (o *OmnicoreDeviceConfig) GetAcknowledged() bool {
	if o == nil || IsNil(o.Acknowledged) {
		var ret bool
		return ret
	}
	return *o.Acknowledged
}

// GetAcknowledgedOk returns a tuple with the Acknowledged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDeviceConfig) GetAcknowledgedOk() (*bool, bool) {
	if o == nil || IsNil(o.Acknowledged) {
		return nil, false
	}
	return o.Acknowledged, true
}

// HasAcknowledged returns a boolean if a field has been set.
func (o *OmnicoreDeviceConfig) HasAcknowledged() bool {
	if o != nil && !IsNil(o.Acknowledged) {
		return true
	}

	return false
}

// SetAcknowledged gets a reference to the given bool and assigns it to the Acknowledged field.
func (o *OmnicoreDeviceConfig) SetAcknowledged(v bool) {
	o.Acknowledged = &v
}

// GetBinaryData returns the BinaryData field value if set, zero value otherwise.
func (o *OmnicoreDeviceConfig) GetBinaryData() string {
	if o == nil || IsNil(o.BinaryData) {
		var ret string
		return ret
	}
	return *o.BinaryData
}

// GetBinaryDataOk returns a tuple with the BinaryData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDeviceConfig) GetBinaryDataOk() (*string, bool) {
	if o == nil || IsNil(o.BinaryData) {
		return nil, false
	}
	return o.BinaryData, true
}

// HasBinaryData returns a boolean if a field has been set.
func (o *OmnicoreDeviceConfig) HasBinaryData() bool {
	if o != nil && !IsNil(o.BinaryData) {
		return true
	}

	return false
}

// SetBinaryData gets a reference to the given string and assigns it to the BinaryData field.
func (o *OmnicoreDeviceConfig) SetBinaryData(v string) {
	o.BinaryData = &v
}

// GetCloudUpdateTime returns the CloudUpdateTime field value if set, zero value otherwise.
func (o *OmnicoreDeviceConfig) GetCloudUpdateTime() string {
	if o == nil || IsNil(o.CloudUpdateTime) {
		var ret string
		return ret
	}
	return *o.CloudUpdateTime
}

// GetCloudUpdateTimeOk returns a tuple with the CloudUpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDeviceConfig) GetCloudUpdateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CloudUpdateTime) {
		return nil, false
	}
	return o.CloudUpdateTime, true
}

// HasCloudUpdateTime returns a boolean if a field has been set.
func (o *OmnicoreDeviceConfig) HasCloudUpdateTime() bool {
	if o != nil && !IsNil(o.CloudUpdateTime) {
		return true
	}

	return false
}

// SetCloudUpdateTime gets a reference to the given string and assigns it to the CloudUpdateTime field.
func (o *OmnicoreDeviceConfig) SetCloudUpdateTime(v string) {
	o.CloudUpdateTime = &v
}

// GetDeviceAckTime returns the DeviceAckTime field value if set, zero value otherwise.
func (o *OmnicoreDeviceConfig) GetDeviceAckTime() string {
	if o == nil || IsNil(o.DeviceAckTime) {
		var ret string
		return ret
	}
	return *o.DeviceAckTime
}

// GetDeviceAckTimeOk returns a tuple with the DeviceAckTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDeviceConfig) GetDeviceAckTimeOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceAckTime) {
		return nil, false
	}
	return o.DeviceAckTime, true
}

// HasDeviceAckTime returns a boolean if a field has been set.
func (o *OmnicoreDeviceConfig) HasDeviceAckTime() bool {
	if o != nil && !IsNil(o.DeviceAckTime) {
		return true
	}

	return false
}

// SetDeviceAckTime gets a reference to the given string and assigns it to the DeviceAckTime field.
func (o *OmnicoreDeviceConfig) SetDeviceAckTime(v string) {
	o.DeviceAckTime = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *OmnicoreDeviceConfig) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDeviceConfig) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *OmnicoreDeviceConfig) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *OmnicoreDeviceConfig) SetVersion(v int32) {
	o.Version = &v
}

func (o OmnicoreDeviceConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OmnicoreDeviceConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Acknowledged) {
		toSerialize["acknowledged"] = o.Acknowledged
	}
	if !IsNil(o.BinaryData) {
		toSerialize["binaryData"] = o.BinaryData
	}
	if !IsNil(o.CloudUpdateTime) {
		toSerialize["cloudUpdateTime"] = o.CloudUpdateTime
	}
	if !IsNil(o.DeviceAckTime) {
		toSerialize["deviceAckTime"] = o.DeviceAckTime
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableOmnicoreDeviceConfig struct {
	value *OmnicoreDeviceConfig
	isSet bool
}

func (v NullableOmnicoreDeviceConfig) Get() *OmnicoreDeviceConfig {
	return v.value
}

func (v *NullableOmnicoreDeviceConfig) Set(val *OmnicoreDeviceConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableOmnicoreDeviceConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableOmnicoreDeviceConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOmnicoreDeviceConfig(val *OmnicoreDeviceConfig) *NullableOmnicoreDeviceConfig {
	return &NullableOmnicoreDeviceConfig{value: val, isSet: true}
}

func (v NullableOmnicoreDeviceConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOmnicoreDeviceConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


