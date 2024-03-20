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
	"os"
)

// checks if the TcpUdpImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TcpUdpImage{}

// TcpUdpImage struct for TcpUdpImage
type TcpUdpImage struct {
	// Indicates if the image is valid
	Valid bool `json:"valid"`
	// The content type of the image
	ContentType *string `json:"contentType,omitempty"`
	// The storage key of the image
	StorageKey *string `json:"storageKey,omitempty"`
	// The link to the image
	Link *string `json:"link,omitempty"`
	// The binary data of the image
	Data **os.File `json:"data,omitempty"`
}

// NewTcpUdpImage instantiates a new TcpUdpImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTcpUdpImage(valid bool) *TcpUdpImage {
	this := TcpUdpImage{}
	this.Valid = valid
	return &this
}

// NewTcpUdpImageWithDefaults instantiates a new TcpUdpImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTcpUdpImageWithDefaults() *TcpUdpImage {
	this := TcpUdpImage{}
	return &this
}

// GetValid returns the Valid field value
func (o *TcpUdpImage) GetValid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Valid
}

// GetValidOk returns a tuple with the Valid field value
// and a boolean to check if the value has been set.
func (o *TcpUdpImage) GetValidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Valid, true
}

// SetValid sets field value
func (o *TcpUdpImage) SetValid(v bool) {
	o.Valid = v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *TcpUdpImage) GetContentType() string {
	if o == nil || IsNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TcpUdpImage) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *TcpUdpImage) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *TcpUdpImage) SetContentType(v string) {
	o.ContentType = &v
}

// GetStorageKey returns the StorageKey field value if set, zero value otherwise.
func (o *TcpUdpImage) GetStorageKey() string {
	if o == nil || IsNil(o.StorageKey) {
		var ret string
		return ret
	}
	return *o.StorageKey
}

// GetStorageKeyOk returns a tuple with the StorageKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TcpUdpImage) GetStorageKeyOk() (*string, bool) {
	if o == nil || IsNil(o.StorageKey) {
		return nil, false
	}
	return o.StorageKey, true
}

// HasStorageKey returns a boolean if a field has been set.
func (o *TcpUdpImage) HasStorageKey() bool {
	if o != nil && !IsNil(o.StorageKey) {
		return true
	}

	return false
}

// SetStorageKey gets a reference to the given string and assigns it to the StorageKey field.
func (o *TcpUdpImage) SetStorageKey(v string) {
	o.StorageKey = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *TcpUdpImage) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TcpUdpImage) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *TcpUdpImage) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *TcpUdpImage) SetLink(v string) {
	o.Link = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TcpUdpImage) GetData() *os.File {
	if o == nil || IsNil(o.Data) {
		var ret *os.File
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TcpUdpImage) GetDataOk() (**os.File, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TcpUdpImage) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given *os.File and assigns it to the Data field.
func (o *TcpUdpImage) SetData(v *os.File) {
	o.Data = &v
}

func (o TcpUdpImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TcpUdpImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["valid"] = o.Valid
	if !IsNil(o.ContentType) {
		toSerialize["contentType"] = o.ContentType
	}
	if !IsNil(o.StorageKey) {
		toSerialize["storageKey"] = o.StorageKey
	}
	// skip: link is readOnly
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableTcpUdpImage struct {
	value *TcpUdpImage
	isSet bool
}

func (v NullableTcpUdpImage) Get() *TcpUdpImage {
	return v.value
}

func (v *NullableTcpUdpImage) Set(val *TcpUdpImage) {
	v.value = val
	v.isSet = true
}

func (v NullableTcpUdpImage) IsSet() bool {
	return v.isSet
}

func (v *NullableTcpUdpImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTcpUdpImage(val *TcpUdpImage) *NullableTcpUdpImage {
	return &NullableTcpUdpImage{value: val, isSet: true}
}

func (v NullableTcpUdpImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTcpUdpImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


