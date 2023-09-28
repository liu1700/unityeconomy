# \PurchasesAPI

All URIs are relative to *https://economy.services.api.unity.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MakeVirtualPurchase**](PurchasesAPI.md#MakeVirtualPurchase) | **Post** /v2/projects/{projectId}/players/{playerId}/purchases/virtual | Make virtual purchase
[**RedeemAppleAppStorePurchase**](PurchasesAPI.md#RedeemAppleAppStorePurchase) | **Post** /v2/projects/{projectId}/players/{playerId}/purchases/appleappstore | Redeem Apple App Store purchase
[**RedeemGooglePlayPurchase**](PurchasesAPI.md#RedeemGooglePlayPurchase) | **Post** /v2/projects/{projectId}/players/{playerId}/purchases/googleplaystore | Redeem Google Play purchase



## MakeVirtualPurchase

> PlayerPurchaseVirtualResponse MakeVirtualPurchase(ctx, projectId, playerId).PlayerPurchaseVirtualRequest(playerPurchaseVirtualRequest).ConfigAssignmentHash(configAssignmentHash).UnityInstallationId(unityInstallationId).AnalyticsUserId(analyticsUserId).Execute()

Make virtual purchase



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/liu1700/unityeconomy"
)

func main() {
    projectId := "projectId_example" // string | ID of the project.
    playerId := "playerId_example" // string | ID of the player.
    playerPurchaseVirtualRequest := *openapiclient.NewPlayerPurchaseVirtualRequest("Id_example") // PlayerPurchaseVirtualRequest | 
    configAssignmentHash := "configAssignmentHash_example" // string | Configuration assignment hash that should be used when processing this request. (optional)
    unityInstallationId := "unityInstallationId_example" // string | Unique identifier that identifies an installation on the client’s device. The same player can have different installationIds if they have the game installed on different devices. It is available to all Unity packages that integrate with the Services SDK Core package. (optional)
    analyticsUserId := "analyticsUserId_example" // string | A unique string that identifies the player and is consistent across their subsequent play sessions for analytics purposes. It is the primary user identifier and it comes from the Core package. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PurchasesAPI.MakeVirtualPurchase(context.Background(), projectId, playerId).PlayerPurchaseVirtualRequest(playerPurchaseVirtualRequest).ConfigAssignmentHash(configAssignmentHash).UnityInstallationId(unityInstallationId).AnalyticsUserId(analyticsUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurchasesAPI.MakeVirtualPurchase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MakeVirtualPurchase`: PlayerPurchaseVirtualResponse
    fmt.Fprintf(os.Stdout, "Response from `PurchasesAPI.MakeVirtualPurchase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project. | 
**playerId** | **string** | ID of the player. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMakeVirtualPurchaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **playerPurchaseVirtualRequest** | [**PlayerPurchaseVirtualRequest**](PlayerPurchaseVirtualRequest.md) |  | 
 **configAssignmentHash** | **string** | Configuration assignment hash that should be used when processing this request. | 
 **unityInstallationId** | **string** | Unique identifier that identifies an installation on the client’s device. The same player can have different installationIds if they have the game installed on different devices. It is available to all Unity packages that integrate with the Services SDK Core package. | 
 **analyticsUserId** | **string** | A unique string that identifies the player and is consistent across their subsequent play sessions for analytics purposes. It is the primary user identifier and it comes from the Core package. | 

### Return type

[**PlayerPurchaseVirtualResponse**](PlayerPurchaseVirtualResponse.md)

### Authorization

[Client](../README.md#Client)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RedeemAppleAppStorePurchase

> PlayerPurchaseAppleappstoreResponse RedeemAppleAppStorePurchase(ctx, projectId, playerId).PlayerPurchaseAppleappstoreRequest(playerPurchaseAppleappstoreRequest).ConfigAssignmentHash(configAssignmentHash).UnityInstallationId(unityInstallationId).AnalyticsUserId(analyticsUserId).Execute()

Redeem Apple App Store purchase



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/liu1700/unityeconomy"
)

func main() {
    projectId := "projectId_example" // string | ID of the project.
    playerId := "playerId_example" // string | ID of the player.
    playerPurchaseAppleappstoreRequest := *openapiclient.NewPlayerPurchaseAppleappstoreRequest("Id_example", "Receipt_example", int32(123), "USD") // PlayerPurchaseAppleappstoreRequest | 
    configAssignmentHash := "configAssignmentHash_example" // string | Configuration assignment hash that should be used when processing this request. (optional)
    unityInstallationId := "unityInstallationId_example" // string | Unique identifier that identifies an installation on the client’s device. The same player can have different installationIds if they have the game installed on different devices. It is available to all Unity packages that integrate with the Services SDK Core package. (optional)
    analyticsUserId := "analyticsUserId_example" // string | A unique string that identifies the player and is consistent across their subsequent play sessions for analytics purposes. It is the primary user identifier and it comes from the Core package. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PurchasesAPI.RedeemAppleAppStorePurchase(context.Background(), projectId, playerId).PlayerPurchaseAppleappstoreRequest(playerPurchaseAppleappstoreRequest).ConfigAssignmentHash(configAssignmentHash).UnityInstallationId(unityInstallationId).AnalyticsUserId(analyticsUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurchasesAPI.RedeemAppleAppStorePurchase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RedeemAppleAppStorePurchase`: PlayerPurchaseAppleappstoreResponse
    fmt.Fprintf(os.Stdout, "Response from `PurchasesAPI.RedeemAppleAppStorePurchase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project. | 
**playerId** | **string** | ID of the player. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRedeemAppleAppStorePurchaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **playerPurchaseAppleappstoreRequest** | [**PlayerPurchaseAppleappstoreRequest**](PlayerPurchaseAppleappstoreRequest.md) |  | 
 **configAssignmentHash** | **string** | Configuration assignment hash that should be used when processing this request. | 
 **unityInstallationId** | **string** | Unique identifier that identifies an installation on the client’s device. The same player can have different installationIds if they have the game installed on different devices. It is available to all Unity packages that integrate with the Services SDK Core package. | 
 **analyticsUserId** | **string** | A unique string that identifies the player and is consistent across their subsequent play sessions for analytics purposes. It is the primary user identifier and it comes from the Core package. | 

### Return type

[**PlayerPurchaseAppleappstoreResponse**](PlayerPurchaseAppleappstoreResponse.md)

### Authorization

[Client](../README.md#Client)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RedeemGooglePlayPurchase

> PlayerPurchaseGoogleplaystoreResponse RedeemGooglePlayPurchase(ctx, projectId, playerId).PlayerPurchaseGoogleplaystoreRequest(playerPurchaseGoogleplaystoreRequest).ConfigAssignmentHash(configAssignmentHash).UnityInstallationId(unityInstallationId).AnalyticsUserId(analyticsUserId).Execute()

Redeem Google Play purchase



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/liu1700/unityeconomy"
)

func main() {
    projectId := "projectId_example" // string | ID of the project.
    playerId := "playerId_example" // string | ID of the player.
    playerPurchaseGoogleplaystoreRequest := *openapiclient.NewPlayerPurchaseGoogleplaystoreRequest("Id_example", "PurchaseData_example", "PurchaseDataSignature_example", int32(123), "USD") // PlayerPurchaseGoogleplaystoreRequest | 
    configAssignmentHash := "configAssignmentHash_example" // string | Configuration assignment hash that should be used when processing this request. (optional)
    unityInstallationId := "unityInstallationId_example" // string | Unique identifier that identifies an installation on the client’s device. The same player can have different installationIds if they have the game installed on different devices. It is available to all Unity packages that integrate with the Services SDK Core package. (optional)
    analyticsUserId := "analyticsUserId_example" // string | A unique string that identifies the player and is consistent across their subsequent play sessions for analytics purposes. It is the primary user identifier and it comes from the Core package. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PurchasesAPI.RedeemGooglePlayPurchase(context.Background(), projectId, playerId).PlayerPurchaseGoogleplaystoreRequest(playerPurchaseGoogleplaystoreRequest).ConfigAssignmentHash(configAssignmentHash).UnityInstallationId(unityInstallationId).AnalyticsUserId(analyticsUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurchasesAPI.RedeemGooglePlayPurchase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RedeemGooglePlayPurchase`: PlayerPurchaseGoogleplaystoreResponse
    fmt.Fprintf(os.Stdout, "Response from `PurchasesAPI.RedeemGooglePlayPurchase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project. | 
**playerId** | **string** | ID of the player. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRedeemGooglePlayPurchaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **playerPurchaseGoogleplaystoreRequest** | [**PlayerPurchaseGoogleplaystoreRequest**](PlayerPurchaseGoogleplaystoreRequest.md) |  | 
 **configAssignmentHash** | **string** | Configuration assignment hash that should be used when processing this request. | 
 **unityInstallationId** | **string** | Unique identifier that identifies an installation on the client’s device. The same player can have different installationIds if they have the game installed on different devices. It is available to all Unity packages that integrate with the Services SDK Core package. | 
 **analyticsUserId** | **string** | A unique string that identifies the player and is consistent across their subsequent play sessions for analytics purposes. It is the primary user identifier and it comes from the Core package. | 

### Return type

[**PlayerPurchaseGoogleplaystoreResponse**](PlayerPurchaseGoogleplaystoreResponse.md)

### Authorization

[Client](../README.md#Client)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

