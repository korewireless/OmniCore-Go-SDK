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

// checks if the OmnicoreDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OmnicoreDevice{}

// OmnicoreDevice struct for OmnicoreDevice
type OmnicoreDevice struct {
	Blocked *bool `json:"blocked,omitempty"`
	Capresent *bool `json:"capresent,omitempty"`
	ClientOnline *bool `json:"clientOnline,omitempty"`
	Config *OmnicoreDeviceConfig `json:"config,omitempty"`
	CreatedOn *string `json:"createdOn,omitempty"`
	Credentials []OmnicoreDeviceCredential `json:"credentials,omitempty"`
	DeviceErrors *string `json:"deviceErrors,omitempty"`
	Gateway []string `json:"gateway,omitempty"`
	GatewayConfig *OmnicoreGatewayConfig `json:"gatewayConfig,omitempty"`
	Id string `json:"id"`
	IsGateway *bool `json:"isGateway,omitempty"`
	LastConfigAckTime *string `json:"lastConfigAckTime,omitempty"`
	LastConfigSendTime *string `json:"lastConfigSendTime,omitempty"`
	LastErrorStatus *OmnicoreErrorStatus `json:"lastErrorStatus,omitempty"`
	LastErrorTime *string `json:"lastErrorTime,omitempty"`
	LastEventTime *string `json:"lastEventTime,omitempty"`
	LastHeartbeatTime *string `json:"lastHeartbeatTime,omitempty"`
	LastStateTime *string `json:"lastStateTime,omitempty"`
	LogLevel *string `json:"logLevel,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
	Name *string `json:"name,omitempty"`
	NumId *string `json:"numId,omitempty"`
	Parent string `json:"parent"`
	Registry string `json:"registry"`
	State *OmnicoreDeviceState `json:"state,omitempty"`
	Subscription string `json:"subscription"`
	Subscriptions []string `json:"subscriptions,omitempty"`
	UpdatedOn *string `json:"updatedOn,omitempty"`
}

// NewOmnicoreDevice instantiates a new OmnicoreDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOmnicoreDevice(id string, parent string, registry string, subscription string) *OmnicoreDevice {
	this := OmnicoreDevice{}
	this.Id = id
	this.Parent = parent
	this.Registry = registry
	this.Subscription = subscription
	return &this
}

// NewOmnicoreDeviceWithDefaults instantiates a new OmnicoreDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOmnicoreDeviceWithDefaults() *OmnicoreDevice {
	this := OmnicoreDevice{}
	return &this
}

// GetBlocked returns the Blocked field value if set, zero value otherwise.
func (o *OmnicoreDevice) GetBlocked() bool {
	if o == nil || IsNil(o.Blocked) {
		var ret bool
		return ret
	}
	return *o.Blocked
}

// GetBlockedOk returns a tuple with the Blocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDevice) GetBlockedOk() (*bool, bool) {
	if o == nil || IsNil(o.Blocked) {
		return nil, false
	}
	return o.Blocked, true
}

// HasBlocked returns a boolean if a field has been set.
func (o *OmnicoreDevice) HasBlocked() bool {
	if o != nil && !IsNil(o.Blocked) {
		return true
	}

	return false
}

// SetBlocked gets a reference to the given bool and assigns it to the Blocked field.
func (o *OmnicoreDevice) SetBlocked(v bool) {
	o.Blocked = &v
}

// GetCapresent returns the Capresent field value if set, zero value otherwise.
func (o *OmnicoreDevice) GetCapresent() bool {
	if o == nil || IsNil(o.Capresent) {
		var ret bool
		return ret
	}
	return *o.Capresent
}

// GetCapresentOk returns a tuple with the Capresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDevice) GetCapresentOk() (*bool, bool) {
	if o == nil || IsNil(o.Capresent) {
		return nil, false
	}
	return o.Capresent, true
}

// HasCapresent returns a boolean if a field has been set.
func (o *OmnicoreDevice) HasCapresent() bool {
	if o != nil && !IsNil(o.Capresent) {
		return true
	}

	return false
}

// SetCapresent gets a reference to the given bool and assigns it to the Capresent field.
func (o *OmnicoreDevice) SetCapresent(v bool) {
	o.Capresent = &v
}

// GetClientOnline returns the ClientOnline field value if set, zero value otherwise.
func (o *OmnicoreDevice) GetClientOnline() bool {
	if o == nil || IsNil(o.ClientOnline) {
		var ret bool
		return ret
	}
	return *o.ClientOnline
}

// GetClientOnlineOk returns a tuple with the ClientOnline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDevice) GetClientOnlineOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientOnline) {
		return nil, false
	}
	return o.ClientOnline, true
}

// HasClientOnline returns a boolean if a field has been set.
func (o *OmnicoreDevice) HasClientOnline() bool {
	if o != nil && !IsNil(o.ClientOnline) {
		return true
	}

	return false
}

// SetClientOnline gets a reference to the given bool and assigns it to the ClientOnline field.
func (o *OmnicoreDevice) SetClientOnline(v bool) {
	o.ClientOnline = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *OmnicoreDevice) GetConfig() OmnicoreDeviceConfig {
	if o == nil || IsNil(o.Config) {
		var ret OmnicoreDeviceConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDevice) GetConfigOk() (*OmnicoreDeviceConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *OmnicoreDevice) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given OmnicoreDeviceConfig and assigns it to the Config field.
func (o *OmnicoreDevice) SetConfig(v OmnicoreDeviceConfig) {
	o.Config = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *OmnicoreDevice) GetCreatedOn() string {
	if o == nil || IsNil(o.CreatedOn) {
		var ret string
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDevice) GetCreatedOnOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedOn) {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *OmnicoreDevice) HasCreatedOn() bool {
	if o != nil && !IsNil(o.CreatedOn) {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given string and assigns it to the CreatedOn field.
func (o *OmnicoreDevice) SetCreatedOn(v string) {
	o.CreatedOn = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *OmnicoreDevice) GetCredentials() []OmnicoreDeviceCredential {
	if o == nil || IsNil(o.Credentials) {
		var ret []OmnicoreDeviceCredential
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDevice) GetCredentialsOk() ([]OmnicoreDeviceCredential, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *OmnicoreDevice) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given []OmnicoreDeviceCredential and assigns it to the Credentials field.
func (o *OmnicoreDevice) SetCredentials(v []OmnicoreDeviceCredential) {
	o.Credentials = v
}

// GetDeviceErrors returns the DeviceErrors field value if set, zero value otherwise.
func (o *OmnicoreDevice) GetDeviceErrors() string {
	if o == nil || IsNil(o.DeviceErrors) {
		var ret string
		return ret
	}
	return *o.DeviceErrors
}

// GetDeviceErrorsOk returns a tuple with the DeviceErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDevice) GetDeviceErrorsOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceErrors) {
		return nil, false
	}
	return o.DeviceErrors, true
}

// HasDeviceErrors returns a boolean if a field has been set.
func (o *OmnicoreDevice) HasDeviceErrors() bool {
	if o != nil && !IsNil(o.DeviceErrors) {
		return true
	}

	return false
}

// SetDeviceErrors gets a reference to the given string and assigns it to the DeviceErrors field.
func (o *OmnicoreDevice) SetDeviceErrors(v string) {
	o.DeviceErrors = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *OmnicoreDevice) GetGateway() []string {
	if o == nil || IsNil(o.Gateway) {
		var ret []string
		return ret
	}
	return o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDevice) GetGatewayOk() ([]string, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *OmnicoreDevice) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given []string and assigns it to the Gateway field.
func (o *OmnicoreDevice) SetGateway(v []string) {
	o.Gateway = v
}

// GetGatewayConfig returns the GatewayConfig field value if set, zero value otherwise.
func (o *OmnicoreDevice) GetGatewayConfig() OmnicoreGatewayConfig {
	if o == nil || IsNil(o.GatewayConfig) {
		var ret OmnicoreGatewayConfig
		return ret
	}
	return *o.GatewayConfig
}

// GetGatewayConfigOk returns a tuple with the GatewayConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDevice) GetGatewayConfigOk() (*OmnicoreGatewayConfig, bool) {
	if o == nil || IsNil(o.GatewayConfig) {
		return nil, false
	}
	return o.GatewayConfig, true
}

// HasGatewayConfig returns a boolean if a field has been set.
func (o *OmnicoreDevice) HasGatewayConfig() bool {
	if o != nil && !IsNil(o.GatewayConfig) {
		return true
	}

	return false
}

// SetGatewayConfig gets a reference to the given OmnicoreGatewayConfig and assigns it to the GatewayConfig field.
func (o *OmnicoreDevice) SetGatewayConfig(v OmnicoreGatewayConfig) {
	o.GatewayConfig = &v
}

// GetId returns the Id field value
func (o *OmnicoreDevice) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OmnicoreDevice) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OmnicoreDevice) SetId(v string) {
	o.Id = v
}

// GetIsGateway returns the IsGateway field value if set, zero value otherwise.
func (o *OmnicoreDevice) GetIsGateway() bool {
	if o == nil || IsNil(o.IsGateway) {
		var ret bool
		return ret
	}
	return *o.IsGateway
}

// GetIsGatewayOk returns a tuple with the IsGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDevice) GetIsGatewayOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGateway) {
		return nil, false
	}
	return o.IsGateway, true
}

// HasIsGateway returns a boolean if a field has been set.
func (o *OmnicoreDevice) HasIsGateway() bool {
	if o != nil && !IsNil(o.IsGateway) {
		return true
	}

	return false
}

// SetIsGateway gets a reference to the given bool and assigns it to the IsGateway field.
func (o *OmnicoreDevice) SetIsGateway(v bool) {
	o.IsGateway = &v
}

// GetLastConfigAckTime returns the LastConfigAckTime field value if set, zero value otherwise.
func (o *OmnicoreDevice) GetLastConfigAckTime() string {
	if o == nil || IsNil(o.LastConfigAckTime) {
		var ret string
		return ret
	}
	return *o.LastConfigAckTime
}

// GetLastConfigAckTimeOk returns a tuple with the LastConfigAckTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDevice) GetLastConfigAckTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LastConfigAckTime) {
		return nil, false
	}
	return o.LastConfigAckTime, true
}

// HasLastConfigAckTime returns a boolean if a field has been set.
func (o *OmnicoreDevice) HasLastConfigAckTime() bool {
	if o != nil && !IsNil(o.LastConfigAckTime) {
		return true
	}

	return false
}

// SetLastConfigAckTime gets a reference to the given string and assigns it to the LastConfigAckTime field.
func (o *OmnicoreDevice) SetLastConfigAckTime(v string) {
	o.LastConfigAckTime = &v
}

// GetLastConfigSendTime returns the LastConfigSendTime field value if set, zero value otherwise.
func (o *OmnicoreDevice) GetLastConfigSendTime() string {
	if o == nil || IsNil(o.LastConfigSendTime) {
		var ret string
		return ret
	}
	return *o.LastConfigSendTime
}

// GetLastConfigSendTimeOk returns a tuple with the LastConfigSendTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDevice) GetLastConfigSendTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LastConfigSendTime) {
		return nil, false
	}
	return o.LastConfigSendTime, true
}

// HasLastConfigSendTime returns a boolean if a field has been set.
func (o *OmnicoreDevice) HasLastConfigSendTime() bool {
	if o != nil && !IsNil(o.LastConfigSendTime) {
		return true
	}

	return false
}

// SetLastConfigSendTime gets a reference to the given string and assigns it to the LastConfigSendTime field.
func (o *OmnicoreDevice) SetLastConfigSendTime(v string) {
	o.LastConfigSendTime = &v
}

// GetLastErrorStatus returns the LastErrorStatus field value if set, zero value otherwise.
func (o *OmnicoreDevice) GetLastErrorStatus() OmnicoreErrorStatus {
	if o == nil || IsNil(o.LastErrorStatus) {
		var ret OmnicoreErrorStatus
		return ret
	}
	return *o.LastErrorStatus
}

// GetLastErrorStatusOk returns a tuple with the LastErrorStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDevice) GetLastErrorStatusOk() (*OmnicoreErrorStatus, bool) {
	if o == nil || IsNil(o.LastErrorStatus) {
		return nil, false
	}
	return o.LastErrorStatus, true
}

// HasLastErrorStatus returns a boolean if a field has been set.
func (o *OmnicoreDevice) HasLastErrorStatus() bool {
	if o != nil && !IsNil(o.LastErrorStatus) {
		return true
	}

	return false
}

// SetLastErrorStatus gets a reference to the given OmnicoreErrorStatus and assigns it to the LastErrorStatus field.
func (o *OmnicoreDevice) SetLastErrorStatus(v OmnicoreErrorStatus) {
	o.LastErrorStatus = &v
}

// GetLastErrorTime returns the LastErrorTime field value if set, zero value otherwise.
func (o *OmnicoreDevice) GetLastErrorTime() string {
	if o == nil || IsNil(o.LastErrorTime) {
		var ret string
		return ret
	}
	return *o.LastErrorTime
}

// GetLastErrorTimeOk returns a tuple with the LastErrorTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDevice) GetLastErrorTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LastErrorTime) {
		return nil, false
	}
	return o.LastErrorTime, true
}

// HasLastErrorTime returns a boolean if a field has been set.
func (o *OmnicoreDevice) HasLastErrorTime() bool {
	if o != nil && !IsNil(o.LastErrorTime) {
		return true
	}

	return false
}

// SetLastErrorTime gets a reference to the given string and assigns it to the LastErrorTime field.
func (o *OmnicoreDevice) SetLastErrorTime(v string) {
	o.LastErrorTime = &v
}

// GetLastEventTime returns the LastEventTime field value if set, zero value otherwise.
func (o *OmnicoreDevice) GetLastEventTime() string {
	if o == nil || IsNil(o.LastEventTime) {
		var ret string
		return ret
	}
	return *o.LastEventTime
}

// GetLastEventTimeOk returns a tuple with the LastEventTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDevice) GetLastEventTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LastEventTime) {
		return nil, false
	}
	return o.LastEventTime, true
}

// HasLastEventTime returns a boolean if a field has been set.
func (o *OmnicoreDevice) HasLastEventTime() bool {
	if o != nil && !IsNil(o.LastEventTime) {
		return true
	}

	return false
}

// SetLastEventTime gets a reference to the given string and assigns it to the LastEventTime field.
func (o *OmnicoreDevice) SetLastEventTime(v string) {
	o.LastEventTime = &v
}

// GetLastHeartbeatTime returns the LastHeartbeatTime field value if set, zero value otherwise.
func (o *OmnicoreDevice) GetLastHeartbeatTime() string {
	if o == nil || IsNil(o.LastHeartbeatTime) {
		var ret string
		return ret
	}
	return *o.LastHeartbeatTime
}

// GetLastHeartbeatTimeOk returns a tuple with the LastHeartbeatTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDevice) GetLastHeartbeatTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LastHeartbeatTime) {
		return nil, false
	}
	return o.LastHeartbeatTime, true
}

// HasLastHeartbeatTime returns a boolean if a field has been set.
func (o *OmnicoreDevice) HasLastHeartbeatTime() bool {
	if o != nil && !IsNil(o.LastHeartbeatTime) {
		return true
	}

	return false
}

// SetLastHeartbeatTime gets a reference to the given string and assigns it to the LastHeartbeatTime field.
func (o *OmnicoreDevice) SetLastHeartbeatTime(v string) {
	o.LastHeartbeatTime = &v
}

// GetLastStateTime returns the LastStateTime field value if set, zero value otherwise.
func (o *OmnicoreDevice) GetLastStateTime() string {
	if o == nil || IsNil(o.LastStateTime) {
		var ret string
		return ret
	}
	return *o.LastStateTime
}

// GetLastStateTimeOk returns a tuple with the LastStateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDevice) GetLastStateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LastStateTime) {
		return nil, false
	}
	return o.LastStateTime, true
}

// HasLastStateTime returns a boolean if a field has been set.
func (o *OmnicoreDevice) HasLastStateTime() bool {
	if o != nil && !IsNil(o.LastStateTime) {
		return true
	}

	return false
}

// SetLastStateTime gets a reference to the given string and assigns it to the LastStateTime field.
func (o *OmnicoreDevice) SetLastStateTime(v string) {
	o.LastStateTime = &v
}

// GetLogLevel returns the LogLevel field value if set, zero value otherwise.
func (o *OmnicoreDevice) GetLogLevel() string {
	if o == nil || IsNil(o.LogLevel) {
		var ret string
		return ret
	}
	return *o.LogLevel
}

// GetLogLevelOk returns a tuple with the LogLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDevice) GetLogLevelOk() (*string, bool) {
	if o == nil || IsNil(o.LogLevel) {
		return nil, false
	}
	return o.LogLevel, true
}

// HasLogLevel returns a boolean if a field has been set.
func (o *OmnicoreDevice) HasLogLevel() bool {
	if o != nil && !IsNil(o.LogLevel) {
		return true
	}

	return false
}

// SetLogLevel gets a reference to the given string and assigns it to the LogLevel field.
func (o *OmnicoreDevice) SetLogLevel(v string) {
	o.LogLevel = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *OmnicoreDevice) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDevice) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *OmnicoreDevice) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *OmnicoreDevice) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OmnicoreDevice) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDevice) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OmnicoreDevice) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OmnicoreDevice) SetName(v string) {
	o.Name = &v
}

// GetNumId returns the NumId field value if set, zero value otherwise.
func (o *OmnicoreDevice) GetNumId() string {
	if o == nil || IsNil(o.NumId) {
		var ret string
		return ret
	}
	return *o.NumId
}

// GetNumIdOk returns a tuple with the NumId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDevice) GetNumIdOk() (*string, bool) {
	if o == nil || IsNil(o.NumId) {
		return nil, false
	}
	return o.NumId, true
}

// HasNumId returns a boolean if a field has been set.
func (o *OmnicoreDevice) HasNumId() bool {
	if o != nil && !IsNil(o.NumId) {
		return true
	}

	return false
}

// SetNumId gets a reference to the given string and assigns it to the NumId field.
func (o *OmnicoreDevice) SetNumId(v string) {
	o.NumId = &v
}

// GetParent returns the Parent field value
func (o *OmnicoreDevice) GetParent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Parent
}

// GetParentOk returns a tuple with the Parent field value
// and a boolean to check if the value has been set.
func (o *OmnicoreDevice) GetParentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parent, true
}

// SetParent sets field value
func (o *OmnicoreDevice) SetParent(v string) {
	o.Parent = v
}

// GetRegistry returns the Registry field value
func (o *OmnicoreDevice) GetRegistry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Registry
}

// GetRegistryOk returns a tuple with the Registry field value
// and a boolean to check if the value has been set.
func (o *OmnicoreDevice) GetRegistryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Registry, true
}

// SetRegistry sets field value
func (o *OmnicoreDevice) SetRegistry(v string) {
	o.Registry = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *OmnicoreDevice) GetState() OmnicoreDeviceState {
	if o == nil || IsNil(o.State) {
		var ret OmnicoreDeviceState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDevice) GetStateOk() (*OmnicoreDeviceState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *OmnicoreDevice) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given OmnicoreDeviceState and assigns it to the State field.
func (o *OmnicoreDevice) SetState(v OmnicoreDeviceState) {
	o.State = &v
}

// GetSubscription returns the Subscription field value
func (o *OmnicoreDevice) GetSubscription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
func (o *OmnicoreDevice) GetSubscriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscription, true
}

// SetSubscription sets field value
func (o *OmnicoreDevice) SetSubscription(v string) {
	o.Subscription = v
}

// GetSubscriptions returns the Subscriptions field value if set, zero value otherwise.
func (o *OmnicoreDevice) GetSubscriptions() []string {
	if o == nil || IsNil(o.Subscriptions) {
		var ret []string
		return ret
	}
	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDevice) GetSubscriptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Subscriptions) {
		return nil, false
	}
	return o.Subscriptions, true
}

// HasSubscriptions returns a boolean if a field has been set.
func (o *OmnicoreDevice) HasSubscriptions() bool {
	if o != nil && !IsNil(o.Subscriptions) {
		return true
	}

	return false
}

// SetSubscriptions gets a reference to the given []string and assigns it to the Subscriptions field.
func (o *OmnicoreDevice) SetSubscriptions(v []string) {
	o.Subscriptions = v
}

// GetUpdatedOn returns the UpdatedOn field value if set, zero value otherwise.
func (o *OmnicoreDevice) GetUpdatedOn() string {
	if o == nil || IsNil(o.UpdatedOn) {
		var ret string
		return ret
	}
	return *o.UpdatedOn
}

// GetUpdatedOnOk returns a tuple with the UpdatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreDevice) GetUpdatedOnOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedOn) {
		return nil, false
	}
	return o.UpdatedOn, true
}

// HasUpdatedOn returns a boolean if a field has been set.
func (o *OmnicoreDevice) HasUpdatedOn() bool {
	if o != nil && !IsNil(o.UpdatedOn) {
		return true
	}

	return false
}

// SetUpdatedOn gets a reference to the given string and assigns it to the UpdatedOn field.
func (o *OmnicoreDevice) SetUpdatedOn(v string) {
	o.UpdatedOn = &v
}

func (o OmnicoreDevice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OmnicoreDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Blocked) {
		toSerialize["blocked"] = o.Blocked
	}
	if !IsNil(o.Capresent) {
		toSerialize["capresent"] = o.Capresent
	}
	if !IsNil(o.ClientOnline) {
		toSerialize["clientOnline"] = o.ClientOnline
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.CreatedOn) {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.DeviceErrors) {
		toSerialize["deviceErrors"] = o.DeviceErrors
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !IsNil(o.GatewayConfig) {
		toSerialize["gatewayConfig"] = o.GatewayConfig
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.IsGateway) {
		toSerialize["isGateway"] = o.IsGateway
	}
	if !IsNil(o.LastConfigAckTime) {
		toSerialize["lastConfigAckTime"] = o.LastConfigAckTime
	}
	if !IsNil(o.LastConfigSendTime) {
		toSerialize["lastConfigSendTime"] = o.LastConfigSendTime
	}
	if !IsNil(o.LastErrorStatus) {
		toSerialize["lastErrorStatus"] = o.LastErrorStatus
	}
	if !IsNil(o.LastErrorTime) {
		toSerialize["lastErrorTime"] = o.LastErrorTime
	}
	if !IsNil(o.LastEventTime) {
		toSerialize["lastEventTime"] = o.LastEventTime
	}
	if !IsNil(o.LastHeartbeatTime) {
		toSerialize["lastHeartbeatTime"] = o.LastHeartbeatTime
	}
	if !IsNil(o.LastStateTime) {
		toSerialize["lastStateTime"] = o.LastStateTime
	}
	if !IsNil(o.LogLevel) {
		toSerialize["logLevel"] = o.LogLevel
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NumId) {
		toSerialize["numId"] = o.NumId
	}
	toSerialize["parent"] = o.Parent
	toSerialize["registry"] = o.Registry
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	toSerialize["subscription"] = o.Subscription
	if !IsNil(o.Subscriptions) {
		toSerialize["subscriptions"] = o.Subscriptions
	}
	if !IsNil(o.UpdatedOn) {
		toSerialize["updatedOn"] = o.UpdatedOn
	}
	return toSerialize, nil
}

type NullableOmnicoreDevice struct {
	value *OmnicoreDevice
	isSet bool
}

func (v NullableOmnicoreDevice) Get() *OmnicoreDevice {
	return v.value
}

func (v *NullableOmnicoreDevice) Set(val *OmnicoreDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableOmnicoreDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableOmnicoreDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOmnicoreDevice(val *OmnicoreDevice) *NullableOmnicoreDevice {
	return &NullableOmnicoreDevice{value: val, isSet: true}
}

func (v NullableOmnicoreDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOmnicoreDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


