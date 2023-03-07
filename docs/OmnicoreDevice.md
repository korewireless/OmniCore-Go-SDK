# OmnicoreDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocked** | Pointer to **bool** |  | [optional] 
**Capresent** | Pointer to **bool** |  | [optional] 
**ClientOnline** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to [**OmnicoreDeviceConfig**](OmnicoreDeviceConfig.md) |  | [optional] 
**CreatedOn** | Pointer to **string** |  | [optional] 
**Credentials** | Pointer to [**[]OmnicoreDeviceCredential**](OmnicoreDeviceCredential.md) |  | [optional] 
**DeviceErrors** | Pointer to **string** |  | [optional] 
**Gateway** | Pointer to **[]string** |  | [optional] 
**GatewayConfig** | Pointer to [**OmnicoreGatewayConfig**](OmnicoreGatewayConfig.md) |  | [optional] 
**Id** | **string** |  | 
**IsGateway** | Pointer to **bool** |  | [optional] 
**LastConfigAckTime** | Pointer to **string** |  | [optional] 
**LastConfigSendTime** | Pointer to **string** |  | [optional] 
**LastErrorStatus** | Pointer to [**OmnicoreErrorStatus**](OmnicoreErrorStatus.md) |  | [optional] 
**LastErrorTime** | Pointer to **string** |  | [optional] 
**LastEventTime** | Pointer to **string** |  | [optional] 
**LastHeartbeatTime** | Pointer to **string** |  | [optional] 
**LastStateTime** | Pointer to **string** |  | [optional] 
**LogLevel** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NumId** | Pointer to **string** |  | [optional] 
**Parent** | **string** |  | 
**Registry** | **string** |  | 
**State** | Pointer to [**OmnicoreDeviceState**](OmnicoreDeviceState.md) |  | [optional] 
**Subscription** | **string** |  | 
**Subscriptions** | Pointer to **[]string** |  | [optional] 
**UpdatedOn** | Pointer to **string** |  | [optional] 

## Methods

### NewOmnicoreDevice

`func NewOmnicoreDevice(id string, parent string, registry string, subscription string, ) *OmnicoreDevice`

NewOmnicoreDevice instantiates a new OmnicoreDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOmnicoreDeviceWithDefaults

`func NewOmnicoreDeviceWithDefaults() *OmnicoreDevice`

NewOmnicoreDeviceWithDefaults instantiates a new OmnicoreDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocked

`func (o *OmnicoreDevice) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *OmnicoreDevice) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *OmnicoreDevice) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *OmnicoreDevice) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetCapresent

`func (o *OmnicoreDevice) GetCapresent() bool`

GetCapresent returns the Capresent field if non-nil, zero value otherwise.

### GetCapresentOk

`func (o *OmnicoreDevice) GetCapresentOk() (*bool, bool)`

GetCapresentOk returns a tuple with the Capresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapresent

`func (o *OmnicoreDevice) SetCapresent(v bool)`

SetCapresent sets Capresent field to given value.

### HasCapresent

`func (o *OmnicoreDevice) HasCapresent() bool`

HasCapresent returns a boolean if a field has been set.

### GetClientOnline

`func (o *OmnicoreDevice) GetClientOnline() bool`

GetClientOnline returns the ClientOnline field if non-nil, zero value otherwise.

### GetClientOnlineOk

`func (o *OmnicoreDevice) GetClientOnlineOk() (*bool, bool)`

GetClientOnlineOk returns a tuple with the ClientOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOnline

`func (o *OmnicoreDevice) SetClientOnline(v bool)`

SetClientOnline sets ClientOnline field to given value.

### HasClientOnline

`func (o *OmnicoreDevice) HasClientOnline() bool`

HasClientOnline returns a boolean if a field has been set.

### GetConfig

`func (o *OmnicoreDevice) GetConfig() OmnicoreDeviceConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *OmnicoreDevice) GetConfigOk() (*OmnicoreDeviceConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *OmnicoreDevice) SetConfig(v OmnicoreDeviceConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *OmnicoreDevice) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedOn

`func (o *OmnicoreDevice) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *OmnicoreDevice) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *OmnicoreDevice) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *OmnicoreDevice) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetCredentials

`func (o *OmnicoreDevice) GetCredentials() []OmnicoreDeviceCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *OmnicoreDevice) GetCredentialsOk() (*[]OmnicoreDeviceCredential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *OmnicoreDevice) SetCredentials(v []OmnicoreDeviceCredential)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *OmnicoreDevice) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetDeviceErrors

`func (o *OmnicoreDevice) GetDeviceErrors() string`

GetDeviceErrors returns the DeviceErrors field if non-nil, zero value otherwise.

### GetDeviceErrorsOk

`func (o *OmnicoreDevice) GetDeviceErrorsOk() (*string, bool)`

