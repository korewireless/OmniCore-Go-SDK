/*
OmniCore Model and State Management API

Testing RegistryApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package OmniCore

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/korewireless/OmniCore-Go-SDK"
)

func Test_OmniCore_RegistryApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RegistryApiService CreateRegistry", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId string

		resp, httpRes, err := apiClient.RegistryApi.CreateRegistry(context.Background(), subscriptionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RegistryApiService DeleteRegistry", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId string
		var registryId string

		resp, httpRes, err := apiClient.RegistryApi.DeleteRegistry(context.Background(), subscriptionId, registryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RegistryApiService GetRegistries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId string

		resp, httpRes, err := apiClient.RegistryApi.GetRegistries(context.Background(), subscriptionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RegistryApiService GetRegistry", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId string
		var registryId string

		resp, httpRes, err := apiClient.RegistryApi.GetRegistry(context.Background(), subscriptionId, registryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RegistryApiService SendBroadcastToDevices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionid string
		var registryId string

		resp, httpRes, err := apiClient.RegistryApi.SendBroadcastToDevices(context.Background(), subscriptionid, registryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RegistryApiService UpdateRegistry", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId string
		var registryId string

		resp, httpRes, err := apiClient.RegistryApi.UpdateRegistry(context.Background(), subscriptionId, registryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
