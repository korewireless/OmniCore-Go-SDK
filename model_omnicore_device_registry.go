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

// checks if the OmnicoreDeviceRegistry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OmnicoreDeviceRegistry{}

// OmnicoreDeviceRegistry struct for OmnicoreDeviceRegistry
type OmnicoreDeviceRegistry struct {
	CreatedOn *string `json:"createdOn,omitempty"`
	Credentials []OmnicoreRegistryCredential `json:"credentials,omitempty"`
	EventNotificationConfigs []OmnicoreEventNotificationConfig `json:"eventNotificationConfigs,omitempty"`
	HttpConfig *OmnicoreHttpConfig `json:"httpConfig,omitempty"`
	Id string `json:"id"`
	LogLevel *string `json:"logLevel,omitempty"`
	LogNotificationConfig *OmnicoreStateNotificationConfig `json:"logNotificationConfig,omitempty"`
	MqttConfig *OmnicoreMqttConfig `json:"mqttConfig,omitempty"`
	Name *string `json:"name,omitempty"`
	NumberOfDevices *int32 `json:"numberOfDevices,omitempty"`
	NumberOfGateways *int32 `json:"numberOfGateways,omitempty"`
	Parent string `json:"parent"`
	StateNotificationConfig *OmnicoreStateNotificationConfig `json:"stateNotificationConfig,omitempty"`
	UpdatedOn *string `json:"updatedOn,omitempty"`
}

// NewOmnicoreDeviceRegistry instantiates a new OmnicoreDeviceRegistry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOmnicoreDeviceRegistry(id string, parent string) *OmnicoreDeviceRegistry {
	this := OmnicoreDeviceRegistry{}
	this.Id = id
	this.Parent = parent
	return &this
}

// NewOmnicoreDeviceRegistryWithDefaults instantiates a new OmnicoreDeviceRegistry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOmnicoreDeviceRegistryWithDefaults() *OmnicoreDeviceRegistry {
	this := OmnicoreDeviceRegistry{}
	return &this
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *OmnicoreDeviceRegistry) GetCreatedOn() string {
	if o == nil || IsNil(o.CreatedOn) {
		var ret string
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDeviceRegistry) GetCreatedOnOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedOn) {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *OmnicoreDeviceRegistry) HasCreatedOn() bool {
	if o != nil && !IsNil(o.CreatedOn) {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given string and assigns it to the CreatedOn field.
func (o *OmnicoreDeviceRegistry) SetCreatedOn(v string) {
	o.CreatedOn = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *OmnicoreDeviceRegistry) GetCredentials() []OmnicoreRegistryCredential {
	if o == nil || IsNil(o.Credentials) {
		var ret []OmnicoreRegistryCredential
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDeviceRegistry) GetCredentialsOk() ([]OmnicoreRegistryCredential, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *OmnicoreDeviceRegistry) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given []OmnicoreRegistryCredential and assigns it to the Credentials field.
func (o *OmnicoreDeviceRegistry) SetCredentials(v []OmnicoreRegistryCredential) {
	o.Credentials = v
}

// GetEventNotificationConfigs returns the EventNotificationConfigs field value if set, zero value otherwise.
func (o *OmnicoreDeviceRegistry) GetEventNotificationConfigs() []OmnicoreEventNotificationConfig {
	if o == nil || IsNil(o.EventNotificationConfigs) {
		var ret []OmnicoreEventNotificationConfig
		return ret
	}
	return o.EventNotificationConfigs
}

// GetEventNotificationConfigsOk returns a tuple with the EventNotificationConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDeviceRegistry) GetEventNotificationConfigsOk() ([]OmnicoreEventNotificationConfig, bool) {
	if o == nil || IsNil(o.EventNotificationConfigs) {
		return nil, false
	}
	return o.EventNotificationConfigs, true
}

// HasEventNotificationConfigs returns a boolean if a field has been set.
func (o *OmnicoreDeviceRegistry) HasEventNotificationConfigs() bool {
	if o != nil && !IsNil(o.EventNotificationConfigs) {
		return true
	}

	return false
}

// SetEventNotificationConfigs gets a reference to the given []OmnicoreEventNotificationConfig and assigns it to the EventNotificationConfigs field.
func (o *OmnicoreDeviceRegistry) SetEventNotificationConfigs(v []OmnicoreEventNotificationConfig) {
	o.EventNotificationConfigs = v
}

// GetHttpConfig returns the HttpConfig field value if set, zero value otherwise.
func (o *OmnicoreDeviceRegistry) GetHttpConfig() OmnicoreHttpConfig {
	if o == nil || IsNil(o.HttpConfig) {
		var ret OmnicoreHttpConfig
		return ret
	}
	return *o.HttpConfig
}

// GetHttpConfigOk returns a tuple with the HttpConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDeviceRegistry) GetHttpConfigOk() (*OmnicoreHttpConfig, bool) {
	if o == nil || IsNil(o.HttpConfig) {
		return nil, false
	}
	return o.HttpConfig, true
}

// HasHttpConfig returns a boolean if a field has been set.
func (o *OmnicoreDeviceRegistry) HasHttpConfig() bool {
	if o != nil && !IsNil(o.HttpConfig) {
		return true
	}

	return false
}

// SetHttpConfig gets a reference to the given OmnicoreHttpConfig and assigns it to the HttpConfig field.
func (o *OmnicoreDeviceRegistry) SetHttpConfig(v OmnicoreHttpConfig) {
	o.HttpConfig = &v
}

// GetId returns the Id field value
func (o *OmnicoreDeviceRegistry) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OmnicoreDeviceRegistry) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OmnicoreDeviceRegistry) SetId(v string) {
	o.Id = v
}

