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

// checks if the CreateRegistry200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRegistry200Response{}

// CreateRegistry200Response struct for CreateRegistry200Response
type CreateRegistry200Response struct {
	Id string `json:"id"`
	Name *string `json:"name,omitempty"`
	Parent string `json:"parent"`
	CreatedOn *string `json:"createdOn,omitempty"`
	UpdatedOn *string `json:"updatedOn,omitempty"`
	Credentials []RegistryCredential `json:"credentials,omitempty"`
	HttpConfig *HttpConfig `json:"httpConfig,omitempty"`
	MqttConfig *MqttConfig `json:"mqttConfig,omitempty"`
	LogLevel *LogLevel `json:"logLevel,omitempty"`
	EventNotificationConfigs []EventNotificationConfig `json:"eventNotificationConfigs,omitempty"`
	LogNotificationConfig *NotificationConfig `json:"logNotificationConfig,omitempty"`
	StateNotificationConfig *NotificationConfig `json:"stateNotificationConfig,omitempty"`
	NumberOfDevices *int32 `json:"numberOfDevices,omitempty"`
	NumberOfGateways *int32 `json:"numberOfGateways,omitempty"`
}

// NewCreateRegistry200Response instantiates a new CreateRegistry200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRegistry200Response(id string, parent string) *CreateRegistry200Response {
	this := CreateRegistry200Response{}
	this.Id = id
	this.Parent = parent
	return &this
}

// NewCreateRegistry200ResponseWithDefaults instantiates a new CreateRegistry200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRegistry200ResponseWithDefaults() *CreateRegistry200Response {
	this := CreateRegistry200Response{}
	return &this
}

// GetId returns the Id field value
func (o *CreateRegistry200Response) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateRegistry200Response) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateRegistry200Response) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateRegistry200Response) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRegistry200Response) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateRegistry200Response) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateRegistry200Response) SetName(v string) {
	o.Name = &v
}

// GetParent returns the Parent field value
func (o *CreateRegistry200Response) GetParent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Parent
}

// GetParentOk returns a tuple with the Parent field value
// and a boolean to check if the value has been set.
func (o *CreateRegistry200Response) GetParentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parent, true
}

// SetParent sets field value
func (o *CreateRegistry200Response) SetParent(v string) {
	o.Parent = v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *CreateRegistry200Response) GetCreatedOn() string {
	if o == nil || IsNil(o.CreatedOn) {
		var ret string
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRegistry200Response) GetCreatedOnOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedOn) {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *CreateRegistry200Response) HasCreatedOn() bool {
	if o != nil && !IsNil(o.CreatedOn) {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given string and assigns it to the CreatedOn field.
func (o *CreateRegistry200Response) SetCreatedOn(v string) {
	o.CreatedOn = &v
}

// GetUpdatedOn returns the UpdatedOn field value if set, zero value otherwise.
func (o *CreateRegistry200Response) GetUpdatedOn() string {
	if o == nil || IsNil(o.UpdatedOn) {
		var ret string
		return ret
	}
	return *o.UpdatedOn
}

// GetUpdatedOnOk returns a tuple with the UpdatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRegistry200Response) GetUpdatedOnOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedOn) {
		return nil, false
	}
	return o.UpdatedOn, true
}

// HasUpdatedOn returns a boolean if a field has been set.
func (o *CreateRegistry200Response) HasUpdatedOn() bool {
	if o != nil && !IsNil(o.UpdatedOn) {
		return true
	}

	return false
}

// SetUpdatedOn gets a reference to the given string and assigns it to the UpdatedOn field.
func (o *CreateRegistry200Response) SetUpdatedOn(v string) {
	o.UpdatedOn = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *CreateRegistry200Response) GetCredentials() []RegistryCredential {
	if o == nil || IsNil(o.Credentials) {
		var ret []RegistryCredential
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRegistry200Response) GetCredentialsOk() ([]RegistryCredential, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *CreateRegistry200Response) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given []RegistryCredential and assigns it to the Credentials field.
func (o *CreateRegistry200Response) SetCredentials(v []RegistryCredential) {
	o.Credentials = v
}

// GetHttpConfig returns the HttpConfig field value if set, zero value otherwise.
func (o *CreateRegistry200Response) GetHttpConfig() HttpConfig {
	if o == nil || IsNil(o.HttpConfig) {
		var ret HttpConfig
		return ret
	}
	return *o.HttpConfig
}

// GetHttpConfigOk returns a tuple with the HttpConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRegistry200Response) GetHttpConfigOk() (*HttpConfig, bool) {
	if o == nil || IsNil(o.HttpConfig) {
		return nil, false
	}
	return o.HttpConfig, true
}

