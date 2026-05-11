# License

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ActivatedAt** | Pointer to **int64** |  | [optional] 
**LicenseType** | Pointer to **string** |  | [optional] 
**LicenseExpiration** | Pointer to **int64** |  | [optional] 
**LicenseMaxAssets** | Pointer to **int64** |  | [optional] 
**LicenseLiveAssetCount** | Pointer to **int64** |  | [optional] 
**LicenseProjectAssetCount** | Pointer to **int64** |  | [optional] 
**ViaReseller** | Pointer to **bool** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**AddressStreet** | Pointer to **string** |  | [optional] 
**AddressCountry** | Pointer to **string** |  | [optional] 
**AddressCity** | Pointer to **string** |  | [optional] 
**AddressRegion** | Pointer to **string** |  | [optional] 
**AddressPostal** | Pointer to **string** |  | [optional] 
**SubscriptionPeriodStart** | Pointer to **int64** |  | [optional] 
**SubscriptionPeriodEnd** | Pointer to **int64** |  | [optional] 
**SubscriptionCancelAt** | Pointer to **int64** |  | [optional] 
**SubscriptionCanceledAt** | Pointer to **int64** |  | [optional] 
**Settings** | Pointer to **map[string]interface{}** |  | [optional] 
**SsoType** | Pointer to **string** |  | [optional] 
**SsoDomain** | Pointer to **string** |  | [optional] 
**SsoMode** | Pointer to **string** |  | [optional] 
**SsoLoginMessage** | Pointer to **string** |  | [optional] 
**SsoLoginIssuerUrl** | Pointer to **string** |  | [optional] 
**SsoLoginLoginUrl** | Pointer to **string** |  | [optional] 
**SsoLoginLogoutUrl** | Pointer to **string** |  | [optional] 
**SsoDefaultRole** | Pointer to **string** |  | [optional] 
**Partner** | Pointer to **string** |  | [optional] 

## Methods

### NewLicense

`func NewLicense() *License`

NewLicense instantiates a new License object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseWithDefaults

`func NewLicenseWithDefaults() *License`

NewLicenseWithDefaults instantiates a new License object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *License) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *License) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *License) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *License) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *License) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *License) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *License) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *License) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *License) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *License) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *License) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *License) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *License) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *License) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *License) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *License) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActivatedAt

`func (o *License) GetActivatedAt() int64`

GetActivatedAt returns the ActivatedAt field if non-nil, zero value otherwise.

### GetActivatedAtOk

`func (o *License) GetActivatedAtOk() (*int64, bool)`

GetActivatedAtOk returns a tuple with the ActivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedAt

`func (o *License) SetActivatedAt(v int64)`

SetActivatedAt sets ActivatedAt field to given value.

### HasActivatedAt

`func (o *License) HasActivatedAt() bool`

HasActivatedAt returns a boolean if a field has been set.

### GetLicenseType

