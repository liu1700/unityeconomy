/*
Economy API

# Introduction   This document outlines the API specification for the Economy API.   The Economy service allows the game client to retrieve and modify a player's economy resources in the cloud. # Concepts   ## Resources   Economy currently allows interaction with the following types of resources:   - Currencies: A resource that, when defined, contains two parameters: Initial and Max. The Initial parameter determines how much of the currency a game assigns to a player upon first interacting with the Economy system. The Max parameter determines how much of the currency the player is allowed to have.   - Inventory Items: A resource that doesn't have any set parameters; its intended use is to indicate the ownership or acquisition of an item in-game, for example, Sword and Shield.     A game client can add, remove or update the associated data of an instance of a configured inventory item from the player's inventory.   - Virtual Purchases: A transactional resource to implement a shop or trade feature. Allows the player to buy items/currencies using the previously defined currencies or inventory items.     A game client can redeem a virtual purchase and the player's account updates with the rewards if the costs criteria are met.   - Real Money Purchases: A transactional resource with the intended use to facilitate a shop or trade feature. Allows the player to buy any amount of items/currencies through an in-app purchase. Only uses the previously defined currencies or inventory items.     A game client can redeem a real money purchase and the player's account updates with the rewards.    The above resources also have an optional Custom Data parameter that can be populated with JSON data from the dashboard to allow clients to read bespoke data.   ## Writelock   The WriteLock is an integer that is automatically incremented serverside whenever a request that changes the stored value of a player's account or inventory.   The purpose of the WriteLock is to help prevent requests from the same or other game clients happening out-of-sync.   This parameter is optional, but when supplied with a request, the service does a comparison with the stored WriteLock on the server, and returns an error on mismatch.   ## Rate Limits   The API has rate limiting in place. Requests are limited on a per-player basis up to 150 requests per minute.   The API responds with a `429` HTTP status code if the requests exceed the rate limit.   Responses with a `429` status code include a `Retry-After` header to be used in conjunction with a client's retry logic, the value is the number of seconds until a request for the given player is accepted. 

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unityeconomy

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// PurchasesAPIService PurchasesAPI service
type PurchasesAPIService service

type ApiMakeVirtualPurchaseRequest struct {
	ctx context.Context
	ApiService *PurchasesAPIService
	projectId string
	playerId string
	playerPurchaseVirtualRequest *PlayerPurchaseVirtualRequest
	configAssignmentHash *string
	unityInstallationId *string
	analyticsUserId *string
}

func (r ApiMakeVirtualPurchaseRequest) PlayerPurchaseVirtualRequest(playerPurchaseVirtualRequest PlayerPurchaseVirtualRequest) ApiMakeVirtualPurchaseRequest {
	r.playerPurchaseVirtualRequest = &playerPurchaseVirtualRequest
	return r
}

// Configuration assignment hash that should be used when processing this request.
func (r ApiMakeVirtualPurchaseRequest) ConfigAssignmentHash(configAssignmentHash string) ApiMakeVirtualPurchaseRequest {
	r.configAssignmentHash = &configAssignmentHash
	return r
}

// Unique identifier that identifies an installation on the client’s device. The same player can have different installationIds if they have the game installed on different devices. It is available to all Unity packages that integrate with the Services SDK Core package.
func (r ApiMakeVirtualPurchaseRequest) UnityInstallationId(unityInstallationId string) ApiMakeVirtualPurchaseRequest {
	r.unityInstallationId = &unityInstallationId
	return r
}

// A unique string that identifies the player and is consistent across their subsequent play sessions for analytics purposes. It is the primary user identifier and it comes from the Core package.
func (r ApiMakeVirtualPurchaseRequest) AnalyticsUserId(analyticsUserId string) ApiMakeVirtualPurchaseRequest {
	r.analyticsUserId = &analyticsUserId
	return r
}

func (r ApiMakeVirtualPurchaseRequest) Execute() (*PlayerPurchaseVirtualResponse, *http.Response, error) {
	return r.ApiService.MakeVirtualPurchaseExecute(r)
}

/*
MakeVirtualPurchase Make virtual purchase

Make a virtual purchase for a player.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId ID of the project.
 @param playerId ID of the player.
 @return ApiMakeVirtualPurchaseRequest
*/
func (a *PurchasesAPIService) MakeVirtualPurchase(ctx context.Context, projectId string, playerId string) ApiMakeVirtualPurchaseRequest {
	return ApiMakeVirtualPurchaseRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		playerId: playerId,
	}
}