// HasHttpConfig returns a boolean if a field has been set.
func (o *CreateRegistry200Response) HasHttpConfig() bool {
	if o != nil && !IsNil(o.HttpConfig) {
		return true
	}

	return false
}

// SetHttpConfig gets a reference to the given HttpConfig and assigns it to the HttpConfig field.
func (o *CreateRegistry200Response) SetHttpConfig(v HttpConfig) {
	o.HttpConfig = &v
}

// GetMqttConfig returns the MqttConfig field value if set, zero value otherwise.
func (o *CreateRegistry200Response) GetMqttConfig() MqttConfig {
	if o == nil || IsNil(o.MqttConfig) {
		var ret MqttConfig
		return ret
	}
	return *o.MqttConfig
}

// GetMqttConfigOk returns a tuple with the MqttConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRegistry200Response) GetMqttConfigOk() (*MqttConfig, bool) {
	if o == nil || IsNil(o.MqttConfig) {
		return nil, false
	}
	return o.MqttConfig, true
}

// HasMqttConfig returns a boolean if a field has been set.
func (o *CreateRegistry200Response) HasMqttConfig() bool {
	if o != nil && !IsNil(o.MqttConfig) {
		return true
	}

	return false
}

// SetMqttConfig gets a reference to the given MqttConfig and assigns it to the MqttConfig field.
func (o *CreateRegistry200Response) SetMqttConfig(v MqttConfig) {
	o.MqttConfig = &v
}

// GetLogLevel returns the LogLevel field value if set, zero value otherwise.
func (o *CreateRegistry200Response) GetLogLevel() LogLevel {
	if o == nil || IsNil(o.LogLevel) {
		var ret LogLevel
		return ret
	}
	return *o.LogLevel
}

// GetLogLevelOk returns a tuple with the LogLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRegistry200Response) GetLogLevelOk() (*LogLevel, bool) {
	if o == nil || IsNil(o.LogLevel) {
		return nil, false
	}
	return o.LogLevel, true
}

// HasLogLevel returns a boolean if a field has been set.
func (o *CreateRegistry200Response) HasLogLevel() bool {
	if o != nil && !IsNil(o.LogLevel) {
		return true
	}

	return false
}

// SetLogLevel gets a reference to the given LogLevel and assigns it to the LogLevel field.
func (o *CreateRegistry200Response) SetLogLevel(v LogLevel) {
	o.LogLevel = &v
}

// GetEventNotificationConfigs returns the EventNotificationConfigs field value if set, zero value otherwise.
func (o *CreateRegistry200Response) GetEventNotificationConfigs() []EventNotificationConfig {
	if o == nil || IsNil(o.EventNotificationConfigs) {
		var ret []EventNotificationConfig
		return ret
	}
	return o.EventNotificationConfigs
}

// GetEventNotificationConfigsOk returns a tuple with the EventNotificationConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRegistry200Response) GetEventNotificationConfigsOk() ([]EventNotificationConfig, bool) {
	if o == nil || IsNil(o.EventNotificationConfigs) {
		return nil, false
	}
	return o.EventNotificationConfigs, true
}

// HasEventNotificationConfigs returns a boolean if a field has been set.
func (o *CreateRegistry200Response) HasEventNotificationConfigs() bool {
	if o != nil && !IsNil(o.EventNotificationConfigs) {
		return true
	}

	return false
}

// SetEventNotificationConfigs gets a reference to the given []EventNotificationConfig and assigns it to the EventNotificationConfigs field.
func (o *CreateRegistry200Response) SetEventNotificationConfigs(v []EventNotificationConfig) {
	o.EventNotificationConfigs = v
}

// GetLogNotificationConfig returns the LogNotificationConfig field value if set, zero value otherwise.
func (o *CreateRegistry200Response) GetLogNotificationConfig() NotificationConfig {
	if o == nil || IsNil(o.LogNotificationConfig) {
		var ret NotificationConfig
		return ret
	}
	return *o.LogNotificationConfig
}

// GetLogNotificationConfigOk returns a tuple with the LogNotificationConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRegistry200Response) GetLogNotificationConfigOk() (*NotificationConfig, bool) {
	if o == nil || IsNil(o.LogNotificationConfig) {
		return nil, false
	}
	return o.LogNotificationConfig, true
}

// HasLogNotificationConfig returns a boolean if a field has been set.
func (o *CreateRegistry200Response) HasLogNotificationConfig() bool {
	if o != nil && !IsNil(o.LogNotificationConfig) {
		return true
	}

	return false
}

