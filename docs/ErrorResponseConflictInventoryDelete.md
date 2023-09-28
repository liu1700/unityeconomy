# ErrorResponseConflictInventoryDelete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Title** | **string** |  | 
**Status** | **int32** | e.g 409 | 
**Code** | **int32** | For example 10204 | 
**Detail** | **string** |  | 
**Instance** | Pointer to **string** |  | [optional] 
**Data** | [**ErrorResponseConflictInventoryDeleteData**](ErrorResponseConflictInventoryDeleteData.md) |  | 

## Methods

### NewErrorResponseConflictInventoryDelete

`func NewErrorResponseConflictInventoryDelete(type_ string, title string, status int32, code int32, detail string, data ErrorResponseConflictInventoryDeleteData, ) *ErrorResponseConflictInventoryDelete`

NewErrorResponseConflictInventoryDelete instantiates a new ErrorResponseConflictInventoryDelete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseConflictInventoryDeleteWithDefaults

`func NewErrorResponseConflictInventoryDeleteWithDefaults() *ErrorResponseConflictInventoryDelete`

NewErrorResponseConflictInventoryDeleteWithDefaults instantiates a new ErrorResponseConflictInventoryDelete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ErrorResponseConflictInventoryDelete) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorResponseConflictInventoryDelete) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorResponseConflictInventoryDelete) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *ErrorResponseConflictInventoryDelete) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ErrorResponseConflictInventoryDelete) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ErrorResponseConflictInventoryDelete) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetStatus

`func (o *ErrorResponseConflictInventoryDelete) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ErrorResponseConflictInventoryDelete) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ErrorResponseConflictInventoryDelete) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetCode

`func (o *ErrorResponseConflictInventoryDelete) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorResponseConflictInventoryDelete) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorResponseConflictInventoryDelete) SetCode(v int32)`

SetCode sets Code field to given value.


### GetDetail

`func (o *ErrorResponseConflictInventoryDelete) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ErrorResponseConflictInventoryDelete) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ErrorResponseConflictInventoryDelete) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetInstance

`func (o *ErrorResponseConflictInventoryDelete) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ErrorResponseConflictInventoryDelete) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ErrorResponseConflictInventoryDelete) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ErrorResponseConflictInventoryDelete) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetData

`func (o *ErrorResponseConflictInventoryDelete) GetData() ErrorResponseConflictInventoryDeleteData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ErrorResponseConflictInventoryDelete) GetDataOk() (*ErrorResponseConflictInventoryDeleteData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ErrorResponseConflictInventoryDelete) SetData(v ErrorResponseConflictInventoryDeleteData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


