# GetExportsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]ExportDetail**](ExportDetail.md) |  | [optional] 

## Methods

### NewGetExportsResponse

`func NewGetExportsResponse() *GetExportsResponse`

NewGetExportsResponse instantiates a new GetExportsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetExportsResponseWithDefaults

`func NewGetExportsResponseWithDefaults() *GetExportsResponse`

NewGetExportsResponseWithDefaults instantiates a new GetExportsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *GetExportsResponse) GetDetails() []ExportDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetExportsResponse) GetDetailsOk() (*[]ExportDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetExportsResponse) SetDetails(v []ExportDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GetExportsResponse) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


