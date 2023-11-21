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

// checks if the CreateVaultKeyBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateVaultKeyBody{}

// CreateVaultKeyBody struct for CreateVaultKeyBody
type CreateVaultKeyBody struct {
	Name *string `json:"name,omitempty"`
	Key *string `json:"key,omitempty"`
}

// NewCreateVaultKeyBody instantiates a new CreateVaultKeyBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVaultKeyBody() *CreateVaultKeyBody {
	this := CreateVaultKeyBody{}
	return &this
}

// NewCreateVaultKeyBodyWithDefaults instantiates a new CreateVaultKeyBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVaultKeyBodyWithDefaults() *CreateVaultKeyBody {
	this := CreateVaultKeyBody{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateVaultKeyBody) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVaultKeyBody) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateVaultKeyBody) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateVaultKeyBody) SetName(v string) {
	o.Name = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *CreateVaultKeyBody) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVaultKeyBody) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *CreateVaultKeyBody) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *CreateVaultKeyBody) SetKey(v string) {
	o.Key = &v
}

func (o CreateVaultKeyBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateVaultKeyBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	return toSerialize, nil
}

type NullableCreateVaultKeyBody struct {
	value *CreateVaultKeyBody
	isSet bool
}

func (v NullableCreateVaultKeyBody) Get() *CreateVaultKeyBody {
	return v.value
}

func (v *NullableCreateVaultKeyBody) Set(val *CreateVaultKeyBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVaultKeyBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVaultKeyBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVaultKeyBody(val *CreateVaultKeyBody) *NullableCreateVaultKeyBody {
	return &NullableCreateVaultKeyBody{value: val, isSet: true}
}

func (v NullableCreateVaultKeyBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVaultKeyBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