`func (o *License) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *License) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *License) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *License) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetLicenseExpiration

`func (o *License) GetLicenseExpiration() int64`

GetLicenseExpiration returns the LicenseExpiration field if non-nil, zero value otherwise.

### GetLicenseExpirationOk

`func (o *License) GetLicenseExpirationOk() (*int64, bool)`

GetLicenseExpirationOk returns a tuple with the LicenseExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseExpiration

`func (o *License) SetLicenseExpiration(v int64)`

SetLicenseExpiration sets LicenseExpiration field to given value.

### HasLicenseExpiration

`func (o *License) HasLicenseExpiration() bool`

HasLicenseExpiration returns a boolean if a field has been set.

### GetLicenseMaxAssets

`func (o *License) GetLicenseMaxAssets() int64`

GetLicenseMaxAssets returns the LicenseMaxAssets field if non-nil, zero value otherwise.

### GetLicenseMaxAssetsOk

`func (o *License) GetLicenseMaxAssetsOk() (*int64, bool)`

GetLicenseMaxAssetsOk returns a tuple with the LicenseMaxAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseMaxAssets

`func (o *License) SetLicenseMaxAssets(v int64)`

SetLicenseMaxAssets sets LicenseMaxAssets field to given value.

### HasLicenseMaxAssets

`func (o *License) HasLicenseMaxAssets() bool`

HasLicenseMaxAssets returns a boolean if a field has been set.

### GetLicenseLiveAssetCount

`func (o *License) GetLicenseLiveAssetCount() int64`

GetLicenseLiveAssetCount returns the LicenseLiveAssetCount field if non-nil, zero value otherwise.

### GetLicenseLiveAssetCountOk

`func (o *License) GetLicenseLiveAssetCountOk() (*int64, bool)`

GetLicenseLiveAssetCountOk returns a tuple with the LicenseLiveAssetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseLiveAssetCount

`func (o *License) SetLicenseLiveAssetCount(v int64)`

SetLicenseLiveAssetCount sets LicenseLiveAssetCount field to given value.

### HasLicenseLiveAssetCount

`func (o *License) HasLicenseLiveAssetCount() bool`

HasLicenseLiveAssetCount returns a boolean if a field has been set.

### GetLicenseProjectAssetCount

`func (o *License) GetLicenseProjectAssetCount() int64`

GetLicenseProjectAssetCount returns the LicenseProjectAssetCount field if non-nil, zero value otherwise.

### GetLicenseProjectAssetCountOk

`func (o *License) GetLicenseProjectAssetCountOk() (*int64, bool)`

GetLicenseProjectAssetCountOk returns a tuple with the LicenseProjectAssetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseProjectAssetCount

`func (o *License) SetLicenseProjectAssetCount(v int64)`

SetLicenseProjectAssetCount sets LicenseProjectAssetCount field to given value.

### HasLicenseProjectAssetCount

`func (o *License) HasLicenseProjectAssetCount() bool`

HasLicenseProjectAssetCount returns a boolean if a field has been set.

### GetViaReseller

`func (o *License) GetViaReseller() bool`

GetViaReseller returns the ViaReseller field if non-nil, zero value otherwise.

### GetViaResellerOk

`func (o *License) GetViaResellerOk() (*bool, bool)`

GetViaResellerOk returns a tuple with the ViaReseller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViaReseller

`func (o *License) SetViaReseller(v bool)`

SetViaReseller sets ViaReseller field to given value.

### HasViaReseller

`func (o *License) HasViaReseller() bool`

HasViaReseller returns a boolean if a field has been set.

### GetPhone

`func (o *License) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *License) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *License) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *License) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetAddressStreet

`func (o *License) GetAddressStreet() string`

GetAddressStreet returns the AddressStreet field if non-nil, zero value otherwise.

### GetAddressStreetOk

`func (o *License) GetAddressStreetOk() (*string, bool)`

GetAddressStreetOk returns a tuple with the AddressStreet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressStreet

`func (o *License) SetAddressStreet(v string)`

SetAddressStreet sets AddressStreet field to given value.

### HasAddressStreet

`func (o *License) HasAddressStreet() bool`

HasAddressStreet returns a boolean if a field has been set.

### GetAddressCountry

`func (o *License) GetAddressCountry() string`

GetAddressCountry returns the AddressCountry field if non-nil, zero value otherwise.

### GetAddressCountryOk

`func (o *License) GetAddressCountryOk() (*string, bool)`

GetAddressCountryOk returns a tuple with the AddressCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCountry

`func (o *License) SetAddressCountry(v string)`

SetAddressCountry sets AddressCountry field to given value.

### HasAddressCountry

`func (o *License) HasAddressCountry() bool`

HasAddressCountry returns a boolean if a field has been set.

### GetAddressCity

`func (o *License) GetAddressCity() string`

GetAddressCity returns the AddressCity field if non-nil, zero value otherwise.

### GetAddressCityOk

`func (o *License) GetAddressCityOk() (*string, bool)`

GetAddressCityOk returns a tuple with the AddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCity

`func (o *License) SetAddressCity(v string)`

SetAddressCity sets AddressCity field to given value.

### HasAddressCity

`func (o *License) HasAddressCity() bool`

HasAddressCity returns a boolean if a field has been set.

### GetAddressRegion

`func (o *License) GetAddressRegion() string`

GetAddressRegion returns the AddressRegion field if non-nil, zero value otherwise.

### GetAddressRegionOk

`func (o *License) GetAddressRegionOk() (*string, bool)`

GetAddressRegionOk returns a tuple with the AddressRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressRegion

`func (o *License) SetAddressRegion(v string)`

SetAddressRegion sets AddressRegion field to given value.

### HasAddressRegion

`func (o *License) HasAddressRegion() bool`

HasAddressRegion returns a boolean if a field has been set.

### GetAddressPostal

`func (o *License) GetAddressPostal() string`

GetAddressPostal returns the AddressPostal field if non-nil, zero value otherwise.

### GetAddressPostalOk

`func (o *License) GetAddressPostalOk() (*string, bool)`

GetAddressPostalOk returns a tuple with the AddressPostal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressPostal

`func (o *License) SetAddressPostal(v string)`

SetAddressPostal sets AddressPostal field to given value.

### HasAddressPostal

`func (o *License) HasAddressPostal() bool`

HasAddressPostal returns a boolean if a field has been set.

### GetSubscriptionPeriodStart

`func (o *License) GetSubscriptionPeriodStart() int64`

GetSubscriptionPeriodStart returns the SubscriptionPeriodStart field if non-nil, zero value otherwise.

### GetSubscriptionPeriodStartOk

`func (o *License) GetSubscriptionPeriodStartOk() (*int64, bool)`

GetSubscriptionPeriodStartOk returns a tuple with the SubscriptionPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPeriodStart

`func (o *License) SetSubscriptionPeriodStart(v int64)`

SetSubscriptionPeriodStart sets SubscriptionPeriodStart field to given value.

### HasSubscriptionPeriodStart

`func (o *License) HasSubscriptionPeriodStart() bool`

HasSubscriptionPeriodStart returns a boolean if a field has been set.

### GetSubscriptionPeriodEnd

`func (o *License) GetSubscriptionPeriodEnd() int64`

GetSubscriptionPeriodEnd returns the SubscriptionPeriodEnd field if non-nil, zero value otherwise.

### GetSubscriptionPeriodEndOk

`func (o *License) GetSubscriptionPeriodEndOk() (*int64, bool)`

GetSubscriptionPeriodEndOk returns a tuple with the SubscriptionPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPeriodEnd

`func (o *License) SetSubscriptionPeriodEnd(v int64)`

SetSubscriptionPeriodEnd sets SubscriptionPeriodEnd field to given value.

### HasSubscriptionPeriodEnd

`func (o *License) HasSubscriptionPeriodEnd() bool`

HasSubscriptionPeriodEnd returns a boolean if a field has been set.

### GetSubscriptionCancelAt

`func (o *License) GetSubscriptionCancelAt() int64`

GetSubscriptionCancelAt returns the SubscriptionCancelAt field if non-nil, zero value otherwise.

### GetSubscriptionCancelAtOk

`func (o *License) GetSubscriptionCancelAtOk() (*int64, bool)`

GetSubscriptionCancelAtOk returns a tuple with the SubscriptionCancelAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionCancelAt

`func (o *License) SetSubscriptionCancelAt(v int64)`

SetSubscriptionCancelAt sets SubscriptionCancelAt field to given value.

### HasSubscriptionCancelAt

`func (o *License) HasSubscriptionCancelAt() bool`

HasSubscriptionCancelAt returns a boolean if a field has been set.

### GetSubscriptionCanceledAt

`func (o *License) GetSubscriptionCanceledAt() int64`

GetSubscriptionCanceledAt returns the SubscriptionCanceledAt field if non-nil, zero value otherwise.

### GetSubscriptionCanceledAtOk

`func (o *License) GetSubscriptionCanceledAtOk() (*int64, bool)`

GetSubscriptionCanceledAtOk returns a tuple with the SubscriptionCanceledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionCanceledAt

`func (o *License) SetSubscriptionCanceledAt(v int64)`

SetSubscriptionCanceledAt sets SubscriptionCanceledAt field to given value.

### HasSubscriptionCanceledAt

`func (o *License) HasSubscriptionCanceledAt() bool`

HasSubscriptionCanceledAt returns a boolean if a field has been set.

### GetSettings

