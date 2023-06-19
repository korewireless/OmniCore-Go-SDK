# MetricsDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NoOfMessagesFor30Minutes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**NoOfMessagesFor48Hours** | Pointer to **[]map[string]interface{}** |  | [optional] 
**BillableBytesReceived** | Pointer to **int32** |  | [optional] 
**BillableBytesSent** | Pointer to **int32** |  | [optional] 
**BillableMessageSize** | Pointer to **int32** |  | [optional] 
**BytesReceived** | Pointer to **int32** |  | [optional] 
**BytesSent** | Pointer to **int32** |  | [optional] 
**MessageSize** | Pointer to **int32** |  | [optional] 
**NoOfAckMessages** | Pointer to **int32** |  | [optional] 
**NoOfCommandMessages** | Pointer to **int32** |  | [optional] 
**NoOfConfigMessages** | Pointer to **int32** |  | [optional] 
**NoOfDeviceConnectionsFailed** | Pointer to **int32** |  | [optional] 
**NoOfDevices** | Pointer to **int32** |  | [optional] 
**NoOfDisConnections** | Pointer to **int32** |  | [optional] 
**NoOfEventMessages** | Pointer to **int32** |  | [optional] 
**NoOfGatewayConnectionsFailed** | Pointer to **int32** |  | [optional] 
**NoOfGateways** | Pointer to **int32** |  | [optional] 
**NoOfLoopBackMessages** | Pointer to **int32** |  | [optional] 
**NoOfMessages** | Pointer to **int32** |  | [optional] 
**NoOfPublishErrors** | Pointer to **int32** |  | [optional] 
**NoOfRegistries** | Pointer to **int32** |  | [optional] 
**NoOfStateMessages** | Pointer to **int32** |  | [optional] 
**NoOfSubscribe** | Pointer to **int32** |  | [optional] 
**NoOfSuccessfulConnections** | Pointer to **int32** |  | [optional] 
**NoOfUnSubscribe** | Pointer to **int32** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 

## Methods

### NewMetricsDetails

`func NewMetricsDetails() *MetricsDetails`

NewMetricsDetails instantiates a new MetricsDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsDetailsWithDefaults

`func NewMetricsDetailsWithDefaults() *MetricsDetails`

NewMetricsDetailsWithDefaults instantiates a new MetricsDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNoOfMessagesFor30Minutes

`func (o *MetricsDetails) GetNoOfMessagesFor30Minutes() []map[string]interface{}`

GetNoOfMessagesFor30Minutes returns the NoOfMessagesFor30Minutes field if non-nil, zero value otherwise.

### GetNoOfMessagesFor30MinutesOk

`func (o *MetricsDetails) GetNoOfMessagesFor30MinutesOk() (*[]map[string]interface{}, bool)`

GetNoOfMessagesFor30MinutesOk returns a tuple with the NoOfMessagesFor30Minutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfMessagesFor30Minutes

`func (o *MetricsDetails) SetNoOfMessagesFor30Minutes(v []map[string]interface{})`

SetNoOfMessagesFor30Minutes sets NoOfMessagesFor30Minutes field to given value.

### HasNoOfMessagesFor30Minutes

`func (o *MetricsDetails) HasNoOfMessagesFor30Minutes() bool`

HasNoOfMessagesFor30Minutes returns a boolean if a field has been set.

### GetNoOfMessagesFor48Hours

`func (o *MetricsDetails) GetNoOfMessagesFor48Hours() []map[string]interface{}`

GetNoOfMessagesFor48Hours returns the NoOfMessagesFor48Hours field if non-nil, zero value otherwise.

### GetNoOfMessagesFor48HoursOk

`func (o *MetricsDetails) GetNoOfMessagesFor48HoursOk() (*[]map[string]interface{}, bool)`

GetNoOfMessagesFor48HoursOk returns a tuple with the NoOfMessagesFor48Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfMessagesFor48Hours

`func (o *MetricsDetails) SetNoOfMessagesFor48Hours(v []map[string]interface{})`

SetNoOfMessagesFor48Hours sets NoOfMessagesFor48Hours field to given value.

### HasNoOfMessagesFor48Hours

`func (o *MetricsDetails) HasNoOfMessagesFor48Hours() bool`

HasNoOfMessagesFor48Hours returns a boolean if a field has been set.

### GetBillableBytesReceived

`func (o *MetricsDetails) GetBillableBytesReceived() int32`

