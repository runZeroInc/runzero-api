# ExportAssetsJSON200ResponseOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**[]Asset**](Asset.md) |  | [optional] 
**NextKey** | Pointer to **string** | The key to use for the next page of results | [optional] 

## Methods

### NewExportAssetsJSON200ResponseOneOf

`func NewExportAssetsJSON200ResponseOneOf() *ExportAssetsJSON200ResponseOneOf`

NewExportAssetsJSON200ResponseOneOf instantiates a new ExportAssetsJSON200ResponseOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportAssetsJSON200ResponseOneOfWithDefaults

`func NewExportAssetsJSON200ResponseOneOfWithDefaults() *ExportAssetsJSON200ResponseOneOf`

NewExportAssetsJSON200ResponseOneOfWithDefaults instantiates a new ExportAssetsJSON200ResponseOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *ExportAssetsJSON200ResponseOneOf) GetAssets() []Asset`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *ExportAssetsJSON200ResponseOneOf) GetAssetsOk() (*[]Asset, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *ExportAssetsJSON200ResponseOneOf) SetAssets(v []Asset)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *ExportAssetsJSON200ResponseOneOf) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetNextKey

`func (o *ExportAssetsJSON200ResponseOneOf) GetNextKey() string`

GetNextKey returns the NextKey field if non-nil, zero value otherwise.

### GetNextKeyOk

`func (o *ExportAssetsJSON200ResponseOneOf) GetNextKeyOk() (*string, bool)`

GetNextKeyOk returns a tuple with the NextKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextKey

`func (o *ExportAssetsJSON200ResponseOneOf) SetNextKey(v string)`

SetNextKey sets NextKey field to given value.

### HasNextKey

`func (o *ExportAssetsJSON200ResponseOneOf) HasNextKey() bool`

HasNextKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


