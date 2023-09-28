# ErrorResponseConflictInventoryUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempted** | [**AddInventoryRequest**](AddInventoryRequest.md) |  | 
**Existing** | [**NullableInventoryResponse**](InventoryResponse.md) |  | 

## Methods

### NewErrorResponseConflictInventoryUpdateData

`func NewErrorResponseConflictInventoryUpdateData(attempted AddInventoryRequest, existing NullableInventoryResponse, ) *ErrorResponseConflictInventoryUpdateData`

NewErrorResponseConflictInventoryUpdateData instantiates a new ErrorResponseConflictInventoryUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseConflictInventoryUpdateDataWithDefaults

`func NewErrorResponseConflictInventoryUpdateDataWithDefaults() *ErrorResponseConflictInventoryUpdateData`

NewErrorResponseConflictInventoryUpdateDataWithDefaults instantiates a new ErrorResponseConflictInventoryUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempted

`func (o *ErrorResponseConflictInventoryUpdateData) GetAttempted() AddInventoryRequest`

GetAttempted returns the Attempted field if non-nil, zero value otherwise.

### GetAttemptedOk

`func (o *ErrorResponseConflictInventoryUpdateData) GetAttemptedOk() (*AddInventoryRequest, bool)`

GetAttemptedOk returns a tuple with the Attempted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempted

`func (o *ErrorResponseConflictInventoryUpdateData) SetAttempted(v AddInventoryRequest)`

SetAttempted sets Attempted field to given value.


### GetExisting

`func (o *ErrorResponseConflictInventoryUpdateData) GetExisting() InventoryResponse`

GetExisting returns the Existing field if non-nil, zero value otherwise.

### GetExistingOk

`func (o *ErrorResponseConflictInventoryUpdateData) GetExistingOk() (*InventoryResponse, bool)`

GetExistingOk returns a tuple with the Existing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExisting

`func (o *ErrorResponseConflictInventoryUpdateData) SetExisting(v InventoryResponse)`

SetExisting sets Existing field to given value.


### SetExistingNil

`func (o *ErrorResponseConflictInventoryUpdateData) SetExistingNil(b bool)`

 SetExistingNil sets the value for Existing to be an explicit nil

### UnsetExisting
`func (o *ErrorResponseConflictInventoryUpdateData) UnsetExisting()`

UnsetExisting ensures that no value is present for Existing, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


