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

// checks if the CreateConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateConfiguration{}

// CreateConfiguration struct for CreateConfiguration
type CreateConfiguration struct {
	Data *string `json:"data,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewCreateConfiguration instantiates a new CreateConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateConfiguration() *CreateConfiguration {
	this := CreateConfiguration{}
	return &this
}

// NewCreateConfigurationWithDefaults instantiates a new CreateConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateConfigurationWithDefaults() *CreateConfiguration {
	this := CreateConfiguration{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateConfiguration) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConfiguration) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateConfiguration) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *CreateConfiguration) SetData(v string) {
	o.Data = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateConfiguration) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConfiguration) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateConfiguration) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateConfiguration) SetType(v string) {
	o.Type = &v
}

func (o CreateConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableCreateConfiguration struct {
	value *CreateConfiguration
	isSet bool
}

func (v NullableCreateConfiguration) Get() *CreateConfiguration {
	return v.value
}

func (v *NullableCreateConfiguration) Set(val *CreateConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateConfiguration(val *CreateConfiguration) *NullableCreateConfiguration {
	return &NullableCreateConfiguration{value: val, isSet: true}
}

func (v NullableCreateConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


