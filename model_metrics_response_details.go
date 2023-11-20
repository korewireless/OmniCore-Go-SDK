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

// checks if the MetricsResponseDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsResponseDetails{}

// MetricsResponseDetails struct for MetricsResponseDetails
type MetricsResponseDetails struct {
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	NoOfFiles *int32 `json:"noOfFiles,omitempty"`
	FileSize *float32 `json:"fileSize,omitempty"`
	Noofoperations *int32 `json:"noofoperations,omitempty"`
	Operations []OperationMetrics `json:"Operations,omitempty"`
	DetailsForTimePeriod []MetricsData `json:"detailsForTimePeriod,omitempty"`
}

// NewMetricsResponseDetails instantiates a new MetricsResponseDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsResponseDetails() *MetricsResponseDetails {
	this := MetricsResponseDetails{}
	return &this
}

// NewMetricsResponseDetailsWithDefaults instantiates a new MetricsResponseDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsResponseDetailsWithDefaults() *MetricsResponseDetails {
	this := MetricsResponseDetails{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *MetricsResponseDetails) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsResponseDetails) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *MetricsResponseDetails) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *MetricsResponseDetails) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetNoOfFiles returns the NoOfFiles field value if set, zero value otherwise.
func (o *MetricsResponseDetails) GetNoOfFiles() int32 {
	if o == nil || IsNil(o.NoOfFiles) {
		var ret int32
		return ret
	}
	return *o.NoOfFiles
}

// GetNoOfFilesOk returns a tuple with the NoOfFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsResponseDetails) GetNoOfFilesOk() (*int32, bool) {
	if o == nil || IsNil(o.NoOfFiles) {
		return nil, false
	}
	return o.NoOfFiles, true
}

// HasNoOfFiles returns a boolean if a field has been set.
func (o *MetricsResponseDetails) HasNoOfFiles() bool {
	if o != nil && !IsNil(o.NoOfFiles) {
		return true
	}

	return false
}

// SetNoOfFiles gets a reference to the given int32 and assigns it to the NoOfFiles field.
func (o *MetricsResponseDetails) SetNoOfFiles(v int32) {
	o.NoOfFiles = &v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *MetricsResponseDetails) GetFileSize() float32 {
	if o == nil || IsNil(o.FileSize) {
		var ret float32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsResponseDetails) GetFileSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.FileSize) {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *MetricsResponseDetails) HasFileSize() bool {
	if o != nil && !IsNil(o.FileSize) {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given float32 and assigns it to the FileSize field.
func (o *MetricsResponseDetails) SetFileSize(v float32) {
	o.FileSize = &v
}

// GetNoofoperations returns the Noofoperations field value if set, zero value otherwise.
func (o *MetricsResponseDetails) GetNoofoperations() int32 {
	if o == nil || IsNil(o.Noofoperations) {
		var ret int32
		return ret
	}
	return *o.Noofoperations
}

// GetNoofoperationsOk returns a tuple with the Noofoperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsResponseDetails) GetNoofoperationsOk() (*int32, bool) {
	if o == nil || IsNil(o.Noofoperations) {
		return nil, false
	}
	return o.Noofoperations, true
}

// HasNoofoperations returns a boolean if a field has been set.
func (o *MetricsResponseDetails) HasNoofoperations() bool {
	if o != nil && !IsNil(o.Noofoperations) {
		return true
	}

	return false
}

// SetNoofoperations gets a reference to the given int32 and assigns it to the Noofoperations field.
func (o *MetricsResponseDetails) SetNoofoperations(v int32) {
	o.Noofoperations = &v
}

// GetOperations returns the Operations field value if set, zero value otherwise.
func (o *MetricsResponseDetails) GetOperations() []OperationMetrics {
	if o == nil || IsNil(o.Operations) {
		var ret []OperationMetrics
		return ret
	}
	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsResponseDetails) GetOperationsOk() ([]OperationMetrics, bool) {
	if o == nil || IsNil(o.Operations) {
		return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *MetricsResponseDetails) HasOperations() bool {
	if o != nil && !IsNil(o.Operations) {
		return true
	}

	return false
}

// SetOperations gets a reference to the given []OperationMetrics and assigns it to the Operations field.
func (o *MetricsResponseDetails) SetOperations(v []OperationMetrics) {
	o.Operations = v
}

// GetDetailsForTimePeriod returns the DetailsForTimePeriod field value if set, zero value otherwise.
func (o *MetricsResponseDetails) GetDetailsForTimePeriod() []MetricsData {
	if o == nil || IsNil(o.DetailsForTimePeriod) {
		var ret []MetricsData
		return ret
	}
	return o.DetailsForTimePeriod
}

// GetDetailsForTimePeriodOk returns a tuple with the DetailsForTimePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsResponseDetails) GetDetailsForTimePeriodOk() ([]MetricsData, bool) {
	if o == nil || IsNil(o.DetailsForTimePeriod) {
		return nil, false
	}
	return o.DetailsForTimePeriod, true
}

// HasDetailsForTimePeriod returns a boolean if a field has been set.
func (o *MetricsResponseDetails) HasDetailsForTimePeriod() bool {
	if o != nil && !IsNil(o.DetailsForTimePeriod) {
		return true
	}

	return false
}

// SetDetailsForTimePeriod gets a reference to the given []MetricsData and assigns it to the DetailsForTimePeriod field.
func (o *MetricsResponseDetails) SetDetailsForTimePeriod(v []MetricsData) {
	o.DetailsForTimePeriod = v
}

func (o MetricsResponseDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsResponseDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !IsNil(o.NoOfFiles) {
		toSerialize["noOfFiles"] = o.NoOfFiles
	}
	if !IsNil(o.FileSize) {
		toSerialize["fileSize"] = o.FileSize
	}
	if !IsNil(o.Noofoperations) {
		toSerialize["noofoperations"] = o.Noofoperations
	}
	if !IsNil(o.Operations) {
		toSerialize["Operations"] = o.Operations
	}
	if !IsNil(o.DetailsForTimePeriod) {
		toSerialize["detailsForTimePeriod"] = o.DetailsForTimePeriod
	}
	return toSerialize, nil
}

type NullableMetricsResponseDetails struct {
	value *MetricsResponseDetails
	isSet bool
}

func (v NullableMetricsResponseDetails) Get() *MetricsResponseDetails {
	return v.value
}

func (v *NullableMetricsResponseDetails) Set(val *MetricsResponseDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsResponseDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsResponseDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsResponseDetails(val *MetricsResponseDetails) *NullableMetricsResponseDetails {
	return &NullableMetricsResponseDetails{value: val, isSet: true}
}

func (v NullableMetricsResponseDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsResponseDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


