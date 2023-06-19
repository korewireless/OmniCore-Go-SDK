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

// checks if the ListSinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSinks{}

// ListSinks struct for ListSinks
type ListSinks struct {
	Sinks []Sink `json:"sinks,omitempty"`
}

// NewListSinks instantiates a new ListSinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSinks() *ListSinks {
	this := ListSinks{}
	return &this
}

// NewListSinksWithDefaults instantiates a new ListSinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSinksWithDefaults() *ListSinks {
	this := ListSinks{}
	return &this
}

// GetSinks returns the Sinks field value if set, zero value otherwise.
func (o *ListSinks) GetSinks() []Sink {
	if o == nil || IsNil(o.Sinks) {
		var ret []Sink
		return ret
	}
	return o.Sinks
}

// GetSinksOk returns a tuple with the Sinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSinks) GetSinksOk() ([]Sink, bool) {
	if o == nil || IsNil(o.Sinks) {
		return nil, false
	}
	return o.Sinks, true
}

// HasSinks returns a boolean if a field has been set.
func (o *ListSinks) HasSinks() bool {
	if o != nil && !IsNil(o.Sinks) {
		return true
	}

	return false
}

// SetSinks gets a reference to the given []Sink and assigns it to the Sinks field.
func (o *ListSinks) SetSinks(v []Sink) {
	o.Sinks = v
}

func (o ListSinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sinks) {
		toSerialize["sinks"] = o.Sinks
	}
	return toSerialize, nil
}

type NullableListSinks struct {
	value *ListSinks
	isSet bool
}

func (v NullableListSinks) Get() *ListSinks {
	return v.value
}

func (v *NullableListSinks) Set(val *ListSinks) {
	v.value = val
	v.isSet = true
}

func (v NullableListSinks) IsSet() bool {
	return v.isSet
}

func (v *NullableListSinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSinks(val *ListSinks) *NullableListSinks {
	return &NullableListSinks{value: val, isSet: true}
}

func (v NullableListSinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


