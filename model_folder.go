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

// checks if the Folder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Folder{}

// Folder struct for Folder
type Folder struct {
	NoOfFiles *int32 `json:"noOfFiles,omitempty"`
	FileSize *float32 `json:"fileSize,omitempty"`
	Noofoperations *int32 `json:"noofoperations,omitempty"`
	Updatedon *string `json:"updatedon,omitempty"`
	Registryid *string `json:"registryid,omitempty"`
}

// NewFolder instantiates a new Folder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFolder() *Folder {
	this := Folder{}
	return &this
}

// NewFolderWithDefaults instantiates a new Folder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFolderWithDefaults() *Folder {
	this := Folder{}
	return &this
}

// GetNoOfFiles returns the NoOfFiles field value if set, zero value otherwise.
func (o *Folder) GetNoOfFiles() int32 {
	if o == nil || IsNil(o.NoOfFiles) {
		var ret int32
		return ret
	}
	return *o.NoOfFiles
}

// GetNoOfFilesOk returns a tuple with the NoOfFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folder) GetNoOfFilesOk() (*int32, bool) {
	if o == nil || IsNil(o.NoOfFiles) {
		return nil, false
	}
	return o.NoOfFiles, true
}

// HasNoOfFiles returns a boolean if a field has been set.
func (o *Folder) HasNoOfFiles() bool {
	if o != nil && !IsNil(o.NoOfFiles) {
		return true
	}

	return false
}

// SetNoOfFiles gets a reference to the given int32 and assigns it to the NoOfFiles field.
func (o *Folder) SetNoOfFiles(v int32) {
	o.NoOfFiles = &v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *Folder) GetFileSize() float32 {
	if o == nil || IsNil(o.FileSize) {
		var ret float32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folder) GetFileSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.FileSize) {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *Folder) HasFileSize() bool {
	if o != nil && !IsNil(o.FileSize) {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given float32 and assigns it to the FileSize field.
func (o *Folder) SetFileSize(v float32) {
	o.FileSize = &v
}

// GetNoofoperations returns the Noofoperations field value if set, zero value otherwise.
func (o *Folder) GetNoofoperations() int32 {
	if o == nil || IsNil(o.Noofoperations) {
		var ret int32
		return ret
	}
	return *o.Noofoperations
}

// GetNoofoperationsOk returns a tuple with the Noofoperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folder) GetNoofoperationsOk() (*int32, bool) {
	if o == nil || IsNil(o.Noofoperations) {
		return nil, false
	}
	return o.Noofoperations, true
}

// HasNoofoperations returns a boolean if a field has been set.
func (o *Folder) HasNoofoperations() bool {
	if o != nil && !IsNil(o.Noofoperations) {
		return true
	}

	return false
}

// SetNoofoperations gets a reference to the given int32 and assigns it to the Noofoperations field.
func (o *Folder) SetNoofoperations(v int32) {
	o.Noofoperations = &v
}

// GetUpdatedon returns the Updatedon field value if set, zero value otherwise.
func (o *Folder) GetUpdatedon() string {
	if o == nil || IsNil(o.Updatedon) {
		var ret string
		return ret
	}
	return *o.Updatedon
}

// GetUpdatedonOk returns a tuple with the Updatedon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folder) GetUpdatedonOk() (*string, bool) {
	if o == nil || IsNil(o.Updatedon) {
		return nil, false
	}
	return o.Updatedon, true
}

// HasUpdatedon returns a boolean if a field has been set.
func (o *Folder) HasUpdatedon() bool {
	if o != nil && !IsNil(o.Updatedon) {
		return true
	}

	return false
}

// SetUpdatedon gets a reference to the given string and assigns it to the Updatedon field.
func (o *Folder) SetUpdatedon(v string) {
	o.Updatedon = &v
}

// GetRegistryid returns the Registryid field value if set, zero value otherwise.
func (o *Folder) GetRegistryid() string {
	if o == nil || IsNil(o.Registryid) {
		var ret string
		return ret
	}
	return *o.Registryid
}

// GetRegistryidOk returns a tuple with the Registryid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folder) GetRegistryidOk() (*string, bool) {
	if o == nil || IsNil(o.Registryid) {
		return nil, false
	}
	return o.Registryid, true
}

// HasRegistryid returns a boolean if a field has been set.
func (o *Folder) HasRegistryid() bool {
	if o != nil && !IsNil(o.Registryid) {
		return true
	}

	return false
}

// SetRegistryid gets a reference to the given string and assigns it to the Registryid field.
func (o *Folder) SetRegistryid(v string) {
	o.Registryid = &v
}

func (o Folder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Folder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NoOfFiles) {
		toSerialize["noOfFiles"] = o.NoOfFiles
	}
	if !IsNil(o.FileSize) {
		toSerialize["fileSize"] = o.FileSize
	}
	if !IsNil(o.Noofoperations) {
		toSerialize["noofoperations"] = o.Noofoperations
	}
	if !IsNil(o.Updatedon) {
		toSerialize["updatedon"] = o.Updatedon
	}
	if !IsNil(o.Registryid) {
		toSerialize["registryid"] = o.Registryid
	}
	return toSerialize, nil
}

type NullableFolder struct {
	value *Folder
	isSet bool
}

func (v NullableFolder) Get() *Folder {
	return v.value
}

func (v *NullableFolder) Set(val *Folder) {
	v.value = val
	v.isSet = true
}

func (v NullableFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFolder(val *Folder) *NullableFolder {
	return &NullableFolder{value: val, isSet: true}
}

func (v NullableFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


