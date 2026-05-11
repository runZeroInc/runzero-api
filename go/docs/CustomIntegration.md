# CustomIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The unique name of the custom integration, without spaces. | 
**Icon** | Pointer to **string** | Base64 encoded png with maximum size 256x256 pixels | [optional] 
**Description** | Pointer to **string** | A text description of the custom integration | [optional] 
**Id** | **string** | The unique ID of the object | 
**ClientId** | **string** | The unique ID of the runZero client/customer account that owns the object | 
**CreatedById** | **string** | The unique ID of the entity that created the object | 
**CreatedAt** | **time.Time** | A timestamp indicating creation time of the object | 
**UpdatedAt** | **time.Time** | A timestamp indicating last modified time of the object | 
**DestroyedAt** | Pointer to **time.Time** | A timestamp indicating deletion time of the object | [optional] 

## Methods

### NewCustomIntegration

`func NewCustomIntegration(name string, id string, clientId string, createdById string, createdAt time.Time, updatedAt time.Time, ) *CustomIntegration`

NewCustomIntegration instantiates a new CustomIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomIntegrationWithDefaults

`func NewCustomIntegrationWithDefaults() *CustomIntegration`

NewCustomIntegrationWithDefaults instantiates a new CustomIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomIntegration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomIntegration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomIntegration) SetName(v string)`

SetName sets Name field to given value.


### GetIcon

`func (o *CustomIntegration) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CustomIntegration) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CustomIntegration) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *CustomIntegration) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetDescription

`func (o *CustomIntegration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomIntegration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomIntegration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomIntegration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *CustomIntegration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomIntegration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomIntegration) SetId(v string)`

SetId sets Id field to given value.


### GetClientId

`func (o *CustomIntegration) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CustomIntegration) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CustomIntegration) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetCreatedById

`func (o *CustomIntegration) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *CustomIntegration) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *CustomIntegration) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetCreatedAt

`func (o *CustomIntegration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomIntegration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomIntegration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CustomIntegration) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CustomIntegration) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CustomIntegration) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDestroyedAt

`func (o *CustomIntegration) GetDestroyedAt() time.Time`

GetDestroyedAt returns the DestroyedAt field if non-nil, zero value otherwise.

### GetDestroyedAtOk

`func (o *CustomIntegration) GetDestroyedAtOk() (*time.Time, bool)`

GetDestroyedAtOk returns a tuple with the DestroyedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyedAt

`func (o *CustomIntegration) SetDestroyedAt(v time.Time)`

SetDestroyedAt sets DestroyedAt field to given value.

### HasDestroyedAt

`func (o *CustomIntegration) HasDestroyedAt() bool`

HasDestroyedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


