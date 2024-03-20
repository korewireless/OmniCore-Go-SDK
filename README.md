# Go API client for OmniCore

This is an OmniCore Model and State Management server.

## Overview

- API version: 1.8.16
- Package version: 1.8.16
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

All URIs are relative to *https://api.korewireless.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DeviceApi* | [**BindDevice**](docs/DeviceApi.md#binddevice) | **Post** /omnicore/subscriptions/{subscriptionId}/registries/{registryId}/bindDeviceToGateway | 
*DeviceApi* | [**BindDevices**](docs/DeviceApi.md#binddevices) | **Post** /omnicore/subscriptions/{subscriptionId}/registries/{registryId}/bindDevicesToGateway | 
*DeviceApi* | [**BlockDeviceCommuncation**](docs/DeviceApi.md#blockdevicecommuncation) | **Put** /omnicore/subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/communication | 
*DeviceApi* | [**CreateDevice**](docs/DeviceApi.md#createdevice) | **Post** /omnicore/subscriptions/{subscriptionId}/registries/{registryId}/devices | 
*DeviceApi* | [**DeleteDevice**](docs/DeviceApi.md#deletedevice) | **Delete** /omnicore/subscriptions/{subscriptionId}/registries/{registryId}/devices/{deviceId} | 
*DeviceApi* | [**GetConfig**](docs/DeviceApi.md#getconfig) | **Get** /omnicore/subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/configVersions | 
*DeviceApi* | [**GetDevice**](docs/DeviceApi.md#getdevice) | **Get** /omnicore/subscriptions/{subscriptionId}/registries/{registryId}/devices/{deviceId} | 
*DeviceApi* | [**GetDevices**](docs/DeviceApi.md#getdevices) | **Get** /omnicore/subscriptions/{subscriptionId}/registries/{registryId}/devices | 
*DeviceApi* | [**GetDevicesLastSeen**](docs/DeviceApi.md#getdeviceslastseen) | **Get** /omnicore/subscriptions/{subscriptionId}/devices | 
*DeviceApi* | [**GetStates**](docs/DeviceApi.md#getstates) | **Get** /omnicore/subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/states | 
*DeviceApi* | [**SendCommandToDevice**](docs/DeviceApi.md#sendcommandtodevice) | **Post** /omnicore/subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/sendCommandToDevice | 
*DeviceApi* | [**UnBindDevice**](docs/DeviceApi.md#unbinddevice) | **Post** /omnicore/subscriptions/{subscriptionId}/registries/{registryId}/unbindDeviceFromGateway | 
*DeviceApi* | [**UnBindDevices**](docs/DeviceApi.md#unbinddevices) | **Post** /omnicore/subscriptions/{subscriptionId}/registries/{registryId}/unbindDevicesFromGateway | 
*DeviceApi* | [**UpdateConfigurationToDevice**](docs/DeviceApi.md#updateconfigurationtodevice) | **Post** /omnicore/subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/updateConfigurationToDevice | 
*DeviceApi* | [**UpdateCustomOnboardRequest**](docs/DeviceApi.md#updatecustomonboardrequest) | **Post** /omnicore/subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/updateCustomOnboardRequest | 
*DeviceApi* | [**UpdateDevice**](docs/DeviceApi.md#updatedevice) | **Patch** /omnicore/subscriptions/{subscriptionId}/registries/{registryId}/devices/{deviceId} | 
*MetricsApi* | [**GetMetrics**](docs/MetricsApi.md#getmetrics) | **Get** /omnicore/subscriptions/{subscriptionId}/metrics | 
*RegistryApi* | [**CreateRegistry**](docs/RegistryApi.md#createregistry) | **Post** /omnicore/subscriptions/{subscriptionId}/registries | 
*RegistryApi* | [**DeleteRegistry**](docs/RegistryApi.md#deleteregistry) | **Delete** /omnicore/subscriptions/{subscriptionId}/registries/{registryId} | 
*RegistryApi* | [**GetRegistries**](docs/RegistryApi.md#getregistries) | **Get** /omnicore/subscriptions/{subscriptionId}/registries | 
*RegistryApi* | [**GetRegistry**](docs/RegistryApi.md#getregistry) | **Get** /omnicore/subscriptions/{subscriptionId}/registries/{registryId} | 
*RegistryApi* | [**SendBroadcastToDevices**](docs/RegistryApi.md#sendbroadcasttodevices) | **Post** /omnicore/subscriptions/{subscriptionid}/registries/{registryId}/sendBroadcastToDevice | 
*RegistryApi* | [**UpdateRegistry**](docs/RegistryApi.md#updateregistry) | **Patch** /omnicore/subscriptions/{subscriptionId}/registries/{registryId} | 
*SinkApi* | [**CreateSink**](docs/SinkApi.md#createsink) | **Post** /omnicore/subscriptions/{subscriptionId}/sinks | 
*SinkApi* | [**DeleteSink**](docs/SinkApi.md#deletesink) | **Delete** /omnicore/subscriptions/{subscriptionId}/sinks/{sinkId} | 
*SinkApi* | [**GetASink**](docs/SinkApi.md#getasink) | **Get** /omnicore/subscriptions/{subscriptionId}/sinks/{sinkId} | 
*SinkApi* | [**GetSinks**](docs/SinkApi.md#getsinks) | **Get** /omnicore/subscriptions/{subscriptionId}/sinks | Get All Sinks
*VaultApi* | [**CreateVaultConfiguration**](docs/VaultApi.md#createvaultconfiguration) | **Post** /vault/subscriptions/{subscriptionid}/configurations | 
*VaultApi* | [**CreateVaultKey**](docs/VaultApi.md#createvaultkey) | **Post** /vault/subscriptions/{subscriptionid}/encryptionkeys | 
*VaultApi* | [**DeleteConfiguration**](docs/VaultApi.md#deleteconfiguration) | **Delete** /vault/subscriptions/{subscriptionid}/configurations/{configid} | 
*VaultApi* | [**DeleteVaultKey**](docs/VaultApi.md#deletevaultkey) | **Delete** /vault/subscriptions/{subscriptionid}/encryptionkeys/{keyid} | 
*VaultApi* | [**EnableEncryption**](docs/VaultApi.md#enableencryption) | **Post** /vault/subscriptions/{subscriptionid}/encryption | 
*VaultApi* | [**GetExports**](docs/VaultApi.md#getexports) | **Get** /vault/subscriptions/{subscriptionid}/exports | 
*VaultApi* | [**GetRegistryData**](docs/VaultApi.md#getregistrydata) | **Get** /vault/subscriptions/{subscriptionid}/folders | 
*VaultApi* | [**GetReplays**](docs/VaultApi.md#getreplays) | **Get** /vault/subscriptions/{subscriptionid}/replays | 
*VaultApi* | [**GetVaultAudit**](docs/VaultApi.md#getvaultaudit) | **Get** /vault/subscriptions/{subscriptionid}/audit | 
*VaultApi* | [**GetVaultConfigurations**](docs/VaultApi.md#getvaultconfigurations) | **Get** /vault/subscriptions/{subscriptionid}/configurations | 
*VaultApi* | [**GetVaultFiles**](docs/VaultApi.md#getvaultfiles) | **Get** /vault/subscriptions/{subscriptionid}/registry/{registryid}/files | 
*VaultApi* | [**GetVaultKeys**](docs/VaultApi.md#getvaultkeys) | **Get** /vault/subscriptions/{subscriptionid}/encryptionkeys | 
*VaultApi* | [**GetVaultMetrics**](docs/VaultApi.md#getvaultmetrics) | **Get** /vault/subscriptions/{subscriptionid}/metrics | 
*VaultApi* | [**GetVaultStatus**](docs/VaultApi.md#getvaultstatus) | **Get** /vault/subscriptions/{subscriptionid}/status | 
*VaultApi* | [**SetRetention**](docs/VaultApi.md#setretention) | **Post** /vault/subscriptions/{subscriptionid}/retention | 
*VaultApi* | [**StartExport**](docs/VaultApi.md#startexport) | **Post** /vault/subscriptions/{subscriptionid}/exports | 
*VaultApi* | [**StartReplay**](docs/VaultApi.md#startreplay) | **Post** /vault/subscriptions/{subscriptionid}/replays | 


## Documentation For Models

 - [Audit](docs/Audit.md)
 - [AuditResult](docs/AuditResult.md)
 - [BindRequest](docs/BindRequest.md)
 - [BindRequestIdsGateway](docs/BindRequestIdsGateway.md)
 - [BlockCommunicationBody](docs/BlockCommunicationBody.md)
 - [Config](docs/Config.md)
 - [Configurations](docs/Configurations.md)
 - [CreateConfiguration](docs/CreateConfiguration.md)
 - [CreateVaultKeyBody](docs/CreateVaultKeyBody.md)
 - [CustomOnboard](docs/CustomOnboard.md)
 - [CustomOnboardTcpUdpModelDetails](docs/CustomOnboardTcpUdpModelDetails.md)
 - [Device](docs/Device.md)
 - [DeviceCommand](docs/DeviceCommand.md)
 - [DeviceConfig](docs/DeviceConfig.md)
 - [DeviceConfiguration](docs/DeviceConfiguration.md)
 - [DeviceCredential](docs/DeviceCredential.md)
 - [DeviceOnline](docs/DeviceOnline.md)
 - [DeviceRegistry](docs/DeviceRegistry.md)
 - [DeviceState](docs/DeviceState.md)
 - [EnableEncryptionBody](docs/EnableEncryptionBody.md)
 - [ErrorFrame](docs/ErrorFrame.md)
 - [ErrorStatus](docs/ErrorStatus.md)
 - [EventNotificationConfig](docs/EventNotificationConfig.md)
 - [ExportDetail](docs/ExportDetail.md)
 - [FileDetail](docs/FileDetail.md)
 - [FileDetails](docs/FileDetails.md)
 - [Folder](docs/Folder.md)
 - [FolderData](docs/FolderData.md)
 - [Frame](docs/Frame.md)
 - [GatewayConfig](docs/GatewayConfig.md)
 - [GenericErrorResponse](docs/GenericErrorResponse.md)
 - [GetExportsResponse](docs/GetExportsResponse.md)
 - [GetKeysResponse](docs/GetKeysResponse.md)
 - [GetReplaysResponse](docs/GetReplaysResponse.md)
 - [HttpConfig](docs/HttpConfig.md)
 - [Info](docs/Info.md)
 - [ListDeviceConfigVersionsResponse](docs/ListDeviceConfigVersionsResponse.md)
 - [ListDeviceRegistries](docs/ListDeviceRegistries.md)
 - [ListDeviceStatesResponse](docs/ListDeviceStatesResponse.md)
 - [ListDevicesOnlineResponse](docs/ListDevicesOnlineResponse.md)
 - [ListDevicesResponse](docs/ListDevicesResponse.md)
 - [ListSinks](docs/ListSinks.md)
 - [LogLevel](docs/LogLevel.md)
 - [Metrics](docs/Metrics.md)
 - [MetricsData](docs/MetricsData.md)
 - [MetricsDetails](docs/MetricsDetails.md)
 - [MetricsLogs](docs/MetricsLogs.md)
 - [MetricsResponse](docs/MetricsResponse.md)
 - [MetricsResponseDetails](docs/MetricsResponseDetails.md)
 - [MqttConfig](docs/MqttConfig.md)
 - [NotificationConfig](docs/NotificationConfig.md)
 - [OperationMetrics](docs/OperationMetrics.md)
 - [Policy](docs/Policy.md)
 - [PublicKeyCertificate](docs/PublicKeyCertificate.md)
 - [PublicKeyCredential](docs/PublicKeyCredential.md)
 - [RegistryCredential](docs/RegistryCredential.md)
 - [Replay](docs/Replay.md)
 - [ReplayBody](docs/ReplayBody.md)
 - [SetRetentionBody](docs/SetRetentionBody.md)
 - [Sink](docs/Sink.md)
 - [StartExportBody](docs/StartExportBody.md)
 - [TcpUdpImage](docs/TcpUdpImage.md)
 - [TcpUdpModel](docs/TcpUdpModel.md)
 - [TcpUdpPortDetail](docs/TcpUdpPortDetail.md)
 - [VaultConfiguration](docs/VaultConfiguration.md)
 - [VaultEncryptionKey](docs/VaultEncryptionKey.md)
 - [VaultStatus](docs/VaultStatus.md)
 - [VaultStatusDetails](docs/VaultStatusDetails.md)
 - [X509CertificateDetails](docs/X509CertificateDetails.md)


## Documentation For Authorization



### bearerAuth

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


### apiKey

- **Type**: API key
- **API key parameter name**: x-api-key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: x-api-key and passed in as the auth context for each request.


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

omnicoresupport@korewireless.com

