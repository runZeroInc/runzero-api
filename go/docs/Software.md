# Software

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SoftwareId** | Pointer to **string** |  | [optional] 
**SoftwareAssetId** | Pointer to **string** |  | [optional] 
**SoftwareOrganizationId** | Pointer to **string** |  | [optional] 
**SoftwareSourceId** | Pointer to **int32** |  | [optional] 
**SoftwareCreatedAt** | Pointer to **int64** |  | [optional] 
**SoftwareUpdatedAt** | Pointer to **int64** |  | [optional] 
**SoftwareServiceAddress** | Pointer to **NullableString** |  | [optional] 
**SoftwareServiceTransport** | Pointer to **string** |  | [optional] 
**SoftwareServicePort** | Pointer to **int64** |  | [optional] 
**SoftwareInstalledAt** | Pointer to **int64** |  | [optional] 
**SoftwareInstalledSize** | Pointer to **int64** |  | [optional] 
**SoftwareInstalledFrom** | Pointer to **string** |  | [optional] 
**SoftwareAppId** | Pointer to **string** |  | [optional] 
**SoftwarePart** | Pointer to **string** |  | [optional] 
**SoftwareVendor** | Pointer to **string** |  | [optional] 
**SoftwareProduct** | Pointer to **string** |  | [optional] 
**SoftwareVersion** | Pointer to **string** |  | [optional] 
**SoftwareUpdate** | Pointer to **string** |  | [optional] 
**SoftwareEdition** | Pointer to **string** |  | [optional] 
**SoftwareLanguage** | Pointer to **string** |  | [optional] 
**SoftwareSwEdition** | Pointer to **string** |  | [optional] 
**SoftwareTargetSw** | Pointer to **string** |  | [optional] 
**SoftwareTargetHw** | Pointer to **string** |  | [optional] 
**SoftwareOther** | Pointer to **string** |  | [optional] 
**SoftwareCpe23** | Pointer to **string** |  | [optional] 
**SoftwareAttributes** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSoftware

`func NewSoftware(id string, ) *Software`

NewSoftware instantiates a new Software object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareWithDefaults

`func NewSoftwareWithDefaults() *Software`

