# AssetExportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**[]Asset**](Asset.md) |  | [optional] 
**NextKey** | Pointer to **string** | The key to use for the next page of results | [optional] 

## Methods

### NewAssetExportResponse

`func NewAssetExportResponse() *AssetExportResponse`

NewAssetExportResponse instantiates a new AssetExportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetExportResponseWithDefaults

`func NewAssetExportResponseWithDefaults() *AssetExportResponse`

NewAssetExportResponseWithDefaults instantiates a new AssetExportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *AssetExportResponse) GetAssets() []Asset`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *AssetExportResponse) GetAssetsOk() (*[]Asset, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *AssetExportResponse) SetAssets(v []Asset)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *AssetExportResponse) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetNextKey

`func (o *AssetExportResponse) GetNextKey() string`

GetNextKey returns the NextKey field if non-nil, zero value otherwise.

### GetNextKeyOk

`func (o *AssetExportResponse) GetNextKeyOk() (*string, bool)`

GetNextKeyOk returns a tuple with the NextKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextKey

`func (o *AssetExportResponse) SetNextKey(v string)`

SetNextKey sets NextKey field to given value.

### HasNextKey

`func (o *AssetExportResponse) HasNextKey() bool`

HasNextKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


