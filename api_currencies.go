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


// CurrenciesAPIService CurrenciesAPI service
type CurrenciesAPIService service

type ApiDecrementPlayerCurrencyBalanceRequest struct {
	ctx context.Context
	ApiService *CurrenciesAPIService
	projectId string
	playerId string
	currencyId string
	currencyModifyBalanceRequest *CurrencyModifyBalanceRequest
	configAssignmentHash *string
	unityInstallationId *string
	analyticsUserId *string
}

func (r ApiDecrementPlayerCurrencyBalanceRequest) CurrencyModifyBalanceRequest(currencyModifyBalanceRequest CurrencyModifyBalanceRequest) ApiDecrementPlayerCurrencyBalanceRequest {
	r.currencyModifyBalanceRequest = &currencyModifyBalanceRequest
	return r
}

// Configuration assignment hash that should be used when processing this request.
func (r ApiDecrementPlayerCurrencyBalanceRequest) ConfigAssignmentHash(configAssignmentHash string) ApiDecrementPlayerCurrencyBalanceRequest {
	r.configAssignmentHash = &configAssignmentHash
	return r
}

// Unique identifier that identifies an installation on the client’s device. The same player can have different installationIds if they have the game installed on different devices. It is available to all Unity packages that integrate with the Services SDK Core package.
func (r ApiDecrementPlayerCurrencyBalanceRequest) UnityInstallationId(unityInstallationId string) ApiDecrementPlayerCurrencyBalanceRequest {
	r.unityInstallationId = &unityInstallationId
	return r
}

// A unique string that identifies the player and is consistent across their subsequent play sessions for analytics purposes. It is the primary user identifier and it comes from the Core package.
func (r ApiDecrementPlayerCurrencyBalanceRequest) AnalyticsUserId(analyticsUserId string) ApiDecrementPlayerCurrencyBalanceRequest {
	r.analyticsUserId = &analyticsUserId
	return r
}

func (r ApiDecrementPlayerCurrencyBalanceRequest) Execute() (*CurrencyBalanceResponse, *http.Response, error) {
	return r.ApiService.DecrementPlayerCurrencyBalanceExecute(r)
}

/*
DecrementPlayerCurrencyBalance Decrement currency balance

Decrements a player's currency balance by a given value.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId ID of the project.
 @param playerId ID of the player.
 @param currencyId Resource ID of the currency.
 @return ApiDecrementPlayerCurrencyBalanceRequest
*/
func (a *CurrenciesAPIService) DecrementPlayerCurrencyBalance(ctx context.Context, projectId string, playerId string, currencyId string) ApiDecrementPlayerCurrencyBalanceRequest {
	return ApiDecrementPlayerCurrencyBalanceRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		playerId: playerId,
		currencyId: currencyId,
	}
}

