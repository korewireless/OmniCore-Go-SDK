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

// checks if the OmnicoreRegistryCredential type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OmnicoreRegistryCredential{}

// OmnicoreRegistryCredential struct for OmnicoreRegistryCredential
type OmnicoreRegistryCredential struct {
	Id *string `json:"id,omitempty"`
	PublicKeyCertificate *OmnicorePublicKeyCertificate `json:"publicKeyCertificate,omitempty"`
}

// NewOmnicoreRegistryCredential instantiates a new OmnicoreRegistryCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOmnicoreRegistryCredential() *OmnicoreRegistryCredential {
	this := OmnicoreRegistryCredential{}
	return &this
}

// NewOmnicoreRegistryCredentialWithDefaults instantiates a new OmnicoreRegistryCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOmnicoreRegistryCredentialWithDefaults() *OmnicoreRegistryCredential {
	this := OmnicoreRegistryCredential{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OmnicoreRegistryCredential) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreRegistryCredential) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OmnicoreRegistryCredential) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OmnicoreRegistryCredential) SetId(v string) {
	o.Id = &v
}

// GetPublicKeyCertificate returns the PublicKeyCertificate field value if set, zero value otherwise.
func (o *OmnicoreRegistryCredential) GetPublicKeyCertificate() OmnicorePublicKeyCertificate {
	if o == nil || IsNil(o.PublicKeyCertificate) {
		var ret OmnicorePublicKeyCertificate
		return ret
	}
	return *o.PublicKeyCertificate
}

// GetPublicKeyCertificateOk returns a tuple with the PublicKeyCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreRegistryCredential) GetPublicKeyCertificateOk() (*OmnicorePublicKeyCertificate, bool) {
	if o == nil || IsNil(o.PublicKeyCertificate) {
		return nil, false
	}
	return o.PublicKeyCertificate, true
}

// HasPublicKeyCertificate returns a boolean if a field has been set.
func (o *OmnicoreRegistryCredential) HasPublicKeyCertificate() bool {
	if o != nil && !IsNil(o.PublicKeyCertificate) {
		return true
	}

	return false
}

// SetPublicKeyCertificate gets a reference to the given OmnicorePublicKeyCertificate and assigns it to the PublicKeyCertificate field.
func (o *OmnicoreRegistryCredential) SetPublicKeyCertificate(v OmnicorePublicKeyCertificate) {
	o.PublicKeyCertificate = &v
}

func (o OmnicoreRegistryCredential) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OmnicoreRegistryCredential) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.PublicKeyCertificate) {
		toSerialize["publicKeyCertificate"] = o.PublicKeyCertificate
	}
	return toSerialize, nil
}

type NullableOmnicoreRegistryCredential struct {
	value *OmnicoreRegistryCredential
	isSet bool
}

func (v NullableOmnicoreRegistryCredential) Get() *OmnicoreRegistryCredential {
	return v.value
}

func (v *NullableOmnicoreRegistryCredential) Set(val *OmnicoreRegistryCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableOmnicoreRegistryCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableOmnicoreRegistryCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOmnicoreRegistryCredential(val *OmnicoreRegistryCredential) *NullableOmnicoreRegistryCredential {
	return &NullableOmnicoreRegistryCredential{value: val, isSet: true}
}

func (v NullableOmnicoreRegistryCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOmnicoreRegistryCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


