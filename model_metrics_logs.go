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

// checks if the MetricsLogs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsLogs{}

// MetricsLogs struct for MetricsLogs
type MetricsLogs struct {
	NoOfFiles *int32 `json:"noOfFiles,omitempty"`
	FileSize *float32 `json:"fileSize,omitempty"`
	Noofoperations *int32 `json:"noofoperations,omitempty"`
	Updatedon *string `json:"updatedon,omitempty"`
	NoOfReplays *int32 `json:"noOfReplays,omitempty"`
	NoOfExports *int32 `json:"noOfExports,omitempty"`
}

// NewMetricsLogs instantiates a new MetricsLogs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsLogs() *MetricsLogs {
	this := MetricsLogs{}
	return &this
}

// NewMetricsLogsWithDefaults instantiates a new MetricsLogs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsLogsWithDefaults() *MetricsLogs {
	this := MetricsLogs{}
	return &this
}

// GetNoOfFiles returns the NoOfFiles field value if set, zero value otherwise.
func (o *MetricsLogs) GetNoOfFiles() int32 {
	if o == nil || IsNil(o.NoOfFiles) {
		var ret int32
		return ret
	}
	return *o.NoOfFiles
}

// GetNoOfFilesOk returns a tuple with the NoOfFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsLogs) GetNoOfFilesOk() (*int32, bool) {
	if o == nil || IsNil(o.NoOfFiles) {
		return nil, false
	}
	return o.NoOfFiles, true
}

// HasNoOfFiles returns a boolean if a field has been set.
func (o *MetricsLogs) HasNoOfFiles() bool {
	if o != nil && !IsNil(o.NoOfFiles) {
		return true
	}

	return false
}

// SetNoOfFiles gets a reference to the given int32 and assigns it to the NoOfFiles field.
func (o *MetricsLogs) SetNoOfFiles(v int32) {
	o.NoOfFiles = &v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *MetricsLogs) GetFileSize() float32 {
	if o == nil || IsNil(o.FileSize) {
		var ret float32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsLogs) GetFileSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.FileSize) {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *MetricsLogs) HasFileSize() bool {
	if o != nil && !IsNil(o.FileSize) {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given float32 and assigns it to the FileSize field.
func (o *MetricsLogs) SetFileSize(v float32) {
	o.FileSize = &v
}

// GetNoofoperations returns the Noofoperations field value if set, zero value otherwise.
func (o *MetricsLogs) GetNoofoperations() int32 {
	if o == nil || IsNil(o.Noofoperations) {
		var ret int32
		return ret
	}
	return *o.Noofoperations
}

// GetNoofoperationsOk returns a tuple with the Noofoperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsLogs) GetNoofoperationsOk() (*int32, bool) {
	if o == nil || IsNil(o.Noofoperations) {
		return nil, false
	}
	return o.Noofoperations, true
}

// HasNoofoperations returns a boolean if a field has been set.
func (o *MetricsLogs) HasNoofoperations() bool {
	if o != nil && !IsNil(o.Noofoperations) {
		return true
	}

	return false
}

// SetNoofoperations gets a reference to the given int32 and assigns it to the Noofoperations field.
func (o *MetricsLogs) SetNoofoperations(v int32) {
	o.Noofoperations = &v
}

// GetUpdatedon returns the Updatedon field value if set, zero value otherwise.
func (o *MetricsLogs) GetUpdatedon() string {
	if o == nil || IsNil(o.Updatedon) {
		var ret string
		return ret
	}
	return *o.Updatedon
}

// GetUpdatedonOk returns a tuple with the Updatedon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsLogs) GetUpdatedonOk() (*string, bool) {
	if o == nil || IsNil(o.Updatedon) {
		return nil, false
	}
	return o.Updatedon, true
}

// HasUpdatedon returns a boolean if a field has been set.
func (o *MetricsLogs) HasUpdatedon() bool {
	if o != nil && !IsNil(o.Updatedon) {
		return true
	}

	return false
}

// SetUpdatedon gets a reference to the given string and assigns it to the Updatedon field.
func (o *MetricsLogs) SetUpdatedon(v string) {
	o.Updatedon = &v
}

// GetNoOfReplays returns the NoOfReplays field value if set, zero value otherwise.
func (o *MetricsLogs) GetNoOfReplays() int32 {
	if o == nil || IsNil(o.NoOfReplays) {
		var ret int32
		return ret
	}
	return *o.NoOfReplays
}

// GetNoOfReplaysOk returns a tuple with the NoOfReplays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsLogs) GetNoOfReplaysOk() (*int32, bool) {
	if o == nil || IsNil(o.NoOfReplays) {
		return nil, false
	}
	return o.NoOfReplays, true
}

// HasNoOfReplays returns a boolean if a field has been set.
func (o *MetricsLogs) HasNoOfReplays() bool {
	if o != nil && !IsNil(o.NoOfReplays) {
		return true
	}

	return false
}

// SetNoOfReplays gets a reference to the given int32 and assigns it to the NoOfReplays field.
func (o *MetricsLogs) SetNoOfReplays(v int32) {
	o.NoOfReplays = &v
}

// GetNoOfExports returns the NoOfExports field value if set, zero value otherwise.
func (o *MetricsLogs) GetNoOfExports() int32 {
	if o == nil || IsNil(o.NoOfExports) {
		var ret int32
		return ret
	}
	return *o.NoOfExports
}

// GetNoOfExportsOk returns a tuple with the NoOfExports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsLogs) GetNoOfExportsOk() (*int32, bool) {
	if o == nil || IsNil(o.NoOfExports) {
		return nil, false
	}
	return o.NoOfExports, true
}

// HasNoOfExports returns a boolean if a field has been set.
func (o *MetricsLogs) HasNoOfExports() bool {
	if o != nil && !IsNil(o.NoOfExports) {
		return true
	}

	return false
}

// SetNoOfExports gets a reference to the given int32 and assigns it to the NoOfExports field.
func (o *MetricsLogs) SetNoOfExports(v int32) {
	o.NoOfExports = &v
}

func (o MetricsLogs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsLogs) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.NoOfReplays) {
		toSerialize["noOfReplays"] = o.NoOfReplays
	}
	if !IsNil(o.NoOfExports) {
		toSerialize["noOfExports"] = o.NoOfExports
	}
	return toSerialize, nil
}

type NullableMetricsLogs struct {
	value *MetricsLogs
	isSet bool
}

func (v NullableMetricsLogs) Get() *MetricsLogs {
	return v.value
}

func (v *NullableMetricsLogs) Set(val *MetricsLogs) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsLogs) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsLogs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsLogs(val *MetricsLogs) *NullableMetricsLogs {
	return &NullableMetricsLogs{value: val, isSet: true}
}

func (v NullableMetricsLogs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsLogs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


