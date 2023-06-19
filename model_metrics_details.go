/* Copyright 2018-2020 KoreWireless
 *
 * This is part of the KoreWireless Omnicore SDK.
 * It is licensed under the BSD 3-Clause license; you may not use this file
 * except in compliance with the License.
 *
 * You may obtain a copy of the License at:
 *  https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */


package OmniCore

import (
	"encoding/json"
)

// checks if the MetricsDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsDetails{}

// MetricsDetails struct for MetricsDetails
type MetricsDetails struct {
	NoOfMessagesFor30Minutes []map[string]interface{} `json:"NoOfMessagesFor30Minutes,omitempty"`
	NoOfMessagesFor48Hours []map[string]interface{} `json:"NoOfMessagesFor48Hours,omitempty"`
	BillableBytesReceived *int32 `json:"billableBytesReceived,omitempty"`
	BillableBytesSent *int32 `json:"billableBytesSent,omitempty"`
	BillableMessageSize *int32 `json:"billableMessageSize,omitempty"`
	BytesReceived *int32 `json:"bytesReceived,omitempty"`
	BytesSent *int32 `json:"bytesSent,omitempty"`
	MessageSize *int32 `json:"messageSize,omitempty"`
	NoOfAckMessages *int32 `json:"noOfAckMessages,omitempty"`
	NoOfCommandMessages *int32 `json:"noOfCommandMessages,omitempty"`
	NoOfConfigMessages *int32 `json:"noOfConfigMessages,omitempty"`
	NoOfDeviceConnectionsFailed *int32 `json:"noOfDeviceConnectionsFailed,omitempty"`
	NoOfDevices *int32 `json:"noOfDevices,omitempty"`
	NoOfDisConnections *int32 `json:"noOfDisConnections,omitempty"`
	NoOfEventMessages *int32 `json:"noOfEventMessages,omitempty"`
	NoOfGatewayConnectionsFailed *int32 `json:"noOfGatewayConnectionsFailed,omitempty"`
	NoOfGateways *int32 `json:"noOfGateways,omitempty"`
	NoOfLoopBackMessages *int32 `json:"noOfLoopBackMessages,omitempty"`
	NoOfMessages *int32 `json:"noOfMessages,omitempty"`
	NoOfPublishErrors *int32 `json:"noOfPublishErrors,omitempty"`
	NoOfRegistries *int32 `json:"noOfRegistries,omitempty"`
	NoOfStateMessages *int32 `json:"noOfStateMessages,omitempty"`
	NoOfSubscribe *int32 `json:"noOfSubscribe,omitempty"`
	NoOfSuccessfulConnections *int32 `json:"noOfSuccessfulConnections,omitempty"`
	NoOfUnSubscribe *int32 `json:"noOfUnSubscribe,omitempty"`
	SubscriptionId *string `json:"subscriptionId,omitempty"`
}

// NewMetricsDetails instantiates a new MetricsDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsDetails() *MetricsDetails {
	this := MetricsDetails{}
	return &this
}

// NewMetricsDetailsWithDefaults instantiates a new MetricsDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsDetailsWithDefaults() *MetricsDetails {
	this := MetricsDetails{}
	return &this
}

// GetNoOfMessagesFor30Minutes returns the NoOfMessagesFor30Minutes field value if set, zero value otherwise.
func (o *MetricsDetails) GetNoOfMessagesFor30Minutes() []map[string]interface{} {
	if o == nil || IsNil(o.NoOfMessagesFor30Minutes) {
		var ret []map[string]interface{}
		return ret
	}
	return o.NoOfMessagesFor30Minutes
}

// GetNoOfMessagesFor30MinutesOk returns a tuple with the NoOfMessagesFor30Minutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDetails) GetNoOfMessagesFor30MinutesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.NoOfMessagesFor30Minutes) {
		return nil, false
	}
	return o.NoOfMessagesFor30Minutes, true
}

// HasNoOfMessagesFor30Minutes returns a boolean if a field has been set.
func (o *MetricsDetails) HasNoOfMessagesFor30Minutes() bool {
	if o != nil && !IsNil(o.NoOfMessagesFor30Minutes) {
		return true
	}

	return false
}