// GetLogLevel returns the LogLevel field value if set, zero value otherwise.
func (o *OmnicoreDeviceRegistry) GetLogLevel() string {
	if o == nil || IsNil(o.LogLevel) {
		var ret string
		return ret
	}
	return *o.LogLevel
}

// GetLogLevelOk returns a tuple with the LogLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDeviceRegistry) GetLogLevelOk() (*string, bool) {
	if o == nil || IsNil(o.LogLevel) {
		return nil, false
	}
	return o.LogLevel, true
}

// HasLogLevel returns a boolean if a field has been set.
func (o *OmnicoreDeviceRegistry) HasLogLevel() bool {
	if o != nil && !IsNil(o.LogLevel) {
		return true
	}

	return false
}

// SetLogLevel gets a reference to the given string and assigns it to the LogLevel field.
func (o *OmnicoreDeviceRegistry) SetLogLevel(v string) {
	o.LogLevel = &v
}

// GetLogNotificationConfig returns the LogNotificationConfig field value if set, zero value otherwise.
func (o *OmnicoreDeviceRegistry) GetLogNotificationConfig() OmnicoreStateNotificationConfig {
	if o == nil || IsNil(o.LogNotificationConfig) {
		var ret OmnicoreStateNotificationConfig
		return ret
	}
	return *o.LogNotificationConfig
}

// GetLogNotificationConfigOk returns a tuple with the LogNotificationConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDeviceRegistry) GetLogNotificationConfigOk() (*OmnicoreStateNotificationConfig, bool) {
	if o == nil || IsNil(o.LogNotificationConfig) {
		return nil, false
	}
	return o.LogNotificationConfig, true
}

// HasLogNotificationConfig returns a boolean if a field has been set.
func (o *OmnicoreDeviceRegistry) HasLogNotificationConfig() bool {
	if o != nil && !IsNil(o.LogNotificationConfig) {
		return true
	}

	return false
}

// SetLogNotificationConfig gets a reference to the given OmnicoreStateNotificationConfig and assigns it to the LogNotificationConfig field.
func (o *OmnicoreDeviceRegistry) SetLogNotificationConfig(v OmnicoreStateNotificationConfig) {
	o.LogNotificationConfig = &v
}

