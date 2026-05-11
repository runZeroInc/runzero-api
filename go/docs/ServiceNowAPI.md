# \ServiceNowAPI

All URIs are relative to *https://console.runzero.com/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SnowExportAssetsCSV**](ServiceNowAPI.md#SnowExportAssetsCSV) | **Get** /export/org/assets.servicenow.csv | Export an asset inventory as CSV for ServiceNow integration
[**SnowExportAssetsJSON**](ServiceNowAPI.md#SnowExportAssetsJSON) | **Get** /export/org/assets.servicenow.json | Exports the asset inventory as JSON
[**SnowExportServicesCSV**](ServiceNowAPI.md#SnowExportServicesCSV) | **Get** /export/org/services.servicenow.csv | Export a service inventory as CSV for ServiceNow integration
[**SnowServiceGraphExportAssetsJSON**](ServiceNowAPI.md#SnowServiceGraphExportAssetsJSON) | **Get** /export/org/assets.servicegraph.json | Exports the asset inventory as JSON



## SnowExportAssetsCSV

> *os.File SnowExportAssetsCSV(ctx).Oid(oid).Execute()

Export an asset inventory as CSV for ServiceNow integration

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
	resp, r, err := apiClient.ServiceNowAPI.SnowExportAssetsCSV(context.Background()).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceNowAPI.SnowExportAssetsCSV``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SnowExportAssetsCSV`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ServiceNowAPI.SnowExportAssetsCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSnowExportAssetsCSVRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 

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


## SnowExportAssetsJSON

> []AssetServiceNow SnowExportAssetsJSON(ctx).Oid(oid).Execute()

Exports the asset inventory as JSON

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
	resp, r, err := apiClient.ServiceNowAPI.SnowExportAssetsJSON(context.Background()).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceNowAPI.SnowExportAssetsJSON``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SnowExportAssetsJSON`: []AssetServiceNow
	fmt.Fprintf(os.Stdout, "Response from `ServiceNowAPI.SnowExportAssetsJSON`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSnowExportAssetsJSONRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 

### Return type

[**[]AssetServiceNow**](AssetServiceNow.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnowExportServicesCSV

> *os.File SnowExportServicesCSV(ctx).Oid(oid).Execute()

Export a service inventory as CSV for ServiceNow integration

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
	resp, r, err := apiClient.ServiceNowAPI.SnowExportServicesCSV(context.Background()).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceNowAPI.SnowExportServicesCSV``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SnowExportServicesCSV`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ServiceNowAPI.SnowExportServicesCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSnowExportServicesCSVRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 

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


## SnowServiceGraphExportAssetsJSON

> []AssetServiceNow SnowServiceGraphExportAssetsJSON(ctx).Oid(oid).Search(search).Execute()

Exports the asset inventory as JSON

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
	resp, r, err := apiClient.ServiceNowAPI.SnowServiceGraphExportAssetsJSON(context.Background()).Oid(oid).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceNowAPI.SnowServiceGraphExportAssetsJSON``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SnowServiceGraphExportAssetsJSON`: []AssetServiceNow
	fmt.Fprintf(os.Stdout, "Response from `ServiceNowAPI.SnowServiceGraphExportAssetsJSON`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSnowServiceGraphExportAssetsJSONRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | an optional search string for filtering results | 

### Return type

[**[]AssetServiceNow**](AssetServiceNow.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

