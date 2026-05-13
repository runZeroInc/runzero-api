# ImportTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ExcludeUnknown** | Pointer to **bool** | Instructs the data ingestion process whether to skip assets which do not merge into an existing asset in the asset inventory | [optional] [default to false]
**Tags** | Pointer to **[]string** | Arbitrary string tag values which are applied to the asset data import task created. | [optional] 

## Methods

### NewImportTask

`func NewImportTask(name string, ) *ImportTask`

NewImportTask instantiates a new ImportTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportTaskWithDefaults

`func NewImportTaskWithDefaults() *ImportTask`

NewImportTaskWithDefaults instantiates a new ImportTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ImportTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImportTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImportTask) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ImportTask) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ImportTask) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ImportTask) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ImportTask) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExcludeUnknown

`func (o *ImportTask) GetExcludeUnknown() bool`

GetExcludeUnknown returns the ExcludeUnknown field if non-nil, zero value otherwise.

### GetExcludeUnknownOk

`func (o *ImportTask) GetExcludeUnknownOk() (*bool, bool)`

GetExcludeUnknownOk returns a tuple with the ExcludeUnknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeUnknown

`func (o *ImportTask) SetExcludeUnknown(v bool)`

SetExcludeUnknown sets ExcludeUnknown field to given value.

### HasExcludeUnknown

`func (o *ImportTask) HasExcludeUnknown() bool`

HasExcludeUnknown returns a boolean if a field has been set.

### GetTags

`func (o *ImportTask) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ImportTask) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ImportTask) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ImportTask) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


