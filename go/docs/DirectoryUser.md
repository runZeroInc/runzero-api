# DirectoryUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**DirectoryUserId** | Pointer to **string** |  | [optional] 
**DirectoryUserOrganizationId** | Pointer to **string** |  | [optional] 
**DirectoryUserSiteId** | Pointer to **string** |  | [optional] 
**DirectoryUserSourceId** | Pointer to **int32** |  | [optional] 
**DirectoryUserCreatedAt** | Pointer to **int64** |  | [optional] 
**DirectoryUserUpdatedAt** | Pointer to **int64** |  | [optional] 
**DirectoryUserUserId** | Pointer to **string** |  | [optional] 
**DirectoryUserGroupIds** | Pointer to **[]string** |  | [optional] 
**DirectoryUserDisplayName** | Pointer to **string** |  | [optional] 
**DirectoryUserName** | Pointer to **string** |  | [optional] 
**DirectoryUserFirstName** | Pointer to **string** |  | [optional] 
**DirectoryUserLastName** | Pointer to **string** |  | [optional] 
**DirectoryUserDescription** | Pointer to **string** |  | [optional] 
**DirectoryUserEmail** | Pointer to **string** |  | [optional] 
**DirectoryUserEmailNormalized** | Pointer to **string** |  | [optional] 
**DirectoryUserPhone** | Pointer to **string** |  | [optional] 
**DirectoryUserTitle** | Pointer to **string** |  | [optional] 
**DirectoryUserLocation** | Pointer to **string** |  | [optional] 
**DirectoryUserLastLogonAt** | Pointer to **int64** |  | [optional] 
**DirectoryUserAttributes** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewDirectoryUser

`func NewDirectoryUser(id string, ) *DirectoryUser`

NewDirectoryUser instantiates a new DirectoryUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryUserWithDefaults

`func NewDirectoryUserWithDefaults() *DirectoryUser`

