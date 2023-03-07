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

// checks if the UpdateDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDevice{}

// UpdateDevice struct for UpdateDevice
type UpdateDevice struct {
	Blocked *bool `json:"blocked,omitempty"`
	Credentials []DeviceCredential `json:"credentials,omitempty"`
	LogLevel *LogLevel `json:"logLevel,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
}

// NewUpdateDevice instantiates a new UpdateDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDevice() *UpdateDevice {
	this := UpdateDevice{}
	return &this
}

// NewUpdateDeviceWithDefaults instantiates a new UpdateDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceWithDefaults() *UpdateDevice {
	this := UpdateDevice{}
	return &this
}

// GetBlocked returns the Blocked field value if set, zero value otherwise.
func (o *UpdateDevice) GetBlocked() bool {
	if o == nil || IsNil(o.Blocked) {
		var ret bool
		return ret
	}
	return *o.Blocked
}

// GetBlockedOk returns a tuple with the Blocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDevice) GetBlockedOk() (*bool, bool) {
	if o == nil || IsNil(o.Blocked) {
		return nil, false
	}
	return o.Blocked, true
}

// HasBlocked returns a boolean if a field has been set.
func (o *UpdateDevice) HasBlocked() bool {
	if o != nil && !IsNil(o.Blocked) {
		return true
	}

	return false
}

// SetBlocked gets a reference to the given bool and assigns it to the Blocked field.
func (o *UpdateDevice) SetBlocked(v bool) {
	o.Blocked = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *UpdateDevice) GetCredentials() []DeviceCredential {
	if o == nil || IsNil(o.Credentials) {
		var ret []DeviceCredential
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDevice) GetCredentialsOk() ([]DeviceCredential, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *UpdateDevice) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given []DeviceCredential and assigns it to the Credentials field.
func (o *UpdateDevice) SetCredentials(v []DeviceCredential) {
	o.Credentials = v
}

// GetLogLevel returns the LogLevel field value if set, zero value otherwise.
func (o *UpdateDevice) GetLogLevel() LogLevel {
	if o == nil || IsNil(o.LogLevel) {
		var ret LogLevel
		return ret
	}
	return *o.LogLevel
}

// GetLogLevelOk returns a tuple with the LogLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDevice) GetLogLevelOk() (*LogLevel, bool) {
	if o == nil || IsNil(o.LogLevel) {
		return nil, false
	}
	return o.LogLevel, true
}

// HasLogLevel returns a boolean if a field has been set.
func (o *UpdateDevice) HasLogLevel() bool {
	if o != nil && !IsNil(o.LogLevel) {
		return true
	}

	return false
}

// SetLogLevel gets a reference to the given LogLevel and assigns it to the LogLevel field.
func (o *UpdateDevice) SetLogLevel(v LogLevel) {
	o.LogLevel = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UpdateDevice) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDevice) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UpdateDevice) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *UpdateDevice) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

func (o UpdateDevice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Blocked) {
		toSerialize["blocked"] = o.Blocked
	}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.LogLevel) {
		toSerialize["logLevel"] = o.LogLevel
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableUpdateDevice struct {
	value *UpdateDevice
	isSet bool
}

func (v NullableUpdateDevice) Get() *UpdateDevice {
	return v.value
}

func (v *NullableUpdateDevice) Set(val *UpdateDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDevice(val *UpdateDevice) *NullableUpdateDevice {
	return &NullableUpdateDevice{value: val, isSet: true}
}

func (v NullableUpdateDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


