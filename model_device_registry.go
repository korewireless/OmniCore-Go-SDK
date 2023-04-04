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

// checks if the DeviceRegistry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceRegistry{}

// DeviceRegistry struct for DeviceRegistry
type DeviceRegistry struct {
	Id string `json:"id"`
	Name *string `json:"name,omitempty"`
	Parent *string `json:"parent,omitempty"`
	CreatedOn *string `json:"createdOn,omitempty"`
	UpdatedOn *string `json:"updatedOn,omitempty"`
	Credentials []RegistryCredential `json:"credentials,omitempty"`
	HttpConfig *HttpConfig `json:"httpConfig,omitempty"`
	MqttConfig *MqttConfig `json:"mqttConfig,omitempty"`
	LogLevel *LogLevel `json:"logLevel,omitempty"`
	EventNotificationConfigs []EventNotificationConfig `json:"eventNotificationConfigs,omitempty"`
	LogNotificationConfig *NotificationConfig `json:"logNotificationConfig,omitempty"`
	StateNotificationConfig *NotificationConfig `json:"stateNotificationConfig,omitempty"`
	JitrNotificationConfig *NotificationConfig `json:"jitrNotificationConfig,omitempty"`
	NumberOfDevices *int32 `json:"numberOfDevices,omitempty"`
	NumberOfGateways *int32 `json:"numberOfGateways,omitempty"`
}

// NewDeviceRegistry instantiates a new DeviceRegistry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceRegistry(id string) *DeviceRegistry {
	this := DeviceRegistry{}
	this.Id = id
	return &this
}

// NewDeviceRegistryWithDefaults instantiates a new DeviceRegistry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceRegistryWithDefaults() *DeviceRegistry {
	this := DeviceRegistry{}
	return &this
}

// GetId returns the Id field value
func (o *DeviceRegistry) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeviceRegistry) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeviceRegistry) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeviceRegistry) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRegistry) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeviceRegistry) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeviceRegistry) SetName(v string) {
	o.Name = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *DeviceRegistry) GetParent() string {
	if o == nil || IsNil(o.Parent) {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRegistry) GetParentOk() (*string, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *DeviceRegistry) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *DeviceRegistry) SetParent(v string) {
	o.Parent = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *DeviceRegistry) GetCreatedOn() string {
	if o == nil || IsNil(o.CreatedOn) {
		var ret string
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRegistry) GetCreatedOnOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedOn) {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *DeviceRegistry) HasCreatedOn() bool {
	if o != nil && !IsNil(o.CreatedOn) {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given string and assigns it to the CreatedOn field.
func (o *DeviceRegistry) SetCreatedOn(v string) {
	o.CreatedOn = &v
}

// GetUpdatedOn returns the UpdatedOn field value if set, zero value otherwise.
func (o *DeviceRegistry) GetUpdatedOn() string {
	if o == nil || IsNil(o.UpdatedOn) {
		var ret string
		return ret
	}
	return *o.UpdatedOn
}

// GetUpdatedOnOk returns a tuple with the UpdatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRegistry) GetUpdatedOnOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedOn) {
		return nil, false
	}
	return o.UpdatedOn, true
}

// HasUpdatedOn returns a boolean if a field has been set.
func (o *DeviceRegistry) HasUpdatedOn() bool {
	if o != nil && !IsNil(o.UpdatedOn) {
		return true
	}

	return false
}

// SetUpdatedOn gets a reference to the given string and assigns it to the UpdatedOn field.
func (o *DeviceRegistry) SetUpdatedOn(v string) {
	o.UpdatedOn = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *DeviceRegistry) GetCredentials() []RegistryCredential {
	if o == nil || IsNil(o.Credentials) {
		var ret []RegistryCredential
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRegistry) GetCredentialsOk() ([]RegistryCredential, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *DeviceRegistry) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given []RegistryCredential and assigns it to the Credentials field.
func (o *DeviceRegistry) SetCredentials(v []RegistryCredential) {
	o.Credentials = v
}

// GetHttpConfig returns the HttpConfig field value if set, zero value otherwise.
func (o *DeviceRegistry) GetHttpConfig() HttpConfig {
	if o == nil || IsNil(o.HttpConfig) {
		var ret HttpConfig
		return ret
	}
	return *o.HttpConfig
}

// GetHttpConfigOk returns a tuple with the HttpConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRegistry) GetHttpConfigOk() (*HttpConfig, bool) {
	if o == nil || IsNil(o.HttpConfig) {
		return nil, false
	}
	return o.HttpConfig, true
}

// HasHttpConfig returns a boolean if a field has been set.
func (o *DeviceRegistry) HasHttpConfig() bool {
	if o != nil && !IsNil(o.HttpConfig) {
		return true
	}

	return false
}

// SetHttpConfig gets a reference to the given HttpConfig and assigns it to the HttpConfig field.
func (o *DeviceRegistry) SetHttpConfig(v HttpConfig) {
	o.HttpConfig = &v
}

