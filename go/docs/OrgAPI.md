# \OrgAPI

All URIs are relative to *https://console.runzero.com/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgCustomIntegration**](OrgAPI.md#GetOrgCustomIntegration) | **Get** /org/custom-integrations/{customIntegrationId} | Get single custom integration
[**GetOrgCustomIntegrations**](OrgAPI.md#GetOrgCustomIntegrations) | **Get** /org/custom-integrations | Get all custom integrations



## GetOrgCustomIntegration

> CustomIntegration GetOrgCustomIntegration(ctx, customIntegrationId).Execute()

Get single custom integration

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
	customIntegrationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the custom integration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgAPI.GetOrgCustomIntegration(context.Background(), customIntegrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgAPI.GetOrgCustomIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgCustomIntegration`: CustomIntegration
	fmt.Fprintf(os.Stdout, "Response from `OrgAPI.GetOrgCustomIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customIntegrationId** | **string** | UUID of the custom integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgCustomIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomIntegration**](CustomIntegration.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgCustomIntegrations

> CustomIntegration GetOrgCustomIntegrations(ctx).Execute()

Get all custom integrations

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgAPI.GetOrgCustomIntegrations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgAPI.GetOrgCustomIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgCustomIntegrations`: CustomIntegration
	fmt.Fprintf(os.Stdout, "Response from `OrgAPI.GetOrgCustomIntegrations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgCustomIntegrationsRequest struct via the builder pattern


### Return type

[**CustomIntegration**](CustomIntegration.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