GetBillableBytesReceived returns the BillableBytesReceived field if non-nil, zero value otherwise.

### GetBillableBytesReceivedOk

`func (o *MetricsDetails) GetBillableBytesReceivedOk() (*int32, bool)`

GetBillableBytesReceivedOk returns a tuple with the BillableBytesReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableBytesReceived

`func (o *MetricsDetails) SetBillableBytesReceived(v int32)`

SetBillableBytesReceived sets BillableBytesReceived field to given value.

### HasBillableBytesReceived

`func (o *MetricsDetails) HasBillableBytesReceived() bool`

HasBillableBytesReceived returns a boolean if a field has been set.

### GetBillableBytesSent

`func (o *MetricsDetails) GetBillableBytesSent() int32`

GetBillableBytesSent returns the BillableBytesSent field if non-nil, zero value otherwise.

### GetBillableBytesSentOk

`func (o *MetricsDetails) GetBillableBytesSentOk() (*int32, bool)`

GetBillableBytesSentOk returns a tuple with the BillableBytesSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableBytesSent

`func (o *MetricsDetails) SetBillableBytesSent(v int32)`

SetBillableBytesSent sets BillableBytesSent field to given value.

### HasBillableBytesSent

`func (o *MetricsDetails) HasBillableBytesSent() bool`

HasBillableBytesSent returns a boolean if a field has been set.

### GetBillableMessageSize

`func (o *MetricsDetails) GetBillableMessageSize() int32`

GetBillableMessageSize returns the BillableMessageSize field if non-nil, zero value otherwise.

### GetBillableMessageSizeOk

`func (o *MetricsDetails) GetBillableMessageSizeOk() (*int32, bool)`

GetBillableMessageSizeOk returns a tuple with the BillableMessageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableMessageSize

`func (o *MetricsDetails) SetBillableMessageSize(v int32)`

SetBillableMessageSize sets BillableMessageSize field to given value.

### HasBillableMessageSize

`func (o *MetricsDetails) HasBillableMessageSize() bool`

HasBillableMessageSize returns a boolean if a field has been set.

### GetBytesReceived

`func (o *MetricsDetails) GetBytesReceived() int32`

GetBytesReceived returns the BytesReceived field if non-nil, zero value otherwise.

### GetBytesReceivedOk

`func (o *MetricsDetails) GetBytesReceivedOk() (*int32, bool)`

GetBytesReceivedOk returns a tuple with the BytesReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesReceived

`func (o *MetricsDetails) SetBytesReceived(v int32)`

SetBytesReceived sets BytesReceived field to given value.

### HasBytesReceived

`func (o *MetricsDetails) HasBytesReceived() bool`

HasBytesReceived returns a boolean if a field has been set.

### GetBytesSent

`func (o *MetricsDetails) GetBytesSent() int32`

GetBytesSent returns the BytesSent field if non-nil, zero value otherwise.

### GetBytesSentOk

`func (o *MetricsDetails) GetBytesSentOk() (*int32, bool)`

GetBytesSentOk returns a tuple with the BytesSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesSent

`func (o *MetricsDetails) SetBytesSent(v int32)`

SetBytesSent sets BytesSent field to given value.

### HasBytesSent

`func (o *MetricsDetails) HasBytesSent() bool`

HasBytesSent returns a boolean if a field has been set.

### GetMessageSize

`func (o *MetricsDetails) GetMessageSize() int32`

GetMessageSize returns the MessageSize field if non-nil, zero value otherwise.

### GetMessageSizeOk

`func (o *MetricsDetails) GetMessageSizeOk() (*int32, bool)`

GetMessageSizeOk returns a tuple with the MessageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSize

`func (o *MetricsDetails) SetMessageSize(v int32)`

SetMessageSize sets MessageSize field to given value.

### HasMessageSize

`func (o *MetricsDetails) HasMessageSize() bool`

HasMessageSize returns a boolean if a field has been set.

### GetNoOfAckMessages

`func (o *MetricsDetails) GetNoOfAckMessages() int32`

GetNoOfAckMessages returns the NoOfAckMessages field if non-nil, zero value otherwise.

### GetNoOfAckMessagesOk

`func (o *MetricsDetails) GetNoOfAckMessagesOk() (*int32, bool)`

GetNoOfAckMessagesOk returns a tuple with the NoOfAckMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfAckMessages

