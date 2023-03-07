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

// checks if the BindRequestIdsGateway type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BindRequestIdsGateway{}

// BindRequestIdsGateway struct for BindRequestIdsGateway
type BindRequestIdsGateway struct {
	DeviceIds []string `json:"deviceIds"`
	GatewayId string `json:"gatewayId"`
}

// NewBindRequestIdsGateway instantiates a new BindRequestIdsGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBindRequestIdsGateway(deviceIds []string, gatewayId string) *BindRequestIdsGateway {
	this := BindRequestIdsGateway{}
	this.DeviceIds = deviceIds
	this.GatewayId = gatewayId
	return &this
}

// NewBindRequestIdsGatewayWithDefaults instantiates a new BindRequestIdsGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBindRequestIdsGatewayWithDefaults() *BindRequestIdsGateway {
	this := BindRequestIdsGateway{}
	return &this
}

// GetDeviceIds returns the DeviceIds field value
func (o *BindRequestIdsGateway) GetDeviceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DeviceIds
}

// GetDeviceIdsOk returns a tuple with the DeviceIds field value
// and a boolean to check if the value has been set.
func (o *BindRequestIdsGateway) GetDeviceIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceIds, true
}

// SetDeviceIds sets field value
func (o *BindRequestIdsGateway) SetDeviceIds(v []string) {
	o.DeviceIds = v
}

// GetGatewayId returns the GatewayId field value
func (o *BindRequestIdsGateway) GetGatewayId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value
// and a boolean to check if the value has been set.
func (o *BindRequestIdsGateway) GetGatewayIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayId, true
}

// SetGatewayId sets field value
func (o *BindRequestIdsGateway) SetGatewayId(v string) {
	o.GatewayId = v
}

func (o BindRequestIdsGateway) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BindRequestIdsGateway) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deviceIds"] = o.DeviceIds
	toSerialize["gatewayId"] = o.GatewayId
	return toSerialize, nil
}

type NullableBindRequestIdsGateway struct {
	value *BindRequestIdsGateway
	isSet bool
}

func (v NullableBindRequestIdsGateway) Get() *BindRequestIdsGateway {
	return v.value
}

func (v *NullableBindRequestIdsGateway) Set(val *BindRequestIdsGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableBindRequestIdsGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableBindRequestIdsGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBindRequestIdsGateway(val *BindRequestIdsGateway) *NullableBindRequestIdsGateway {
	return &NullableBindRequestIdsGateway{value: val, isSet: true}
}

func (v NullableBindRequestIdsGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBindRequestIdsGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


