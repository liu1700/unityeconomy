# CurrencyExchangeItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the currency. | 
**Amount** | **int64** | Amount of currency added or deducted. | 

## Methods

### NewCurrencyExchangeItem

`func NewCurrencyExchangeItem(id string, amount int64, ) *CurrencyExchangeItem`

NewCurrencyExchangeItem instantiates a new CurrencyExchangeItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyExchangeItemWithDefaults

`func NewCurrencyExchangeItemWithDefaults() *CurrencyExchangeItem`

NewCurrencyExchangeItemWithDefaults instantiates a new CurrencyExchangeItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CurrencyExchangeItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CurrencyExchangeItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CurrencyExchangeItem) SetId(v string)`

SetId sets Id field to given value.


### GetAmount

`func (o *CurrencyExchangeItem) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CurrencyExchangeItem) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CurrencyExchangeItem) SetAmount(v int64)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


