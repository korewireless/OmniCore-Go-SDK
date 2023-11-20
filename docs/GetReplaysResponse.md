# GetReplaysResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]Replay**](Replay.md) |  | [optional] 

## Methods

### NewGetReplaysResponse

`func NewGetReplaysResponse() *GetReplaysResponse`

NewGetReplaysResponse instantiates a new GetReplaysResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReplaysResponseWithDefaults

`func NewGetReplaysResponseWithDefaults() *GetReplaysResponse`

NewGetReplaysResponseWithDefaults instantiates a new GetReplaysResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *GetReplaysResponse) GetDetails() []Replay`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetReplaysResponse) GetDetailsOk() (*[]Replay, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetReplaysResponse) SetDetails(v []Replay)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GetReplaysResponse) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


