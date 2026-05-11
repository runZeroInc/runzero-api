# \CiscoSNTCAPI

All URIs are relative to *https://console.runzero.com/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportAssetsCiscoCSV**](CiscoSNTCAPI.md#ExportAssetsCiscoCSV) | **Get** /export/org/assets.cisco.csv | Cisco serial number and model name export for Cisco Smart Net Total Care Service.



## ExportAssetsCiscoCSV

> *os.File ExportAssetsCiscoCSV(ctx).Oid(oid).Search(search).Execute()

Cisco serial number and model name export for Cisco Smart Net Total Care Service.

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
	search := "search_example" // string | an optional search string for filtering results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CiscoSNTCAPI.ExportAssetsCiscoCSV(context.Background()).Oid(oid).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CiscoSNTCAPI.ExportAssetsCiscoCSV``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportAssetsCiscoCSV`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `CiscoSNTCAPI.ExportAssetsCiscoCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportAssetsCiscoCSVRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | an optional search string for filtering results | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

