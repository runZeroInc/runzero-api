# VMwareCredentialFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**Password** | **string** |  | 
**Insecure** | Pointer to **string** |  | [optional] 
**Thumbprints** | Pointer to **string** |  | [optional] 

## Methods

### NewVMwareCredentialFields

`func NewVMwareCredentialFields(username string, password string, ) *VMwareCredentialFields`

NewVMwareCredentialFields instantiates a new VMwareCredentialFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMwareCredentialFieldsWithDefaults

`func NewVMwareCredentialFieldsWithDefaults() *VMwareCredentialFields`

NewVMwareCredentialFieldsWithDefaults instantiates a new VMwareCredentialFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *VMwareCredentialFields) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VMwareCredentialFields) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VMwareCredentialFields) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *VMwareCredentialFields) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VMwareCredentialFields) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VMwareCredentialFields) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetInsecure

`func (o *VMwareCredentialFields) GetInsecure() string`

GetInsecure returns the Insecure field if non-nil, zero value otherwise.

### GetInsecureOk

`func (o *VMwareCredentialFields) GetInsecureOk() (*string, bool)`

GetInsecureOk returns a tuple with the Insecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecure

`func (o *VMwareCredentialFields) SetInsecure(v string)`

SetInsecure sets Insecure field to given value.

### HasInsecure

`func (o *VMwareCredentialFields) HasInsecure() bool`

HasInsecure returns a boolean if a field has been set.

### GetThumbprints

`func (o *VMwareCredentialFields) GetThumbprints() string`

GetThumbprints returns the Thumbprints field if non-nil, zero value otherwise.

### GetThumbprintsOk

`func (o *VMwareCredentialFields) GetThumbprintsOk() (*string, bool)`

GetThumbprintsOk returns a tuple with the Thumbprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprints

`func (o *VMwareCredentialFields) SetThumbprints(v string)`

SetThumbprints sets Thumbprints field to given value.

### HasThumbprints

`func (o *VMwareCredentialFields) HasThumbprints() bool`

HasThumbprints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


