# \CurrenciesAPI

All URIs are relative to *https://economy.services.api.unity.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DecrementPlayerCurrencyBalance**](CurrenciesAPI.md#DecrementPlayerCurrencyBalance) | **Post** /v2/projects/{projectId}/players/{playerId}/currencies/{currencyId}/decrement | Decrement currency balance
[**GetPlayerCurrencies**](CurrenciesAPI.md#GetPlayerCurrencies) | **Get** /v2/projects/{projectId}/players/{playerId}/currencies | Player currency balances
[**IncrementPlayerCurrencyBalance**](CurrenciesAPI.md#IncrementPlayerCurrencyBalance) | **Post** /v2/projects/{projectId}/players/{playerId}/currencies/{currencyId}/increment | Increment currency balance
[**SetPlayerCurrencyBalance**](CurrenciesAPI.md#SetPlayerCurrencyBalance) | **Put** /v2/projects/{projectId}/players/{playerId}/currencies/{currencyId} | Set currency balance



## DecrementPlayerCurrencyBalance

> CurrencyBalanceResponse DecrementPlayerCurrencyBalance(ctx, projectId, playerId, currencyId).CurrencyModifyBalanceRequest(currencyModifyBalanceRequest).ConfigAssignmentHash(configAssignmentHash).UnityInstallationId(unityInstallationId).AnalyticsUserId(analyticsUserId).Execute()

Decrement currency balance



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
    currencyId := "currencyId_example" // string | Resource ID of the currency.
    currencyModifyBalanceRequest := *openapiclient.NewCurrencyModifyBalanceRequest(int64(123)) // CurrencyModifyBalanceRequest | 
    configAssignmentHash := "configAssignmentHash_example" // string | Configuration assignment hash that should be used when processing this request. (optional)
    unityInstallationId := "unityInstallationId_example" // string | Unique identifier that identifies an installation on the client’s device. The same player can have different installationIds if they have the game installed on different devices. It is available to all Unity packages that integrate with the Services SDK Core package. (optional)
    analyticsUserId := "analyticsUserId_example" // string | A unique string that identifies the player and is consistent across their subsequent play sessions for analytics purposes. It is the primary user identifier and it comes from the Core package. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CurrenciesAPI.DecrementPlayerCurrencyBalance(context.Background(), projectId, playerId, currencyId).CurrencyModifyBalanceRequest(currencyModifyBalanceRequest).ConfigAssignmentHash(configAssignmentHash).UnityInstallationId(unityInstallationId).AnalyticsUserId(analyticsUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesAPI.DecrementPlayerCurrencyBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DecrementPlayerCurrencyBalance`: CurrencyBalanceResponse
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesAPI.DecrementPlayerCurrencyBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project. | 
**playerId** | **string** | ID of the player. | 
**currencyId** | **string** | Resource ID of the currency. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDecrementPlayerCurrencyBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **currencyModifyBalanceRequest** | [**CurrencyModifyBalanceRequest**](CurrencyModifyBalanceRequest.md) |  | 
 **configAssignmentHash** | **string** | Configuration assignment hash that should be used when processing this request. | 
 **unityInstallationId** | **string** | Unique identifier that identifies an installation on the client’s device. The same player can have different installationIds if they have the game installed on different devices. It is available to all Unity packages that integrate with the Services SDK Core package. | 
 **analyticsUserId** | **string** | A unique string that identifies the player and is consistent across their subsequent play sessions for analytics purposes. It is the primary user identifier and it comes from the Core package. | 

### Return type

[**CurrencyBalanceResponse**](CurrencyBalanceResponse.md)

### Authorization

[Client](../README.md#Client)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayerCurrencies

> PlayerCurrencyBalanceResponse GetPlayerCurrencies(ctx, projectId, playerId).ConfigAssignmentHash(configAssignmentHash).UnityInstallationId(unityInstallationId).AnalyticsUserId(analyticsUserId).After(after).Limit(limit).Execute()

Player currency balances



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
    configAssignmentHash := "configAssignmentHash_example" // string | Configuration assignment hash that should be used when processing this request. (optional)
    unityInstallationId := "unityInstallationId_example" // string | Unique identifier that identifies an installation on the client’s device. The same player can have different installationIds if they have the game installed on different devices. It is available to all Unity packages that integrate with the Services SDK Core package. (optional)
    analyticsUserId := "analyticsUserId_example" // string | A unique string that identifies the player and is consistent across their subsequent play sessions for analytics purposes. It is the primary user identifier and it comes from the Core package. (optional)
    after := "after_example" // string | The currency ID after which to retrieve the next page of balances. (optional)
    limit := int32(56) // int32 | Number of currencies to be returned. Defaults to 20. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CurrenciesAPI.GetPlayerCurrencies(context.Background(), projectId, playerId).ConfigAssignmentHash(configAssignmentHash).UnityInstallationId(unityInstallationId).AnalyticsUserId(analyticsUserId).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesAPI.GetPlayerCurrencies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlayerCurrencies`: PlayerCurrencyBalanceResponse
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesAPI.GetPlayerCurrencies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project. | 
**playerId** | **string** | ID of the player. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayerCurrenciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **configAssignmentHash** | **string** | Configuration assignment hash that should be used when processing this request. | 
 **unityInstallationId** | **string** | Unique identifier that identifies an installation on the client’s device. The same player can have different installationIds if they have the game installed on different devices. It is available to all Unity packages that integrate with the Services SDK Core package. | 
 **analyticsUserId** | **string** | A unique string that identifies the player and is consistent across their subsequent play sessions for analytics purposes. It is the primary user identifier and it comes from the Core package. | 
 **after** | **string** | The currency ID after which to retrieve the next page of balances. | 
 **limit** | **int32** | Number of currencies to be returned. Defaults to 20. | 

### Return type

[**PlayerCurrencyBalanceResponse**](PlayerCurrencyBalanceResponse.md)

### Authorization

[Client](../README.md#Client)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IncrementPlayerCurrencyBalance

> CurrencyBalanceResponse IncrementPlayerCurrencyBalance(ctx, projectId, playerId, currencyId).CurrencyModifyBalanceRequest(currencyModifyBalanceRequest).ConfigAssignmentHash(configAssignmentHash).UnityInstallationId(unityInstallationId).AnalyticsUserId(analyticsUserId).Execute()

Increment currency balance



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
    currencyId := "currencyId_example" // string | Resource ID of the currency.
    currencyModifyBalanceRequest := *openapiclient.NewCurrencyModifyBalanceRequest(int64(123)) // CurrencyModifyBalanceRequest | 
    configAssignmentHash := "configAssignmentHash_example" // string | Configuration assignment hash that should be used when processing this request. (optional)
    unityInstallationId := "unityInstallationId_example" // string | Unique identifier that identifies an installation on the client’s device. The same player can have different installationIds if they have the game installed on different devices. It is available to all Unity packages that integrate with the Services SDK Core package. (optional)
    analyticsUserId := "analyticsUserId_example" // string | A unique string that identifies the player and is consistent across their subsequent play sessions for analytics purposes. It is the primary user identifier and it comes from the Core package. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CurrenciesAPI.IncrementPlayerCurrencyBalance(context.Background(), projectId, playerId, currencyId).CurrencyModifyBalanceRequest(currencyModifyBalanceRequest).ConfigAssignmentHash(configAssignmentHash).UnityInstallationId(unityInstallationId).AnalyticsUserId(analyticsUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesAPI.IncrementPlayerCurrencyBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IncrementPlayerCurrencyBalance`: CurrencyBalanceResponse
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesAPI.IncrementPlayerCurrencyBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project. | 
**playerId** | **string** | ID of the player. | 
**currencyId** | **string** | Resource ID of the currency. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIncrementPlayerCurrencyBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **currencyModifyBalanceRequest** | [**CurrencyModifyBalanceRequest**](CurrencyModifyBalanceRequest.md) |  | 
 **configAssignmentHash** | **string** | Configuration assignment hash that should be used when processing this request. | 
 **unityInstallationId** | **string** | Unique identifier that identifies an installation on the client’s device. The same player can have different installationIds if they have the game installed on different devices. It is available to all Unity packages that integrate with the Services SDK Core package. | 
 **analyticsUserId** | **string** | A unique string that identifies the player and is consistent across their subsequent play sessions for analytics purposes. It is the primary user identifier and it comes from the Core package. | 

### Return type

[**CurrencyBalanceResponse**](CurrencyBalanceResponse.md)

### Authorization

[Client](../README.md#Client)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPlayerCurrencyBalance

> CurrencyBalanceResponse SetPlayerCurrencyBalance(ctx, projectId, playerId, currencyId).CurrencyBalanceRequest(currencyBalanceRequest).ConfigAssignmentHash(configAssignmentHash).UnityInstallationId(unityInstallationId).AnalyticsUserId(analyticsUserId).Execute()

Set currency balance



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
    currencyId := "currencyId_example" // string | Resource ID of the currency.
    currencyBalanceRequest := *openapiclient.NewCurrencyBalanceRequest(int64(123)) // CurrencyBalanceRequest | 
    configAssignmentHash := "configAssignmentHash_example" // string | Configuration assignment hash that should be used when processing this request. (optional)
    unityInstallationId := "unityInstallationId_example" // string | Unique identifier that identifies an installation on the client’s device. The same player can have different installationIds if they have the game installed on different devices. It is available to all Unity packages that integrate with the Services SDK Core package. (optional)
    analyticsUserId := "analyticsUserId_example" // string | A unique string that identifies the player and is consistent across their subsequent play sessions for analytics purposes. It is the primary user identifier and it comes from the Core package. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CurrenciesAPI.SetPlayerCurrencyBalance(context.Background(), projectId, playerId, currencyId).CurrencyBalanceRequest(currencyBalanceRequest).ConfigAssignmentHash(configAssignmentHash).UnityInstallationId(unityInstallationId).AnalyticsUserId(analyticsUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesAPI.SetPlayerCurrencyBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetPlayerCurrencyBalance`: CurrencyBalanceResponse
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesAPI.SetPlayerCurrencyBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project. | 
**playerId** | **string** | ID of the player. | 
**currencyId** | **string** | Resource ID of the currency. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPlayerCurrencyBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **currencyBalanceRequest** | [**CurrencyBalanceRequest**](CurrencyBalanceRequest.md) |  | 
 **configAssignmentHash** | **string** | Configuration assignment hash that should be used when processing this request. | 
 **unityInstallationId** | **string** | Unique identifier that identifies an installation on the client’s device. The same player can have different installationIds if they have the game installed on different devices. It is available to all Unity packages that integrate with the Services SDK Core package. | 
 **analyticsUserId** | **string** | A unique string that identifies the player and is consistent across their subsequent play sessions for analytics purposes. It is the primary user identifier and it comes from the Core package. | 

### Return type

[**CurrencyBalanceResponse**](CurrencyBalanceResponse.md)

### Authorization

[Client](../README.md#Client)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

