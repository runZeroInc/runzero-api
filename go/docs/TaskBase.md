# TaskBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**TemplateId** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**AgentId** | Pointer to **string** |  | [optional] 
**HostedZoneId** | Pointer to **string** | The ID of the Hosted Zone which executes the task. If the  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**CruncherId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedByUserId** | Pointer to **string** |  | [optional] 
**CustomIntegrationId** | Pointer to **string** | The ID of the custom integration source, if the last task executed with this template was an import of Asset Data. | [optional] 
**SourceId** | Pointer to **int32** | The numeric ID of the data source, if the task executed with this template is a runZero scan or third party data connection import. | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Params** | Pointer to **map[string]string** |  | [optional] 
**Stats** | Pointer to **map[string]interface{}** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**Recur** | Pointer to **bool** |  | [optional] 
**RecurFrequency** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **int64** |  | [optional] 
**RecurLast** | Pointer to **int64** |  | [optional] 
**RecurNext** | Pointer to **int64** |  | [optional] 
**RecurLastTaskId** | Pointer to **string** |  | [optional] 

## Methods

### NewTaskBase

`func NewTaskBase() *TaskBase`

NewTaskBase instantiates a new TaskBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskBaseWithDefaults

`func NewTaskBaseWithDefaults() *TaskBase`

NewTaskBaseWithDefaults instantiates a new TaskBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskBase) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaskBase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TaskBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaskBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *TaskBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaskBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaskBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaskBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTemplateId

`func (o *TaskBase) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *TaskBase) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *TaskBase) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *TaskBase) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetClientId

`func (o *TaskBase) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TaskBase) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TaskBase) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *TaskBase) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *TaskBase) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *TaskBase) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *TaskBase) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *TaskBase) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetAgentId

`func (o *TaskBase) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *TaskBase) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *TaskBase) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *TaskBase) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetHostedZoneId

`func (o *TaskBase) GetHostedZoneId() string`

GetHostedZoneId returns the HostedZoneId field if non-nil, zero value otherwise.

### GetHostedZoneIdOk

`func (o *TaskBase) GetHostedZoneIdOk() (*string, bool)`

GetHostedZoneIdOk returns a tuple with the HostedZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedZoneId

`func (o *TaskBase) SetHostedZoneId(v string)`

SetHostedZoneId sets HostedZoneId field to given value.

### HasHostedZoneId

`func (o *TaskBase) HasHostedZoneId() bool`

HasHostedZoneId returns a boolean if a field has been set.

### GetSiteId

`func (o *TaskBase) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *TaskBase) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *TaskBase) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *TaskBase) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetCruncherId

`func (o *TaskBase) GetCruncherId() string`

GetCruncherId returns the CruncherId field if non-nil, zero value otherwise.

### GetCruncherIdOk

`func (o *TaskBase) GetCruncherIdOk() (*string, bool)`

GetCruncherIdOk returns a tuple with the CruncherId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCruncherId

`func (o *TaskBase) SetCruncherId(v string)`

SetCruncherId sets CruncherId field to given value.

### HasCruncherId

`func (o *TaskBase) HasCruncherId() bool`

HasCruncherId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TaskBase) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskBase) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskBase) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TaskBase) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *TaskBase) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TaskBase) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TaskBase) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *TaskBase) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *TaskBase) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *TaskBase) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *TaskBase) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *TaskBase) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetCustomIntegrationId

`func (o *TaskBase) GetCustomIntegrationId() string`

GetCustomIntegrationId returns the CustomIntegrationId field if non-nil, zero value otherwise.

### GetCustomIntegrationIdOk

`func (o *TaskBase) GetCustomIntegrationIdOk() (*string, bool)`

GetCustomIntegrationIdOk returns a tuple with the CustomIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomIntegrationId

`func (o *TaskBase) SetCustomIntegrationId(v string)`

SetCustomIntegrationId sets CustomIntegrationId field to given value.

### HasCustomIntegrationId

`func (o *TaskBase) HasCustomIntegrationId() bool`

HasCustomIntegrationId returns a boolean if a field has been set.

### GetSourceId

`func (o *TaskBase) GetSourceId() int32`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *TaskBase) GetSourceIdOk() (*int32, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *TaskBase) SetSourceId(v int32)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *TaskBase) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TaskBase) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TaskBase) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TaskBase) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TaskBase) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetType

`func (o *TaskBase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaskBase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaskBase) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TaskBase) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *TaskBase) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskBase) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskBase) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaskBase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetError

`func (o *TaskBase) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TaskBase) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TaskBase) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *TaskBase) HasError() bool`

HasError returns a boolean if a field has been set.

### GetParams

`func (o *TaskBase) GetParams() map[string]string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *TaskBase) GetParamsOk() (*map[string]string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *TaskBase) SetParams(v map[string]string)`

SetParams sets Params field to given value.

### HasParams

`func (o *TaskBase) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetStats

`func (o *TaskBase) GetStats() map[string]interface{}`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *TaskBase) GetStatsOk() (*map[string]interface{}, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *TaskBase) SetStats(v map[string]interface{})`

SetStats sets Stats field to given value.

### HasStats

`func (o *TaskBase) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetHidden

`func (o *TaskBase) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *TaskBase) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *TaskBase) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *TaskBase) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetParentId

`func (o *TaskBase) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *TaskBase) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *TaskBase) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *TaskBase) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetRecur

`func (o *TaskBase) GetRecur() bool`

GetRecur returns the Recur field if non-nil, zero value otherwise.

### GetRecurOk

`func (o *TaskBase) GetRecurOk() (*bool, bool)`

GetRecurOk returns a tuple with the Recur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecur

`func (o *TaskBase) SetRecur(v bool)`

SetRecur sets Recur field to given value.

### HasRecur

`func (o *TaskBase) HasRecur() bool`

HasRecur returns a boolean if a field has been set.

### GetRecurFrequency

`func (o *TaskBase) GetRecurFrequency() string`

GetRecurFrequency returns the RecurFrequency field if non-nil, zero value otherwise.

### GetRecurFrequencyOk

`func (o *TaskBase) GetRecurFrequencyOk() (*string, bool)`

GetRecurFrequencyOk returns a tuple with the RecurFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurFrequency

`func (o *TaskBase) SetRecurFrequency(v string)`

SetRecurFrequency sets RecurFrequency field to given value.

### HasRecurFrequency

`func (o *TaskBase) HasRecurFrequency() bool`

HasRecurFrequency returns a boolean if a field has been set.

### GetStartTime

`func (o *TaskBase) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TaskBase) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TaskBase) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TaskBase) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetRecurLast

`func (o *TaskBase) GetRecurLast() int64`

GetRecurLast returns the RecurLast field if non-nil, zero value otherwise.

### GetRecurLastOk

`func (o *TaskBase) GetRecurLastOk() (*int64, bool)`

GetRecurLastOk returns a tuple with the RecurLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurLast

`func (o *TaskBase) SetRecurLast(v int64)`

SetRecurLast sets RecurLast field to given value.

### HasRecurLast

`func (o *TaskBase) HasRecurLast() bool`

HasRecurLast returns a boolean if a field has been set.

### GetRecurNext

`func (o *TaskBase) GetRecurNext() int64`

GetRecurNext returns the RecurNext field if non-nil, zero value otherwise.

### GetRecurNextOk

`func (o *TaskBase) GetRecurNextOk() (*int64, bool)`

GetRecurNextOk returns a tuple with the RecurNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurNext

`func (o *TaskBase) SetRecurNext(v int64)`

SetRecurNext sets RecurNext field to given value.

### HasRecurNext

`func (o *TaskBase) HasRecurNext() bool`

HasRecurNext returns a boolean if a field has been set.

### GetRecurLastTaskId

`func (o *TaskBase) GetRecurLastTaskId() string`

GetRecurLastTaskId returns the RecurLastTaskId field if non-nil, zero value otherwise.

### GetRecurLastTaskIdOk

`func (o *TaskBase) GetRecurLastTaskIdOk() (*string, bool)`

GetRecurLastTaskIdOk returns a tuple with the RecurLastTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurLastTaskId

`func (o *TaskBase) SetRecurLastTaskId(v string)`

SetRecurLastTaskId sets RecurLastTaskId field to given value.

### HasRecurLastTaskId

`func (o *TaskBase) HasRecurLastTaskId() bool`

HasRecurLastTaskId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


