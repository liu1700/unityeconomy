# PlayerPurchaseGoogleplaystoreRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the purchase. | 
**PurchaseData** | **string** | A JSON encoded string returned from a successful in-app billing purchase. | 
**PurchaseDataSignature** | **string** | A signature of the &#x60;purchaseData&#x60; returned from a successful in-app billing purchase. | 
**LocalCost** | **int32** | The cost of the purchase as an integer in the minor currency format, for example, $1.99 USD would be 199. | 
**LocalCurrency** | **string** | The ISO-4217 currency code with which the player purchased the IAP. | 

## Methods

### NewPlayerPurchaseGoogleplaystoreRequest

`func NewPlayerPurchaseGoogleplaystoreRequest(id string, purchaseData string, purchaseDataSignature string, localCost int32, localCurrency string, ) *PlayerPurchaseGoogleplaystoreRequest`

NewPlayerPurchaseGoogleplaystoreRequest instantiates a new PlayerPurchaseGoogleplaystoreRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerPurchaseGoogleplaystoreRequestWithDefaults

`func NewPlayerPurchaseGoogleplaystoreRequestWithDefaults() *PlayerPurchaseGoogleplaystoreRequest`

NewPlayerPurchaseGoogleplaystoreRequestWithDefaults instantiates a new PlayerPurchaseGoogleplaystoreRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlayerPurchaseGoogleplaystoreRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlayerPurchaseGoogleplaystoreRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlayerPurchaseGoogleplaystoreRequest) SetId(v string)`

SetId sets Id field to given value.


### GetPurchaseData

`func (o *PlayerPurchaseGoogleplaystoreRequest) GetPurchaseData() string`

GetPurchaseData returns the PurchaseData field if non-nil, zero value otherwise.

### GetPurchaseDataOk

`func (o *PlayerPurchaseGoogleplaystoreRequest) GetPurchaseDataOk() (*string, bool)`

GetPurchaseDataOk returns a tuple with the PurchaseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseData

`func (o *PlayerPurchaseGoogleplaystoreRequest) SetPurchaseData(v string)`

SetPurchaseData sets PurchaseData field to given value.


### GetPurchaseDataSignature

`func (o *PlayerPurchaseGoogleplaystoreRequest) GetPurchaseDataSignature() string`

GetPurchaseDataSignature returns the PurchaseDataSignature field if non-nil, zero value otherwise.

### GetPurchaseDataSignatureOk

`func (o *PlayerPurchaseGoogleplaystoreRequest) GetPurchaseDataSignatureOk() (*string, bool)`

GetPurchaseDataSignatureOk returns a tuple with the PurchaseDataSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDataSignature

`func (o *PlayerPurchaseGoogleplaystoreRequest) SetPurchaseDataSignature(v string)`

SetPurchaseDataSignature sets PurchaseDataSignature field to given value.


### GetLocalCost

`func (o *PlayerPurchaseGoogleplaystoreRequest) GetLocalCost() int32`

GetLocalCost returns the LocalCost field if non-nil, zero value otherwise.

### GetLocalCostOk

`func (o *PlayerPurchaseGoogleplaystoreRequest) GetLocalCostOk() (*int32, bool)`

GetLocalCostOk returns a tuple with the LocalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCost

`func (o *PlayerPurchaseGoogleplaystoreRequest) SetLocalCost(v int32)`

SetLocalCost sets LocalCost field to given value.


### GetLocalCurrency

`func (o *PlayerPurchaseGoogleplaystoreRequest) GetLocalCurrency() string`

GetLocalCurrency returns the LocalCurrency field if non-nil, zero value otherwise.

### GetLocalCurrencyOk

`func (o *PlayerPurchaseGoogleplaystoreRequest) GetLocalCurrencyOk() (*string, bool)`

GetLocalCurrencyOk returns a tuple with the LocalCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCurrency

`func (o *PlayerPurchaseGoogleplaystoreRequest) SetLocalCurrency(v string)`

SetLocalCurrency sets LocalCurrency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