`func (o *License) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *License) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *License) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *License) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetSsoType

`func (o *License) GetSsoType() string`

GetSsoType returns the SsoType field if non-nil, zero value otherwise.

### GetSsoTypeOk

`func (o *License) GetSsoTypeOk() (*string, bool)`

GetSsoTypeOk returns a tuple with the SsoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoType

`func (o *License) SetSsoType(v string)`

SetSsoType sets SsoType field to given value.

### HasSsoType

`func (o *License) HasSsoType() bool`

HasSsoType returns a boolean if a field has been set.

### GetSsoDomain

`func (o *License) GetSsoDomain() string`

GetSsoDomain returns the SsoDomain field if non-nil, zero value otherwise.

### GetSsoDomainOk

`func (o *License) GetSsoDomainOk() (*string, bool)`

GetSsoDomainOk returns a tuple with the SsoDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoDomain

`func (o *License) SetSsoDomain(v string)`

SetSsoDomain sets SsoDomain field to given value.

### HasSsoDomain

`func (o *License) HasSsoDomain() bool`

HasSsoDomain returns a boolean if a field has been set.

### GetSsoMode

`func (o *License) GetSsoMode() string`

GetSsoMode returns the SsoMode field if non-nil, zero value otherwise.

### GetSsoModeOk

`func (o *License) GetSsoModeOk() (*string, bool)`

GetSsoModeOk returns a tuple with the SsoMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoMode

`func (o *License) SetSsoMode(v string)`

SetSsoMode sets SsoMode field to given value.

### HasSsoMode

`func (o *License) HasSsoMode() bool`

HasSsoMode returns a boolean if a field has been set.

### GetSsoLoginMessage

`func (o *License) GetSsoLoginMessage() string`

GetSsoLoginMessage returns the SsoLoginMessage field if non-nil, zero value otherwise.

### GetSsoLoginMessageOk

`func (o *License) GetSsoLoginMessageOk() (*string, bool)`

GetSsoLoginMessageOk returns a tuple with the SsoLoginMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoLoginMessage

`func (o *License) SetSsoLoginMessage(v string)`

SetSsoLoginMessage sets SsoLoginMessage field to given value.

### HasSsoLoginMessage

`func (o *License) HasSsoLoginMessage() bool`

HasSsoLoginMessage returns a boolean if a field has been set.

### GetSsoLoginIssuerUrl

`func (o *License) GetSsoLoginIssuerUrl() string`

GetSsoLoginIssuerUrl returns the SsoLoginIssuerUrl field if non-nil, zero value otherwise.

### GetSsoLoginIssuerUrlOk

`func (o *License) GetSsoLoginIssuerUrlOk() (*string, bool)`

GetSsoLoginIssuerUrlOk returns a tuple with the SsoLoginIssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoLoginIssuerUrl

`func (o *License) SetSsoLoginIssuerUrl(v string)`

SetSsoLoginIssuerUrl sets SsoLoginIssuerUrl field to given value.

### HasSsoLoginIssuerUrl

`func (o *License) HasSsoLoginIssuerUrl() bool`

HasSsoLoginIssuerUrl returns a boolean if a field has been set.

### GetSsoLoginLoginUrl

`func (o *License) GetSsoLoginLoginUrl() string`

GetSsoLoginLoginUrl returns the SsoLoginLoginUrl field if non-nil, zero value otherwise.

### GetSsoLoginLoginUrlOk

`func (o *License) GetSsoLoginLoginUrlOk() (*string, bool)`

GetSsoLoginLoginUrlOk returns a tuple with the SsoLoginLoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoLoginLoginUrl

`func (o *License) SetSsoLoginLoginUrl(v string)`

SetSsoLoginLoginUrl sets SsoLoginLoginUrl field to given value.

### HasSsoLoginLoginUrl

`func (o *License) HasSsoLoginLoginUrl() bool`

HasSsoLoginLoginUrl returns a boolean if a field has been set.

### GetSsoLoginLogoutUrl

`func (o *License) GetSsoLoginLogoutUrl() string`

GetSsoLoginLogoutUrl returns the SsoLoginLogoutUrl field if non-nil, zero value otherwise.

### GetSsoLoginLogoutUrlOk

`func (o *License) GetSsoLoginLogoutUrlOk() (*string, bool)`

GetSsoLoginLogoutUrlOk returns a tuple with the SsoLoginLogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoLoginLogoutUrl

`func (o *License) SetSsoLoginLogoutUrl(v string)`

SetSsoLoginLogoutUrl sets SsoLoginLogoutUrl field to given value.

### HasSsoLoginLogoutUrl

`func (o *License) HasSsoLoginLogoutUrl() bool`

HasSsoLoginLogoutUrl returns a boolean if a field has been set.

### GetSsoDefaultRole

`func (o *License) GetSsoDefaultRole() string`

GetSsoDefaultRole returns the SsoDefaultRole field if non-nil, zero value otherwise.

### GetSsoDefaultRoleOk

`func (o *License) GetSsoDefaultRoleOk() (*string, bool)`

GetSsoDefaultRoleOk returns a tuple with the SsoDefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoDefaultRole

`func (o *License) SetSsoDefaultRole(v string)`

SetSsoDefaultRole sets SsoDefaultRole field to given value.

### HasSsoDefaultRole

`func (o *License) HasSsoDefaultRole() bool`

HasSsoDefaultRole returns a boolean if a field has been set.

### GetPartner

`func (o *License) GetPartner() string`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *License) GetPartnerOk() (*string, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *License) SetPartner(v string)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *License) HasPartner() bool`

HasPartner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


