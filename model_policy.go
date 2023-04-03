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

// checks if the Policy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Policy{}

// Policy struct for Policy
type Policy struct {
	Connect bool `json:"Connect"`
	PublishState bool `json:"PublishState"`
	PublishEvents bool `json:"PublishEvents"`
	PublishEventsRegex string `json:"PublishEventsRegex"`
	PublishLoopback bool `json:"PublishLoopback"`
	SubscribeCommand bool `json:"SubscribeCommand"`
	SubscribeCommandRegex string `json:"SubscribeCommandRegex"`
	SubscribeBroadcast bool `json:"SubscribeBroadcast"`
	SubscribeBroadcastRegex string `json:"SubscribeBroadcastRegex"`
	SubscribeConfig bool `json:"SubscribeConfig"`
}

// NewPolicy instantiates a new Policy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicy(connect bool, publishState bool, publishEvents bool, publishEventsRegex string, publishLoopback bool, subscribeCommand bool, subscribeCommandRegex string, subscribeBroadcast bool, subscribeBroadcastRegex string, subscribeConfig bool) *Policy {
	this := Policy{}
	this.Connect = connect
	this.PublishState = publishState
	this.PublishEvents = publishEvents
	this.PublishEventsRegex = publishEventsRegex
	this.PublishLoopback = publishLoopback
	this.SubscribeCommand = subscribeCommand
	this.SubscribeCommandRegex = subscribeCommandRegex
	this.SubscribeBroadcast = subscribeBroadcast
	this.SubscribeBroadcastRegex = subscribeBroadcastRegex
	this.SubscribeConfig = subscribeConfig
	return &this
}

// NewPolicyWithDefaults instantiates a new Policy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyWithDefaults() *Policy {
	this := Policy{}
	return &this
}

// GetConnect returns the Connect field value
func (o *Policy) GetConnect() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Connect
}

// GetConnectOk returns a tuple with the Connect field value
// and a boolean to check if the value has been set.
func (o *Policy) GetConnectOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connect, true
}

// SetConnect sets field value
func (o *Policy) SetConnect(v bool) {
	o.Connect = v
}

// GetPublishState returns the PublishState field value
func (o *Policy) GetPublishState() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PublishState
}

// GetPublishStateOk returns a tuple with the PublishState field value
// and a boolean to check if the value has been set.
func (o *Policy) GetPublishStateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublishState, true
}

// SetPublishState sets field value
func (o *Policy) SetPublishState(v bool) {
	o.PublishState = v
}

// GetPublishEvents returns the PublishEvents field value
func (o *Policy) GetPublishEvents() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PublishEvents
}

// GetPublishEventsOk returns a tuple with the PublishEvents field value
// and a boolean to check if the value has been set.
func (o *Policy) GetPublishEventsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublishEvents, true
}

// SetPublishEvents sets field value
func (o *Policy) SetPublishEvents(v bool) {
	o.PublishEvents = v
}

// GetPublishEventsRegex returns the PublishEventsRegex field value
func (o *Policy) GetPublishEventsRegex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublishEventsRegex
}

// GetPublishEventsRegexOk returns a tuple with the PublishEventsRegex field value
// and a boolean to check if the value has been set.
func (o *Policy) GetPublishEventsRegexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublishEventsRegex, true
}

// SetPublishEventsRegex sets field value
func (o *Policy) SetPublishEventsRegex(v string) {
	o.PublishEventsRegex = v
}

// GetPublishLoopback returns the PublishLoopback field value
func (o *Policy) GetPublishLoopback() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PublishLoopback
}

// GetPublishLoopbackOk returns a tuple with the PublishLoopback field value
// and a boolean to check if the value has been set.
func (o *Policy) GetPublishLoopbackOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublishLoopback, true
}

// SetPublishLoopback sets field value
func (o *Policy) SetPublishLoopback(v bool) {
	o.PublishLoopback = v
}

// GetSubscribeCommand returns the SubscribeCommand field value
func (o *Policy) GetSubscribeCommand() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SubscribeCommand
}

// GetSubscribeCommandOk returns a tuple with the SubscribeCommand field value
// and a boolean to check if the value has been set.
func (o *Policy) GetSubscribeCommandOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscribeCommand, true
}

// SetSubscribeCommand sets field value
func (o *Policy) SetSubscribeCommand(v bool) {
	o.SubscribeCommand = v
}

// GetSubscribeCommandRegex returns the SubscribeCommandRegex field value
func (o *Policy) GetSubscribeCommandRegex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscribeCommandRegex
}

// GetSubscribeCommandRegexOk returns a tuple with the SubscribeCommandRegex field value
// and a boolean to check if the value has been set.
func (o *Policy) GetSubscribeCommandRegexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscribeCommandRegex, true
}

// SetSubscribeCommandRegex sets field value
func (o *Policy) SetSubscribeCommandRegex(v string) {
	o.SubscribeCommandRegex = v
}

// GetSubscribeBroadcast returns the SubscribeBroadcast field value
func (o *Policy) GetSubscribeBroadcast() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SubscribeBroadcast
}

// GetSubscribeBroadcastOk returns a tuple with the SubscribeBroadcast field value
// and a boolean to check if the value has been set.
func (o *Policy) GetSubscribeBroadcastOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscribeBroadcast, true
}

// SetSubscribeBroadcast sets field value
func (o *Policy) SetSubscribeBroadcast(v bool) {
	o.SubscribeBroadcast = v
}

// GetSubscribeBroadcastRegex returns the SubscribeBroadcastRegex field value
func (o *Policy) GetSubscribeBroadcastRegex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscribeBroadcastRegex
}

// GetSubscribeBroadcastRegexOk returns a tuple with the SubscribeBroadcastRegex field value
// and a boolean to check if the value has been set.
func (o *Policy) GetSubscribeBroadcastRegexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscribeBroadcastRegex, true
}

// SetSubscribeBroadcastRegex sets field value
func (o *Policy) SetSubscribeBroadcastRegex(v string) {
	o.SubscribeBroadcastRegex = v
}

// GetSubscribeConfig returns the SubscribeConfig field value
func (o *Policy) GetSubscribeConfig() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SubscribeConfig
}

// GetSubscribeConfigOk returns a tuple with the SubscribeConfig field value
// and a boolean to check if the value has been set.
func (o *Policy) GetSubscribeConfigOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscribeConfig, true
}

// SetSubscribeConfig sets field value
func (o *Policy) SetSubscribeConfig(v bool) {
	o.SubscribeConfig = v
}

func (o Policy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Policy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Connect"] = o.Connect
	toSerialize["PublishState"] = o.PublishState
	toSerialize["PublishEvents"] = o.PublishEvents
	toSerialize["PublishEventsRegex"] = o.PublishEventsRegex
	toSerialize["PublishLoopback"] = o.PublishLoopback
	toSerialize["SubscribeCommand"] = o.SubscribeCommand
	toSerialize["SubscribeCommandRegex"] = o.SubscribeCommandRegex
	toSerialize["SubscribeBroadcast"] = o.SubscribeBroadcast
	toSerialize["SubscribeBroadcastRegex"] = o.SubscribeBroadcastRegex
	toSerialize["SubscribeConfig"] = o.SubscribeConfig
	return toSerialize, nil
}

type NullablePolicy struct {
	value *Policy
	isSet bool
}

func (v NullablePolicy) Get() *Policy {
	return v.value
}

func (v *NullablePolicy) Set(val *Policy) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicy(val *Policy) *NullablePolicy {
	return &NullablePolicy{value: val, isSet: true}
}

func (v NullablePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


