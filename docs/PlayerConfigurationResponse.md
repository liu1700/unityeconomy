# PlayerConfigurationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**PlayerConfigurationResponseMetadata**](PlayerConfigurationResponseMetadata.md) |  | [optional] 
**Results** | [**[]PlayerConfigurationResponseResultsInner**](PlayerConfigurationResponseResultsInner.md) | Array of resource definitions. | 

## Methods

### NewPlayerConfigurationResponse

`func NewPlayerConfigurationResponse(results []PlayerConfigurationResponseResultsInner, ) *PlayerConfigurationResponse`

NewPlayerConfigurationResponse instantiates a new PlayerConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerConfigurationResponseWithDefaults

`func NewPlayerConfigurationResponseWithDefaults() *PlayerConfigurationResponse`

NewPlayerConfigurationResponseWithDefaults instantiates a new PlayerConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *PlayerConfigurationResponse) GetMetadata() PlayerConfigurationResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PlayerConfigurationResponse) GetMetadataOk() (*PlayerConfigurationResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PlayerConfigurationResponse) SetMetadata(v PlayerConfigurationResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PlayerConfigurationResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetResults

`func (o *PlayerConfigurationResponse) GetResults() []PlayerConfigurationResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PlayerConfigurationResponse) GetResultsOk() (*[]PlayerConfigurationResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PlayerConfigurationResponse) SetResults(v []PlayerConfigurationResponseResultsInner)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


