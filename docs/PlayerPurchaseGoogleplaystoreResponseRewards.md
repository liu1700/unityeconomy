# PlayerPurchaseGoogleplaystoreResponseRewards

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | [**[]CurrencyExchangeItem**](CurrencyExchangeItem.md) | Currency that was credited in the purchase. | 
**Inventory** | [**[]InventoryExchangeItem**](InventoryExchangeItem.md) | Inventory that was credited in the purchase. | 

## Methods

### NewPlayerPurchaseGoogleplaystoreResponseRewards

`func NewPlayerPurchaseGoogleplaystoreResponseRewards(currency []CurrencyExchangeItem, inventory []InventoryExchangeItem, ) *PlayerPurchaseGoogleplaystoreResponseRewards`

NewPlayerPurchaseGoogleplaystoreResponseRewards instantiates a new PlayerPurchaseGoogleplaystoreResponseRewards object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerPurchaseGoogleplaystoreResponseRewardsWithDefaults

`func NewPlayerPurchaseGoogleplaystoreResponseRewardsWithDefaults() *PlayerPurchaseGoogleplaystoreResponseRewards`

NewPlayerPurchaseGoogleplaystoreResponseRewardsWithDefaults instantiates a new PlayerPurchaseGoogleplaystoreResponseRewards object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *PlayerPurchaseGoogleplaystoreResponseRewards) GetCurrency() []CurrencyExchangeItem`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PlayerPurchaseGoogleplaystoreResponseRewards) GetCurrencyOk() (*[]CurrencyExchangeItem, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PlayerPurchaseGoogleplaystoreResponseRewards) SetCurrency(v []CurrencyExchangeItem)`

SetCurrency sets Currency field to given value.


### GetInventory

`func (o *PlayerPurchaseGoogleplaystoreResponseRewards) GetInventory() []InventoryExchangeItem`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *PlayerPurchaseGoogleplaystoreResponseRewards) GetInventoryOk() (*[]InventoryExchangeItem, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *PlayerPurchaseGoogleplaystoreResponseRewards) SetInventory(v []InventoryExchangeItem)`

SetInventory sets Inventory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


