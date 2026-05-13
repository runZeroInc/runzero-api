# TaskOptions

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
**HostedZoneName** | Pointer to **string** | The string &#39;auto&#39; will use any available hosted zone. Otherwise, provide the string name (hostedzone1) of the hosted zone. | [optional] 

## Methods

### NewTaskOptions

`func NewTaskOptions() *TaskOptions`

NewTaskOptions instantiates a new TaskOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskOptionsWithDefaults

`func NewTaskOptionsWithDefaults() *TaskOptions`

NewTaskOptionsWithDefaults instantiates a new TaskOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskOptions) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskOptions) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskOptions) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaskOptions) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TaskOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskOptions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaskOptions) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *TaskOptions) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaskOptions) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaskOptions) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaskOptions) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTemplateId

`func (o *TaskOptions) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *TaskOptions) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *TaskOptions) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *TaskOptions) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetClientId

`func (o *TaskOptions) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TaskOptions) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TaskOptions) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *TaskOptions) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *TaskOptions) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *TaskOptions) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *TaskOptions) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *TaskOptions) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetAgentId

`func (o *TaskOptions) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *TaskOptions) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *TaskOptions) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *TaskOptions) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetHostedZoneId

`func (o *TaskOptions) GetHostedZoneId() string`

GetHostedZoneId returns the HostedZoneId field if non-nil, zero value otherwise.

### GetHostedZoneIdOk

`func (o *TaskOptions) GetHostedZoneIdOk() (*string, bool)`

GetHostedZoneIdOk returns a tuple with the HostedZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedZoneId

`func (o *TaskOptions) SetHostedZoneId(v string)`

SetHostedZoneId sets HostedZoneId field to given value.

### HasHostedZoneId

`func (o *TaskOptions) HasHostedZoneId() bool`

HasHostedZoneId returns a boolean if a field has been set.

### GetSiteId

`func (o *TaskOptions) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *TaskOptions) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *TaskOptions) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *TaskOptions) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetCruncherId

`func (o *TaskOptions) GetCruncherId() string`

GetCruncherId returns the CruncherId field if non-nil, zero value otherwise.

### GetCruncherIdOk

`func (o *TaskOptions) GetCruncherIdOk() (*string, bool)`

GetCruncherIdOk returns a tuple with the CruncherId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCruncherId

`func (o *TaskOptions) SetCruncherId(v string)`

SetCruncherId sets CruncherId field to given value.

### HasCruncherId

`func (o *TaskOptions) HasCruncherId() bool`

HasCruncherId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TaskOptions) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskOptions) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskOptions) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TaskOptions) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *TaskOptions) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TaskOptions) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TaskOptions) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *TaskOptions) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *TaskOptions) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *TaskOptions) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *TaskOptions) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *TaskOptions) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetCustomIntegrationId

`func (o *TaskOptions) GetCustomIntegrationId() string`

GetCustomIntegrationId returns the CustomIntegrationId field if non-nil, zero value otherwise.

### GetCustomIntegrationIdOk

`func (o *TaskOptions) GetCustomIntegrationIdOk() (*string, bool)`

GetCustomIntegrationIdOk returns a tuple with the CustomIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomIntegrationId

`func (o *TaskOptions) SetCustomIntegrationId(v string)`

SetCustomIntegrationId sets CustomIntegrationId field to given value.

### HasCustomIntegrationId

`func (o *TaskOptions) HasCustomIntegrationId() bool`

HasCustomIntegrationId returns a boolean if a field has been set.

### GetSourceId

`func (o *TaskOptions) GetSourceId() int32`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *TaskOptions) GetSourceIdOk() (*int32, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *TaskOptions) SetSourceId(v int32)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *TaskOptions) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TaskOptions) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TaskOptions) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TaskOptions) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TaskOptions) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetType

`func (o *TaskOptions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaskOptions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaskOptions) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TaskOptions) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *TaskOptions) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskOptions) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskOptions) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaskOptions) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetError

`func (o *TaskOptions) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TaskOptions) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TaskOptions) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *TaskOptions) HasError() bool`

HasError returns a boolean if a field has been set.

### GetParams

`func (o *TaskOptions) GetParams() map[string]string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *TaskOptions) GetParamsOk() (*map[string]string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *TaskOptions) SetParams(v map[string]string)`

SetParams sets Params field to given value.

### HasParams

`func (o *TaskOptions) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetStats

