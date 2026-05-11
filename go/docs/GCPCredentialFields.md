# GCPCredentialFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CrossProject** | Pointer to **string** |  | [optional] 
**AuthProviderX509CertUrl** | Pointer to **string** |  | [optional] 
**AuthUri** | Pointer to **string** |  | [optional] 
**ClientEmail** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientX509CertUrl** | Pointer to **string** |  | [optional] 
**PrivateKey** | Pointer to **string** | base64 encoded private key, beginning with -----BEGIN PRIVATE KEY----- | [optional] 
**PrivateKeyId** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**TokenUri** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UniverseDomain** | Pointer to **string** |  | [optional] 

## Methods

### NewGCPCredentialFields

`func NewGCPCredentialFields() *GCPCredentialFields`

NewGCPCredentialFields instantiates a new GCPCredentialFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCPCredentialFieldsWithDefaults

`func NewGCPCredentialFieldsWithDefaults() *GCPCredentialFields`

NewGCPCredentialFieldsWithDefaults instantiates a new GCPCredentialFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrossProject

`func (o *GCPCredentialFields) GetCrossProject() string`

GetCrossProject returns the CrossProject field if non-nil, zero value otherwise.

### GetCrossProjectOk

`func (o *GCPCredentialFields) GetCrossProjectOk() (*string, bool)`

GetCrossProjectOk returns a tuple with the CrossProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossProject

`func (o *GCPCredentialFields) SetCrossProject(v string)`

SetCrossProject sets CrossProject field to given value.

### HasCrossProject

`func (o *GCPCredentialFields) HasCrossProject() bool`

HasCrossProject returns a boolean if a field has been set.

### GetAuthProviderX509CertUrl

`func (o *GCPCredentialFields) GetAuthProviderX509CertUrl() string`

GetAuthProviderX509CertUrl returns the AuthProviderX509CertUrl field if non-nil, zero value otherwise.

### GetAuthProviderX509CertUrlOk

`func (o *GCPCredentialFields) GetAuthProviderX509CertUrlOk() (*string, bool)`

GetAuthProviderX509CertUrlOk returns a tuple with the AuthProviderX509CertUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProviderX509CertUrl

`func (o *GCPCredentialFields) SetAuthProviderX509CertUrl(v string)`

SetAuthProviderX509CertUrl sets AuthProviderX509CertUrl field to given value.

### HasAuthProviderX509CertUrl

`func (o *GCPCredentialFields) HasAuthProviderX509CertUrl() bool`

HasAuthProviderX509CertUrl returns a boolean if a field has been set.

### GetAuthUri

`func (o *GCPCredentialFields) GetAuthUri() string`

GetAuthUri returns the AuthUri field if non-nil, zero value otherwise.

### GetAuthUriOk

`func (o *GCPCredentialFields) GetAuthUriOk() (*string, bool)`

GetAuthUriOk returns a tuple with the AuthUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUri

`func (o *GCPCredentialFields) SetAuthUri(v string)`

SetAuthUri sets AuthUri field to given value.

### HasAuthUri

`func (o *GCPCredentialFields) HasAuthUri() bool`

HasAuthUri returns a boolean if a field has been set.

### GetClientEmail

`func (o *GCPCredentialFields) GetClientEmail() string`

GetClientEmail returns the ClientEmail field if non-nil, zero value otherwise.

### GetClientEmailOk

`func (o *GCPCredentialFields) GetClientEmailOk() (*string, bool)`

GetClientEmailOk returns a tuple with the ClientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEmail

`func (o *GCPCredentialFields) SetClientEmail(v string)`

SetClientEmail sets ClientEmail field to given value.

### HasClientEmail

`func (o *GCPCredentialFields) HasClientEmail() bool`

HasClientEmail returns a boolean if a field has been set.

### GetClientId

`func (o *GCPCredentialFields) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GCPCredentialFields) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GCPCredentialFields) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *GCPCredentialFields) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientX509CertUrl

`func (o *GCPCredentialFields) GetClientX509CertUrl() string`

GetClientX509CertUrl returns the ClientX509CertUrl field if non-nil, zero value otherwise.

### GetClientX509CertUrlOk

`func (o *GCPCredentialFields) GetClientX509CertUrlOk() (*string, bool)`

GetClientX509CertUrlOk returns a tuple with the ClientX509CertUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientX509CertUrl

`func (o *GCPCredentialFields) SetClientX509CertUrl(v string)`

SetClientX509CertUrl sets ClientX509CertUrl field to given value.

### HasClientX509CertUrl

`func (o *GCPCredentialFields) HasClientX509CertUrl() bool`

HasClientX509CertUrl returns a boolean if a field has been set.

### GetPrivateKey

`func (o *GCPCredentialFields) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *GCPCredentialFields) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *GCPCredentialFields) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *GCPCredentialFields) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPrivateKeyId

`func (o *GCPCredentialFields) GetPrivateKeyId() string`

GetPrivateKeyId returns the PrivateKeyId field if non-nil, zero value otherwise.

### GetPrivateKeyIdOk

`func (o *GCPCredentialFields) GetPrivateKeyIdOk() (*string, bool)`

GetPrivateKeyIdOk returns a tuple with the PrivateKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyId

`func (o *GCPCredentialFields) SetPrivateKeyId(v string)`

SetPrivateKeyId sets PrivateKeyId field to given value.

### HasPrivateKeyId

`func (o *GCPCredentialFields) HasPrivateKeyId() bool`

HasPrivateKeyId returns a boolean if a field has been set.

### GetProjectId

`func (o *GCPCredentialFields) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GCPCredentialFields) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GCPCredentialFields) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *GCPCredentialFields) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetTokenUri

`func (o *GCPCredentialFields) GetTokenUri() string`

GetTokenUri returns the TokenUri field if non-nil, zero value otherwise.

### GetTokenUriOk

`func (o *GCPCredentialFields) GetTokenUriOk() (*string, bool)`

GetTokenUriOk returns a tuple with the TokenUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUri

`func (o *GCPCredentialFields) SetTokenUri(v string)`

SetTokenUri sets TokenUri field to given value.

### HasTokenUri

`func (o *GCPCredentialFields) HasTokenUri() bool`

HasTokenUri returns a boolean if a field has been set.

### GetType

`func (o *GCPCredentialFields) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GCPCredentialFields) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GCPCredentialFields) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GCPCredentialFields) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUniverseDomain

`func (o *GCPCredentialFields) GetUniverseDomain() string`

GetUniverseDomain returns the UniverseDomain field if non-nil, zero value otherwise.

### GetUniverseDomainOk

`func (o *GCPCredentialFields) GetUniverseDomainOk() (*string, bool)`

GetUniverseDomainOk returns a tuple with the UniverseDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseDomain

`func (o *GCPCredentialFields) SetUniverseDomain(v string)`

SetUniverseDomain sets UniverseDomain field to given value.

### HasUniverseDomain

`func (o *GCPCredentialFields) HasUniverseDomain() bool`

HasUniverseDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


