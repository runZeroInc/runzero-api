# AssetOwnershipType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Reference** | Pointer to **int64** |  | [optional] 
**Order** | Pointer to **int64** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 

## Methods

### NewAssetOwnershipType

`func NewAssetOwnershipType(id string, name string, ) *AssetOwnershipType`

NewAssetOwnershipType instantiates a new AssetOwnershipType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetOwnershipTypeWithDefaults

`func NewAssetOwnershipTypeWithDefaults() *AssetOwnershipType`

NewAssetOwnershipTypeWithDefaults instantiates a new AssetOwnershipType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AssetOwnershipType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssetOwnershipType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssetOwnershipType) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AssetOwnershipType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssetOwnershipType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssetOwnershipType) SetName(v string)`

SetName sets Name field to given value.


### GetReference

`func (o *AssetOwnershipType) GetReference() int64`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *AssetOwnershipType) GetReferenceOk() (*int64, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *AssetOwnershipType) SetReference(v int64)`

SetReference sets Reference field to given value.

### HasReference

`func (o *AssetOwnershipType) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetOrder

`func (o *AssetOwnershipType) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *AssetOwnershipType) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *AssetOwnershipType) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *AssetOwnershipType) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetHidden

`func (o *AssetOwnershipType) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *AssetOwnershipType) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *AssetOwnershipType) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *AssetOwnershipType) HasHidden() bool`

HasHidden returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


