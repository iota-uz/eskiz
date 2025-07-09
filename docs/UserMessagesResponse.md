# UserMessagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**PaginationData**](PaginationData.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewUserMessagesResponse

`func NewUserMessagesResponse() *UserMessagesResponse`

NewUserMessagesResponse instantiates a new UserMessagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserMessagesResponseWithDefaults

`func NewUserMessagesResponseWithDefaults() *UserMessagesResponse`

NewUserMessagesResponseWithDefaults instantiates a new UserMessagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *UserMessagesResponse) GetData() PaginationData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserMessagesResponse) GetDataOk() (*PaginationData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserMessagesResponse) SetData(v PaginationData)`

SetData sets Data field to given value.

### HasData

`func (o *UserMessagesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *UserMessagesResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserMessagesResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserMessagesResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserMessagesResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


