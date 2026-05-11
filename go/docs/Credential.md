# Credential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ClientId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | The service the credentials are for. | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**CreatedById** | Pointer to **string** |  | [optional] 
**CreatedByEmail** | Pointer to **string** |  | [optional] 
**Acl** | Pointer to **map[string]interface{}** |  | [optional] 
**Global** | Pointer to **bool** |  | [optional] 
**Cidrs** | Pointer to **[]string** |  | [optional] 
**LastUsedAt** | Pointer to **int64** |  | [optional] 
**LastUsedById** | Pointer to **string** |  | [optional] 

## Methods

### NewCredential

`func NewCredential(id string, ) *Credential`

NewCredential instantiates a new Credential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialWithDefaults

`func NewCredentialWithDefaults() *Credential`

NewCredentialWithDefaults instantiates a new Credential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Credential) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Credential) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Credential) SetId(v string)`

SetId sets Id field to given value.


### GetClientId

`func (o *Credential) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Credential) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Credential) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Credential) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetName

`func (o *Credential) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Credential) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Credential) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Credential) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Credential) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Credential) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Credential) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Credential) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Credential) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Credential) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Credential) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Credential) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedById

`func (o *Credential) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *Credential) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *Credential) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *Credential) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetCreatedByEmail

`func (o *Credential) GetCreatedByEmail() string`

GetCreatedByEmail returns the CreatedByEmail field if non-nil, zero value otherwise.

### GetCreatedByEmailOk

`func (o *Credential) GetCreatedByEmailOk() (*string, bool)`

GetCreatedByEmailOk returns a tuple with the CreatedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByEmail

`func (o *Credential) SetCreatedByEmail(v string)`

SetCreatedByEmail sets CreatedByEmail field to given value.

### HasCreatedByEmail

`func (o *Credential) HasCreatedByEmail() bool`

HasCreatedByEmail returns a boolean if a field has been set.

### GetAcl

`func (o *Credential) GetAcl() map[string]interface{}`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *Credential) GetAclOk() (*map[string]interface{}, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *Credential) SetAcl(v map[string]interface{})`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *Credential) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### GetGlobal

`func (o *Credential) GetGlobal() bool`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *Credential) GetGlobalOk() (*bool, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *Credential) SetGlobal(v bool)`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *Credential) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.

### GetCidrs

`func (o *Credential) GetCidrs() []string`

GetCidrs returns the Cidrs field if non-nil, zero value otherwise.

### GetCidrsOk

`func (o *Credential) GetCidrsOk() (*[]string, bool)`

GetCidrsOk returns a tuple with the Cidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrs

`func (o *Credential) SetCidrs(v []string)`

SetCidrs sets Cidrs field to given value.

### HasCidrs

`func (o *Credential) HasCidrs() bool`

HasCidrs returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *Credential) GetLastUsedAt() int64`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *Credential) GetLastUsedAtOk() (*int64, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *Credential) SetLastUsedAt(v int64)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *Credential) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetLastUsedById

`func (o *Credential) GetLastUsedById() string`

GetLastUsedById returns the LastUsedById field if non-nil, zero value otherwise.

### GetLastUsedByIdOk

`func (o *Credential) GetLastUsedByIdOk() (*string, bool)`

GetLastUsedByIdOk returns a tuple with the LastUsedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedById

`func (o *Credential) SetLastUsedById(v string)`

SetLastUsedById sets LastUsedById field to given value.

### HasLastUsedById

`func (o *Credential) HasLastUsedById() bool`

HasLastUsedById returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


