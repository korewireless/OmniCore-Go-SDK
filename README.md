# Go API client for OmniCore

This is an Omnicore Model and State Management server.

## Overview

- API version: 1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [http://www.korewireless.com](http://www.korewireless.com)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import OmniCore "github.com/korewireless/OmniCore-Go-SDK"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), OmniCore.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), OmniCore.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), OmniCore.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), OmniCore.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://demo-api.omnicore.cloud.korewireless.com/model-state-management*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DeviceApi* | [**BindDevice**](docs/DeviceApi.md#binddevice) | **Post** /subscriptions/{subscriptionId}/registries/{registryId}/bindDeviceToGateway | Gateway - Bind a Device to Gateway
*DeviceApi* | [**BindDevices**](docs/DeviceApi.md#binddevices) | **Post** /subscriptions/{subscriptionId}/registries/{registryId}/bindDevicesToGateway | Gateway - Bind Devices to Gateway
*DeviceApi* | [**BlockDeviceCommuncation**](docs/DeviceApi.md#blockdevicecommuncation) | **Put** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/communication | Block Device Communication
*DeviceApi* | [**CreateDevice**](docs/DeviceApi.md#createdevice) | **Post** /subscriptions/{subscriptionId}/registries/{registryId}/devices | Add New Device
*DeviceApi* | [**DeleteDevice**](docs/DeviceApi.md#deletedevice) | **Delete** /subscriptions/{subscriptionId}/registries/{registryId}/devices/{deviceId} | Remove Device
*DeviceApi* | [**GetConfig**](docs/DeviceApi.md#getconfig) | **Get** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/configVersions | Get Config
*DeviceApi* | [**GetDevice**](docs/DeviceApi.md#getdevice) | **Get** /subscriptions/{subscriptionId}/registries/{registryId}/devices/{deviceId} | Get Device
*DeviceApi* | [**GetDevices**](docs/DeviceApi.md#getdevices) | **Get** /subscriptions/{subscriptionId}/registries/{registryId}/devices | Get All Devices
*DeviceApi* | [**GetStates**](docs/DeviceApi.md#getstates) | **Get** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/states | Get States
*DeviceApi* | [**SendCommandToDevice**](docs/DeviceApi.md#sendcommandtodevice) | **Post** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/sendCommandToDevice | Send Command
*DeviceApi* | [**SendConfigurationToDevice**](docs/DeviceApi.md#sendconfigurationtodevice) | **Post** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/sendConfigurationToDevice | Send Config
*DeviceApi* | [**UnBindDevice**](docs/DeviceApi.md#unbinddevice) | **Post** /subscriptions/{subscriptionId}/registries/{registryId}/unbindDeviceFromGateway | Gateway - UnBind a Device from Gateway
*DeviceApi* | [**UnBindDevices**](docs/DeviceApi.md#unbinddevices) | **Post** /subscriptions/{subscriptionId}/registries/{registryId}/unbindDevicesFromGateway | Gateway - UnBind Devices from Gateway
*DeviceApi* | [**UpdateDevice**](docs/DeviceApi.md#updatedevice) | **Patch** /subscriptions/{subscriptionId}/registries/{registryId}/devices/{deviceId} | Modify Device
*RegistryApi* | [**CreateRegistry**](docs/RegistryApi.md#createregistry) | **Post** /subscriptions/{subscriptionId}/registries | Add New Registry
*RegistryApi* | [**DeleteRegistry**](docs/RegistryApi.md#deleteregistry) | **Delete** /subscriptions/{subscriptionId}/registries/{registryId} | Remove Registry
*RegistryApi* | [**GetRegistries**](docs/RegistryApi.md#getregistries) | **Get** /subscriptions/{subscriptionId}/registries | Get All Registries
*RegistryApi* | [**GetRegistry**](docs/RegistryApi.md#getregistry) | **Get** /subscriptions/{subscriptionId}/registries/{registryId} | Get Registry
*RegistryApi* | [**UpdateRegistry**](docs/RegistryApi.md#updateregistry) | **Patch** /subscriptions/{subscriptionId}/registries/{registryId} | Modify Registry


## Documentation For Models

 - [CreateRegistryRequest](docs/CreateRegistryRequest.md)
 - [OmnicoreBindRequest](docs/OmnicoreBindRequest.md)
 - [OmnicoreBindRequestIdsGateway](docs/OmnicoreBindRequestIdsGateway.md)
 - [OmnicoreBlockCommunicationBody](docs/OmnicoreBlockCommunicationBody.md)
 - [OmnicoreCommand](docs/OmnicoreCommand.md)
 - [OmnicoreConfig](docs/OmnicoreConfig.md)
 - [OmnicoreDevice](docs/OmnicoreDevice.md)
 - [OmnicoreDeviceConfig](docs/OmnicoreDeviceConfig.md)
 - [OmnicoreDeviceCredential](docs/OmnicoreDeviceCredential.md)
 - [OmnicoreDeviceRegistry](docs/OmnicoreDeviceRegistry.md)
 - [OmnicoreDeviceState](docs/OmnicoreDeviceState.md)
 - [OmnicoreErrorFrame](docs/OmnicoreErrorFrame.md)
 - [OmnicoreErrorStatus](docs/OmnicoreErrorStatus.md)
 - [OmnicoreEventNotificationConfig](docs/OmnicoreEventNotificationConfig.md)
 - [OmnicoreGatewayConfig](docs/OmnicoreGatewayConfig.md)
 - [OmnicoreGenericErrorResponse](docs/OmnicoreGenericErrorResponse.md)
 - [OmnicoreHttpConfig](docs/OmnicoreHttpConfig.md)
 - [OmnicoreInfo](docs/OmnicoreInfo.md)
 - [OmnicoreListDeviceConfigVersionsResponse](docs/OmnicoreListDeviceConfigVersionsResponse.md)
 - [OmnicoreListDeviceRegistriesResponse](docs/OmnicoreListDeviceRegistriesResponse.md)
 - [OmnicoreListDeviceStatesResponse](docs/OmnicoreListDeviceStatesResponse.md)
 - [OmnicoreListDevicesResponse](docs/OmnicoreListDevicesResponse.md)
 - [OmnicoreMqttConfig](docs/OmnicoreMqttConfig.md)
 - [OmnicoreNewDevice](docs/OmnicoreNewDevice.md)
 - [OmnicoreNewRegistry](docs/OmnicoreNewRegistry.md)
 - [OmnicorePublicKeyCertificate](docs/OmnicorePublicKeyCertificate.md)
 - [OmnicorePublicKeyCredential](docs/OmnicorePublicKeyCredential.md)
 - [OmnicoreRegistryCredential](docs/OmnicoreRegistryCredential.md)
 - [OmnicoreStateNotificationConfig](docs/OmnicoreStateNotificationConfig.md)
 - [OmnicoreUpdateDevice](docs/OmnicoreUpdateDevice.md)
 - [OmnicoreUpdateRegistry](docs/OmnicoreUpdateRegistry.md)
 - [OmnicoreX509CertificateDetails](docs/OmnicoreX509CertificateDetails.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@korewireless.com

