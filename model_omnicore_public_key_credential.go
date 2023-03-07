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

// checks if the OmnicorePublicKeyCredential type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OmnicorePublicKeyCredential{}

// OmnicorePublicKeyCredential struct for OmnicorePublicKeyCredential
type OmnicorePublicKeyCredential struct {
	// Format: The format of the key.  Possible values:   \"UNSPECIFIED_PUBLIC_KEY_FORMAT\" - The format has not been specified. This is an invalid default value and must not be used.   \"RSA_PEM\" - An RSA public key encoded in base64, and wrapped by `-----BEGIN PUBLIC KEY-----` and `-----END PUBLIC KEY-----`. This can be used to verify `RS256` signatures in JWT tokens ([RFC7518]( https://www.ietf.org/rfc/rfc7518.txt)).   \"RSA_X509_PEM\" - As RSA_PEM, but wrapped in an X.509v3 certificate ([RFC5280]( https://www.ietf.org/rfc/rfc5280.txt)), encoded in base64, and wrapped by `-----BEGIN CERTIFICATE-----` and `-----END CERTIFICATE-----`.   \"ES256_PEM\" - Public key for the ECDSA algorithm using P-256 and SHA-256, encoded in base64, and wrapped by `-----BEGIN PUBLIC KEY-----` and `-----END PUBLIC KEY-----`. This can be used to verify JWT tokens with the `ES256` algorithm ([RFC7518](https://www.ietf.org/rfc/rfc7518.txt)). This curve is defined in [OpenSSL](https://www.openssl.org/) as the `prime256v1` curve.   \"ES256_X509_PEM\" - As ES256_PEM, but wrapped in an X.509v3 certificate ([RFC5280]( https://www.ietf.org/rfc/rfc5280.txt)), encoded in base64, and wrapped by `-----BEGIN CERTIFICATE-----` and `-----END CERTIFICATE-----`.
	Format string `json:"format"`
	// Key: The key data.
	Key *string `json:"key,omitempty"`
}

// NewOmnicorePublicKeyCredential instantiates a new OmnicorePublicKeyCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOmnicorePublicKeyCredential(format string) *OmnicorePublicKeyCredential {
	this := OmnicorePublicKeyCredential{}
	this.Format = format
	return &this
}

// NewOmnicorePublicKeyCredentialWithDefaults instantiates a new OmnicorePublicKeyCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOmnicorePublicKeyCredentialWithDefaults() *OmnicorePublicKeyCredential {
	this := OmnicorePublicKeyCredential{}
	return &this
}

// GetFormat returns the Format field value
func (o *OmnicorePublicKeyCredential) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *OmnicorePublicKeyCredential) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *OmnicorePublicKeyCredential) SetFormat(v string) {
	o.Format = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *OmnicorePublicKeyCredential) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmnicorePublicKeyCredential) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *OmnicorePublicKeyCredential) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *OmnicorePublicKeyCredential) SetKey(v string) {
	o.Key = &v
}

func (o OmnicorePublicKeyCredential) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OmnicorePublicKeyCredential) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["format"] = o.Format
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	return toSerialize, nil
}

type NullableOmnicorePublicKeyCredential struct {
	value *OmnicorePublicKeyCredential
	isSet bool
}

func (v NullableOmnicorePublicKeyCredential) Get() *OmnicorePublicKeyCredential {
	return v.value
}

func (v *NullableOmnicorePublicKeyCredential) Set(val *OmnicorePublicKeyCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableOmnicorePublicKeyCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableOmnicorePublicKeyCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOmnicorePublicKeyCredential(val *OmnicorePublicKeyCredential) *NullableOmnicorePublicKeyCredential {
	return &NullableOmnicorePublicKeyCredential{value: val, isSet: true}
}

func (v NullableOmnicorePublicKeyCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOmnicorePublicKeyCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


