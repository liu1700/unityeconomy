# ValidationErrorBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** |  | 
**Messages** | **[]string** |  | 

## Methods

### NewValidationErrorBody

`func NewValidationErrorBody(field string, messages []string, ) *ValidationErrorBody`

NewValidationErrorBody instantiates a new ValidationErrorBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationErrorBodyWithDefaults

`func NewValidationErrorBodyWithDefaults() *ValidationErrorBody`

NewValidationErrorBodyWithDefaults instantiates a new ValidationErrorBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *ValidationErrorBody) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ValidationErrorBody) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ValidationErrorBody) SetField(v string)`

SetField sets Field field to given value.


### GetMessages

`func (o *ValidationErrorBody) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ValidationErrorBody) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ValidationErrorBody) SetMessages(v []string)`

SetMessages sets Messages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


