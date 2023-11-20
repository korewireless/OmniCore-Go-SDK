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

// checks if the Frame type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Frame{}

// Frame struct for Frame
type Frame struct {
	Details *string `json:"details,omitempty"`
}

// NewFrame instantiates a new Frame object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrame() *Frame {
	this := Frame{}
	return &this
}

// NewFrameWithDefaults instantiates a new Frame object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrameWithDefaults() *Frame {
	this := Frame{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *Frame) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Frame) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *Frame) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *Frame) SetDetails(v string) {
	o.Details = &v
}

func (o Frame) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Frame) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableFrame struct {
	value *Frame
	isSet bool
}

func (v NullableFrame) Get() *Frame {
	return v.value
}

func (v *NullableFrame) Set(val *Frame) {
	v.value = val
	v.isSet = true
}

func (v NullableFrame) IsSet() bool {
	return v.isSet
}

func (v *NullableFrame) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrame(val *Frame) *NullableFrame {
	return &NullableFrame{value: val, isSet: true}
}

func (v NullableFrame) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrame) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


