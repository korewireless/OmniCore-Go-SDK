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

// checks if the VaultStatusDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VaultStatusDetails{}

// VaultStatusDetails struct for VaultStatusDetails
type VaultStatusDetails struct {
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	StorageType *string `json:"storageType,omitempty"`
	IsEnabled *bool `json:"isEnabled,omitempty"`
	IsCreated *bool `json:"isCreated,omitempty"`
	Updatedon *string `json:"updatedon,omitempty"`
	CreatedOn *string `json:"createdOn,omitempty"`
	MqttId *string `json:"MqttId,omitempty"`
	RetentionPeriod *int32 `json:"retentionPeriod,omitempty"`
	EncryptionKeyId *int32 `json:"encryptionKeyId,omitempty"`
	IsEncrypted *bool `json:"isEncrypted,omitempty"`
}

// NewVaultStatusDetails instantiates a new VaultStatusDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVaultStatusDetails() *VaultStatusDetails {
	this := VaultStatusDetails{}
	return &this
}

// NewVaultStatusDetailsWithDefaults instantiates a new VaultStatusDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVaultStatusDetailsWithDefaults() *VaultStatusDetails {
	this := VaultStatusDetails{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *VaultStatusDetails) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultStatusDetails) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *VaultStatusDetails) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *VaultStatusDetails) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetStorageType returns the StorageType field value if set, zero value otherwise.
func (o *VaultStatusDetails) GetStorageType() string {
	if o == nil || IsNil(o.StorageType) {
		var ret string
		return ret
	}
	return *o.StorageType
}

// GetStorageTypeOk returns a tuple with the StorageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultStatusDetails) GetStorageTypeOk() (*string, bool) {
	if o == nil || IsNil(o.StorageType) {
		return nil, false
	}
	return o.StorageType, true
}

// HasStorageType returns a boolean if a field has been set.
func (o *VaultStatusDetails) HasStorageType() bool {
	if o != nil && !IsNil(o.StorageType) {
		return true
	}

	return false
}

// SetStorageType gets a reference to the given string and assigns it to the StorageType field.
func (o *VaultStatusDetails) SetStorageType(v string) {
	o.StorageType = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *VaultStatusDetails) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultStatusDetails) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *VaultStatusDetails) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *VaultStatusDetails) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetIsCreated returns the IsCreated field value if set, zero value otherwise.
func (o *VaultStatusDetails) GetIsCreated() bool {
	if o == nil || IsNil(o.IsCreated) {
		var ret bool
		return ret
	}
	return *o.IsCreated
}

// GetIsCreatedOk returns a tuple with the IsCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultStatusDetails) GetIsCreatedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCreated) {
		return nil, false
	}
	return o.IsCreated, true
}

// HasIsCreated returns a boolean if a field has been set.
func (o *VaultStatusDetails) HasIsCreated() bool {
	if o != nil && !IsNil(o.IsCreated) {
		return true
	}

	return false
}

// SetIsCreated gets a reference to the given bool and assigns it to the IsCreated field.
func (o *VaultStatusDetails) SetIsCreated(v bool) {
	o.IsCreated = &v
}

// GetUpdatedon returns the Updatedon field value if set, zero value otherwise.
func (o *VaultStatusDetails) GetUpdatedon() string {
	if o == nil || IsNil(o.Updatedon) {
		var ret string
		return ret
	}
	return *o.Updatedon
}

// GetUpdatedonOk returns a tuple with the Updatedon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultStatusDetails) GetUpdatedonOk() (*string, bool) {
	if o == nil || IsNil(o.Updatedon) {
		return nil, false
	}
	return o.Updatedon, true
}

// HasUpdatedon returns a boolean if a field has been set.
func (o *VaultStatusDetails) HasUpdatedon() bool {
	if o != nil && !IsNil(o.Updatedon) {
		return true
	}

	return false
}

// SetUpdatedon gets a reference to the given string and assigns it to the Updatedon field.
func (o *VaultStatusDetails) SetUpdatedon(v string) {
	o.Updatedon = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *VaultStatusDetails) GetCreatedOn() string {
	if o == nil || IsNil(o.CreatedOn) {
		var ret string
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultStatusDetails) GetCreatedOnOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedOn) {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *VaultStatusDetails) HasCreatedOn() bool {
	if o != nil && !IsNil(o.CreatedOn) {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given string and assigns it to the CreatedOn field.
func (o *VaultStatusDetails) SetCreatedOn(v string) {
	o.CreatedOn = &v
}

// GetMqttId returns the MqttId field value if set, zero value otherwise.
func (o *VaultStatusDetails) GetMqttId() string {
	if o == nil || IsNil(o.MqttId) {
		var ret string
		return ret
	}
	return *o.MqttId
}

// GetMqttIdOk returns a tuple with the MqttId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultStatusDetails) GetMqttIdOk() (*string, bool) {
	if o == nil || IsNil(o.MqttId) {
		return nil, false
	}
	return o.MqttId, true
}

// HasMqttId returns a boolean if a field has been set.
func (o *VaultStatusDetails) HasMqttId() bool {
	if o != nil && !IsNil(o.MqttId) {
		return true
	}

	return false
}

// SetMqttId gets a reference to the given string and assigns it to the MqttId field.
func (o *VaultStatusDetails) SetMqttId(v string) {
	o.MqttId = &v
}

// GetRetentionPeriod returns the RetentionPeriod field value if set, zero value otherwise.
func (o *VaultStatusDetails) GetRetentionPeriod() int32 {
	if o == nil || IsNil(o.RetentionPeriod) {
		var ret int32
		return ret
	}
	return *o.RetentionPeriod
}

// GetRetentionPeriodOk returns a tuple with the RetentionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultStatusDetails) GetRetentionPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.RetentionPeriod) {
		return nil, false
	}
	return o.RetentionPeriod, true
}