// Execute executes the request
//  @return CurrencyBalanceResponse
func (a *CurrenciesAPIService) DecrementPlayerCurrencyBalanceExecute(r ApiDecrementPlayerCurrencyBalanceRequest) (*CurrencyBalanceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CurrencyBalanceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CurrenciesAPIService.DecrementPlayerCurrencyBalance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/projects/{projectId}/players/{playerId}/currencies/{currencyId}/decrement"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"playerId"+"}", url.PathEscape(parameterValueToString(r.playerId, "playerId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"currencyId"+"}", url.PathEscape(parameterValueToString(r.currencyId, "currencyId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.currencyModifyBalanceRequest == nil {
		return localVarReturnValue, nil, reportError("currencyModifyBalanceRequest is required and must be specified")
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
	localVarPostBody = r.currencyModifyBalanceRequest
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
		if localVarHTTPResponse.StatusCode == 409 {
			var v ErrorResponseConflictCurrencyBalance
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

type ApiGetPlayerCurrenciesRequest struct {
	ctx context.Context
	ApiService *CurrenciesAPIService
	projectId string
	playerId string
	configAssignmentHash *string
	unityInstallationId *string
	analyticsUserId *string
	after *string
	limit *int32
}

// Configuration assignment hash that should be used when processing this request.
func (r ApiGetPlayerCurrenciesRequest) ConfigAssignmentHash(configAssignmentHash string) ApiGetPlayerCurrenciesRequest {
	r.configAssignmentHash = &configAssignmentHash
	return r
}

// Unique identifier that identifies an installation on the client’s device. The same player can have different installationIds if they have the game installed on different devices. It is available to all Unity packages that integrate with the Services SDK Core package.
func (r ApiGetPlayerCurrenciesRequest) UnityInstallationId(unityInstallationId string) ApiGetPlayerCurrenciesRequest {
	r.unityInstallationId = &unityInstallationId
	return r
}

// A unique string that identifies the player and is consistent across their subsequent play sessions for analytics purposes. It is the primary user identifier and it comes from the Core package.
func (r ApiGetPlayerCurrenciesRequest) AnalyticsUserId(analyticsUserId string) ApiGetPlayerCurrenciesRequest {
	r.analyticsUserId = &analyticsUserId
	return r
}

// The currency ID after which to retrieve the next page of balances.
func (r ApiGetPlayerCurrenciesRequest) After(after string) ApiGetPlayerCurrenciesRequest {
	r.after = &after
	return r
}

// Number of currencies to be returned. Defaults to 20.
func (r ApiGetPlayerCurrenciesRequest) Limit(limit int32) ApiGetPlayerCurrenciesRequest {
	r.limit = &limit
	return r
}

func (r ApiGetPlayerCurrenciesRequest) Execute() (*PlayerCurrencyBalanceResponse, *http.Response, error) {
	return r.ApiService.GetPlayerCurrenciesExecute(r)
}

/*
GetPlayerCurrencies Player currency balances

Get a list of currency balances for a player.
Results ordered in ascending currency ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId ID of the project.
 @param playerId ID of the player.
 @return ApiGetPlayerCurrenciesRequest
*/
func (a *CurrenciesAPIService) GetPlayerCurrencies(ctx context.Context, projectId string, playerId string) ApiGetPlayerCurrenciesRequest {
	return ApiGetPlayerCurrenciesRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		playerId: playerId,
	}
}

// Execute executes the request
//  @return PlayerCurrencyBalanceResponse
func (a *CurrenciesAPIService) GetPlayerCurrenciesExecute(r ApiGetPlayerCurrenciesRequest) (*PlayerCurrencyBalanceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PlayerCurrencyBalanceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CurrenciesAPIService.GetPlayerCurrencies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/projects/{projectId}/players/{playerId}/currencies"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"playerId"+"}", url.PathEscape(parameterValueToString(r.playerId, "playerId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.configAssignmentHash != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "configAssignmentHash", r.configAssignmentHash, "")
	}
	if r.after != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "after", r.after, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiIncrementPlayerCurrencyBalanceRequest struct {
	ctx context.Context
	ApiService *CurrenciesAPIService
	projectId string
	playerId string
	currencyId string
	currencyModifyBalanceRequest *CurrencyModifyBalanceRequest
	configAssignmentHash *string
	unityInstallationId *string
	analyticsUserId *string
}

// 
func (r ApiIncrementPlayerCurrencyBalanceRequest) CurrencyModifyBalanceRequest(currencyModifyBalanceRequest CurrencyModifyBalanceRequest) ApiIncrementPlayerCurrencyBalanceRequest {
	r.currencyModifyBalanceRequest = &currencyModifyBalanceRequest
	return r
}

// Configuration assignment hash that should be used when processing this request.
func (r ApiIncrementPlayerCurrencyBalanceRequest) ConfigAssignmentHash(configAssignmentHash string) ApiIncrementPlayerCurrencyBalanceRequest {
	r.configAssignmentHash = &configAssignmentHash
	return r
}

// Unique identifier that identifies an installation on the client’s device. The same player can have different installationIds if they have the game installed on different devices. It is available to all Unity packages that integrate with the Services SDK Core package.
func (r ApiIncrementPlayerCurrencyBalanceRequest) UnityInstallationId(unityInstallationId string) ApiIncrementPlayerCurrencyBalanceRequest {
	r.unityInstallationId = &unityInstallationId
	return r
}

// A unique string that identifies the player and is consistent across their subsequent play sessions for analytics purposes. It is the primary user identifier and it comes from the Core package.
func (r ApiIncrementPlayerCurrencyBalanceRequest) AnalyticsUserId(analyticsUserId string) ApiIncrementPlayerCurrencyBalanceRequest {
	r.analyticsUserId = &analyticsUserId
	return r
}

func (r ApiIncrementPlayerCurrencyBalanceRequest) Execute() (*CurrencyBalanceResponse, *http.Response, error) {
	return r.ApiService.IncrementPlayerCurrencyBalanceExecute(r)
}

/*
IncrementPlayerCurrencyBalance Increment currency balance

Increment a player's currency balance by a given value.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId ID of the project.
 @param playerId ID of the player.
 @param currencyId Resource ID of the currency.
 @return ApiIncrementPlayerCurrencyBalanceRequest
*/
func (a *CurrenciesAPIService) IncrementPlayerCurrencyBalance(ctx context.Context, projectId string, playerId string, currencyId string) ApiIncrementPlayerCurrencyBalanceRequest {
	return ApiIncrementPlayerCurrencyBalanceRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		playerId: playerId,
		currencyId: currencyId,
	}
}

