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

// checks if the OmnicoreEventNotificationConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OmnicoreEventNotificationConfig{}

// OmnicoreEventNotificationConfig struct for OmnicoreEventNotificationConfig
type OmnicoreEventNotificationConfig struct {
	// PubsubTopicName: A Cloud Pub/Sub topic name. For example, `projects/myProject/topics/deviceEvents`.
	PubsubTopicName *string `json:"pubsubTopicName,omitempty"`
	// SubfolderMatches: If the subfolder name matches this string exactly, this configuration will be used. The string must not include the leading '/' character. If empty, all strings are matched. This field is used only for telemetry events; subfolders are not supported for state changes.
	SubfolderMatches *string `json:"subfolderMatches,omitempty"`
}

// NewOmnicoreEventNotificationConfig instantiates a new OmnicoreEventNotificationConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOmnicoreEventNotificationConfig() *OmnicoreEventNotificationConfig {
	this := OmnicoreEventNotificationConfig{}
	return &this
}

// NewOmnicoreEventNotificationConfigWithDefaults instantiates a new OmnicoreEventNotificationConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOmnicoreEventNotificationConfigWithDefaults() *OmnicoreEventNotificationConfig {
	this := OmnicoreEventNotificationConfig{}
	return &this
}

// GetPubsubTopicName returns the PubsubTopicName field value if set, zero value otherwise.
func (o *OmnicoreEventNotificationConfig) GetPubsubTopicName() string {
	if o == nil || IsNil(o.PubsubTopicName) {
		var ret string
		return ret
	}
	return *o.PubsubTopicName
}

// GetPubsubTopicNameOk returns a tuple with the PubsubTopicName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreEventNotificationConfig) GetPubsubTopicNameOk() (*string, bool) {
	if o == nil || IsNil(o.PubsubTopicName) {
		return nil, false
	}
	return o.PubsubTopicName, true
}

// HasPubsubTopicName returns a boolean if a field has been set.
func (o *OmnicoreEventNotificationConfig) HasPubsubTopicName() bool {
	if o != nil && !IsNil(o.PubsubTopicName) {
		return true
	}

	return false
}

// SetPubsubTopicName gets a reference to the given string and assigns it to the PubsubTopicName field.
func (o *OmnicoreEventNotificationConfig) SetPubsubTopicName(v string) {
	o.PubsubTopicName = &v
}

// GetSubfolderMatches returns the SubfolderMatches field value if set, zero value otherwise.
func (o *OmnicoreEventNotificationConfig) GetSubfolderMatches() string {
	if o == nil || IsNil(o.SubfolderMatches) {
		var ret string
		return ret
	}
	return *o.SubfolderMatches
}

// GetSubfolderMatchesOk returns a tuple with the SubfolderMatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicoreEventNotificationConfig) GetSubfolderMatchesOk() (*string, bool) {
	if o == nil || IsNil(o.SubfolderMatches) {
		return nil, false
	}
	return o.SubfolderMatches, true
}

// HasSubfolderMatches returns a boolean if a field has been set.
func (o *OmnicoreEventNotificationConfig) HasSubfolderMatches() bool {
	if o != nil && !IsNil(o.SubfolderMatches) {
		return true
	}

	return false
}

// SetSubfolderMatches gets a reference to the given string and assigns it to the SubfolderMatches field.
func (o *OmnicoreEventNotificationConfig) SetSubfolderMatches(v string) {
	o.SubfolderMatches = &v
}

func (o OmnicoreEventNotificationConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OmnicoreEventNotificationConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PubsubTopicName) {
		toSerialize["pubsubTopicName"] = o.PubsubTopicName
	}
	if !IsNil(o.SubfolderMatches) {
		toSerialize["subfolderMatches"] = o.SubfolderMatches
	}
	return toSerialize, nil
}

type NullableOmnicoreEventNotificationConfig struct {
	value *OmnicoreEventNotificationConfig
	isSet bool
}

func (v NullableOmnicoreEventNotificationConfig) Get() *OmnicoreEventNotificationConfig {
	return v.value
}

func (v *NullableOmnicoreEventNotificationConfig) Set(val *OmnicoreEventNotificationConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableOmnicoreEventNotificationConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableOmnicoreEventNotificationConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOmnicoreEventNotificationConfig(val *OmnicoreEventNotificationConfig) *NullableOmnicoreEventNotificationConfig {
	return &NullableOmnicoreEventNotificationConfig{value: val, isSet: true}
}

func (v NullableOmnicoreEventNotificationConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOmnicoreEventNotificationConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


