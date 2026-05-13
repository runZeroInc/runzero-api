# ImportAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Any value which can uniquely identify the asset within the custom integration. | 
**RunZeroID** | Pointer to **string** | The unique identifier of the runZero asset to merge into. | [optional] 
**NetworkInterfaces** | Pointer to [**[]NetworkInterface**](NetworkInterface.md) | The asset&#39;s networking configuration. | [optional] 
**Hostnames** | Pointer to **[]string** | Represents hostnames the asset is assigned or reachable at. These can be fully-qualified hostnames with the domain name, or a short hostname. | [optional] 
**Domain** | Pointer to **string** | Represents a single domain name which could be applied to all non-fqdns in the hostnames field. | [optional] 
**FirstSeenTS** | Pointer to **time.Time** | Represents the earliest time the asset was seen by the custom integration reporting it, using a date string as defined by RFC 3339, section 5.6. | [optional] 
**Os** | Pointer to **string** | The name of the asset&#39;s operating system. It is advisable to keep the data clean by normalizing to existing values when possible. | [optional] 
**OsVersion** | Pointer to **string** | The version of the asset&#39;s operating system. It is advisable to keep the data clean by normalizing to existing values when possible. | [optional] 
**Manufacturer** | Pointer to **string** | The manufacturer of the operating system of the asset. It is advisable to keep the data clean by normalizing to existing values when possible. | [optional] 
**Model** | Pointer to **string** | The hardware model of the asset. It is advisable to keep the data clean by normalizing to existing values when possible. | [optional] 
**Tags** | Pointer to **[]string** | Arbitrary string tags applied to the asset. | [optional] 
**DeviceType** | Pointer to **string** |  | [optional] 
**CustomAttributes** | Pointer to [**map[string]CustomAttributesValue**](CustomAttributesValue.md) | Flat map of arbitrary string key/value pairs representing custom attribute data not described in properties above. Note the maximum number of keys and length of values. Additionally, property names may only be 256 characters long. | [optional] 

## Methods

### NewImportAsset

`func NewImportAsset(id string, ) *ImportAsset`

NewImportAsset instantiates a new ImportAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportAssetWithDefaults

`func NewImportAssetWithDefaults() *ImportAsset`

NewImportAssetWithDefaults instantiates a new ImportAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImportAsset) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImportAsset) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImportAsset) SetId(v string)`

SetId sets Id field to given value.


### GetRunZeroID

`func (o *ImportAsset) GetRunZeroID() string`

GetRunZeroID returns the RunZeroID field if non-nil, zero value otherwise.

### GetRunZeroIDOk

`func (o *ImportAsset) GetRunZeroIDOk() (*string, bool)`

GetRunZeroIDOk returns a tuple with the RunZeroID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunZeroID

`func (o *ImportAsset) SetRunZeroID(v string)`

SetRunZeroID sets RunZeroID field to given value.

### HasRunZeroID

`func (o *ImportAsset) HasRunZeroID() bool`

HasRunZeroID returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *ImportAsset) GetNetworkInterfaces() []NetworkInterface`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *ImportAsset) GetNetworkInterfacesOk() (*[]NetworkInterface, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *ImportAsset) SetNetworkInterfaces(v []NetworkInterface)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *ImportAsset) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetHostnames

`func (o *ImportAsset) GetHostnames() []string`

GetHostnames returns the Hostnames field if non-nil, zero value otherwise.

### GetHostnamesOk

`func (o *ImportAsset) GetHostnamesOk() (*[]string, bool)`

GetHostnamesOk returns a tuple with the Hostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnames

`func (o *ImportAsset) SetHostnames(v []string)`

SetHostnames sets Hostnames field to given value.

### HasHostnames

`func (o *ImportAsset) HasHostnames() bool`

HasHostnames returns a boolean if a field has been set.

### GetDomain

`func (o *ImportAsset) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ImportAsset) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ImportAsset) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ImportAsset) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetFirstSeenTS

`func (o *ImportAsset) GetFirstSeenTS() time.Time`

GetFirstSeenTS returns the FirstSeenTS field if non-nil, zero value otherwise.

### GetFirstSeenTSOk

`func (o *ImportAsset) GetFirstSeenTSOk() (*time.Time, bool)`

GetFirstSeenTSOk returns a tuple with the FirstSeenTS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenTS

`func (o *ImportAsset) SetFirstSeenTS(v time.Time)`

SetFirstSeenTS sets FirstSeenTS field to given value.

### HasFirstSeenTS

`func (o *ImportAsset) HasFirstSeenTS() bool`

HasFirstSeenTS returns a boolean if a field has been set.

### GetOs

`func (o *ImportAsset) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *ImportAsset) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *ImportAsset) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *ImportAsset) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetOsVersion

`func (o *ImportAsset) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *ImportAsset) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *ImportAsset) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *ImportAsset) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetManufacturer

`func (o *ImportAsset) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *ImportAsset) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *ImportAsset) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *ImportAsset) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetModel

`func (o *ImportAsset) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ImportAsset) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ImportAsset) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ImportAsset) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetTags

`func (o *ImportAsset) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ImportAsset) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ImportAsset) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ImportAsset) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDeviceType

`func (o *ImportAsset) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *ImportAsset) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *ImportAsset) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *ImportAsset) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetCustomAttributes

`func (o *ImportAsset) GetCustomAttributes() map[string]CustomAttributesValue`

GetCustomAttributes returns the CustomAttributes field if non-nil, zero value otherwise.

### GetCustomAttributesOk

`func (o *ImportAsset) GetCustomAttributesOk() (*map[string]CustomAttributesValue, bool)`

GetCustomAttributesOk returns a tuple with the CustomAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributes

`func (o *ImportAsset) SetCustomAttributes(v map[string]CustomAttributesValue)`

SetCustomAttributes sets CustomAttributes field to given value.

### HasCustomAttributes

`func (o *ImportAsset) HasCustomAttributes() bool`

HasCustomAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


