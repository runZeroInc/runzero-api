# CredentialOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Acl** | Pointer to **map[string]interface{}** |  | [optional] 
**Global** | Pointer to **bool** |  | [optional] 
**Cidrs** | Pointer to **string** |  | [optional] 
**Secret** | Pointer to [**CredentialFields**](CredentialFields.md) |  | [optional] 

## Methods

### NewCredentialOptions

`func NewCredentialOptions() *CredentialOptions`

NewCredentialOptions instantiates a new CredentialOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialOptionsWithDefaults

`func NewCredentialOptionsWithDefaults() *CredentialOptions`

NewCredentialOptionsWithDefaults instantiates a new CredentialOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CredentialOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CredentialOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CredentialOptions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CredentialOptions) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *CredentialOptions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CredentialOptions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CredentialOptions) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CredentialOptions) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAcl

`func (o *CredentialOptions) GetAcl() map[string]interface{}`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *CredentialOptions) GetAclOk() (*map[string]interface{}, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *CredentialOptions) SetAcl(v map[string]interface{})`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *CredentialOptions) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### GetGlobal

`func (o *CredentialOptions) GetGlobal() bool`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *CredentialOptions) GetGlobalOk() (*bool, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *CredentialOptions) SetGlobal(v bool)`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *CredentialOptions) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.

### GetCidrs

`func (o *CredentialOptions) GetCidrs() string`

GetCidrs returns the Cidrs field if non-nil, zero value otherwise.

### GetCidrsOk

`func (o *CredentialOptions) GetCidrsOk() (*string, bool)`

GetCidrsOk returns a tuple with the Cidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrs

`func (o *CredentialOptions) SetCidrs(v string)`

SetCidrs sets Cidrs field to given value.

### HasCidrs

`func (o *CredentialOptions) HasCidrs() bool`

HasCidrs returns a boolean if a field has been set.

### GetSecret

`func (o *CredentialOptions) GetSecret() CredentialFields`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *CredentialOptions) GetSecretOk() (*CredentialFields, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *CredentialOptions) SetSecret(v CredentialFields)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *CredentialOptions) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


