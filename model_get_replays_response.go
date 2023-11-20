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

// checks if the GetReplaysResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetReplaysResponse{}

// GetReplaysResponse struct for GetReplaysResponse
type GetReplaysResponse struct {
	Details []Replay `json:"details,omitempty"`
}

// NewGetReplaysResponse instantiates a new GetReplaysResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetReplaysResponse() *GetReplaysResponse {
	this := GetReplaysResponse{}
	return &this
}

// NewGetReplaysResponseWithDefaults instantiates a new GetReplaysResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetReplaysResponseWithDefaults() *GetReplaysResponse {
	this := GetReplaysResponse{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *GetReplaysResponse) GetDetails() []Replay {
	if o == nil || IsNil(o.Details) {
		var ret []Replay
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReplaysResponse) GetDetailsOk() ([]Replay, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *GetReplaysResponse) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []Replay and assigns it to the Details field.
func (o *GetReplaysResponse) SetDetails(v []Replay) {
	o.Details = v
}

func (o GetReplaysResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetReplaysResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableGetReplaysResponse struct {
	value *GetReplaysResponse
	isSet bool
}

func (v NullableGetReplaysResponse) Get() *GetReplaysResponse {
	return v.value
}

func (v *NullableGetReplaysResponse) Set(val *GetReplaysResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetReplaysResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetReplaysResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetReplaysResponse(val *GetReplaysResponse) *NullableGetReplaysResponse {
	return &NullableGetReplaysResponse{value: val, isSet: true}
}

func (v NullableGetReplaysResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetReplaysResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