// SetNoOfMessagesFor30Minutes gets a reference to the given []map[string]interface{} and assigns it to the NoOfMessagesFor30Minutes field.
func (o *MetricsDetails) SetNoOfMessagesFor30Minutes(v []map[string]interface{}) {
	o.NoOfMessagesFor30Minutes = v
}

// GetNoOfMessagesFor48Hours returns the NoOfMessagesFor48Hours field value if set, zero value otherwise.
func (o *MetricsDetails) GetNoOfMessagesFor48Hours() []map[string]interface{} {
	if o == nil || IsNil(o.NoOfMessagesFor48Hours) {
		var ret []map[string]interface{}
		return ret
	}
	return o.NoOfMessagesFor48Hours
}

// GetNoOfMessagesFor48HoursOk returns a tuple with the NoOfMessagesFor48Hours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDetails) GetNoOfMessagesFor48HoursOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.NoOfMessagesFor48Hours) {
		return nil, false
	}
	return o.NoOfMessagesFor48Hours, true
}

// HasNoOfMessagesFor48Hours returns a boolean if a field has been set.
func (o *MetricsDetails) HasNoOfMessagesFor48Hours() bool {
	if o != nil && !IsNil(o.NoOfMessagesFor48Hours) {
		return true
	}

	return false
}

// SetNoOfMessagesFor48Hours gets a reference to the given []map[string]interface{} and assigns it to the NoOfMessagesFor48Hours field.
func (o *MetricsDetails) SetNoOfMessagesFor48Hours(v []map[string]interface{}) {
	o.NoOfMessagesFor48Hours = v
}

// GetBillableBytesReceived returns the BillableBytesReceived field value if set, zero value otherwise.
func (o *MetricsDetails) GetBillableBytesReceived() int32 {
	if o == nil || IsNil(o.BillableBytesReceived) {
		var ret int32
		return ret
	}
	return *o.BillableBytesReceived
}

// GetBillableBytesReceivedOk returns a tuple with the BillableBytesReceived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDetails) GetBillableBytesReceivedOk() (*int32, bool) {
	if o == nil || IsNil(o.BillableBytesReceived) {
		return nil, false
	}
	return o.BillableBytesReceived, true
}

// HasBillableBytesReceived returns a boolean if a field has been set.
func (o *MetricsDetails) HasBillableBytesReceived() bool {
	if o != nil && !IsNil(o.BillableBytesReceived) {
		return true
	}

	return false
}

// SetBillableBytesReceived gets a reference to the given int32 and assigns it to the BillableBytesReceived field.
func (o *MetricsDetails) SetBillableBytesReceived(v int32) {
	o.BillableBytesReceived = &v
}

// GetBillableBytesSent returns the BillableBytesSent field value if set, zero value otherwise.
func (o *MetricsDetails) GetBillableBytesSent() int32 {
	if o == nil || IsNil(o.BillableBytesSent) {
		var ret int32
		return ret
	}
	return *o.BillableBytesSent
}

// GetBillableBytesSentOk returns a tuple with the BillableBytesSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDetails) GetBillableBytesSentOk() (*int32, bool) {
	if o == nil || IsNil(o.BillableBytesSent) {
		return nil, false
	}
	return o.BillableBytesSent, true
}

// HasBillableBytesSent returns a boolean if a field has been set.
func (o *MetricsDetails) HasBillableBytesSent() bool {
	if o != nil && !IsNil(o.BillableBytesSent) {
		return true
	}

	return false
}

// SetBillableBytesSent gets a reference to the given int32 and assigns it to the BillableBytesSent field.
func (o *MetricsDetails) SetBillableBytesSent(v int32) {
	o.BillableBytesSent = &v
}

// GetBillableMessageSize returns the BillableMessageSize field value if set, zero value otherwise.
func (o *MetricsDetails) GetBillableMessageSize() int32 {
	if o == nil || IsNil(o.BillableMessageSize) {
		var ret int32
		return ret
	}
	return *o.BillableMessageSize
}

// GetBillableMessageSizeOk returns a tuple with the BillableMessageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDetails) GetBillableMessageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.BillableMessageSize) {
		return nil, false
	}
	return o.BillableMessageSize, true
}

// HasBillableMessageSize returns a boolean if a field has been set.
func (o *MetricsDetails) HasBillableMessageSize() bool {
	if o != nil && !IsNil(o.BillableMessageSize) {
		return true
	}

	return false
}

