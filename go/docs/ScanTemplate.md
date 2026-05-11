# ScanTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the template. | 
**Name** | Pointer to **string** | The name of the template. | [optional] 
**Description** | Pointer to **string** | The description of the template. | [optional] 
**ClientId** | Pointer to **string** | ID of the account which owns the template. | [optional] 
**OrganizationId** | **string** | ID of the organization the template is available in. | 
**AgentId** | Pointer to **string** | ID of the explorer which may execute the template. | [optional] 
**SiteId** | Pointer to **string** | ID of the site the template is being used in. | [optional] 
**CruncherId** | Pointer to **string** | ID of the runZero cruncher the task is executing on. | [optional] 
**CreatedAt** | Pointer to **int64** | Unix timestamp value indicating when the template was created. | [optional] 
**CreatedBy** | Pointer to **string** | The username of the account which created the template. | [optional] 
**CreatedByUserId** | Pointer to **string** | The ID of the account which created the template. | [optional] 
**UpdatedAt** | Pointer to **int64** | Unix timestamp value indicating when the template was last modified. | [optional] 
**Type** | Pointer to **string** | The type of task the template creates. | [optional] 
**Status** | Pointer to **string** | The status of the last task using the template. | [optional] 
**Error** | Pointer to **string** | The error message, if any, of the last task using the template. | [optional] 
**Params** | Pointer to **map[string]string** | A number of task parameter values. Currently there is no authoritative list of in-use values. See existing templates for examples. | [optional] 
**Stats** | Pointer to **map[string]interface{}** | A map of statistics about the last task executed with the template. Currently there is no authoritative list of in-use values. See existing templates for examples. | [optional] 
**Hidden** | Pointer to **bool** | A flag indicating whether the item is hidden from common view. | [optional] 
**ParentId** | Pointer to **string** | The ID of the parent entity of the task scheduled. | [optional] 
**Recur** | Pointer to **bool** | A flag representing whether derived tasks are scheduled to repeat. | [optional] 
**RecurFrequency** | Pointer to **string** | A string time duration value representing execution frequency, if scheduled to repeat. You may use values including as once, hourly, daily, weekly, monthly, continuous  | [optional] 
**StartTime** | Pointer to **int64** | Unix timestamp representing the next execution time. | [optional] 
**RecurLast** | Pointer to **int64** | Unix timestamp representing the last execution if scheduled to repeat. | [optional] 
**RecurNext** | Pointer to **int64** | Unix timestamp representing the next execution if scheduled to repeat. | [optional] 
**RecurLastTaskId** | Pointer to **string** | The ID of the task that last executed if scheduled to repeat. | [optional] 
**GracePeriod** | Pointer to **string** | Additional time beyond hard expiration deadline by which the task may still be allowed to execute. | [optional] 
**CustomIntegrationId** | Pointer to **string** | The ID of the custom integration source, if the last task executed with this template was an import of Asset Data. | [optional] 
**SourceId** | Pointer to **string** | The numeric ID of the data source, if the task executed with this template is a runZero scan or third party data connection import. | [optional] 
**TemplateId** | Pointer to **string** | The ID of the template. | [optional] 
**SizeSite** | Pointer to **int64** | The size in assets of the site the last task the template was executed against. | [optional] 
**SizeData** | Pointer to **int64** | The total size of result data of the last task the template was used with. | [optional] 
**SizeResults** | Pointer to **int64** | The number of results in the last task the template was used with. | [optional] 
**HostedZoneId** | Pointer to **string** | The ID of the hosted zone that ran the last task the template was used with. | [optional] 
**LinkedTaskCount** | Pointer to **int32** | The number of tasks derived from the template. | [optional] 
**Global** | **bool** | Whether the template is globally available to all organizations. | 
**Acl** | **map[string]interface{}** | A map of IDs to strings which describe how the template may be accessed. Currently there is no authoritative list of in-use values. See existing templates for examples. | 

## Methods

### NewScanTemplate

`func NewScanTemplate(id string, organizationId string, global bool, acl map[string]interface{}, ) *ScanTemplate`

NewScanTemplate instantiates a new ScanTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanTemplateWithDefaults

`func NewScanTemplateWithDefaults() *ScanTemplate`

NewScanTemplateWithDefaults instantiates a new ScanTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScanTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScanTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScanTemplate) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ScanTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScanTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScanTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScanTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ScanTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScanTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScanTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScanTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetClientId

`func (o *ScanTemplate) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ScanTemplate) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ScanTemplate) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ScanTemplate) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ScanTemplate) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ScanTemplate) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ScanTemplate) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetAgentId

`func (o *ScanTemplate) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *ScanTemplate) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *ScanTemplate) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *ScanTemplate) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetSiteId

