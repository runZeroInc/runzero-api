# AssetServiceNow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | **string** |  | 
**Organization** | Pointer to **string** |  | [optional] 
**Site** | Pointer to **string** |  | [optional] 
**DetectedBy** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**SysClassName** | Pointer to **string** |  | [optional] 
**OsVendor** | Pointer to **string** |  | [optional] 
**OsProduct** | Pointer to **string** |  | [optional] 
**OsVersion** | Pointer to **string** |  | [optional] 
**HwVendor** | Pointer to **string** |  | [optional] 
**HwProduct** | Pointer to **string** |  | [optional] 
**HwVersion** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **NullableString** |  | [optional] 
**AddressesScope** | Pointer to **string** |  | [optional] 
**AddressesExtra** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 
**MacManufacturer** | Pointer to **string** |  | [optional] 
**NewestMacAge** | Pointer to **string** |  | [optional] 
**Macs** | Pointer to **string** |  | [optional] 
**MacVendors** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**Domains** | Pointer to **string** |  | [optional] 
**ServiceCount** | Pointer to **int64** |  | [optional] 
**ServiceCountTcp** | Pointer to **int64** |  | [optional] 
**ServiceCountUdp** | Pointer to **int64** |  | [optional] 
**ServiceCountArp** | Pointer to **int64** |  | [optional] 
**ServiceCountIcmp** | Pointer to **int64** |  | [optional] 
**LowestTtl** | Pointer to **int64** |  | [optional] 
**LowestRtt** | Pointer to **int64** |  | [optional] 
**Alive** | Pointer to **bool** |  | [optional] 
**FirstDiscovered** | Pointer to **string** |  | [optional] 
**LastDiscovered** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 

## Methods

### NewAssetServiceNow

`func NewAssetServiceNow(assetId string, ) *AssetServiceNow`

NewAssetServiceNow instantiates a new AssetServiceNow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetServiceNowWithDefaults

`func NewAssetServiceNowWithDefaults() *AssetServiceNow`

NewAssetServiceNowWithDefaults instantiates a new AssetServiceNow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *AssetServiceNow) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AssetServiceNow) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AssetServiceNow) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetOrganization

`func (o *AssetServiceNow) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *AssetServiceNow) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *AssetServiceNow) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *AssetServiceNow) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSite

`func (o *AssetServiceNow) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *AssetServiceNow) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *AssetServiceNow) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *AssetServiceNow) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetDetectedBy

`func (o *AssetServiceNow) GetDetectedBy() string`

GetDetectedBy returns the DetectedBy field if non-nil, zero value otherwise.

### GetDetectedByOk

`func (o *AssetServiceNow) GetDetectedByOk() (*string, bool)`

GetDetectedByOk returns a tuple with the DetectedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedBy

`func (o *AssetServiceNow) SetDetectedBy(v string)`

SetDetectedBy sets DetectedBy field to given value.

### HasDetectedBy

`func (o *AssetServiceNow) HasDetectedBy() bool`

HasDetectedBy returns a boolean if a field has been set.

### GetType

`func (o *AssetServiceNow) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AssetServiceNow) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AssetServiceNow) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AssetServiceNow) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSysClassName

`func (o *AssetServiceNow) GetSysClassName() string`

GetSysClassName returns the SysClassName field if non-nil, zero value otherwise.

### GetSysClassNameOk

`func (o *AssetServiceNow) GetSysClassNameOk() (*string, bool)`

GetSysClassNameOk returns a tuple with the SysClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysClassName

`func (o *AssetServiceNow) SetSysClassName(v string)`

SetSysClassName sets SysClassName field to given value.

### HasSysClassName

`func (o *AssetServiceNow) HasSysClassName() bool`

HasSysClassName returns a boolean if a field has been set.

### GetOsVendor

`func (o *AssetServiceNow) GetOsVendor() string`

GetOsVendor returns the OsVendor field if non-nil, zero value otherwise.

### GetOsVendorOk

`func (o *AssetServiceNow) GetOsVendorOk() (*string, bool)`

GetOsVendorOk returns a tuple with the OsVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVendor

`func (o *AssetServiceNow) SetOsVendor(v string)`

SetOsVendor sets OsVendor field to given value.

### HasOsVendor

`func (o *AssetServiceNow) HasOsVendor() bool`

HasOsVendor returns a boolean if a field has been set.

### GetOsProduct

`func (o *AssetServiceNow) GetOsProduct() string`

GetOsProduct returns the OsProduct field if non-nil, zero value otherwise.

### GetOsProductOk

`func (o *AssetServiceNow) GetOsProductOk() (*string, bool)`

GetOsProductOk returns a tuple with the OsProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsProduct

`func (o *AssetServiceNow) SetOsProduct(v string)`

SetOsProduct sets OsProduct field to given value.

### HasOsProduct

`func (o *AssetServiceNow) HasOsProduct() bool`

HasOsProduct returns a boolean if a field has been set.

### GetOsVersion

`func (o *AssetServiceNow) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *AssetServiceNow) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *AssetServiceNow) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *AssetServiceNow) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetHwVendor

