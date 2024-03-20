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

// checks if the CustomOnboard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomOnboard{}

// CustomOnboard struct for CustomOnboard
type CustomOnboard struct {
	Id string `json:"id"`
	Name *string `json:"name,omitempty"`
	NumId *string `json:"numId,omitempty"`
	Parent *string `json:"parent,omitempty"`
	Registry *string `json:"registry,omitempty"`
	Blocked *bool `json:"blocked,omitempty"`
	Capresent *bool `json:"capresent,omitempty"`
	Subscription *string `json:"subscription,omitempty"`
	CreatedOn *string `json:"createdOn,omitempty"`
	UpdatedOn *string `json:"updatedOn,omitempty"`
	Credentials []DeviceCredential `json:"credentials,omitempty"`
	Gateway []string `json:"gateway,omitempty"`
	GatewayConfig *GatewayConfig `json:"gatewayConfig,omitempty"`
	IsGateway *bool `json:"isGateway,omitempty"`
	DeviceErrors *string `json:"deviceErrors,omitempty"`
	ClientOnline *bool `json:"clientOnline,omitempty"`
	LastConfigAckTime *string `json:"lastConfigAckTime,omitempty"`
	LastConfigSendTime *string `json:"lastConfigSendTime,omitempty"`
	LastErrorStatus *ErrorStatus `json:"lastErrorStatus,omitempty"`
	LastErrorTime *string `json:"lastErrorTime,omitempty"`
	LastEventTime *string `json:"lastEventTime,omitempty"`
	LastHeartbeatTime *string `json:"lastHeartbeatTime,omitempty"`
	LastStateTime *string `json:"lastStateTime,omitempty"`
	LogLevel *LogLevel `json:"logLevel,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
	Config *DeviceConfig `json:"config,omitempty"`
	State *DeviceState `json:"state,omitempty"`
	Policy *Policy `json:"policy,omitempty"`
	CustomOnboardData *string `json:"customOnboardData,omitempty"`
	IsApprove *bool `json:"isApprove,omitempty"`
	TcpUdpModelDetails *ComponentsSchemasTcpUdpModel `json:"tcpUdpModelDetails,omitempty"`
	TcpUdpModelId *uint32 `json:"tcpUdpModelId,omitempty"`
	IsTcpUdpDevice *bool `json:"isTcpUdpDevice,omitempty"`
}

// NewCustomOnboard instantiates a new CustomOnboard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomOnboard(id string) *CustomOnboard {
	this := CustomOnboard{}
	this.Id = id
	return &this
}

// NewCustomOnboardWithDefaults instantiates a new CustomOnboard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomOnboardWithDefaults() *CustomOnboard {
	this := CustomOnboard{}
	return &this
}

// GetId returns the Id field value
func (o *CustomOnboard) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomOnboard) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomOnboard) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomOnboard) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomOnboard) SetName(v string) {
	o.Name = &v
}

// GetNumId returns the NumId field value if set, zero value otherwise.
func (o *CustomOnboard) GetNumId() string {
	if o == nil || IsNil(o.NumId) {
		var ret string
		return ret
	}
	return *o.NumId
}

// GetNumIdOk returns a tuple with the NumId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetNumIdOk() (*string, bool) {
	if o == nil || IsNil(o.NumId) {
		return nil, false
	}
	return o.NumId, true
}

// HasNumId returns a boolean if a field has been set.
func (o *CustomOnboard) HasNumId() bool {
	if o != nil && !IsNil(o.NumId) {
		return true
	}

	return false
}

// SetNumId gets a reference to the given string and assigns it to the NumId field.
func (o *CustomOnboard) SetNumId(v string) {
	o.NumId = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *CustomOnboard) GetParent() string {
	if o == nil || IsNil(o.Parent) {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetParentOk() (*string, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *CustomOnboard) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *CustomOnboard) SetParent(v string) {
	o.Parent = &v
}

// GetRegistry returns the Registry field value if set, zero value otherwise.
func (o *CustomOnboard) GetRegistry() string {
	if o == nil || IsNil(o.Registry) {
		var ret string
		return ret
	}
	return *o.Registry
}

// GetRegistryOk returns a tuple with the Registry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetRegistryOk() (*string, bool) {
	if o == nil || IsNil(o.Registry) {
		return nil, false
	}
	return o.Registry, true
}

// HasRegistry returns a boolean if a field has been set.
func (o *CustomOnboard) HasRegistry() bool {
	if o != nil && !IsNil(o.Registry) {
		return true
	}

	return false
}

// SetRegistry gets a reference to the given string and assigns it to the Registry field.
func (o *CustomOnboard) SetRegistry(v string) {
	o.Registry = &v
}

// GetBlocked returns the Blocked field value if set, zero value otherwise.
func (o *CustomOnboard) GetBlocked() bool {
	if o == nil || IsNil(o.Blocked) {
		var ret bool
		return ret
	}
	return *o.Blocked
}

// GetBlockedOk returns a tuple with the Blocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetBlockedOk() (*bool, bool) {
	if o == nil || IsNil(o.Blocked) {
		return nil, false
	}
	return o.Blocked, true
}

// HasBlocked returns a boolean if a field has been set.
func (o *CustomOnboard) HasBlocked() bool {
	if o != nil && !IsNil(o.Blocked) {
		return true
	}

	return false
}

// SetBlocked gets a reference to the given bool and assigns it to the Blocked field.
func (o *CustomOnboard) SetBlocked(v bool) {
	o.Blocked = &v
}

// GetCapresent returns the Capresent field value if set, zero value otherwise.
func (o *CustomOnboard) GetCapresent() bool {
	if o == nil || IsNil(o.Capresent) {
		var ret bool
		return ret
	}
	return *o.Capresent
}

// GetCapresentOk returns a tuple with the Capresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetCapresentOk() (*bool, bool) {
	if o == nil || IsNil(o.Capresent) {
		return nil, false
	}
	return o.Capresent, true
}

// HasCapresent returns a boolean if a field has been set.
func (o *CustomOnboard) HasCapresent() bool {
	if o != nil && !IsNil(o.Capresent) {
		return true
	}

	return false
}

// SetCapresent gets a reference to the given bool and assigns it to the Capresent field.
func (o *CustomOnboard) SetCapresent(v bool) {
	o.Capresent = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *CustomOnboard) GetSubscription() string {
	if o == nil || IsNil(o.Subscription) {
		var ret string
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetSubscriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *CustomOnboard) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given string and assigns it to the Subscription field.
func (o *CustomOnboard) SetSubscription(v string) {
	o.Subscription = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *CustomOnboard) GetCreatedOn() string {
	if o == nil || IsNil(o.CreatedOn) {
		var ret string
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetCreatedOnOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedOn) {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *CustomOnboard) HasCreatedOn() bool {
	if o != nil && !IsNil(o.CreatedOn) {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given string and assigns it to the CreatedOn field.
func (o *CustomOnboard) SetCreatedOn(v string) {
	o.CreatedOn = &v
}

// GetUpdatedOn returns the UpdatedOn field value if set, zero value otherwise.
func (o *CustomOnboard) GetUpdatedOn() string {
	if o == nil || IsNil(o.UpdatedOn) {
		var ret string
		return ret
	}
	return *o.UpdatedOn
}

// GetUpdatedOnOk returns a tuple with the UpdatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetUpdatedOnOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedOn) {
		return nil, false
	}
	return o.UpdatedOn, true
}

// HasUpdatedOn returns a boolean if a field has been set.
func (o *CustomOnboard) HasUpdatedOn() bool {
	if o != nil && !IsNil(o.UpdatedOn) {
		return true
	}

	return false
}

// SetUpdatedOn gets a reference to the given string and assigns it to the UpdatedOn field.
func (o *CustomOnboard) SetUpdatedOn(v string) {
	o.UpdatedOn = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *CustomOnboard) GetCredentials() []DeviceCredential {
	if o == nil || IsNil(o.Credentials) {
		var ret []DeviceCredential
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetCredentialsOk() ([]DeviceCredential, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *CustomOnboard) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given []DeviceCredential and assigns it to the Credentials field.
func (o *CustomOnboard) SetCredentials(v []DeviceCredential) {
	o.Credentials = v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *CustomOnboard) GetGateway() []string {
	if o == nil || IsNil(o.Gateway) {
		var ret []string
		return ret
	}
	return o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetGatewayOk() ([]string, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *CustomOnboard) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given []string and assigns it to the Gateway field.
func (o *CustomOnboard) SetGateway(v []string) {
	o.Gateway = v
}

// GetGatewayConfig returns the GatewayConfig field value if set, zero value otherwise.
func (o *CustomOnboard) GetGatewayConfig() GatewayConfig {
	if o == nil || IsNil(o.GatewayConfig) {
		var ret GatewayConfig
		return ret
	}
	return *o.GatewayConfig
}

// GetGatewayConfigOk returns a tuple with the GatewayConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetGatewayConfigOk() (*GatewayConfig, bool) {
	if o == nil || IsNil(o.GatewayConfig) {
		return nil, false
	}
	return o.GatewayConfig, true
}

// HasGatewayConfig returns a boolean if a field has been set.
func (o *CustomOnboard) HasGatewayConfig() bool {
	if o != nil && !IsNil(o.GatewayConfig) {
		return true
	}

	return false
}

// SetGatewayConfig gets a reference to the given GatewayConfig and assigns it to the GatewayConfig field.
func (o *CustomOnboard) SetGatewayConfig(v GatewayConfig) {
	o.GatewayConfig = &v
}

// GetIsGateway returns the IsGateway field value if set, zero value otherwise.
func (o *CustomOnboard) GetIsGateway() bool {
	if o == nil || IsNil(o.IsGateway) {
		var ret bool
		return ret
	}
	return *o.IsGateway
}

// GetIsGatewayOk returns a tuple with the IsGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetIsGatewayOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGateway) {
		return nil, false
	}
	return o.IsGateway, true
}

// HasIsGateway returns a boolean if a field has been set.
func (o *CustomOnboard) HasIsGateway() bool {
	if o != nil && !IsNil(o.IsGateway) {
		return true
	}

	return false
}

// SetIsGateway gets a reference to the given bool and assigns it to the IsGateway field.
func (o *CustomOnboard) SetIsGateway(v bool) {
	o.IsGateway = &v
}

// GetDeviceErrors returns the DeviceErrors field value if set, zero value otherwise.
func (o *CustomOnboard) GetDeviceErrors() string {
	if o == nil || IsNil(o.DeviceErrors) {
		var ret string
		return ret
	}
	return *o.DeviceErrors
}

// GetDeviceErrorsOk returns a tuple with the DeviceErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetDeviceErrorsOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceErrors) {
		return nil, false
	}
	return o.DeviceErrors, true
}

// HasDeviceErrors returns a boolean if a field has been set.
func (o *CustomOnboard) HasDeviceErrors() bool {
	if o != nil && !IsNil(o.DeviceErrors) {
		return true
	}

	return false
}

// SetDeviceErrors gets a reference to the given string and assigns it to the DeviceErrors field.
func (o *CustomOnboard) SetDeviceErrors(v string) {
	o.DeviceErrors = &v
}

// GetClientOnline returns the ClientOnline field value if set, zero value otherwise.
func (o *CustomOnboard) GetClientOnline() bool {
	if o == nil || IsNil(o.ClientOnline) {
		var ret bool
		return ret
	}
	return *o.ClientOnline
}

// GetClientOnlineOk returns a tuple with the ClientOnline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetClientOnlineOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientOnline) {
		return nil, false
	}
	return o.ClientOnline, true
}

// HasClientOnline returns a boolean if a field has been set.
func (o *CustomOnboard) HasClientOnline() bool {
	if o != nil && !IsNil(o.ClientOnline) {
		return true
	}

	return false
}

// SetClientOnline gets a reference to the given bool and assigns it to the ClientOnline field.
func (o *CustomOnboard) SetClientOnline(v bool) {
	o.ClientOnline = &v
}

// GetLastConfigAckTime returns the LastConfigAckTime field value if set, zero value otherwise.
func (o *CustomOnboard) GetLastConfigAckTime() string {
	if o == nil || IsNil(o.LastConfigAckTime) {
		var ret string
		return ret
	}
	return *o.LastConfigAckTime
}

// GetLastConfigAckTimeOk returns a tuple with the LastConfigAckTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetLastConfigAckTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LastConfigAckTime) {
		return nil, false
	}
	return o.LastConfigAckTime, true
}

// HasLastConfigAckTime returns a boolean if a field has been set.
func (o *CustomOnboard) HasLastConfigAckTime() bool {
	if o != nil && !IsNil(o.LastConfigAckTime) {
		return true
	}

	return false
}

// SetLastConfigAckTime gets a reference to the given string and assigns it to the LastConfigAckTime field.
func (o *CustomOnboard) SetLastConfigAckTime(v string) {
	o.LastConfigAckTime = &v
}

// GetLastConfigSendTime returns the LastConfigSendTime field value if set, zero value otherwise.
func (o *CustomOnboard) GetLastConfigSendTime() string {
	if o == nil || IsNil(o.LastConfigSendTime) {
		var ret string
		return ret
	}
	return *o.LastConfigSendTime
}

// GetLastConfigSendTimeOk returns a tuple with the LastConfigSendTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetLastConfigSendTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LastConfigSendTime) {
		return nil, false
	}
	return o.LastConfigSendTime, true
}

// HasLastConfigSendTime returns a boolean if a field has been set.
func (o *CustomOnboard) HasLastConfigSendTime() bool {
	if o != nil && !IsNil(o.LastConfigSendTime) {
		return true
	}

	return false
}

// SetLastConfigSendTime gets a reference to the given string and assigns it to the LastConfigSendTime field.
func (o *CustomOnboard) SetLastConfigSendTime(v string) {
	o.LastConfigSendTime = &v
}

// GetLastErrorStatus returns the LastErrorStatus field value if set, zero value otherwise.
func (o *CustomOnboard) GetLastErrorStatus() ErrorStatus {
	if o == nil || IsNil(o.LastErrorStatus) {
		var ret ErrorStatus
		return ret
	}
	return *o.LastErrorStatus
}

// GetLastErrorStatusOk returns a tuple with the LastErrorStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetLastErrorStatusOk() (*ErrorStatus, bool) {
	if o == nil || IsNil(o.LastErrorStatus) {
		return nil, false
	}
	return o.LastErrorStatus, true
}

// HasLastErrorStatus returns a boolean if a field has been set.
func (o *CustomOnboard) HasLastErrorStatus() bool {
	if o != nil && !IsNil(o.LastErrorStatus) {
		return true
	}

	return false
}

// SetLastErrorStatus gets a reference to the given ErrorStatus and assigns it to the LastErrorStatus field.
func (o *CustomOnboard) SetLastErrorStatus(v ErrorStatus) {
	o.LastErrorStatus = &v
}

// GetLastErrorTime returns the LastErrorTime field value if set, zero value otherwise.
func (o *CustomOnboard) GetLastErrorTime() string {
	if o == nil || IsNil(o.LastErrorTime) {
		var ret string
		return ret
	}
	return *o.LastErrorTime
}

// GetLastErrorTimeOk returns a tuple with the LastErrorTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetLastErrorTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LastErrorTime) {
		return nil, false
	}
	return o.LastErrorTime, true
}

// HasLastErrorTime returns a boolean if a field has been set.
func (o *CustomOnboard) HasLastErrorTime() bool {
	if o != nil && !IsNil(o.LastErrorTime) {
		return true
	}

	return false
}

// SetLastErrorTime gets a reference to the given string and assigns it to the LastErrorTime field.
func (o *CustomOnboard) SetLastErrorTime(v string) {
	o.LastErrorTime = &v
}

// GetLastEventTime returns the LastEventTime field value if set, zero value otherwise.
func (o *CustomOnboard) GetLastEventTime() string {
	if o == nil || IsNil(o.LastEventTime) {
		var ret string
		return ret
	}
	return *o.LastEventTime
}

// GetLastEventTimeOk returns a tuple with the LastEventTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetLastEventTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LastEventTime) {
		return nil, false
	}
	return o.LastEventTime, true
}

// HasLastEventTime returns a boolean if a field has been set.
func (o *CustomOnboard) HasLastEventTime() bool {
	if o != nil && !IsNil(o.LastEventTime) {
		return true
	}

	return false
}

// SetLastEventTime gets a reference to the given string and assigns it to the LastEventTime field.
func (o *CustomOnboard) SetLastEventTime(v string) {
	o.LastEventTime = &v
}

// GetLastHeartbeatTime returns the LastHeartbeatTime field value if set, zero value otherwise.
func (o *CustomOnboard) GetLastHeartbeatTime() string {
	if o == nil || IsNil(o.LastHeartbeatTime) {
		var ret string
		return ret
	}
	return *o.LastHeartbeatTime
}

// GetLastHeartbeatTimeOk returns a tuple with the LastHeartbeatTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetLastHeartbeatTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LastHeartbeatTime) {
		return nil, false
	}
	return o.LastHeartbeatTime, true
}

// HasLastHeartbeatTime returns a boolean if a field has been set.
func (o *CustomOnboard) HasLastHeartbeatTime() bool {
	if o != nil && !IsNil(o.LastHeartbeatTime) {
		return true
	}

	return false
}

// SetLastHeartbeatTime gets a reference to the given string and assigns it to the LastHeartbeatTime field.
func (o *CustomOnboard) SetLastHeartbeatTime(v string) {
	o.LastHeartbeatTime = &v
}

// GetLastStateTime returns the LastStateTime field value if set, zero value otherwise.
func (o *CustomOnboard) GetLastStateTime() string {
	if o == nil || IsNil(o.LastStateTime) {
		var ret string
		return ret
	}
	return *o.LastStateTime
}

// GetLastStateTimeOk returns a tuple with the LastStateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetLastStateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LastStateTime) {
		return nil, false
	}
	return o.LastStateTime, true
}

// HasLastStateTime returns a boolean if a field has been set.
func (o *CustomOnboard) HasLastStateTime() bool {
	if o != nil && !IsNil(o.LastStateTime) {
		return true
	}

	return false
}

// SetLastStateTime gets a reference to the given string and assigns it to the LastStateTime field.
func (o *CustomOnboard) SetLastStateTime(v string) {
	o.LastStateTime = &v
}

// GetLogLevel returns the LogLevel field value if set, zero value otherwise.
func (o *CustomOnboard) GetLogLevel() LogLevel {
	if o == nil || IsNil(o.LogLevel) {
		var ret LogLevel
		return ret
	}
	return *o.LogLevel
}

// GetLogLevelOk returns a tuple with the LogLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetLogLevelOk() (*LogLevel, bool) {
	if o == nil || IsNil(o.LogLevel) {
		return nil, false
	}
	return o.LogLevel, true
}

// HasLogLevel returns a boolean if a field has been set.
func (o *CustomOnboard) HasLogLevel() bool {
	if o != nil && !IsNil(o.LogLevel) {
		return true
	}

	return false
}

// SetLogLevel gets a reference to the given LogLevel and assigns it to the LogLevel field.
func (o *CustomOnboard) SetLogLevel(v LogLevel) {
	o.LogLevel = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CustomOnboard) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CustomOnboard) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *CustomOnboard) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *CustomOnboard) GetConfig() DeviceConfig {
	if o == nil || IsNil(o.Config) {
		var ret DeviceConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetConfigOk() (*DeviceConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *CustomOnboard) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given DeviceConfig and assigns it to the Config field.
func (o *CustomOnboard) SetConfig(v DeviceConfig) {
	o.Config = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CustomOnboard) GetState() DeviceState {
	if o == nil || IsNil(o.State) {
		var ret DeviceState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetStateOk() (*DeviceState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CustomOnboard) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given DeviceState and assigns it to the State field.
func (o *CustomOnboard) SetState(v DeviceState) {
	o.State = &v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *CustomOnboard) GetPolicy() Policy {
	if o == nil || IsNil(o.Policy) {
		var ret Policy
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetPolicyOk() (*Policy, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *CustomOnboard) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given Policy and assigns it to the Policy field.
func (o *CustomOnboard) SetPolicy(v Policy) {
	o.Policy = &v
}

// GetCustomOnboardData returns the CustomOnboardData field value if set, zero value otherwise.
func (o *CustomOnboard) GetCustomOnboardData() string {
	if o == nil || IsNil(o.CustomOnboardData) {
		var ret string
		return ret
	}
	return *o.CustomOnboardData
}

// GetCustomOnboardDataOk returns a tuple with the CustomOnboardData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetCustomOnboardDataOk() (*string, bool) {
	if o == nil || IsNil(o.CustomOnboardData) {
		return nil, false
	}
	return o.CustomOnboardData, true
}

// HasCustomOnboardData returns a boolean if a field has been set.
func (o *CustomOnboard) HasCustomOnboardData() bool {
	if o != nil && !IsNil(o.CustomOnboardData) {
		return true
	}

	return false
}

// SetCustomOnboardData gets a reference to the given string and assigns it to the CustomOnboardData field.
func (o *CustomOnboard) SetCustomOnboardData(v string) {
	o.CustomOnboardData = &v
}

// GetIsApprove returns the IsApprove field value if set, zero value otherwise.
func (o *CustomOnboard) GetIsApprove() bool {
	if o == nil || IsNil(o.IsApprove) {
		var ret bool
		return ret
	}
	return *o.IsApprove
}

// GetIsApproveOk returns a tuple with the IsApprove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetIsApproveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsApprove) {
		return nil, false
	}
	return o.IsApprove, true
}

// HasIsApprove returns a boolean if a field has been set.
func (o *CustomOnboard) HasIsApprove() bool {
	if o != nil && !IsNil(o.IsApprove) {
		return true
	}

	return false
}

// SetIsApprove gets a reference to the given bool and assigns it to the IsApprove field.
func (o *CustomOnboard) SetIsApprove(v bool) {
	o.IsApprove = &v
}

// GetTcpUdpModelDetails returns the TcpUdpModelDetails field value if set, zero value otherwise.
func (o *CustomOnboard) GetTcpUdpModelDetails() ComponentsSchemasTcpUdpModel {
	if o == nil || IsNil(o.TcpUdpModelDetails) {
		var ret ComponentsSchemasTcpUdpModel
		return ret
	}
	return *o.TcpUdpModelDetails
}

// GetTcpUdpModelDetailsOk returns a tuple with the TcpUdpModelDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetTcpUdpModelDetailsOk() (*ComponentsSchemasTcpUdpModel, bool) {
	if o == nil || IsNil(o.TcpUdpModelDetails) {
		return nil, false
	}
	return o.TcpUdpModelDetails, true
}

// HasTcpUdpModelDetails returns a boolean if a field has been set.
func (o *CustomOnboard) HasTcpUdpModelDetails() bool {
	if o != nil && !IsNil(o.TcpUdpModelDetails) {
		return true
	}

	return false
}

// SetTcpUdpModelDetails gets a reference to the given ComponentsSchemasTcpUdpModel and assigns it to the TcpUdpModelDetails field.
func (o *CustomOnboard) SetTcpUdpModelDetails(v ComponentsSchemasTcpUdpModel) {
	o.TcpUdpModelDetails = &v
}

// GetTcpUdpModelId returns the TcpUdpModelId field value if set, zero value otherwise.
func (o *CustomOnboard) GetTcpUdpModelId() uint32 {
	if o == nil || IsNil(o.TcpUdpModelId) {
		var ret uint32
		return ret
	}
	return *o.TcpUdpModelId
}

// GetTcpUdpModelIdOk returns a tuple with the TcpUdpModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetTcpUdpModelIdOk() (*uint32, bool) {
	if o == nil || IsNil(o.TcpUdpModelId) {
		return nil, false
	}
	return o.TcpUdpModelId, true
}

// HasTcpUdpModelId returns a boolean if a field has been set.
func (o *CustomOnboard) HasTcpUdpModelId() bool {
	if o != nil && !IsNil(o.TcpUdpModelId) {
		return true
	}

	return false
}

// SetTcpUdpModelId gets a reference to the given uint32 and assigns it to the TcpUdpModelId field.
func (o *CustomOnboard) SetTcpUdpModelId(v uint32) {
	o.TcpUdpModelId = &v
}

// GetIsTcpUdpDevice returns the IsTcpUdpDevice field value if set, zero value otherwise.
func (o *CustomOnboard) GetIsTcpUdpDevice() bool {
	if o == nil || IsNil(o.IsTcpUdpDevice) {
		var ret bool
		return ret
	}
	return *o.IsTcpUdpDevice
}

// GetIsTcpUdpDeviceOk returns a tuple with the IsTcpUdpDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboard) GetIsTcpUdpDeviceOk() (*bool, bool) {
	if o == nil || IsNil(o.IsTcpUdpDevice) {
		return nil, false
	}
	return o.IsTcpUdpDevice, true
}

// HasIsTcpUdpDevice returns a boolean if a field has been set.
func (o *CustomOnboard) HasIsTcpUdpDevice() bool {
	if o != nil && !IsNil(o.IsTcpUdpDevice) {
		return true
	}

	return false
}

// SetIsTcpUdpDevice gets a reference to the given bool and assigns it to the IsTcpUdpDevice field.
func (o *CustomOnboard) SetIsTcpUdpDevice(v bool) {
	o.IsTcpUdpDevice = &v
}

func (o CustomOnboard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomOnboard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	// skip: name is readOnly
	// skip: numId is readOnly
	// skip: parent is readOnly
	// skip: registry is readOnly
	if !IsNil(o.Blocked) {
		toSerialize["blocked"] = o.Blocked
	}
	// skip: capresent is readOnly
	// skip: subscription is readOnly
	// skip: createdOn is readOnly
	// skip: updatedOn is readOnly
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !IsNil(o.GatewayConfig) {
		toSerialize["gatewayConfig"] = o.GatewayConfig
	}
	if !IsNil(o.IsGateway) {
		toSerialize["isGateway"] = o.IsGateway
	}
	// skip: deviceErrors is readOnly
	// skip: clientOnline is readOnly
	// skip: lastConfigAckTime is readOnly
	// skip: lastConfigSendTime is readOnly
	if !IsNil(o.LastErrorStatus) {
		toSerialize["lastErrorStatus"] = o.LastErrorStatus
	}
	// skip: lastErrorTime is readOnly
	// skip: lastEventTime is readOnly
	// skip: lastHeartbeatTime is readOnly
	// skip: lastStateTime is readOnly
	if !IsNil(o.LogLevel) {
		toSerialize["logLevel"] = o.LogLevel
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	if !IsNil(o.CustomOnboardData) {
		toSerialize["customOnboardData"] = o.CustomOnboardData
	}
	if !IsNil(o.IsApprove) {
		toSerialize["isApprove"] = o.IsApprove
	}
	if !IsNil(o.TcpUdpModelDetails) {
		toSerialize["tcpUdpModelDetails"] = o.TcpUdpModelDetails
	}
	if !IsNil(o.TcpUdpModelId) {
		toSerialize["tcpUdpModelId"] = o.TcpUdpModelId
	}
	if !IsNil(o.IsTcpUdpDevice) {
		toSerialize["isTcpUdpDevice"] = o.IsTcpUdpDevice
	}
	return toSerialize, nil
}

type NullableCustomOnboard struct {
	value *CustomOnboard
	isSet bool
}

func (v NullableCustomOnboard) Get() *CustomOnboard {
	return v.value
}

func (v *NullableCustomOnboard) Set(val *CustomOnboard) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomOnboard) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomOnboard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomOnboard(val *CustomOnboard) *NullableCustomOnboard {
	return &NullableCustomOnboard{value: val, isSet: true}
}

func (v NullableCustomOnboard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomOnboard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


