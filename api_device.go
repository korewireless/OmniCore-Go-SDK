/* Copyright 2018-2020 KoreWireless
 *
 * This is part of the KoreWireless Omnicore SDK.
 * It is licensed under the BSD 3-Clause license; you may not use this file
 * except in compliance with the License.
 *
 * You may obtain a copy of the License at:
 *  https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */


package OmniCore

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// DeviceApiService DeviceApi service
type DeviceApiService service

type ApiBindDeviceRequest struct {
	ctx context.Context
	ApiService *DeviceApiService
	subscriptionId string
	registryId string
	device *BindRequest
}

// application/json
func (r ApiBindDeviceRequest) Device(device BindRequest) ApiBindDeviceRequest {
	r.device = &device
	return r
}

func (r ApiBindDeviceRequest) Execute() (*Info, *http.Response, error) {
	return r.ApiService.BindDeviceExecute(r)
}

/*
BindDevice Method for BindDevice

Bind  a device to a gateway under a registry

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subscriptionId Subscription ID
 @param registryId Registry ID
 @return ApiBindDeviceRequest
*/
func (a *DeviceApiService) BindDevice(ctx context.Context, subscriptionId string, registryId string) ApiBindDeviceRequest {
	return ApiBindDeviceRequest{
		ApiService: a,
		ctx: ctx,
		subscriptionId: subscriptionId,
		registryId: registryId,
	}
}

