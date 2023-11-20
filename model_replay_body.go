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

// checks if the ReplayBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplayBody{}

// ReplayBody struct for ReplayBody
type ReplayBody struct {
	Registry *string `json:"registry,omitempty"`
	StartTime *int32 `json:"startTime,omitempty"`
	EndTime *int32 `json:"endTime,omitempty"`
	Destination *string `json:"destination,omitempty"`
	Source *string `json:"source,omitempty"`
}

// NewReplayBody instantiates a new ReplayBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplayBody() *ReplayBody {
	this := ReplayBody{}
	return &this
}

// NewReplayBodyWithDefaults instantiates a new ReplayBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplayBodyWithDefaults() *ReplayBody {
	this := ReplayBody{}
	return &this
}

// GetRegistry returns the Registry field value if set, zero value otherwise.
func (o *ReplayBody) GetRegistry() string {
	if o == nil || IsNil(o.Registry) {
		var ret string
		return ret
	}
	return *o.Registry
}

// GetRegistryOk returns a tuple with the Registry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplayBody) GetRegistryOk() (*string, bool) {
	if o == nil || IsNil(o.Registry) {
		return nil, false
	}
	return o.Registry, true
}

// HasRegistry returns a boolean if a field has been set.
func (o *ReplayBody) HasRegistry() bool {
	if o != nil && !IsNil(o.Registry) {
		return true
	}

	return false
}

// SetRegistry gets a reference to the given string and assigns it to the Registry field.
func (o *ReplayBody) SetRegistry(v string) {
	o.Registry = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ReplayBody) GetStartTime() int32 {
	if o == nil || IsNil(o.StartTime) {
		var ret int32
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplayBody) GetStartTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ReplayBody) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int32 and assigns it to the StartTime field.
func (o *ReplayBody) SetStartTime(v int32) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *ReplayBody) GetEndTime() int32 {
	if o == nil || IsNil(o.EndTime) {
		var ret int32
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplayBody) GetEndTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *ReplayBody) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int32 and assigns it to the EndTime field.
func (o *ReplayBody) SetEndTime(v int32) {
	o.EndTime = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *ReplayBody) GetDestination() string {
	if o == nil || IsNil(o.Destination) {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplayBody) GetDestinationOk() (*string, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *ReplayBody) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *ReplayBody) SetDestination(v string) {
	o.Destination = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ReplayBody) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplayBody) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ReplayBody) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ReplayBody) SetSource(v string) {
	o.Source = &v
}

func (o ReplayBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplayBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Registry) {
		toSerialize["registry"] = o.Registry
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	return toSerialize, nil
}

type NullableReplayBody struct {
	value *ReplayBody
	isSet bool
}

func (v NullableReplayBody) Get() *ReplayBody {
	return v.value
}

func (v *NullableReplayBody) Set(val *ReplayBody) {
	v.value = val
	v.isSet = true
}

func (v NullableReplayBody) IsSet() bool {
	return v.isSet
}

func (v *NullableReplayBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplayBody(val *ReplayBody) *NullableReplayBody {
	return &NullableReplayBody{value: val, isSet: true}
}

func (v NullableReplayBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplayBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


