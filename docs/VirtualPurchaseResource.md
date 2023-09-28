# VirtualPurchaseResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier for the resource. | 
**Name** | **string** | Name of the resource. | 
**Type** | **string** | Type of the item, for example &#x60;CURRENCY&#x60;, &#x60;INVENTORY_ITEM&#x60;, &#x60;VIRTUAL_PURCHASE&#x60;, &#x60;MONEY_PURCHASE&#x60;. | 
**Created** | [**ModifiedMetadata**](ModifiedMetadata.md) |  | 
**Modified** | [**ModifiedMetadata**](ModifiedMetadata.md) |  | 
**Costs** | Pointer to [**[]Cost**](Cost.md) | The costs deducted from the player when making the purchase. A cost is an ID of a currency or inventory item, plus an amount. | [optional] 
**Rewards** | Pointer to [**[]Reward**](Reward.md) | The rewards credited to the player when making the purchase. A reward is composed of the ID of a currency or inventory item, the amount of that currency or item, and the default instance data (for inventory items). | [optional] 
**CustomData** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewVirtualPurchaseResource

`func NewVirtualPurchaseResource(id string, name string, type_ string, created ModifiedMetadata, modified ModifiedMetadata, ) *VirtualPurchaseResource`

NewVirtualPurchaseResource instantiates a new VirtualPurchaseResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualPurchaseResourceWithDefaults

`func NewVirtualPurchaseResourceWithDefaults() *VirtualPurchaseResource`

NewVirtualPurchaseResourceWithDefaults instantiates a new VirtualPurchaseResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VirtualPurchaseResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualPurchaseResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualPurchaseResource) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *VirtualPurchaseResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualPurchaseResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualPurchaseResource) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *VirtualPurchaseResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VirtualPurchaseResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VirtualPurchaseResource) SetType(v string)`

SetType sets Type field to given value.


### GetCreated

`func (o *VirtualPurchaseResource) GetCreated() ModifiedMetadata`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *VirtualPurchaseResource) GetCreatedOk() (*ModifiedMetadata, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *VirtualPurchaseResource) SetCreated(v ModifiedMetadata)`

SetCreated sets Created field to given value.


### GetModified

`func (o *VirtualPurchaseResource) GetModified() ModifiedMetadata`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *VirtualPurchaseResource) GetModifiedOk() (*ModifiedMetadata, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *VirtualPurchaseResource) SetModified(v ModifiedMetadata)`

SetModified sets Modified field to given value.


### GetCosts

`func (o *VirtualPurchaseResource) GetCosts() []Cost`

GetCosts returns the Costs field if non-nil, zero value otherwise.

### GetCostsOk

`func (o *VirtualPurchaseResource) GetCostsOk() (*[]Cost, bool)`

GetCostsOk returns a tuple with the Costs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosts

`func (o *VirtualPurchaseResource) SetCosts(v []Cost)`

SetCosts sets Costs field to given value.

### HasCosts

`func (o *VirtualPurchaseResource) HasCosts() bool`

HasCosts returns a boolean if a field has been set.

### GetRewards

`func (o *VirtualPurchaseResource) GetRewards() []Reward`

GetRewards returns the Rewards field if non-nil, zero value otherwise.

### GetRewardsOk

`func (o *VirtualPurchaseResource) GetRewardsOk() (*[]Reward, bool)`

GetRewardsOk returns a tuple with the Rewards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewards

`func (o *VirtualPurchaseResource) SetRewards(v []Reward)`

SetRewards sets Rewards field to given value.

### HasRewards

`func (o *VirtualPurchaseResource) HasRewards() bool`

HasRewards returns a boolean if a field has been set.

### GetCustomData

`func (o *VirtualPurchaseResource) GetCustomData() map[string]interface{}`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *VirtualPurchaseResource) GetCustomDataOk() (*map[string]interface{}, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *VirtualPurchaseResource) SetCustomData(v map[string]interface{})`

SetCustomData sets CustomData field to given value.

### HasCustomData

`func (o *VirtualPurchaseResource) HasCustomData() bool`

HasCustomData returns a boolean if a field has been set.

### SetCustomDataNil

`func (o *VirtualPurchaseResource) SetCustomDataNil(b bool)`

 SetCustomDataNil sets the value for CustomData to be an explicit nil

### UnsetCustomData
`func (o *VirtualPurchaseResource) UnsetCustomData()`

UnsetCustomData ensures that no value is present for CustomData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


