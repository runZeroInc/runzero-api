# AWSCredentialFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | **string** |  | 
**SecretAccessKey** | **string** |  | 
**Regions** | Pointer to **string** |  | [optional] 
**UseCrossAccountOrg** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 

## Methods

### NewAWSCredentialFields

`func NewAWSCredentialFields(accessKey string, secretAccessKey string, ) *AWSCredentialFields`

NewAWSCredentialFields instantiates a new AWSCredentialFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSCredentialFieldsWithDefaults

`func NewAWSCredentialFieldsWithDefaults() *AWSCredentialFields`

NewAWSCredentialFieldsWithDefaults instantiates a new AWSCredentialFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *AWSCredentialFields) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *AWSCredentialFields) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *AWSCredentialFields) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetSecretAccessKey

`func (o *AWSCredentialFields) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *AWSCredentialFields) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *AWSCredentialFields) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.


### GetRegions

`func (o *AWSCredentialFields) GetRegions() string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *AWSCredentialFields) GetRegionsOk() (*string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *AWSCredentialFields) SetRegions(v string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *AWSCredentialFields) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetUseCrossAccountOrg

`func (o *AWSCredentialFields) GetUseCrossAccountOrg() string`

GetUseCrossAccountOrg returns the UseCrossAccountOrg field if non-nil, zero value otherwise.

### GetUseCrossAccountOrgOk

`func (o *AWSCredentialFields) GetUseCrossAccountOrgOk() (*string, bool)`

GetUseCrossAccountOrgOk returns a tuple with the UseCrossAccountOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCrossAccountOrg

`func (o *AWSCredentialFields) SetUseCrossAccountOrg(v string)`

SetUseCrossAccountOrg sets UseCrossAccountOrg field to given value.

### HasUseCrossAccountOrg

`func (o *AWSCredentialFields) HasUseCrossAccountOrg() bool`

HasUseCrossAccountOrg returns a boolean if a field has been set.

### GetRole

`func (o *AWSCredentialFields) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AWSCredentialFields) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AWSCredentialFields) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AWSCredentialFields) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


