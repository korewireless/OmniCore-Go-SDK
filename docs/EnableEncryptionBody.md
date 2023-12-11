# EnableEncryptionBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**IsEncrypted** | Pointer to **bool** |  | [optional] 
**EncryptionKeyId** | Pointer to **int32** |  | [optional] 

## Methods

### NewEnableEncryptionBody

`func NewEnableEncryptionBody() *EnableEncryptionBody`

NewEnableEncryptionBody instantiates a new EnableEncryptionBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableEncryptionBodyWithDefaults

`func NewEnableEncryptionBodyWithDefaults() *EnableEncryptionBody`

NewEnableEncryptionBodyWithDefaults instantiates a new EnableEncryptionBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *EnableEncryptionBody) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *EnableEncryptionBody) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *EnableEncryptionBody) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *EnableEncryptionBody) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetIsEncrypted

`func (o *EnableEncryptionBody) GetIsEncrypted() bool`

GetIsEncrypted returns the IsEncrypted field if non-nil, zero value otherwise.

### GetIsEncryptedOk

`func (o *EnableEncryptionBody) GetIsEncryptedOk() (*bool, bool)`

GetIsEncryptedOk returns a tuple with the IsEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEncrypted

`func (o *EnableEncryptionBody) SetIsEncrypted(v bool)`

SetIsEncrypted sets IsEncrypted field to given value.

### HasIsEncrypted

`func (o *EnableEncryptionBody) HasIsEncrypted() bool`

HasIsEncrypted returns a boolean if a field has been set.

### GetEncryptionKeyId

`func (o *EnableEncryptionBody) GetEncryptionKeyId() int32`

GetEncryptionKeyId returns the EncryptionKeyId field if non-nil, zero value otherwise.

### GetEncryptionKeyIdOk

`func (o *EnableEncryptionBody) GetEncryptionKeyIdOk() (*int32, bool)`

GetEncryptionKeyIdOk returns a tuple with the EncryptionKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeyId

`func (o *EnableEncryptionBody) SetEncryptionKeyId(v int32)`

SetEncryptionKeyId sets EncryptionKeyId field to given value.

### HasEncryptionKeyId

`func (o *EnableEncryptionBody) HasEncryptionKeyId() bool`

HasEncryptionKeyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


