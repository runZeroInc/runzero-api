# CredentialFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | **string** |  | 
**SecretAccessKey** | **string** |  | 
**Regions** | Pointer to **string** |  | [optional] 
**UseCrossAccountOrg** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**ClientId** | **string** |  | 
**ClientSecret** | **string** |  | 
**TenantId** | **string** |  | 
**Environment** | **string** |  | 
**MultiSubscription** | Pointer to **string** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**Username** | **string** |  | 
**Password** | **string** |  | 
**ApiUrl** | **string** |  | 
**AccessSecret** | **string** |  | 
**CrossProject** | Pointer to **string** |  | [optional] 
**AuthProviderX509CertUrl** | Pointer to **string** |  | [optional] 
**AuthUri** | Pointer to **string** |  | [optional] 
**ClientEmail** | Pointer to **string** |  | [optional] 
**ClientX509CertUrl** | Pointer to **string** |  | [optional] 
**PrivateKey** | **string** |  | 
**PrivateKeyId** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**TokenUri** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UniverseDomain** | Pointer to **string** |  | [optional] 
**Delegate** | **string** |  | 
**CustomerId** | Pointer to **string** |  | [optional] 
**Insecure** | Pointer to **string** |  | [optional] 
**Thumbprints** | Pointer to **string** |  | [optional] 
**Url** | **string** |  | 
**BaseDn** | **string** |  | 
**LegacyTls** | Pointer to **string** |  | [optional] 
**DbConnectionString** | **string** |  | 
**ApiKey** | **string** |  | 
**Hostname** | **string** |  | 
**SecretKey** | **string** |  | 
**Community** | **string** |  | 
**Context** | Pointer to **string** |  | [optional] 
**AuthProtocol** | Pointer to **string** |  | [optional] 
**AuthPassphrase** | Pointer to **string** |  | [optional] 
**PrivacyProtocol** | Pointer to **string** |  | [optional] 
**PrivacyPassphrase** | Pointer to **string** |  | [optional] 
**ApiToken** | **string** |  | 
**AuthUrl** | **string** |  | 
**X509Certificate** | **string** |  | 

## Methods

### NewCredentialFields

`func NewCredentialFields(accessKey string, secretAccessKey string, clientId string, clientSecret string, tenantId string, environment string, username string, password string, apiUrl string, accessSecret string, privateKey string, delegate string, url string, baseDn string, dbConnectionString string, apiKey string, hostname string, secretKey string, community string, apiToken string, authUrl string, x509Certificate string, ) *CredentialFields`

NewCredentialFields instantiates a new CredentialFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialFieldsWithDefaults

`func NewCredentialFieldsWithDefaults() *CredentialFields`

NewCredentialFieldsWithDefaults instantiates a new CredentialFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *CredentialFields) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *CredentialFields) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *CredentialFields) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetSecretAccessKey

`func (o *CredentialFields) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *CredentialFields) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *CredentialFields) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.


### GetRegions

`func (o *CredentialFields) GetRegions() string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *CredentialFields) GetRegionsOk() (*string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *CredentialFields) SetRegions(v string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *CredentialFields) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetUseCrossAccountOrg

`func (o *CredentialFields) GetUseCrossAccountOrg() string`

GetUseCrossAccountOrg returns the UseCrossAccountOrg field if non-nil, zero value otherwise.

### GetUseCrossAccountOrgOk

`func (o *CredentialFields) GetUseCrossAccountOrgOk() (*string, bool)`

GetUseCrossAccountOrgOk returns a tuple with the UseCrossAccountOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCrossAccountOrg

`func (o *CredentialFields) SetUseCrossAccountOrg(v string)`

SetUseCrossAccountOrg sets UseCrossAccountOrg field to given value.

### HasUseCrossAccountOrg

`func (o *CredentialFields) HasUseCrossAccountOrg() bool`

HasUseCrossAccountOrg returns a boolean if a field has been set.

### GetRole

`func (o *CredentialFields) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CredentialFields) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CredentialFields) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *CredentialFields) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetClientId

`func (o *CredentialFields) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CredentialFields) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CredentialFields) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *CredentialFields) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *CredentialFields) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *CredentialFields) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetTenantId

`func (o *CredentialFields) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CredentialFields) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CredentialFields) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetEnvironment

`func (o *CredentialFields) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CredentialFields) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CredentialFields) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetMultiSubscription

`func (o *CredentialFields) GetMultiSubscription() string`

GetMultiSubscription returns the MultiSubscription field if non-nil, zero value otherwise.

### GetMultiSubscriptionOk

`func (o *CredentialFields) GetMultiSubscriptionOk() (*string, bool)`

GetMultiSubscriptionOk returns a tuple with the MultiSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiSubscription

