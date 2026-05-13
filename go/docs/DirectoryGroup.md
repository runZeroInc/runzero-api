# DirectoryGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**DirectoryGroupId** | Pointer to **string** |  | [optional] 
**DirectoryGroupOrganizationId** | Pointer to **string** |  | [optional] 
**DirectoryGroupSiteId** | Pointer to **string** |  | [optional] 
**DirectoryGroupSourceId** | Pointer to **int32** |  | [optional] 
**DirectoryGroupCreatedAt** | Pointer to **int64** |  | [optional] 
**DirectoryGroupUpdatedAt** | Pointer to **int64** |  | [optional] 
**DirectoryGroupGroupId** | Pointer to **string** |  | [optional] 
**DirectoryGroupUserCount** | Pointer to **int64** |  | [optional] 
**DirectoryGroupName** | Pointer to **string** |  | [optional] 
**DirectoryGroupDisplayName** | Pointer to **string** |  | [optional] 
**DirectoryGroupDescription** | Pointer to **string** |  | [optional] 
**DirectoryGroupEmail** | Pointer to **string** |  | [optional] 
**DirectoryGroupEmailNormalized** | Pointer to **string** |  | [optional] 
**DirectoryGroupAttributes** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewDirectoryGroup

`func NewDirectoryGroup(id string, ) *DirectoryGroup`

NewDirectoryGroup instantiates a new DirectoryGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryGroupWithDefaults

`func NewDirectoryGroupWithDefaults() *DirectoryGroup`

