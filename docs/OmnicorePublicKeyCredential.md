# OmnicorePublicKeyCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | **string** | Format: The format of the key.  Possible values:   \&quot;UNSPECIFIED_PUBLIC_KEY_FORMAT\&quot; - The format has not been specified. This is an invalid default value and must not be used.   \&quot;RSA_PEM\&quot; - An RSA public key encoded in base64, and wrapped by &#x60;-----BEGIN PUBLIC KEY-----&#x60; and &#x60;-----END PUBLIC KEY-----&#x60;. This can be used to verify &#x60;RS256&#x60; signatures in JWT tokens ([RFC7518]( https://www.ietf.org/rfc/rfc7518.txt)).   \&quot;RSA_X509_PEM\&quot; - As RSA_PEM, but wrapped in an X.509v3 certificate ([RFC5280]( https://www.ietf.org/rfc/rfc5280.txt)), encoded in base64, and wrapped by &#x60;-----BEGIN CERTIFICATE-----&#x60; and &#x60;-----END CERTIFICATE-----&#x60;.   \&quot;ES256_PEM\&quot; - Public key for the ECDSA algorithm using P-256 and SHA-256, encoded in base64, and wrapped by &#x60;-----BEGIN PUBLIC KEY-----&#x60; and &#x60;-----END PUBLIC KEY-----&#x60;. This can be used to verify JWT tokens with the &#x60;ES256&#x60; algorithm ([RFC7518](https://www.ietf.org/rfc/rfc7518.txt)). This curve is defined in [OpenSSL](https://www.openssl.org/) as the &#x60;prime256v1&#x60; curve.   \&quot;ES256_X509_PEM\&quot; - As ES256_PEM, but wrapped in an X.509v3 certificate ([RFC5280]( https://www.ietf.org/rfc/rfc5280.txt)), encoded in base64, and wrapped by &#x60;-----BEGIN CERTIFICATE-----&#x60; and &#x60;-----END CERTIFICATE-----&#x60;. | 
**Key** | Pointer to **string** | Key: The key data. | [optional] 

## Methods

### NewOmnicorePublicKeyCredential

`func NewOmnicorePublicKeyCredential(format string, ) *OmnicorePublicKeyCredential`

NewOmnicorePublicKeyCredential instantiates a new OmnicorePublicKeyCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOmnicorePublicKeyCredentialWithDefaults

`func NewOmnicorePublicKeyCredentialWithDefaults() *OmnicorePublicKeyCredential`

NewOmnicorePublicKeyCredentialWithDefaults instantiates a new OmnicorePublicKeyCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *OmnicorePublicKeyCredential) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *OmnicorePublicKeyCredential) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *OmnicorePublicKeyCredential) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetKey

`func (o *OmnicorePublicKeyCredential) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *OmnicorePublicKeyCredential) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *OmnicorePublicKeyCredential) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *OmnicorePublicKeyCredential) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


