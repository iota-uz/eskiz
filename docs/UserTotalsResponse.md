# UserTotalsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]UserTotalsItem**](UserTotalsItem.md) |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUserTotalsResponse

`func NewUserTotalsResponse() *UserTotalsResponse`

NewUserTotalsResponse instantiates a new UserTotalsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserTotalsResponseWithDefaults

`func NewUserTotalsResponseWithDefaults() *UserTotalsResponse`

NewUserTotalsResponseWithDefaults instantiates a new UserTotalsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UserTotalsResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserTotalsResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserTotalsResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserTotalsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *UserTotalsResponse) GetData() []UserTotalsItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserTotalsResponse) GetDataOk() (*[]UserTotalsItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserTotalsResponse) SetData(v []UserTotalsItem)`

SetData sets Data field to given value.

### HasData

`func (o *UserTotalsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetId

`func (o *UserTotalsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserTotalsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserTotalsResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserTotalsResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *UserTotalsResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UserTotalsResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


