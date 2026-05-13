# GoogleWorkspaceCredentialFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delegate** | **string** |  | 
**CustomerId** | Pointer to **string** |  | [optional] 
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

### NewGoogleWorkspaceCredentialFields

`func NewGoogleWorkspaceCredentialFields(delegate string, ) *GoogleWorkspaceCredentialFields`

NewGoogleWorkspaceCredentialFields instantiates a new GoogleWorkspaceCredentialFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleWorkspaceCredentialFieldsWithDefaults

`func NewGoogleWorkspaceCredentialFieldsWithDefaults() *GoogleWorkspaceCredentialFields`

NewGoogleWorkspaceCredentialFieldsWithDefaults instantiates a new GoogleWorkspaceCredentialFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegate

`func (o *GoogleWorkspaceCredentialFields) GetDelegate() string`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *GoogleWorkspaceCredentialFields) GetDelegateOk() (*string, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *GoogleWorkspaceCredentialFields) SetDelegate(v string)`

SetDelegate sets Delegate field to given value.


### GetCustomerId

`func (o *GoogleWorkspaceCredentialFields) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *GoogleWorkspaceCredentialFields) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *GoogleWorkspaceCredentialFields) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *GoogleWorkspaceCredentialFields) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetAuthProviderX509CertUrl

`func (o *GoogleWorkspaceCredentialFields) GetAuthProviderX509CertUrl() string`

GetAuthProviderX509CertUrl returns the AuthProviderX509CertUrl field if non-nil, zero value otherwise.

### GetAuthProviderX509CertUrlOk

`func (o *GoogleWorkspaceCredentialFields) GetAuthProviderX509CertUrlOk() (*string, bool)`

GetAuthProviderX509CertUrlOk returns a tuple with the AuthProviderX509CertUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProviderX509CertUrl

`func (o *GoogleWorkspaceCredentialFields) SetAuthProviderX509CertUrl(v string)`

SetAuthProviderX509CertUrl sets AuthProviderX509CertUrl field to given value.

### HasAuthProviderX509CertUrl

`func (o *GoogleWorkspaceCredentialFields) HasAuthProviderX509CertUrl() bool`

HasAuthProviderX509CertUrl returns a boolean if a field has been set.

### GetAuthUri

`func (o *GoogleWorkspaceCredentialFields) GetAuthUri() string`

GetAuthUri returns the AuthUri field if non-nil, zero value otherwise.

### GetAuthUriOk

`func (o *GoogleWorkspaceCredentialFields) GetAuthUriOk() (*string, bool)`

GetAuthUriOk returns a tuple with the AuthUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUri

`func (o *GoogleWorkspaceCredentialFields) SetAuthUri(v string)`

SetAuthUri sets AuthUri field to given value.

### HasAuthUri

`func (o *GoogleWorkspaceCredentialFields) HasAuthUri() bool`

HasAuthUri returns a boolean if a field has been set.

### GetClientEmail

`func (o *GoogleWorkspaceCredentialFields) GetClientEmail() string`

GetClientEmail returns the ClientEmail field if non-nil, zero value otherwise.

### GetClientEmailOk

`func (o *GoogleWorkspaceCredentialFields) GetClientEmailOk() (*string, bool)`

GetClientEmailOk returns a tuple with the ClientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEmail

`func (o *GoogleWorkspaceCredentialFields) SetClientEmail(v string)`

SetClientEmail sets ClientEmail field to given value.

### HasClientEmail

`func (o *GoogleWorkspaceCredentialFields) HasClientEmail() bool`

HasClientEmail returns a boolean if a field has been set.

### GetClientId

`func (o *GoogleWorkspaceCredentialFields) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GoogleWorkspaceCredentialFields) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GoogleWorkspaceCredentialFields) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *GoogleWorkspaceCredentialFields) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientX509CertUrl

`func (o *GoogleWorkspaceCredentialFields) GetClientX509CertUrl() string`

GetClientX509CertUrl returns the ClientX509CertUrl field if non-nil, zero value otherwise.

### GetClientX509CertUrlOk

`func (o *GoogleWorkspaceCredentialFields) GetClientX509CertUrlOk() (*string, bool)`

GetClientX509CertUrlOk returns a tuple with the ClientX509CertUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientX509CertUrl

`func (o *GoogleWorkspaceCredentialFields) SetClientX509CertUrl(v string)`

SetClientX509CertUrl sets ClientX509CertUrl field to given value.

### HasClientX509CertUrl

`func (o *GoogleWorkspaceCredentialFields) HasClientX509CertUrl() bool`

HasClientX509CertUrl returns a boolean if a field has been set.

### GetPrivateKey

`func (o *GoogleWorkspaceCredentialFields) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *GoogleWorkspaceCredentialFields) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *GoogleWorkspaceCredentialFields) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *GoogleWorkspaceCredentialFields) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPrivateKeyId

`func (o *GoogleWorkspaceCredentialFields) GetPrivateKeyId() string`

GetPrivateKeyId returns the PrivateKeyId field if non-nil, zero value otherwise.

### GetPrivateKeyIdOk

`func (o *GoogleWorkspaceCredentialFields) GetPrivateKeyIdOk() (*string, bool)`

GetPrivateKeyIdOk returns a tuple with the PrivateKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyId

`func (o *GoogleWorkspaceCredentialFields) SetPrivateKeyId(v string)`

SetPrivateKeyId sets PrivateKeyId field to given value.

### HasPrivateKeyId

`func (o *GoogleWorkspaceCredentialFields) HasPrivateKeyId() bool`

HasPrivateKeyId returns a boolean if a field has been set.

### GetProjectId

`func (o *GoogleWorkspaceCredentialFields) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GoogleWorkspaceCredentialFields) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GoogleWorkspaceCredentialFields) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *GoogleWorkspaceCredentialFields) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetTokenUri

`func (o *GoogleWorkspaceCredentialFields) GetTokenUri() string`

GetTokenUri returns the TokenUri field if non-nil, zero value otherwise.

### GetTokenUriOk

`func (o *GoogleWorkspaceCredentialFields) GetTokenUriOk() (*string, bool)`

GetTokenUriOk returns a tuple with the TokenUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUri

`func (o *GoogleWorkspaceCredentialFields) SetTokenUri(v string)`

SetTokenUri sets TokenUri field to given value.

### HasTokenUri

`func (o *GoogleWorkspaceCredentialFields) HasTokenUri() bool`

HasTokenUri returns a boolean if a field has been set.

### GetType

`func (o *GoogleWorkspaceCredentialFields) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GoogleWorkspaceCredentialFields) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GoogleWorkspaceCredentialFields) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GoogleWorkspaceCredentialFields) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUniverseDomain

`func (o *GoogleWorkspaceCredentialFields) GetUniverseDomain() string`

GetUniverseDomain returns the UniverseDomain field if non-nil, zero value otherwise.

### GetUniverseDomainOk

`func (o *GoogleWorkspaceCredentialFields) GetUniverseDomainOk() (*string, bool)`

GetUniverseDomainOk returns a tuple with the UniverseDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseDomain

`func (o *GoogleWorkspaceCredentialFields) SetUniverseDomain(v string)`

SetUniverseDomain sets UniverseDomain field to given value.

### HasUniverseDomain

`func (o *GoogleWorkspaceCredentialFields) HasUniverseDomain() bool`

HasUniverseDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


