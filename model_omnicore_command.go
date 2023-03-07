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

// checks if the OmnicoreCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OmnicoreCommand{}

// OmnicoreCommand struct for OmnicoreCommand
type OmnicoreCommand struct {
	BinaryData string `json:"binaryData"`
	Subfolder *string `json:"subfolder,omitempty"`
}

// NewOmnicoreCommand instantiates a new OmnicoreCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOmnicoreCommand(binaryData string) *OmnicoreCommand {
	this := OmnicoreCommand{}
	this.BinaryData = binaryData
	return &this
}

// NewOmnicoreCommandWithDefaults instantiates a new OmnicoreCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOmnicoreCommandWithDefaults() *OmnicoreCommand {
	this := OmnicoreCommand{}
	return &this
}

// GetBinaryData returns the BinaryData field value
func (o *OmnicoreCommand) GetBinaryData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BinaryData
}

// GetBinaryDataOk returns a tuple with the BinaryData field value
// and a boolean to check if the value has been set.
func (o *OmnicoreCommand) GetBinaryDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BinaryData, true
}

// SetBinaryData sets field value
func (o *OmnicoreCommand) SetBinaryData(v string) {
	o.BinaryData = v
}

// GetSubfolder returns the Subfolder field value if set, zero value otherwise.
func (o *OmnicoreCommand) GetSubfolder() string {
	if o == nil || IsNil(o.Subfolder) {
		var ret string
		return ret
	}
	return *o.Subfolder
}

// GetSubfolderOk returns a tuple with the Subfolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreCommand) GetSubfolderOk() (*string, bool) {
	if o == nil || IsNil(o.Subfolder) {
		return nil, false
	}
	return o.Subfolder, true
}

// HasSubfolder returns a boolean if a field has been set.
func (o *OmnicoreCommand) HasSubfolder() bool {
	if o != nil && !IsNil(o.Subfolder) {
		return true
	}

	return false
}

// SetSubfolder gets a reference to the given string and assigns it to the Subfolder field.
func (o *OmnicoreCommand) SetSubfolder(v string) {
	o.Subfolder = &v
}

func (o OmnicoreCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OmnicoreCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["binaryData"] = o.BinaryData
	if !IsNil(o.Subfolder) {
		toSerialize["subfolder"] = o.Subfolder
	}
	return toSerialize, nil
}

type NullableOmnicoreCommand struct {
	value *OmnicoreCommand
	isSet bool
}

func (v NullableOmnicoreCommand) Get() *OmnicoreCommand {
	return v.value
}

func (v *NullableOmnicoreCommand) Set(val *OmnicoreCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableOmnicoreCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableOmnicoreCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOmnicoreCommand(val *OmnicoreCommand) *NullableOmnicoreCommand {
	return &NullableOmnicoreCommand{value: val, isSet: true}
}

func (v NullableOmnicoreCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOmnicoreCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


