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
	TotalExports *int32 `json:"TotalExports,omitempty"`
	TotalReplays *int32 `json:"TotalReplays,omitempty"`
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

// GetTotalExports returns the TotalExports field value if set, zero value otherwise.
func (o *MetricsData) GetTotalExports() int32 {
	if o == nil || IsNil(o.TotalExports) {
		var ret int32
		return ret
	}
	return *o.TotalExports
}

// GetTotalExportsOk returns a tuple with the TotalExports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsData) GetTotalExportsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalExports) {
		return nil, false
	}
	return o.TotalExports, true
}

// HasTotalExports returns a boolean if a field has been set.
func (o *MetricsData) HasTotalExports() bool {
	if o != nil && !IsNil(o.TotalExports) {
		return true
	}

	return false
}

// SetTotalExports gets a reference to the given int32 and assigns it to the TotalExports field.
func (o *MetricsData) SetTotalExports(v int32) {
	o.TotalExports = &v
}

// GetTotalReplays returns the TotalReplays field value if set, zero value otherwise.
func (o *MetricsData) GetTotalReplays() int32 {
	if o == nil || IsNil(o.TotalReplays) {
		var ret int32
		return ret
	}
	return *o.TotalReplays
}

// GetTotalReplaysOk returns a tuple with the TotalReplays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsData) GetTotalReplaysOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalReplays) {
		return nil, false
	}
	return o.TotalReplays, true
}

// HasTotalReplays returns a boolean if a field has been set.
func (o *MetricsData) HasTotalReplays() bool {
	if o != nil && !IsNil(o.TotalReplays) {
		return true
	}

	return false
}

// SetTotalReplays gets a reference to the given int32 and assigns it to the TotalReplays field.
func (o *MetricsData) SetTotalReplays(v int32) {
	o.TotalReplays = &v
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
	if !IsNil(o.TotalExports) {
		toSerialize["TotalExports"] = o.TotalExports
	}
	if !IsNil(o.TotalReplays) {
		toSerialize["TotalReplays"] = o.TotalReplays
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


