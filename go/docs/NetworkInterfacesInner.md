# NetworkInterfacesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4Addresses** | Pointer to **[]string** | Represents IPV4 addresses. Addresses are ordered from most to least likely to uniquely identify the asset. | [optional] 
**Ipv6Addresses** | Pointer to **[]string** | Represents the IPV6 addresses. Addresses are ordered from most to least likely to uniquely identify the asset. | [optional] 
**MacAddress** | Pointer to **string** | Represents a MAC address in IEEE 802 MAC/EUI-48, or EUI-64 form in one of the following formats:   01:23:45:67:89:AB   01:23:45:67:89:ab:cd:ef   01-23-45-67-89-ab   01-23-45-67-89-ab-cd-ef   0123.4567.89ab   0123.4567.89ab.cdef   0123 4567 89ab cdEF  | [optional] 

## Methods

### NewNetworkInterfacesInner

`func NewNetworkInterfacesInner() *NetworkInterfacesInner`

NewNetworkInterfacesInner instantiates a new NetworkInterfacesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkInterfacesInnerWithDefaults

`func NewNetworkInterfacesInnerWithDefaults() *NetworkInterfacesInner`

NewNetworkInterfacesInnerWithDefaults instantiates a new NetworkInterfacesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4Addresses

`func (o *NetworkInterfacesInner) GetIpv4Addresses() []string`

GetIpv4Addresses returns the Ipv4Addresses field if non-nil, zero value otherwise.

### GetIpv4AddressesOk

`func (o *NetworkInterfacesInner) GetIpv4AddressesOk() (*[]string, bool)`

GetIpv4AddressesOk returns a tuple with the Ipv4Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addresses

`func (o *NetworkInterfacesInner) SetIpv4Addresses(v []string)`

SetIpv4Addresses sets Ipv4Addresses field to given value.

### HasIpv4Addresses

`func (o *NetworkInterfacesInner) HasIpv4Addresses() bool`

HasIpv4Addresses returns a boolean if a field has been set.

### GetIpv6Addresses

`func (o *NetworkInterfacesInner) GetIpv6Addresses() []string`

GetIpv6Addresses returns the Ipv6Addresses field if non-nil, zero value otherwise.

### GetIpv6AddressesOk

`func (o *NetworkInterfacesInner) GetIpv6AddressesOk() (*[]string, bool)`

GetIpv6AddressesOk returns a tuple with the Ipv6Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addresses

`func (o *NetworkInterfacesInner) SetIpv6Addresses(v []string)`

SetIpv6Addresses sets Ipv6Addresses field to given value.

### HasIpv6Addresses

`func (o *NetworkInterfacesInner) HasIpv6Addresses() bool`

HasIpv6Addresses returns a boolean if a field has been set.

### GetMacAddress

`func (o *NetworkInterfacesInner) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *NetworkInterfacesInner) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *NetworkInterfacesInner) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *NetworkInterfacesInner) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


