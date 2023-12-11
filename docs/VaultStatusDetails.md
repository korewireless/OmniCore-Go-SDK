# VaultStatusDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | Pointer to **string** |  | [optional] 
**StorageType** | Pointer to **string** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**IsCreated** | Pointer to **bool** |  | [optional] 
**Updatedon** | Pointer to **string** |  | [optional] 
**CreatedOn** | Pointer to **string** |  | [optional] 
**MqttId** | Pointer to **string** |  | [optional] 
**RetentionPeriod** | Pointer to **int32** |  | [optional] 
**EncryptionKeyId** | Pointer to **int32** |  | [optional] 
**IsEncrypted** | Pointer to **bool** |  | [optional] 

## Methods

### NewVaultStatusDetails

`func NewVaultStatusDetails() *VaultStatusDetails`

NewVaultStatusDetails instantiates a new VaultStatusDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultStatusDetailsWithDefaults

`func NewVaultStatusDetailsWithDefaults() *VaultStatusDetails`

NewVaultStatusDetailsWithDefaults instantiates a new VaultStatusDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *VaultStatusDetails) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *VaultStatusDetails) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *VaultStatusDetails) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *VaultStatusDetails) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetStorageType

`func (o *VaultStatusDetails) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *VaultStatusDetails) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *VaultStatusDetails) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *VaultStatusDetails) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetIsEnabled

`func (o *VaultStatusDetails) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *VaultStatusDetails) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *VaultStatusDetails) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *VaultStatusDetails) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsCreated

`func (o *VaultStatusDetails) GetIsCreated() bool`

GetIsCreated returns the IsCreated field if non-nil, zero value otherwise.

### GetIsCreatedOk

`func (o *VaultStatusDetails) GetIsCreatedOk() (*bool, bool)`

GetIsCreatedOk returns a tuple with the IsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCreated

`func (o *VaultStatusDetails) SetIsCreated(v bool)`

SetIsCreated sets IsCreated field to given value.

### HasIsCreated

`func (o *VaultStatusDetails) HasIsCreated() bool`

HasIsCreated returns a boolean if a field has been set.

### GetUpdatedon

`func (o *VaultStatusDetails) GetUpdatedon() string`

GetUpdatedon returns the Updatedon field if non-nil, zero value otherwise.

### GetUpdatedonOk

`func (o *VaultStatusDetails) GetUpdatedonOk() (*string, bool)`

GetUpdatedonOk returns a tuple with the Updatedon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedon

`func (o *VaultStatusDetails) SetUpdatedon(v string)`

SetUpdatedon sets Updatedon field to given value.

### HasUpdatedon

`func (o *VaultStatusDetails) HasUpdatedon() bool`

HasUpdatedon returns a boolean if a field has been set.

### GetCreatedOn

`func (o *VaultStatusDetails) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *VaultStatusDetails) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *VaultStatusDetails) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *VaultStatusDetails) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetMqttId

`func (o *VaultStatusDetails) GetMqttId() string`

GetMqttId returns the MqttId field if non-nil, zero value otherwise.

### GetMqttIdOk

`func (o *VaultStatusDetails) GetMqttIdOk() (*string, bool)`

GetMqttIdOk returns a tuple with the MqttId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttId

`func (o *VaultStatusDetails) SetMqttId(v string)`

SetMqttId sets MqttId field to given value.

### HasMqttId

`func (o *VaultStatusDetails) HasMqttId() bool`

HasMqttId returns a boolean if a field has been set.

### GetRetentionPeriod

`func (o *VaultStatusDetails) GetRetentionPeriod() int32`

GetRetentionPeriod returns the RetentionPeriod field if non-nil, zero value otherwise.

### GetRetentionPeriodOk

`func (o *VaultStatusDetails) GetRetentionPeriodOk() (*int32, bool)`

GetRetentionPeriodOk returns a tuple with the RetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriod

`func (o *VaultStatusDetails) SetRetentionPeriod(v int32)`

SetRetentionPeriod sets RetentionPeriod field to given value.

### HasRetentionPeriod

`func (o *VaultStatusDetails) HasRetentionPeriod() bool`

HasRetentionPeriod returns a boolean if a field has been set.

### GetEncryptionKeyId

`func (o *VaultStatusDetails) GetEncryptionKeyId() int32`

GetEncryptionKeyId returns the EncryptionKeyId field if non-nil, zero value otherwise.

### GetEncryptionKeyIdOk

`func (o *VaultStatusDetails) GetEncryptionKeyIdOk() (*int32, bool)`

GetEncryptionKeyIdOk returns a tuple with the EncryptionKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeyId

`func (o *VaultStatusDetails) SetEncryptionKeyId(v int32)`

SetEncryptionKeyId sets EncryptionKeyId field to given value.

### HasEncryptionKeyId

`func (o *VaultStatusDetails) HasEncryptionKeyId() bool`

HasEncryptionKeyId returns a boolean if a field has been set.

### GetIsEncrypted

`func (o *VaultStatusDetails) GetIsEncrypted() bool`

GetIsEncrypted returns the IsEncrypted field if non-nil, zero value otherwise.

### GetIsEncryptedOk

`func (o *VaultStatusDetails) GetIsEncryptedOk() (*bool, bool)`

GetIsEncryptedOk returns a tuple with the IsEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEncrypted

`func (o *VaultStatusDetails) SetIsEncrypted(v bool)`

SetIsEncrypted sets IsEncrypted field to given value.

### HasIsEncrypted

`func (o *VaultStatusDetails) HasIsEncrypted() bool`

HasIsEncrypted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


