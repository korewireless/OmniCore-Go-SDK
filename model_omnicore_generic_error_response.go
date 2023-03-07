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

// checks if the OmnicoreGenericErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OmnicoreGenericErrorResponse{}

// OmnicoreGenericErrorResponse struct for OmnicoreGenericErrorResponse
type OmnicoreGenericErrorResponse struct {
	Error OmnicoreErrorFrame `json:"error"`
}

// NewOmnicoreGenericErrorResponse instantiates a new OmnicoreGenericErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOmnicoreGenericErrorResponse(error_ OmnicoreErrorFrame) *OmnicoreGenericErrorResponse {
	this := OmnicoreGenericErrorResponse{}
	this.Error = error_
	return &this
}

// NewOmnicoreGenericErrorResponseWithDefaults instantiates a new OmnicoreGenericErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOmnicoreGenericErrorResponseWithDefaults() *OmnicoreGenericErrorResponse {
	this := OmnicoreGenericErrorResponse{}
	return &this
}

// GetError returns the Error field value
func (o *OmnicoreGenericErrorResponse) GetError() OmnicoreErrorFrame {
	if o == nil {
		var ret OmnicoreErrorFrame
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *OmnicoreGenericErrorResponse) GetErrorOk() (*OmnicoreErrorFrame, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *OmnicoreGenericErrorResponse) SetError(v OmnicoreErrorFrame) {
	o.Error = v
}

func (o OmnicoreGenericErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OmnicoreGenericErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

type NullableOmnicoreGenericErrorResponse struct {
	value *OmnicoreGenericErrorResponse
	isSet bool
}

func (v NullableOmnicoreGenericErrorResponse) Get() *OmnicoreGenericErrorResponse {
	return v.value
}

func (v *NullableOmnicoreGenericErrorResponse) Set(val *OmnicoreGenericErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOmnicoreGenericErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOmnicoreGenericErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOmnicoreGenericErrorResponse(val *OmnicoreGenericErrorResponse) *NullableOmnicoreGenericErrorResponse {
	return &NullableOmnicoreGenericErrorResponse{value: val, isSet: true}
}

func (v NullableOmnicoreGenericErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOmnicoreGenericErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