`func (o *MetricsDetails) SetNoOfAckMessages(v int32)`

SetNoOfAckMessages sets NoOfAckMessages field to given value.

### HasNoOfAckMessages

`func (o *MetricsDetails) HasNoOfAckMessages() bool`

HasNoOfAckMessages returns a boolean if a field has been set.

### GetNoOfCommandMessages

`func (o *MetricsDetails) GetNoOfCommandMessages() int32`

GetNoOfCommandMessages returns the NoOfCommandMessages field if non-nil, zero value otherwise.

### GetNoOfCommandMessagesOk

`func (o *MetricsDetails) GetNoOfCommandMessagesOk() (*int32, bool)`

GetNoOfCommandMessagesOk returns a tuple with the NoOfCommandMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfCommandMessages

`func (o *MetricsDetails) SetNoOfCommandMessages(v int32)`

SetNoOfCommandMessages sets NoOfCommandMessages field to given value.

### HasNoOfCommandMessages

`func (o *MetricsDetails) HasNoOfCommandMessages() bool`

HasNoOfCommandMessages returns a boolean if a field has been set.

### GetNoOfConfigMessages

`func (o *MetricsDetails) GetNoOfConfigMessages() int32`

GetNoOfConfigMessages returns the NoOfConfigMessages field if non-nil, zero value otherwise.

### GetNoOfConfigMessagesOk

`func (o *MetricsDetails) GetNoOfConfigMessagesOk() (*int32, bool)`

GetNoOfConfigMessagesOk returns a tuple with the NoOfConfigMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfConfigMessages

`func (o *MetricsDetails) SetNoOfConfigMessages(v int32)`

SetNoOfConfigMessages sets NoOfConfigMessages field to given value.

### HasNoOfConfigMessages

`func (o *MetricsDetails) HasNoOfConfigMessages() bool`

HasNoOfConfigMessages returns a boolean if a field has been set.

### GetNoOfDeviceConnectionsFailed

`func (o *MetricsDetails) GetNoOfDeviceConnectionsFailed() int32`

GetNoOfDeviceConnectionsFailed returns the NoOfDeviceConnectionsFailed field if non-nil, zero value otherwise.

### GetNoOfDeviceConnectionsFailedOk

`func (o *MetricsDetails) GetNoOfDeviceConnectionsFailedOk() (*int32, bool)`

GetNoOfDeviceConnectionsFailedOk returns a tuple with the NoOfDeviceConnectionsFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfDeviceConnectionsFailed

`func (o *MetricsDetails) SetNoOfDeviceConnectionsFailed(v int32)`

SetNoOfDeviceConnectionsFailed sets NoOfDeviceConnectionsFailed field to given value.

### HasNoOfDeviceConnectionsFailed

`func (o *MetricsDetails) HasNoOfDeviceConnectionsFailed() bool`

HasNoOfDeviceConnectionsFailed returns a boolean if a field has been set.

### GetNoOfDevices

`func (o *MetricsDetails) GetNoOfDevices() int32`

GetNoOfDevices returns the NoOfDevices field if non-nil, zero value otherwise.

### GetNoOfDevicesOk

`func (o *MetricsDetails) GetNoOfDevicesOk() (*int32, bool)`

GetNoOfDevicesOk returns a tuple with the NoOfDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfDevices

`func (o *MetricsDetails) SetNoOfDevices(v int32)`

SetNoOfDevices sets NoOfDevices field to given value.

### HasNoOfDevices

`func (o *MetricsDetails) HasNoOfDevices() bool`

HasNoOfDevices returns a boolean if a field has been set.

### GetNoOfDisConnections

`func (o *MetricsDetails) GetNoOfDisConnections() int32`

GetNoOfDisConnections returns the NoOfDisConnections field if non-nil, zero value otherwise.

### GetNoOfDisConnectionsOk

`func (o *MetricsDetails) GetNoOfDisConnectionsOk() (*int32, bool)`

GetNoOfDisConnectionsOk returns a tuple with the NoOfDisConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfDisConnections

`func (o *MetricsDetails) SetNoOfDisConnections(v int32)`

SetNoOfDisConnections sets NoOfDisConnections field to given value.

### HasNoOfDisConnections

`func (o *MetricsDetails) HasNoOfDisConnections() bool`

HasNoOfDisConnections returns a boolean if a field has been set.

### GetNoOfEventMessages

