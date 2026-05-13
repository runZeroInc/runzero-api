# Certificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedAt** | Pointer to **int64** | When the certificate was last seen in a response from a service, and updated. | [optional] 
**OrganizationId** | **string** |  | 
**Names** | Pointer to **[]string** | An assembled list of all names in the certificate. | [optional] 
**SelfSigned** | Pointer to **bool** | Whether the certificate appears to be self-signed based on subject and authority. | [optional] 
**Hidden** | Pointer to **bool** | Whether the certificate has been hidden from the default certificates view. | [optional] 
**Serial** | Pointer to **string** | The serial number of the certificate. | [optional] 
**ValidityStart** | Pointer to **int64** | When the certificate becomes valid. | [optional] 
**ValidityEnd** | Pointer to **int64** | When the certificate ceases to be valid. | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 
**PublicKeyBits** | Pointer to **int32** |  | [optional] 
**PublicKeyAlgorithm** | Pointer to **string** |  | [optional] 
**Signature** | Pointer to **string** |  | [optional] 
**SignatureAlgorithm** | Pointer to **string** |  | [optional] 
**FpBkhash** | Pointer to **string** | The badkeys hash. | [optional] 
**FpSha1** | Pointer to **string** | The SHA1 hash of the certificate. | [optional] 
**FpSha256** | Pointer to **string** | The SHA256 hash of the certificate. | [optional] 
**FpMd5** | Pointer to **string** | The MD5 hash of the certificate (for SSH). | [optional] 
**Subject** | Pointer to **string** | The subject of the certificate. | [optional] 
**Cn** | Pointer to **string** | The Common Name field from the certificate (no longer used by web browsers). | [optional] 
**Version** | Pointer to **int32** | The version of the certificate. | [optional] 
**Issuer** | Pointer to **string** | The authority which issued the certificate. | [optional] 
**SubjectKeyId** | Pointer to **string** | The key ID of the subject of the certificate. | [optional] 
**AuthorityKeyId** | Pointer to **string** | The key ID of the authority which signed the certificate. | [optional] 
**OcspServer** | Pointer to **[]string** | Zero or more OCSP server URLs. | [optional] 
**CrlDistributionPoints** | Pointer to **[]string** | Zero or more URLs of CRLs. | [optional] 
**IssuingCertificateUrl** | Pointer to **[]string** | Zero or more URLs where the issuing certificate can be found. | [optional] 
**IsCa** | Pointer to **bool** | Whether the certificate claims to be a Certificate Authority. | [optional] 
**KeyUsage** | Pointer to **[]string** | Valid purposes the certificate&#39;s key can be used for. | [optional] 
**ExtKeyUsage** | Pointer to **[]string** | Additional purposes the certificate&#39;s key can be used for. | [optional] 
**SanDnsNames** | Pointer to **[]string** | Subject Alternative Name hostnames. | [optional] 
**SanIpAddresses** | Pointer to **[]string** | Subject Alternative Name IP addresses. | [optional] 
**SanEmailAddresses** | Pointer to **[]string** | Subject Alternative Name email addresses. | [optional] 
**SanUris** | Pointer to **[]string** | Subject Alternative Name URIs. | [optional] 
**PublicKeyParameters** | Pointer to **map[string]interface{}** | Parameters specific to the public key type. | [optional] 

## Methods

### NewCertificate

`func NewCertificate(id string, organizationId string, ) *Certificate`

NewCertificate instantiates a new Certificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateWithDefaults

`func NewCertificateWithDefaults() *Certificate`

NewCertificateWithDefaults instantiates a new Certificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Certificate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Certificate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Certificate) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *Certificate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Certificate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Certificate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Certificate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Certificate) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Certificate) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Certificate) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Certificate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Certificate) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Certificate) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Certificate) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Certificate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Certificate) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Certificate) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Certificate) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetNames

`func (o *Certificate) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *Certificate) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *Certificate) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *Certificate) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetSelfSigned

`func (o *Certificate) GetSelfSigned() bool`

GetSelfSigned returns the SelfSigned field if non-nil, zero value otherwise.

### GetSelfSignedOk

`func (o *Certificate) GetSelfSignedOk() (*bool, bool)`

GetSelfSignedOk returns a tuple with the SelfSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfSigned

`func (o *Certificate) SetSelfSigned(v bool)`

SetSelfSigned sets SelfSigned field to given value.

### HasSelfSigned

`func (o *Certificate) HasSelfSigned() bool`

HasSelfSigned returns a boolean if a field has been set.

### GetHidden

`func (o *Certificate) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *Certificate) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *Certificate) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *Certificate) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetSerial

