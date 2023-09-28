# PlayerPurchaseVirtualRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the purchase. | 
**PlayersInventoryItemIds** | Pointer to **[]string** | IDs of the player&#39;s inventory items that should be used for any item costs associated with the purchase. | [optional] 

## Methods

### NewPlayerPurchaseVirtualRequest

`func NewPlayerPurchaseVirtualRequest(id string, ) *PlayerPurchaseVirtualRequest`

NewPlayerPurchaseVirtualRequest instantiates a new PlayerPurchaseVirtualRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerPurchaseVirtualRequestWithDefaults

`func NewPlayerPurchaseVirtualRequestWithDefaults() *PlayerPurchaseVirtualRequest`

NewPlayerPurchaseVirtualRequestWithDefaults instantiates a new PlayerPurchaseVirtualRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlayerPurchaseVirtualRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlayerPurchaseVirtualRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlayerPurchaseVirtualRequest) SetId(v string)`

SetId sets Id field to given value.


### GetPlayersInventoryItemIds

`func (o *PlayerPurchaseVirtualRequest) GetPlayersInventoryItemIds() []string`

GetPlayersInventoryItemIds returns the PlayersInventoryItemIds field if non-nil, zero value otherwise.

### GetPlayersInventoryItemIdsOk

`func (o *PlayerPurchaseVirtualRequest) GetPlayersInventoryItemIdsOk() (*[]string, bool)`

GetPlayersInventoryItemIdsOk returns a tuple with the PlayersInventoryItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayersInventoryItemIds

`func (o *PlayerPurchaseVirtualRequest) SetPlayersInventoryItemIds(v []string)`

SetPlayersInventoryItemIds sets PlayersInventoryItemIds field to given value.

### HasPlayersInventoryItemIds

`func (o *PlayerPurchaseVirtualRequest) HasPlayersInventoryItemIds() bool`

HasPlayersInventoryItemIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