// GetMqttConfig returns the MqttConfig field value if set, zero value otherwise.
func (o *DeviceRegistry) GetMqttConfig() MqttConfig {
	if o == nil || IsNil(o.MqttConfig) {
		var ret MqttConfig
		return ret
	}
	return *o.MqttConfig
}

// GetMqttConfigOk returns a tuple with the MqttConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRegistry) GetMqttConfigOk() (*MqttConfig, bool) {
	if o == nil || IsNil(o.MqttConfig) {
		return nil, false
	}
	return o.MqttConfig, true
}

// HasMqttConfig returns a boolean if a field has been set.
func (o *DeviceRegistry) HasMqttConfig() bool {
	if o != nil && !IsNil(o.MqttConfig) {
		return true
	}

	return false
}

// SetMqttConfig gets a reference to the given MqttConfig and assigns it to the MqttConfig field.
func (o *DeviceRegistry) SetMqttConfig(v MqttConfig) {
	o.MqttConfig = &v
}

// GetLogLevel returns the LogLevel field value if set, zero value otherwise.
func (o *DeviceRegistry) GetLogLevel() LogLevel {
	if o == nil || IsNil(o.LogLevel) {
		var ret LogLevel
		return ret
	}
	return *o.LogLevel
}

// GetLogLevelOk returns a tuple with the LogLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRegistry) GetLogLevelOk() (*LogLevel, bool) {
	if o == nil || IsNil(o.LogLevel) {
		return nil, false
	}
	return o.LogLevel, true
}

// HasLogLevel returns a boolean if a field has been set.
func (o *DeviceRegistry) HasLogLevel() bool {
	if o != nil && !IsNil(o.LogLevel) {
		return true
	}

	return false
}

// SetLogLevel gets a reference to the given LogLevel and assigns it to the LogLevel field.
func (o *DeviceRegistry) SetLogLevel(v LogLevel) {
	o.LogLevel = &v
}

// GetEventNotificationConfigs returns the EventNotificationConfigs field value if set, zero value otherwise.
func (o *DeviceRegistry) GetEventNotificationConfigs() []EventNotificationConfig {
	if o == nil || IsNil(o.EventNotificationConfigs) {
		var ret []EventNotificationConfig
		return ret
	}
	return o.EventNotificationConfigs
}

// GetEventNotificationConfigsOk returns a tuple with the EventNotificationConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRegistry) GetEventNotificationConfigsOk() ([]EventNotificationConfig, bool) {
	if o == nil || IsNil(o.EventNotificationConfigs) {
		return nil, false
	}
	return o.EventNotificationConfigs, true
}

// HasEventNotificationConfigs returns a boolean if a field has been set.
func (o *DeviceRegistry) HasEventNotificationConfigs() bool {
	if o != nil && !IsNil(o.EventNotificationConfigs) {
		return true
	}

	return false
}

// SetEventNotificationConfigs gets a reference to the given []EventNotificationConfig and assigns it to the EventNotificationConfigs field.
func (o *DeviceRegistry) SetEventNotificationConfigs(v []EventNotificationConfig) {
	o.EventNotificationConfigs = v
}

// GetLogNotificationConfig returns the LogNotificationConfig field value if set, zero value otherwise.
func (o *DeviceRegistry) GetLogNotificationConfig() NotificationConfig {
	if o == nil || IsNil(o.LogNotificationConfig) {
		var ret NotificationConfig
		return ret
	}
	return *o.LogNotificationConfig
}

// GetLogNotificationConfigOk returns a tuple with the LogNotificationConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRegistry) GetLogNotificationConfigOk() (*NotificationConfig, bool) {
	if o == nil || IsNil(o.LogNotificationConfig) {
		return nil, false
	}
	return o.LogNotificationConfig, true
}

// HasLogNotificationConfig returns a boolean if a field has been set.
func (o *DeviceRegistry) HasLogNotificationConfig() bool {
	if o != nil && !IsNil(o.LogNotificationConfig) {
		return true
	}

	return false
}

// SetLogNotificationConfig gets a reference to the given NotificationConfig and assigns it to the LogNotificationConfig field.
func (o *DeviceRegistry) SetLogNotificationConfig(v NotificationConfig) {
	o.LogNotificationConfig = &v
}

// GetStateNotificationConfig returns the StateNotificationConfig field value if set, zero value otherwise.
func (o *DeviceRegistry) GetStateNotificationConfig() NotificationConfig {
	if o == nil || IsNil(o.StateNotificationConfig) {
		var ret NotificationConfig
		return ret
	}
	return *o.StateNotificationConfig
}

// GetStateNotificationConfigOk returns a tuple with the StateNotificationConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRegistry) GetStateNotificationConfigOk() (*NotificationConfig, bool) {
	if o == nil || IsNil(o.StateNotificationConfig) {
		return nil, false
	}
	return o.StateNotificationConfig, true
}

// HasStateNotificationConfig returns a boolean if a field has been set.
func (o *DeviceRegistry) HasStateNotificationConfig() bool {
	if o != nil && !IsNil(o.StateNotificationConfig) {
		return true
	}

	return false
}

