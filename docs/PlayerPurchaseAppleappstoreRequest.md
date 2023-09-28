# PlayerPurchaseAppleappstoreRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the purchase. | 
**Receipt** | **string** | Receipt data returned from the Apple App Store as a result of a successful purchase. This should be base64 encoded. | 
**LocalCost** | **int32** | The cost of the purchase as an integer in the minor currency format, for example, $1.99 USD would be 199. | 
**LocalCurrency** | **string** | The ISO-4217 currency code with which the player purchased the IAP. | 

## Methods

### NewPlayerPurchaseAppleappstoreRequest

`func NewPlayerPurchaseAppleappstoreRequest(id string, receipt string, localCost int32, localCurrency string, ) *PlayerPurchaseAppleappstoreRequest`

NewPlayerPurchaseAppleappstoreRequest instantiates a new PlayerPurchaseAppleappstoreRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerPurchaseAppleappstoreRequestWithDefaults

`func NewPlayerPurchaseAppleappstoreRequestWithDefaults() *PlayerPurchaseAppleappstoreRequest`

NewPlayerPurchaseAppleappstoreRequestWithDefaults instantiates a new PlayerPurchaseAppleappstoreRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlayerPurchaseAppleappstoreRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlayerPurchaseAppleappstoreRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlayerPurchaseAppleappstoreRequest) SetId(v string)`

SetId sets Id field to given value.


### GetReceipt

`func (o *PlayerPurchaseAppleappstoreRequest) GetReceipt() string`

GetReceipt returns the Receipt field if non-nil, zero value otherwise.

### GetReceiptOk

`func (o *PlayerPurchaseAppleappstoreRequest) GetReceiptOk() (*string, bool)`

GetReceiptOk returns a tuple with the Receipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceipt

`func (o *PlayerPurchaseAppleappstoreRequest) SetReceipt(v string)`

SetReceipt sets Receipt field to given value.


### GetLocalCost

`func (o *PlayerPurchaseAppleappstoreRequest) GetLocalCost() int32`

GetLocalCost returns the LocalCost field if non-nil, zero value otherwise.

### GetLocalCostOk

`func (o *PlayerPurchaseAppleappstoreRequest) GetLocalCostOk() (*int32, bool)`

GetLocalCostOk returns a tuple with the LocalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCost

`func (o *PlayerPurchaseAppleappstoreRequest) SetLocalCost(v int32)`

SetLocalCost sets LocalCost field to given value.


### GetLocalCurrency

`func (o *PlayerPurchaseAppleappstoreRequest) GetLocalCurrency() string`

GetLocalCurrency returns the LocalCurrency field if non-nil, zero value otherwise.

### GetLocalCurrencyOk

`func (o *PlayerPurchaseAppleappstoreRequest) GetLocalCurrencyOk() (*string, bool)`

GetLocalCurrencyOk returns a tuple with the LocalCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCurrency

`func (o *PlayerPurchaseAppleappstoreRequest) SetLocalCurrency(v string)`

SetLocalCurrency sets LocalCurrency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


