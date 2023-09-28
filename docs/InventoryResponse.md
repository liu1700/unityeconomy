# InventoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlayersInventoryItemId** | **string** | ID of the player&#39;s inventory item. | 
**InventoryItemId** | **string** | Resource ID of the inventory item configuration associated with this instance. | 
**InstanceData** | Pointer to **map[string]interface{}** | Instance data. Max size when serialized 5 KB. | [optional] 
**WriteLock** | **string** | The write lock for the inventory item instance. | 
**Created** | [**ModifiedMetadata**](ModifiedMetadata.md) |  | 
**Modified** | [**ModifiedMetadata**](ModifiedMetadata.md) |  | 

## Methods

### NewInventoryResponse

`func NewInventoryResponse(playersInventoryItemId string, inventoryItemId string, writeLock string, created ModifiedMetadata, modified ModifiedMetadata, ) *InventoryResponse`

NewInventoryResponse instantiates a new InventoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryResponseWithDefaults

`func NewInventoryResponseWithDefaults() *InventoryResponse`

NewInventoryResponseWithDefaults instantiates a new InventoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayersInventoryItemId

`func (o *InventoryResponse) GetPlayersInventoryItemId() string`

GetPlayersInventoryItemId returns the PlayersInventoryItemId field if non-nil, zero value otherwise.

### GetPlayersInventoryItemIdOk

`func (o *InventoryResponse) GetPlayersInventoryItemIdOk() (*string, bool)`

GetPlayersInventoryItemIdOk returns a tuple with the PlayersInventoryItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayersInventoryItemId

`func (o *InventoryResponse) SetPlayersInventoryItemId(v string)`

SetPlayersInventoryItemId sets PlayersInventoryItemId field to given value.


### GetInventoryItemId

`func (o *InventoryResponse) GetInventoryItemId() string`

GetInventoryItemId returns the InventoryItemId field if non-nil, zero value otherwise.

### GetInventoryItemIdOk

`func (o *InventoryResponse) GetInventoryItemIdOk() (*string, bool)`

GetInventoryItemIdOk returns a tuple with the InventoryItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryItemId

`func (o *InventoryResponse) SetInventoryItemId(v string)`

SetInventoryItemId sets InventoryItemId field to given value.


### GetInstanceData

`func (o *InventoryResponse) GetInstanceData() map[string]interface{}`

GetInstanceData returns the InstanceData field if non-nil, zero value otherwise.

### GetInstanceDataOk

`func (o *InventoryResponse) GetInstanceDataOk() (*map[string]interface{}, bool)`

GetInstanceDataOk returns a tuple with the InstanceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceData

`func (o *InventoryResponse) SetInstanceData(v map[string]interface{})`

SetInstanceData sets InstanceData field to given value.

### HasInstanceData

`func (o *InventoryResponse) HasInstanceData() bool`

HasInstanceData returns a boolean if a field has been set.

### SetInstanceDataNil

`func (o *InventoryResponse) SetInstanceDataNil(b bool)`

 SetInstanceDataNil sets the value for InstanceData to be an explicit nil

### UnsetInstanceData
`func (o *InventoryResponse) UnsetInstanceData()`

UnsetInstanceData ensures that no value is present for InstanceData, not even an explicit nil
### GetWriteLock

`func (o *InventoryResponse) GetWriteLock() string`

GetWriteLock returns the WriteLock field if non-nil, zero value otherwise.

### GetWriteLockOk

`func (o *InventoryResponse) GetWriteLockOk() (*string, bool)`

GetWriteLockOk returns a tuple with the WriteLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteLock

`func (o *InventoryResponse) SetWriteLock(v string)`

SetWriteLock sets WriteLock field to given value.


### GetCreated

`func (o *InventoryResponse) GetCreated() ModifiedMetadata`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *InventoryResponse) GetCreatedOk() (*ModifiedMetadata, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *InventoryResponse) SetCreated(v ModifiedMetadata)`

SetCreated sets Created field to given value.


### GetModified

`func (o *InventoryResponse) GetModified() ModifiedMetadata`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *InventoryResponse) GetModifiedOk() (*ModifiedMetadata, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *InventoryResponse) SetModified(v ModifiedMetadata)`

SetModified sets Modified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


