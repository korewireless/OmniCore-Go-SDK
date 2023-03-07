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

// checks if the OmnicoreInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OmnicoreInfo{}

// OmnicoreInfo struct for OmnicoreInfo
type OmnicoreInfo struct {
	Info *string `json:"info,omitempty"`
}

// NewOmnicoreInfo instantiates a new OmnicoreInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOmnicoreInfo() *OmnicoreInfo {
	this := OmnicoreInfo{}
	return &this
}

// NewOmnicoreInfoWithDefaults instantiates a new OmnicoreInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOmnicoreInfoWithDefaults() *OmnicoreInfo {
	this := OmnicoreInfo{}
	return &this
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *OmnicoreInfo) GetInfo() string {
	if o == nil || IsNil(o.Info) {
		var ret string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreInfo) GetInfoOk() (*string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *OmnicoreInfo) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given string and assigns it to the Info field.
func (o *OmnicoreInfo) SetInfo(v string) {
	o.Info = &v
}

func (o OmnicoreInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OmnicoreInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Info) {
		toSerialize["info"] = o.Info
	}
	return toSerialize, nil
}

type NullableOmnicoreInfo struct {
	value *OmnicoreInfo
	isSet bool
}

func (v NullableOmnicoreInfo) Get() *OmnicoreInfo {
	return v.value
}

func (v *NullableOmnicoreInfo) Set(val *OmnicoreInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOmnicoreInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOmnicoreInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOmnicoreInfo(val *OmnicoreInfo) *NullableOmnicoreInfo {
	return &NullableOmnicoreInfo{value: val, isSet: true}
}

func (v NullableOmnicoreInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOmnicoreInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


