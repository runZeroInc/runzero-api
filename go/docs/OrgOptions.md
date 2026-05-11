# OrgOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**ExpirationAssetsStale** | Pointer to **float64** |  | [optional] 
**ExpirationAssetsOffline** | Pointer to **float64** |  | [optional] 
**ExpirationIntegrationAttributes** | Pointer to **float64** |  | [optional] 
**ExpirationScans** | Pointer to **float64** |  | [optional] 
**ExpirationVulnerabilities** | Pointer to **float64** |  | [optional] 
**KeepLatestIntegrationAttributes** | Pointer to **string** |  | [optional] 

## Methods

### NewOrgOptions

`func NewOrgOptions() *OrgOptions`

NewOrgOptions instantiates a new OrgOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgOptionsWithDefaults

`func NewOrgOptionsWithDefaults() *OrgOptions`

NewOrgOptionsWithDefaults instantiates a new OrgOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrgOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrgOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrgOptions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrgOptions) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *OrgOptions) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrgOptions) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrgOptions) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrgOptions) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParentId

`func (o *OrgOptions) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *OrgOptions) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *OrgOptions) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *OrgOptions) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetExpirationAssetsStale

`func (o *OrgOptions) GetExpirationAssetsStale() float64`

GetExpirationAssetsStale returns the ExpirationAssetsStale field if non-nil, zero value otherwise.

### GetExpirationAssetsStaleOk

`func (o *OrgOptions) GetExpirationAssetsStaleOk() (*float64, bool)`

GetExpirationAssetsStaleOk returns a tuple with the ExpirationAssetsStale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationAssetsStale

`func (o *OrgOptions) SetExpirationAssetsStale(v float64)`

SetExpirationAssetsStale sets ExpirationAssetsStale field to given value.

### HasExpirationAssetsStale

`func (o *OrgOptions) HasExpirationAssetsStale() bool`

HasExpirationAssetsStale returns a boolean if a field has been set.

### GetExpirationAssetsOffline

`func (o *OrgOptions) GetExpirationAssetsOffline() float64`

GetExpirationAssetsOffline returns the ExpirationAssetsOffline field if non-nil, zero value otherwise.

### GetExpirationAssetsOfflineOk

`func (o *OrgOptions) GetExpirationAssetsOfflineOk() (*float64, bool)`

GetExpirationAssetsOfflineOk returns a tuple with the ExpirationAssetsOffline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationAssetsOffline

`func (o *OrgOptions) SetExpirationAssetsOffline(v float64)`

SetExpirationAssetsOffline sets ExpirationAssetsOffline field to given value.

### HasExpirationAssetsOffline

`func (o *OrgOptions) HasExpirationAssetsOffline() bool`

HasExpirationAssetsOffline returns a boolean if a field has been set.

### GetExpirationIntegrationAttributes

`func (o *OrgOptions) GetExpirationIntegrationAttributes() float64`

GetExpirationIntegrationAttributes returns the ExpirationIntegrationAttributes field if non-nil, zero value otherwise.

### GetExpirationIntegrationAttributesOk

`func (o *OrgOptions) GetExpirationIntegrationAttributesOk() (*float64, bool)`

GetExpirationIntegrationAttributesOk returns a tuple with the ExpirationIntegrationAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationIntegrationAttributes

`func (o *OrgOptions) SetExpirationIntegrationAttributes(v float64)`

SetExpirationIntegrationAttributes sets ExpirationIntegrationAttributes field to given value.

### HasExpirationIntegrationAttributes

`func (o *OrgOptions) HasExpirationIntegrationAttributes() bool`

HasExpirationIntegrationAttributes returns a boolean if a field has been set.

### GetExpirationScans

`func (o *OrgOptions) GetExpirationScans() float64`

GetExpirationScans returns the ExpirationScans field if non-nil, zero value otherwise.

### GetExpirationScansOk

`func (o *OrgOptions) GetExpirationScansOk() (*float64, bool)`

GetExpirationScansOk returns a tuple with the ExpirationScans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationScans

`func (o *OrgOptions) SetExpirationScans(v float64)`

SetExpirationScans sets ExpirationScans field to given value.

### HasExpirationScans

`func (o *OrgOptions) HasExpirationScans() bool`

HasExpirationScans returns a boolean if a field has been set.

### GetExpirationVulnerabilities

`func (o *OrgOptions) GetExpirationVulnerabilities() float64`

GetExpirationVulnerabilities returns the ExpirationVulnerabilities field if non-nil, zero value otherwise.

### GetExpirationVulnerabilitiesOk

`func (o *OrgOptions) GetExpirationVulnerabilitiesOk() (*float64, bool)`

GetExpirationVulnerabilitiesOk returns a tuple with the ExpirationVulnerabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationVulnerabilities

`func (o *OrgOptions) SetExpirationVulnerabilities(v float64)`

SetExpirationVulnerabilities sets ExpirationVulnerabilities field to given value.

### HasExpirationVulnerabilities

`func (o *OrgOptions) HasExpirationVulnerabilities() bool`

HasExpirationVulnerabilities returns a boolean if a field has been set.

### GetKeepLatestIntegrationAttributes

`func (o *OrgOptions) GetKeepLatestIntegrationAttributes() string`

GetKeepLatestIntegrationAttributes returns the KeepLatestIntegrationAttributes field if non-nil, zero value otherwise.

### GetKeepLatestIntegrationAttributesOk

`func (o *OrgOptions) GetKeepLatestIntegrationAttributesOk() (*string, bool)`

GetKeepLatestIntegrationAttributesOk returns a tuple with the KeepLatestIntegrationAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepLatestIntegrationAttributes

`func (o *OrgOptions) SetKeepLatestIntegrationAttributes(v string)`

SetKeepLatestIntegrationAttributes sets KeepLatestIntegrationAttributes field to given value.

### HasKeepLatestIntegrationAttributes

`func (o *OrgOptions) HasKeepLatestIntegrationAttributes() bool`

HasKeepLatestIntegrationAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


