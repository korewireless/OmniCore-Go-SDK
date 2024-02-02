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

// checks if the DeviceOnline type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceOnline{}

// DeviceOnline struct for DeviceOnline
type DeviceOnline struct {
	Id string `json:"id"`
	Registry *string `json:"registry,omitempty"`
	ClientOnline *bool `json:"clientOnline,omitempty"`
	LastHeartbeatTime *string `json:"lastHeartbeatTime,omitempty"`
	Subscription *string `json:"subscription,omitempty"`
}

// NewDeviceOnline instantiates a new DeviceOnline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceOnline(id string) *DeviceOnline {
	this := DeviceOnline{}
	this.Id = id
	return &this
}

// NewDeviceOnlineWithDefaults instantiates a new DeviceOnline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceOnlineWithDefaults() *DeviceOnline {
	this := DeviceOnline{}
	return &this
}

// GetId returns the Id field value
func (o *DeviceOnline) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeviceOnline) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeviceOnline) SetId(v string) {
	o.Id = v
}

// GetRegistry returns the Registry field value if set, zero value otherwise.
func (o *DeviceOnline) GetRegistry() string {
	if o == nil || IsNil(o.Registry) {
		var ret string
		return ret
	}
	return *o.Registry
}

// GetRegistryOk returns a tuple with the Registry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceOnline) GetRegistryOk() (*string, bool) {
	if o == nil || IsNil(o.Registry) {
		return nil, false
	}
	return o.Registry, true
}

// HasRegistry returns a boolean if a field has been set.
func (o *DeviceOnline) HasRegistry() bool {
	if o != nil && !IsNil(o.Registry) {
		return true
	}

	return false
}

// SetRegistry gets a reference to the given string and assigns it to the Registry field.
func (o *DeviceOnline) SetRegistry(v string) {
	o.Registry = &v
}

// GetClientOnline returns the ClientOnline field value if set, zero value otherwise.
func (o *DeviceOnline) GetClientOnline() bool {
	if o == nil || IsNil(o.ClientOnline) {
		var ret bool
		return ret
	}
	return *o.ClientOnline
}

// GetClientOnlineOk returns a tuple with the ClientOnline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceOnline) GetClientOnlineOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientOnline) {
		return nil, false
	}
	return o.ClientOnline, true
}

// HasClientOnline returns a boolean if a field has been set.
func (o *DeviceOnline) HasClientOnline() bool {
	if o != nil && !IsNil(o.ClientOnline) {
		return true
	}

	return false
}

// SetClientOnline gets a reference to the given bool and assigns it to the ClientOnline field.
func (o *DeviceOnline) SetClientOnline(v bool) {
	o.ClientOnline = &v
}

// GetLastHeartbeatTime returns the LastHeartbeatTime field value if set, zero value otherwise.
func (o *DeviceOnline) GetLastHeartbeatTime() string {
	if o == nil || IsNil(o.LastHeartbeatTime) {
		var ret string
		return ret
	}
	return *o.LastHeartbeatTime
}

// GetLastHeartbeatTimeOk returns a tuple with the LastHeartbeatTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceOnline) GetLastHeartbeatTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LastHeartbeatTime) {
		return nil, false
	}
	return o.LastHeartbeatTime, true
}

// HasLastHeartbeatTime returns a boolean if a field has been set.
func (o *DeviceOnline) HasLastHeartbeatTime() bool {
	if o != nil && !IsNil(o.LastHeartbeatTime) {
		return true
	}

	return false
}

// SetLastHeartbeatTime gets a reference to the given string and assigns it to the LastHeartbeatTime field.
func (o *DeviceOnline) SetLastHeartbeatTime(v string) {
	o.LastHeartbeatTime = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *DeviceOnline) GetSubscription() string {
	if o == nil || IsNil(o.Subscription) {
		var ret string
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceOnline) GetSubscriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *DeviceOnline) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given string and assigns it to the Subscription field.
func (o *DeviceOnline) SetSubscription(v string) {
	o.Subscription = &v
}

func (o DeviceOnline) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceOnline) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	// skip: registry is readOnly
	// skip: clientOnline is readOnly
	// skip: lastHeartbeatTime is readOnly
	// skip: subscription is readOnly
	return toSerialize, nil
}

type NullableDeviceOnline struct {
	value *DeviceOnline
	isSet bool
}

func (v NullableDeviceOnline) Get() *DeviceOnline {
	return v.value
}

func (v *NullableDeviceOnline) Set(val *DeviceOnline) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceOnline) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceOnline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceOnline(val *DeviceOnline) *NullableDeviceOnline {
	return &NullableDeviceOnline{value: val, isSet: true}
}

func (v NullableDeviceOnline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceOnline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


