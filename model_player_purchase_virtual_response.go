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

// checks if the PlayerPurchaseVirtualResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlayerPurchaseVirtualResponse{}

// PlayerPurchaseVirtualResponse struct for PlayerPurchaseVirtualResponse
type PlayerPurchaseVirtualResponse struct {
	Costs PlayerPurchaseVirtualResponseCosts `json:"costs"`
	Rewards PlayerPurchaseVirtualResponseRewards `json:"rewards"`
}

// NewPlayerPurchaseVirtualResponse instantiates a new PlayerPurchaseVirtualResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlayerPurchaseVirtualResponse(costs PlayerPurchaseVirtualResponseCosts, rewards PlayerPurchaseVirtualResponseRewards) *PlayerPurchaseVirtualResponse {
	this := PlayerPurchaseVirtualResponse{}
	this.Costs = costs
	this.Rewards = rewards
	return &this
}

// NewPlayerPurchaseVirtualResponseWithDefaults instantiates a new PlayerPurchaseVirtualResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlayerPurchaseVirtualResponseWithDefaults() *PlayerPurchaseVirtualResponse {
	this := PlayerPurchaseVirtualResponse{}
	return &this
}

// GetCosts returns the Costs field value
func (o *PlayerPurchaseVirtualResponse) GetCosts() PlayerPurchaseVirtualResponseCosts {
	if o == nil {
		var ret PlayerPurchaseVirtualResponseCosts
		return ret
	}

	return o.Costs
}

// GetCostsOk returns a tuple with the Costs field value
// and a boolean to check if the value has been set.
func (o *PlayerPurchaseVirtualResponse) GetCostsOk() (*PlayerPurchaseVirtualResponseCosts, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Costs, true
}

// SetCosts sets field value
func (o *PlayerPurchaseVirtualResponse) SetCosts(v PlayerPurchaseVirtualResponseCosts) {
	o.Costs = v
}

// GetRewards returns the Rewards field value
func (o *PlayerPurchaseVirtualResponse) GetRewards() PlayerPurchaseVirtualResponseRewards {
	if o == nil {
		var ret PlayerPurchaseVirtualResponseRewards
		return ret
	}

	return o.Rewards
}

// GetRewardsOk returns a tuple with the Rewards field value
// and a boolean to check if the value has been set.
func (o *PlayerPurchaseVirtualResponse) GetRewardsOk() (*PlayerPurchaseVirtualResponseRewards, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rewards, true
}

// SetRewards sets field value
func (o *PlayerPurchaseVirtualResponse) SetRewards(v PlayerPurchaseVirtualResponseRewards) {
	o.Rewards = v
}

func (o PlayerPurchaseVirtualResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlayerPurchaseVirtualResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["costs"] = o.Costs
	toSerialize["rewards"] = o.Rewards
	return toSerialize, nil
}

type NullablePlayerPurchaseVirtualResponse struct {
	value *PlayerPurchaseVirtualResponse
	isSet bool
}

func (v NullablePlayerPurchaseVirtualResponse) Get() *PlayerPurchaseVirtualResponse {
	return v.value
}

func (v *NullablePlayerPurchaseVirtualResponse) Set(val *PlayerPurchaseVirtualResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePlayerPurchaseVirtualResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePlayerPurchaseVirtualResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlayerPurchaseVirtualResponse(val *PlayerPurchaseVirtualResponse) *NullablePlayerPurchaseVirtualResponse {
	return &NullablePlayerPurchaseVirtualResponse{value: val, isSet: true}
}

func (v NullablePlayerPurchaseVirtualResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlayerPurchaseVirtualResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


