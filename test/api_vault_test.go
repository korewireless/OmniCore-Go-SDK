/*
OmniCore Model and State Management API

Testing VaultApiService

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

func Test_OmniCore_VaultApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test VaultApiService CreateVaultConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionid string

		resp, httpRes, err := apiClient.VaultApi.CreateVaultConfiguration(context.Background(), subscriptionid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VaultApiService CreateVaultKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionid string

		resp, httpRes, err := apiClient.VaultApi.CreateVaultKey(context.Background(), subscriptionid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VaultApiService DeleteConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionid string
		var configid string

		resp, httpRes, err := apiClient.VaultApi.DeleteConfiguration(context.Background(), subscriptionid, configid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VaultApiService DeleteVaultKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionid string
		var keyid string

		resp, httpRes, err := apiClient.VaultApi.DeleteVaultKey(context.Background(), subscriptionid, keyid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VaultApiService EnableEncryption", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionid string

		resp, httpRes, err := apiClient.VaultApi.EnableEncryption(context.Background(), subscriptionid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VaultApiService GetExports", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionid string

		resp, httpRes, err := apiClient.VaultApi.GetExports(context.Background(), subscriptionid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VaultApiService GetRegistryData", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionid string

		resp, httpRes, err := apiClient.VaultApi.GetRegistryData(context.Background(), subscriptionid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VaultApiService GetReplays", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionid string

		resp, httpRes, err := apiClient.VaultApi.GetReplays(context.Background(), subscriptionid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VaultApiService GetVaultAudit", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionid string

		resp, httpRes, err := apiClient.VaultApi.GetVaultAudit(context.Background(), subscriptionid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VaultApiService GetVaultConfigurations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionid string

		resp, httpRes, err := apiClient.VaultApi.GetVaultConfigurations(context.Background(), subscriptionid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VaultApiService GetVaultFiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionid string
		var registryid string

		resp, httpRes, err := apiClient.VaultApi.GetVaultFiles(context.Background(), subscriptionid, registryid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VaultApiService GetVaultKeys", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionid string

		resp, httpRes, err := apiClient.VaultApi.GetVaultKeys(context.Background(), subscriptionid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VaultApiService GetVaultMetrics", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionid string

		resp, httpRes, err := apiClient.VaultApi.GetVaultMetrics(context.Background(), subscriptionid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VaultApiService GetVaultStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionid string

		resp, httpRes, err := apiClient.VaultApi.GetVaultStatus(context.Background(), subscriptionid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VaultApiService SetRetention", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionid string

		resp, httpRes, err := apiClient.VaultApi.SetRetention(context.Background(), subscriptionid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VaultApiService StartExport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionid string

		resp, httpRes, err := apiClient.VaultApi.StartExport(context.Background(), subscriptionid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VaultApiService StartReplay", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionid string

		resp, httpRes, err := apiClient.VaultApi.StartReplay(context.Background(), subscriptionid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
