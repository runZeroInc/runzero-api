# X509CredentialFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**X509Certificate** | **string** |  | 
**PrivateKey** | **string** |  | 

## Methods

### NewX509CredentialFields

`func NewX509CredentialFields(x509Certificate string, privateKey string, ) *X509CredentialFields`

NewX509CredentialFields instantiates a new X509CredentialFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewX509CredentialFieldsWithDefaults

`func NewX509CredentialFieldsWithDefaults() *X509CredentialFields`

NewX509CredentialFieldsWithDefaults instantiates a new X509CredentialFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetX509Certificate

`func (o *X509CredentialFields) GetX509Certificate() string`

GetX509Certificate returns the X509Certificate field if non-nil, zero value otherwise.

### GetX509CertificateOk

`func (o *X509CredentialFields) GetX509CertificateOk() (*string, bool)`

GetX509CertificateOk returns a tuple with the X509Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509Certificate

`func (o *X509CredentialFields) SetX509Certificate(v string)`

SetX509Certificate sets X509Certificate field to given value.


### GetPrivateKey

`func (o *X509CredentialFields) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *X509CredentialFields) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *X509CredentialFields) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


