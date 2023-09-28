# ErrorResponseConflictCurrencyBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Title** | **string** |  | 
**Status** | **int32** | For example 409 | 
**Code** | **int32** | For example 10204 | 
**Detail** | **string** |  | 
**Instance** | Pointer to **string** |  | [optional] 
**Data** | [**ErrorResponseConflictCurrencyBalanceData**](ErrorResponseConflictCurrencyBalanceData.md) |  | 

## Methods

### NewErrorResponseConflictCurrencyBalance

`func NewErrorResponseConflictCurrencyBalance(type_ string, title string, status int32, code int32, detail string, data ErrorResponseConflictCurrencyBalanceData, ) *ErrorResponseConflictCurrencyBalance`

NewErrorResponseConflictCurrencyBalance instantiates a new ErrorResponseConflictCurrencyBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseConflictCurrencyBalanceWithDefaults

`func NewErrorResponseConflictCurrencyBalanceWithDefaults() *ErrorResponseConflictCurrencyBalance`

NewErrorResponseConflictCurrencyBalanceWithDefaults instantiates a new ErrorResponseConflictCurrencyBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ErrorResponseConflictCurrencyBalance) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorResponseConflictCurrencyBalance) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorResponseConflictCurrencyBalance) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *ErrorResponseConflictCurrencyBalance) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ErrorResponseConflictCurrencyBalance) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ErrorResponseConflictCurrencyBalance) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetStatus

`func (o *ErrorResponseConflictCurrencyBalance) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ErrorResponseConflictCurrencyBalance) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ErrorResponseConflictCurrencyBalance) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetCode

`func (o *ErrorResponseConflictCurrencyBalance) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorResponseConflictCurrencyBalance) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorResponseConflictCurrencyBalance) SetCode(v int32)`

SetCode sets Code field to given value.


### GetDetail

`func (o *ErrorResponseConflictCurrencyBalance) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ErrorResponseConflictCurrencyBalance) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ErrorResponseConflictCurrencyBalance) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetInstance

`func (o *ErrorResponseConflictCurrencyBalance) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ErrorResponseConflictCurrencyBalance) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ErrorResponseConflictCurrencyBalance) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ErrorResponseConflictCurrencyBalance) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetData

`func (o *ErrorResponseConflictCurrencyBalance) GetData() ErrorResponseConflictCurrencyBalanceData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ErrorResponseConflictCurrencyBalance) GetDataOk() (*ErrorResponseConflictCurrencyBalanceData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ErrorResponseConflictCurrencyBalance) SetData(v ErrorResponseConflictCurrencyBalanceData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