NewSoftwareWithDefaults instantiates a new Software object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Software) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Software) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Software) SetId(v string)`

SetId sets Id field to given value.


### GetSoftwareId

`func (o *Software) GetSoftwareId() string`

GetSoftwareId returns the SoftwareId field if non-nil, zero value otherwise.

### GetSoftwareIdOk

`func (o *Software) GetSoftwareIdOk() (*string, bool)`

GetSoftwareIdOk returns a tuple with the SoftwareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareId

`func (o *Software) SetSoftwareId(v string)`

SetSoftwareId sets SoftwareId field to given value.

### HasSoftwareId

`func (o *Software) HasSoftwareId() bool`

HasSoftwareId returns a boolean if a field has been set.

### GetSoftwareAssetId

`func (o *Software) GetSoftwareAssetId() string`

GetSoftwareAssetId returns the SoftwareAssetId field if non-nil, zero value otherwise.

### GetSoftwareAssetIdOk

`func (o *Software) GetSoftwareAssetIdOk() (*string, bool)`

GetSoftwareAssetIdOk returns a tuple with the SoftwareAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareAssetId

`func (o *Software) SetSoftwareAssetId(v string)`

SetSoftwareAssetId sets SoftwareAssetId field to given value.

### HasSoftwareAssetId

`func (o *Software) HasSoftwareAssetId() bool`

HasSoftwareAssetId returns a boolean if a field has been set.

### GetSoftwareOrganizationId

`func (o *Software) GetSoftwareOrganizationId() string`

GetSoftwareOrganizationId returns the SoftwareOrganizationId field if non-nil, zero value otherwise.

### GetSoftwareOrganizationIdOk

`func (o *Software) GetSoftwareOrganizationIdOk() (*string, bool)`

GetSoftwareOrganizationIdOk returns a tuple with the SoftwareOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareOrganizationId

`func (o *Software) SetSoftwareOrganizationId(v string)`

SetSoftwareOrganizationId sets SoftwareOrganizationId field to given value.

### HasSoftwareOrganizationId

`func (o *Software) HasSoftwareOrganizationId() bool`

HasSoftwareOrganizationId returns a boolean if a field has been set.

### GetSoftwareSourceId

`func (o *Software) GetSoftwareSourceId() int32`

GetSoftwareSourceId returns the SoftwareSourceId field if non-nil, zero value otherwise.

### GetSoftwareSourceIdOk

`func (o *Software) GetSoftwareSourceIdOk() (*int32, bool)`

GetSoftwareSourceIdOk returns a tuple with the SoftwareSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareSourceId

`func (o *Software) SetSoftwareSourceId(v int32)`

SetSoftwareSourceId sets SoftwareSourceId field to given value.

### HasSoftwareSourceId

`func (o *Software) HasSoftwareSourceId() bool`

HasSoftwareSourceId returns a boolean if a field has been set.

### GetSoftwareCreatedAt

`func (o *Software) GetSoftwareCreatedAt() int64`

GetSoftwareCreatedAt returns the SoftwareCreatedAt field if non-nil, zero value otherwise.

### GetSoftwareCreatedAtOk

`func (o *Software) GetSoftwareCreatedAtOk() (*int64, bool)`

GetSoftwareCreatedAtOk returns a tuple with the SoftwareCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareCreatedAt

`func (o *Software) SetSoftwareCreatedAt(v int64)`

SetSoftwareCreatedAt sets SoftwareCreatedAt field to given value.

### HasSoftwareCreatedAt

`func (o *Software) HasSoftwareCreatedAt() bool`

HasSoftwareCreatedAt returns a boolean if a field has been set.

### GetSoftwareUpdatedAt

`func (o *Software) GetSoftwareUpdatedAt() int64`

GetSoftwareUpdatedAt returns the SoftwareUpdatedAt field if non-nil, zero value otherwise.

### GetSoftwareUpdatedAtOk

`func (o *Software) GetSoftwareUpdatedAtOk() (*int64, bool)`

GetSoftwareUpdatedAtOk returns a tuple with the SoftwareUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpdatedAt

`func (o *Software) SetSoftwareUpdatedAt(v int64)`

SetSoftwareUpdatedAt sets SoftwareUpdatedAt field to given value.

### HasSoftwareUpdatedAt

`func (o *Software) HasSoftwareUpdatedAt() bool`

HasSoftwareUpdatedAt returns a boolean if a field has been set.

### GetSoftwareServiceAddress

`func (o *Software) GetSoftwareServiceAddress() string`

GetSoftwareServiceAddress returns the SoftwareServiceAddress field if non-nil, zero value otherwise.

### GetSoftwareServiceAddressOk

`func (o *Software) GetSoftwareServiceAddressOk() (*string, bool)`

GetSoftwareServiceAddressOk returns a tuple with the SoftwareServiceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareServiceAddress

`func (o *Software) SetSoftwareServiceAddress(v string)`

SetSoftwareServiceAddress sets SoftwareServiceAddress field to given value.

### HasSoftwareServiceAddress

`func (o *Software) HasSoftwareServiceAddress() bool`

HasSoftwareServiceAddress returns a boolean if a field has been set.

### SetSoftwareServiceAddressNil

`func (o *Software) SetSoftwareServiceAddressNil(b bool)`

 SetSoftwareServiceAddressNil sets the value for SoftwareServiceAddress to be an explicit nil

### UnsetSoftwareServiceAddress
`func (o *Software) UnsetSoftwareServiceAddress()`

UnsetSoftwareServiceAddress ensures that no value is present for SoftwareServiceAddress, not even an explicit nil
### GetSoftwareServiceTransport

`func (o *Software) GetSoftwareServiceTransport() string`

GetSoftwareServiceTransport returns the SoftwareServiceTransport field if non-nil, zero value otherwise.

### GetSoftwareServiceTransportOk

`func (o *Software) GetSoftwareServiceTransportOk() (*string, bool)`

GetSoftwareServiceTransportOk returns a tuple with the SoftwareServiceTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareServiceTransport

`func (o *Software) SetSoftwareServiceTransport(v string)`

SetSoftwareServiceTransport sets SoftwareServiceTransport field to given value.

### HasSoftwareServiceTransport

`func (o *Software) HasSoftwareServiceTransport() bool`

HasSoftwareServiceTransport returns a boolean if a field has been set.

### GetSoftwareServicePort

`func (o *Software) GetSoftwareServicePort() int64`

GetSoftwareServicePort returns the SoftwareServicePort field if non-nil, zero value otherwise.

### GetSoftwareServicePortOk

`func (o *Software) GetSoftwareServicePortOk() (*int64, bool)`

GetSoftwareServicePortOk returns a tuple with the SoftwareServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareServicePort

`func (o *Software) SetSoftwareServicePort(v int64)`

SetSoftwareServicePort sets SoftwareServicePort field to given value.

### HasSoftwareServicePort

`func (o *Software) HasSoftwareServicePort() bool`

HasSoftwareServicePort returns a boolean if a field has been set.

### GetSoftwareInstalledAt

`func (o *Software) GetSoftwareInstalledAt() int64`

GetSoftwareInstalledAt returns the SoftwareInstalledAt field if non-nil, zero value otherwise.

### GetSoftwareInstalledAtOk

`func (o *Software) GetSoftwareInstalledAtOk() (*int64, bool)`

GetSoftwareInstalledAtOk returns a tuple with the SoftwareInstalledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareInstalledAt

`func (o *Software) SetSoftwareInstalledAt(v int64)`

SetSoftwareInstalledAt sets SoftwareInstalledAt field to given value.

### HasSoftwareInstalledAt

`func (o *Software) HasSoftwareInstalledAt() bool`

HasSoftwareInstalledAt returns a boolean if a field has been set.

### GetSoftwareInstalledSize

`func (o *Software) GetSoftwareInstalledSize() int64`

GetSoftwareInstalledSize returns the SoftwareInstalledSize field if non-nil, zero value otherwise.

### GetSoftwareInstalledSizeOk

`func (o *Software) GetSoftwareInstalledSizeOk() (*int64, bool)`

GetSoftwareInstalledSizeOk returns a tuple with the SoftwareInstalledSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareInstalledSize

`func (o *Software) SetSoftwareInstalledSize(v int64)`

SetSoftwareInstalledSize sets SoftwareInstalledSize field to given value.

### HasSoftwareInstalledSize

`func (o *Software) HasSoftwareInstalledSize() bool`

HasSoftwareInstalledSize returns a boolean if a field has been set.

### GetSoftwareInstalledFrom

`func (o *Software) GetSoftwareInstalledFrom() string`

GetSoftwareInstalledFrom returns the SoftwareInstalledFrom field if non-nil, zero value otherwise.

### GetSoftwareInstalledFromOk

`func (o *Software) GetSoftwareInstalledFromOk() (*string, bool)`

GetSoftwareInstalledFromOk returns a tuple with the SoftwareInstalledFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareInstalledFrom

`func (o *Software) SetSoftwareInstalledFrom(v string)`

SetSoftwareInstalledFrom sets SoftwareInstalledFrom field to given value.

### HasSoftwareInstalledFrom

`func (o *Software) HasSoftwareInstalledFrom() bool`

HasSoftwareInstalledFrom returns a boolean if a field has been set.

### GetSoftwareAppId

`func (o *Software) GetSoftwareAppId() string`

GetSoftwareAppId returns the SoftwareAppId field if non-nil, zero value otherwise.

### GetSoftwareAppIdOk

`func (o *Software) GetSoftwareAppIdOk() (*string, bool)`

GetSoftwareAppIdOk returns a tuple with the SoftwareAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareAppId

`func (o *Software) SetSoftwareAppId(v string)`

SetSoftwareAppId sets SoftwareAppId field to given value.

### HasSoftwareAppId

`func (o *Software) HasSoftwareAppId() bool`

HasSoftwareAppId returns a boolean if a field has been set.

### GetSoftwarePart

`func (o *Software) GetSoftwarePart() string`

GetSoftwarePart returns the SoftwarePart field if non-nil, zero value otherwise.

### GetSoftwarePartOk

`func (o *Software) GetSoftwarePartOk() (*string, bool)`

GetSoftwarePartOk returns a tuple with the SoftwarePart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwarePart

`func (o *Software) SetSoftwarePart(v string)`

SetSoftwarePart sets SoftwarePart field to given value.

### HasSoftwarePart

`func (o *Software) HasSoftwarePart() bool`

HasSoftwarePart returns a boolean if a field has been set.

### GetSoftwareVendor

`func (o *Software) GetSoftwareVendor() string`

GetSoftwareVendor returns the SoftwareVendor field if non-nil, zero value otherwise.

### GetSoftwareVendorOk

`func (o *Software) GetSoftwareVendorOk() (*string, bool)`

GetSoftwareVendorOk returns a tuple with the SoftwareVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVendor

`func (o *Software) SetSoftwareVendor(v string)`

SetSoftwareVendor sets SoftwareVendor field to given value.

### HasSoftwareVendor

`func (o *Software) HasSoftwareVendor() bool`

HasSoftwareVendor returns a boolean if a field has been set.

### GetSoftwareProduct

`func (o *Software) GetSoftwareProduct() string`

GetSoftwareProduct returns the SoftwareProduct field if non-nil, zero value otherwise.

### GetSoftwareProductOk

`func (o *Software) GetSoftwareProductOk() (*string, bool)`

GetSoftwareProductOk returns a tuple with the SoftwareProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareProduct

`func (o *Software) SetSoftwareProduct(v string)`

SetSoftwareProduct sets SoftwareProduct field to given value.

### HasSoftwareProduct

`func (o *Software) HasSoftwareProduct() bool`

HasSoftwareProduct returns a boolean if a field has been set.

### GetSoftwareVersion

`func (o *Software) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *Software) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *Software) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *Software) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### GetSoftwareUpdate

