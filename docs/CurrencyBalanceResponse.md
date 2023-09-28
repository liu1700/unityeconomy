# CurrencyBalanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyId** | **string** | Resource ID of the currency. | 
**Balance** | **int64** | The player&#39;s balance. | 
**WriteLock** | **string** | The write lock for the currency balance. | 
**Created** | [**ModifiedMetadata**](ModifiedMetadata.md) |  | 
**Modified** | [**ModifiedMetadata**](ModifiedMetadata.md) |  | 

## Methods

### NewCurrencyBalanceResponse

`func NewCurrencyBalanceResponse(currencyId string, balance int64, writeLock string, created ModifiedMetadata, modified ModifiedMetadata, ) *CurrencyBalanceResponse`

NewCurrencyBalanceResponse instantiates a new CurrencyBalanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyBalanceResponseWithDefaults

`func NewCurrencyBalanceResponseWithDefaults() *CurrencyBalanceResponse`

NewCurrencyBalanceResponseWithDefaults instantiates a new CurrencyBalanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyId

`func (o *CurrencyBalanceResponse) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *CurrencyBalanceResponse) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *CurrencyBalanceResponse) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.


### GetBalance

`func (o *CurrencyBalanceResponse) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *CurrencyBalanceResponse) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *CurrencyBalanceResponse) SetBalance(v int64)`

SetBalance sets Balance field to given value.


### GetWriteLock

`func (o *CurrencyBalanceResponse) GetWriteLock() string`

GetWriteLock returns the WriteLock field if non-nil, zero value otherwise.

### GetWriteLockOk

`func (o *CurrencyBalanceResponse) GetWriteLockOk() (*string, bool)`

GetWriteLockOk returns a tuple with the WriteLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteLock

`func (o *CurrencyBalanceResponse) SetWriteLock(v string)`

SetWriteLock sets WriteLock field to given value.


### GetCreated

`func (o *CurrencyBalanceResponse) GetCreated() ModifiedMetadata`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CurrencyBalanceResponse) GetCreatedOk() (*ModifiedMetadata, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CurrencyBalanceResponse) SetCreated(v ModifiedMetadata)`

SetCreated sets Created field to given value.


### GetModified

`func (o *CurrencyBalanceResponse) GetModified() ModifiedMetadata`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *CurrencyBalanceResponse) GetModifiedOk() (*ModifiedMetadata, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *CurrencyBalanceResponse) SetModified(v ModifiedMetadata)`

SetModified sets Modified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