NewDirectoryGroupWithDefaults instantiates a new DirectoryGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DirectoryGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DirectoryGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DirectoryGroup) SetId(v string)`

SetId sets Id field to given value.


### GetDirectoryGroupId

`func (o *DirectoryGroup) GetDirectoryGroupId() string`

GetDirectoryGroupId returns the DirectoryGroupId field if non-nil, zero value otherwise.

### GetDirectoryGroupIdOk

`func (o *DirectoryGroup) GetDirectoryGroupIdOk() (*string, bool)`

GetDirectoryGroupIdOk returns a tuple with the DirectoryGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryGroupId

`func (o *DirectoryGroup) SetDirectoryGroupId(v string)`

SetDirectoryGroupId sets DirectoryGroupId field to given value.

### HasDirectoryGroupId

`func (o *DirectoryGroup) HasDirectoryGroupId() bool`

HasDirectoryGroupId returns a boolean if a field has been set.

### GetDirectoryGroupOrganizationId

`func (o *DirectoryGroup) GetDirectoryGroupOrganizationId() string`

GetDirectoryGroupOrganizationId returns the DirectoryGroupOrganizationId field if non-nil, zero value otherwise.

### GetDirectoryGroupOrganizationIdOk

`func (o *DirectoryGroup) GetDirectoryGroupOrganizationIdOk() (*string, bool)`

GetDirectoryGroupOrganizationIdOk returns a tuple with the DirectoryGroupOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryGroupOrganizationId

`func (o *DirectoryGroup) SetDirectoryGroupOrganizationId(v string)`

SetDirectoryGroupOrganizationId sets DirectoryGroupOrganizationId field to given value.

### HasDirectoryGroupOrganizationId

`func (o *DirectoryGroup) HasDirectoryGroupOrganizationId() bool`

HasDirectoryGroupOrganizationId returns a boolean if a field has been set.

### GetDirectoryGroupSiteId

`func (o *DirectoryGroup) GetDirectoryGroupSiteId() string`

GetDirectoryGroupSiteId returns the DirectoryGroupSiteId field if non-nil, zero value otherwise.

### GetDirectoryGroupSiteIdOk

`func (o *DirectoryGroup) GetDirectoryGroupSiteIdOk() (*string, bool)`

GetDirectoryGroupSiteIdOk returns a tuple with the DirectoryGroupSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryGroupSiteId

`func (o *DirectoryGroup) SetDirectoryGroupSiteId(v string)`

SetDirectoryGroupSiteId sets DirectoryGroupSiteId field to given value.

### HasDirectoryGroupSiteId

`func (o *DirectoryGroup) HasDirectoryGroupSiteId() bool`

HasDirectoryGroupSiteId returns a boolean if a field has been set.

### GetDirectoryGroupSourceId

`func (o *DirectoryGroup) GetDirectoryGroupSourceId() int32`

GetDirectoryGroupSourceId returns the DirectoryGroupSourceId field if non-nil, zero value otherwise.

### GetDirectoryGroupSourceIdOk

`func (o *DirectoryGroup) GetDirectoryGroupSourceIdOk() (*int32, bool)`

GetDirectoryGroupSourceIdOk returns a tuple with the DirectoryGroupSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryGroupSourceId

`func (o *DirectoryGroup) SetDirectoryGroupSourceId(v int32)`

SetDirectoryGroupSourceId sets DirectoryGroupSourceId field to given value.

### HasDirectoryGroupSourceId

`func (o *DirectoryGroup) HasDirectoryGroupSourceId() bool`

HasDirectoryGroupSourceId returns a boolean if a field has been set.

### GetDirectoryGroupCreatedAt

`func (o *DirectoryGroup) GetDirectoryGroupCreatedAt() int64`

GetDirectoryGroupCreatedAt returns the DirectoryGroupCreatedAt field if non-nil, zero value otherwise.

### GetDirectoryGroupCreatedAtOk

`func (o *DirectoryGroup) GetDirectoryGroupCreatedAtOk() (*int64, bool)`

GetDirectoryGroupCreatedAtOk returns a tuple with the DirectoryGroupCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryGroupCreatedAt

`func (o *DirectoryGroup) SetDirectoryGroupCreatedAt(v int64)`

SetDirectoryGroupCreatedAt sets DirectoryGroupCreatedAt field to given value.

### HasDirectoryGroupCreatedAt

`func (o *DirectoryGroup) HasDirectoryGroupCreatedAt() bool`

HasDirectoryGroupCreatedAt returns a boolean if a field has been set.

### GetDirectoryGroupUpdatedAt

`func (o *DirectoryGroup) GetDirectoryGroupUpdatedAt() int64`

GetDirectoryGroupUpdatedAt returns the DirectoryGroupUpdatedAt field if non-nil, zero value otherwise.

### GetDirectoryGroupUpdatedAtOk

`func (o *DirectoryGroup) GetDirectoryGroupUpdatedAtOk() (*int64, bool)`

GetDirectoryGroupUpdatedAtOk returns a tuple with the DirectoryGroupUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryGroupUpdatedAt

`func (o *DirectoryGroup) SetDirectoryGroupUpdatedAt(v int64)`

SetDirectoryGroupUpdatedAt sets DirectoryGroupUpdatedAt field to given value.

### HasDirectoryGroupUpdatedAt

`func (o *DirectoryGroup) HasDirectoryGroupUpdatedAt() bool`

HasDirectoryGroupUpdatedAt returns a boolean if a field has been set.

### GetDirectoryGroupGroupId

`func (o *DirectoryGroup) GetDirectoryGroupGroupId() string`

GetDirectoryGroupGroupId returns the DirectoryGroupGroupId field if non-nil, zero value otherwise.

### GetDirectoryGroupGroupIdOk

`func (o *DirectoryGroup) GetDirectoryGroupGroupIdOk() (*string, bool)`

GetDirectoryGroupGroupIdOk returns a tuple with the DirectoryGroupGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryGroupGroupId

`func (o *DirectoryGroup) SetDirectoryGroupGroupId(v string)`

SetDirectoryGroupGroupId sets DirectoryGroupGroupId field to given value.

### HasDirectoryGroupGroupId

`func (o *DirectoryGroup) HasDirectoryGroupGroupId() bool`

HasDirectoryGroupGroupId returns a boolean if a field has been set.

### GetDirectoryGroupUserCount

`func (o *DirectoryGroup) GetDirectoryGroupUserCount() int64`

GetDirectoryGroupUserCount returns the DirectoryGroupUserCount field if non-nil, zero value otherwise.

### GetDirectoryGroupUserCountOk

`func (o *DirectoryGroup) GetDirectoryGroupUserCountOk() (*int64, bool)`

GetDirectoryGroupUserCountOk returns a tuple with the DirectoryGroupUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryGroupUserCount

`func (o *DirectoryGroup) SetDirectoryGroupUserCount(v int64)`

SetDirectoryGroupUserCount sets DirectoryGroupUserCount field to given value.

### HasDirectoryGroupUserCount

`func (o *DirectoryGroup) HasDirectoryGroupUserCount() bool`

HasDirectoryGroupUserCount returns a boolean if a field has been set.

### GetDirectoryGroupName

`func (o *DirectoryGroup) GetDirectoryGroupName() string`

GetDirectoryGroupName returns the DirectoryGroupName field if non-nil, zero value otherwise.

### GetDirectoryGroupNameOk

`func (o *DirectoryGroup) GetDirectoryGroupNameOk() (*string, bool)`

GetDirectoryGroupNameOk returns a tuple with the DirectoryGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryGroupName

`func (o *DirectoryGroup) SetDirectoryGroupName(v string)`

SetDirectoryGroupName sets DirectoryGroupName field to given value.

### HasDirectoryGroupName

`func (o *DirectoryGroup) HasDirectoryGroupName() bool`

HasDirectoryGroupName returns a boolean if a field has been set.

### GetDirectoryGroupDisplayName

`func (o *DirectoryGroup) GetDirectoryGroupDisplayName() string`

GetDirectoryGroupDisplayName returns the DirectoryGroupDisplayName field if non-nil, zero value otherwise.

### GetDirectoryGroupDisplayNameOk

`func (o *DirectoryGroup) GetDirectoryGroupDisplayNameOk() (*string, bool)`

GetDirectoryGroupDisplayNameOk returns a tuple with the DirectoryGroupDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryGroupDisplayName

`func (o *DirectoryGroup) SetDirectoryGroupDisplayName(v string)`

SetDirectoryGroupDisplayName sets DirectoryGroupDisplayName field to given value.

### HasDirectoryGroupDisplayName

`func (o *DirectoryGroup) HasDirectoryGroupDisplayName() bool`

HasDirectoryGroupDisplayName returns a boolean if a field has been set.

### GetDirectoryGroupDescription

`func (o *DirectoryGroup) GetDirectoryGroupDescription() string`

GetDirectoryGroupDescription returns the DirectoryGroupDescription field if non-nil, zero value otherwise.

### GetDirectoryGroupDescriptionOk

`func (o *DirectoryGroup) GetDirectoryGroupDescriptionOk() (*string, bool)`

GetDirectoryGroupDescriptionOk returns a tuple with the DirectoryGroupDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryGroupDescription

`func (o *DirectoryGroup) SetDirectoryGroupDescription(v string)`

SetDirectoryGroupDescription sets DirectoryGroupDescription field to given value.

### HasDirectoryGroupDescription

`func (o *DirectoryGroup) HasDirectoryGroupDescription() bool`

HasDirectoryGroupDescription returns a boolean if a field has been set.

### GetDirectoryGroupEmail

`func (o *DirectoryGroup) GetDirectoryGroupEmail() string`

GetDirectoryGroupEmail returns the DirectoryGroupEmail field if non-nil, zero value otherwise.

### GetDirectoryGroupEmailOk

`func (o *DirectoryGroup) GetDirectoryGroupEmailOk() (*string, bool)`

GetDirectoryGroupEmailOk returns a tuple with the DirectoryGroupEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryGroupEmail

`func (o *DirectoryGroup) SetDirectoryGroupEmail(v string)`

SetDirectoryGroupEmail sets DirectoryGroupEmail field to given value.

### HasDirectoryGroupEmail

`func (o *DirectoryGroup) HasDirectoryGroupEmail() bool`

HasDirectoryGroupEmail returns a boolean if a field has been set.

### GetDirectoryGroupEmailNormalized

`func (o *DirectoryGroup) GetDirectoryGroupEmailNormalized() string`

GetDirectoryGroupEmailNormalized returns the DirectoryGroupEmailNormalized field if non-nil, zero value otherwise.

### GetDirectoryGroupEmailNormalizedOk

`func (o *DirectoryGroup) GetDirectoryGroupEmailNormalizedOk() (*string, bool)`

GetDirectoryGroupEmailNormalizedOk returns a tuple with the DirectoryGroupEmailNormalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryGroupEmailNormalized

`func (o *DirectoryGroup) SetDirectoryGroupEmailNormalized(v string)`

SetDirectoryGroupEmailNormalized sets DirectoryGroupEmailNormalized field to given value.

### HasDirectoryGroupEmailNormalized

`func (o *DirectoryGroup) HasDirectoryGroupEmailNormalized() bool`

HasDirectoryGroupEmailNormalized returns a boolean if a field has been set.

### GetDirectoryGroupAttributes

`func (o *DirectoryGroup) GetDirectoryGroupAttributes() map[string]string`

GetDirectoryGroupAttributes returns the DirectoryGroupAttributes field if non-nil, zero value otherwise.

### GetDirectoryGroupAttributesOk

`func (o *DirectoryGroup) GetDirectoryGroupAttributesOk() (*map[string]string, bool)`

GetDirectoryGroupAttributesOk returns a tuple with the DirectoryGroupAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryGroupAttributes

`func (o *DirectoryGroup) SetDirectoryGroupAttributes(v map[string]string)`

SetDirectoryGroupAttributes sets DirectoryGroupAttributes field to given value.

### HasDirectoryGroupAttributes

`func (o *DirectoryGroup) HasDirectoryGroupAttributes() bool`

HasDirectoryGroupAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


