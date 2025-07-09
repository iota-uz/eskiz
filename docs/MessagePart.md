# MessagePart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accept** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Submit** | Pointer to **int** |  | [optional] 
**Delivery** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewMessagePart

`func NewMessagePart() *MessagePart`

NewMessagePart instantiates a new MessagePart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagePartWithDefaults

`func NewMessagePartWithDefaults() *MessagePart`

NewMessagePartWithDefaults instantiates a new MessagePart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccept

`func (o *MessagePart) GetAccept() time.Time`

GetAccept returns the Accept field if non-nil, zero value otherwise.

### GetAcceptOk

`func (o *MessagePart) GetAcceptOk() (*time.Time, bool)`

GetAcceptOk returns a tuple with the Accept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccept

`func (o *MessagePart) SetAccept(v time.Time)`

SetAccept sets Accept field to given value.

### HasAccept

`func (o *MessagePart) HasAccept() bool`

HasAccept returns a boolean if a field has been set.

### GetStatus

`func (o *MessagePart) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MessagePart) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MessagePart) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MessagePart) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubmit

`func (o *MessagePart) GetSubmit() int`

GetSubmit returns the Submit field if non-nil, zero value otherwise.

### GetSubmitOk

`func (o *MessagePart) GetSubmitOk() (*int, bool)`

GetSubmitOk returns a tuple with the Submit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmit

`func (o *MessagePart) SetSubmit(v int)`

SetSubmit sets Submit field to given value.

### HasSubmit

`func (o *MessagePart) HasSubmit() bool`

HasSubmit returns a boolean if a field has been set.

### GetDelivery

`func (o *MessagePart) GetDelivery() time.Time`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *MessagePart) GetDeliveryOk() (*time.Time, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *MessagePart) SetDelivery(v time.Time)`

SetDelivery sets Delivery field to given value.

### HasDelivery

`func (o *MessagePart) HasDelivery() bool`

HasDelivery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