// Execute executes the request
//  @return CurrencyBalanceResponse
func (a *CurrenciesAPIService) IncrementPlayerCurrencyBalanceExecute(r ApiIncrementPlayerCurrencyBalanceRequest) (*CurrencyBalanceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CurrencyBalanceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CurrenciesAPIService.IncrementPlayerCurrencyBalance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/projects/{projectId}/players/{playerId}/currencies/{currencyId}/increment"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"playerId"+"}", url.PathEscape(parameterValueToString(r.playerId, "playerId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"currencyId"+"}", url.PathEscape(parameterValueToString(r.currencyId, "currencyId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.currencyModifyBalanceRequest == nil {
		return localVarReturnValue, nil, reportError("currencyModifyBalanceRequest is required and must be specified")
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
	localVarPostBody = r.currencyModifyBalanceRequest
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
		if localVarHTTPResponse.StatusCode == 409 {
			var v ErrorResponseConflictCurrencyBalance
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

type ApiSetPlayerCurrencyBalanceRequest struct {
	ctx context.Context
	ApiService *CurrenciesAPIService
	projectId string
	playerId string
	currencyId string
	currencyBalanceRequest *CurrencyBalanceRequest
	configAssignmentHash *string
	unityInstallationId *string
	analyticsUserId *string
}

func (r ApiSetPlayerCurrencyBalanceRequest) CurrencyBalanceRequest(currencyBalanceRequest CurrencyBalanceRequest) ApiSetPlayerCurrencyBalanceRequest {
	r.currencyBalanceRequest = &currencyBalanceRequest
	return r
}

// Configuration assignment hash that should be used when processing this request.
func (r ApiSetPlayerCurrencyBalanceRequest) ConfigAssignmentHash(configAssignmentHash string) ApiSetPlayerCurrencyBalanceRequest {
	r.configAssignmentHash = &configAssignmentHash
	return r
}

// Unique identifier that identifies an installation on the client’s device. The same player can have different installationIds if they have the game installed on different devices. It is available to all Unity packages that integrate with the Services SDK Core package.
func (r ApiSetPlayerCurrencyBalanceRequest) UnityInstallationId(unityInstallationId string) ApiSetPlayerCurrencyBalanceRequest {
	r.unityInstallationId = &unityInstallationId
	return r
}

// A unique string that identifies the player and is consistent across their subsequent play sessions for analytics purposes. It is the primary user identifier and it comes from the Core package.
func (r ApiSetPlayerCurrencyBalanceRequest) AnalyticsUserId(analyticsUserId string) ApiSetPlayerCurrencyBalanceRequest {
	r.analyticsUserId = &analyticsUserId
	return r
}

func (r ApiSetPlayerCurrencyBalanceRequest) Execute() (*CurrencyBalanceResponse, *http.Response, error) {
	return r.ApiService.SetPlayerCurrencyBalanceExecute(r)
}

/*
SetPlayerCurrencyBalance Set currency balance

Set a player's currency balance.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId ID of the project.
 @param playerId ID of the player.
 @param currencyId Resource ID of the currency.
 @return ApiSetPlayerCurrencyBalanceRequest
*/
func (a *CurrenciesAPIService) SetPlayerCurrencyBalance(ctx context.Context, projectId string, playerId string, currencyId string) ApiSetPlayerCurrencyBalanceRequest {
	return ApiSetPlayerCurrencyBalanceRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		playerId: playerId,
		currencyId: currencyId,
	}
}

// Execute executes the request
//  @return CurrencyBalanceResponse
func (a *CurrenciesAPIService) SetPlayerCurrencyBalanceExecute(r ApiSetPlayerCurrencyBalanceRequest) (*CurrencyBalanceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CurrencyBalanceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CurrenciesAPIService.SetPlayerCurrencyBalance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/projects/{projectId}/players/{playerId}/currencies/{currencyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"playerId"+"}", url.PathEscape(parameterValueToString(r.playerId, "playerId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"currencyId"+"}", url.PathEscape(parameterValueToString(r.currencyId, "currencyId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.currencyBalanceRequest == nil {
		return localVarReturnValue, nil, reportError("currencyBalanceRequest is required and must be specified")
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
	localVarPostBody = r.currencyBalanceRequest
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
		if localVarHTTPResponse.StatusCode == 409 {
			var v ErrorResponseConflictCurrencyBalance
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