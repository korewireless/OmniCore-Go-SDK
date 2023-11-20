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

// checks if the EnableVault type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnableVault{}

// EnableVault struct for EnableVault
type EnableVault struct {
	Type string `json:"type"`
	Action string `json:"action"`
}

// NewEnableVault instantiates a new EnableVault object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnableVault(type_ string, action string) *EnableVault {
	this := EnableVault{}
	this.Type = type_
	this.Action = action
	return &this
}

// NewEnableVaultWithDefaults instantiates a new EnableVault object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnableVaultWithDefaults() *EnableVault {
	this := EnableVault{}
	return &this
}

// GetType returns the Type field value
func (o *EnableVault) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EnableVault) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EnableVault) SetType(v string) {
	o.Type = v
}

// GetAction returns the Action field value
func (o *EnableVault) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *EnableVault) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *EnableVault) SetAction(v string) {
	o.Action = v
}

func (o EnableVault) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnableVault) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["action"] = o.Action
	return toSerialize, nil
}

type NullableEnableVault struct {
	value *EnableVault
	isSet bool
}

func (v NullableEnableVault) Get() *EnableVault {
	return v.value
}

func (v *NullableEnableVault) Set(val *EnableVault) {
	v.value = val
	v.isSet = true
}

func (v NullableEnableVault) IsSet() bool {
	return v.isSet
}

func (v *NullableEnableVault) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnableVault(val *EnableVault) *NullableEnableVault {
	return &NullableEnableVault{value: val, isSet: true}
}

func (v NullableEnableVault) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnableVault) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


