# \ExportAPI

All URIs are relative to *https://console.runzero.com/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportAssetMetricsJSON**](ExportAPI.md#ExportAssetMetricsJSON) | **Get** /org/metrics | Export asset metrics
[**ExportAssetTopHWCSV**](ExportAPI.md#ExportAssetTopHWCSV) | **Get** /org/assets/hw.csv | Top asset hardware products as CSV
[**ExportAssetTopOSCSV**](ExportAPI.md#ExportAssetTopOSCSV) | **Get** /org/assets/os.csv | Top asset operating systems as CSV
[**ExportAssetTopTagsCSV**](ExportAPI.md#ExportAssetTopTagsCSV) | **Get** /org/assets/tags.csv | Top asset tags as CSV
[**ExportAssetTopTypesCSV**](ExportAPI.md#ExportAssetTopTypesCSV) | **Get** /org/assets/type.csv | Top asset types as CSV
[**ExportAssetsCSV**](ExportAPI.md#ExportAssetsCSV) | **Get** /export/org/assets.csv | Asset inventory as CSV
[**ExportAssetsJSON**](ExportAPI.md#ExportAssetsJSON) | **Get** /export/org/assets.json | Exports the asset inventory
[**ExportAssetsJSONL**](ExportAPI.md#ExportAssetsJSONL) | **Get** /export/org/assets.jsonl | Asset inventory as JSON line-delimited
[**ExportAssetsNmapXML**](ExportAPI.md#ExportAssetsNmapXML) | **Get** /export/org/assets.nmap.xml | Asset inventory as Nmap-style XML
[**ExportCertificatesCSV**](ExportAPI.md#ExportCertificatesCSV) | **Get** /export/org/certificates.csv | Export the certificate inventory as CSV
[**ExportCertificatesJSON**](ExportAPI.md#ExportCertificatesJSON) | **Get** /export/org/certificates.json | Export the certificate inventory as JSON
[**ExportCertificatesJSONL**](ExportAPI.md#ExportCertificatesJSONL) | **Get** /export/org/certificates.jsonl | Export the certificate inventory as JSONL line-delimited
[**ExportDirectoryGroupsCSV**](ExportAPI.md#ExportDirectoryGroupsCSV) | **Get** /export/org/groups.csv | Group inventory as CSV
[**ExportDirectoryGroupsJSON**](ExportAPI.md#ExportDirectoryGroupsJSON) | **Get** /export/org/groups.json | Exports the group inventory
[**ExportDirectoryGroupsJSONL**](ExportAPI.md#ExportDirectoryGroupsJSONL) | **Get** /export/org/groups.jsonl | Group inventory as JSON line-delimited
[**ExportDirectoryUsersCSV**](ExportAPI.md#ExportDirectoryUsersCSV) | **Get** /export/org/users.csv | User inventory as CSV
[**ExportDirectoryUsersJSON**](ExportAPI.md#ExportDirectoryUsersJSON) | **Get** /export/org/users.json | Exports the user inventory
[**ExportDirectoryUsersJSONL**](ExportAPI.md#ExportDirectoryUsersJSONL) | **Get** /export/org/users.jsonl | User inventory as JSON line-delimited
[**ExportFindingsCSV**](ExportAPI.md#ExportFindingsCSV) | **Get** /export/org/findings.csv | Export findings as CSV
[**ExportFindingsJSON**](ExportAPI.md#ExportFindingsJSON) | **Get** /export/org/findings.json | Export findings as JSON
[**ExportFindingsJSONL**](ExportAPI.md#ExportFindingsJSONL) | **Get** /export/org/findings.jsonl | Export findings as JSON line-delimited
[**ExportSNMPARPCacheCSV**](ExportAPI.md#ExportSNMPARPCacheCSV) | **Get** /export/org/snmp.arpcache.csv | SNMP ARP cache data as CSV
[**ExportServicesCSV**](ExportAPI.md#ExportServicesCSV) | **Get** /export/org/services.csv | Service inventory as CSV
[**ExportServicesJSON**](ExportAPI.md#ExportServicesJSON) | **Get** /export/org/services.json | Service inventory as JSON
[**ExportServicesJSONL**](ExportAPI.md#ExportServicesJSONL) | **Get** /export/org/services.jsonl | Service inventory as JSON line-delimited
[**ExportServicesTopProductsCSV**](ExportAPI.md#ExportServicesTopProductsCSV) | **Get** /org/services/products.csv | Top service products as CSV
[**ExportServicesTopProtocolsCSV**](ExportAPI.md#ExportServicesTopProtocolsCSV) | **Get** /org/services/protocols.csv | Top service protocols as CSV
[**ExportServicesTopTCPCSV**](ExportAPI.md#ExportServicesTopTCPCSV) | **Get** /org/services/tcp.csv | Top TCP services as CSV
[**ExportServicesTopUDPCSV**](ExportAPI.md#ExportServicesTopUDPCSV) | **Get** /org/services/udp.csv | Top UDP services as CSV
[**ExportSitesCSV**](ExportAPI.md#ExportSitesCSV) | **Get** /export/org/sites.csv | Site list as CSV
[**ExportSitesJSON**](ExportAPI.md#ExportSitesJSON) | **Get** /export/org/sites.json | Export all sites
[**ExportSitesJSONL**](ExportAPI.md#ExportSitesJSONL) | **Get** /export/org/sites.jsonl | Site list as JSON line-delimited
[**ExportSoftwareCSV**](ExportAPI.md#ExportSoftwareCSV) | **Get** /export/org/software.csv | Software inventory as CSV
[**ExportSoftwareJSON**](ExportAPI.md#ExportSoftwareJSON) | **Get** /export/org/software.json | Exports the software inventory
[**ExportSoftwareJSONL**](ExportAPI.md#ExportSoftwareJSONL) | **Get** /export/org/software.jsonl | Software inventory as JSON line-delimited
[**ExportSubnetUtilizationStatsCSV**](ExportAPI.md#ExportSubnetUtilizationStatsCSV) | **Get** /export/org/subnet.stats.csv | Subnet utilization statistics as as CSV
[**ExportTasksJSON**](ExportAPI.md#ExportTasksJSON) | **Get** /export/org/tasks.json | Exports organization tasks
[**ExportTasksJSONL**](ExportAPI.md#ExportTasksJSONL) | **Get** /export/org/tasks.jsonl | Organization tasks as JSON line-delimited
[**ExportVulnerabilitiesCSV**](ExportAPI.md#ExportVulnerabilitiesCSV) | **Get** /export/org/vulnerabilities.csv | Export the vulnerability inventory as CSV
[**ExportVulnerabilitiesJSON**](ExportAPI.md#ExportVulnerabilitiesJSON) | **Get** /export/org/vulnerabilities.json | Export the vulnerability inventory as JSON
[**ExportVulnerabilitiesJSONL**](ExportAPI.md#ExportVulnerabilitiesJSONL) | **Get** /export/org/vulnerabilities.jsonl | Export the vulnerability inventory as JSON line-delimited
[**ExportWirelessCSV**](ExportAPI.md#ExportWirelessCSV) | **Get** /export/org/wireless.csv | Wireless inventory as CSV
[**ExportWirelessJSON**](ExportAPI.md#ExportWirelessJSON) | **Get** /export/org/wireless.json | Wireless inventory as JSON
[**ExportWirelessJSONL**](ExportAPI.md#ExportWirelessJSONL) | **Get** /export/org/wireless.jsonl | Wireless inventory as JSON line-delimited



## ExportAssetMetricsJSON

> map[string]AssetMetric ExportAssetMetricsJSON(ctx).Oid(oid).Execute()

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
	resp, r, err := apiClient.ExportAPI.ExportAssetMetricsJSON(context.Background()).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportAssetMetricsJSON``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportAssetMetricsJSON`: map[string]AssetMetric
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportAssetMetricsJSON`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportAssetMetricsJSONRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 

### Return type

[**map[string]AssetMetric**](AssetMetric.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportAssetTopHWCSV

> *os.File ExportAssetTopHWCSV(ctx).Oid(oid).Execute()

Top asset hardware products as CSV

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
	resp, r, err := apiClient.ExportAPI.ExportAssetTopHWCSV(context.Background()).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportAssetTopHWCSV``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportAssetTopHWCSV`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportAssetTopHWCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportAssetTopHWCSVRequest struct via the builder pattern


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


## ExportAssetTopOSCSV

> *os.File ExportAssetTopOSCSV(ctx).Oid(oid).Execute()

Top asset operating systems as CSV

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
	resp, r, err := apiClient.ExportAPI.ExportAssetTopOSCSV(context.Background()).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportAssetTopOSCSV``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportAssetTopOSCSV`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportAssetTopOSCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportAssetTopOSCSVRequest struct via the builder pattern


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


## ExportAssetTopTagsCSV

> *os.File ExportAssetTopTagsCSV(ctx).Oid(oid).Execute()

Top asset tags as CSV

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
	resp, r, err := apiClient.ExportAPI.ExportAssetTopTagsCSV(context.Background()).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportAssetTopTagsCSV``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportAssetTopTagsCSV`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportAssetTopTagsCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportAssetTopTagsCSVRequest struct via the builder pattern


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


## ExportAssetTopTypesCSV

> *os.File ExportAssetTopTypesCSV(ctx).Oid(oid).Execute()

Top asset types as CSV

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
	resp, r, err := apiClient.ExportAPI.ExportAssetTopTypesCSV(context.Background()).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportAssetTopTypesCSV``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportAssetTopTypesCSV`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportAssetTopTypesCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportAssetTopTypesCSVRequest struct via the builder pattern


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


## ExportAssetsCSV

> *os.File ExportAssetsCSV(ctx).Oid(oid).Search(search).Execute()

Asset inventory as CSV

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
	resp, r, err := apiClient.ExportAPI.ExportAssetsCSV(context.Background()).Oid(oid).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportAssetsCSV``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportAssetsCSV`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportAssetsCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportAssetsCSVRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 

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


## ExportAssetsJSON

> AssetExportResponse ExportAssetsJSON(ctx).Oid(oid).Search(search).Fields(fields).PageSize(pageSize).StartKey(startKey).Execute()

Exports the asset inventory

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
	pageSize := int32(56) // int32 | The number of results to return per request. (optional)
	startKey := "startKey_example" // string | The value to use for requesting the next page when requesting paginated results.  This should be the value of the `next_key` attribute returned in the previous response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportAPI.ExportAssetsJSON(context.Background()).Oid(oid).Search(search).Fields(fields).PageSize(pageSize).StartKey(startKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportAssetsJSON``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportAssetsJSON`: AssetExportResponse
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportAssetsJSON`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportAssetsJSONRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 
 **fields** | **string** | A list of fields to export, comma-separated | 
 **pageSize** | **int32** | The number of results to return per request. | 
 **startKey** | **string** | The value to use for requesting the next page when requesting paginated results.  This should be the value of the &#x60;next_key&#x60; attribute returned in the previous response. | 

### Return type

[**AssetExportResponse**](AssetExportResponse.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportAssetsJSONL

> *os.File ExportAssetsJSONL(ctx).Oid(oid).Search(search).Fields(fields).Execute()

Asset inventory as JSON line-delimited

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
	resp, r, err := apiClient.ExportAPI.ExportAssetsJSONL(context.Background()).Oid(oid).Search(search).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportAssetsJSONL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportAssetsJSONL`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportAssetsJSONL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportAssetsJSONLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 
 **fields** | **string** | A list of fields to export, comma-separated | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportAssetsNmapXML

> *os.File ExportAssetsNmapXML(ctx).Oid(oid).Search(search).Execute()

Asset inventory as Nmap-style XML

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
	resp, r, err := apiClient.ExportAPI.ExportAssetsNmapXML(context.Background()).Oid(oid).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportAssetsNmapXML``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportAssetsNmapXML`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportAssetsNmapXML`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportAssetsNmapXMLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportCertificatesCSV

> *os.File ExportCertificatesCSV(ctx).Oid(oid).Search(search).Execute()

Export the certificate inventory as CSV

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
	resp, r, err := apiClient.ExportAPI.ExportCertificatesCSV(context.Background()).Oid(oid).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportCertificatesCSV``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportCertificatesCSV`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportCertificatesCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportCertificatesCSVRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 

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
	resp, r, err := apiClient.ExportAPI.ExportCertificatesJSON(context.Background()).Oid(oid).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportCertificatesJSON``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportCertificatesJSON`: []Certificate
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportCertificatesJSON`: %v\n", resp)
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


## ExportCertificatesJSONL

> *os.File ExportCertificatesJSONL(ctx).Oid(oid).Search(search).Execute()

Export the certificate inventory as JSONL line-delimited

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
	resp, r, err := apiClient.ExportAPI.ExportCertificatesJSONL(context.Background()).Oid(oid).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportCertificatesJSONL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportCertificatesJSONL`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportCertificatesJSONL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportCertificatesJSONLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportDirectoryGroupsCSV

> *os.File ExportDirectoryGroupsCSV(ctx).Oid(oid).Search(search).Execute()

Group inventory as CSV

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
	resp, r, err := apiClient.ExportAPI.ExportDirectoryGroupsCSV(context.Background()).Oid(oid).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportDirectoryGroupsCSV``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportDirectoryGroupsCSV`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportDirectoryGroupsCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportDirectoryGroupsCSVRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 

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


## ExportDirectoryGroupsJSON

> []DirectoryGroup ExportDirectoryGroupsJSON(ctx).Oid(oid).Search(search).Fields(fields).Execute()

Exports the group inventory

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
	resp, r, err := apiClient.ExportAPI.ExportDirectoryGroupsJSON(context.Background()).Oid(oid).Search(search).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportDirectoryGroupsJSON``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportDirectoryGroupsJSON`: []DirectoryGroup
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportDirectoryGroupsJSON`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportDirectoryGroupsJSONRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 
 **fields** | **string** | A list of fields to export, comma-separated | 

### Return type

[**[]DirectoryGroup**](DirectoryGroup.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportDirectoryGroupsJSONL

> *os.File ExportDirectoryGroupsJSONL(ctx).Oid(oid).Search(search).Fields(fields).Execute()

Group inventory as JSON line-delimited

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
	resp, r, err := apiClient.ExportAPI.ExportDirectoryGroupsJSONL(context.Background()).Oid(oid).Search(search).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportDirectoryGroupsJSONL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportDirectoryGroupsJSONL`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportDirectoryGroupsJSONL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportDirectoryGroupsJSONLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 
 **fields** | **string** | A list of fields to export, comma-separated | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportDirectoryUsersCSV

> *os.File ExportDirectoryUsersCSV(ctx).Oid(oid).Search(search).Execute()

User inventory as CSV

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
	resp, r, err := apiClient.ExportAPI.ExportDirectoryUsersCSV(context.Background()).Oid(oid).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportDirectoryUsersCSV``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportDirectoryUsersCSV`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportDirectoryUsersCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportDirectoryUsersCSVRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 

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


## ExportDirectoryUsersJSON

> []DirectoryUser ExportDirectoryUsersJSON(ctx).Oid(oid).Search(search).Fields(fields).Execute()

Exports the user inventory

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
	resp, r, err := apiClient.ExportAPI.ExportDirectoryUsersJSON(context.Background()).Oid(oid).Search(search).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportDirectoryUsersJSON``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportDirectoryUsersJSON`: []DirectoryUser
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportDirectoryUsersJSON`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportDirectoryUsersJSONRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 
 **fields** | **string** | A list of fields to export, comma-separated | 

### Return type

[**[]DirectoryUser**](DirectoryUser.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportDirectoryUsersJSONL

> *os.File ExportDirectoryUsersJSONL(ctx).Oid(oid).Search(search).Fields(fields).Execute()

User inventory as JSON line-delimited

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
	resp, r, err := apiClient.ExportAPI.ExportDirectoryUsersJSONL(context.Background()).Oid(oid).Search(search).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportDirectoryUsersJSONL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportDirectoryUsersJSONL`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportDirectoryUsersJSONL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportDirectoryUsersJSONLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 
 **fields** | **string** | A list of fields to export, comma-separated | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportFindingsCSV

> *os.File ExportFindingsCSV(ctx).Oid(oid).Search(search).Execute()

Export findings as CSV

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
	resp, r, err := apiClient.ExportAPI.ExportFindingsCSV(context.Background()).Oid(oid).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportFindingsCSV``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportFindingsCSV`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportFindingsCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportFindingsCSVRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 

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


## ExportFindingsJSON

> []Finding ExportFindingsJSON(ctx).Oid(oid).Search(search).Execute()

Export findings as JSON

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
	resp, r, err := apiClient.ExportAPI.ExportFindingsJSON(context.Background()).Oid(oid).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportFindingsJSON``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportFindingsJSON`: []Finding
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportFindingsJSON`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportFindingsJSONRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 

### Return type

[**[]Finding**](Finding.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportFindingsJSONL

> []Finding ExportFindingsJSONL(ctx).Oid(oid).Search(search).Execute()

Export findings as JSON line-delimited

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
	resp, r, err := apiClient.ExportAPI.ExportFindingsJSONL(context.Background()).Oid(oid).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportFindingsJSONL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportFindingsJSONL`: []Finding
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportFindingsJSONL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportFindingsJSONLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 

### Return type

[**[]Finding**](Finding.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportSNMPARPCacheCSV

> *os.File ExportSNMPARPCacheCSV(ctx).Oid(oid).Execute()

SNMP ARP cache data as CSV

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
	resp, r, err := apiClient.ExportAPI.ExportSNMPARPCacheCSV(context.Background()).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportSNMPARPCacheCSV``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportSNMPARPCacheCSV`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportSNMPARPCacheCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportSNMPARPCacheCSVRequest struct via the builder pattern


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


## ExportServicesCSV

> *os.File ExportServicesCSV(ctx).Oid(oid).Search(search).Execute()

Service inventory as CSV

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
	resp, r, err := apiClient.ExportAPI.ExportServicesCSV(context.Background()).Oid(oid).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportServicesCSV``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportServicesCSV`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportServicesCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportServicesCSVRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 

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


## ExportServicesJSON

> ServiceExportResponse ExportServicesJSON(ctx).Oid(oid).Search(search).Fields(fields).PageSize(pageSize).StartKey(startKey).Execute()

Service inventory as JSON

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
	pageSize := int32(56) // int32 | The number of results to return per request. (optional)
	startKey := "startKey_example" // string | The value to use for requesting the next page when requesting paginated results.  This should be the value of the `next_key` attribute returned in the previous response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportAPI.ExportServicesJSON(context.Background()).Oid(oid).Search(search).Fields(fields).PageSize(pageSize).StartKey(startKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportServicesJSON``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportServicesJSON`: ServiceExportResponse
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportServicesJSON`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportServicesJSONRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 
 **fields** | **string** | A list of fields to export, comma-separated | 
 **pageSize** | **int32** | The number of results to return per request. | 
 **startKey** | **string** | The value to use for requesting the next page when requesting paginated results.  This should be the value of the &#x60;next_key&#x60; attribute returned in the previous response. | 

### Return type

[**ServiceExportResponse**](ServiceExportResponse.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportServicesJSONL

> *os.File ExportServicesJSONL(ctx).Oid(oid).Search(search).Fields(fields).Execute()

Service inventory as JSON line-delimited

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
	resp, r, err := apiClient.ExportAPI.ExportServicesJSONL(context.Background()).Oid(oid).Search(search).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportServicesJSONL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportServicesJSONL`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportServicesJSONL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportServicesJSONLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 
 **fields** | **string** | A list of fields to export, comma-separated | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportServicesTopProductsCSV

> *os.File ExportServicesTopProductsCSV(ctx).Oid(oid).Execute()

Top service products as CSV

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
	resp, r, err := apiClient.ExportAPI.ExportServicesTopProductsCSV(context.Background()).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportServicesTopProductsCSV``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportServicesTopProductsCSV`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportServicesTopProductsCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportServicesTopProductsCSVRequest struct via the builder pattern


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


## ExportServicesTopProtocolsCSV

> *os.File ExportServicesTopProtocolsCSV(ctx).Oid(oid).Execute()

Top service protocols as CSV

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
	resp, r, err := apiClient.ExportAPI.ExportServicesTopProtocolsCSV(context.Background()).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportServicesTopProtocolsCSV``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportServicesTopProtocolsCSV`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportServicesTopProtocolsCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportServicesTopProtocolsCSVRequest struct via the builder pattern


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


## ExportServicesTopTCPCSV

> *os.File ExportServicesTopTCPCSV(ctx).Oid(oid).Execute()

Top TCP services as CSV

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
	resp, r, err := apiClient.ExportAPI.ExportServicesTopTCPCSV(context.Background()).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportServicesTopTCPCSV``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportServicesTopTCPCSV`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportServicesTopTCPCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportServicesTopTCPCSVRequest struct via the builder pattern


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


## ExportServicesTopUDPCSV

> *os.File ExportServicesTopUDPCSV(ctx).Oid(oid).Execute()

Top UDP services as CSV

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
	resp, r, err := apiClient.ExportAPI.ExportServicesTopUDPCSV(context.Background()).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportServicesTopUDPCSV``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportServicesTopUDPCSV`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportServicesTopUDPCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportServicesTopUDPCSVRequest struct via the builder pattern


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


## ExportSitesCSV

> *os.File ExportSitesCSV(ctx).Oid(oid).Execute()

Site list as CSV

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
	resp, r, err := apiClient.ExportAPI.ExportSitesCSV(context.Background()).Oid(oid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportSitesCSV``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportSitesCSV`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportSitesCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportSitesCSVRequest struct via the builder pattern


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


## ExportSitesJSON

> []Site ExportSitesJSON(ctx).Oid(oid).Search(search).Fields(fields).Execute()

Export all sites

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
	resp, r, err := apiClient.ExportAPI.ExportSitesJSON(context.Background()).Oid(oid).Search(search).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportSitesJSON``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportSitesJSON`: []Site
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportSitesJSON`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportSitesJSONRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 
 **fields** | **string** | A list of fields to export, comma-separated | 

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


## ExportSitesJSONL

> *os.File ExportSitesJSONL(ctx).Oid(oid).Search(search).Fields(fields).Execute()

Site list as JSON line-delimited

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
	resp, r, err := apiClient.ExportAPI.ExportSitesJSONL(context.Background()).Oid(oid).Search(search).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportSitesJSONL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportSitesJSONL`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportSitesJSONL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportSitesJSONLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 
 **fields** | **string** | A list of fields to export, comma-separated | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportSoftwareCSV

> *os.File ExportSoftwareCSV(ctx).Oid(oid).Search(search).Execute()

Software inventory as CSV

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
	resp, r, err := apiClient.ExportAPI.ExportSoftwareCSV(context.Background()).Oid(oid).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportSoftwareCSV``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportSoftwareCSV`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportSoftwareCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportSoftwareCSVRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 

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


## ExportSoftwareJSON

> SoftwareExportResponse ExportSoftwareJSON(ctx).Oid(oid).Search(search).Fields(fields).PageSize(pageSize).StartKey(startKey).Execute()

Exports the software inventory

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
	pageSize := int32(56) // int32 | The number of results to return per request. (optional)
	startKey := "startKey_example" // string | The value to use for requesting the next page when requesting paginated results.  This should be the value of the `next_key` attribute returned in the previous response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportAPI.ExportSoftwareJSON(context.Background()).Oid(oid).Search(search).Fields(fields).PageSize(pageSize).StartKey(startKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportSoftwareJSON``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportSoftwareJSON`: SoftwareExportResponse
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportSoftwareJSON`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportSoftwareJSONRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 
 **fields** | **string** | A list of fields to export, comma-separated | 
 **pageSize** | **int32** | The number of results to return per request. | 
 **startKey** | **string** | The value to use for requesting the next page when requesting paginated results.  This should be the value of the &#x60;next_key&#x60; attribute returned in the previous response. | 

### Return type

[**SoftwareExportResponse**](SoftwareExportResponse.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportSoftwareJSONL

> *os.File ExportSoftwareJSONL(ctx).Oid(oid).Search(search).Fields(fields).Execute()

Software inventory as JSON line-delimited

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
	resp, r, err := apiClient.ExportAPI.ExportSoftwareJSONL(context.Background()).Oid(oid).Search(search).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportSoftwareJSONL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportSoftwareJSONL`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportSoftwareJSONL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportSoftwareJSONLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 
 **fields** | **string** | A list of fields to export, comma-separated | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportSubnetUtilizationStatsCSV

> *os.File ExportSubnetUtilizationStatsCSV(ctx).Oid(oid).Mask(mask).Execute()

Subnet utilization statistics as as CSV

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
	mask := "mask_example" // string | an optional subnet mask size (ex:24) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportAPI.ExportSubnetUtilizationStatsCSV(context.Background()).Oid(oid).Mask(mask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportSubnetUtilizationStatsCSV``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportSubnetUtilizationStatsCSV`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportSubnetUtilizationStatsCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportSubnetUtilizationStatsCSVRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **mask** | **string** | an optional subnet mask size (ex:24) | 

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


## ExportTasksJSON

> []Task ExportTasksJSON(ctx).Oid(oid).Search(search).Fields(fields).Execute()

Exports organization tasks

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
	resp, r, err := apiClient.ExportAPI.ExportTasksJSON(context.Background()).Oid(oid).Search(search).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportTasksJSON``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportTasksJSON`: []Task
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportTasksJSON`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportTasksJSONRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 
 **fields** | **string** | A list of fields to export, comma-separated | 

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


## ExportTasksJSONL

> *os.File ExportTasksJSONL(ctx).Oid(oid).Search(search).Fields(fields).Execute()

Organization tasks as JSON line-delimited

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
	resp, r, err := apiClient.ExportAPI.ExportTasksJSONL(context.Background()).Oid(oid).Search(search).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportTasksJSONL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportTasksJSONL`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportTasksJSONL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportTasksJSONLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 
 **fields** | **string** | A list of fields to export, comma-separated | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportVulnerabilitiesCSV

> *os.File ExportVulnerabilitiesCSV(ctx).Oid(oid).Search(search).Execute()

Export the vulnerability inventory as CSV

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
	resp, r, err := apiClient.ExportAPI.ExportVulnerabilitiesCSV(context.Background()).Oid(oid).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportVulnerabilitiesCSV``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportVulnerabilitiesCSV`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportVulnerabilitiesCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportVulnerabilitiesCSVRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 

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


## ExportVulnerabilitiesJSON

> VulnerabilityExportResponse ExportVulnerabilitiesJSON(ctx).Oid(oid).Search(search).Fields(fields).PageSize(pageSize).StartKey(startKey).Execute()

Export the vulnerability inventory as JSON

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
	pageSize := int32(56) // int32 | The number of results to return per request. (optional)
	startKey := "startKey_example" // string | The value to use for requesting the next page when requesting paginated results.  This should be the value of the `next_key` attribute returned in the previous response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportAPI.ExportVulnerabilitiesJSON(context.Background()).Oid(oid).Search(search).Fields(fields).PageSize(pageSize).StartKey(startKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportVulnerabilitiesJSON``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportVulnerabilitiesJSON`: VulnerabilityExportResponse
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportVulnerabilitiesJSON`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportVulnerabilitiesJSONRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 
 **fields** | **string** | A list of fields to export, comma-separated | 
 **pageSize** | **int32** | The number of results to return per request. | 
 **startKey** | **string** | The value to use for requesting the next page when requesting paginated results.  This should be the value of the &#x60;next_key&#x60; attribute returned in the previous response. | 

### Return type

[**VulnerabilityExportResponse**](VulnerabilityExportResponse.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportVulnerabilitiesJSONL

> *os.File ExportVulnerabilitiesJSONL(ctx).Oid(oid).Search(search).Fields(fields).Execute()

Export the vulnerability inventory as JSON line-delimited

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
	resp, r, err := apiClient.ExportAPI.ExportVulnerabilitiesJSONL(context.Background()).Oid(oid).Search(search).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportVulnerabilitiesJSONL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportVulnerabilitiesJSONL`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportVulnerabilitiesJSONL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportVulnerabilitiesJSONLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 
 **fields** | **string** | A list of fields to export, comma-separated | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportWirelessCSV

> *os.File ExportWirelessCSV(ctx).Oid(oid).Search(search).Execute()

Wireless inventory as CSV

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
	resp, r, err := apiClient.ExportAPI.ExportWirelessCSV(context.Background()).Oid(oid).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportWirelessCSV``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportWirelessCSV`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportWirelessCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportWirelessCSVRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 

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


## ExportWirelessJSON

> WirelessExportResponse ExportWirelessJSON(ctx).Oid(oid).Search(search).Fields(fields).PageSize(pageSize).StartKey(startKey).Execute()

Wireless inventory as JSON

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
	pageSize := int32(56) // int32 | The number of results to return per request. (optional)
	startKey := "startKey_example" // string | The value to use for requesting the next page when requesting paginated results.  This should be the value of the `next_key` attribute returned in the previous response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportAPI.ExportWirelessJSON(context.Background()).Oid(oid).Search(search).Fields(fields).PageSize(pageSize).StartKey(startKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportWirelessJSON``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportWirelessJSON`: WirelessExportResponse
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportWirelessJSON`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportWirelessJSONRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 
 **fields** | **string** | A list of fields to export, comma-separated | 
 **pageSize** | **int32** | The number of results to return per request. | 
 **startKey** | **string** | The value to use for requesting the next page when requesting paginated results.  This should be the value of the &#x60;next_key&#x60; attribute returned in the previous response. | 

### Return type

[**WirelessExportResponse**](WirelessExportResponse.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportWirelessJSONL

> *os.File ExportWirelessJSONL(ctx).Oid(oid).Search(search).Fields(fields).Execute()

Wireless inventory as JSON line-delimited

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
	resp, r, err := apiClient.ExportAPI.ExportWirelessJSONL(context.Background()).Oid(oid).Search(search).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportWirelessJSONL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportWirelessJSONL`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.ExportWirelessJSONL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportWirelessJSONLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | The current Organization | 
 **search** | **string** | A search query in runZero search query syntax | 
 **fields** | **string** | A list of fields to export, comma-separated | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