`func (o *CredentialFields) SetMultiSubscription(v string)`

SetMultiSubscription sets MultiSubscription field to given value.

### HasMultiSubscription

`func (o *CredentialFields) HasMultiSubscription() bool`

HasMultiSubscription returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *CredentialFields) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CredentialFields) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CredentialFields) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *CredentialFields) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetUsername

`func (o *CredentialFields) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CredentialFields) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CredentialFields) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *CredentialFields) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CredentialFields) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CredentialFields) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetApiUrl

`func (o *CredentialFields) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *CredentialFields) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *CredentialFields) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.


### GetAccessSecret

`func (o *CredentialFields) GetAccessSecret() string`

GetAccessSecret returns the AccessSecret field if non-nil, zero value otherwise.

### GetAccessSecretOk

`func (o *CredentialFields) GetAccessSecretOk() (*string, bool)`

GetAccessSecretOk returns a tuple with the AccessSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessSecret

`func (o *CredentialFields) SetAccessSecret(v string)`

SetAccessSecret sets AccessSecret field to given value.


### GetCrossProject

`func (o *CredentialFields) GetCrossProject() string`

GetCrossProject returns the CrossProject field if non-nil, zero value otherwise.

### GetCrossProjectOk

`func (o *CredentialFields) GetCrossProjectOk() (*string, bool)`

GetCrossProjectOk returns a tuple with the CrossProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossProject

`func (o *CredentialFields) SetCrossProject(v string)`

SetCrossProject sets CrossProject field to given value.

### HasCrossProject

`func (o *CredentialFields) HasCrossProject() bool`

HasCrossProject returns a boolean if a field has been set.

### GetAuthProviderX509CertUrl

`func (o *CredentialFields) GetAuthProviderX509CertUrl() string`

GetAuthProviderX509CertUrl returns the AuthProviderX509CertUrl field if non-nil, zero value otherwise.

### GetAuthProviderX509CertUrlOk

`func (o *CredentialFields) GetAuthProviderX509CertUrlOk() (*string, bool)`

GetAuthProviderX509CertUrlOk returns a tuple with the AuthProviderX509CertUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProviderX509CertUrl

`func (o *CredentialFields) SetAuthProviderX509CertUrl(v string)`

SetAuthProviderX509CertUrl sets AuthProviderX509CertUrl field to given value.

### HasAuthProviderX509CertUrl

`func (o *CredentialFields) HasAuthProviderX509CertUrl() bool`

HasAuthProviderX509CertUrl returns a boolean if a field has been set.

### GetAuthUri

`func (o *CredentialFields) GetAuthUri() string`

GetAuthUri returns the AuthUri field if non-nil, zero value otherwise.

### GetAuthUriOk

`func (o *CredentialFields) GetAuthUriOk() (*string, bool)`

GetAuthUriOk returns a tuple with the AuthUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUri

`func (o *CredentialFields) SetAuthUri(v string)`

SetAuthUri sets AuthUri field to given value.

### HasAuthUri

`func (o *CredentialFields) HasAuthUri() bool`

HasAuthUri returns a boolean if a field has been set.

### GetClientEmail

`func (o *CredentialFields) GetClientEmail() string`

GetClientEmail returns the ClientEmail field if non-nil, zero value otherwise.

### GetClientEmailOk

`func (o *CredentialFields) GetClientEmailOk() (*string, bool)`

GetClientEmailOk returns a tuple with the ClientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEmail

`func (o *CredentialFields) SetClientEmail(v string)`

SetClientEmail sets ClientEmail field to given value.

### HasClientEmail

`func (o *CredentialFields) HasClientEmail() bool`

HasClientEmail returns a boolean if a field has been set.

### GetClientX509CertUrl

`func (o *CredentialFields) GetClientX509CertUrl() string`

GetClientX509CertUrl returns the ClientX509CertUrl field if non-nil, zero value otherwise.

### GetClientX509CertUrlOk

`func (o *CredentialFields) GetClientX509CertUrlOk() (*string, bool)`

GetClientX509CertUrlOk returns a tuple with the ClientX509CertUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientX509CertUrl

`func (o *CredentialFields) SetClientX509CertUrl(v string)`

SetClientX509CertUrl sets ClientX509CertUrl field to given value.

### HasClientX509CertUrl

`func (o *CredentialFields) HasClientX509CertUrl() bool`

HasClientX509CertUrl returns a boolean if a field has been set.

### GetPrivateKey

