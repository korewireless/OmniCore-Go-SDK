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

// checks if the ListDeviceRegistries type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDeviceRegistries{}

// ListDeviceRegistries struct for ListDeviceRegistries
type ListDeviceRegistries struct {
	DeviceRegistries []DeviceRegistry `json:"deviceRegistries"`
	PageNumber *int32 `json:"pageNumber,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewListDeviceRegistries instantiates a new ListDeviceRegistries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDeviceRegistries(deviceRegistries []DeviceRegistry) *ListDeviceRegistries {
	this := ListDeviceRegistries{}
	this.DeviceRegistries = deviceRegistries
	return &this
}

// NewListDeviceRegistriesWithDefaults instantiates a new ListDeviceRegistries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDeviceRegistriesWithDefaults() *ListDeviceRegistries {
	this := ListDeviceRegistries{}
	return &this
}

// GetDeviceRegistries returns the DeviceRegistries field value
func (o *ListDeviceRegistries) GetDeviceRegistries() []DeviceRegistry {
	if o == nil {
		var ret []DeviceRegistry
		return ret
	}

	return o.DeviceRegistries
}

// GetDeviceRegistriesOk returns a tuple with the DeviceRegistries field value
// and a boolean to check if the value has been set.
func (o *ListDeviceRegistries) GetDeviceRegistriesOk() ([]DeviceRegistry, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceRegistries, true
}

// SetDeviceRegistries sets field value
func (o *ListDeviceRegistries) SetDeviceRegistries(v []DeviceRegistry) {
	o.DeviceRegistries = v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *ListDeviceRegistries) GetPageNumber() int32 {
	if o == nil || IsNil(o.PageNumber) {
		var ret int32
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDeviceRegistries) GetPageNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *ListDeviceRegistries) HasPageNumber() bool {
	if o != nil && !IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int32 and assigns it to the PageNumber field.
func (o *ListDeviceRegistries) SetPageNumber(v int32) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ListDeviceRegistries) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDeviceRegistries) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ListDeviceRegistries) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *ListDeviceRegistries) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ListDeviceRegistries) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDeviceRegistries) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ListDeviceRegistries) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *ListDeviceRegistries) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o ListDeviceRegistries) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDeviceRegistries) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deviceRegistries"] = o.DeviceRegistries
	if !IsNil(o.PageNumber) {
		toSerialize["pageNumber"] = o.PageNumber
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableListDeviceRegistries struct {
	value *ListDeviceRegistries
	isSet bool
}

func (v NullableListDeviceRegistries) Get() *ListDeviceRegistries {
	return v.value
}

func (v *NullableListDeviceRegistries) Set(val *ListDeviceRegistries) {
	v.value = val
	v.isSet = true
}

func (v NullableListDeviceRegistries) IsSet() bool {
	return v.isSet
}

func (v *NullableListDeviceRegistries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDeviceRegistries(val *ListDeviceRegistries) *NullableListDeviceRegistries {
	return &NullableListDeviceRegistries{value: val, isSet: true}
}

func (v NullableListDeviceRegistries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDeviceRegistries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


