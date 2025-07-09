# SendSmsBatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackUrl** | Pointer to **string** |  | [optional] 
**DispatchId** | Pointer to **float64** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**Messages** | Pointer to [**[]SendSmsBatchRequestMessagesInner**](SendSmsBatchRequestMessagesInner.md) |  | [optional] 

## Methods

### NewSendSmsBatchRequest

`func NewSendSmsBatchRequest() *SendSmsBatchRequest`

NewSendSmsBatchRequest instantiates a new SendSmsBatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendSmsBatchRequestWithDefaults

`func NewSendSmsBatchRequestWithDefaults() *SendSmsBatchRequest`

NewSendSmsBatchRequestWithDefaults instantiates a new SendSmsBatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackUrl

`func (o *SendSmsBatchRequest) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *SendSmsBatchRequest) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *SendSmsBatchRequest) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *SendSmsBatchRequest) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetDispatchId

`func (o *SendSmsBatchRequest) GetDispatchId() float64`

GetDispatchId returns the DispatchId field if non-nil, zero value otherwise.

### GetDispatchIdOk

`func (o *SendSmsBatchRequest) GetDispatchIdOk() (*float64, bool)`

GetDispatchIdOk returns a tuple with the DispatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispatchId

`func (o *SendSmsBatchRequest) SetDispatchId(v float64)`

SetDispatchId sets DispatchId field to given value.

### HasDispatchId

`func (o *SendSmsBatchRequest) HasDispatchId() bool`

HasDispatchId returns a boolean if a field has been set.

### GetFrom

`func (o *SendSmsBatchRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SendSmsBatchRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SendSmsBatchRequest) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *SendSmsBatchRequest) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetMessages

`func (o *SendSmsBatchRequest) GetMessages() []SendSmsBatchRequestMessagesInner`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *SendSmsBatchRequest) GetMessagesOk() (*[]SendSmsBatchRequestMessagesInner, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *SendSmsBatchRequest) SetMessages(v []SendSmsBatchRequestMessagesInner)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *SendSmsBatchRequest) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


