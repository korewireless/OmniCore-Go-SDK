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

// checks if the MetricsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsData{}

// MetricsData struct for MetricsData
type MetricsData struct {
	Data []MetricsLogs `json:"data,omitempty"`
	TotalExportSize *float32 `json:"TotalExportSize,omitempty"`
	TotalReplaySize *float32 `json:"TotalReplaySize,omitempty"`
}

// NewMetricsData instantiates a new MetricsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsData() *MetricsData {
	this := MetricsData{}
	return &this
}

// NewMetricsDataWithDefaults instantiates a new MetricsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsDataWithDefaults() *MetricsData {
	this := MetricsData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *MetricsData) GetData() []MetricsLogs {
	if o == nil || IsNil(o.Data) {
		var ret []MetricsLogs
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsData) GetDataOk() ([]MetricsLogs, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *MetricsData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []MetricsLogs and assigns it to the Data field.
func (o *MetricsData) SetData(v []MetricsLogs) {
	o.Data = v
}

// GetTotalExportSize returns the TotalExportSize field value if set, zero value otherwise.
func (o *MetricsData) GetTotalExportSize() float32 {
	if o == nil || IsNil(o.TotalExportSize) {
		var ret float32
		return ret
	}
	return *o.TotalExportSize
}

// GetTotalExportSizeOk returns a tuple with the TotalExportSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsData) GetTotalExportSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalExportSize) {
		return nil, false
	}
	return o.TotalExportSize, true
}

// HasTotalExportSize returns a boolean if a field has been set.
func (o *MetricsData) HasTotalExportSize() bool {
	if o != nil && !IsNil(o.TotalExportSize) {
		return true
	}

	return false
}

// SetTotalExportSize gets a reference to the given float32 and assigns it to the TotalExportSize field.
func (o *MetricsData) SetTotalExportSize(v float32) {
	o.TotalExportSize = &v
}

// GetTotalReplaySize returns the TotalReplaySize field value if set, zero value otherwise.
func (o *MetricsData) GetTotalReplaySize() float32 {
	if o == nil || IsNil(o.TotalReplaySize) {
		var ret float32
		return ret
	}
	return *o.TotalReplaySize
}

// GetTotalReplaySizeOk returns a tuple with the TotalReplaySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsData) GetTotalReplaySizeOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalReplaySize) {
		return nil, false
	}
	return o.TotalReplaySize, true
}

// HasTotalReplaySize returns a boolean if a field has been set.
func (o *MetricsData) HasTotalReplaySize() bool {
	if o != nil && !IsNil(o.TotalReplaySize) {
		return true
	}

	return false
}

// SetTotalReplaySize gets a reference to the given float32 and assigns it to the TotalReplaySize field.
func (o *MetricsData) SetTotalReplaySize(v float32) {
	o.TotalReplaySize = &v
}

func (o MetricsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalExportSize) {
		toSerialize["TotalExportSize"] = o.TotalExportSize
	}
	if !IsNil(o.TotalReplaySize) {
		toSerialize["TotalReplaySize"] = o.TotalReplaySize
	}
	return toSerialize, nil
}

type NullableMetricsData struct {
	value *MetricsData
	isSet bool
}

func (v NullableMetricsData) Get() *MetricsData {
	return v.value
}

func (v *NullableMetricsData) Set(val *MetricsData) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsData) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsData(val *MetricsData) *NullableMetricsData {
	return &NullableMetricsData{value: val, isSet: true}
}

func (v NullableMetricsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


