# ExportServicesJSON200ResponseOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Services** | Pointer to [**[]Service**](Service.md) |  | [optional] 
**NextKey** | Pointer to **string** | The key to use for the next page of results | [optional] 

## Methods

### NewExportServicesJSON200ResponseOneOf

`func NewExportServicesJSON200ResponseOneOf() *ExportServicesJSON200ResponseOneOf`

NewExportServicesJSON200ResponseOneOf instantiates a new ExportServicesJSON200ResponseOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportServicesJSON200ResponseOneOfWithDefaults

`func NewExportServicesJSON200ResponseOneOfWithDefaults() *ExportServicesJSON200ResponseOneOf`

NewExportServicesJSON200ResponseOneOfWithDefaults instantiates a new ExportServicesJSON200ResponseOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServices

`func (o *ExportServicesJSON200ResponseOneOf) GetServices() []Service`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ExportServicesJSON200ResponseOneOf) GetServicesOk() (*[]Service, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ExportServicesJSON200ResponseOneOf) SetServices(v []Service)`

SetServices sets Services field to given value.

### HasServices

`func (o *ExportServicesJSON200ResponseOneOf) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetNextKey

`func (o *ExportServicesJSON200ResponseOneOf) GetNextKey() string`

GetNextKey returns the NextKey field if non-nil, zero value otherwise.

### GetNextKeyOk

`func (o *ExportServicesJSON200ResponseOneOf) GetNextKeyOk() (*string, bool)`

GetNextKeyOk returns a tuple with the NextKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextKey

`func (o *ExportServicesJSON200ResponseOneOf) SetNextKey(v string)`

SetNextKey sets NextKey field to given value.

### HasNextKey

`func (o *ExportServicesJSON200ResponseOneOf) HasNextKey() bool`

HasNextKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