`func (o *TaskOptions) GetStats() map[string]interface{}`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *TaskOptions) GetStatsOk() (*map[string]interface{}, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *TaskOptions) SetStats(v map[string]interface{})`

SetStats sets Stats field to given value.

### HasStats

`func (o *TaskOptions) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetHidden

`func (o *TaskOptions) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *TaskOptions) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *TaskOptions) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *TaskOptions) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetParentId

`func (o *TaskOptions) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *TaskOptions) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *TaskOptions) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *TaskOptions) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetRecur

`func (o *TaskOptions) GetRecur() bool`

GetRecur returns the Recur field if non-nil, zero value otherwise.

### GetRecurOk

`func (o *TaskOptions) GetRecurOk() (*bool, bool)`

GetRecurOk returns a tuple with the Recur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecur

`func (o *TaskOptions) SetRecur(v bool)`

SetRecur sets Recur field to given value.

### HasRecur

`func (o *TaskOptions) HasRecur() bool`

HasRecur returns a boolean if a field has been set.

### GetRecurFrequency

`func (o *TaskOptions) GetRecurFrequency() string`

GetRecurFrequency returns the RecurFrequency field if non-nil, zero value otherwise.

### GetRecurFrequencyOk

`func (o *TaskOptions) GetRecurFrequencyOk() (*string, bool)`

GetRecurFrequencyOk returns a tuple with the RecurFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurFrequency

`func (o *TaskOptions) SetRecurFrequency(v string)`

SetRecurFrequency sets RecurFrequency field to given value.

### HasRecurFrequency

`func (o *TaskOptions) HasRecurFrequency() bool`

HasRecurFrequency returns a boolean if a field has been set.

### GetStartTime

`func (o *TaskOptions) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TaskOptions) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TaskOptions) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TaskOptions) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetRecurLast

`func (o *TaskOptions) GetRecurLast() int64`

GetRecurLast returns the RecurLast field if non-nil, zero value otherwise.

### GetRecurLastOk

`func (o *TaskOptions) GetRecurLastOk() (*int64, bool)`

GetRecurLastOk returns a tuple with the RecurLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurLast

`func (o *TaskOptions) SetRecurLast(v int64)`

SetRecurLast sets RecurLast field to given value.

### HasRecurLast

`func (o *TaskOptions) HasRecurLast() bool`

HasRecurLast returns a boolean if a field has been set.

### GetRecurNext

`func (o *TaskOptions) GetRecurNext() int64`

GetRecurNext returns the RecurNext field if non-nil, zero value otherwise.

### GetRecurNextOk

`func (o *TaskOptions) GetRecurNextOk() (*int64, bool)`

GetRecurNextOk returns a tuple with the RecurNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurNext

`func (o *TaskOptions) SetRecurNext(v int64)`

SetRecurNext sets RecurNext field to given value.

### HasRecurNext

`func (o *TaskOptions) HasRecurNext() bool`

HasRecurNext returns a boolean if a field has been set.

### GetRecurLastTaskId

`func (o *TaskOptions) GetRecurLastTaskId() string`

GetRecurLastTaskId returns the RecurLastTaskId field if non-nil, zero value otherwise.

### GetRecurLastTaskIdOk

`func (o *TaskOptions) GetRecurLastTaskIdOk() (*string, bool)`

GetRecurLastTaskIdOk returns a tuple with the RecurLastTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurLastTaskId

`func (o *TaskOptions) SetRecurLastTaskId(v string)`

SetRecurLastTaskId sets RecurLastTaskId field to given value.

### HasRecurLastTaskId

`func (o *TaskOptions) HasRecurLastTaskId() bool`

HasRecurLastTaskId returns a boolean if a field has been set.

### GetHostedZoneName

`func (o *TaskOptions) GetHostedZoneName() string`

GetHostedZoneName returns the HostedZoneName field if non-nil, zero value otherwise.

### GetHostedZoneNameOk

`func (o *TaskOptions) GetHostedZoneNameOk() (*string, bool)`

GetHostedZoneNameOk returns a tuple with the HostedZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedZoneName

`func (o *TaskOptions) SetHostedZoneName(v string)`

SetHostedZoneName sets HostedZoneName field to given value.

### HasHostedZoneName

`func (o *TaskOptions) HasHostedZoneName() bool`

HasHostedZoneName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


