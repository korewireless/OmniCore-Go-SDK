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

// checks if the PublicKeyCertificate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicKeyCertificate{}

// PublicKeyCertificate struct for PublicKeyCertificate
type PublicKeyCertificate struct {
	// Certificate: The certificate data.
	Certificate *string `json:"certificate,omitempty"`
	// Format: The certificate format.  Possible values:   \"UNSPECIFIED_PUBLIC_KEY_CERTIFICATE_FORMAT\" - The format has not been specified. This is an invalid default value and must not be used.   \"X509_CERTIFICATE_PEM\" - An X.509v3 certificate ([RFC5280](https://www.ietf.org/rfc/rfc5280.txt)), encoded in base64, and wrapped by `-----BEGIN CERTIFICATE-----` and `-----END CERTIFICATE-----`.
	Format *string `json:"format,omitempty"`
	X509Details *X509CertificateDetails `json:"x509Details,omitempty"`
}

// NewPublicKeyCertificate instantiates a new PublicKeyCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicKeyCertificate() *PublicKeyCertificate {
	this := PublicKeyCertificate{}
	return &this
}

// NewPublicKeyCertificateWithDefaults instantiates a new PublicKeyCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicKeyCertificateWithDefaults() *PublicKeyCertificate {
	this := PublicKeyCertificate{}
	return &this
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *PublicKeyCertificate) GetCertificate() string {
	if o == nil || IsNil(o.Certificate) {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicKeyCertificate) GetCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.Certificate) {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *PublicKeyCertificate) HasCertificate() bool {
	if o != nil && !IsNil(o.Certificate) {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *PublicKeyCertificate) SetCertificate(v string) {
	o.Certificate = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *PublicKeyCertificate) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicKeyCertificate) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *PublicKeyCertificate) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *PublicKeyCertificate) SetFormat(v string) {
	o.Format = &v
}

// GetX509Details returns the X509Details field value if set, zero value otherwise.
func (o *PublicKeyCertificate) GetX509Details() X509CertificateDetails {
	if o == nil || IsNil(o.X509Details) {
		var ret X509CertificateDetails
		return ret
	}
	return *o.X509Details
}

// GetX509DetailsOk returns a tuple with the X509Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicKeyCertificate) GetX509DetailsOk() (*X509CertificateDetails, bool) {
	if o == nil || IsNil(o.X509Details) {
		return nil, false
	}
	return o.X509Details, true
}

// HasX509Details returns a boolean if a field has been set.
func (o *PublicKeyCertificate) HasX509Details() bool {
	if o != nil && !IsNil(o.X509Details) {
		return true
	}

	return false
}

// SetX509Details gets a reference to the given X509CertificateDetails and assigns it to the X509Details field.
func (o *PublicKeyCertificate) SetX509Details(v X509CertificateDetails) {
	o.X509Details = &v
}

func (o PublicKeyCertificate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicKeyCertificate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Certificate) {
		toSerialize["certificate"] = o.Certificate
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.X509Details) {
		toSerialize["x509Details"] = o.X509Details
	}
	return toSerialize, nil
}

type NullablePublicKeyCertificate struct {
	value *PublicKeyCertificate
	isSet bool
}

func (v NullablePublicKeyCertificate) Get() *PublicKeyCertificate {
	return v.value
}

func (v *NullablePublicKeyCertificate) Set(val *PublicKeyCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicKeyCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicKeyCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicKeyCertificate(val *PublicKeyCertificate) *NullablePublicKeyCertificate {
	return &NullablePublicKeyCertificate{value: val, isSet: true}
}

func (v NullablePublicKeyCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicKeyCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


