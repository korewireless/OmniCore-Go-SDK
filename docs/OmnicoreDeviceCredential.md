# OmnicoreDeviceCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationTime** | Pointer to **string** | ExpirationTime: [Optional] The time at which this credential becomes invalid. This credential will be ignored for new client authentication requests after this timestamp; however, it will not be automatically deleted. | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**PublicKey** | Pointer to [**OmnicorePublicKeyCredential**](OmnicorePublicKeyCredential.md) |  | [optional] 

## Methods

### NewOmnicoreDeviceCredential

`func NewOmnicoreDeviceCredential() *OmnicoreDeviceCredential`

NewOmnicoreDeviceCredential instantiates a new OmnicoreDeviceCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOmnicoreDeviceCredentialWithDefaults

`func NewOmnicoreDeviceCredentialWithDefaults() *OmnicoreDeviceCredential`

NewOmnicoreDeviceCredentialWithDefaults instantiates a new OmnicoreDeviceCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationTime

`func (o *OmnicoreDeviceCredential) GetExpirationTime() string`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *OmnicoreDeviceCredential) GetExpirationTimeOk() (*string, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *OmnicoreDeviceCredential) SetExpirationTime(v string)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *OmnicoreDeviceCredential) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetId

`func (o *OmnicoreDeviceCredential) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OmnicoreDeviceCredential) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OmnicoreDeviceCredential) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OmnicoreDeviceCredential) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPublicKey

`func (o *OmnicoreDeviceCredential) GetPublicKey() OmnicorePublicKeyCredential`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *OmnicoreDeviceCredential) GetPublicKeyOk() (*OmnicorePublicKeyCredential, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *OmnicoreDeviceCredential) SetPublicKey(v OmnicorePublicKeyCredential)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *OmnicoreDeviceCredential) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