NewDirectoryUserWithDefaults instantiates a new DirectoryUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DirectoryUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DirectoryUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DirectoryUser) SetId(v string)`

SetId sets Id field to given value.


### GetDirectoryUserId

`func (o *DirectoryUser) GetDirectoryUserId() string`

GetDirectoryUserId returns the DirectoryUserId field if non-nil, zero value otherwise.

### GetDirectoryUserIdOk

`func (o *DirectoryUser) GetDirectoryUserIdOk() (*string, bool)`

GetDirectoryUserIdOk returns a tuple with the DirectoryUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryUserId

`func (o *DirectoryUser) SetDirectoryUserId(v string)`

SetDirectoryUserId sets DirectoryUserId field to given value.

### HasDirectoryUserId

`func (o *DirectoryUser) HasDirectoryUserId() bool`

HasDirectoryUserId returns a boolean if a field has been set.

### GetDirectoryUserOrganizationId

`func (o *DirectoryUser) GetDirectoryUserOrganizationId() string`

GetDirectoryUserOrganizationId returns the DirectoryUserOrganizationId field if non-nil, zero value otherwise.

### GetDirectoryUserOrganizationIdOk

`func (o *DirectoryUser) GetDirectoryUserOrganizationIdOk() (*string, bool)`

GetDirectoryUserOrganizationIdOk returns a tuple with the DirectoryUserOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryUserOrganizationId

`func (o *DirectoryUser) SetDirectoryUserOrganizationId(v string)`

SetDirectoryUserOrganizationId sets DirectoryUserOrganizationId field to given value.

### HasDirectoryUserOrganizationId

`func (o *DirectoryUser) HasDirectoryUserOrganizationId() bool`

HasDirectoryUserOrganizationId returns a boolean if a field has been set.

### GetDirectoryUserSiteId

`func (o *DirectoryUser) GetDirectoryUserSiteId() string`

GetDirectoryUserSiteId returns the DirectoryUserSiteId field if non-nil, zero value otherwise.

### GetDirectoryUserSiteIdOk

`func (o *DirectoryUser) GetDirectoryUserSiteIdOk() (*string, bool)`

GetDirectoryUserSiteIdOk returns a tuple with the DirectoryUserSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryUserSiteId

`func (o *DirectoryUser) SetDirectoryUserSiteId(v string)`

SetDirectoryUserSiteId sets DirectoryUserSiteId field to given value.

### HasDirectoryUserSiteId

`func (o *DirectoryUser) HasDirectoryUserSiteId() bool`

HasDirectoryUserSiteId returns a boolean if a field has been set.

### GetDirectoryUserSourceId

`func (o *DirectoryUser) GetDirectoryUserSourceId() int32`

GetDirectoryUserSourceId returns the DirectoryUserSourceId field if non-nil, zero value otherwise.

### GetDirectoryUserSourceIdOk

`func (o *DirectoryUser) GetDirectoryUserSourceIdOk() (*int32, bool)`

GetDirectoryUserSourceIdOk returns a tuple with the DirectoryUserSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryUserSourceId

`func (o *DirectoryUser) SetDirectoryUserSourceId(v int32)`

SetDirectoryUserSourceId sets DirectoryUserSourceId field to given value.

### HasDirectoryUserSourceId

`func (o *DirectoryUser) HasDirectoryUserSourceId() bool`

HasDirectoryUserSourceId returns a boolean if a field has been set.

### GetDirectoryUserCreatedAt

`func (o *DirectoryUser) GetDirectoryUserCreatedAt() int64`

GetDirectoryUserCreatedAt returns the DirectoryUserCreatedAt field if non-nil, zero value otherwise.

### GetDirectoryUserCreatedAtOk

`func (o *DirectoryUser) GetDirectoryUserCreatedAtOk() (*int64, bool)`

GetDirectoryUserCreatedAtOk returns a tuple with the DirectoryUserCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryUserCreatedAt

`func (o *DirectoryUser) SetDirectoryUserCreatedAt(v int64)`

SetDirectoryUserCreatedAt sets DirectoryUserCreatedAt field to given value.

### HasDirectoryUserCreatedAt

`func (o *DirectoryUser) HasDirectoryUserCreatedAt() bool`

HasDirectoryUserCreatedAt returns a boolean if a field has been set.

### GetDirectoryUserUpdatedAt

`func (o *DirectoryUser) GetDirectoryUserUpdatedAt() int64`

GetDirectoryUserUpdatedAt returns the DirectoryUserUpdatedAt field if non-nil, zero value otherwise.

### GetDirectoryUserUpdatedAtOk

`func (o *DirectoryUser) GetDirectoryUserUpdatedAtOk() (*int64, bool)`

GetDirectoryUserUpdatedAtOk returns a tuple with the DirectoryUserUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryUserUpdatedAt

`func (o *DirectoryUser) SetDirectoryUserUpdatedAt(v int64)`

SetDirectoryUserUpdatedAt sets DirectoryUserUpdatedAt field to given value.

### HasDirectoryUserUpdatedAt

`func (o *DirectoryUser) HasDirectoryUserUpdatedAt() bool`

HasDirectoryUserUpdatedAt returns a boolean if a field has been set.

### GetDirectoryUserUserId

`func (o *DirectoryUser) GetDirectoryUserUserId() string`

GetDirectoryUserUserId returns the DirectoryUserUserId field if non-nil, zero value otherwise.

### GetDirectoryUserUserIdOk

`func (o *DirectoryUser) GetDirectoryUserUserIdOk() (*string, bool)`

GetDirectoryUserUserIdOk returns a tuple with the DirectoryUserUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryUserUserId

`func (o *DirectoryUser) SetDirectoryUserUserId(v string)`

SetDirectoryUserUserId sets DirectoryUserUserId field to given value.

### HasDirectoryUserUserId

`func (o *DirectoryUser) HasDirectoryUserUserId() bool`

HasDirectoryUserUserId returns a boolean if a field has been set.

### GetDirectoryUserGroupIds

`func (o *DirectoryUser) GetDirectoryUserGroupIds() []string`

GetDirectoryUserGroupIds returns the DirectoryUserGroupIds field if non-nil, zero value otherwise.

### GetDirectoryUserGroupIdsOk

`func (o *DirectoryUser) GetDirectoryUserGroupIdsOk() (*[]string, bool)`

GetDirectoryUserGroupIdsOk returns a tuple with the DirectoryUserGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryUserGroupIds

`func (o *DirectoryUser) SetDirectoryUserGroupIds(v []string)`

SetDirectoryUserGroupIds sets DirectoryUserGroupIds field to given value.

### HasDirectoryUserGroupIds

`func (o *DirectoryUser) HasDirectoryUserGroupIds() bool`

HasDirectoryUserGroupIds returns a boolean if a field has been set.

### GetDirectoryUserDisplayName

`func (o *DirectoryUser) GetDirectoryUserDisplayName() string`

GetDirectoryUserDisplayName returns the DirectoryUserDisplayName field if non-nil, zero value otherwise.

### GetDirectoryUserDisplayNameOk

`func (o *DirectoryUser) GetDirectoryUserDisplayNameOk() (*string, bool)`

GetDirectoryUserDisplayNameOk returns a tuple with the DirectoryUserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryUserDisplayName

`func (o *DirectoryUser) SetDirectoryUserDisplayName(v string)`

SetDirectoryUserDisplayName sets DirectoryUserDisplayName field to given value.

### HasDirectoryUserDisplayName

`func (o *DirectoryUser) HasDirectoryUserDisplayName() bool`

HasDirectoryUserDisplayName returns a boolean if a field has been set.

### GetDirectoryUserName

`func (o *DirectoryUser) GetDirectoryUserName() string`

GetDirectoryUserName returns the DirectoryUserName field if non-nil, zero value otherwise.

### GetDirectoryUserNameOk

`func (o *DirectoryUser) GetDirectoryUserNameOk() (*string, bool)`

GetDirectoryUserNameOk returns a tuple with the DirectoryUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryUserName

`func (o *DirectoryUser) SetDirectoryUserName(v string)`

SetDirectoryUserName sets DirectoryUserName field to given value.

### HasDirectoryUserName

`func (o *DirectoryUser) HasDirectoryUserName() bool`

HasDirectoryUserName returns a boolean if a field has been set.

### GetDirectoryUserFirstName

`func (o *DirectoryUser) GetDirectoryUserFirstName() string`

GetDirectoryUserFirstName returns the DirectoryUserFirstName field if non-nil, zero value otherwise.

### GetDirectoryUserFirstNameOk

`func (o *DirectoryUser) GetDirectoryUserFirstNameOk() (*string, bool)`

GetDirectoryUserFirstNameOk returns a tuple with the DirectoryUserFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryUserFirstName

`func (o *DirectoryUser) SetDirectoryUserFirstName(v string)`

SetDirectoryUserFirstName sets DirectoryUserFirstName field to given value.

### HasDirectoryUserFirstName

`func (o *DirectoryUser) HasDirectoryUserFirstName() bool`

HasDirectoryUserFirstName returns a boolean if a field has been set.

### GetDirectoryUserLastName

`func (o *DirectoryUser) GetDirectoryUserLastName() string`

GetDirectoryUserLastName returns the DirectoryUserLastName field if non-nil, zero value otherwise.

### GetDirectoryUserLastNameOk

`func (o *DirectoryUser) GetDirectoryUserLastNameOk() (*string, bool)`

GetDirectoryUserLastNameOk returns a tuple with the DirectoryUserLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryUserLastName

`func (o *DirectoryUser) SetDirectoryUserLastName(v string)`

SetDirectoryUserLastName sets DirectoryUserLastName field to given value.

### HasDirectoryUserLastName

`func (o *DirectoryUser) HasDirectoryUserLastName() bool`

HasDirectoryUserLastName returns a boolean if a field has been set.

### GetDirectoryUserDescription

`func (o *DirectoryUser) GetDirectoryUserDescription() string`

GetDirectoryUserDescription returns the DirectoryUserDescription field if non-nil, zero value otherwise.

### GetDirectoryUserDescriptionOk

`func (o *DirectoryUser) GetDirectoryUserDescriptionOk() (*string, bool)`

GetDirectoryUserDescriptionOk returns a tuple with the DirectoryUserDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryUserDescription

`func (o *DirectoryUser) SetDirectoryUserDescription(v string)`

SetDirectoryUserDescription sets DirectoryUserDescription field to given value.

### HasDirectoryUserDescription

`func (o *DirectoryUser) HasDirectoryUserDescription() bool`

HasDirectoryUserDescription returns a boolean if a field has been set.

### GetDirectoryUserEmail

`func (o *DirectoryUser) GetDirectoryUserEmail() string`

GetDirectoryUserEmail returns the DirectoryUserEmail field if non-nil, zero value otherwise.

### GetDirectoryUserEmailOk

`func (o *DirectoryUser) GetDirectoryUserEmailOk() (*string, bool)`

GetDirectoryUserEmailOk returns a tuple with the DirectoryUserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryUserEmail

`func (o *DirectoryUser) SetDirectoryUserEmail(v string)`

SetDirectoryUserEmail sets DirectoryUserEmail field to given value.

### HasDirectoryUserEmail

`func (o *DirectoryUser) HasDirectoryUserEmail() bool`

HasDirectoryUserEmail returns a boolean if a field has been set.

### GetDirectoryUserEmailNormalized

`func (o *DirectoryUser) GetDirectoryUserEmailNormalized() string`

GetDirectoryUserEmailNormalized returns the DirectoryUserEmailNormalized field if non-nil, zero value otherwise.

### GetDirectoryUserEmailNormalizedOk

`func (o *DirectoryUser) GetDirectoryUserEmailNormalizedOk() (*string, bool)`

GetDirectoryUserEmailNormalizedOk returns a tuple with the DirectoryUserEmailNormalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryUserEmailNormalized

`func (o *DirectoryUser) SetDirectoryUserEmailNormalized(v string)`

SetDirectoryUserEmailNormalized sets DirectoryUserEmailNormalized field to given value.

### HasDirectoryUserEmailNormalized

`func (o *DirectoryUser) HasDirectoryUserEmailNormalized() bool`

HasDirectoryUserEmailNormalized returns a boolean if a field has been set.

### GetDirectoryUserPhone

`func (o *DirectoryUser) GetDirectoryUserPhone() string`

GetDirectoryUserPhone returns the DirectoryUserPhone field if non-nil, zero value otherwise.

### GetDirectoryUserPhoneOk

`func (o *DirectoryUser) GetDirectoryUserPhoneOk() (*string, bool)`

GetDirectoryUserPhoneOk returns a tuple with the DirectoryUserPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryUserPhone

`func (o *DirectoryUser) SetDirectoryUserPhone(v string)`

SetDirectoryUserPhone sets DirectoryUserPhone field to given value.

### HasDirectoryUserPhone

`func (o *DirectoryUser) HasDirectoryUserPhone() bool`

HasDirectoryUserPhone returns a boolean if a field has been set.

### GetDirectoryUserTitle

`func (o *DirectoryUser) GetDirectoryUserTitle() string`

GetDirectoryUserTitle returns the DirectoryUserTitle field if non-nil, zero value otherwise.

### GetDirectoryUserTitleOk

`func (o *DirectoryUser) GetDirectoryUserTitleOk() (*string, bool)`

GetDirectoryUserTitleOk returns a tuple with the DirectoryUserTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryUserTitle

`func (o *DirectoryUser) SetDirectoryUserTitle(v string)`

SetDirectoryUserTitle sets DirectoryUserTitle field to given value.

### HasDirectoryUserTitle

`func (o *DirectoryUser) HasDirectoryUserTitle() bool`

HasDirectoryUserTitle returns a boolean if a field has been set.

### GetDirectoryUserLocation

`func (o *DirectoryUser) GetDirectoryUserLocation() string`

GetDirectoryUserLocation returns the DirectoryUserLocation field if non-nil, zero value otherwise.

### GetDirectoryUserLocationOk

`func (o *DirectoryUser) GetDirectoryUserLocationOk() (*string, bool)`

GetDirectoryUserLocationOk returns a tuple with the DirectoryUserLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryUserLocation

`func (o *DirectoryUser) SetDirectoryUserLocation(v string)`

SetDirectoryUserLocation sets DirectoryUserLocation field to given value.

### HasDirectoryUserLocation

`func (o *DirectoryUser) HasDirectoryUserLocation() bool`

HasDirectoryUserLocation returns a boolean if a field has been set.

### GetDirectoryUserLastLogonAt

`func (o *DirectoryUser) GetDirectoryUserLastLogonAt() int64`

GetDirectoryUserLastLogonAt returns the DirectoryUserLastLogonAt field if non-nil, zero value otherwise.

### GetDirectoryUserLastLogonAtOk

`func (o *DirectoryUser) GetDirectoryUserLastLogonAtOk() (*int64, bool)`

GetDirectoryUserLastLogonAtOk returns a tuple with the DirectoryUserLastLogonAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryUserLastLogonAt

`func (o *DirectoryUser) SetDirectoryUserLastLogonAt(v int64)`

SetDirectoryUserLastLogonAt sets DirectoryUserLastLogonAt field to given value.

### HasDirectoryUserLastLogonAt

`func (o *DirectoryUser) HasDirectoryUserLastLogonAt() bool`

HasDirectoryUserLastLogonAt returns a boolean if a field has been set.

### GetDirectoryUserAttributes

`func (o *DirectoryUser) GetDirectoryUserAttributes() map[string]string`

GetDirectoryUserAttributes returns the DirectoryUserAttributes field if non-nil, zero value otherwise.

### GetDirectoryUserAttributesOk

`func (o *DirectoryUser) GetDirectoryUserAttributesOk() (*map[string]string, bool)`

GetDirectoryUserAttributesOk returns a tuple with the DirectoryUserAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryUserAttributes

`func (o *DirectoryUser) SetDirectoryUserAttributes(v map[string]string)`

SetDirectoryUserAttributes sets DirectoryUserAttributes field to given value.

### HasDirectoryUserAttributes

`func (o *DirectoryUser) HasDirectoryUserAttributes() bool`

HasDirectoryUserAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


