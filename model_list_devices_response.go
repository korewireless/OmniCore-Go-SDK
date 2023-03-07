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

// checks if the ListDevicesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDevicesResponse{}

// ListDevicesResponse struct for ListDevicesResponse
type ListDevicesResponse struct {
	Devices []Device `json:"devices"`
	PageNumber *int32 `json:"pageNumber,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewListDevicesResponse instantiates a new ListDevicesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDevicesResponse(devices []Device) *ListDevicesResponse {
	this := ListDevicesResponse{}
	this.Devices = devices
	return &this
}

// NewListDevicesResponseWithDefaults instantiates a new ListDevicesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDevicesResponseWithDefaults() *ListDevicesResponse {
	this := ListDevicesResponse{}
	return &this
}

// GetDevices returns the Devices field value
func (o *ListDevicesResponse) GetDevices() []Device {
	if o == nil {
		var ret []Device
		return ret
	}

	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value
// and a boolean to check if the value has been set.
func (o *ListDevicesResponse) GetDevicesOk() ([]Device, bool) {
	if o == nil {
		return nil, false
	}
	return o.Devices, true
}

// SetDevices sets field value
func (o *ListDevicesResponse) SetDevices(v []Device) {
	o.Devices = v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *ListDevicesResponse) GetPageNumber() int32 {
	if o == nil || IsNil(o.PageNumber) {
		var ret int32
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDevicesResponse) GetPageNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *ListDevicesResponse) HasPageNumber() bool {
	if o != nil && !IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int32 and assigns it to the PageNumber field.
func (o *ListDevicesResponse) SetPageNumber(v int32) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ListDevicesResponse) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDevicesResponse) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ListDevicesResponse) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *ListDevicesResponse) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ListDevicesResponse) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDevicesResponse) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ListDevicesResponse) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *ListDevicesResponse) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o ListDevicesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDevicesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["devices"] = o.Devices
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

type NullableListDevicesResponse struct {
	value *ListDevicesResponse
	isSet bool
}

func (v NullableListDevicesResponse) Get() *ListDevicesResponse {
	return v.value
}

func (v *NullableListDevicesResponse) Set(val *ListDevicesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListDevicesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListDevicesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDevicesResponse(val *ListDevicesResponse) *NullableListDevicesResponse {
	return &NullableListDevicesResponse{value: val, isSet: true}
}

func (v NullableListDevicesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDevicesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


