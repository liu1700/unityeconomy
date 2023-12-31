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

// checks if the VirtualPurchaseResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualPurchaseResource{}

// VirtualPurchaseResource The definition for a economy resource that represents a virtual purchase.
type VirtualPurchaseResource struct {
	// Identifier for the resource.
	Id string `json:"id"`
	// Name of the resource.
	Name string `json:"name"`
	// Type of the item, for example `CURRENCY`, `INVENTORY_ITEM`, `VIRTUAL_PURCHASE`, `MONEY_PURCHASE`.
	Type string `json:"type"`
	Created ModifiedMetadata `json:"created"`
	Modified ModifiedMetadata `json:"modified"`
	// The costs deducted from the player when making the purchase. A cost is an ID of a currency or inventory item, plus an amount.
	Costs []Cost `json:"costs,omitempty"`
	// The rewards credited to the player when making the purchase. A reward is composed of the ID of a currency or inventory item, the amount of that currency or item, and the default instance data (for inventory items).
	Rewards []Reward `json:"rewards,omitempty"`
	CustomData map[string]interface{} `json:"customData,omitempty"`
}

// NewVirtualPurchaseResource instantiates a new VirtualPurchaseResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualPurchaseResource(id string, name string, type_ string, created ModifiedMetadata, modified ModifiedMetadata) *VirtualPurchaseResource {
	this := VirtualPurchaseResource{}
	this.Id = id
	this.Name = name
	this.Type = type_
	this.Created = created
	this.Modified = modified
	return &this
}

// NewVirtualPurchaseResourceWithDefaults instantiates a new VirtualPurchaseResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualPurchaseResourceWithDefaults() *VirtualPurchaseResource {
	this := VirtualPurchaseResource{}
	return &this
}

// GetId returns the Id field value
func (o *VirtualPurchaseResource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VirtualPurchaseResource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VirtualPurchaseResource) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *VirtualPurchaseResource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VirtualPurchaseResource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VirtualPurchaseResource) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *VirtualPurchaseResource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VirtualPurchaseResource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VirtualPurchaseResource) SetType(v string) {
	o.Type = v
}

// GetCreated returns the Created field value
func (o *VirtualPurchaseResource) GetCreated() ModifiedMetadata {
	if o == nil {
		var ret ModifiedMetadata
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *VirtualPurchaseResource) GetCreatedOk() (*ModifiedMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *VirtualPurchaseResource) SetCreated(v ModifiedMetadata) {
	o.Created = v
}

// GetModified returns the Modified field value
func (o *VirtualPurchaseResource) GetModified() ModifiedMetadata {
	if o == nil {
		var ret ModifiedMetadata
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *VirtualPurchaseResource) GetModifiedOk() (*ModifiedMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *VirtualPurchaseResource) SetModified(v ModifiedMetadata) {
	o.Modified = v
}

// GetCosts returns the Costs field value if set, zero value otherwise.
func (o *VirtualPurchaseResource) GetCosts() []Cost {
	if o == nil || IsNil(o.Costs) {
		var ret []Cost
		return ret
	}
	return o.Costs
}

// GetCostsOk returns a tuple with the Costs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualPurchaseResource) GetCostsOk() ([]Cost, bool) {
	if o == nil || IsNil(o.Costs) {
		return nil, false
	}
	return o.Costs, true
}

// HasCosts returns a boolean if a field has been set.
func (o *VirtualPurchaseResource) HasCosts() bool {
	if o != nil && !IsNil(o.Costs) {
		return true
	}

	return false
}

// SetCosts gets a reference to the given []Cost and assigns it to the Costs field.
func (o *VirtualPurchaseResource) SetCosts(v []Cost) {
	o.Costs = v
}

// GetRewards returns the Rewards field value if set, zero value otherwise.
func (o *VirtualPurchaseResource) GetRewards() []Reward {
	if o == nil || IsNil(o.Rewards) {
		var ret []Reward
		return ret
	}
	return o.Rewards
}

// GetRewardsOk returns a tuple with the Rewards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualPurchaseResource) GetRewardsOk() ([]Reward, bool) {
	if o == nil || IsNil(o.Rewards) {
		return nil, false
	}
	return o.Rewards, true
}

// HasRewards returns a boolean if a field has been set.
func (o *VirtualPurchaseResource) HasRewards() bool {
	if o != nil && !IsNil(o.Rewards) {
		return true
	}

	return false
}

// SetRewards gets a reference to the given []Reward and assigns it to the Rewards field.
func (o *VirtualPurchaseResource) SetRewards(v []Reward) {
	o.Rewards = v
}

// GetCustomData returns the CustomData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualPurchaseResource) GetCustomData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomData
}

// GetCustomDataOk returns a tuple with the CustomData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualPurchaseResource) GetCustomDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomData) {
		return map[string]interface{}{}, false
	}
	return o.CustomData, true
}

// HasCustomData returns a boolean if a field has been set.
func (o *VirtualPurchaseResource) HasCustomData() bool {
	if o != nil && IsNil(o.CustomData) {
		return true
	}

	return false
}

// SetCustomData gets a reference to the given map[string]interface{} and assigns it to the CustomData field.
func (o *VirtualPurchaseResource) SetCustomData(v map[string]interface{}) {
	o.CustomData = v
}

func (o VirtualPurchaseResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualPurchaseResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["created"] = o.Created
	toSerialize["modified"] = o.Modified
	if !IsNil(o.Costs) {
		toSerialize["costs"] = o.Costs
	}
	if !IsNil(o.Rewards) {
		toSerialize["rewards"] = o.Rewards
	}
	if o.CustomData != nil {
		toSerialize["customData"] = o.CustomData
	}
	return toSerialize, nil
}

type NullableVirtualPurchaseResource struct {
	value *VirtualPurchaseResource
	isSet bool
}

func (v NullableVirtualPurchaseResource) Get() *VirtualPurchaseResource {
	return v.value
}

func (v *NullableVirtualPurchaseResource) Set(val *VirtualPurchaseResource) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualPurchaseResource) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualPurchaseResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualPurchaseResource(val *VirtualPurchaseResource) *NullableVirtualPurchaseResource {
	return &NullableVirtualPurchaseResource{value: val, isSet: true}
}

func (v NullableVirtualPurchaseResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualPurchaseResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


