# SmsCallbackWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** | UUID запроса, возвращаемый при отправке SMS | [optional] 
**MessageId** | Pointer to **string** | ID сообщения от SMS провайдера | [optional] 
**UserSmsId** | Pointer to **string** | Пользовательский идентификатор SMS | [optional] 
**Country** | Pointer to **string** | Код страны получателя | [optional] 
**PhoneNumber** | Pointer to **string** | Номер телефона получателя | [optional] 
**SmsCount** | Pointer to **string** | Количество SMS частей | [optional] 
**Status** | Pointer to **string** | Статус доставки SMS | [optional] 
**StatusDate** | Pointer to **string** | Дата и время получения статуса | [optional] 

## Methods

### NewSmsCallbackWebhook

`func NewSmsCallbackWebhook() *SmsCallbackWebhook`

NewSmsCallbackWebhook instantiates a new SmsCallbackWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsCallbackWebhookWithDefaults

`func NewSmsCallbackWebhookWithDefaults() *SmsCallbackWebhook`

NewSmsCallbackWebhookWithDefaults instantiates a new SmsCallbackWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *SmsCallbackWebhook) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SmsCallbackWebhook) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SmsCallbackWebhook) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *SmsCallbackWebhook) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetMessageId

`func (o *SmsCallbackWebhook) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *SmsCallbackWebhook) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *SmsCallbackWebhook) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *SmsCallbackWebhook) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetUserSmsId

`func (o *SmsCallbackWebhook) GetUserSmsId() string`

GetUserSmsId returns the UserSmsId field if non-nil, zero value otherwise.

### GetUserSmsIdOk

`func (o *SmsCallbackWebhook) GetUserSmsIdOk() (*string, bool)`

GetUserSmsIdOk returns a tuple with the UserSmsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSmsId

`func (o *SmsCallbackWebhook) SetUserSmsId(v string)`

SetUserSmsId sets UserSmsId field to given value.

### HasUserSmsId

`func (o *SmsCallbackWebhook) HasUserSmsId() bool`

HasUserSmsId returns a boolean if a field has been set.

### GetCountry

`func (o *SmsCallbackWebhook) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SmsCallbackWebhook) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SmsCallbackWebhook) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *SmsCallbackWebhook) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *SmsCallbackWebhook) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *SmsCallbackWebhook) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *SmsCallbackWebhook) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *SmsCallbackWebhook) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetSmsCount

`func (o *SmsCallbackWebhook) GetSmsCount() string`

GetSmsCount returns the SmsCount field if non-nil, zero value otherwise.

### GetSmsCountOk

`func (o *SmsCallbackWebhook) GetSmsCountOk() (*string, bool)`

GetSmsCountOk returns a tuple with the SmsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsCount

`func (o *SmsCallbackWebhook) SetSmsCount(v string)`

SetSmsCount sets SmsCount field to given value.

### HasSmsCount

`func (o *SmsCallbackWebhook) HasSmsCount() bool`

HasSmsCount returns a boolean if a field has been set.

### GetStatus

`func (o *SmsCallbackWebhook) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SmsCallbackWebhook) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SmsCallbackWebhook) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SmsCallbackWebhook) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDate

`func (o *SmsCallbackWebhook) GetStatusDate() string`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *SmsCallbackWebhook) GetStatusDateOk() (*string, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *SmsCallbackWebhook) SetStatusDate(v string)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *SmsCallbackWebhook) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


