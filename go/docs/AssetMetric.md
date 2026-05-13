# AssetMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | Pointer to **string** |  | [optional] 
**Metrics** | Pointer to **map[string]interface{}** | Free-form metric values keyed by metric name. | [optional] 

## Methods

### NewAssetMetric

`func NewAssetMetric() *AssetMetric`

NewAssetMetric instantiates a new AssetMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetMetricWithDefaults

`func NewAssetMetricWithDefaults() *AssetMetric`

NewAssetMetricWithDefaults instantiates a new AssetMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *AssetMetric) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *AssetMetric) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *AssetMetric) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *AssetMetric) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetMetrics

`func (o *AssetMetric) GetMetrics() map[string]interface{}`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *AssetMetric) GetMetricsOk() (*map[string]interface{}, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *AssetMetric) SetMetrics(v map[string]interface{})`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *AssetMetric) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