`func (o *AssetServiceNow) GetHwVendor() string`

GetHwVendor returns the HwVendor field if non-nil, zero value otherwise.

### GetHwVendorOk

`func (o *AssetServiceNow) GetHwVendorOk() (*string, bool)`

GetHwVendorOk returns a tuple with the HwVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwVendor

`func (o *AssetServiceNow) SetHwVendor(v string)`

SetHwVendor sets HwVendor field to given value.

### HasHwVendor

`func (o *AssetServiceNow) HasHwVendor() bool`

HasHwVendor returns a boolean if a field has been set.

### GetHwProduct

`func (o *AssetServiceNow) GetHwProduct() string`

GetHwProduct returns the HwProduct field if non-nil, zero value otherwise.

### GetHwProductOk

`func (o *AssetServiceNow) GetHwProductOk() (*string, bool)`

GetHwProductOk returns a tuple with the HwProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwProduct

`func (o *AssetServiceNow) SetHwProduct(v string)`

SetHwProduct sets HwProduct field to given value.

### HasHwProduct

`func (o *AssetServiceNow) HasHwProduct() bool`

HasHwProduct returns a boolean if a field has been set.

### GetHwVersion

`func (o *AssetServiceNow) GetHwVersion() string`

GetHwVersion returns the HwVersion field if non-nil, zero value otherwise.

### GetHwVersionOk

`func (o *AssetServiceNow) GetHwVersionOk() (*string, bool)`

GetHwVersionOk returns a tuple with the HwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwVersion

`func (o *AssetServiceNow) SetHwVersion(v string)`

SetHwVersion sets HwVersion field to given value.

### HasHwVersion

`func (o *AssetServiceNow) HasHwVersion() bool`

HasHwVersion returns a boolean if a field has been set.

### GetIpAddress

`func (o *AssetServiceNow) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *AssetServiceNow) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *AssetServiceNow) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *AssetServiceNow) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *AssetServiceNow) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *AssetServiceNow) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetAddressesScope

`func (o *AssetServiceNow) GetAddressesScope() string`

GetAddressesScope returns the AddressesScope field if non-nil, zero value otherwise.

### GetAddressesScopeOk

`func (o *AssetServiceNow) GetAddressesScopeOk() (*string, bool)`

GetAddressesScopeOk returns a tuple with the AddressesScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressesScope

`func (o *AssetServiceNow) SetAddressesScope(v string)`

SetAddressesScope sets AddressesScope field to given value.

### HasAddressesScope

`func (o *AssetServiceNow) HasAddressesScope() bool`

HasAddressesScope returns a boolean if a field has been set.

### GetAddressesExtra

`func (o *AssetServiceNow) GetAddressesExtra() string`

GetAddressesExtra returns the AddressesExtra field if non-nil, zero value otherwise.

### GetAddressesExtraOk

`func (o *AssetServiceNow) GetAddressesExtraOk() (*string, bool)`

GetAddressesExtraOk returns a tuple with the AddressesExtra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressesExtra

`func (o *AssetServiceNow) SetAddressesExtra(v string)`

SetAddressesExtra sets AddressesExtra field to given value.

### HasAddressesExtra

`func (o *AssetServiceNow) HasAddressesExtra() bool`

HasAddressesExtra returns a boolean if a field has been set.

### GetMacAddress

`func (o *AssetServiceNow) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *AssetServiceNow) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *AssetServiceNow) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *AssetServiceNow) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMacManufacturer

`func (o *AssetServiceNow) GetMacManufacturer() string`

GetMacManufacturer returns the MacManufacturer field if non-nil, zero value otherwise.

### GetMacManufacturerOk

`func (o *AssetServiceNow) GetMacManufacturerOk() (*string, bool)`

GetMacManufacturerOk returns a tuple with the MacManufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacManufacturer

`func (o *AssetServiceNow) SetMacManufacturer(v string)`

SetMacManufacturer sets MacManufacturer field to given value.

### HasMacManufacturer

`func (o *AssetServiceNow) HasMacManufacturer() bool`

HasMacManufacturer returns a boolean if a field has been set.

