# CurrencyModifyBalanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyId** | Pointer to **string** | Resource ID of the currency. | [optional] 
**Amount** | **int64** | The value by which to increment or decrement. Zero is allowed but results in no change to the currency balance. | 
**WriteLock** | Pointer to **string** | The write lock for the currency balance. | [optional] 

## Methods

### NewCurrencyModifyBalanceRequest

`func NewCurrencyModifyBalanceRequest(amount int64, ) *CurrencyModifyBalanceRequest`

NewCurrencyModifyBalanceRequest instantiates a new CurrencyModifyBalanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyModifyBalanceRequestWithDefaults

`func NewCurrencyModifyBalanceRequestWithDefaults() *CurrencyModifyBalanceRequest`

NewCurrencyModifyBalanceRequestWithDefaults instantiates a new CurrencyModifyBalanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyId

`func (o *CurrencyModifyBalanceRequest) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *CurrencyModifyBalanceRequest) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *CurrencyModifyBalanceRequest) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *CurrencyModifyBalanceRequest) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetAmount

`func (o *CurrencyModifyBalanceRequest) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CurrencyModifyBalanceRequest) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CurrencyModifyBalanceRequest) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetWriteLock

`func (o *CurrencyModifyBalanceRequest) GetWriteLock() string`

GetWriteLock returns the WriteLock field if non-nil, zero value otherwise.

### GetWriteLockOk

`func (o *CurrencyModifyBalanceRequest) GetWriteLockOk() (*string, bool)`

GetWriteLockOk returns a tuple with the WriteLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteLock

`func (o *CurrencyModifyBalanceRequest) SetWriteLock(v string)`

SetWriteLock sets WriteLock field to given value.

### HasWriteLock

`func (o *CurrencyModifyBalanceRequest) HasWriteLock() bool`

HasWriteLock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