// SetBillableMessageSize gets a reference to the given int32 and assigns it to the BillableMessageSize field.
func (o *MetricsDetails) SetBillableMessageSize(v int32) {
	o.BillableMessageSize = &v
}

// GetBytesReceived returns the BytesReceived field value if set, zero value otherwise.
func (o *MetricsDetails) GetBytesReceived() int32 {
	if o == nil || IsNil(o.BytesReceived) {
		var ret int32
		return ret
	}
	return *o.BytesReceived
}

// GetBytesReceivedOk returns a tuple with the BytesReceived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDetails) GetBytesReceivedOk() (*int32, bool) {
	if o == nil || IsNil(o.BytesReceived) {
		return nil, false
	}
	return o.BytesReceived, true
}

// HasBytesReceived returns a boolean if a field has been set.
func (o *MetricsDetails) HasBytesReceived() bool {
	if o != nil && !IsNil(o.BytesReceived) {
		return true
	}

	return false
}

// SetBytesReceived gets a reference to the given int32 and assigns it to the BytesReceived field.
func (o *MetricsDetails) SetBytesReceived(v int32) {
	o.BytesReceived = &v
}

// GetBytesSent returns the BytesSent field value if set, zero value otherwise.
func (o *MetricsDetails) GetBytesSent() int32 {
	if o == nil || IsNil(o.BytesSent) {
		var ret int32
		return ret
	}
	return *o.BytesSent
}

// GetBytesSentOk returns a tuple with the BytesSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDetails) GetBytesSentOk() (*int32, bool) {
	if o == nil || IsNil(o.BytesSent) {
		return nil, false
	}
	return o.BytesSent, true
}

// HasBytesSent returns a boolean if a field has been set.
func (o *MetricsDetails) HasBytesSent() bool {
	if o != nil && !IsNil(o.BytesSent) {
		return true
	}

	return false
}

// SetBytesSent gets a reference to the given int32 and assigns it to the BytesSent field.
func (o *MetricsDetails) SetBytesSent(v int32) {
	o.BytesSent = &v
}

// GetMessageSize returns the MessageSize field value if set, zero value otherwise.
func (o *MetricsDetails) GetMessageSize() int32 {
	if o == nil || IsNil(o.MessageSize) {
		var ret int32
		return ret
	}
	return *o.MessageSize
}

// GetMessageSizeOk returns a tuple with the MessageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDetails) GetMessageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageSize) {
		return nil, false
	}
	return o.MessageSize, true
}

// HasMessageSize returns a boolean if a field has been set.
func (o *MetricsDetails) HasMessageSize() bool {
	if o != nil && !IsNil(o.MessageSize) {
		return true
	}

	return false
}

// SetMessageSize gets a reference to the given int32 and assigns it to the MessageSize field.
func (o *MetricsDetails) SetMessageSize(v int32) {
	o.MessageSize = &v
}

// GetNoOfAckMessages returns the NoOfAckMessages field value if set, zero value otherwise.
func (o *MetricsDetails) GetNoOfAckMessages() int32 {
	if o == nil || IsNil(o.NoOfAckMessages) {
		var ret int32
		return ret
	}
	return *o.NoOfAckMessages
}

// GetNoOfAckMessagesOk returns a tuple with the NoOfAckMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDetails) GetNoOfAckMessagesOk() (*int32, bool) {
	if o == nil || IsNil(o.NoOfAckMessages) {
		return nil, false
	}
	return o.NoOfAckMessages, true
}

// HasNoOfAckMessages returns a boolean if a field has been set.
func (o *MetricsDetails) HasNoOfAckMessages() bool {
	if o != nil && !IsNil(o.NoOfAckMessages) {
		return true
	}

	return false
}

// SetNoOfAckMessages gets a reference to the given int32 and assigns it to the NoOfAckMessages field.
func (o *MetricsDetails) SetNoOfAckMessages(v int32) {
	o.NoOfAckMessages = &v
}

// GetNoOfCommandMessages returns the NoOfCommandMessages field value if set, zero value otherwise.
func (o *MetricsDetails) GetNoOfCommandMessages() int32 {
	if o == nil || IsNil(o.NoOfCommandMessages) {
		var ret int32
		return ret
	}
	return *o.NoOfCommandMessages
}

