# OrganizationAPIKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ClientId** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**LastUsedAt** | Pointer to **int64** |  | [optional] 
**LastUsedIp** | Pointer to **string** |  | [optional] 
**LastUsedUa** | Pointer to **string** |  | [optional] 
**Counter** | Pointer to **int64** |  | [optional] 
**UsageToday** | Pointer to **int64** |  | [optional] 
**UsageLimit** | Pointer to **int64** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Inactive** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewOrganizationAPIKey

`func NewOrganizationAPIKey(id string, ) *OrganizationAPIKey`

NewOrganizationAPIKey instantiates a new OrganizationAPIKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationAPIKeyWithDefaults

`func NewOrganizationAPIKeyWithDefaults() *OrganizationAPIKey`

NewOrganizationAPIKeyWithDefaults instantiates a new OrganizationAPIKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationAPIKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationAPIKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationAPIKey) SetId(v string)`

SetId sets Id field to given value.


### GetClientId

`func (o *OrganizationAPIKey) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OrganizationAPIKey) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OrganizationAPIKey) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OrganizationAPIKey) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *OrganizationAPIKey) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OrganizationAPIKey) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OrganizationAPIKey) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *OrganizationAPIKey) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrganizationAPIKey) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationAPIKey) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationAPIKey) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrganizationAPIKey) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *OrganizationAPIKey) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *OrganizationAPIKey) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *OrganizationAPIKey) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *OrganizationAPIKey) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetComment

`func (o *OrganizationAPIKey) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *OrganizationAPIKey) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *OrganizationAPIKey) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *OrganizationAPIKey) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *OrganizationAPIKey) GetLastUsedAt() int64`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *OrganizationAPIKey) GetLastUsedAtOk() (*int64, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *OrganizationAPIKey) SetLastUsedAt(v int64)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *OrganizationAPIKey) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetLastUsedIp

`func (o *OrganizationAPIKey) GetLastUsedIp() string`

GetLastUsedIp returns the LastUsedIp field if non-nil, zero value otherwise.

### GetLastUsedIpOk

`func (o *OrganizationAPIKey) GetLastUsedIpOk() (*string, bool)`

GetLastUsedIpOk returns a tuple with the LastUsedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedIp

`func (o *OrganizationAPIKey) SetLastUsedIp(v string)`

SetLastUsedIp sets LastUsedIp field to given value.

### HasLastUsedIp

`func (o *OrganizationAPIKey) HasLastUsedIp() bool`

HasLastUsedIp returns a boolean if a field has been set.

### GetLastUsedUa

`func (o *OrganizationAPIKey) GetLastUsedUa() string`

GetLastUsedUa returns the LastUsedUa field if non-nil, zero value otherwise.

### GetLastUsedUaOk

`func (o *OrganizationAPIKey) GetLastUsedUaOk() (*string, bool)`

GetLastUsedUaOk returns a tuple with the LastUsedUa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedUa

`func (o *OrganizationAPIKey) SetLastUsedUa(v string)`

SetLastUsedUa sets LastUsedUa field to given value.

### HasLastUsedUa

`func (o *OrganizationAPIKey) HasLastUsedUa() bool`

HasLastUsedUa returns a boolean if a field has been set.

### GetCounter

`func (o *OrganizationAPIKey) GetCounter() int64`

GetCounter returns the Counter field if non-nil, zero value otherwise.

### GetCounterOk

`func (o *OrganizationAPIKey) GetCounterOk() (*int64, bool)`

GetCounterOk returns a tuple with the Counter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounter

`func (o *OrganizationAPIKey) SetCounter(v int64)`

SetCounter sets Counter field to given value.

### HasCounter

`func (o *OrganizationAPIKey) HasCounter() bool`

HasCounter returns a boolean if a field has been set.

### GetUsageToday

`func (o *OrganizationAPIKey) GetUsageToday() int64`

GetUsageToday returns the UsageToday field if non-nil, zero value otherwise.

### GetUsageTodayOk

`func (o *OrganizationAPIKey) GetUsageTodayOk() (*int64, bool)`

GetUsageTodayOk returns a tuple with the UsageToday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageToday

`func (o *OrganizationAPIKey) SetUsageToday(v int64)`

SetUsageToday sets UsageToday field to given value.

### HasUsageToday

`func (o *OrganizationAPIKey) HasUsageToday() bool`

HasUsageToday returns a boolean if a field has been set.

### GetUsageLimit

`func (o *OrganizationAPIKey) GetUsageLimit() int64`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *OrganizationAPIKey) GetUsageLimitOk() (*int64, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *OrganizationAPIKey) SetUsageLimit(v int64)`

SetUsageLimit sets UsageLimit field to given value.

### HasUsageLimit

`func (o *OrganizationAPIKey) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### GetToken

`func (o *OrganizationAPIKey) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *OrganizationAPIKey) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *OrganizationAPIKey) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *OrganizationAPIKey) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetInactive

`func (o *OrganizationAPIKey) GetInactive() bool`

GetInactive returns the Inactive field if non-nil, zero value otherwise.

### GetInactiveOk

`func (o *OrganizationAPIKey) GetInactiveOk() (*bool, bool)`

GetInactiveOk returns a tuple with the Inactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactive

`func (o *OrganizationAPIKey) SetInactive(v bool)`

SetInactive sets Inactive field to given value.

### HasInactive

`func (o *OrganizationAPIKey) HasInactive() bool`

HasInactive returns a boolean if a field has been set.

### GetType

`func (o *OrganizationAPIKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrganizationAPIKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrganizationAPIKey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OrganizationAPIKey) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


