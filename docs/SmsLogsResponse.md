# SmsLogsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to [**[]LogMessage**](LogMessage.md) |  | [optional] 
**MessagesCount** | Pointer to **int** |  | [optional] 

## Methods

### NewSmsLogsResponse

`func NewSmsLogsResponse() *SmsLogsResponse`

NewSmsLogsResponse instantiates a new SmsLogsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsLogsResponseWithDefaults

`func NewSmsLogsResponseWithDefaults() *SmsLogsResponse`

NewSmsLogsResponseWithDefaults instantiates a new SmsLogsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *SmsLogsResponse) GetMessages() []LogMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *SmsLogsResponse) GetMessagesOk() (*[]LogMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *SmsLogsResponse) SetMessages(v []LogMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *SmsLogsResponse) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetMessagesCount

`func (o *SmsLogsResponse) GetMessagesCount() int`

GetMessagesCount returns the MessagesCount field if non-nil, zero value otherwise.

### GetMessagesCountOk

`func (o *SmsLogsResponse) GetMessagesCountOk() (*int, bool)`

GetMessagesCountOk returns a tuple with the MessagesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesCount

`func (o *SmsLogsResponse) SetMessagesCount(v int)`

SetMessagesCount sets MessagesCount field to given value.

### HasMessagesCount

`func (o *SmsLogsResponse) HasMessagesCount() bool`

HasMessagesCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


