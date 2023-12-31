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

// checks if the CurrencyModifyBalanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CurrencyModifyBalanceRequest{}

// CurrencyModifyBalanceRequest struct for CurrencyModifyBalanceRequest
type CurrencyModifyBalanceRequest struct {
	// Resource ID of the currency.
	// Deprecated
	CurrencyId *string `json:"currencyId,omitempty"`
	// The value by which to increment or decrement. Zero is allowed but results in no change to the currency balance.
	Amount int64 `json:"amount"`
	// The write lock for the currency balance.
	WriteLock *string `json:"writeLock,omitempty"`
}

// NewCurrencyModifyBalanceRequest instantiates a new CurrencyModifyBalanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrencyModifyBalanceRequest(amount int64) *CurrencyModifyBalanceRequest {
	this := CurrencyModifyBalanceRequest{}
	this.Amount = amount
	return &this
}

// NewCurrencyModifyBalanceRequestWithDefaults instantiates a new CurrencyModifyBalanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrencyModifyBalanceRequestWithDefaults() *CurrencyModifyBalanceRequest {
	this := CurrencyModifyBalanceRequest{}
	return &this
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise.
// Deprecated
func (o *CurrencyModifyBalanceRequest) GetCurrencyId() string {
	if o == nil || IsNil(o.CurrencyId) {
		var ret string
		return ret
	}
	return *o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CurrencyModifyBalanceRequest) GetCurrencyIdOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyId) {
		return nil, false
	}
	return o.CurrencyId, true
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *CurrencyModifyBalanceRequest) HasCurrencyId() bool {
	if o != nil && !IsNil(o.CurrencyId) {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given string and assigns it to the CurrencyId field.
// Deprecated
func (o *CurrencyModifyBalanceRequest) SetCurrencyId(v string) {
	o.CurrencyId = &v
}

// GetAmount returns the Amount field value
func (o *CurrencyModifyBalanceRequest) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CurrencyModifyBalanceRequest) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CurrencyModifyBalanceRequest) SetAmount(v int64) {
	o.Amount = v
}

// GetWriteLock returns the WriteLock field value if set, zero value otherwise.
func (o *CurrencyModifyBalanceRequest) GetWriteLock() string {
	if o == nil || IsNil(o.WriteLock) {
		var ret string
		return ret
	}
	return *o.WriteLock
}

// GetWriteLockOk returns a tuple with the WriteLock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrencyModifyBalanceRequest) GetWriteLockOk() (*string, bool) {
	if o == nil || IsNil(o.WriteLock) {
		return nil, false
	}
	return o.WriteLock, true
}

// HasWriteLock returns a boolean if a field has been set.
func (o *CurrencyModifyBalanceRequest) HasWriteLock() bool {
	if o != nil && !IsNil(o.WriteLock) {
		return true
	}

	return false
}

// SetWriteLock gets a reference to the given string and assigns it to the WriteLock field.
func (o *CurrencyModifyBalanceRequest) SetWriteLock(v string) {
	o.WriteLock = &v
}

func (o CurrencyModifyBalanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CurrencyModifyBalanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrencyId) {
		toSerialize["currencyId"] = o.CurrencyId
	}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.WriteLock) {
		toSerialize["writeLock"] = o.WriteLock
	}
	return toSerialize, nil
}

type NullableCurrencyModifyBalanceRequest struct {
	value *CurrencyModifyBalanceRequest
	isSet bool
}

func (v NullableCurrencyModifyBalanceRequest) Get() *CurrencyModifyBalanceRequest {
	return v.value
}

func (v *NullableCurrencyModifyBalanceRequest) Set(val *CurrencyModifyBalanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrencyModifyBalanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrencyModifyBalanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrencyModifyBalanceRequest(val *CurrencyModifyBalanceRequest) *NullableCurrencyModifyBalanceRequest {
	return &NullableCurrencyModifyBalanceRequest{value: val, isSet: true}
}

func (v NullableCurrencyModifyBalanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrencyModifyBalanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