### GetNewestMacAge

`func (o *AssetServiceNow) GetNewestMacAge() string`

GetNewestMacAge returns the NewestMacAge field if non-nil, zero value otherwise.

### GetNewestMacAgeOk

`func (o *AssetServiceNow) GetNewestMacAgeOk() (*string, bool)`

GetNewestMacAgeOk returns a tuple with the NewestMacAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewestMacAge

`func (o *AssetServiceNow) SetNewestMacAge(v string)`

SetNewestMacAge sets NewestMacAge field to given value.

### HasNewestMacAge

`func (o *AssetServiceNow) HasNewestMacAge() bool`

HasNewestMacAge returns a boolean if a field has been set.

### GetMacs

`func (o *AssetServiceNow) GetMacs() string`

GetMacs returns the Macs field if non-nil, zero value otherwise.

### GetMacsOk

`func (o *AssetServiceNow) GetMacsOk() (*string, bool)`

GetMacsOk returns a tuple with the Macs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacs

`func (o *AssetServiceNow) SetMacs(v string)`

SetMacs sets Macs field to given value.

### HasMacs

`func (o *AssetServiceNow) HasMacs() bool`

HasMacs returns a boolean if a field has been set.

### GetMacVendors

`func (o *AssetServiceNow) GetMacVendors() string`

GetMacVendors returns the MacVendors field if non-nil, zero value otherwise.

### GetMacVendorsOk

`func (o *AssetServiceNow) GetMacVendorsOk() (*string, bool)`

GetMacVendorsOk returns a tuple with the MacVendors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacVendors

`func (o *AssetServiceNow) SetMacVendors(v string)`

SetMacVendors sets MacVendors field to given value.

### HasMacVendors

`func (o *AssetServiceNow) HasMacVendors() bool`

HasMacVendors returns a boolean if a field has been set.

### GetName

`func (o *AssetServiceNow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssetServiceNow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssetServiceNow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssetServiceNow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *AssetServiceNow) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AssetServiceNow) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AssetServiceNow) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AssetServiceNow) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDomains

`func (o *AssetServiceNow) GetDomains() string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *AssetServiceNow) GetDomainsOk() (*string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *AssetServiceNow) SetDomains(v string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *AssetServiceNow) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetServiceCount

`func (o *AssetServiceNow) GetServiceCount() int64`

GetServiceCount returns the ServiceCount field if non-nil, zero value otherwise.

### GetServiceCountOk

`func (o *AssetServiceNow) GetServiceCountOk() (*int64, bool)`

GetServiceCountOk returns a tuple with the ServiceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCount

`func (o *AssetServiceNow) SetServiceCount(v int64)`

SetServiceCount sets ServiceCount field to given value.

### HasServiceCount

`func (o *AssetServiceNow) HasServiceCount() bool`

HasServiceCount returns a boolean if a field has been set.

### GetServiceCountTcp

`func (o *AssetServiceNow) GetServiceCountTcp() int64`

GetServiceCountTcp returns the ServiceCountTcp field if non-nil, zero value otherwise.

### GetServiceCountTcpOk

`func (o *AssetServiceNow) GetServiceCountTcpOk() (*int64, bool)`

GetServiceCountTcpOk returns a tuple with the ServiceCountTcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCountTcp

`func (o *AssetServiceNow) SetServiceCountTcp(v int64)`

SetServiceCountTcp sets ServiceCountTcp field to given value.

### HasServiceCountTcp

`func (o *AssetServiceNow) HasServiceCountTcp() bool`

HasServiceCountTcp returns a boolean if a field has been set.

### GetServiceCountUdp

`func (o *AssetServiceNow) GetServiceCountUdp() int64`

GetServiceCountUdp returns the ServiceCountUdp field if non-nil, zero value otherwise.

### GetServiceCountUdpOk

`func (o *AssetServiceNow) GetServiceCountUdpOk() (*int64, bool)`

GetServiceCountUdpOk returns a tuple with the ServiceCountUdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCountUdp

`func (o *AssetServiceNow) SetServiceCountUdp(v int64)`

SetServiceCountUdp sets ServiceCountUdp field to given value.

### HasServiceCountUdp

`func (o *AssetServiceNow) HasServiceCountUdp() bool`

HasServiceCountUdp returns a boolean if a field has been set.

### GetServiceCountArp

`func (o *AssetServiceNow) GetServiceCountArp() int64`

GetServiceCountArp returns the ServiceCountArp field if non-nil, zero value otherwise.

### GetServiceCountArpOk

`func (o *AssetServiceNow) GetServiceCountArpOk() (*int64, bool)`

