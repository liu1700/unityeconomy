# \ConfigurationAPI

All URIs are relative to *https://economy.services.api.unity.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPlayerConfiguration**](ConfigurationAPI.md#GetPlayerConfiguration) | **Get** /v2/projects/{projectId}/players/{playerId}/config/resources | Get player&#39;s configuration



## GetPlayerConfiguration

> PlayerConfigurationResponse GetPlayerConfiguration(ctx, projectId, playerId).ConfigAssignmentHash(configAssignmentHash).UnityInstallationId(unityInstallationId).AnalyticsUserId(analyticsUserId).Execute()

Get player's configuration



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationAPI.GetPlayerConfiguration(context.Background(), projectId, playerId).ConfigAssignmentHash(configAssignmentHash).UnityInstallationId(unityInstallationId).AnalyticsUserId(analyticsUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GetPlayerConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlayerConfiguration`: PlayerConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GetPlayerConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project. | 
**playerId** | **string** | ID of the player. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayerConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **configAssignmentHash** | **string** | Configuration assignment hash that should be used when processing this request. | 
 **unityInstallationId** | **string** | Unique identifier that identifies an installation on the client’s device. The same player can have different installationIds if they have the game installed on different devices. It is available to all Unity packages that integrate with the Services SDK Core package. | 
 **analyticsUserId** | **string** | A unique string that identifies the player and is consistent across their subsequent play sessions for analytics purposes. It is the primary user identifier and it comes from the Core package. | 

### Return type

[**PlayerConfigurationResponse**](PlayerConfigurationResponse.md)

### Authorization

[Client](../README.md#Client)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