`func (o *CredentialFields) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CredentialFields) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CredentialFields) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetPrivateKeyId

`func (o *CredentialFields) GetPrivateKeyId() string`

GetPrivateKeyId returns the PrivateKeyId field if non-nil, zero value otherwise.

### GetPrivateKeyIdOk

`func (o *CredentialFields) GetPrivateKeyIdOk() (*string, bool)`

GetPrivateKeyIdOk returns a tuple with the PrivateKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyId

`func (o *CredentialFields) SetPrivateKeyId(v string)`

SetPrivateKeyId sets PrivateKeyId field to given value.

### HasPrivateKeyId

`func (o *CredentialFields) HasPrivateKeyId() bool`

HasPrivateKeyId returns a boolean if a field has been set.

### GetProjectId

`func (o *CredentialFields) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CredentialFields) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CredentialFields) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CredentialFields) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetTokenUri

`func (o *CredentialFields) GetTokenUri() string`

GetTokenUri returns the TokenUri field if non-nil, zero value otherwise.

### GetTokenUriOk

`func (o *CredentialFields) GetTokenUriOk() (*string, bool)`

GetTokenUriOk returns a tuple with the TokenUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUri

`func (o *CredentialFields) SetTokenUri(v string)`

SetTokenUri sets TokenUri field to given value.

### HasTokenUri

`func (o *CredentialFields) HasTokenUri() bool`

HasTokenUri returns a boolean if a field has been set.

### GetType

`func (o *CredentialFields) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CredentialFields) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CredentialFields) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CredentialFields) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUniverseDomain

`func (o *CredentialFields) GetUniverseDomain() string`

GetUniverseDomain returns the UniverseDomain field if non-nil, zero value otherwise.

### GetUniverseDomainOk

`func (o *CredentialFields) GetUniverseDomainOk() (*string, bool)`

GetUniverseDomainOk returns a tuple with the UniverseDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseDomain

`func (o *CredentialFields) SetUniverseDomain(v string)`

SetUniverseDomain sets UniverseDomain field to given value.

### HasUniverseDomain

`func (o *CredentialFields) HasUniverseDomain() bool`

HasUniverseDomain returns a boolean if a field has been set.

### GetDelegate

`func (o *CredentialFields) GetDelegate() string`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *CredentialFields) GetDelegateOk() (*string, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *CredentialFields) SetDelegate(v string)`

SetDelegate sets Delegate field to given value.


### GetCustomerId

`func (o *CredentialFields) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CredentialFields) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CredentialFields) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *CredentialFields) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetInsecure

`func (o *CredentialFields) GetInsecure() string`

GetInsecure returns the Insecure field if non-nil, zero value otherwise.

### GetInsecureOk

`func (o *CredentialFields) GetInsecureOk() (*string, bool)`

GetInsecureOk returns a tuple with the Insecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecure

`func (o *CredentialFields) SetInsecure(v string)`

SetInsecure sets Insecure field to given value.

### HasInsecure

`func (o *CredentialFields) HasInsecure() bool`

HasInsecure returns a boolean if a field has been set.

### GetThumbprints

`func (o *CredentialFields) GetThumbprints() string`

GetThumbprints returns the Thumbprints field if non-nil, zero value otherwise.

### GetThumbprintsOk

`func (o *CredentialFields) GetThumbprintsOk() (*string, bool)`

GetThumbprintsOk returns a tuple with the Thumbprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprints

`func (o *CredentialFields) SetThumbprints(v string)`

SetThumbprints sets Thumbprints field to given value.

### HasThumbprints

`func (o *CredentialFields) HasThumbprints() bool`

HasThumbprints returns a boolean if a field has been set.

### GetUrl

`func (o *CredentialFields) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CredentialFields) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CredentialFields) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetBaseDn

`func (o *CredentialFields) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *CredentialFields) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *CredentialFields) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.


### GetLegacyTls

`func (o *CredentialFields) GetLegacyTls() string`

GetLegacyTls returns the LegacyTls field if non-nil, zero value otherwise.

### GetLegacyTlsOk

`func (o *CredentialFields) GetLegacyTlsOk() (*string, bool)`

GetLegacyTlsOk returns a tuple with the LegacyTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyTls

`func (o *CredentialFields) SetLegacyTls(v string)`

SetLegacyTls sets LegacyTls field to given value.

### HasLegacyTls

`func (o *CredentialFields) HasLegacyTls() bool`

HasLegacyTls returns a boolean if a field has been set.

### GetDbConnectionString

`func (o *CredentialFields) GetDbConnectionString() string`

GetDbConnectionString returns the DbConnectionString field if non-nil, zero value otherwise.

### GetDbConnectionStringOk

`func (o *CredentialFields) GetDbConnectionStringOk() (*string, bool)`

GetDbConnectionStringOk returns a tuple with the DbConnectionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbConnectionString

`func (o *CredentialFields) SetDbConnectionString(v string)`

SetDbConnectionString sets DbConnectionString field to given value.


### GetApiKey

`func (o *CredentialFields) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *CredentialFields) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *CredentialFields) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetHostname

