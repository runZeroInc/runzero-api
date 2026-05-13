# BaseCustomIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The unique name of the custom integration, without spaces. | [optional] 
**Icon** | Pointer to **string** | Base64 encoded png with maximum size 256x256 pixels | [optional] 
**Description** | Pointer to **string** | A text description of the custom integration | [optional] 

## Methods

### NewBaseCustomIntegration

`func NewBaseCustomIntegration() *BaseCustomIntegration`

NewBaseCustomIntegration instantiates a new BaseCustomIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseCustomIntegrationWithDefaults

`func NewBaseCustomIntegrationWithDefaults() *BaseCustomIntegration`

NewBaseCustomIntegrationWithDefaults instantiates a new BaseCustomIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BaseCustomIntegration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseCustomIntegration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseCustomIntegration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BaseCustomIntegration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIcon

`func (o *BaseCustomIntegration) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *BaseCustomIntegration) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *BaseCustomIntegration) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *BaseCustomIntegration) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetDescription

`func (o *BaseCustomIntegration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BaseCustomIntegration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BaseCustomIntegration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BaseCustomIntegration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