`func (o *ScanTemplate) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ScanTemplate) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ScanTemplate) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ScanTemplate) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetCruncherId

`func (o *ScanTemplate) GetCruncherId() string`

GetCruncherId returns the CruncherId field if non-nil, zero value otherwise.

### GetCruncherIdOk

`func (o *ScanTemplate) GetCruncherIdOk() (*string, bool)`

GetCruncherIdOk returns a tuple with the CruncherId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCruncherId

`func (o *ScanTemplate) SetCruncherId(v string)`

SetCruncherId sets CruncherId field to given value.

### HasCruncherId

`func (o *ScanTemplate) HasCruncherId() bool`

HasCruncherId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ScanTemplate) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ScanTemplate) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ScanTemplate) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ScanTemplate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ScanTemplate) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ScanTemplate) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ScanTemplate) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ScanTemplate) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *ScanTemplate) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *ScanTemplate) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *ScanTemplate) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *ScanTemplate) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ScanTemplate) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ScanTemplate) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ScanTemplate) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ScanTemplate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetType

`func (o *ScanTemplate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScanTemplate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScanTemplate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ScanTemplate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *ScanTemplate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScanTemplate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScanTemplate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ScanTemplate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetError

`func (o *ScanTemplate) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ScanTemplate) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ScanTemplate) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ScanTemplate) HasError() bool`

HasError returns a boolean if a field has been set.

### GetParams

`func (o *ScanTemplate) GetParams() map[string]string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ScanTemplate) GetParamsOk() (*map[string]string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ScanTemplate) SetParams(v map[string]string)`

SetParams sets Params field to given value.

### HasParams

`func (o *ScanTemplate) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetStats

`func (o *ScanTemplate) GetStats() map[string]interface{}`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ScanTemplate) GetStatsOk() (*map[string]interface{}, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ScanTemplate) SetStats(v map[string]interface{})`

SetStats sets Stats field to given value.

### HasStats

`func (o *ScanTemplate) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetHidden

`func (o *ScanTemplate) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *ScanTemplate) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *ScanTemplate) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *ScanTemplate) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetParentId

`func (o *ScanTemplate) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ScanTemplate) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ScanTemplate) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ScanTemplate) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetRecur

`func (o *ScanTemplate) GetRecur() bool`

GetRecur returns the Recur field if non-nil, zero value otherwise.

### GetRecurOk

`func (o *ScanTemplate) GetRecurOk() (*bool, bool)`

GetRecurOk returns a tuple with the Recur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecur

`func (o *ScanTemplate) SetRecur(v bool)`

SetRecur sets Recur field to given value.

### HasRecur

`func (o *ScanTemplate) HasRecur() bool`

HasRecur returns a boolean if a field has been set.

### GetRecurFrequency

`func (o *ScanTemplate) GetRecurFrequency() string`

GetRecurFrequency returns the RecurFrequency field if non-nil, zero value otherwise.

### GetRecurFrequencyOk

`func (o *ScanTemplate) GetRecurFrequencyOk() (*string, bool)`

GetRecurFrequencyOk returns a tuple with the RecurFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurFrequency

`func (o *ScanTemplate) SetRecurFrequency(v string)`

SetRecurFrequency sets RecurFrequency field to given value.

### HasRecurFrequency

`func (o *ScanTemplate) HasRecurFrequency() bool`

HasRecurFrequency returns a boolean if a field has been set.

### GetStartTime

`func (o *ScanTemplate) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ScanTemplate) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ScanTemplate) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ScanTemplate) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetRecurLast

`func (o *ScanTemplate) GetRecurLast() int64`

GetRecurLast returns the RecurLast field if non-nil, zero value otherwise.

### GetRecurLastOk

`func (o *ScanTemplate) GetRecurLastOk() (*int64, bool)`

GetRecurLastOk returns a tuple with the RecurLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurLast

`func (o *ScanTemplate) SetRecurLast(v int64)`

SetRecurLast sets RecurLast field to given value.

### HasRecurLast

`func (o *ScanTemplate) HasRecurLast() bool`

HasRecurLast returns a boolean if a field has been set.

### GetRecurNext

`func (o *ScanTemplate) GetRecurNext() int64`

GetRecurNext returns the RecurNext field if non-nil, zero value otherwise.

### GetRecurNextOk

`func (o *ScanTemplate) GetRecurNextOk() (*int64, bool)`

GetRecurNextOk returns a tuple with the RecurNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurNext

`func (o *ScanTemplate) SetRecurNext(v int64)`

SetRecurNext sets RecurNext field to given value.

### HasRecurNext

`func (o *ScanTemplate) HasRecurNext() bool`

HasRecurNext returns a boolean if a field has been set.

### GetRecurLastTaskId

`func (o *ScanTemplate) GetRecurLastTaskId() string`

GetRecurLastTaskId returns the RecurLastTaskId field if non-nil, zero value otherwise.

### GetRecurLastTaskIdOk

`func (o *ScanTemplate) GetRecurLastTaskIdOk() (*string, bool)`

GetRecurLastTaskIdOk returns a tuple with the RecurLastTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurLastTaskId

`func (o *ScanTemplate) SetRecurLastTaskId(v string)`

SetRecurLastTaskId sets RecurLastTaskId field to given value.

### HasRecurLastTaskId

`func (o *ScanTemplate) HasRecurLastTaskId() bool`

HasRecurLastTaskId returns a boolean if a field has been set.

