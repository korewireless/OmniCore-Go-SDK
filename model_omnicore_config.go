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

// checks if the OmnicoreConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OmnicoreConfig{}

// OmnicoreConfig struct for OmnicoreConfig
type OmnicoreConfig struct {
	BinaryData string `json:"binaryData"`
	Subfolder *string `json:"subfolder,omitempty"`
	VersionToUpdate *string `json:"versionToUpdate,omitempty"`
}

// NewOmnicoreConfig instantiates a new OmnicoreConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOmnicoreConfig(binaryData string) *OmnicoreConfig {
	this := OmnicoreConfig{}
	this.BinaryData = binaryData
	return &this
}

// NewOmnicoreConfigWithDefaults instantiates a new OmnicoreConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOmnicoreConfigWithDefaults() *OmnicoreConfig {
	this := OmnicoreConfig{}
	return &this
}

// GetBinaryData returns the BinaryData field value
func (o *OmnicoreConfig) GetBinaryData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BinaryData
}

// GetBinaryDataOk returns a tuple with the BinaryData field value
// and a boolean to check if the value has been set.
func (o *OmnicoreConfig) GetBinaryDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BinaryData, true
}

// SetBinaryData sets field value
func (o *OmnicoreConfig) SetBinaryData(v string) {
	o.BinaryData = v
}

// GetSubfolder returns the Subfolder field value if set, zero value otherwise.
func (o *OmnicoreConfig) GetSubfolder() string {
	if o == nil || IsNil(o.Subfolder) {
		var ret string
		return ret
	}
	return *o.Subfolder
}

// GetSubfolderOk returns a tuple with the Subfolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreConfig) GetSubfolderOk() (*string, bool) {
	if o == nil || IsNil(o.Subfolder) {
		return nil, false
	}
	return o.Subfolder, true
}

// HasSubfolder returns a boolean if a field has been set.
func (o *OmnicoreConfig) HasSubfolder() bool {
	if o != nil && !IsNil(o.Subfolder) {
		return true
	}

	return false
}

// SetSubfolder gets a reference to the given string and assigns it to the Subfolder field.
func (o *OmnicoreConfig) SetSubfolder(v string) {
	o.Subfolder = &v
}

// GetVersionToUpdate returns the VersionToUpdate field value if set, zero value otherwise.
func (o *OmnicoreConfig) GetVersionToUpdate() string {
	if o == nil || IsNil(o.VersionToUpdate) {
		var ret string
		return ret
	}
	return *o.VersionToUpdate
}

// GetVersionToUpdateOk returns a tuple with the VersionToUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreConfig) GetVersionToUpdateOk() (*string, bool) {
	if o == nil || IsNil(o.VersionToUpdate) {
		return nil, false
	}
	return o.VersionToUpdate, true
}

// HasVersionToUpdate returns a boolean if a field has been set.
func (o *OmnicoreConfig) HasVersionToUpdate() bool {
	if o != nil && !IsNil(o.VersionToUpdate) {
		return true
	}

	return false
}

// SetVersionToUpdate gets a reference to the given string and assigns it to the VersionToUpdate field.
func (o *OmnicoreConfig) SetVersionToUpdate(v string) {
	o.VersionToUpdate = &v
}

func (o OmnicoreConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OmnicoreConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["binaryData"] = o.BinaryData
	if !IsNil(o.Subfolder) {
		toSerialize["subfolder"] = o.Subfolder
	}
	if !IsNil(o.VersionToUpdate) {
		toSerialize["versionToUpdate"] = o.VersionToUpdate
	}
	return toSerialize, nil
}

type NullableOmnicoreConfig struct {
	value *OmnicoreConfig
	isSet bool
}

func (v NullableOmnicoreConfig) Get() *OmnicoreConfig {
	return v.value
}

func (v *NullableOmnicoreConfig) Set(val *OmnicoreConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableOmnicoreConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableOmnicoreConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOmnicoreConfig(val *OmnicoreConfig) *NullableOmnicoreConfig {
	return &NullableOmnicoreConfig{value: val, isSet: true}
}

func (v NullableOmnicoreConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOmnicoreConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


