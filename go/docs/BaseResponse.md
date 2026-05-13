# BaseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the object | 
**ClientId** | **string** | The unique ID of the runZero client/customer account that owns the object | 
**CreatedById** | **string** | The unique ID of the entity that created the object | 
**CreatedAt** | **time.Time** | A timestamp indicating creation time of the object | 
**UpdatedAt** | **time.Time** | A timestamp indicating last modified time of the object | 
**DestroyedAt** | Pointer to **time.Time** | A timestamp indicating deletion time of the object | [optional] 

## Methods

### NewBaseResponse

`func NewBaseResponse(id string, clientId string, createdById string, createdAt time.Time, updatedAt time.Time, ) *BaseResponse`

NewBaseResponse instantiates a new BaseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseResponseWithDefaults

`func NewBaseResponseWithDefaults() *BaseResponse`

NewBaseResponseWithDefaults instantiates a new BaseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseResponse) SetId(v string)`

SetId sets Id field to given value.


### GetClientId

`func (o *BaseResponse) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *BaseResponse) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *BaseResponse) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetCreatedById

`func (o *BaseResponse) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *BaseResponse) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *BaseResponse) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetCreatedAt

`func (o *BaseResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BaseResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BaseResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BaseResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BaseResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BaseResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDestroyedAt

`func (o *BaseResponse) GetDestroyedAt() time.Time`

GetDestroyedAt returns the DestroyedAt field if non-nil, zero value otherwise.

### GetDestroyedAtOk

`func (o *BaseResponse) GetDestroyedAtOk() (*time.Time, bool)`

GetDestroyedAtOk returns a tuple with the DestroyedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyedAt

`func (o *BaseResponse) SetDestroyedAt(v time.Time)`

SetDestroyedAt sets DestroyedAt field to given value.

### HasDestroyedAt

`func (o *BaseResponse) HasDestroyedAt() bool`

HasDestroyedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