// SetStateNotificationConfig gets a reference to the given NotificationConfig and assigns it to the StateNotificationConfig field.
func (o *DeviceRegistry) SetStateNotificationConfig(v NotificationConfig) {
	o.StateNotificationConfig = &v
}

// GetJitrNotificationConfig returns the JitrNotificationConfig field value if set, zero value otherwise.
func (o *DeviceRegistry) GetJitrNotificationConfig() NotificationConfig {
	if o == nil || IsNil(o.JitrNotificationConfig) {
		var ret NotificationConfig
		return ret
	}
	return *o.JitrNotificationConfig
}

// GetJitrNotificationConfigOk returns a tuple with the JitrNotificationConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRegistry) GetJitrNotificationConfigOk() (*NotificationConfig, bool) {
	if o == nil || IsNil(o.JitrNotificationConfig) {
		return nil, false
	}
	return o.JitrNotificationConfig, true
}

// HasJitrNotificationConfig returns a boolean if a field has been set.
func (o *DeviceRegistry) HasJitrNotificationConfig() bool {
	if o != nil && !IsNil(o.JitrNotificationConfig) {
		return true
	}

	return false
}

// SetJitrNotificationConfig gets a reference to the given NotificationConfig and assigns it to the JitrNotificationConfig field.
func (o *DeviceRegistry) SetJitrNotificationConfig(v NotificationConfig) {
	o.JitrNotificationConfig = &v
}

// GetNumberOfDevices returns the NumberOfDevices field value if set, zero value otherwise.
func (o *DeviceRegistry) GetNumberOfDevices() int32 {
	if o == nil || IsNil(o.NumberOfDevices) {
		var ret int32
		return ret
	}
	return *o.NumberOfDevices
}

// GetNumberOfDevicesOk returns a tuple with the NumberOfDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRegistry) GetNumberOfDevicesOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfDevices) {
		return nil, false
	}
	return o.NumberOfDevices, true
}

// HasNumberOfDevices returns a boolean if a field has been set.
func (o *DeviceRegistry) HasNumberOfDevices() bool {
	if o != nil && !IsNil(o.NumberOfDevices) {
		return true
	}

	return false
}

// SetNumberOfDevices gets a reference to the given int32 and assigns it to the NumberOfDevices field.
func (o *DeviceRegistry) SetNumberOfDevices(v int32) {
	o.NumberOfDevices = &v
}

// GetNumberOfGateways returns the NumberOfGateways field value if set, zero value otherwise.
func (o *DeviceRegistry) GetNumberOfGateways() int32 {
	if o == nil || IsNil(o.NumberOfGateways) {
		var ret int32
		return ret
	}
	return *o.NumberOfGateways
}

// GetNumberOfGatewaysOk returns a tuple with the NumberOfGateways field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRegistry) GetNumberOfGatewaysOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfGateways) {
		return nil, false
	}
	return o.NumberOfGateways, true
}

// HasNumberOfGateways returns a boolean if a field has been set.
func (o *DeviceRegistry) HasNumberOfGateways() bool {
	if o != nil && !IsNil(o.NumberOfGateways) {
		return true
	}

	return false
}

// SetNumberOfGateways gets a reference to the given int32 and assigns it to the NumberOfGateways field.
func (o *DeviceRegistry) SetNumberOfGateways(v int32) {
	o.NumberOfGateways = &v
}

func (o DeviceRegistry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceRegistry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	// skip: name is readOnly
	// skip: parent is readOnly
	// skip: createdOn is readOnly
	// skip: updatedOn is readOnly
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.HttpConfig) {
		toSerialize["httpConfig"] = o.HttpConfig
	}
	if !IsNil(o.MqttConfig) {
		toSerialize["mqttConfig"] = o.MqttConfig
	}
	if !IsNil(o.LogLevel) {
		toSerialize["logLevel"] = o.LogLevel
	}
	if !IsNil(o.EventNotificationConfigs) {
		toSerialize["eventNotificationConfigs"] = o.EventNotificationConfigs
	}
	if !IsNil(o.LogNotificationConfig) {
		toSerialize["logNotificationConfig"] = o.LogNotificationConfig
	}
	if !IsNil(o.StateNotificationConfig) {
		toSerialize["stateNotificationConfig"] = o.StateNotificationConfig
	}
	if !IsNil(o.JitrNotificationConfig) {
		toSerialize["jitrNotificationConfig"] = o.JitrNotificationConfig
	}
	// skip: numberOfDevices is readOnly
	// skip: numberOfGateways is readOnly
	return toSerialize, nil
}

type NullableDeviceRegistry struct {
	value *DeviceRegistry
	isSet bool
}

func (v NullableDeviceRegistry) Get() *DeviceRegistry {
	return v.value
}

func (v *NullableDeviceRegistry) Set(val *DeviceRegistry) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceRegistry) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceRegistry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceRegistry(val *DeviceRegistry) *NullableDeviceRegistry {
	return &NullableDeviceRegistry{value: val, isSet: true}
}

func (v NullableDeviceRegistry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceRegistry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


