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

// checks if the ErrorFrame type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorFrame{}

// ErrorFrame struct for ErrorFrame
type ErrorFrame struct {
	Code int32 `json:"code"`
	Details string `json:"details"`
	Message string `json:"message"`
	Status string `json:"status"`
}

// NewErrorFrame instantiates a new ErrorFrame object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorFrame(code int32, details string, message string, status string) *ErrorFrame {
	this := ErrorFrame{}
	this.Code = code
	this.Details = details
	this.Message = message
	this.Status = status
	return &this
}

// NewErrorFrameWithDefaults instantiates a new ErrorFrame object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorFrameWithDefaults() *ErrorFrame {
	this := ErrorFrame{}
	return &this
}

// GetCode returns the Code field value
func (o *ErrorFrame) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ErrorFrame) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ErrorFrame) SetCode(v int32) {
	o.Code = v
}

// GetDetails returns the Details field value
func (o *ErrorFrame) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *ErrorFrame) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *ErrorFrame) SetDetails(v string) {
	o.Details = v
}

// GetMessage returns the Message field value
func (o *ErrorFrame) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ErrorFrame) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ErrorFrame) SetMessage(v string) {
	o.Message = v
}

// GetStatus returns the Status field value
func (o *ErrorFrame) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ErrorFrame) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ErrorFrame) SetStatus(v string) {
	o.Status = v
}

func (o ErrorFrame) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorFrame) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["details"] = o.Details
	toSerialize["message"] = o.Message
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableErrorFrame struct {
	value *ErrorFrame
	isSet bool
}

func (v NullableErrorFrame) Get() *ErrorFrame {
	return v.value
}

func (v *NullableErrorFrame) Set(val *ErrorFrame) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorFrame) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorFrame) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorFrame(val *ErrorFrame) *NullableErrorFrame {
	return &NullableErrorFrame{value: val, isSet: true}
}

func (v NullableErrorFrame) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorFrame) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


