# \SplunkAPI

All URIs are relative to *https://console.runzero.com/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SplunkAssetSyncCreatedJSON**](SplunkAPI.md#SplunkAssetSyncCreatedJSON) | **Get** /export/org/assets/sync/created/assets.json | Exports the asset inventory in a sync-friendly manner using created_at as a checkpoint. Requires the Splunk entitlement.
[**SplunkAssetSyncUpdatedJSON**](SplunkAPI.md#SplunkAssetSyncUpdatedJSON) | **Get** /export/org/assets/sync/updated/assets.json | Exports the asset inventory in a sync-friendly manner using updated_at as a checkpoint. Requires the Splunk entitlement.



## SplunkAssetSyncCreatedJSON

> AssetsWithCheckpoint SplunkAssetSyncCreatedJSON(ctx).Oid(oid).Search(search).Fields(fields).Since(since).Execute()

Exports the asset inventory in a sync-friendly manner using created_at as a checkpoint. Requires the Splunk entitlement.

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
	fields := "fields_example" // string | an optional list of fields to export, comma-separated (optional)
	since := int64(1576300370) // int64 | an optional unix timestamp to use as a checkpoint (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SplunkAPI.SplunkAssetSyncCreatedJSON(context.Background()).Oid(oid).Search(search).Fields(fields).Since(since).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SplunkAPI.SplunkAssetSyncCreatedJSON``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SplunkAssetSyncCreatedJSON`: AssetsWithCheckpoint
	fmt.Fprintf(os.Stdout, "Response from `SplunkAPI.SplunkAssetSyncCreatedJSON`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSplunkAssetSyncCreatedJSONRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | an optional search string for filtering results | 
 **fields** | **string** | an optional list of fields to export, comma-separated | 
 **since** | **int64** | an optional unix timestamp to use as a checkpoint | 

### Return type

[**AssetsWithCheckpoint**](AssetsWithCheckpoint.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SplunkAssetSyncUpdatedJSON

> AssetsWithCheckpoint SplunkAssetSyncUpdatedJSON(ctx).Oid(oid).Search(search).Fields(fields).Since(since).Execute()

Exports the asset inventory in a sync-friendly manner using updated_at as a checkpoint. Requires the Splunk entitlement.

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
	fields := "fields_example" // string | an optional list of fields to export, comma-separated (optional)
	since := int64(1576300370) // int64 | an optional unix timestamp to use as a checkpoint (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SplunkAPI.SplunkAssetSyncUpdatedJSON(context.Background()).Oid(oid).Search(search).Fields(fields).Since(since).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SplunkAPI.SplunkAssetSyncUpdatedJSON``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SplunkAssetSyncUpdatedJSON`: AssetsWithCheckpoint
	fmt.Fprintf(os.Stdout, "Response from `SplunkAPI.SplunkAssetSyncUpdatedJSON`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSplunkAssetSyncUpdatedJSONRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | an optional search string for filtering results | 
 **fields** | **string** | an optional list of fields to export, comma-separated | 
 **since** | **int64** | an optional unix timestamp to use as a checkpoint | 

### Return type

[**AssetsWithCheckpoint**](AssetsWithCheckpoint.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