`func (o *MetricsDetails) GetNoOfEventMessages() int32`

GetNoOfEventMessages returns the NoOfEventMessages field if non-nil, zero value otherwise.

### GetNoOfEventMessagesOk

`func (o *MetricsDetails) GetNoOfEventMessagesOk() (*int32, bool)`

GetNoOfEventMessagesOk returns a tuple with the NoOfEventMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfEventMessages

`func (o *MetricsDetails) SetNoOfEventMessages(v int32)`

SetNoOfEventMessages sets NoOfEventMessages field to given value.

### HasNoOfEventMessages

`func (o *MetricsDetails) HasNoOfEventMessages() bool`

HasNoOfEventMessages returns a boolean if a field has been set.

### GetNoOfGatewayConnectionsFailed

`func (o *MetricsDetails) GetNoOfGatewayConnectionsFailed() int32`

GetNoOfGatewayConnectionsFailed returns the NoOfGatewayConnectionsFailed field if non-nil, zero value otherwise.

### GetNoOfGatewayConnectionsFailedOk

`func (o *MetricsDetails) GetNoOfGatewayConnectionsFailedOk() (*int32, bool)`

GetNoOfGatewayConnectionsFailedOk returns a tuple with the NoOfGatewayConnectionsFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfGatewayConnectionsFailed

`func (o *MetricsDetails) SetNoOfGatewayConnectionsFailed(v int32)`

SetNoOfGatewayConnectionsFailed sets NoOfGatewayConnectionsFailed field to given value.

### HasNoOfGatewayConnectionsFailed

`func (o *MetricsDetails) HasNoOfGatewayConnectionsFailed() bool`

HasNoOfGatewayConnectionsFailed returns a boolean if a field has been set.

### GetNoOfGateways

`func (o *MetricsDetails) GetNoOfGateways() int32`

GetNoOfGateways returns the NoOfGateways field if non-nil, zero value otherwise.

### GetNoOfGatewaysOk

`func (o *MetricsDetails) GetNoOfGatewaysOk() (*int32, bool)`

GetNoOfGatewaysOk returns a tuple with the NoOfGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfGateways

`func (o *MetricsDetails) SetNoOfGateways(v int32)`

SetNoOfGateways sets NoOfGateways field to given value.

### HasNoOfGateways

`func (o *MetricsDetails) HasNoOfGateways() bool`

HasNoOfGateways returns a boolean if a field has been set.

### GetNoOfLoopBackMessages

`func (o *MetricsDetails) GetNoOfLoopBackMessages() int32`

GetNoOfLoopBackMessages returns the NoOfLoopBackMessages field if non-nil, zero value otherwise.

### GetNoOfLoopBackMessagesOk

`func (o *MetricsDetails) GetNoOfLoopBackMessagesOk() (*int32, bool)`

GetNoOfLoopBackMessagesOk returns a tuple with the NoOfLoopBackMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfLoopBackMessages

`func (o *MetricsDetails) SetNoOfLoopBackMessages(v int32)`

SetNoOfLoopBackMessages sets NoOfLoopBackMessages field to given value.

### HasNoOfLoopBackMessages

`func (o *MetricsDetails) HasNoOfLoopBackMessages() bool`

HasNoOfLoopBackMessages returns a boolean if a field has been set.

### GetNoOfMessages

`func (o *MetricsDetails) GetNoOfMessages() int32`

GetNoOfMessages returns the NoOfMessages field if non-nil, zero value otherwise.

### GetNoOfMessagesOk

`func (o *MetricsDetails) GetNoOfMessagesOk() (*int32, bool)`

GetNoOfMessagesOk returns a tuple with the NoOfMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfMessages

`func (o *MetricsDetails) SetNoOfMessages(v int32)`

SetNoOfMessages sets NoOfMessages field to given value.

### HasNoOfMessages

`func (o *MetricsDetails) HasNoOfMessages() bool`

HasNoOfMessages returns a boolean if a field has been set.

### GetNoOfPublishErrors

`func (o *MetricsDetails) GetNoOfPublishErrors() int32`

GetNoOfPublishErrors returns the NoOfPublishErrors field if non-nil, zero value otherwise.

### GetNoOfPublishErrorsOk

`func (o *MetricsDetails) GetNoOfPublishErrorsOk() (*int32, bool)`

GetNoOfPublishErrorsOk returns a tuple with the NoOfPublishErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfPublishErrors

`func (o *MetricsDetails) SetNoOfPublishErrors(v int32)`

SetNoOfPublishErrors sets NoOfPublishErrors field to given value.

### HasNoOfPublishErrors

`func (o *MetricsDetails) HasNoOfPublishErrors() bool`

HasNoOfPublishErrors returns a boolean if a field has been set.

### GetNoOfRegistries

`func (o *MetricsDetails) GetNoOfRegistries() int32`

GetNoOfRegistries returns the NoOfRegistries field if non-nil, zero value otherwise.

### GetNoOfRegistriesOk

`func (o *MetricsDetails) GetNoOfRegistriesOk() (*int32, bool)`

GetNoOfRegistriesOk returns a tuple with the NoOfRegistries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfRegistries

`func (o *MetricsDetails) SetNoOfRegistries(v int32)`

SetNoOfRegistries sets NoOfRegistries field to given value.

### HasNoOfRegistries

`func (o *MetricsDetails) HasNoOfRegistries() bool`

HasNoOfRegistries returns a boolean if a field has been set.

### GetNoOfStateMessages

`func (o *MetricsDetails) GetNoOfStateMessages() int32`

GetNoOfStateMessages returns the NoOfStateMessages field if non-nil, zero value otherwise.

### GetNoOfStateMessagesOk

`func (o *MetricsDetails) GetNoOfStateMessagesOk() (*int32, bool)`

GetNoOfStateMessagesOk returns a tuple with the NoOfStateMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfStateMessages

`func (o *MetricsDetails) SetNoOfStateMessages(v int32)`

SetNoOfStateMessages sets NoOfStateMessages field to given value.

### HasNoOfStateMessages

`func (o *MetricsDetails) HasNoOfStateMessages() bool`

HasNoOfStateMessages returns a boolean if a field has been set.

### GetNoOfSubscribe

`func (o *MetricsDetails) GetNoOfSubscribe() int32`

GetNoOfSubscribe returns the NoOfSubscribe field if non-nil, zero value otherwise.

### GetNoOfSubscribeOk

`func (o *MetricsDetails) GetNoOfSubscribeOk() (*int32, bool)`

GetNoOfSubscribeOk returns a tuple with the NoOfSubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfSubscribe

`func (o *MetricsDetails) SetNoOfSubscribe(v int32)`

SetNoOfSubscribe sets NoOfSubscribe field to given value.

### HasNoOfSubscribe

`func (o *MetricsDetails) HasNoOfSubscribe() bool`

HasNoOfSubscribe returns a boolean if a field has been set.

### GetNoOfSuccessfulConnections

`func (o *MetricsDetails) GetNoOfSuccessfulConnections() int32`

GetNoOfSuccessfulConnections returns the NoOfSuccessfulConnections field if non-nil, zero value otherwise.

### GetNoOfSuccessfulConnectionsOk

`func (o *MetricsDetails) GetNoOfSuccessfulConnectionsOk() (*int32, bool)`

GetNoOfSuccessfulConnectionsOk returns a tuple with the NoOfSuccessfulConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfSuccessfulConnections

`func (o *MetricsDetails) SetNoOfSuccessfulConnections(v int32)`

SetNoOfSuccessfulConnections sets NoOfSuccessfulConnections field to given value.

### HasNoOfSuccessfulConnections

`func (o *MetricsDetails) HasNoOfSuccessfulConnections() bool`

HasNoOfSuccessfulConnections returns a boolean if a field has been set.

### GetNoOfUnSubscribe

`func (o *MetricsDetails) GetNoOfUnSubscribe() int32`

GetNoOfUnSubscribe returns the NoOfUnSubscribe field if non-nil, zero value otherwise.

### GetNoOfUnSubscribeOk

`func (o *MetricsDetails) GetNoOfUnSubscribeOk() (*int32, bool)`

GetNoOfUnSubscribeOk returns a tuple with the NoOfUnSubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfUnSubscribe

`func (o *MetricsDetails) SetNoOfUnSubscribe(v int32)`

SetNoOfUnSubscribe sets NoOfUnSubscribe field to given value.

### HasNoOfUnSubscribe

`func (o *MetricsDetails) HasNoOfUnSubscribe() bool`

HasNoOfUnSubscribe returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *MetricsDetails) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *MetricsDetails) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *MetricsDetails) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *MetricsDetails) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