// SetLogNotificationConfig gets a reference to the given NotificationConfig and assigns it to the LogNotificationConfig field.
func (o *CreateRegistry200Response) SetLogNotificationConfig(v NotificationConfig) {
	o.LogNotificationConfig = &v
}

// GetStateNotificationConfig returns the StateNotificationConfig field value if set, zero value otherwise.
func (o *CreateRegistry200Response) GetStateNotificationConfig() NotificationConfig {
	if o == nil || IsNil(o.StateNotificationConfig) {
		var ret NotificationConfig
		return ret
	}
	return *o.StateNotificationConfig
}

// GetStateNotificationConfigOk returns a tuple with the StateNotificationConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRegistry200Response) GetStateNotificationConfigOk() (*NotificationConfig, bool) {
	if o == nil || IsNil(o.StateNotificationConfig) {
		return nil, false
	}
	return o.StateNotificationConfig, true
}

// HasStateNotificationConfig returns a boolean if a field has been set.
func (o *CreateRegistry200Response) HasStateNotificationConfig() bool {
	if o != nil && !IsNil(o.StateNotificationConfig) {
		return true
	}

	return false
}

// SetStateNotificationConfig gets a reference to the given NotificationConfig and assigns it to the StateNotificationConfig field.
func (o *CreateRegistry200Response) SetStateNotificationConfig(v NotificationConfig) {
	o.StateNotificationConfig = &v
}

// GetNumberOfDevices returns the NumberOfDevices field value if set, zero value otherwise.
func (o *CreateRegistry200Response) GetNumberOfDevices() int32 {
	if o == nil || IsNil(o.NumberOfDevices) {
		var ret int32
		return ret
	}
	return *o.NumberOfDevices
}

// GetNumberOfDevicesOk returns a tuple with the NumberOfDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRegistry200Response) GetNumberOfDevicesOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfDevices) {
		return nil, false
	}
	return o.NumberOfDevices, true
}

// HasNumberOfDevices returns a boolean if a field has been set.
func (o *CreateRegistry200Response) HasNumberOfDevices() bool {
	if o != nil && !IsNil(o.NumberOfDevices) {
		return true
	}

	return false
}

// SetNumberOfDevices gets a reference to the given int32 and assigns it to the NumberOfDevices field.
func (o *CreateRegistry200Response) SetNumberOfDevices(v int32) {
	o.NumberOfDevices = &v
}

// GetNumberOfGateways returns the NumberOfGateways field value if set, zero value otherwise.
func (o *CreateRegistry200Response) GetNumberOfGateways() int32 {
	if o == nil || IsNil(o.NumberOfGateways) {
		var ret int32
		return ret
	}
	return *o.NumberOfGateways
}

// GetNumberOfGatewaysOk returns a tuple with the NumberOfGateways field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRegistry200Response) GetNumberOfGatewaysOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfGateways) {
		return nil, false
	}
	return o.NumberOfGateways, true
}

// HasNumberOfGateways returns a boolean if a field has been set.
func (o *CreateRegistry200Response) HasNumberOfGateways() bool {
	if o != nil && !IsNil(o.NumberOfGateways) {
		return true
	}

	return false
}

// SetNumberOfGateways gets a reference to the given int32 and assigns it to the NumberOfGateways field.
func (o *CreateRegistry200Response) SetNumberOfGateways(v int32) {
	o.NumberOfGateways = &v
}

func (o CreateRegistry200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRegistry200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["parent"] = o.Parent
	if !IsNil(o.CreatedOn) {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if !IsNil(o.UpdatedOn) {
		toSerialize["updatedOn"] = o.UpdatedOn
	}
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
	if !IsNil(o.NumberOfDevices) {
		toSerialize["numberOfDevices"] = o.NumberOfDevices
	}
	if !IsNil(o.NumberOfGateways) {
		toSerialize["numberOfGateways"] = o.NumberOfGateways
	}
	return toSerialize, nil
}

type NullableCreateRegistry200Response struct {
	value *CreateRegistry200Response
	isSet bool
}

func (v NullableCreateRegistry200Response) Get() *CreateRegistry200Response {
	return v.value
}

func (v *NullableCreateRegistry200Response) Set(val *CreateRegistry200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRegistry200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRegistry200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRegistry200Response(val *CreateRegistry200Response) *NullableCreateRegistry200Response {
	return &NullableCreateRegistry200Response{value: val, isSet: true}
}

func (v NullableCreateRegistry200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRegistry200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


