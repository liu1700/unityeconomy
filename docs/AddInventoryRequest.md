# AddInventoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InventoryItemId** | **string** | Resource ID of the inventory item. | 
**PlayersInventoryItemId** | Pointer to **string** | A &#x60;playersInventoryItemId&#x60; for the item being created. If not given, Economy automatically generates the ID. An ID must be unique for a player. | [optional] 
**InstanceData** | Pointer to **map[string]interface{}** | Instance data to be saved against the new inventory item. Max size when serialized 5 KB. | [optional] 

## Methods

### NewAddInventoryRequest

`func NewAddInventoryRequest(inventoryItemId string, ) *AddInventoryRequest`

NewAddInventoryRequest instantiates a new AddInventoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddInventoryRequestWithDefaults

`func NewAddInventoryRequestWithDefaults() *AddInventoryRequest`

NewAddInventoryRequestWithDefaults instantiates a new AddInventoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInventoryItemId

`func (o *AddInventoryRequest) GetInventoryItemId() string`

GetInventoryItemId returns the InventoryItemId field if non-nil, zero value otherwise.

### GetInventoryItemIdOk

`func (o *AddInventoryRequest) GetInventoryItemIdOk() (*string, bool)`

GetInventoryItemIdOk returns a tuple with the InventoryItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryItemId

`func (o *AddInventoryRequest) SetInventoryItemId(v string)`

SetInventoryItemId sets InventoryItemId field to given value.


### GetPlayersInventoryItemId

`func (o *AddInventoryRequest) GetPlayersInventoryItemId() string`

GetPlayersInventoryItemId returns the PlayersInventoryItemId field if non-nil, zero value otherwise.

### GetPlayersInventoryItemIdOk

`func (o *AddInventoryRequest) GetPlayersInventoryItemIdOk() (*string, bool)`

GetPlayersInventoryItemIdOk returns a tuple with the PlayersInventoryItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayersInventoryItemId

`func (o *AddInventoryRequest) SetPlayersInventoryItemId(v string)`

SetPlayersInventoryItemId sets PlayersInventoryItemId field to given value.

### HasPlayersInventoryItemId

`func (o *AddInventoryRequest) HasPlayersInventoryItemId() bool`

HasPlayersInventoryItemId returns a boolean if a field has been set.

### GetInstanceData

`func (o *AddInventoryRequest) GetInstanceData() map[string]interface{}`

GetInstanceData returns the InstanceData field if non-nil, zero value otherwise.

### GetInstanceDataOk

`func (o *AddInventoryRequest) GetInstanceDataOk() (*map[string]interface{}, bool)`

GetInstanceDataOk returns a tuple with the InstanceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceData

`func (o *AddInventoryRequest) SetInstanceData(v map[string]interface{})`

SetInstanceData sets InstanceData field to given value.

### HasInstanceData

`func (o *AddInventoryRequest) HasInstanceData() bool`

HasInstanceData returns a boolean if a field has been set.

### SetInstanceDataNil

`func (o *AddInventoryRequest) SetInstanceDataNil(b bool)`

 SetInstanceDataNil sets the value for InstanceData to be an explicit nil

### UnsetInstanceData
`func (o *AddInventoryRequest) UnsetInstanceData()`

UnsetInstanceData ensures that no value is present for InstanceData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


