# ExportSoftwareJSON200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Software** | Pointer to [**[]Software**](Software.md) |  | [optional] 
**NextKey** | Pointer to **string** | The key to use for the next page of results | [optional] 

## Methods

### NewExportSoftwareJSON200Response

`func NewExportSoftwareJSON200Response() *ExportSoftwareJSON200Response`

NewExportSoftwareJSON200Response instantiates a new ExportSoftwareJSON200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportSoftwareJSON200ResponseWithDefaults

`func NewExportSoftwareJSON200ResponseWithDefaults() *ExportSoftwareJSON200Response`

NewExportSoftwareJSON200ResponseWithDefaults instantiates a new ExportSoftwareJSON200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSoftware

`func (o *ExportSoftwareJSON200Response) GetSoftware() []Software`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *ExportSoftwareJSON200Response) GetSoftwareOk() (*[]Software, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware

`func (o *ExportSoftwareJSON200Response) SetSoftware(v []Software)`

SetSoftware sets Software field to given value.

### HasSoftware

`func (o *ExportSoftwareJSON200Response) HasSoftware() bool`

HasSoftware returns a boolean if a field has been set.

### GetNextKey

`func (o *ExportSoftwareJSON200Response) GetNextKey() string`

GetNextKey returns the NextKey field if non-nil, zero value otherwise.

### GetNextKeyOk

`func (o *ExportSoftwareJSON200Response) GetNextKeyOk() (*string, bool)`

GetNextKeyOk returns a tuple with the NextKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextKey

`func (o *ExportSoftwareJSON200Response) SetNextKey(v string)`

SetNextKey sets NextKey field to given value.

### HasNextKey

`func (o *ExportSoftwareJSON200Response) HasNextKey() bool`

HasNextKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


