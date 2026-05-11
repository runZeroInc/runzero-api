# GroupPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **int64** |  | [optional] 
**OrgDefaultRole** | Pointer to **string** |  | [optional] 
**OrgRoles** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGroupPost

`func NewGroupPost() *GroupPost`

NewGroupPost instantiates a new GroupPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupPostWithDefaults

`func NewGroupPostWithDefaults() *GroupPost`

NewGroupPostWithDefaults instantiates a new GroupPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GroupPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupPost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupPost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *GroupPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupPost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupPost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExpiresAt

`func (o *GroupPost) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GroupPost) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GroupPost) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GroupPost) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetOrgDefaultRole

`func (o *GroupPost) GetOrgDefaultRole() string`

GetOrgDefaultRole returns the OrgDefaultRole field if non-nil, zero value otherwise.

### GetOrgDefaultRoleOk

`func (o *GroupPost) GetOrgDefaultRoleOk() (*string, bool)`

GetOrgDefaultRoleOk returns a tuple with the OrgDefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgDefaultRole

`func (o *GroupPost) SetOrgDefaultRole(v string)`

SetOrgDefaultRole sets OrgDefaultRole field to given value.

### HasOrgDefaultRole

`func (o *GroupPost) HasOrgDefaultRole() bool`

HasOrgDefaultRole returns a boolean if a field has been set.

### GetOrgRoles

`func (o *GroupPost) GetOrgRoles() map[string]interface{}`

GetOrgRoles returns the OrgRoles field if non-nil, zero value otherwise.

### GetOrgRolesOk

`func (o *GroupPost) GetOrgRolesOk() (*map[string]interface{}, bool)`

GetOrgRolesOk returns a tuple with the OrgRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgRoles

`func (o *GroupPost) SetOrgRoles(v map[string]interface{})`

SetOrgRoles sets OrgRoles field to given value.

### HasOrgRoles

`func (o *GroupPost) HasOrgRoles() bool`

HasOrgRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


