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

// checks if the RegistryCertificate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistryCertificate{}

// RegistryCertificate struct for RegistryCertificate
type RegistryCertificate struct {
	Credentials RegistryCredential `json:"credentials"`
}

// NewRegistryCertificate instantiates a new RegistryCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistryCertificate(credentials RegistryCredential) *RegistryCertificate {
	this := RegistryCertificate{}
	this.Credentials = credentials
	return &this
}

// NewRegistryCertificateWithDefaults instantiates a new RegistryCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistryCertificateWithDefaults() *RegistryCertificate {
	this := RegistryCertificate{}
	return &this
}

// GetCredentials returns the Credentials field value
func (o *RegistryCertificate) GetCredentials() RegistryCredential {
	if o == nil {
		var ret RegistryCredential
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *RegistryCertificate) GetCredentialsOk() (*RegistryCredential, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *RegistryCertificate) SetCredentials(v RegistryCredential) {
	o.Credentials = v
}

func (o RegistryCertificate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistryCertificate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["credentials"] = o.Credentials
	return toSerialize, nil
}

type NullableRegistryCertificate struct {
	value *RegistryCertificate
	isSet bool
}

func (v NullableRegistryCertificate) Get() *RegistryCertificate {
	return v.value
}

func (v *NullableRegistryCertificate) Set(val *RegistryCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistryCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistryCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistryCertificate(val *RegistryCertificate) *NullableRegistryCertificate {
	return &NullableRegistryCertificate{value: val, isSet: true}
}

func (v NullableRegistryCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistryCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