GetDeviceErrorsOk returns a tuple with the DeviceErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceErrors

`func (o *OmnicoreDevice) SetDeviceErrors(v string)`

SetDeviceErrors sets DeviceErrors field to given value.

### HasDeviceErrors

`func (o *OmnicoreDevice) HasDeviceErrors() bool`

HasDeviceErrors returns a boolean if a field has been set.

### GetGateway

`func (o *OmnicoreDevice) GetGateway() []string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *OmnicoreDevice) GetGatewayOk() (*[]string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *OmnicoreDevice) SetGateway(v []string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *OmnicoreDevice) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetGatewayConfig

`func (o *OmnicoreDevice) GetGatewayConfig() OmnicoreGatewayConfig`

GetGatewayConfig returns the GatewayConfig field if non-nil, zero value otherwise.

### GetGatewayConfigOk

`func (o *OmnicoreDevice) GetGatewayConfigOk() (*OmnicoreGatewayConfig, bool)`

GetGatewayConfigOk returns a tuple with the GatewayConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayConfig

`func (o *OmnicoreDevice) SetGatewayConfig(v OmnicoreGatewayConfig)`

SetGatewayConfig sets GatewayConfig field to given value.

### HasGatewayConfig

`func (o *OmnicoreDevice) HasGatewayConfig() bool`

HasGatewayConfig returns a boolean if a field has been set.

### GetId

`func (o *OmnicoreDevice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OmnicoreDevice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OmnicoreDevice) SetId(v string)`

SetId sets Id field to given value.


### GetIsGateway

`func (o *OmnicoreDevice) GetIsGateway() bool`

GetIsGateway returns the IsGateway field if non-nil, zero value otherwise.

### GetIsGatewayOk

`func (o *OmnicoreDevice) GetIsGatewayOk() (*bool, bool)`

GetIsGatewayOk returns a tuple with the IsGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGateway

`func (o *OmnicoreDevice) SetIsGateway(v bool)`

SetIsGateway sets IsGateway field to given value.

### HasIsGateway

`func (o *OmnicoreDevice) HasIsGateway() bool`

HasIsGateway returns a boolean if a field has been set.

### GetLastConfigAckTime

`func (o *OmnicoreDevice) GetLastConfigAckTime() string`

GetLastConfigAckTime returns the LastConfigAckTime field if non-nil, zero value otherwise.

### GetLastConfigAckTimeOk

`func (o *OmnicoreDevice) GetLastConfigAckTimeOk() (*string, bool)`

GetLastConfigAckTimeOk returns a tuple with the LastConfigAckTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConfigAckTime

`func (o *OmnicoreDevice) SetLastConfigAckTime(v string)`

SetLastConfigAckTime sets LastConfigAckTime field to given value.

### HasLastConfigAckTime

`func (o *OmnicoreDevice) HasLastConfigAckTime() bool`

HasLastConfigAckTime returns a boolean if a field has been set.

### GetLastConfigSendTime

`func (o *OmnicoreDevice) GetLastConfigSendTime() string`

GetLastConfigSendTime returns the LastConfigSendTime field if non-nil, zero value otherwise.

### GetLastConfigSendTimeOk

`func (o *OmnicoreDevice) GetLastConfigSendTimeOk() (*string, bool)`

GetLastConfigSendTimeOk returns a tuple with the LastConfigSendTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConfigSendTime

`func (o *OmnicoreDevice) SetLastConfigSendTime(v string)`

SetLastConfigSendTime sets LastConfigSendTime field to given value.

### HasLastConfigSendTime

`func (o *OmnicoreDevice) HasLastConfigSendTime() bool`

HasLastConfigSendTime returns a boolean if a field has been set.

### GetLastErrorStatus

`func (o *OmnicoreDevice) GetLastErrorStatus() OmnicoreErrorStatus`

GetLastErrorStatus returns the LastErrorStatus field if non-nil, zero value otherwise.

### GetLastErrorStatusOk

`func (o *OmnicoreDevice) GetLastErrorStatusOk() (*OmnicoreErrorStatus, bool)`

GetLastErrorStatusOk returns a tuple with the LastErrorStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorStatus

`func (o *OmnicoreDevice) SetLastErrorStatus(v OmnicoreErrorStatus)`

SetLastErrorStatus sets LastErrorStatus field to given value.

### HasLastErrorStatus

`func (o *OmnicoreDevice) HasLastErrorStatus() bool`

HasLastErrorStatus returns a boolean if a field has been set.

### GetLastErrorTime

`func (o *OmnicoreDevice) GetLastErrorTime() string`

GetLastErrorTime returns the LastErrorTime field if non-nil, zero value otherwise.

### GetLastErrorTimeOk

`func (o *OmnicoreDevice) GetLastErrorTimeOk() (*string, bool)`

GetLastErrorTimeOk returns a tuple with the LastErrorTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorTime

`func (o *OmnicoreDevice) SetLastErrorTime(v string)`

SetLastErrorTime sets LastErrorTime field to given value.

### HasLastErrorTime

`func (o *OmnicoreDevice) HasLastErrorTime() bool`

HasLastErrorTime returns a boolean if a field has been set.

### GetLastEventTime

`func (o *OmnicoreDevice) GetLastEventTime() string`

GetLastEventTime returns the LastEventTime field if non-nil, zero value otherwise.

### GetLastEventTimeOk

`func (o *OmnicoreDevice) GetLastEventTimeOk() (*string, bool)`

GetLastEventTimeOk returns a tuple with the LastEventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEventTime

`func (o *OmnicoreDevice) SetLastEventTime(v string)`

SetLastEventTime sets LastEventTime field to given value.

### HasLastEventTime

`func (o *OmnicoreDevice) HasLastEventTime() bool`

HasLastEventTime returns a boolean if a field has been set.

### GetLastHeartbeatTime

`func (o *OmnicoreDevice) GetLastHeartbeatTime() string`

GetLastHeartbeatTime returns the LastHeartbeatTime field if non-nil, zero value otherwise.

### GetLastHeartbeatTimeOk

`func (o *OmnicoreDevice) GetLastHeartbeatTimeOk() (*string, bool)`

GetLastHeartbeatTimeOk returns a tuple with the LastHeartbeatTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeartbeatTime

`func (o *OmnicoreDevice) SetLastHeartbeatTime(v string)`

SetLastHeartbeatTime sets LastHeartbeatTime field to given value.

### HasLastHeartbeatTime

`func (o *OmnicoreDevice) HasLastHeartbeatTime() bool`

HasLastHeartbeatTime returns a boolean if a field has been set.

### GetLastStateTime

`func (o *OmnicoreDevice) GetLastStateTime() string`

GetLastStateTime returns the LastStateTime field if non-nil, zero value otherwise.

### GetLastStateTimeOk

`func (o *OmnicoreDevice) GetLastStateTimeOk() (*string, bool)`

GetLastStateTimeOk returns a tuple with the LastStateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStateTime

`func (o *OmnicoreDevice) SetLastStateTime(v string)`

SetLastStateTime sets LastStateTime field to given value.

### HasLastStateTime

`func (o *OmnicoreDevice) HasLastStateTime() bool`

HasLastStateTime returns a boolean if a field has been set.

### GetLogLevel

`func (o *OmnicoreDevice) GetLogLevel() string`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *OmnicoreDevice) GetLogLevelOk() (*string, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *OmnicoreDevice) SetLogLevel(v string)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *OmnicoreDevice) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetMetadata

`func (o *OmnicoreDevice) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OmnicoreDevice) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OmnicoreDevice) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OmnicoreDevice) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *OmnicoreDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OmnicoreDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OmnicoreDevice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OmnicoreDevice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumId

