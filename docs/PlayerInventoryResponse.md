# PlayerInventoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]InventoryResponse**](InventoryResponse.md) | List of player&#39;s inventory items. | 
**Links** | [**PlayerCurrencyBalanceResponseLinks**](PlayerCurrencyBalanceResponseLinks.md) |  | 

## Methods

### NewPlayerInventoryResponse

`func NewPlayerInventoryResponse(results []InventoryResponse, links PlayerCurrencyBalanceResponseLinks, ) *PlayerInventoryResponse`

NewPlayerInventoryResponse instantiates a new PlayerInventoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerInventoryResponseWithDefaults

`func NewPlayerInventoryResponseWithDefaults() *PlayerInventoryResponse`

NewPlayerInventoryResponseWithDefaults instantiates a new PlayerInventoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *PlayerInventoryResponse) GetResults() []InventoryResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PlayerInventoryResponse) GetResultsOk() (*[]InventoryResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PlayerInventoryResponse) SetResults(v []InventoryResponse)`

SetResults sets Results field to given value.


### GetLinks

`func (o *PlayerInventoryResponse) GetLinks() PlayerCurrencyBalanceResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PlayerInventoryResponse) GetLinksOk() (*PlayerCurrencyBalanceResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PlayerInventoryResponse) SetLinks(v PlayerCurrencyBalanceResponseLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