// GetMqttConfig returns the MqttConfig field value if set, zero value otherwise.
func (o *OmnicoreDeviceRegistry) GetMqttConfig() OmnicoreMqttConfig {
	if o == nil || IsNil(o.MqttConfig) {
		var ret OmnicoreMqttConfig
		return ret
	}
	return *o.MqttConfig
}

// GetMqttConfigOk returns a tuple with the MqttConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDeviceRegistry) GetMqttConfigOk() (*OmnicoreMqttConfig, bool) {
	if o == nil || IsNil(o.MqttConfig) {
		return nil, false
	}
	return o.MqttConfig, true
}

// HasMqttConfig returns a boolean if a field has been set.
func (o *OmnicoreDeviceRegistry) HasMqttConfig() bool {
	if o != nil && !IsNil(o.MqttConfig) {
		return true
	}

	return false
}

// SetMqttConfig gets a reference to the given OmnicoreMqttConfig and assigns it to the MqttConfig field.
func (o *OmnicoreDeviceRegistry) SetMqttConfig(v OmnicoreMqttConfig) {
	o.MqttConfig = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OmnicoreDeviceRegistry) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDeviceRegistry) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OmnicoreDeviceRegistry) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OmnicoreDeviceRegistry) SetName(v string) {
	o.Name = &v
}

// GetNumberOfDevices returns the NumberOfDevices field value if set, zero value otherwise.
func (o *OmnicoreDeviceRegistry) GetNumberOfDevices() int32 {
	if o == nil || IsNil(o.NumberOfDevices) {
		var ret int32
		return ret
	}
	return *o.NumberOfDevices
}

// GetNumberOfDevicesOk returns a tuple with the NumberOfDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDeviceRegistry) GetNumberOfDevicesOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfDevices) {
		return nil, false
	}
	return o.NumberOfDevices, true
}

// HasNumberOfDevices returns a boolean if a field has been set.
func (o *OmnicoreDeviceRegistry) HasNumberOfDevices() bool {
	if o != nil && !IsNil(o.NumberOfDevices) {
		return true
	}

	return false
}

// SetNumberOfDevices gets a reference to the given int32 and assigns it to the NumberOfDevices field.
func (o *OmnicoreDeviceRegistry) SetNumberOfDevices(v int32) {
	o.NumberOfDevices = &v
}

// GetNumberOfGateways returns the NumberOfGateways field value if set, zero value otherwise.
func (o *OmnicoreDeviceRegistry) GetNumberOfGateways() int32 {
	if o == nil || IsNil(o.NumberOfGateways) {
		var ret int32
		return ret
	}
	return *o.NumberOfGateways
}

// GetNumberOfGatewaysOk returns a tuple with the NumberOfGateways field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDeviceRegistry) GetNumberOfGatewaysOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfGateways) {
		return nil, false
	}
	return o.NumberOfGateways, true
}

// HasNumberOfGateways returns a boolean if a field has been set.
func (o *OmnicoreDeviceRegistry) HasNumberOfGateways() bool {
	if o != nil && !IsNil(o.NumberOfGateways) {
		return true
	}

	return false
}

// SetNumberOfGateways gets a reference to the given int32 and assigns it to the NumberOfGateways field.
func (o *OmnicoreDeviceRegistry) SetNumberOfGateways(v int32) {
	o.NumberOfGateways = &v
}

// GetParent returns the Parent field value
func (o *OmnicoreDeviceRegistry) GetParent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Parent
}

// GetParentOk returns a tuple with the Parent field value
// and a boolean to check if the value has been set.
func (o *OmnicoreDeviceRegistry) GetParentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parent, true
}

// SetParent sets field value
func (o *OmnicoreDeviceRegistry) SetParent(v string) {
	o.Parent = v
}

// GetStateNotificationConfig returns the StateNotificationConfig field value if set, zero value otherwise.
func (o *OmnicoreDeviceRegistry) GetStateNotificationConfig() OmnicoreStateNotificationConfig {
	if o == nil || IsNil(o.StateNotificationConfig) {
		var ret OmnicoreStateNotificationConfig
		return ret
	}
	return *o.StateNotificationConfig
}

