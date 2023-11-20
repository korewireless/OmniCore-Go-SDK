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

// checks if the VaultConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VaultConfiguration{}

// VaultConfiguration struct for VaultConfiguration
type VaultConfiguration struct {
	Id *int32 `json:"id,omitempty"`
	Subscription *string `json:"subscription,omitempty"`
	Type *string `json:"type,omitempty"`
	Data *string `json:"data,omitempty"`
}

// NewVaultConfiguration instantiates a new VaultConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVaultConfiguration() *VaultConfiguration {
	this := VaultConfiguration{}
	return &this
}

// NewVaultConfigurationWithDefaults instantiates a new VaultConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVaultConfigurationWithDefaults() *VaultConfiguration {
	this := VaultConfiguration{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VaultConfiguration) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultConfiguration) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VaultConfiguration) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *VaultConfiguration) SetId(v int32) {
	o.Id = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *VaultConfiguration) GetSubscription() string {
	if o == nil || IsNil(o.Subscription) {
		var ret string
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultConfiguration) GetSubscriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *VaultConfiguration) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given string and assigns it to the Subscription field.
func (o *VaultConfiguration) SetSubscription(v string) {
	o.Subscription = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VaultConfiguration) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultConfiguration) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VaultConfiguration) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VaultConfiguration) SetType(v string) {
	o.Type = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *VaultConfiguration) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultConfiguration) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *VaultConfiguration) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *VaultConfiguration) SetData(v string) {
	o.Data = &v
}

func (o VaultConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VaultConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Subscription) {
		toSerialize["subscription"] = o.Subscription
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableVaultConfiguration struct {
	value *VaultConfiguration
	isSet bool
}

func (v NullableVaultConfiguration) Get() *VaultConfiguration {
	return v.value
}

func (v *NullableVaultConfiguration) Set(val *VaultConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableVaultConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableVaultConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVaultConfiguration(val *VaultConfiguration) *NullableVaultConfiguration {
	return &NullableVaultConfiguration{value: val, isSet: true}
}

func (v NullableVaultConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVaultConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


