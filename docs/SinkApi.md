# \SinkApi

All URIs are relative to *https://api.korewireless.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSink**](SinkApi.md#CreateSink) | **Post** /omnicore/subscriptions/{subscriptionId}/sinks | 
[**DeleteSink**](SinkApi.md#DeleteSink) | **Delete** /omnicore/subscriptions/{subscriptionId}/sinks/{sinkId} | 
[**GetASink**](SinkApi.md#GetASink) | **Get** /omnicore/subscriptions/{subscriptionId}/sinks/{sinkId} | 
[**GetSinks**](SinkApi.md#GetSinks) | **Get** /omnicore/subscriptions/{subscriptionId}/sinks | Get All Sinks



## CreateSink

> Sink CreateSink(ctx, subscriptionId).Sink(sink).Execute()





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
    sink := *openapiclient.NewSink() // Sink |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SinkApi.CreateSink(context.Background(), subscriptionId).Sink(sink).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SinkApi.CreateSink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSink`: Sink
    fmt.Fprintf(os.Stdout, "Response from `SinkApi.CreateSink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sink** | [**Sink**](Sink.md) |  | 

### Return type

[**Sink**](Sink.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSink

> map[string]interface{} DeleteSink(ctx, subscriptionId, sinkId).Execute()





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
    sinkId := "sinkId_example" // string | Sink ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SinkApi.DeleteSink(context.Background(), subscriptionId, sinkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SinkApi.DeleteSink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSink`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SinkApi.DeleteSink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 
**sinkId** | **string** | Sink ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetASink

> Sink GetASink(ctx, subscriptionId, sinkId).Execute()





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
    sinkId := "sinkId_example" // string | Sink ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SinkApi.GetASink(context.Background(), subscriptionId, sinkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SinkApi.GetASink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetASink`: Sink
    fmt.Fprintf(os.Stdout, "Response from `SinkApi.GetASink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 
**sinkId** | **string** | Sink ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetASinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Sink**](Sink.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSinks

> ListSinks GetSinks(ctx, subscriptionId).Execute()

Get All Sinks



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SinkApi.GetSinks(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SinkApi.GetSinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSinks`: ListSinks
    fmt.Fprintf(os.Stdout, "Response from `SinkApi.GetSinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListSinks**](ListSinks.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