// GetNoOfCommandMessagesOk returns a tuple with the NoOfCommandMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDetails) GetNoOfCommandMessagesOk() (*int32, bool) {
	if o == nil || IsNil(o.NoOfCommandMessages) {
		return nil, false
	}
	return o.NoOfCommandMessages, true
}

// HasNoOfCommandMessages returns a boolean if a field has been set.
func (o *MetricsDetails) HasNoOfCommandMessages() bool {
	if o != nil && !IsNil(o.NoOfCommandMessages) {
		return true
	}

	return false
}

// SetNoOfCommandMessages gets a reference to the given int32 and assigns it to the NoOfCommandMessages field.
func (o *MetricsDetails) SetNoOfCommandMessages(v int32) {
	o.NoOfCommandMessages = &v
}

// GetNoOfConfigMessages returns the NoOfConfigMessages field value if set, zero value otherwise.
func (o *MetricsDetails) GetNoOfConfigMessages() int32 {
	if o == nil || IsNil(o.NoOfConfigMessages) {
		var ret int32
		return ret
	}
	return *o.NoOfConfigMessages
}

// GetNoOfConfigMessagesOk returns a tuple with the NoOfConfigMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDetails) GetNoOfConfigMessagesOk() (*int32, bool) {
	if o == nil || IsNil(o.NoOfConfigMessages) {
		return nil, false
	}
	return o.NoOfConfigMessages, true
}

// HasNoOfConfigMessages returns a boolean if a field has been set.
func (o *MetricsDetails) HasNoOfConfigMessages() bool {
	if o != nil && !IsNil(o.NoOfConfigMessages) {
		return true
	}

	return false
}

// SetNoOfConfigMessages gets a reference to the given int32 and assigns it to the NoOfConfigMessages field.
func (o *MetricsDetails) SetNoOfConfigMessages(v int32) {
	o.NoOfConfigMessages = &v
}

// GetNoOfDeviceConnectionsFailed returns the NoOfDeviceConnectionsFailed field value if set, zero value otherwise.
func (o *MetricsDetails) GetNoOfDeviceConnectionsFailed() int32 {
	if o == nil || IsNil(o.NoOfDeviceConnectionsFailed) {
		var ret int32
		return ret
	}
	return *o.NoOfDeviceConnectionsFailed
}

// GetNoOfDeviceConnectionsFailedOk returns a tuple with the NoOfDeviceConnectionsFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDetails) GetNoOfDeviceConnectionsFailedOk() (*int32, bool) {
	if o == nil || IsNil(o.NoOfDeviceConnectionsFailed) {
		return nil, false
	}
	return o.NoOfDeviceConnectionsFailed, true
}

// HasNoOfDeviceConnectionsFailed returns a boolean if a field has been set.
func (o *MetricsDetails) HasNoOfDeviceConnectionsFailed() bool {
	if o != nil && !IsNil(o.NoOfDeviceConnectionsFailed) {
		return true
	}

	return false
}

// SetNoOfDeviceConnectionsFailed gets a reference to the given int32 and assigns it to the NoOfDeviceConnectionsFailed field.
func (o *MetricsDetails) SetNoOfDeviceConnectionsFailed(v int32) {
	o.NoOfDeviceConnectionsFailed = &v
}

// GetNoOfDevices returns the NoOfDevices field value if set, zero value otherwise.
func (o *MetricsDetails) GetNoOfDevices() int32 {
	if o == nil || IsNil(o.NoOfDevices) {
		var ret int32
		return ret
	}
	return *o.NoOfDevices
}

// GetNoOfDevicesOk returns a tuple with the NoOfDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDetails) GetNoOfDevicesOk() (*int32, bool) {
	if o == nil || IsNil(o.NoOfDevices) {
		return nil, false
	}
	return o.NoOfDevices, true
}

// HasNoOfDevices returns a boolean if a field has been set.
func (o *MetricsDetails) HasNoOfDevices() bool {
	if o != nil && !IsNil(o.NoOfDevices) {
		return true
	}

	return false
}

// SetNoOfDevices gets a reference to the given int32 and assigns it to the NoOfDevices field.
func (o *MetricsDetails) SetNoOfDevices(v int32) {
	o.NoOfDevices = &v
}

// GetNoOfDisConnections returns the NoOfDisConnections field value if set, zero value otherwise.
func (o *MetricsDetails) GetNoOfDisConnections() int32 {
	if o == nil || IsNil(o.NoOfDisConnections) {
		var ret int32
		return ret
	}
	return *o.NoOfDisConnections
}

