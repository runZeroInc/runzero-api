# SoftwarePage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Software** | Pointer to [**[]Software**](Software.md) |  | [optional] 
**NextKey** | Pointer to **string** | The key to use for the next page of results | [optional] 

## Methods

### NewSoftwarePage

`func NewSoftwarePage() *SoftwarePage`

NewSoftwarePage instantiates a new SoftwarePage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarePageWithDefaults

`func NewSoftwarePageWithDefaults() *SoftwarePage`

NewSoftwarePageWithDefaults instantiates a new SoftwarePage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSoftware

`func (o *SoftwarePage) GetSoftware() []Software`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *SoftwarePage) GetSoftwareOk() (*[]Software, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware

`func (o *SoftwarePage) SetSoftware(v []Software)`

SetSoftware sets Software field to given value.

### HasSoftware

`func (o *SoftwarePage) HasSoftware() bool`

HasSoftware returns a boolean if a field has been set.

### GetNextKey

`func (o *SoftwarePage) GetNextKey() string`

GetNextKey returns the NextKey field if non-nil, zero value otherwise.

### GetNextKeyOk

`func (o *SoftwarePage) GetNextKeyOk() (*string, bool)`

GetNextKeyOk returns a tuple with the NextKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextKey

`func (o *SoftwarePage) SetNextKey(v string)`

SetNextKey sets NextKey field to given value.

### HasNextKey

`func (o *SoftwarePage) HasNextKey() bool`

HasNextKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


