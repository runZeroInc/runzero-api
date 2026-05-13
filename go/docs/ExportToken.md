# ExportToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **interface{}** |  | 
**Token** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**LastUsedAt** | Pointer to **int64** |  | [optional] 
**LastUsedBy** | Pointer to **string** |  | [optional] 
**Counter** | Pointer to **int64** |  | [optional] 

## Methods

### NewExportToken

`func NewExportToken(id interface{}, ) *ExportToken`

NewExportToken instantiates a new ExportToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportTokenWithDefaults

`func NewExportTokenWithDefaults() *ExportToken`

NewExportTokenWithDefaults instantiates a new ExportToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExportToken) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExportToken) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExportToken) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *ExportToken) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ExportToken) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetToken

`func (o *ExportToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ExportToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ExportToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ExportToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetComment

`func (o *ExportToken) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ExportToken) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ExportToken) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ExportToken) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ExportToken) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExportToken) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExportToken) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ExportToken) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *ExportToken) GetLastUsedAt() int64`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *ExportToken) GetLastUsedAtOk() (*int64, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *ExportToken) SetLastUsedAt(v int64)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *ExportToken) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetLastUsedBy

`func (o *ExportToken) GetLastUsedBy() string`

GetLastUsedBy returns the LastUsedBy field if non-nil, zero value otherwise.

### GetLastUsedByOk

`func (o *ExportToken) GetLastUsedByOk() (*string, bool)`

GetLastUsedByOk returns a tuple with the LastUsedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedBy

`func (o *ExportToken) SetLastUsedBy(v string)`

SetLastUsedBy sets LastUsedBy field to given value.

### HasLastUsedBy

`func (o *ExportToken) HasLastUsedBy() bool`

HasLastUsedBy returns a boolean if a field has been set.

### GetCounter

`func (o *ExportToken) GetCounter() int64`

GetCounter returns the Counter field if non-nil, zero value otherwise.

### GetCounterOk

`func (o *ExportToken) GetCounterOk() (*int64, bool)`

GetCounterOk returns a tuple with the Counter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounter

`func (o *ExportToken) SetCounter(v int64)`

SetCounter sets Counter field to given value.

### HasCounter

`func (o *ExportToken) HasCounter() bool`

HasCounter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


