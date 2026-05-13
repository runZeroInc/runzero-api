# HostedZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the hosted zone | 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** | Whether the hosted zone is enabled | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The last modification time of the hosted zone | [optional] 
**ProcessorId** | Pointer to **string** | The processor ID assigned to the hosted zone | [optional] 
**ExplorersConcurrency** | Pointer to **int64** | The number of concurrent explorer tasks that can be executed | [optional] 
**ExplorersTotal** | Pointer to **int64** | The number of explorers available in the zone | [optional] 
**TasksActive** | Pointer to **int64** | The number of tasks executing in the zone | [optional] 
**TasksWaiting** | Pointer to **int64** | The number of tasks waiting to execute in the zone | [optional] 
**OrganizationId** | Pointer to **string** | The ID of the organization the hosted zone is assigned to | [optional] 

## Methods

### NewHostedZone

`func NewHostedZone(id string, ) *HostedZone`

NewHostedZone instantiates a new HostedZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostedZoneWithDefaults

`func NewHostedZoneWithDefaults() *HostedZone`

NewHostedZoneWithDefaults instantiates a new HostedZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HostedZone) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostedZone) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostedZone) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *HostedZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HostedZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HostedZone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HostedZone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *HostedZone) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *HostedZone) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *HostedZone) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *HostedZone) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *HostedZone) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HostedZone) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HostedZone) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HostedZone) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetProcessorId

`func (o *HostedZone) GetProcessorId() string`

GetProcessorId returns the ProcessorId field if non-nil, zero value otherwise.

### GetProcessorIdOk

`func (o *HostedZone) GetProcessorIdOk() (*string, bool)`

GetProcessorIdOk returns a tuple with the ProcessorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorId

`func (o *HostedZone) SetProcessorId(v string)`

SetProcessorId sets ProcessorId field to given value.

### HasProcessorId

`func (o *HostedZone) HasProcessorId() bool`

HasProcessorId returns a boolean if a field has been set.

### GetExplorersConcurrency

`func (o *HostedZone) GetExplorersConcurrency() int64`

GetExplorersConcurrency returns the ExplorersConcurrency field if non-nil, zero value otherwise.

### GetExplorersConcurrencyOk

`func (o *HostedZone) GetExplorersConcurrencyOk() (*int64, bool)`

GetExplorersConcurrencyOk returns a tuple with the ExplorersConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplorersConcurrency

`func (o *HostedZone) SetExplorersConcurrency(v int64)`

SetExplorersConcurrency sets ExplorersConcurrency field to given value.

### HasExplorersConcurrency

`func (o *HostedZone) HasExplorersConcurrency() bool`

HasExplorersConcurrency returns a boolean if a field has been set.

### GetExplorersTotal

`func (o *HostedZone) GetExplorersTotal() int64`

GetExplorersTotal returns the ExplorersTotal field if non-nil, zero value otherwise.

### GetExplorersTotalOk

`func (o *HostedZone) GetExplorersTotalOk() (*int64, bool)`

GetExplorersTotalOk returns a tuple with the ExplorersTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplorersTotal

`func (o *HostedZone) SetExplorersTotal(v int64)`

SetExplorersTotal sets ExplorersTotal field to given value.

### HasExplorersTotal

`func (o *HostedZone) HasExplorersTotal() bool`

HasExplorersTotal returns a boolean if a field has been set.

### GetTasksActive

`func (o *HostedZone) GetTasksActive() int64`

GetTasksActive returns the TasksActive field if non-nil, zero value otherwise.

### GetTasksActiveOk

`func (o *HostedZone) GetTasksActiveOk() (*int64, bool)`

GetTasksActiveOk returns a tuple with the TasksActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksActive

`func (o *HostedZone) SetTasksActive(v int64)`

SetTasksActive sets TasksActive field to given value.

### HasTasksActive

`func (o *HostedZone) HasTasksActive() bool`

HasTasksActive returns a boolean if a field has been set.

### GetTasksWaiting

`func (o *HostedZone) GetTasksWaiting() int64`

GetTasksWaiting returns the TasksWaiting field if non-nil, zero value otherwise.

### GetTasksWaitingOk

`func (o *HostedZone) GetTasksWaitingOk() (*int64, bool)`

GetTasksWaitingOk returns a tuple with the TasksWaiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksWaiting

`func (o *HostedZone) SetTasksWaiting(v int64)`

SetTasksWaiting sets TasksWaiting field to given value.

### HasTasksWaiting

`func (o *HostedZone) HasTasksWaiting() bool`

HasTasksWaiting returns a boolean if a field has been set.

### GetOrganizationId

`func (o *HostedZone) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *HostedZone) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *HostedZone) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *HostedZone) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