`func (o *Certificate) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *Certificate) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *Certificate) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *Certificate) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetValidityStart

`func (o *Certificate) GetValidityStart() int64`

GetValidityStart returns the ValidityStart field if non-nil, zero value otherwise.

### GetValidityStartOk

`func (o *Certificate) GetValidityStartOk() (*int64, bool)`

GetValidityStartOk returns a tuple with the ValidityStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityStart

`func (o *Certificate) SetValidityStart(v int64)`

SetValidityStart sets ValidityStart field to given value.

### HasValidityStart

`func (o *Certificate) HasValidityStart() bool`

HasValidityStart returns a boolean if a field has been set.

### GetValidityEnd

`func (o *Certificate) GetValidityEnd() int64`

GetValidityEnd returns the ValidityEnd field if non-nil, zero value otherwise.

### GetValidityEndOk

`func (o *Certificate) GetValidityEndOk() (*int64, bool)`

GetValidityEndOk returns a tuple with the ValidityEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityEnd

`func (o *Certificate) SetValidityEnd(v int64)`

SetValidityEnd sets ValidityEnd field to given value.

### HasValidityEnd

`func (o *Certificate) HasValidityEnd() bool`

HasValidityEnd returns a boolean if a field has been set.

### GetPublicKey

`func (o *Certificate) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *Certificate) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *Certificate) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *Certificate) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetPublicKeyBits

`func (o *Certificate) GetPublicKeyBits() int32`

GetPublicKeyBits returns the PublicKeyBits field if non-nil, zero value otherwise.

### GetPublicKeyBitsOk

`func (o *Certificate) GetPublicKeyBitsOk() (*int32, bool)`

GetPublicKeyBitsOk returns a tuple with the PublicKeyBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyBits

`func (o *Certificate) SetPublicKeyBits(v int32)`

SetPublicKeyBits sets PublicKeyBits field to given value.

### HasPublicKeyBits

`func (o *Certificate) HasPublicKeyBits() bool`

HasPublicKeyBits returns a boolean if a field has been set.

### GetPublicKeyAlgorithm

`func (o *Certificate) GetPublicKeyAlgorithm() string`

GetPublicKeyAlgorithm returns the PublicKeyAlgorithm field if non-nil, zero value otherwise.

### GetPublicKeyAlgorithmOk

`func (o *Certificate) GetPublicKeyAlgorithmOk() (*string, bool)`

GetPublicKeyAlgorithmOk returns a tuple with the PublicKeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyAlgorithm

`func (o *Certificate) SetPublicKeyAlgorithm(v string)`

SetPublicKeyAlgorithm sets PublicKeyAlgorithm field to given value.

### HasPublicKeyAlgorithm

`func (o *Certificate) HasPublicKeyAlgorithm() bool`

HasPublicKeyAlgorithm returns a boolean if a field has been set.

### GetSignature

`func (o *Certificate) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *Certificate) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *Certificate) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *Certificate) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *Certificate) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *Certificate) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *Certificate) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *Certificate) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### GetFpBkhash

`func (o *Certificate) GetFpBkhash() string`

GetFpBkhash returns the FpBkhash field if non-nil, zero value otherwise.

### GetFpBkhashOk

`func (o *Certificate) GetFpBkhashOk() (*string, bool)`

GetFpBkhashOk returns a tuple with the FpBkhash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFpBkhash

`func (o *Certificate) SetFpBkhash(v string)`

SetFpBkhash sets FpBkhash field to given value.

### HasFpBkhash

`func (o *Certificate) HasFpBkhash() bool`

HasFpBkhash returns a boolean if a field has been set.

### GetFpSha1

`func (o *Certificate) GetFpSha1() string`

GetFpSha1 returns the FpSha1 field if non-nil, zero value otherwise.

### GetFpSha1Ok

`func (o *Certificate) GetFpSha1Ok() (*string, bool)`

GetFpSha1Ok returns a tuple with the FpSha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFpSha1

`func (o *Certificate) SetFpSha1(v string)`

SetFpSha1 sets FpSha1 field to given value.

### HasFpSha1

`func (o *Certificate) HasFpSha1() bool`

HasFpSha1 returns a boolean if a field has been set.

### GetFpSha256

`func (o *Certificate) GetFpSha256() string`

GetFpSha256 returns the FpSha256 field if non-nil, zero value otherwise.

### GetFpSha256Ok

`func (o *Certificate) GetFpSha256Ok() (*string, bool)`

GetFpSha256Ok returns a tuple with the FpSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFpSha256

`func (o *Certificate) SetFpSha256(v string)`

SetFpSha256 sets FpSha256 field to given value.

### HasFpSha256

`func (o *Certificate) HasFpSha256() bool`

HasFpSha256 returns a boolean if a field has been set.

### GetFpMd5

`func (o *Certificate) GetFpMd5() string`

GetFpMd5 returns the FpMd5 field if non-nil, zero value otherwise.

### GetFpMd5Ok

`func (o *Certificate) GetFpMd5Ok() (*string, bool)`

GetFpMd5Ok returns a tuple with the FpMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFpMd5

`func (o *Certificate) SetFpMd5(v string)`

SetFpMd5 sets FpMd5 field to given value.

### HasFpMd5

`func (o *Certificate) HasFpMd5() bool`

HasFpMd5 returns a boolean if a field has been set.

### GetSubject

`func (o *Certificate) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Certificate) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Certificate) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Certificate) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetCn

`func (o *Certificate) GetCn() string`

GetCn returns the Cn field if non-nil, zero value otherwise.

### GetCnOk

`func (o *Certificate) GetCnOk() (*string, bool)`

GetCnOk returns a tuple with the Cn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCn

`func (o *Certificate) SetCn(v string)`

SetCn sets Cn field to given value.

### HasCn

`func (o *Certificate) HasCn() bool`

HasCn returns a boolean if a field has been set.

### GetVersion

`func (o *Certificate) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Certificate) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Certificate) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Certificate) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetIssuer

`func (o *Certificate) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *Certificate) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *Certificate) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *Certificate) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSubjectKeyId

`func (o *Certificate) GetSubjectKeyId() string`

GetSubjectKeyId returns the SubjectKeyId field if non-nil, zero value otherwise.

### GetSubjectKeyIdOk

`func (o *Certificate) GetSubjectKeyIdOk() (*string, bool)`

GetSubjectKeyIdOk returns a tuple with the SubjectKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectKeyId

`func (o *Certificate) SetSubjectKeyId(v string)`

SetSubjectKeyId sets SubjectKeyId field to given value.

### HasSubjectKeyId

`func (o *Certificate) HasSubjectKeyId() bool`

HasSubjectKeyId returns a boolean if a field has been set.

### GetAuthorityKeyId

`func (o *Certificate) GetAuthorityKeyId() string`

GetAuthorityKeyId returns the AuthorityKeyId field if non-nil, zero value otherwise.

### GetAuthorityKeyIdOk

`func (o *Certificate) GetAuthorityKeyIdOk() (*string, bool)`

GetAuthorityKeyIdOk returns a tuple with the AuthorityKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityKeyId

`func (o *Certificate) SetAuthorityKeyId(v string)`

SetAuthorityKeyId sets AuthorityKeyId field to given value.

### HasAuthorityKeyId

`func (o *Certificate) HasAuthorityKeyId() bool`

HasAuthorityKeyId returns a boolean if a field has been set.

### GetOcspServer

`func (o *Certificate) GetOcspServer() []string`

GetOcspServer returns the OcspServer field if non-nil, zero value otherwise.

### GetOcspServerOk

`func (o *Certificate) GetOcspServerOk() (*[]string, bool)`

GetOcspServerOk returns a tuple with the OcspServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspServer

`func (o *Certificate) SetOcspServer(v []string)`

SetOcspServer sets OcspServer field to given value.

### HasOcspServer

`func (o *Certificate) HasOcspServer() bool`

HasOcspServer returns a boolean if a field has been set.

### GetCrlDistributionPoints

`func (o *Certificate) GetCrlDistributionPoints() []string`

GetCrlDistributionPoints returns the CrlDistributionPoints field if non-nil, zero value otherwise.

### GetCrlDistributionPointsOk

`func (o *Certificate) GetCrlDistributionPointsOk() (*[]string, bool)`

GetCrlDistributionPointsOk returns a tuple with the CrlDistributionPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlDistributionPoints

`func (o *Certificate) SetCrlDistributionPoints(v []string)`

SetCrlDistributionPoints sets CrlDistributionPoints field to given value.

### HasCrlDistributionPoints

`func (o *Certificate) HasCrlDistributionPoints() bool`

HasCrlDistributionPoints returns a boolean if a field has been set.

### GetIssuingCertificateUrl

`func (o *Certificate) GetIssuingCertificateUrl() []string`

GetIssuingCertificateUrl returns the IssuingCertificateUrl field if non-nil, zero value otherwise.

### GetIssuingCertificateUrlOk

`func (o *Certificate) GetIssuingCertificateUrlOk() (*[]string, bool)`

GetIssuingCertificateUrlOk returns a tuple with the IssuingCertificateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCertificateUrl

`func (o *Certificate) SetIssuingCertificateUrl(v []string)`

SetIssuingCertificateUrl sets IssuingCertificateUrl field to given value.

### HasIssuingCertificateUrl

`func (o *Certificate) HasIssuingCertificateUrl() bool`

HasIssuingCertificateUrl returns a boolean if a field has been set.

### GetIsCa

`func (o *Certificate) GetIsCa() bool`

GetIsCa returns the IsCa field if non-nil, zero value otherwise.

