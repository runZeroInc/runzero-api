# \DefaultAPI

All URIs are relative to *https://console.runzero.com/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportAssetMetricsJSON**](DefaultAPI.md#ExportAssetMetricsJSON) | **Get** /org/metrics | Export asset metrics



## ExportAssetMetricsJSON

> map[string]ExportAssetMetricsJSON200ResponseValue ExportAssetMetricsJSON(ctx).Oid(oid).Execute()

Export asset metrics

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ExportAssetMetricsJSON(context.Background()).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ExportAssetMetricsJSON``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportAssetMetricsJSON`: map[string]ExportAssetMetricsJSON200ResponseValue
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ExportAssetMetricsJSON`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportAssetMetricsJSONRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 

### Return type

[**map[string]ExportAssetMetricsJSON200ResponseValue**](ExportAssetMetricsJSON200ResponseValue.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

