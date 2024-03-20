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

// checks if the TcpUdpPortDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TcpUdpPortDetail{}

// TcpUdpPortDetail struct for TcpUdpPortDetail
type TcpUdpPortDetail struct {
	// Indicates if the port is enabled
	Enabled bool `json:"enabled"`
	// The port number
	Port *int32 `json:"port,omitempty"`
}

// NewTcpUdpPortDetail instantiates a new TcpUdpPortDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTcpUdpPortDetail(enabled bool) *TcpUdpPortDetail {
	this := TcpUdpPortDetail{}
	this.Enabled = enabled
	return &this
}

// NewTcpUdpPortDetailWithDefaults instantiates a new TcpUdpPortDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTcpUdpPortDetailWithDefaults() *TcpUdpPortDetail {
	this := TcpUdpPortDetail{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *TcpUdpPortDetail) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *TcpUdpPortDetail) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *TcpUdpPortDetail) SetEnabled(v bool) {
	o.Enabled = v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *TcpUdpPortDetail) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TcpUdpPortDetail) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *TcpUdpPortDetail) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *TcpUdpPortDetail) SetPort(v int32) {
	o.Port = &v
}

func (o TcpUdpPortDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TcpUdpPortDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	return toSerialize, nil
}

type NullableTcpUdpPortDetail struct {
	value *TcpUdpPortDetail
	isSet bool
}

func (v NullableTcpUdpPortDetail) Get() *TcpUdpPortDetail {
	return v.value
}

func (v *NullableTcpUdpPortDetail) Set(val *TcpUdpPortDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableTcpUdpPortDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableTcpUdpPortDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTcpUdpPortDetail(val *TcpUdpPortDetail) *NullableTcpUdpPortDetail {
	return &NullableTcpUdpPortDetail{value: val, isSet: true}
}

func (v NullableTcpUdpPortDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTcpUdpPortDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


