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

// checks if the ListDeviceRegistriesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDeviceRegistriesResponse{}

// ListDeviceRegistriesResponse struct for ListDeviceRegistriesResponse
type ListDeviceRegistriesResponse struct {
	DeviceRegistries []DeviceRegistry `json:"deviceRegistries"`
	PageNumber *int32 `json:"pageNumber,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewListDeviceRegistriesResponse instantiates a new ListDeviceRegistriesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDeviceRegistriesResponse(deviceRegistries []DeviceRegistry) *ListDeviceRegistriesResponse {
	this := ListDeviceRegistriesResponse{}
	this.DeviceRegistries = deviceRegistries
	return &this
}

// NewListDeviceRegistriesResponseWithDefaults instantiates a new ListDeviceRegistriesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDeviceRegistriesResponseWithDefaults() *ListDeviceRegistriesResponse {
	this := ListDeviceRegistriesResponse{}
	return &this
}

// GetDeviceRegistries returns the DeviceRegistries field value
func (o *ListDeviceRegistriesResponse) GetDeviceRegistries() []DeviceRegistry {
	if o == nil {
		var ret []DeviceRegistry
		return ret
	}

	return o.DeviceRegistries
}

// GetDeviceRegistriesOk returns a tuple with the DeviceRegistries field value
// and a boolean to check if the value has been set.
func (o *ListDeviceRegistriesResponse) GetDeviceRegistriesOk() ([]DeviceRegistry, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceRegistries, true
}

// SetDeviceRegistries sets field value
func (o *ListDeviceRegistriesResponse) SetDeviceRegistries(v []DeviceRegistry) {
	o.DeviceRegistries = v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *ListDeviceRegistriesResponse) GetPageNumber() int32 {
	if o == nil || IsNil(o.PageNumber) {
		var ret int32
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDeviceRegistriesResponse) GetPageNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *ListDeviceRegistriesResponse) HasPageNumber() bool {
	if o != nil && !IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int32 and assigns it to the PageNumber field.
func (o *ListDeviceRegistriesResponse) SetPageNumber(v int32) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ListDeviceRegistriesResponse) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDeviceRegistriesResponse) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ListDeviceRegistriesResponse) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *ListDeviceRegistriesResponse) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ListDeviceRegistriesResponse) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDeviceRegistriesResponse) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ListDeviceRegistriesResponse) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *ListDeviceRegistriesResponse) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o ListDeviceRegistriesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDeviceRegistriesResponse) ToMap() (map[string]interface{}, error) {
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

type NullableListDeviceRegistriesResponse struct {
	value *ListDeviceRegistriesResponse
	isSet bool
}

func (v NullableListDeviceRegistriesResponse) Get() *ListDeviceRegistriesResponse {
	return v.value
}

func (v *NullableListDeviceRegistriesResponse) Set(val *ListDeviceRegistriesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListDeviceRegistriesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListDeviceRegistriesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDeviceRegistriesResponse(val *ListDeviceRegistriesResponse) *NullableListDeviceRegistriesResponse {
	return &NullableListDeviceRegistriesResponse{value: val, isSet: true}
}

func (v NullableListDeviceRegistriesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDeviceRegistriesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