`func (o *CredentialFields) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *CredentialFields) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *CredentialFields) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetSecretKey

`func (o *CredentialFields) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *CredentialFields) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *CredentialFields) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### GetCommunity

`func (o *CredentialFields) GetCommunity() string`

GetCommunity returns the Community field if non-nil, zero value otherwise.

### GetCommunityOk

`func (o *CredentialFields) GetCommunityOk() (*string, bool)`

GetCommunityOk returns a tuple with the Community field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunity

`func (o *CredentialFields) SetCommunity(v string)`

SetCommunity sets Community field to given value.


### GetContext

`func (o *CredentialFields) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *CredentialFields) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *CredentialFields) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *CredentialFields) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetAuthProtocol

`func (o *CredentialFields) GetAuthProtocol() string`

GetAuthProtocol returns the AuthProtocol field if non-nil, zero value otherwise.

### GetAuthProtocolOk

`func (o *CredentialFields) GetAuthProtocolOk() (*string, bool)`

GetAuthProtocolOk returns a tuple with the AuthProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProtocol

`func (o *CredentialFields) SetAuthProtocol(v string)`

SetAuthProtocol sets AuthProtocol field to given value.

### HasAuthProtocol

`func (o *CredentialFields) HasAuthProtocol() bool`

HasAuthProtocol returns a boolean if a field has been set.

### GetAuthPassphrase

`func (o *CredentialFields) GetAuthPassphrase() string`

GetAuthPassphrase returns the AuthPassphrase field if non-nil, zero value otherwise.

### GetAuthPassphraseOk

`func (o *CredentialFields) GetAuthPassphraseOk() (*string, bool)`

GetAuthPassphraseOk returns a tuple with the AuthPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPassphrase

`func (o *CredentialFields) SetAuthPassphrase(v string)`

SetAuthPassphrase sets AuthPassphrase field to given value.

### HasAuthPassphrase

`func (o *CredentialFields) HasAuthPassphrase() bool`

HasAuthPassphrase returns a boolean if a field has been set.

### GetPrivacyProtocol

`func (o *CredentialFields) GetPrivacyProtocol() string`

GetPrivacyProtocol returns the PrivacyProtocol field if non-nil, zero value otherwise.

### GetPrivacyProtocolOk

`func (o *CredentialFields) GetPrivacyProtocolOk() (*string, bool)`

GetPrivacyProtocolOk returns a tuple with the PrivacyProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyProtocol

`func (o *CredentialFields) SetPrivacyProtocol(v string)`

SetPrivacyProtocol sets PrivacyProtocol field to given value.

### HasPrivacyProtocol

`func (o *CredentialFields) HasPrivacyProtocol() bool`

HasPrivacyProtocol returns a boolean if a field has been set.

### GetPrivacyPassphrase

`func (o *CredentialFields) GetPrivacyPassphrase() string`

GetPrivacyPassphrase returns the PrivacyPassphrase field if non-nil, zero value otherwise.

### GetPrivacyPassphraseOk

`func (o *CredentialFields) GetPrivacyPassphraseOk() (*string, bool)`

GetPrivacyPassphraseOk returns a tuple with the PrivacyPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPassphrase

`func (o *CredentialFields) SetPrivacyPassphrase(v string)`

SetPrivacyPassphrase sets PrivacyPassphrase field to given value.

### HasPrivacyPassphrase

`func (o *CredentialFields) HasPrivacyPassphrase() bool`

HasPrivacyPassphrase returns a boolean if a field has been set.

### GetApiToken

`func (o *CredentialFields) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *CredentialFields) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *CredentialFields) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.


### GetAuthUrl

`func (o *CredentialFields) GetAuthUrl() string`

GetAuthUrl returns the AuthUrl field if non-nil, zero value otherwise.

### GetAuthUrlOk

`func (o *CredentialFields) GetAuthUrlOk() (*string, bool)`

GetAuthUrlOk returns a tuple with the AuthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUrl

`func (o *CredentialFields) SetAuthUrl(v string)`

SetAuthUrl sets AuthUrl field to given value.


### GetX509Certificate

`func (o *CredentialFields) GetX509Certificate() string`

GetX509Certificate returns the X509Certificate field if non-nil, zero value otherwise.

### GetX509CertificateOk

`func (o *CredentialFields) GetX509CertificateOk() (*string, bool)`

GetX509CertificateOk returns a tuple with the X509Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509Certificate

`func (o *CredentialFields) SetX509Certificate(v string)`

SetX509Certificate sets X509Certificate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


