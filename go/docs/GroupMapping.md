# GroupMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**GroupId** | **string** |  | 
**GroupName** | Pointer to **string** |  | [optional] 
**SsoAttribute** | **string** |  | 
**SsoValue** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**CreatedByEmail** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 

## Methods

### NewGroupMapping

`func NewGroupMapping(id string, groupId string, ssoAttribute string, ssoValue string, ) *GroupMapping`

NewGroupMapping instantiates a new GroupMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupMappingWithDefaults

`func NewGroupMappingWithDefaults() *GroupMapping`

NewGroupMappingWithDefaults instantiates a new GroupMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupMapping) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupMapping) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupMapping) SetId(v string)`

SetId sets Id field to given value.


### GetGroupId

`func (o *GroupMapping) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupMapping) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupMapping) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetGroupName

`func (o *GroupMapping) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *GroupMapping) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *GroupMapping) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *GroupMapping) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetSsoAttribute

`func (o *GroupMapping) GetSsoAttribute() string`

GetSsoAttribute returns the SsoAttribute field if non-nil, zero value otherwise.

### GetSsoAttributeOk

`func (o *GroupMapping) GetSsoAttributeOk() (*string, bool)`

GetSsoAttributeOk returns a tuple with the SsoAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoAttribute

`func (o *GroupMapping) SetSsoAttribute(v string)`

SetSsoAttribute sets SsoAttribute field to given value.


### GetSsoValue

`func (o *GroupMapping) GetSsoValue() string`

GetSsoValue returns the SsoValue field if non-nil, zero value otherwise.

### GetSsoValueOk

`func (o *GroupMapping) GetSsoValueOk() (*string, bool)`

GetSsoValueOk returns a tuple with the SsoValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoValue

`func (o *GroupMapping) SetSsoValue(v string)`

SetSsoValue sets SsoValue field to given value.


### GetDescription

`func (o *GroupMapping) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupMapping) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupMapping) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupMapping) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedByEmail

`func (o *GroupMapping) GetCreatedByEmail() string`

GetCreatedByEmail returns the CreatedByEmail field if non-nil, zero value otherwise.

### GetCreatedByEmailOk

`func (o *GroupMapping) GetCreatedByEmailOk() (*string, bool)`

GetCreatedByEmailOk returns a tuple with the CreatedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByEmail

`func (o *GroupMapping) SetCreatedByEmail(v string)`

SetCreatedByEmail sets CreatedByEmail field to given value.

### HasCreatedByEmail

`func (o *GroupMapping) HasCreatedByEmail() bool`

HasCreatedByEmail returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GroupMapping) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GroupMapping) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GroupMapping) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GroupMapping) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GroupMapping) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GroupMapping) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GroupMapping) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GroupMapping) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


