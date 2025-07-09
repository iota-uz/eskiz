# LogMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**To** | Pointer to **string** |  | [optional] 
**DispatchId** | Pointer to **string** |  | [optional] 
**UserSmsId** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**FullMessage** | Pointer to **string** |  | [optional] 
**LevelName** | Pointer to **string** |  | [optional] 

## Methods

### NewLogMessage

`func NewLogMessage() *LogMessage`

NewLogMessage instantiates a new LogMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogMessageWithDefaults

`func NewLogMessageWithDefaults() *LogMessage`

NewLogMessageWithDefaults instantiates a new LogMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *LogMessage) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *LogMessage) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *LogMessage) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *LogMessage) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetFrom

`func (o *LogMessage) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *LogMessage) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *LogMessage) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *LogMessage) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *LogMessage) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *LogMessage) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *LogMessage) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *LogMessage) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetDispatchId

`func (o *LogMessage) GetDispatchId() string`

GetDispatchId returns the DispatchId field if non-nil, zero value otherwise.

### GetDispatchIdOk

`func (o *LogMessage) GetDispatchIdOk() (*string, bool)`

GetDispatchIdOk returns a tuple with the DispatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispatchId

`func (o *LogMessage) SetDispatchId(v string)`

SetDispatchId sets DispatchId field to given value.

### HasDispatchId

`func (o *LogMessage) HasDispatchId() bool`

HasDispatchId returns a boolean if a field has been set.

### GetUserSmsId

`func (o *LogMessage) GetUserSmsId() string`

GetUserSmsId returns the UserSmsId field if non-nil, zero value otherwise.

### GetUserSmsIdOk

`func (o *LogMessage) GetUserSmsIdOk() (*string, bool)`

GetUserSmsIdOk returns a tuple with the UserSmsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSmsId

`func (o *LogMessage) SetUserSmsId(v string)`

SetUserSmsId sets UserSmsId field to given value.

### HasUserSmsId

`func (o *LogMessage) HasUserSmsId() bool`

HasUserSmsId returns a boolean if a field has been set.

### GetTimestamp

`func (o *LogMessage) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LogMessage) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LogMessage) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LogMessage) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetMessage

`func (o *LogMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LogMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LogMessage) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *LogMessage) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetFullMessage

`func (o *LogMessage) GetFullMessage() string`

GetFullMessage returns the FullMessage field if non-nil, zero value otherwise.

### GetFullMessageOk

`func (o *LogMessage) GetFullMessageOk() (*string, bool)`

GetFullMessageOk returns a tuple with the FullMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullMessage

`func (o *LogMessage) SetFullMessage(v string)`

SetFullMessage sets FullMessage field to given value.

### HasFullMessage

`func (o *LogMessage) HasFullMessage() bool`

HasFullMessage returns a boolean if a field has been set.

### GetLevelName

`func (o *LogMessage) GetLevelName() string`

GetLevelName returns the LevelName field if non-nil, zero value otherwise.

### GetLevelNameOk

`func (o *LogMessage) GetLevelNameOk() (*string, bool)`

GetLevelNameOk returns a tuple with the LevelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelName

`func (o *LogMessage) SetLevelName(v string)`

SetLevelName sets LevelName field to given value.

### HasLevelName

`func (o *LogMessage) HasLevelName() bool`

HasLevelName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


