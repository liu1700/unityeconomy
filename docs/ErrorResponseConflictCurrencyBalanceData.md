# ErrorResponseConflictCurrencyBalanceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempted** | [**CurrencyBalanceRequest**](CurrencyBalanceRequest.md) |  | 
**Existing** | [**NullableCurrencyBalanceResponse**](CurrencyBalanceResponse.md) |  | 

## Methods

### NewErrorResponseConflictCurrencyBalanceData

`func NewErrorResponseConflictCurrencyBalanceData(attempted CurrencyBalanceRequest, existing NullableCurrencyBalanceResponse, ) *ErrorResponseConflictCurrencyBalanceData`

NewErrorResponseConflictCurrencyBalanceData instantiates a new ErrorResponseConflictCurrencyBalanceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseConflictCurrencyBalanceDataWithDefaults

`func NewErrorResponseConflictCurrencyBalanceDataWithDefaults() *ErrorResponseConflictCurrencyBalanceData`

NewErrorResponseConflictCurrencyBalanceDataWithDefaults instantiates a new ErrorResponseConflictCurrencyBalanceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempted

`func (o *ErrorResponseConflictCurrencyBalanceData) GetAttempted() CurrencyBalanceRequest`

GetAttempted returns the Attempted field if non-nil, zero value otherwise.

### GetAttemptedOk

`func (o *ErrorResponseConflictCurrencyBalanceData) GetAttemptedOk() (*CurrencyBalanceRequest, bool)`

GetAttemptedOk returns a tuple with the Attempted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempted

`func (o *ErrorResponseConflictCurrencyBalanceData) SetAttempted(v CurrencyBalanceRequest)`

SetAttempted sets Attempted field to given value.


### GetExisting

`func (o *ErrorResponseConflictCurrencyBalanceData) GetExisting() CurrencyBalanceResponse`

GetExisting returns the Existing field if non-nil, zero value otherwise.

### GetExistingOk

`func (o *ErrorResponseConflictCurrencyBalanceData) GetExistingOk() (*CurrencyBalanceResponse, bool)`

GetExistingOk returns a tuple with the Existing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExisting

`func (o *ErrorResponseConflictCurrencyBalanceData) SetExisting(v CurrencyBalanceResponse)`

SetExisting sets Existing field to given value.


### SetExistingNil

`func (o *ErrorResponseConflictCurrencyBalanceData) SetExistingNil(b bool)`

 SetExistingNil sets the value for Existing to be an explicit nil

### UnsetExisting
`func (o *ErrorResponseConflictCurrencyBalanceData) UnsetExisting()`

UnsetExisting ensures that no value is present for Existing, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


