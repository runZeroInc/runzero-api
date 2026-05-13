# WirelessPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Wireless** | Pointer to [**[]Wireless**](Wireless.md) |  | [optional] 
**NextKey** | Pointer to **string** | The key to use for the next page of results | [optional] 

## Methods

### NewWirelessPage

`func NewWirelessPage() *WirelessPage`

NewWirelessPage instantiates a new WirelessPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWirelessPageWithDefaults

`func NewWirelessPageWithDefaults() *WirelessPage`

NewWirelessPageWithDefaults instantiates a new WirelessPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWireless

`func (o *WirelessPage) GetWireless() []Wireless`

GetWireless returns the Wireless field if non-nil, zero value otherwise.

### GetWirelessOk

`func (o *WirelessPage) GetWirelessOk() (*[]Wireless, bool)`

GetWirelessOk returns a tuple with the Wireless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireless

`func (o *WirelessPage) SetWireless(v []Wireless)`

SetWireless sets Wireless field to given value.

### HasWireless

`func (o *WirelessPage) HasWireless() bool`

HasWireless returns a boolean if a field has been set.

### GetNextKey

`func (o *WirelessPage) GetNextKey() string`

GetNextKey returns the NextKey field if non-nil, zero value otherwise.

### GetNextKeyOk

`func (o *WirelessPage) GetNextKeyOk() (*string, bool)`

GetNextKeyOk returns a tuple with the NextKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextKey

`func (o *WirelessPage) SetNextKey(v string)`

SetNextKey sets NextKey field to given value.

### HasNextKey

`func (o *WirelessPage) HasNextKey() bool`

HasNextKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