// GetNoOfDisConnectionsOk returns a tuple with the NoOfDisConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDetails) GetNoOfDisConnectionsOk() (*int32, bool) {
	if o == nil || IsNil(o.NoOfDisConnections) {
		return nil, false
	}
	return o.NoOfDisConnections, true
}

// HasNoOfDisConnections returns a boolean if a field has been set.
func (o *MetricsDetails) HasNoOfDisConnections() bool {
	if o != nil && !IsNil(o.NoOfDisConnections) {
		return true
	}

	return false
}

// SetNoOfDisConnections gets a reference to the given int32 and assigns it to the NoOfDisConnections field.
func (o *MetricsDetails) SetNoOfDisConnections(v int32) {
	o.NoOfDisConnections = &v
}

// GetNoOfEventMessages returns the NoOfEventMessages field value if set, zero value otherwise.
func (o *MetricsDetails) GetNoOfEventMessages() int32 {
	if o == nil || IsNil(o.NoOfEventMessages) {
		var ret int32
		return ret
	}
	return *o.NoOfEventMessages
}

// GetNoOfEventMessagesOk returns a tuple with the NoOfEventMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDetails) GetNoOfEventMessagesOk() (*int32, bool) {
	if o == nil || IsNil(o.NoOfEventMessages) {
		return nil, false
	}
	return o.NoOfEventMessages, true
}

// HasNoOfEventMessages returns a boolean if a field has been set.
func (o *MetricsDetails) HasNoOfEventMessages() bool {
	if o != nil && !IsNil(o.NoOfEventMessages) {
		return true
	}

	return false
}

// SetNoOfEventMessages gets a reference to the given int32 and assigns it to the NoOfEventMessages field.
func (o *MetricsDetails) SetNoOfEventMessages(v int32) {
	o.NoOfEventMessages = &v
}

// GetNoOfGatewayConnectionsFailed returns the NoOfGatewayConnectionsFailed field value if set, zero value otherwise.
func (o *MetricsDetails) GetNoOfGatewayConnectionsFailed() int32 {
	if o == nil || IsNil(o.NoOfGatewayConnectionsFailed) {
		var ret int32
		return ret
	}
	return *o.NoOfGatewayConnectionsFailed
}

// GetNoOfGatewayConnectionsFailedOk returns a tuple with the NoOfGatewayConnectionsFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDetails) GetNoOfGatewayConnectionsFailedOk() (*int32, bool) {
	if o == nil || IsNil(o.NoOfGatewayConnectionsFailed) {
		return nil, false
	}
	return o.NoOfGatewayConnectionsFailed, true
}

// HasNoOfGatewayConnectionsFailed returns a boolean if a field has been set.
func (o *MetricsDetails) HasNoOfGatewayConnectionsFailed() bool {
	if o != nil && !IsNil(o.NoOfGatewayConnectionsFailed) {
		return true
	}

	return false
}

// SetNoOfGatewayConnectionsFailed gets a reference to the given int32 and assigns it to the NoOfGatewayConnectionsFailed field.
func (o *MetricsDetails) SetNoOfGatewayConnectionsFailed(v int32) {
	o.NoOfGatewayConnectionsFailed = &v
}

// GetNoOfGateways returns the NoOfGateways field value if set, zero value otherwise.
func (o *MetricsDetails) GetNoOfGateways() int32 {
	if o == nil || IsNil(o.NoOfGateways) {
		var ret int32
		return ret
	}
	return *o.NoOfGateways
}

// GetNoOfGatewaysOk returns a tuple with the NoOfGateways field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDetails) GetNoOfGatewaysOk() (*int32, bool) {
	if o == nil || IsNil(o.NoOfGateways) {
		return nil, false
	}
	return o.NoOfGateways, true
}

// HasNoOfGateways returns a boolean if a field has been set.
func (o *MetricsDetails) HasNoOfGateways() bool {
	if o != nil && !IsNil(o.NoOfGateways) {
		return true
	}

	return false
}

// SetNoOfGateways gets a reference to the given int32 and assigns it to the NoOfGateways field.
func (o *MetricsDetails) SetNoOfGateways(v int32) {
	o.NoOfGateways = &v
}

