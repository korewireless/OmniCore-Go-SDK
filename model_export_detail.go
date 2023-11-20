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

// checks if the ExportDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportDetail{}

// ExportDetail struct for ExportDetail
type ExportDetail struct {
	Subscription *string `json:"subscription,omitempty"`
	Name *string `json:"name,omitempty"`
	Source *string `json:"source,omitempty"`
	Status *string `json:"status,omitempty"`
	Destination *string `json:"destination,omitempty"`
	CreatedOn *string `json:"createdOn,omitempty"`
	Nooffiles *int32 `json:"nooffiles,omitempty"`
	FileSize *float32 `json:"fileSize,omitempty"`
	Done *bool `json:"done,omitempty"`
}

// NewExportDetail instantiates a new ExportDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportDetail() *ExportDetail {
	this := ExportDetail{}
	return &this
}

// NewExportDetailWithDefaults instantiates a new ExportDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportDetailWithDefaults() *ExportDetail {
	this := ExportDetail{}
	return &this
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *ExportDetail) GetSubscription() string {
	if o == nil || IsNil(o.Subscription) {
		var ret string
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportDetail) GetSubscriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *ExportDetail) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given string and assigns it to the Subscription field.
func (o *ExportDetail) SetSubscription(v string) {
	o.Subscription = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExportDetail) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportDetail) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExportDetail) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExportDetail) SetName(v string) {
	o.Name = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ExportDetail) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportDetail) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ExportDetail) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ExportDetail) SetSource(v string) {
	o.Source = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ExportDetail) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportDetail) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ExportDetail) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ExportDetail) SetStatus(v string) {
	o.Status = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *ExportDetail) GetDestination() string {
	if o == nil || IsNil(o.Destination) {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportDetail) GetDestinationOk() (*string, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *ExportDetail) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *ExportDetail) SetDestination(v string) {
	o.Destination = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *ExportDetail) GetCreatedOn() string {
	if o == nil || IsNil(o.CreatedOn) {
		var ret string
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportDetail) GetCreatedOnOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedOn) {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *ExportDetail) HasCreatedOn() bool {
	if o != nil && !IsNil(o.CreatedOn) {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given string and assigns it to the CreatedOn field.
func (o *ExportDetail) SetCreatedOn(v string) {
	o.CreatedOn = &v
}

// GetNooffiles returns the Nooffiles field value if set, zero value otherwise.
func (o *ExportDetail) GetNooffiles() int32 {
	if o == nil || IsNil(o.Nooffiles) {
		var ret int32
		return ret
	}
	return *o.Nooffiles
}

// GetNooffilesOk returns a tuple with the Nooffiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportDetail) GetNooffilesOk() (*int32, bool) {
	if o == nil || IsNil(o.Nooffiles) {
		return nil, false
	}
	return o.Nooffiles, true
}

// HasNooffiles returns a boolean if a field has been set.
func (o *ExportDetail) HasNooffiles() bool {
	if o != nil && !IsNil(o.Nooffiles) {
		return true
	}

	return false
}

// SetNooffiles gets a reference to the given int32 and assigns it to the Nooffiles field.
func (o *ExportDetail) SetNooffiles(v int32) {
	o.Nooffiles = &v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *ExportDetail) GetFileSize() float32 {
	if o == nil || IsNil(o.FileSize) {
		var ret float32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportDetail) GetFileSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.FileSize) {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *ExportDetail) HasFileSize() bool {
	if o != nil && !IsNil(o.FileSize) {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given float32 and assigns it to the FileSize field.
func (o *ExportDetail) SetFileSize(v float32) {
	o.FileSize = &v
}

// GetDone returns the Done field value if set, zero value otherwise.
func (o *ExportDetail) GetDone() bool {
	if o == nil || IsNil(o.Done) {
		var ret bool
		return ret
	}
	return *o.Done
}

// GetDoneOk returns a tuple with the Done field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportDetail) GetDoneOk() (*bool, bool) {
	if o == nil || IsNil(o.Done) {
		return nil, false
	}
	return o.Done, true
}

// HasDone returns a boolean if a field has been set.
func (o *ExportDetail) HasDone() bool {
	if o != nil && !IsNil(o.Done) {
		return true
	}

	return false
}

// SetDone gets a reference to the given bool and assigns it to the Done field.
func (o *ExportDetail) SetDone(v bool) {
	o.Done = &v
}

func (o ExportDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExportDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subscription) {
		toSerialize["subscription"] = o.Subscription
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !IsNil(o.CreatedOn) {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if !IsNil(o.Nooffiles) {
		toSerialize["nooffiles"] = o.Nooffiles
	}
	if !IsNil(o.FileSize) {
		toSerialize["fileSize"] = o.FileSize
	}
	if !IsNil(o.Done) {
		toSerialize["done"] = o.Done
	}
	return toSerialize, nil
}

type NullableExportDetail struct {
	value *ExportDetail
	isSet bool
}

func (v NullableExportDetail) Get() *ExportDetail {
	return v.value
}

func (v *NullableExportDetail) Set(val *ExportDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableExportDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableExportDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportDetail(val *ExportDetail) *NullableExportDetail {
	return &NullableExportDetail{value: val, isSet: true}
}

func (v NullableExportDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


