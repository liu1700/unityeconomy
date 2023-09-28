# InventoryDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WriteLock** | Pointer to **string** | The write lock for the inventory item instance. This property has been deprecated. Please use the &#x60;writeLock&#x60; query parameter instead. | [optional] 

## Methods

### NewInventoryDeleteRequest

`func NewInventoryDeleteRequest() *InventoryDeleteRequest`

NewInventoryDeleteRequest instantiates a new InventoryDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryDeleteRequestWithDefaults

`func NewInventoryDeleteRequestWithDefaults() *InventoryDeleteRequest`

NewInventoryDeleteRequestWithDefaults instantiates a new InventoryDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWriteLock

`func (o *InventoryDeleteRequest) GetWriteLock() string`

GetWriteLock returns the WriteLock field if non-nil, zero value otherwise.

### GetWriteLockOk

`func (o *InventoryDeleteRequest) GetWriteLockOk() (*string, bool)`

GetWriteLockOk returns a tuple with the WriteLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteLock

`func (o *InventoryDeleteRequest) SetWriteLock(v string)`

SetWriteLock sets WriteLock field to given value.

### HasWriteLock

`func (o *InventoryDeleteRequest) HasWriteLock() bool`

HasWriteLock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