// Execute executes the request
//  @return Info
func (a *DeviceApiService) BindDeviceExecute(r ApiBindDeviceRequest) (*Info, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Info
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceApiService.BindDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscriptions/{subscriptionId}/registries/{registryId}/bindDeviceToGateway"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", url.PathEscape(parameterValueToString(r.subscriptionId, "subscriptionId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"registryId"+"}", url.PathEscape(parameterValueToString(r.registryId, "registryId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.device == nil {
		return localVarReturnValue, nil, reportError("device is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.device
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiBindDevicesRequest struct {
	ctx context.Context
	ApiService *DeviceApiService
	subscriptionId string
	registryId string
	device *BindRequestIdsGateway
}

// application/json
func (r ApiBindDevicesRequest) Device(device BindRequestIdsGateway) ApiBindDevicesRequest {
	r.device = &device
	return r
}

func (r ApiBindDevicesRequest) Execute() (*Info, *http.Response, error) {
	return r.ApiService.BindDevicesExecute(r)
}

/*
BindDevices Method for BindDevices

Bind devices to a gateway under a registry

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subscriptionId Subscription ID
 @param registryId Registry ID
 @return ApiBindDevicesRequest
*/
func (a *DeviceApiService) BindDevices(ctx context.Context, subscriptionId string, registryId string) ApiBindDevicesRequest {
	return ApiBindDevicesRequest{
		ApiService: a,
		ctx: ctx,
		subscriptionId: subscriptionId,
		registryId: registryId,
	}
}

// Execute executes the request
//  @return Info
func (a *DeviceApiService) BindDevicesExecute(r ApiBindDevicesRequest) (*Info, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Info
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceApiService.BindDevices")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscriptions/{subscriptionId}/registries/{registryId}/bindDevicesToGateway"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", url.PathEscape(parameterValueToString(r.subscriptionId, "subscriptionId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"registryId"+"}", url.PathEscape(parameterValueToString(r.registryId, "registryId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.device == nil {
		return localVarReturnValue, nil, reportError("device is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.device
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiBlockDeviceCommuncationRequest struct {
	ctx context.Context
	ApiService *DeviceApiService
	subscriptionid string
	registryId string
	deviceId string
	device *BlockCommunicationBody
}

// application/json
func (r ApiBlockDeviceCommuncationRequest) Device(device BlockCommunicationBody) ApiBlockDeviceCommuncationRequest {
	r.device = &device
	return r
}

func (r ApiBlockDeviceCommuncationRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.BlockDeviceCommuncationExecute(r)
}

/*
BlockDeviceCommuncation Method for BlockDeviceCommuncation

Blocks All Communication From A Device

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subscriptionid Subscription ID
 @param registryId Registry ID
 @param deviceId Device ID
 @return ApiBlockDeviceCommuncationRequest
*/
func (a *DeviceApiService) BlockDeviceCommuncation(ctx context.Context, subscriptionid string, registryId string, deviceId string) ApiBlockDeviceCommuncationRequest {
	return ApiBlockDeviceCommuncationRequest{
		ApiService: a,
		ctx: ctx,
		subscriptionid: subscriptionid,
		registryId: registryId,
		deviceId: deviceId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *DeviceApiService) BlockDeviceCommuncationExecute(r ApiBlockDeviceCommuncationRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceApiService.BlockDeviceCommuncation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/communication"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionid"+"}", url.PathEscape(parameterValueToString(r.subscriptionid, "subscriptionid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"registryId"+"}", url.PathEscape(parameterValueToString(r.registryId, "registryId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"deviceId"+"}", url.PathEscape(parameterValueToString(r.deviceId, "deviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.device == nil {
		return localVarReturnValue, nil, reportError("device is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.device
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreateDeviceRequest struct {
	ctx context.Context
	ApiService *DeviceApiService
	subscriptionId string
	registryId string
	device *NewDevice
}

// application/json
func (r ApiCreateDeviceRequest) Device(device NewDevice) ApiCreateDeviceRequest {
	r.device = &device
	return r
}

func (r ApiCreateDeviceRequest) Execute() (*Device, *http.Response, error) {
	return r.ApiService.CreateDeviceExecute(r)
}

/*
CreateDevice Method for CreateDevice

Create a device under a registry

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subscriptionId Subscription ID
 @param registryId Registry ID
 @return ApiCreateDeviceRequest
*/
func (a *DeviceApiService) CreateDevice(ctx context.Context, subscriptionId string, registryId string) ApiCreateDeviceRequest {
	return ApiCreateDeviceRequest{
		ApiService: a,
		ctx: ctx,
		subscriptionId: subscriptionId,
		registryId: registryId,
	}
}

// Execute executes the request
//  @return Device
func (a *DeviceApiService) CreateDeviceExecute(r ApiCreateDeviceRequest) (*Device, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Device
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceApiService.CreateDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscriptions/{subscriptionId}/registries/{registryId}/devices"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", url.PathEscape(parameterValueToString(r.subscriptionId, "subscriptionId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"registryId"+"}", url.PathEscape(parameterValueToString(r.registryId, "registryId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.device == nil {
		return localVarReturnValue, nil, reportError("device is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.device
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteDeviceRequest struct {
	ctx context.Context
	ApiService *DeviceApiService
	subscriptionId string
	registryId string
	deviceId string
}

func (r ApiDeleteDeviceRequest) Execute() (*Info, *http.Response, error) {
	return r.ApiService.DeleteDeviceExecute(r)
}

/*
DeleteDevice Method for DeleteDevice

Delete a device under a registry

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subscriptionId Subscription ID
 @param registryId Registry ID
 @param deviceId Device ID
 @return ApiDeleteDeviceRequest
*/
func (a *DeviceApiService) DeleteDevice(ctx context.Context, subscriptionId string, registryId string, deviceId string) ApiDeleteDeviceRequest {
	return ApiDeleteDeviceRequest{
		ApiService: a,
		ctx: ctx,
		subscriptionId: subscriptionId,
		registryId: registryId,
		deviceId: deviceId,
	}
}

// Execute executes the request
//  @return Info
func (a *DeviceApiService) DeleteDeviceExecute(r ApiDeleteDeviceRequest) (*Info, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Info
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceApiService.DeleteDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscriptions/{subscriptionId}/registries/{registryId}/devices/{deviceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", url.PathEscape(parameterValueToString(r.subscriptionId, "subscriptionId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"registryId"+"}", url.PathEscape(parameterValueToString(r.registryId, "registryId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"deviceId"+"}", url.PathEscape(parameterValueToString(r.deviceId, "deviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetConfigRequest struct {
	ctx context.Context
	ApiService *DeviceApiService
	subscriptionid string
	registryId string
	deviceId string
	numVersions *int32
}

// Device ID
func (r ApiGetConfigRequest) NumVersions(numVersions int32) ApiGetConfigRequest {
	r.numVersions = &numVersions
	return r
}

func (r ApiGetConfigRequest) Execute() (*ListDeviceConfigVersionsResponse, *http.Response, error) {
	return r.ApiService.GetConfigExecute(r)
}

/*
GetConfig Method for GetConfig

Get Configs Of Devices

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subscriptionid Subscription ID
 @param registryId Registry ID
 @param deviceId Device ID
 @return ApiGetConfigRequest
*/
func (a *DeviceApiService) GetConfig(ctx context.Context, subscriptionid string, registryId string, deviceId string) ApiGetConfigRequest {
	return ApiGetConfigRequest{
		ApiService: a,
		ctx: ctx,
		subscriptionid: subscriptionid,
		registryId: registryId,
		deviceId: deviceId,
	}
}

// Execute executes the request
//  @return ListDeviceConfigVersionsResponse
func (a *DeviceApiService) GetConfigExecute(r ApiGetConfigRequest) (*ListDeviceConfigVersionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListDeviceConfigVersionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceApiService.GetConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/configVersions"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionid"+"}", url.PathEscape(parameterValueToString(r.subscriptionid, "subscriptionid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"registryId"+"}", url.PathEscape(parameterValueToString(r.registryId, "registryId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"deviceId"+"}", url.PathEscape(parameterValueToString(r.deviceId, "deviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.numVersions == nil {
		return localVarReturnValue, nil, reportError("numVersions is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "numVersions", r.numVersions, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetDeviceRequest struct {
	ctx context.Context
	ApiService *DeviceApiService
	registryId string
	subscriptionId string
	deviceId string
}

func (r ApiGetDeviceRequest) Execute() (*Device, *http.Response, error) {
	return r.ApiService.GetDeviceExecute(r)
}

/*
GetDevice Method for GetDevice

Get a device under a registry

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param registryId Registry ID
 @param subscriptionId Subscription ID
 @param deviceId Device ID
 @return ApiGetDeviceRequest
*/
func (a *DeviceApiService) GetDevice(ctx context.Context, registryId string, subscriptionId string, deviceId string) ApiGetDeviceRequest {
	return ApiGetDeviceRequest{
		ApiService: a,
		ctx: ctx,
		registryId: registryId,
		subscriptionId: subscriptionId,
		deviceId: deviceId,
	}
}

// Execute executes the request
//  @return Device
func (a *DeviceApiService) GetDeviceExecute(r ApiGetDeviceRequest) (*Device, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Device
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceApiService.GetDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscriptions/{subscriptionId}/registries/{registryId}/devices/{deviceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"registryId"+"}", url.PathEscape(parameterValueToString(r.registryId, "registryId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", url.PathEscape(parameterValueToString(r.subscriptionId, "subscriptionId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"deviceId"+"}", url.PathEscape(parameterValueToString(r.deviceId, "deviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetDevicesRequest struct {
	ctx context.Context
	ApiService *DeviceApiService
	registryId string
	subscriptionId string
	pageNumber *int32
	pageSize *int32
	fieldMask *string
	deviceIds *[]string
	deviceNumIds *[]string
	gatewayListOptionsAssociationsDeviceId *string
	gatewayListOptionsAssociationsGatewayId *string
	gatewayListOptionsGatewayType *string
}

// Page Number
func (r ApiGetDevicesRequest) PageNumber(pageNumber int32) ApiGetDevicesRequest {
	r.pageNumber = &pageNumber
	return r
}

// The maximum number of devices to return in the response. If this value is zero, the service will select a default size. 
func (r ApiGetDevicesRequest) PageSize(pageSize int32) ApiGetDevicesRequest {
	r.pageSize = &pageSize
	return r
}

// The fields of the Device resource to be returned to the response. The fields id and numId are always returned, along with any other fields specified. A comma-separated list of fully qualified names of fields. Example: 
func (r ApiGetDevicesRequest) FieldMask(fieldMask string) ApiGetDevicesRequest {
	r.fieldMask = &fieldMask
	return r
}

// A list of device string IDs. For example, [&#39;device0&#39;, &#39;device12&#39;]. If empty, this field is ignored. Maximum IDs: 10,000
func (r ApiGetDevicesRequest) DeviceIds(deviceIds []string) ApiGetDevicesRequest {
	r.deviceIds = &deviceIds
	return r
}

// A list of device numeric IDs. If empty, this field is ignored. Maximum IDs: 10,000.
func (r ApiGetDevicesRequest) DeviceNumIds(deviceNumIds []string) ApiGetDevicesRequest {
	r.deviceNumIds = &deviceNumIds
	return r
}

// If set, returns only the gateways with which the specified device is associated. The device ID can be numeric (num_id) or the user-defined string (id). For example, if 456 is specified, returns only the gateways to which the device with num_id 456 is bound.
func (r ApiGetDevicesRequest) GatewayListOptionsAssociationsDeviceId(gatewayListOptionsAssociationsDeviceId string) ApiGetDevicesRequest {
	r.gatewayListOptionsAssociationsDeviceId = &gatewayListOptionsAssociationsDeviceId
	return r
}

// If set, only devices associated with the specified gateway are returned. The gateway ID can be numeric (num_id) or the user-defined string (id). For example, if 123 is specified, only devices bound to the gateway with num_id 123 are returned
func (r ApiGetDevicesRequest) GatewayListOptionsAssociationsGatewayId(gatewayListOptionsAssociationsGatewayId string) ApiGetDevicesRequest {
	r.gatewayListOptionsAssociationsGatewayId = &gatewayListOptionsAssociationsGatewayId
	return r
}

// If GATEWAY is specified, only gateways are returned. If NON_GATEWAY is specified, only non-gateway devices are returned. If GATEWAY_TYPE_UNSPECIFIED is specified, all devices are returned.
func (r ApiGetDevicesRequest) GatewayListOptionsGatewayType(gatewayListOptionsGatewayType string) ApiGetDevicesRequest {
	r.gatewayListOptionsGatewayType = &gatewayListOptionsGatewayType
	return r
}

func (r ApiGetDevicesRequest) Execute() (*ListDevicesResponse, *http.Response, error) {
	return r.ApiService.GetDevicesExecute(r)
}

/*
GetDevices Method for GetDevices

Get all devices under a registry

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param registryId Registry ID
 @param subscriptionId Subscription ID
 @return ApiGetDevicesRequest
*/
func (a *DeviceApiService) GetDevices(ctx context.Context, registryId string, subscriptionId string) ApiGetDevicesRequest {
	return ApiGetDevicesRequest{
		ApiService: a,
		ctx: ctx,
		registryId: registryId,
		subscriptionId: subscriptionId,
	}
}

// Execute executes the request
//  @return ListDevicesResponse
func (a *DeviceApiService) GetDevicesExecute(r ApiGetDevicesRequest) (*ListDevicesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListDevicesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceApiService.GetDevices")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscriptions/{subscriptionId}/registries/{registryId}/devices"
	localVarPath = strings.Replace(localVarPath, "{"+"registryId"+"}", url.PathEscape(parameterValueToString(r.registryId, "registryId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", url.PathEscape(parameterValueToString(r.subscriptionId, "subscriptionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNumber", r.pageNumber, "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "")
	}
	if r.fieldMask != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fieldMask", r.fieldMask, "")
	}
	if r.deviceIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "deviceIds", r.deviceIds, "csv")
	}
	if r.deviceNumIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "deviceNumIds", r.deviceNumIds, "csv")
	}
	if r.gatewayListOptionsAssociationsDeviceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "gatewayListOptions.associationsDeviceId", r.gatewayListOptionsAssociationsDeviceId, "")
	}
	if r.gatewayListOptionsAssociationsGatewayId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "gatewayListOptions.associationsGatewayId", r.gatewayListOptionsAssociationsGatewayId, "")
	}
	if r.gatewayListOptionsGatewayType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "gatewayListOptions.gatewayType", r.gatewayListOptionsGatewayType, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetStatesRequest struct {
	ctx context.Context
	ApiService *DeviceApiService
	subscriptionid string
	registryId string
	deviceId string
	numStates *int32
}

// The number of states to list. States are listed in descending order of update time. The maximum number of states retained is 10. If this value is zero, it will return all the states available.
func (r ApiGetStatesRequest) NumStates(numStates int32) ApiGetStatesRequest {
	r.numStates = &numStates
	return r
}

func (r ApiGetStatesRequest) Execute() (*ListDeviceStatesResponse, *http.Response, error) {
	return r.ApiService.GetStatesExecute(r)
}

/*
GetStates Method for GetStates

Get States Of Devices

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subscriptionid Subscription ID
 @param registryId Registry ID
 @param deviceId Device ID
 @return ApiGetStatesRequest
*/
func (a *DeviceApiService) GetStates(ctx context.Context, subscriptionid string, registryId string, deviceId string) ApiGetStatesRequest {
	return ApiGetStatesRequest{
		ApiService: a,
		ctx: ctx,
		subscriptionid: subscriptionid,
		registryId: registryId,
		deviceId: deviceId,
	}
}

// Execute executes the request
//  @return ListDeviceStatesResponse
func (a *DeviceApiService) GetStatesExecute(r ApiGetStatesRequest) (*ListDeviceStatesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListDeviceStatesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceApiService.GetStates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/states"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionid"+"}", url.PathEscape(parameterValueToString(r.subscriptionid, "subscriptionid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"registryId"+"}", url.PathEscape(parameterValueToString(r.registryId, "registryId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"deviceId"+"}", url.PathEscape(parameterValueToString(r.deviceId, "deviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.numStates != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "numStates", r.numStates, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSendCommandToDeviceRequest struct {
	ctx context.Context
	ApiService *DeviceApiService
	subscriptionid string
	registryId string
	deviceId string
	device *DeviceCommand
}

// application/json
func (r ApiSendCommandToDeviceRequest) Device(device DeviceCommand) ApiSendCommandToDeviceRequest {
	r.device = &device
	return r
}

func (r ApiSendCommandToDeviceRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.SendCommandToDeviceExecute(r)
}

/*
SendCommandToDevice Method for SendCommandToDevice

Send A Command To A Device

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subscriptionid Subscription ID
 @param registryId Registry ID
 @param deviceId Device ID
 @return ApiSendCommandToDeviceRequest
*/
func (a *DeviceApiService) SendCommandToDevice(ctx context.Context, subscriptionid string, registryId string, deviceId string) ApiSendCommandToDeviceRequest {
	return ApiSendCommandToDeviceRequest{
		ApiService: a,
		ctx: ctx,
		subscriptionid: subscriptionid,
		registryId: registryId,
		deviceId: deviceId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *DeviceApiService) SendCommandToDeviceExecute(r ApiSendCommandToDeviceRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceApiService.SendCommandToDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/sendCommandToDevice"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionid"+"}", url.PathEscape(parameterValueToString(r.subscriptionid, "subscriptionid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"registryId"+"}", url.PathEscape(parameterValueToString(r.registryId, "registryId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"deviceId"+"}", url.PathEscape(parameterValueToString(r.deviceId, "deviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.device == nil {
		return localVarReturnValue, nil, reportError("device is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.device
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSendConfigurationToDeviceRequest struct {
	ctx context.Context
	ApiService *DeviceApiService
	subscriptionid string
	registryId string
	deviceId string
	device *DeviceConfiguration
}

// application/json
func (r ApiSendConfigurationToDeviceRequest) Device(device DeviceConfiguration) ApiSendConfigurationToDeviceRequest {
	r.device = &device
	return r
}

func (r ApiSendConfigurationToDeviceRequest) Execute() (*DeviceConfig, *http.Response, error) {
	return r.ApiService.SendConfigurationToDeviceExecute(r)
}

/*
SendConfigurationToDevice Method for SendConfigurationToDevice

Send A Config To A Device

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subscriptionid Subscription ID
 @param registryId Registry ID
 @param deviceId Device ID
 @return ApiSendConfigurationToDeviceRequest
*/
func (a *DeviceApiService) SendConfigurationToDevice(ctx context.Context, subscriptionid string, registryId string, deviceId string) ApiSendConfigurationToDeviceRequest {
	return ApiSendConfigurationToDeviceRequest{
		ApiService: a,
		ctx: ctx,
		subscriptionid: subscriptionid,
		registryId: registryId,
		deviceId: deviceId,
	}
}

// Execute executes the request
//  @return DeviceConfig
func (a *DeviceApiService) SendConfigurationToDeviceExecute(r ApiSendConfigurationToDeviceRequest) (*DeviceConfig, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DeviceConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceApiService.SendConfigurationToDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/sendConfigurationToDevice"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionid"+"}", url.PathEscape(parameterValueToString(r.subscriptionid, "subscriptionid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"registryId"+"}", url.PathEscape(parameterValueToString(r.registryId, "registryId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"deviceId"+"}", url.PathEscape(parameterValueToString(r.deviceId, "deviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.device == nil {
		return localVarReturnValue, nil, reportError("device is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.device
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUnBindDeviceRequest struct {
	ctx context.Context
	ApiService *DeviceApiService
	subscriptionId string
	registryId string
	device *BindRequest
}

// application/json
func (r ApiUnBindDeviceRequest) Device(device BindRequest) ApiUnBindDeviceRequest {
	r.device = &device
	return r
}

func (r ApiUnBindDeviceRequest) Execute() (*Info, *http.Response, error) {
	return r.ApiService.UnBindDeviceExecute(r)
}

/*
UnBindDevice Method for UnBindDevice

UnBind  a device from a gateway under a registry

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subscriptionId Subscription ID
 @param registryId Registry ID
 @return ApiUnBindDeviceRequest
*/
func (a *DeviceApiService) UnBindDevice(ctx context.Context, subscriptionId string, registryId string) ApiUnBindDeviceRequest {
	return ApiUnBindDeviceRequest{
		ApiService: a,
		ctx: ctx,
		subscriptionId: subscriptionId,
		registryId: registryId,
	}
}

// Execute executes the request
//  @return Info
func (a *DeviceApiService) UnBindDeviceExecute(r ApiUnBindDeviceRequest) (*Info, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Info
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceApiService.UnBindDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscriptions/{subscriptionId}/registries/{registryId}/unbindDeviceFromGateway"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", url.PathEscape(parameterValueToString(r.subscriptionId, "subscriptionId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"registryId"+"}", url.PathEscape(parameterValueToString(r.registryId, "registryId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.device == nil {
		return localVarReturnValue, nil, reportError("device is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.device
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUnBindDevicesRequest struct {
	ctx context.Context
	ApiService *DeviceApiService
	subscriptionId string
	registryId string
	device *BindRequestIdsGateway
}

// application/json
func (r ApiUnBindDevicesRequest) Device(device BindRequestIdsGateway) ApiUnBindDevicesRequest {
	r.device = &device
	return r
}

func (r ApiUnBindDevicesRequest) Execute() (*Info, *http.Response, error) {
	return r.ApiService.UnBindDevicesExecute(r)
}

/*
UnBindDevices Method for UnBindDevices

UnBind devices from a gateway under a registry

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subscriptionId Subscription ID
 @param registryId Registry ID
 @return ApiUnBindDevicesRequest
*/
func (a *DeviceApiService) UnBindDevices(ctx context.Context, subscriptionId string, registryId string) ApiUnBindDevicesRequest {
	return ApiUnBindDevicesRequest{
		ApiService: a,
		ctx: ctx,
		subscriptionId: subscriptionId,
		registryId: registryId,
	}
}

// Execute executes the request
//  @return Info
func (a *DeviceApiService) UnBindDevicesExecute(r ApiUnBindDevicesRequest) (*Info, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Info
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceApiService.UnBindDevices")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscriptions/{subscriptionId}/registries/{registryId}/unbindDevicesFromGateway"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", url.PathEscape(parameterValueToString(r.subscriptionId, "subscriptionId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"registryId"+"}", url.PathEscape(parameterValueToString(r.registryId, "registryId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.device == nil {
		return localVarReturnValue, nil, reportError("device is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.device
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateDeviceRequest struct {
	ctx context.Context
	ApiService *DeviceApiService
	subscriptionId string
	registryId string
	deviceId string
	updateMask *string
	device *UpdateDevice
}

// Required. Only updates the device fields indicated by this mask. The field mask must not be empty, and it must not contain fields that are immutable or only set by the server. Mutable top-level fields: credentials,log_level, blocked, and metadata
func (r ApiUpdateDeviceRequest) UpdateMask(updateMask string) ApiUpdateDeviceRequest {
	r.updateMask = &updateMask
	return r
}

// application/json
func (r ApiUpdateDeviceRequest) Device(device UpdateDevice) ApiUpdateDeviceRequest {
	r.device = &device
	return r
}

func (r ApiUpdateDeviceRequest) Execute() (*Device, *http.Response, error) {
	return r.ApiService.UpdateDeviceExecute(r)
}

/*
UpdateDevice Method for UpdateDevice

Modify device under a registry

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subscriptionId Subscription ID
 @param registryId Registry ID
 @param deviceId Device ID
 @return ApiUpdateDeviceRequest
*/
func (a *DeviceApiService) UpdateDevice(ctx context.Context, subscriptionId string, registryId string, deviceId string) ApiUpdateDeviceRequest {
	return ApiUpdateDeviceRequest{
		ApiService: a,
		ctx: ctx,
		subscriptionId: subscriptionId,
		registryId: registryId,
		deviceId: deviceId,
	}
}

// Execute executes the request
//  @return Device
func (a *DeviceApiService) UpdateDeviceExecute(r ApiUpdateDeviceRequest) (*Device, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Device
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceApiService.UpdateDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscriptions/{subscriptionId}/registries/{registryId}/devices/{deviceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", url.PathEscape(parameterValueToString(r.subscriptionId, "subscriptionId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"registryId"+"}", url.PathEscape(parameterValueToString(r.registryId, "registryId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"deviceId"+"}", url.PathEscape(parameterValueToString(r.deviceId, "deviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateMask == nil {
		return localVarReturnValue, nil, reportError("updateMask is required and must be specified")
	}
	if r.device == nil {
		return localVarReturnValue, nil, reportError("device is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "updateMask", r.updateMask, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.device
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GenericErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
