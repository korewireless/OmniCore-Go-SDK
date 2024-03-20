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
	"time"
)

// checks if the CustomOnboardTcpUdpModelDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomOnboardTcpUdpModelDetails{}

// CustomOnboardTcpUdpModelDetails struct for CustomOnboardTcpUdpModelDetails
type CustomOnboardTcpUdpModelDetails struct {
	// The ID of the TCP/UDP model
	Id *int32 `json:"id,omitempty"`
	// The name of the model
	ModelName *string `json:"modelName,omitempty"`
	// The manufacturer of the model
	Manufacturer *string `json:"manufacturer,omitempty"`
	Image TcpUdpImage `json:"image"`
	TcpDetails TcpUdpPortDetail `json:"tcpDetails"`
	UdpDetails TcpUdpPortDetail `json:"udpDetails"`
	// Additional metadata in raw JSON format
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// The creation timestamp of the model
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The last update timestamp of the model
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewCustomOnboardTcpUdpModelDetails instantiates a new CustomOnboardTcpUdpModelDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomOnboardTcpUdpModelDetails(image TcpUdpImage, tcpDetails TcpUdpPortDetail, udpDetails TcpUdpPortDetail) *CustomOnboardTcpUdpModelDetails {
	this := CustomOnboardTcpUdpModelDetails{}
	this.Image = image
	this.TcpDetails = tcpDetails
	this.UdpDetails = udpDetails
	return &this
}

// NewCustomOnboardTcpUdpModelDetailsWithDefaults instantiates a new CustomOnboardTcpUdpModelDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomOnboardTcpUdpModelDetailsWithDefaults() *CustomOnboardTcpUdpModelDetails {
	this := CustomOnboardTcpUdpModelDetails{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomOnboardTcpUdpModelDetails) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboardTcpUdpModelDetails) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomOnboardTcpUdpModelDetails) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CustomOnboardTcpUdpModelDetails) SetId(v int32) {
	o.Id = &v
}

// GetModelName returns the ModelName field value if set, zero value otherwise.
func (o *CustomOnboardTcpUdpModelDetails) GetModelName() string {
	if o == nil || IsNil(o.ModelName) {
		var ret string
		return ret
	}
	return *o.ModelName
}

// GetModelNameOk returns a tuple with the ModelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboardTcpUdpModelDetails) GetModelNameOk() (*string, bool) {
	if o == nil || IsNil(o.ModelName) {
		return nil, false
	}
	return o.ModelName, true
}

// HasModelName returns a boolean if a field has been set.
func (o *CustomOnboardTcpUdpModelDetails) HasModelName() bool {
	if o != nil && !IsNil(o.ModelName) {
		return true
	}

	return false
}

// SetModelName gets a reference to the given string and assigns it to the ModelName field.
func (o *CustomOnboardTcpUdpModelDetails) SetModelName(v string) {
	o.ModelName = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise.
func (o *CustomOnboardTcpUdpModelDetails) GetManufacturer() string {
	if o == nil || IsNil(o.Manufacturer) {
		var ret string
		return ret
	}
	return *o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboardTcpUdpModelDetails) GetManufacturerOk() (*string, bool) {
	if o == nil || IsNil(o.Manufacturer) {
		return nil, false
	}
	return o.Manufacturer, true
}

// HasManufacturer returns a boolean if a field has been set.
func (o *CustomOnboardTcpUdpModelDetails) HasManufacturer() bool {
	if o != nil && !IsNil(o.Manufacturer) {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given string and assigns it to the Manufacturer field.
func (o *CustomOnboardTcpUdpModelDetails) SetManufacturer(v string) {
	o.Manufacturer = &v
}

// GetImage returns the Image field value
func (o *CustomOnboardTcpUdpModelDetails) GetImage() TcpUdpImage {
	if o == nil {
		var ret TcpUdpImage
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *CustomOnboardTcpUdpModelDetails) GetImageOk() (*TcpUdpImage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *CustomOnboardTcpUdpModelDetails) SetImage(v TcpUdpImage) {
	o.Image = v
}

// GetTcpDetails returns the TcpDetails field value
func (o *CustomOnboardTcpUdpModelDetails) GetTcpDetails() TcpUdpPortDetail {
	if o == nil {
		var ret TcpUdpPortDetail
		return ret
	}

	return o.TcpDetails
}

// GetTcpDetailsOk returns a tuple with the TcpDetails field value
// and a boolean to check if the value has been set.
func (o *CustomOnboardTcpUdpModelDetails) GetTcpDetailsOk() (*TcpUdpPortDetail, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TcpDetails, true
}

// SetTcpDetails sets field value
func (o *CustomOnboardTcpUdpModelDetails) SetTcpDetails(v TcpUdpPortDetail) {
	o.TcpDetails = v
}

// GetUdpDetails returns the UdpDetails field value
func (o *CustomOnboardTcpUdpModelDetails) GetUdpDetails() TcpUdpPortDetail {
	if o == nil {
		var ret TcpUdpPortDetail
		return ret
	}

	return o.UdpDetails
}

// GetUdpDetailsOk returns a tuple with the UdpDetails field value
// and a boolean to check if the value has been set.
func (o *CustomOnboardTcpUdpModelDetails) GetUdpDetailsOk() (*TcpUdpPortDetail, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UdpDetails, true
}

// SetUdpDetails sets field value
func (o *CustomOnboardTcpUdpModelDetails) SetUdpDetails(v TcpUdpPortDetail) {
	o.UdpDetails = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CustomOnboardTcpUdpModelDetails) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboardTcpUdpModelDetails) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CustomOnboardTcpUdpModelDetails) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *CustomOnboardTcpUdpModelDetails) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CustomOnboardTcpUdpModelDetails) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboardTcpUdpModelDetails) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CustomOnboardTcpUdpModelDetails) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CustomOnboardTcpUdpModelDetails) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CustomOnboardTcpUdpModelDetails) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOnboardTcpUdpModelDetails) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CustomOnboardTcpUdpModelDetails) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CustomOnboardTcpUdpModelDetails) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o CustomOnboardTcpUdpModelDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomOnboardTcpUdpModelDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	if !IsNil(o.ModelName) {
		toSerialize["modelName"] = o.ModelName
	}
	if !IsNil(o.Manufacturer) {
		toSerialize["manufacturer"] = o.Manufacturer
	}
	toSerialize["image"] = o.Image
	toSerialize["tcpDetails"] = o.TcpDetails
	toSerialize["udpDetails"] = o.UdpDetails
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	// skip: createdAt is readOnly
	// skip: updatedAt is readOnly
	return toSerialize, nil
}

type NullableCustomOnboardTcpUdpModelDetails struct {
	value *CustomOnboardTcpUdpModelDetails
	isSet bool
}

func (v NullableCustomOnboardTcpUdpModelDetails) Get() *CustomOnboardTcpUdpModelDetails {
	return v.value
}

func (v *NullableCustomOnboardTcpUdpModelDetails) Set(val *CustomOnboardTcpUdpModelDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomOnboardTcpUdpModelDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomOnboardTcpUdpModelDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomOnboardTcpUdpModelDetails(val *CustomOnboardTcpUdpModelDetails) *NullableCustomOnboardTcpUdpModelDetails {
	return &NullableCustomOnboardTcpUdpModelDetails{value: val, isSet: true}
}

func (v NullableCustomOnboardTcpUdpModelDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomOnboardTcpUdpModelDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


