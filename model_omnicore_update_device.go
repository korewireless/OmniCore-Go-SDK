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

// checks if the OmnicoreUpdateDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OmnicoreUpdateDevice{}

// OmnicoreUpdateDevice struct for OmnicoreUpdateDevice
type OmnicoreUpdateDevice struct {
	Blocked *bool `json:"blocked,omitempty"`
	Credentials []OmnicoreDeviceCredential `json:"credentials,omitempty"`
	LogLevel *string `json:"logLevel,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
}

// NewOmnicoreUpdateDevice instantiates a new OmnicoreUpdateDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOmnicoreUpdateDevice() *OmnicoreUpdateDevice {
	this := OmnicoreUpdateDevice{}
	return &this
}

// NewOmnicoreUpdateDeviceWithDefaults instantiates a new OmnicoreUpdateDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOmnicoreUpdateDeviceWithDefaults() *OmnicoreUpdateDevice {
	this := OmnicoreUpdateDevice{}
	return &this
}

// GetBlocked returns the Blocked field value if set, zero value otherwise.
func (o *OmnicoreUpdateDevice) GetBlocked() bool {
	if o == nil || IsNil(o.Blocked) {
		var ret bool
		return ret
	}
	return *o.Blocked
}

// GetBlockedOk returns a tuple with the Blocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreUpdateDevice) GetBlockedOk() (*bool, bool) {
	if o == nil || IsNil(o.Blocked) {
		return nil, false
	}
	return o.Blocked, true
}

// HasBlocked returns a boolean if a field has been set.
func (o *OmnicoreUpdateDevice) HasBlocked() bool {
	if o != nil && !IsNil(o.Blocked) {
		return true
	}

	return false
}

// SetBlocked gets a reference to the given bool and assigns it to the Blocked field.
func (o *OmnicoreUpdateDevice) SetBlocked(v bool) {
	o.Blocked = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *OmnicoreUpdateDevice) GetCredentials() []OmnicoreDeviceCredential {
	if o == nil || IsNil(o.Credentials) {
		var ret []OmnicoreDeviceCredential
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreUpdateDevice) GetCredentialsOk() ([]OmnicoreDeviceCredential, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *OmnicoreUpdateDevice) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given []OmnicoreDeviceCredential and assigns it to the Credentials field.
func (o *OmnicoreUpdateDevice) SetCredentials(v []OmnicoreDeviceCredential) {
	o.Credentials = v
}

// GetLogLevel returns the LogLevel field value if set, zero value otherwise.
func (o *OmnicoreUpdateDevice) GetLogLevel() string {
	if o == nil || IsNil(o.LogLevel) {
		var ret string
		return ret
	}
	return *o.LogLevel
}

// GetLogLevelOk returns a tuple with the LogLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreUpdateDevice) GetLogLevelOk() (*string, bool) {
	if o == nil || IsNil(o.LogLevel) {
		return nil, false
	}
	return o.LogLevel, true
}

// HasLogLevel returns a boolean if a field has been set.
func (o *OmnicoreUpdateDevice) HasLogLevel() bool {
	if o != nil && !IsNil(o.LogLevel) {
		return true
	}

	return false
}

// SetLogLevel gets a reference to the given string and assigns it to the LogLevel field.
func (o *OmnicoreUpdateDevice) SetLogLevel(v string) {
	o.LogLevel = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *OmnicoreUpdateDevice) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreUpdateDevice) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *OmnicoreUpdateDevice) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *OmnicoreUpdateDevice) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

func (o OmnicoreUpdateDevice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OmnicoreUpdateDevice) ToMap() (map[string]interface{}, error) {
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

type NullableOmnicoreUpdateDevice struct {
	value *OmnicoreUpdateDevice
	isSet bool
}

func (v NullableOmnicoreUpdateDevice) Get() *OmnicoreUpdateDevice {
	return v.value
}

func (v *NullableOmnicoreUpdateDevice) Set(val *OmnicoreUpdateDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableOmnicoreUpdateDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableOmnicoreUpdateDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOmnicoreUpdateDevice(val *OmnicoreUpdateDevice) *NullableOmnicoreUpdateDevice {
	return &NullableOmnicoreUpdateDevice{value: val, isSet: true}
}

func (v NullableOmnicoreUpdateDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOmnicoreUpdateDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


