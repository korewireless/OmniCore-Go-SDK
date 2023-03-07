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

// checks if the X509CertificateDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &X509CertificateDetails{}

// X509CertificateDetails struct for X509CertificateDetails
type X509CertificateDetails struct {
	// ExpiryTime: The time the certificate becomes invalid.
	ExpiryTime *string `json:"expiryTime,omitempty"`
	// Issuer: The entity that signed the certificate.
	Issuer *string `json:"issuer,omitempty"`
	// PublicKeyType: The type of public key in the certificate.
	PublicKeyType *string `json:"publicKeyType,omitempty"`
	// SignatureAlgorithm: The algorithm used to sign the certificate.
	SignatureAlgorithm *string `json:"signatureAlgorithm,omitempty"`
	// StartTime: The time the certificate becomes valid.
	StartTime *string `json:"startTime,omitempty"`
	// Subject: The entity the certificate and public key belong to.
	Subject *string `json:"subject,omitempty"`
}

// NewX509CertificateDetails instantiates a new X509CertificateDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewX509CertificateDetails() *X509CertificateDetails {
	this := X509CertificateDetails{}
	return &this
}

// NewX509CertificateDetailsWithDefaults instantiates a new X509CertificateDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewX509CertificateDetailsWithDefaults() *X509CertificateDetails {
	this := X509CertificateDetails{}
	return &this
}

// GetExpiryTime returns the ExpiryTime field value if set, zero value otherwise.
func (o *X509CertificateDetails) GetExpiryTime() string {
	if o == nil || IsNil(o.ExpiryTime) {
		var ret string
		return ret
	}
	return *o.ExpiryTime
}

// GetExpiryTimeOk returns a tuple with the ExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509CertificateDetails) GetExpiryTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiryTime) {
		return nil, false
	}
	return o.ExpiryTime, true
}

// HasExpiryTime returns a boolean if a field has been set.
func (o *X509CertificateDetails) HasExpiryTime() bool {
	if o != nil && !IsNil(o.ExpiryTime) {
		return true
	}

	return false
}

// SetExpiryTime gets a reference to the given string and assigns it to the ExpiryTime field.
func (o *X509CertificateDetails) SetExpiryTime(v string) {
	o.ExpiryTime = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *X509CertificateDetails) GetIssuer() string {
	if o == nil || IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509CertificateDetails) GetIssuerOk() (*string, bool) {
	if o == nil || IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *X509CertificateDetails) HasIssuer() bool {
	if o != nil && !IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *X509CertificateDetails) SetIssuer(v string) {
	o.Issuer = &v
}

// GetPublicKeyType returns the PublicKeyType field value if set, zero value otherwise.
func (o *X509CertificateDetails) GetPublicKeyType() string {
	if o == nil || IsNil(o.PublicKeyType) {
		var ret string
		return ret
	}
	return *o.PublicKeyType
}

// GetPublicKeyTypeOk returns a tuple with the PublicKeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509CertificateDetails) GetPublicKeyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKeyType) {
		return nil, false
	}
	return o.PublicKeyType, true
}

// HasPublicKeyType returns a boolean if a field has been set.
func (o *X509CertificateDetails) HasPublicKeyType() bool {
	if o != nil && !IsNil(o.PublicKeyType) {
		return true
	}

	return false
}

// SetPublicKeyType gets a reference to the given string and assigns it to the PublicKeyType field.
func (o *X509CertificateDetails) SetPublicKeyType(v string) {
	o.PublicKeyType = &v
}

// GetSignatureAlgorithm returns the SignatureAlgorithm field value if set, zero value otherwise.
func (o *X509CertificateDetails) GetSignatureAlgorithm() string {
	if o == nil || IsNil(o.SignatureAlgorithm) {
		var ret string
		return ret
	}
	return *o.SignatureAlgorithm
}

// GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509CertificateDetails) GetSignatureAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureAlgorithm) {
		return nil, false
	}
	return o.SignatureAlgorithm, true
}

// HasSignatureAlgorithm returns a boolean if a field has been set.
func (o *X509CertificateDetails) HasSignatureAlgorithm() bool {
	if o != nil && !IsNil(o.SignatureAlgorithm) {
		return true
	}

	return false
}

// SetSignatureAlgorithm gets a reference to the given string and assigns it to the SignatureAlgorithm field.
func (o *X509CertificateDetails) SetSignatureAlgorithm(v string) {
	o.SignatureAlgorithm = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *X509CertificateDetails) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509CertificateDetails) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *X509CertificateDetails) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *X509CertificateDetails) SetStartTime(v string) {
	o.StartTime = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *X509CertificateDetails) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509CertificateDetails) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *X509CertificateDetails) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *X509CertificateDetails) SetSubject(v string) {
	o.Subject = &v
}

func (o X509CertificateDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o X509CertificateDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpiryTime) {
		toSerialize["expiryTime"] = o.ExpiryTime
	}
	if !IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	if !IsNil(o.PublicKeyType) {
		toSerialize["publicKeyType"] = o.PublicKeyType
	}
	if !IsNil(o.SignatureAlgorithm) {
		toSerialize["signatureAlgorithm"] = o.SignatureAlgorithm
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	return toSerialize, nil
}

type NullableX509CertificateDetails struct {
	value *X509CertificateDetails
	isSet bool
}

func (v NullableX509CertificateDetails) Get() *X509CertificateDetails {
	return v.value
}

func (v *NullableX509CertificateDetails) Set(val *X509CertificateDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableX509CertificateDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableX509CertificateDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableX509CertificateDetails(val *X509CertificateDetails) *NullableX509CertificateDetails {
	return &NullableX509CertificateDetails{value: val, isSet: true}
}

func (v NullableX509CertificateDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableX509CertificateDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


