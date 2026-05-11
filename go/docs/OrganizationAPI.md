# \OrganizationAPI

All URIs are relative to *https://console.runzero.com/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkRemoveCustomIntegration**](OrganizationAPI.md#BulkRemoveCustomIntegration) | **Post** /org/custom-integrations/{custom_integration_id}/bulk/remove | Remove custom integration from a list of assets
[**ClearBulkAssetOwners**](OrganizationAPI.md#ClearBulkAssetOwners) | **Post** /org/assets/bulk/clearOwners | Clear all owners across multiple assets based on a search query
[**ClearBulkAssetTags**](OrganizationAPI.md#ClearBulkAssetTags) | **Post** /org/assets/bulk/clearTags | Clear all tags across multiple assets based on a search query
[**CreateSample**](OrganizationAPI.md#CreateSample) | **Put** /org/sites/{site_id}/sample | Create a traffic sampling task for a given site
[**CreateScan**](OrganizationAPI.md#CreateScan) | **Put** /org/sites/{site_id}/scan | Create a scan task for a given site
[**CreateSite**](OrganizationAPI.md#CreateSite) | **Put** /org/sites | Create a new site
[**GetAgent**](OrganizationAPI.md#GetAgent) | **Get** /org/agents/{agent_id} | Get details for a single agent. Legacy path for /org/explorers/{explorer_id}
[**GetAgents**](OrganizationAPI.md#GetAgents) | **Get** /org/agents | Get all agents. Legacy path for /org/explorers
[**GetAsset**](OrganizationAPI.md#GetAsset) | **Get** /org/assets/{asset_id} | Get asset details
[**GetAssets**](OrganizationAPI.md#GetAssets) | **Get** /org/assets | Get all assets
[**GetExplorer**](OrganizationAPI.md#GetExplorer) | **Get** /org/explorers/{explorer_id} | Get details for a single explorer.
[**GetExplorers**](OrganizationAPI.md#GetExplorers) | **Get** /org/explorers | Get all explorers
[**GetHostedZone**](OrganizationAPI.md#GetHostedZone) | **Get** /org/hosted-zones/{hosted_zone_id} | Get details for a single hosted zone.
[**GetHostedZones**](OrganizationAPI.md#GetHostedZones) | **Get** /org/hosted-zones | Get all hosted zones
[**GetKey**](OrganizationAPI.md#GetKey) | **Get** /org/key | Get API key details
[**GetOrgCustomIntegration**](OrganizationAPI.md#GetOrgCustomIntegration) | **Get** /org/custom-integrations/{customIntegrationId} | Get single custom integration
[**GetOrgCustomIntegrations**](OrganizationAPI.md#GetOrgCustomIntegrations) | **Get** /org/custom-integrations | Get all custom integrations
[**GetOrganization**](OrganizationAPI.md#GetOrganization) | **Get** /org | Get organization details
[**GetService**](OrganizationAPI.md#GetService) | **Get** /org/services/{service_id} | Get service details
[**GetServices**](OrganizationAPI.md#GetServices) | **Get** /org/services | Get all services
[**GetSite**](OrganizationAPI.md#GetSite) | **Get** /org/sites/{site_id} | Get site details
[**GetSites**](OrganizationAPI.md#GetSites) | **Get** /org/sites | Get all sites
[**GetTask**](OrganizationAPI.md#GetTask) | **Get** /org/tasks/{task_id} | Get task details
[**GetTaskChangeReport**](OrganizationAPI.md#GetTaskChangeReport) | **Get** /org/tasks/{task_id}/changes | Returns a temporary task change report data url
[**GetTaskLog**](OrganizationAPI.md#GetTaskLog) | **Get** /org/tasks/{task_id}/log | Returns a temporary task log data url
[**GetTaskScanData**](OrganizationAPI.md#GetTaskScanData) | **Get** /org/tasks/{task_id}/data | Returns a temporary task scan data url
[**GetTasks**](OrganizationAPI.md#GetTasks) | **Get** /org/tasks | Get all tasks (last 1000)
[**GetWirelessLAN**](OrganizationAPI.md#GetWirelessLAN) | **Get** /org/wireless/{wireless_id} | Get wireless LAN details
[**GetWirelessLANs**](OrganizationAPI.md#GetWirelessLANs) | **Get** /org/wireless | Get all wireless LANs
[**HideTask**](OrganizationAPI.md#HideTask) | **Post** /org/tasks/{task_id}/hide | Signal that a completed task should be hidden
[**ImportNessusScanData**](OrganizationAPI.md#ImportNessusScanData) | **Put** /org/sites/{site_id}/import/nessus | Import a Nessus scan data file into a site
[**ImportPacketData**](OrganizationAPI.md#ImportPacketData) | **Put** /org/sites/{site_id}/import/packet | Import a packet capture file into a site
[**ImportScanData**](OrganizationAPI.md#ImportScanData) | **Put** /org/sites/{site_id}/import | Import a scan data file into a site
[**MergeAssets**](OrganizationAPI.md#MergeAssets) | **Patch** /org/assets/merge | Merge multiple assets
[**RemoveAgent**](OrganizationAPI.md#RemoveAgent) | **Delete** /org/agents/{agent_id} | Remove and uninstall an agent. Legacy path for /org/explorers/{explorer_id}
[**RemoveAsset**](OrganizationAPI.md#RemoveAsset) | **Delete** /org/assets/{asset_id} | Remove an asset
[**RemoveAssetSource**](OrganizationAPI.md#RemoveAssetSource) | **Delete** /org/assets/{asset_id}/sources/{source_id}/remove | Remove single source from asset
[**RemoveBulkAssets**](OrganizationAPI.md#RemoveBulkAssets) | **Post** /org/assets/bulk/delete | Removes multiple assets by ID
[**RemoveCustomIntegration**](OrganizationAPI.md#RemoveCustomIntegration) | **Delete** /org/assets/{asset_id}/custom-integrations/{custom_integration_id}/remove | Remove single custom integration from asset
[**RemoveExplorer**](OrganizationAPI.md#RemoveExplorer) | **Delete** /org/explorers/{explorer_id} | Remove and uninstall an explorer
[**RemoveKey**](OrganizationAPI.md#RemoveKey) | **Delete** /org/key | Remove the current API key
[**RemoveService**](OrganizationAPI.md#RemoveService) | **Delete** /org/services/{service_id} | Remove a service
[**RemoveSite**](OrganizationAPI.md#RemoveSite) | **Delete** /org/sites/{site_id} | Remove a site and associated assets
[**RemoveWirelessLAN**](OrganizationAPI.md#RemoveWirelessLAN) | **Delete** /org/wireless/{wireless_id} | Remove a wireless LAN
[**RotateKey**](OrganizationAPI.md#RotateKey) | **Patch** /org/key/rotate | Rotate the API key secret and return the updated key
[**StopTask**](OrganizationAPI.md#StopTask) | **Post** /org/tasks/{task_id}/stop | Signal that a task should be stopped or canceled.This will also remove recurring and scheduled tasks
[**UpdateAgentSettings**](OrganizationAPI.md#UpdateAgentSettings) | **Patch** /org/agents/{agent_id} | Update the settings associated with the agent. Legacy path for /org/explorers/{explorer_id}
[**UpdateAssetComments**](OrganizationAPI.md#UpdateAssetComments) | **Patch** /org/assets/{asset_id}/comments | Update asset comments
[**UpdateAssetCriticality**](OrganizationAPI.md#UpdateAssetCriticality) | **Patch** /org/assets/{asset_id}/criticality | Update asset criticality
[**UpdateAssetOwners**](OrganizationAPI.md#UpdateAssetOwners) | **Patch** /org/assets/{asset_id}/owners | Update asset owners
[**UpdateAssetTags**](OrganizationAPI.md#UpdateAssetTags) | **Patch** /org/assets/{asset_id}/tags | Update asset tags
[**UpdateBulkAssetCriticality**](OrganizationAPI.md#UpdateBulkAssetCriticality) | **Patch** /org/assets/bulk/criticality | Update criticality across multiple assets based on a search query
[**UpdateBulkAssetOwners**](OrganizationAPI.md#UpdateBulkAssetOwners) | **Patch** /org/assets/bulk/owners | Update asset owners across multiple assets based on a search query
[**UpdateBulkAssetTags**](OrganizationAPI.md#UpdateBulkAssetTags) | **Patch** /org/assets/bulk/tags | Update tags across multiple assets based on a search query
[**UpdateExplorerSettings**](OrganizationAPI.md#UpdateExplorerSettings) | **Patch** /org/explorers/{explorer_id} | Update the settings associated with the Explorer
[**UpdateOrganization**](OrganizationAPI.md#UpdateOrganization) | **Patch** /org | Update organization details
[**UpdateSite**](OrganizationAPI.md#UpdateSite) | **Patch** /org/sites/{site_id} | Update a site definition
[**UpdateTask**](OrganizationAPI.md#UpdateTask) | **Patch** /org/tasks/{task_id} | Update task parameters
[**UpgradeAgent**](OrganizationAPI.md#UpgradeAgent) | **Post** /org/agents/{agent_id}/update | Force an agent to update and restart. Legacy path for /org/explorers/{explorer_id}/update
[**UpgradeExplorer**](OrganizationAPI.md#UpgradeExplorer) | **Post** /org/explorers/{explorer_id}/update | Force an explorer to update and restart



## BulkRemoveCustomIntegration

> BulkRemoveCustomIntegration(ctx, customIntegrationId).AssetIDs(assetIDs).Execute()

Remove custom integration from a list of assets

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
	assetIDs := *openapiclient.NewAssetIDs([]string{"453C191F-644E-4EA8-9727-0E81E5275C35"}) // AssetIDs | list of asset IDs to remove

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.BulkRemoveCustomIntegration(context.Background(), customIntegrationId).AssetIDs(assetIDs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.BulkRemoveCustomIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customIntegrationId** | **string** | UUID of the custom integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkRemoveCustomIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assetIDs** | [**AssetIDs**](AssetIDs.md) | list of asset IDs to remove | 

### Return type

 (empty response body)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClearBulkAssetOwners

> BulkAssetUpdateResult ClearBulkAssetOwners(ctx).SearchQuery(searchQuery).Oid(oid).Execute()

Clear all owners across multiple assets based on a search query

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
	searchQuery := *openapiclient.NewSearchQuery("alive:true and os:windows") // SearchQuery | search query to filter
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.ClearBulkAssetOwners(context.Background()).SearchQuery(searchQuery).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.ClearBulkAssetOwners``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClearBulkAssetOwners`: BulkAssetUpdateResult
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.ClearBulkAssetOwners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClearBulkAssetOwnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchQuery** | [**SearchQuery**](SearchQuery.md) | search query to filter | 
 **oid** | **string** | The current Organization | 

### Return type

[**BulkAssetUpdateResult**](BulkAssetUpdateResult.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClearBulkAssetTags

> map[string]interface{} ClearBulkAssetTags(ctx).SearchQuery(searchQuery).Oid(oid).Execute()

Clear all tags across multiple assets based on a search query

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
	searchQuery := *openapiclient.NewSearchQuery("alive:true and os:windows") // SearchQuery | search query to filter
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.ClearBulkAssetTags(context.Background()).SearchQuery(searchQuery).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.ClearBulkAssetTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClearBulkAssetTags`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.ClearBulkAssetTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClearBulkAssetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchQuery** | [**SearchQuery**](SearchQuery.md) | search query to filter | 
 **oid** | **string** | The current Organization | 

### Return type

**map[string]interface{}**

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSample

> Task CreateSample(ctx, siteId).SampleOptions(sampleOptions).Oid(oid).Execute()

Create a traffic sampling task for a given site

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
	siteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID or name of the site to scan
	sampleOptions := *openapiclient.NewSampleOptions("eth0,wlan0") // SampleOptions | 
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.CreateSample(context.Background(), siteId).SampleOptions(sampleOptions).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.CreateSample``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSample`: Task
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.CreateSample`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | UUID or name of the site to scan | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSampleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sampleOptions** | [**SampleOptions**](SampleOptions.md) |  | 
 **oid** | **string** | The current Organization | 

### Return type

[**Task**](Task.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateScan

> Task CreateScan(ctx, siteId).ScanOptions(scanOptions).Oid(oid).Execute()

Create a scan task for a given site

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
	siteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID or name of the site to scan
	scanOptions := *openapiclient.NewScanOptions("defaults") // ScanOptions | 
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.CreateScan(context.Background(), siteId).ScanOptions(scanOptions).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.CreateScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateScan`: Task
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.CreateScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | UUID or name of the site to scan | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scanOptions** | [**ScanOptions**](ScanOptions.md) |  | 
 **oid** | **string** | The current Organization | 

### Return type

[**Task**](Task.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSite

> Site CreateSite(ctx).SiteOptions(siteOptions).Oid(oid).Execute()

Create a new site

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
	siteOptions := *openapiclient.NewSiteOptions("New Site") // SiteOptions | site definition
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.CreateSite(context.Background()).SiteOptions(siteOptions).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.CreateSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSite`: Site
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.CreateSite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **siteOptions** | [**SiteOptions**](SiteOptions.md) | site definition | 
 **oid** | **string** | The current Organization | 

### Return type

[**Site**](Site.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgent

> Agent GetAgent(ctx, agentId).Oid(oid).Execute()

Get details for a single agent. Legacy path for /org/explorers/{explorer_id}

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
	agentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the agent
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.GetAgent(context.Background(), agentId).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgent`: Agent
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | UUID of the agent | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oid** | **string** | The current Organization | 

### Return type

[**Agent**](Agent.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgents

> []Agent GetAgents(ctx).Oid(oid).Execute()

Get all agents. Legacy path for /org/explorers

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
	resp, r, err := apiClient.OrganizationAPI.GetAgents(context.Background()).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgents`: []Agent
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetAgents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 

### Return type

[**[]Agent**](Agent.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAsset

> Asset GetAsset(ctx, assetId).Oid(oid).Execute()

Get asset details

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
	assetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the asset to retrieve
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.GetAsset(context.Background(), assetId).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAsset`: Asset
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | UUID of the asset to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oid** | **string** | The current Organization | 

### Return type

[**Asset**](Asset.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssets

> []Asset GetAssets(ctx).Oid(oid).Search(search).Fields(fields).Execute()

Get all assets

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
	fields := "fields_example" // string | A list of fields to export, comma-separated (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.GetAssets(context.Background()).Oid(oid).Search(search).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssets`: []Asset
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetAssets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 
 **fields** | **string** | A list of fields to export, comma-separated | 

### Return type

[**[]Asset**](Asset.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExplorer

> Agent GetExplorer(ctx, explorerId).Oid(oid).Execute()

Get details for a single explorer.



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
	explorerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the explorer
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.GetExplorer(context.Background(), explorerId).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetExplorer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExplorer`: Agent
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetExplorer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**explorerId** | **string** | UUID of the explorer | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExplorerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oid** | **string** | The current Organization | 

### Return type

[**Agent**](Agent.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExplorers

> []Agent GetExplorers(ctx).Oid(oid).Execute()

Get all explorers



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
	resp, r, err := apiClient.OrganizationAPI.GetExplorers(context.Background()).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetExplorers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExplorers`: []Agent
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetExplorers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExplorersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 

### Return type

[**[]Agent**](Agent.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostedZone

> HostedZone GetHostedZone(ctx, hostedZoneId).Oid(oid).Execute()

Get details for a single hosted zone.



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
	hostedZoneId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the hosted zone
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.GetHostedZone(context.Background(), hostedZoneId).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetHostedZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostedZone`: HostedZone
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetHostedZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostedZoneId** | **string** | UUID of the hosted zone | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostedZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oid** | **string** | The current Organization | 

### Return type

[**HostedZone**](HostedZone.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostedZones

> []HostedZone GetHostedZones(ctx).Oid(oid).Execute()

Get all hosted zones



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
	resp, r, err := apiClient.OrganizationAPI.GetHostedZones(context.Background()).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetHostedZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostedZones`: []HostedZone
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetHostedZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHostedZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 

### Return type

[**[]HostedZone**](HostedZone.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKey

> OrganizationAPIKey GetKey(ctx).Oid(oid).Execute()

Get API key details

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
	resp, r, err := apiClient.OrganizationAPI.GetKey(context.Background()).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKey`: OrganizationAPIKey
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 

### Return type

[**OrganizationAPIKey**](OrganizationAPIKey.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	resp, r, err := apiClient.OrganizationAPI.GetOrgCustomIntegration(context.Background(), customIntegrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetOrgCustomIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgCustomIntegration`: CustomIntegration
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetOrgCustomIntegration`: %v\n", resp)
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
	resp, r, err := apiClient.OrganizationAPI.GetOrgCustomIntegrations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetOrgCustomIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgCustomIntegrations`: CustomIntegration
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetOrgCustomIntegrations`: %v\n", resp)
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


## GetOrganization

> Organization GetOrganization(ctx).Oid(oid).Execute()

Get organization details

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
	resp, r, err := apiClient.OrganizationAPI.GetOrganization(context.Background()).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganization`: Organization
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 

### Return type

[**Organization**](Organization.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetService

> Service GetService(ctx, serviceId).Oid(oid).Execute()

Get service details

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
	serviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the service to retrieve
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.GetService(context.Background(), serviceId).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetService`: Service
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | UUID of the service to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oid** | **string** | The current Organization | 

### Return type

[**Service**](Service.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServices

> []Service GetServices(ctx).Oid(oid).Search(search).Execute()

Get all services

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
	resp, r, err := apiClient.OrganizationAPI.GetServices(context.Background()).Oid(oid).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServices`: []Service
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | an optional search string for filtering results | 

### Return type

[**[]Service**](Service.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSite

> Site GetSite(ctx, siteId).Oid(oid).Execute()

Get site details

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
	siteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID or name of the site
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.GetSite(context.Background(), siteId).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSite`: Site
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | UUID or name of the site | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oid** | **string** | The current Organization | 

### Return type

[**Site**](Site.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSites

> []Site GetSites(ctx).Oid(oid).Execute()

Get all sites

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
	resp, r, err := apiClient.OrganizationAPI.GetSites(context.Background()).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSites`: []Site
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetSites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 

### Return type

[**[]Site**](Site.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTask

> Task GetTask(ctx, taskId).Oid(oid).Execute()

Get task details

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
	taskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the task to retrieve
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.GetTask(context.Background(), taskId).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTask`: Task
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | UUID of the task to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oid** | **string** | The current Organization | 

### Return type

[**Task**](Task.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskChangeReport

> GetTaskChangeReport(ctx, taskId).Oid(oid).Execute()

Returns a temporary task change report data url

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
	taskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the task
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.GetTaskChangeReport(context.Background(), taskId).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetTaskChangeReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | UUID of the task | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskChangeReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oid** | **string** | The current Organization | 

### Return type

 (empty response body)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskLog

> GetTaskLog(ctx, taskId).Oid(oid).Execute()

Returns a temporary task log data url

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
	taskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the task
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.GetTaskLog(context.Background(), taskId).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetTaskLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | UUID of the task | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oid** | **string** | The current Organization | 

### Return type

 (empty response body)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskScanData

> GetTaskScanData(ctx, taskId).Oid(oid).Execute()

Returns a temporary task scan data url

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
	taskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the task
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.GetTaskScanData(context.Background(), taskId).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetTaskScanData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | UUID of the task | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskScanDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oid** | **string** | The current Organization | 

### Return type

 (empty response body)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTasks

> []Task GetTasks(ctx).Oid(oid).Status(status).Search(search).Execute()

Get all tasks (last 1000)

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
	status := "status_example" // string | an optional status string for filtering results (optional)
	search := "search_example" // string | an optional search string for filtering results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.GetTasks(context.Background()).Oid(oid).Status(status).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTasks`: []Task
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **status** | **string** | an optional status string for filtering results | 
 **search** | **string** | an optional search string for filtering results | 

### Return type

[**[]Task**](Task.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWirelessLAN

> Wireless GetWirelessLAN(ctx, wirelessId).Oid(oid).Execute()

Get wireless LAN details

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
	wirelessId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the wireless LAN to retrieve
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.GetWirelessLAN(context.Background(), wirelessId).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetWirelessLAN``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWirelessLAN`: Wireless
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetWirelessLAN`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wirelessId** | **string** | UUID of the wireless LAN to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWirelessLANRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oid** | **string** | The current Organization | 

### Return type

[**Wireless**](Wireless.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWirelessLANs

> []Wireless GetWirelessLANs(ctx).Oid(oid).Search(search).Execute()

Get all wireless LANs

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
	resp, r, err := apiClient.OrganizationAPI.GetWirelessLANs(context.Background()).Oid(oid).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetWirelessLANs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWirelessLANs`: []Wireless
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetWirelessLANs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWirelessLANsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | an optional search string for filtering results | 

### Return type

[**[]Wireless**](Wireless.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HideTask

> Task HideTask(ctx, taskId).Oid(oid).Execute()

Signal that a completed task should be hidden

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
	taskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the task to hide
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.HideTask(context.Background(), taskId).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.HideTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HideTask`: Task
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.HideTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | UUID of the task to hide | 

### Other Parameters

Other parameters are passed through a pointer to a apiHideTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oid** | **string** | The current Organization | 

### Return type

[**Task**](Task.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportNessusScanData

> Task ImportNessusScanData(ctx, siteId).Body(body).Oid(oid).Execute()

Import a Nessus scan data file into a site

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
	siteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID or name of the site to import Nessus scan data into
	body := os.NewFile(1234, "some_file") // *os.File | 
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.ImportNessusScanData(context.Background(), siteId).Body(body).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.ImportNessusScanData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportNessusScanData`: Task
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.ImportNessusScanData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | UUID or name of the site to import Nessus scan data into | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportNessusScanDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** |  | 
 **oid** | **string** | The current Organization | 

### Return type

[**Task**](Task.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportPacketData

> Task ImportPacketData(ctx, siteId).Body(body).Oid(oid).Execute()

Import a packet capture file into a site

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
	siteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID or name of the site to import packet capture into
	body := os.NewFile(1234, "some_file") // *os.File | 
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.ImportPacketData(context.Background(), siteId).Body(body).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.ImportPacketData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportPacketData`: Task
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.ImportPacketData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | UUID or name of the site to import packet capture into | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportPacketDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** |  | 
 **oid** | **string** | The current Organization | 

### Return type

[**Task**](Task.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportScanData

> Task ImportScanData(ctx, siteId).Body(body).Oid(oid).Execute()

Import a scan data file into a site

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
	siteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID or name of the site to import scan data into
	body := os.NewFile(1234, "some_file") // *os.File | 
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.ImportScanData(context.Background(), siteId).Body(body).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.ImportScanData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportScanData`: Task
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.ImportScanData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | UUID or name of the site to import scan data into | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportScanDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** |  | 
 **oid** | **string** | The current Organization | 

### Return type

[**Task**](Task.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MergeAssets

> map[string]interface{} MergeAssets(ctx).AssetIDs(assetIDs).Oid(oid).Execute()

Merge multiple assets

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
	assetIDs := *openapiclient.NewAssetIDs([]string{"453C191F-644E-4EA8-9727-0E81E5275C35"}) // AssetIDs | List of Asset IDs to merge
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.MergeAssets(context.Background()).AssetIDs(assetIDs).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.MergeAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MergeAssets`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.MergeAssets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMergeAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetIDs** | [**AssetIDs**](AssetIDs.md) | List of Asset IDs to merge | 
 **oid** | **string** | The current Organization | 

### Return type

**map[string]interface{}**

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAgent

> RemoveAgent(ctx, agentId).Oid(oid).Execute()

Remove and uninstall an agent. Legacy path for /org/explorers/{explorer_id}

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
	agentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the agent to remove
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.RemoveAgent(context.Background(), agentId).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.RemoveAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | UUID of the agent to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oid** | **string** | The current Organization | 

### Return type

 (empty response body)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAsset

> RemoveAsset(ctx, assetId).Oid(oid).Execute()

Remove an asset

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
	assetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the asset to remove
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.RemoveAsset(context.Background(), assetId).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.RemoveAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | UUID of the asset to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oid** | **string** | The current Organization | 

### Return type

 (empty response body)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAssetSource

> RemoveAssetSource(ctx, assetId, sourceId).Execute()

Remove single source from asset

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
	assetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the asset to update
	sourceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the source

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.RemoveAssetSource(context.Background(), assetId, sourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.RemoveAssetSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | UUID of the asset to update | 
**sourceId** | **string** | UUID of the source | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAssetSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveBulkAssets

> RemoveBulkAssets(ctx).AssetIDs(assetIDs).Oid(oid).Execute()

Removes multiple assets by ID

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
	assetIDs := *openapiclient.NewAssetIDs([]string{"453C191F-644E-4EA8-9727-0E81E5275C35"}) // AssetIDs | list of asset IDs to remove
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.RemoveBulkAssets(context.Background()).AssetIDs(assetIDs).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.RemoveBulkAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveBulkAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetIDs** | [**AssetIDs**](AssetIDs.md) | list of asset IDs to remove | 
 **oid** | **string** | The current Organization | 

### Return type

 (empty response body)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveCustomIntegration

> RemoveCustomIntegration(ctx, assetId, customIntegrationId).Execute()

Remove single custom integration from asset

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
	assetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the asset to update
	customIntegrationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the custom integration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.RemoveCustomIntegration(context.Background(), assetId, customIntegrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.RemoveCustomIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | UUID of the asset to update | 
**customIntegrationId** | **string** | UUID of the custom integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveCustomIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveExplorer

> RemoveExplorer(ctx, explorerId).Oid(oid).Execute()

Remove and uninstall an explorer



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
	explorerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the explorer to remove
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.RemoveExplorer(context.Background(), explorerId).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.RemoveExplorer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**explorerId** | **string** | UUID of the explorer to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveExplorerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oid** | **string** | The current Organization | 

### Return type

 (empty response body)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveKey

> RemoveKey(ctx).Oid(oid).Execute()

Remove the current API key

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
	r, err := apiClient.OrganizationAPI.RemoveKey(context.Background()).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.RemoveKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 

### Return type

 (empty response body)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveService

> RemoveService(ctx, serviceId).Oid(oid).Execute()

Remove a service

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
	serviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the service to remove
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.RemoveService(context.Background(), serviceId).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.RemoveService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | UUID of the service to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oid** | **string** | The current Organization | 

### Return type

 (empty response body)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveSite

> RemoveSite(ctx, siteId).Oid(oid).Execute()

Remove a site and associated assets

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
	siteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID or name of the site to remove
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.RemoveSite(context.Background(), siteId).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.RemoveSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | UUID or name of the site to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oid** | **string** | The current Organization | 

### Return type

 (empty response body)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveWirelessLAN

> RemoveWirelessLAN(ctx, wirelessId).Oid(oid).Execute()

Remove a wireless LAN

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
	wirelessId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the wireless LAN to remove
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.RemoveWirelessLAN(context.Background(), wirelessId).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.RemoveWirelessLAN``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wirelessId** | **string** | UUID of the wireless LAN to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveWirelessLANRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oid** | **string** | The current Organization | 

### Return type

 (empty response body)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateKey

> OrganizationAPIKey RotateKey(ctx).Oid(oid).Execute()

Rotate the API key secret and return the updated key

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
	resp, r, err := apiClient.OrganizationAPI.RotateKey(context.Background()).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.RotateKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateKey`: OrganizationAPIKey
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.RotateKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRotateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 

### Return type

[**OrganizationAPIKey**](OrganizationAPIKey.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopTask

> Task StopTask(ctx, taskId).Oid(oid).Execute()

Signal that a task should be stopped or canceled.This will also remove recurring and scheduled tasks

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
	taskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the task to stop
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.StopTask(context.Background(), taskId).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.StopTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopTask`: Task
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.StopTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | UUID of the task to stop | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oid** | **string** | The current Organization | 

### Return type

[**Task**](Task.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAgentSettings

> Agent UpdateAgentSettings(ctx, agentId).AgentPatchedSettings(agentPatchedSettings).Oid(oid).Execute()

Update the settings associated with the agent. Legacy path for /org/explorers/{explorer_id}

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
	agentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the agent to update
	agentPatchedSettings := *openapiclient.NewAgentPatchedSettings() // AgentPatchedSettings | The updated settings to apply to the agent
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.UpdateAgentSettings(context.Background(), agentId).AgentPatchedSettings(agentPatchedSettings).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.UpdateAgentSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAgentSettings`: Agent
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.UpdateAgentSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | UUID of the agent to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAgentSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentPatchedSettings** | [**AgentPatchedSettings**](AgentPatchedSettings.md) | The updated settings to apply to the agent | 
 **oid** | **string** | The current Organization | 

### Return type

[**Agent**](Agent.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssetComments

> Asset UpdateAssetComments(ctx, assetId).AssetComments(assetComments).Oid(oid).Execute()

Update asset comments

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
	assetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the asset to update
	assetComments := *openapiclient.NewAssetComments("Sales Laptop") // AssetComments | comments to apply to the asset
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.UpdateAssetComments(context.Background(), assetId).AssetComments(assetComments).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.UpdateAssetComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAssetComments`: Asset
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.UpdateAssetComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | UUID of the asset to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssetCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assetComments** | [**AssetComments**](AssetComments.md) | comments to apply to the asset | 
 **oid** | **string** | The current Organization | 

### Return type

[**Asset**](Asset.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssetCriticality

> Asset UpdateAssetCriticality(ctx, assetId).AssetCriticality(assetCriticality).Oid(oid).Execute()

Update asset criticality

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
	assetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the asset to update
	assetCriticality := *openapiclient.NewAssetCriticality("high") // AssetCriticality | comments to apply to the asset
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.UpdateAssetCriticality(context.Background(), assetId).AssetCriticality(assetCriticality).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.UpdateAssetCriticality``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAssetCriticality`: Asset
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.UpdateAssetCriticality`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | UUID of the asset to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssetCriticalityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assetCriticality** | [**AssetCriticality**](AssetCriticality.md) | comments to apply to the asset | 
 **oid** | **string** | The current Organization | 

### Return type

[**Asset**](Asset.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssetOwners

> Asset UpdateAssetOwners(ctx, assetId).AssetOwnerships(assetOwnerships).Oid(oid).Execute()

Update asset owners

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
	assetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the asset to update
	assetOwnerships := *openapiclient.NewAssetOwnerships() // AssetOwnerships | list of ownerships to apply to the asset
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.UpdateAssetOwners(context.Background(), assetId).AssetOwnerships(assetOwnerships).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.UpdateAssetOwners``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAssetOwners`: Asset
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.UpdateAssetOwners`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | UUID of the asset to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssetOwnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assetOwnerships** | [**AssetOwnerships**](AssetOwnerships.md) | list of ownerships to apply to the asset | 
 **oid** | **string** | The current Organization | 

### Return type

[**Asset**](Asset.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssetTags

> Asset UpdateAssetTags(ctx, assetId).AssetTags(assetTags).Oid(oid).Execute()

Update asset tags

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
	assetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the asset to update
	assetTags := *openapiclient.NewAssetTags("ThisTag=Value -OldTag") // AssetTags | tags to apply to the asset
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.UpdateAssetTags(context.Background(), assetId).AssetTags(assetTags).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.UpdateAssetTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAssetTags`: Asset
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.UpdateAssetTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | UUID of the asset to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assetTags** | [**AssetTags**](AssetTags.md) | tags to apply to the asset | 
 **oid** | **string** | The current Organization | 

### Return type

[**Asset**](Asset.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBulkAssetCriticality

> map[string]interface{} UpdateBulkAssetCriticality(ctx).AssetCriticalityWithSearch(assetCriticalityWithSearch).Oid(oid).Execute()

Update criticality across multiple assets based on a search query

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
	assetCriticalityWithSearch := *openapiclient.NewAssetCriticalityWithSearch("high", "alive:true and os:windows") // AssetCriticalityWithSearch | search query to filter and criticality to apply
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.UpdateBulkAssetCriticality(context.Background()).AssetCriticalityWithSearch(assetCriticalityWithSearch).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.UpdateBulkAssetCriticality``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBulkAssetCriticality`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.UpdateBulkAssetCriticality`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBulkAssetCriticalityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetCriticalityWithSearch** | [**AssetCriticalityWithSearch**](AssetCriticalityWithSearch.md) | search query to filter and criticality to apply | 
 **oid** | **string** | The current Organization | 

### Return type

**map[string]interface{}**

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBulkAssetOwners

> BulkAssetUpdateResult UpdateBulkAssetOwners(ctx).AssetOwnershipsWithSearch(assetOwnershipsWithSearch).Oid(oid).Execute()

Update asset owners across multiple assets based on a search query

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
	assetOwnershipsWithSearch := *openapiclient.NewAssetOwnershipsWithSearch("alive:true and os:windows", []openapiclient.AssetOwnershipsWithSearchOwnershipsInner{*openapiclient.NewAssetOwnershipsWithSearchOwnershipsInner()}) // AssetOwnershipsWithSearch | search query to filter and ownerships to apply
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.UpdateBulkAssetOwners(context.Background()).AssetOwnershipsWithSearch(assetOwnershipsWithSearch).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.UpdateBulkAssetOwners``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBulkAssetOwners`: BulkAssetUpdateResult
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.UpdateBulkAssetOwners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBulkAssetOwnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetOwnershipsWithSearch** | [**AssetOwnershipsWithSearch**](AssetOwnershipsWithSearch.md) | search query to filter and ownerships to apply | 
 **oid** | **string** | The current Organization | 

### Return type

[**BulkAssetUpdateResult**](BulkAssetUpdateResult.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBulkAssetTags

> map[string]interface{} UpdateBulkAssetTags(ctx).AssetTagsWithSearch(assetTagsWithSearch).Oid(oid).Execute()

Update tags across multiple assets based on a search query

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
	assetTagsWithSearch := *openapiclient.NewAssetTagsWithSearch("ThisTag=Value -OldTag", "alive:true and os:windows") // AssetTagsWithSearch | search query to filter and tags to apply
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.UpdateBulkAssetTags(context.Background()).AssetTagsWithSearch(assetTagsWithSearch).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.UpdateBulkAssetTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBulkAssetTags`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.UpdateBulkAssetTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBulkAssetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetTagsWithSearch** | [**AssetTagsWithSearch**](AssetTagsWithSearch.md) | search query to filter and tags to apply | 
 **oid** | **string** | The current Organization | 

### Return type

**map[string]interface{}**

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExplorerSettings

> Agent UpdateExplorerSettings(ctx, explorerId).AgentPatchedSettings(agentPatchedSettings).Oid(oid).Execute()

Update the settings associated with the Explorer



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
	explorerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the explorer to update
	agentPatchedSettings := *openapiclient.NewAgentPatchedSettings() // AgentPatchedSettings | The updated settings to apply to the Explorer
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.UpdateExplorerSettings(context.Background(), explorerId).AgentPatchedSettings(agentPatchedSettings).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.UpdateExplorerSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateExplorerSettings`: Agent
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.UpdateExplorerSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**explorerId** | **string** | UUID of the explorer to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExplorerSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentPatchedSettings** | [**AgentPatchedSettings**](AgentPatchedSettings.md) | The updated settings to apply to the Explorer | 
 **oid** | **string** | The current Organization | 

### Return type

[**Agent**](Agent.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganization

> Organization UpdateOrganization(ctx).OrgOptions(orgOptions).Oid(oid).Execute()

Update organization details

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
	orgOptions := *openapiclient.NewOrgOptions() // OrgOptions | organization options
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.UpdateOrganization(context.Background()).OrgOptions(orgOptions).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.UpdateOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganization`: Organization
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.UpdateOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgOptions** | [**OrgOptions**](OrgOptions.md) | organization options | 
 **oid** | **string** | The current Organization | 

### Return type

[**Organization**](Organization.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSite

> Site UpdateSite(ctx, siteId).SiteOptions(siteOptions).Oid(oid).Execute()

Update a site definition

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
	siteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID or name of the site to update
	siteOptions := *openapiclient.NewSiteOptions("New Site") // SiteOptions | site object
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.UpdateSite(context.Background(), siteId).SiteOptions(siteOptions).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.UpdateSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSite`: Site
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.UpdateSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | UUID or name of the site to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteOptions** | [**SiteOptions**](SiteOptions.md) | site object | 
 **oid** | **string** | The current Organization | 

### Return type

[**Site**](Site.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTask

> Task UpdateTask(ctx, taskId).TaskOptions(taskOptions).Oid(oid).Execute()

Update task parameters

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
	taskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the task to update
	taskOptions := *openapiclient.NewTaskOptions() // TaskOptions | task object
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.UpdateTask(context.Background(), taskId).TaskOptions(taskOptions).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.UpdateTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTask`: Task
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.UpdateTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | UUID of the task to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taskOptions** | [**TaskOptions**](TaskOptions.md) | task object | 
 **oid** | **string** | The current Organization | 

### Return type

[**Task**](Task.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeAgent

> UpgradeAgent(ctx, agentId).Oid(oid).Execute()

Force an agent to update and restart. Legacy path for /org/explorers/{explorer_id}/update

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
	agentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the agent to update
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.UpgradeAgent(context.Background(), agentId).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.UpgradeAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | UUID of the agent to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oid** | **string** | The current Organization | 

### Return type

 (empty response body)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeExplorer

> UpgradeExplorer(ctx, explorerId).Oid(oid).Execute()

Force an explorer to update and restart



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
	explorerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the explorer to update
	oid := "oid_example" // string | The current Organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.UpgradeExplorer(context.Background(), explorerId).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.UpgradeExplorer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**explorerId** | **string** | UUID of the explorer to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeExplorerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oid** | **string** | The current Organization | 

### Return type

 (empty response body)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

