# ScanTemplateOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the template. | 
**Description** | Pointer to **string** | Description of the template. | [optional] 
**OrganizationId** | **string** | The ID of the organization the template will be created in | 
**Params** | Pointer to **map[string]string** | A number of scan parameter values. Currently there is no authoritative list of acceptable values. See existing templates for examples. | [optional] 
**Global** | **bool** | Whether the template is globally available to all organizations. | 
**Acl** | **map[string]interface{}** | A map of IDs to strings which describe how the template may be accessed. Currently there is no authoritative list of acceptable values. See existing templates for examples. | 

## Methods

### NewScanTemplateOptions

`func NewScanTemplateOptions(name string, organizationId string, global bool, acl map[string]interface{}, ) *ScanTemplateOptions`

NewScanTemplateOptions instantiates a new ScanTemplateOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanTemplateOptionsWithDefaults

`func NewScanTemplateOptionsWithDefaults() *ScanTemplateOptions`

NewScanTemplateOptionsWithDefaults instantiates a new ScanTemplateOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ScanTemplateOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScanTemplateOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScanTemplateOptions) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ScanTemplateOptions) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScanTemplateOptions) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScanTemplateOptions) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScanTemplateOptions) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ScanTemplateOptions) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ScanTemplateOptions) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ScanTemplateOptions) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetParams

`func (o *ScanTemplateOptions) GetParams() map[string]string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ScanTemplateOptions) GetParamsOk() (*map[string]string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ScanTemplateOptions) SetParams(v map[string]string)`

SetParams sets Params field to given value.

### HasParams

`func (o *ScanTemplateOptions) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetGlobal

`func (o *ScanTemplateOptions) GetGlobal() bool`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *ScanTemplateOptions) GetGlobalOk() (*bool, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *ScanTemplateOptions) SetGlobal(v bool)`

SetGlobal sets Global field to given value.


### GetAcl

`func (o *ScanTemplateOptions) GetAcl() map[string]interface{}`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *ScanTemplateOptions) GetAclOk() (*map[string]interface{}, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *ScanTemplateOptions) SetAcl(v map[string]interface{})`

SetAcl sets Acl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


