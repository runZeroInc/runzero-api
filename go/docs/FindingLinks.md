# FindingLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cves** | Pointer to [**[]FindingLinksCvesInner**](FindingLinksCvesInner.md) |  | [optional] 
**RzBlog** | Pointer to [**[]FindingLinksCvesInner**](FindingLinksCvesInner.md) | URLs of relevant runZero blog postings. | [optional] 
**Urls** | Pointer to [**[]FindingLinksCvesInner**](FindingLinksCvesInner.md) | Additional relevant URLs. | [optional] 

## Methods

### NewFindingLinks

`func NewFindingLinks() *FindingLinks`

NewFindingLinks instantiates a new FindingLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingLinksWithDefaults

`func NewFindingLinksWithDefaults() *FindingLinks`

NewFindingLinksWithDefaults instantiates a new FindingLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCves

`func (o *FindingLinks) GetCves() []FindingLinksCvesInner`

GetCves returns the Cves field if non-nil, zero value otherwise.

### GetCvesOk

`func (o *FindingLinks) GetCvesOk() (*[]FindingLinksCvesInner, bool)`

GetCvesOk returns a tuple with the Cves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCves

`func (o *FindingLinks) SetCves(v []FindingLinksCvesInner)`

SetCves sets Cves field to given value.

### HasCves

`func (o *FindingLinks) HasCves() bool`

HasCves returns a boolean if a field has been set.

### GetRzBlog

`func (o *FindingLinks) GetRzBlog() []FindingLinksCvesInner`

GetRzBlog returns the RzBlog field if non-nil, zero value otherwise.

### GetRzBlogOk

`func (o *FindingLinks) GetRzBlogOk() (*[]FindingLinksCvesInner, bool)`

GetRzBlogOk returns a tuple with the RzBlog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRzBlog

`func (o *FindingLinks) SetRzBlog(v []FindingLinksCvesInner)`

SetRzBlog sets RzBlog field to given value.

### HasRzBlog

`func (o *FindingLinks) HasRzBlog() bool`

HasRzBlog returns a boolean if a field has been set.

### GetUrls

`func (o *FindingLinks) GetUrls() []FindingLinksCvesInner`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *FindingLinks) GetUrlsOk() (*[]FindingLinksCvesInner, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *FindingLinks) SetUrls(v []FindingLinksCvesInner)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *FindingLinks) HasUrls() bool`

HasUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


