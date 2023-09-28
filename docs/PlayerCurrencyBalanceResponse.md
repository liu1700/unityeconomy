# PlayerCurrencyBalanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]CurrencyBalanceResponse**](CurrencyBalanceResponse.md) | List of currency balances. | 
**Links** | [**PlayerCurrencyBalanceResponseLinks**](PlayerCurrencyBalanceResponseLinks.md) |  | 

## Methods

### NewPlayerCurrencyBalanceResponse

`func NewPlayerCurrencyBalanceResponse(results []CurrencyBalanceResponse, links PlayerCurrencyBalanceResponseLinks, ) *PlayerCurrencyBalanceResponse`

NewPlayerCurrencyBalanceResponse instantiates a new PlayerCurrencyBalanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerCurrencyBalanceResponseWithDefaults

`func NewPlayerCurrencyBalanceResponseWithDefaults() *PlayerCurrencyBalanceResponse`

NewPlayerCurrencyBalanceResponseWithDefaults instantiates a new PlayerCurrencyBalanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *PlayerCurrencyBalanceResponse) GetResults() []CurrencyBalanceResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PlayerCurrencyBalanceResponse) GetResultsOk() (*[]CurrencyBalanceResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PlayerCurrencyBalanceResponse) SetResults(v []CurrencyBalanceResponse)`

SetResults sets Results field to given value.


### GetLinks

`func (o *PlayerCurrencyBalanceResponse) GetLinks() PlayerCurrencyBalanceResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PlayerCurrencyBalanceResponse) GetLinksOk() (*PlayerCurrencyBalanceResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PlayerCurrencyBalanceResponse) SetLinks(v PlayerCurrencyBalanceResponseLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


