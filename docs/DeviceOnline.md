# DeviceOnline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Registry** | Pointer to **string** |  | [optional] [readonly] 
**ClientOnline** | Pointer to **bool** |  | [optional] [readonly] 
**LastHeartbeatTime** | Pointer to **string** |  | [optional] [readonly] 
**Subscription** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewDeviceOnline

`func NewDeviceOnline(id string, ) *DeviceOnline`

NewDeviceOnline instantiates a new DeviceOnline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceOnlineWithDefaults

`func NewDeviceOnlineWithDefaults() *DeviceOnline`

NewDeviceOnlineWithDefaults instantiates a new DeviceOnline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceOnline) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceOnline) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceOnline) SetId(v string)`

SetId sets Id field to given value.


### GetRegistry

`func (o *DeviceOnline) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *DeviceOnline) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *DeviceOnline) SetRegistry(v string)`

SetRegistry sets Registry field to given value.

### HasRegistry

`func (o *DeviceOnline) HasRegistry() bool`

HasRegistry returns a boolean if a field has been set.

### GetClientOnline

`func (o *DeviceOnline) GetClientOnline() bool`

GetClientOnline returns the ClientOnline field if non-nil, zero value otherwise.

### GetClientOnlineOk

`func (o *DeviceOnline) GetClientOnlineOk() (*bool, bool)`

GetClientOnlineOk returns a tuple with the ClientOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOnline

`func (o *DeviceOnline) SetClientOnline(v bool)`

SetClientOnline sets ClientOnline field to given value.

### HasClientOnline

`func (o *DeviceOnline) HasClientOnline() bool`

HasClientOnline returns a boolean if a field has been set.

### GetLastHeartbeatTime

`func (o *DeviceOnline) GetLastHeartbeatTime() string`

GetLastHeartbeatTime returns the LastHeartbeatTime field if non-nil, zero value otherwise.

### GetLastHeartbeatTimeOk

`func (o *DeviceOnline) GetLastHeartbeatTimeOk() (*string, bool)`

GetLastHeartbeatTimeOk returns a tuple with the LastHeartbeatTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeartbeatTime

`func (o *DeviceOnline) SetLastHeartbeatTime(v string)`

SetLastHeartbeatTime sets LastHeartbeatTime field to given value.

### HasLastHeartbeatTime

`func (o *DeviceOnline) HasLastHeartbeatTime() bool`

HasLastHeartbeatTime returns a boolean if a field has been set.

### GetSubscription

`func (o *DeviceOnline) GetSubscription() string`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *DeviceOnline) GetSubscriptionOk() (*string, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *DeviceOnline) SetSubscription(v string)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *DeviceOnline) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