`func (o *Software) GetSoftwareUpdate() string`

GetSoftwareUpdate returns the SoftwareUpdate field if non-nil, zero value otherwise.

### GetSoftwareUpdateOk

`func (o *Software) GetSoftwareUpdateOk() (*string, bool)`

GetSoftwareUpdateOk returns a tuple with the SoftwareUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpdate

`func (o *Software) SetSoftwareUpdate(v string)`

SetSoftwareUpdate sets SoftwareUpdate field to given value.

### HasSoftwareUpdate

`func (o *Software) HasSoftwareUpdate() bool`

HasSoftwareUpdate returns a boolean if a field has been set.

### GetSoftwareEdition

`func (o *Software) GetSoftwareEdition() string`

GetSoftwareEdition returns the SoftwareEdition field if non-nil, zero value otherwise.

### GetSoftwareEditionOk

`func (o *Software) GetSoftwareEditionOk() (*string, bool)`

GetSoftwareEditionOk returns a tuple with the SoftwareEdition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareEdition

`func (o *Software) SetSoftwareEdition(v string)`

SetSoftwareEdition sets SoftwareEdition field to given value.

### HasSoftwareEdition

`func (o *Software) HasSoftwareEdition() bool`

HasSoftwareEdition returns a boolean if a field has been set.

