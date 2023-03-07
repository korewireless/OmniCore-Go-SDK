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

// checks if the ErrorStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorStatus{}

// ErrorStatus struct for ErrorStatus
type ErrorStatus struct {
	Code *int32 `json:"code,omitempty"`
	Details *string `json:"details,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewErrorStatus instantiates a new ErrorStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorStatus() *ErrorStatus {
	this := ErrorStatus{}
	return &this
}

// NewErrorStatusWithDefaults instantiates a new ErrorStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorStatusWithDefaults() *ErrorStatus {
	this := ErrorStatus{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ErrorStatus) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorStatus) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ErrorStatus) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *ErrorStatus) SetCode(v int32) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ErrorStatus) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorStatus) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ErrorStatus) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *ErrorStatus) SetDetails(v string) {
	o.Details = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ErrorStatus) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorStatus) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ErrorStatus) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ErrorStatus) SetMessage(v string) {
	o.Message = &v
}

func (o ErrorStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorStatus) ToMap() (map[string]interface{}, error) {
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

type NullableErrorStatus struct {
	value *ErrorStatus
	isSet bool
}

func (v NullableErrorStatus) Get() *ErrorStatus {
	return v.value
}

func (v *NullableErrorStatus) Set(val *ErrorStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorStatus(val *ErrorStatus) *NullableErrorStatus {
	return &NullableErrorStatus{value: val, isSet: true}
}

func (v NullableErrorStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


