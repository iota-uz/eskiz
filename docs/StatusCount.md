# StatusCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Total** | Pointer to **int** |  | [optional] 

## Methods

### NewStatusCount

`func NewStatusCount() *StatusCount`

NewStatusCount instantiates a new StatusCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusCountWithDefaults

`func NewStatusCountWithDefaults() *StatusCount`

NewStatusCountWithDefaults instantiates a new StatusCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *StatusCount) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StatusCount) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StatusCount) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StatusCount) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotal

`func (o *StatusCount) GetTotal() int`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *StatusCount) GetTotalOk() (*int, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *StatusCount) SetTotal(v int)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *StatusCount) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


