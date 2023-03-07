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

// checks if the RegistryCredential type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistryCredential{}

// RegistryCredential struct for RegistryCredential
type RegistryCredential struct {
	PublicKeyCertificate *PublicKeyCertificate `json:"publicKeyCertificate,omitempty"`
}

// NewRegistryCredential instantiates a new RegistryCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistryCredential() *RegistryCredential {
	this := RegistryCredential{}
	return &this
}

// NewRegistryCredentialWithDefaults instantiates a new RegistryCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistryCredentialWithDefaults() *RegistryCredential {
	this := RegistryCredential{}
	return &this
}

// GetPublicKeyCertificate returns the PublicKeyCertificate field value if set, zero value otherwise.
func (o *RegistryCredential) GetPublicKeyCertificate() PublicKeyCertificate {
	if o == nil || IsNil(o.PublicKeyCertificate) {
		var ret PublicKeyCertificate
		return ret
	}
	return *o.PublicKeyCertificate
}

// GetPublicKeyCertificateOk returns a tuple with the PublicKeyCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryCredential) GetPublicKeyCertificateOk() (*PublicKeyCertificate, bool) {
	if o == nil || IsNil(o.PublicKeyCertificate) {
		return nil, false
	}
	return o.PublicKeyCertificate, true
}

// HasPublicKeyCertificate returns a boolean if a field has been set.
func (o *RegistryCredential) HasPublicKeyCertificate() bool {
	if o != nil && !IsNil(o.PublicKeyCertificate) {
		return true
	}

	return false
}

// SetPublicKeyCertificate gets a reference to the given PublicKeyCertificate and assigns it to the PublicKeyCertificate field.
func (o *RegistryCredential) SetPublicKeyCertificate(v PublicKeyCertificate) {
	o.PublicKeyCertificate = &v
}

func (o RegistryCredential) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistryCredential) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PublicKeyCertificate) {
		toSerialize["publicKeyCertificate"] = o.PublicKeyCertificate
	}
	return toSerialize, nil
}

type NullableRegistryCredential struct {
	value *RegistryCredential
	isSet bool
}

func (v NullableRegistryCredential) Get() *RegistryCredential {
	return v.value
}

func (v *NullableRegistryCredential) Set(val *RegistryCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistryCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistryCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistryCredential(val *RegistryCredential) *NullableRegistryCredential {
	return &NullableRegistryCredential{value: val, isSet: true}
}

func (v NullableRegistryCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistryCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


