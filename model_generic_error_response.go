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

// checks if the GenericErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenericErrorResponse{}

// GenericErrorResponse struct for GenericErrorResponse
type GenericErrorResponse struct {
	Error ErrorFrame `json:"error"`
}

// NewGenericErrorResponse instantiates a new GenericErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericErrorResponse(error_ ErrorFrame) *GenericErrorResponse {
	this := GenericErrorResponse{}
	this.Error = error_
	return &this
}

// NewGenericErrorResponseWithDefaults instantiates a new GenericErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericErrorResponseWithDefaults() *GenericErrorResponse {
	this := GenericErrorResponse{}
	return &this
}

// GetError returns the Error field value
func (o *GenericErrorResponse) GetError() ErrorFrame {
	if o == nil {
		var ret ErrorFrame
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *GenericErrorResponse) GetErrorOk() (*ErrorFrame, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *GenericErrorResponse) SetError(v ErrorFrame) {
	o.Error = v
}

func (o GenericErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenericErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

type NullableGenericErrorResponse struct {
	value *GenericErrorResponse
	isSet bool
}

func (v NullableGenericErrorResponse) Get() *GenericErrorResponse {
	return v.value
}

func (v *NullableGenericErrorResponse) Set(val *GenericErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericErrorResponse(val *GenericErrorResponse) *NullableGenericErrorResponse {
	return &NullableGenericErrorResponse{value: val, isSet: true}
}

func (v NullableGenericErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


