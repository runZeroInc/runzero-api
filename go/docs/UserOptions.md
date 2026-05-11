# UserOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**ClientAdmin** | Pointer to **bool** |  | [optional] 
**OrgDefaultRole** | Pointer to **string** |  | [optional] 
**OrgRoles** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUserOptions

`func NewUserOptions() *UserOptions`

NewUserOptions instantiates a new UserOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserOptionsWithDefaults

`func NewUserOptionsWithDefaults() *UserOptions`

NewUserOptionsWithDefaults instantiates a new UserOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *UserOptions) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserOptions) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserOptions) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserOptions) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UserOptions) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserOptions) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserOptions) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserOptions) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *UserOptions) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserOptions) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserOptions) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserOptions) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetClientAdmin

`func (o *UserOptions) GetClientAdmin() bool`

GetClientAdmin returns the ClientAdmin field if non-nil, zero value otherwise.

### GetClientAdminOk

`func (o *UserOptions) GetClientAdminOk() (*bool, bool)`

GetClientAdminOk returns a tuple with the ClientAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAdmin

`func (o *UserOptions) SetClientAdmin(v bool)`

SetClientAdmin sets ClientAdmin field to given value.

### HasClientAdmin

`func (o *UserOptions) HasClientAdmin() bool`

HasClientAdmin returns a boolean if a field has been set.

### GetOrgDefaultRole

`func (o *UserOptions) GetOrgDefaultRole() string`

GetOrgDefaultRole returns the OrgDefaultRole field if non-nil, zero value otherwise.

### GetOrgDefaultRoleOk

`func (o *UserOptions) GetOrgDefaultRoleOk() (*string, bool)`

GetOrgDefaultRoleOk returns a tuple with the OrgDefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgDefaultRole

`func (o *UserOptions) SetOrgDefaultRole(v string)`

SetOrgDefaultRole sets OrgDefaultRole field to given value.

### HasOrgDefaultRole

`func (o *UserOptions) HasOrgDefaultRole() bool`

HasOrgDefaultRole returns a boolean if a field has been set.

### GetOrgRoles

`func (o *UserOptions) GetOrgRoles() map[string]interface{}`

GetOrgRoles returns the OrgRoles field if non-nil, zero value otherwise.

### GetOrgRolesOk

`func (o *UserOptions) GetOrgRolesOk() (*map[string]interface{}, bool)`

GetOrgRolesOk returns a tuple with the OrgRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgRoles

`func (o *UserOptions) SetOrgRoles(v map[string]interface{})`

SetOrgRoles sets OrgRoles field to given value.

### HasOrgRoles

`func (o *UserOptions) HasOrgRoles() bool`

HasOrgRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