// Execute executes the request
//  @return PlayerPurchaseVirtualResponse
func (a *PurchasesAPIService) MakeVirtualPurchaseExecute(r ApiMakeVirtualPurchaseRequest) (*PlayerPurchaseVirtualResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PlayerPurchaseVirtualResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PurchasesAPIService.MakeVirtualPurchase")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/projects/{projectId}/players/{playerId}/purchases/virtual"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"playerId"+"}", url.PathEscape(parameterValueToString(r.playerId, "playerId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.playerPurchaseVirtualRequest == nil {
		return localVarReturnValue, nil, reportError("playerPurchaseVirtualRequest is required and must be specified")
	}

	if r.configAssignmentHash != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "configAssignmentHash", r.configAssignmentHash, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.unityInstallationId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "unity-installation-id", r.unityInstallationId, "")
	}
	if r.analyticsUserId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "analytics-user-id", r.analyticsUserId, "")
	}
	// body params
	localVarPostBody = r.playerPurchaseVirtualRequest
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
			var v SetPlayerCurrencyBalance400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v BasicErrorResponse
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
			var v BasicErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v BasicErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v BasicErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v BasicErrorResponse
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

type ApiRedeemAppleAppStorePurchaseRequest struct {
	ctx context.Context
	ApiService *PurchasesAPIService
	projectId string
	playerId string
	playerPurchaseAppleappstoreRequest *PlayerPurchaseAppleappstoreRequest
	configAssignmentHash *string
	unityInstallationId *string
	analyticsUserId *string
}

func (r ApiRedeemAppleAppStorePurchaseRequest) PlayerPurchaseAppleappstoreRequest(playerPurchaseAppleappstoreRequest PlayerPurchaseAppleappstoreRequest) ApiRedeemAppleAppStorePurchaseRequest {
	r.playerPurchaseAppleappstoreRequest = &playerPurchaseAppleappstoreRequest
	return r
}

// Configuration assignment hash that should be used when processing this request.
func (r ApiRedeemAppleAppStorePurchaseRequest) ConfigAssignmentHash(configAssignmentHash string) ApiRedeemAppleAppStorePurchaseRequest {
	r.configAssignmentHash = &configAssignmentHash
	return r
}

// Unique identifier that identifies an installation on the client’s device. The same player can have different installationIds if they have the game installed on different devices. It is available to all Unity packages that integrate with the Services SDK Core package.
func (r ApiRedeemAppleAppStorePurchaseRequest) UnityInstallationId(unityInstallationId string) ApiRedeemAppleAppStorePurchaseRequest {
	r.unityInstallationId = &unityInstallationId
	return r
}

// A unique string that identifies the player and is consistent across their subsequent play sessions for analytics purposes. It is the primary user identifier and it comes from the Core package.
func (r ApiRedeemAppleAppStorePurchaseRequest) AnalyticsUserId(analyticsUserId string) ApiRedeemAppleAppStorePurchaseRequest {
	r.analyticsUserId = &analyticsUserId
	return r
}

func (r ApiRedeemAppleAppStorePurchaseRequest) Execute() (*PlayerPurchaseAppleappstoreResponse, *http.Response, error) {
	return r.ApiService.RedeemAppleAppStorePurchaseExecute(r)
}

/*
RedeemAppleAppStorePurchase Redeem Apple App Store purchase

Redeem an Apple App Store purchase for a player.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId ID of the project.
 @param playerId ID of the player.
 @return ApiRedeemAppleAppStorePurchaseRequest
*/
func (a *PurchasesAPIService) RedeemAppleAppStorePurchase(ctx context.Context, projectId string, playerId string) ApiRedeemAppleAppStorePurchaseRequest {
	return ApiRedeemAppleAppStorePurchaseRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		playerId: playerId,
	}
}

// Execute executes the request
//  @return PlayerPurchaseAppleappstoreResponse
func (a *PurchasesAPIService) RedeemAppleAppStorePurchaseExecute(r ApiRedeemAppleAppStorePurchaseRequest) (*PlayerPurchaseAppleappstoreResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PlayerPurchaseAppleappstoreResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PurchasesAPIService.RedeemAppleAppStorePurchase")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/projects/{projectId}/players/{playerId}/purchases/appleappstore"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"playerId"+"}", url.PathEscape(parameterValueToString(r.playerId, "playerId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.playerPurchaseAppleappstoreRequest == nil {
		return localVarReturnValue, nil, reportError("playerPurchaseAppleappstoreRequest is required and must be specified")
	}

	if r.configAssignmentHash != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "configAssignmentHash", r.configAssignmentHash, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.unityInstallationId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "unity-installation-id", r.unityInstallationId, "")
	}
	if r.analyticsUserId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "analytics-user-id", r.analyticsUserId, "")
	}
	// body params
	localVarPostBody = r.playerPurchaseAppleappstoreRequest
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
			var v SetPlayerCurrencyBalance400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v BasicErrorResponse
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
			var v BasicErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v RedeemAppleAppStorePurchase422Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v BasicErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v BasicErrorResponse
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

