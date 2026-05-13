# GroupPut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **int64** |  | [optional] 
**OrgDefaultRole** | Pointer to **string** |  | [optional] 
**OrgRoles** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGroupPut

`func NewGroupPut() *GroupPut`

NewGroupPut instantiates a new GroupPut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupPutWithDefaults

`func NewGroupPutWithDefaults() *GroupPut`

NewGroupPutWithDefaults instantiates a new GroupPut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupPut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupPut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupPut) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupPut) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *GroupPut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupPut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupPut) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupPut) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *GroupPut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupPut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupPut) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupPut) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExpiresAt

`func (o *GroupPut) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GroupPut) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GroupPut) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GroupPut) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetOrgDefaultRole

`func (o *GroupPut) GetOrgDefaultRole() string`

GetOrgDefaultRole returns the OrgDefaultRole field if non-nil, zero value otherwise.

### GetOrgDefaultRoleOk

`func (o *GroupPut) GetOrgDefaultRoleOk() (*string, bool)`

GetOrgDefaultRoleOk returns a tuple with the OrgDefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgDefaultRole

`func (o *GroupPut) SetOrgDefaultRole(v string)`

SetOrgDefaultRole sets OrgDefaultRole field to given value.

### HasOrgDefaultRole

`func (o *GroupPut) HasOrgDefaultRole() bool`

HasOrgDefaultRole returns a boolean if a field has been set.

### GetOrgRoles

`func (o *GroupPut) GetOrgRoles() map[string]interface{}`

GetOrgRoles returns the OrgRoles field if non-nil, zero value otherwise.

### GetOrgRolesOk

`func (o *GroupPut) GetOrgRolesOk() (*map[string]interface{}, bool)`

GetOrgRolesOk returns a tuple with the OrgRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgRoles

`func (o *GroupPut) SetOrgRoles(v map[string]interface{})`

SetOrgRoles sets OrgRoles field to given value.

### HasOrgRoles

`func (o *GroupPut) HasOrgRoles() bool`

HasOrgRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


