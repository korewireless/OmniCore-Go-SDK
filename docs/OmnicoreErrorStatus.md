# OmnicoreErrorStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewOmnicoreErrorStatus

`func NewOmnicoreErrorStatus() *OmnicoreErrorStatus`

NewOmnicoreErrorStatus instantiates a new OmnicoreErrorStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOmnicoreErrorStatusWithDefaults

`func NewOmnicoreErrorStatusWithDefaults() *OmnicoreErrorStatus`

NewOmnicoreErrorStatusWithDefaults instantiates a new OmnicoreErrorStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *OmnicoreErrorStatus) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OmnicoreErrorStatus) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OmnicoreErrorStatus) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *OmnicoreErrorStatus) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetails

`func (o *OmnicoreErrorStatus) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *OmnicoreErrorStatus) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *OmnicoreErrorStatus) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *OmnicoreErrorStatus) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetMessage

`func (o *OmnicoreErrorStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OmnicoreErrorStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OmnicoreErrorStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OmnicoreErrorStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


