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

// checks if the NotificationConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationConfig{}

// NotificationConfig struct for NotificationConfig
type NotificationConfig struct {
	// PubsubTopicName: A Cloud Pub/Sub topic name. For example, `projects/myProject/topics/deviceEvents`.
	PubsubTopicName *string `json:"pubsubTopicName,omitempty"`
}

// NewNotificationConfig instantiates a new NotificationConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationConfig() *NotificationConfig {
	this := NotificationConfig{}
	return &this
}

// NewNotificationConfigWithDefaults instantiates a new NotificationConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationConfigWithDefaults() *NotificationConfig {
	this := NotificationConfig{}
	return &this
}

// GetPubsubTopicName returns the PubsubTopicName field value if set, zero value otherwise.
func (o *NotificationConfig) GetPubsubTopicName() string {
	if o == nil || IsNil(o.PubsubTopicName) {
		var ret string
		return ret
	}
	return *o.PubsubTopicName
}

// GetPubsubTopicNameOk returns a tuple with the PubsubTopicName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationConfig) GetPubsubTopicNameOk() (*string, bool) {
	if o == nil || IsNil(o.PubsubTopicName) {
		return nil, false
	}
	return o.PubsubTopicName, true
}

// HasPubsubTopicName returns a boolean if a field has been set.
func (o *NotificationConfig) HasPubsubTopicName() bool {
	if o != nil && !IsNil(o.PubsubTopicName) {
		return true
	}

	return false
}

// SetPubsubTopicName gets a reference to the given string and assigns it to the PubsubTopicName field.
func (o *NotificationConfig) SetPubsubTopicName(v string) {
	o.PubsubTopicName = &v
}

func (o NotificationConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PubsubTopicName) {
		toSerialize["pubsubTopicName"] = o.PubsubTopicName
	}
	return toSerialize, nil
}

type NullableNotificationConfig struct {
	value *NotificationConfig
	isSet bool
}

func (v NullableNotificationConfig) Get() *NotificationConfig {
	return v.value
}

func (v *NullableNotificationConfig) Set(val *NotificationConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationConfig(val *NotificationConfig) *NullableNotificationConfig {
	return &NullableNotificationConfig{value: val, isSet: true}
}

func (v NullableNotificationConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


