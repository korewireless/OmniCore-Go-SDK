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

// checks if the OmnicoreBlockCommunicationBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OmnicoreBlockCommunicationBody{}

// OmnicoreBlockCommunicationBody struct for OmnicoreBlockCommunicationBody
type OmnicoreBlockCommunicationBody struct {
	Isblocked *bool `json:"isblocked,omitempty"`
}

// NewOmnicoreBlockCommunicationBody instantiates a new OmnicoreBlockCommunicationBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOmnicoreBlockCommunicationBody() *OmnicoreBlockCommunicationBody {
	this := OmnicoreBlockCommunicationBody{}
	return &this
}

// NewOmnicoreBlockCommunicationBodyWithDefaults instantiates a new OmnicoreBlockCommunicationBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOmnicoreBlockCommunicationBodyWithDefaults() *OmnicoreBlockCommunicationBody {
	this := OmnicoreBlockCommunicationBody{}
	return &this
}

// GetIsblocked returns the Isblocked field value if set, zero value otherwise.
func (o *OmnicoreBlockCommunicationBody) GetIsblocked() bool {
	if o == nil || IsNil(o.Isblocked) {
		var ret bool
		return ret
	}
	return *o.Isblocked
}

// GetIsblockedOk returns a tuple with the Isblocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreBlockCommunicationBody) GetIsblockedOk() (*bool, bool) {
	if o == nil || IsNil(o.Isblocked) {
		return nil, false
	}
	return o.Isblocked, true
}

// HasIsblocked returns a boolean if a field has been set.
func (o *OmnicoreBlockCommunicationBody) HasIsblocked() bool {
	if o != nil && !IsNil(o.Isblocked) {
		return true
	}

	return false
}

// SetIsblocked gets a reference to the given bool and assigns it to the Isblocked field.
func (o *OmnicoreBlockCommunicationBody) SetIsblocked(v bool) {
	o.Isblocked = &v
}

func (o OmnicoreBlockCommunicationBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OmnicoreBlockCommunicationBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Isblocked) {
		toSerialize["isblocked"] = o.Isblocked
	}
	return toSerialize, nil
}

type NullableOmnicoreBlockCommunicationBody struct {
	value *OmnicoreBlockCommunicationBody
	isSet bool
}

func (v NullableOmnicoreBlockCommunicationBody) Get() *OmnicoreBlockCommunicationBody {
	return v.value
}

func (v *NullableOmnicoreBlockCommunicationBody) Set(val *OmnicoreBlockCommunicationBody) {
	v.value = val
	v.isSet = true
}

func (v NullableOmnicoreBlockCommunicationBody) IsSet() bool {
	return v.isSet
}

func (v *NullableOmnicoreBlockCommunicationBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOmnicoreBlockCommunicationBody(val *OmnicoreBlockCommunicationBody) *NullableOmnicoreBlockCommunicationBody {
	return &NullableOmnicoreBlockCommunicationBody{value: val, isSet: true}
}

func (v NullableOmnicoreBlockCommunicationBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOmnicoreBlockCommunicationBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


