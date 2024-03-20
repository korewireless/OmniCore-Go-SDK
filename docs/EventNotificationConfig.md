# EventNotificationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PubsubTopicName** | Pointer to **string** | PubsubTopicName: A Topic name. For example, &#x60;projects/myProject/topics/deviceEvents&#x60;. | [optional] 
**SubfolderMatches** | Pointer to **string** | SubfolderMatches: If the subfolder name matches this string exactly, this configuration will be used. The string must not include the leading &#39;/&#39; character. If empty, all strings are matched. This field is used only for telemetry events; subfolders are not supported for state changes. | [optional] 
**RoleArn** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Stream** | Pointer to **string** |  | [optional] 

## Methods

### NewEventNotificationConfig

`func NewEventNotificationConfig() *EventNotificationConfig`

NewEventNotificationConfig instantiates a new EventNotificationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventNotificationConfigWithDefaults

`func NewEventNotificationConfigWithDefaults() *EventNotificationConfig`

NewEventNotificationConfigWithDefaults instantiates a new EventNotificationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubsubTopicName

`func (o *EventNotificationConfig) GetPubsubTopicName() string`

GetPubsubTopicName returns the PubsubTopicName field if non-nil, zero value otherwise.

### GetPubsubTopicNameOk

`func (o *EventNotificationConfig) GetPubsubTopicNameOk() (*string, bool)`

GetPubsubTopicNameOk returns a tuple with the PubsubTopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubsubTopicName

`func (o *EventNotificationConfig) SetPubsubTopicName(v string)`

SetPubsubTopicName sets PubsubTopicName field to given value.

### HasPubsubTopicName

`func (o *EventNotificationConfig) HasPubsubTopicName() bool`

HasPubsubTopicName returns a boolean if a field has been set.

### GetSubfolderMatches

`func (o *EventNotificationConfig) GetSubfolderMatches() string`

GetSubfolderMatches returns the SubfolderMatches field if non-nil, zero value otherwise.

### GetSubfolderMatchesOk

`func (o *EventNotificationConfig) GetSubfolderMatchesOk() (*string, bool)`

GetSubfolderMatchesOk returns a tuple with the SubfolderMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubfolderMatches

`func (o *EventNotificationConfig) SetSubfolderMatches(v string)`

SetSubfolderMatches sets SubfolderMatches field to given value.

### HasSubfolderMatches

`func (o *EventNotificationConfig) HasSubfolderMatches() bool`

HasSubfolderMatches returns a boolean if a field has been set.

### GetRoleArn

`func (o *EventNotificationConfig) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *EventNotificationConfig) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *EventNotificationConfig) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *EventNotificationConfig) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### GetRegion

`func (o *EventNotificationConfig) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *EventNotificationConfig) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *EventNotificationConfig) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *EventNotificationConfig) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStream

`func (o *EventNotificationConfig) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *EventNotificationConfig) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *EventNotificationConfig) SetStream(v string)`

SetStream sets Stream field to given value.

### HasStream

`func (o *EventNotificationConfig) HasStream() bool`

HasStream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


