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

// checks if the VaultEncryptionKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VaultEncryptionKey{}

// VaultEncryptionKey struct for VaultEncryptionKey
type VaultEncryptionKey struct {
	Subscription *string `json:"subscription,omitempty"`
	Name *string `json:"name,omitempty"`
	Id *string `json:"id,omitempty"`
	Key *string `json:"key,omitempty"`
}

// NewVaultEncryptionKey instantiates a new VaultEncryptionKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVaultEncryptionKey() *VaultEncryptionKey {
	this := VaultEncryptionKey{}
	return &this
}

// NewVaultEncryptionKeyWithDefaults instantiates a new VaultEncryptionKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVaultEncryptionKeyWithDefaults() *VaultEncryptionKey {
	this := VaultEncryptionKey{}
	return &this
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *VaultEncryptionKey) GetSubscription() string {
	if o == nil || IsNil(o.Subscription) {
		var ret string
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultEncryptionKey) GetSubscriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *VaultEncryptionKey) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given string and assigns it to the Subscription field.
func (o *VaultEncryptionKey) SetSubscription(v string) {
	o.Subscription = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VaultEncryptionKey) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultEncryptionKey) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VaultEncryptionKey) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VaultEncryptionKey) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VaultEncryptionKey) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultEncryptionKey) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VaultEncryptionKey) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VaultEncryptionKey) SetId(v string) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *VaultEncryptionKey) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultEncryptionKey) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *VaultEncryptionKey) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *VaultEncryptionKey) SetKey(v string) {
	o.Key = &v
}

func (o VaultEncryptionKey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VaultEncryptionKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subscription) {
		toSerialize["subscription"] = o.Subscription
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	return toSerialize, nil
}

type NullableVaultEncryptionKey struct {
	value *VaultEncryptionKey
	isSet bool
}

func (v NullableVaultEncryptionKey) Get() *VaultEncryptionKey {
	return v.value
}

func (v *NullableVaultEncryptionKey) Set(val *VaultEncryptionKey) {
	v.value = val
	v.isSet = true
}

func (v NullableVaultEncryptionKey) IsSet() bool {
	return v.isSet
}

func (v *NullableVaultEncryptionKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVaultEncryptionKey(val *VaultEncryptionKey) *NullableVaultEncryptionKey {
	return &NullableVaultEncryptionKey{value: val, isSet: true}
}

func (v NullableVaultEncryptionKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVaultEncryptionKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


