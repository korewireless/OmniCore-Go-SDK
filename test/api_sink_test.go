/*
OmniCore Model and State Management API

Testing SinkApiService

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

func Test_OmniCore_SinkApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SinkApiService CreateSink", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId string

		resp, httpRes, err := apiClient.SinkApi.CreateSink(context.Background(), subscriptionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SinkApiService DeleteSink", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId string
		var sinkId string

		resp, httpRes, err := apiClient.SinkApi.DeleteSink(context.Background(), subscriptionId, sinkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SinkApiService GetSink", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId string
		var sinkId string

		resp, httpRes, err := apiClient.SinkApi.GetSink(context.Background(), subscriptionId, sinkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SinkApiService GetSinks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId string

		resp, httpRes, err := apiClient.SinkApi.GetSinks(context.Background(), subscriptionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
