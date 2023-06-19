# Sink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Subscription** | Pointer to **string** |  | [optional] [readonly] 
**Sink** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**ListSinksSinksInnerConfig**](ListSinksSinksInnerConfig.md) |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] [readonly] 
**Createdon** | Pointer to **string** |  | [optional] [readonly] 
**Updatedon** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewSink

`func NewSink() *Sink`

NewSink instantiates a new Sink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSinkWithDefaults

`func NewSinkWithDefaults() *Sink`

NewSinkWithDefaults instantiates a new Sink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Sink) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Sink) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Sink) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Sink) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSubscription

`func (o *Sink) GetSubscription() string`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *Sink) GetSubscriptionOk() (*string, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *Sink) SetSubscription(v string)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *Sink) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetSink

`func (o *Sink) GetSink() string`

GetSink returns the Sink field if non-nil, zero value otherwise.

### GetSinkOk

`func (o *Sink) GetSinkOk() (*string, bool)`

GetSinkOk returns a tuple with the Sink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSink

`func (o *Sink) SetSink(v string)`

SetSink sets Sink field to given value.

### HasSink

`func (o *Sink) HasSink() bool`

HasSink returns a boolean if a field has been set.

### GetConfig

`func (o *Sink) GetConfig() ListSinksSinksInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Sink) GetConfigOk() (*ListSinksSinksInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Sink) SetConfig(v ListSinksSinksInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Sink) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *Sink) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Sink) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Sink) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Sink) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedon

`func (o *Sink) GetCreatedon() string`

GetCreatedon returns the Createdon field if non-nil, zero value otherwise.

### GetCreatedonOk

`func (o *Sink) GetCreatedonOk() (*string, bool)`

GetCreatedonOk returns a tuple with the Createdon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedon

`func (o *Sink) SetCreatedon(v string)`

SetCreatedon sets Createdon field to given value.

### HasCreatedon

`func (o *Sink) HasCreatedon() bool`

HasCreatedon returns a boolean if a field has been set.

### GetUpdatedon

`func (o *Sink) GetUpdatedon() string`

GetUpdatedon returns the Updatedon field if non-nil, zero value otherwise.

### GetUpdatedonOk

`func (o *Sink) GetUpdatedonOk() (*string, bool)`

GetUpdatedonOk returns a tuple with the Updatedon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedon

`func (o *Sink) SetUpdatedon(v string)`

SetUpdatedon sets Updatedon field to given value.

### HasUpdatedon

`func (o *Sink) HasUpdatedon() bool`

HasUpdatedon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


