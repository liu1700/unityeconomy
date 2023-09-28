# ModifiedMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **NullableTime** | Date time in ISO 8601 format. &#x60;null&#x60; if there is no associated value. | 

## Methods

### NewModifiedMetadata

`func NewModifiedMetadata(date NullableTime, ) *ModifiedMetadata`

NewModifiedMetadata instantiates a new ModifiedMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifiedMetadataWithDefaults

`func NewModifiedMetadataWithDefaults() *ModifiedMetadata`

NewModifiedMetadataWithDefaults instantiates a new ModifiedMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *ModifiedMetadata) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ModifiedMetadata) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ModifiedMetadata) SetDate(v time.Time)`

SetDate sets Date field to given value.


### SetDateNil

`func (o *ModifiedMetadata) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *ModifiedMetadata) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


