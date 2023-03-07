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

// checks if the OmnicoreListDeviceStatesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OmnicoreListDeviceStatesResponse{}

// OmnicoreListDeviceStatesResponse struct for OmnicoreListDeviceStatesResponse
type OmnicoreListDeviceStatesResponse struct {
	DeviceStates []OmnicoreDeviceState `json:"deviceStates,omitempty"`
}

// NewOmnicoreListDeviceStatesResponse instantiates a new OmnicoreListDeviceStatesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOmnicoreListDeviceStatesResponse() *OmnicoreListDeviceStatesResponse {
	this := OmnicoreListDeviceStatesResponse{}
	return &this
}

// NewOmnicoreListDeviceStatesResponseWithDefaults instantiates a new OmnicoreListDeviceStatesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOmnicoreListDeviceStatesResponseWithDefaults() *OmnicoreListDeviceStatesResponse {
	this := OmnicoreListDeviceStatesResponse{}
	return &this
}

// GetDeviceStates returns the DeviceStates field value if set, zero value otherwise.
func (o *OmnicoreListDeviceStatesResponse) GetDeviceStates() []OmnicoreDeviceState {
	if o == nil || IsNil(o.DeviceStates) {
		var ret []OmnicoreDeviceState
		return ret
	}
	return o.DeviceStates
}

// GetDeviceStatesOk returns a tuple with the DeviceStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreListDeviceStatesResponse) GetDeviceStatesOk() ([]OmnicoreDeviceState, bool) {
	if o == nil || IsNil(o.DeviceStates) {
		return nil, false
	}
	return o.DeviceStates, true
}

// HasDeviceStates returns a boolean if a field has been set.
func (o *OmnicoreListDeviceStatesResponse) HasDeviceStates() bool {
	if o != nil && !IsNil(o.DeviceStates) {
		return true
	}

	return false
}

// SetDeviceStates gets a reference to the given []OmnicoreDeviceState and assigns it to the DeviceStates field.
func (o *OmnicoreListDeviceStatesResponse) SetDeviceStates(v []OmnicoreDeviceState) {
	o.DeviceStates = v
}

func (o OmnicoreListDeviceStatesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OmnicoreListDeviceStatesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceStates) {
		toSerialize["deviceStates"] = o.DeviceStates
	}
	return toSerialize, nil
}

type NullableOmnicoreListDeviceStatesResponse struct {
	value *OmnicoreListDeviceStatesResponse
	isSet bool
}

func (v NullableOmnicoreListDeviceStatesResponse) Get() *OmnicoreListDeviceStatesResponse {
	return v.value
}

func (v *NullableOmnicoreListDeviceStatesResponse) Set(val *OmnicoreListDeviceStatesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOmnicoreListDeviceStatesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOmnicoreListDeviceStatesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOmnicoreListDeviceStatesResponse(val *OmnicoreListDeviceStatesResponse) *NullableOmnicoreListDeviceStatesResponse {
	return &NullableOmnicoreListDeviceStatesResponse{value: val, isSet: true}
}

func (v NullableOmnicoreListDeviceStatesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOmnicoreListDeviceStatesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


