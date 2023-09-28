# PlayerPurchaseVirtualResponseCosts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to [**[]CurrencyExchangeItem**](CurrencyExchangeItem.md) | Currency that was deducted in the purchase. | [optional] 
**Inventory** | [**[]InventoryExchangeItem**](InventoryExchangeItem.md) | Inventory that was deducted in the purchase. | 

## Methods

### NewPlayerPurchaseVirtualResponseCosts

`func NewPlayerPurchaseVirtualResponseCosts(inventory []InventoryExchangeItem, ) *PlayerPurchaseVirtualResponseCosts`

NewPlayerPurchaseVirtualResponseCosts instantiates a new PlayerPurchaseVirtualResponseCosts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerPurchaseVirtualResponseCostsWithDefaults

`func NewPlayerPurchaseVirtualResponseCostsWithDefaults() *PlayerPurchaseVirtualResponseCosts`

NewPlayerPurchaseVirtualResponseCostsWithDefaults instantiates a new PlayerPurchaseVirtualResponseCosts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *PlayerPurchaseVirtualResponseCosts) GetCurrency() []CurrencyExchangeItem`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PlayerPurchaseVirtualResponseCosts) GetCurrencyOk() (*[]CurrencyExchangeItem, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PlayerPurchaseVirtualResponseCosts) SetCurrency(v []CurrencyExchangeItem)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PlayerPurchaseVirtualResponseCosts) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetInventory

`func (o *PlayerPurchaseVirtualResponseCosts) GetInventory() []InventoryExchangeItem`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *PlayerPurchaseVirtualResponseCosts) GetInventoryOk() (*[]InventoryExchangeItem, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *PlayerPurchaseVirtualResponseCosts) SetInventory(v []InventoryExchangeItem)`

SetInventory sets Inventory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


