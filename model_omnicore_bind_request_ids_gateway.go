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

// checks if the OmnicoreBindRequestIdsGateway type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OmnicoreBindRequestIdsGateway{}

// OmnicoreBindRequestIdsGateway struct for OmnicoreBindRequestIdsGateway
type OmnicoreBindRequestIdsGateway struct {
	DeviceIds []string `json:"deviceIds"`
	GatewayId string `json:"gatewayId"`
}

// NewOmnicoreBindRequestIdsGateway instantiates a new OmnicoreBindRequestIdsGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOmnicoreBindRequestIdsGateway(deviceIds []string, gatewayId string) *OmnicoreBindRequestIdsGateway {
	this := OmnicoreBindRequestIdsGateway{}
	this.DeviceIds = deviceIds
	this.GatewayId = gatewayId
	return &this
}

// NewOmnicoreBindRequestIdsGatewayWithDefaults instantiates a new OmnicoreBindRequestIdsGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOmnicoreBindRequestIdsGatewayWithDefaults() *OmnicoreBindRequestIdsGateway {
	this := OmnicoreBindRequestIdsGateway{}
	return &this
}

// GetDeviceIds returns the DeviceIds field value
func (o *OmnicoreBindRequestIdsGateway) GetDeviceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DeviceIds
}

// GetDeviceIdsOk returns a tuple with the DeviceIds field value
// and a boolean to check if the value has been set.
func (o *OmnicoreBindRequestIdsGateway) GetDeviceIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceIds, true
}

// SetDeviceIds sets field value
func (o *OmnicoreBindRequestIdsGateway) SetDeviceIds(v []string) {
	o.DeviceIds = v
}

// GetGatewayId returns the GatewayId field value
func (o *OmnicoreBindRequestIdsGateway) GetGatewayId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value
// and a boolean to check if the value has been set.
func (o *OmnicoreBindRequestIdsGateway) GetGatewayIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayId, true
}

// SetGatewayId sets field value
func (o *OmnicoreBindRequestIdsGateway) SetGatewayId(v string) {
	o.GatewayId = v
}

func (o OmnicoreBindRequestIdsGateway) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OmnicoreBindRequestIdsGateway) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deviceIds"] = o.DeviceIds
	toSerialize["gatewayId"] = o.GatewayId
	return toSerialize, nil
}

type NullableOmnicoreBindRequestIdsGateway struct {
	value *OmnicoreBindRequestIdsGateway
	isSet bool
}

func (v NullableOmnicoreBindRequestIdsGateway) Get() *OmnicoreBindRequestIdsGateway {
	return v.value
}

func (v *NullableOmnicoreBindRequestIdsGateway) Set(val *OmnicoreBindRequestIdsGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableOmnicoreBindRequestIdsGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableOmnicoreBindRequestIdsGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOmnicoreBindRequestIdsGateway(val *OmnicoreBindRequestIdsGateway) *NullableOmnicoreBindRequestIdsGateway {
	return &NullableOmnicoreBindRequestIdsGateway{value: val, isSet: true}
}

func (v NullableOmnicoreBindRequestIdsGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOmnicoreBindRequestIdsGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


