# \InventoryAPI

All URIs are relative to *https://economy.services.api.unity.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddInventoryItem**](InventoryAPI.md#AddInventoryItem) | **Post** /v2/projects/{projectId}/players/{playerId}/inventory | Add inventory item
[**DeleteInventoryItem**](InventoryAPI.md#DeleteInventoryItem) | **Delete** /v2/projects/{projectId}/players/{playerId}/inventory/{playersInventoryItemId} | Delete player&#39;s inventory item
[**GetPlayerInventory**](InventoryAPI.md#GetPlayerInventory) | **Get** /v2/projects/{projectId}/players/{playerId}/inventory | List player inventory
[**UpdateInventoryItem**](InventoryAPI.md#UpdateInventoryItem) | **Put** /v2/projects/{projectId}/players/{playerId}/inventory/{playersInventoryItemId} | Update player&#39;s inventory item



## AddInventoryItem

> InventoryResponse AddInventoryItem(ctx, projectId, playerId).AddInventoryRequest(addInventoryRequest).ConfigAssignmentHash(configAssignmentHash).UnityInstallationId(unityInstallationId).AnalyticsUserId(analyticsUserId).Execute()

Add inventory item



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
    addInventoryRequest := *openapiclient.NewAddInventoryRequest("InventoryItemId_example") // AddInventoryRequest | 
    configAssignmentHash := "configAssignmentHash_example" // string | Configuration assignment hash that should be used when processing this request. (optional)
    unityInstallationId := "unityInstallationId_example" // string | Unique identifier that identifies an installation on the client’s device. The same player can have different installationIds if they have the game installed on different devices. It is available to all Unity packages that integrate with the Services SDK Core package. (optional)
    analyticsUserId := "analyticsUserId_example" // string | A unique string that identifies the player and is consistent across their subsequent play sessions for analytics purposes. It is the primary user identifier and it comes from the Core package. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryAPI.AddInventoryItem(context.Background(), projectId, playerId).AddInventoryRequest(addInventoryRequest).ConfigAssignmentHash(configAssignmentHash).UnityInstallationId(unityInstallationId).AnalyticsUserId(analyticsUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.AddInventoryItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddInventoryItem`: InventoryResponse
    fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.AddInventoryItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project. | 
**playerId** | **string** | ID of the player. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddInventoryItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addInventoryRequest** | [**AddInventoryRequest**](AddInventoryRequest.md) |  | 
 **configAssignmentHash** | **string** | Configuration assignment hash that should be used when processing this request. | 
 **unityInstallationId** | **string** | Unique identifier that identifies an installation on the client’s device. The same player can have different installationIds if they have the game installed on different devices. It is available to all Unity packages that integrate with the Services SDK Core package. | 
 **analyticsUserId** | **string** | A unique string that identifies the player and is consistent across their subsequent play sessions for analytics purposes. It is the primary user identifier and it comes from the Core package. | 

### Return type

[**InventoryResponse**](InventoryResponse.md)

### Authorization

[Client](../README.md#Client)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInventoryItem

> DeleteInventoryItem(ctx, projectId, playerId, playersInventoryItemId).ConfigAssignmentHash(configAssignmentHash).UnityInstallationId(unityInstallationId).AnalyticsUserId(analyticsUserId).WriteLock(writeLock).InventoryDeleteRequest(inventoryDeleteRequest).Execute()

Delete player's inventory item



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
    playersInventoryItemId := "playersInventoryItemId_example" // string | The `playersInventoryItemId` of the item to be updated.
    configAssignmentHash := "configAssignmentHash_example" // string | Configuration assignment hash that should be used when processing this request. (optional)
    unityInstallationId := "unityInstallationId_example" // string | Unique identifier that identifies an installation on the client’s device. The same player can have different installationIds if they have the game installed on different devices. It is available to all Unity packages that integrate with the Services SDK Core package. (optional)
    analyticsUserId := "analyticsUserId_example" // string | A unique string that identifies the player and is consistent across their subsequent play sessions for analytics purposes. It is the primary user identifier and it comes from the Core package. (optional)
    writeLock := "writeLock_example" // string | The writelock for a database entry. (optional)
    inventoryDeleteRequest := *openapiclient.NewInventoryDeleteRequest() // InventoryDeleteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.InventoryAPI.DeleteInventoryItem(context.Background(), projectId, playerId, playersInventoryItemId).ConfigAssignmentHash(configAssignmentHash).UnityInstallationId(unityInstallationId).AnalyticsUserId(analyticsUserId).WriteLock(writeLock).InventoryDeleteRequest(inventoryDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.DeleteInventoryItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project. | 
**playerId** | **string** | ID of the player. | 
**playersInventoryItemId** | **string** | The &#x60;playersInventoryItemId&#x60; of the item to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInventoryItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **configAssignmentHash** | **string** | Configuration assignment hash that should be used when processing this request. | 
 **unityInstallationId** | **string** | Unique identifier that identifies an installation on the client’s device. The same player can have different installationIds if they have the game installed on different devices. It is available to all Unity packages that integrate with the Services SDK Core package. | 
 **analyticsUserId** | **string** | A unique string that identifies the player and is consistent across their subsequent play sessions for analytics purposes. It is the primary user identifier and it comes from the Core package. | 
 **writeLock** | **string** | The writelock for a database entry. | 
 **inventoryDeleteRequest** | [**InventoryDeleteRequest**](InventoryDeleteRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[Client](../README.md#Client)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayerInventory

> PlayerInventoryResponse GetPlayerInventory(ctx, projectId, playerId).ConfigAssignmentHash(configAssignmentHash).UnityInstallationId(unityInstallationId).AnalyticsUserId(analyticsUserId).After(after).Limit(limit).PlayersInventoryItemIds(playersInventoryItemIds).InventoryItemIds(inventoryItemIds).Execute()

List player inventory



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
    after := "after_example" // string | The `playersInventoryItemId` after which to retrieve the next page of balances. (optional)
    limit := int32(56) // int32 | Number of items to be returned. Defaults to 20. (optional)
    playersInventoryItemIds := []string{"Inner_example"} // []string | List of `playersInventoryItemIds` in array notation, for example, `playersInventoryItemIds[]=ID1&playersInventoryItemIds[]=ID2`. (optional)
    inventoryItemIds := []string{"Inner_example"} // []string | List of inventory item IDs in array notation, for example, `inventoryItemIds[]=ID1&inventoryItemIds[]=ID2`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryAPI.GetPlayerInventory(context.Background(), projectId, playerId).ConfigAssignmentHash(configAssignmentHash).UnityInstallationId(unityInstallationId).AnalyticsUserId(analyticsUserId).After(after).Limit(limit).PlayersInventoryItemIds(playersInventoryItemIds).InventoryItemIds(inventoryItemIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.GetPlayerInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlayerInventory`: PlayerInventoryResponse
    fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.GetPlayerInventory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project. | 
**playerId** | **string** | ID of the player. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayerInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **configAssignmentHash** | **string** | Configuration assignment hash that should be used when processing this request. | 
 **unityInstallationId** | **string** | Unique identifier that identifies an installation on the client’s device. The same player can have different installationIds if they have the game installed on different devices. It is available to all Unity packages that integrate with the Services SDK Core package. | 
 **analyticsUserId** | **string** | A unique string that identifies the player and is consistent across their subsequent play sessions for analytics purposes. It is the primary user identifier and it comes from the Core package. | 
 **after** | **string** | The &#x60;playersInventoryItemId&#x60; after which to retrieve the next page of balances. | 
 **limit** | **int32** | Number of items to be returned. Defaults to 20. | 
 **playersInventoryItemIds** | **[]string** | List of &#x60;playersInventoryItemIds&#x60; in array notation, for example, &#x60;playersInventoryItemIds[]&#x3D;ID1&amp;playersInventoryItemIds[]&#x3D;ID2&#x60;. | 
 **inventoryItemIds** | **[]string** | List of inventory item IDs in array notation, for example, &#x60;inventoryItemIds[]&#x3D;ID1&amp;inventoryItemIds[]&#x3D;ID2&#x60;. | 

### Return type

[**PlayerInventoryResponse**](PlayerInventoryResponse.md)

### Authorization

[Client](../README.md#Client)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInventoryItem

> InventoryResponse UpdateInventoryItem(ctx, projectId, playerId, playersInventoryItemId).InventoryRequestUpdate(inventoryRequestUpdate).ConfigAssignmentHash(configAssignmentHash).UnityInstallationId(unityInstallationId).AnalyticsUserId(analyticsUserId).Execute()

Update player's inventory item



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
    playersInventoryItemId := "playersInventoryItemId_example" // string | The `playersInventoryItemId` of the item to be updated.
    inventoryRequestUpdate := *openapiclient.NewInventoryRequestUpdate(map[string]interface{}(123)) // InventoryRequestUpdate | 
    configAssignmentHash := "configAssignmentHash_example" // string | Configuration assignment hash that should be used when processing this request. (optional)
    unityInstallationId := "unityInstallationId_example" // string | Unique identifier that identifies an installation on the client’s device. The same player can have different installationIds if they have the game installed on different devices. It is available to all Unity packages that integrate with the Services SDK Core package. (optional)
    analyticsUserId := "analyticsUserId_example" // string | A unique string that identifies the player and is consistent across their subsequent play sessions for analytics purposes. It is the primary user identifier and it comes from the Core package. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryAPI.UpdateInventoryItem(context.Background(), projectId, playerId, playersInventoryItemId).InventoryRequestUpdate(inventoryRequestUpdate).ConfigAssignmentHash(configAssignmentHash).UnityInstallationId(unityInstallationId).AnalyticsUserId(analyticsUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.UpdateInventoryItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInventoryItem`: InventoryResponse
    fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.UpdateInventoryItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project. | 
**playerId** | **string** | ID of the player. | 
**playersInventoryItemId** | **string** | The &#x60;playersInventoryItemId&#x60; of the item to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInventoryItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **inventoryRequestUpdate** | [**InventoryRequestUpdate**](InventoryRequestUpdate.md) |  | 
 **configAssignmentHash** | **string** | Configuration assignment hash that should be used when processing this request. | 
 **unityInstallationId** | **string** | Unique identifier that identifies an installation on the client’s device. The same player can have different installationIds if they have the game installed on different devices. It is available to all Unity packages that integrate with the Services SDK Core package. | 
 **analyticsUserId** | **string** | A unique string that identifies the player and is consistent across their subsequent play sessions for analytics purposes. It is the primary user identifier and it comes from the Core package. | 

### Return type

[**InventoryResponse**](InventoryResponse.md)

### Authorization

[Client](../README.md#Client)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

