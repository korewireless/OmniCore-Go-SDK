# Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionParameter** | **string** |  | 
**Region** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 

## Methods

### NewConfig

`func NewConfig(connectionParameter string, ) *Config`

NewConfig instantiates a new Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigWithDefaults

`func NewConfigWithDefaults() *Config`

NewConfigWithDefaults instantiates a new Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionParameter

`func (o *Config) GetConnectionParameter() string`

GetConnectionParameter returns the ConnectionParameter field if non-nil, zero value otherwise.

### GetConnectionParameterOk

`func (o *Config) GetConnectionParameterOk() (*string, bool)`

GetConnectionParameterOk returns a tuple with the ConnectionParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionParameter

`func (o *Config) SetConnectionParameter(v string)`

SetConnectionParameter sets ConnectionParameter field to given value.


### GetRegion

`func (o *Config) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Config) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Config) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Config) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetExternalId

`func (o *Config) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Config) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Config) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Config) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


