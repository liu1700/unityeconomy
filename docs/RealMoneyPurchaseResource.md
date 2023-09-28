# RealMoneyPurchaseResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the resource. | 
**Name** | **string** | A descriptive name for the resource. | 
**Type** | **string** | Type of the item, for example &#x60;CURRENCY&#x60;, &#x60;INVENTORY_ITEM&#x60;, &#x60;VIRTUAL_PURCHASE&#x60;, &#x60;MONEY_PURCHASE&#x60;. | 
**StoreIdentifiers** | [**RealMoneyPurchaseResourceStoreIdentifiers**](RealMoneyPurchaseResourceStoreIdentifiers.md) |  | 
**Rewards** | [**[]Reward**](Reward.md) | The rewards credited to the player when making the purchase. A reward is composed of the ID of a currency or inventory item, the amount of that currency or item, and the default instance data (for inventory items). | 
**CustomData** | **map[string]interface{}** | Max size when serialised 5 KB. | 
**Created** | [**ModifiedMetadata**](ModifiedMetadata.md) |  | 
**Modified** | [**ModifiedMetadata**](ModifiedMetadata.md) |  | 

## Methods

### NewRealMoneyPurchaseResource

`func NewRealMoneyPurchaseResource(id string, name string, type_ string, storeIdentifiers RealMoneyPurchaseResourceStoreIdentifiers, rewards []Reward, customData map[string]interface{}, created ModifiedMetadata, modified ModifiedMetadata, ) *RealMoneyPurchaseResource`

NewRealMoneyPurchaseResource instantiates a new RealMoneyPurchaseResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealMoneyPurchaseResourceWithDefaults

`func NewRealMoneyPurchaseResourceWithDefaults() *RealMoneyPurchaseResource`

NewRealMoneyPurchaseResourceWithDefaults instantiates a new RealMoneyPurchaseResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RealMoneyPurchaseResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RealMoneyPurchaseResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RealMoneyPurchaseResource) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *RealMoneyPurchaseResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RealMoneyPurchaseResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RealMoneyPurchaseResource) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *RealMoneyPurchaseResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealMoneyPurchaseResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealMoneyPurchaseResource) SetType(v string)`

SetType sets Type field to given value.


### GetStoreIdentifiers

`func (o *RealMoneyPurchaseResource) GetStoreIdentifiers() RealMoneyPurchaseResourceStoreIdentifiers`

GetStoreIdentifiers returns the StoreIdentifiers field if non-nil, zero value otherwise.

### GetStoreIdentifiersOk

`func (o *RealMoneyPurchaseResource) GetStoreIdentifiersOk() (*RealMoneyPurchaseResourceStoreIdentifiers, bool)`

GetStoreIdentifiersOk returns a tuple with the StoreIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreIdentifiers

`func (o *RealMoneyPurchaseResource) SetStoreIdentifiers(v RealMoneyPurchaseResourceStoreIdentifiers)`

SetStoreIdentifiers sets StoreIdentifiers field to given value.


### GetRewards

`func (o *RealMoneyPurchaseResource) GetRewards() []Reward`

GetRewards returns the Rewards field if non-nil, zero value otherwise.

### GetRewardsOk

`func (o *RealMoneyPurchaseResource) GetRewardsOk() (*[]Reward, bool)`

GetRewardsOk returns a tuple with the Rewards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewards

`func (o *RealMoneyPurchaseResource) SetRewards(v []Reward)`

SetRewards sets Rewards field to given value.


### GetCustomData

`func (o *RealMoneyPurchaseResource) GetCustomData() map[string]interface{}`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *RealMoneyPurchaseResource) GetCustomDataOk() (*map[string]interface{}, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *RealMoneyPurchaseResource) SetCustomData(v map[string]interface{})`

SetCustomData sets CustomData field to given value.


### SetCustomDataNil

`func (o *RealMoneyPurchaseResource) SetCustomDataNil(b bool)`

 SetCustomDataNil sets the value for CustomData to be an explicit nil

### UnsetCustomData
`func (o *RealMoneyPurchaseResource) UnsetCustomData()`

UnsetCustomData ensures that no value is present for CustomData, not even an explicit nil
### GetCreated

`func (o *RealMoneyPurchaseResource) GetCreated() ModifiedMetadata`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RealMoneyPurchaseResource) GetCreatedOk() (*ModifiedMetadata, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RealMoneyPurchaseResource) SetCreated(v ModifiedMetadata)`

SetCreated sets Created field to given value.


### GetModified

`func (o *RealMoneyPurchaseResource) GetModified() ModifiedMetadata`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *RealMoneyPurchaseResource) GetModifiedOk() (*ModifiedMetadata, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *RealMoneyPurchaseResource) SetModified(v ModifiedMetadata)`

SetModified sets Modified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


