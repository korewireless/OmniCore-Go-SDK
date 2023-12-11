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

// checks if the EnableEncryptionBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnableEncryptionBody{}

// EnableEncryptionBody struct for EnableEncryptionBody
type EnableEncryptionBody struct {
	Action *string `json:"action,omitempty"`
	IsEncrypted *bool `json:"isEncrypted,omitempty"`
	EncryptionKeyId *int32 `json:"encryptionKeyId,omitempty"`
}

// NewEnableEncryptionBody instantiates a new EnableEncryptionBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnableEncryptionBody() *EnableEncryptionBody {
	this := EnableEncryptionBody{}
	return &this
}

// NewEnableEncryptionBodyWithDefaults instantiates a new EnableEncryptionBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnableEncryptionBodyWithDefaults() *EnableEncryptionBody {
	this := EnableEncryptionBody{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *EnableEncryptionBody) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableEncryptionBody) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *EnableEncryptionBody) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *EnableEncryptionBody) SetAction(v string) {
	o.Action = &v
}

// GetIsEncrypted returns the IsEncrypted field value if set, zero value otherwise.
func (o *EnableEncryptionBody) GetIsEncrypted() bool {
	if o == nil || IsNil(o.IsEncrypted) {
		var ret bool
		return ret
	}
	return *o.IsEncrypted
}

// GetIsEncryptedOk returns a tuple with the IsEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableEncryptionBody) GetIsEncryptedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEncrypted) {
		return nil, false
	}
	return o.IsEncrypted, true
}

// HasIsEncrypted returns a boolean if a field has been set.
func (o *EnableEncryptionBody) HasIsEncrypted() bool {
	if o != nil && !IsNil(o.IsEncrypted) {
		return true
	}

	return false
}

// SetIsEncrypted gets a reference to the given bool and assigns it to the IsEncrypted field.
func (o *EnableEncryptionBody) SetIsEncrypted(v bool) {
	o.IsEncrypted = &v
}

// GetEncryptionKeyId returns the EncryptionKeyId field value if set, zero value otherwise.
func (o *EnableEncryptionBody) GetEncryptionKeyId() int32 {
	if o == nil || IsNil(o.EncryptionKeyId) {
		var ret int32
		return ret
	}
	return *o.EncryptionKeyId
}

// GetEncryptionKeyIdOk returns a tuple with the EncryptionKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableEncryptionBody) GetEncryptionKeyIdOk() (*int32, bool) {
	if o == nil || IsNil(o.EncryptionKeyId) {
		return nil, false
	}
	return o.EncryptionKeyId, true
}

// HasEncryptionKeyId returns a boolean if a field has been set.
func (o *EnableEncryptionBody) HasEncryptionKeyId() bool {
	if o != nil && !IsNil(o.EncryptionKeyId) {
		return true
	}

	return false
}

// SetEncryptionKeyId gets a reference to the given int32 and assigns it to the EncryptionKeyId field.
func (o *EnableEncryptionBody) SetEncryptionKeyId(v int32) {
	o.EncryptionKeyId = &v
}

func (o EnableEncryptionBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnableEncryptionBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.IsEncrypted) {
		toSerialize["isEncrypted"] = o.IsEncrypted
	}
	if !IsNil(o.EncryptionKeyId) {
		toSerialize["encryptionKeyId"] = o.EncryptionKeyId
	}
	return toSerialize, nil
}

type NullableEnableEncryptionBody struct {
	value *EnableEncryptionBody
	isSet bool
}

func (v NullableEnableEncryptionBody) Get() *EnableEncryptionBody {
	return v.value
}

func (v *NullableEnableEncryptionBody) Set(val *EnableEncryptionBody) {
	v.value = val
	v.isSet = true
}

func (v NullableEnableEncryptionBody) IsSet() bool {
	return v.isSet
}

func (v *NullableEnableEncryptionBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnableEncryptionBody(val *EnableEncryptionBody) *NullableEnableEncryptionBody {
	return &NullableEnableEncryptionBody{value: val, isSet: true}
}

func (v NullableEnableEncryptionBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnableEncryptionBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


