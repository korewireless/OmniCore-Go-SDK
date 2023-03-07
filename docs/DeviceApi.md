# \DeviceApi

All URIs are relative to *https://demo-api.omnicore.cloud.korewireless.com/model-state-management*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BindDevice**](DeviceApi.md#BindDevice) | **Post** /subscriptions/{subscriptionId}/registries/{registryId}/bindDeviceToGateway | Gateway - Bind a Device to Gateway
[**BindDevices**](DeviceApi.md#BindDevices) | **Post** /subscriptions/{subscriptionId}/registries/{registryId}/bindDevicesToGateway | Gateway - Bind Devices to Gateway
[**BlockDeviceCommuncation**](DeviceApi.md#BlockDeviceCommuncation) | **Put** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/communication | Block Device Communication
[**CreateDevice**](DeviceApi.md#CreateDevice) | **Post** /subscriptions/{subscriptionId}/registries/{registryId}/devices | Add New Device
[**DeleteDevice**](DeviceApi.md#DeleteDevice) | **Delete** /subscriptions/{subscriptionId}/registries/{registryId}/devices/{deviceId} | Remove Device
[**GetConfig**](DeviceApi.md#GetConfig) | **Get** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/configVersions | Get Config
[**GetDevice**](DeviceApi.md#GetDevice) | **Get** /subscriptions/{subscriptionId}/registries/{registryId}/devices/{deviceId} | Get Device
[**GetDevices**](DeviceApi.md#GetDevices) | **Get** /subscriptions/{subscriptionId}/registries/{registryId}/devices | Get All Devices
[**GetStates**](DeviceApi.md#GetStates) | **Get** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/states | Get States
[**SendCommandToDevice**](DeviceApi.md#SendCommandToDevice) | **Post** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/sendCommandToDevice | Send Command
[**SendConfigurationToDevice**](DeviceApi.md#SendConfigurationToDevice) | **Post** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/sendConfigurationToDevice | Send Config
[**UnBindDevice**](DeviceApi.md#UnBindDevice) | **Post** /subscriptions/{subscriptionId}/registries/{registryId}/unbindDeviceFromGateway | Gateway - UnBind a Device from Gateway
[**UnBindDevices**](DeviceApi.md#UnBindDevices) | **Post** /subscriptions/{subscriptionId}/registries/{registryId}/unbindDevicesFromGateway | Gateway - UnBind Devices from Gateway
[**UpdateDevice**](DeviceApi.md#UpdateDevice) | **Patch** /subscriptions/{subscriptionId}/registries/{registryId}/devices/{deviceId} | Modify Device



## BindDevice

> OmnicoreInfo BindDevice(ctx, subscriptionId, registryId).Device(device).Execute()

Gateway - Bind a Device to Gateway



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
    device := *openapiclient.NewOmnicoreBindRequest("DeviceId_example", "GatewayId_example") // OmnicoreBindRequest | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.BindDevice(context.Background(), subscriptionId, registryId).Device(device).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.BindDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BindDevice`: OmnicoreInfo
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.BindDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBindDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **device** | [**OmnicoreBindRequest**](OmnicoreBindRequest.md) | application/json | 

### Return type

[**OmnicoreInfo**](OmnicoreInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BindDevices

> OmnicoreInfo BindDevices(ctx, subscriptionId, registryId).Device(device).Execute()

Gateway - Bind Devices to Gateway



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
    device := *openapiclient.NewOmnicoreBindRequestIdsGateway([]string{"DeviceIds_example"}, "GatewayId_example") // OmnicoreBindRequestIdsGateway | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.BindDevices(context.Background(), subscriptionId, registryId).Device(device).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.BindDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BindDevices`: OmnicoreInfo
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.BindDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBindDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **device** | [**OmnicoreBindRequestIdsGateway**](OmnicoreBindRequestIdsGateway.md) | application/json | 

### Return type

[**OmnicoreInfo**](OmnicoreInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlockDeviceCommuncation

> map[string]interface{} BlockDeviceCommuncation(ctx, subscriptionid, registryId, deviceId).Device(device).Execute()

Block Device Communication



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
    registryId := "registryId_example" // string | Registry ID
    deviceId := "deviceId_example" // string | Device ID
    device := *openapiclient.NewOmnicoreBlockCommunicationBody() // OmnicoreBlockCommunicationBody | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.BlockDeviceCommuncation(context.Background(), subscriptionid, registryId, deviceId).Device(device).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.BlockDeviceCommuncation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BlockDeviceCommuncation`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.BlockDeviceCommuncation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBlockDeviceCommuncationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **device** | [**OmnicoreBlockCommunicationBody**](OmnicoreBlockCommunicationBody.md) | application/json | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDevice

> OmnicoreDevice CreateDevice(ctx, subscriptionId, registryId).Device(device).Execute()

Add New Device



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
    device := *openapiclient.NewOmnicoreNewDevice("Id_example") // OmnicoreNewDevice | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.CreateDevice(context.Background(), subscriptionId, registryId).Device(device).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.CreateDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDevice`: OmnicoreDevice
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.CreateDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **device** | [**OmnicoreNewDevice**](OmnicoreNewDevice.md) | application/json | 

### Return type

[**OmnicoreDevice**](OmnicoreDevice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDevice

> OmnicoreInfo DeleteDevice(ctx, subscriptionId, registryId, deviceId).Execute()

Remove Device



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
    deviceId := "deviceId_example" // string | Device ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.DeleteDevice(context.Background(), subscriptionId, registryId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.DeleteDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDevice`: OmnicoreInfo
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.DeleteDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceRequest struct via the builder pattern


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


## GetConfig

> OmnicoreListDeviceConfigVersionsResponse GetConfig(ctx, subscriptionid, registryId, deviceId).NumVersions(numVersions).Execute()

Get Config



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
    registryId := "registryId_example" // string | Registry ID
    deviceId := "deviceId_example" // string | Device ID
    numVersions := int32(56) // int32 | Device ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.GetConfig(context.Background(), subscriptionid, registryId, deviceId).NumVersions(numVersions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfig`: OmnicoreListDeviceConfigVersionsResponse
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **numVersions** | **int32** | Device ID | 

### Return type

[**OmnicoreListDeviceConfigVersionsResponse**](OmnicoreListDeviceConfigVersionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevice

> OmnicoreDevice GetDevice(ctx, registryId, subscriptionId, deviceId).Execute()

Get Device



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
    registryId := "registryId_example" // string | Registry ID
    subscriptionId := "subscriptionId_example" // string | Subscription ID
    deviceId := "deviceId_example" // string | Device ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.GetDevice(context.Background(), registryId, subscriptionId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDevice`: OmnicoreDevice
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string** | Registry ID | 
**subscriptionId** | **string** | Subscription ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OmnicoreDevice**](OmnicoreDevice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevices

> OmnicoreListDevicesResponse GetDevices(ctx, registryId, subscriptionId).PageNumber(pageNumber).PageSize(pageSize).FieldMask(fieldMask).DeviceIds(deviceIds).DeviceNumIds(deviceNumIds).GatewayListOptionsAssociationsDeviceId(gatewayListOptionsAssociationsDeviceId).GatewayListOptionsAssociationsGatewayId(gatewayListOptionsAssociationsGatewayId).GatewayListOptionsGatewayType(gatewayListOptionsGatewayType).Execute()

Get All Devices



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
    registryId := "registryId_example" // string | Registry ID
    subscriptionId := "subscriptionId_example" // string | Subscription ID
    pageNumber := int32(56) // int32 | Page Number (optional)
    pageSize := int32(56) // int32 | The maximum number of devices to return in the response. If this value is zero, the service will select a default size.  (optional)
    fieldMask := "fieldMask_example" // string | The fields of the Device resource to be returned to the response. The fields id and numId are always returned, along with any other fields specified. A comma-separated list of fully qualified names of fields. Example:  (optional)
    deviceIds := []string{"Inner_example"} // []string | A list of device string IDs. For example, ['device0', 'device12']. If empty, this field is ignored. Maximum IDs: 10,000 (optional)
    deviceNumIds := []string{"Inner_example"} // []string | A list of device numeric IDs. If empty, this field is ignored. Maximum IDs: 10,000. (optional)
    gatewayListOptionsAssociationsDeviceId := "gatewayListOptionsAssociationsDeviceId_example" // string | If set, returns only the gateways with which the specified device is associated. The device ID can be numeric (num_id) or the user-defined string (id). For example, if 456 is specified, returns only the gateways to which the device with num_id 456 is bound. (optional)
    gatewayListOptionsAssociationsGatewayId := "gatewayListOptionsAssociationsGatewayId_example" // string | If set, only devices associated with the specified gateway are returned. The gateway ID can be numeric (num_id) or the user-defined string (id). For example, if 123 is specified, only devices bound to the gateway with num_id 123 are returned (optional)
    gatewayListOptionsGatewayType := "gatewayListOptionsGatewayType_example" // string | If GATEWAY is specified, only gateways are returned. If NON_GATEWAY is specified, only non-gateway devices are returned. If GATEWAY_TYPE_UNSPECIFIED is specified, all devices are returned. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.GetDevices(context.Background(), registryId, subscriptionId).PageNumber(pageNumber).PageSize(pageSize).FieldMask(fieldMask).DeviceIds(deviceIds).DeviceNumIds(deviceNumIds).GatewayListOptionsAssociationsDeviceId(gatewayListOptionsAssociationsDeviceId).GatewayListOptionsAssociationsGatewayId(gatewayListOptionsAssociationsGatewayId).GatewayListOptionsGatewayType(gatewayListOptionsGatewayType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDevices`: OmnicoreListDevicesResponse
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string** | Registry ID | 
**subscriptionId** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageNumber** | **int32** | Page Number | 
 **pageSize** | **int32** | The maximum number of devices to return in the response. If this value is zero, the service will select a default size.  | 
 **fieldMask** | **string** | The fields of the Device resource to be returned to the response. The fields id and numId are always returned, along with any other fields specified. A comma-separated list of fully qualified names of fields. Example:  | 
 **deviceIds** | **[]string** | A list of device string IDs. For example, [&#39;device0&#39;, &#39;device12&#39;]. If empty, this field is ignored. Maximum IDs: 10,000 | 
 **deviceNumIds** | **[]string** | A list of device numeric IDs. If empty, this field is ignored. Maximum IDs: 10,000. | 
 **gatewayListOptionsAssociationsDeviceId** | **string** | If set, returns only the gateways with which the specified device is associated. The device ID can be numeric (num_id) or the user-defined string (id). For example, if 456 is specified, returns only the gateways to which the device with num_id 456 is bound. | 
 **gatewayListOptionsAssociationsGatewayId** | **string** | If set, only devices associated with the specified gateway are returned. The gateway ID can be numeric (num_id) or the user-defined string (id). For example, if 123 is specified, only devices bound to the gateway with num_id 123 are returned | 
 **gatewayListOptionsGatewayType** | **string** | If GATEWAY is specified, only gateways are returned. If NON_GATEWAY is specified, only non-gateway devices are returned. If GATEWAY_TYPE_UNSPECIFIED is specified, all devices are returned. | 

### Return type

[**OmnicoreListDevicesResponse**](OmnicoreListDevicesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStates

> OmnicoreListDeviceStatesResponse GetStates(ctx, subscriptionid, registryId, deviceId).NumStates(numStates).Execute()

Get States



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
    registryId := "registryId_example" // string | Registry ID
    deviceId := "deviceId_example" // string | Device ID
    numStates := int32(56) // int32 | The number of states to list. States are listed in descending order of update time. The maximum number of states retained is 10. If this value is zero, it will return all the states available. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.GetStates(context.Background(), subscriptionid, registryId, deviceId).NumStates(numStates).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStates`: OmnicoreListDeviceStatesResponse
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **numStates** | **int32** | The number of states to list. States are listed in descending order of update time. The maximum number of states retained is 10. If this value is zero, it will return all the states available. | 

### Return type

[**OmnicoreListDeviceStatesResponse**](OmnicoreListDeviceStatesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendCommandToDevice

> map[string]interface{} SendCommandToDevice(ctx, subscriptionid, registryId, deviceId).Device(device).Execute()

Send Command



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
    registryId := "registryId_example" // string | Registry ID
    deviceId := "deviceId_example" // string | Device ID
    device := *openapiclient.NewOmnicoreCommand("BinaryData_example") // OmnicoreCommand | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.SendCommandToDevice(context.Background(), subscriptionid, registryId, deviceId).Device(device).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.SendCommandToDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendCommandToDevice`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.SendCommandToDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendCommandToDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **device** | [**OmnicoreCommand**](OmnicoreCommand.md) | application/json | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendConfigurationToDevice

> OmnicoreDeviceConfig SendConfigurationToDevice(ctx, subscriptionid, registryId, deviceId).Device(device).Execute()

Send Config



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
    registryId := "registryId_example" // string | Registry ID
    deviceId := "deviceId_example" // string | Device ID
    device := *openapiclient.NewOmnicoreConfig("BinaryData_example") // OmnicoreConfig | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.SendConfigurationToDevice(context.Background(), subscriptionid, registryId, deviceId).Device(device).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.SendConfigurationToDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendConfigurationToDevice`: OmnicoreDeviceConfig
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.SendConfigurationToDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendConfigurationToDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **device** | [**OmnicoreConfig**](OmnicoreConfig.md) | application/json | 

### Return type

[**OmnicoreDeviceConfig**](OmnicoreDeviceConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnBindDevice

> OmnicoreInfo UnBindDevice(ctx, subscriptionId, registryId).Device(device).Execute()

Gateway - UnBind a Device from Gateway



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
    device := *openapiclient.NewOmnicoreBindRequest("DeviceId_example", "GatewayId_example") // OmnicoreBindRequest | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.UnBindDevice(context.Background(), subscriptionId, registryId).Device(device).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.UnBindDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnBindDevice`: OmnicoreInfo
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.UnBindDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnBindDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **device** | [**OmnicoreBindRequest**](OmnicoreBindRequest.md) | application/json | 

### Return type

[**OmnicoreInfo**](OmnicoreInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnBindDevices

> OmnicoreInfo UnBindDevices(ctx, subscriptionId, registryId).Device(device).Execute()

Gateway - UnBind Devices from Gateway



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
    device := *openapiclient.NewOmnicoreBindRequestIdsGateway([]string{"DeviceIds_example"}, "GatewayId_example") // OmnicoreBindRequestIdsGateway | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.UnBindDevices(context.Background(), subscriptionId, registryId).Device(device).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.UnBindDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnBindDevices`: OmnicoreInfo
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.UnBindDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnBindDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **device** | [**OmnicoreBindRequestIdsGateway**](OmnicoreBindRequestIdsGateway.md) | application/json | 

### Return type

[**OmnicoreInfo**](OmnicoreInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDevice

> OmnicoreDevice UpdateDevice(ctx, subscriptionId, registryId, deviceId).UpdateMask(updateMask).Device(device).Execute()

Modify Device



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
    deviceId := "deviceId_example" // string | Device ID
    updateMask := "updateMask_example" // string | Required. Only updates the device fields indicated by this mask. The field mask must not be empty, and it must not contain fields that are immutable or only set by the server. Mutable top-level fields: credentials,log_level, blocked, and metadata
    device := *openapiclient.NewOmnicoreUpdateDevice() // OmnicoreUpdateDevice | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.UpdateDevice(context.Background(), subscriptionId, registryId, deviceId).UpdateMask(updateMask).Device(device).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.UpdateDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDevice`: OmnicoreDevice
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.UpdateDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateMask** | **string** | Required. Only updates the device fields indicated by this mask. The field mask must not be empty, and it must not contain fields that are immutable or only set by the server. Mutable top-level fields: credentials,log_level, blocked, and metadata | 
 **device** | [**OmnicoreUpdateDevice**](OmnicoreUpdateDevice.md) | application/json | 

### Return type

[**OmnicoreDevice**](OmnicoreDevice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