### GetGracePeriod

`func (o *ScanTemplate) GetGracePeriod() string`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *ScanTemplate) GetGracePeriodOk() (*string, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *ScanTemplate) SetGracePeriod(v string)`

SetGracePeriod sets GracePeriod field to given value.

### HasGracePeriod

`func (o *ScanTemplate) HasGracePeriod() bool`

HasGracePeriod returns a boolean if a field has been set.

### GetCustomIntegrationId

`func (o *ScanTemplate) GetCustomIntegrationId() string`

GetCustomIntegrationId returns the CustomIntegrationId field if non-nil, zero value otherwise.

### GetCustomIntegrationIdOk

`func (o *ScanTemplate) GetCustomIntegrationIdOk() (*string, bool)`

GetCustomIntegrationIdOk returns a tuple with the CustomIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomIntegrationId

`func (o *ScanTemplate) SetCustomIntegrationId(v string)`

SetCustomIntegrationId sets CustomIntegrationId field to given value.

### HasCustomIntegrationId

`func (o *ScanTemplate) HasCustomIntegrationId() bool`

HasCustomIntegrationId returns a boolean if a field has been set.

### GetSourceId

`func (o *ScanTemplate) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ScanTemplate) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ScanTemplate) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *ScanTemplate) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetTemplateId

`func (o *ScanTemplate) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *ScanTemplate) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *ScanTemplate) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *ScanTemplate) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetSizeSite

`func (o *ScanTemplate) GetSizeSite() int64`

GetSizeSite returns the SizeSite field if non-nil, zero value otherwise.

### GetSizeSiteOk

`func (o *ScanTemplate) GetSizeSiteOk() (*int64, bool)`

GetSizeSiteOk returns a tuple with the SizeSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeSite

`func (o *ScanTemplate) SetSizeSite(v int64)`

SetSizeSite sets SizeSite field to given value.

### HasSizeSite

`func (o *ScanTemplate) HasSizeSite() bool`

HasSizeSite returns a boolean if a field has been set.

### GetSizeData

`func (o *ScanTemplate) GetSizeData() int64`

GetSizeData returns the SizeData field if non-nil, zero value otherwise.

### GetSizeDataOk

`func (o *ScanTemplate) GetSizeDataOk() (*int64, bool)`

GetSizeDataOk returns a tuple with the SizeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeData

`func (o *ScanTemplate) SetSizeData(v int64)`

SetSizeData sets SizeData field to given value.

### HasSizeData

`func (o *ScanTemplate) HasSizeData() bool`

HasSizeData returns a boolean if a field has been set.

### GetSizeResults

`func (o *ScanTemplate) GetSizeResults() int64`

GetSizeResults returns the SizeResults field if non-nil, zero value otherwise.

### GetSizeResultsOk

`func (o *ScanTemplate) GetSizeResultsOk() (*int64, bool)`

GetSizeResultsOk returns a tuple with the SizeResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeResults

`func (o *ScanTemplate) SetSizeResults(v int64)`

SetSizeResults sets SizeResults field to given value.

### HasSizeResults

`func (o *ScanTemplate) HasSizeResults() bool`

HasSizeResults returns a boolean if a field has been set.

### GetHostedZoneId

`func (o *ScanTemplate) GetHostedZoneId() string`

GetHostedZoneId returns the HostedZoneId field if non-nil, zero value otherwise.

### GetHostedZoneIdOk

`func (o *ScanTemplate) GetHostedZoneIdOk() (*string, bool)`

GetHostedZoneIdOk returns a tuple with the HostedZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedZoneId

`func (o *ScanTemplate) SetHostedZoneId(v string)`

SetHostedZoneId sets HostedZoneId field to given value.

### HasHostedZoneId

`func (o *ScanTemplate) HasHostedZoneId() bool`

HasHostedZoneId returns a boolean if a field has been set.

### GetLinkedTaskCount

`func (o *ScanTemplate) GetLinkedTaskCount() int32`

GetLinkedTaskCount returns the LinkedTaskCount field if non-nil, zero value otherwise.

### GetLinkedTaskCountOk

`func (o *ScanTemplate) GetLinkedTaskCountOk() (*int32, bool)`

GetLinkedTaskCountOk returns a tuple with the LinkedTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedTaskCount

`func (o *ScanTemplate) SetLinkedTaskCount(v int32)`

SetLinkedTaskCount sets LinkedTaskCount field to given value.

### HasLinkedTaskCount

`func (o *ScanTemplate) HasLinkedTaskCount() bool`

HasLinkedTaskCount returns a boolean if a field has been set.

### GetGlobal

`func (o *ScanTemplate) GetGlobal() bool`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *ScanTemplate) GetGlobalOk() (*bool, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *ScanTemplate) SetGlobal(v bool)`

SetGlobal sets Global field to given value.


### GetAcl

`func (o *ScanTemplate) GetAcl() map[string]interface{}`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *ScanTemplate) GetAclOk() (*map[string]interface{}, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *ScanTemplate) SetAcl(v map[string]interface{})`

SetAcl sets Acl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


