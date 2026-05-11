# WirelessExportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Wireless** | Pointer to [**[]Wireless**](Wireless.md) |  | [optional] 
**NextKey** | Pointer to **string** | The key to use for the next page of results | [optional] 

## Methods

### NewWirelessExportResponse

`func NewWirelessExportResponse() *WirelessExportResponse`

NewWirelessExportResponse instantiates a new WirelessExportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWirelessExportResponseWithDefaults

`func NewWirelessExportResponseWithDefaults() *WirelessExportResponse`

NewWirelessExportResponseWithDefaults instantiates a new WirelessExportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWireless

`func (o *WirelessExportResponse) GetWireless() []Wireless`

GetWireless returns the Wireless field if non-nil, zero value otherwise.

### GetWirelessOk

`func (o *WirelessExportResponse) GetWirelessOk() (*[]Wireless, bool)`

GetWirelessOk returns a tuple with the Wireless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireless

`func (o *WirelessExportResponse) SetWireless(v []Wireless)`

SetWireless sets Wireless field to given value.

### HasWireless

`func (o *WirelessExportResponse) HasWireless() bool`

HasWireless returns a boolean if a field has been set.

### GetNextKey

`func (o *WirelessExportResponse) GetNextKey() string`

GetNextKey returns the NextKey field if non-nil, zero value otherwise.

### GetNextKeyOk

`func (o *WirelessExportResponse) GetNextKeyOk() (*string, bool)`

GetNextKeyOk returns a tuple with the NextKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextKey

`func (o *WirelessExportResponse) SetNextKey(v string)`

SetNextKey sets NextKey field to given value.

### HasNextKey

`func (o *WirelessExportResponse) HasNextKey() bool`

HasNextKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