`func (o *OmnicoreDevice) GetNumId() string`

GetNumId returns the NumId field if non-nil, zero value otherwise.

### GetNumIdOk

`func (o *OmnicoreDevice) GetNumIdOk() (*string, bool)`

GetNumIdOk returns a tuple with the NumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumId

`func (o *OmnicoreDevice) SetNumId(v string)`

SetNumId sets NumId field to given value.

### HasNumId

`func (o *OmnicoreDevice) HasNumId() bool`

HasNumId returns a boolean if a field has been set.

### GetParent

`func (o *OmnicoreDevice) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *OmnicoreDevice) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *OmnicoreDevice) SetParent(v string)`

SetParent sets Parent field to given value.


### GetRegistry

`func (o *OmnicoreDevice) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *OmnicoreDevice) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *OmnicoreDevice) SetRegistry(v string)`

SetRegistry sets Registry field to given value.


### GetState

`func (o *OmnicoreDevice) GetState() OmnicoreDeviceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OmnicoreDevice) GetStateOk() (*OmnicoreDeviceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OmnicoreDevice) SetState(v OmnicoreDeviceState)`

SetState sets State field to given value.

### HasState

`func (o *OmnicoreDevice) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubscription

`func (o *OmnicoreDevice) GetSubscription() string`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *OmnicoreDevice) GetSubscriptionOk() (*string, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *OmnicoreDevice) SetSubscription(v string)`

SetSubscription sets Subscription field to given value.


### GetSubscriptions

`func (o *OmnicoreDevice) GetSubscriptions() []string`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *OmnicoreDevice) GetSubscriptionsOk() (*[]string, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *OmnicoreDevice) SetSubscriptions(v []string)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *OmnicoreDevice) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *OmnicoreDevice) GetUpdatedOn() string`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *OmnicoreDevice) GetUpdatedOnOk() (*string, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *OmnicoreDevice) SetUpdatedOn(v string)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *OmnicoreDevice) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


