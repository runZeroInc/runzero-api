# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**SourceId** | Pointer to **string** |  | [optional] 
**SourceName** | Pointer to **string** |  | [optional] 
**SourceType** | Pointer to **string** |  | [optional] 
**TargetId** | Pointer to **string** |  | [optional] 
**TargetName** | Pointer to **string** |  | [optional] 
**TargetType** | Pointer to **string** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Details** | Pointer to **map[string]interface{}** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**ProcessorId** | Pointer to **string** |  | [optional] 
**ProcessedAt** | Pointer to **int64** |  | [optional] 

## Methods

### NewEvent

`func NewEvent() *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Event) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Event) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Event) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Event) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Event) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Event) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Event) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Event) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetClientId

`func (o *Event) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Event) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Event) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Event) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Event) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Event) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Event) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Event) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetSiteId

`func (o *Event) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Event) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Event) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *Event) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetAction

`func (o *Event) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Event) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Event) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *Event) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetSourceId

`func (o *Event) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *Event) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *Event) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *Event) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceName

`func (o *Event) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *Event) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *Event) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *Event) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### GetSourceType

`func (o *Event) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *Event) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *Event) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *Event) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetTargetId

`func (o *Event) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *Event) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *Event) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *Event) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetTargetName

`func (o *Event) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *Event) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *Event) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *Event) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetTargetType

`func (o *Event) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *Event) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *Event) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *Event) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetSuccess

`func (o *Event) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *Event) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *Event) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *Event) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetDetails

`func (o *Event) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Event) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Event) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Event) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetState

`func (o *Event) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Event) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Event) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Event) HasState() bool`

HasState returns a boolean if a field has been set.

### GetProcessorId

`func (o *Event) GetProcessorId() string`

GetProcessorId returns the ProcessorId field if non-nil, zero value otherwise.

### GetProcessorIdOk

`func (o *Event) GetProcessorIdOk() (*string, bool)`

GetProcessorIdOk returns a tuple with the ProcessorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorId

`func (o *Event) SetProcessorId(v string)`

SetProcessorId sets ProcessorId field to given value.

### HasProcessorId

`func (o *Event) HasProcessorId() bool`

HasProcessorId returns a boolean if a field has been set.

### GetProcessedAt

`func (o *Event) GetProcessedAt() int64`

GetProcessedAt returns the ProcessedAt field if non-nil, zero value otherwise.

### GetProcessedAtOk

`func (o *Event) GetProcessedAtOk() (*int64, bool)`

GetProcessedAtOk returns a tuple with the ProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedAt

`func (o *Event) SetProcessedAt(v int64)`

SetProcessedAt sets ProcessedAt field to given value.

### HasProcessedAt

`func (o *Event) HasProcessedAt() bool`

HasProcessedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


