# SoftwareExportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Software** | Pointer to [**[]Software**](Software.md) |  | [optional] 
**NextKey** | Pointer to **string** | The key to use for the next page of results | [optional] 

## Methods

### NewSoftwareExportResponse

`func NewSoftwareExportResponse() *SoftwareExportResponse`

NewSoftwareExportResponse instantiates a new SoftwareExportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareExportResponseWithDefaults

`func NewSoftwareExportResponseWithDefaults() *SoftwareExportResponse`

NewSoftwareExportResponseWithDefaults instantiates a new SoftwareExportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSoftware

`func (o *SoftwareExportResponse) GetSoftware() []Software`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *SoftwareExportResponse) GetSoftwareOk() (*[]Software, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware

`func (o *SoftwareExportResponse) SetSoftware(v []Software)`

SetSoftware sets Software field to given value.

### HasSoftware

`func (o *SoftwareExportResponse) HasSoftware() bool`

HasSoftware returns a boolean if a field has been set.

### GetNextKey

`func (o *SoftwareExportResponse) GetNextKey() string`

GetNextKey returns the NextKey field if non-nil, zero value otherwise.

### GetNextKeyOk

`func (o *SoftwareExportResponse) GetNextKeyOk() (*string, bool)`

GetNextKeyOk returns a tuple with the NextKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextKey

`func (o *SoftwareExportResponse) SetNextKey(v string)`

SetNextKey sets NextKey field to given value.

### HasNextKey

`func (o *SoftwareExportResponse) HasNextKey() bool`

HasNextKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