### GetSoftwareLanguage

`func (o *Software) GetSoftwareLanguage() string`

GetSoftwareLanguage returns the SoftwareLanguage field if non-nil, zero value otherwise.

### GetSoftwareLanguageOk

`func (o *Software) GetSoftwareLanguageOk() (*string, bool)`

GetSoftwareLanguageOk returns a tuple with the SoftwareLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareLanguage

`func (o *Software) SetSoftwareLanguage(v string)`

SetSoftwareLanguage sets SoftwareLanguage field to given value.

### HasSoftwareLanguage

`func (o *Software) HasSoftwareLanguage() bool`

HasSoftwareLanguage returns a boolean if a field has been set.

### GetSoftwareSwEdition

`func (o *Software) GetSoftwareSwEdition() string`

GetSoftwareSwEdition returns the SoftwareSwEdition field if non-nil, zero value otherwise.

### GetSoftwareSwEditionOk

`func (o *Software) GetSoftwareSwEditionOk() (*string, bool)`

GetSoftwareSwEditionOk returns a tuple with the SoftwareSwEdition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareSwEdition

`func (o *Software) SetSoftwareSwEdition(v string)`

SetSoftwareSwEdition sets SoftwareSwEdition field to given value.

### HasSoftwareSwEdition

`func (o *Software) HasSoftwareSwEdition() bool`

