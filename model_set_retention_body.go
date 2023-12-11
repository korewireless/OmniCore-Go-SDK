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

// checks if the SetRetentionBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetRetentionBody{}

// SetRetentionBody struct for SetRetentionBody
type SetRetentionBody struct {
	RetentionPeriod *string `json:"retentionPeriod,omitempty"`
}

// NewSetRetentionBody instantiates a new SetRetentionBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetRetentionBody() *SetRetentionBody {
	this := SetRetentionBody{}
	return &this
}

// NewSetRetentionBodyWithDefaults instantiates a new SetRetentionBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetRetentionBodyWithDefaults() *SetRetentionBody {
	this := SetRetentionBody{}
	return &this
}

// GetRetentionPeriod returns the RetentionPeriod field value if set, zero value otherwise.
func (o *SetRetentionBody) GetRetentionPeriod() string {
	if o == nil || IsNil(o.RetentionPeriod) {
		var ret string
		return ret
	}
	return *o.RetentionPeriod
}

// GetRetentionPeriodOk returns a tuple with the RetentionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetRetentionBody) GetRetentionPeriodOk() (*string, bool) {
	if o == nil || IsNil(o.RetentionPeriod) {
		return nil, false
	}
	return o.RetentionPeriod, true
}

// HasRetentionPeriod returns a boolean if a field has been set.
func (o *SetRetentionBody) HasRetentionPeriod() bool {
	if o != nil && !IsNil(o.RetentionPeriod) {
		return true
	}

	return false
}

// SetRetentionPeriod gets a reference to the given string and assigns it to the RetentionPeriod field.
func (o *SetRetentionBody) SetRetentionPeriod(v string) {
	o.RetentionPeriod = &v
}

func (o SetRetentionBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetRetentionBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RetentionPeriod) {
		toSerialize["retentionPeriod"] = o.RetentionPeriod
	}
	return toSerialize, nil
}

type NullableSetRetentionBody struct {
	value *SetRetentionBody
	isSet bool
}

func (v NullableSetRetentionBody) Get() *SetRetentionBody {
	return v.value
}

func (v *NullableSetRetentionBody) Set(val *SetRetentionBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSetRetentionBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSetRetentionBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetRetentionBody(val *SetRetentionBody) *NullableSetRetentionBody {
	return &NullableSetRetentionBody{value: val, isSet: true}
}

func (v NullableSetRetentionBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetRetentionBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


