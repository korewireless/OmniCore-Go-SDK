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

// checks if the BindRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BindRequest{}

// BindRequest struct for BindRequest
type BindRequest struct {
	DeviceId string `json:"deviceId"`
	GatewayId string `json:"gatewayId"`
}

// NewBindRequest instantiates a new BindRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBindRequest(deviceId string, gatewayId string) *BindRequest {
	this := BindRequest{}
	this.DeviceId = deviceId
	this.GatewayId = gatewayId
	return &this
}

// NewBindRequestWithDefaults instantiates a new BindRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBindRequestWithDefaults() *BindRequest {
	this := BindRequest{}
	return &this
}

// GetDeviceId returns the DeviceId field value
func (o *BindRequest) GetDeviceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value
// and a boolean to check if the value has been set.
func (o *BindRequest) GetDeviceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceId, true
}

// SetDeviceId sets field value
func (o *BindRequest) SetDeviceId(v string) {
	o.DeviceId = v
}

// GetGatewayId returns the GatewayId field value
func (o *BindRequest) GetGatewayId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value
// and a boolean to check if the value has been set.
func (o *BindRequest) GetGatewayIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayId, true
}

// SetGatewayId sets field value
func (o *BindRequest) SetGatewayId(v string) {
	o.GatewayId = v
}

func (o BindRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BindRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deviceId"] = o.DeviceId
	toSerialize["gatewayId"] = o.GatewayId
	return toSerialize, nil
}

type NullableBindRequest struct {
	value *BindRequest
	isSet bool
}

func (v NullableBindRequest) Get() *BindRequest {
	return v.value
}

func (v *NullableBindRequest) Set(val *BindRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBindRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBindRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBindRequest(val *BindRequest) *NullableBindRequest {
	return &NullableBindRequest{value: val, isSet: true}
}

func (v NullableBindRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBindRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


