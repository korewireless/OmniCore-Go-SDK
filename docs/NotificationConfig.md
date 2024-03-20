# NotificationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PubsubTopicName** | Pointer to **string** | PubsubTopicName: A Topic name. For example, &#x60;projects/myProject/topics/deviceEvents&#x60;. | [optional] 
**RoleArn** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Stream** | Pointer to **string** |  | [optional] 

## Methods

### NewNotificationConfig

`func NewNotificationConfig() *NotificationConfig`

NewNotificationConfig instantiates a new NotificationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationConfigWithDefaults

`func NewNotificationConfigWithDefaults() *NotificationConfig`

NewNotificationConfigWithDefaults instantiates a new NotificationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubsubTopicName

`func (o *NotificationConfig) GetPubsubTopicName() string`

GetPubsubTopicName returns the PubsubTopicName field if non-nil, zero value otherwise.

### GetPubsubTopicNameOk

`func (o *NotificationConfig) GetPubsubTopicNameOk() (*string, bool)`

GetPubsubTopicNameOk returns a tuple with the PubsubTopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubsubTopicName

`func (o *NotificationConfig) SetPubsubTopicName(v string)`

SetPubsubTopicName sets PubsubTopicName field to given value.

### HasPubsubTopicName

`func (o *NotificationConfig) HasPubsubTopicName() bool`

HasPubsubTopicName returns a boolean if a field has been set.

### GetRoleArn

`func (o *NotificationConfig) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *NotificationConfig) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *NotificationConfig) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *NotificationConfig) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### GetRegion

`func (o *NotificationConfig) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NotificationConfig) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NotificationConfig) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *NotificationConfig) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStream

`func (o *NotificationConfig) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *NotificationConfig) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *NotificationConfig) SetStream(v string)`

SetStream sets Stream field to given value.

### HasStream

`func (o *NotificationConfig) HasStream() bool`

HasStream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


