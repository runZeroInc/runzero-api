# Finding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Category** | Pointer to **string** |  | [optional] 
**OrganizationId** | **string** |  | 
**FindingCode** | Pointer to **string** | A runZero-assigned code for the finding. | [optional] 
**Name** | Pointer to **string** | The title of the finding. | [optional] 
**Description** | Pointer to **string** | Additional information about the finding. | [optional] 
**Solution** | Pointer to **string** | Remediation information. | [optional] 
**Links** | Pointer to [**FindingLinks**](FindingLinks.md) |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**LastDetectedAt** | Pointer to **int64** |  | [optional] 
**InstanceCount** | Pointer to **int64** |  | [optional] 
**RiskRank** | Pointer to **string** |  | [optional] 
**RiskRankValue** | Pointer to **int32** | 0 &#x3D; info, 4 &#x3D; critical | [optional] 

## Methods

### NewFinding

`func NewFinding(id string, organizationId string, ) *Finding`

NewFinding instantiates a new Finding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingWithDefaults

`func NewFindingWithDefaults() *Finding`

NewFindingWithDefaults instantiates a new Finding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Finding) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Finding) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Finding) SetId(v string)`

SetId sets Id field to given value.


### GetCategory

`func (o *Finding) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Finding) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Finding) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Finding) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Finding) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Finding) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Finding) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetFindingCode

`func (o *Finding) GetFindingCode() string`

GetFindingCode returns the FindingCode field if non-nil, zero value otherwise.

### GetFindingCodeOk

`func (o *Finding) GetFindingCodeOk() (*string, bool)`

GetFindingCodeOk returns a tuple with the FindingCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingCode

`func (o *Finding) SetFindingCode(v string)`

SetFindingCode sets FindingCode field to given value.

### HasFindingCode

`func (o *Finding) HasFindingCode() bool`

HasFindingCode returns a boolean if a field has been set.

### GetName

`func (o *Finding) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Finding) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Finding) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Finding) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Finding) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Finding) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Finding) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Finding) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSolution

`func (o *Finding) GetSolution() string`

GetSolution returns the Solution field if non-nil, zero value otherwise.

### GetSolutionOk

`func (o *Finding) GetSolutionOk() (*string, bool)`

GetSolutionOk returns a tuple with the Solution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolution

`func (o *Finding) SetSolution(v string)`

SetSolution sets Solution field to given value.

### HasSolution

`func (o *Finding) HasSolution() bool`

HasSolution returns a boolean if a field has been set.

### GetLinks

`func (o *Finding) GetLinks() FindingLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Finding) GetLinksOk() (*FindingLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Finding) SetLinks(v FindingLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Finding) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Finding) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Finding) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Finding) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Finding) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Finding) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Finding) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Finding) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Finding) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLastDetectedAt

`func (o *Finding) GetLastDetectedAt() int64`

GetLastDetectedAt returns the LastDetectedAt field if non-nil, zero value otherwise.

### GetLastDetectedAtOk

`func (o *Finding) GetLastDetectedAtOk() (*int64, bool)`

GetLastDetectedAtOk returns a tuple with the LastDetectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDetectedAt

`func (o *Finding) SetLastDetectedAt(v int64)`

SetLastDetectedAt sets LastDetectedAt field to given value.

### HasLastDetectedAt

`func (o *Finding) HasLastDetectedAt() bool`

HasLastDetectedAt returns a boolean if a field has been set.

### GetInstanceCount

`func (o *Finding) GetInstanceCount() int64`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *Finding) GetInstanceCountOk() (*int64, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *Finding) SetInstanceCount(v int64)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *Finding) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetRiskRank

`func (o *Finding) GetRiskRank() string`

GetRiskRank returns the RiskRank field if non-nil, zero value otherwise.

### GetRiskRankOk

`func (o *Finding) GetRiskRankOk() (*string, bool)`

GetRiskRankOk returns a tuple with the RiskRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskRank

`func (o *Finding) SetRiskRank(v string)`

SetRiskRank sets RiskRank field to given value.

### HasRiskRank

`func (o *Finding) HasRiskRank() bool`

HasRiskRank returns a boolean if a field has been set.

### GetRiskRankValue

`func (o *Finding) GetRiskRankValue() int32`

GetRiskRankValue returns the RiskRankValue field if non-nil, zero value otherwise.

### GetRiskRankValueOk

`func (o *Finding) GetRiskRankValueOk() (*int32, bool)`

GetRiskRankValueOk returns a tuple with the RiskRankValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskRankValue

`func (o *Finding) SetRiskRankValue(v int32)`

SetRiskRankValue sets RiskRankValue field to given value.

### HasRiskRankValue

`func (o *Finding) HasRiskRankValue() bool`

HasRiskRankValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