GetServiceCountArpOk returns a tuple with the ServiceCountArp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCountArp

`func (o *AssetServiceNow) SetServiceCountArp(v int64)`

SetServiceCountArp sets ServiceCountArp field to given value.

### HasServiceCountArp

`func (o *AssetServiceNow) HasServiceCountArp() bool`

HasServiceCountArp returns a boolean if a field has been set.

### GetServiceCountIcmp

`func (o *AssetServiceNow) GetServiceCountIcmp() int64`

GetServiceCountIcmp returns the ServiceCountIcmp field if non-nil, zero value otherwise.

### GetServiceCountIcmpOk

`func (o *AssetServiceNow) GetServiceCountIcmpOk() (*int64, bool)`

GetServiceCountIcmpOk returns a tuple with the ServiceCountIcmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCountIcmp

`func (o *AssetServiceNow) SetServiceCountIcmp(v int64)`

SetServiceCountIcmp sets ServiceCountIcmp field to given value.

### HasServiceCountIcmp

`func (o *AssetServiceNow) HasServiceCountIcmp() bool`

HasServiceCountIcmp returns a boolean if a field has been set.

### GetLowestTtl

`func (o *AssetServiceNow) GetLowestTtl() int64`

GetLowestTtl returns the LowestTtl field if non-nil, zero value otherwise.

### GetLowestTtlOk

`func (o *AssetServiceNow) GetLowestTtlOk() (*int64, bool)`

GetLowestTtlOk returns a tuple with the LowestTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowestTtl

`func (o *AssetServiceNow) SetLowestTtl(v int64)`

SetLowestTtl sets LowestTtl field to given value.

### HasLowestTtl

`func (o *AssetServiceNow) HasLowestTtl() bool`

HasLowestTtl returns a boolean if a field has been set.

### GetLowestRtt

`func (o *AssetServiceNow) GetLowestRtt() int64`

GetLowestRtt returns the LowestRtt field if non-nil, zero value otherwise.

### GetLowestRttOk

`func (o *AssetServiceNow) GetLowestRttOk() (*int64, bool)`

GetLowestRttOk returns a tuple with the LowestRtt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowestRtt

`func (o *AssetServiceNow) SetLowestRtt(v int64)`

SetLowestRtt sets LowestRtt field to given value.

### HasLowestRtt

`func (o *AssetServiceNow) HasLowestRtt() bool`

HasLowestRtt returns a boolean if a field has been set.

### GetAlive

`func (o *AssetServiceNow) GetAlive() bool`

GetAlive returns the Alive field if non-nil, zero value otherwise.

### GetAliveOk

`func (o *AssetServiceNow) GetAliveOk() (*bool, bool)`

GetAliveOk returns a tuple with the Alive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlive

`func (o *AssetServiceNow) SetAlive(v bool)`

SetAlive sets Alive field to given value.

### HasAlive

`func (o *AssetServiceNow) HasAlive() bool`

HasAlive returns a boolean if a field has been set.

### GetFirstDiscovered

`func (o *AssetServiceNow) GetFirstDiscovered() string`

GetFirstDiscovered returns the FirstDiscovered field if non-nil, zero value otherwise.

### GetFirstDiscoveredOk

`func (o *AssetServiceNow) GetFirstDiscoveredOk() (*string, bool)`

GetFirstDiscoveredOk returns a tuple with the FirstDiscovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstDiscovered

`func (o *AssetServiceNow) SetFirstDiscovered(v string)`

SetFirstDiscovered sets FirstDiscovered field to given value.

### HasFirstDiscovered

`func (o *AssetServiceNow) HasFirstDiscovered() bool`

HasFirstDiscovered returns a boolean if a field has been set.

### GetLastDiscovered

`func (o *AssetServiceNow) GetLastDiscovered() string`

GetLastDiscovered returns the LastDiscovered field if non-nil, zero value otherwise.

### GetLastDiscoveredOk

`func (o *AssetServiceNow) GetLastDiscoveredOk() (*string, bool)`

GetLastDiscoveredOk returns a tuple with the LastDiscovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDiscovered

`func (o *AssetServiceNow) SetLastDiscovered(v string)`

SetLastDiscovered sets LastDiscovered field to given value.

### HasLastDiscovered

`func (o *AssetServiceNow) HasLastDiscovered() bool`

HasLastDiscovered returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AssetServiceNow) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AssetServiceNow) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AssetServiceNow) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AssetServiceNow) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetComments

`func (o *AssetServiceNow) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *AssetServiceNow) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *AssetServiceNow) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *AssetServiceNow) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


