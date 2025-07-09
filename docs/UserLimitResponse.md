# UserLimitResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**UserBalance**](UserBalance.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewUserLimitResponse

`func NewUserLimitResponse() *UserLimitResponse`

NewUserLimitResponse instantiates a new UserLimitResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserLimitResponseWithDefaults

`func NewUserLimitResponseWithDefaults() *UserLimitResponse`

NewUserLimitResponseWithDefaults instantiates a new UserLimitResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *UserLimitResponse) GetData() UserBalance`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserLimitResponse) GetDataOk() (*UserBalance, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserLimitResponse) SetData(v UserBalance)`

SetData sets Data field to given value.

### HasData

`func (o *UserLimitResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *UserLimitResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserLimitResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserLimitResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserLimitResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


