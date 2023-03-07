# \RegistryApi

All URIs are relative to *https://demo-api.omnicore.cloud.korewireless.com/model-state-management*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRegistry**](RegistryApi.md#CreateRegistry) | **Post** /subscriptions/{subscriptionId}/registries | Add New Registry
[**DeleteRegistry**](RegistryApi.md#DeleteRegistry) | **Delete** /subscriptions/{subscriptionId}/registries/{registryId} | Remove Registry
[**GetRegistries**](RegistryApi.md#GetRegistries) | **Get** /subscriptions/{subscriptionId}/registries | Get All Registries
[**GetRegistry**](RegistryApi.md#GetRegistry) | **Get** /subscriptions/{subscriptionId}/registries/{registryId} | Get Registry
[**UpdateRegistry**](RegistryApi.md#UpdateRegistry) | **Patch** /subscriptions/{subscriptionId}/registries/{registryId} | Modify Registry



## CreateRegistry

> OmnicoreDeviceRegistry CreateRegistry(ctx, subscriptionId).Registry(registry).Execute()

Add New Registry



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
    subscriptionId := "subscriptionId_example" // string | Subscription ID
    registry := *openapiclient.NewCreateRegistryRequest("Id_example") // CreateRegistryRequest | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.CreateRegistry(context.Background(), subscriptionId).Registry(registry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.CreateRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRegistry`: OmnicoreDeviceRegistry
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.CreateRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **registry** | [**CreateRegistryRequest**](CreateRegistryRequest.md) | application/json | 

### Return type

[**OmnicoreDeviceRegistry**](OmnicoreDeviceRegistry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRegistry

> OmnicoreInfo DeleteRegistry(ctx, subscriptionId, registryId).Execute()

Remove Registry



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
    subscriptionId := "subscriptionId_example" // string | Subscription ID
    registryId := "registryId_example" // string | Registry ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.DeleteRegistry(context.Background(), subscriptionId, registryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.DeleteRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRegistry`: OmnicoreInfo
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.DeleteRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OmnicoreInfo**](OmnicoreInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegistries

> OmnicoreListDeviceRegistriesResponse GetRegistries(ctx, subscriptionId).PageNumber(pageNumber).PageSize(pageSize).Execute()

Get All Registries



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
    subscriptionId := "subscriptionId_example" // string | Subscription ID
    pageNumber := int32(56) // int32 | Page Number (optional)
    pageSize := int32(56) // int32 | Page Size (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.GetRegistries(context.Background(), subscriptionId).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.GetRegistries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegistries`: OmnicoreListDeviceRegistriesResponse
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.GetRegistries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **int32** | Page Number | 
 **pageSize** | **int32** | Page Size | 

### Return type

[**OmnicoreListDeviceRegistriesResponse**](OmnicoreListDeviceRegistriesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegistry

> OmnicoreDeviceRegistry GetRegistry(ctx, subscriptionId, registryId).Execute()

Get Registry



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
    subscriptionId := "subscriptionId_example" // string | Subscription ID
    registryId := "registryId_example" // string | Registry ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.GetRegistry(context.Background(), subscriptionId, registryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.GetRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegistry`: OmnicoreDeviceRegistry
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.GetRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OmnicoreDeviceRegistry**](OmnicoreDeviceRegistry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRegistry

> OmnicoreDeviceRegistry UpdateRegistry(ctx, subscriptionId, registryId).UpdateMask(updateMask).Registry(registry).Execute()

Modify Registry



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
    subscriptionId := "subscriptionId_example" // string | Subscription ID
    registryId := "registryId_example" // string | Registry ID
    updateMask := "updateMask_example" // string | values to be updated: eventNotificationConfigs,stateNotificationConfig.pubsub_topic_name,logNotificationConfig.pubsub_topic_name,mqttConfig.mqtt_enabled_state,httpConfig.http_enabled_state,logLevel,credentials
    registry := *openapiclient.NewOmnicoreUpdateRegistry() // OmnicoreUpdateRegistry | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.UpdateRegistry(context.Background(), subscriptionId, registryId).UpdateMask(updateMask).Registry(registry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.UpdateRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRegistry`: OmnicoreDeviceRegistry
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.UpdateRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateMask** | **string** | values to be updated: eventNotificationConfigs,stateNotificationConfig.pubsub_topic_name,logNotificationConfig.pubsub_topic_name,mqttConfig.mqtt_enabled_state,httpConfig.http_enabled_state,logLevel,credentials | 
 **registry** | [**OmnicoreUpdateRegistry**](OmnicoreUpdateRegistry.md) | application/json | 

### Return type

[**OmnicoreDeviceRegistry**](OmnicoreDeviceRegistry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

