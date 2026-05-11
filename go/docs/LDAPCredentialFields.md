# LDAPCredentialFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Username** | **string** |  | 
**Password** | **string** |  | 
**BaseDn** | **string** |  | 
**Insecure** | Pointer to **string** |  | [optional] 
**LegacyTls** | Pointer to **string** |  | [optional] 
**Thumbprints** | Pointer to **string** |  | [optional] 

## Methods

### NewLDAPCredentialFields

`func NewLDAPCredentialFields(url string, username string, password string, baseDn string, ) *LDAPCredentialFields`

NewLDAPCredentialFields instantiates a new LDAPCredentialFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLDAPCredentialFieldsWithDefaults

`func NewLDAPCredentialFieldsWithDefaults() *LDAPCredentialFields`

NewLDAPCredentialFieldsWithDefaults instantiates a new LDAPCredentialFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *LDAPCredentialFields) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LDAPCredentialFields) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LDAPCredentialFields) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUsername

`func (o *LDAPCredentialFields) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *LDAPCredentialFields) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *LDAPCredentialFields) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *LDAPCredentialFields) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LDAPCredentialFields) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LDAPCredentialFields) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetBaseDn

`func (o *LDAPCredentialFields) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *LDAPCredentialFields) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *LDAPCredentialFields) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.


### GetInsecure

`func (o *LDAPCredentialFields) GetInsecure() string`

GetInsecure returns the Insecure field if non-nil, zero value otherwise.

### GetInsecureOk

`func (o *LDAPCredentialFields) GetInsecureOk() (*string, bool)`

GetInsecureOk returns a tuple with the Insecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecure

`func (o *LDAPCredentialFields) SetInsecure(v string)`

SetInsecure sets Insecure field to given value.

### HasInsecure

`func (o *LDAPCredentialFields) HasInsecure() bool`

HasInsecure returns a boolean if a field has been set.

### GetLegacyTls

`func (o *LDAPCredentialFields) GetLegacyTls() string`

GetLegacyTls returns the LegacyTls field if non-nil, zero value otherwise.

### GetLegacyTlsOk

`func (o *LDAPCredentialFields) GetLegacyTlsOk() (*string, bool)`

GetLegacyTlsOk returns a tuple with the LegacyTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyTls

`func (o *LDAPCredentialFields) SetLegacyTls(v string)`

SetLegacyTls sets LegacyTls field to given value.

### HasLegacyTls

`func (o *LDAPCredentialFields) HasLegacyTls() bool`

HasLegacyTls returns a boolean if a field has been set.

### GetThumbprints

`func (o *LDAPCredentialFields) GetThumbprints() string`

GetThumbprints returns the Thumbprints field if non-nil, zero value otherwise.

### GetThumbprintsOk

`func (o *LDAPCredentialFields) GetThumbprintsOk() (*string, bool)`

GetThumbprintsOk returns a tuple with the Thumbprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprints

`func (o *LDAPCredentialFields) SetThumbprints(v string)`

SetThumbprints sets Thumbprints field to given value.

### HasThumbprints

`func (o *LDAPCredentialFields) HasThumbprints() bool`

HasThumbprints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