type ApiRedeemGooglePlayPurchaseRequest struct {
	ctx context.Context
	ApiService *PurchasesAPIService
	projectId string
	playerId string
	playerPurchaseGoogleplaystoreRequest *PlayerPurchaseGoogleplaystoreRequest
	configAssignmentHash *string
	unityInstallationId *string
	analyticsUserId *string
}

func (r ApiRedeemGooglePlayPurchaseRequest) PlayerPurchaseGoogleplaystoreRequest(playerPurchaseGoogleplaystoreRequest PlayerPurchaseGoogleplaystoreRequest) ApiRedeemGooglePlayPurchaseRequest {
	r.playerPurchaseGoogleplaystoreRequest = &playerPurchaseGoogleplaystoreRequest
	return r
}

// Configuration assignment hash that should be used when processing this request.
func (r ApiRedeemGooglePlayPurchaseRequest) ConfigAssignmentHash(configAssignmentHash string) ApiRedeemGooglePlayPurchaseRequest {
	r.configAssignmentHash = &configAssignmentHash
	return r
}

// Unique identifier that identifies an installation on the client’s device. The same player can have different installationIds if they have the game installed on different devices. It is available to all Unity packages that integrate with the Services SDK Core package.
func (r ApiRedeemGooglePlayPurchaseRequest) UnityInstallationId(unityInstallationId string) ApiRedeemGooglePlayPurchaseRequest {
	r.unityInstallationId = &unityInstallationId
	return r
}

// A unique string that identifies the player and is consistent across their subsequent play sessions for analytics purposes. It is the primary user identifier and it comes from the Core package.
func (r ApiRedeemGooglePlayPurchaseRequest) AnalyticsUserId(analyticsUserId string) ApiRedeemGooglePlayPurchaseRequest {
	r.analyticsUserId = &analyticsUserId
	return r
}

func (r ApiRedeemGooglePlayPurchaseRequest) Execute() (*PlayerPurchaseGoogleplaystoreResponse, *http.Response, error) {
	return r.ApiService.RedeemGooglePlayPurchaseExecute(r)
}

/*
RedeemGooglePlayPurchase Redeem Google Play purchase

Redeem a Google Play store purchase for a player.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId ID of the project.
 @param playerId ID of the player.
 @return ApiRedeemGooglePlayPurchaseRequest
*/
func (a *PurchasesAPIService) RedeemGooglePlayPurchase(ctx context.Context, projectId string, playerId string) ApiRedeemGooglePlayPurchaseRequest {
	return ApiRedeemGooglePlayPurchaseRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		playerId: playerId,
	}
}

// Execute executes the request
//  @return PlayerPurchaseGoogleplaystoreResponse
func (a *PurchasesAPIService) RedeemGooglePlayPurchaseExecute(r ApiRedeemGooglePlayPurchaseRequest) (*PlayerPurchaseGoogleplaystoreResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PlayerPurchaseGoogleplaystoreResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PurchasesAPIService.RedeemGooglePlayPurchase")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/projects/{projectId}/players/{playerId}/purchases/googleplaystore"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"playerId"+"}", url.PathEscape(parameterValueToString(r.playerId, "playerId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.playerPurchaseGoogleplaystoreRequest == nil {
		return localVarReturnValue, nil, reportError("playerPurchaseGoogleplaystoreRequest is required and must be specified")
	}

	if r.configAssignmentHash != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "configAssignmentHash", r.configAssignmentHash, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.unityInstallationId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "unity-installation-id", r.unityInstallationId, "")
	}
	if r.analyticsUserId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "analytics-user-id", r.analyticsUserId, "")
	}
	// body params
	localVarPostBody = r.playerPurchaseGoogleplaystoreRequest
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
			var v SetPlayerCurrencyBalance400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v BasicErrorResponse
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
			var v BasicErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v RedeemGooglePlayPurchase422Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v BasicErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v BasicErrorResponse
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