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

// checks if the DeviceCredential type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceCredential{}

// DeviceCredential struct for DeviceCredential
type DeviceCredential struct {
	// ExpirationTime: [Optional] The time at which this credential becomes invalid. This credential will be ignored for new client authentication requests after this timestamp; however, it will not be automatically deleted.
	ExpirationTime *string `json:"expirationTime,omitempty"`
	PublicKey *PublicKeyCredential `json:"publicKey,omitempty"`
}

// NewDeviceCredential instantiates a new DeviceCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceCredential() *DeviceCredential {
	this := DeviceCredential{}
	return &this
}

// NewDeviceCredentialWithDefaults instantiates a new DeviceCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceCredentialWithDefaults() *DeviceCredential {
	this := DeviceCredential{}
	return &this
}

// GetExpirationTime returns the ExpirationTime field value if set, zero value otherwise.
func (o *DeviceCredential) GetExpirationTime() string {
	if o == nil || IsNil(o.ExpirationTime) {
		var ret string
		return ret
	}
	return *o.ExpirationTime
}

// GetExpirationTimeOk returns a tuple with the ExpirationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCredential) GetExpirationTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationTime) {
		return nil, false
	}
	return o.ExpirationTime, true
}

// HasExpirationTime returns a boolean if a field has been set.
func (o *DeviceCredential) HasExpirationTime() bool {
	if o != nil && !IsNil(o.ExpirationTime) {
		return true
	}

	return false
}

// SetExpirationTime gets a reference to the given string and assigns it to the ExpirationTime field.
func (o *DeviceCredential) SetExpirationTime(v string) {
	o.ExpirationTime = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *DeviceCredential) GetPublicKey() PublicKeyCredential {
	if o == nil || IsNil(o.PublicKey) {
		var ret PublicKeyCredential
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCredential) GetPublicKeyOk() (*PublicKeyCredential, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *DeviceCredential) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given PublicKeyCredential and assigns it to the PublicKey field.
func (o *DeviceCredential) SetPublicKey(v PublicKeyCredential) {
	o.PublicKey = &v
}

func (o DeviceCredential) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceCredential) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpirationTime) {
		toSerialize["expirationTime"] = o.ExpirationTime
	}
	if !IsNil(o.PublicKey) {
		toSerialize["publicKey"] = o.PublicKey
	}
	return toSerialize, nil
}

type NullableDeviceCredential struct {
	value *DeviceCredential
	isSet bool
}

func (v NullableDeviceCredential) Get() *DeviceCredential {
	return v.value
}

func (v *NullableDeviceCredential) Set(val *DeviceCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceCredential(val *DeviceCredential) *NullableDeviceCredential {
	return &NullableDeviceCredential{value: val, isSet: true}
}

func (v NullableDeviceCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


