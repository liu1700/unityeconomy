# CurrencyBalanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyId** | Pointer to **string** | Resource ID of the currency. | [optional] 
**Balance** | **int64** | The player&#39;s balance. | 
**WriteLock** | Pointer to **string** | The write lock for the currency balance. | [optional] 

## Methods

### NewCurrencyBalanceRequest

`func NewCurrencyBalanceRequest(balance int64, ) *CurrencyBalanceRequest`

NewCurrencyBalanceRequest instantiates a new CurrencyBalanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyBalanceRequestWithDefaults

`func NewCurrencyBalanceRequestWithDefaults() *CurrencyBalanceRequest`

NewCurrencyBalanceRequestWithDefaults instantiates a new CurrencyBalanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyId

`func (o *CurrencyBalanceRequest) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *CurrencyBalanceRequest) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *CurrencyBalanceRequest) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *CurrencyBalanceRequest) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetBalance

`func (o *CurrencyBalanceRequest) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *CurrencyBalanceRequest) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *CurrencyBalanceRequest) SetBalance(v int64)`

SetBalance sets Balance field to given value.


### GetWriteLock

`func (o *CurrencyBalanceRequest) GetWriteLock() string`

GetWriteLock returns the WriteLock field if non-nil, zero value otherwise.

### GetWriteLockOk

`func (o *CurrencyBalanceRequest) GetWriteLockOk() (*string, bool)`

GetWriteLockOk returns a tuple with the WriteLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteLock

`func (o *CurrencyBalanceRequest) SetWriteLock(v string)`

SetWriteLock sets WriteLock field to given value.

### HasWriteLock

`func (o *CurrencyBalanceRequest) HasWriteLock() bool`

HasWriteLock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


