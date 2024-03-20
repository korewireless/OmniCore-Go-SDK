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

// checks if the TcpUdpModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TcpUdpModel{}

// TcpUdpModel struct for TcpUdpModel
type TcpUdpModel struct {
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

// NewTcpUdpModel instantiates a new TcpUdpModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTcpUdpModel(image TcpUdpImage, tcpDetails TcpUdpPortDetail, udpDetails TcpUdpPortDetail) *TcpUdpModel {
	this := TcpUdpModel{}
	this.Image = image
	this.TcpDetails = tcpDetails
	this.UdpDetails = udpDetails
	return &this
}

// NewTcpUdpModelWithDefaults instantiates a new TcpUdpModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTcpUdpModelWithDefaults() *TcpUdpModel {
	this := TcpUdpModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TcpUdpModel) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TcpUdpModel) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TcpUdpModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TcpUdpModel) SetId(v int32) {
	o.Id = &v
}

// GetModelName returns the ModelName field value if set, zero value otherwise.
func (o *TcpUdpModel) GetModelName() string {
	if o == nil || IsNil(o.ModelName) {
		var ret string
		return ret
	}
	return *o.ModelName
}

// GetModelNameOk returns a tuple with the ModelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TcpUdpModel) GetModelNameOk() (*string, bool) {
	if o == nil || IsNil(o.ModelName) {
		return nil, false
	}
	return o.ModelName, true
}

// HasModelName returns a boolean if a field has been set.
func (o *TcpUdpModel) HasModelName() bool {
	if o != nil && !IsNil(o.ModelName) {
		return true
	}

	return false
}

// SetModelName gets a reference to the given string and assigns it to the ModelName field.
func (o *TcpUdpModel) SetModelName(v string) {
	o.ModelName = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise.
func (o *TcpUdpModel) GetManufacturer() string {
	if o == nil || IsNil(o.Manufacturer) {
		var ret string
		return ret
	}
	return *o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TcpUdpModel) GetManufacturerOk() (*string, bool) {
	if o == nil || IsNil(o.Manufacturer) {
		return nil, false
	}
	return o.Manufacturer, true
}

// HasManufacturer returns a boolean if a field has been set.
func (o *TcpUdpModel) HasManufacturer() bool {
	if o != nil && !IsNil(o.Manufacturer) {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given string and assigns it to the Manufacturer field.
func (o *TcpUdpModel) SetManufacturer(v string) {
	o.Manufacturer = &v
}

// GetImage returns the Image field value
func (o *TcpUdpModel) GetImage() TcpUdpImage {
	if o == nil {
		var ret TcpUdpImage
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *TcpUdpModel) GetImageOk() (*TcpUdpImage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *TcpUdpModel) SetImage(v TcpUdpImage) {
	o.Image = v
}

// GetTcpDetails returns the TcpDetails field value
func (o *TcpUdpModel) GetTcpDetails() TcpUdpPortDetail {
	if o == nil {
		var ret TcpUdpPortDetail
		return ret
	}

	return o.TcpDetails
}

// GetTcpDetailsOk returns a tuple with the TcpDetails field value
// and a boolean to check if the value has been set.
func (o *TcpUdpModel) GetTcpDetailsOk() (*TcpUdpPortDetail, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TcpDetails, true
}

// SetTcpDetails sets field value
func (o *TcpUdpModel) SetTcpDetails(v TcpUdpPortDetail) {
	o.TcpDetails = v
}

// GetUdpDetails returns the UdpDetails field value
func (o *TcpUdpModel) GetUdpDetails() TcpUdpPortDetail {
	if o == nil {
		var ret TcpUdpPortDetail
		return ret
	}

	return o.UdpDetails
}

// GetUdpDetailsOk returns a tuple with the UdpDetails field value
// and a boolean to check if the value has been set.
func (o *TcpUdpModel) GetUdpDetailsOk() (*TcpUdpPortDetail, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UdpDetails, true
}

// SetUdpDetails sets field value
func (o *TcpUdpModel) SetUdpDetails(v TcpUdpPortDetail) {
	o.UdpDetails = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *TcpUdpModel) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TcpUdpModel) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *TcpUdpModel) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *TcpUdpModel) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TcpUdpModel) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TcpUdpModel) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TcpUdpModel) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *TcpUdpModel) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TcpUdpModel) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TcpUdpModel) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TcpUdpModel) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *TcpUdpModel) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o TcpUdpModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TcpUdpModel) ToMap() (map[string]interface{}, error) {
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

type NullableTcpUdpModel struct {
	value *TcpUdpModel
	isSet bool
}

func (v NullableTcpUdpModel) Get() *TcpUdpModel {
	return v.value
}

func (v *NullableTcpUdpModel) Set(val *TcpUdpModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTcpUdpModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTcpUdpModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTcpUdpModel(val *TcpUdpModel) *NullableTcpUdpModel {
	return &NullableTcpUdpModel{value: val, isSet: true}
}

func (v NullableTcpUdpModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTcpUdpModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


