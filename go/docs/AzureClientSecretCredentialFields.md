# AzureClientSecretCredentialFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** |  | 
**ClientSecret** | **string** |  | 
**TenantId** | **string** |  | 
**Environment** | **string** |  | 
**MultiSubscription** | Pointer to **string** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 

## Methods

### NewAzureClientSecretCredentialFields

`func NewAzureClientSecretCredentialFields(clientId string, clientSecret string, tenantId string, environment string, ) *AzureClientSecretCredentialFields`

NewAzureClientSecretCredentialFields instantiates a new AzureClientSecretCredentialFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureClientSecretCredentialFieldsWithDefaults

`func NewAzureClientSecretCredentialFieldsWithDefaults() *AzureClientSecretCredentialFields`

NewAzureClientSecretCredentialFieldsWithDefaults instantiates a new AzureClientSecretCredentialFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *AzureClientSecretCredentialFields) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AzureClientSecretCredentialFields) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AzureClientSecretCredentialFields) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *AzureClientSecretCredentialFields) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AzureClientSecretCredentialFields) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AzureClientSecretCredentialFields) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetTenantId

`func (o *AzureClientSecretCredentialFields) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureClientSecretCredentialFields) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureClientSecretCredentialFields) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetEnvironment

`func (o *AzureClientSecretCredentialFields) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AzureClientSecretCredentialFields) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AzureClientSecretCredentialFields) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetMultiSubscription

`func (o *AzureClientSecretCredentialFields) GetMultiSubscription() string`

GetMultiSubscription returns the MultiSubscription field if non-nil, zero value otherwise.

### GetMultiSubscriptionOk

`func (o *AzureClientSecretCredentialFields) GetMultiSubscriptionOk() (*string, bool)`

GetMultiSubscriptionOk returns a tuple with the MultiSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiSubscription

`func (o *AzureClientSecretCredentialFields) SetMultiSubscription(v string)`

SetMultiSubscription sets MultiSubscription field to given value.

### HasMultiSubscription

`func (o *AzureClientSecretCredentialFields) HasMultiSubscription() bool`

HasMultiSubscription returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *AzureClientSecretCredentialFields) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureClientSecretCredentialFields) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureClientSecretCredentialFields) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *AzureClientSecretCredentialFields) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