// GetNoOfLoopBackMessages returns the NoOfLoopBackMessages field value if set, zero value otherwise.
func (o *MetricsDetails) GetNoOfLoopBackMessages() int32 {
	if o == nil || IsNil(o.NoOfLoopBackMessages) {
		var ret int32
		return ret
	}
	return *o.NoOfLoopBackMessages
}

// GetNoOfLoopBackMessagesOk returns a tuple with the NoOfLoopBackMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDetails) GetNoOfLoopBackMessagesOk() (*int32, bool) {
	if o == nil || IsNil(o.NoOfLoopBackMessages) {
		return nil, false
	}
	return o.NoOfLoopBackMessages, true
}

// HasNoOfLoopBackMessages returns a boolean if a field has been set.
func (o *MetricsDetails) HasNoOfLoopBackMessages() bool {
	if o != nil && !IsNil(o.NoOfLoopBackMessages) {
		return true
	}

	return false
}

// SetNoOfLoopBackMessages gets a reference to the given int32 and assigns it to the NoOfLoopBackMessages field.
func (o *MetricsDetails) SetNoOfLoopBackMessages(v int32) {
	o.NoOfLoopBackMessages = &v
}

// GetNoOfMessages returns the NoOfMessages field value if set, zero value otherwise.
func (o *MetricsDetails) GetNoOfMessages() int32 {
	if o == nil || IsNil(o.NoOfMessages) {
		var ret int32
		return ret
	}
	return *o.NoOfMessages
}

// GetNoOfMessagesOk returns a tuple with the NoOfMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDetails) GetNoOfMessagesOk() (*int32, bool) {
	if o == nil || IsNil(o.NoOfMessages) {
		return nil, false
	}
	return o.NoOfMessages, true
}

// HasNoOfMessages returns a boolean if a field has been set.
func (o *MetricsDetails) HasNoOfMessages() bool {
	if o != nil && !IsNil(o.NoOfMessages) {
		return true
	}

	return false
}

// SetNoOfMessages gets a reference to the given int32 and assigns it to the NoOfMessages field.
func (o *MetricsDetails) SetNoOfMessages(v int32) {
	o.NoOfMessages = &v
}

// GetNoOfPublishErrors returns the NoOfPublishErrors field value if set, zero value otherwise.
func (o *MetricsDetails) GetNoOfPublishErrors() int32 {
	if o == nil || IsNil(o.NoOfPublishErrors) {
		var ret int32
		return ret
	}
	return *o.NoOfPublishErrors
}

// GetNoOfPublishErrorsOk returns a tuple with the NoOfPublishErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDetails) GetNoOfPublishErrorsOk() (*int32, bool) {
	if o == nil || IsNil(o.NoOfPublishErrors) {
		return nil, false
	}
	return o.NoOfPublishErrors, true
}

// HasNoOfPublishErrors returns a boolean if a field has been set.
func (o *MetricsDetails) HasNoOfPublishErrors() bool {
	if o != nil && !IsNil(o.NoOfPublishErrors) {
		return true
	}

	return false
}

// SetNoOfPublishErrors gets a reference to the given int32 and assigns it to the NoOfPublishErrors field.
func (o *MetricsDetails) SetNoOfPublishErrors(v int32) {
	o.NoOfPublishErrors = &v
}

// GetNoOfRegistries returns the NoOfRegistries field value if set, zero value otherwise.
func (o *MetricsDetails) GetNoOfRegistries() int32 {
	if o == nil || IsNil(o.NoOfRegistries) {
		var ret int32
		return ret
	}
	return *o.NoOfRegistries
}

// GetNoOfRegistriesOk returns a tuple with the NoOfRegistries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDetails) GetNoOfRegistriesOk() (*int32, bool) {
	if o == nil || IsNil(o.NoOfRegistries) {
		return nil, false
	}
	return o.NoOfRegistries, true
}

// HasNoOfRegistries returns a boolean if a field has been set.
func (o *MetricsDetails) HasNoOfRegistries() bool {
	if o != nil && !IsNil(o.NoOfRegistries) {
		return true
	}

	return false
}

// SetNoOfRegistries gets a reference to the given int32 and assigns it to the NoOfRegistries field.
func (o *MetricsDetails) SetNoOfRegistries(v int32) {
	o.NoOfRegistries = &v
}

