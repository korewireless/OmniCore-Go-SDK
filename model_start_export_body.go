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

// checks if the StartExportBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartExportBody{}

// StartExportBody struct for StartExportBody
type StartExportBody struct {
	Source *string `json:"source,omitempty"`
	Destination *string `json:"destination,omitempty"`
}

// NewStartExportBody instantiates a new StartExportBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartExportBody() *StartExportBody {
	this := StartExportBody{}
	return &this
}

// NewStartExportBodyWithDefaults instantiates a new StartExportBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartExportBodyWithDefaults() *StartExportBody {
	this := StartExportBody{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *StartExportBody) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartExportBody) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *StartExportBody) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *StartExportBody) SetSource(v string) {
	o.Source = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *StartExportBody) GetDestination() string {
	if o == nil || IsNil(o.Destination) {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartExportBody) GetDestinationOk() (*string, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *StartExportBody) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *StartExportBody) SetDestination(v string) {
	o.Destination = &v
}

func (o StartExportBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartExportBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	return toSerialize, nil
}

type NullableStartExportBody struct {
	value *StartExportBody
	isSet bool
}

func (v NullableStartExportBody) Get() *StartExportBody {
	return v.value
}

func (v *NullableStartExportBody) Set(val *StartExportBody) {
	v.value = val
	v.isSet = true
}

func (v NullableStartExportBody) IsSet() bool {
	return v.isSet
}

func (v *NullableStartExportBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartExportBody(val *StartExportBody) *NullableStartExportBody {
	return &NullableStartExportBody{value: val, isSet: true}
}

func (v NullableStartExportBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartExportBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