### GetIsCaOk

`func (o *Certificate) GetIsCaOk() (*bool, bool)`

GetIsCaOk returns a tuple with the IsCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCa

`func (o *Certificate) SetIsCa(v bool)`

SetIsCa sets IsCa field to given value.

### HasIsCa

`func (o *Certificate) HasIsCa() bool`

HasIsCa returns a boolean if a field has been set.

### GetKeyUsage

`func (o *Certificate) GetKeyUsage() []string`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *Certificate) GetKeyUsageOk() (*[]string, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *Certificate) SetKeyUsage(v []string)`

SetKeyUsage sets KeyUsage field to given value.

### HasKeyUsage

`func (o *Certificate) HasKeyUsage() bool`

HasKeyUsage returns a boolean if a field has been set.

### GetExtKeyUsage

`func (o *Certificate) GetExtKeyUsage() []string`

GetExtKeyUsage returns the ExtKeyUsage field if non-nil, zero value otherwise.

### GetExtKeyUsageOk

`func (o *Certificate) GetExtKeyUsageOk() (*[]string, bool)`

GetExtKeyUsageOk returns a tuple with the ExtKeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtKeyUsage

`func (o *Certificate) SetExtKeyUsage(v []string)`

SetExtKeyUsage sets ExtKeyUsage field to given value.

### HasExtKeyUsage

`func (o *Certificate) HasExtKeyUsage() bool`

HasExtKeyUsage returns a boolean if a field has been set.

### GetSanDnsNames

`func (o *Certificate) GetSanDnsNames() []string`

GetSanDnsNames returns the SanDnsNames field if non-nil, zero value otherwise.

### GetSanDnsNamesOk

`func (o *Certificate) GetSanDnsNamesOk() (*[]string, bool)`

GetSanDnsNamesOk returns a tuple with the SanDnsNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanDnsNames

`func (o *Certificate) SetSanDnsNames(v []string)`

SetSanDnsNames sets SanDnsNames field to given value.

### HasSanDnsNames

`func (o *Certificate) HasSanDnsNames() bool`

HasSanDnsNames returns a boolean if a field has been set.

### GetSanIpAddresses

`func (o *Certificate) GetSanIpAddresses() []*string`

GetSanIpAddresses returns the SanIpAddresses field if non-nil, zero value otherwise.

### GetSanIpAddressesOk

`func (o *Certificate) GetSanIpAddressesOk() (*[]*string, bool)`

GetSanIpAddressesOk returns a tuple with the SanIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanIpAddresses

`func (o *Certificate) SetSanIpAddresses(v []*string)`

SetSanIpAddresses sets SanIpAddresses field to given value.

### HasSanIpAddresses

`func (o *Certificate) HasSanIpAddresses() bool`

HasSanIpAddresses returns a boolean if a field has been set.

### GetSanEmailAddresses

`func (o *Certificate) GetSanEmailAddresses() []string`

GetSanEmailAddresses returns the SanEmailAddresses field if non-nil, zero value otherwise.

### GetSanEmailAddressesOk

`func (o *Certificate) GetSanEmailAddressesOk() (*[]string, bool)`

GetSanEmailAddressesOk returns a tuple with the SanEmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanEmailAddresses

`func (o *Certificate) SetSanEmailAddresses(v []string)`

SetSanEmailAddresses sets SanEmailAddresses field to given value.

### HasSanEmailAddresses

`func (o *Certificate) HasSanEmailAddresses() bool`

HasSanEmailAddresses returns a boolean if a field has been set.

### GetSanUris

`func (o *Certificate) GetSanUris() []string`

GetSanUris returns the SanUris field if non-nil, zero value otherwise.

### GetSanUrisOk

`func (o *Certificate) GetSanUrisOk() (*[]string, bool)`

GetSanUrisOk returns a tuple with the SanUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanUris

`func (o *Certificate) SetSanUris(v []string)`

SetSanUris sets SanUris field to given value.

### HasSanUris

`func (o *Certificate) HasSanUris() bool`

HasSanUris returns a boolean if a field has been set.

### GetPublicKeyParameters

`func (o *Certificate) GetPublicKeyParameters() map[string]interface{}`

GetPublicKeyParameters returns the PublicKeyParameters field if non-nil, zero value otherwise.

### GetPublicKeyParametersOk

`func (o *Certificate) GetPublicKeyParametersOk() (*map[string]interface{}, bool)`

GetPublicKeyParametersOk returns a tuple with the PublicKeyParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyParameters

`func (o *Certificate) SetPublicKeyParameters(v map[string]interface{})`

SetPublicKeyParameters sets PublicKeyParameters field to given value.

### HasPublicKeyParameters

`func (o *Certificate) HasPublicKeyParameters() bool`

HasPublicKeyParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