// GetNoOfStateMessages returns the NoOfStateMessages field value if set, zero value otherwise.
func (o *MetricsDetails) GetNoOfStateMessages() int32 {
	if o == nil || IsNil(o.NoOfStateMessages) {
		var ret int32
		return ret
	}
	return *o.NoOfStateMessages
}

// GetNoOfStateMessagesOk returns a tuple with the NoOfStateMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDetails) GetNoOfStateMessagesOk() (*int32, bool) {
	if o == nil || IsNil(o.NoOfStateMessages) {
		return nil, false
	}
	return o.NoOfStateMessages, true
}

// HasNoOfStateMessages returns a boolean if a field has been set.
func (o *MetricsDetails) HasNoOfStateMessages() bool {
	if o != nil && !IsNil(o.NoOfStateMessages) {
		return true
	}

	return false
}

// SetNoOfStateMessages gets a reference to the given int32 and assigns it to the NoOfStateMessages field.
func (o *MetricsDetails) SetNoOfStateMessages(v int32) {
	o.NoOfStateMessages = &v
}

// GetNoOfSubscribe returns the NoOfSubscribe field value if set, zero value otherwise.
func (o *MetricsDetails) GetNoOfSubscribe() int32 {
	if o == nil || IsNil(o.NoOfSubscribe) {
		var ret int32
		return ret
	}
	return *o.NoOfSubscribe
}

// GetNoOfSubscribeOk returns a tuple with the NoOfSubscribe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDetails) GetNoOfSubscribeOk() (*int32, bool) {
	if o == nil || IsNil(o.NoOfSubscribe) {
		return nil, false
	}
	return o.NoOfSubscribe, true
}

// HasNoOfSubscribe returns a boolean if a field has been set.
func (o *MetricsDetails) HasNoOfSubscribe() bool {
	if o != nil && !IsNil(o.NoOfSubscribe) {
		return true
	}

	return false
}

// SetNoOfSubscribe gets a reference to the given int32 and assigns it to the NoOfSubscribe field.
func (o *MetricsDetails) SetNoOfSubscribe(v int32) {
	o.NoOfSubscribe = &v
}

// GetNoOfSuccessfulConnections returns the NoOfSuccessfulConnections field value if set, zero value otherwise.
func (o *MetricsDetails) GetNoOfSuccessfulConnections() int32 {
	if o == nil || IsNil(o.NoOfSuccessfulConnections) {
		var ret int32
		return ret
	}
	return *o.NoOfSuccessfulConnections
}

// GetNoOfSuccessfulConnectionsOk returns a tuple with the NoOfSuccessfulConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDetails) GetNoOfSuccessfulConnectionsOk() (*int32, bool) {
	if o == nil || IsNil(o.NoOfSuccessfulConnections) {
		return nil, false
	}
	return o.NoOfSuccessfulConnections, true
}

// HasNoOfSuccessfulConnections returns a boolean if a field has been set.
func (o *MetricsDetails) HasNoOfSuccessfulConnections() bool {
	if o != nil && !IsNil(o.NoOfSuccessfulConnections) {
		return true
	}

	return false
}

// SetNoOfSuccessfulConnections gets a reference to the given int32 and assigns it to the NoOfSuccessfulConnections field.
func (o *MetricsDetails) SetNoOfSuccessfulConnections(v int32) {
	o.NoOfSuccessfulConnections = &v
}

// GetNoOfUnSubscribe returns the NoOfUnSubscribe field value if set, zero value otherwise.
func (o *MetricsDetails) GetNoOfUnSubscribe() int32 {
	if o == nil || IsNil(o.NoOfUnSubscribe) {
		var ret int32
		return ret
	}
	return *o.NoOfUnSubscribe
}

// GetNoOfUnSubscribeOk returns a tuple with the NoOfUnSubscribe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDetails) GetNoOfUnSubscribeOk() (*int32, bool) {
	if o == nil || IsNil(o.NoOfUnSubscribe) {
		return nil, false
	}
	return o.NoOfUnSubscribe, true
}

// HasNoOfUnSubscribe returns a boolean if a field has been set.
func (o *MetricsDetails) HasNoOfUnSubscribe() bool {
	if o != nil && !IsNil(o.NoOfUnSubscribe) {
		return true
	}

	return false
}

