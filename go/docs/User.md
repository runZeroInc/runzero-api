# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ClientId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**ClientAdmin** | Pointer to **bool** |  | [optional] 
**OrgDefaultRole** | Pointer to **string** |  | [optional] 
**OrgRoles** | Pointer to **map[string]interface{}** |  | [optional] 
**PasswordEnabledAt** | Pointer to **int64** |  | [optional] 
**InviteTokenExpiration** | Pointer to **int64** |  | [optional] 
**ResetTokenExpiration** | Pointer to **int64** |  | [optional] 
**TermsAgreedAsOf** | Pointer to **int64** |  | [optional] 
**LastLoginIp** | Pointer to **NullableString** |  | [optional] 
**LastLoginAt** | Pointer to **int64** |  | [optional] 
**LastLoginUa** | Pointer to **string** |  | [optional] 
**LastActivityAt** | Pointer to **int64** |  | [optional] 
**SsoOnly** | Pointer to **bool** |  | [optional] 
**LoginFailures** | Pointer to **int64** |  | [optional] 
**Actions** | Pointer to **int64** |  | [optional] 
**LastActionAt** | Pointer to **int64** |  | [optional] 
**MfaEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewUser

`func NewUser(id string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.


### GetClientId

`func (o *User) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *User) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *User) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *User) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *User) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *User) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *User) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *User) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *User) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *User) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *User) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *User) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *User) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *User) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *User) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFirstName

`func (o *User) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *User) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *User) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *User) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *User) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *User) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *User) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *User) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetClientAdmin

`func (o *User) GetClientAdmin() bool`

GetClientAdmin returns the ClientAdmin field if non-nil, zero value otherwise.

### GetClientAdminOk

`func (o *User) GetClientAdminOk() (*bool, bool)`

GetClientAdminOk returns a tuple with the ClientAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAdmin

`func (o *User) SetClientAdmin(v bool)`

SetClientAdmin sets ClientAdmin field to given value.

### HasClientAdmin

`func (o *User) HasClientAdmin() bool`

HasClientAdmin returns a boolean if a field has been set.

### GetOrgDefaultRole

`func (o *User) GetOrgDefaultRole() string`

GetOrgDefaultRole returns the OrgDefaultRole field if non-nil, zero value otherwise.

### GetOrgDefaultRoleOk

`func (o *User) GetOrgDefaultRoleOk() (*string, bool)`

GetOrgDefaultRoleOk returns a tuple with the OrgDefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgDefaultRole

`func (o *User) SetOrgDefaultRole(v string)`

SetOrgDefaultRole sets OrgDefaultRole field to given value.

### HasOrgDefaultRole

`func (o *User) HasOrgDefaultRole() bool`

HasOrgDefaultRole returns a boolean if a field has been set.

### GetOrgRoles

`func (o *User) GetOrgRoles() map[string]interface{}`

GetOrgRoles returns the OrgRoles field if non-nil, zero value otherwise.

### GetOrgRolesOk

`func (o *User) GetOrgRolesOk() (*map[string]interface{}, bool)`

GetOrgRolesOk returns a tuple with the OrgRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgRoles

`func (o *User) SetOrgRoles(v map[string]interface{})`

SetOrgRoles sets OrgRoles field to given value.

### HasOrgRoles

`func (o *User) HasOrgRoles() bool`

HasOrgRoles returns a boolean if a field has been set.

### GetPasswordEnabledAt

`func (o *User) GetPasswordEnabledAt() int64`

GetPasswordEnabledAt returns the PasswordEnabledAt field if non-nil, zero value otherwise.

### GetPasswordEnabledAtOk

`func (o *User) GetPasswordEnabledAtOk() (*int64, bool)`

GetPasswordEnabledAtOk returns a tuple with the PasswordEnabledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordEnabledAt

`func (o *User) SetPasswordEnabledAt(v int64)`

SetPasswordEnabledAt sets PasswordEnabledAt field to given value.

### HasPasswordEnabledAt

`func (o *User) HasPasswordEnabledAt() bool`

HasPasswordEnabledAt returns a boolean if a field has been set.

### GetInviteTokenExpiration

`func (o *User) GetInviteTokenExpiration() int64`

GetInviteTokenExpiration returns the InviteTokenExpiration field if non-nil, zero value otherwise.

### GetInviteTokenExpirationOk

`func (o *User) GetInviteTokenExpirationOk() (*int64, bool)`

GetInviteTokenExpirationOk returns a tuple with the InviteTokenExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteTokenExpiration

`func (o *User) SetInviteTokenExpiration(v int64)`

SetInviteTokenExpiration sets InviteTokenExpiration field to given value.

### HasInviteTokenExpiration

`func (o *User) HasInviteTokenExpiration() bool`

HasInviteTokenExpiration returns a boolean if a field has been set.

### GetResetTokenExpiration

`func (o *User) GetResetTokenExpiration() int64`

GetResetTokenExpiration returns the ResetTokenExpiration field if non-nil, zero value otherwise.

### GetResetTokenExpirationOk

`func (o *User) GetResetTokenExpirationOk() (*int64, bool)`

GetResetTokenExpirationOk returns a tuple with the ResetTokenExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetTokenExpiration

`func (o *User) SetResetTokenExpiration(v int64)`

SetResetTokenExpiration sets ResetTokenExpiration field to given value.

### HasResetTokenExpiration

`func (o *User) HasResetTokenExpiration() bool`

HasResetTokenExpiration returns a boolean if a field has been set.

### GetTermsAgreedAsOf

`func (o *User) GetTermsAgreedAsOf() int64`

GetTermsAgreedAsOf returns the TermsAgreedAsOf field if non-nil, zero value otherwise.

### GetTermsAgreedAsOfOk

`func (o *User) GetTermsAgreedAsOfOk() (*int64, bool)`

GetTermsAgreedAsOfOk returns a tuple with the TermsAgreedAsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsAgreedAsOf

`func (o *User) SetTermsAgreedAsOf(v int64)`

SetTermsAgreedAsOf sets TermsAgreedAsOf field to given value.

### HasTermsAgreedAsOf

`func (o *User) HasTermsAgreedAsOf() bool`

HasTermsAgreedAsOf returns a boolean if a field has been set.

### GetLastLoginIp

`func (o *User) GetLastLoginIp() string`

GetLastLoginIp returns the LastLoginIp field if non-nil, zero value otherwise.

### GetLastLoginIpOk

`func (o *User) GetLastLoginIpOk() (*string, bool)`

GetLastLoginIpOk returns a tuple with the LastLoginIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginIp

`func (o *User) SetLastLoginIp(v string)`

SetLastLoginIp sets LastLoginIp field to given value.

### HasLastLoginIp

`func (o *User) HasLastLoginIp() bool`

HasLastLoginIp returns a boolean if a field has been set.

### SetLastLoginIpNil

`func (o *User) SetLastLoginIpNil(b bool)`

 SetLastLoginIpNil sets the value for LastLoginIp to be an explicit nil

### UnsetLastLoginIp
`func (o *User) UnsetLastLoginIp()`

UnsetLastLoginIp ensures that no value is present for LastLoginIp, not even an explicit nil
### GetLastLoginAt

`func (o *User) GetLastLoginAt() int64`

GetLastLoginAt returns the LastLoginAt field if non-nil, zero value otherwise.

### GetLastLoginAtOk

`func (o *User) GetLastLoginAtOk() (*int64, bool)`

GetLastLoginAtOk returns a tuple with the LastLoginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginAt

`func (o *User) SetLastLoginAt(v int64)`

SetLastLoginAt sets LastLoginAt field to given value.

### HasLastLoginAt

`func (o *User) HasLastLoginAt() bool`

HasLastLoginAt returns a boolean if a field has been set.

### GetLastLoginUa

`func (o *User) GetLastLoginUa() string`

GetLastLoginUa returns the LastLoginUa field if non-nil, zero value otherwise.

### GetLastLoginUaOk

`func (o *User) GetLastLoginUaOk() (*string, bool)`

GetLastLoginUaOk returns a tuple with the LastLoginUa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginUa

`func (o *User) SetLastLoginUa(v string)`

SetLastLoginUa sets LastLoginUa field to given value.

### HasLastLoginUa

`func (o *User) HasLastLoginUa() bool`

HasLastLoginUa returns a boolean if a field has been set.

### GetLastActivityAt

`func (o *User) GetLastActivityAt() int64`

GetLastActivityAt returns the LastActivityAt field if non-nil, zero value otherwise.

### GetLastActivityAtOk

`func (o *User) GetLastActivityAtOk() (*int64, bool)`

GetLastActivityAtOk returns a tuple with the LastActivityAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityAt

`func (o *User) SetLastActivityAt(v int64)`

SetLastActivityAt sets LastActivityAt field to given value.

### HasLastActivityAt

`func (o *User) HasLastActivityAt() bool`

HasLastActivityAt returns a boolean if a field has been set.

### GetSsoOnly

`func (o *User) GetSsoOnly() bool`

GetSsoOnly returns the SsoOnly field if non-nil, zero value otherwise.

### GetSsoOnlyOk

`func (o *User) GetSsoOnlyOk() (*bool, bool)`

GetSsoOnlyOk returns a tuple with the SsoOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoOnly

`func (o *User) SetSsoOnly(v bool)`

SetSsoOnly sets SsoOnly field to given value.

### HasSsoOnly

`func (o *User) HasSsoOnly() bool`

HasSsoOnly returns a boolean if a field has been set.

### GetLoginFailures

`func (o *User) GetLoginFailures() int64`

GetLoginFailures returns the LoginFailures field if non-nil, zero value otherwise.

### GetLoginFailuresOk

`func (o *User) GetLoginFailuresOk() (*int64, bool)`

GetLoginFailuresOk returns a tuple with the LoginFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginFailures

`func (o *User) SetLoginFailures(v int64)`

SetLoginFailures sets LoginFailures field to given value.

### HasLoginFailures

`func (o *User) HasLoginFailures() bool`

HasLoginFailures returns a boolean if a field has been set.

### GetActions

`func (o *User) GetActions() int64`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *User) GetActionsOk() (*int64, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *User) SetActions(v int64)`

SetActions sets Actions field to given value.

### HasActions

`func (o *User) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetLastActionAt

`func (o *User) GetLastActionAt() int64`

GetLastActionAt returns the LastActionAt field if non-nil, zero value otherwise.

### GetLastActionAtOk

`func (o *User) GetLastActionAtOk() (*int64, bool)`

GetLastActionAtOk returns a tuple with the LastActionAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActionAt

`func (o *User) SetLastActionAt(v int64)`

SetLastActionAt sets LastActionAt field to given value.

### HasLastActionAt

`func (o *User) HasLastActionAt() bool`

HasLastActionAt returns a boolean if a field has been set.

### GetMfaEnabled

`func (o *User) GetMfaEnabled() bool`

GetMfaEnabled returns the MfaEnabled field if non-nil, zero value otherwise.

### GetMfaEnabledOk

`func (o *User) GetMfaEnabledOk() (*bool, bool)`

GetMfaEnabledOk returns a tuple with the MfaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaEnabled

`func (o *User) SetMfaEnabled(v bool)`

SetMfaEnabled sets MfaEnabled field to given value.

### HasMfaEnabled

`func (o *User) HasMfaEnabled() bool`

HasMfaEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


