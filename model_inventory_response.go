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

// checks if the InventoryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryResponse{}

// InventoryResponse struct for InventoryResponse
type InventoryResponse struct {
	// ID of the player's inventory item.
	PlayersInventoryItemId string `json:"playersInventoryItemId"`
	// Resource ID of the inventory item configuration associated with this instance.
	InventoryItemId string `json:"inventoryItemId"`
	// Instance data. Max size when serialized 5 KB.
	InstanceData map[string]interface{} `json:"instanceData,omitempty"`
	// The write lock for the inventory item instance.
	WriteLock string `json:"writeLock"`
	Created ModifiedMetadata `json:"created"`
	Modified ModifiedMetadata `json:"modified"`
}

// NewInventoryResponse instantiates a new InventoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryResponse(playersInventoryItemId string, inventoryItemId string, writeLock string, created ModifiedMetadata, modified ModifiedMetadata) *InventoryResponse {
	this := InventoryResponse{}
	this.PlayersInventoryItemId = playersInventoryItemId
	this.InventoryItemId = inventoryItemId
	this.WriteLock = writeLock
	this.Created = created
	this.Modified = modified
	return &this
}

// NewInventoryResponseWithDefaults instantiates a new InventoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryResponseWithDefaults() *InventoryResponse {
	this := InventoryResponse{}
	return &this
}

// GetPlayersInventoryItemId returns the PlayersInventoryItemId field value
func (o *InventoryResponse) GetPlayersInventoryItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlayersInventoryItemId
}

// GetPlayersInventoryItemIdOk returns a tuple with the PlayersInventoryItemId field value
// and a boolean to check if the value has been set.
func (o *InventoryResponse) GetPlayersInventoryItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlayersInventoryItemId, true
}

// SetPlayersInventoryItemId sets field value
func (o *InventoryResponse) SetPlayersInventoryItemId(v string) {
	o.PlayersInventoryItemId = v
}

// GetInventoryItemId returns the InventoryItemId field value
func (o *InventoryResponse) GetInventoryItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InventoryItemId
}

// GetInventoryItemIdOk returns a tuple with the InventoryItemId field value
// and a boolean to check if the value has been set.
func (o *InventoryResponse) GetInventoryItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InventoryItemId, true
}

// SetInventoryItemId sets field value
func (o *InventoryResponse) SetInventoryItemId(v string) {
	o.InventoryItemId = v
}

// GetInstanceData returns the InstanceData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryResponse) GetInstanceData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.InstanceData
}

// GetInstanceDataOk returns a tuple with the InstanceData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryResponse) GetInstanceDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.InstanceData) {
		return map[string]interface{}{}, false
	}
	return o.InstanceData, true
}

// HasInstanceData returns a boolean if a field has been set.
func (o *InventoryResponse) HasInstanceData() bool {
	if o != nil && IsNil(o.InstanceData) {
		return true
	}

	return false
}

// SetInstanceData gets a reference to the given map[string]interface{} and assigns it to the InstanceData field.
func (o *InventoryResponse) SetInstanceData(v map[string]interface{}) {
	o.InstanceData = v
}

// GetWriteLock returns the WriteLock field value
func (o *InventoryResponse) GetWriteLock() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WriteLock
}

// GetWriteLockOk returns a tuple with the WriteLock field value
// and a boolean to check if the value has been set.
func (o *InventoryResponse) GetWriteLockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WriteLock, true
}

// SetWriteLock sets field value
func (o *InventoryResponse) SetWriteLock(v string) {
	o.WriteLock = v
}

// GetCreated returns the Created field value
func (o *InventoryResponse) GetCreated() ModifiedMetadata {
	if o == nil {
		var ret ModifiedMetadata
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *InventoryResponse) GetCreatedOk() (*ModifiedMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *InventoryResponse) SetCreated(v ModifiedMetadata) {
	o.Created = v
}

// GetModified returns the Modified field value
func (o *InventoryResponse) GetModified() ModifiedMetadata {
	if o == nil {
		var ret ModifiedMetadata
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *InventoryResponse) GetModifiedOk() (*ModifiedMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *InventoryResponse) SetModified(v ModifiedMetadata) {
	o.Modified = v
}

func (o InventoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["playersInventoryItemId"] = o.PlayersInventoryItemId
	toSerialize["inventoryItemId"] = o.InventoryItemId
	if o.InstanceData != nil {
		toSerialize["instanceData"] = o.InstanceData
	}
	toSerialize["writeLock"] = o.WriteLock
	toSerialize["created"] = o.Created
	toSerialize["modified"] = o.Modified
	return toSerialize, nil
}

type NullableInventoryResponse struct {
	value *InventoryResponse
	isSet bool
}

func (v NullableInventoryResponse) Get() *InventoryResponse {
	return v.value
}

func (v *NullableInventoryResponse) Set(val *InventoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryResponse(val *InventoryResponse) *NullableInventoryResponse {
	return &NullableInventoryResponse{value: val, isSet: true}
}

func (v NullableInventoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


