# CurrencyResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier for the resource. | 
**Name** | **string** | Name of the resource. | 
**Type** | **string** | Type of the resource, for example &#x60;CURRENCY&#x60;, &#x60;INVENTORY_ITEM&#x60;, &#x60;VIRTUAL_PURCHASE&#x60;, &#x60;MONEY_PURCHASE&#x60;. | 
**Created** | [**ModifiedMetadata**](ModifiedMetadata.md) |  | 
**Modified** | [**ModifiedMetadata**](ModifiedMetadata.md) |  | 
**Initial** | Pointer to **int32** |  | [optional] 
**Max** | Pointer to **int32** |  | [optional] 
**CustomData** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCurrencyResource

`func NewCurrencyResource(id string, name string, type_ string, created ModifiedMetadata, modified ModifiedMetadata, ) *CurrencyResource`

NewCurrencyResource instantiates a new CurrencyResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyResourceWithDefaults

`func NewCurrencyResourceWithDefaults() *CurrencyResource`

NewCurrencyResourceWithDefaults instantiates a new CurrencyResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CurrencyResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CurrencyResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CurrencyResource) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CurrencyResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CurrencyResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CurrencyResource) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CurrencyResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CurrencyResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CurrencyResource) SetType(v string)`

SetType sets Type field to given value.


### GetCreated

`func (o *CurrencyResource) GetCreated() ModifiedMetadata`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CurrencyResource) GetCreatedOk() (*ModifiedMetadata, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CurrencyResource) SetCreated(v ModifiedMetadata)`

SetCreated sets Created field to given value.


### GetModified

`func (o *CurrencyResource) GetModified() ModifiedMetadata`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *CurrencyResource) GetModifiedOk() (*ModifiedMetadata, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *CurrencyResource) SetModified(v ModifiedMetadata)`

SetModified sets Modified field to given value.


### GetInitial

`func (o *CurrencyResource) GetInitial() int32`

GetInitial returns the Initial field if non-nil, zero value otherwise.

### GetInitialOk

`func (o *CurrencyResource) GetInitialOk() (*int32, bool)`

GetInitialOk returns a tuple with the Initial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitial

`func (o *CurrencyResource) SetInitial(v int32)`

SetInitial sets Initial field to given value.

### HasInitial

`func (o *CurrencyResource) HasInitial() bool`

HasInitial returns a boolean if a field has been set.

### GetMax

`func (o *CurrencyResource) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *CurrencyResource) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *CurrencyResource) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *CurrencyResource) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetCustomData

`func (o *CurrencyResource) GetCustomData() map[string]interface{}`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *CurrencyResource) GetCustomDataOk() (*map[string]interface{}, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *CurrencyResource) SetCustomData(v map[string]interface{})`

SetCustomData sets CustomData field to given value.

### HasCustomData

`func (o *CurrencyResource) HasCustomData() bool`

HasCustomData returns a boolean if a field has been set.

### SetCustomDataNil

`func (o *CurrencyResource) SetCustomDataNil(b bool)`

 SetCustomDataNil sets the value for CustomData to be an explicit nil

### UnsetCustomData
`func (o *CurrencyResource) UnsetCustomData()`

UnsetCustomData ensures that no value is present for CustomData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


