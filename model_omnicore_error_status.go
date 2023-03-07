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

// checks if the OmnicoreErrorStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OmnicoreErrorStatus{}

// OmnicoreErrorStatus struct for OmnicoreErrorStatus
type OmnicoreErrorStatus struct {
	Code *int32 `json:"code,omitempty"`
	Details *string `json:"details,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewOmnicoreErrorStatus instantiates a new OmnicoreErrorStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOmnicoreErrorStatus() *OmnicoreErrorStatus {
	this := OmnicoreErrorStatus{}
	return &this
}

// NewOmnicoreErrorStatusWithDefaults instantiates a new OmnicoreErrorStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOmnicoreErrorStatusWithDefaults() *OmnicoreErrorStatus {
	this := OmnicoreErrorStatus{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OmnicoreErrorStatus) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreErrorStatus) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OmnicoreErrorStatus) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *OmnicoreErrorStatus) SetCode(v int32) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *OmnicoreErrorStatus) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreErrorStatus) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *OmnicoreErrorStatus) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *OmnicoreErrorStatus) SetDetails(v string) {
	o.Details = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *OmnicoreErrorStatus) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreErrorStatus) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *OmnicoreErrorStatus) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *OmnicoreErrorStatus) SetMessage(v string) {
	o.Message = &v
}

func (o OmnicoreErrorStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OmnicoreErrorStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableOmnicoreErrorStatus struct {
	value *OmnicoreErrorStatus
	isSet bool
}

func (v NullableOmnicoreErrorStatus) Get() *OmnicoreErrorStatus {
	return v.value
}

func (v *NullableOmnicoreErrorStatus) Set(val *OmnicoreErrorStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableOmnicoreErrorStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableOmnicoreErrorStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOmnicoreErrorStatus(val *OmnicoreErrorStatus) *NullableOmnicoreErrorStatus {
	return &NullableOmnicoreErrorStatus{value: val, isSet: true}
}

func (v NullableOmnicoreErrorStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOmnicoreErrorStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


