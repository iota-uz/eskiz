# DispatchStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]StatusCount**](StatusCount.md) |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDispatchStatusResponse

`func NewDispatchStatusResponse() *DispatchStatusResponse`

NewDispatchStatusResponse instantiates a new DispatchStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDispatchStatusResponseWithDefaults

`func NewDispatchStatusResponseWithDefaults() *DispatchStatusResponse`

NewDispatchStatusResponseWithDefaults instantiates a new DispatchStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DispatchStatusResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DispatchStatusResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DispatchStatusResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DispatchStatusResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *DispatchStatusResponse) GetData() []StatusCount`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DispatchStatusResponse) GetDataOk() (*[]StatusCount, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DispatchStatusResponse) SetData(v []StatusCount)`

SetData sets Data field to given value.

### HasData

`func (o *DispatchStatusResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetId

`func (o *DispatchStatusResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DispatchStatusResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DispatchStatusResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DispatchStatusResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DispatchStatusResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DispatchStatusResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