// SetNoOfUnSubscribe gets a reference to the given int32 and assigns it to the NoOfUnSubscribe field.
func (o *MetricsDetails) SetNoOfUnSubscribe(v int32) {
	o.NoOfUnSubscribe = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *MetricsDetails) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDetails) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *MetricsDetails) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *MetricsDetails) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

func (o MetricsDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NoOfMessagesFor30Minutes) {
		toSerialize["NoOfMessagesFor30Minutes"] = o.NoOfMessagesFor30Minutes
	}
	if !IsNil(o.NoOfMessagesFor48Hours) {
		toSerialize["NoOfMessagesFor48Hours"] = o.NoOfMessagesFor48Hours
	}
	if !IsNil(o.BillableBytesReceived) {
		toSerialize["billableBytesReceived"] = o.BillableBytesReceived
	}
	if !IsNil(o.BillableBytesSent) {
		toSerialize["billableBytesSent"] = o.BillableBytesSent
	}
	if !IsNil(o.BillableMessageSize) {
		toSerialize["billableMessageSize"] = o.BillableMessageSize
	}
	if !IsNil(o.BytesReceived) {
		toSerialize["bytesReceived"] = o.BytesReceived
	}
	if !IsNil(o.BytesSent) {
		toSerialize["bytesSent"] = o.BytesSent
	}
	if !IsNil(o.MessageSize) {
		toSerialize["messageSize"] = o.MessageSize
	}
	if !IsNil(o.NoOfAckMessages) {
		toSerialize["noOfAckMessages"] = o.NoOfAckMessages
	}
	if !IsNil(o.NoOfCommandMessages) {
		toSerialize["noOfCommandMessages"] = o.NoOfCommandMessages
	}
	if !IsNil(o.NoOfConfigMessages) {
		toSerialize["noOfConfigMessages"] = o.NoOfConfigMessages
	}
	if !IsNil(o.NoOfDeviceConnectionsFailed) {
		toSerialize["noOfDeviceConnectionsFailed"] = o.NoOfDeviceConnectionsFailed
	}
	if !IsNil(o.NoOfDevices) {
		toSerialize["noOfDevices"] = o.NoOfDevices
	}
	if !IsNil(o.NoOfDisConnections) {
		toSerialize["noOfDisConnections"] = o.NoOfDisConnections
	}
	if !IsNil(o.NoOfEventMessages) {
		toSerialize["noOfEventMessages"] = o.NoOfEventMessages
	}
	if !IsNil(o.NoOfGatewayConnectionsFailed) {
		toSerialize["noOfGatewayConnectionsFailed"] = o.NoOfGatewayConnectionsFailed
	}
	if !IsNil(o.NoOfGateways) {
		toSerialize["noOfGateways"] = o.NoOfGateways
	}
	if !IsNil(o.NoOfLoopBackMessages) {
		toSerialize["noOfLoopBackMessages"] = o.NoOfLoopBackMessages
	}
	if !IsNil(o.NoOfMessages) {
		toSerialize["noOfMessages"] = o.NoOfMessages
	}
	if !IsNil(o.NoOfPublishErrors) {
		toSerialize["noOfPublishErrors"] = o.NoOfPublishErrors
	}
	if !IsNil(o.NoOfRegistries) {
		toSerialize["noOfRegistries"] = o.NoOfRegistries
	}
	if !IsNil(o.NoOfStateMessages) {
		toSerialize["noOfStateMessages"] = o.NoOfStateMessages
	}
	if !IsNil(o.NoOfSubscribe) {
		toSerialize["noOfSubscribe"] = o.NoOfSubscribe
	}
	if !IsNil(o.NoOfSuccessfulConnections) {
		toSerialize["noOfSuccessfulConnections"] = o.NoOfSuccessfulConnections
	}
	if !IsNil(o.NoOfUnSubscribe) {
		toSerialize["noOfUnSubscribe"] = o.NoOfUnSubscribe
	}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	return toSerialize, nil
}

type NullableMetricsDetails struct {
	value *MetricsDetails
	isSet bool
}

func (v NullableMetricsDetails) Get() *MetricsDetails {
	return v.value
}

func (v *NullableMetricsDetails) Set(val *MetricsDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsDetails(val *MetricsDetails) *NullableMetricsDetails {
	return &NullableMetricsDetails{value: val, isSet: true}
}

func (v NullableMetricsDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


