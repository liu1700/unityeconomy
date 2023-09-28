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

// checks if the InventoryDeleteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryDeleteRequest{}

// InventoryDeleteRequest The request body has been deprecated in favor of using query parameters due to inconsistent HTTP client support for sending request bodies for DELETE requests. Please use the corresponding query parameters instead, which will take precedence over any request body properties if both are present.
type InventoryDeleteRequest struct {
	// The write lock for the inventory item instance. This property has been deprecated. Please use the `writeLock` query parameter instead.
	// Deprecated
	WriteLock *string `json:"writeLock,omitempty"`
}

// NewInventoryDeleteRequest instantiates a new InventoryDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryDeleteRequest() *InventoryDeleteRequest {
	this := InventoryDeleteRequest{}
	return &this
}

// NewInventoryDeleteRequestWithDefaults instantiates a new InventoryDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryDeleteRequestWithDefaults() *InventoryDeleteRequest {
	this := InventoryDeleteRequest{}
	return &this
}

// GetWriteLock returns the WriteLock field value if set, zero value otherwise.
// Deprecated
func (o *InventoryDeleteRequest) GetWriteLock() string {
	if o == nil || IsNil(o.WriteLock) {
		var ret string
		return ret
	}
	return *o.WriteLock
}

// GetWriteLockOk returns a tuple with the WriteLock field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *InventoryDeleteRequest) GetWriteLockOk() (*string, bool) {
	if o == nil || IsNil(o.WriteLock) {
		return nil, false
	}
	return o.WriteLock, true
}

// HasWriteLock returns a boolean if a field has been set.
func (o *InventoryDeleteRequest) HasWriteLock() bool {
	if o != nil && !IsNil(o.WriteLock) {
		return true
	}

	return false
}

// SetWriteLock gets a reference to the given string and assigns it to the WriteLock field.
// Deprecated
func (o *InventoryDeleteRequest) SetWriteLock(v string) {
	o.WriteLock = &v
}

func (o InventoryDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryDeleteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WriteLock) {
		toSerialize["writeLock"] = o.WriteLock
	}
	return toSerialize, nil
}

type NullableInventoryDeleteRequest struct {
	value *InventoryDeleteRequest
	isSet bool
}

func (v NullableInventoryDeleteRequest) Get() *InventoryDeleteRequest {
	return v.value
}

func (v *NullableInventoryDeleteRequest) Set(val *InventoryDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryDeleteRequest(val *InventoryDeleteRequest) *NullableInventoryDeleteRequest {
	return &NullableInventoryDeleteRequest{value: val, isSet: true}
}

func (v NullableInventoryDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


