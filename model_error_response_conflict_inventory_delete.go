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

// checks if the ErrorResponseConflictInventoryDelete type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorResponseConflictInventoryDelete{}

// ErrorResponseConflictInventoryDelete An error response sent back upon player inventory item conflict.
type ErrorResponseConflictInventoryDelete struct {
	Type string `json:"type"`
	Title string `json:"title"`
	// e.g 409
	Status int32 `json:"status"`
	// For example 10204
	Code int32 `json:"code"`
	Detail string `json:"detail"`
	Instance *string `json:"instance,omitempty"`
	Data ErrorResponseConflictInventoryDeleteData `json:"data"`
}

// NewErrorResponseConflictInventoryDelete instantiates a new ErrorResponseConflictInventoryDelete object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorResponseConflictInventoryDelete(type_ string, title string, status int32, code int32, detail string, data ErrorResponseConflictInventoryDeleteData) *ErrorResponseConflictInventoryDelete {
	this := ErrorResponseConflictInventoryDelete{}
	this.Type = type_
	this.Title = title
	this.Status = status
	this.Code = code
	this.Detail = detail
	this.Data = data
	return &this
}

// NewErrorResponseConflictInventoryDeleteWithDefaults instantiates a new ErrorResponseConflictInventoryDelete object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorResponseConflictInventoryDeleteWithDefaults() *ErrorResponseConflictInventoryDelete {
	this := ErrorResponseConflictInventoryDelete{}
	return &this
}

// GetType returns the Type field value
func (o *ErrorResponseConflictInventoryDelete) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ErrorResponseConflictInventoryDelete) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ErrorResponseConflictInventoryDelete) SetType(v string) {
	o.Type = v
}

// GetTitle returns the Title field value
func (o *ErrorResponseConflictInventoryDelete) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ErrorResponseConflictInventoryDelete) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ErrorResponseConflictInventoryDelete) SetTitle(v string) {
	o.Title = v
}

// GetStatus returns the Status field value
func (o *ErrorResponseConflictInventoryDelete) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ErrorResponseConflictInventoryDelete) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ErrorResponseConflictInventoryDelete) SetStatus(v int32) {
	o.Status = v
}

// GetCode returns the Code field value
func (o *ErrorResponseConflictInventoryDelete) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ErrorResponseConflictInventoryDelete) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ErrorResponseConflictInventoryDelete) SetCode(v int32) {
	o.Code = v
}

// GetDetail returns the Detail field value
func (o *ErrorResponseConflictInventoryDelete) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *ErrorResponseConflictInventoryDelete) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *ErrorResponseConflictInventoryDelete) SetDetail(v string) {
	o.Detail = v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *ErrorResponseConflictInventoryDelete) GetInstance() string {
	if o == nil || IsNil(o.Instance) {
		var ret string
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseConflictInventoryDelete) GetInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.Instance) {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *ErrorResponseConflictInventoryDelete) HasInstance() bool {
	if o != nil && !IsNil(o.Instance) {
		return true
	}

	return false
}

// SetInstance gets a reference to the given string and assigns it to the Instance field.
func (o *ErrorResponseConflictInventoryDelete) SetInstance(v string) {
	o.Instance = &v
}

// GetData returns the Data field value
func (o *ErrorResponseConflictInventoryDelete) GetData() ErrorResponseConflictInventoryDeleteData {
	if o == nil {
		var ret ErrorResponseConflictInventoryDeleteData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ErrorResponseConflictInventoryDelete) GetDataOk() (*ErrorResponseConflictInventoryDeleteData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ErrorResponseConflictInventoryDelete) SetData(v ErrorResponseConflictInventoryDeleteData) {
	o.Data = v
}

func (o ErrorResponseConflictInventoryDelete) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorResponseConflictInventoryDelete) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["title"] = o.Title
	toSerialize["status"] = o.Status
	toSerialize["code"] = o.Code
	toSerialize["detail"] = o.Detail
	if !IsNil(o.Instance) {
		toSerialize["instance"] = o.Instance
	}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableErrorResponseConflictInventoryDelete struct {
	value *ErrorResponseConflictInventoryDelete
	isSet bool
}

func (v NullableErrorResponseConflictInventoryDelete) Get() *ErrorResponseConflictInventoryDelete {
	return v.value
}

func (v *NullableErrorResponseConflictInventoryDelete) Set(val *ErrorResponseConflictInventoryDelete) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorResponseConflictInventoryDelete) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorResponseConflictInventoryDelete) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorResponseConflictInventoryDelete(val *ErrorResponseConflictInventoryDelete) *NullableErrorResponseConflictInventoryDelete {
	return &NullableErrorResponseConflictInventoryDelete{value: val, isSet: true}
}

func (v NullableErrorResponseConflictInventoryDelete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorResponseConflictInventoryDelete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