// GetStateNotificationConfigOk returns a tuple with the StateNotificationConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDeviceRegistry) GetStateNotificationConfigOk() (*OmnicoreStateNotificationConfig, bool) {
	if o == nil || IsNil(o.StateNotificationConfig) {
		return nil, false
	}
	return o.StateNotificationConfig, true
}

// HasStateNotificationConfig returns a boolean if a field has been set.
func (o *OmnicoreDeviceRegistry) HasStateNotificationConfig() bool {
	if o != nil && !IsNil(o.StateNotificationConfig) {
		return true
	}

	return false
}

// SetStateNotificationConfig gets a reference to the given OmnicoreStateNotificationConfig and assigns it to the StateNotificationConfig field.
func (o *OmnicoreDeviceRegistry) SetStateNotificationConfig(v OmnicoreStateNotificationConfig) {
	o.StateNotificationConfig = &v
}

// GetUpdatedOn returns the UpdatedOn field value if set, zero value otherwise.
func (o *OmnicoreDeviceRegistry) GetUpdatedOn() string {
	if o == nil || IsNil(o.UpdatedOn) {
		var ret string
		return ret
	}
	return *o.UpdatedOn
}

// GetUpdatedOnOk returns a tuple with the UpdatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDeviceRegistry) GetUpdatedOnOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedOn) {
		return nil, false
	}
	return o.UpdatedOn, true
}

// HasUpdatedOn returns a boolean if a field has been set.
func (o *OmnicoreDeviceRegistry) HasUpdatedOn() bool {
	if o != nil && !IsNil(o.UpdatedOn) {
		return true
	}

	return false
}

// SetUpdatedOn gets a reference to the given string and assigns it to the UpdatedOn field.
func (o *OmnicoreDeviceRegistry) SetUpdatedOn(v string) {
	o.UpdatedOn = &v
}

func (o OmnicoreDeviceRegistry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OmnicoreDeviceRegistry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedOn) {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.EventNotificationConfigs) {
		toSerialize["eventNotificationConfigs"] = o.EventNotificationConfigs
	}
	if !IsNil(o.HttpConfig) {
		toSerialize["httpConfig"] = o.HttpConfig
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.LogLevel) {
		toSerialize["logLevel"] = o.LogLevel
	}
	if !IsNil(o.LogNotificationConfig) {
		toSerialize["logNotificationConfig"] = o.LogNotificationConfig
	}
	if !IsNil(o.MqttConfig) {
		toSerialize["mqttConfig"] = o.MqttConfig
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NumberOfDevices) {
		toSerialize["numberOfDevices"] = o.NumberOfDevices
	}
	if !IsNil(o.NumberOfGateways) {
		toSerialize["numberOfGateways"] = o.NumberOfGateways
	}
	toSerialize["parent"] = o.Parent
	if !IsNil(o.StateNotificationConfig) {
		toSerialize["stateNotificationConfig"] = o.StateNotificationConfig
	}
	if !IsNil(o.UpdatedOn) {
		toSerialize["updatedOn"] = o.UpdatedOn
	}
	return toSerialize, nil
}

type NullableOmnicoreDeviceRegistry struct {
	value *OmnicoreDeviceRegistry
	isSet bool
}

func (v NullableOmnicoreDeviceRegistry) Get() *OmnicoreDeviceRegistry {
	return v.value
}

func (v *NullableOmnicoreDeviceRegistry) Set(val *OmnicoreDeviceRegistry) {
	v.value = val
	v.isSet = true
}

func (v NullableOmnicoreDeviceRegistry) IsSet() bool {
	return v.isSet
}

func (v *NullableOmnicoreDeviceRegistry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOmnicoreDeviceRegistry(val *OmnicoreDeviceRegistry) *NullableOmnicoreDeviceRegistry {
	return &NullableOmnicoreDeviceRegistry{value: val, isSet: true}
}

func (v NullableOmnicoreDeviceRegistry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOmnicoreDeviceRegistry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


