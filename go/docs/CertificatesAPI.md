# \CertificatesAPI

All URIs are relative to *https://console.runzero.com/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportCertificatesJSON**](CertificatesAPI.md#ExportCertificatesJSON) | **Get** /export/org/certificates.json | Export the certificate inventory as JSON



## ExportCertificatesJSON

> []Certificate ExportCertificatesJSON(ctx).Oid(oid).Search(search).Execute()

Export the certificate inventory as JSON

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	oid := "oid_example" // string | The current Organization (optional)
	search := "search_example" // string | A search query in runZero search query syntax (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificatesAPI.ExportCertificatesJSON(context.Background()).Oid(oid).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.ExportCertificatesJSON``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportCertificatesJSON`: []Certificate
	fmt.Fprintf(os.Stdout, "Response from `CertificatesAPI.ExportCertificatesJSON`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportCertificatesJSONRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 

### Return type

[**[]Certificate**](Certificate.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

