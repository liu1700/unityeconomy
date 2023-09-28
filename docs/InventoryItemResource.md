# InventoryItemResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier for the resource. | 
**Name** | **string** | Name of the resource. | 
**Type** | **string** | Type of the resource, for example &#x60;CURRENCY&#x60;, &#x60;INVENTORY_ITEM&#x60;, &#x60;VIRTUAL_PURCHASE&#x60;, &#x60;MONEY_PURCHASE&#x60;. | 
**Created** | [**ModifiedMetadata**](ModifiedMetadata.md) |  | 
**Modified** | [**ModifiedMetadata**](ModifiedMetadata.md) |  | 
**CustomData** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewInventoryItemResource

`func NewInventoryItemResource(id string, name string, type_ string, created ModifiedMetadata, modified ModifiedMetadata, ) *InventoryItemResource`

NewInventoryItemResource instantiates a new InventoryItemResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryItemResourceWithDefaults

`func NewInventoryItemResourceWithDefaults() *InventoryItemResource`

NewInventoryItemResourceWithDefaults instantiates a new InventoryItemResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InventoryItemResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InventoryItemResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InventoryItemResource) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *InventoryItemResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InventoryItemResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InventoryItemResource) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *InventoryItemResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InventoryItemResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InventoryItemResource) SetType(v string)`

SetType sets Type field to given value.


### GetCreated

`func (o *InventoryItemResource) GetCreated() ModifiedMetadata`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *InventoryItemResource) GetCreatedOk() (*ModifiedMetadata, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *InventoryItemResource) SetCreated(v ModifiedMetadata)`

SetCreated sets Created field to given value.


### GetModified

`func (o *InventoryItemResource) GetModified() ModifiedMetadata`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *InventoryItemResource) GetModifiedOk() (*ModifiedMetadata, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *InventoryItemResource) SetModified(v ModifiedMetadata)`

SetModified sets Modified field to given value.


### GetCustomData

`func (o *InventoryItemResource) GetCustomData() map[string]interface{}`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *InventoryItemResource) GetCustomDataOk() (*map[string]interface{}, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *InventoryItemResource) SetCustomData(v map[string]interface{})`

SetCustomData sets CustomData field to given value.

### HasCustomData

`func (o *InventoryItemResource) HasCustomData() bool`

HasCustomData returns a boolean if a field has been set.

### SetCustomDataNil

`func (o *InventoryItemResource) SetCustomDataNil(b bool)`

 SetCustomDataNil sets the value for CustomData to be an explicit nil

### UnsetCustomData
`func (o *InventoryItemResource) UnsetCustomData()`

UnsetCustomData ensures that no value is present for CustomData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


