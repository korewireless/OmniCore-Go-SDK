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

// checks if the Audit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Audit{}

// Audit struct for Audit
type Audit struct {
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	Operation *string `json:"operation,omitempty"`
	Actor *string `json:"actor,omitempty"`
	Updatedon *string `json:"updatedon,omitempty"`
}

// NewAudit instantiates a new Audit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAudit() *Audit {
	this := Audit{}
	return &this
}

// NewAuditWithDefaults instantiates a new Audit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditWithDefaults() *Audit {
	this := Audit{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *Audit) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audit) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *Audit) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *Audit) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *Audit) GetOperation() string {
	if o == nil || IsNil(o.Operation) {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audit) GetOperationOk() (*string, bool) {
	if o == nil || IsNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *Audit) HasOperation() bool {
	if o != nil && !IsNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *Audit) SetOperation(v string) {
	o.Operation = &v
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *Audit) GetActor() string {
	if o == nil || IsNil(o.Actor) {
		var ret string
		return ret
	}
	return *o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audit) GetActorOk() (*string, bool) {
	if o == nil || IsNil(o.Actor) {
		return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *Audit) HasActor() bool {
	if o != nil && !IsNil(o.Actor) {
		return true
	}

	return false
}

// SetActor gets a reference to the given string and assigns it to the Actor field.
func (o *Audit) SetActor(v string) {
	o.Actor = &v
}

// GetUpdatedon returns the Updatedon field value if set, zero value otherwise.
func (o *Audit) GetUpdatedon() string {
	if o == nil || IsNil(o.Updatedon) {
		var ret string
		return ret
	}
	return *o.Updatedon
}

// GetUpdatedonOk returns a tuple with the Updatedon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audit) GetUpdatedonOk() (*string, bool) {
	if o == nil || IsNil(o.Updatedon) {
		return nil, false
	}
	return o.Updatedon, true
}

// HasUpdatedon returns a boolean if a field has been set.
func (o *Audit) HasUpdatedon() bool {
	if o != nil && !IsNil(o.Updatedon) {
		return true
	}

	return false
}

// SetUpdatedon gets a reference to the given string and assigns it to the Updatedon field.
func (o *Audit) SetUpdatedon(v string) {
	o.Updatedon = &v
}

func (o Audit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Audit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !IsNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	if !IsNil(o.Actor) {
		toSerialize["actor"] = o.Actor
	}
	if !IsNil(o.Updatedon) {
		toSerialize["updatedon"] = o.Updatedon
	}
	return toSerialize, nil
}

type NullableAudit struct {
	value *Audit
	isSet bool
}

func (v NullableAudit) Get() *Audit {
	return v.value
}

func (v *NullableAudit) Set(val *Audit) {
	v.value = val
	v.isSet = true
}

func (v NullableAudit) IsSet() bool {
	return v.isSet
}

func (v *NullableAudit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAudit(val *Audit) *NullableAudit {
	return &NullableAudit{value: val, isSet: true}
}

func (v NullableAudit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAudit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


