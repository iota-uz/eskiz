# TemplatesListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Result** | Pointer to [**[]TemplateItem**](TemplateItem.md) |  | [optional] 

## Methods

### NewTemplatesListResponse

`func NewTemplatesListResponse() *TemplatesListResponse`

NewTemplatesListResponse instantiates a new TemplatesListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatesListResponseWithDefaults

`func NewTemplatesListResponseWithDefaults() *TemplatesListResponse`

NewTemplatesListResponseWithDefaults instantiates a new TemplatesListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *TemplatesListResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *TemplatesListResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *TemplatesListResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *TemplatesListResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetResult

`func (o *TemplatesListResponse) GetResult() []TemplateItem`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TemplatesListResponse) GetResultOk() (*[]TemplateItem, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TemplatesListResponse) SetResult(v []TemplateItem)`

SetResult sets Result field to given value.

### HasResult

`func (o *TemplatesListResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


