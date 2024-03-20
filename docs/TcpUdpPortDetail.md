# TcpUdpPortDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Indicates if the port is enabled | 
**Port** | Pointer to **int32** | The port number | [optional] 

## Methods

### NewTcpUdpPortDetail

`func NewTcpUdpPortDetail(enabled bool, ) *TcpUdpPortDetail`

NewTcpUdpPortDetail instantiates a new TcpUdpPortDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTcpUdpPortDetailWithDefaults

`func NewTcpUdpPortDetailWithDefaults() *TcpUdpPortDetail`

NewTcpUdpPortDetailWithDefaults instantiates a new TcpUdpPortDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *TcpUdpPortDetail) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TcpUdpPortDetail) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TcpUdpPortDetail) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetPort

`func (o *TcpUdpPortDetail) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TcpUdpPortDetail) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TcpUdpPortDetail) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *TcpUdpPortDetail) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


