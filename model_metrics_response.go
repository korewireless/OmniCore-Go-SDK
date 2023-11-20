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

// checks if the MetricsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsResponse{}

// MetricsResponse struct for MetricsResponse
type MetricsResponse struct {
	Details *MetricsResponseDetails `json:"details,omitempty"`
}

// NewMetricsResponse instantiates a new MetricsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsResponse() *MetricsResponse {
	this := MetricsResponse{}
	return &this
}

// NewMetricsResponseWithDefaults instantiates a new MetricsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsResponseWithDefaults() *MetricsResponse {
	this := MetricsResponse{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *MetricsResponse) GetDetails() MetricsResponseDetails {
	if o == nil || IsNil(o.Details) {
		var ret MetricsResponseDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsResponse) GetDetailsOk() (*MetricsResponseDetails, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *MetricsResponse) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given MetricsResponseDetails and assigns it to the Details field.
func (o *MetricsResponse) SetDetails(v MetricsResponseDetails) {
	o.Details = &v
}

func (o MetricsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableMetricsResponse struct {
	value *MetricsResponse
	isSet bool
}

func (v NullableMetricsResponse) Get() *MetricsResponse {
	return v.value
}

func (v *NullableMetricsResponse) Set(val *MetricsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsResponse(val *MetricsResponse) *NullableMetricsResponse {
	return &NullableMetricsResponse{value: val, isSet: true}
}

func (v NullableMetricsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


