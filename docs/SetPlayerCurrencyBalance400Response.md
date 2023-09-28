# SetPlayerCurrencyBalance400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Title** | **string** |  | 
**Status** | **int32** |  | 
**Code** | **int32** |  | 
**Detail** | **string** |  | 
**Instance** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **[]map[string]interface{}** | Machine readable service specific errors. | [optional] 
**Errors** | [**[]ValidationErrorBody**](ValidationErrorBody.md) |  | 

## Methods

### NewSetPlayerCurrencyBalance400Response

`func NewSetPlayerCurrencyBalance400Response(type_ string, title string, status int32, code int32, detail string, errors []ValidationErrorBody, ) *SetPlayerCurrencyBalance400Response`

NewSetPlayerCurrencyBalance400Response instantiates a new SetPlayerCurrencyBalance400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetPlayerCurrencyBalance400ResponseWithDefaults

`func NewSetPlayerCurrencyBalance400ResponseWithDefaults() *SetPlayerCurrencyBalance400Response`

NewSetPlayerCurrencyBalance400ResponseWithDefaults instantiates a new SetPlayerCurrencyBalance400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SetPlayerCurrencyBalance400Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SetPlayerCurrencyBalance400Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SetPlayerCurrencyBalance400Response) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *SetPlayerCurrencyBalance400Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SetPlayerCurrencyBalance400Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SetPlayerCurrencyBalance400Response) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetStatus

`func (o *SetPlayerCurrencyBalance400Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SetPlayerCurrencyBalance400Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SetPlayerCurrencyBalance400Response) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetCode

`func (o *SetPlayerCurrencyBalance400Response) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SetPlayerCurrencyBalance400Response) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SetPlayerCurrencyBalance400Response) SetCode(v int32)`

SetCode sets Code field to given value.


### GetDetail

`func (o *SetPlayerCurrencyBalance400Response) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *SetPlayerCurrencyBalance400Response) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *SetPlayerCurrencyBalance400Response) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetInstance

`func (o *SetPlayerCurrencyBalance400Response) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *SetPlayerCurrencyBalance400Response) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *SetPlayerCurrencyBalance400Response) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *SetPlayerCurrencyBalance400Response) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetDetails

`func (o *SetPlayerCurrencyBalance400Response) GetDetails() []map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *SetPlayerCurrencyBalance400Response) GetDetailsOk() (*[]map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *SetPlayerCurrencyBalance400Response) SetDetails(v []map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *SetPlayerCurrencyBalance400Response) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetErrors

`func (o *SetPlayerCurrencyBalance400Response) GetErrors() []ValidationErrorBody`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SetPlayerCurrencyBalance400Response) GetErrorsOk() (*[]ValidationErrorBody, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SetPlayerCurrencyBalance400Response) SetErrors(v []ValidationErrorBody)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


