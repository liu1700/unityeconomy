/*
Economy API

Testing PurchasesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package unityeconomy

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/liu1700/unityeconomy"
)

func Test_unityeconomy_PurchasesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PurchasesAPIService MakeVirtualPurchase", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string
		var playerId string

		resp, httpRes, err := apiClient.PurchasesAPI.MakeVirtualPurchase(context.Background(), projectId, playerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PurchasesAPIService RedeemAppleAppStorePurchase", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string
		var playerId string

		resp, httpRes, err := apiClient.PurchasesAPI.RedeemAppleAppStorePurchase(context.Background(), projectId, playerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PurchasesAPIService RedeemGooglePlayPurchase", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string
		var playerId string

		resp, httpRes, err := apiClient.PurchasesAPI.RedeemGooglePlayPurchase(context.Background(), projectId, playerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}