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

// checks if the OmnicoreNewDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OmnicoreNewDevice{}

// OmnicoreNewDevice struct for OmnicoreNewDevice
type OmnicoreNewDevice struct {
	Blocked *bool `json:"blocked,omitempty"`
	Config *OmnicoreDeviceConfig `json:"config,omitempty"`
	Credentials []OmnicoreDeviceCredential `json:"credentials,omitempty"`
	GatewayConfig *OmnicoreGatewayConfig `json:"gatewayConfig,omitempty"`
	Id string `json:"id"`
	LogLevel *string `json:"logLevel,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
}

// NewOmnicoreNewDevice instantiates a new OmnicoreNewDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOmnicoreNewDevice(id string) *OmnicoreNewDevice {
	this := OmnicoreNewDevice{}
	this.Id = id
	return &this
}

// NewOmnicoreNewDeviceWithDefaults instantiates a new OmnicoreNewDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOmnicoreNewDeviceWithDefaults() *OmnicoreNewDevice {
	this := OmnicoreNewDevice{}
	return &this
}

// GetBlocked returns the Blocked field value if set, zero value otherwise.
func (o *OmnicoreNewDevice) GetBlocked() bool {
	if o == nil || IsNil(o.Blocked) {
		var ret bool
		return ret
	}
	return *o.Blocked
}

// GetBlockedOk returns a tuple with the Blocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreNewDevice) GetBlockedOk() (*bool, bool) {
	if o == nil || IsNil(o.Blocked) {
		return nil, false
	}
	return o.Blocked, true
}

// HasBlocked returns a boolean if a field has been set.
func (o *OmnicoreNewDevice) HasBlocked() bool {
	if o != nil && !IsNil(o.Blocked) {
		return true
	}

	return false
}

// SetBlocked gets a reference to the given bool and assigns it to the Blocked field.
func (o *OmnicoreNewDevice) SetBlocked(v bool) {
	o.Blocked = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *OmnicoreNewDevice) GetConfig() OmnicoreDeviceConfig {
	if o == nil || IsNil(o.Config) {
		var ret OmnicoreDeviceConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreNewDevice) GetConfigOk() (*OmnicoreDeviceConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *OmnicoreNewDevice) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given OmnicoreDeviceConfig and assigns it to the Config field.
func (o *OmnicoreNewDevice) SetConfig(v OmnicoreDeviceConfig) {
	o.Config = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *OmnicoreNewDevice) GetCredentials() []OmnicoreDeviceCredential {
	if o == nil || IsNil(o.Credentials) {
		var ret []OmnicoreDeviceCredential
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreNewDevice) GetCredentialsOk() ([]OmnicoreDeviceCredential, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *OmnicoreNewDevice) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given []OmnicoreDeviceCredential and assigns it to the Credentials field.
func (o *OmnicoreNewDevice) SetCredentials(v []OmnicoreDeviceCredential) {
	o.Credentials = v
}

// GetGatewayConfig returns the GatewayConfig field value if set, zero value otherwise.
func (o *OmnicoreNewDevice) GetGatewayConfig() OmnicoreGatewayConfig {
	if o == nil || IsNil(o.GatewayConfig) {
		var ret OmnicoreGatewayConfig
		return ret
	}
	return *o.GatewayConfig
}

// GetGatewayConfigOk returns a tuple with the GatewayConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreNewDevice) GetGatewayConfigOk() (*OmnicoreGatewayConfig, bool) {
	if o == nil || IsNil(o.GatewayConfig) {
		return nil, false
	}
	return o.GatewayConfig, true
}

// HasGatewayConfig returns a boolean if a field has been set.
func (o *OmnicoreNewDevice) HasGatewayConfig() bool {
	if o != nil && !IsNil(o.GatewayConfig) {
		return true
	}

	return false
}

// SetGatewayConfig gets a reference to the given OmnicoreGatewayConfig and assigns it to the GatewayConfig field.
func (o *OmnicoreNewDevice) SetGatewayConfig(v OmnicoreGatewayConfig) {
	o.GatewayConfig = &v
}

// GetId returns the Id field value
func (o *OmnicoreNewDevice) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OmnicoreNewDevice) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OmnicoreNewDevice) SetId(v string) {
	o.Id = v
}

// GetLogLevel returns the LogLevel field value if set, zero value otherwise.
func (o *OmnicoreNewDevice) GetLogLevel() string {
	if o == nil || IsNil(o.LogLevel) {
		var ret string
		return ret
	}
	return *o.LogLevel
}

// GetLogLevelOk returns a tuple with the LogLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreNewDevice) GetLogLevelOk() (*string, bool) {
	if o == nil || IsNil(o.LogLevel) {
		return nil, false
	}
	return o.LogLevel, true
}

// HasLogLevel returns a boolean if a field has been set.
func (o *OmnicoreNewDevice) HasLogLevel() bool {
	if o != nil && !IsNil(o.LogLevel) {
		return true
	}

	return false
}

// SetLogLevel gets a reference to the given string and assigns it to the LogLevel field.
func (o *OmnicoreNewDevice) SetLogLevel(v string) {
	o.LogLevel = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *OmnicoreNewDevice) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreNewDevice) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *OmnicoreNewDevice) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *OmnicoreNewDevice) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

func (o OmnicoreNewDevice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OmnicoreNewDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Blocked) {
		toSerialize["blocked"] = o.Blocked
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.GatewayConfig) {
		toSerialize["gatewayConfig"] = o.GatewayConfig
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.LogLevel) {
		toSerialize["logLevel"] = o.LogLevel
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableOmnicoreNewDevice struct {
	value *OmnicoreNewDevice
	isSet bool
}

func (v NullableOmnicoreNewDevice) Get() *OmnicoreNewDevice {
	return v.value
}

func (v *NullableOmnicoreNewDevice) Set(val *OmnicoreNewDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableOmnicoreNewDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableOmnicoreNewDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOmnicoreNewDevice(val *OmnicoreNewDevice) *NullableOmnicoreNewDevice {
	return &NullableOmnicoreNewDevice{value: val, isSet: true}
}

func (v NullableOmnicoreNewDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOmnicoreNewDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


