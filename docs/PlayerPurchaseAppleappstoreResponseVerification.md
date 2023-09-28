# PlayerPurchaseAppleappstoreResponseVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Status of the receipt verification.  * &#x60;VALID&#x60;: The purchase was valid.  * &#x60;VALID_NOT_REDEEMED&#x60;: The purchase was valid but seen before, and had not yet been redeemed.  * &#x60;INVALID_ALREADY_REDEEMED&#x60;: The purchase has already been redeemed.  * &#x60;INVALID_VERIFICATION_FAILED&#x60;: The receipt verification Service returned that the receipt data was not valid.  * &#x60;INVALID_ANOTHER_PLAYER&#x60;: The receipt has previously been used by a different player and validated.  * &#x60;INVALID_CONFIGURATION&#x60;: The service configuration is invalid, further information in the details section of the response.  * &#x60;INVALID_PRODUCT_ID_MISMATCH&#x60;: The purchase configuration store product identifier does not match the one in the receipt. * &#x60;CURRENCY_MAX_EXCEEDED&#x60;: Could not add the rewards because one or more currencies would be taken over the specified maximum balance. | 
**Store** | [**PlayerPurchaseAppleappstoreResponseVerificationStore**](PlayerPurchaseAppleappstoreResponseVerificationStore.md) |  | 

## Methods

### NewPlayerPurchaseAppleappstoreResponseVerification

`func NewPlayerPurchaseAppleappstoreResponseVerification(status string, store PlayerPurchaseAppleappstoreResponseVerificationStore, ) *PlayerPurchaseAppleappstoreResponseVerification`

NewPlayerPurchaseAppleappstoreResponseVerification instantiates a new PlayerPurchaseAppleappstoreResponseVerification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerPurchaseAppleappstoreResponseVerificationWithDefaults

`func NewPlayerPurchaseAppleappstoreResponseVerificationWithDefaults() *PlayerPurchaseAppleappstoreResponseVerification`

NewPlayerPurchaseAppleappstoreResponseVerificationWithDefaults instantiates a new PlayerPurchaseAppleappstoreResponseVerification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PlayerPurchaseAppleappstoreResponseVerification) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlayerPurchaseAppleappstoreResponseVerification) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlayerPurchaseAppleappstoreResponseVerification) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStore

`func (o *PlayerPurchaseAppleappstoreResponseVerification) GetStore() PlayerPurchaseAppleappstoreResponseVerificationStore`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *PlayerPurchaseAppleappstoreResponseVerification) GetStoreOk() (*PlayerPurchaseAppleappstoreResponseVerificationStore, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *PlayerPurchaseAppleappstoreResponseVerification) SetStore(v PlayerPurchaseAppleappstoreResponseVerificationStore)`

SetStore sets Store field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


