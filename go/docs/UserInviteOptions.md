# UserInviteOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**ClientAdmin** | Pointer to **bool** |  | [optional] 
**OrgDefaultRole** | Pointer to **string** |  | [optional] 
**OrgRoles** | Pointer to **map[string]interface{}** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewUserInviteOptions

`func NewUserInviteOptions() *UserInviteOptions`

NewUserInviteOptions instantiates a new UserInviteOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInviteOptionsWithDefaults

`func NewUserInviteOptionsWithDefaults() *UserInviteOptions`

NewUserInviteOptionsWithDefaults instantiates a new UserInviteOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *UserInviteOptions) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserInviteOptions) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserInviteOptions) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserInviteOptions) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UserInviteOptions) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserInviteOptions) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserInviteOptions) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserInviteOptions) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *UserInviteOptions) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserInviteOptions) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserInviteOptions) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserInviteOptions) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetClientAdmin

`func (o *UserInviteOptions) GetClientAdmin() bool`

GetClientAdmin returns the ClientAdmin field if non-nil, zero value otherwise.

### GetClientAdminOk

`func (o *UserInviteOptions) GetClientAdminOk() (*bool, bool)`

GetClientAdminOk returns a tuple with the ClientAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAdmin

`func (o *UserInviteOptions) SetClientAdmin(v bool)`

SetClientAdmin sets ClientAdmin field to given value.

### HasClientAdmin

`func (o *UserInviteOptions) HasClientAdmin() bool`

HasClientAdmin returns a boolean if a field has been set.

### GetOrgDefaultRole

`func (o *UserInviteOptions) GetOrgDefaultRole() string`

GetOrgDefaultRole returns the OrgDefaultRole field if non-nil, zero value otherwise.

### GetOrgDefaultRoleOk

`func (o *UserInviteOptions) GetOrgDefaultRoleOk() (*string, bool)`

GetOrgDefaultRoleOk returns a tuple with the OrgDefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgDefaultRole

`func (o *UserInviteOptions) SetOrgDefaultRole(v string)`

SetOrgDefaultRole sets OrgDefaultRole field to given value.

### HasOrgDefaultRole

`func (o *UserInviteOptions) HasOrgDefaultRole() bool`

HasOrgDefaultRole returns a boolean if a field has been set.

### GetOrgRoles

`func (o *UserInviteOptions) GetOrgRoles() map[string]interface{}`

GetOrgRoles returns the OrgRoles field if non-nil, zero value otherwise.

### GetOrgRolesOk

`func (o *UserInviteOptions) GetOrgRolesOk() (*map[string]interface{}, bool)`

GetOrgRolesOk returns a tuple with the OrgRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgRoles

`func (o *UserInviteOptions) SetOrgRoles(v map[string]interface{})`

SetOrgRoles sets OrgRoles field to given value.

### HasOrgRoles

`func (o *UserInviteOptions) HasOrgRoles() bool`

HasOrgRoles returns a boolean if a field has been set.

### GetSubject

`func (o *UserInviteOptions) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *UserInviteOptions) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *UserInviteOptions) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *UserInviteOptions) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetMessage

`func (o *UserInviteOptions) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UserInviteOptions) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UserInviteOptions) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UserInviteOptions) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


