# SmsStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**SmsStatusData**](SmsStatusData.md) |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSmsStatusResponse

`func NewSmsStatusResponse() *SmsStatusResponse`

NewSmsStatusResponse instantiates a new SmsStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsStatusResponseWithDefaults

`func NewSmsStatusResponseWithDefaults() *SmsStatusResponse`

NewSmsStatusResponseWithDefaults instantiates a new SmsStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SmsStatusResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SmsStatusResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SmsStatusResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SmsStatusResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *SmsStatusResponse) GetData() SmsStatusData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SmsStatusResponse) GetDataOk() (*SmsStatusData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SmsStatusResponse) SetData(v SmsStatusData)`

SetData sets Data field to given value.

### HasData

`func (o *SmsStatusResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetId

`func (o *SmsStatusResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SmsStatusResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SmsStatusResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SmsStatusResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SmsStatusResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SmsStatusResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


