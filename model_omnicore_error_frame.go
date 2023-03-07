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

// checks if the OmnicoreErrorFrame type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OmnicoreErrorFrame{}

// OmnicoreErrorFrame struct for OmnicoreErrorFrame
type OmnicoreErrorFrame struct {
	Code int32 `json:"code"`
	Details string `json:"details"`
	Message string `json:"message"`
	Status string `json:"status"`
}

// NewOmnicoreErrorFrame instantiates a new OmnicoreErrorFrame object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOmnicoreErrorFrame(code int32, details string, message string, status string) *OmnicoreErrorFrame {
	this := OmnicoreErrorFrame{}
	this.Code = code
	this.Details = details
	this.Message = message
	this.Status = status
	return &this
}

// NewOmnicoreErrorFrameWithDefaults instantiates a new OmnicoreErrorFrame object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOmnicoreErrorFrameWithDefaults() *OmnicoreErrorFrame {
	this := OmnicoreErrorFrame{}
	return &this
}

// GetCode returns the Code field value
func (o *OmnicoreErrorFrame) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *OmnicoreErrorFrame) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *OmnicoreErrorFrame) SetCode(v int32) {
	o.Code = v
}

// GetDetails returns the Details field value
func (o *OmnicoreErrorFrame) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *OmnicoreErrorFrame) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *OmnicoreErrorFrame) SetDetails(v string) {
	o.Details = v
}

// GetMessage returns the Message field value
func (o *OmnicoreErrorFrame) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *OmnicoreErrorFrame) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *OmnicoreErrorFrame) SetMessage(v string) {
	o.Message = v
}

// GetStatus returns the Status field value
func (o *OmnicoreErrorFrame) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OmnicoreErrorFrame) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *OmnicoreErrorFrame) SetStatus(v string) {
	o.Status = v
}

func (o OmnicoreErrorFrame) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OmnicoreErrorFrame) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["details"] = o.Details
	toSerialize["message"] = o.Message
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableOmnicoreErrorFrame struct {
	value *OmnicoreErrorFrame
	isSet bool
}

func (v NullableOmnicoreErrorFrame) Get() *OmnicoreErrorFrame {
	return v.value
}

func (v *NullableOmnicoreErrorFrame) Set(val *OmnicoreErrorFrame) {
	v.value = val
	v.isSet = true
}

func (v NullableOmnicoreErrorFrame) IsSet() bool {
	return v.isSet
}

func (v *NullableOmnicoreErrorFrame) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOmnicoreErrorFrame(val *OmnicoreErrorFrame) *NullableOmnicoreErrorFrame {
	return &NullableOmnicoreErrorFrame{value: val, isSet: true}
}

func (v NullableOmnicoreErrorFrame) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOmnicoreErrorFrame) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


