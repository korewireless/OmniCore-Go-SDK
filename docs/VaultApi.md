# \VaultApi

All URIs are relative to *https://api.korewireless.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVaultConfiguration**](VaultApi.md#CreateVaultConfiguration) | **Post** /vault/subscriptions/{subscriptionid}/configurations | 
[**CreateVaultKey**](VaultApi.md#CreateVaultKey) | **Post** /vault/subscriptions/{subscriptionid}/encryptionkeys | 
[**DeleteConfiguration**](VaultApi.md#DeleteConfiguration) | **Delete** /vault/subscriptions/{subscriptionid}/configurations/{configid} | 
[**DeleteVaultKey**](VaultApi.md#DeleteVaultKey) | **Delete** /vault/subscriptions/{subscriptionid}/encryptionkeys/{keyid} | 
[**EnableEncryption**](VaultApi.md#EnableEncryption) | **Post** /vault/subscriptions/{subscriptionid}/encryption | 
[**GetExports**](VaultApi.md#GetExports) | **Get** /vault/subscriptions/{subscriptionid}/exports | 
[**GetRegistryData**](VaultApi.md#GetRegistryData) | **Get** /vault/subscriptions/{subscriptionid}/folders | 
[**GetReplays**](VaultApi.md#GetReplays) | **Get** /vault/subscriptions/{subscriptionid}/replays | 
[**GetVaultAudit**](VaultApi.md#GetVaultAudit) | **Get** /vault/subscriptions/{subscriptionid}/audit | 
[**GetVaultConfigurations**](VaultApi.md#GetVaultConfigurations) | **Get** /vault/subscriptions/{subscriptionid}/configurations | 
[**GetVaultFiles**](VaultApi.md#GetVaultFiles) | **Get** /vault/subscriptions/{subscriptionid}/registry/{registryid}/files | 
[**GetVaultKeys**](VaultApi.md#GetVaultKeys) | **Get** /vault/subscriptions/{subscriptionid}/encryptionkeys | 
[**GetVaultMetrics**](VaultApi.md#GetVaultMetrics) | **Get** /vault/subscriptions/{subscriptionid}/metrics | 
[**GetVaultStatus**](VaultApi.md#GetVaultStatus) | **Get** /vault/subscriptions/{subscriptionid}/status | 
[**SetRetention**](VaultApi.md#SetRetention) | **Post** /vault/subscriptions/{subscriptionid}/retention | 
[**StartExport**](VaultApi.md#StartExport) | **Post** /vault/subscriptions/{subscriptionid}/exports | 
[**StartReplay**](VaultApi.md#StartReplay) | **Post** /vault/subscriptions/{subscriptionid}/replays | 



## CreateVaultConfiguration

> Frame CreateVaultConfiguration(ctx, subscriptionid).CreateConfiguration(createConfiguration).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/korewireless/OmniCore-Go-SDK"
)

func main() {
    subscriptionid := "subscriptionid_example" // string | Subscription ID
    createConfiguration := *openapiclient.NewCreateConfiguration() // CreateConfiguration | application/json (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultApi.CreateVaultConfiguration(context.Background(), subscriptionid).CreateConfiguration(createConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.CreateVaultConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultConfiguration`: Frame
    fmt.Fprintf(os.Stdout, "Response from `VaultApi.CreateVaultConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createConfiguration** | [**CreateConfiguration**](CreateConfiguration.md) | application/json | 

### Return type

[**Frame**](Frame.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVaultKey

> Frame CreateVaultKey(ctx, subscriptionid).CreateVaultKeyBody(createVaultKeyBody).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/korewireless/OmniCore-Go-SDK"
)

func main() {
    subscriptionid := "subscriptionid_example" // string | Subscription ID
    createVaultKeyBody := *openapiclient.NewCreateVaultKeyBody() // CreateVaultKeyBody | application/json (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultApi.CreateVaultKey(context.Background(), subscriptionid).CreateVaultKeyBody(createVaultKeyBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.CreateVaultKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultKey`: Frame
    fmt.Fprintf(os.Stdout, "Response from `VaultApi.CreateVaultKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createVaultKeyBody** | [**CreateVaultKeyBody**](CreateVaultKeyBody.md) | application/json | 

### Return type

[**Frame**](Frame.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConfiguration

> Frame DeleteConfiguration(ctx, subscriptionid, configid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/korewireless/OmniCore-Go-SDK"
)

func main() {
    subscriptionid := "subscriptionid_example" // string | Subscription ID
    configid := "configid_example" // string | config id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultApi.DeleteConfiguration(context.Background(), subscriptionid, configid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.DeleteConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteConfiguration`: Frame
    fmt.Fprintf(os.Stdout, "Response from `VaultApi.DeleteConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** | Subscription ID | 
**configid** | **string** | config id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Frame**](Frame.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVaultKey

> Frame DeleteVaultKey(ctx, subscriptionid, keyid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/korewireless/OmniCore-Go-SDK"
)

func main() {
    subscriptionid := "subscriptionid_example" // string | Subscription ID
    keyid := "keyid_example" // string | key id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultApi.DeleteVaultKey(context.Background(), subscriptionid, keyid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.DeleteVaultKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVaultKey`: Frame
    fmt.Fprintf(os.Stdout, "Response from `VaultApi.DeleteVaultKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** | Subscription ID | 
**keyid** | **string** | key id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVaultKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Frame**](Frame.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableEncryption

> Frame EnableEncryption(ctx, subscriptionid).EnableEncryptionBody(enableEncryptionBody).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/korewireless/OmniCore-Go-SDK"
)

func main() {
    subscriptionid := "subscriptionid_example" // string | Subscription ID
    enableEncryptionBody := *openapiclient.NewEnableEncryptionBody() // EnableEncryptionBody | application/json (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultApi.EnableEncryption(context.Background(), subscriptionid).EnableEncryptionBody(enableEncryptionBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.EnableEncryption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableEncryption`: Frame
    fmt.Fprintf(os.Stdout, "Response from `VaultApi.EnableEncryption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableEncryptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enableEncryptionBody** | [**EnableEncryptionBody**](EnableEncryptionBody.md) | application/json | 

### Return type

[**Frame**](Frame.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExports

> GetExportsResponse GetExports(ctx, subscriptionid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/korewireless/OmniCore-Go-SDK"
)

func main() {
    subscriptionid := "subscriptionid_example" // string | Subscription ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultApi.GetExports(context.Background(), subscriptionid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.GetExports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExports`: GetExportsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultApi.GetExports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetExportsResponse**](GetExportsResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegistryData

> FolderData GetRegistryData(ctx, subscriptionid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/korewireless/OmniCore-Go-SDK"
)

func main() {
    subscriptionid := "subscriptionid_example" // string | Subscription ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultApi.GetRegistryData(context.Background(), subscriptionid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.GetRegistryData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegistryData`: FolderData
    fmt.Fprintf(os.Stdout, "Response from `VaultApi.GetRegistryData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistryDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FolderData**](FolderData.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReplays

> GetReplaysResponse GetReplays(ctx, subscriptionid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/korewireless/OmniCore-Go-SDK"
)

func main() {
    subscriptionid := "subscriptionid_example" // string | Subscription ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultApi.GetReplays(context.Background(), subscriptionid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.GetReplays``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReplays`: GetReplaysResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultApi.GetReplays`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReplaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetReplaysResponse**](GetReplaysResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVaultAudit

> AuditResult GetVaultAudit(ctx, subscriptionid).PageNumber(pageNumber).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/korewireless/OmniCore-Go-SDK"
)

func main() {
    subscriptionid := "subscriptionid_example" // string | Subscription ID
    pageNumber := int32(56) // int32 | Page Number (optional)
    pageSize := int32(56) // int32 | Page Size (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultApi.GetVaultAudit(context.Background(), subscriptionid).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.GetVaultAudit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultAudit`: AuditResult
    fmt.Fprintf(os.Stdout, "Response from `VaultApi.GetVaultAudit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultAuditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **int32** | Page Number | 
 **pageSize** | **int32** | Page Size | 

### Return type

[**AuditResult**](AuditResult.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVaultConfigurations

> Configurations GetVaultConfigurations(ctx, subscriptionid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/korewireless/OmniCore-Go-SDK"
)

func main() {
    subscriptionid := "subscriptionid_example" // string | Subscription ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultApi.GetVaultConfigurations(context.Background(), subscriptionid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.GetVaultConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultConfigurations`: Configurations
    fmt.Fprintf(os.Stdout, "Response from `VaultApi.GetVaultConfigurations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Configurations**](Configurations.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVaultFiles

> FileDetails GetVaultFiles(ctx, subscriptionid, registryid).FileType(fileType).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/korewireless/OmniCore-Go-SDK"
)

func main() {
    subscriptionid := "subscriptionid_example" // string | Subscription ID
    registryid := "registryid_example" // string | registry ID
    fileType := "fileType_example" // string | file type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultApi.GetVaultFiles(context.Background(), subscriptionid, registryid).FileType(fileType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.GetVaultFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultFiles`: FileDetails
    fmt.Fprintf(os.Stdout, "Response from `VaultApi.GetVaultFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** | Subscription ID | 
**registryid** | **string** | registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fileType** | **string** | file type | 

### Return type

[**FileDetails**](FileDetails.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVaultKeys

> GetKeysResponse GetVaultKeys(ctx, subscriptionid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/korewireless/OmniCore-Go-SDK"
)

func main() {
    subscriptionid := "subscriptionid_example" // string | Subscription ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultApi.GetVaultKeys(context.Background(), subscriptionid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.GetVaultKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultKeys`: GetKeysResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultApi.GetVaultKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetKeysResponse**](GetKeysResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVaultMetrics

> MetricsResponse GetVaultMetrics(ctx, subscriptionid).StartTime(startTime).EndTime(endTime).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/korewireless/OmniCore-Go-SDK"
)

func main() {
    subscriptionid := "subscriptionid_example" // string | Subscription ID
    startTime := "startTime_example" // string | start time (optional)
    endTime := "endTime_example" // string | end time (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultApi.GetVaultMetrics(context.Background(), subscriptionid).StartTime(startTime).EndTime(endTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.GetVaultMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMetrics`: MetricsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultApi.GetVaultMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **string** | start time | 
 **endTime** | **string** | end time | 

### Return type

[**MetricsResponse**](MetricsResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVaultStatus

> VaultStatus GetVaultStatus(ctx, subscriptionid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/korewireless/OmniCore-Go-SDK"
)

func main() {
    subscriptionid := "subscriptionid_example" // string | Subscription ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultApi.GetVaultStatus(context.Background(), subscriptionid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.GetVaultStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultStatus`: VaultStatus
    fmt.Fprintf(os.Stdout, "Response from `VaultApi.GetVaultStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VaultStatus**](VaultStatus.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetRetention

> Frame SetRetention(ctx, subscriptionid).SetRetentionBody(setRetentionBody).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/korewireless/OmniCore-Go-SDK"
)

func main() {
    subscriptionid := "subscriptionid_example" // string | Subscription ID
    setRetentionBody := *openapiclient.NewSetRetentionBody() // SetRetentionBody | application/json (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultApi.SetRetention(context.Background(), subscriptionid).SetRetentionBody(setRetentionBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.SetRetention``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetRetention`: Frame
    fmt.Fprintf(os.Stdout, "Response from `VaultApi.SetRetention`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetRetentionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setRetentionBody** | [**SetRetentionBody**](SetRetentionBody.md) | application/json | 

### Return type

[**Frame**](Frame.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartExport

> Frame StartExport(ctx, subscriptionid).StartExportBody(startExportBody).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/korewireless/OmniCore-Go-SDK"
)

func main() {
    subscriptionid := "subscriptionid_example" // string | Subscription ID
    startExportBody := *openapiclient.NewStartExportBody() // StartExportBody | application/json (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultApi.StartExport(context.Background(), subscriptionid).StartExportBody(startExportBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.StartExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartExport`: Frame
    fmt.Fprintf(os.Stdout, "Response from `VaultApi.StartExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startExportBody** | [**StartExportBody**](StartExportBody.md) | application/json | 

### Return type

[**Frame**](Frame.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartReplay

> string StartReplay(ctx, subscriptionid).ReplayBody(replayBody).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/korewireless/OmniCore-Go-SDK"
)

func main() {
    subscriptionid := "subscriptionid_example" // string | Subscription ID
    replayBody := *openapiclient.NewReplayBody() // ReplayBody | application/json (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultApi.StartReplay(context.Background(), subscriptionid).ReplayBody(replayBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.StartReplay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartReplay`: string
    fmt.Fprintf(os.Stdout, "Response from `VaultApi.StartReplay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartReplayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replayBody** | [**ReplayBody**](ReplayBody.md) | application/json | 

### Return type

**string**

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

