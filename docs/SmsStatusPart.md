# SmsStatusPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **int** |  | [optional] 
**Accepted** | Pointer to **bool** |  | [optional] 
**DlrTime** | Pointer to **time.Time** |  | [optional] 
**DlrState** | Pointer to **string** |  | [optional] 
**PartIndex** | Pointer to **int** |  | [optional] 
**AcceptTime** | Pointer to **time.Time** |  | [optional] 
**AcceptStatus** | Pointer to **int** |  | [optional] 

## Methods

### NewSmsStatusPart

`func NewSmsStatusPart() *SmsStatusPart`

NewSmsStatusPart instantiates a new SmsStatusPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsStatusPartWithDefaults

`func NewSmsStatusPartWithDefaults() *SmsStatusPart`

NewSmsStatusPartWithDefaults instantiates a new SmsStatusPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *SmsStatusPart) GetGroup() int`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SmsStatusPart) GetGroupOk() (*int, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SmsStatusPart) SetGroup(v int)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *SmsStatusPart) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetAccepted

`func (o *SmsStatusPart) GetAccepted() bool`

GetAccepted returns the Accepted field if non-nil, zero value otherwise.

### GetAcceptedOk

`func (o *SmsStatusPart) GetAcceptedOk() (*bool, bool)`

GetAcceptedOk returns a tuple with the Accepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepted

`func (o *SmsStatusPart) SetAccepted(v bool)`

SetAccepted sets Accepted field to given value.

### HasAccepted

`func (o *SmsStatusPart) HasAccepted() bool`

HasAccepted returns a boolean if a field has been set.

### GetDlrTime

`func (o *SmsStatusPart) GetDlrTime() time.Time`

GetDlrTime returns the DlrTime field if non-nil, zero value otherwise.

### GetDlrTimeOk

`func (o *SmsStatusPart) GetDlrTimeOk() (*time.Time, bool)`

GetDlrTimeOk returns a tuple with the DlrTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlrTime

`func (o *SmsStatusPart) SetDlrTime(v time.Time)`

SetDlrTime sets DlrTime field to given value.

### HasDlrTime

`func (o *SmsStatusPart) HasDlrTime() bool`

HasDlrTime returns a boolean if a field has been set.

### GetDlrState

`func (o *SmsStatusPart) GetDlrState() string`

GetDlrState returns the DlrState field if non-nil, zero value otherwise.

### GetDlrStateOk

`func (o *SmsStatusPart) GetDlrStateOk() (*string, bool)`

GetDlrStateOk returns a tuple with the DlrState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlrState

`func (o *SmsStatusPart) SetDlrState(v string)`

SetDlrState sets DlrState field to given value.

### HasDlrState

`func (o *SmsStatusPart) HasDlrState() bool`

HasDlrState returns a boolean if a field has been set.

### GetPartIndex

`func (o *SmsStatusPart) GetPartIndex() int`

GetPartIndex returns the PartIndex field if non-nil, zero value otherwise.

### GetPartIndexOk

`func (o *SmsStatusPart) GetPartIndexOk() (*int, bool)`

GetPartIndexOk returns a tuple with the PartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartIndex

`func (o *SmsStatusPart) SetPartIndex(v int)`

SetPartIndex sets PartIndex field to given value.

### HasPartIndex

`func (o *SmsStatusPart) HasPartIndex() bool`

HasPartIndex returns a boolean if a field has been set.

### GetAcceptTime

`func (o *SmsStatusPart) GetAcceptTime() time.Time`

GetAcceptTime returns the AcceptTime field if non-nil, zero value otherwise.

### GetAcceptTimeOk

`func (o *SmsStatusPart) GetAcceptTimeOk() (*time.Time, bool)`

GetAcceptTimeOk returns a tuple with the AcceptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptTime

`func (o *SmsStatusPart) SetAcceptTime(v time.Time)`

SetAcceptTime sets AcceptTime field to given value.

### HasAcceptTime

`func (o *SmsStatusPart) HasAcceptTime() bool`

HasAcceptTime returns a boolean if a field has been set.

### GetAcceptStatus

`func (o *SmsStatusPart) GetAcceptStatus() int`

GetAcceptStatus returns the AcceptStatus field if non-nil, zero value otherwise.

### GetAcceptStatusOk

`func (o *SmsStatusPart) GetAcceptStatusOk() (*int, bool)`

GetAcceptStatusOk returns a tuple with the AcceptStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptStatus

`func (o *SmsStatusPart) SetAcceptStatus(v int)`

SetAcceptStatus sets AcceptStatus field to given value.

### HasAcceptStatus

`func (o *SmsStatusPart) HasAcceptStatus() bool`

HasAcceptStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


