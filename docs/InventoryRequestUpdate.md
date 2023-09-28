# InventoryRequestUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceData** | **map[string]interface{}** | Instance data to be saved against the inventory item. Max size when serialized 5 KB. | 
**WriteLock** | Pointer to **string** | The write lock for the inventory item instance. | [optional] 

## Methods

### NewInventoryRequestUpdate

`func NewInventoryRequestUpdate(instanceData map[string]interface{}, ) *InventoryRequestUpdate`

NewInventoryRequestUpdate instantiates a new InventoryRequestUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryRequestUpdateWithDefaults

`func NewInventoryRequestUpdateWithDefaults() *InventoryRequestUpdate`

NewInventoryRequestUpdateWithDefaults instantiates a new InventoryRequestUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceData

`func (o *InventoryRequestUpdate) GetInstanceData() map[string]interface{}`

GetInstanceData returns the InstanceData field if non-nil, zero value otherwise.

### GetInstanceDataOk

`func (o *InventoryRequestUpdate) GetInstanceDataOk() (*map[string]interface{}, bool)`

GetInstanceDataOk returns a tuple with the InstanceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceData

`func (o *InventoryRequestUpdate) SetInstanceData(v map[string]interface{})`

SetInstanceData sets InstanceData field to given value.


### SetInstanceDataNil

`func (o *InventoryRequestUpdate) SetInstanceDataNil(b bool)`

 SetInstanceDataNil sets the value for InstanceData to be an explicit nil

### UnsetInstanceData
`func (o *InventoryRequestUpdate) UnsetInstanceData()`

UnsetInstanceData ensures that no value is present for InstanceData, not even an explicit nil
### GetWriteLock

`func (o *InventoryRequestUpdate) GetWriteLock() string`

GetWriteLock returns the WriteLock field if non-nil, zero value otherwise.

### GetWriteLockOk

`func (o *InventoryRequestUpdate) GetWriteLockOk() (*string, bool)`

GetWriteLockOk returns a tuple with the WriteLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteLock

`func (o *InventoryRequestUpdate) SetWriteLock(v string)`

SetWriteLock sets WriteLock field to given value.

### HasWriteLock

`func (o *InventoryRequestUpdate) HasWriteLock() bool`

HasWriteLock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


