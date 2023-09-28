# InventoryExchangeItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the inventory item. | 
**Amount** | **int32** | Number of player inventory items. | 
**PlayersInventoryItemIds** | **[]string** | The &#x60;playersInventoryItemIds&#x60; for the player&#39;s items to be added or removed. | 

## Methods

### NewInventoryExchangeItem

`func NewInventoryExchangeItem(id string, amount int32, playersInventoryItemIds []string, ) *InventoryExchangeItem`

NewInventoryExchangeItem instantiates a new InventoryExchangeItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryExchangeItemWithDefaults

`func NewInventoryExchangeItemWithDefaults() *InventoryExchangeItem`

NewInventoryExchangeItemWithDefaults instantiates a new InventoryExchangeItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InventoryExchangeItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InventoryExchangeItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InventoryExchangeItem) SetId(v string)`

SetId sets Id field to given value.


### GetAmount

`func (o *InventoryExchangeItem) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InventoryExchangeItem) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InventoryExchangeItem) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetPlayersInventoryItemIds

`func (o *InventoryExchangeItem) GetPlayersInventoryItemIds() []string`

GetPlayersInventoryItemIds returns the PlayersInventoryItemIds field if non-nil, zero value otherwise.

### GetPlayersInventoryItemIdsOk

`func (o *InventoryExchangeItem) GetPlayersInventoryItemIdsOk() (*[]string, bool)`

GetPlayersInventoryItemIdsOk returns a tuple with the PlayersInventoryItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayersInventoryItemIds

`func (o *InventoryExchangeItem) SetPlayersInventoryItemIds(v []string)`

SetPlayersInventoryItemIds sets PlayersInventoryItemIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