HasSoftwareSwEdition returns a boolean if a field has been set.

### GetSoftwareTargetSw

`func (o *Software) GetSoftwareTargetSw() string`

GetSoftwareTargetSw returns the SoftwareTargetSw field if non-nil, zero value otherwise.

### GetSoftwareTargetSwOk

`func (o *Software) GetSoftwareTargetSwOk() (*string, bool)`

GetSoftwareTargetSwOk returns a tuple with the SoftwareTargetSw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTargetSw

`func (o *Software) SetSoftwareTargetSw(v string)`

SetSoftwareTargetSw sets SoftwareTargetSw field to given value.

### HasSoftwareTargetSw

`func (o *Software) HasSoftwareTargetSw() bool`

HasSoftwareTargetSw returns a boolean if a field has been set.

### GetSoftwareTargetHw

`func (o *Software) GetSoftwareTargetHw() string`

GetSoftwareTargetHw returns the SoftwareTargetHw field if non-nil, zero value otherwise.

### GetSoftwareTargetHwOk

`func (o *Software) GetSoftwareTargetHwOk() (*string, bool)`

GetSoftwareTargetHwOk returns a tuple with the SoftwareTargetHw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTargetHw

`func (o *Software) SetSoftwareTargetHw(v string)`

SetSoftwareTargetHw sets SoftwareTargetHw field to given value.

### HasSoftwareTargetHw

`func (o *Software) HasSoftwareTargetHw() bool`

HasSoftwareTargetHw returns a boolean if a field has been set.

### GetSoftwareOther

`func (o *Software) GetSoftwareOther() string`

GetSoftwareOther returns the SoftwareOther field if non-nil, zero value otherwise.

### GetSoftwareOtherOk

`func (o *Software) GetSoftwareOtherOk() (*string, bool)`

GetSoftwareOtherOk returns a tuple with the SoftwareOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareOther

`func (o *Software) SetSoftwareOther(v string)`

SetSoftwareOther sets SoftwareOther field to given value.

### HasSoftwareOther

`func (o *Software) HasSoftwareOther() bool`

HasSoftwareOther returns a boolean if a field has been set.

### GetSoftwareCpe23

`func (o *Software) GetSoftwareCpe23() string`

GetSoftwareCpe23 returns the SoftwareCpe23 field if non-nil, zero value otherwise.

### GetSoftwareCpe23Ok

`func (o *Software) GetSoftwareCpe23Ok() (*string, bool)`

GetSoftwareCpe23Ok returns a tuple with the SoftwareCpe23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareCpe23

`func (o *Software) SetSoftwareCpe23(v string)`

SetSoftwareCpe23 sets SoftwareCpe23 field to given value.

### HasSoftwareCpe23

`func (o *Software) HasSoftwareCpe23() bool`

HasSoftwareCpe23 returns a boolean if a field has been set.

### GetSoftwareAttributes

`func (o *Software) GetSoftwareAttributes() map[string]string`

GetSoftwareAttributes returns the SoftwareAttributes field if non-nil, zero value otherwise.

### GetSoftwareAttributesOk

`func (o *Software) GetSoftwareAttributesOk() (*map[string]string, bool)`

GetSoftwareAttributesOk returns a tuple with the SoftwareAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareAttributes

`func (o *Software) SetSoftwareAttributes(v map[string]string)`

SetSoftwareAttributes sets SoftwareAttributes field to given value.

### HasSoftwareAttributes

`func (o *Software) HasSoftwareAttributes() bool`

HasSoftwareAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


