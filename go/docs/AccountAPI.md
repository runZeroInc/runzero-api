# \AccountAPI

All URIs are relative to *https://console.runzero.com/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountAssetOwnershipTypes**](AccountAPI.md#CreateAccountAssetOwnershipTypes) | **Post** /account/assets/ownership-types | Create new asset ownership types
[**CreateAccountCredential**](AccountAPI.md#CreateAccountCredential) | **Put** /account/credentials | Create a new credential
[**CreateAccountCustomIntegration**](AccountAPI.md#CreateAccountCustomIntegration) | **Post** /account/custom-integrations | Create a new custom integration
[**CreateAccountCustomIntegrationAndID**](AccountAPI.md#CreateAccountCustomIntegrationAndID) | **Put** /account/custom-integrations/{customIntegrationId} | Replace custom integration at provided ID
[**CreateAccountGroup**](AccountAPI.md#CreateAccountGroup) | **Post** /account/groups | Create a new group
[**CreateAccountGroupMapping**](AccountAPI.md#CreateAccountGroupMapping) | **Post** /account/sso/groups | Create a new SSO group mapping
[**CreateAccountKey**](AccountAPI.md#CreateAccountKey) | **Put** /account/keys | Create a new key
[**CreateAccountOrganization**](AccountAPI.md#CreateAccountOrganization) | **Put** /account/orgs | Create a new organization
[**CreateAccountOrganizationExportToken**](AccountAPI.md#CreateAccountOrganizationExportToken) | **Post** /account/orgs/{org_id}/exportTokens | Create a new export token for an organization
[**CreateAccountScanTemplate**](AccountAPI.md#CreateAccountScanTemplate) | **Post** /account/tasks/templates | Create a new scan template
[**CreateAccountUser**](AccountAPI.md#CreateAccountUser) | **Put** /account/users | Create a new user account
[**CreateAccountUserInvite**](AccountAPI.md#CreateAccountUserInvite) | **Put** /account/users/invite | Create a new user account and send an email invite
[**DeleteAccountAssetOwnershipType**](AccountAPI.md#DeleteAccountAssetOwnershipType) | **Delete** /account/assets/ownership-types/{ownership_type_id} | Delete a single asset ownership type
[**DeleteAccountAssetOwnershipTypes**](AccountAPI.md#DeleteAccountAssetOwnershipTypes) | **Delete** /account/assets/ownership-types | Delete asset ownership types
[**DeleteAccountCustomIntegration**](AccountAPI.md#DeleteAccountCustomIntegration) | **Delete** /account/custom-integrations/{customIntegrationId} | Delete an custom integration
[**DeleteAccountOrganizationExportToken**](AccountAPI.md#DeleteAccountOrganizationExportToken) | **Delete** /account/orgs/{org_id}/exportTokens/{key_id} | Removes the export token from the specified organization
[**DeleteAccountOrganizationExportTokenDeprecated**](AccountAPI.md#DeleteAccountOrganizationExportTokenDeprecated) | **Delete** /account/orgs/{org_id}/exportToken | Removes the export token from the specified organization
[**ExportEventsJSON**](AccountAPI.md#ExportEventsJSON) | **Get** /account/events.json | System event log as JSON
[**ExportEventsJSONL**](AccountAPI.md#ExportEventsJSONL) | **Get** /account/events.jsonl | System event log as JSON line-delimited
[**GetAPIToken**](AccountAPI.md#GetAPIToken) | **Post** /account/api/token | Generate an access token using an API client
[**GetAccountAgents**](AccountAPI.md#GetAccountAgents) | **Get** /account/agents | Get all agents across all organizations
[**GetAccountAssetOwnershipTypes**](AccountAPI.md#GetAccountAssetOwnershipTypes) | **Get** /account/assets/ownership-types | Get all asset ownership types
[**GetAccountCredential**](AccountAPI.md#GetAccountCredential) | **Get** /account/credentials/{credential_id} | Get credential details
[**GetAccountCredentials**](AccountAPI.md#GetAccountCredentials) | **Get** /account/credentials | Get all account credentials
[**GetAccountCustomIntegration**](AccountAPI.md#GetAccountCustomIntegration) | **Get** /account/custom-integrations/{customIntegrationId} | Get single custom integration
[**GetAccountCustomIntegrations**](AccountAPI.md#GetAccountCustomIntegrations) | **Get** /account/custom-integrations | Get all custom integrations
[**GetAccountGroup**](AccountAPI.md#GetAccountGroup) | **Get** /account/groups/{group_id} | Get group details
[**GetAccountGroupMapping**](AccountAPI.md#GetAccountGroupMapping) | **Get** /account/sso/groups/{group_mapping_id} | Get SSO group mapping details
[**GetAccountGroupMappings**](AccountAPI.md#GetAccountGroupMappings) | **Get** /account/sso/groups | Get all SSO group mappings
[**GetAccountGroups**](AccountAPI.md#GetAccountGroups) | **Get** /account/groups | Get all groups
[**GetAccountKey**](AccountAPI.md#GetAccountKey) | **Get** /account/keys/{key_id} | Get key details
[**GetAccountKeys**](AccountAPI.md#GetAccountKeys) | **Get** /account/keys | Get all active API keys
[**GetAccountLicense**](AccountAPI.md#GetAccountLicense) | **Get** /account/license | Get license details
[**GetAccountOrganization**](AccountAPI.md#GetAccountOrganization) | **Get** /account/orgs/{org_id} | Get organization details
[**GetAccountOrganizationExportToken**](AccountAPI.md#GetAccountOrganizationExportToken) | **Get** /account/orgs/{org_id}/exportTokens/{key_id} | Get export token details
[**GetAccountOrganizationExportTokens**](AccountAPI.md#GetAccountOrganizationExportTokens) | **Get** /account/orgs/{org_id}/exportTokens | Get all active export tokens for an organization
[**GetAccountOrganizations**](AccountAPI.md#GetAccountOrganizations) | **Get** /account/orgs | Get all organization details
[**GetAccountScanTemplate**](AccountAPI.md#GetAccountScanTemplate) | **Get** /account/tasks/templates/{scan_template_id} | Get scan template details
[**GetAccountScanTemplates**](AccountAPI.md#GetAccountScanTemplates) | **Get** /account/tasks/templates | Get all scan templates across all organizations (up to 1000)
[**GetAccountSites**](AccountAPI.md#GetAccountSites) | **Get** /account/sites | Get all sites details across all organizations
[**GetAccountTasks**](AccountAPI.md#GetAccountTasks) | **Get** /account/tasks | Get all task details across all organizations (up to 1000)
[**GetAccountUser**](AccountAPI.md#GetAccountUser) | **Get** /account/users/{user_id} | Get user details
[**GetAccountUsers**](AccountAPI.md#GetAccountUsers) | **Get** /account/users | Get all users
[**RemoveAccountCredential**](AccountAPI.md#RemoveAccountCredential) | **Delete** /account/credentials/{credential_id} | Remove this credential
[**RemoveAccountGroup**](AccountAPI.md#RemoveAccountGroup) | **Delete** /account/groups/{group_id} | Remove this group
[**RemoveAccountGroupMapping**](AccountAPI.md#RemoveAccountGroupMapping) | **Delete** /account/sso/groups/{group_mapping_id} | Remove this SSO group mapping
[**RemoveAccountKey**](AccountAPI.md#RemoveAccountKey) | **Delete** /account/keys/{key_id} | Remove this key
[**RemoveAccountOrganization**](AccountAPI.md#RemoveAccountOrganization) | **Delete** /account/orgs/{org_id} | Remove this organization
[**RemoveAccountScanTemplate**](AccountAPI.md#RemoveAccountScanTemplate) | **Delete** /account/tasks/templates/{scan_template_id} | Remove scan template
[**RemoveAccountUser**](AccountAPI.md#RemoveAccountUser) | **Delete** /account/users/{user_id} | Remove this user
[**ResetAccountUserLockout**](AccountAPI.md#ResetAccountUserLockout) | **Patch** /account/users/{user_id}/resetLockout | Resets the user&#39;s lockout status
[**ResetAccountUserMFA**](AccountAPI.md#ResetAccountUserMFA) | **Patch** /account/users/{user_id}/resetMFA | Resets the user&#39;s MFA tokens
[**ResetAccountUserPassword**](AccountAPI.md#ResetAccountUserPassword) | **Patch** /account/users/{user_id}/resetPassword | Sends the user a password reset email
[**RotateAPIToken**](AccountAPI.md#RotateAPIToken) | **Post** /account/api/rotate | Rotate the API client secret
[**RotateAccountKey**](AccountAPI.md#RotateAccountKey) | **Patch** /account/keys/{key_id}/rotate | Rotates the key secret
[**RotateAccountOrganizationExportToken**](AccountAPI.md#RotateAccountOrganizationExportToken) | **Patch** /account/orgs/{org_id}/exportTokens/{key_id}/rotate | Rotates an organization export token and returns the updated token
[**RotateAccountOrganizationExportTokenDeprecated**](AccountAPI.md#RotateAccountOrganizationExportTokenDeprecated) | **Patch** /account/orgs/{org_id}/exportToken/rotate | Rotates an organization export token and returns the updated token
[**UpdateAccountAssetOwnershipType**](AccountAPI.md#UpdateAccountAssetOwnershipType) | **Patch** /account/assets/ownership-types/{ownership_type_id} | Update a single asset ownership type
[**UpdateAccountAssetOwnershipTypes**](AccountAPI.md#UpdateAccountAssetOwnershipTypes) | **Put** /account/assets/ownership-types | Update asset ownership types
[**UpdateAccountCustomIntegration**](AccountAPI.md#UpdateAccountCustomIntegration) | **Patch** /account/custom-integrations/{customIntegrationId} | Update a single custom integration
[**UpdateAccountGroup**](AccountAPI.md#UpdateAccountGroup) | **Put** /account/groups | Update an existing group
[**UpdateAccountGroupMapping**](AccountAPI.md#UpdateAccountGroupMapping) | **Put** /account/sso/groups | Update an existing SSO group mapping
[**UpdateAccountOrganization**](AccountAPI.md#UpdateAccountOrganization) | **Patch** /account/orgs/{org_id} | Update organization details
[**UpdateAccountScanTemplate**](AccountAPI.md#UpdateAccountScanTemplate) | **Put** /account/tasks/templates | Update scan template
[**UpdateAccountUser**](AccountAPI.md#UpdateAccountUser) | **Patch** /account/users/{user_id} | Update a user&#39;s details



## CreateAccountAssetOwnershipTypes

> []AssetOwnershipType CreateAccountAssetOwnershipTypes(ctx).AssetOwnershipTypePost(assetOwnershipTypePost).Execute()

Create new asset ownership types

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
	assetOwnershipTypePost := []openapiclient.AssetOwnershipTypePost{*openapiclient.NewAssetOwnershipTypePost("Asset Owner")} // []AssetOwnershipTypePost | array of asset ownership types

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CreateAccountAssetOwnershipTypes(context.Background()).AssetOwnershipTypePost(assetOwnershipTypePost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CreateAccountAssetOwnershipTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountAssetOwnershipTypes`: []AssetOwnershipType
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CreateAccountAssetOwnershipTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountAssetOwnershipTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetOwnershipTypePost** | [**[]AssetOwnershipTypePost**](AssetOwnershipTypePost.md) | array of asset ownership types | 

### Return type

[**[]AssetOwnershipType**](AssetOwnershipType.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountCredential

> Credential CreateAccountCredential(ctx).CredentialOptions(credentialOptions).Execute()

Create a new credential

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
	credentialOptions := *openapiclient.NewCredentialOptions() // CredentialOptions | credential parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CreateAccountCredential(context.Background()).CredentialOptions(credentialOptions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CreateAccountCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountCredential`: Credential
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CreateAccountCredential`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentialOptions** | [**CredentialOptions**](CredentialOptions.md) | credential parameters | 

### Return type

[**Credential**](Credential.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountCustomIntegration

> CustomIntegration CreateAccountCustomIntegration(ctx).CustomIntegrationCreate(customIntegrationCreate).Execute()

Create a new custom integration

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
	customIntegrationCreate := *openapiclient.NewCustomIntegrationCreate("my-custom-integration") // CustomIntegrationCreate | The description of the custom integration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CreateAccountCustomIntegration(context.Background()).CustomIntegrationCreate(customIntegrationCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CreateAccountCustomIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountCustomIntegration`: CustomIntegration
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CreateAccountCustomIntegration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountCustomIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customIntegrationCreate** | [**CustomIntegrationCreate**](CustomIntegrationCreate.md) | The description of the custom integration | 

### Return type

[**CustomIntegration**](CustomIntegration.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountCustomIntegrationAndID

> CustomIntegration CreateAccountCustomIntegrationAndID(ctx, customIntegrationId).CustomIntegrationCreate(customIntegrationCreate).Execute()

Replace custom integration at provided ID

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
	customIntegrationCreate := *openapiclient.NewCustomIntegrationCreate("my-custom-integration") // CustomIntegrationCreate | single custom integration type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CreateAccountCustomIntegrationAndID(context.Background(), customIntegrationId).CustomIntegrationCreate(customIntegrationCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CreateAccountCustomIntegrationAndID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountCustomIntegrationAndID`: CustomIntegration
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CreateAccountCustomIntegrationAndID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customIntegrationId** | **string** | UUID of the custom integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountCustomIntegrationAndIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customIntegrationCreate** | [**CustomIntegrationCreate**](CustomIntegrationCreate.md) | single custom integration type | 

### Return type

[**CustomIntegration**](CustomIntegration.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountGroup

> Group CreateAccountGroup(ctx).GroupPost(groupPost).Execute()

Create a new group

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
	groupPost := *openapiclient.NewGroupPost() // GroupPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CreateAccountGroup(context.Background()).GroupPost(groupPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CreateAccountGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountGroup`: Group
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CreateAccountGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupPost** | [**GroupPost**](GroupPost.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountGroupMapping

> GroupMapping CreateAccountGroupMapping(ctx).GroupMapping(groupMapping).Execute()

Create a new SSO group mapping

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
	groupMapping := *openapiclient.NewGroupMapping("f6cfb91a-52ea-4a86-bf9a-5a891a26f52b", "2b096711-4d28-4417-8635-64af4f62c1ae", "basic-attribute", "basic-attribute-value") // GroupMapping | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CreateAccountGroupMapping(context.Background()).GroupMapping(groupMapping).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CreateAccountGroupMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountGroupMapping`: GroupMapping
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CreateAccountGroupMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountGroupMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupMapping** | [**GroupMapping**](GroupMapping.md) |  | 

### Return type

[**GroupMapping**](GroupMapping.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountKey

> OrganizationAPIKey CreateAccountKey(ctx).APIKeyOptions(aPIKeyOptions).Execute()

Create a new key

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
	aPIKeyOptions := *openapiclient.NewAPIKeyOptions() // APIKeyOptions | key parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CreateAccountKey(context.Background()).APIKeyOptions(aPIKeyOptions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CreateAccountKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountKey`: OrganizationAPIKey
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CreateAccountKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aPIKeyOptions** | [**APIKeyOptions**](APIKeyOptions.md) | key parameters | 

### Return type

[**OrganizationAPIKey**](OrganizationAPIKey.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountOrganization

> Organization CreateAccountOrganization(ctx).OrgOptions(orgOptions).Execute()

Create a new organization

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
	orgOptions := *openapiclient.NewOrgOptions() // OrgOptions | organization definition

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CreateAccountOrganization(context.Background()).OrgOptions(orgOptions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CreateAccountOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountOrganization`: Organization
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CreateAccountOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgOptions** | [**OrgOptions**](OrgOptions.md) | organization definition | 

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


## CreateAccountOrganizationExportToken

> ExportToken CreateAccountOrganizationExportToken(ctx, orgId).ExportTokenOptions(exportTokenOptions).Execute()

Create a new export token for an organization

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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the organization to create an export token for
	exportTokenOptions := *openapiclient.NewExportTokenOptions() // ExportTokenOptions | export token parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CreateAccountOrganizationExportToken(context.Background(), orgId).ExportTokenOptions(exportTokenOptions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CreateAccountOrganizationExportToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountOrganizationExportToken`: ExportToken
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CreateAccountOrganizationExportToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | UUID of the organization to create an export token for | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountOrganizationExportTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportTokenOptions** | [**ExportTokenOptions**](ExportTokenOptions.md) | export token parameters | 

### Return type

[**ExportToken**](ExportToken.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountScanTemplate

> ScanTemplate CreateAccountScanTemplate(ctx).ScanTemplateOptions(scanTemplateOptions).Execute()

Create a new scan template

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
	scanTemplateOptions := *openapiclient.NewScanTemplateOptions("My Scan Template", "f6cfb91a-52ea-4a86-bf9a-5a891a26f52b", false, map[string]interface{}{"key": interface{}(123)}) // ScanTemplateOptions | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CreateAccountScanTemplate(context.Background()).ScanTemplateOptions(scanTemplateOptions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CreateAccountScanTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountScanTemplate`: ScanTemplate
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CreateAccountScanTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountScanTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scanTemplateOptions** | [**ScanTemplateOptions**](ScanTemplateOptions.md) |  | 

### Return type

[**ScanTemplate**](ScanTemplate.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountUser

> User CreateAccountUser(ctx).UserOptions(userOptions).Execute()

Create a new user account

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
	userOptions := *openapiclient.NewUserOptions() // UserOptions | user parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CreateAccountUser(context.Background()).UserOptions(userOptions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CreateAccountUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountUser`: User
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CreateAccountUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userOptions** | [**UserOptions**](UserOptions.md) | user parameters | 

### Return type

[**User**](User.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountUserInvite

> User CreateAccountUserInvite(ctx).UserInviteOptions(userInviteOptions).Execute()

Create a new user account and send an email invite

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
	userInviteOptions := *openapiclient.NewUserInviteOptions() // UserInviteOptions | user invite parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CreateAccountUserInvite(context.Background()).UserInviteOptions(userInviteOptions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CreateAccountUserInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountUserInvite`: User
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CreateAccountUserInvite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountUserInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userInviteOptions** | [**UserInviteOptions**](UserInviteOptions.md) | user invite parameters | 

### Return type

[**User**](User.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccountAssetOwnershipType

> []AssetOwnershipType DeleteAccountAssetOwnershipType(ctx, ownershipTypeId).Execute()

Delete a single asset ownership type

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
	ownershipTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the asset ownership type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.DeleteAccountAssetOwnershipType(context.Background(), ownershipTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.DeleteAccountAssetOwnershipType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAccountAssetOwnershipType`: []AssetOwnershipType
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.DeleteAccountAssetOwnershipType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownershipTypeId** | **string** | UUID of the asset ownership type | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountAssetOwnershipTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AssetOwnershipType**](AssetOwnershipType.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccountAssetOwnershipTypes

> []AssetOwnershipType DeleteAccountAssetOwnershipTypes(ctx).RequestBody(requestBody).Execute()

Delete asset ownership types

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
	requestBody := []string{"Property_example"} // []string | Array of ids

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.DeleteAccountAssetOwnershipTypes(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.DeleteAccountAssetOwnershipTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAccountAssetOwnershipTypes`: []AssetOwnershipType
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.DeleteAccountAssetOwnershipTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountAssetOwnershipTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** | Array of ids | 

### Return type

[**[]AssetOwnershipType**](AssetOwnershipType.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccountCustomIntegration

> CustomIntegration DeleteAccountCustomIntegration(ctx, customIntegrationId).Execute()

Delete an custom integration

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
	resp, r, err := apiClient.AccountAPI.DeleteAccountCustomIntegration(context.Background(), customIntegrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.DeleteAccountCustomIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAccountCustomIntegration`: CustomIntegration
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.DeleteAccountCustomIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customIntegrationId** | **string** | UUID of the custom integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountCustomIntegrationRequest struct via the builder pattern


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


## DeleteAccountOrganizationExportToken

> DeleteAccountOrganizationExportToken(ctx, orgId, keyId).Execute()

Removes the export token from the specified organization

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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the organization to retrieve
	keyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the export token ID to remove

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountAPI.DeleteAccountOrganizationExportToken(context.Background(), orgId, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.DeleteAccountOrganizationExportToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | UUID of the organization to retrieve | 
**keyId** | **string** | UUID of the export token ID to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountOrganizationExportTokenRequest struct via the builder pattern


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


## DeleteAccountOrganizationExportTokenDeprecated

> DeleteAccountOrganizationExportTokenDeprecated(ctx, orgId).Execute()

Removes the export token from the specified organization



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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the organization to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountAPI.DeleteAccountOrganizationExportTokenDeprecated(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.DeleteAccountOrganizationExportTokenDeprecated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | UUID of the organization to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountOrganizationExportTokenDeprecatedRequest struct via the builder pattern


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


## ExportEventsJSON

> EventExportResponse ExportEventsJSON(ctx).Search(search).Fields(fields).PageSize(pageSize).StartKey(startKey).Execute()

System event log as JSON

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
	search := "search_example" // string | an optional search string for filtering results (optional)
	fields := "fields_example" // string | an optional list of fields to export, comma-separated (optional)
	pageSize := int32(56) // int32 | The number of results to return per request. (optional)
	startKey := "startKey_example" // string | The value to use for requesting the next page when requesting paginated results.  This should be the value of the `next_key` attribute returned in the previous response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.ExportEventsJSON(context.Background()).Search(search).Fields(fields).PageSize(pageSize).StartKey(startKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.ExportEventsJSON``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportEventsJSON`: EventExportResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.ExportEventsJSON`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportEventsJSONRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | an optional search string for filtering results | 
 **fields** | **string** | an optional list of fields to export, comma-separated | 
 **pageSize** | **int32** | The number of results to return per request. | 
 **startKey** | **string** | The value to use for requesting the next page when requesting paginated results.  This should be the value of the &#x60;next_key&#x60; attribute returned in the previous response. | 

### Return type

[**EventExportResponse**](EventExportResponse.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportEventsJSONL

> *os.File ExportEventsJSONL(ctx).Search(search).Fields(fields).Execute()

System event log as JSON line-delimited

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
	search := "search_example" // string | an optional search string for filtering results (optional)
	fields := "fields_example" // string | an optional list of fields to export, comma-separated (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.ExportEventsJSONL(context.Background()).Search(search).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.ExportEventsJSONL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportEventsJSONL`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.ExportEventsJSONL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportEventsJSONLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | an optional search string for filtering results | 
 **fields** | **string** | an optional list of fields to export, comma-separated | 

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


## GetAPIToken

> AccessToken GetAPIToken(ctx).GrantType(grantType).ClientId(clientId).ClientSecret(clientSecret).Execute()

Generate an access token using an API client

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
	grantType := "grantType_example" // string | 
	clientId := "clientId_example" // string | 
	clientSecret := "clientSecret_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetAPIToken(context.Background()).GrantType(grantType).ClientId(clientId).ClientSecret(clientSecret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAPIToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAPIToken`: AccessToken
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAPIToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAPITokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** |  | 
 **clientId** | **string** |  | 
 **clientSecret** | **string** |  | 

### Return type

[**AccessToken**](AccessToken.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountAgents

> []Agent GetAccountAgents(ctx).Search(search).Execute()

Get all agents across all organizations

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
	search := "search_example" // string | an optional search string for filtering results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetAccountAgents(context.Background()).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountAgents`: []Agent
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountAgents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | an optional search string for filtering results | 

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


## GetAccountAssetOwnershipTypes

> []AssetOwnershipType GetAccountAssetOwnershipTypes(ctx).Execute()

Get all asset ownership types

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
	resp, r, err := apiClient.AccountAPI.GetAccountAssetOwnershipTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountAssetOwnershipTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountAssetOwnershipTypes`: []AssetOwnershipType
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountAssetOwnershipTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountAssetOwnershipTypesRequest struct via the builder pattern


### Return type

[**[]AssetOwnershipType**](AssetOwnershipType.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountCredential

> Credential GetAccountCredential(ctx, credentialId).Execute()

Get credential details

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
	credentialId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the credential to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetAccountCredential(context.Background(), credentialId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountCredential`: Credential
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialId** | **string** | UUID of the credential to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Credential**](Credential.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountCredentials

> []Credential GetAccountCredentials(ctx).Search(search).Execute()

Get all account credentials

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
	search := "search_example" // string | an optional search string for filtering results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetAccountCredentials(context.Background()).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountCredentials`: []Credential
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | an optional search string for filtering results | 

### Return type

[**[]Credential**](Credential.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountCustomIntegration

> CustomIntegration GetAccountCustomIntegration(ctx, customIntegrationId).Execute()

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
	resp, r, err := apiClient.AccountAPI.GetAccountCustomIntegration(context.Background(), customIntegrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountCustomIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountCustomIntegration`: CustomIntegration
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountCustomIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customIntegrationId** | **string** | UUID of the custom integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountCustomIntegrationRequest struct via the builder pattern


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


## GetAccountCustomIntegrations

> CustomIntegration GetAccountCustomIntegrations(ctx).Execute()

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
	resp, r, err := apiClient.AccountAPI.GetAccountCustomIntegrations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountCustomIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountCustomIntegrations`: CustomIntegration
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountCustomIntegrations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountCustomIntegrationsRequest struct via the builder pattern


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


## GetAccountGroup

> Group GetAccountGroup(ctx, groupId).Execute()

Get group details

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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetAccountGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountGroup`: Group
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | UUID of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Group**](Group.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountGroupMapping

> GroupMapping GetAccountGroupMapping(ctx, groupMappingId).Execute()

Get SSO group mapping details

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
	groupMappingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the SSO group mapping

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetAccountGroupMapping(context.Background(), groupMappingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountGroupMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountGroupMapping`: GroupMapping
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountGroupMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupMappingId** | **string** | UUID of the SSO group mapping | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountGroupMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupMapping**](GroupMapping.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountGroupMappings

> GroupMapping GetAccountGroupMappings(ctx).Execute()

Get all SSO group mappings

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
	resp, r, err := apiClient.AccountAPI.GetAccountGroupMappings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountGroupMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountGroupMappings`: GroupMapping
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountGroupMappings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountGroupMappingsRequest struct via the builder pattern


### Return type

[**GroupMapping**](GroupMapping.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountGroups

> Group GetAccountGroups(ctx).Execute()

Get all groups

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
	resp, r, err := apiClient.AccountAPI.GetAccountGroups(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountGroups`: Group
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountGroupsRequest struct via the builder pattern


### Return type

[**Group**](Group.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountKey

> OrganizationAPIKey GetAccountKey(ctx, keyId).Execute()

Get key details

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
	keyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the key to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetAccountKey(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountKey`: OrganizationAPIKey
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | UUID of the key to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetAccountKeys

> []OrganizationAPIKey GetAccountKeys(ctx).Execute()

Get all active API keys

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
	resp, r, err := apiClient.AccountAPI.GetAccountKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountKeys`: []OrganizationAPIKey
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountKeysRequest struct via the builder pattern


### Return type

[**[]OrganizationAPIKey**](OrganizationAPIKey.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountLicense

> License GetAccountLicense(ctx).Execute()

Get license details

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
	resp, r, err := apiClient.AccountAPI.GetAccountLicense(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountLicense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountLicense`: License
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountLicense`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountLicenseRequest struct via the builder pattern


### Return type

[**License**](License.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountOrganization

> Organization GetAccountOrganization(ctx, orgId).Execute()

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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the organization to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetAccountOrganization(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountOrganization`: Organization
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | UUID of the organization to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetAccountOrganizationExportToken

> ExportToken GetAccountOrganizationExportToken(ctx, orgId, keyId).Execute()

Get export token details

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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the organization to retrieve
	keyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the export token ID to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetAccountOrganizationExportToken(context.Background(), orgId, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountOrganizationExportToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountOrganizationExportToken`: ExportToken
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountOrganizationExportToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | UUID of the organization to retrieve | 
**keyId** | **string** | UUID of the export token ID to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountOrganizationExportTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ExportToken**](ExportToken.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountOrganizationExportTokens

> []ExportToken GetAccountOrganizationExportTokens(ctx, orgId).Execute()

Get all active export tokens for an organization

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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the organization to retrieve export tokens for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetAccountOrganizationExportTokens(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountOrganizationExportTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountOrganizationExportTokens`: []ExportToken
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountOrganizationExportTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | UUID of the organization to retrieve export tokens for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountOrganizationExportTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ExportToken**](ExportToken.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountOrganizations

> []Organization GetAccountOrganizations(ctx).Search(search).Execute()

Get all organization details

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
	search := "search_example" // string | an optional search string for filtering results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetAccountOrganizations(context.Background()).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountOrganizations`: []Organization
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountOrganizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | an optional search string for filtering results | 

### Return type

[**[]Organization**](Organization.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountScanTemplate

> ScanTemplate GetAccountScanTemplate(ctx, scanTemplateId).Execute()

Get scan template details

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
	scanTemplateId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the scan template to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetAccountScanTemplate(context.Background(), scanTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountScanTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountScanTemplate`: ScanTemplate
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountScanTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scanTemplateId** | **string** | UUID of the scan template to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountScanTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScanTemplate**](ScanTemplate.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountScanTemplates

> []ScanTemplate GetAccountScanTemplates(ctx).Search(search).Execute()

Get all scan templates across all organizations (up to 1000)

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
	search := "search_example" // string | an optional search string for filtering results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetAccountScanTemplates(context.Background()).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountScanTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountScanTemplates`: []ScanTemplate
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountScanTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountScanTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | an optional search string for filtering results | 

### Return type

[**[]ScanTemplate**](ScanTemplate.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountSites

> []Site GetAccountSites(ctx).Search(search).Execute()

Get all sites details across all organizations

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
	search := "search_example" // string | an optional search string for filtering results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetAccountSites(context.Background()).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountSites`: []Site
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountSites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | an optional search string for filtering results | 

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


## GetAccountTasks

> []Task GetAccountTasks(ctx).Search(search).Execute()

Get all task details across all organizations (up to 1000)

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
	search := "search_example" // string | an optional search string for filtering results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetAccountTasks(context.Background()).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountTasks`: []Task
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


## GetAccountUser

> User GetAccountUser(ctx, userId).Execute()

Get user details

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the user to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetAccountUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountUser`: User
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | UUID of the user to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountUsers

> []User GetAccountUsers(ctx).Execute()

Get all users

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
	resp, r, err := apiClient.AccountAPI.GetAccountUsers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountUsers`: []User
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountUsersRequest struct via the builder pattern


### Return type

[**[]User**](User.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAccountCredential

> RemoveAccountCredential(ctx, credentialId).Execute()

Remove this credential

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
	credentialId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the credential to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountAPI.RemoveAccountCredential(context.Background(), credentialId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.RemoveAccountCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialId** | **string** | UUID of the credential to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAccountCredentialRequest struct via the builder pattern


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


## RemoveAccountGroup

> RemoveAccountGroup(ctx, groupId).Execute()

Remove this group

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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountAPI.RemoveAccountGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.RemoveAccountGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | UUID of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAccountGroupRequest struct via the builder pattern


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


## RemoveAccountGroupMapping

> RemoveAccountGroupMapping(ctx, groupMappingId).Execute()

Remove this SSO group mapping

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
	groupMappingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the SSO group mapping

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountAPI.RemoveAccountGroupMapping(context.Background(), groupMappingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.RemoveAccountGroupMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupMappingId** | **string** | UUID of the SSO group mapping | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAccountGroupMappingRequest struct via the builder pattern


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


## RemoveAccountKey

> RemoveAccountKey(ctx, keyId).Execute()

Remove this key

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
	keyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the key to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountAPI.RemoveAccountKey(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.RemoveAccountKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | UUID of the key to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAccountKeyRequest struct via the builder pattern


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


## RemoveAccountOrganization

> RemoveAccountOrganization(ctx, orgId).Execute()

Remove this organization

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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the organization to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountAPI.RemoveAccountOrganization(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.RemoveAccountOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | UUID of the organization to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAccountOrganizationRequest struct via the builder pattern


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


## RemoveAccountScanTemplate

> ScanTemplate RemoveAccountScanTemplate(ctx, scanTemplateId).Execute()

Remove scan template

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
	scanTemplateId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the scan template to remove

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.RemoveAccountScanTemplate(context.Background(), scanTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.RemoveAccountScanTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveAccountScanTemplate`: ScanTemplate
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.RemoveAccountScanTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scanTemplateId** | **string** | UUID of the scan template to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAccountScanTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScanTemplate**](ScanTemplate.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAccountUser

> RemoveAccountUser(ctx, userId).Execute()

Remove this user

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the user to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountAPI.RemoveAccountUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.RemoveAccountUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | UUID of the user to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAccountUserRequest struct via the builder pattern


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


## ResetAccountUserLockout

> User ResetAccountUserLockout(ctx, userId).Execute()

Resets the user's lockout status

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the user to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.ResetAccountUserLockout(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.ResetAccountUserLockout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetAccountUserLockout`: User
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.ResetAccountUserLockout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | UUID of the user to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetAccountUserLockoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetAccountUserMFA

> User ResetAccountUserMFA(ctx, userId).Execute()

Resets the user's MFA tokens

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the user to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.ResetAccountUserMFA(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.ResetAccountUserMFA``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetAccountUserMFA`: User
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.ResetAccountUserMFA`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | UUID of the user to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetAccountUserMFARequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetAccountUserPassword

> User ResetAccountUserPassword(ctx, userId).Execute()

Sends the user a password reset email

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the user to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.ResetAccountUserPassword(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.ResetAccountUserPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetAccountUserPassword`: User
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.ResetAccountUserPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | UUID of the user to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetAccountUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateAPIToken

> APIClientCredentials RotateAPIToken(ctx).GrantType(grantType).ClientId(clientId).ClientSecret(clientSecret).Execute()

Rotate the API client secret

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
	grantType := "grantType_example" // string | 
	clientId := "clientId_example" // string | 
	clientSecret := "clientSecret_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.RotateAPIToken(context.Background()).GrantType(grantType).ClientId(clientId).ClientSecret(clientSecret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.RotateAPIToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateAPIToken`: APIClientCredentials
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.RotateAPIToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRotateAPITokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** |  | 
 **clientId** | **string** |  | 
 **clientSecret** | **string** |  | 

### Return type

[**APIClientCredentials**](APIClientCredentials.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateAccountKey

> OrganizationAPIKey RotateAccountKey(ctx, keyId).Execute()

Rotates the key secret

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
	keyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the key to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.RotateAccountKey(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.RotateAccountKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateAccountKey`: OrganizationAPIKey
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.RotateAccountKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | UUID of the key to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRotateAccountKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## RotateAccountOrganizationExportToken

> ExportToken RotateAccountOrganizationExportToken(ctx, orgId, keyId).Execute()

Rotates an organization export token and returns the updated token

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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the organization to retrieve
	keyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the export token ID to rotate

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.RotateAccountOrganizationExportToken(context.Background(), orgId, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.RotateAccountOrganizationExportToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateAccountOrganizationExportToken`: ExportToken
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.RotateAccountOrganizationExportToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | UUID of the organization to retrieve | 
**keyId** | **string** | UUID of the export token ID to rotate | 

### Other Parameters

Other parameters are passed through a pointer to a apiRotateAccountOrganizationExportTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ExportToken**](ExportToken.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateAccountOrganizationExportTokenDeprecated

> ExportToken RotateAccountOrganizationExportTokenDeprecated(ctx, orgId).Execute()

Rotates an organization export token and returns the updated token



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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the organization to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.RotateAccountOrganizationExportTokenDeprecated(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.RotateAccountOrganizationExportTokenDeprecated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateAccountOrganizationExportTokenDeprecated`: ExportToken
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.RotateAccountOrganizationExportTokenDeprecated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | UUID of the organization to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRotateAccountOrganizationExportTokenDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExportToken**](ExportToken.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountAssetOwnershipType

> AssetOwnershipType UpdateAccountAssetOwnershipType(ctx, ownershipTypeId).AssetOwnershipTypePost(assetOwnershipTypePost).Execute()

Update a single asset ownership type

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
	ownershipTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the asset ownership type
	assetOwnershipTypePost := *openapiclient.NewAssetOwnershipTypePost("Asset Owner") // AssetOwnershipTypePost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UpdateAccountAssetOwnershipType(context.Background(), ownershipTypeId).AssetOwnershipTypePost(assetOwnershipTypePost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UpdateAccountAssetOwnershipType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccountAssetOwnershipType`: AssetOwnershipType
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UpdateAccountAssetOwnershipType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownershipTypeId** | **string** | UUID of the asset ownership type | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountAssetOwnershipTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assetOwnershipTypePost** | [**AssetOwnershipTypePost**](AssetOwnershipTypePost.md) |  | 

### Return type

[**AssetOwnershipType**](AssetOwnershipType.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountAssetOwnershipTypes

> []AssetOwnershipType UpdateAccountAssetOwnershipTypes(ctx).AssetOwnershipType(assetOwnershipType).Execute()

Update asset ownership types

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
	assetOwnershipType := []openapiclient.AssetOwnershipType{*openapiclient.NewAssetOwnershipType("f6cfb91a-52ea-4a86-bf9a-5a891a26f52b", "Asset Owner")} // []AssetOwnershipType | array of asset ownership types

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UpdateAccountAssetOwnershipTypes(context.Background()).AssetOwnershipType(assetOwnershipType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UpdateAccountAssetOwnershipTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccountAssetOwnershipTypes`: []AssetOwnershipType
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UpdateAccountAssetOwnershipTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountAssetOwnershipTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetOwnershipType** | [**[]AssetOwnershipType**](AssetOwnershipType.md) | array of asset ownership types | 

### Return type

[**[]AssetOwnershipType**](AssetOwnershipType.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountCustomIntegration

> CustomIntegration UpdateAccountCustomIntegration(ctx, customIntegrationId).BaseCustomIntegration(baseCustomIntegration).Execute()

Update a single custom integration

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
	baseCustomIntegration := *openapiclient.NewBaseCustomIntegration() // BaseCustomIntegration | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UpdateAccountCustomIntegration(context.Background(), customIntegrationId).BaseCustomIntegration(baseCustomIntegration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UpdateAccountCustomIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccountCustomIntegration`: CustomIntegration
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UpdateAccountCustomIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customIntegrationId** | **string** | UUID of the custom integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountCustomIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **baseCustomIntegration** | [**BaseCustomIntegration**](BaseCustomIntegration.md) |  | 

### Return type

[**CustomIntegration**](CustomIntegration.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountGroup

> Group UpdateAccountGroup(ctx).GroupPut(groupPut).Execute()

Update an existing group

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
	groupPut := *openapiclient.NewGroupPut() // GroupPut | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UpdateAccountGroup(context.Background()).GroupPut(groupPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UpdateAccountGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccountGroup`: Group
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UpdateAccountGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupPut** | [**GroupPut**](GroupPut.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountGroupMapping

> GroupMapping UpdateAccountGroupMapping(ctx).GroupMapping(groupMapping).Execute()

Update an existing SSO group mapping

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
	groupMapping := *openapiclient.NewGroupMapping("f6cfb91a-52ea-4a86-bf9a-5a891a26f52b", "2b096711-4d28-4417-8635-64af4f62c1ae", "basic-attribute", "basic-attribute-value") // GroupMapping | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UpdateAccountGroupMapping(context.Background()).GroupMapping(groupMapping).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UpdateAccountGroupMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccountGroupMapping`: GroupMapping
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UpdateAccountGroupMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountGroupMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupMapping** | [**GroupMapping**](GroupMapping.md) |  | 

### Return type

[**GroupMapping**](GroupMapping.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountOrganization

> Organization UpdateAccountOrganization(ctx, orgId).OrgOptions(orgOptions).Execute()

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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the organization to retrieve
	orgOptions := *openapiclient.NewOrgOptions() // OrgOptions | organization options

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UpdateAccountOrganization(context.Background(), orgId).OrgOptions(orgOptions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UpdateAccountOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccountOrganization`: Organization
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UpdateAccountOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | UUID of the organization to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgOptions** | [**OrgOptions**](OrgOptions.md) | organization options | 

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


## UpdateAccountScanTemplate

> ScanTemplate UpdateAccountScanTemplate(ctx).ScanTemplate(scanTemplate).Execute()

Update scan template

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
	scanTemplate := *openapiclient.NewScanTemplate("e77602e0-3fb8-4734-aef9-fbc6fdcb0fa8", "e77602e0-3fb8-4734-aef9-fbc6fdcb0fa8", false, map[string]interface{}{"key": interface{}(123)}) // ScanTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UpdateAccountScanTemplate(context.Background()).ScanTemplate(scanTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UpdateAccountScanTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccountScanTemplate`: ScanTemplate
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UpdateAccountScanTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountScanTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scanTemplate** | [**ScanTemplate**](ScanTemplate.md) |  | 

### Return type

[**ScanTemplate**](ScanTemplate.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountUser

> User UpdateAccountUser(ctx, userId).UserOptions(userOptions).Execute()

Update a user's details

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the user to retrieve
	userOptions := *openapiclient.NewUserOptions() // UserOptions | user parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UpdateAccountUser(context.Background(), userId).UserOptions(userOptions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UpdateAccountUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccountUser`: User
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UpdateAccountUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | UUID of the user to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userOptions** | [**UserOptions**](UserOptions.md) | user parameters | 

### Return type

[**User**](User.md)

### Authorization

[oauthDefaults](../README.md#oauthDefaults), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

