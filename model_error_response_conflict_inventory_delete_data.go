/*
Economy API

# Introduction   This document outlines the API specification for the Economy API.   The Economy service allows the game client to retrieve and modify a player's economy resources in the cloud. # Concepts   ## Resources   Economy currently allows interaction with the following types of resources:   - Currencies: A resource that, when defined, contains two parameters: Initial and Max. The Initial parameter determines how much of the currency a game assigns to a player upon first interacting with the Economy system. The Max parameter determines how much of the currency the player is allowed to have.   - Inventory Items: A resource that doesn't have any set parameters; its intended use is to indicate the ownership or acquisition of an item in-game, for example, Sword and Shield.     A game client can add, remove or update the associated data of an instance of a configured inventory item from the player's inventory.   - Virtual Purchases: A transactional resource to implement a shop or trade feature. Allows the player to buy items/currencies using the previously defined currencies or inventory items.     A game client can redeem a virtual purchase and the player's account updates with the rewards if the costs criteria are met.   - Real Money Purchases: A transactional resource with the intended use to facilitate a shop or trade feature. Allows the player to buy any amount of items/currencies through an in-app purchase. Only uses the previously defined currencies or inventory items.     A game client can redeem a real money purchase and the player's account updates with the rewards.    The above resources also have an optional Custom Data parameter that can be populated with JSON data from the dashboard to allow clients to read bespoke data.   ## Writelock   The WriteLock is an integer that is automatically incremented serverside whenever a request that changes the stored value of a player's account or inventory.   The purpose of the WriteLock is to help prevent requests from the same or other game clients happening out-of-sync.   This parameter is optional, but when supplied with a request, the service does a comparison with the stored WriteLock on the server, and returns an error on mismatch.   ## Rate Limits   The API has rate limiting in place. Requests are limited on a per-player basis up to 150 requests per minute.   The API responds with a `429` HTTP status code if the requests exceed the rate limit.   Responses with a `429` status code include a `Retry-After` header to be used in conjunction with a client's retry logic, the value is the number of seconds until a request for the given player is accepted. 

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unityeconomy

import (
	"encoding/json"
)

// checks if the ErrorResponseConflictInventoryDeleteData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorResponseConflictInventoryDeleteData{}

// ErrorResponseConflictInventoryDeleteData struct for ErrorResponseConflictInventoryDeleteData
type ErrorResponseConflictInventoryDeleteData struct {
	Attempted ErrorResponseConflictInventoryDeleteDataAttempted `json:"attempted"`
	Existing InventoryResponse `json:"existing"`
}

// NewErrorResponseConflictInventoryDeleteData instantiates a new ErrorResponseConflictInventoryDeleteData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorResponseConflictInventoryDeleteData(attempted ErrorResponseConflictInventoryDeleteDataAttempted, existing InventoryResponse) *ErrorResponseConflictInventoryDeleteData {
	this := ErrorResponseConflictInventoryDeleteData{}
	this.Attempted = attempted
	this.Existing = existing
	return &this
}

// NewErrorResponseConflictInventoryDeleteDataWithDefaults instantiates a new ErrorResponseConflictInventoryDeleteData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorResponseConflictInventoryDeleteDataWithDefaults() *ErrorResponseConflictInventoryDeleteData {
	this := ErrorResponseConflictInventoryDeleteData{}
	return &this
}

// GetAttempted returns the Attempted field value
func (o *ErrorResponseConflictInventoryDeleteData) GetAttempted() ErrorResponseConflictInventoryDeleteDataAttempted {
	if o == nil {
		var ret ErrorResponseConflictInventoryDeleteDataAttempted
		return ret
	}

	return o.Attempted
}

// GetAttemptedOk returns a tuple with the Attempted field value
// and a boolean to check if the value has been set.
func (o *ErrorResponseConflictInventoryDeleteData) GetAttemptedOk() (*ErrorResponseConflictInventoryDeleteDataAttempted, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attempted, true
}

// SetAttempted sets field value
func (o *ErrorResponseConflictInventoryDeleteData) SetAttempted(v ErrorResponseConflictInventoryDeleteDataAttempted) {
	o.Attempted = v
}

// GetExisting returns the Existing field value
func (o *ErrorResponseConflictInventoryDeleteData) GetExisting() InventoryResponse {
	if o == nil {
		var ret InventoryResponse
		return ret
	}

	return o.Existing
}

// GetExistingOk returns a tuple with the Existing field value
// and a boolean to check if the value has been set.
func (o *ErrorResponseConflictInventoryDeleteData) GetExistingOk() (*InventoryResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Existing, true
}

// SetExisting sets field value
func (o *ErrorResponseConflictInventoryDeleteData) SetExisting(v InventoryResponse) {
	o.Existing = v
}

func (o ErrorResponseConflictInventoryDeleteData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorResponseConflictInventoryDeleteData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attempted"] = o.Attempted
	toSerialize["existing"] = o.Existing
	return toSerialize, nil
}

type NullableErrorResponseConflictInventoryDeleteData struct {
	value *ErrorResponseConflictInventoryDeleteData
	isSet bool
}

func (v NullableErrorResponseConflictInventoryDeleteData) Get() *ErrorResponseConflictInventoryDeleteData {
	return v.value
}

func (v *NullableErrorResponseConflictInventoryDeleteData) Set(val *ErrorResponseConflictInventoryDeleteData) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorResponseConflictInventoryDeleteData) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorResponseConflictInventoryDeleteData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorResponseConflictInventoryDeleteData(val *ErrorResponseConflictInventoryDeleteData) *NullableErrorResponseConflictInventoryDeleteData {
	return &NullableErrorResponseConflictInventoryDeleteData{value: val, isSet: true}
}

func (v NullableErrorResponseConflictInventoryDeleteData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorResponseConflictInventoryDeleteData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


