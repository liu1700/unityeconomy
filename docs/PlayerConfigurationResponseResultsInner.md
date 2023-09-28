# PlayerConfigurationResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the resource. | 
**Name** | **string** | A descriptive name for the resource. | 
**Type** | **string** | Type of the resource, for example &#x60;CURRENCY&#x60;, &#x60;INVENTORY_ITEM&#x60;, &#x60;VIRTUAL_PURCHASE&#x60;, &#x60;MONEY_PURCHASE&#x60;. | 
**Created** | [**ModifiedMetadata**](ModifiedMetadata.md) |  | 
**Modified** | [**ModifiedMetadata**](ModifiedMetadata.md) |  | 
**Initial** | Pointer to **int32** |  | [optional] 
**Max** | Pointer to **int32** |  | [optional] 
**CustomData** | **map[string]interface{}** | Max size when serialised 5 KB. | 
**Costs** | Pointer to [**[]Cost**](Cost.md) | The costs deducted from the player when making the purchase. A cost is an ID of a currency or inventory item, plus an amount. | [optional] 
**Rewards** | [**[]Reward**](Reward.md) | The rewards credited to the player when making the purchase. A reward is composed of the ID of a currency or inventory item, the amount of that currency or item, and the default instance data (for inventory items). | 
**StoreIdentifiers** | [**RealMoneyPurchaseResourceStoreIdentifiers**](RealMoneyPurchaseResourceStoreIdentifiers.md) |  | 

## Methods

### NewPlayerConfigurationResponseResultsInner

`func NewPlayerConfigurationResponseResultsInner(id string, name string, type_ string, created ModifiedMetadata, modified ModifiedMetadata, customData map[string]interface{}, rewards []Reward, storeIdentifiers RealMoneyPurchaseResourceStoreIdentifiers, ) *PlayerConfigurationResponseResultsInner`

NewPlayerConfigurationResponseResultsInner instantiates a new PlayerConfigurationResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerConfigurationResponseResultsInnerWithDefaults

`func NewPlayerConfigurationResponseResultsInnerWithDefaults() *PlayerConfigurationResponseResultsInner`

NewPlayerConfigurationResponseResultsInnerWithDefaults instantiates a new PlayerConfigurationResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlayerConfigurationResponseResultsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlayerConfigurationResponseResultsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlayerConfigurationResponseResultsInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PlayerConfigurationResponseResultsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlayerConfigurationResponseResultsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlayerConfigurationResponseResultsInner) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *PlayerConfigurationResponseResultsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlayerConfigurationResponseResultsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlayerConfigurationResponseResultsInner) SetType(v string)`

SetType sets Type field to given value.


### GetCreated

`func (o *PlayerConfigurationResponseResultsInner) GetCreated() ModifiedMetadata`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PlayerConfigurationResponseResultsInner) GetCreatedOk() (*ModifiedMetadata, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PlayerConfigurationResponseResultsInner) SetCreated(v ModifiedMetadata)`

SetCreated sets Created field to given value.


### GetModified

`func (o *PlayerConfigurationResponseResultsInner) GetModified() ModifiedMetadata`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *PlayerConfigurationResponseResultsInner) GetModifiedOk() (*ModifiedMetadata, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *PlayerConfigurationResponseResultsInner) SetModified(v ModifiedMetadata)`

SetModified sets Modified field to given value.


### GetInitial

`func (o *PlayerConfigurationResponseResultsInner) GetInitial() int32`

GetInitial returns the Initial field if non-nil, zero value otherwise.

### GetInitialOk

`func (o *PlayerConfigurationResponseResultsInner) GetInitialOk() (*int32, bool)`

GetInitialOk returns a tuple with the Initial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitial

`func (o *PlayerConfigurationResponseResultsInner) SetInitial(v int32)`

SetInitial sets Initial field to given value.

### HasInitial

`func (o *PlayerConfigurationResponseResultsInner) HasInitial() bool`

HasInitial returns a boolean if a field has been set.

### GetMax

`func (o *PlayerConfigurationResponseResultsInner) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *PlayerConfigurationResponseResultsInner) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *PlayerConfigurationResponseResultsInner) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *PlayerConfigurationResponseResultsInner) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetCustomData

`func (o *PlayerConfigurationResponseResultsInner) GetCustomData() map[string]interface{}`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *PlayerConfigurationResponseResultsInner) GetCustomDataOk() (*map[string]interface{}, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *PlayerConfigurationResponseResultsInner) SetCustomData(v map[string]interface{})`

SetCustomData sets CustomData field to given value.


### SetCustomDataNil

`func (o *PlayerConfigurationResponseResultsInner) SetCustomDataNil(b bool)`

 SetCustomDataNil sets the value for CustomData to be an explicit nil

### UnsetCustomData
`func (o *PlayerConfigurationResponseResultsInner) UnsetCustomData()`

UnsetCustomData ensures that no value is present for CustomData, not even an explicit nil
### GetCosts

`func (o *PlayerConfigurationResponseResultsInner) GetCosts() []Cost`

GetCosts returns the Costs field if non-nil, zero value otherwise.

### GetCostsOk

`func (o *PlayerConfigurationResponseResultsInner) GetCostsOk() (*[]Cost, bool)`

GetCostsOk returns a tuple with the Costs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosts

`func (o *PlayerConfigurationResponseResultsInner) SetCosts(v []Cost)`

SetCosts sets Costs field to given value.

### HasCosts

`func (o *PlayerConfigurationResponseResultsInner) HasCosts() bool`

HasCosts returns a boolean if a field has been set.

### GetRewards

`func (o *PlayerConfigurationResponseResultsInner) GetRewards() []Reward`

GetRewards returns the Rewards field if non-nil, zero value otherwise.

### GetRewardsOk

`func (o *PlayerConfigurationResponseResultsInner) GetRewardsOk() (*[]Reward, bool)`

GetRewardsOk returns a tuple with the Rewards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewards

`func (o *PlayerConfigurationResponseResultsInner) SetRewards(v []Reward)`

SetRewards sets Rewards field to given value.


### GetStoreIdentifiers

`func (o *PlayerConfigurationResponseResultsInner) GetStoreIdentifiers() RealMoneyPurchaseResourceStoreIdentifiers`

GetStoreIdentifiers returns the StoreIdentifiers field if non-nil, zero value otherwise.

### GetStoreIdentifiersOk

`func (o *PlayerConfigurationResponseResultsInner) GetStoreIdentifiersOk() (*RealMoneyPurchaseResourceStoreIdentifiers, bool)`

GetStoreIdentifiersOk returns a tuple with the StoreIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreIdentifiers

`func (o *PlayerConfigurationResponseResultsInner) SetStoreIdentifiers(v RealMoneyPurchaseResourceStoreIdentifiers)`

SetStoreIdentifiers sets StoreIdentifiers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


