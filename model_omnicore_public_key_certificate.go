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

// checks if the OmnicorePublicKeyCertificate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OmnicorePublicKeyCertificate{}

// OmnicorePublicKeyCertificate struct for OmnicorePublicKeyCertificate
type OmnicorePublicKeyCertificate struct {
	// Certificate: The certificate data.
	Certificate *string `json:"certificate,omitempty"`
	// Format: The certificate format.  Possible values:   \"UNSPECIFIED_PUBLIC_KEY_CERTIFICATE_FORMAT\" - The format has not been specified. This is an invalid default value and must not be used.   \"X509_CERTIFICATE_PEM\" - An X.509v3 certificate ([RFC5280](https://www.ietf.org/rfc/rfc5280.txt)), encoded in base64, and wrapped by `-----BEGIN CERTIFICATE-----` and `-----END CERTIFICATE-----`.
	Format *string `json:"format,omitempty"`
	X509Details *OmnicoreX509CertificateDetails `json:"x509Details,omitempty"`
}

// NewOmnicorePublicKeyCertificate instantiates a new OmnicorePublicKeyCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOmnicorePublicKeyCertificate() *OmnicorePublicKeyCertificate {
	this := OmnicorePublicKeyCertificate{}
	return &this
}

// NewOmnicorePublicKeyCertificateWithDefaults instantiates a new OmnicorePublicKeyCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOmnicorePublicKeyCertificateWithDefaults() *OmnicorePublicKeyCertificate {
	this := OmnicorePublicKeyCertificate{}
	return &this
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *OmnicorePublicKeyCertificate) GetCertificate() string {
	if o == nil || IsNil(o.Certificate) {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicorePublicKeyCertificate) GetCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.Certificate) {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *OmnicorePublicKeyCertificate) HasCertificate() bool {
	if o != nil && !IsNil(o.Certificate) {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *OmnicorePublicKeyCertificate) SetCertificate(v string) {
	o.Certificate = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *OmnicorePublicKeyCertificate) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicorePublicKeyCertificate) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *OmnicorePublicKeyCertificate) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *OmnicorePublicKeyCertificate) SetFormat(v string) {
	o.Format = &v
}

// GetX509Details returns the X509Details field value if set, zero value otherwise.
func (o *OmnicorePublicKeyCertificate) GetX509Details() OmnicoreX509CertificateDetails {
	if o == nil || IsNil(o.X509Details) {
		var ret OmnicoreX509CertificateDetails
		return ret
	}
	return *o.X509Details
}

// GetX509DetailsOk returns a tuple with the X509Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicorePublicKeyCertificate) GetX509DetailsOk() (*OmnicoreX509CertificateDetails, bool) {
	if o == nil || IsNil(o.X509Details) {
		return nil, false
	}
	return o.X509Details, true
}

// HasX509Details returns a boolean if a field has been set.
func (o *OmnicorePublicKeyCertificate) HasX509Details() bool {
	if o != nil && !IsNil(o.X509Details) {
		return true
	}

	return false
}

// SetX509Details gets a reference to the given OmnicoreX509CertificateDetails and assigns it to the X509Details field.
func (o *OmnicorePublicKeyCertificate) SetX509Details(v OmnicoreX509CertificateDetails) {
	o.X509Details = &v
}

func (o OmnicorePublicKeyCertificate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OmnicorePublicKeyCertificate) ToMap() (map[string]interface{}, error) {
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

type NullableOmnicorePublicKeyCertificate struct {
	value *OmnicorePublicKeyCertificate
	isSet bool
}

func (v NullableOmnicorePublicKeyCertificate) Get() *OmnicorePublicKeyCertificate {
	return v.value
}

func (v *NullableOmnicorePublicKeyCertificate) Set(val *OmnicorePublicKeyCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableOmnicorePublicKeyCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableOmnicorePublicKeyCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOmnicorePublicKeyCertificate(val *OmnicorePublicKeyCertificate) *NullableOmnicorePublicKeyCertificate {
	return &NullableOmnicorePublicKeyCertificate{value: val, isSet: true}
}

func (v NullableOmnicorePublicKeyCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOmnicorePublicKeyCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


