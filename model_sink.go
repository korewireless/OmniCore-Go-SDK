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

// checks if the Sink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Sink{}

// Sink struct for Sink
type Sink struct {
	Id *string `json:"id,omitempty"`
	Subscription *string `json:"subscription,omitempty"`
	Sink *string `json:"sink,omitempty"`
	Config *Config `json:"config,omitempty"`
	Status *bool `json:"status,omitempty"`
	Createdon *string `json:"createdon,omitempty"`
	Updatedon *string `json:"updatedon,omitempty"`
}

// NewSink instantiates a new Sink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSink() *Sink {
	this := Sink{}
	return &this
}

// NewSinkWithDefaults instantiates a new Sink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSinkWithDefaults() *Sink {
	this := Sink{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Sink) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sink) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Sink) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Sink) SetId(v string) {
	o.Id = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *Sink) GetSubscription() string {
	if o == nil || IsNil(o.Subscription) {
		var ret string
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sink) GetSubscriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *Sink) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given string and assigns it to the Subscription field.
func (o *Sink) SetSubscription(v string) {
	o.Subscription = &v
}

// GetSink returns the Sink field value if set, zero value otherwise.
func (o *Sink) GetSink() string {
	if o == nil || IsNil(o.Sink) {
		var ret string
		return ret
	}
	return *o.Sink
}

// GetSinkOk returns a tuple with the Sink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sink) GetSinkOk() (*string, bool) {
	if o == nil || IsNil(o.Sink) {
		return nil, false
	}
	return o.Sink, true
}

// HasSink returns a boolean if a field has been set.
func (o *Sink) HasSink() bool {
	if o != nil && !IsNil(o.Sink) {
		return true
	}

	return false
}

// SetSink gets a reference to the given string and assigns it to the Sink field.
func (o *Sink) SetSink(v string) {
	o.Sink = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *Sink) GetConfig() Config {
	if o == nil || IsNil(o.Config) {
		var ret Config
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sink) GetConfigOk() (*Config, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *Sink) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given Config and assigns it to the Config field.
func (o *Sink) SetConfig(v Config) {
	o.Config = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Sink) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sink) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Sink) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *Sink) SetStatus(v bool) {
	o.Status = &v
}

// GetCreatedon returns the Createdon field value if set, zero value otherwise.
func (o *Sink) GetCreatedon() string {
	if o == nil || IsNil(o.Createdon) {
		var ret string
		return ret
	}
	return *o.Createdon
}

// GetCreatedonOk returns a tuple with the Createdon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sink) GetCreatedonOk() (*string, bool) {
	if o == nil || IsNil(o.Createdon) {
		return nil, false
	}
	return o.Createdon, true
}

// HasCreatedon returns a boolean if a field has been set.
func (o *Sink) HasCreatedon() bool {
	if o != nil && !IsNil(o.Createdon) {
		return true
	}

	return false
}

// SetCreatedon gets a reference to the given string and assigns it to the Createdon field.
func (o *Sink) SetCreatedon(v string) {
	o.Createdon = &v
}

// GetUpdatedon returns the Updatedon field value if set, zero value otherwise.
func (o *Sink) GetUpdatedon() string {
	if o == nil || IsNil(o.Updatedon) {
		var ret string
		return ret
	}
	return *o.Updatedon
}

// GetUpdatedonOk returns a tuple with the Updatedon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sink) GetUpdatedonOk() (*string, bool) {
	if o == nil || IsNil(o.Updatedon) {
		return nil, false
	}
	return o.Updatedon, true
}

// HasUpdatedon returns a boolean if a field has been set.
func (o *Sink) HasUpdatedon() bool {
	if o != nil && !IsNil(o.Updatedon) {
		return true
	}

	return false
}

// SetUpdatedon gets a reference to the given string and assigns it to the Updatedon field.
func (o *Sink) SetUpdatedon(v string) {
	o.Updatedon = &v
}

func (o Sink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Sink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: subscription is readOnly
	if !IsNil(o.Sink) {
		toSerialize["sink"] = o.Sink
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	// skip: status is readOnly
	// skip: createdon is readOnly
	// skip: updatedon is readOnly
	return toSerialize, nil
}

type NullableSink struct {
	value *Sink
	isSet bool
}

func (v NullableSink) Get() *Sink {
	return v.value
}

func (v *NullableSink) Set(val *Sink) {
	v.value = val
	v.isSet = true
}

func (v NullableSink) IsSet() bool {
	return v.isSet
}

func (v *NullableSink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSink(val *Sink) *NullableSink {
	return &NullableSink{value: val, isSet: true}
}

func (v NullableSink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


