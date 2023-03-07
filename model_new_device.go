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

// checks if the NewDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewDevice{}

// NewDevice struct for NewDevice
type NewDevice struct {
	Id string `json:"id"`
	Blocked *bool `json:"blocked,omitempty"`
	Credentials []DeviceCredential `json:"credentials,omitempty"`
	GatewayConfig *GatewayConfig `json:"gatewayConfig,omitempty"`
	Config *DeviceConfig `json:"config,omitempty"`
	LogLevel *LogLevel `json:"logLevel,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
}

// NewNewDevice instantiates a new NewDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewDevice(id string) *NewDevice {
	this := NewDevice{}
	this.Id = id
	return &this
}

// NewNewDeviceWithDefaults instantiates a new NewDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewDeviceWithDefaults() *NewDevice {
	this := NewDevice{}
	return &this
}

// GetId returns the Id field value
func (o *NewDevice) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NewDevice) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NewDevice) SetId(v string) {
	o.Id = v
}

// GetBlocked returns the Blocked field value if set, zero value otherwise.
func (o *NewDevice) GetBlocked() bool {
	if o == nil || IsNil(o.Blocked) {
		var ret bool
		return ret
	}
	return *o.Blocked
}

// GetBlockedOk returns a tuple with the Blocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewDevice) GetBlockedOk() (*bool, bool) {
	if o == nil || IsNil(o.Blocked) {
		return nil, false
	}
	return o.Blocked, true
}

// HasBlocked returns a boolean if a field has been set.
func (o *NewDevice) HasBlocked() bool {
	if o != nil && !IsNil(o.Blocked) {
		return true
	}

	return false
}

// SetBlocked gets a reference to the given bool and assigns it to the Blocked field.
func (o *NewDevice) SetBlocked(v bool) {
	o.Blocked = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *NewDevice) GetCredentials() []DeviceCredential {
	if o == nil || IsNil(o.Credentials) {
		var ret []DeviceCredential
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewDevice) GetCredentialsOk() ([]DeviceCredential, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *NewDevice) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given []DeviceCredential and assigns it to the Credentials field.
func (o *NewDevice) SetCredentials(v []DeviceCredential) {
	o.Credentials = v
}

// GetGatewayConfig returns the GatewayConfig field value if set, zero value otherwise.
func (o *NewDevice) GetGatewayConfig() GatewayConfig {
	if o == nil || IsNil(o.GatewayConfig) {
		var ret GatewayConfig
		return ret
	}
	return *o.GatewayConfig
}

// GetGatewayConfigOk returns a tuple with the GatewayConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewDevice) GetGatewayConfigOk() (*GatewayConfig, bool) {
	if o == nil || IsNil(o.GatewayConfig) {
		return nil, false
	}
	return o.GatewayConfig, true
}

// HasGatewayConfig returns a boolean if a field has been set.
func (o *NewDevice) HasGatewayConfig() bool {
	if o != nil && !IsNil(o.GatewayConfig) {
		return true
	}

	return false
}

// SetGatewayConfig gets a reference to the given GatewayConfig and assigns it to the GatewayConfig field.
func (o *NewDevice) SetGatewayConfig(v GatewayConfig) {
	o.GatewayConfig = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *NewDevice) GetConfig() DeviceConfig {
	if o == nil || IsNil(o.Config) {
		var ret DeviceConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewDevice) GetConfigOk() (*DeviceConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *NewDevice) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given DeviceConfig and assigns it to the Config field.
func (o *NewDevice) SetConfig(v DeviceConfig) {
	o.Config = &v
}

// GetLogLevel returns the LogLevel field value if set, zero value otherwise.
func (o *NewDevice) GetLogLevel() LogLevel {
	if o == nil || IsNil(o.LogLevel) {
		var ret LogLevel
		return ret
	}
	return *o.LogLevel
}

// GetLogLevelOk returns a tuple with the LogLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewDevice) GetLogLevelOk() (*LogLevel, bool) {
	if o == nil || IsNil(o.LogLevel) {
		return nil, false
	}
	return o.LogLevel, true
}

// HasLogLevel returns a boolean if a field has been set.
func (o *NewDevice) HasLogLevel() bool {
	if o != nil && !IsNil(o.LogLevel) {
		return true
	}

	return false
}

// SetLogLevel gets a reference to the given LogLevel and assigns it to the LogLevel field.
func (o *NewDevice) SetLogLevel(v LogLevel) {
	o.LogLevel = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *NewDevice) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewDevice) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *NewDevice) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *NewDevice) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

func (o NewDevice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Blocked) {
		toSerialize["blocked"] = o.Blocked
	}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.GatewayConfig) {
		toSerialize["gatewayConfig"] = o.GatewayConfig
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.LogLevel) {
		toSerialize["logLevel"] = o.LogLevel
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableNewDevice struct {
	value *NewDevice
	isSet bool
}

func (v NullableNewDevice) Get() *NewDevice {
	return v.value
}

func (v *NullableNewDevice) Set(val *NewDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableNewDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableNewDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewDevice(val *NewDevice) *NullableNewDevice {
	return &NullableNewDevice{value: val, isSet: true}
}

func (v NullableNewDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


