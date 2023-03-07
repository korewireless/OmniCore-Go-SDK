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

// checks if the DeviceCertificate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceCertificate{}

// DeviceCertificate struct for DeviceCertificate
type DeviceCertificate struct {
	Credentials DeviceCredential `json:"credentials"`
}

// NewDeviceCertificate instantiates a new DeviceCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceCertificate(credentials DeviceCredential) *DeviceCertificate {
	this := DeviceCertificate{}
	this.Credentials = credentials
	return &this
}

// NewDeviceCertificateWithDefaults instantiates a new DeviceCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceCertificateWithDefaults() *DeviceCertificate {
	this := DeviceCertificate{}
	return &this
}

// GetCredentials returns the Credentials field value
func (o *DeviceCertificate) GetCredentials() DeviceCredential {
	if o == nil {
		var ret DeviceCredential
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *DeviceCertificate) GetCredentialsOk() (*DeviceCredential, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *DeviceCertificate) SetCredentials(v DeviceCredential) {
	o.Credentials = v
}

func (o DeviceCertificate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceCertificate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["credentials"] = o.Credentials
	return toSerialize, nil
}

type NullableDeviceCertificate struct {
	value *DeviceCertificate
	isSet bool
}

func (v NullableDeviceCertificate) Get() *DeviceCertificate {
	return v.value
}

func (v *NullableDeviceCertificate) Set(val *DeviceCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceCertificate(val *DeviceCertificate) *NullableDeviceCertificate {
	return &NullableDeviceCertificate{value: val, isSet: true}
}

func (v NullableDeviceCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


