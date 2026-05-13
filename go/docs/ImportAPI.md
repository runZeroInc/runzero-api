# \ImportAPI

All URIs are relative to *https://console.runzero.com/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImportCustomIntegrationAssets**](ImportAPI.md#ImportCustomIntegrationAssets) | **Post** /import/org/{orgID}/assets | Import assets described by a registered custom integration



## ImportCustomIntegrationAssets

> Task ImportCustomIntegrationAssets(ctx, orgID).SiteId(siteId).CustomIntegrationId(customIntegrationId).ImportTask(importTask).AssetData(assetData).Execute()

Import assets described by a registered custom integration



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
	orgID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the organization to import the assets into
	siteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the site assets are to be imported into.
	customIntegrationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the registered custom integration which produced the asset data. Uniqueness is not checked/enforced. See /account/custom-integrations api.
	importTask := *openapiclient.NewImportTask("my import task") // ImportTask | 
	assetData := os.NewFile(1234, "some_file") // *os.File | A gzip (not .tar.gz) compressed file containing ImportAsset objects. The file data may be a JSON array of ImportAsset objects, e.g. [{},{},...] or JSONL format, with a single JSON representation of an ImportAsset object on each new line, e.g. {}\\\\n{}\\\\n... 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImportAPI.ImportCustomIntegrationAssets(context.Background(), orgID).SiteId(siteId).CustomIntegrationId(customIntegrationId).ImportTask(importTask).AssetData(assetData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportAPI.ImportCustomIntegrationAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportCustomIntegrationAssets`: Task
	fmt.Fprintf(os.Stdout, "Response from `ImportAPI.ImportCustomIntegrationAssets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgID** | **string** | The ID of the organization to import the assets into | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportCustomIntegrationAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteId** | **string** | The ID of the site assets are to be imported into. | 
 **customIntegrationId** | **string** | The unique ID of the registered custom integration which produced the asset data. Uniqueness is not checked/enforced. See /account/custom-integrations api. | 
 **importTask** | [**ImportTask**](ImportTask.md) |  | 
 **assetData** | ***os.File** | A gzip (not .tar.gz) compressed file containing ImportAsset objects. The file data may be a JSON array of ImportAsset objects, e.g. [{},{},...] or JSONL format, with a single JSON representation of an ImportAsset object on each new line, e.g. {}\\\\n{}\\\\n...  | 

### Return type

[**Task**](Task.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

