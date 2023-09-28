# Reward

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | The ID of a currency or inventory item resource that is credited as part of the purchase. | 
**Amount** | **int32** | The amount of the resource credited as part of the purchase. | 
**DefaultInstanceData** | Pointer to **map[string]interface{}** | When the reward is an inventory item resource, this specifies instance data automatically saved against the new inventory item. Max size when serialised 5 KB. This property has been deprecated and will only return \&quot;null\&quot;. | [optional] 

## Methods

### NewReward

`func NewReward(resourceId string, amount int32, ) *Reward`

NewReward instantiates a new Reward object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardWithDefaults

`func NewRewardWithDefaults() *Reward`

NewRewardWithDefaults instantiates a new Reward object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *Reward) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *Reward) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *Reward) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetAmount

`func (o *Reward) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Reward) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Reward) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetDefaultInstanceData

`func (o *Reward) GetDefaultInstanceData() map[string]interface{}`

GetDefaultInstanceData returns the DefaultInstanceData field if non-nil, zero value otherwise.

### GetDefaultInstanceDataOk

`func (o *Reward) GetDefaultInstanceDataOk() (*map[string]interface{}, bool)`

GetDefaultInstanceDataOk returns a tuple with the DefaultInstanceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceData

`func (o *Reward) SetDefaultInstanceData(v map[string]interface{})`

SetDefaultInstanceData sets DefaultInstanceData field to given value.

### HasDefaultInstanceData

`func (o *Reward) HasDefaultInstanceData() bool`

HasDefaultInstanceData returns a boolean if a field has been set.

### SetDefaultInstanceDataNil

`func (o *Reward) SetDefaultInstanceDataNil(b bool)`

 SetDefaultInstanceDataNil sets the value for DefaultInstanceData to be an explicit nil

### UnsetDefaultInstanceData
`func (o *Reward) UnsetDefaultInstanceData()`

UnsetDefaultInstanceData ensures that no value is present for DefaultInstanceData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