// HasRetentionPeriod returns a boolean if a field has been set.
func (o *VaultStatusDetails) HasRetentionPeriod() bool {
	if o != nil && !IsNil(o.RetentionPeriod) {
		return true
	}

	return false
}

// SetRetentionPeriod gets a reference to the given int32 and assigns it to the RetentionPeriod field.
func (o *VaultStatusDetails) SetRetentionPeriod(v int32) {
	o.RetentionPeriod = &v
}

// GetEncryptionKeyId returns the EncryptionKeyId field value if set, zero value otherwise.
func (o *VaultStatusDetails) GetEncryptionKeyId() int32 {
	if o == nil || IsNil(o.EncryptionKeyId) {
		var ret int32
		return ret
	}
	return *o.EncryptionKeyId
}

// GetEncryptionKeyIdOk returns a tuple with the EncryptionKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultStatusDetails) GetEncryptionKeyIdOk() (*int32, bool) {
	if o == nil || IsNil(o.EncryptionKeyId) {
		return nil, false
	}
	return o.EncryptionKeyId, true
}

// HasEncryptionKeyId returns a boolean if a field has been set.
func (o *VaultStatusDetails) HasEncryptionKeyId() bool {
	if o != nil && !IsNil(o.EncryptionKeyId) {
		return true
	}

	return false
}

// SetEncryptionKeyId gets a reference to the given int32 and assigns it to the EncryptionKeyId field.
func (o *VaultStatusDetails) SetEncryptionKeyId(v int32) {
	o.EncryptionKeyId = &v
}

// GetIsEncrypted returns the IsEncrypted field value if set, zero value otherwise.
func (o *VaultStatusDetails) GetIsEncrypted() bool {
	if o == nil || IsNil(o.IsEncrypted) {
		var ret bool
		return ret
	}
	return *o.IsEncrypted
}

// GetIsEncryptedOk returns a tuple with the IsEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultStatusDetails) GetIsEncryptedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEncrypted) {
		return nil, false
	}
	return o.IsEncrypted, true
}

// HasIsEncrypted returns a boolean if a field has been set.
func (o *VaultStatusDetails) HasIsEncrypted() bool {
	if o != nil && !IsNil(o.IsEncrypted) {
		return true
	}

	return false
}

// SetIsEncrypted gets a reference to the given bool and assigns it to the IsEncrypted field.
func (o *VaultStatusDetails) SetIsEncrypted(v bool) {
	o.IsEncrypted = &v
}

func (o VaultStatusDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VaultStatusDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !IsNil(o.StorageType) {
		toSerialize["storageType"] = o.StorageType
	}
	if !IsNil(o.IsEnabled) {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if !IsNil(o.IsCreated) {
		toSerialize["isCreated"] = o.IsCreated
	}
	if !IsNil(o.Updatedon) {
		toSerialize["updatedon"] = o.Updatedon
	}
	if !IsNil(o.CreatedOn) {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if !IsNil(o.MqttId) {
		toSerialize["MqttId"] = o.MqttId
	}
	if !IsNil(o.RetentionPeriod) {
		toSerialize["retentionPeriod"] = o.RetentionPeriod
	}
	if !IsNil(o.EncryptionKeyId) {
		toSerialize["encryptionKeyId"] = o.EncryptionKeyId
	}
	if !IsNil(o.IsEncrypted) {
		toSerialize["isEncrypted"] = o.IsEncrypted
	}
	return toSerialize, nil
}

type NullableVaultStatusDetails struct {
	value *VaultStatusDetails
	isSet bool
}

func (v NullableVaultStatusDetails) Get() *VaultStatusDetails {
	return v.value
}

func (v *NullableVaultStatusDetails) Set(val *VaultStatusDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableVaultStatusDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableVaultStatusDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVaultStatusDetails(val *VaultStatusDetails) *NullableVaultStatusDetails {
	return &NullableVaultStatusDetails{value: val, isSet: true}
}

func (v NullableVaultStatusDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVaultStatusDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


