# OmnicoreStateNotificationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PubsubTopicName** | Pointer to **string** | PubsubTopicName: A Cloud Pub/Sub topic name. For example, &#x60;projects/myProject/topics/deviceEvents&#x60;. | [optional] 

## Methods

### NewOmnicoreStateNotificationConfig

`func NewOmnicoreStateNotificationConfig() *OmnicoreStateNotificationConfig`

NewOmnicoreStateNotificationConfig instantiates a new OmnicoreStateNotificationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOmnicoreStateNotificationConfigWithDefaults

`func NewOmnicoreStateNotificationConfigWithDefaults() *OmnicoreStateNotificationConfig`

NewOmnicoreStateNotificationConfigWithDefaults instantiates a new OmnicoreStateNotificationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubsubTopicName

`func (o *OmnicoreStateNotificationConfig) GetPubsubTopicName() string`

GetPubsubTopicName returns the PubsubTopicName field if non-nil, zero value otherwise.

### GetPubsubTopicNameOk

`func (o *OmnicoreStateNotificationConfig) GetPubsubTopicNameOk() (*string, bool)`

GetPubsubTopicNameOk returns a tuple with the PubsubTopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubsubTopicName

`func (o *OmnicoreStateNotificationConfig) SetPubsubTopicName(v string)`

SetPubsubTopicName sets PubsubTopicName field to given value.

### HasPubsubTopicName

`func (o *OmnicoreStateNotificationConfig) HasPubsubTopicName() bool`

HasPubsubTopicName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


